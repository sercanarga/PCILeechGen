// Package firmware handles COE generation and config space operations.
package firmware

import (
	"fmt"
	"strings"

	"github.com/sercanarga/pcileechgen/internal/pci"
)

// shadowCfgSpaceWords is the BRAM size used by pcileech shadow config space (4KB = 1024 DWORDs).
const shadowCfgSpaceWords = 1024

// formatCOE writes a COE file from a slice of uint32 words.
func formatCOE(header string, words []uint32) string {
	var sb strings.Builder
	sb.WriteString(header)
	sb.WriteString("memory_initialization_radix=16;\n")
	sb.WriteString("memory_initialization_vector=\n")

	for i, w := range words {
		if i < len(words)-1 {
			sb.WriteString(fmt.Sprintf("%08x,\n", w))
		} else {
			sb.WriteString(fmt.Sprintf("%08x;\n", w))
		}
	}
	return sb.String()
}

// GenerateConfigSpaceCOE generates the pcileech_cfgspace.coe file content.
// Always outputs 1024 DWORDs (4KB) to match the shadow config space BRAM size.
// If the donor config space is smaller (e.g. 256 bytes), the remaining words are zero-filled.
func GenerateConfigSpaceCOE(cs *pci.ConfigSpace) string {
	words := make([]uint32, shadowCfgSpaceWords)

	donorWords := cs.Size / 4
	for i := 0; i < donorWords && i < shadowCfgSpaceWords; i++ {
		words[i] = cs.ReadU32(i * 4)
	}

	return formatCOE(
		"; PCILeechGen - PCI Configuration Space (4KB shadow)\n"+
			"; Generated from donor device config space\n"+
			";\n",
		words,
	)
}

// GenerateWritemaskCOE generates the pcileech_cfgspace_writemask.coe file.
// Always outputs 1024 DWORDs (4KB) to match the shadow config space DROM size.
// Defines which bits are writable per PCI spec.
func GenerateWritemaskCOE(cs *pci.ConfigSpace) string {
	masks := make([]uint32, shadowCfgSpaceWords)

	// PCI Header writable fields (Type 0 header)
	masks[0x04/4] = 0x0000FFFF // Command register (lower 16 bits writable)
	masks[0x0C/4] = 0x0000FF00 // Latency timer
	masks[0x3C/4] = 0x000000FF // Interrupt Line

	// BAR registers: writable (bits above size alignment)
	for i := 0; i < 6; i++ {
		barOffset := 0x10 + (i * 4)
		barValue := cs.BAR(i)
		if barValue == 0 {
			continue
		}

		if barValue&0x01 != 0 {
			masks[barOffset/4] = 0xFFFFFFFC // IO BAR
		} else {
			masks[barOffset/4] = 0xFFFFFFF0 // Memory BAR
		}
	}

	// Expansion ROM BAR
	masks[0x30/4] = 0xFFFFF801

	// Apply capability-specific writemasks (legacy space)
	applyCapabilityWritemasks(cs, masks)

	// Apply extended capability writemasks (0x100+)
	applyExtCapabilityWritemasks(cs, masks)

	return formatCOE(
		"; PCILeechGen - Configuration Space Write Mask (4KB shadow)\n"+
			"; 1 = writable bit, 0 = read-only bit\n"+
			";\n",
		masks,
	)
}

// applyCapabilityWritemasks applies writemasks for known PCI capabilities.
func applyCapabilityWritemasks(cs *pci.ConfigSpace, masks []uint32) {
	caps := pci.ParseCapabilities(cs)
	for _, cap := range caps {
		switch cap.ID {
		case pci.CapIDPowerManagement:
			// PM Control/Status register at cap+4 is partially writable
			if cap.Offset+4 < pci.ConfigSpaceLegacySize {
				masks[(cap.Offset+4)/4] = 0x00008103 // PowerState bits + PME_En + PME_Status
			}
		case pci.CapIDMSI:
			// MSI Message Control is partially writable
			if cap.Offset+4 < pci.ConfigSpaceLegacySize {
				masks[(cap.Offset)/4] |= 0x00710000 // Enable + MultiMsg Enable
			}
		case pci.CapIDMSIX:
			// MSI-X Message Control
			if cap.Offset < pci.ConfigSpaceLegacySize {
				masks[(cap.Offset)/4] |= 0xC0000000 // Enable + Function Mask
			}
		case pci.CapIDPCIExpress:
			// PCIe Device Control at cap+8
			if cap.Offset+8 < pci.ConfigSpaceLegacySize {
				masks[(cap.Offset+8)/4] = 0x0000FFFF
			}
			// PCIe Link Control at cap+16 (0x10)
			if cap.Offset+16 < pci.ConfigSpaceLegacySize {
				masks[(cap.Offset+16)/4] = 0x0000FFFF
			}
		}
	}
}

// applyExtCapabilityWritemasks applies writemasks for PCIe extended capabilities.
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
			// AER: Uncorrectable Error Status at cap+4 (RW1C)
			if wordIdx+1 < len(masks) {
				masks[wordIdx+1] = 0xFFFFFFFF
			}
			// AER: Uncorrectable Error Mask at cap+8
			if wordIdx+2 < len(masks) {
				masks[wordIdx+2] = 0xFFFFFFFF
			}
			// AER: Uncorrectable Error Severity at cap+12
			if wordIdx+3 < len(masks) {
				masks[wordIdx+3] = 0xFFFFFFFF
			}
			// AER: Correctable Error Status at cap+16 (RW1C)
			if wordIdx+4 < len(masks) {
				masks[wordIdx+4] = 0xFFFFFFFF
			}
			// AER: Correctable Error Mask at cap+20
			if wordIdx+5 < len(masks) {
				masks[wordIdx+5] = 0xFFFFFFFF
			}
		case pci.ExtCapIDLTR:
			// LTR: Max Snoop/No-Snoop Latency at cap+4
			if wordIdx+1 < len(masks) {
				masks[wordIdx+1] = 0xFFFFFFFF
			}
		}
	}
}

// GenerateBarZeroCOE generates the pcileech_bar_zero4k.coe file.
// This is a 4KB zero-filled COE file used for BAR response data.
func GenerateBarZeroCOE() string {
	words := make([]uint32, shadowCfgSpaceWords)
	return formatCOE(
		"; PCILeechGen - BAR Zero 4KB\n",
		words,
	)
}
