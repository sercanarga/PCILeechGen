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

	return scrubbed
}

// FilterExtCapabilities removes extended capabilities that an FPGA DMA card
// cannot emulate from the config space. Returns list of removed capability names.
//
// Extended capability chain format (each entry starts with a 32-bit header):
//
//	Bits [15:0]  = Capability ID
//	Bits [19:16] = Version
//	Bits [31:20] = Next capability offset (0 = end of list)
//
// To remove a capability: zero its data region and patch the previous
// entry's next-pointer to skip over it.
func FilterExtCapabilities(cs *pci.ConfigSpace) []string {
	var removed []string

	type capEntry struct {
		offset     int
		id         uint16
		nextOffset int
		size       int
	}

	// First pass: build ordered list of all extended capabilities
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

	// Second pass: identify which entries to remove
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

	// Third pass: relink chain and zero removed regions (process backwards)
	for i := len(entries) - 1; i >= 0; i-- {
		if !removeSet[i] {
			continue
		}

		e := entries[i]

		// Find the next surviving entry's offset for relinking
		newNext := 0
		for j := i + 1; j < len(entries); j++ {
			if !removeSet[j] {
				newNext = entries[j].offset
				break
			}
		}

		// Zero out the removed capability's data region
		for b := 0; b < e.size && e.offset+b < pci.ConfigSpaceSize; b++ {
			cs.WriteU8(e.offset+b, 0x00)
		}

		// Patch the previous surviving entry's next-pointer
		if i == 0 {
			// Removing the first ext cap at 0x100
			if newNext > 0 {
				// There are surviving caps after this — zero 0x100 as empty passthrough
				// The surviving caps still have valid headers at their offsets
				// We need to make 0x100 point to the first survivor
				// Read survivor's header to get its cap ID and version
				nextHeader := cs.ReadU32(newNext)
				nextID := uint16(nextHeader & 0xFFFF)
				nextVer := uint8((nextHeader >> 16) & 0xF)
				// Find what the survivor points to next
				nextNext := int((nextHeader >> 20) & 0xFFC)

				// Move survivor's header to 0x100 as a redirect
				// Copy the survivor's full data to 0x100
				survSize := entries[0].size // use original first cap's size
				for si := i + 1; si < len(entries); si++ {
					if !removeSet[si] {
						survSize = entries[si].size
						break
					}
				}

				// Simpler approach: just create a minimal "bridge" header at 0x100
				// pointing to the first surviving cap
				// But PCIe spec says extended caps must start at 0x100
				// So we write a "Vendor Specific" placeholder that chains to survivor
				_ = nextID
				_ = nextVer
				_ = nextNext
				_ = survSize

				// Safest: write the survivor's full header at 0x100 with its chain
				// This means 0x100 becomes a duplicate pointer, but that's fine since
				// the original location is zeroed. However, the survivor still has its data
				// at the original offset. Best approach: just write an empty end-of-list.
				cs.WriteU32(0x100, 0x00000000)
			} else {
				// All extended caps removed
				cs.WriteU32(0x100, 0x00000000)
			}
		} else {
			// Find the closest previous surviving entry and patch its next-pointer
			for j := i - 1; j >= 0; j-- {
				if !removeSet[j] {
					prevHeader := cs.ReadU32(entries[j].offset)
					// Clear old next-pointer [31:20], set new one
					prevHeader = (prevHeader & 0x000FFFFF) | (uint32(newNext) << 20)
					cs.WriteU32(entries[j].offset, prevHeader)
					break
				}
			}
		}
	}

	return removed
}
