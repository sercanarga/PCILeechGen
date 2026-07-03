package scrub

import (
	"fmt"
	"log/slog"

	"github.com/sercanarga/pcileechgen/internal/firmware/overlay"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

// standard caps the FPGA can't emulate - pruned from the linked list
var unsafeStandardCaps = map[uint8]string{
	pci.CapIDVPD:               "VPD",
	pci.CapIDAGP:               "AGP",
	pci.CapIDAGP8x:             "AGP 8x",
	pci.CapIDSlotID:            "Slot ID",
	pci.CapIDCompactPCIHotSwap: "CompactPCI HotSwap",
	pci.CapIDCompactPCI:        "CompactPCI",
	pci.CapIDPCIHotPlug:        "PCI Hot-Plug",
	pci.CapIDHyperTransport:    "HyperTransport",
	pci.CapIDDebugPort:         "Debug Port",
	pci.CapIDSATADataIndex:     "SATA Data/Index",
	pci.CapIDEnhancedAlloc:     "Enhanced Allocation",
	pci.CapIDFlatteningPortal:  "Flattening Portal",
	pci.CapIDPCIX:              "PCI-X",
	pci.CapIDBridgeSubsysVID:   "Bridge Subsystem VID",
	pci.CapIDSecureDevice:      "Secure Device",
	pci.CapIDAdvancedFeatures:  "Advanced Features",
}

// PruneStandardCaps unlinks unsupported caps and returns what was removed.
func PruneStandardCaps(cs *pci.ConfigSpace, om *overlay.Map) []string {
	if !cs.HasCapabilities() {
		return nil
	}

	var removed []string

	prevNextOff := 0x34
	ptr := int(cs.ReadU8(0x34)) & 0xFC

	visited := make(map[int]bool)
	for ptr != 0 && ptr < pci.ConfigSpaceLegacySize && !visited[ptr] {
		visited[ptr] = true
		capID := cs.ReadU8(ptr)
		nextPtr := int(cs.ReadU8(ptr+1)) & 0xFC

		if _, bad := unsafeStandardCaps[capID]; bad {
			name := unsafeStandardCaps[capID]
			slog.Info("pruning standard cap", "name", name, "id", fmt.Sprintf("0x%02X", capID), "offset", fmt.Sprintf("0x%02X", ptr))

			om.WriteU8(prevNextOff, uint8(nextPtr),
				fmt.Sprintf("prune cap 0x%02X (%s): relink", capID, name))

			// Zero the whole pruned cap body, not just the header. The relink
			// already unlinks it from the chain, but the body bytes (VPD data,
			// HyperTransport regs, ...) otherwise linger in config space and a
			// raw scan finds the orphaned structure - a forensic tell real
			// hardware never leaves. Bound to the next cap (or a small default
			// for the last cap) so neighbours are untouched.
			end := nextPtr
			if end <= ptr {
				end = ptr + capSizeAt(cs, capID, ptr)
			}
			if end > pci.ConfigSpaceLegacySize {
				end = pci.ConfigSpaceLegacySize
			}
			for b := ptr; b < end; b++ {
				om.WriteU8(b, 0, fmt.Sprintf("zero pruned cap 0x%02X body", capID))
			}

			removed = append(removed, fmt.Sprintf("%s (0x%02X) at 0x%02X", name, capID, ptr))
		} else {
			prevNextOff = ptr + 1
		}
		ptr = nextPtr
	}

	return removed
}

// ValidateCapChain checks for loops and bad pointers in the cap list.
func ValidateCapChain(cs *pci.ConfigSpace) error {
	if !cs.HasCapabilities() {
		return nil
	}

	visited := make(map[int]bool)
	ptr := int(cs.ReadU8(0x34)) & 0xFC
	count := 0

	for ptr != 0 {
		if ptr < 0x40 || ptr >= pci.ConfigSpaceLegacySize {
			return fmt.Errorf("capability pointer 0x%02X out of range [0x40, 0xFF]", ptr)
		}
		if ptr&0x03 != 0 {
			return fmt.Errorf("capability pointer 0x%02X not DWORD-aligned", ptr)
		}
		if visited[ptr] {
			return fmt.Errorf("capability chain loop at offset 0x%02X", ptr)
		}
		visited[ptr] = true
		count++
		if count > 48 {
			return fmt.Errorf("capability chain too long (%d entries), likely corrupted", count)
		}
		ptr = int(cs.ReadU8(ptr+1)) & 0xFC
	}
	return nil
}
