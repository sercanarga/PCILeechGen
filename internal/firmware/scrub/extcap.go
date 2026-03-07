package scrub

import (
	"fmt"
	"log/slog"

	"github.com/sercanarga/pcileechgen/internal/pci"
)

// extCapFilter describes a dangerous ext cap and why it must be removed.
type extCapFilter struct {
	Name   string
	Reason string
}

// unsafeExtCaps lists extended capabilities requiring active FPGA hardware
// emulation that the device cannot provide. Read-only / informational caps
// (AER, DSN, LTR, SecondaryPCIe, …) are safe to keep.
//
// Caps whose control registers are neutralised by other scrub passes
// (e.g. L1PM Substates → scrubASPMPass) must NOT appear here.
var unsafeExtCaps = map[uint16]extCapFilter{
	pci.ExtCapIDSRIOV:        {"SR-IOV", "VF creation requires FPGA hardware support"},
	pci.ExtCapIDMRIOV:        {"MR-IOV", "multi-root IOV management unsupported"},
	pci.ExtCapIDResizableBAR: {"Resizable BAR", "dynamic BAR resizing unsupported"},
	pci.ExtCapIDATS:          {"ATS", "IOMMU address translation unsupported"},
	pci.ExtCapIDPageRequest:  {"Page Request", "IOMMU page fault handling unsupported"},
	pci.ExtCapIDPASID:        {"PASID", "process address space management unsupported"},
	pci.ExtCapIDDPC:          {"DPC", "downstream port containment signaling unsupported"},
	pci.ExtCapIDPTM:          {"PTM", "hardware precision timestamping unsupported"},
	pci.ExtCapIDMulticast:    {"Multicast", "TLP multicast routing unsupported"},
}

// IsUnsafeExtCap returns true if the given ext cap ID will be filtered out.
func IsUnsafeExtCap(id uint16) bool {
	_, ok := unsafeExtCaps[id]
	return ok
}

// UnsafeExtCapName returns the human name for a filtered cap, or empty string.
func UnsafeExtCapName(id uint16) string {
	if f, ok := unsafeExtCaps[id]; ok {
		return f.Name
	}
	return ""
}

// UnsafeExtCapReason returns why the cap is filtered, or empty string.
func UnsafeExtCapReason(id uint16) string {
	if f, ok := unsafeExtCaps[id]; ok {
		return f.Reason
	}
	return ""
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
			reason := UnsafeExtCapReason(e.id)
			removed = append(removed, fmt.Sprintf("%s (0x%04x) at offset 0x%03x", name, e.id, e.offset))
			slog.Info("filtering ext cap",
				"name", name, "id", fmt.Sprintf("0x%04x", e.id),
				"offset", fmt.Sprintf("0x%03x", e.offset), "reason", reason)
		}
	}

	if len(removeSet) == 0 {
		return nil
	}

	relinkExtCapChain(cs, entries, removeSet)
	return removed
}
