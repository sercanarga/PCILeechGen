package output

import (
	"testing"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/donor/behavior"
	"github.com/sercanarga/pcileechgen/internal/firmware"
)

func TestBuildSVConfig_UsesDonorHistogram(t *testing.T) {
	ctx := makeDonorContext(0x144D, 0xA808, ccNVMe)
	b, err := board.Find("PCIeSquirrel")
	if err != nil {
		t.Fatalf("board.Find: %v", err)
	}
	ids := firmware.ExtractDeviceIDs(ctx.ConfigSpace, ctx.ExtCapabilities)
	ow := NewOutputWriter(t.TempDir(), "", 0, 0)
	scrubbed, entropy, _ := ow.scrubAndVary(ctx, b, ids)

	cfg, err := ow.buildSVConfig(ctx, scrubbed, ids, entropy, b)
	if err != nil {
		t.Fatalf("buildSVConfig (no histogram): %v", err)
	}
	if cfg.LatencyConfig.HasHistogram {
		t.Error("expected HasHistogram=false without a donor histogram")
	}

	h := &behavior.TimingHistogram{
		SampleCount:  500,
		MinCycles:    2,
		MaxCycles:    9,
		MedianCycles: 4,
	}
	for i := range h.Buckets {
		h.Buckets[i] = uint8(i)
		h.CDF[i] = uint8((i + 1) * 16)
	}
	ow.TimingHistogram = h
	cfg2, err := ow.buildSVConfig(ctx, scrubbed, ids, entropy, b)
	if err != nil {
		t.Fatalf("buildSVConfig (with histogram): %v", err)
	}
	if !cfg2.LatencyConfig.HasHistogram {
		t.Error("expected HasHistogram=true when a donor histogram is supplied")
	}
	if cfg2.LatencyConfig.CDF != h.CDF {
		t.Errorf("expected CDF derived from donor histogram, got %v", cfg2.LatencyConfig.CDF)
	}
}
