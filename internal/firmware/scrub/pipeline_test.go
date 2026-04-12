package scrub

import (
	"testing"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/firmware/overlay"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

func TestDefaultPipeline_HasAllPasses(t *testing.T) {
	pipeline := defaultPipeline()
	if len(pipeline) != 16 {
		t.Errorf("expected 16 passes, got %d", len(pipeline))
	}

	names := make(map[string]bool)
	for _, p := range pipeline {
		name := p.Name()
		if name == "" {
			t.Error("pass with empty name")
		}
		if names[name] {
			t.Errorf("duplicate pass name: %s", name)
		}
		names[name] = true
	}
}

func TestDefaultPipeline_NamesAreUnique(t *testing.T) {
	pipeline := defaultPipeline()
	seen := make(map[string]bool)
	for _, p := range pipeline {
		if seen[p.Name()] {
			t.Errorf("duplicate pass name: %s", p.Name())
		}
		seen[p.Name()] = true
	}
}

func TestClearMiscPass(t *testing.T) {
	cs := pci.NewConfigSpaceFromBytes(make([]byte, 256))
	cs.Data[0x0F] = 0xFF // BIST
	cs.Data[0x3C] = 0x0A // Interrupt Line
	cs.Data[0x0D] = 0x40 // Latency Timer
	cs.Data[0x0C] = 0x10 // Cache Line Size

	om := overlay.NewMap(cs)
	pass := &clearMiscPass{}
	pass.Apply(cs, nil, om, ctxFor(cs))

	if cs.Data[0x0F] != 0 {
		t.Error("BIST should be cleared")
	}
	if cs.Data[0x3C] != 0 {
		t.Error("Interrupt Line should be cleared")
	}
	if cs.Data[0x0D] != 0 {
		t.Error("Latency Timer should be cleared")
	}
	if cs.Data[0x0C] != 0 {
		t.Error("Cache Line Size should be cleared")
	}
}

func TestSanitizeCmdStatusPass(t *testing.T) {
	data := make([]byte, 256)
	// Command = 0xFFFF (all bits set)
	data[0x04] = 0xFF
	data[0x05] = 0xFF
	// Status = 0xFFFF
	data[0x06] = 0xFF
	data[0x07] = 0xFF

	cs := pci.NewConfigSpaceFromBytes(data)
	om := overlay.NewMap(cs)
	pass := &sanitizeCmdStatusPass{}
	pass.Apply(cs, nil, om, ctxFor(cs))

	cmd := cs.Command()
	if cmd != (0xFFFF & cmdMask) {
		t.Errorf("Command should be masked, got 0x%04X", cmd)
	}
	status := cs.Status()
	if status != (0xFFFF & statusMask) {
		t.Errorf("Status should be masked, got 0x%04X", status)
	}
}

func TestFilterExtCapsPass_SmallCS(t *testing.T) {
	// config space smaller than 4K should be a no-op
	cs := pci.NewConfigSpaceFromBytes(make([]byte, 256))
	om := overlay.NewMap(cs)
	pass := &filterExtCapsPass{}
	pass.Apply(cs, nil, om, ctxFor(cs)) // should not panic
}

func TestClampBARsPass(t *testing.T) {
	data := make([]byte, 256)
	// BAR0 = memory BAR with large size (mask=0xFFF00000 -> 1MB)
	data[0x10] = 0x00
	data[0x11] = 0x00
	data[0x12] = 0x10 // some address
	data[0x13] = 0x80

	cs := pci.NewConfigSpaceFromBytes(data)
	om := overlay.NewMap(cs)
	pass := &clampBARsPass{}
	pass.Apply(cs, nil, om, ctxFor(cs))
	// should not panic; BAR clamping applies 4KB mask
}

func TestClampLinkPass_NilBoard(t *testing.T) {
	data := make([]byte, 256)
	cs := pci.NewConfigSpaceFromBytes(data)
	om := overlay.NewMap(cs)
	pass := &clampLinkPass{}
	// nil board uses defaults
	b := &board.Board{}
	pass.Apply(cs, b, om, ctxFor(cs)) // should not panic
}

func TestValidateCapChainPass_EmptyCS(t *testing.T) {
	cs := pci.NewConfigSpaceFromBytes(make([]byte, 256))
	om := overlay.NewMap(cs)
	pass := &validateCapChainPass{}
	pass.Apply(cs, nil, om, ctxFor(cs)) // should not panic
}
