// Package scrub sanitizes donor config space for safe FPGA replay.
package scrub

import (
	"fmt"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/firmware/overlay"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

const (
	cmdMask    = 0x0547 // keep BusMaster, IO, Memory, SERR, ParityErr
	cmdForce   = 0x0006 // always set BME(2) + MSE(1) — donor may lack these
	statusMask = 0x06F0 // keep 66MHz, FastB2B, CapList, DevSel bits
)

// ext caps the FPGA can't emulate
var unsafeExtCaps = map[uint16]string{
	pci.ExtCapIDSRIOV:         "SR-IOV",
	pci.ExtCapIDMRIOV:         "MR-IOV",
	pci.ExtCapIDResizableBAR:  "Resizable BAR",
	pci.ExtCapIDATS:           "ATS",
	pci.ExtCapIDPageRequest:   "Page Request",
	pci.ExtCapIDPASID:         "PASID",
	pci.ExtCapIDL1PMSubstates: "L1 PM Substates",
	pci.ExtCapIDDPC:           "DPC",
	pci.ExtCapIDPTM:           "PTM",
	pci.ExtCapIDSecondaryPCIe: "Secondary PCIe",
	pci.ExtCapIDMulticast:     "Multicast",
}

const BRAMSize = 4096

const bar0SizeMask = 0xFFFFF000 // 4 KB aligned

func IsUnsafeExtCap(id uint16) bool {
	_, ok := unsafeExtCaps[id]
	return ok
}

func UnsafeExtCapName(id uint16) string {
	if name, ok := unsafeExtCaps[id]; ok {
		return name
	}
	return ""
}

// min sizes (bytes) for standard PCI caps
var capMinSize = map[uint8]int{
	pci.CapIDPowerManagement: 8,  // PM: 2 header + 2 PMC + 2 PMCSR + 2 data
	pci.CapIDMSIX:            12, // MSI-X: 2 header + 2 ctl + 4 table + 4 PBA
	pci.CapIDPCIExpress:      60, // PCIe: 0x3C typical for v2
	pci.CapIDVendorSpecific:  3,  // at least header + length byte
}

// computeMSISize returns the actual MSI cap size (10-24 bytes).
func computeMSISize(cs *pci.ConfigSpace, capOffset int) int {
	if capOffset+4 > cs.Size {
		return 10
	}
	msgCtl := cs.ReadU16(capOffset + 2)
	is64bit := (msgCtl & (1 << 7)) != 0
	hasMasking := (msgCtl & (1 << 8)) != 0

	size := 10 // header + msgctl + addr + data
	if is64bit {
		size += 4 // upper address dword
	}
	if hasMasking {
		size += 2 + 8 // reserved + mask + pending
	}
	return size
}

// capSizeAt returns byte size for a cap; MSI is variable, rest is static.
func capSizeAt(cs *pci.ConfigSpace, id uint8, offset int) int {
	if id == pci.CapIDMSI {
		return computeMSISize(cs, offset)
	}
	if s, ok := capMinSize[id]; ok {
		return s
	}
	return 8
}

// zeroVendorRegisters clears 0x40-0xFF bytes not belonging to any cap.
func zeroVendorRegisters(cs *pci.ConfigSpace, om *overlay.Map) {
	covered := make([]bool, pci.ConfigSpaceLegacySize)
	for i := 0; i < 0x40; i++ {
		covered[i] = true
	}

	caps := pci.ParseCapabilities(cs)
	for _, cap := range caps {
		size := capSizeAt(cs, cap.ID, cap.Offset)
		for i := cap.Offset; i < cap.Offset+size && i < pci.ConfigSpaceLegacySize; i++ {
			covered[i] = true
		}
	}

	for i := 0x40; i < pci.ConfigSpaceLegacySize; i++ {
		if !covered[i] {
			om.WriteU8(i, 0x00, "clear uncovered vendor register")
		}
	}
}

// ScrubConfigSpace shorthand — returns scrubbed copy, discards diff.
func ScrubConfigSpace(cs *pci.ConfigSpace, b *board.Board) *pci.ConfigSpace {
	scrubbed, _ := ScrubConfigSpaceWithOverlay(cs, b)
	return scrubbed
}

// ScrubConfigSpaceWithOverlay scrubs the config space and returns both the
// scrubbed copy and an overlay map recording every change.
func ScrubConfigSpaceWithOverlay(cs *pci.ConfigSpace, b *board.Board) (*pci.ConfigSpace, *overlay.Map) {
	scrubbed := cs.Clone()
	om := overlay.NewMap(scrubbed)

	for _, pass := range defaultPipeline() {
		pass.Apply(scrubbed, b, om)
	}

	return scrubbed, om
}

// clampBARsToFPGA shrinks all memory BARs to 4 KB.
func clampBARsToFPGA(cs *pci.ConfigSpace, om *overlay.Map) {
	for i := 0; i < 6; i++ {
		barOffset := 0x10 + (i * 4)
		barVal := cs.BAR(i)
		if barVal == 0 {
			continue
		}

		if barVal&0x01 != 0 {
			continue
		}

		is64bit := (barVal & 0x06) == 0x04
		newBar := bar0SizeMask | (barVal & 0x0F)
		om.WriteU32(barOffset, newBar, fmt.Sprintf("clamp BAR%d to 4 KB", i))

		if is64bit && i < 5 {
			om.WriteU32(barOffset+4, 0x00000000, fmt.Sprintf("clear BAR%d upper 32 bits", i+1))
			i++ // skip upper half
		}
	}
}

// relocateMSIXToBRAM moves MSI-X table/PBA offsets to 0x1000+ so the
// separate BRAM module can serve them. Keeps MSI-X enabled.
func relocateMSIXToBRAM(cs *pci.ConfigSpace, om *overlay.Map) {
	caps := pci.ParseCapabilities(cs)
	for _, cap := range caps {
		if cap.ID != pci.CapIDMSIX {
			continue
		}
		if cap.Offset+12 > pci.ConfigSpaceLegacySize {
			continue
		}

		info := pci.ParseMSIXCap(cs)
		if info == nil {
			continue
		}

		tableSize := info.TableSize * 16
		pbaSize := (info.TableSize + 63) / 64 * 8
		if pbaSize < 8 {
			pbaSize = 8
		}

		newTableOffset := uint32(0x1000)
		newPBAOffset := newTableOffset + uint32(tableSize)
		newPBAOffset = (newPBAOffset + 7) &^ 7

		tableReg := (newTableOffset & 0xFFFFFFF8) | uint32(info.TableBIR)
		pbaReg := (newPBAOffset & 0xFFFFFFF8) | uint32(info.PBABIR)

		om.WriteU32(cap.Offset+4, tableReg,
			fmt.Sprintf("relocate MSI-X table to 0x%X (%d vectors)", newTableOffset, info.TableSize))
		om.WriteU32(cap.Offset+8, pbaReg,
			fmt.Sprintf("relocate MSI-X PBA to 0x%X (%d bytes)", newPBAOffset, pbaSize))

		msgCtl := cs.ReadU16(cap.Offset + 2)
		msgCtl |= 0x8000
		msgCtl &= ^uint16(0x4000)
		om.WriteU16(cap.Offset+2, msgCtl, "MSI-X enable (BRAM replicated)")

		break
	}
}
