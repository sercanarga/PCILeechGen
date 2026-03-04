package overlay

import (
	"strings"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/pci"
)

func testCS() *pci.ConfigSpace {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	cs.WriteU16(0x00, 0x8086)
	cs.WriteU16(0x02, 0x1533)
	cs.WriteU8(0x08, 0x03)
	cs.WriteU32(0x10, 0xFFFFF004)
	return cs
}

func TestWriteU32(t *testing.T) {
	cs := testCS()
	om := NewMap(cs)
	om.WriteU32(0x10, 0x00000000, "clear BAR0")
	if cs.ReadU32(0x10) != 0 {
		t.Error("WriteU32 should modify underlying config space")
	}
	if om.Count() != 1 {
		t.Errorf("count: got %d, want 1", om.Count())
	}
	e := om.Diff()[0]
	if e.OldValue != 0xFFFFF004 || e.NewValue != 0 {
		t.Errorf("diff values wrong: old=%X new=%X", e.OldValue, e.NewValue)
	}
	if e.Reason != "clear BAR0" {
		t.Errorf("reason: got %q", e.Reason)
	}
}

func TestWriteU32_NoOp(t *testing.T) {
	cs := testCS()
	om := NewMap(cs)
	om.WriteU32(0x00, cs.ReadU32(0x00), "should not record")
	if om.Count() != 0 {
		t.Error("writing same value should not record an entry")
	}
}

func TestWriteU16(t *testing.T) {
	cs := testCS()
	om := NewMap(cs)
	om.WriteU16(0x04, 0x0007, "enable bus master")
	if cs.ReadU16(0x04) != 0x0007 {
		t.Error("WriteU16 should modify config space")
	}
	if om.Count() != 1 {
		t.Errorf("count: got %d, want 1", om.Count())
	}
}

func TestWriteU8(t *testing.T) {
	cs := testCS()
	om := NewMap(cs)
	om.WriteU8(0x08, 0x05, "bump revision")
	if cs.ReadU8(0x08) != 0x05 {
		t.Error("WriteU8 should modify config space")
	}
}

func TestZeroRange(t *testing.T) {
	cs := testCS()
	// Write some data to vendor-specific range
	cs.WriteU32(0xC0, 0xDEADBEEF)
	cs.WriteU32(0xC4, 0xCAFEBABE)

	om := NewMap(cs)
	om.ZeroRange(0xC0, 0xD0, "clear vendor-specific")

	if cs.ReadU32(0xC0) != 0 || cs.ReadU32(0xC4) != 0 {
		t.Error("ZeroRange should zero the range")
	}
	if om.Count() != 1 {
		t.Errorf("count: got %d, want 1 (range entry)", om.Count())
	}
}

func TestZeroRange_AlreadyZero(t *testing.T) {
	cs := testCS()
	om := NewMap(cs)
	om.ZeroRange(0xC0, 0xD0, "should not record")
	if om.Count() != 0 {
		t.Error("zeroing already-zero range should not record")
	}
}

func TestFormatDiff(t *testing.T) {
	cs := testCS()
	om := NewMap(cs)
	om.WriteU32(0x10, 0, "clear BAR0")
	om.WriteU8(0x08, 0x05, "bump rev")

	report := om.FormatDiff()
	if !strings.Contains(report, "2 changes") {
		t.Error("report should mention change count")
	}
	if !strings.Contains(report, "clear BAR0") {
		t.Error("report should contain reason")
	}
}

func TestFormatDiff_Empty(t *testing.T) {
	cs := testCS()
	om := NewMap(cs)
	report := om.FormatDiff()
	if !strings.Contains(report, "No modifications") {
		t.Error("empty diff should say no modifications")
	}
}

func TestMultipleChanges(t *testing.T) {
	cs := testCS()
	om := NewMap(cs)
	om.WriteU32(0x10, 0, "clear BAR0")
	om.WriteU16(0x04, 0x0007, "enable bus master")
	om.WriteU8(0x3C, 0xFF, "set interrupt line")

	if om.Count() != 3 {
		t.Errorf("count: got %d, want 3", om.Count())
	}
	diff := om.Diff()
	if diff[0].Offset != 0x10 || diff[1].Offset != 0x04 || diff[2].Offset != 0x3C {
		t.Error("diff should preserve chronological order")
	}
}
