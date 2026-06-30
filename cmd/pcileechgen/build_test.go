package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

func TestDonorTimingHistogram_FromTrace(t *testing.T) {
	var b strings.Builder
	ts := 1000.0
	for i := 0; i < 64; i++ {
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
