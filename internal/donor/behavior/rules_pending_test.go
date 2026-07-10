package behavior

import (
	"math"
	"testing"
	"time"

	"github.com/sercanarga/pcileechgen/internal/donor/mmio"
)

func TestReplayRetriggerKeepsFirstPendingDeadline(t *testing.T) {
	set := validReviewerRuleSet()
	set.Rules = []Rule{{
		ID: "queue", State: "idle", Access: AccessWrite, Width: 4,
		Offset: 0x20, Value: 1, ValueMask: math.MaxUint32,
		DelayedEvents: []DelayedEvent{{
			DelayCycles: 5,
			Updates:     []RegisterUpdate{{Offset: 0x24, Width: 4, Value: 1, Mask: math.MaxUint32}},
		}},
		Confidence: 1, Provenance: []string{"test"},
	}}
	trace := &mmio.TraceResult{
		BARIndex: 0, BARSize: 0x1000, Duration: 70 * time.Nanosecond,
		Records: []mmio.AccessRecord{
			{Type: mmio.AccessWrite, Width: 4, Offset: 0x20, Value: 1},
			{Type: mmio.AccessWrite, Width: 4, Offset: 0x20, Value: 1, Timestamp: 20 * time.Nanosecond},
			{Type: mmio.AccessRead, Width: 4, Offset: 0x24, Value: 1, Timestamp: 50 * time.Nanosecond},
			{Type: mmio.AccessWrite, Width: 4, Offset: 0x20, Value: 1, Timestamp: 70 * time.Nanosecond},
		},
	}
	result, err := Replay(set, trace)
	if err != nil {
		t.Fatalf("Replay: %v", err)
	}
	if len(result.ReadMismatches) != 0 {
		t.Fatalf("first deadline was restarted or lost: %+v", result.ReadMismatches)
	}
	if len(result.MatchedRules) != 3 {
		t.Fatalf("matched rules = %v, want initial, retrigger, and post-deadline trigger", result.MatchedRules)
	}
}

func simultaneousReviewerRules() *RuleSet {
	set := validReviewerRuleSet()
	set.Rules = []Rule{
		{
			ID: "low", State: "idle", Access: AccessWrite, Width: 4,
			Offset: 0x20, Value: 1, ValueMask: math.MaxUint32,
			DelayedEvents: []DelayedEvent{{
				DelayCycles: 5,
				Updates:     []RegisterUpdate{{Offset: 0x24, Width: 4, Value: 0x05, Mask: 0x0f}},
				NextState:   "first",
			}},
			Confidence: 1, Provenance: []string{"test"},
		},
		{
			ID: "high", State: "idle", Access: AccessWrite, Width: 4,
			Offset: 0x28, Value: 1, ValueMask: math.MaxUint32,
			DelayedEvents: []DelayedEvent{{
				DelayCycles: 4,
				Updates:     []RegisterUpdate{{Offset: 0x24, Width: 4, Value: 0xa0, Mask: 0xf0}},
				NextState:   "second",
			}},
			Confidence: 1, Provenance: []string{"test"},
		},
	}
	return set
}

func TestReplayComposesCrossRuleEventsDueTogether(t *testing.T) {
	set := simultaneousReviewerRules()
	if err := Validate(set); err != nil {
		t.Fatalf("disjoint cross-rule events should validate: %v", err)
	}
	trace := &mmio.TraceResult{
		BARIndex: 0, BARSize: 0x1000, Duration: 50 * time.Nanosecond,
		Records: []mmio.AccessRecord{
			{Type: mmio.AccessWrite, Width: 4, Offset: 0x20, Value: 1},
			{Type: mmio.AccessWrite, Width: 4, Offset: 0x28, Value: 1, Timestamp: 10 * time.Nanosecond},
			{Type: mmio.AccessRead, Width: 4, Offset: 0x24, Value: 0xa5, Timestamp: 50 * time.Nanosecond},
		},
	}
	result, err := Replay(set, trace)
	if err != nil {
		t.Fatalf("Replay: %v", err)
	}
	if len(result.ReadMismatches) != 0 || result.Registers[0x24] != 0xa5 {
		t.Fatalf("cross-rule composed result = %+v", result)
	}
	if result.TerminalState != "second" {
		t.Fatalf("terminal state = %q, want stable later-event state second", result.TerminalState)
	}
}

func TestValidateRejectsContradictoryCrossRuleDelayedUpdates(t *testing.T) {
	set := simultaneousReviewerRules()
	set.Rules[1].DelayedEvents[0].Updates[0] = RegisterUpdate{Offset: 0x24, Width: 4, Value: 0, Mask: 1}
	if err := Validate(set); err == nil {
		t.Fatal("potentially simultaneous contradictory cross-rule updates should be rejected")
	}
}

func TestInferDoesNotCrossTraceSessionBoundary(t *testing.T) {
	first := &mmio.TraceResult{
		BDF: "0000:03:00.0", BARIndex: 0, BARSize: 0x1000,
		Records: []mmio.AccessRecord{{
			Type: mmio.AccessWrite, Width: 4, Offset: 0x20, Value: 1,
		}},
	}
	second := &mmio.TraceResult{
		BDF: "0000:03:00.0", BARIndex: 0, BARSize: 0x1000,
		Records: []mmio.AccessRecord{
			{Type: mmio.AccessRead, Width: 4, Offset: 0x24, Value: 0},
			{Type: mmio.AccessRead, Width: 4, Offset: 0x24, Value: 1, Timestamp: time.Microsecond},
		},
	}
	set, err := Infer(first, second)
	if err != nil {
		t.Fatalf("Infer: %v", err)
	}
	if len(set.Rules) != 0 {
		t.Fatalf("inference crossed capture session boundary: %+v", set.Rules)
	}
}

func TestEngineAndReplayComposeTriggerWriteBeforeSameRegisterUpdate(t *testing.T) {
	set := validReviewerRuleSet()
	set.Rules = []Rule{{
		ID: "compose-self", State: "idle", Access: AccessWrite, Width: 4,
		Offset: 0x20, Value: 0x12345678, ValueMask: math.MaxUint32,
		Updates: []RegisterUpdate{{
			Offset: 0x20, Width: 4, Value: 0x0000ab00, Mask: 0x0000ff00,
		}},
		Confidence: 1, Provenance: []string{"test"},
	}}
	engine, err := NewEngine(set)
	if err != nil {
		t.Fatalf("NewEngine: %v", err)
	}
	record := mmio.AccessRecord{Type: mmio.AccessWrite, Width: 4, Offset: 0x20, Value: 0x12345678}
	if err := engine.Apply(record); err != nil {
		t.Fatalf("Apply: %v", err)
	}
	if got, ok := engine.Register(0x20); !ok || got != 0x1234ab78 {
		t.Fatalf("engine composed value = %#x, present=%v, want 0x1234ab78", got, ok)
	}
	result, err := Replay(set, &mmio.TraceResult{
		BARIndex: 0, BARSize: 0x1000, Records: []mmio.AccessRecord{record},
	})
	if err != nil {
		t.Fatalf("Replay: %v", err)
	}
	if got := result.Registers[0x20]; got != 0x1234ab78 {
		t.Fatalf("replay composed value = %#x, want 0x1234ab78", got)
	}
}

func TestEngineAndReplayApplyWritePolicyBeforeSameRegisterRuleUpdate(t *testing.T) {
	set := validReviewerRuleSet()
	set.InitialRegisters = []RegisterValue{{
		Offset: 0x20, Width: 4, Value: 0xa5f0cc33,
		WritePolicy: &RegisterWritePolicy{RWMask: 0x0000000f, W1CMask: 0x000000f0},
	}}
	set.Rules = []Rule{{
		ID: "compose-policy", State: "idle", Access: AccessWrite, Width: 4,
		Offset: 0x20, Value: 0x123456a5, ValueMask: math.MaxUint32,
		Updates: []RegisterUpdate{{
			Offset: 0x20, Width: 4, Value: 0x00005a00, Mask: 0x0000ff00,
		}},
		Confidence: 1, Provenance: []string{"test"},
	}}
	record := mmio.AccessRecord{
		Type: mmio.AccessWrite, Width: 4, Offset: 0x20, Value: 0x123456a5,
	}
	engine, err := NewEngine(set)
	if err != nil {
		t.Fatalf("NewEngine: %v", err)
	}
	if err := engine.Apply(record); err != nil {
		t.Fatalf("Apply: %v", err)
	}
	if got, ok := engine.Register(0x20); !ok || got != 0xa5f05a15 {
		t.Fatalf("engine policy composition = %#x, present=%v, want 0xa5f05a15", got, ok)
	}
	result, err := Replay(set, &mmio.TraceResult{
		BARIndex: 0, BARSize: 0x1000, Records: []mmio.AccessRecord{record},
	})
	if err != nil {
		t.Fatalf("Replay: %v", err)
	}
	if got := result.Registers[0x20]; got != 0xa5f05a15 {
		t.Fatalf("replay policy composition = %#x, want 0xa5f05a15", got)
	}
}

func TestReplayMasksExtraneousBitsPerSimultaneousDelayedContribution(t *testing.T) {
	set := validReviewerRuleSet()
	set.InitialRegisters = []RegisterValue{{Offset: 0x24, Width: 4, Value: 3}}
	set.Rules = []Rule{
		{
			ID: "clear-bit0", State: "idle", Access: AccessWrite, Width: 4,
			Offset: 0x20, Value: 1, ValueMask: math.MaxUint32,
			DelayedEvents: []DelayedEvent{{
				DelayCycles: 5,
				Updates:     []RegisterUpdate{{Offset: 0x24, Width: 4, Value: 2, Mask: 1}},
			}},
			Confidence: 1, Provenance: []string{"test"},
		},
		{
			ID: "clear-bit1", State: "idle", Access: AccessWrite, Width: 4,
			Offset: 0x28, Value: 1, ValueMask: math.MaxUint32,
			DelayedEvents: []DelayedEvent{{
				DelayCycles: 4,
				Updates:     []RegisterUpdate{{Offset: 0x24, Width: 4, Value: 0, Mask: 2}},
			}},
			Confidence: 1, Provenance: []string{"test"},
		},
	}
	result, err := Replay(set, &mmio.TraceResult{
		BARIndex: 0, BARSize: 0x1000,
		Records: []mmio.AccessRecord{
			{Type: mmio.AccessWrite, Width: 4, Offset: 0x20, Value: 1},
			{Type: mmio.AccessWrite, Width: 4, Offset: 0x28, Value: 1, Timestamp: 10 * time.Nanosecond},
			{Type: mmio.AccessRead, Width: 4, Offset: 0x24, Value: 0, Timestamp: 50 * time.Nanosecond},
		},
	})
	if err != nil {
		t.Fatalf("Replay: %v", err)
	}
	if len(result.ReadMismatches) != 0 || result.Registers[0x24] != 0 {
		t.Fatalf("extraneous contribution bits leaked: %+v", result)
	}
}
