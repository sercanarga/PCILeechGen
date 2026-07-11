package svgen

import (
	"bytes"
	"math"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/donor/behavior"
)

func generatedRuleSet(updateValue uint64) *behavior.RuleSet {
	return &behavior.RuleSet{
		Version:      behavior.RuleSchemaVersion,
		BARIndex:     0,
		BARSize:      0x1000,
		ClockHz:      100_000_000,
		InitialState: "idle",
		Rules: []behavior.Rule{{
			ID: "enable-ready", State: "idle", Access: behavior.AccessKind("write"),
			Width: 4, Offset: 0x20, Value: 1, ValueMask: math.MaxUint32,
			NextState: "waiting",
			DelayedEvents: []behavior.DelayedEvent{{
				DelayCycles: 9,
				Updates:     []behavior.RegisterUpdate{{Offset: 0x24, Width: 4, Value: updateValue, Mask: math.MaxUint32}},
				NextState:   "ready",
			}},
			Confidence: 1, Provenance: []string{"test fixture"},
		}},
	}
}

func TestGenerateBarImplDeviceSV_RuleChangesAlterGeneratedFSM(t *testing.T) {
	readyLow := testConfig()
	readyLow.BehaviorRules = generatedRuleSet(1)
	lowSV, err := GenerateBarImplDeviceSV(readyLow)
	if err != nil {
		t.Fatalf("GenerateBarImplDeviceSV ready-low rule: %v", err)
	}

	readyHigh := testConfig()
	readyHigh.BehaviorRules = generatedRuleSet(2)
	highSV, err := GenerateBarImplDeviceSV(readyHigh)
	if err != nil {
		t.Fatalf("GenerateBarImplDeviceSV ready-high rule: %v", err)
	}
	if bytes.Equal([]byte(lowSV), []byte(highSV)) {
		t.Fatal("changing a learned register update did not alter generated RTL")
	}

	highSVAgain, err := GenerateBarImplDeviceSV(readyHigh)
	if err != nil {
		t.Fatalf("GenerateBarImplDeviceSV repeated: %v", err)
	}
	if highSV != highSVAgain {
		t.Fatal("same learned rules generated nondeterministic RTL")
	}
}

func TestGenerateBarImplDeviceSV_RejectsUnsupportedRuleWidth(t *testing.T) {
	cfg := testConfig()
	cfg.BehaviorRules = generatedRuleSet(1)
	cfg.BehaviorRules.Rules[0].Width = 8
	if _, err := GenerateBarImplDeviceSV(cfg); err == nil {
		t.Fatal("expected explicitly unsupported 8-byte RTL trigger to fail generation")
	}
}
