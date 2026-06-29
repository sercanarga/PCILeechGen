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
//
// when shadow config space is enabled (cfgtlp_zero=0), the shadow BRAM is the
// sole source for config space reads. BAR sizing probes (host writes 0xFFFFFFFF
// then reads back) go through the BRAM, not the IP core. the writemask must
// limit writable bits to match the BAR size mask so the host reads back the
// correct size-encoded value (e.g. 0xFFFFF000 for 4KB; variable Bar0Size supported via boards).
//
// identity registers (VID/DID, ClassCode, SubsysIDs) are read-only to prevent
// host writes from corrupting device identity.
func GenerateWritemaskCOE(cs *pci.ConfigSpace) string {
	masks := make([]uint32, shadowCfgSpaceWords)

	// Standard capability window (0x40-0xFF): writable by default; the loop
	// below locks the individual read-only fields. The shadow is read-modify-
	// write (wr_dina = mask ? wr_data : rd_data), so a writable bit stores the
	// host value and reads it back, while a locked bit ignores the write and
	// returns the baked value - exactly real-HW RW vs RO semantics. Control
	// bits a driver writes and reads back (e.g. PMCSR power state, MSI/MSI-X
	// enable, PCIe Device/Link Control) MUST stay writable, or the OS sees its
	// write not take and stalls; only genuinely read-only fields are locked.
	for i := 0x40 / 4; i < 0x100/4; i++ {
		masks[i] = 0xFFFFFFFF
	}

	// Lock the read-only fields inside the capability window. The window is
	// writable by default (above) so control/status bits stay in sync with the
	// Xilinx IP core, but a capability's structural fields (ID, next pointer)
	// and its read-only registers (PMC, the PCIe Capabilities/DevCap/LinkCap
	// registers, the MSI-X Table/PBA offset registers, ...) never accept writes
	// on real hardware. Leaving them writable lets a detector write one and read
	// it back changed - a write-to-RO-then-readback tell real silicon never
	// produces. Lock only the unambiguous read-only fields, bounded to each
	// capability's own extent so a short cap can't clobber its neighbour.
	for _, cap := range pci.ParseCapabilities(cs) {
		hdr := cap.Offset / 4
		if hdr < 0x40/4 || hdr >= 0x100/4 {
			continue
		}
		capLen := len(cap.Data)
		inCap := func(rel int) (int, bool) {
			if rel+4 > capLen {
				return 0, false
			}
			dw := (cap.Offset + rel) / 4
			return dw, dw >= 0x40/4 && dw < 0x100/4
		}
		lockRO := func(rel int) { // whole DWORD at cap.Offset+rel read-only
			if dw, ok := inCap(rel); ok {
				masks[dw] = 0x00000000
			}
		}
		lockStatusHalf := func(rel int) { // upper 16 bits RO (status), lower 16 (control) stay RW
			if dw, ok := inCap(rel); ok {
				masks[dw] &^= 0xFFFF0000
			}
		}
		masks[hdr] &^= 0x0000FFFF // ID + next pointer are always read-only
		switch cap.ID {
		case pci.CapIDPowerManagement:
			masks[hdr] = 0x00000000 // header DWORD also holds PMC (RO); PMCSR at +4 stays writable
		case pci.CapIDMSI:
			// Message Control (upper 16): only MSI Enable (bit0) and Multiple
			// Message Enable (bits 6:4) are writable; the rest (64-bit/PVM
			// capable, Multiple Message Capable) is read-only. Address/Data
			// registers at +4 onward stay writable for the OS to program.
			masks[hdr] = 0x00710000
		case pci.CapIDPCIExpress:
			masks[hdr] = 0x00000000 // header DWORD holds the PCIe Capabilities register (RO)
			lockRO(0x04)            // Device Capabilities (RO)
			lockStatusHalf(0x08)    // Device Status (RO/RW1C); Device Control stays RW
			lockRO(0x0C)            // Link Capabilities (RO)
			lockStatusHalf(0x10)    // Link Status (RO/RW1C); Link Control stays RW
			lockStatusHalf(0x18)    // Slot Status (RO/RW1C); Slot Control stays RW
			lockStatusHalf(0x1C)    // Root Capabilities (RO, upper); Root Control stays RW
			lockRO(0x20)            // Root Status (RO/RW1C)
			lockRO(0x24)            // Device Capabilities 2 (RO)
			lockRO(0x2C)            // Link Capabilities 2 (RO)
		case pci.CapIDMSIX:
			masks[hdr] = 0xC0000000 // keep only MSI-X Enable + Function Mask writable
			lockRO(0x04)            // Table Offset/BIR (RO)
			lockRO(0x08)            // PBA Offset/BIR (RO)
		}
	}

	// extended config space (0x100-0xFFF): read-only.
	// scrubber already neutralized all writable extended cap fields.

	// header type 0 registers (DWORD index = byte offset / 4)
	masks[0] = 0x00000000 // 0x00: VID:DID (read-only identity)
	masks[1] = 0x0000FFFF // 0x04: Command (RW); Status (upper 16) is RO/RW1C - writes ignored like real HW
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
		} else {
			// memory BAR: type+prefetch bits [3:0] read-only
			masks[dw] = barVal & 0xFFFFFFF0
		}
	}

	masks[10] = 0x00000000 // 0x28: CardBus CIS (read-only)
	masks[11] = 0x00000000 // 0x2C: SubsysVID:SubsysDID (read-only identity)
	masks[12] = 0x00000000 // 0x30: Expansion ROM (not implemented on FPGA)
	masks[13] = 0x00000000 // 0x34: CapPtr + reserved (read-only)
	masks[14] = 0x00000000 // 0x38: reserved
	masks[15] = 0x000000FF // 0x3C: IntLine writable, IntPin/MinGnt/MaxLat RO

	return formatCOE(
		"; PCILeechGen - Configuration Space Write Mask (4KB shadow)\n"+
			"; 1 = writable bit, 0 = read-only bit\n"+
			"; BAR masks match scrubbed BAR sizes for correct BAR sizing\n"+
			"; Identity registers (VID/DID, ClassCode, SubsysIDs) are read-only\n"+
			";\n",
		masks,
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

// GenerateBarInitHex outputs a donor BAR snapshot in $readmemh format, one
// DWORD per line, sized to sizeBytes/4 words (padded with zero, truncated to
// fit). The generic BRAM fallback ($readmemh into bar_mem) uses this so an
// unknown-class device returns the donor's real register values instead of all
// zeros - a driver/detector reading a known donor register no longer sees 0.
func GenerateBarInitHex(barData []byte, sizeBytes int) string {
	words := sizeBytes / 4
	if words <= 0 {
		words = 1024
	}

	var sb strings.Builder
	sb.WriteString("// PCILeechGen - BAR Snapshot Init (generic BRAM seed)\n")
	sb.WriteString(fmt.Sprintf("// %d DWORDs from donor BAR memory\n", words))
	for i := 0; i < words; i++ {
		off := i * 4
		var w uint32
		if off+4 <= len(barData) {
			w = util.ReadLE32(barData, off)
		}
		sb.WriteString(fmt.Sprintf("%08X\n", w))
	}
	return sb.String()
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
