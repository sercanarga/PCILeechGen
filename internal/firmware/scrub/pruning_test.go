package scrub

import (
	"testing"

	"github.com/sercanarga/pcileechgen/internal/firmware/overlay"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

func TestPruneStandardCaps_RemovesVPD(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Data[0x06] = 0x10 // Status: capabilities list present

	// Build a chain: PM(0x40) -> VPD(0x50) -> PCIe(0x60) -> 0
	cs.WriteU8(0x34, 0x40) // cap pointer

	// PM at 0x40
	cs.WriteU8(0x40, pci.CapIDPowerManagement)
	cs.WriteU8(0x41, 0x50) // next -> VPD

	// VPD at 0x50
	cs.WriteU8(0x50, pci.CapIDVPD)
	cs.WriteU8(0x51, 0x60) // next -> PCIe

	// PCIe at 0x60
	cs.WriteU8(0x60, pci.CapIDPCIExpress)
	cs.WriteU8(0x61, 0x00) // end

	om := overlay.NewMap(cs)
	removed := PruneStandardCaps(cs, om)

	if len(removed) != 1 {
		t.Fatalf("expected 1 removed cap, got %d: %v", len(removed), removed)
	}

	// PM's next should now point to PCIe (0x60), skipping VPD
	pmNext := cs.ReadU8(0x41)
	if pmNext != 0x60 {
		t.Errorf("PM next should be 0x60 (PCIe), got 0x%02X", pmNext)
	}

	// VPD at 0x50 should be zeroed
	if cs.ReadU8(0x50) != 0 {
		t.Error("VPD cap ID should be zeroed")
	}
}

// The whole pruned cap body must be zeroed, not just the header, so no
// orphaned cap structure lingers in config space for a raw scan to find.
func TestPruneStandardCaps_ZeroesBody(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Data[0x06] = 0x10
	cs.WriteU8(0x34, 0x40)
	// PM(0x40) -> VPD(0x50) -> PCIe(0x60)
	cs.WriteU8(0x40, pci.CapIDPowerManagement)
	cs.WriteU8(0x41, 0x50)
	cs.WriteU8(0x50, pci.CapIDVPD)
	cs.WriteU8(0x51, 0x60)
	cs.WriteU8(0x60, pci.CapIDPCIExpress)
	cs.WriteU8(0x61, 0x00)
	// VPD body bytes (0x52..0x5F) carry leftover donor data.
	for b := 0x52; b < 0x60; b++ {
		cs.WriteU8(b, 0xAA)
	}

	om := overlay.NewMap(cs)
	if removed := PruneStandardCaps(cs, om); len(removed) != 1 {
		t.Fatalf("expected 1 removed cap, got %d", len(removed))
	}

	for b := 0x50; b < 0x60; b++ {
		if cs.ReadU8(b) != 0 {
			t.Errorf("pruned VPD body byte 0x%02X = 0x%02X, want 0 (orphan structure must be wiped)", b, cs.ReadU8(b))
		}
	}
	// The following kept cap (PCIe at 0x60) must be untouched.
	if cs.ReadU8(0x60) != pci.CapIDPCIExpress {
		t.Errorf("neighbour cap at 0x60 clobbered: 0x%02X", cs.ReadU8(0x60))
	}
}

func TestPruneStandardCaps_NothingToRemove(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Data[0x06] = 0x10

	cs.WriteU8(0x34, 0x40)
	cs.WriteU8(0x40, pci.CapIDPCIExpress)
	cs.WriteU8(0x41, 0x00)

	om := overlay.NewMap(cs)
	removed := PruneStandardCaps(cs, om)

	if len(removed) != 0 {
		t.Errorf("expected 0 removed, got %d", len(removed))
	}
}

func TestValidateCapChain_Valid(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Data[0x06] = 0x10

	cs.WriteU8(0x34, 0x40)
	cs.WriteU8(0x40, pci.CapIDPowerManagement)
	cs.WriteU8(0x41, 0x50)
	cs.WriteU8(0x50, pci.CapIDPCIExpress)
	cs.WriteU8(0x51, 0x00)

	if err := ValidateCapChain(cs); err != nil {
		t.Errorf("expected valid chain, got: %v", err)
	}
}

func TestValidateCapChain_Loop(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Data[0x06] = 0x10

	cs.WriteU8(0x34, 0x40)
	cs.WriteU8(0x40, pci.CapIDPowerManagement)
	cs.WriteU8(0x41, 0x40) // loop!

	if err := ValidateCapChain(cs); err == nil {
		t.Error("expected error for loop")
	}
}

func TestValidateCapChain_OutOfBounds(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Data[0x06] = 0x10

	cs.WriteU8(0x34, 0x04) // out of bounds (< 0x40)

	if err := ValidateCapChain(cs); err == nil {
		t.Error("expected error for out-of-bounds pointer")
	}
}
