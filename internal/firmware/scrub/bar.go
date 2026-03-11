package scrub

import (
	"github.com/sercanarga/pcileechgen/internal/firmware"
	"github.com/sercanarga/pcileechgen/internal/firmware/devclass"
	"github.com/sercanarga/pcileechgen/internal/util"
)

// ScrubBarContent patches BAR data with device-class quirks using default BRAM size.
func ScrubBarContent(barContents map[int][]byte, classCode uint32, vendorID uint16) {
	ScrubBarContentWithBRAM(barContents, classCode, vendorID, BRAMSize)
}

// ScrubBarContentWithBRAM is like ScrubBarContent but takes a custom BRAM size.
func ScrubBarContentWithBRAM(barContents map[int][]byte, classCode uint32, vendorID uint16, bramSize int) {
	data := firmware.LargestBar(barContents)
	if data == nil {
		return
	}

	strategy := devclass.StrategyForClassAndVendor(classCode, vendorID)
	if strategy == nil {
		return
	}

	strategy.ScrubBAR(data)

	// xHCI also needs BRAM-aware register clamping
	if strategy.DeviceClass() == devclass.ClassXHCI {
		scrubXHCIBar0(data, bramSize)
	}
}

// scrubXHCIBar0 clamps xHCI BAR0 registers to fit BRAM and fakes a running controller.
func scrubXHCIBar0(data []byte, bramSize int) {
	if len(data) < 0x20 {
		return
	}

	capLen := xhciFixCapLength(data)
	xhciFixHCIVersion(data)

	if capLen+0x40 > len(data) {
		return
	}

	maxSlots, maxIntrs, maxPorts := xhciReadStructParams(data)
	xhciClampScratchpads(data)
	xhciClampXECP(data, bramSize)

	maxSlots = xhciClampDBOFF(data, capLen, maxSlots, bramSize)
	maxIntrs = xhciClampRTSOFF(data, capLen, maxIntrs, bramSize)
	maxPorts = xhciClampPorts(capLen, maxPorts, bramSize)

	// write back clamped HCSPARAMS1
	hcsparams1 := uint32(maxSlots) | (uint32(maxIntrs) << 8) | (uint32(maxPorts) << 24)
	util.WriteLE32(data, 0x04, hcsparams1)

	xhciSetOperationalState(data, capLen, maxSlots)
}

func xhciFixCapLength(data []byte) int {
	capLen := int(data[0x00])
	if capLen == 0 || capLen > 0x40 {
		capLen = 0x20
	}
	data[0x00] = byte(capLen)
	return capLen
}

func xhciFixHCIVersion(data []byte) {
	hciVer := uint16(data[0x02]) | uint16(data[0x03])<<8
	if hciVer < 0x0100 {
		data[0x02] = 0x00
		data[0x03] = 0x01 // 0x0100 = xHCI 1.0
	}
}

func xhciReadStructParams(data []byte) (maxSlots, maxIntrs, maxPorts int) {
	hcsparams1 := util.ReadLE32(data, 0x04)
	maxSlots = int(hcsparams1 & 0xFF)
	if maxSlots == 0 {
		maxSlots = 32
	}
	maxIntrs = int((hcsparams1 >> 8) & 0x7FF)
	maxPorts = int((hcsparams1 >> 24) & 0xFF)
	return
}

// xhciClampScratchpads zeroes scratchpad counts in HCSPARAMS2, preserving SPR (bit 26).
func xhciClampScratchpads(data []byte) {
	hcsparams2 := util.ReadLE32(data, 0x08)
	hcsparams2 &= ^uint32(0xFBE00000)
	util.WriteLE32(data, 0x08, hcsparams2)
}

func xhciClampXECP(data []byte, bramSize int) {
	hccparams1 := util.ReadLE32(data, 0x10)
	xecp := int((hccparams1 >> 16) & 0xFFFF)
	if xecp*4 >= bramSize {
		hccparams1 &= 0x0000FFFF
		util.WriteLE32(data, 0x10, hccparams1)
	}
}

// xhciClampDBOFF clamps doorbell offset to fit BRAM, shrinking MaxSlots if needed.
func xhciClampDBOFF(data []byte, capLen, maxSlots, bramSize int) int {
	dboff := util.ReadLE32(data, 0x14) & ^uint32(0x03)
	doorbellSize := (maxSlots + 1) * 4

	if int(dboff)+doorbellSize > bramSize {
		newDBOFF := bramSize - doorbellSize
		if newDBOFF < 0 {
			newDBOFF = capLen + 0x20
		}
		newDBOFF = newDBOFF & ^0x1F
		if newDBOFF < capLen+0x20 {
			// can't fit, shrink MaxSlots
			available := bramSize - (capLen + 0x20)
			if available < 8 {
				available = 8
			}
			maxSlots = available/4 - 1
			if maxSlots < 1 {
				maxSlots = 1
			}
			doorbellSize = (maxSlots + 1) * 4
			newDBOFF = bramSize - doorbellSize
			if newDBOFF < 0 {
				newDBOFF = capLen + 0x20
			}
			newDBOFF = newDBOFF & ^0x1F
		}
		util.WriteLE32(data, 0x14, uint32(newDBOFF))
	}
	return maxSlots
}

func xhciClampRTSOFF(data []byte, capLen, maxIntrs, bramSize int) int {
	rtsoff := int(util.ReadLE32(data, 0x18) & ^uint32(0x1F))

	if rtsoff > 0 && maxIntrs > 0 {
		remaining := bramSize - rtsoff - 0x20
		if remaining < 0x20 {
			remaining = 0x20
		}
		maxFit := remaining / 0x20
		if maxFit < 1 {
			maxFit = 1
		}
		if maxIntrs > maxFit {
			maxIntrs = maxFit
		}
	}
	if maxIntrs < 1 {
		maxIntrs = 1
	}

	runtimeSize := 0x20 + maxIntrs*0x20
	if rtsoff+runtimeSize > bramSize {
		newRTSOFF := capLen + 0x20
		newRTSOFF = (newRTSOFF + 0x1F) & ^0x1F
		if newRTSOFF+runtimeSize > bramSize {
			newRTSOFF = bramSize - runtimeSize
			newRTSOFF = newRTSOFF & ^0x1F
		}
		rtsoff = newRTSOFF
		util.WriteLE32(data, 0x18, uint32(rtsoff))

		remaining := bramSize - rtsoff - 0x20
		if remaining < 0x20 {
			remaining = 0x20
		}
		maxFit := remaining / 0x20
		if maxFit < 1 {
			maxFit = 1
		}
		if maxIntrs > maxFit {
			maxIntrs = maxFit
		}
	}
	return maxIntrs
}

func xhciClampPorts(capLen, maxPorts, bramSize int) int {
	portBase := capLen + 0x400
	if portBase < bramSize {
		maxPortsFit := (bramSize - portBase) / 0x10
		if maxPorts > maxPortsFit {
			maxPorts = maxPortsFit
		}
	}
	if maxPorts < 1 {
		maxPorts = 1
	}
	return maxPorts
}

// xhciSetOperationalState makes the controller look like it's running.
func xhciSetOperationalState(data []byte, capLen, maxSlots int) {

	util.WriteLE32(data, capLen+0x08, 0x01)

	util.WriteLE32(data, capLen+0x14, 0x00)
	util.WriteLE32(data, capLen+0x18, 0x00)
	util.WriteLE32(data, capLen+0x1C, 0x00)

	config := util.ReadLE32(data, capLen+0x38)
	config = (config & 0xFFFFFF00) | uint32(maxSlots)
	util.WriteLE32(data, capLen+0x38, config)

	usbcmd := util.ReadLE32(data, capLen)
	usbcmd |= 0x01
	usbcmd &= ^uint32(0x02)
	util.WriteLE32(data, capLen, usbcmd)

	usbsts := util.ReadLE32(data, capLen+4)
	usbsts &= ^uint32(0x01 | 0x04)
	util.WriteLE32(data, capLen+4, usbsts)
}
