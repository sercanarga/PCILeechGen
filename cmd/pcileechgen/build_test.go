package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/donor/mmio"
	"github.com/sercanarga/pcileechgen/internal/pci"
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
	for i := range 5 {
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

func TestDonorTimingHistogram_FromTrace(t *testing.T) {
	var b strings.Builder
	ts := 1000.0
	for i := range 64 {
		ts += 0.000002 * float64(1+(i%5))
		fmt.Fprintf(&b, "R 4 %.6f 2 0xf780010c 0x%x 0x0 0\n", ts, i)
	}
	tracePath := filepath.Join(t.TempDir(), "mmiotrace.txt")
	if err := os.WriteFile(tracePath, []byte(b.String()), 0o644); err != nil {
		t.Fatalf("write trace: %v", err)
	}

	ctx := &donor.DeviceContext{
		Device:      pci.PCIDevice{VendorID: 0x144D, DeviceID: 0xA808, ClassCode: 0x010802},
		BARs:        []pci.BAR{{Index: 2, Size: 4096, Address: 0xf7800000}},
		BARContents: map[int][]byte{2: make([]byte, 4096)},
	}

	h, err := donorTimingHistogram(tracePath, ctx)
	if err != nil {
		t.Fatalf("donorTimingHistogram: %v", err)
	}
	if h == nil || h.SampleCount == 0 {
		t.Fatalf("expected a populated histogram, got %+v", h)
	}
}
