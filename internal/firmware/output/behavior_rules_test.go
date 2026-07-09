package output

import (
	"math"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/donor/behavior"
	"github.com/sercanarga/pcileechgen/internal/firmware"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

func TestBuildSVConfig_ConsumesPersistedBehaviorRules(t *testing.T) {
	rules := &behavior.RuleSet{
		Version:      behavior.RuleSchemaVersion,
		BARIndex:     0,
		BARSize:      0x1000,
		ClockHz:      100_000_000,
		InitialState: "idle",
		Rules: []behavior.Rule{{
			ID: "enable-ready", State: "idle", Access: behavior.AccessKind("write"),
			Width: 4, Offset: 0x20, Value: 1, ValueMask: math.MaxUint32,
			DelayedEvents: []behavior.DelayedEvent{{
				DelayCycles: 7,
				Updates: []behavior.RegisterUpdate{{Offset: 0x24, Width: 4, Value: 1, Mask: math.MaxUint32}},
				NextState: "ready",
			}},
			Confidence: 1, Provenance: []string{"test fixture"},
		}},
	}
	ctx := &donor.DeviceContext{
		Device:        pci.PCIDevice{VendorID: 0x1234, DeviceID: 0x5678, ClassCode: 0xff0000},
		ConfigSpace:   pci.NewConfigSpace(),
		BARs:          []pci.BAR{{Index: 0, Size: 0x1000, Type: pci.BARTypeMem32}},
		BARContents:   map[int][]byte{0: make([]byte, 0x1000)},
		BehaviorRules: rules,
	}
	ids := firmware.DeviceIDs{VendorID: 0x1234, DeviceID: 0x5678, ClassCode: 0xff0000}
	writer := NewOutputWriter(t.TempDir(), "", 1, 1)
	cfg, err := writer.buildSVConfig(ctx, ctx.ConfigSpace, ids, 0x12345678, &board.Board{BRAMSize: 0x1000})
	if err != nil {
		t.Fatalf("buildSVConfig: %v", err)
	}
	if cfg.BehaviorRules == nil {
		t.Fatal("build discarded persisted behavior rules")
	}
	if len(cfg.BehaviorRules.Rules) != 1 || cfg.BehaviorRules.Rules[0].ID != "enable-ready" {
		t.Fatalf("SV config behavior rules = %+v", cfg.BehaviorRules)
	}
	if got := cfg.BehaviorRules.Rules[0].DelayedEvents[0].DelayCycles; got != 7 {
		t.Fatalf("SV config delay = %d, want 7", got)
	}
}
