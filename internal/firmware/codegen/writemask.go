package codegen

import (
	"github.com/sercanarga/pcileechgen/internal/pci"
)

func applyCapabilityWritemasks(cs *pci.ConfigSpace, masks []uint32) {
	caps := pci.ParseCapabilities(cs)
	for _, cap := range caps {
		switch cap.ID {
		case pci.CapIDPowerManagement:
			if cap.Offset+4+4 <= pci.ConfigSpaceLegacySize {
				masks[(cap.Offset+4)/4] = 0x00008103 // PowerState + PME_En + PME_Status
			}
		case pci.CapIDMSI:
			applyMSIWritemask(cs, cap, masks)
		case pci.CapIDMSIX:
			if cap.Offset < pci.ConfigSpaceLegacySize {
				masks[(cap.Offset)/4] |= 0xC0000000 // Enable + Function Mask
			}
		case pci.CapIDPCIExpress:
			applyPCIeWritemask(cap, masks)
		}
	}
}

// applyMSIWritemask sets writemask bits for MSI capability fields.
func applyMSIWritemask(cs *pci.ConfigSpace, cap pci.Capability, masks []uint32) {
	if cap.Offset+4 <= pci.ConfigSpaceLegacySize {
		masks[(cap.Offset)/4] |= 0x00710000 // Enable + MultiMsg Enable
	}

	msgCtl := cs.ReadU16(cap.Offset + 2)
	is64Bit := (msgCtl & 0x0080) != 0
	hasMask := (msgCtl & 0x0100) != 0

	if cap.Offset+0x04+4 <= pci.ConfigSpaceLegacySize {
		masks[(cap.Offset+0x04)/4] = 0xFFFFFFFC // addr low
	}

	if is64Bit {
		if cap.Offset+0x08+4 <= pci.ConfigSpaceLegacySize {
			masks[(cap.Offset+0x08)/4] = 0xFFFFFFFF // addr high
		}
		if cap.Offset+0x0C+4 <= pci.ConfigSpaceLegacySize {
			masks[(cap.Offset+0x0C)/4] = 0x0000FFFF // data
		}
		if hasMask && cap.Offset+0x10+4 <= pci.ConfigSpaceLegacySize {
			masks[(cap.Offset+0x10)/4] = 0xFFFFFFFF // mask bits
		}
	} else {
		if cap.Offset+0x08+4 <= pci.ConfigSpaceLegacySize {
			masks[(cap.Offset+0x08)/4] = 0x0000FFFF // data
		}
		if hasMask && cap.Offset+0x0C+4 <= pci.ConfigSpaceLegacySize {
			masks[(cap.Offset+0x0C)/4] = 0xFFFFFFFF // mask bits
		}
	}
}

// applyPCIeWritemask sets writemask bits for PCIe capability fields.
func applyPCIeWritemask(cap pci.Capability, masks []uint32) {
	if cap.Offset+8+4 <= pci.ConfigSpaceLegacySize {
		// DevCtl (bits 15:0) + DevStatus RW1C (bits 19:16)
		masks[(cap.Offset+8)/4] = 0x000FFFFF
	}
	if cap.Offset+16+4 <= pci.ConfigSpaceLegacySize {
		masks[(cap.Offset+16)/4] = 0x0000FFFF // LinkCtl
	}
}

func applyExtCapabilityWritemasks(cs *pci.ConfigSpace, masks []uint32) {
	if cs.Size < pci.ConfigSpaceSize {
		return
	}

	extCaps := pci.ParseExtCapabilities(cs)
	for _, cap := range extCaps {
		wordIdx := cap.Offset / 4
		if wordIdx >= len(masks) {
			continue
		}

		switch cap.ID {
		case pci.ExtCapIDAER:
			// UE status, mask, severity + CE status, mask
			for i := 1; i <= 5 && wordIdx+i < len(masks); i++ {
				masks[wordIdx+i] = 0xFFFFFFFF
			}
		case pci.ExtCapIDLTR:
			if wordIdx+1 < len(masks) {
				masks[wordIdx+1] = 0xFFFFFFFF
			}
		}
	}
}
