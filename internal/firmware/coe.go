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

// GenerateBarContentCOE generates the pcileech_bar_zero4k.coe file.
// If BAR memory contents are available from the donor device, they are used to
// populate the BRAM. Otherwise, the file is zero-filled as a safe fallback.
// The first active memory BAR's content is used (pcileech-fpga has a single 4KB BRAM).
func GenerateBarContentCOE(barContents map[int][]byte) string {
	words := make([]uint32, shadowCfgSpaceWords)

	// Find the first BAR with content and populate the BRAM
	if len(barContents) > 0 {
		// Use lowest-indexed BAR content available
		bestIdx := -1
		for idx := range barContents {
			if bestIdx == -1 || idx < bestIdx {
				bestIdx = idx
			}
		}
		if bestIdx >= 0 {
			data := barContents[bestIdx]
			// Convert bytes to uint32 words (little-endian)
			for i := 0; i+3 < len(data) && i/4 < len(words); i += 4 {
				words[i/4] = uint32(data[i]) |
					uint32(data[i+1])<<8 |
					uint32(data[i+2])<<16 |
					uint32(data[i+3])<<24
			}
		}
	}

	header := "; PCILeechGen - BAR Content (4KB shadow)\n"
	if len(barContents) > 0 {
		header += "; Populated from donor device BAR memory\n"
	} else {
		header += "; Zero-filled (no donor BAR data available)\n"
	}
	header += ";\n"

	return formatCOE(header, words)
}

// ScrubBarContent patches BAR data for device-class-specific quirks.
// Call before GenerateBarContentCOE.
func ScrubBarContent(barContents map[int][]byte, classCode uint32) {
	data := lowestBarData(barContents)
	if data == nil {
		return
	}
	switch classCode {
	case 0x010802: // NVMe
		scrubNVMeBar0(data)
	case 0x0C0330: // xHCI USB 3.0
		scrubXHCIBar0(data)
	}
}

// lowestBarData returns the byte slice for the lowest-indexed BAR, or nil.
func lowestBarData(barContents map[int][]byte) []byte {
	bestIdx := -1
	for idx := range barContents {
		if bestIdx == -1 || idx < bestIdx {
			bestIdx = idx
		}
	}
	if bestIdx < 0 {
		return nil
	}
	return barContents[bestIdx]
}

// readLE32 reads a little-endian uint32 from a byte slice at the given offset.
func readLE32(data []byte, off int) uint32 {
	return uint32(data[off]) | uint32(data[off+1])<<8 |
		uint32(data[off+2])<<16 | uint32(data[off+3])<<24
}

// writeLE32 writes a little-endian uint32 to a byte slice at the given offset.
func writeLE32(data []byte, off int, val uint32) {
	data[off] = byte(val)
	data[off+1] = byte(val >> 8)
	data[off+2] = byte(val >> 16)
	data[off+3] = byte(val >> 24)
}

// scrubNVMeBar0 sets CSTS.RDY=1 and CC.EN=1 in the BAR0 snapshot.
// stornvme.sys polls CSTS.RDY after writing CC.EN — static BRAM can't
// flip the bit, so we pre-load it.
func scrubNVMeBar0(data []byte) {
	if len(data) < 0x20 {
		return
	}

	// CSTS @ 0x1C: RDY=1, clear CFS + SHST
	csts := readLE32(data, 0x1C)
	csts |= 0x01
	csts &= ^uint32(0x02 | 0x0C)
	writeLE32(data, 0x1C, csts)

	// CC @ 0x14: EN=1 (coherent with RDY)
	cc := readLE32(data, 0x14)
	cc |= 0x01
	writeLE32(data, 0x14, cc)
}

const bramSize = 0x1000 // 4KB BAR BRAM

// scrubXHCIBar0 patches xHCI BAR0 registers to fit within the 4KB BRAM
// and fakes a running controller state (R/S=1, HCH=0).
func scrubXHCIBar0(data []byte) {
	if len(data) < 0x20 {
		return
	}

	capLen := int(data[0x00]) // operational regs base
	if capLen == 0 || capLen > 0x40 {
		capLen = 0x20
	}

	if capLen+0x40 > len(data) {
		return
	}

	// HCSPARAMS1 (0x04)
	hcsparams1 := readLE32(data, 0x04)
	maxSlots := int(hcsparams1 & 0xFF)
	if maxSlots == 0 {
		maxSlots = 32
	}
	maxIntrs := int((hcsparams1 >> 8) & 0x7FF)
	maxPorts := int((hcsparams1 >> 24) & 0xFF)

	// HCSPARAMS2 (0x08): nuke scratchpad counts, BRAM can't handle them
	hcsparams2 := readLE32(data, 0x08)
	hcsparams2 &= ^uint32(0xFFE00000)
	writeLE32(data, 0x08, hcsparams2)

	// HCCPARAMS1 (0x10): kill xECP if it points outside BRAM
	hccparams1 := readLE32(data, 0x10)
	xecp := int((hccparams1 >> 16) & 0xFFFF)
	if xecp*4 >= bramSize {
		hccparams1 &= 0x0000FFFF
		writeLE32(data, 0x10, hccparams1)
	}

	// DBOFF (0x14): doorbell array, (MaxSlots+1)*4 bytes
	dboff := readLE32(data, 0x14) & ^uint32(0x03)
	doorbellSize := (maxSlots + 1) * 4

	if int(dboff)+doorbellSize > bramSize {
		newDBOFF := bramSize - doorbellSize
		newDBOFF = newDBOFF & ^0x1F // align down 32B
		if newDBOFF < capLen+0x20 {
			// not enough room, shrink MaxSlots
			available := bramSize - (capLen + 0x20)
			maxSlots = available/4 - 1
			if maxSlots < 1 {
				maxSlots = 1
			}
			doorbellSize = (maxSlots + 1) * 4
			newDBOFF = bramSize - doorbellSize
			newDBOFF = newDBOFF & ^0x1F
		}
		writeLE32(data, 0x14, uint32(newDBOFF))
	}

	// RTSOFF (0x18): runtime regs, each interrupter takes 0x20 bytes
	rtsoff := int(readLE32(data, 0x18) & ^uint32(0x1F))

	// clamp MaxIntrs to fit
	if rtsoff > 0 && maxIntrs > 0 {
		maxFit := (bramSize - rtsoff - 0x20) / 0x20
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
		newRTSOFF = (newRTSOFF + 0x1F) & ^0x1F // align up 32B
		if newRTSOFF+runtimeSize > bramSize {
			newRTSOFF = bramSize - runtimeSize
			newRTSOFF = newRTSOFF & ^0x1F
		}
		rtsoff = newRTSOFF
		writeLE32(data, 0x18, uint32(rtsoff))

		// re-check after moving
		maxFit := (bramSize - rtsoff - 0x20) / 0x20
		if maxFit < 1 {
			maxFit = 1
		}
		if maxIntrs > maxFit {
			maxIntrs = maxFit
		}
	}

	// MaxPorts: port regs start at capLen+0x400, 0x10 each
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

	// write back clamped HCSPARAMS1
	hcsparams1 = uint32(maxSlots) | (uint32(maxIntrs) << 8) | (uint32(maxPorts) << 24)
	writeLE32(data, 0x04, hcsparams1)

	// PAGESIZE: 4KB
	writeLE32(data, capLen+0x08, 0x01)

	// DNCTRL + CRCR: clear, irrelevant for static BRAM
	writeLE32(data, capLen+0x14, 0x00)
	writeLE32(data, capLen+0x18, 0x00)
	writeLE32(data, capLen+0x1C, 0x00)

	// CONFIG: MaxSlotsEn = clamped MaxSlots
	config := readLE32(data, capLen+0x38)
	config = (config & 0xFFFFFF00) | uint32(maxSlots)
	writeLE32(data, capLen+0x38, config)

	// USBCMD: R/S=1, HCRST=0
	usbcmd := readLE32(data, capLen)
	usbcmd |= 0x01
	usbcmd &= ^uint32(0x02)
	writeLE32(data, capLen, usbcmd)

	// USBSTS: HCH=0, HSE=0
	usbsts := readLE32(data, capLen+4)
	usbsts &= ^uint32(0x01 | 0x04)
	writeLE32(data, capLen+4, usbsts)
}
