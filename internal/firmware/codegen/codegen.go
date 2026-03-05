// Package codegen emits COE and HEX files for Vivado BRAM init.
package codegen

import (
	"fmt"
	"strings"

	"github.com/sercanarga/pcileechgen/internal/firmware"
	"github.com/sercanarga/pcileechgen/internal/pci"
	"github.com/sercanarga/pcileechgen/internal/util"
)

const shadowCfgSpaceWords = 1024 // 4KB shadow BRAM

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

// GenerateConfigSpaceCOE outputs 1024 DWORDs (4KB) for pcileech_cfgspace.coe.
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

// GenerateWritemaskCOE outputs the writemask COE (1=writable, 0=read-only).
func GenerateWritemaskCOE(cs *pci.ConfigSpace) string {
	masks := make([]uint32, shadowCfgSpaceWords)

	masks[0x04/4] = 0x0000FFFF // Command
	masks[0x0C/4] = 0x0000FF00 // Latency Timer
	masks[0x3C/4] = 0x000000FF // Interrupt Line

	// BARs
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

	masks[0x30/4] = 0xFFFFF801 // Expansion ROM

	applyCapabilityWritemasks(cs, masks)
	applyExtCapabilityWritemasks(cs, masks)

	return formatCOE(
		"; PCILeechGen - Configuration Space Write Mask (4KB shadow)\n"+
			"; 1 = writable bit, 0 = read-only bit\n"+
			";\n",
		masks,
	)
}

func applyCapabilityWritemasks(cs *pci.ConfigSpace, masks []uint32) {
	caps := pci.ParseCapabilities(cs)
	for _, cap := range caps {
		switch cap.ID {
		case pci.CapIDPowerManagement:
			if cap.Offset+4+4 <= pci.ConfigSpaceLegacySize {
				masks[(cap.Offset+4)/4] = 0x00008103 // PowerState + PME_En + PME_Status
			}
		case pci.CapIDMSI:
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
		case pci.CapIDMSIX:
			if cap.Offset < pci.ConfigSpaceLegacySize {
				masks[(cap.Offset)/4] |= 0xC0000000 // Enable + Function Mask
			}
		case pci.CapIDPCIExpress:
			if cap.Offset+8+4 <= pci.ConfigSpaceLegacySize {
				masks[(cap.Offset+8)/4] = 0x0000FFFF // DevCtl
			}
			if cap.Offset+16+4 <= pci.ConfigSpaceLegacySize {
				masks[(cap.Offset+16)/4] = 0x0000FFFF // LinkCtl
			}
		}
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

// GenerateBarContentCOE outputs BAR shadow COE from the lowest BAR.
func GenerateBarContentCOE(barContents map[int][]byte) string {
	words := make([]uint32, shadowCfgSpaceWords)

	data := firmware.LowestBarData(barContents)
	if data != nil {
		for i := 0; i+4 <= len(data) && i/4 < len(words); i += 4 {
			words[i/4] = util.ReadLE32(data, i)
		}
	}

	header := "; PCILeechGen - BAR Content (4KB shadow)\n"
	if data != nil {
		header += "; Populated from donor device BAR memory\n"
	} else {
		header += "; Zero-filled (no donor BAR data available)\n"
	}
	header += ";\n"

	return formatCOE(header, words)
}

// GenerateConfigSpaceHex outputs config space in $readmemh format (1024 DWORDs).
func GenerateConfigSpaceHex(cs *pci.ConfigSpace) string {
	var sb strings.Builder
	sb.WriteString("// PCILeechGen - Config Space Init (4KB = 1024 DWORDs)\n")
	sb.WriteString(fmt.Sprintf("// Device: %04X:%04X\n", cs.VendorID(), cs.DeviceID()))

	for i := 0; i < shadowCfgSpaceWords; i++ {
		word := cs.ReadU32(i * 4)
		sb.WriteString(fmt.Sprintf("%08X // [%03X]\n", word, i*4))
	}

	return sb.String()
}

// GenerateMSIXTableHex outputs MSI-X table entries in $readmemh format.
func GenerateMSIXTableHex(entries []pci.MSIXEntry) string {
	var sb strings.Builder
	sb.WriteString("// PCILeechGen - MSI-X Table Init\n")
	sb.WriteString(fmt.Sprintf("// %d vectors (%d DWORDs)\n", len(entries), len(entries)*4))

	for i, e := range entries {
		ctrl := e.Control | 0x01 // masked on init
		sb.WriteString(fmt.Sprintf("%08X // [%d] addr_lo\n", e.AddrLo, i))
		sb.WriteString(fmt.Sprintf("%08X // [%d] addr_hi\n", e.AddrHi, i))
		sb.WriteString(fmt.Sprintf("%08X // [%d] data\n", e.Data, i))
		sb.WriteString(fmt.Sprintf("%08X // [%d] control (masked)\n", ctrl, i))
	}

	return sb.String()
}
