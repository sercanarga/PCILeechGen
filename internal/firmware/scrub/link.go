package scrub

import (
	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/firmware"
	"github.com/sercanarga/pcileechgen/internal/firmware/overlay"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

// clampLinkCapability caps link speed/width to board limits.
func clampLinkCapability(cs *pci.ConfigSpace, b *board.Board, om *overlay.Map, caps []pci.Capability) {
	for _, cap := range caps {
		if cap.ID != pci.CapIDPCIExpress {
			continue
		}

		var maxSpeed uint8
		if b != nil {
			maxSpeed = b.MaxLinkSpeedOrDefault()
		} else {
			maxSpeed = firmware.LinkSpeedGen2
		}
		maxWidth := uint8(0)
		if b != nil && b.PCIeLanes > 0 {
			maxWidth = uint8(b.PCIeLanes)
		}

		// Link Capabilities (cap+0x0C)
		if cap.Offset+0x0C+4 <= pci.ConfigSpaceLegacySize {
			linkCap := cs.ReadU32(cap.Offset + 0x0C)
			newLinkCap := linkCap
			if speed := uint8(newLinkCap & 0x0F); speed > maxSpeed {
				newLinkCap = (newLinkCap & 0xFFFFFFF0) | uint32(maxSpeed)
			}
			if maxWidth > 0 {
				if width := uint8((newLinkCap >> 4) & 0x3F); width > maxWidth {
					newLinkCap = (newLinkCap & 0xFFFFFC0F) | (uint32(maxWidth) << 4)
				}
			}
			om.WriteU32(cap.Offset+0x0C, newLinkCap, "clamp Link Capabilities")
		}

		// Link Status (cap+0x12)
		if cap.Offset+0x12+2 <= pci.ConfigSpaceLegacySize {
			ls := cs.ReadU16(cap.Offset + 0x12)
			newLS := ls
			if speed := uint8(newLS & 0x0F); speed > maxSpeed {
				newLS = (newLS & 0xFFF0) | uint16(maxSpeed)
			}
			if maxWidth > 0 {
				if width := uint8((newLS >> 4) & 0x3F); width > maxWidth {
					newLS = (newLS & 0xFC0F) | (uint16(maxWidth) << 4)
				}
			}
			om.WriteU16(cap.Offset+0x12, newLS, "clamp Link Status")
		}

		// LinkCtl2 (cap+0x30) - target speed
		if cap.Offset+0x30+2 <= pci.ConfigSpaceLegacySize {
			lc2 := cs.ReadU16(cap.Offset + 0x30)
			newLC2 := lc2
			speed := uint8(newLC2 & 0x0F)
			if speed == 0 || speed > maxSpeed {
				newLC2 = (newLC2 & 0xFFF0) | uint16(maxSpeed)
			}
			om.WriteU16(cap.Offset+0x30, newLC2, "clamp Link Control 2 target speed")
		}

		// LinkCap2 (cap+0x2C) - strip unsupported speeds from the donor vector
		if cap.Offset+0x2C+4 <= pci.ConfigSpaceLegacySize {
			lc2 := cs.ReadU32(cap.Offset + 0x2C)
			if lc2 != 0 {
				donorVec := lc2 & 0xFE // bits [7:1]
				var mask uint32
				for s := uint8(1); s <= maxSpeed; s++ {
					mask |= 1 << s
				}
				clampedVec := donorVec & mask
				if clampedVec == 0 {
					clampedVec = 0x02 // at least Gen1
				}
				om.WriteU32(cap.Offset+0x2C, (lc2&0xFFFFFF01)|clampedVec, "clamp Link Capabilities 2 speed vector")
			}
		}

		break
	}
}

// clampDeviceCapability forces MPS=128B, disables phantoms, ext tags, and
// all power-management-triggering features the FPGA can't emulate.
func clampDeviceCapability(cs *pci.ConfigSpace, om *overlay.Map, caps []pci.Capability) {
	for _, cap := range caps {
		if cap.ID != pci.CapIDPCIExpress {
			continue
		}

		// DevCap (cap+0x04)
		if cap.Offset+0x04+4 <= pci.ConfigSpaceLegacySize {
			devCap := cs.ReadU32(cap.Offset + 0x04)
			newDevCap := devCap
			newDevCap &= ^uint32(0x07) // MPS -> 128B
			newDevCap &= ^uint32(0x18) // phantom functions off
			newDevCap &= ^uint32(0x20) // extended tag off
			newDevCap &= ^uint32(1 << 28) // FLR capable off - prevents Windows from issuing function level reset
			om.WriteU32(cap.Offset+0x04, newDevCap, "clamp Device Capabilities (MPS/phantom/exttag/FLR)")
		}

		// DevCtl (cap+0x08)
		if cap.Offset+0x08+2 <= pci.ConfigSpaceLegacySize {
			devCtl := cs.ReadU16(cap.Offset + 0x08)
			newDevCtl := devCtl
			newDevCtl &= ^uint16(0x00E0) // MPS -> 128B
			newDevCtl &= ^uint16(0x0100) // ext tag off
			newDevCtl &= ^uint16(0x0200) // phantom off
			newDevCtl |= 0x0010          // Relaxed Ordering Enable
			if mrrs := (newDevCtl >> 12) & 0x07; mrrs > 2 {
				newDevCtl = (newDevCtl & 0x8FFF) | (2 << 12)
			}
			newDevCtl &= ^uint16(1 << 15) // initiate FLR off
			om.WriteU16(cap.Offset+0x08, newDevCtl, "clamp Device Control (MPS/MRRS/phantom/exttag)")
		}

		// DevCap2 (cap+0x24)
		if cap.Offset+0x24+4 <= pci.ConfigSpaceLegacySize {
			devCap2 := cs.ReadU32(cap.Offset + 0x24)
			newDevCap2 := devCap2
			newDevCap2 &= ^uint32(1 << 11)  // LTR Mechanism Supported off
			newDevCap2 &= ^uint32(0x03 << 18) // OBFF Supported off (bits 19:18)
			newDevCap2 &= ^uint32(1 << 16)  // 10-bit tag completer off
			newDevCap2 &= ^uint32(1 << 17)  // 10-bit tag requester off
			newDevCap2 &= ^uint32(1 << 28)  // FRS Supported off
			om.WriteU32(cap.Offset+0x24, newDevCap2, "clamp Device Capabilities 2 (LTR/OBFF/tags/FRS)")
		}

		// DevCtl2 (cap+0x28) - clear OBFF Enable and Completion Timeout Value
		if cap.Offset+0x28+2 <= pci.ConfigSpaceLegacySize {
			devCtl2 := cs.ReadU16(cap.Offset + 0x28)
			newDevCtl2 := devCtl2
			newDevCtl2 &= ^uint16(0x03 << 13) // OBFF Enable off (bits 14:13)
			newDevCtl2 &= ^uint16(0x0F)        // Completion Timeout Value = default
			om.WriteU16(cap.Offset+0x28, newDevCtl2, "clear OBFF Enable + Completion Timeout Value")
		}

		break
	}
}

