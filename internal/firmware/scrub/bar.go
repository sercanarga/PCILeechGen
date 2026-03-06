package scrub

import (
	"github.com/sercanarga/pcileechgen/internal/firmware"
	"github.com/sercanarga/pcileechgen/internal/firmware/devclass"
	"github.com/sercanarga/pcileechgen/internal/util"
)

// ScrubBarContent patches BAR data for device-class-specific quirks.
// Call before GenerateBarContentCOE.
func ScrubBarContent(barContents map[int][]byte, classCode uint32) {
	data := firmware.LowestBar(barContents)
	if data == nil {
		return
	}

	strategy := devclass.StrategyForClass(classCode)
	if strategy != nil {
		strategy.ScrubBAR(data)
	}

	// xHCI needs BRAM-size-aware clamping that the strategy can't do alone
	if strategy != nil && strategy.DeviceClass() == devclass.ClassXHCI {
		scrubXHCIBar0(data)
	}
}

// scrubXHCIBar0 clamps xHCI BAR0 registers to 4KB BRAM
// and sets R/S=1, HCH=0 so the driver sees a running controller.
func scrubXHCIBar0(data []byte) {
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
	xhciClampXECP(data)

	maxSlots = xhciClampDBOFF(data, capLen, maxSlots)
	maxIntrs = xhciClampRTSOFF(data, capLen, maxIntrs)
	maxPorts = xhciClampPorts(data, capLen, maxPorts)

	// Write back clamped HCSPARAMS1
	hcsparams1 := uint32(maxSlots) | (uint32(maxIntrs) << 8) | (uint32(maxPorts) << 24)
	util.WriteLE32(data, 0x04, hcsparams1)

	xhciSetOperationalState(data, capLen, maxSlots)
}

// xhciFixCapLength ensures a valid CAPLENGTH field.
func xhciFixCapLength(data []byte) int {
	capLen := int(data[0x00])
	if capLen == 0 || capLen > 0x40 {
		capLen = 0x20
	}
	data[0x00] = byte(capLen)
	return capLen
}

// xhciFixHCIVersion ensures HCIVERSION >= 0x0100 (xHCI 1.0).
func xhciFixHCIVersion(data []byte) {
	hciVer := uint16(data[0x02]) | uint16(data[0x03])<<8
	if hciVer < 0x0100 {
		data[0x02] = 0x00
		data[0x03] = 0x01 // 0x0100 = xHCI 1.0
	}
}

// xhciReadStructParams extracts MaxSlots, MaxIntrs, MaxPorts from HCSPARAMS1.
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

// xhciClampScratchpads zeroes scratchpad counts in HCSPARAMS2 — BRAM can't handle them.
func xhciClampScratchpads(data []byte) {
	hcsparams2 := util.ReadLE32(data, 0x08)
	hcsparams2 &= ^uint32(0xFFE00000)
	util.WriteLE32(data, 0x08, hcsparams2)
}

// xhciClampXECP zeroes the xECP pointer in HCCPARAMS1 if it points outside BRAM.
func xhciClampXECP(data []byte) {
	hccparams1 := util.ReadLE32(data, 0x10)
	xecp := int((hccparams1 >> 16) & 0xFFFF)
	if xecp*4 >= BRAMSize {
		hccparams1 &= 0x0000FFFF
		util.WriteLE32(data, 0x10, hccparams1)
	}
}

// xhciClampDBOFF clamps doorbell offset so the array fits in BRAM.
// Returns updated maxSlots if shrinking was needed.
func xhciClampDBOFF(data []byte, capLen, maxSlots int) int {
	dboff := util.ReadLE32(data, 0x14) & ^uint32(0x03)
	doorbellSize := (maxSlots + 1) * 4

	if int(dboff)+doorbellSize > BRAMSize {
		newDBOFF := BRAMSize - doorbellSize
		if newDBOFF < 0 {
			newDBOFF = capLen + 0x20
		}
		newDBOFF = newDBOFF & ^0x1F // align down 32B
		if newDBOFF < capLen+0x20 {
			// not enough room, shrink MaxSlots
			available := BRAMSize - (capLen + 0x20)
			if available < 8 {
				available = 8 // minimum 2 doorbell slots
			}
			maxSlots = available/4 - 1
			if maxSlots < 1 {
				maxSlots = 1
			}
			doorbellSize = (maxSlots + 1) * 4
			newDBOFF = BRAMSize - doorbellSize
			if newDBOFF < 0 {
				newDBOFF = capLen + 0x20
			}
			newDBOFF = newDBOFF & ^0x1F
		}
		util.WriteLE32(data, 0x14, uint32(newDBOFF))
	}
	return maxSlots
}

// xhciClampRTSOFF clamps runtime register offset and MaxIntrs to fit BRAM.
func xhciClampRTSOFF(data []byte, capLen, maxIntrs int) int {
	rtsoff := int(util.ReadLE32(data, 0x18) & ^uint32(0x1F))

	// Clamp MaxIntrs to fit
	if rtsoff > 0 && maxIntrs > 0 {
		remaining := BRAMSize - rtsoff - 0x20
		if remaining < 0x20 {
			remaining = 0x20 // at least 1 interrupter
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
	if rtsoff+runtimeSize > BRAMSize {
		newRTSOFF := capLen + 0x20
		newRTSOFF = (newRTSOFF + 0x1F) & ^0x1F // align up 32B
		if newRTSOFF+runtimeSize > BRAMSize {
			newRTSOFF = BRAMSize - runtimeSize
			newRTSOFF = newRTSOFF & ^0x1F
		}
		rtsoff = newRTSOFF
		util.WriteLE32(data, 0x18, uint32(rtsoff))

		// re-check after moving
		remaining := BRAMSize - rtsoff - 0x20
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

// xhciClampPorts clamps MaxPorts so port registers fit within BRAM.
func xhciClampPorts(data []byte, capLen, maxPorts int) int {
	portBase := capLen + 0x400
	if portBase < BRAMSize {
		maxPortsFit := (BRAMSize - portBase) / 0x10
		if maxPorts > maxPortsFit {
			maxPorts = maxPortsFit
		}
	}
	if maxPorts < 1 {
		maxPorts = 1
	}
	return maxPorts
}

// xhciSetOperationalState sets page size, clears stale state, and ensures
// the controller appears running (R/S=1, HCH=0).
func xhciSetOperationalState(data []byte, capLen, maxSlots int) {
	// PAGESIZE: 4KB
	util.WriteLE32(data, capLen+0x08, 0x01)

	// DNCTRL + CRCR: clear, irrelevant for static BRAM
	util.WriteLE32(data, capLen+0x14, 0x00)
	util.WriteLE32(data, capLen+0x18, 0x00)
	util.WriteLE32(data, capLen+0x1C, 0x00)

	// CONFIG: MaxSlotsEn = clamped MaxSlots
	config := util.ReadLE32(data, capLen+0x38)
	config = (config & 0xFFFFFF00) | uint32(maxSlots)
	util.WriteLE32(data, capLen+0x38, config)

	// USBCMD: R/S=1, HCRST=0
	usbcmd := util.ReadLE32(data, capLen)
	usbcmd |= 0x01
	usbcmd &= ^uint32(0x02)
	util.WriteLE32(data, capLen, usbcmd)

	// USBSTS: HCH=0, HSE=0
	usbsts := util.ReadLE32(data, capLen+4)
	usbsts &= ^uint32(0x01 | 0x04)
	util.WriteLE32(data, capLen+4, usbsts)
}
