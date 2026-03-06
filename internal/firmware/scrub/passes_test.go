package scrub

import (
	"testing"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/firmware/overlay"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

// --- Pass tests ---

func TestClearMiscPass_AllFields(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceLegacySize
	cs.WriteU8(0x0F, 0xFF) // BIST
	cs.WriteU8(0x3C, 0x0A) // Interrupt Line
	cs.WriteU8(0x0D, 0x40) // Latency Timer
	cs.WriteU8(0x0C, 0x10) // Cache Line Size

	om := overlay.NewMap(cs)
	p := &clearMiscPass{}
	p.Apply(cs, nil, om)

	if cs.ReadU8(0x0F) != 0 {
		t.Error("BIST should be cleared")
	}
	if cs.ReadU8(0x3C) != 0 {
		t.Error("Interrupt Line should be cleared")
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
	p.Apply(cs, nil, om)

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
	cs.WriteU16(0x4A, 0xFFFF) // Device Status — dirty
	cs.WriteU16(0x52, 0xFFFF) // Link Status — dirty

	om := overlay.NewMap(cs)
	p := &scrubPCIeCapPass{}
	p.Apply(cs, nil, om)

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
	cs.WriteU16(0x44, 0x8003) // D3, PME_Status set

	om := overlay.NewMap(cs)
	p := &scrubPMCapPass{}
	p.Apply(cs, nil, om)

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
	p.Apply(cs, nil, om)

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
	p.Apply(cs, nil, om) // should be no-op
}

func TestFilterExtCapsPass(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	// Write a safe cap
	cs.WriteU32(0x100, uint32(pci.ExtCapIDAER)|(1<<16))
	om := overlay.NewMap(cs)
	p := &filterExtCapsPass{}
	p.Apply(cs, nil, om) // should not remove AER
}

func TestClampBARsPass_Memory(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	cs.WriteU32(0x10, 0xFFF00004)
	om := overlay.NewMap(cs)
	p := &clampBARsPass{}
	p.Apply(cs, nil, om)
}

func TestRelocateMSIXPass(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	om := overlay.NewMap(cs)
	p := &relocateMSIXPass{}
	p.Apply(cs, nil, om) // no MSI-X — should be no-op
}

func TestClampLinkPass(t *testing.T) {
	cs := makeTestCS()
	b := &board.Board{PCIeLanes: 1, MaxLinkSpeed: 2}
	om := overlay.NewMap(cs)
	p := &clampLinkPass{}
	p.Apply(cs, b, om)
}

func TestClampDeviceCapPass(t *testing.T) {
	cs := makeTestCS()
	om := overlay.NewMap(cs)
	p := &clampDeviceCapPass{}
	p.Apply(cs, nil, om)
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
	p.Apply(cs, nil, om)
}

func TestApplyVendorQuirksPass(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	om := overlay.NewMap(cs)
	p := &applyVendorQuirksPass{}
	p.Apply(cs, nil, om)
}

func TestPruneStdCapsPass(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	om := overlay.NewMap(cs)
	p := &pruneStdCapsPass{}
	p.Apply(cs, nil, om)
}

func TestValidateCapChainPass(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	om := overlay.NewMap(cs)
	p := &validateCapChainPass{}
	p.Apply(cs, nil, om)
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
