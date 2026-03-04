package scrub

import (
	"testing"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

func TestScrubConfigSpaceWithOverlay_TracksMods(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.WriteU8(0x0F, 0xFF)    // BIST non-zero
	cs.WriteU8(0x3C, 0x0A)    // Interrupt Line non-zero
	cs.WriteU16(0x04, 0xFFFF) // Command all bits set

	scrubbed, om := ScrubConfigSpaceWithOverlay(cs, nil)

	if om.Count() == 0 {
		t.Error("overlay should record modifications")
	}

	// BIST should be zeroed
	if scrubbed.ReadU8(0x0F) != 0x00 {
		t.Error("BIST should be zeroed after scrub")
	}

	// Interrupt Line should be zeroed
	if scrubbed.ReadU8(0x3C) != 0x00 {
		t.Error("Interrupt Line should be zeroed after scrub")
	}
}

func TestScrubConfigSpace_IsSameAsOverlayVersion(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	cs.WriteU16(0x00, 0x8086) // VendorID
	cs.WriteU16(0x02, 0x1234) // DeviceID
	cs.WriteU8(0x0F, 0xFF)
	cs.WriteU16(0x04, 0xFFFF)

	plain := ScrubConfigSpace(cs, nil)
	overlay, _ := ScrubConfigSpaceWithOverlay(cs, nil)

	// Both should produce identical results
	for i := 0; i < cs.Size; i++ {
		if plain.ReadU8(i) != overlay.ReadU8(i) {
			t.Errorf("mismatch at offset 0x%02X: plain=0x%02X overlay=0x%02X",
				i, plain.ReadU8(i), overlay.ReadU8(i))
		}
	}
}

func TestScrubConfigSpaceWithOverlay_BoardAwareLinkSpeed(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	// Status register: set capabilities list bit (bit 4)
	cs.WriteU16(0x06, 0x0010)
	// Capability pointer
	cs.WriteU8(0x34, 0x40)
	// PCIe Capability at offset 0x40
	cs.WriteU8(0x40, byte(pci.CapIDPCIExpress)) // PCIe cap ID
	cs.WriteU8(0x41, 0x00)                      // next cap = 0
	// Link Capabilities at cap+0x0C = 0x4C: Gen3, x4
	cs.WriteU32(0x4C, 0x00000043)

	b := &board.Board{
		MaxLinkSpeed: 2, // Limit to Gen2
	}

	scrubbed, _ := ScrubConfigSpaceWithOverlay(cs, b)

	linkCap := scrubbed.ReadU32(0x4C)
	speed := linkCap & 0x0F
	if speed > 2 {
		t.Errorf("link speed should be clamped to Gen2, got %d", speed)
	}
}
