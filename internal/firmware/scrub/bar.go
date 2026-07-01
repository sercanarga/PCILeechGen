package scrub

import (
	"math/bits"

	"github.com/sercanarga/pcileechgen/internal/firmware"
	"github.com/sercanarga/pcileechgen/internal/firmware/devclass"
	"github.com/sercanarga/pcileechgen/internal/util"
)

func ScrubBarContent(barContents map[int][]byte, classCode uint32, vendorID uint16, bramSize int) {
	ScrubBarContentWithBRAM(barContents, classCode, vendorID, bramSize)
}

func ScrubBarContentWithBRAM(barContents map[int][]byte, classCode uint32, vendorID uint16, bramSize int) {
	data := firmware.LargestBar(barContents)
	if data == nil {
		return
	}
	if bramSize > 0 && len(data) > bramSize {
		data = data[:bramSize]
	}

	strategy := devclass.StrategyForClassAndVendor(classCode, vendorID)
	if strategy == nil {
		return
	}

	strategy.ScrubBAR(data)

	if strategy.DeviceClass() == devclass.ClassXHCI {
		scrubXHCIBar0(data, bramSize)
	}
	if strategy.DeviceClass() == devclass.ClassSATA {
		scrubSATABar0(data, bramSize)
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

	maxSlots = xhciClampDBOFF(data, maxSlots, bramSize)
	maxIntrs = xhciClampRTSOFF(data, maxIntrs, bramSize)
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

// clampCountToFit shrinks claimed to the largest count whose region
// [base, base+claimed*itemSize) fits within bramSize, floored at 1.
func clampCountToFit(base, itemSize, claimed, bramSize int) int {
	if base < bramSize {
		maxFit := (bramSize - base) / itemSize
		if claimed > maxFit {
			claimed = maxFit
		}
	}
	if claimed < 1 {
		claimed = 1
	}
	return claimed
}

// xhciClampDBOFF shrinks MaxSlots so the doorbell array fits within BRAM.
// DBOFF itself is never relocated: bar_controller.sv.tmpl and barmodel's
// buildXHCIBARModel hardcode the doorbell array at a fixed offset (see
// devclass/xhci.go's ScrubBAR), so moving it here would desync the SV
// wiring from what the driver was told.
func xhciClampDBOFF(data []byte, maxSlots, bramSize int) int {
	dboff := int(util.ReadLE32(data, 0x14) & ^uint32(0x03))
	maxSlots = clampCountToFit(dboff, 4, maxSlots+1, bramSize) - 1
	if maxSlots < 1 {
		maxSlots = 1
	}
	return maxSlots
}

// xhciClampRTSOFF shrinks MaxIntrs so the runtime interrupter register sets
// (starting at RTSOFF+0x20) fit within BRAM. RTSOFF itself is never
// relocated, for the same reason DBOFF isn't -- see xhciClampDBOFF.
func xhciClampRTSOFF(data []byte, maxIntrs, bramSize int) int {
	rtsoff := int(util.ReadLE32(data, 0x18) & ^uint32(0x1F))
	return clampCountToFit(rtsoff+0x20, 0x20, maxIntrs, bramSize)
}

func xhciClampPorts(capLen, maxPorts, bramSize int) int {
	return clampCountToFit(capLen+0x400, 0x10, maxPorts, bramSize)
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

// scrubSATABar0 clamps AHCI PI (Ports Implemented) to fit BRAM and fakes
// "no device present" for any port that fits but has no real donor data.
func scrubSATABar0(data []byte, bramSize int) {
	if len(data) < 0x10 {
		return
	}

	pi := util.ReadLE32(data, 0x0C)
	claimedPorts := bits.OnesCount32(pi)
	if claimedPorts == 0 {
		return
	}

	maxPorts := sataClampPorts(claimedPorts, bramSize)
	if maxPorts < claimedPorts {
		// shrink to only the low N ports (0..maxPorts-1) that fit
		pi &= (uint32(1) << uint(maxPorts)) - 1
		util.WriteLE32(data, 0x0C, pi)
	}

	// Port 0 always has real profile/donor data. Any other implemented
	// port that fits in BRAM but was never actually captured (its PxSSTS
	// block is still zero) needs an explicit "no device" state, otherwise
	// the host driver sees a phantom drive with garbage register content.
	for port := 1; port < 32; port++ {
		if pi&(1<<uint(port)) == 0 {
			continue
		}
		base := 0x100 + port*0x80
		if base+0x80 > len(data) {
			break
		}
		if util.ReadLE32(data, base+0x28) != 0 {
			continue // real donor data already shows a device here
		}
		sataMarkPortEmpty(data, base)
	}
}

// sataClampPorts clamps the AHCI port count to what fits within BRAM. Each
// port block is 0x80 bytes starting at 0x100.
// ponytail: assumes PI is the common contiguous 0..N-1 port bitmask; a
// sparse PI (spec-legal but rare) would need per-bit offset checks instead
// of a popcount-derived port count.
func sataClampPorts(claimedPorts, bramSize int) int {
	return clampCountToFit(0x100, 0x80, claimedPorts, bramSize)
}

// sataMarkPortEmpty writes the canonical "no device" register state for an
// AHCI port block that has no real donor data backing it.
func sataMarkPortEmpty(data []byte, base int) {
	util.WriteLE32(data, base+0x28, 0x00000000) // PxSSTS - no device detected
	util.WriteLE32(data, base+0x24, 0xFFFFFFFF) // PxSIG - no device signature
	tfd := util.ReadLE32(data, base+0x20)
	tfd &^= 0x81 // PxTFD - clear BSY/ERR so driver doesn't hang on phantom port
	util.WriteLE32(data, base+0x20, tfd)
}
