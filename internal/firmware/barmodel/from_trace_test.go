package barmodel

import (
	"strings"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/donor/mmio"
)

// traceReport builds a BarModelReport with one BAR covering the given
// registers. Each register is described by its tag, offset, values, and
// last_value.
func traceReport(regs []mmio.RegisterReport) *mmio.BarModelReport {
	br := mmio.BarReport{BAR: 0, Base: 0x80000000, Size: 0x1000}
	br.Registers = append(br.Registers, regs...)
	return &mmio.BarModelReport{Bars: []mmio.BarReport{br}}
}

func rep(bar, off int, tag string, values []uint32, last uint32, reads, writes int) mmio.RegisterReport {
	vs := append([]uint32(nil), values...)
	return mmio.RegisterReport{
		BAR: bar, Offset: uint32(off), Tag: tag,
		Values: vs, LastValue: last, ReadCount: reads, WriteCount: writes,
	}
}

// TestFromTraceReport_StaticAndRW_ResetValuesAndProvenance: happy path. A
// static register gets reset = the single observed value and RWMask=0; an RW
// register gets reset = LastValue and RWMask=0xFFFFFFFF. Both carry
// provenance citing the trace report.
func TestFromTraceReport_StaticAndRW_ResetValuesAndProvenance(t *testing.T) {
	report := traceReport([]mmio.RegisterReport{
		rep(0, 0x00, "static", []uint32{0x0040FF17}, 0x0040FF17, 1, 0),
		rep(0, 0x14, "rw", []uint32{0x00460001}, 0x00460001, 1, 1),
	})

	sug, err := FromTraceReport(report)
	if err != nil {
		t.Fatalf("FromTraceReport: %v", err)
	}
	if sug == nil || sug.Model == nil {
		t.Fatal("nil suggestion")
	}
	if len(sug.Model.Registers) != 2 {
		t.Fatalf("registers: got %d, want 2", len(sug.Model.Registers))
	}

	byOff := map[uint32]BARRegister{}
	for _, r := range sug.Model.Registers {
		byOff[r.Offset] = r
	}

	if r, ok := byOff[0x00]; !ok {
		t.Fatal("missing static reg at 0x00")
	} else {
		if r.Reset != 0x0040FF17 {
			t.Errorf("static reset: got 0x%08X, want 0x0040FF17", r.Reset)
		}
		if r.RWMask != 0 {
			t.Errorf("static RWMask: got 0x%08X, want 0", r.RWMask)
		}
	}

	if r, ok := byOff[0x14]; !ok {
		t.Fatal("missing rw reg at 0x14")
	} else {
		if r.Reset != 0x00460001 {
			t.Errorf("rw reset: got 0x%08X, want 0x00460001 (LastValue)", r.Reset)
		}
		if r.RWMask != 0xFFFFFFFF {
			t.Errorf("rw RWMask: got 0x%08X, want 0xFFFFFFFF", r.RWMask)
		}
	}

	// Provenance.
	if sug.Provenance == nil {
		t.Fatal("nil provenance")
	}
	if sug.Provenance.Source != "trace" {
		t.Errorf("provenance source: got %q, want \"trace\"", sug.Provenance.Source)
	}
	if sug.Provenance.ReportDigest == "" {
		t.Error("empty report digest")
	}
	if len(sug.Provenance.Registers) != 2 {
		t.Fatalf("provenance registers: got %d, want 2", len(sug.Provenance.Registers))
	}
	provByOff := map[uint32]RegisterProvenance{}
	for _, p := range sug.Provenance.Registers {
		provByOff[p.Offset] = p
	}
	if p, ok := provByOff[0x00]; !ok {
		t.Fatal("missing provenance for 0x00")
	} else {
		if p.Tag != "static" || p.ReadCount != 1 || p.WriteCount != 0 || p.SampleCount != 1 {
			t.Errorf("provenance 0x00: %#v", p)
		}
	}
	if p, ok := provByOff[0x14]; !ok {
		t.Fatal("missing provenance for 0x14")
	} else {
		if p.Tag != "rw" || p.ReadCount != 1 || p.WriteCount != 1 || p.SampleCount != 1 {
			t.Errorf("provenance 0x14: %#v", p)
		}
	}
}

// TestFromTraceReport_ConflictingValuesNotStableState: failure QA scenario. A
// volatile offset (multiple conflicting read values) and a counter must NOT
// appear as stable registers. They must be omitted and warned about.
func TestFromTraceReport_ConflictingValuesNotStableState(t *testing.T) {
	report := traceReport([]mmio.RegisterReport{
		rep(0, 0x40, "volatile", []uint32{0xAA, 0x11, 0xEE}, 0xEE, 3, 0),
		rep(0, 0x30, "counter", []uint32{1, 2, 3, 4}, 4, 4, 0),
		rep(0, 0x1C, "polling", []uint32{0, 0, 0, 0, 0}, 0, 5, 0),
		rep(0, 0x50, "unknown", []uint32{0xDEAD}, 0xDEAD, 1, 0),
		rep(0, 0x00, "static", []uint32{0xABCD}, 0xABCD, 1, 0),
	})

	sug, err := FromTraceReport(report)
	if err != nil {
		t.Fatalf("FromTraceReport: %v", err)
	}

	// Only the static register should be emitted.
	if len(sug.Model.Registers) != 1 {
		t.Fatalf("emitted registers: got %d, want 1 (only static)", len(sug.Model.Registers))
	}
	if sug.Model.Registers[0].Offset != 0x00 {
		t.Errorf("emitted offset: got 0x%X, want 0x00", sug.Model.Registers[0].Offset)
	}

	// The volatile/counter/polling/unknown offsets must NOT be present as
	// stable HDL state.
	emitted := map[uint32]bool{}
	for _, r := range sug.Model.Registers {
		emitted[r.Offset] = true
	}
	for _, bad := range []uint32{0x40, 0x30, 0x1C, 0x50} {
		if emitted[bad] {
			t.Errorf("offset 0x%X must NOT be emitted as stable state", bad)
		}
	}

	// Warnings must mention each omitted offset.
	if len(sug.Warnings) != 4 {
		t.Fatalf("warnings: got %d, want 4", len(sug.Warnings))
	}
	warned := map[string]bool{}
	for _, w := range sug.Warnings {
		warned[w.Message] = true
	}
	for _, off := range []string{"0x00000040", "0x00000030", "0x0000001C", "0x00000050"} {
		found := false
		for msg := range warned {
			if strings.Contains(msg, off) {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("no warning cites offset %s", off)
		}
	}

	// Provenance must still record every offset (emitted or not) so a
	// reviewer can see what was considered and rejected.
	if len(sug.Provenance.Registers) != 5 {
		t.Errorf("provenance registers: got %d, want 5 (all considered)", len(sug.Provenance.Registers))
	}
}

// TestFromTraceReport_Determinism: the same report must yield byte-identical
// suggestions across calls. Field order is fixed by structs; registers and
// provenance are sorted by offset.
func TestFromTraceReport_Determinism(t *testing.T) {
	build := func() *TraceSuggestion {
		return mustSuggest(t, traceReport([]mmio.RegisterReport{
			rep(0, 0x14, "rw", []uint32{0x1, 0x1}, 0x1, 1, 1),
			rep(0, 0x00, "static", []uint32{0x2}, 0x2, 1, 0),
			rep(0, 0x30, "volatile", []uint32{0xA, 0xB}, 0xB, 2, 0),
		}))
	}
	a, b := build(), build()
	if a == nil || b == nil {
		t.Fatal("nil suggestion")
	}
	if a.Provenance.ReportDigest != b.Provenance.ReportDigest {
		t.Errorf("digest differs: %s vs %s", a.Provenance.ReportDigest, b.Provenance.ReportDigest)
	}
	if len(a.Model.Registers) != len(b.Model.Registers) {
		t.Fatalf("register count differs")
	}
	for i := range a.Model.Registers {
		if !regsEqual(a.Model.Registers[i], b.Model.Registers[i]) {
			t.Fatalf("register %d differs: %#v vs %#v", i, a.Model.Registers[i], b.Model.Registers[i])
		}
	}
	if len(a.Provenance.Registers) != len(b.Provenance.Registers) {
		t.Fatalf("provenance count differs")
	}
	for i := range a.Provenance.Registers {
		if a.Provenance.Registers[i] != b.Provenance.Registers[i] {
			t.Fatalf("provenance %d differs: %#v vs %#v", i, a.Provenance.Registers[i], b.Provenance.Registers[i])
		}
	}
	// Sorted by offset.
	if a.Model.Registers[0].Offset != 0x00 || a.Model.Registers[1].Offset != 0x14 {
		t.Errorf("registers not sorted: %#v", a.Model.Registers)
	}
}

// TestSuggestFromTrace_DefaultOff: the gated entry point returns nil with no
// warnings and no error when Enabled is false. This proves HDL emission from
// traces is default-OFF.
func TestSuggestFromTrace_DefaultOff(t *testing.T) {
	report := traceReport([]mmio.RegisterReport{
		rep(0, 0x00, "static", []uint32{0xABCD}, 0xABCD, 1, 0),
	})
	got, err := SuggestFromTrace(report, TraceSuggestionOptions{Enabled: false})
	if err != nil {
		t.Fatalf("SuggestFromTrace disabled: %v", err)
	}
	if got != nil {
		t.Fatalf("disabled SuggestFromTrace must return nil, got %#v", got)
	}
	// Explicit opt-in works.
	got, err = SuggestFromTrace(report, TraceSuggestionOptions{Enabled: true})
	if err != nil {
		t.Fatalf("SuggestFromTrace enabled: %v", err)
	}
	if got == nil || got.Model == nil || len(got.Model.Registers) != 1 {
		t.Fatalf("enabled SuggestFromTrace must return the suggestion, got %#v", got)
	}
}

// TestFromTraceReport_NilReport: nil input is a typed error, never a panic.
func TestFromTraceReport_NilReport(t *testing.T) {
	if _, err := FromTraceReport(nil); err == nil {
		t.Fatal("nil report must error")
	}
}

func mustSuggest(t *testing.T, report *mmio.BarModelReport) *TraceSuggestion {
	t.Helper()
	sug, err := FromTraceReport(report)
	if err != nil {
		t.Fatalf("FromTraceReport: %v", err)
	}
	if sug == nil {
		t.Fatal("nil suggestion")
	}
	return sug
}

// regsEqual compares two BARRegisters field-by-field, including the
// SequentialValues slice (BARRegister contains a slice so == is not allowed).
func regsEqual(a, b BARRegister) bool {
	if a.Offset != b.Offset || a.Width != b.Width || a.Reset != b.Reset ||
		a.RWMask != b.RWMask || a.Name != b.Name || a.IsRW1C != b.IsRW1C ||
		a.IsFSMDriven != b.IsFSMDriven || a.SequentialRead != b.SequentialRead ||
		len(a.SequentialValues) != len(b.SequentialValues) {
		return false
	}
	for i := range a.SequentialValues {
		if a.SequentialValues[i] != b.SequentialValues[i] {
			return false
		}
	}
	return true
}