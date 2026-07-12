package behavior

import (
	"bytes"
	"encoding/json"
	"math"
	"testing"
	"time"

	"github.com/sercanarga/pcileechgen/internal/donor/mmio"
)

func transitionTrace(polls int) *mmio.TraceResult {
	records := []mmio.AccessRecord{
		{BDF: "0000:03:00.0", BARIndex: 2, Address: 0xf7800020, Offset: 0x20, Width: 4, Type: mmio.AccessWrite, Value: 1, Timestamp: 0},
	}
	for i := range polls {
		records = append(records, mmio.AccessRecord{
			BDF: "0000:03:00.0", BARIndex: 2, Address: 0xf7800024, Offset: 0x24, Width: 4,
			Type: mmio.AccessRead, Value: 0, Timestamp: time.Duration(i+1) * 10 * time.Microsecond,
		})
	}
	records = append(records, mmio.AccessRecord{
		BDF: "0000:03:00.0", BARIndex: 2, Address: 0xf7800024, Offset: 0x24, Width: 4,
		Type: mmio.AccessRead, Value: 1, Timestamp: time.Duration(polls+1) * 10 * time.Microsecond,
	})
	return &mmio.TraceResult{
		BDF: "0000:03:00.0", BARIndex: 2, BARSize: 0x4000, Records: records,
	}
}

func TestInfer_DeterministicRulesCarryConfidenceAndProvenance(t *testing.T) {
	first, err := Infer(transitionTrace(3), transitionTrace(5))
	if err != nil {
		t.Fatalf("Infer first: %v", err)
	}
	second, err := Infer(transitionTrace(3), transitionTrace(5))
	if err != nil {
		t.Fatalf("Infer second: %v", err)
	}
	if err := Validate(first); err != nil {
		t.Fatalf("inferred rules are invalid: %v", err)
	}
	if len(first.Rules) == 0 {
		t.Fatal("expected a repeated write-triggered status transition to produce a rule")
	}

	firstJSON, err := json.Marshal(first)
	if err != nil {
		t.Fatalf("marshal first rules: %v", err)
	}
	secondJSON, err := json.Marshal(second)
	if err != nil {
		t.Fatalf("marshal second rules: %v", err)
	}
	if !bytes.Equal(firstJSON, secondJSON) {
		t.Fatalf("identical traces produced nondeterministic JSON:\nfirst:  %s\nsecond: %s", firstJSON, secondJSON)
	}

	var document struct {
		Version int `json:"version"`
		Rules   []struct {
			Confidence float64         `json:"confidence"`
			Provenance json.RawMessage `json:"provenance"`
		} `json:"rules"`
	}
	if err := json.Unmarshal(firstJSON, &document); err != nil {
		t.Fatalf("decode rule schema: %v", err)
	}
	if document.Version != RuleSchemaVersion {
		t.Fatalf("schema version = %d, want %d", document.Version, RuleSchemaVersion)
	}
	for i, rule := range document.Rules {
		if rule.Confidence <= 0 || rule.Confidence > 1 {
			t.Errorf("rule %d confidence = %v, want (0,1]", i, rule.Confidence)
		}
		if len(rule.Provenance) == 0 || bytes.Equal(rule.Provenance, []byte("null")) || bytes.Equal(rule.Provenance, []byte(`""`)) {
			t.Errorf("rule %d has no provenance: %s", i, rule.Provenance)
		}
	}
	var schema map[string]any
	if err := json.Unmarshal(firstJSON, &schema); err != nil {
		t.Fatalf("decode canonical schema: %v", err)
	}
	for _, key := range []string{"version", "bar_index", "bar_size", "clock_hz", "initial_state", "rules"} {
		if _, ok := schema[key]; !ok {
			t.Errorf("canonical rule JSON missing snake_case field %q: %v", key, schema)
		}
	}
	ruleObjects, ok := schema["rules"].([]any)
	if !ok || len(ruleObjects) == 0 {
		t.Fatalf("canonical rules field = %#v", schema["rules"])
	}
	ruleObject, ok := ruleObjects[0].(map[string]any)
	if !ok {
		t.Fatalf("canonical first rule = %#v", ruleObjects[0])
	}
	for _, key := range []string{"id", "state", "access", "width", "offset", "value_mask", "confidence", "provenance"} {
		if _, ok := ruleObject[key]; !ok {
			t.Errorf("canonical rule JSON missing snake_case field %q: %v", key, ruleObject)
		}
	}
}

func TestInfer_RejectsConflictingObservedTransitions(t *testing.T) {
	readyOne := transitionTrace(3)
	readyTwo := transitionTrace(3)
	readyTwo.Records[len(readyTwo.Records)-1].Value = 2
	if _, err := Infer(readyOne, readyTwo); err == nil {
		t.Fatal("expected conflicting observations of the same trigger and status register to be rejected")
	}
}

func TestValidate_RejectsConflictingRules(t *testing.T) {
	rules := &RuleSet{
		Version:      RuleSchemaVersion,
		BARSize:      0x1000,
		ClockHz:      100_000_000,
		InitialState: "idle",
		Rules: []Rule{
			{ID: "to-a", State: "idle", Access: AccessKind("write"), Width: 4, Offset: 0x20, Value: 1, ValueMask: math.MaxUint32, NextState: "a", Confidence: 1, Provenance: []string{"test"}},
			{ID: "to-b", State: "idle", Access: AccessKind("write"), Width: 4, Offset: 0x20, Value: 1, ValueMask: math.MaxUint32, NextState: "b", Confidence: 1, Provenance: []string{"test"}},
		},
	}
	if err := Validate(rules); err == nil {
		t.Fatal("expected indistinguishable triggers with conflicting effects to be rejected")
	}
}

func TestValidate_RejectsUnboundedDelay(t *testing.T) {
	rules := &RuleSet{
		Version:      RuleSchemaVersion,
		BARSize:      0x1000,
		ClockHz:      100_000_000,
		InitialState: "idle",
		Rules: []Rule{{
			ID: "unbounded", State: "idle", Access: AccessKind("write"), Width: 4,
			Offset: 0x20, Value: 1, ValueMask: math.MaxUint32,
			DelayedEvents: []DelayedEvent{{DelayCycles: MaxDelayCycles + 1, NextState: "ready"}},
			Confidence:    1, Provenance: []string{"test"},
		}},
	}
	if err := Validate(rules); err == nil {
		t.Fatal("expected an unbounded learned delay to be rejected")
	}
	rules.Rules[0].DelayedEvents[0].DelayCycles = MaxDelayCycles
	if err := Validate(rules); err != nil {
		t.Fatalf("maximum bounded delay should be valid: %v", err)
	}
}

func TestEngine_DelayedTransitionIsPollingInsensitive(t *testing.T) {
	rules := &RuleSet{
		Version:      RuleSchemaVersion,
		BARSize:      0x1000,
		ClockHz:      100_000_000,
		InitialState: "idle",
		Rules: []Rule{{
			ID: "start", State: "idle", Access: AccessKind("write"), Width: 4,
			Offset: 0x20, Value: 1, ValueMask: math.MaxUint32, NextState: "waiting",
			DelayedEvents: []DelayedEvent{{
				DelayCycles: 5,
				Updates:     []RegisterUpdate{{Offset: 0x24, Width: 4, Value: 1, Mask: math.MaxUint32}},
				NextState:   "ready",
			}},
			Confidence: 1, Provenance: []string{"test"},
		}},
	}

	run := func(t *testing.T, polls int) (*Engine, uint64) {
		t.Helper()
		engine, err := NewEngine(rules)
		if err != nil {
			t.Fatalf("NewEngine: %v", err)
		}
		if err := engine.Apply(mmio.AccessRecord{Type: mmio.AccessWrite, Width: 4, Offset: 0x20, Value: 1}); err != nil {
			t.Fatalf("apply trigger: %v", err)
		}
		if err := engine.Advance(4); err != nil {
			t.Fatalf("advance: %v", err)
		}
		for i := range polls {
			if err := engine.Apply(mmio.AccessRecord{Type: mmio.AccessRead, Width: 4, Offset: 0x24}); err != nil {
				t.Fatalf("apply poll %d: %v", i, err)
			}
		}
		if engine.State() != "waiting" {
			t.Fatalf("state before final cycle = %q, want waiting", engine.State())
		}
		if err := engine.Advance(1); err != nil {
			t.Fatalf("advance: %v", err)
		}
		value, ok := engine.Register(0x24)
		if !ok {
			t.Fatal("status register missing after delayed update")
		}
		return engine, value
	}

	withoutPolls, valueWithoutPolls := run(t, 0)
	withPolls, valueWithPolls := run(t, 100)
	if withoutPolls.State() != "ready" || withPolls.State() != "ready" {
		t.Fatalf("final states = %q and %q, want ready", withoutPolls.State(), withPolls.State())
	}
	if valueWithoutPolls != 1 || valueWithPolls != 1 {
		t.Fatalf("status values = %#x and %#x, want 1", valueWithoutPolls, valueWithPolls)
	}
}

func TestReplay_PollCountDoesNotChangeObservedTransition(t *testing.T) {
	rules := &RuleSet{
		Version: RuleSchemaVersion, BARSize: 0x1000, ClockHz: 100_000_000, InitialState: "idle",
		Rules: []Rule{{
			ID: "start", State: "idle", Access: AccessKind("write"), Width: 4,
			Offset: 0x20, Value: 1, ValueMask: math.MaxUint32, NextState: "waiting",
			DelayedEvents: []DelayedEvent{{
				DelayCycles: 5,
				Updates:     []RegisterUpdate{{Offset: 0x24, Width: 4, Value: 1, Mask: math.MaxUint32}},
				NextState:   "ready",
			}},
			Confidence: 1, Provenance: []string{"test"},
		}},
	}
	traceWithPolls := func(polls int) *mmio.TraceResult {
		records := []mmio.AccessRecord{
			{Type: mmio.AccessWrite, Width: 4, Offset: 0x20, Value: 1},
		}
		for range polls {
			records = append(records, mmio.AccessRecord{
				Type: mmio.AccessRead, Width: 4, Offset: 0x24, Value: 0, Timestamp: 10 * time.Nanosecond,
			})
		}
		records = append(records, mmio.AccessRecord{
			Type: mmio.AccessRead, Width: 4, Offset: 0x24, Value: 1, Timestamp: 50 * time.Nanosecond,
		})
		return &mmio.TraceResult{BARSize: 0x1000, Duration: 50 * time.Nanosecond, Records: records}
	}

	for _, polls := range []int{0, 1, 100} {
		result, err := Replay(rules, traceWithPolls(polls))
		if err != nil {
			t.Errorf("Replay with %d redundant polls: %v", polls, err)
			continue
		}
		if result == nil {
			t.Errorf("Replay with %d redundant polls returned nil result", polls)
			continue
		}
		if result.TerminalState != "ready" {
			t.Errorf("Replay with %d polls terminal state = %q, want ready", polls, result.TerminalState)
		}
		if got := result.Registers[0x24]; got != 1 {
			t.Errorf("Replay with %d polls status = %#x, want 1", polls, got)
		}
		if len(result.MatchedRules) != 1 || result.MatchedRules[0] != "start" {
			t.Errorf("Replay with %d polls matched rules = %v, want [start]", polls, result.MatchedRules)
		}
		if len(result.ReadMismatches) != 0 {
			t.Errorf("Replay with %d polls mismatches = %+v", polls, result.ReadMismatches)
		}
	}
}
