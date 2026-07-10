package svgen

import (
	"math"
	"strings"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/donor/behavior"
	"github.com/sercanarga/pcileechgen/internal/donor/mmio"
	"github.com/sercanarga/pcileechgen/internal/firmware/barmodel"
)

func reviewerCompilerModel() *barmodel.BARModel {
	return &barmodel.BARModel{
		Size: 0x1000,
		Registers: []barmodel.BARRegister{
			{Offset: 0x24, Width: 4, Reset: 0xdeadbeef, Name: "STATUS0"},
			{Offset: 0x28, Width: 4, Reset: 0xcafebabe, Name: "STATUS1"},
		},
	}
}

func reviewerCompilerRules() *behavior.RuleSet {
	return &behavior.RuleSet{
		Version: behavior.RuleSchemaVersion, BARSize: 0x1000,
		ClockHz: 100_000_000, InitialState: "idle",
		Rules: []behavior.Rule{{
			ID: "compose", State: "idle", Access: behavior.AccessWrite, Width: 4,
			Offset: 0x20, Value: 1, ValueMask: math.MaxUint32,
			Updates: []behavior.RegisterUpdate{
				{Offset: 0x24, Width: 4, Value: 0x00000005, Mask: 0x0000000f},
				{Offset: 0x24, Width: 4, Value: 0x000000a0, Mask: 0x000000f0},
			},
			DelayedEvents: []behavior.DelayedEvent{
				{DelayCycles: 10, Updates: []behavior.RegisterUpdate{
					{Offset: 0x24, Width: 4, Value: 0x00000300, Mask: 0x00000f00},
				}},
				{DelayCycles: 10, Updates: []behavior.RegisterUpdate{
					{Offset: 0x24, Width: 4, Value: 0x0000c000, Mask: 0x0000f000},
				}},
			},
			Confidence: 1, Provenance: []string{"test"},
		}},
	}
}

func TestCompileBehaviorRulesCoalescesDisjointRegisterUpdates(t *testing.T) {
	compiled, err := CompileBehaviorRules(reviewerCompilerRules(), reviewerCompilerModel())
	if err != nil {
		t.Fatalf("CompileBehaviorRules: %v", err)
	}
	if len(compiled.Rules) != 1 || len(compiled.Rules[0].Updates) != 1 {
		t.Fatalf("compiled immediate updates = %+v, want one coalesced update", compiled.Rules)
	}
	immediate := compiled.Rules[0].Updates[0]
	if immediate.Offset != 0x24 || immediate.Mask != 0xff || immediate.Value != 0xa5 {
		t.Fatalf("coalesced immediate update = %+v, want offset 0x24 mask 0xff value 0xa5", immediate)
	}
	if len(compiled.Rules[0].DelayedEvents) != 1 || len(compiled.Rules[0].DelayedEvents[0].Updates) != 1 {
		t.Fatalf("compiled delayed updates = %+v, want one coalesced update", compiled.Rules[0].DelayedEvents)
	}
	delayed := compiled.Rules[0].DelayedEvents[0].Updates[0]
	if delayed.Offset != 0x24 || delayed.Mask != 0xff00 || delayed.Value != 0xc300 {
		t.Fatalf("coalesced delayed update = %+v, want offset 0x24 mask 0xff00 value 0xc300", delayed)
	}
}

func TestCompileBehaviorRulesRescalesDelayWithCeilingAndRejectsOverflow(t *testing.T) {
	set := reviewerCompilerRules()
	compiled, err := CompileBehaviorRules(set, reviewerCompilerModel())
	if err != nil {
		t.Fatalf("CompileBehaviorRules: %v", err)
	}
	if got := compiled.Rules[0].DelayedEvents[0].DelayCycles; got != 13 {
		t.Fatalf("125MHz delay = %d, want ceil(10*125/100)=13", got)
	}

	set.ClockHz = 1
	set.Rules[0].DelayedEvents[0].DelayCycles = behavior.MaxDelayCycles
	if _, err := CompileBehaviorRules(set, reviewerCompilerModel()); err == nil {
		t.Fatal("expected rescaled delay above hardware bound to be rejected")
	}
}

func TestEngineReplayAndCompiledRTLShareInitialRegisterSeeds(t *testing.T) {
	set := reviewerCompilerRules()
	set.InitialRegisters = []behavior.RegisterValue{{Offset: 0x24, Width: 4, Value: 0x11223344}}
	set.Rules[0].Updates = []behavior.RegisterUpdate{
		{Offset: 0x24, Width: 4, Value: 1, Mask: 1},
		{Offset: 0x28, Width: 4, Value: 1, Mask: 1},
	}
	set.Rules[0].DelayedEvents = nil

	compiled, err := CompileBehaviorRules(set, reviewerCompilerModel())
	if err != nil {
		t.Fatalf("CompileBehaviorRules: %v", err)
	}
	resets := make(map[uint32]uint32)
	for _, register := range compiled.Registers {
		resets[register.Offset] = register.Reset
	}
	if resets[0x24] != 0x11223344 || resets[0x28] != 0 {
		t.Fatalf("compiled resets = %#v, want explicit 0x11223344 and implicit zero", resets)
	}

	engine, err := behavior.NewEngine(set)
	if err != nil {
		t.Fatalf("NewEngine: %v", err)
	}
	if got, ok := engine.Register(0x24); !ok || got != 0x11223344 {
		t.Fatalf("explicit engine seed = %#x, present=%v, want 0x11223344", got, ok)
	}
	if got, _ := engine.Register(0x28); got != 0 {
		t.Fatalf("implicit engine seed = %#x, want zero", got)
	}

	trace := &mmio.TraceResult{
		BARIndex: 0, BARSize: 0x1000,
		Records: []mmio.AccessRecord{
			{Type: mmio.AccessRead, Width: 4, Offset: 0x24, Value: 0x11223344},
			{Type: mmio.AccessRead, Width: 4, Offset: 0x28, Value: 0},
		},
	}
	result, err := behavior.Replay(set, trace)
	if err != nil {
		t.Fatalf("Replay: %v", err)
	}
	if len(result.ReadMismatches) != 0 {
		t.Fatalf("initial reads disagree with compiled reset seeds: %+v", result.ReadMismatches)
	}
}

func TestGenerateBehaviorRTLGuardsTriggersOnEventDeadline(t *testing.T) {
	cfg := testConfig()
	cfg.BARModel = reviewerCompilerModel()
	cfg.BehaviorRules = generatedRuleSet(1)
	generated, err := GenerateBarImplDeviceSV(cfg)
	if err != nil {
		t.Fatalf("GenerateBarImplDeviceSV: %v", err)
	}
	if !strings.Contains(generated, "behavior_event_due") {
		t.Fatal("generated RTL has no event deadline signal")
	}
	if !strings.Contains(generated, "!behavior_event_due") {
		t.Fatal("generated RTL does not suppress triggers on an event deadline")
	}
}

func TestGenerateBehaviorRTLPreservesFirstPendingDeadlineAndComposesDueUpdates(t *testing.T) {
	set := &behavior.RuleSet{
		Version: behavior.RuleSchemaVersion, BARSize: 0x1000,
		ClockHz: 100_000_000, InitialState: "idle",
		Rules: []behavior.Rule{
			{
				ID: "low", State: "idle", Access: behavior.AccessWrite, Width: 4,
				Offset: 0x20, Value: 1, ValueMask: math.MaxUint32,
				DelayedEvents: []behavior.DelayedEvent{{
					DelayCycles: 5,
					Updates: []behavior.RegisterUpdate{{Offset: 0x24, Width: 4, Value: 5, Mask: 0x0f}},
					NextState: "first",
				}},
				Confidence: 1, Provenance: []string{"test"},
			},
			{
				ID: "high", State: "idle", Access: behavior.AccessWrite, Width: 4,
				Offset: 0x28, Value: 1, ValueMask: math.MaxUint32,
				DelayedEvents: []behavior.DelayedEvent{{
					DelayCycles: 4,
					Updates: []behavior.RegisterUpdate{{Offset: 0x24, Width: 4, Value: 0xa0, Mask: 0xf0}},
					NextState: "second",
				}},
				Confidence: 1, Provenance: []string{"test"},
			},
		},
	}
	compiled, err := CompileBehaviorRules(set, reviewerCompilerModel())
	if err != nil {
		t.Fatalf("CompileBehaviorRules: %v", err)
	}
	if len(compiled.DueRegisters) != 1 || compiled.DueRegisters[0].Offset != 0x24 ||
		len(compiled.DueRegisters[0].Contributions) != 2 {
		t.Fatalf("compiled due-register composition = %+v", compiled.DueRegisters)
	}
	secondState := -1
	for index, state := range compiled.States {
		if state == "second" {
			secondState = index
		}
	}
	if len(compiled.AllEvents) != 2 || compiled.AllEvents[1].NextState != secondState {
		t.Fatalf("compiled static event order = %+v, states=%v", compiled.AllEvents, compiled.States)
	}
	cfg := testConfig()
	cfg.DeviceClass = ""
	cfg.ClassCode = 0xff0000
	cfg.BARModel = reviewerCompilerModel()
	cfg.BehaviorRules = set
	generated, err := GenerateBarImplDeviceSV(cfg)
	if err != nil {
		t.Fatalf("GenerateBarImplDeviceSV: %v", err)
	}
	if !strings.Contains(generated, "if (!rule_delay_pending_0)") {
		t.Fatal("generated RTL rearms an already pending delayed event")
	}
	if !strings.Contains(generated, "behavior_due_mask_00000024") ||
		!strings.Contains(generated, "behavior_due_value_00000024") {
		t.Fatal("generated RTL has no global due-update composition for register 0x24")
	}
	assignment := "if (behavior_due_mask_00000024 != 32'h00000000) reg_0x00000024 <="
	if strings.Count(generated, assignment) != 1 {
		t.Fatalf("global delayed register assignment count = %d, want 1", strings.Count(generated, assignment))
	}
}

func TestValidateBehaviorRuleOwnershipRejectsClassOwnedInitialRegister(t *testing.T) {
	set := &behavior.RuleSet{
		Version: behavior.RuleSchemaVersion, BARSize: 0x1000,
		ClockHz: 100_000_000, InitialState: "idle",
		InitialRegisters: []behavior.RegisterValue{{
			Offset: 0x24, Width: 4, Value: 1,
		}},
	}
	if err := ValidateBehaviorRuleOwnership(set, "audio"); err == nil {
		t.Fatal("audio-owned register seed should be rejected")
	}
}

func TestCompileAndGenerateSameRegisterUpdateUsesTriggerWriteAsBase(t *testing.T) {
	set := &behavior.RuleSet{
		Version: behavior.RuleSchemaVersion, BARSize: 0x1000,
		ClockHz: 100_000_000, InitialState: "idle",
		Rules: []behavior.Rule{{
			ID: "compose-self", State: "idle", Access: behavior.AccessWrite, Width: 4,
			Offset: 0x20, Value: 0x12345678, ValueMask: math.MaxUint32,
			Updates: []behavior.RegisterUpdate{{
				Offset: 0x20, Width: 4, Value: 0x0000ab00, Mask: 0x0000ff00,
			}},
			Confidence: 1, Provenance: []string{"test"},
		}},
	}
	model := &barmodel.BARModel{
		Size: 0x1000,
		Registers: []barmodel.BARRegister{{
			Offset: 0x20, Width: 4, RWMask: math.MaxUint32, Name: "CONTROL",
		}},
	}
	compiled, err := CompileBehaviorRules(set, model)
	if err != nil {
		t.Fatalf("CompileBehaviorRules: %v", err)
	}
	if len(compiled.Rules) != 1 || len(compiled.Rules[0].Updates) != 1 ||
		!compiled.Rules[0].Updates[0].ComposeTriggerWrite {
		t.Fatalf("compiled same-register update = %+v", compiled.Rules)
	}
	cfg := testConfig()
	cfg.DeviceClass = ""
	cfg.ClassCode = 0xff0000
	cfg.BARModel = model
	cfg.BehaviorRules = set
	generated, err := GenerateBarImplDeviceSV(cfg)
	if err != nil {
		t.Fatalf("GenerateBarImplDeviceSV: %v", err)
	}
	if !strings.Contains(generated, "wr_data & ~32'h0000FF00") {
		t.Fatal("generated same-register update does not compose from trigger write data")
	}
}

func TestWritePolicyCompositionMatchesEngineCompilerAndRTL(t *testing.T) {
	set := &behavior.RuleSet{
		Version: behavior.RuleSchemaVersion, BARSize: 0x1000,
		ClockHz: 100_000_000, InitialState: "idle",
		InitialRegisters: []behavior.RegisterValue{{
			Offset: 0x20, Width: 4, Value: 0xa5f0cc33,
			WritePolicy: &behavior.RegisterWritePolicy{RWMask: 0x0000000f, W1CMask: 0x000000f0},
		}},
		Rules: []behavior.Rule{{
			ID: "compose-policy", State: "idle", Access: behavior.AccessWrite, Width: 4,
			Offset: 0x20, Value: 0x123456a5, ValueMask: math.MaxUint32,
			Updates: []behavior.RegisterUpdate{{
				Offset: 0x20, Width: 4, Value: 0x00005a00, Mask: 0x0000ff00,
			}},
			Confidence: 1, Provenance: []string{"test"},
		}},
	}
	model := &barmodel.BARModel{
		Size: 0x1000,
		Registers: []barmodel.BARRegister{{
			Offset: 0x20, Width: 4, Reset: 0xa5f0cc33, RWMask: 0x0000000f,
			W1CMask: 0x000000f0, Name: "CONTROL",
		}},
	}
	compiled, err := CompileBehaviorRules(set, model)
	if err != nil {
		t.Fatalf("CompileBehaviorRules: %v", err)
	}
	update := compiled.Rules[0].Updates[0]
	if !update.ComposeTriggerWrite || update.TriggerRWMask != 0x0000000f ||
		update.TriggerW1CMask != 0x000000f0 {
		t.Fatalf("compiled write policy = %+v", update)
	}
	cfg := testConfig()
	cfg.DeviceClass = ""
	cfg.ClassCode = 0xff0000
	cfg.BARModel = model
	cfg.BehaviorRules = set
	generated, err := GenerateBarImplDeviceSV(cfg)
	if err != nil {
		t.Fatalf("GenerateBarImplDeviceSV: %v", err)
	}
	if !strings.Contains(generated, "behavior_wr_be_mask") ||
		!strings.Contains(generated, "32'h0000000F") ||
		!strings.Contains(generated, "32'h000000F0") {
		t.Fatal("generated RTL does not apply canonical RW/W1C write policy before rule overlay")
	}
}

func TestCompileAndGenerateDueValuesMaskEachContribution(t *testing.T) {
	set := &behavior.RuleSet{
		Version: behavior.RuleSchemaVersion, BARSize: 0x1000,
		ClockHz: 100_000_000, InitialState: "idle",
		InitialRegisters: []behavior.RegisterValue{{Offset: 0x24, Width: 4, Value: 3}},
		Rules: []behavior.Rule{
			{
				ID: "clear-bit0", State: "idle", Access: behavior.AccessWrite, Width: 4,
				Offset: 0x20, Value: 1, ValueMask: math.MaxUint32,
				DelayedEvents: []behavior.DelayedEvent{{
					DelayCycles: 5,
					Updates: []behavior.RegisterUpdate{{Offset: 0x24, Width: 4, Value: 2, Mask: 1}},
				}},
				Confidence: 1, Provenance: []string{"test"},
			},
			{
				ID: "clear-bit1", State: "idle", Access: behavior.AccessWrite, Width: 4,
				Offset: 0x28, Value: 1, ValueMask: math.MaxUint32,
				DelayedEvents: []behavior.DelayedEvent{{
					DelayCycles: 4,
					Updates: []behavior.RegisterUpdate{{Offset: 0x24, Width: 4, Value: 0, Mask: 2}},
				}},
				Confidence: 1, Provenance: []string{"test"},
			},
		},
	}
	model := &barmodel.BARModel{
		Size: 0x1000,
		Registers: []barmodel.BARRegister{{
			Offset: 0x24, Width: 4, Reset: 3, Name: "STATUS",
		}},
	}
	compiled, err := CompileBehaviorRules(set, model)
	if err != nil {
		t.Fatalf("CompileBehaviorRules: %v", err)
	}
	if len(compiled.DueRegisters) != 1 || len(compiled.DueRegisters[0].Contributions) != 2 {
		t.Fatalf("compiled contributions = %+v", compiled.DueRegisters)
	}
	for _, contribution := range compiled.DueRegisters[0].Contributions {
		if contribution.Value != 0 {
			t.Fatalf("compiled contribution leaked value outside mask: %+v", contribution)
		}
	}
	cfg := testConfig()
	cfg.DeviceClass = ""
	cfg.ClassCode = 0xff0000
	cfg.BARModel = model
	cfg.BehaviorRules = set
	generated, err := GenerateBarImplDeviceSV(cfg)
	if err != nil {
		t.Fatalf("GenerateBarImplDeviceSV: %v", err)
	}
	for _, term := range []string{
		"({32{rule_event_due_0}} & 32'h00000000)",
		"({32{rule_event_due_1}} & 32'h00000000)",
	} {
		if !strings.Contains(generated, term) {
			t.Fatalf("generated due value is not contribution-masked: missing %q", term)
		}
	}
}
