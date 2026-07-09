package main

import (
	"encoding/json"
	"math"
	"os"
	"path/filepath"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/donor/behavior"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

func TestLoadDonorContext_AttachesValidatedBehaviorRules(t *testing.T) {
	dir := t.TempDir()
	contextPath := filepath.Join(dir, "device_context.json")
	rulesPath := filepath.Join(dir, "behavior_rules.json")
	ctx := &donor.DeviceContext{ConfigSpace: pci.NewConfigSpace()}
	if err := donor.SaveContext(ctx, contextPath); err != nil {
		t.Fatalf("save context: %v", err)
	}
	rules := &behavior.RuleSet{
		Version: behavior.RuleSchemaVersion, BARIndex: 0, BARSize: 0x1000,
		ClockHz: 100_000_000, InitialState: "idle",
		Rules: []behavior.Rule{{
			ID: "enable", State: "idle", Access: behavior.AccessKind("write"), Width: 4,
			Offset: 0x20, Value: 1, ValueMask: math.MaxUint32, NextState: "ready",
			Confidence: 1, Provenance: []string{"test fixture"},
		}},
	}
	data, err := json.Marshal(rules)
	if err != nil {
		t.Fatalf("marshal rules: %v", err)
	}
	if err := os.WriteFile(rulesPath, data, 0o644); err != nil {
		t.Fatalf("write rules: %v", err)
	}

	previous := buildOpts
	t.Cleanup(func() { buildOpts = previous })
	buildOpts = buildFlags{fromJSON: contextPath, behaviorRules: rulesPath}
	loaded, err := loadDonorContext()
	if err != nil {
		t.Fatalf("loadDonorContext: %v", err)
	}
	if loaded.BehaviorRules == nil || len(loaded.BehaviorRules.Rules) != 1 {
		t.Fatalf("loaded behavior rules = %+v", loaded.BehaviorRules)
	}
	if loaded.BehaviorRules.Rules[0].ID != "enable" {
		t.Fatalf("loaded rule ID = %q, want enable", loaded.BehaviorRules.Rules[0].ID)
	}
}

func TestLoadDonorContext_RejectsInvalidBehaviorRules(t *testing.T) {
	dir := t.TempDir()
	contextPath := filepath.Join(dir, "device_context.json")
	rulesPath := filepath.Join(dir, "behavior_rules.json")
	if err := donor.SaveContext(&donor.DeviceContext{ConfigSpace: pci.NewConfigSpace()}, contextPath); err != nil {
		t.Fatalf("save context: %v", err)
	}
	if err := os.WriteFile(rulesPath, []byte(`{"version":999,"bar_size":4096,"initial_state":"idle"}`), 0o644); err != nil {
		t.Fatalf("write invalid rules: %v", err)
	}

	previous := buildOpts
	t.Cleanup(func() { buildOpts = previous })
	buildOpts = buildFlags{fromJSON: contextPath, behaviorRules: rulesPath}
	if _, err := loadDonorContext(); err == nil {
		t.Fatal("expected unsupported rule schema version to be rejected")
	}
}
