package behavior

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/sercanarga/pcileechgen/internal/donor/mmio"
)

const (
	RuleSchemaVersion = 1
	DefaultRuleClockHz uint64 = 125_000_000
	MaxDelayCycles uint32 = 1_000_000
	MaxRules = 256
	MaxDelayedEvents = 64
	maxUpdatesPerAction = 16
	MaxInitialRegisters = 1024
)

type AccessKind string

const (
	AccessRead         AccessKind = "read"
	AccessWrite        AccessKind = "write"
	UnknownInputIgnore            = "ignore"
)

type RuleSet struct {
	Version          int             `json:"version"`
	BDF              string          `json:"bdf,omitempty"`
	BARIndex         int             `json:"bar_index"`
	BARSize          int             `json:"bar_size"`
	ClockHz          uint64          `json:"clock_hz"`
	InitialState     string          `json:"initial_state"`
	InitialRegisters []RegisterValue `json:"initial_registers,omitempty"`
	UnknownInputPolicy string          `json:"unknown_input_policy,omitempty"`
	Rules            []Rule          `json:"rules"`
}

type RegisterValue struct {
	Offset      uint32               `json:"offset"`
	Width       uint8                `json:"width"`
	Value       uint64               `json:"value"`
	WritePolicy *RegisterWritePolicy `json:"write_policy,omitempty"`
}

type RegisterWritePolicy struct {
	RWMask  uint64 `json:"rw_mask"`
	W1CMask uint64 `json:"w1c_mask"`
}

type Rule struct {
	ID            string          `json:"id"`
	State         string          `json:"state"`
	Access        AccessKind      `json:"access"`
	Width         uint8           `json:"width"`
	Offset        uint32          `json:"offset"`
	Value         uint64          `json:"value"`
	ValueMask     uint64          `json:"value_mask"`
	NextState     string          `json:"next_state,omitempty"`
	Updates       []RegisterUpdate `json:"updates,omitempty"`
	DelayedEvents []DelayedEvent  `json:"delayed_events,omitempty"`
	Confidence    float64         `json:"confidence"`
	Provenance    []string        `json:"provenance,omitempty"`
}

type RegisterUpdate struct {
	Offset uint32 `json:"offset"`
	Width  uint8  `json:"width"`
	Value  uint64 `json:"value"`
	Mask   uint64 `json:"mask"`
}

type DelayedEvent struct {
	DelayCycles uint32           `json:"delay_cycles"`
	Updates     []RegisterUpdate `json:"updates,omitempty"`
	NextState   string           `json:"next_state,omitempty"`
}

type ReadMismatch struct {
	RecordIndex int    `json:"record_index"`
	Offset      uint32 `json:"offset"`
	Expected    uint64 `json:"expected"`
	Actual      uint64 `json:"actual"`
}

type ReplayResult struct {
	TerminalState  string            `json:"terminal_state"`
	Registers      map[uint32]uint64 `json:"registers"`
	MatchedRules   []string          `json:"matched_rules"`
	ReadMismatches []ReadMismatch    `json:"read_mismatches,omitempty"`
}

func widthMask(width uint8) (uint64, bool) {
	switch width {
	case 1:
		return 0xff, true
	case 2:
		return 0xffff, true
	case 4:
		return 0xffffffff, true
	case 8:
		return ^uint64(0), true
	default:
		return 0, false
	}
}

func validateUpdate(update RegisterUpdate, barSize int) error {
	validBits, ok := widthMask(update.Width)
	if !ok {
		return fmt.Errorf("unsupported update width %d", update.Width)
	}
	if uint64(update.Offset)+uint64(update.Width) > uint64(barSize) {
		return fmt.Errorf("update at %#x width %d is outside BAR size %#x", update.Offset, update.Width, barSize)
	}
	if update.Offset%uint32(update.Width) != 0 {
		return fmt.Errorf("update at %#x width %d is misaligned", update.Offset, update.Width)
	}
	if update.Mask == 0 {
		return fmt.Errorf("update at %#x has an empty mask", update.Offset)
	}
	if update.Mask&^validBits != 0 || update.Value&^validBits != 0 {
		return fmt.Errorf("update at %#x exceeds its %d-byte width", update.Offset, update.Width)
	}
	return nil
}

func validateUpdateGroup(updates []RegisterUpdate, barSize int) error {
	for i, update := range updates {
		if err := validateUpdate(update, barSize); err != nil {
			return err
		}
		for j := range i {
			previous := updates[j]
			previousEnd := uint64(previous.Offset) + uint64(previous.Width)
			updateEnd := uint64(update.Offset) + uint64(update.Width)
			if uint64(previous.Offset) >= updateEnd || uint64(update.Offset) >= previousEnd {
				continue
			}
			if previous.Offset != update.Offset || previous.Width != update.Width {
				return fmt.Errorf("updates at %#x and %#x overlap different register widths", previous.Offset, update.Offset)
			}
			overlap := previous.Mask & update.Mask
			if ((previous.Value ^ update.Value) & overlap) != 0 {
				return fmt.Errorf("updates at %#x assign contradictory overlapping masks", update.Offset)
			}
		}
	}
	return nil
}

func Validate(set *RuleSet) error {
	if set == nil {
		return fmt.Errorf("behavior rule set is nil")
	}
	if set.Version != RuleSchemaVersion {
		return fmt.Errorf("unsupported behavior rule schema version %d (want %d)", set.Version, RuleSchemaVersion)
	}
	if set.BARIndex < 0 || set.BARIndex > 5 {
		return fmt.Errorf("invalid behavior BAR index %d", set.BARIndex)
	}
	if set.BARSize <= 0 {
		return fmt.Errorf("behavior BAR size must be positive")
	}
	if set.ClockHz == 0 || set.ClockHz > 1_000_000_000 {
		return fmt.Errorf("behavior clock_hz must be between 1 and 1000000000")
	}
	if set.InitialState == "" {
		return fmt.Errorf("behavior initial_state is required")
	}
	if len(set.InitialRegisters) > MaxInitialRegisters {
		return fmt.Errorf("initial register count %d exceeds maximum %d", len(set.InitialRegisters), MaxInitialRegisters)
	}
	for i, initial := range set.InitialRegisters {
		validBits, ok := widthMask(initial.Width)
		if !ok {
			return fmt.Errorf("initial register at %#x has unsupported width %d", initial.Offset, initial.Width)
		}
		if uint64(initial.Offset)+uint64(initial.Width) > uint64(set.BARSize) {
			return fmt.Errorf("initial register at %#x is outside BAR size %#x", initial.Offset, set.BARSize)
		}
		if initial.Offset%uint32(initial.Width) != 0 || initial.Value&^validBits != 0 {
			return fmt.Errorf("initial register at %#x is invalid for width %d", initial.Offset, initial.Width)
		}
		if initial.WritePolicy != nil {
			if initial.WritePolicy.RWMask&^validBits != 0 || initial.WritePolicy.W1CMask&^validBits != 0 {
				return fmt.Errorf("initial register at %#x has write-policy bits outside width %d", initial.Offset, initial.Width)
			}
			if initial.WritePolicy.RWMask&initial.WritePolicy.W1CMask != 0 {
				return fmt.Errorf("initial register at %#x has overlapping RW/W1C masks", initial.Offset)
			}
		}
		for j := range i {
			previous := set.InitialRegisters[j]
			if uint64(previous.Offset) < uint64(initial.Offset)+uint64(initial.Width) &&
				uint64(initial.Offset) < uint64(previous.Offset)+uint64(previous.Width) {
				return fmt.Errorf("initial registers at %#x and %#x overlap", previous.Offset, initial.Offset)
			}
		}
	}
	if set.UnknownInputPolicy != "" && set.UnknownInputPolicy != UnknownInputIgnore {
		return fmt.Errorf("unsupported unknown_input_policy %q", set.UnknownInputPolicy)
	}
	if len(set.Rules) > MaxRules {
		return fmt.Errorf("behavior rule count %d exceeds maximum %d", len(set.Rules), MaxRules)
	}

	ids := make(map[string]struct{}, len(set.Rules))
	seenRules := make([]Rule, 0, len(set.Rules))
	delayedCount := 0
	var allDelayedUpdates []RegisterUpdate
	for i, rule := range set.Rules {
		if rule.ID == "" {
			return fmt.Errorf("rule %d has no id", i)
		}
		if _, exists := ids[rule.ID]; exists {
			return fmt.Errorf("duplicate behavior rule id %q", rule.ID)
		}
		ids[rule.ID] = struct{}{}
		if rule.State == "" {
			return fmt.Errorf("rule %q has no state", rule.ID)
		}
		if rule.Access != AccessRead && rule.Access != AccessWrite {
			return fmt.Errorf("rule %q has invalid access %q", rule.ID, rule.Access)
		}
		validBits, ok := widthMask(rule.Width)
		if !ok {
			return fmt.Errorf("rule %q has unsupported width %d", rule.ID, rule.Width)
		}
		if uint64(rule.Offset)+uint64(rule.Width) > uint64(set.BARSize) {
			return fmt.Errorf("rule %q access is outside BAR size %#x", rule.ID, set.BARSize)
		}
		if rule.Offset%uint32(rule.Width) != 0 {
			return fmt.Errorf("rule %q access is misaligned", rule.ID)
		}
		if rule.ValueMask&^validBits != 0 || rule.Value&^validBits != 0 {
			return fmt.Errorf("rule %q match exceeds its %d-byte width", rule.ID, rule.Width)
		}
		if rule.Confidence < 0 || rule.Confidence > 1 {
			return fmt.Errorf("rule %q confidence must be between 0 and 1", rule.ID)
		}
		for _, previous := range seenRules {
			sameKey := previous.State == rule.State && previous.Access == rule.Access &&
				previous.Width == rule.Width && previous.Offset == rule.Offset
			if sameKey && ((previous.Value^rule.Value)&previous.ValueMask&rule.ValueMask) == 0 {
				return fmt.Errorf("rules %q and %q have overlapping matches", previous.ID, rule.ID)
			}
		}
		seenRules = append(seenRules, rule)
		if len(rule.Updates) > maxUpdatesPerAction {
			return fmt.Errorf("rule %q has too many immediate updates", rule.ID)
		}
		if err := validateUpdateGroup(rule.Updates, set.BARSize); err != nil {
			return fmt.Errorf("rule %q: %w", rule.ID, err)
		}
		delayUpdates := make(map[uint32][]RegisterUpdate)
		delayStates := make(map[uint32]string)
		for _, event := range rule.DelayedEvents {
			delayedCount++
			if event.DelayCycles == 0 || event.DelayCycles > MaxDelayCycles {
				return fmt.Errorf("rule %q delay %d is outside 1..%d cycles", rule.ID, event.DelayCycles, MaxDelayCycles)
			}
			if len(event.Updates) == 0 && event.NextState == "" {
				return fmt.Errorf("rule %q has an empty delayed event", rule.ID)
			}
			if len(event.Updates) > maxUpdatesPerAction {
				return fmt.Errorf("rule %q delayed event has too many updates", rule.ID)
			}
			if err := validateUpdateGroup(event.Updates, set.BARSize); err != nil {
				return fmt.Errorf("rule %q delayed event: %w", rule.ID, err)
			}
			if state := delayStates[event.DelayCycles]; state != "" && event.NextState != "" && state != event.NextState {
				return fmt.Errorf("rule %q has contradictory states at delay %d", rule.ID, event.DelayCycles)
			}
			if event.NextState != "" {
				delayStates[event.DelayCycles] = event.NextState
			}
			delayUpdates[event.DelayCycles] = append(delayUpdates[event.DelayCycles], event.Updates...)
			allDelayedUpdates = append(allDelayedUpdates, event.Updates...)
			if len(delayUpdates[event.DelayCycles]) > maxUpdatesPerAction {
				return fmt.Errorf("rule %q has too many combined updates at delay %d", rule.ID, event.DelayCycles)
			}
		}
		for delay, updates := range delayUpdates {
			if err := validateUpdateGroup(updates, set.BARSize); err != nil {
				return fmt.Errorf("rule %q events at delay %d: %w", rule.ID, delay, err)
			}
		}
	}
	if err := validateUpdateGroup(allDelayedUpdates, set.BARSize); err != nil {
		return fmt.Errorf("delayed events: %w", err)
	}
	if delayedCount > MaxDelayedEvents {
		return fmt.Errorf("behavior delayed event count %d exceeds maximum %d", delayedCount, MaxDelayedEvents)
	}
	return nil
}

func LoadRuleSet(path string) (*RuleSet, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read behavior rules %q: %w", path, err)
	}
	var set RuleSet
	if err := json.Unmarshal(data, &set); err != nil {
		return nil, fmt.Errorf("decode behavior rules %q: %w", path, err)
	}
	if err := Validate(&set); err != nil {
		return nil, fmt.Errorf("validate behavior rules %q: %w", path, err)
	}
	return &set, nil
}

func SaveRuleSet(set *RuleSet, path string) error {
	if err := Validate(set); err != nil {
		return err
	}
	data, err := json.MarshalIndent(set, "", "  ")
	if err != nil {
		return fmt.Errorf("encode behavior rules: %w", err)
	}
	data = append(data, '\n')
	if err := os.WriteFile(path, data, 0o644); err != nil {
		return fmt.Errorf("write behavior rules %q: %w", path, err)
	}
	return nil
}

type pendingEvent struct {
	key       string
	order     int
	remaining uint32
	event     DelayedEvent
}

type Engine struct {
	set     *RuleSet
	state   string
	regs    map[uint32]uint64
	pending []pendingEvent
	matched []string
	suppressNext bool
	eventOrder map[string]int
	writePolicies map[uint32]RegisterWritePolicy
}

func NewEngine(set *RuleSet) (*Engine, error) {
	if err := Validate(set); err != nil {
		return nil, err
	}
	regs := make(map[uint32]uint64, len(set.InitialRegisters))
	for _, rule := range set.Rules {
		for _, update := range rule.Updates {
			if _, exists := regs[update.Offset]; !exists {
				regs[update.Offset] = 0
			}
		}
		for _, event := range rule.DelayedEvents {
			for _, update := range event.Updates {
				if _, exists := regs[update.Offset]; !exists {
					regs[update.Offset] = 0
				}
			}
		}
	}
	writePolicies := make(map[uint32]RegisterWritePolicy)
	for _, initial := range set.InitialRegisters {
		regs[initial.Offset] = initial.Value
		if initial.WritePolicy != nil {
			writePolicies[initial.Offset] = *initial.WritePolicy
		}
	}
	eventOrder := make(map[string]int)
	order := 0
	for _, rule := range set.Rules {
		for eventIndex := range rule.DelayedEvents {
			eventOrder[fmt.Sprintf("%s/%d", rule.ID, eventIndex)] = order
			order++
		}
	}
	return &Engine{
		set: set, state: set.InitialState, regs: regs, eventOrder: eventOrder, writePolicies: writePolicies,
	}, nil
}

func (e *Engine) State() string { return e.state }

func (e *Engine) Register(offset uint32) (uint64, bool) {
	value, ok := e.regs[offset]
	return value, ok
}

func applyUpdates(regs map[uint32]uint64, updates []RegisterUpdate) {
	for _, update := range updates {
		old := regs[update.Offset]
		regs[update.Offset] = (old &^ update.Mask) | (update.Value & update.Mask)
	}
}

func normalWriteValue(oldValue, writeValue uint64, width uint8, policy *RegisterWritePolicy) uint64 {
	validBits, _ := widthMask(width)
	rwMask := validBits
	var w1cMask uint64
	if policy != nil {
		rwMask = policy.RWMask
		w1cMask = policy.W1CMask
	}
	next := (oldValue &^ rwMask) | (writeValue & rwMask)
	next &^= writeValue & w1cMask
	return next & validBits
}

func (e *Engine) Apply(record mmio.AccessRecord) error {
	if e == nil || e.set == nil {
		return fmt.Errorf("behavior engine is nil")
	}
	width := record.Width
	if width == 0 {
		width = 4
	}
	var access AccessKind
	switch record.Type {
	case mmio.AccessRead:
		access = AccessRead
	case mmio.AccessWrite:
		access = AccessWrite
	default:
		return fmt.Errorf("invalid MMIO access type %d", record.Type)
	}
	if e.suppressNext {
		e.suppressNext = false
		return nil
	}
	for _, rule := range e.set.Rules {
		if rule.State != e.state || rule.Access != access || rule.Width != width || rule.Offset != record.Offset {
			continue
		}
		if ((record.Value ^ rule.Value) & rule.ValueMask) != 0 {
			continue
		}
		if access == AccessWrite {
			for _, update := range rule.Updates {
				if update.Offset == record.Offset && update.Width == width {
					policy, hasPolicy := e.writePolicies[record.Offset]
					if hasPolicy {
						e.regs[record.Offset] = normalWriteValue(e.regs[record.Offset], record.Value, width, &policy)
					} else {
						e.regs[record.Offset] = normalWriteValue(e.regs[record.Offset], record.Value, width, nil)
					}
					break
				}
			}
		}
		applyUpdates(e.regs, rule.Updates)
		if rule.NextState != "" {
			e.state = rule.NextState
		}
		for eventIndex, event := range rule.DelayedEvents {
			key := fmt.Sprintf("%s/%d", rule.ID, eventIndex)
			alreadyPending := false
			for _, pending := range e.pending {
				if pending.key == key {
					alreadyPending = true
					break
				}
			}
			if alreadyPending {
				continue
			}
			if len(e.pending) >= MaxDelayedEvents {
				return fmt.Errorf("pending behavior event limit %d exceeded", MaxDelayedEvents)
			}
			e.pending = append(e.pending, pendingEvent{key: key, order: e.eventOrder[key], remaining: event.DelayCycles, event: event})
		}
		e.matched = append(e.matched, rule.ID)
		break
	}
	return nil
}

func (e *Engine) Advance(cycles uint32) error {
	if e == nil {
		return fmt.Errorf("behavior engine is nil")
	}
	exactDeadline := false
	remaining := e.pending[:0]
	var due []pendingEvent
	for _, pending := range e.pending {
		if pending.remaining > cycles {
			pending.remaining -= cycles
			remaining = append(remaining, pending)
			continue
		}
		if pending.remaining == cycles && cycles != 0 {
			exactDeadline = true
		}
		due = append(due, pending)
	}
	sort.Slice(due, func(i, j int) bool { return due[i].order < due[j].order })
	var dueUpdates []RegisterUpdate
	for _, pending := range due {
		dueUpdates = append(dueUpdates, pending.event.Updates...)
		if pending.event.NextState != "" {
			e.state = pending.event.NextState
		}
	}
	applyUpdates(e.regs, dueUpdates)
	e.pending = remaining
	e.suppressNext = exactDeadline
	return nil
}

func durationToCycles(delta time.Duration, clockHz uint64) (uint64, error) {
	if delta < 0 {
		return 0, fmt.Errorf("negative duration")
	}
	seconds := uint64(delta / time.Second)
	remainder := uint64(delta % time.Second)
	maximum := ^uint64(0)
	if seconds != 0 && clockHz > maximum/seconds {
		return 0, fmt.Errorf("duration cycle conversion overflows")
	}
	whole := seconds * clockHz
	if remainder != 0 && clockHz > maximum/remainder {
		return 0, fmt.Errorf("duration cycle conversion overflows")
	}
	fraction := remainder * clockHz / uint64(time.Second)
	if whole > maximum-fraction {
		return 0, fmt.Errorf("duration cycle conversion overflows")
	}
	return whole + fraction, nil
}

func Replay(set *RuleSet, trace *mmio.TraceResult) (*ReplayResult, error) {
	if trace == nil {
		return nil, fmt.Errorf("replay trace is nil")
	}
	engine, err := NewEngine(set)
	if err != nil {
		return nil, err
	}
	if trace.BARIndex != set.BARIndex {
		return nil, fmt.Errorf("trace BAR%d does not match rule BAR%d", trace.BARIndex, set.BARIndex)
	}
	var previous time.Duration
	result := &ReplayResult{}
	for index, record := range trace.Records {
		if record.BARIndex != 0 && record.BARIndex != set.BARIndex {
			return nil, fmt.Errorf("record %d BAR%d does not match rule BAR%d", index, record.BARIndex, set.BARIndex)
		}
		if record.Timestamp < previous {
			return nil, fmt.Errorf("record %d timestamp is not monotonic", index)
		}
		delta := record.Timestamp - previous
		cycles64, cycleErr := durationToCycles(delta, set.ClockHz)
		if cycleErr != nil {
			return nil, fmt.Errorf("record %d: %w", index, cycleErr)
		}
		if cycles64 > uint64(^uint32(0)) {
			return nil, fmt.Errorf("record %d cycle delta %d exceeds engine bound", index, cycles64)
		}
		if err := engine.Advance(uint32(cycles64)); err != nil {
			return nil, err
		}
		previous = record.Timestamp
		if record.Type == mmio.AccessRead {
			actual := engine.regs[record.Offset]
			if actual != record.Value {
				result.ReadMismatches = append(result.ReadMismatches, ReadMismatch{RecordIndex: index, Offset: record.Offset, Expected: record.Value, Actual: actual})
			}
		}
		if err := engine.Apply(record); err != nil {
			return nil, fmt.Errorf("record %d: %w", index, err)
		}
	}
	result.TerminalState = engine.state
	result.Registers = make(map[uint32]uint64, len(engine.regs))
	for offset, value := range engine.regs {
		result.Registers[offset] = value
	}
	result.MatchedRules = append(result.MatchedRules, engine.matched...)
	return result, nil
}

type inferredObservation struct {
	ordinal       int
	triggerOffset uint32
	triggerWidth  uint8
	triggerValue  uint64
	updateOffset  uint32
	updateWidth   uint8
	before        uint64
	after         uint64
	delayCycles   uint32
	sessionIndex  int
	provenance    string
}

func observations(trace *mmio.TraceResult, clockHz uint64, sessionIndex int) []inferredObservation {
	var out []inferredObservation
	ordinal := 0
	for i, record := range trace.Records {
		if record.Type != mmio.AccessWrite || record.Width != 4 || record.Offset%4 != 0 {
			continue
		}
		baseline := make(map[uint32]mmio.AccessRecord)
		for j := i + 1; j < len(trace.Records); j++ {
			candidate := trace.Records[j]
			if candidate.Type == mmio.AccessWrite {
				break
			}
			if candidate.Type != mmio.AccessRead || candidate.Width != 4 || candidate.Offset%4 != 0 {
				continue
			}
			before, exists := baseline[candidate.Offset]
			if !exists {
				baseline[candidate.Offset] = candidate
				continue
			}
			if before.Value == candidate.Value || before.Width != candidate.Width {
				continue
			}
			delta := candidate.Timestamp - record.Timestamp
			if delta < 0 {
				break
			}
			cycles, cycleErr := durationToCycles(delta, clockHz)
			if cycleErr != nil || cycles > uint64(MaxDelayCycles) {
				break
			}
			out = append(out, inferredObservation{
				sessionIndex: sessionIndex,
				ordinal: ordinal, triggerOffset: record.Offset, triggerWidth: record.Width,
				triggerValue: record.Value, updateOffset: candidate.Offset, updateWidth: candidate.Width,
				before: before.Value, after: candidate.Value, delayCycles: uint32(cycles),
				provenance: fmt.Sprintf("session=%d trace=%s bar=%d record=%d", sessionIndex, trace.StartTime.UTC().Format(time.RFC3339Nano), trace.BARIndex, i),
			})
			ordinal++
			break
		}
	}
	return out
}

func Infer(traces ...*mmio.TraceResult) (*RuleSet, error) {
	if len(traces) == 0 {
		return nil, fmt.Errorf("behavior inference requires at least one trace")
	}
	first := traces[0]
	if first == nil || first.BARSize <= 0 {
		return nil, fmt.Errorf("trace 0 has no valid BAR aperture")
	}
	set := &RuleSet{
		Version: RuleSchemaVersion, BDF: first.BDF, BARIndex: first.BARIndex, BARSize: first.BARSize,
		ClockHz: DefaultRuleClockHz, InitialState: "initial", UnknownInputPolicy: UnknownInputIgnore,
	}
	type candidateGroup struct {
		key      string
		ordinal  int
		outcomes map[uint64][]inferredObservation
	}
	groups := make(map[string]*candidateGroup)
	for traceIndex, trace := range traces {
		if trace == nil {
			return nil, fmt.Errorf("trace %d is nil", traceIndex)
		}
		if trace.BDF != first.BDF || trace.BARIndex != first.BARIndex || trace.BARSize != first.BARSize {
			return nil, fmt.Errorf("trace %d target does not match trace 0", traceIndex)
		}
		for _, obs := range observations(trace, set.ClockHz, traceIndex) {
			key := fmt.Sprintf("%d/%d/%x/%d/%d/%x", obs.triggerOffset, obs.triggerWidth, obs.triggerValue, obs.updateOffset, obs.updateWidth, obs.before)
			group, exists := groups[key]
			if !exists {
				group = &candidateGroup{key: key, ordinal: obs.ordinal, outcomes: make(map[uint64][]inferredObservation)}
				groups[key] = group
			}
			if obs.ordinal < group.ordinal {
				group.ordinal = obs.ordinal
			}
			group.outcomes[obs.after] = append(group.outcomes[obs.after], obs)
		}
	}
	ordered := make([]*candidateGroup, 0, len(groups))
	for _, group := range groups {
		ordered = append(ordered, group)
	}
	sort.Slice(ordered, func(i, j int) bool {
		if ordered[i].ordinal == ordered[j].ordinal {
			return ordered[i].key < ordered[j].key
		}
		return ordered[i].ordinal < ordered[j].ordinal
	})
	initialValues := make(map[uint32]uint64)
	for _, group := range ordered {
		if len(group.outcomes) > 1 {
			counterexamples := make([]string, 0, len(group.outcomes))
			for value, observations := range group.outcomes {
				sources := make([]string, 0, len(observations))
				for _, observation := range observations {
					sources = append(sources, observation.provenance)
				}
				sort.Strings(sources)
				counterexamples = append(counterexamples, fmt.Sprintf("value=%#x provenance=%v", value, sources))
			}
			sort.Strings(counterexamples)
			return nil, fmt.Errorf("conflicting observed transitions for %s: %v", group.key, counterexamples)
		}
		var supported []inferredObservation
		for _, observations := range group.outcomes {
			supported = observations
		}
		if len(supported) < 2 {
			continue
		}
		sort.SliceStable(supported, func(i, j int) bool { return supported[i].delayCycles < supported[j].delayCycles })
		obs := supported[len(supported)/2]
		mask := obs.before ^ obs.after
		if mask == 0 {
			continue
		}
		if initial, exists := initialValues[obs.updateOffset]; exists && initial != obs.before {
			return nil, fmt.Errorf("conflicting initial values at %#x", obs.updateOffset)
		}
		if _, exists := initialValues[obs.updateOffset]; !exists {
			initialValues[obs.updateOffset] = obs.before
			set.InitialRegisters = append(set.InitialRegisters, RegisterValue{Offset: obs.updateOffset, Width: obs.updateWidth, Value: obs.before})
		}
		state := "initial"
		sessions := make(map[int]struct{})
		for _, observation := range supported {
			sessions[observation.sessionIndex] = struct{}{}
		}
		confidence := float64(len(sessions)) / float64(len(traces))
		independentEvidence := float64(len(sessions)) / 2
		if independentEvidence > 1 {
			independentEvidence = 1
		}
		if confidence > independentEvidence {
			confidence = independentEvidence
		}
		update := RegisterUpdate{Offset: obs.updateOffset, Width: obs.updateWidth, Value: obs.after, Mask: mask}
		rule := Rule{
			ID: fmt.Sprintf("trace_rule_%d", len(set.Rules)+1), State: state, Access: AccessWrite,
			Width: obs.triggerWidth, Offset: obs.triggerOffset, Value: obs.triggerValue,
			ValueMask: mustWidthMask(obs.triggerWidth), Confidence: confidence,
		}
		for _, source := range supported {
			rule.Provenance = append(rule.Provenance, source.provenance)
		}
		if obs.delayCycles == 0 {
			rule.Updates = []RegisterUpdate{update}
		} else {
			rule.DelayedEvents = []DelayedEvent{{DelayCycles: obs.delayCycles, Updates: []RegisterUpdate{update}}}
		}
		set.Rules = append(set.Rules, rule)
		if len(set.Rules) >= MaxRules {
			break
		}
	}
	if err := Validate(set); err != nil {
		return nil, err
	}
	return set, nil
}

func mustWidthMask(width uint8) uint64 {
	mask, _ := widthMask(width)
	return mask
}
