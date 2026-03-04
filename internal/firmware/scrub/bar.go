package scrub

import (
	"github.com/sercanarga/pcileechgen/internal/firmware"
	"github.com/sercanarga/pcileechgen/internal/util"
)

// ScrubBarContent patches BAR data for device-class-specific quirks.
// Call before GenerateBarContentCOE.
func ScrubBarContent(barContents map[int][]byte, classCode uint32) {
	data := firmware.LowestBarData(barContents)
	if data == nil {
		return
	}
	// Match on base class + subclass only (ignore progIF variations)
	switch classCode >> 8 {
	case 0x0108: // NVMe (any progIF)
		scrubNVMeBar0(data)
	case 0x0C03: // xHCI USB (any progIF — 0x30 standard, others possible)
		scrubXHCIBar0(data)
	}
}

// scrubNVMeBar0 patches NVMe controller registers in the BAR0 snapshot
// so stornvme.sys sees a coherent, ready-to-use controller state.
// Static BRAM can't do real handshakes, but a clean initial state
// avoids early-fail paths in the driver.
func scrubNVMeBar0(data []byte) {
	if len(data) < 0x38 {
		return
	}

	// CSTS @ 0x1C: RDY=1, clear CFS(1), SHST(3:2), NSSRO(4)
	// donor snapshot often has SHST=2 (shutdown complete) which
	// contradicts CC.EN=1 and confuses the driver immediately.
	csts := util.ReadLE32(data, 0x1C)
	csts |= 0x01          // RDY = 1
	csts &= ^uint32(0x1E) // clear bits 4:1 (NSSRO, SHST, CFS)
	util.WriteLE32(data, 0x1C, csts)

	// CC @ 0x14: EN=1 (coherent with RDY)
	cc := util.ReadLE32(data, 0x14)
	cc |= 0x01
	util.WriteLE32(data, 0x14, cc)

	// INTMS/INTMC @ 0x0C, 0x10: clear interrupt masks
	util.WriteLE32(data, 0x0C, 0x00)
	util.WriteLE32(data, 0x10, 0x00)

	// NSSR @ 0x20: clear subsystem reset
	util.WriteLE32(data, 0x20, 0x00)

	// AQA/ASQ/ACQ @ 0x24-0x37: zero out donor queue config.
	// These contain physical addresses from the donor host which
	// are meaningless on the target — driver sets them up itself.
	util.WriteLE32(data, 0x24, 0x00) // AQA
	util.WriteLE32(data, 0x28, 0x00) // ASQ low
	util.WriteLE32(data, 0x2C, 0x00) // ASQ high
	util.WriteLE32(data, 0x30, 0x00) // ACQ low
	util.WriteLE32(data, 0x34, 0x00) // ACQ high
}

// scrubXHCIBar0 clamps xHCI BAR0 registers to 4KB BRAM
// and sets R/S=1, HCH=0 so the driver sees a running controller.
func scrubXHCIBar0(data []byte) {
	if len(data) < 0x20 {
		return
	}

	capLen := int(data[0x00])
	if capLen == 0 || capLen > 0x40 {
		capLen = 0x20
	}
	// always write CAPLENGTH so driver finds operational regs
	data[0x00] = byte(capLen)

	// HCIVERSION (0x02): must be >= 0x0100 for xHCI 1.0
	hciVer := uint16(data[0x02]) | uint16(data[0x03])<<8
	if hciVer < 0x0100 {
		data[0x02] = 0x00
		data[0x03] = 0x01 // 0x0100 = xHCI 1.0
	}

	if capLen+0x40 > len(data) {
		return
	}

	// HCSPARAMS1 (0x04)
	hcsparams1 := util.ReadLE32(data, 0x04)
	maxSlots := int(hcsparams1 & 0xFF)
	if maxSlots == 0 {
		maxSlots = 32
	}
	maxIntrs := int((hcsparams1 >> 8) & 0x7FF)
	maxPorts := int((hcsparams1 >> 24) & 0xFF)

	// HCSPARAMS2 (0x08): nuke scratchpad counts, BRAM can't handle them
	hcsparams2 := util.ReadLE32(data, 0x08)
	hcsparams2 &= ^uint32(0xFFE00000)
	util.WriteLE32(data, 0x08, hcsparams2)

	// HCCPARAMS1 (0x10): kill xECP if it points outside BRAM
	hccparams1 := util.ReadLE32(data, 0x10)
	xecp := int((hccparams1 >> 16) & 0xFFFF)
	if xecp*4 >= BRAMSize {
		hccparams1 &= 0x0000FFFF
		util.WriteLE32(data, 0x10, hccparams1)
	}

	// DBOFF (0x14): doorbell array, (MaxSlots+1)*4 bytes
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

	// RTSOFF (0x18): runtime regs, each interrupter takes 0x20 bytes
	rtsoff := int(util.ReadLE32(data, 0x18) & ^uint32(0x1F))

	// clamp MaxIntrs to fit
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

	// MaxPorts: port regs start at capLen+0x400, 0x10 each
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

	// write back clamped HCSPARAMS1
	hcsparams1 = uint32(maxSlots) | (uint32(maxIntrs) << 8) | (uint32(maxPorts) << 24)
	util.WriteLE32(data, 0x04, hcsparams1)

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
