package mmio

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

// fixtureTrace builds a trace with one register per classification category,
// all inside a single 4KB BAR at base 0x80000000.
func fixtureTrace() []AccessRecord {
	return []AccessRecord{
		// static: one read, never written, single value.
		{Address: 0x80000000, Type: AccessRead, Value: 0x0040FF17},
		// polling: 5 reads, identical value.
		{Address: 0x8000001C, Type: AccessRead, Value: 0},
		{Address: 0x8000001C, Type: AccessRead, Value: 0},
		{Address: 0x8000001C, Type: AccessRead, Value: 0},
		{Address: 0x8000001C, Type: AccessRead, Value: 0},
		{Address: 0x8000001C, Type: AccessRead, Value: 0},
		// counter: monotonic increasing reads.
		{Address: 0x80000030, Type: AccessRead, Value: 1},
		{Address: 0x80000030, Type: AccessRead, Value: 2},
		{Address: 0x80000030, Type: AccessRead, Value: 3},
		{Address: 0x80000030, Type: AccessRead, Value: 4},
		// rw: written and read.
		{Address: 0x80000014, Type: AccessWrite, Value: 0x00460001},
		{Address: 0x80000014, Type: AccessRead, Value: 0x00460001},
		// volatile: multiple distinct non-monotonic reads.
		{Address: 0x80000040, Type: AccessRead, Value: 0xAA},
		{Address: 0x80000040, Type: AccessRead, Value: 0x11},
		{Address: 0x80000040, Type: AccessRead, Value: 0xEE},
		// out-of-BAR: address outside the window.
		{Address: 0x90000000, Type: AccessRead, Value: 0xDEAD},
	}
}

func fixtureBounds() map[int]BarRegion {
	return map[int]BarRegion{0: {Base: 0x80000000, Size: 0x1000}}
}

// TestBuildReport_ClassifiesAllCategories asserts each tag appears exactly
// once and points at the expected offset.
func TestBuildReport_ClassifiesAllCategories(t *testing.T) {
	r, warnings, err := BuildReport(fixtureTrace(), fixtureBounds())
	if err != nil {
		t.Fatalf("BuildReport error: %v", err)
	}
	if len(r.Bars) != 1 {
		t.Fatalf("bars: got %d, want 1", len(r.Bars))
	}
	br := r.Bars[0]
	wantTags := map[uint32]string{
		0x00: "static",
		0x1C: "polling",
		0x30: "counter",
		0x14: "rw",
		0x40: "volatile",
	}
	gotTags := map[string]uint32{}
	for _, reg := range br.Registers {
		gotTags[reg.Tag] = reg.Offset
		if want, ok := wantTags[reg.Offset]; ok && want != reg.Tag {
			t.Errorf("offset 0x%X tag: got %q, want %q", reg.Offset, reg.Tag, want)
		}
	}
	for wantOff, tag := range wantTags {
		gotOff, ok := gotTags[tag]
		if !ok {
			t.Errorf("missing tag %q (want offset 0x%X)", tag, wantOff)
		} else if gotOff != wantOff {
			t.Errorf("tag %q offset: got 0x%X, want 0x%X", tag, gotOff, wantOff)
		}
	}

	// Out-of-BAR quarantine.
	if len(r.Unknown) != 1 || r.Unknown[0].Address != 0x90000000 {
		t.Fatalf("unknown regions: %#v", r.Unknown)
	}
	if len(warnings) != 1 || warnings[0].Kind != "out_of_bar" {
		t.Fatalf("warnings: %#v", warnings)
	}
	if r.TotalOutOfBar != 1 {
		t.Errorf("total_out_of_bar: got %d, want 1", r.TotalOutOfBar)
	}
}

// TestBuildReport_DeterministicJSON proves two runs produce byte-identical
// JSON sorted by BAR then offset.
func TestBuildReport_DeterministicJSON(t *testing.T) {
	r1, _, err := BuildReport(fixtureTrace(), fixtureBounds())
	if err != nil {
		t.Fatal(err)
	}
	r2, _, err := BuildReport(fixtureTrace(), fixtureBounds())
	if err != nil {
		t.Fatal(err)
	}
	dir := t.TempDir()
	p1 := filepath.Join(dir, "a.json")
	p2 := filepath.Join(dir, "b.json")
	if err := r1.WriteJSON(p1); err != nil {
		t.Fatal(err)
	}
	if err := r2.WriteJSON(p2); err != nil {
		t.Fatal(err)
	}
	b1, err := os.ReadFile(p1)
	if err != nil {
		t.Fatal(err)
	}
	b2, err := os.ReadFile(p2)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(b1, b2) {
		t.Fatalf("non-deterministic JSON:\n--- run 1 ---\n%s\n--- run 2 ---\n%s", b1, b2)
	}

	// Registers must be sorted by offset.
	var parsed BarModelReport
	if err := json.Unmarshal(b1, &parsed); err != nil {
		t.Fatal(err)
	}
	for _, br := range parsed.Bars {
		for i := 1; i < len(br.Registers); i++ {
			if br.Registers[i].Offset <= br.Registers[i-1].Offset {
				t.Fatalf("registers not sorted by offset: 0x%X before 0x%X",
					br.Registers[i-1].Offset, br.Registers[i].Offset)
			}
		}
	}
}

// TestBuildReport_OutOfBARQuarantine verifies a trace with only out-of-BAR
// offsets is quarantined without crashing and still yields a report.
func TestBuildReport_OutOfBARQuarantine(t *testing.T) {
	records := []AccessRecord{
		{Address: 0x90000000, Type: AccessRead, Value: 1},
		{Address: 0xA0000000, Type: AccessWrite, Value: 2},
	}
	r, warnings, err := BuildReport(records, fixtureBounds())
	if err != nil {
		t.Fatalf("BuildReport error: %v", err)
	}
	if len(r.Bars) != 1 || len(r.Bars[0].Registers) != 0 {
		t.Fatalf("expected one BAR with zero registers, got %#v", r.Bars)
	}
	if len(r.Unknown) != 2 {
		t.Fatalf("unknown regions: got %d, want 2", len(r.Unknown))
	}
	if len(warnings) == 0 || warnings[0].Kind != "out_of_bar" {
		t.Fatalf("missing out_of_bar warning: %#v", warnings)
	}
}

// TestBuildReport_EmptyBoundsQuarantinesAll: with no BAR bounds, every record
// is quarantined - no panic, no error.
func TestBuildReport_EmptyBoundsQuarantinesAll(t *testing.T) {
	records := fixtureTrace()
	r, warnings, err := BuildReport(records, map[int]BarRegion{})
	if err != nil {
		t.Fatalf("BuildReport error: %v", err)
	}
	if len(r.Bars) != 0 {
		t.Fatalf("bars: got %d, want 0", len(r.Bars))
	}
	if len(r.Unknown) != len(records) {
		t.Fatalf("unknown: got %d, want %d", len(r.Unknown), len(records))
	}
	if len(warnings) == 0 {
		t.Fatal("expected out_of_bar warning")
	}
}

// TestBuildReport_NilRecords: empty input yields an empty but valid report.
func TestBuildReport_NilRecords(t *testing.T) {
	r, warnings, err := BuildReport(nil, fixtureBounds())
	if err != nil {
		t.Fatal(err)
	}
	if r == nil || len(r.Bars) != 1 || len(r.Bars[0].Registers) != 0 {
		t.Fatalf("expected empty report with one empty BAR, got %#v", r)
	}
	if len(warnings) != 0 {
		t.Fatalf("expected no warnings, got %#v", warnings)
	}
}

// TestParseMMIOTraceFile round-trips a fixture file through the parser.
func TestParseMMIOTraceFile(t *testing.T) {
	body := `# mmiotrace fixture
R 4 0.001 0x80000000 0x0040ff17
W 4 0.002 0x80000014 0x00460001
R 4 0.003 0x80000014 0x00460001
not a record line
`
	path := filepath.Join(t.TempDir(), "trace.log")
	if err := os.WriteFile(path, []byte(body), 0644); err != nil {
		t.Fatal(err)
	}
	got, err := ParseMMIOTraceFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 3 {
		t.Fatalf("records: got %d, want 3", len(got))
	}
	if got[0].Address != 0x80000000 || got[0].Type != AccessRead {
		t.Fatalf("record 0: %#v", got[0])
	}
	if got[1].Type != AccessWrite {
		t.Fatalf("record 1 not a write: %#v", got[1])
	}
}