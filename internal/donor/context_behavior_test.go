package donor

import (
	"bytes"
	"math"
	"os"
	"path/filepath"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/donor/behavior"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

func TestDeviceContext_BehaviorRulesRoundTripDeterministically(t *testing.T) {
	ctx := &DeviceContext{
		ConfigSpace: pci.NewConfigSpace(),
		BehaviorRules: &behavior.RuleSet{
			Version:      behavior.RuleSchemaVersion,
			BDF:          "0000:03:00.0",
			BARIndex:     2,
			BARSize:      0x4000,
			ClockHz:      100_000_000,
			InitialState: "idle",
			Rules: []behavior.Rule{{
				ID: "enable-ready", State: "idle", Access: behavior.AccessKind("write"),
				Width: 4, Offset: 0x20, Value: 1, ValueMask: math.MaxUint32,
				NextState: "waiting",
				DelayedEvents: []behavior.DelayedEvent{{
					DelayCycles: 12,
					Updates:     []behavior.RegisterUpdate{{Offset: 0x24, Width: 4, Value: 1, Mask: math.MaxUint32}},
					NextState:   "ready",
				}},
				Confidence: 1, Provenance: []string{"test fixture"},
			}},
		},
	}

	firstPath := filepath.Join(t.TempDir(), "first.json")
	secondPath := filepath.Join(t.TempDir(), "second.json")
	if err := SaveContext(ctx, firstPath); err != nil {
		t.Fatalf("SaveContext first: %v", err)
	}
	loaded, err := LoadContext(firstPath)
	if err != nil {
		t.Fatalf("LoadContext: %v", err)
	}
	if loaded.BehaviorRules == nil {
		t.Fatal("behavior rules were dropped while loading device context")
	}
	if err := behavior.Validate(loaded.BehaviorRules); err != nil {
		t.Fatalf("loaded behavior rules are invalid: %v", err)
	}
	if len(loaded.BehaviorRules.Rules) != 1 || loaded.BehaviorRules.Rules[0].ID != "enable-ready" {
		t.Fatalf("loaded behavior rules = %+v", loaded.BehaviorRules)
	}
	if err := SaveContext(loaded, secondPath); err != nil {
		t.Fatalf("SaveContext second: %v", err)
	}
	firstJSON, err := os.ReadFile(firstPath)
	if err != nil {
		t.Fatalf("read first context: %v", err)
	}
	secondJSON, err := os.ReadFile(secondPath)
	if err != nil {
		t.Fatalf("read second context: %v", err)
	}
	if !bytes.Equal(firstJSON, secondJSON) {
		t.Fatalf("context rule persistence is nondeterministic after round trip")
	}
}
