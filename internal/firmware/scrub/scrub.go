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

const BRAMSize = 4096

const bar0SizeMask = 0xFFFFF000 // 4 KB aligned

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
		size += 8 // mask (4B) + pending (4B)
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

// vendorCapRange defines a known vendor-specific register region worth preserving.
type vendorCapRange struct {
	VendorID uint16
	Start    int
	End      int // exclusive
	Name     string
}

// knownVendorCaps lists vendor-specific register regions that should NOT be zeroed.
// Zeroing these can trigger driver errors or detection (e.g. missing firmware status).
var knownVendorCaps = []vendorCapRange{
	{0x8086, 0x40, 0x60, "Intel PCIe advanced error"},
	{0x8086, 0xE0, 0x100, "Intel device-specific config"},
	{0x10EC, 0x40, 0x60, "Realtek PHY control"},
	{0x10EC, 0x80, 0xA0, "Realtek LED/WOL config"},
	{0x14E4, 0x48, 0x60, "Broadcom device control"},
	{0x144D, 0x40, 0x50, "Samsung NVMe PM region"},
	{0x168C, 0x40, 0x70, "Qualcomm Atheros radio config"},
	{0x1912, 0xF0, 0x100, "Renesas firmware status"},
	{0x1B21, 0x40, 0x60, "ASMedia link training"},
	{0x1B73, 0xA0, 0xC0, "Fresco Logic xHCI quirks"},
	{0x11AB, 0x40, 0x60, "Marvell device control"},
	{0x11AB, 0x70, 0x90, "Marvell PHY/LED config"},
	{0x15B3, 0x60, 0x80, "Mellanox device-specific"},
	{0x1D6A, 0x40, 0x60, "Aquantia PHY control"},
	{0x14C3, 0x40, 0x60, "MediaTek Wi-Fi config"},
	{0x1987, 0x40, 0x50, "Phison NVMe device config"},
	{0x1C5C, 0x40, 0x50, "SK Hynix NVMe config"},
	{0x1344, 0x40, 0x50, "Micron NVMe config"},
	{0x15B7, 0x40, 0x50, "SanDisk/WD NVMe config"},
	{0x1E0F, 0x40, 0x50, "KIOXIA NVMe config"},
}

// vendorWhitelist returns a coverage bitmap for vendor-specific register regions
// that should be preserved for the given vendor ID.
func vendorWhitelist(vid uint16) []bool {
	covered := make([]bool, pci.ConfigSpaceLegacySize)
	for _, vc := range knownVendorCaps {
		if vc.VendorID == vid {
			for i := vc.Start; i < vc.End && i < pci.ConfigSpaceLegacySize; i++ {
				covered[i] = true
			}
		}
	}
	return covered
}

// zeroVendorRegisters clears 0x40-0xFF bytes not belonging to any cap
// or known vendor-specific region. Uses pre-parsed caps from ScrubContext
// to avoid re-parsing the already-modified config space.
func zeroVendorRegisters(cs *pci.ConfigSpace, om *overlay.Map, caps []pci.Capability) {
	covered := make([]bool, pci.ConfigSpaceLegacySize)
	for i := 0; i < 0x40; i++ {
		covered[i] = true
	}

	for _, cap := range caps {
		size := capSizeAt(cs, cap.ID, cap.Offset)
		for i := cap.Offset; i < cap.Offset+size && i < pci.ConfigSpaceLegacySize; i++ {
			covered[i] = true
		}
	}

	// preserve known vendor-specific register regions
	vid := cs.VendorID()
	whitelist := vendorWhitelist(vid)
	for i := 0x40; i < pci.ConfigSpaceLegacySize; i++ {
		if whitelist[i] {
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

	ctx := &ScrubContext{
		Caps:      pci.ParseCapabilities(scrubbed),
		ExtCaps:   pci.ParseExtCapabilities(scrubbed),
		ClassCode: cs.ReadU32(0x08) >> 8, // bytes 0x09-0x0B: ProgIf + SubClass + BaseClass
	}

	for _, pass := range defaultPipeline() {
		pass.Apply(scrubbed, b, om, ctx)
	}

	return scrubbed, om
}

// clampBARsToFPGA shrinks all memory BARs to 4 KB and I/O BARs to 256 bytes.
func clampBARsToFPGA(cs *pci.ConfigSpace, om *overlay.Map) {
	for i := 0; i < 6; i++ {
		barOffset := 0x10 + (i * 4)
		barVal := cs.BAR(i)
		if barVal == 0 {
			continue
		}

		if barVal&0x01 != 0 {
			// I/O BAR: clamp to 256 bytes (mask = 0xFFFFFF01)
			newBar := uint32(0xFFFFFF00) | (barVal & 0x03)
			om.WriteU32(barOffset, newBar, fmt.Sprintf("clamp I/O BAR%d to 256 bytes", i))
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
func relocateMSIXToBRAM(cs *pci.ConfigSpace, om *overlay.Map, caps []pci.Capability) {
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

		// BIR must be 0 — FPGA only serves BAR0
		tableReg := newTableOffset & 0xFFFFFFF8 // BIR = 0
		pbaReg := newPBAOffset & 0xFFFFFFF8     // BIR = 0

		om.WriteU32(cap.Offset+4, tableReg,
			fmt.Sprintf("relocate MSI-X table to 0x%X (%d vectors, BIR=0)", newTableOffset, info.TableSize))
		om.WriteU32(cap.Offset+8, pbaReg,
			fmt.Sprintf("relocate MSI-X PBA to 0x%X (%d bytes, BIR=0)", newPBAOffset, pbaSize))

		msgCtl := cs.ReadU16(cap.Offset + 2)
		msgCtl |= 0x8000
		msgCtl &= ^uint16(0x4000)
		om.WriteU16(cap.Offset+2, msgCtl, "MSI-X enable (BRAM replicated)")

		break
	}
}
