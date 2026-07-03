package output

import (
	"strings"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/firmware"
)

func TestBuildSVConfig_ILAInstance(t *testing.T) {
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
		t.Fatalf("buildSVConfig (no ILA): %v", err)
	}
	if cfg.ILAInstanceSV != "" {
		t.Error("expected no ILA instance when ILADepth=0")
	}

	ow.ILADepth = 1024
	cfg2, err := ow.buildSVConfig(ctx, scrubbed, ids, entropy, b)
	if err != nil {
		t.Fatalf("buildSVConfig (ILA): %v", err)
	}
	if !strings.Contains(cfg2.ILAInstanceSV, "ila_0 u_ila_dbg") {
		t.Errorf("expected ILA instance, got %q", cfg2.ILAInstanceSV)
	}
}
