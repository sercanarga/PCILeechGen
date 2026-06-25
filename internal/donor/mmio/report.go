package mmio

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

// BarRegion is the physical address window for one BAR.
type BarRegion struct {
	Base uint64
	Size uint64
}

// regAcc accumulates per-offset statistics during report building.
type regAcc struct {
	reads    int
	writes   int
	readVals []uint32 // in observation order
	allVals  map[uint32]struct{}
	last     uint32
}

// Warning is a structured, deterministic warning emitted by report builders.
type Warning struct {
	Kind    string `json:"kind"`
	Message string `json:"message"`
}

// String renders a warning as "kind: message".
func (w Warning) String() string { return fmt.Sprintf("%s: %s", w.Kind, w.Message) }

// RegisterReport summarizes observed behavior at one BAR offset.
type RegisterReport struct {
	BAR        int      `json:"bar"`
	Offset     uint32   `json:"offset"`
	Tag        string   `json:"tag"` // static | rw | counter | polling | volatile | unknown
	ReadCount  int      `json:"read_count"`
	WriteCount int      `json:"write_count"`
	Hot        bool     `json:"hot"`
	Values     []uint32 `json:"values,omitempty"` // sorted unique observed values
	LastValue  uint32   `json:"last_value"`
}

// UnknownRegion is a trace record that fell outside every provided BAR window.
type UnknownRegion struct {
	Address uint64 `json:"address"`
	Offset  uint32 `json:"offset,omitempty"`
	Type    string `json:"type"` // "R" | "W"
	Value   uint32 `json:"value"`
}

// BarReport is the per-BAR slice of the evidence report.
type BarReport struct {
	BAR       int              `json:"bar"`
	Base      uint64           `json:"base"`
	Size      uint64           `json:"size"`
	Registers []RegisterReport `json:"registers"`
}

// BarModelReport is the structured BAR evidence report. It is a REPORT only:
// it describes what the trace observed and never dictates aggressive
// generated behavior. Downstream consumers (e.g. barmodel) must treat every
// tag as a hint, not ground truth.
type BarModelReport struct {
	Bars          []BarReport     `json:"bars"`
	Unknown       []UnknownRegion `json:"unknown_regions,omitempty"`
	TotalReads    int             `json:"total_reads"`
	TotalWrites   int             `json:"total_writes"`
	TotalInBar    int             `json:"total_in_bar"`
	TotalOutOfBar int             `json:"total_out_of_bar"`
}

// Classification thresholds. ponytail: these are naive sample-count heuristics,
// not statistically grounded. Upgrade path: track variance, time-windowed
// stability, and confidence intervals before trusting any tag for codegen.
const (
	// hotThreshold: an offset with this many total accesses is "hot".
	hotThreshold = 8
	// pollingMinReads: this many reads of one stable read-only offset => polling.
	pollingMinReads = 4
	// minSamples: below this we refuse to guess and tag "unknown".
	minSamples = 2
)

// BuildReport maps raw trace records into a deterministic BarModelReport using
// the provided BAR bounds. Records outside every BAR window are quarantined
// into UnknownRegions and surfaced via warnings - they never cause an error.
// An empty barBounds map yields only quarantined records.
func BuildReport(records []AccessRecord, barBounds map[int]BarRegion) (*BarModelReport, []Warning, error) {
	report := &BarModelReport{}
	var warnings []Warning

	// Ordered BAR indices for deterministic output.
	bars := make([]int, 0, len(barBounds))
	for idx := range barBounds {
		bars = append(bars, idx)
	}
	sort.Ints(bars)

	// per-(bar,offset) accumulator
	accs := make(map[int]map[uint32]*regAcc)
	for _, b := range bars {
		accs[b] = make(map[uint32]*regAcc)
	}

	for _, rec := range records {
		report.TotalReads += boolToInt(rec.Type == AccessRead)
		report.TotalWrites += boolToInt(rec.Type == AccessWrite)

		bar, off, inBar := locate(rec, barBounds)
		if !inBar {
			report.TotalOutOfBar++
			report.Unknown = append(report.Unknown, UnknownRegion{
				Address: rec.Address,
				Offset:  rec.Offset,
				Type:    rec.Type.String(),
				Value:   rec.Value,
			})
			continue
		}
		report.TotalInBar++

		m := accs[bar]
		a, ok := m[off]
		if !ok {
			a = &regAcc{allVals: make(map[uint32]struct{})}
			m[off] = a
		}
		a.allVals[rec.Value] = struct{}{}
		a.last = rec.Value
		if rec.Type == AccessRead {
			a.reads++
			a.readVals = append(a.readVals, rec.Value)
		} else {
			a.writes++
		}
	}

	if len(report.Unknown) > 0 {
		warnings = append(warnings, Warning{
			Kind:    "out_of_bar",
			Message: fmt.Sprintf("%d trace records fell outside all BAR windows; quarantined to unknown_regions", len(report.Unknown)),
		})
	}

	// Build per-BAR register lists sorted by offset.
	for _, bar := range bars {
		region := barBounds[bar]
		br := BarReport{BAR: bar, Base: region.Base, Size: region.Size}
		for off, a := range accs[bar] {
			br.Registers = append(br.Registers, classifyRegister(bar, off, a))
		}
		sort.Slice(br.Registers, func(i, j int) bool {
			return br.Registers[i].Offset < br.Registers[j].Offset
		})
		report.Bars = append(report.Bars, br)
	}

	return report, warnings, nil
}

// locate maps a record to (bar, offset) using physical Address when present,
// falling back to the legacy Offset field against each BAR's size. ponytail:
// the Address==0 fallback is a guess; once all capture formats emit real
// physical addresses, drop the fallback and require Address.
func locate(rec AccessRecord, barBounds map[int]BarRegion) (int, uint32, bool) {
	if rec.Address != 0 {
		for bar, r := range barBounds {
			if r.Size == 0 {
				continue
			}
			if rec.Address >= r.Base && rec.Address < r.Base+r.Size {
				return bar, uint32(rec.Address - r.Base), true
			}
		}
		return 0, 0, false
	}
	// Legacy offset-only record: assume it belongs to the smallest BAR that
	// contains the offset. Most fixtures have a single BAR so this is fine.
	for bar, r := range barBounds {
		if r.Size == 0 {
			continue
		}
		if uint64(rec.Offset) < r.Size {
			return bar, rec.Offset, true
		}
	}
	return 0, 0, false
}

// classifyRegister applies the conservative tag heuristic. ponytail: order of
// checks encodes priority; do not reorder without revisiting semantics.
func classifyRegister(bar int, off uint32, a *regAcc) RegisterReport {
	vals := make([]uint32, 0, len(a.allVals))
	for v := range a.allVals {
		vals = append(vals, v)
	}
	sort.Slice(vals, func(i, j int) bool { return vals[i] < vals[j] })

	total := a.reads + a.writes
	rep := RegisterReport{
		BAR:        bar,
		Offset:     off,
		ReadCount:  a.reads,
		WriteCount: a.writes,
		Hot:        total >= hotThreshold,
		Values:     vals,
		LastValue:  a.last,
	}

	rep.Tag = tagRegister(a)
	return rep
}

// tagRegister is the pure classification rule. ponytail: naive thresholds; see
// hotThreshold/pollingMinReads/minSamples constants for the ceiling.
func tagRegister(a *regAcc) string {
	written := a.writes > 0
	read := a.reads > 0

	if written && read {
		return "rw"
	}
	if written && !read {
		// Write-only register; no readback evidence. Refuse to over-claim.
		return "unknown"
	}
	if !read {
		// No accesses at all - should not happen, but stay safe.
		return "unknown"
	}

	// Read-only from here.
	distinct := len(a.allVals)
	if distinct == 1 {
		// ponytail: polling vs static split is a count heuristic, not a
		// temporal one. A real polling detector would use inter-arrival
		// timing from rec.Timestamp.
		if a.reads >= pollingMinReads {
			return "polling"
		}
		return "static"
	}
	// Multiple distinct read values: need enough samples to judge pattern.
	// ponytail: minSamples ceiling - below this, "unknown" beats a guess.
	if a.reads < minSamples {
		return "unknown"
	}
	if isMonotonic(a.readVals) {
		return "counter"
	}
	return "volatile"
}

// isMonotonic reports whether the read sequence is strictly increasing or
// strictly decreasing. ponytail: strict monotonicity is too tight for noisy
// counters; relax to non-strict once we have enough samples to tolerate dups.
func isMonotonic(vals []uint32) bool {
	if len(vals) < 2 {
		return false
	}
	inc, dec := true, true
	for i := 1; i < len(vals); i++ {
		if vals[i] <= vals[i-1] {
			inc = false
		}
		if vals[i] >= vals[i-1] {
			dec = false
		}
	}
	return inc || dec
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

// WriteJSON writes the report as deterministic JSON: struct field order fixes
// the key order, and Bars/Registers are pre-sorted by BuildReport. Two runs
// over the same input produce byte-identical output.
func (r *BarModelReport) WriteJSON(path string) error {
	data, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal report: %w", err)
	}
	data = append(data, '\n')
	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("write %s: %w", path, err)
	}
	return nil
}

// ParseMMIOTraceFile reads a simple mmiotrace-like line format
//
//	R <width> <ts> <addr> <val>
//	W <width> <ts> <addr> <val>
//
// and returns the parsed records in file order. Unrecognized lines are
// skipped silently (mmiotrace output includes headers and annotations that
// are not records). ponytail: reuses parseMMIOTraceLine from trace_live.go
// rather than spawning a parallel parser.
func ParseMMIOTraceFile(path string) ([]AccessRecord, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open trace: %w", err)
	}
	defer f.Close()

	var out []AccessRecord
	sc := bufio.NewScanner(f)
	sc.Buffer(make([]byte, 64*1024), 1024*1024)
	lineNo := 0
	for sc.Scan() {
		lineNo++
		line := strings.TrimSpace(sc.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		rec, ok := parseMMIOTraceLine(line, TraceImportOptions{})
		if !ok {
			// Skip non-record lines (mmiotrace headers, annotations).
			continue
		}
		out = append(out, rec)
	}
	if err := sc.Err(); err != nil {
		return nil, fmt.Errorf("read trace at line %d: %w", lineNo, err)
	}
	return out, nil
}
