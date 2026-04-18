package scrub

import (
	"testing"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/firmware/overlay"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

func ctxFor(cs *pci.ConfigSpace) *ScrubContext {
	return &ScrubContext{
		Caps:    pci.ParseCapabilities(cs),
		ExtCaps: pci.ParseExtCapabilities(cs),
	}
}

// --- Pass tests ---

func TestClearMiscPass_AllFields(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceLegacySize
	cs.WriteU8(0x0F, 0xFF) // BIST
	cs.WriteU8(0x3C, 0x0A) // Interrupt Line (preserved, not cleared)
	cs.WriteU8(0x0D, 0x40) // Latency Timer
	cs.WriteU8(0x0C, 0x10) // Cache Line Size

	om := overlay.NewMap(cs)
	p := &clearMiscPass{}
	p.Apply(cs, nil, om, ctxFor(cs))

	if cs.ReadU8(0x0F) != 0 {
		t.Error("BIST should be cleared")
	}
	// Interrupt Line is intentionally NOT cleared (see passes.go comment)
	if cs.ReadU8(0x3C) != 0x0A {
		t.Errorf("Interrupt Line should be preserved, got 0x%02X want 0x0A", cs.ReadU8(0x3C))
	}
	if cs.ReadU8(0x0D) != 0 {
		t.Error("Latency Timer should be cleared")
	}
	if cs.ReadU8(0x0C) != 0 {
		t.Error("Cache Line Size should be cleared")
	}
}

func TestSanitizeCmdStatusPass_Force(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceLegacySize
	cs.WriteU16(0x04, 0xFFFF) // Command with all bits
	cs.WriteU16(0x06, 0xFFFF) // Status with all bits

	om := overlay.NewMap(cs)
	p := &sanitizeCmdStatusPass{}
	p.Apply(cs, nil, om, ctxFor(cs))

	cmd := cs.ReadU16(0x04)
	if cmd&0x06 == 0 {
		t.Error("BME and MSE should be forced on")
	}
}

func TestScrubPCIeCapPass(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceLegacySize
	cs.WriteU16(0x06, 0x0010) // has caps
	cs.WriteU8(0x34, 0x40)
	cs.WriteU8(0x40, pci.CapIDPCIExpress)
	cs.WriteU8(0x41, 0x00)
	cs.WriteU16(0x4A, 0xFFFF) // Device Status - dirty
	cs.WriteU16(0x52, 0xFFFF) // Link Status - dirty

	om := overlay.NewMap(cs)
	p := &scrubPCIeCapPass{}
	p.Apply(cs, nil, om, ctxFor(cs))

	if cs.ReadU16(0x4A) != 0 {
		t.Errorf("Device Status = 0x%04x, want 0", cs.ReadU16(0x4A))
	}
	ls := cs.ReadU16(0x52)
	if ls&0xC000 != 0 {
		t.Error("Link Status RW1C bits should be cleared")
	}
}

func TestScrubPMCapPass(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceLegacySize
	cs.WriteU16(0x06, 0x0010) // has caps
	cs.WriteU8(0x34, 0x40)
	cs.WriteU8(0x40, pci.CapIDPowerManagement)
	cs.WriteU8(0x41, 0x00)
	// PMC with PME from D0, D3hot, D3cold
	cs.WriteU16(0x42, 0xC803)
	// PMCSR: D3, PME_Enable=1, PME_Status=1
	cs.WriteU16(0x44, 0x9103)

	om := overlay.NewMap(cs)
	p := &scrubPMCapPass{}
	p.Apply(cs, nil, om, ctxFor(cs))

	// PMC: PME_Support bits [15:11] must be cleared
	pmc := cs.ReadU16(0x42)
	if pmc&0xF800 != 0 {
		t.Errorf("PMC PME_Support = 0x%04X, should be 0 (prevent D3 transitions)", pmc&0xF800)
	}
	// PMC version should be preserved
	if pmc&0x07 != 3 {
		t.Errorf("PMC version = %d, should be 3", pmc&0x07)
	}

	// PMCSR checks
	pmcsr := cs.ReadU16(0x44)
	if pmcsr&0x03 != 0 {
		t.Error("Power state should be D0 (bits 1:0 = 0)")
	}
	if pmcsr&0x8000 != 0 {
		t.Error("PME_Status should be cleared")
	}
	if pmcsr&0x08 == 0 {
		t.Error("NoSoftReset should be set")
	}
	if pmcsr&0x0100 != 0 {
		t.Error("PME_Enable should be cleared (prevent PME wake)")
	}
}

// verify D3 prevention: scrub should remove all PME hints so Windows
// doesn't aggressively transition the device to D3hot after idle.
func TestScrubPMCapPass_PreventsD3(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceLegacySize
	cs.WriteU16(0x06, 0x0010)
	cs.WriteU8(0x34, 0x40)
	cs.WriteU8(0x40, pci.CapIDPowerManagement)
	cs.WriteU8(0x41, 0x00)
	// donor advertises PME from all D-states
	cs.WriteU16(0x42, 0xF803) // PME from D0, D1, D2, D3hot, D3cold
	cs.WriteU16(0x44, 0x0103) // D3, PME_Enable=1

	om := overlay.NewMap(cs)
	p := &scrubPMCapPass{}
	p.Apply(cs, nil, om, ctxFor(cs))

	pmc := cs.ReadU16(0x42)
	// no PME support from any state
	for bit := 11; bit <= 15; bit++ {
		if pmc&(1<<bit) != 0 {
			t.Errorf("PMC bit %d (PME_Support) should be cleared, PMC=0x%04X", bit, pmc)
		}
	}

	pmcsr := cs.ReadU16(0x44)
	// D0 + PME disabled
	if pmcsr&0x03 != 0 {
		t.Errorf("should be D0, got D%d", pmcsr&0x03)
	}
	if pmcsr&0x0100 != 0 {
		t.Error("PME_Enable should be cleared")
	}
}

func TestScrubAERPass(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	// AER at 0x100
	aerHeader := uint32(pci.ExtCapIDAER) | (1 << 16)
	cs.WriteU32(0x100, aerHeader)
	cs.WriteU32(0x104, 0xFFFFFFFF) // uncorrectable status
	cs.WriteU32(0x110, 0xFFFFFFFF) // correctable status
	cs.WriteU32(0x11C, 0xFFFFFFFF) // root error status

	om := overlay.NewMap(cs)
	p := &scrubAERPass{}
	p.Apply(cs, nil, om, ctxFor(cs))

	if cs.ReadU32(0x104) != 0 {
		t.Error("AER uncorrectable status should be cleared")
	}
	if cs.ReadU32(0x110) != 0 {
		t.Error("AER correctable status should be cleared")
	}
	if cs.ReadU32(0x11C) != 0 {
		t.Error("AER root error status should be cleared")
	}
}

func TestScrubAERPass_SmallCS(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceLegacySize // no ext caps
	om := overlay.NewMap(cs)
	p := &scrubAERPass{}
	p.Apply(cs, nil, om, ctxFor(cs)) // should be no-op
}

func TestFilterExtCapsPass(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	// Write a safe cap
	cs.WriteU32(0x100, uint32(pci.ExtCapIDAER)|(1<<16))
	om := overlay.NewMap(cs)
	p := &filterExtCapsPass{}
	p.Apply(cs, nil, om, ctxFor(cs)) // should not remove AER
}

func TestClampBARsPass_Memory(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	cs.WriteU32(0x10, 0xFFF00004)
	om := overlay.NewMap(cs)
	p := &clampBARsPass{}
	p.Apply(cs, nil, om, ctxFor(cs))
}

func TestRelocateMSIXPass(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	om := overlay.NewMap(cs)
	p := &relocateMSIXPass{}
	p.Apply(cs, nil, om, ctxFor(cs)) // no MSI-X - should be no-op
}

func TestClampLinkPass(t *testing.T) {
	cs := makeTestCS()
	b := &board.Board{PCIeLanes: 1, MaxLinkSpeed: 2}
	om := overlay.NewMap(cs)
	p := &clampLinkPass{}
	p.Apply(cs, b, om, ctxFor(cs))
}

func TestClampDeviceCapPass(t *testing.T) {
	cs := makeTestCS()
	om := overlay.NewMap(cs)
	p := &clampDeviceCapPass{}
	p.Apply(cs, nil, om, ctxFor(cs))
}

func TestZeroVendorPass(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	cs.WriteU16(0x06, 0x0010)
	cs.WriteU8(0x34, 0x40)
	cs.WriteU8(0x40, pci.CapIDPCIExpress)
	cs.WriteU8(0x41, 0x00)
	om := overlay.NewMap(cs)
	p := &zeroVendorPass{}
	p.Apply(cs, nil, om, ctxFor(cs))
}

func TestApplyVendorQuirksPass(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	om := overlay.NewMap(cs)
	p := &applyVendorQuirksPass{}
	p.Apply(cs, nil, om, ctxFor(cs))
}

func TestPruneStdCapsPass(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	om := overlay.NewMap(cs)
	p := &pruneStdCapsPass{}
	p.Apply(cs, nil, om, ctxFor(cs))
}

func TestValidateCapChainPass(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	om := overlay.NewMap(cs)
	p := &validateCapChainPass{}
	p.Apply(cs, nil, om, ctxFor(cs))
}

// --- Pass name tests ---

func TestPassNames(t *testing.T) {
	passes := []struct {
		name string
		pass interface{ Name() string }
	}{
		{"clear misc registers", &clearMiscPass{}},
		{"sanitize Command/Status", &sanitizeCmdStatusPass{}},
		{"scrub PCIe capability", &scrubPCIeCapPass{}},
		{"scrub PM capability", &scrubPMCapPass{}},
		{"scrub AER status", &scrubAERPass{}},
		{"filter ext capabilities", &filterExtCapsPass{}},
		{"clamp BARs to FPGA", &clampBARsPass{}},
		{"relocate MSI-X to BRAM", &relocateMSIXPass{}},
		{"clamp link capability", &clampLinkPass{}},
		{"clamp device capability", &clampDeviceCapPass{}},
		{"zero vendor registers", &zeroVendorPass{}},
		{"vendor quirks", &applyVendorQuirksPass{}},
		{"prune standard caps", &pruneStdCapsPass{}},
		{"validate cap chain", &validateCapChainPass{}},
	}
	for _, tt := range passes {
		if tt.pass.Name() != tt.name {
			t.Errorf("Name() = %q, want %q", tt.pass.Name(), tt.name)
		}
	}
}
