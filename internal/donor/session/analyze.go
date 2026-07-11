package session

import (
	"fmt"
	"sort"
	"time"

	"github.com/sercanarga/pcileechgen/internal/donor/mmio"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

// Confidence identifies how broadly an inferred behavior was observed.
type Confidence string

const (
	ConfidenceInferred Confidence = "inferred"
	ConfidenceObserved Confidence = "observed"
)

// WriteEffect describes an observed relationship between a write and later read.
type WriteEffect string

const (
	WriteEffectUnknown      WriteEffect = "unknown"
	WriteEffectDirectLatch  WriteEffect = "direct latch"
	WriteEffectSelfClearing WriteEffect = "self-clearing"
)

// CaptureData is the normalized input for initialization analysis.
type CaptureData struct {
	Manifest     *Manifest
	Trace        *mmio.TraceResult
	ConfigBefore *pci.ConfigSpace
	ConfigAfter  *pci.ConfigSpace
}

// Phase is a time-bounded group of initialization accesses.
type Phase struct {
	Session int           `json:"session"`
	Index   int           `json:"index"`
	Start   time.Duration `json:"start"`
	End     time.Duration `json:"end"`
	Kind    string        `json:"kind"`
	Records int           `json:"records"`
}

// RegisterEvidence summarizes observations for one BAR offset.
type RegisterEvidence struct {
	Offset         uint32      `json:"offset"`
	Widths         []int       `json:"widths"`
	Reads          int         `json:"reads"`
	Writes         int         `json:"writes"`
	Values         []uint32    `json:"values"`
	Classification string      `json:"classification"`
	Confidence     Confidence  `json:"confidence"`
	FirstPhase     int         `json:"first_phase"`
	Polling        bool        `json:"polling"`
	WriteEffect    WriteEffect `json:"write_effect"`
}

// Dependency is an observed write followed by polling of another register.
type Dependency struct {
	WriteOffset uint32 `json:"write_offset"`
	ReadOffset  uint32 `json:"read_offset"`
	Occurrences int    `json:"occurrences"`
}

// ConfigChange is one changed configuration-space DWORD.
type ConfigChange struct {
	Offset uint32 `json:"offset"`
	Before uint32 `json:"before"`
	After  uint32 `json:"after"`
	Mask   uint32 `json:"mask"`
}

// MACCandidate is a cautiously inferred Ethernet MAC location.
type MACCandidate struct {
	Address    string     `json:"address"`
	Source     string     `json:"source"`
	Confidence Confidence `json:"confidence"`
}

// Analysis is the deterministic result of comparing initialization sessions.
type Analysis struct {
	SessionCount  int                `json:"session_count"`
	Phases        []Phase            `json:"phases"`
	Registers     []RegisterEvidence `json:"registers"`
	Dependencies  []Dependency       `json:"dependencies,omitempty"`
	ConfigChanges []ConfigChange     `json:"config_changes,omitempty"`
	MACCandidates []MACCandidate     `json:"mac_candidates,omitempty"`
}

type registerAggregate struct {
	widths      map[uint8]struct{}
	values      map[uint32]struct{}
	sessions    map[int]struct{}
	reads       int
	writes      int
	firstPhase  int
	polling     bool
	writeEffect WriteEffect
}

// AnalyzeInitialization derives evidence from repeated capture sessions.
func AnalyzeInitialization(captures []CaptureData) Analysis {
	analysis := Analysis{SessionCount: len(captures)}
	registers := make(map[uint32]*registerAggregate)
	dependencies := make(map[[2]uint32]int)
	changes := make(map[string]ConfigChange)
	for sessionIndex, capture := range captures {
		if capture.Trace == nil {
			continue
		}
		phases, phaseByRecord := segmentPhases(sessionIndex, capture.Trace.Records)
		analysis.Phases = append(analysis.Phases, phases...)
		collectRegisterEvidence(registers, sessionIndex, capture.Trace.Records, phaseByRecord)
		collectDependencies(dependencies, capture.Trace.Records)
		collectConfigChanges(changes, capture.ConfigBefore, capture.ConfigAfter)
	}
	analysis.Registers = finishRegisterEvidence(registers)
	analysis.Dependencies = finishDependencies(dependencies)
	analysis.ConfigChanges = finishConfigChanges(changes)
	analysis.MACCandidates = inferMACCandidates(captures, analysis.Registers)
	return analysis
}

func segmentPhases(session int, records []mmio.AccessRecord) ([]Phase, []int) {
	if len(records) == 0 {
		return nil, nil
	}
	phaseByRecord := make([]int, len(records))
	start := 0
	var phases []Phase
	for i := 1; i <= len(records); i++ {
		boundary := i == len(records) || records[i].Timestamp-records[i-1].Timestamp > time.Millisecond
		if !boundary {
			continue
		}
		index := len(phases)
		for j := start; j < i; j++ {
			phaseByRecord[j] = index
		}
		kind := "initialization"
		if records[start].Type == mmio.AccessWrite && hasRepeatedRead(records[start:i]) {
			kind = "command-poll"
		}
		phases = append(phases, Phase{Session: session, Index: index, Start: records[start].Timestamp, End: records[i-1].Timestamp, Kind: kind, Records: i - start})
		start = i
	}
	return phases, phaseByRecord
}

func collectRegisterEvidence(out map[uint32]*registerAggregate, session int, records []mmio.AccessRecord, phaseByRecord []int) {
	for i, record := range records {
		a := out[record.Offset]
		if a == nil {
			a = &registerAggregate{widths: make(map[uint8]struct{}), values: make(map[uint32]struct{}), sessions: make(map[int]struct{}), firstPhase: phaseByRecord[i], writeEffect: WriteEffectUnknown}
			out[record.Offset] = a
		}
		a.widths[record.Width] = struct{}{}
		a.values[record.Value] = struct{}{}
		a.sessions[session] = struct{}{}
		if record.Type == mmio.AccessWrite {
			a.writes++
			for _, later := range records[i+1:] {
				if later.Offset != record.Offset || later.Type != mmio.AccessRead {
					continue
				}
				if later.Value == 0 && record.Value != 0 {
					a.writeEffect = WriteEffectSelfClearing
				} else if later.Value == record.Value && a.writeEffect == WriteEffectUnknown {
					a.writeEffect = WriteEffectDirectLatch
				}
				break
			}
		} else {
			a.reads++
		}
	}
	for offset, a := range out {
		if repeatedReadCount(records, offset) >= 3 {
			a.polling = true
		}
	}
}

func collectDependencies(out map[[2]uint32]int, records []mmio.AccessRecord) {
	for i, record := range records {
		if record.Type != mmio.AccessWrite {
			continue
		}
		limit := i + 17
		if limit > len(records) {
			limit = len(records)
		}
		counts := make(map[uint32]int)
		for _, later := range records[i+1 : limit] {
			if later.Type == mmio.AccessRead && later.Offset != record.Offset {
				counts[later.Offset]++
			}
		}
		for offset, count := range counts {
			if count >= 3 {
				out[[2]uint32{record.Offset, offset}]++
			}
		}
	}
}

func collectConfigChanges(out map[string]ConfigChange, before, after *pci.ConfigSpace) {
	if before == nil || after == nil {
		return
	}
	limit := before.Size
	if after.Size < limit {
		limit = after.Size
	}
	for offset := 0; offset+4 <= limit; offset += 4 {
		oldValue, newValue := before.ReadU32(offset), after.ReadU32(offset)
		if oldValue != newValue {
			change := ConfigChange{Offset: uint32(offset), Before: oldValue, After: newValue, Mask: oldValue ^ newValue}
			out[fmt.Sprintf("%x:%x:%x", offset, oldValue, newValue)] = change
		}
	}
}

func finishRegisterEvidence(input map[uint32]*registerAggregate) []RegisterEvidence {
	result := make([]RegisterEvidence, 0, len(input))
	for offset, a := range input {
		widths := sortedUint8Keys(a.widths)
		values := sortedUint32Keys(a.values)
		classification := "stable"
		switch {
		case a.polling:
			classification = "polling"
		case a.reads == 0:
			classification = "write-only"
		case a.writes > 0:
			classification = "read-write"
		case len(values) > 1:
			classification = "changing"
		}
		confidence := ConfidenceInferred
		if len(a.sessions) >= 2 {
			confidence = ConfidenceObserved
		}
		result = append(result, RegisterEvidence{Offset: offset, Widths: widths, Reads: a.reads, Writes: a.writes, Values: values, Classification: classification, Confidence: confidence, FirstPhase: a.firstPhase, Polling: a.polling, WriteEffect: a.writeEffect})
	}
	sort.Slice(result, func(i, j int) bool { return result[i].Offset < result[j].Offset })
	return result
}

func finishDependencies(input map[[2]uint32]int) []Dependency {
	result := make([]Dependency, 0, len(input))
	for key, count := range input {
		result = append(result, Dependency{WriteOffset: key[0], ReadOffset: key[1], Occurrences: count})
	}
	sort.Slice(result, func(i, j int) bool {
		if result[i].WriteOffset == result[j].WriteOffset {
			return result[i].ReadOffset < result[j].ReadOffset
		}
		return result[i].WriteOffset < result[j].WriteOffset
	})
	return result
}

func finishConfigChanges(input map[string]ConfigChange) []ConfigChange {
	result := make([]ConfigChange, 0, len(input))
	for _, change := range input {
		result = append(result, change)
	}
	sort.Slice(result, func(i, j int) bool {
		if result[i].Offset != result[j].Offset {
			return result[i].Offset < result[j].Offset
		}
		if result[i].Before != result[j].Before {
			return result[i].Before < result[j].Before
		}
		return result[i].After < result[j].After
	})
	return result
}

func inferMACCandidates(captures []CaptureData, registers []RegisterEvidence) []MACCandidate {
	if len(captures) == 0 || captures[0].Manifest == nil || captures[0].Manifest.Device.ClassCode>>16 != 0x02 {
		return nil
	}
	var low, high *RegisterEvidence
	for i := range registers {
		if registers[i].Offset == 0 && len(registers[i].Values) == 1 {
			low = &registers[i]
		}
		if registers[i].Offset == 4 && len(registers[i].Values) == 1 {
			high = &registers[i]
		}
	}
	if low == nil || high == nil {
		return nil
	}
	lo, hi := low.Values[0], high.Values[0]
	if lo == 0 || lo == 0xffffffff || hi&0xffff == 0 || hi&0xffff == 0xffff {
		return nil
	}
	address := fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x", byte(lo), byte(lo>>8), byte(lo>>16), byte(lo>>24), byte(hi), byte(hi>>8))
	return []MACCandidate{{Address: address, Source: "BAR offsets 0x0/0x4", Confidence: ConfidenceInferred}}
}

func hasRepeatedRead(records []mmio.AccessRecord) bool {
	counts := make(map[uint32]int)
	for _, record := range records {
		if record.Type == mmio.AccessRead {
			counts[record.Offset]++
		}
	}
	for _, count := range counts {
		if count >= 3 {
			return true
		}
	}
	return false
}

func repeatedReadCount(records []mmio.AccessRecord, offset uint32) int {
	count := 0
	for _, record := range records {
		if record.Type == mmio.AccessRead && record.Offset == offset {
			count++
		}
	}
	return count
}

func sortedUint8Keys(input map[uint8]struct{}) []int {
	result := make([]int, 0, len(input))
	for value := range input {
		result = append(result, int(value))
	}
	sort.Ints(result)
	return result
}

func sortedUint32Keys(input map[uint32]struct{}) []uint32 {
	result := make([]uint32, 0, len(input))
	for value := range input {
		result = append(result, value)
	}
	sort.Slice(result, func(i, j int) bool { return result[i] < result[j] })
	return result
}
