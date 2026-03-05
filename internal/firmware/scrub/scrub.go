// Package scrub sanitizes donor config space for safe FPGA replay.
package scrub

import (
	"fmt"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/firmware"
	"github.com/sercanarga/pcileechgen/internal/firmware/overlay"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

const (
	cmdMask    = 0x0547 // keep BusMaster, IO, Memory, SERR, ParityErr
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

// clampLinkCapability caps link speed/width to board limits.
func clampLinkCapability(cs *pci.ConfigSpace, b *board.Board, om *overlay.Map) {
	caps := pci.ParseCapabilities(cs)
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

		// LinkCtl2 (cap+0x30) — target speed
		if cap.Offset+0x30+2 <= pci.ConfigSpaceLegacySize {
			lc2 := cs.ReadU16(cap.Offset + 0x30)
			newLC2 := lc2
			speed := uint8(newLC2 & 0x0F)
			if speed == 0 || speed > maxSpeed {
				newLC2 = (newLC2 & 0xFFF0) | uint16(maxSpeed)
			}
			om.WriteU16(cap.Offset+0x30, newLC2, "clamp Link Control 2 target speed")
		}

		// LinkCap2 (cap+0x2C) — strip unsupported speeds from the donor vector
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

// clampDeviceCapability forces MPS=128B, disables phantoms and ext tags.
func clampDeviceCapability(cs *pci.ConfigSpace, om *overlay.Map) {
	caps := pci.ParseCapabilities(cs)
	for _, cap := range caps {
		if cap.ID != pci.CapIDPCIExpress {
			continue
		}

		// DevCap (cap+0x04)
		if cap.Offset+0x04+4 <= pci.ConfigSpaceLegacySize {
			devCap := cs.ReadU32(cap.Offset + 0x04)
			newDevCap := devCap
			newDevCap &= ^uint32(0x07) // MPS → 128B
			newDevCap &= ^uint32(0x18) // phantom functions off
			newDevCap &= ^uint32(0x20) // extended tag off
			om.WriteU32(cap.Offset+0x04, newDevCap, "clamp Device Capabilities (MPS/phantom/exttag)")
		}

		// DevCtl (cap+0x08)
		if cap.Offset+0x08+2 <= pci.ConfigSpaceLegacySize {
			devCtl := cs.ReadU16(cap.Offset + 0x08)
			newDevCtl := devCtl
			newDevCtl &= ^uint16(0x00E0) // MPS → 128B
			newDevCtl &= ^uint16(0x0100) // ext tag off
			newDevCtl &= ^uint16(0x0200) // phantom off
			if mrrs := (newDevCtl >> 12) & 0x07; mrrs > 2 {
				newDevCtl = (newDevCtl & 0x8FFF) | (2 << 12)
			}
			om.WriteU16(cap.Offset+0x08, newDevCtl, "clamp Device Control (MPS/MRRS/phantom/exttag)")
		}

		// DevCap2 (cap+0x24)
		if cap.Offset+0x24+4 <= pci.ConfigSpaceLegacySize {
			devCap2 := cs.ReadU32(cap.Offset + 0x24)
			newDevCap2 := devCap2
			newDevCap2 &= ^uint32(1 << 16) // 10-bit tag completer off
			newDevCap2 &= ^uint32(1 << 17) // 10-bit tag requester off
			om.WriteU32(cap.Offset+0x24, newDevCap2, "clamp Device Capabilities 2 (10-bit tags)")
		}

		break
	}
}

// extCapEntry represents one entry in the extended capability linked list.
type extCapEntry struct {
	offset     int
	id         uint16
	version    uint8
	nextOffset int
	size       int
}

// parseExtCapChain walks the ext cap linked list and returns all entries.
func parseExtCapChain(cs *pci.ConfigSpace) []extCapEntry {
	var entries []extCapEntry
	visited := make(map[int]bool)
	offset := 0x100

	for offset >= 0x100 && offset < pci.ConfigSpaceSize && !visited[offset] {
		visited[offset] = true

		header := cs.ReadU32(offset)
		if header == 0 || header == 0xFFFFFFFF {
			break
		}

		capID := uint16(header & 0xFFFF)
		capVer := uint8((header >> 16) & 0xF)
		nextOff := int((header >> 20) & 0xFFC)

		size := 4
		if nextOff > offset {
			size = nextOff - offset
		} else if nextOff == 0 {
			size = pci.ConfigSpaceSize - offset
		}

		entries = append(entries, extCapEntry{
			offset:     offset,
			id:         capID,
			version:    capVer,
			nextOffset: nextOff,
			size:       size,
		})

		if nextOff == 0 {
			break
		}
		offset = nextOff
	}
	return entries
}

// relinkExtCapChain patches the next-pointers of surviving entries.
func relinkExtCapChain(cs *pci.ConfigSpace, entries []extCapEntry, removeSet map[int]bool) {
	firstSurvivor := -1
	for i := range entries {
		if !removeSet[i] {
			firstSurvivor = i
			break
		}
	}

	// wipe removed regions
	for i := range entries {
		if !removeSet[i] {
			continue
		}
		e := entries[i]
		for b := 0; b < e.size && e.offset+b < pci.ConfigSpaceSize; b++ {
			cs.WriteU8(e.offset+b, 0x00)
		}
	}

	if firstSurvivor < 0 {
		cs.WriteU32(0x100, 0x00000000) // all gone
		return
	}

	// if first entry was removed, relocate first survivor to 0x100
	if removeSet[0] {
		surv := entries[firstSurvivor]

		for b := 0; b < surv.size && b < surv.offset && surv.offset+b < pci.ConfigSpaceSize; b++ {
			cs.WriteU8(0x100+b, cs.ReadU8(surv.offset+b))
		}
		for b := 0; b < surv.size && surv.offset+b < pci.ConfigSpaceSize; b++ {
			cs.WriteU8(surv.offset+b, 0x00)
		}

		newNext := 0
		for j := firstSurvivor + 1; j < len(entries); j++ {
			if !removeSet[j] {
				newNext = entries[j].offset
				break
			}
		}

		hdr := cs.ReadU32(0x100)
		hdr = (hdr & 0x000FFFFF) | (uint32(newNext) << 20)
		cs.WriteU32(0x100, hdr)

		entries[firstSurvivor].offset = 0x100
	}

	// collect survivors and relink next pointers
	var survivors []int
	for i := range entries {
		if !removeSet[i] {
			survivors = append(survivors, i)
		}
	}

	for si := 0; si < len(survivors); si++ {
		idx := survivors[si]
		e := entries[idx]

		newNext := 0
		if si+1 < len(survivors) {
			newNext = entries[survivors[si+1]].offset
		}

		hdr := cs.ReadU32(e.offset)
		hdr = (hdr & 0x000FFFFF) | (uint32(newNext) << 20)
		cs.WriteU32(e.offset, hdr)
	}
}

// FilterExtCapabilities strips unsupported ext caps and relinks the chain.
func FilterExtCapabilities(cs *pci.ConfigSpace) []string {
	entries := parseExtCapChain(cs)
	if len(entries) == 0 {
		return nil
	}

	removeSet := make(map[int]bool)
	var removed []string
	for i, e := range entries {
		if IsUnsafeExtCap(e.id) {
			removeSet[i] = true
			name := UnsafeExtCapName(e.id)
			removed = append(removed, fmt.Sprintf("%s (0x%04x) at offset 0x%03x", name, e.id, e.offset))
		}
	}

	if len(removeSet) == 0 {
		return nil
	}

	relinkExtCapChain(cs, entries, removeSet)
	return removed
}
