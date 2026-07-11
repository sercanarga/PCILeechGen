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
// Shadow BRAM is the sole read source under cfgtlp_zero=0, so masks match BAR
// sizes (sizing probes return size-encoded values); identity regs stay read-only.
func GenerateWritemask(cs *pci.ConfigSpace) []uint32 {
	masks := make([]uint32, shadowCfgSpaceWords)

	// 0x40-0xFF capabilities: fully writable. Locking here would desync shadow BRAM from the IP core.
	for i := 0x40 / 4; i < 0x100/4; i++ {
		masks[i] = 0xFFFFFFFF
	}

	// extended config space (0x100-0xFFF): read-only (scrubber already neutralized writable ext caps).

	// header type 0 registers (DWORD index = byte offset / 4)
	masks[0] = 0x00000000 // 0x00: VID:DID (read-only identity)
	masks[1] = 0xFFFFFFFF // 0x04: Command:Status (OS needs full control)
	masks[2] = 0x00000000 // 0x08: RevisionID:ClassCode (read-only identity)
	masks[3] = 0xFF00FFFF // 0x0C: CLS+LT writable, HeaderType RO, BIST writable

	// BAR registers 0x10-0x24 (DWORD 4-9): size-matching masks
	for i := 0; i < 6; i++ {
		barOffset := 0x10 + (i * 4)
		barVal := cs.ReadU32(barOffset)
		dw := barOffset / 4
		if barVal == 0 {
			masks[dw] = 0x00000000 // unused BAR
			continue
		}
		if barVal&0x01 != 0 {
			// I/O BAR: type bits [1:0] read-only
			masks[dw] = barVal & 0xFFFFFFFC
			continue
		}
		// memory BAR lower dword: type+prefetch bits [3:0] read-only
		masks[dw] = barVal & 0xFFFFFFF0
		// 64-bit memory BAR (bits[2:1]=10b): upper dword is address only, fully writable
		if barVal&0x06 == 0x04 && i+1 < 6 {
			masks[dw+1] = 0xFFFFFFFF
			i++ // upper dword of the 64-bit pair consumed
		}
	}

	masks[10] = 0x00000000 // 0x28: CardBus CIS (read-only)
	masks[11] = 0x00000000 // 0x2C: SubsysVID:SubsysDID (read-only identity)
	masks[12] = 0x00000000 // 0x30: Expansion ROM (not implemented on FPGA)
	masks[13] = 0x00000000 // 0x34: CapPtr + reserved (read-only)
	masks[14] = 0x00000000 // 0x38: reserved
	masks[15] = 0x000000FF // 0x3C: IntLine writable, IntPin/MinGnt/MaxLat RO

	return masks
}

// GenerateWritemaskCOE outputs the writemask COE (1=writable, 0=read-only).
func GenerateWritemaskCOE(cs *pci.ConfigSpace) string {
	return formatCOE(
		"; PCILeechGen - Configuration Space Write Mask (4KB shadow)\n"+
			"; 1 = writable bit, 0 = read-only bit\n"+
			"; BAR masks match scrubbed BAR sizes for correct BAR sizing\n"+
			"; Identity registers (VID/DID, ClassCode, SubsysIDs) are read-only\n"+
			";\n",
		GenerateWritemask(cs),
	)
}

// GenerateBarContentCOE outputs BAR shadow COE from the lowest BAR.
// Note: actual BAR aperture (Bar0Size) may be variable (>4KB on large boards); this seeds 4KB shadow.
func GenerateBarContentCOE(barContents map[int][]byte, size int) string {
	n := 1024
	if size > 0 {
		n = (size + 3) / 4
	}
	words := make([]uint32, n)

	data := firmware.LargestBar(barContents)
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
