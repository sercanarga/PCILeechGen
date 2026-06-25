package mmio

import (
	"reflect"
	"testing"
	"time"
)

func TestDeriveTraceBAROverlay_SequentialAndStatic(t *testing.T) {
	trace := &TraceResult{
		Records: []AccessRecord{
			{Offset: 0x10, Type: AccessRead, Value: 0xDEADBEEF, Timestamp: 1 * time.Millisecond},
			{Offset: 0x10, Type: AccessRead, Value: 0x00000002, Timestamp: 2 * time.Millisecond},
			{Offset: 0x10, Type: AccessRead, Value: 0x00000001, Timestamp: 3 * time.Millisecond},
			{Offset: 0x20, Type: AccessRead, Value: 0xCAFEBABE, Timestamp: 1 * time.Millisecond},
			{Offset: 0x20, Type: AccessRead, Value: 0xCAFEBABE, Timestamp: 2 * time.Millisecond},
		},
	}

	got := DeriveTraceBAROverlay(trace)
	if got == nil {
		t.Fatal("overlay should not be nil")
	}
	if !reflect.DeepEqual(got.Sequential[0x10], []uint32{0xDEADBEEF, 0x00000002, 0x00000001}) {
		t.Fatalf("seq[0x10] = %#v", got.Sequential[0x10])
	}
	if got.Static[0x20] != 0xCAFEBABE {
		t.Fatalf("static[0x20] = 0x%08X", got.Static[0x20])
	}
}

func TestDeriveTraceBAROverlay_SkipsNoisyOffsets(t *testing.T) {
	trace := &TraceResult{
		Records: []AccessRecord{
			{Offset: 0x30, Type: AccessRead, Value: 1},
			{Offset: 0x30, Type: AccessRead, Value: 9},
			{Offset: 0x30, Type: AccessWrite, Value: 4},
			{Offset: 0x30, Type: AccessRead, Value: 2},
		},
	}

	got := DeriveTraceBAROverlay(trace)
	if _, ok := got.Sequential[0x30]; ok {
		t.Fatal("noisy offset should not become sequential")
	}
	if _, ok := got.Static[0x30]; ok {
		t.Fatal("noisy offset should not become static")
	}
}

func TestDeriveTraceBAROverlay_InfersWriteMask(t *testing.T) {
	trace := &TraceResult{
		Records: []AccessRecord{
			{Offset: 0x40, Type: AccessWrite, Value: 0x00000000},
			{Offset: 0x40, Type: AccessWrite, Value: 0x0000000F},
			{Offset: 0x40, Type: AccessWrite, Value: 0x00000007},
		},
	}

	got := DeriveTraceBAROverlay(trace)
	if got.WriteMask[0x40] != 0x0000000F {
		t.Fatalf("write mask for 0x40 = 0x%08X, want 0x0000000F", got.WriteMask[0x40])
	}
}

func TestDeriveTraceBAROverlay_InfersRW1CMask(t *testing.T) {
	trace := &TraceResult{
		Records: []AccessRecord{
			{Offset: 0x44, Type: AccessRead, Value: 0x00000001},
			{Offset: 0x44, Type: AccessWrite, Value: 0x00000000},
			{Offset: 0x44, Type: AccessWrite, Value: 0x00000001},
		},
	}

	got := DeriveTraceBAROverlay(trace)
	if got.WriteMask[0x44] != 0x00000001 {
		t.Fatalf("write mask for 0x44 = 0x%08X, want 0x00000001", got.WriteMask[0x44])
	}
	if got.RW1CMask[0x44] != 0x00000001 {
		t.Fatalf("rw1c mask for 0x44 = 0x%08X, want 0x00000001", got.RW1CMask[0x44])
	}
}

func TestRemapTraceToBAROffsets(t *testing.T) {
	trace := &TraceResult{Records: []AccessRecord{
		{Address: 0x80004010, Offset: 0x10, Type: AccessRead, Value: 1},
		{Address: 0x80004014, Offset: 0x14, Type: AccessRead, Value: 2},
		{Address: 0x90000000, Offset: 0x00, Type: AccessRead, Value: 3},
	}}

	got, err := RemapTraceToBAROffsets(trace, 0x80000000, 0x8000)
	if err != nil {
		t.Fatal(err)
	}
	if len(got.Records) != 2 {
		t.Fatalf("records = %d, want 2", len(got.Records))
	}
	if got.Records[0].Offset != 0x4010 || got.Records[1].Offset != 0x4014 {
		t.Fatalf("offsets not remapped: %#v", got.Records)
	}
}
