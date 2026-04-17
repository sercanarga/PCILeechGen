// Package codegen emits COE and HEX files for Vivado BRAM init.
package codegen

import (
	"fmt"
	"log/slog"
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
// correct size-encoded value (e.g. 0xFFFFF000 for 4KB).
//
// identity registers (VID/DID, ClassCode, SubsysIDs) are read-only to prevent
// host writes from corrupting device identity.
func GenerateWritemaskCOE(cs *pci.ConfigSpace) string {
	masks := make([]uint32, shadowCfgSpaceWords)

	// standard config space capabilities (0x40-0xFF): start fully writable,
	// then lock critical PM and PCIe fields to prevent host from killing the link.
	for i := 0x40 / 4; i < 0x100/4; i++ {
		masks[i] = 0xFFFFFFFF
	}

	// extended config space (0x100-0xFFF): read-only.
	// scrubber already neutralized all writable extended cap fields.

	// header type 0 registers (DWORD index = byte offset / 4)
	masks[0] = 0x00000000  // 0x00: VID:DID (read-only identity)
	masks[1] = 0xFFFFFFFF  // 0x04: Command:Status (OS needs full control)
	masks[2] = 0x00000000  // 0x08: RevisionID:ClassCode (read-only identity)
	masks[3] = 0xFF00FFFF  // 0x0C: CLS+LT writable, HeaderType RO, BIST writable

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

	// lock PM and PCIe capability fields to prevent D3 / ASPM re-enable.
	// without this, Windows writes D3hot to PMCSR after Code 10, the Xilinx
	// IP core transitions to D3, and DMA stops until power cycle.
	lockCapabilityMasks(cs, masks)

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
func GenerateBarContentCOE(barContents map[int][]byte) string {
	words := make([]uint32, shadowCfgSpaceWords)

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

// lockCapabilityMasks walks the capability chain and makes PM and PCIe
// capability DWORDs read-only in the writemask. this prevents the host
// OS from:
//   - writing D3hot to PMCSR (kills TLP processing, DMA stops)
//   - re-enabling ASPM in PCIe LinkCtl (link enters L1, FPGA can't exit)
//   - modifying MPS/MRRS in DevCtl (mismatch with IP core)
//   - setting LTR in DevCtl2 (platform throttles link)
//
// MSI/MSI-X capabilities remain fully writable so the OS can program
// interrupt addresses and data for driver operation.
func lockCapabilityMasks(cs *pci.ConfigSpace, masks []uint32) {
	caps := pci.ParseCapabilities(cs)
	for _, cap := range caps {
		switch cap.ID {
		case pci.CapIDPowerManagement:
			lockPMCap(cap.Offset, masks)
		case pci.CapIDPCIExpress:
			lockPCIeCap(cap.Offset, masks)
		}
	}
}

// lockPMCap makes the PM capability read-only.
// cap+0: [PMC : NextPtr : CapID] - all RO
// cap+4: [BSE+Data : PMCSR] - PowerState bits [1:0] RO (prevents D3)
func lockPMCap(offset int, masks []uint32) {
	setMask := func(off int, mask uint32) {
		dw := off / 4
		if dw >= 0 && dw < len(masks) {
			masks[dw] = mask
		}
	}
	// header + PMC: fully RO
	setMask(offset, 0x00000000)
	// PMCSR: only PME_Status (bit 15) writable for RW1C clearing.
	// PowerState [1:0] = RO, NoSoftReset [3] = RO, PME_En [8] = RO.
	// BSE + Data (upper 16 bits) = RO.
	setMask(offset+4, 0x00008000)

	slog.Debug("locked PM capability writemask",
		"offset", fmt.Sprintf("0x%02X", offset))
}

// lockPCIeCap makes the PCIe capability mostly read-only.
// the scrubber already set correct link speed, ASPM=off, MPS=128, etc.
// locking these prevents Windows from re-enabling ASPM or modifying
// link parameters after scrubbing.
func lockPCIeCap(offset int, masks []uint32) {
	setMask := func(off int, mask uint32) {
		dw := off / 4
		if dw >= 0 && dw < len(masks) {
			masks[dw] = mask
		}
	}

	// cap+0x00: [PCIeCapReg : NextPtr : CapID] - RO
	setMask(offset+0x00, 0x00000000)
	// cap+0x04: DevCap - RO
	setMask(offset+0x04, 0x00000000)
	// cap+0x08: [DevStatus : DevCtl] - RO
	// DevCtl: MPS/MRRS must match IP core, FLR dangerous
	// DevStatus: RW1C bits, but harmless to lock
	setMask(offset+0x08, 0x00000000)
	// cap+0x0C: LinkCap - RO
	setMask(offset+0x0C, 0x00000000)
	// cap+0x10: [LinkStatus : LinkCtl]
	// LinkCtl: ASPM [1:0] and ClockPM [8] MUST be RO
	// allow harmless bits: CommonClock [6], ExtSynch [7], BW interrupts [11:10]
	// LinkStatus: RO + RW1C, allow RW1C clearing
	setMask(offset+0x10, 0xFFFF0CC0)
	// cap+0x14 to cap+0x23: slot/root regs - RO (endpoint, unused)
	for i := 0x14; i <= 0x20; i += 4 {
		setMask(offset+i, 0x00000000)
	}
	// cap+0x24: DevCap2 - RO
	setMask(offset+0x24, 0x00000000)
	// cap+0x28: [DevStatus2 : DevCtl2]
	// DevCtl2 LTR [10] must be RO, rest mostly safe
	setMask(offset+0x28, 0x00000000)
	// cap+0x2C: LinkCap2 - RO
	setMask(offset+0x2C, 0x00000000)
	// cap+0x30: [LinkStatus2 : LinkCtl2]
	// LinkCtl2 target speed [3:0] - RO
	setMask(offset+0x30, 0x00000000)
	// cap+0x34 to cap+0x3B: slot2 regs - RO
	for i := 0x34; i <= 0x38; i += 4 {
		setMask(offset+i, 0x00000000)
	}

	slog.Debug("locked PCIe capability writemask",
		"offset", fmt.Sprintf("0x%02X", offset))
}
