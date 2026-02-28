package pci

import "testing"

func TestParseBARsFromConfigSpace(t *testing.T) {
	cs := NewConfigSpace()

	// BAR0: 32-bit memory at 0xFE000000
	cs.WriteU32(0x10, 0xFE000000)

	// BAR1: IO BAR at 0x0000E001
	cs.WriteU32(0x14, 0x0000E001)

	// BAR2: 64-bit memory at 0x100000000
	cs.WriteU32(0x18, 0x0000000C) // 64-bit, prefetchable
	cs.WriteU32(0x1C, 0x00000001) // upper 32

	// BAR4: disabled
	cs.WriteU32(0x20, 0x00000000)

	bars := ParseBARsFromConfigSpace(cs)

	// BAR0
	if bars[0].Type != BARTypeMem32 {
		t.Errorf("BAR0 type = %q, want mem32", bars[0].Type)
	}
	if bars[0].Address != 0xFE000000 {
		t.Errorf("BAR0 address = 0x%x, want 0xFE000000", bars[0].Address)
	}

	// BAR1
	if bars[1].Type != BARTypeIO {
		t.Errorf("BAR1 type = %q, want io", bars[1].Type)
	}
	if bars[1].Address != 0x0000E000 {
		t.Errorf("BAR1 address = 0x%x, want 0xE000", bars[1].Address)
	}

	// BAR2 should be 64-bit
	if bars[2].Type != BARTypeMem64 {
		t.Errorf("BAR2 type = %q, want mem64", bars[2].Type)
	}
	if !bars[2].Is64Bit {
		t.Error("BAR2 should be 64-bit")
	}
	if !bars[2].Prefetchable {
		t.Error("BAR2 should be prefetchable")
	}
}

func TestParseBARsFromSysfsResource(t *testing.T) {
	lines := []string{
		"0x00000000f7d00000 0x00000000f7dfffff 0x0040200", // BAR0: 1MB memory
		"0x0000000000000000 0x0000000000000000 0x0000000", // BAR1: disabled
		"0x0000000000006001 0x000000000000601f 0x0040101", // BAR2: IO, 31 bytes
		"0x0000000000000000 0x0000000000000000 0x0000000", // BAR3: disabled
		"0x00000000f7c00000 0x00000000f7c3ffff 0x004020c", // BAR4: mem64, prefetch
		"0x0000000000000000 0x0000000000000000 0x0000000", // BAR5: disabled
	}

	bars := ParseBARsFromSysfsResource(lines)

	if len(bars) != 6 {
		t.Fatalf("Expected 6 BARs, got %d", len(bars))
	}

	// BAR0: 1MB memory
	if bars[0].Type != BARTypeMem32 {
		t.Errorf("BAR0 type = %q, want mem32", bars[0].Type)
	}
	if bars[0].Size != 0x100000 {
		t.Errorf("BAR0 size = 0x%x, want 0x100000", bars[0].Size)
	}

	// BAR1: disabled
	if !bars[1].IsDisabled() {
		t.Error("BAR1 should be disabled")
	}

	// BAR2: IO
	if bars[2].Type != BARTypeIO {
		t.Errorf("BAR2 type = %q, want io", bars[2].Type)
	}

	// BAR4: 64-bit prefetchable
	if bars[4].Type != BARTypeMem64 {
		t.Errorf("BAR4 type = %q, want mem64", bars[4].Type)
	}
	if !bars[4].Prefetchable {
		t.Error("BAR4 should be prefetchable")
	}
}

func TestBARIsIOIsMemory(t *testing.T) {
	io := BAR{Type: BARTypeIO}
	if !io.IsIO() {
		t.Error("IO BAR.IsIO() should be true")
	}
	if io.IsMemory() {
		t.Error("IO BAR.IsMemory() should be false")
	}

	mem32 := BAR{Type: BARTypeMem32}
	if !mem32.IsMemory() {
		t.Error("Mem32 BAR.IsMemory() should be true")
	}

	mem64 := BAR{Type: BARTypeMem64}
	if !mem64.IsMemory() {
		t.Error("Mem64 BAR.IsMemory() should be true")
	}
}

func TestBARSizeHuman(t *testing.T) {
	tests := []struct {
		size uint64
		want string
	}{
		{0, "0"},
		{512, "512 B"},
		{1024, "1 KB"},
		{4096, "4 KB"},
		{1048576, "1 MB"},
		{16777216, "16 MB"},
		{1073741824, "1 GB"},
	}

	for _, tt := range tests {
		b := BAR{Size: tt.size}
		got := b.SizeHuman()
		if got != tt.want {
			t.Errorf("SizeHuman(%d) = %q, want %q", tt.size, got, tt.want)
		}
	}
}

func TestBARString(t *testing.T) {
	disabled := BAR{Index: 3, Type: BARTypeDisabled}
	if disabled.String() != "BAR3: [disabled]" {
		t.Errorf("Disabled BAR string = %q", disabled.String())
	}

	mem := BAR{
		Index:        0,
		Type:         BARTypeMem32,
		Address:      0xFE000000,
		Size:         1048576,
		Prefetchable: true,
	}
	s := mem.String()
	if s != "BAR0: mem32 at 0xfe000000, size 1 MB [prefetchable]" {
		t.Errorf("Memory BAR string = %q", s)
	}
}
