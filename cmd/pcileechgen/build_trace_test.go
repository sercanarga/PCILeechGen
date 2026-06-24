package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

func TestParseBuildTraceSpec(t *testing.T) {
	bar, path, err := parseBuildTraceSpec("2=/tmp/trace.log")
	if err != nil {
		t.Fatalf("parseBuildTraceSpec failed: %v", err)
	}
	if bar != 2 || path != "/tmp/trace.log" {
		t.Fatalf("parseBuildTraceSpec = (%d, %q), want (2, /tmp/trace.log)", bar, path)
	}
}

func TestApplyBuildTraceSpecs_AddsTraceToContext(t *testing.T) {
	tmpDir := t.TempDir()
	tracePath := filepath.Join(tmpDir, "bar2.trace")
	if err := os.WriteFile(tracePath, []byte("R 4 0.001 0xfebf001c 0x00000000\nW 4 0.002 0xfebf0024 0x001f001f\n"), 0644); err != nil {
		t.Fatalf("write trace fixture: %v", err)
	}
	ctx := &donor.DeviceContext{
		Device: pci.PCIDevice{BDF: pci.BDF{Domain: 0, Bus: 3, Device: 0, Function: 0}},
		BARs:   []pci.BAR{{Index: 2, Type: pci.BARTypeMem32, Address: 0xfebf0000, Size: 8192}},
	}

	if err := applyBuildTraceSpecs(ctx, []string{"2=" + tracePath}); err != nil {
		t.Fatalf("applyBuildTraceSpecs failed: %v", err)
	}

	trace := ctx.MMIOTraces[2]
	if trace == nil {
		t.Fatal("BAR2 trace missing from context")
	}
	if trace.BDF != "0000:03:00.0" || trace.BARIndex != 2 || trace.BARSize != 8192 {
		t.Fatalf("trace metadata mismatch: %#v", trace)
	}
	if len(trace.Records) != 2 || trace.Records[0].Offset != 0x1C || trace.Records[1].Offset != 0x24 {
		t.Fatalf("trace records mismatch: %#v", trace.Records)
	}
}

func TestApplyBuildTraceSpecs_UsesBARBase(t *testing.T) {
	tmpDir := t.TempDir()
	tracePath := filepath.Join(tmpDir, "bar2.trace")
	if err := os.WriteFile(tracePath, []byte("W 4 0.001 0xfebf1008 0x00000007\n"), 0644); err != nil {
		t.Fatalf("write trace fixture: %v", err)
	}
	ctx := &donor.DeviceContext{
		Device: pci.PCIDevice{BDF: pci.BDF{Domain: 0, Bus: 3, Device: 0, Function: 0}},
		BARs:   []pci.BAR{{Index: 2, Type: pci.BARTypeMem32, Address: 0xfebf0000, Size: 8192}},
	}

	if err := applyBuildTraceSpecs(ctx, []string{"2=" + tracePath}); err != nil {
		t.Fatalf("applyBuildTraceSpecs failed: %v", err)
	}
	if got := ctx.MMIOTraces[2].Records[0].Offset; got != 0x1008 {
		t.Fatalf("trace offset = 0x%X, want 0x1008", got)
	}
}
