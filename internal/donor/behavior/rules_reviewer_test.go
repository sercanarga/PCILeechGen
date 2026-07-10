package behavior

import (
	"math"
	"testing"
	"time"

	"github.com/sercanarga/pcileechgen/internal/donor/mmio"
)

func validReviewerRuleSet() *RuleSet {
	return &RuleSet{
		Version: RuleSchemaVersion, BARSize: 0x1000, ClockHz: 100_000_000, InitialState: "idle",
	}
}

func TestValidateUpdateCompositionAndContradiction(t *testing.T) {
	set := validReviewerRuleSet()
	set.Rules = []Rule{{
		ID: "compose", State: "idle", Access: AccessWrite, Width: 4,
		Offset: 0x20, Value: 1, ValueMask: math.MaxUint32,
		Updates: []RegisterUpdate{
			{Offset: 0x24, Width: 4, Value: 0x00000005, Mask: 0x0000000f},
			{Offset: 0x24, Width: 4, Value: 0x000000a0, Mask: 0x000000f0},
		},
		DelayedEvents: []DelayedEvent{
			{DelayCycles: 3, Updates: []RegisterUpdate{
				{Offset: 0x24, Width: 4, Value: 0x00000300, Mask: 0x00000f00},
			}},
			{DelayCycles: 3, Updates: []RegisterUpdate{
				{Offset: 0x24, Width: 4, Value: 0x0000c000, Mask: 0x0000f000},
			}},
		},
		Confidence: 1, Provenance: []string{"test"},
	}}
	if err := Validate(set); err != nil {
		t.Fatalf("disjoint same-register updates should compose: %v", err)
	}
	engine, err := NewEngine(set)
	if err != nil {
		t.Fatalf("NewEngine: %v", err)
	}
	if err := engine.Apply(mmio.AccessRecord{Type: mmio.AccessWrite, Width: 4, Offset: 0x20, Value: 1}); err != nil {
		t.Fatalf("Apply: %v", err)
	}
	if got, ok := engine.Register(0x24); !ok || got != 0x000000a5 {
		t.Fatalf("immediate composed value = %#x, present=%v, want 0xa5", got, ok)
	}
	if err := engine.Advance(3); err != nil {
		t.Fatalf("Advance: %v", err)
	}
	if got, ok := engine.Register(0x24); !ok || got != 0x0000c3a5 {
		t.Fatalf("delayed composed value = %#x, present=%v, want 0xc3a5", got, ok)
	}

	set.Rules[0].Updates = []RegisterUpdate{
		{Offset: 0x24, Width: 4, Value: 0, Mask: 1},
		{Offset: 0x24, Width: 4, Value: 1, Mask: 1},
	}
	set.Rules[0].DelayedEvents = nil
	if err := Validate(set); err == nil {
		t.Fatal("contradictory overlapping updates to one register should be rejected")
	}
	set.Rules[0].Updates = nil
	set.Rules[0].DelayedEvents = []DelayedEvent{
		{DelayCycles: 3, Updates: []RegisterUpdate{{Offset: 0x24, Width: 4, Value: 0, Mask: 1}}},
		{DelayCycles: 3, Updates: []RegisterUpdate{{Offset: 0x24, Width: 4, Value: 1, Mask: 1}}},
	}
	if err := Validate(set); err == nil {
		t.Fatal("contradictory overlapping updates at one deadline should be rejected")
	}
}

func reviewerInferenceTrace(triggerWidth uint8, triggerOffset uint32, statusWidth uint8, statusOffset uint32) *mmio.TraceResult {
	return &mmio.TraceResult{
		BDF: "0000:03:00.0", BARIndex: 0, BARSize: 0x1000,
		Records: []mmio.AccessRecord{
			{Type: mmio.AccessWrite, Width: triggerWidth, Offset: triggerOffset, Value: 1},
			{Type: mmio.AccessRead, Width: statusWidth, Offset: statusOffset, Value: 0, Timestamp: time.Microsecond},
			{Type: mmio.AccessRead, Width: statusWidth, Offset: statusOffset, Value: 1, Timestamp: 2 * time.Microsecond},
			{Type: mmio.AccessWrite, Width: triggerWidth, Offset: triggerOffset, Value: 1, Timestamp: 3 * time.Microsecond},
			{Type: mmio.AccessRead, Width: statusWidth, Offset: statusOffset, Value: 0, Timestamp: 4 * time.Microsecond},
			{Type: mmio.AccessRead, Width: statusWidth, Offset: statusOffset, Value: 1, Timestamp: 5 * time.Microsecond},
		},
	}
}

func TestInferFiltersRulesUnsupportedByRTL(t *testing.T) {
	cases := []struct {
		name          string
		triggerWidth  uint8
		triggerOffset uint32
		statusWidth   uint8
		statusOffset  uint32
	}{
		{name: "byte trigger", triggerWidth: 1, triggerOffset: 0x20, statusWidth: 4, statusOffset: 0x24},
		{name: "word trigger", triggerWidth: 2, triggerOffset: 0x20, statusWidth: 4, statusOffset: 0x24},
		{name: "qword trigger", triggerWidth: 8, triggerOffset: 0x20, statusWidth: 4, statusOffset: 0x24},
		{name: "unaligned trigger", triggerWidth: 4, triggerOffset: 0x22, statusWidth: 4, statusOffset: 0x24},
		{name: "byte status", triggerWidth: 4, triggerOffset: 0x20, statusWidth: 1, statusOffset: 0x24},
		{name: "word status", triggerWidth: 4, triggerOffset: 0x20, statusWidth: 2, statusOffset: 0x24},
		{name: "qword status", triggerWidth: 4, triggerOffset: 0x20, statusWidth: 8, statusOffset: 0x28},
		{name: "unaligned status", triggerWidth: 4, triggerOffset: 0x20, statusWidth: 4, statusOffset: 0x26},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			set, err := Infer(reviewerInferenceTrace(tc.triggerWidth, tc.triggerOffset, tc.statusWidth, tc.statusOffset))
			if err != nil {
				t.Fatalf("Infer should filter unsupported observation, not fail: %v", err)
			}
			if len(set.Rules) != 0 {
				t.Fatalf("inferred unsupported rules = %+v", set.Rules)
			}
		})
	}
}

func TestReplaySuppressesTriggerAtDelayedEventDeadline(t *testing.T) {
	set := validReviewerRuleSet()
	set.Rules = []Rule{
		{
			ID: "arm", State: "idle", Access: AccessWrite, Width: 4,
			Offset: 0x20, Value: 1, ValueMask: math.MaxUint32,
			DelayedEvents: []DelayedEvent{{
				DelayCycles: 5,
				Updates:     []RegisterUpdate{{Offset: 0x24, Width: 4, Value: 1, Mask: math.MaxUint32}},
				NextState:   "armed",
			}},
			Confidence: 1, Provenance: []string{"test"},
		},
		{
			ID: "fire", State: "armed", Access: AccessWrite, Width: 4,
			Offset: 0x28, Value: 1, ValueMask: math.MaxUint32, NextState: "fired",
			Confidence: 1, Provenance: []string{"test"},
		},
	}
	trace := &mmio.TraceResult{
		BARIndex: 0, BARSize: 0x1000, Duration: 50 * time.Nanosecond,
		Records: []mmio.AccessRecord{
			{Type: mmio.AccessWrite, Width: 4, Offset: 0x20, Value: 1},
			{Type: mmio.AccessWrite, Width: 4, Offset: 0x28, Value: 1, Timestamp: 50 * time.Nanosecond},
		},
	}
	result, err := Replay(set, trace)
	if err != nil {
		t.Fatalf("Replay: %v", err)
	}
	if result.TerminalState != "armed" {
		t.Fatalf("terminal state = %q, want delayed event state armed", result.TerminalState)
	}
	if len(result.MatchedRules) != 1 || result.MatchedRules[0] != "arm" {
		t.Fatalf("same-cycle trigger was not suppressed: matched %v", result.MatchedRules)
	}
	if got := result.Registers[0x24]; got != 1 {
		t.Fatalf("deadline update = %#x, want 1", got)
	}
}

func TestReplayRejectsOverflowingTimestampCycleDelta(t *testing.T) {
	set := validReviewerRuleSet()
	trace := &mmio.TraceResult{
		BARIndex: 0, BARSize: 0x1000,
		Records: []mmio.AccessRecord{{
			Type: mmio.AccessRead, Width: 4, Offset: 0x24,
			Timestamp: 147573952590 * time.Nanosecond,
		}},
	}
	if _, err := Replay(set, trace); err == nil {
		t.Fatal("overflowing timestamp conversion should reject instead of wrapping to an immediate cycle")
	}
}

func TestInferDoesNotWrapHugeDelayIntoImmediateRule(t *testing.T) {
	delta := 147573952590 * time.Nanosecond
	trace := &mmio.TraceResult{
		BDF: "0000:03:00.0", BARIndex: 0, BARSize: 0x1000,
		Records: []mmio.AccessRecord{
			{Type: mmio.AccessWrite, Width: 4, Offset: 0x20, Value: 1},
			{Type: mmio.AccessRead, Width: 4, Offset: 0x24, Value: 0, Timestamp: time.Nanosecond},
			{Type: mmio.AccessRead, Width: 4, Offset: 0x24, Value: 1, Timestamp: delta},
			{Type: mmio.AccessWrite, Width: 4, Offset: 0x20, Value: 1, Timestamp: delta + time.Nanosecond},
			{Type: mmio.AccessRead, Width: 4, Offset: 0x24, Value: 0, Timestamp: delta + 2*time.Nanosecond},
			{Type: mmio.AccessRead, Width: 4, Offset: 0x24, Value: 1, Timestamp: 2*delta + time.Nanosecond},
		},
	}
	set, err := Infer(trace)
	if err != nil {
		t.Fatalf("Infer: %v", err)
	}
	if len(set.Rules) != 0 {
		t.Fatalf("huge delay wrapped into inferred rule: %+v", set.Rules)
	}
}
