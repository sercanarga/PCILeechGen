package barmodel

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sort"

	"github.com/sercanarga/pcileechgen/internal/donor/mmio"
)

// TraceSuggestionOptions gates trace-derived BAR model suggestions. The
// default build path never calls FromTraceReport, so generated HDL from trace
// evidence is DISABLED BY DEFAULT. A caller must explicitly set Enabled=true
// to opt in. ponytail: ceiling -- a single bool flag today; upgrade path adds
// per-BAR opt-in, confidence thresholds, and a reviewer sign-off field.
type TraceSuggestionOptions struct {
	Enabled bool
}

// TraceProvenance records where a trace-derived suggestion came from so a
// reviewer can trace any generated register back to the evidence report.
type TraceProvenance struct {
	// Source is always "trace" for suggestions produced by FromTraceReport.
	Source string `json:"source"`
	// ReportDigest is the lowercase hex SHA-256 of the canonical JSON
	// encoding of the input BarModelReport. Two suggestions with the same
	// digest were derived from identical evidence.
	ReportDigest string `json:"report_digest"`
	// Registers carries per-offset provenance, sorted by Offset. Each entry
	// cites the tag and sample counts that produced the corresponding
	// BARRegister.
	Registers []RegisterProvenance `json:"registers"`
}

// RegisterProvenance is the evidence citation for one emitted register.
type RegisterProvenance struct {
	Offset      uint32 `json:"offset"`
	Tag         string `json:"tag"` // static | rw | counter | polling | volatile | unknown
	ReadCount   int    `json:"read_count"`
	WriteCount  int    `json:"write_count"`
	SampleCount int    `json:"sample_count"` // distinct observed values
}

// TraceSuggestion is a conservative BAR model derived from a trace report,
// together with the provenance trail and any warnings about offsets that were
// intentionally omitted. It is a SUGGESTION only: it is never wired into the
// default generation path. A caller must opt in via TraceSuggestionOptions.
type TraceSuggestion struct {
	Model      *BARModel
	Provenance *TraceProvenance
	Warnings   []mmio.Warning
}

// SuggestFromTrace is the gated entry point. It returns nil (no suggestion,
// no warnings, no error) when opts is nil or opts.Enabled is false. This is
// the function a CLI / build path should call; it keeps trace-derived HDL
// default-OFF without the caller having to remember to branch.
//
// ponytail: ceiling -- the gate is a single bool. Upgrade path: a confidence
// threshold + reviewer sign-off before any register reaches codegen.
func SuggestFromTrace(report *mmio.BarModelReport, opts TraceSuggestionOptions) (*TraceSuggestion, error) {
	if !opts.Enabled {
		return nil, nil
	}
	return FromTraceReport(report)
}

// FromTraceReport converts a trace evidence report (Todo 5's BarModelReport)
// into a conservative BARModel suggestion. It is a pure function over the
// report; it does not touch the default generation path.
//
// Conversion rules (conservative ceiling; see the table in from_trace_test.go):
//
//	static    -> read-only register (RWMask=0), reset = the single observed
//	             value. Safe: host reads are passthrough, host writes ignored.
//	rw        -> RW register (RWMask=0xFFFFFFFF), reset = LastValue. Safe: the
//	             trace already observed the host writing and reading it back.
//	counter   -> OMIT + warn. Counters advance on their own; a frozen reset
//	             value would lie. Unmapped reads return 0 (safe default).
//	polling   -> OMIT + warn. Repeatedly polled RO offset whose value may
//	             change between trace and runtime; freezing it would lie.
//	volatile  -> OMIT + warn. Multiple conflicting read values were observed;
//	             NEVER emit as stable HDL state.
//	unknown   -> OMIT + warn. Insufficient evidence; never assume.
//
// The key guardrail: any offset with conflicting/multiple observed values
// (volatile, counter, polling, unknown) is NEVER emitted as stable HDL state.
// It is omitted so unmapped reads fall through to the safe default (0 / echo).
//
// Determinism: registers are sorted by offset; field order is fixed by the
// struct; the same report yields byte-identical suggestions. Provenance
// includes a SHA-256 digest of the report so two suggestions are comparable.
func FromTraceReport(report *mmio.BarModelReport) (*TraceSuggestion, error) {
	if report == nil {
		return nil, fmt.Errorf("from_trace: nil report")
	}

	digest, err := reportDigest(report)
	if err != nil {
		return nil, fmt.Errorf("from_trace: digest report: %w", err)
	}

	prov := &TraceProvenance{
		Source:       "trace",
		ReportDigest: digest,
	}
	var warnings []mmio.Warning
	var regs []BARRegister

	for _, bar := range report.Bars {
		for _, r := range bar.Registers {
			reg, provEntry, warn, emit := convertRegister(r)
			provEntry.Offset = r.Offset
			provEntry.Tag = r.Tag
			provEntry.ReadCount = r.ReadCount
			provEntry.WriteCount = r.WriteCount
			provEntry.SampleCount = len(r.Values)
			prov.Registers = append(prov.Registers, provEntry)
			if warn != nil {
				warnings = append(warnings, *warn)
			}
			if emit {
				regs = append(regs, *reg)
			}
		}
	}

	// Deterministic order: sort registers by offset, stable.
	sort.SliceStable(regs, func(i, j int) bool {
		return regs[i].Offset < regs[j].Offset
	})
	sort.SliceStable(prov.Registers, func(i, j int) bool {
		return prov.Registers[i].Offset < prov.Registers[j].Offset
	})

	model := &BARModel{Registers: regs}
	// Size is unknown from a trace; leave 0 so the caller (which knows the
	// real BAR size) sets it. ponytail: never guess the BAR size from a trace.
	if len(regs) > 0 {
		validateTraceModel(model)
	}

	return &TraceSuggestion{
		Model:      model,
		Provenance: prov,
		Warnings:   warnings,
	}, nil
}

// convertRegister applies the conservative conversion table for one trace
// register. It returns the BARRegister (nil when omitted), the provenance
// entry, an optional warning, and an emit flag.
func convertRegister(r mmio.RegisterReport) (reg *BARRegister, prov RegisterProvenance, warn *mmio.Warning, emit bool) {
	switch r.Tag {
	case "static":
		// Single observed value, read-only. Safe to freeze.
		return &BARRegister{
			Offset: r.Offset,
			Width:  4,
			Reset:  singleValue(r),
			RWMask: 0,
			Name:   fmt.Sprintf("TRACE_STATIC_0x%08X", r.Offset),
		}, prov, nil, true

	case "rw":
		// Host wrote and read it back. Reset = last observed value so the
		// register powers up in the state the trace saw last.
		return &BARRegister{
			Offset: r.Offset,
			Width:  4,
			Reset:  r.LastValue,
			RWMask: 0xFFFFFFFF,
			Name:   fmt.Sprintf("TRACE_RW_0x%08X", r.Offset),
		}, prov, nil, true

	case "counter", "polling", "volatile":
		w := mmio.Warning{
			Kind: "volatile_omitted",
			Message: fmt.Sprintf(
				"trace tag %q at offset 0x%08X has %d distinct values; omitted from stable HDL state (safe default)",
				r.Tag, r.Offset, len(r.Values)),
		}
		return nil, prov, &w, false

	case "unknown":
		w := mmio.Warning{
			Kind: "unknown_omitted",
			Message: fmt.Sprintf(
				"trace tag %q at offset 0x%08X; insufficient evidence, omitted (safe default)",
				r.Tag, r.Offset),
		}
		return nil, prov, &w, false

	default:
		// Defensive: an unrecognized tag is treated as unknown rather than
		// guessed. ponytail: ceiling -- never assume a new tag's intent.
		w := mmio.Warning{
			Kind: "unknown_omitted",
			Message: fmt.Sprintf(
				"unrecognized trace tag %q at offset 0x%08X; omitted (safe default)",
				r.Tag, r.Offset),
		}
		return nil, prov, &w, false
	}
}

// singleValue returns the single observed value for a static register. When
// the report somehow has zero or multiple values, it returns LastValue as a
// safe fallback and the caller's validateTraceModel / tests catch the
// discrepancy. ponytail: the static tag is only assigned when distinct==1, so
// this branch is defensive, not expected.
func singleValue(r mmio.RegisterReport) uint32 {
	if len(r.Values) == 1 {
		return r.Values[0]
	}
	return r.LastValue
}

// reportDigest is the lowercase hex SHA-256 of the canonical JSON encoding of
// the report. Reuses Todo 3's hashing style (sha256 + hex) without importing
// the provenance package, to keep barmodel a leaf consumer of mmio.
func reportDigest(report *mmio.BarModelReport) (string, error) {
	data, err := json.Marshal(report)
	if err != nil {
		return "", err
	}
	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:]), nil
}

// validateTraceModel checks the invariants the SV codegen assumes, but only
// for the trace-derived subset: DWORD alignment and no duplicate offsets. It
// panics on violation (same contract as validateModel in model.go) so a bad
// converter never silently produces broken HDL.
func validateTraceModel(m *BARModel) {
	seen := make(map[uint32]bool, len(m.Registers))
	for _, r := range m.Registers {
		if r.Offset%4 != 0 {
			panic(fmt.Sprintf("from_trace: register %s at offset 0x%X is not DWORD-aligned", r.Name, r.Offset))
		}
		if seen[r.Offset] {
			panic(fmt.Sprintf("from_trace: duplicate offset 0x%X", r.Offset))
		}
		seen[r.Offset] = true
	}
}