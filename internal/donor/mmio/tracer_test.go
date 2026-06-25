package mmio

import (
	"strings"
	"testing"
	"time"
)

func sampleTrace() *TraceResult {
	return &TraceResult{
		BDF:      "0000:03:00.0",
		BARIndex: 0,
		BARSize:  4096,
		Duration: 5 * time.Second,
		Records: []AccessRecord{
			// Init writes
			{Offset: 0x14, Type: AccessWrite, Value: 0x00460001, Timestamp: 1 * time.Millisecond},
			{Offset: 0x20, Type: AccessWrite, Value: 0x00000000, Timestamp: 2 * time.Millisecond},
			{Offset: 0x24, Type: AccessWrite, Value: 0x001F001F, Timestamp: 3 * time.Millisecond},
			// Polling CSTS
			{Offset: 0x1C, Type: AccessRead, Value: 0x00000000, Timestamp: 10 * time.Millisecond},
			{Offset: 0x1C, Type: AccessRead, Value: 0x00000000, Timestamp: 20 * time.Millisecond},
			{Offset: 0x1C, Type: AccessRead, Value: 0x00000000, Timestamp: 30 * time.Millisecond},
			{Offset: 0x1C, Type: AccessRead, Value: 0x00000001, Timestamp: 40 * time.Millisecond},
			// CAP read
			{Offset: 0x00, Type: AccessRead, Value: 0x0040FF17, Timestamp: 50 * time.Millisecond},
		},
	}
}

func TestParseTraceReader_ParsesMMIOTraceLines(t *testing.T) {
	input := strings.NewReader(`
ignored line
R 4 0.001 0xfebf001c 0x00000000
W 4 0.002 0xfebf0024 0x001f001f
`)

	trace, err := ParseTraceReader(input, TraceImportOptions{
		BDF:      "0000:03:00.0",
		BARIndex: 2,
		BARSize:  4096,
	})
	if err != nil {
		t.Fatalf("ParseTraceReader failed: %v", err)
	}
	if trace.BDF != "0000:03:00.0" || trace.BARIndex != 2 || trace.BARSize != 4096 {
		t.Fatalf("trace identity mismatch: %#v", trace)
	}
	if len(trace.Records) != 2 {
		t.Fatalf("records = %d, want 2", len(trace.Records))
	}
	if trace.Records[0].Offset != 0x1C || trace.Records[0].Type != AccessRead {
		t.Fatalf("first record mismatch: %#v", trace.Records[0])
	}
	if trace.Records[1].Offset != 0x24 || trace.Records[1].Type != AccessWrite {
		t.Fatalf("second record mismatch: %#v", trace.Records[1])
	}
}

func TestParseTraceReader_UsesBARBaseForLargeBAROffsets(t *testing.T) {
	input := strings.NewReader("W 4 0.001 0xfebf1008 0x00000007\n")

	trace, err := ParseTraceReader(input, TraceImportOptions{
		BDF:      "0000:03:00.0",
		BARIndex: 0,
		BARSize:  8192,
		BARBase:  0xfebf0000,
	})
	if err != nil {
		t.Fatalf("ParseTraceReader failed: %v", err)
	}
	if len(trace.Records) != 1 || trace.Records[0].Offset != 0x1008 {
		t.Fatalf("trace offset mismatch: %#v", trace.Records)
	}
}

func TestAnalyze_Basic(t *testing.T) {
	p := Analyze(sampleTrace())
	if p.TotalAccesses != 8 {
		t.Errorf("total accesses: got %d, want 8", p.TotalAccesses)
	}
	if p.TotalReads != 5 {
		t.Errorf("total reads: got %d, want 5", p.TotalReads)
	}
	if p.TotalWrites != 3 {
		t.Errorf("total writes: got %d, want 3", p.TotalWrites)
	}
}

func TestAnalyze_HotRegisters(t *testing.T) {
	p := Analyze(sampleTrace())
	if len(p.HotRegisters) == 0 {
		t.Fatal("no hot registers found")
	}
	// CSTS (0x1C) should be the hottest - 4 reads
	if p.HotRegisters[0].Offset != 0x1C {
		t.Errorf("hottest register: got 0x%X, want 0x1C", p.HotRegisters[0].Offset)
	}
	if p.HotRegisters[0].ReadCount != 4 {
		t.Errorf("CSTS read count: got %d, want 4", p.HotRegisters[0].ReadCount)
	}
}

func TestAnalyze_PollingDetection(t *testing.T) {
	p := Analyze(sampleTrace())
	if len(p.PollingLoops) == 0 {
		t.Fatal("no polling loops detected")
	}
	found := false
	for _, pl := range p.PollingLoops {
		if pl.Offset == 0x1C {
			found = true
			if pl.Count != 4 {
				t.Errorf("CSTS poll count: got %d, want 4", pl.Count)
			}
		}
	}
	if !found {
		t.Error("CSTS polling not detected")
	}
}

func TestAnalyze_InitSequence(t *testing.T) {
	p := Analyze(sampleTrace())
	if len(p.InitSequence) != 3 {
		t.Fatalf("init sequence: got %d entries, want 3", len(p.InitSequence))
	}
	// First write should be CC at 0x14
	if p.InitSequence[0].Offset != 0x14 {
		t.Errorf("first init write: got 0x%X, want 0x14", p.InitSequence[0].Offset)
	}
}

func TestBuildReport_ClassifiesRegisters(t *testing.T) {
	report := BuildLegacyTraceReport(sampleTrace())
	if report.BDF != "0000:03:00.0" || report.BARIndex != 0 {
		t.Fatalf("report identity mismatch: %#v", report)
	}
	if report.TotalReads != 5 || report.TotalWrites != 3 {
		t.Fatalf("report totals = R:%d W:%d, want R:5 W:3", report.TotalReads, report.TotalWrites)
	}

	byOffset := make(map[uint32]RegisterSummary)
	for _, reg := range report.Registers {
		byOffset[reg.Offset] = reg
	}
	if got := byOffset[0x1C].Classification; got != RegisterPolled {
		t.Fatalf("CSTS classification = %s, want %s", got, RegisterPolled)
	}
	if got := byOffset[0x14].Classification; got != RegisterWriteOnly {
		t.Fatalf("CC classification = %s, want %s", got, RegisterWriteOnly)
	}
	if got := byOffset[0x00].Classification; got != RegisterStaticRead {
		t.Fatalf("CAP classification = %s, want %s", got, RegisterStaticRead)
	}
}

func TestAnalyze_Nil(t *testing.T) {
	p := Analyze(nil)
	if p.TotalAccesses != 0 {
		t.Error("nil trace should produce empty pattern")
	}
}

func TestAnalyze_Empty(t *testing.T) {
	p := Analyze(&TraceResult{})
	if p.TotalAccesses != 0 {
		t.Error("empty trace should produce empty pattern")
	}
}

func TestFormatReport(t *testing.T) {
	p := Analyze(sampleTrace())
	report := FormatReport(p)
	if !strings.Contains(report, "MMIO Trace Analysis") {
		t.Error("report should contain header")
	}
	if !strings.Contains(report, "0x01C") {
		t.Error("report should contain CSTS register offset")
	}
	if !strings.Contains(report, "Polling") {
		t.Error("report should contain polling section")
	}
	if !strings.Contains(report, "Init Sequence") {
		t.Error("report should contain init sequence section")
	}
}

func TestFormatReport_Nil(t *testing.T) {
	report := FormatReport(nil)
	if !strings.Contains(report, "No trace data") {
		t.Error("nil pattern should produce 'no data' message")
	}
}
