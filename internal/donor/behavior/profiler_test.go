package behavior

import (
	"strings"
	"testing"
	"time"

	"github.com/sercanarga/pcileechgen/internal/donor/mmio"
)

func nvmeTrace() *mmio.TraceResult {
	return &mmio.TraceResult{
		BDF:      "0000:03:00.0",
		BARIndex: 0,
		BARSize:  4096,
		Duration: 2 * time.Second,
		Records: []mmio.AccessRecord{
			{Offset: 0x00, Type: mmio.AccessRead, Value: 0x0040FF17, Timestamp: 1 * time.Millisecond},
			{Offset: 0x08, Type: mmio.AccessRead, Value: 0x00010400, Timestamp: 2 * time.Millisecond},
			{Offset: 0x14, Type: mmio.AccessWrite, Value: 0x00460001, Timestamp: 5 * time.Millisecond},
			{Offset: 0x1C, Type: mmio.AccessRead, Value: 0x00000000, Timestamp: 10 * time.Millisecond},
			{Offset: 0x1C, Type: mmio.AccessRead, Value: 0x00000001, Timestamp: 20 * time.Millisecond},
			{Offset: 0x24, Type: mmio.AccessWrite, Value: 0x001F001F, Timestamp: 25 * time.Millisecond},
		},
	}
}

func TestFromMMIOTrace_NVMe(t *testing.T) {
	profile := FromMMIOTrace(nvmeTrace(), 0x010802)
	if profile.DeviceBDF != "0000:03:00.0" {
		t.Errorf("BDF: got %q", profile.DeviceBDF)
	}
	if len(profile.InitSequence) == 0 {
		t.Fatal("init sequence should not be empty")
	}
	// First access should be CAP read
	if profile.InitSequence[0].Offset != 0x00 {
		t.Errorf("first init step: got 0x%X, want 0x00", profile.InitSequence[0].Offset)
	}
	if profile.InitSequence[0].Type != "read" {
		t.Error("first step should be a read")
	}
}

func TestFromMMIOTrace_PurposeAnnotation(t *testing.T) {
	profile := FromMMIOTrace(nvmeTrace(), 0x010802)
	found := false
	for _, step := range profile.InitSequence {
		if step.Offset == 0x14 && strings.Contains(step.Purpose, "CC") {
			found = true
		}
	}
	if !found {
		t.Error("CC write should have NVMe purpose annotation")
	}
}

func TestFromMMIOTrace_Stats(t *testing.T) {
	profile := FromMMIOTrace(nvmeTrace(), 0x010802)
	if profile.AccessStats.TotalReads != 4 {
		t.Errorf("reads: got %d, want 4", profile.AccessStats.TotalReads)
	}
	if profile.AccessStats.TotalWrites != 2 {
		t.Errorf("writes: got %d, want 2", profile.AccessStats.TotalWrites)
	}
}

func TestFromMMIOTrace_Nil(t *testing.T) {
	profile := FromMMIOTrace(nil, 0)
	if profile == nil {
		t.Fatal("nil trace should return empty profile, not nil")
	}
}

func TestFormatReport(t *testing.T) {
	profile := FromMMIOTrace(nvmeTrace(), 0x010802)
	report := FormatReport(profile)
	if !strings.Contains(report, "Behavior Profile") {
		t.Error("report missing header")
	}
	if !strings.Contains(report, "Initialization Sequence") {
		t.Error("report missing init sequence")
	}
}

func TestFormatReport_Nil(t *testing.T) {
	report := FormatReport(nil)
	if !strings.Contains(report, "No behavior profile") {
		t.Error("nil profile should say no data")
	}
}
