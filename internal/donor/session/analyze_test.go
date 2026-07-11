package session

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/sercanarga/pcileechgen/internal/donor/mmio"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

func TestAnalyzeInitializationBuildsEvidencePhasesAndDependencies(t *testing.T) {
	captures := []CaptureData{
		analysisFixture(4),
		analysisFixture(6),
	}

	analysis := AnalyzeInitialization(captures)
	if analysis.SessionCount != 2 || len(analysis.Phases) < 2 {
		t.Fatalf("analysis summary = %+v", analysis)
	}
	reset := findRegister(t, analysis.Registers, 0x37)
	if reset.WriteEffect != WriteEffectSelfClearing || reset.Confidence != ConfidenceObserved {
		t.Fatalf("reset evidence = %+v", reset)
	}
	status := findRegister(t, analysis.Registers, 0x3e)
	if !status.Polling || status.Reads != 8 {
		t.Fatalf("status evidence = %+v", status)
	}
	if len(analysis.Dependencies) == 0 || analysis.Dependencies[0].WriteOffset != 0x37 || analysis.Dependencies[0].ReadOffset != 0x3e {
		t.Fatalf("dependencies = %+v", analysis.Dependencies)
	}
	if len(analysis.ConfigChanges) != 1 || analysis.ConfigChanges[0].Offset != 4 {
		t.Fatalf("config changes = %+v", analysis.ConfigChanges)
	}
}

func analysisFixture(terminal uint32) CaptureData {
	before := pci.NewConfigSpace()
	after := pci.NewConfigSpace()
	after.WriteU32(4, 7)
	trace := &mmio.TraceResult{Records: []mmio.AccessRecord{
		{Offset: 0x37, Width: 1, ByteEnable: 8, Type: mmio.AccessWrite, Value: 0x10, Timestamp: 0},
		{Offset: 0x3e, Width: 2, ByteEnable: 0xc, Type: mmio.AccessRead, Value: 1, Timestamp: 10 * time.Microsecond},
		{Offset: 0x3e, Width: 2, ByteEnable: 0xc, Type: mmio.AccessRead, Value: 1, Timestamp: 20 * time.Microsecond},
		{Offset: 0x3e, Width: 2, ByteEnable: 0xc, Type: mmio.AccessRead, Value: 1, Timestamp: 30 * time.Microsecond},
		{Offset: 0x3e, Width: 2, ByteEnable: 0xc, Type: mmio.AccessRead, Value: terminal, Timestamp: 40 * time.Microsecond},
		{Offset: 0x37, Width: 1, ByteEnable: 8, Type: mmio.AccessRead, Value: 0, Timestamp: 50 * time.Microsecond},
		{Offset: 0x40, Width: 4, ByteEnable: 0xf, Type: mmio.AccessWrite, Value: 1, Timestamp: 2 * time.Millisecond},
	}}
	return CaptureData{Manifest: &Manifest{Device: pci.PCIDevice{ClassCode: 0x020000}}, Trace: trace, ConfigBefore: before, ConfigAfter: after}
}

func findRegister(t *testing.T, registers []RegisterEvidence, offset uint32) RegisterEvidence {
	t.Helper()
	for _, register := range registers {
		if register.Offset == offset {
			return register
		}
	}
	t.Fatalf("register 0x%x not found", offset)
	return RegisterEvidence{}
}

func TestAnalysisJSONEncodesWidthsAsNumbers(t *testing.T) {
	analysis := AnalyzeInitialization([]CaptureData{analysisFixture(4)})
	data, err := json.Marshal(analysis)
	if err != nil {
		t.Fatal(err)
	}
	if string(data) == "" || !strings.Contains(string(data), `"widths":[1]`) {
		t.Fatalf("analysis JSON widths are not numeric: %s", data)
	}
}
