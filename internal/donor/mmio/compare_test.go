package mmio

import "testing"

func TestCompareTracesClassifiesAndSortsRegisters(t *testing.T) {
	traces := []*TraceResult{
		{Records: []AccessRecord{
			{Offset: 0x20, Type: AccessRead, Value: 1},
			{Offset: 0x10, Type: AccessRead, Value: 0xdeadbeef},
			{Offset: 0x30, Type: AccessWrite, Value: 7},
			{Offset: 0x30, Type: AccessRead, Value: 7},
			{Offset: 0x40, Type: AccessWrite, Value: 1},
		}},
		{Records: []AccessRecord{
			{Offset: 0x20, Type: AccessRead, Value: 2},
			{Offset: 0x10, Type: AccessRead, Value: 0xdeadbeef},
			{Offset: 0x30, Type: AccessWrite, Value: 9},
			{Offset: 0x30, Type: AccessRead, Value: 9},
			{Offset: 0x40, Type: AccessWrite, Value: 1},
		}},
	}

	got := CompareTraces(traces)
	if len(got) != 4 {
		t.Fatalf("CompareTraces() returned %d registers, want 4", len(got))
	}
	wantOffsets := []uint32{0x10, 0x20, 0x30, 0x40}
	wantClasses := []RegisterClass{RegisterStable, RegisterChanging, RegisterReadAfterWrite, RegisterInitializationWrite}
	for i := range got {
		if got[i].Offset != wantOffsets[i] || got[i].Classification != wantClasses[i] {
			t.Errorf("result[%d] = {offset: 0x%x, class: %q}, want {offset: 0x%x, class: %q}", i, got[i].Offset, got[i].Classification, wantOffsets[i], wantClasses[i])
		}
	}
	if len(got[1].Values) != 2 || got[1].Values[0] != 1 || got[1].Values[1] != 2 {
		t.Fatalf("changing values = %v, want [1 2]", got[1].Values)
	}
}

func TestCompareTracesIgnoresNilTraces(t *testing.T) {
	if got := CompareTraces([]*TraceResult{nil}); len(got) != 0 {
		t.Fatalf("CompareTraces(nil) = %v, want empty", got)
	}
}

func TestCompareTracesRequiresWriteBeforeReadForReadAfterWrite(t *testing.T) {
	trace := &TraceResult{Records: []AccessRecord{
		{Offset: 0x10, Type: AccessRead, Value: 1},
		{Offset: 0x10, Type: AccessWrite, Value: 1},
	}}

	got := CompareTraces([]*TraceResult{trace})

	if len(got) != 1 || got[0].Classification != RegisterStable {
		t.Fatalf("classification = %v, want %q", got, RegisterStable)
	}
}
