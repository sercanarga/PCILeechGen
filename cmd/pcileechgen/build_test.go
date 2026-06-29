package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/sercanarga/pcileechgen/internal/donor/mmio"
)

func TestLoadTimingHistogram_NoPath(t *testing.T) {
	h, err := loadTimingHistogram("")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if h != nil {
		t.Errorf("empty path should yield nil histogram (synthetic defaults), got %+v", h)
	}
}

func TestLoadTimingHistogram_FromTrace(t *testing.T) {
	// 5 reads spaced 1us apart -> 4 inter-read intervals -> usable histogram.
	tr := mmio.TraceResult{
		BDF:      "0000:03:00.0",
		BARIndex: 0,
		BARSize:  4096,
		Duration: 5 * time.Microsecond,
	}
	for i := 0; i < 5; i++ {
		tr.Records = append(tr.Records, mmio.AccessRecord{
			Offset:    0x10,
			Type:      mmio.AccessRead,
			Value:     uint32(i),
			Timestamp: time.Duration(i) * time.Microsecond,
		})
	}
	data, err := json.Marshal(tr)
	if err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(t.TempDir(), "trace.json")
	if err := os.WriteFile(path, data, 0o644); err != nil {
		t.Fatal(err)
	}

	h, err := loadTimingHistogram(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if h == nil || h.SampleCount == 0 {
		t.Fatalf("expected a histogram with samples, got %+v", h)
	}
	if h.MinCycles <= 0 || h.MaxCycles < h.MinCycles {
		t.Errorf("implausible cycle range: min=%d max=%d", h.MinCycles, h.MaxCycles)
	}
}

func TestLoadTimingHistogram_BadFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "bad.json")
	if err := os.WriteFile(path, []byte("{not json"), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := loadTimingHistogram(path); err == nil {
		t.Error("expected error on malformed trace JSON")
	}
}
