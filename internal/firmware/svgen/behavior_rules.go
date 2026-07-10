package svgen

import (
	"fmt"

	"github.com/sercanarga/pcileechgen/internal/donor/behavior"
	"github.com/sercanarga/pcileechgen/internal/firmware/barmodel"
)

const BehaviorRTLClockHz uint64 = 125_000_000

type CompiledBehavior struct {
	StateBits    int
	InitialState int
	States       []string
	Registers    []CompiledBehaviorRegister
	Rules        []CompiledBehaviorRule
	AllEvents    []CompiledBehaviorEvent
	DueRegisters []CompiledDueRegister
}

type CompiledBehaviorRegister struct {
	Offset  uint32
	Reset   uint32
	RWMask  uint32
	W1CMask uint32
}

type CompiledBehaviorRule struct {
	Index          int
	State          int
	NextState      int
	Offset         uint32
	Value          uint32
	ValueMask      uint32
	ByteEnableMask uint8
	Updates        []CompiledBehaviorUpdate
	DelayedEvents  []CompiledBehaviorEvent
}

type CompiledBehaviorUpdate struct {
	Offset uint32
	Value  uint32
	Mask   uint32
	ComposeTriggerWrite bool
	TriggerRWMask         uint32
	TriggerW1CMask        uint32
}
type CompiledDueRegister struct {
	Offset        uint32
	Contributions []CompiledDueContribution
}

type CompiledDueContribution struct {
	EventIndex int
	Value      uint32
	Mask       uint32
}

func ValidateBehaviorRuleOwnership(set *behavior.RuleSet, deviceClass string) error {
	if set == nil || deviceClass != "audio" {
		return nil
	}
	reserved := map[uint32]struct{}{
		0x0c: {}, 0x24: {}, 0x48: {}, 0x4c: {}, 0x58: {}, 0x5c: {}, 0x60: {}, 0x70: {}, 0x74: {},
	}
	for _, initial := range set.InitialRegisters {
		if _, exists := reserved[initial.Offset]; exists {
			return fmt.Errorf("initial register at %#x conflicts with audio device behavior", initial.Offset)
		}
	}
	for _, rule := range set.Rules {
		if _, exists := reserved[rule.Offset]; exists {
			return fmt.Errorf("rule %q trigger at %#x conflicts with audio device behavior", rule.ID, rule.Offset)
		}
		updates := append([]behavior.RegisterUpdate(nil), rule.Updates...)
		for _, event := range rule.DelayedEvents {
			updates = append(updates, event.Updates...)
		}
		for _, update := range updates {
			if _, exists := reserved[update.Offset]; exists {
				return fmt.Errorf("rule %q update at %#x conflicts with audio device behavior", rule.ID, update.Offset)
			}
		}
	}
	return nil
}

func prepareBehaviorConfig(cfg *SVGeneratorConfig) (*SVGeneratorConfig, error) {
	if cfg == nil {
		return nil, fmt.Errorf("SV generator config is nil")
	}
	if cfg.BehaviorRules == nil || cfg.CompiledBehavior != nil {
		return cfg, nil
	}
	prepared := *cfg
	if err := ValidateBehaviorRuleOwnership(prepared.BehaviorRules, prepared.DeviceClass); err != nil {
		return nil, err
	}
	if prepared.BARModel == nil {
		prepared.BARModel = barmodel.BuildBARModel(nil, prepared.ClassCode, nil)
	}
	if cfg.BARModel != nil {
		modelCopy := *cfg.BARModel
		modelCopy.Registers = append([]barmodel.BARRegister(nil), cfg.BARModel.Registers...)
		prepared.BARModel = &modelCopy
	}
	model, err := barmodel.ApplyBehaviorRules(prepared.BARModel, prepared.BehaviorRules)
	if err != nil {
		return nil, err
	}
	prepared.BARModel = model
	prepared.CompiledBehavior, err = CompileBehaviorRules(prepared.BehaviorRules, model)
	if err != nil {
		return nil, err
	}
	return &prepared, nil
}


type CompiledBehaviorEvent struct {
	Index       int
	DelayCycles uint32
	NextState   int
	Updates     []CompiledBehaviorUpdate
}

func CompileBehaviorRules(set *behavior.RuleSet, model *barmodel.BARModel) (*CompiledBehavior, error) {
	if set == nil {
		return nil, nil
	}
	if err := behavior.Validate(set); err != nil {
		return nil, err
	}
	if model == nil {
		return nil, fmt.Errorf("behavior rules require a BAR model")
	}
	stateIndex := make(map[string]int)
	states := make([]string, 0)
	addState := func(state string) int {
		if state == "" {
			return -1
		}
		if index, ok := stateIndex[state]; ok {
			return index
		}
		index := len(states)
		stateIndex[state] = index
		states = append(states, state)
		return index
	}
	initial := addState(set.InitialState)
	for _, rule := range set.Rules {
		addState(rule.State)
		addState(rule.NextState)
		for _, event := range rule.DelayedEvents {
			addState(event.NextState)
		}
	}
	bits := 1
	for (1 << bits) < len(states) {
		bits++
	}
	compiled := &CompiledBehavior{StateBits: bits, InitialState: initial, States: states}
	owned := make(map[uint32]struct{})
	initialValues := make(map[uint32]uint32, len(set.InitialRegisters))
	for _, initialValue := range set.InitialRegisters {
		initialValues[initialValue.Offset] = uint32(initialValue.Value)
	}
	triggerRWMasks := make(map[uint32]uint32, len(model.Registers))
	triggerW1CMasks := make(map[uint32]uint32, len(model.Registers))
	for _, register := range model.Registers {
		triggerRWMasks[register.Offset] = register.RWMask
		triggerW1CMasks[register.Offset] = register.W1CMask
	}
	for _, initialValue := range set.InitialRegisters {
		if initialValue.WritePolicy != nil {
			triggerRWMasks[initialValue.Offset] = uint32(initialValue.WritePolicy.RWMask)
			triggerW1CMasks[initialValue.Offset] = uint32(initialValue.WritePolicy.W1CMask)
		}
	}
	eventIndex := 0
	for ruleIndex, rule := range set.Rules {
		if rule.Access != behavior.AccessWrite || rule.Width != 4 || rule.Offset%4 != 0 {
			return nil, fmt.Errorf("rule %q is not a supported aligned 32-bit write trigger", rule.ID)
		}
		next := stateIndex[rule.State]
		if rule.NextState != "" {
			next = stateIndex[rule.NextState]
		}
		immediate, err := coalesceUpdates(rule.Updates)
		if err != nil {
			return nil, fmt.Errorf("rule %q immediate updates: %w", rule.ID, err)
		}
		compiledRule := CompiledBehaviorRule{
			Index: ruleIndex, State: stateIndex[rule.State], NextState: next,
			Offset: rule.Offset, Value: uint32(rule.Value), ValueMask: uint32(rule.ValueMask),
			ByteEnableMask: 0x0f, Updates: immediate,
		}
		for updateIndex := range compiledRule.Updates {
			update := &compiledRule.Updates[updateIndex]
			if update.Offset == rule.Offset {
				update.ComposeTriggerWrite = true
				update.TriggerRWMask = triggerRWMasks[rule.Offset]
				update.TriggerW1CMask = triggerW1CMasks[rule.Offset]
			}
			owned[update.Offset] = struct{}{}
		}
		type eventGroup struct {
			delay     uint32
			nextState int
			updates   []behavior.RegisterUpdate
		}
		eventGroups := make([]eventGroup, 0, len(rule.DelayedEvents))
		eventByDelay := make(map[uint32]int)
		for _, event := range rule.DelayedEvents {
			delay, scaleErr := scaleDelayCycles(event.DelayCycles, set.ClockHz)
			if scaleErr != nil {
				return nil, fmt.Errorf("rule %q: %w", rule.ID, scaleErr)
			}
			nextEventState := -1
			if event.NextState != "" {
				nextEventState = stateIndex[event.NextState]
			}
			if groupIndex, exists := eventByDelay[delay]; exists {
				group := &eventGroups[groupIndex]
				if group.nextState >= 0 && nextEventState >= 0 && group.nextState != nextEventState {
					return nil, fmt.Errorf("rule %q has conflicting states after delay rescaling", rule.ID)
				}
				if group.nextState < 0 {
					group.nextState = nextEventState
				}
				group.updates = append(group.updates, event.Updates...)
				continue
			}
			eventByDelay[delay] = len(eventGroups)
			eventGroups = append(eventGroups, eventGroup{delay: delay, nextState: nextEventState, updates: append([]behavior.RegisterUpdate(nil), event.Updates...)})
		}
		for _, group := range eventGroups {
			updates, mergeErr := coalesceUpdates(group.updates)
			if mergeErr != nil {
				return nil, fmt.Errorf("rule %q delayed updates: %w", rule.ID, mergeErr)
			}
			compiledEvent := CompiledBehaviorEvent{Index: eventIndex, DelayCycles: group.delay, NextState: group.nextState, Updates: updates}
			eventIndex++
			for _, update := range updates {
				owned[update.Offset] = struct{}{}
			}
			compiledRule.DelayedEvents = append(compiledRule.DelayedEvents, compiledEvent)
			compiled.AllEvents = append(compiled.AllEvents, compiledEvent)
		}
		compiled.Rules = append(compiled.Rules, compiledRule)
	}
	dueByOffset := make(map[uint32][]CompiledDueContribution)
	for _, event := range compiled.AllEvents {
		for _, update := range event.Updates {
			dueByOffset[update.Offset] = append(dueByOffset[update.Offset], CompiledDueContribution{
				EventIndex: event.Index, Value: update.Value, Mask: update.Mask,
			})
		}
	}
	for _, register := range model.Registers {
		if _, ok := owned[register.Offset]; ok {
			compiled.Registers = append(compiled.Registers, CompiledBehaviorRegister{
				Offset: register.Offset, Reset: initialValues[register.Offset], RWMask: register.RWMask, W1CMask: register.W1CMask,
			})
			if contributions := dueByOffset[register.Offset]; len(contributions) > 0 {
				compiled.DueRegisters = append(compiled.DueRegisters, CompiledDueRegister{
					Offset: register.Offset, Contributions: contributions,
				})
			}
		}
	}
	return compiled, nil
}

func coalesceUpdates(updates []behavior.RegisterUpdate) ([]CompiledBehaviorUpdate, error) {
	result := make([]CompiledBehaviorUpdate, 0, len(updates))
	byOffset := make(map[uint32]int)
	for _, update := range updates {
		if update.Width != 4 || update.Offset%4 != 0 {
			return nil, fmt.Errorf("update at %#x is not an aligned 32-bit register update", update.Offset)
		}
		if index, exists := byOffset[update.Offset]; exists {
			previous := &result[index]
			overlap := previous.Mask & uint32(update.Mask)
			if ((previous.Value ^ uint32(update.Value)) & overlap) != 0 {
				return nil, fmt.Errorf("update at %#x has contradictory overlapping masks", update.Offset)
			}
			previous.Value = (previous.Value &^ uint32(update.Mask)) | (uint32(update.Value) & uint32(update.Mask))
			previous.Mask |= uint32(update.Mask)
			continue
		}
		byOffset[update.Offset] = len(result)
		result = append(result, CompiledBehaviorUpdate{Offset: update.Offset, Value: uint32(update.Value) & uint32(update.Mask), Mask: uint32(update.Mask)})
	}
	return result, nil
}

func scaleDelayCycles(delay uint32, sourceClockHz uint64) (uint32, error) {
	if sourceClockHz == 0 {
		return 0, fmt.Errorf("source behavior clock is zero")
	}
	if uint64(delay) > ^uint64(0)/BehaviorRTLClockHz {
		return 0, fmt.Errorf("delay %d overflows RTL clock conversion", delay)
	}
	numerator := uint64(delay) * BehaviorRTLClockHz
	scaled := numerator / sourceClockHz
	if numerator%sourceClockHz != 0 {
		scaled++
	}
	if scaled == 0 {
		scaled = 1
	}
	if scaled > uint64(behavior.MaxDelayCycles) {
		return 0, fmt.Errorf("rescaled delay %d exceeds maximum %d", scaled, behavior.MaxDelayCycles)
	}
	return uint32(scaled), nil
}

func byteEnableMask(mask uint32) uint8 {
	var enables uint8
	for byteIndex := range 4 {
		if mask&(0xff<<uint(byteIndex*8)) != 0 {
			enables |= 1 << uint(byteIndex)
		}
	}
	return enables
}
