package firmware

import (
	"fmt"

	"github.com/sercanarga/pcileechgen/internal/pci"
)

// unsafeExtCaps lists extended capability IDs that FPGA DMA cards cannot
// emulate. These are removed from config space to prevent the host OS from
// interacting with features that don't actually exist on the hardware.
var unsafeExtCaps = map[uint16]string{
	pci.ExtCapIDSRIOV:         "SR-IOV",          // Virtual functions — FPGA has no VF support
	pci.ExtCapIDMRIOV:         "MR-IOV",          // Multi-root IOV — not applicable
	pci.ExtCapIDResizableBAR:  "Resizable BAR",   // FPGA has fixed BRAM size
	pci.ExtCapIDATS:           "ATS",             // Address Translation Services — requires IOMMU interaction
	pci.ExtCapIDPageRequest:   "Page Request",    // Requires ATS
	pci.ExtCapIDPASID:         "PASID",           // Process Address Space ID — requires ATS
	pci.ExtCapIDL1PMSubstates: "L1 PM Substates", // FPGA doesn't implement ASPM L1 substates
	pci.ExtCapIDDPC:           "DPC",             // Downstream Port Containment — endpoint doesn't have this
	pci.ExtCapIDPTM:           "PTM",             // Precision Time Measurement — FPGA has no PTM clock
	pci.ExtCapIDSecondaryPCIe: "Secondary PCIe",  // For bridges only
	pci.ExtCapIDMulticast:     "Multicast",       // Requires switch/bridge support
}

// IsUnsafeExtCap returns true if the given extended capability ID cannot be
// emulated by an FPGA DMA card and should be filtered from config space.
func IsUnsafeExtCap(id uint16) bool {
	_, ok := unsafeExtCaps[id]
	return ok
}

// UnsafeExtCapName returns the human-readable name for an unsafe capability.
func UnsafeExtCapName(id uint16) string {
	if name, ok := unsafeExtCaps[id]; ok {
		return name
	}
	return ""
}

// ScrubConfigSpace cleans potentially dangerous or detection-revealing registers
// from the config space before writing to COE. This prevents leaking donor-specific
// debug/diagnostic data that the DMA card cannot actually implement.
func ScrubConfigSpace(cs *pci.ConfigSpace) *pci.ConfigSpace {
	scrubbed := cs.Clone()

	// Clear BIST register (offset 0x0F) — DMA card cannot run self-test
	scrubbed.WriteU8(0x0F, 0x00)

	// Clear Interrupt Line (offset 0x3C) — will be assigned by host at runtime
	scrubbed.WriteU8(0x3C, 0x00)

	// Clear Latency Timer (offset 0x0D) — not meaningful for PCIe devices
	scrubbed.WriteU8(0x0D, 0x00)

	// Clear Cache Line Size (offset 0x0C) — set by OS
	scrubbed.WriteU8(0x0C, 0x00)

	// Sanitize Command register (offset 0x04) — reset to sane defaults
	// Keep only: IO Space(0), Memory Space(1), Bus Master(2), Parity Error Response(6)
	cmd := scrubbed.Command() & 0x0547
	scrubbed.WriteU16(0x04, cmd)

	// Clear Status register write-1-to-clear bits (offset 0x06)
	// Keep capability list bit (4) and speed bits, clear error bits
	status := scrubbed.Status() & 0x06F0
	scrubbed.WriteU16(0x06, status)

	// Sanitize PCIe Device Status (clear all error/transaction bits)
	caps := pci.ParseCapabilities(scrubbed)
	for _, cap := range caps {
		if cap.ID == pci.CapIDPCIExpress && cap.Offset+10 < pci.ConfigSpaceLegacySize {
			// Device Status at cap+10: clear all RW1C bits
			scrubbed.WriteU16(cap.Offset+10, 0x0000)

			// Link Status at cap+18: clear training bits
			if cap.Offset+18 < pci.ConfigSpaceLegacySize {
				lstatus := scrubbed.ReadU16(cap.Offset + 18)
				lstatus &= 0x3FFF
				scrubbed.WriteU16(cap.Offset+18, lstatus)
			}
		}

		if cap.ID == pci.CapIDPowerManagement && cap.Offset+4 < pci.ConfigSpaceLegacySize {
			// PM Control/Status: set to D0, preserve NoSoftReset
			pmcsr := scrubbed.ReadU16(cap.Offset + 4)
			pmcsr &= 0xFFFC // Clear PowerState bits (set to D0)
			pmcsr &= 0x7FFF // Clear PME_Status
			pmcsr |= 0x0008 // Set NoSoftReset — FPGA preserves state across D3hot→D0
			scrubbed.WriteU16(cap.Offset+4, pmcsr)
		}
	}

	// Scrub extended config space
	if scrubbed.Size >= pci.ConfigSpaceSize {
		// Clean AER error status registers
		extCaps := pci.ParseExtCapabilities(scrubbed)
		for _, cap := range extCaps {
			if cap.ID == pci.ExtCapIDAER {
				if cap.Offset+4+4 <= pci.ConfigSpaceSize {
					scrubbed.WriteU32(cap.Offset+4, 0x00000000) // Uncorrectable Error Status
				}
				if cap.Offset+16+4 <= pci.ConfigSpaceSize {
					scrubbed.WriteU32(cap.Offset+16, 0x00000000) // Correctable Error Status
				}
				if cap.Offset+28+4 <= pci.ConfigSpaceSize {
					scrubbed.WriteU32(cap.Offset+28, 0x00000000) // Root Error Status
				}
			}
		}

		// Filter out capabilities that FPGA cannot emulate
		FilterExtCapabilities(scrubbed)
	}

	// Clamp BAR sizes to FPGA BRAM limit (4 KB)
	clampBARsToFPGA(scrubbed)

	return scrubbed
}

const fpgaBRAMSize = 4096           // pcileech-fpga shadow BAR BRAM size
const fpgaBAR0SizeMask = 0xFFFFF000 // 4 KB aligned BAR mask

// clampBARsToFPGA rewrites memory BAR registers to advertise 4 KB max size,
// matching the actual FPGA BRAM capacity.
func clampBARsToFPGA(cs *pci.ConfigSpace) {
	for i := 0; i < 6; i++ {
		barOffset := 0x10 + (i * 4)
		barVal := cs.BAR(i)
		if barVal == 0 {
			continue
		}

		isIO := barVal&0x01 != 0
		if isIO {
			continue // skip IO BARs
		}

		is64bit := (barVal & 0x06) == 0x04
		// preserve type bits [3:0], apply 4 KB mask
		newBar := fpgaBAR0SizeMask | (barVal & 0x0F)
		cs.WriteU32(barOffset, newBar)

		if is64bit && i < 5 {
			cs.WriteU32(barOffset+4, 0x00000000) // clear upper 32 bits
			i++                                  // skip upper half
		}
	}
}

// FilterExtCapabilities strips extended capabilities the FPGA can't emulate.
// Zeroes removed regions, relinks the chain, and relocates to 0x100 if needed
// (PCIe spec requires ext caps to start there).
func FilterExtCapabilities(cs *pci.ConfigSpace) []string {
	var removed []string

	type capEntry struct {
		offset     int
		id         uint16
		version    uint8
		nextOffset int
		size       int
	}

	// walk the chain and collect all entries
	var entries []capEntry
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

		entries = append(entries, capEntry{
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

	if len(entries) == 0 {
		return nil
	}

	// mark unsafe entries for removal
	removeSet := make(map[int]bool)
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

	// find first survivor (-1 if all removed)
	firstSurvivor := -1
	for i := range entries {
		if !removeSet[i] {
			firstSurvivor = i
			break
		}
	}

	// if 0x100 is being removed, we need to relocate a survivor there
	needsRelocate := removeSet[0] && firstSurvivor > 0

	// zero out removed cap regions
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
		// all caps gone
		cs.WriteU32(0x100, 0x00000000)
		return removed
	}

	if needsRelocate {
		// move first survivor to 0x100
		surv := entries[firstSurvivor]

		// copy data
		for b := 0; b < surv.size && b < surv.offset && surv.offset+b < pci.ConfigSpaceSize; b++ {
			cs.WriteU8(0x100+b, cs.ReadU8(surv.offset+b))
		}

		// clear original location
		for b := 0; b < surv.size && surv.offset+b < pci.ConfigSpaceSize; b++ {
			cs.WriteU8(surv.offset+b, 0x00)
		}

		// find next survivor for chain
		newNext := 0
		for j := firstSurvivor + 1; j < len(entries); j++ {
			if !removeSet[j] {
				newNext = entries[j].offset
				break
			}
		}

		// fix next-pointer at 0x100
		hdr := cs.ReadU32(0x100)
		hdr = (hdr & 0x000FFFFF) | (uint32(newNext) << 20)
		cs.WriteU32(0x100, hdr)

		// update offset so relinking below uses the new location
		entries[firstSurvivor].offset = 0x100
	}

	// relink surviving chain
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

	return removed
}
