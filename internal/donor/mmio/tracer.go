// Package mmio traces BAR register accesses and extracts access patterns.
package mmio

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

// AccessType indicates read or write.
type AccessType uint8

const (
	AccessRead  AccessType = 0
	AccessWrite AccessType = 1
)

func (a AccessType) String() string {
	if a == AccessWrite {
		return "W"
	}
	return "R"
}

// AccessRecord is one captured BAR read or write.
type AccessRecord struct {
	Offset    uint32        // byte offset within the BAR
	Address   uint64        `json:"address,omitempty"` // physical MMIO address when available
	Type      AccessType    // read or write
	Value     uint32        // value read or written
	Timestamp time.Duration // time since trace started
}

// TraceResult is everything captured during one tracing session.
type TraceResult struct {
	BDF       string        // device BDF (e.g. "0000:03:00.0")
	BARIndex  int           // which BAR was traced
	BARSize   int           // BAR size in bytes
	Duration  time.Duration // total trace duration
	Records   []AccessRecord
	StartTime time.Time
}

// AccessPattern is the analyzed summary - hot regs, polls, init writes.
type AccessPattern struct {
	TotalAccesses int
	TotalReads    int
	TotalWrites   int
	HotRegisters  []HotRegister  // sorted by access count, descending
	PollingLoops  []PollingLoop  // detected polling patterns
	InitSequence  []AccessRecord // first N unique writes (likely initialization)
}

// HotRegister is a register with high access count.
type HotRegister struct {
	Offset     uint32
	ReadCount  int
	WriteCount int
	TotalCount int
	LastValue  uint32
	Values     []uint32 // unique values seen
}

// PollingLoop is a repeated-read pattern on one register.
type PollingLoop struct {
	Offset   uint32
	Count    int           // number of polls
	Interval time.Duration // avg interval between reads
}

type RegisterClassification string

const (
	RegisterStaticRead   RegisterClassification = "static_read"
	RegisterVolatileRead RegisterClassification = "volatile_read"
	RegisterPolled       RegisterClassification = "polled"
	RegisterWriteOnly    RegisterClassification = "write_only"
	RegisterReadWrite    RegisterClassification = "read_write"
)

// RegisterSummary is the stable report form for one observed BAR register.
type RegisterSummary struct {
	Offset         uint32                 `json:"offset"`
	ReadCount      int                    `json:"read_count"`
	WriteCount     int                    `json:"write_count"`
	TotalCount     int                    `json:"total_count"`
	LastValue      uint32                 `json:"last_value"`
	Values         []uint32               `json:"values"`
	Classification RegisterClassification `json:"classification"`
}

// TraceReport is a deterministic, JSON-friendly summary of one MMIO trace.
type TraceReport struct {
	BDF           string            `json:"bdf"`
	BARIndex      int               `json:"bar_index"`
	BARSize       int               `json:"bar_size"`
	DurationMS    int64             `json:"duration_ms"`
	TotalAccesses int               `json:"total_accesses"`
	TotalReads    int               `json:"total_reads"`
	TotalWrites   int               `json:"total_writes"`
	Registers     []RegisterSummary `json:"registers"`
	PollingLoops  []PollingLoop     `json:"polling_loops"`
	InitSequence  []AccessRecord    `json:"init_sequence"`
}

// BuildLegacyTraceReport is a deterministic, JSON-friendly summary of one MMIO
// trace. It reflects raw access patterns only and is used by import/report
// flows that intentionally emit the legacy trace schema.
func BuildLegacyTraceReport(trace *TraceResult) *TraceReport {
	pattern := Analyze(trace)
	report := &TraceReport{
		TotalAccesses: pattern.TotalAccesses,
		TotalReads:    pattern.TotalReads,
		TotalWrites:   pattern.TotalWrites,
		PollingLoops:  append([]PollingLoop(nil), pattern.PollingLoops...),
		InitSequence:  append([]AccessRecord(nil), pattern.InitSequence...),
	}
	if trace != nil {
		report.BDF = trace.BDF
		report.BARIndex = trace.BARIndex
		report.BARSize = trace.BARSize
		report.DurationMS = trace.Duration.Milliseconds()
	}

	polled := make(map[uint32]bool, len(pattern.PollingLoops))
	for _, loop := range pattern.PollingLoops {
		polled[loop.Offset] = true
	}
	for _, hot := range pattern.HotRegisters {
		report.Registers = append(report.Registers, RegisterSummary{
			Offset:         hot.Offset,
			ReadCount:      hot.ReadCount,
			WriteCount:     hot.WriteCount,
			TotalCount:     hot.TotalCount,
			LastValue:      hot.LastValue,
			Values:         append([]uint32(nil), hot.Values...),
			Classification: classifyRegister(hot, polled[hot.Offset]),
		})
	}
	sort.Slice(report.Registers, func(i, j int) bool {
		return report.Registers[i].Offset < report.Registers[j].Offset
	})
	return report
}

func classifyRegister(reg HotRegister, polled bool) RegisterClassification {
	if reg.ReadCount > 0 && reg.WriteCount > 0 {
		return RegisterReadWrite
	}
	if polled {
		return RegisterPolled
	}
	if reg.WriteCount > 0 {
		return RegisterWriteOnly
	}
	if len(reg.Values) > 1 {
		return RegisterVolatileRead
	}
	return RegisterStaticRead
}

// --- analysis ---

// Analyze crunches raw records into hot-register stats, polling loops, and init sequence.
func Analyze(trace *TraceResult) *AccessPattern {
	if trace == nil || len(trace.Records) == 0 {
		return &AccessPattern{}
	}

	pattern := &AccessPattern{
		TotalAccesses: len(trace.Records),
	}

	// per-register hit counts
	type regStats struct {
		reads     int
		writes    int
		values    map[uint32]bool
		last      uint32
		readTimes []time.Duration
	}
	stats := make(map[uint32]*regStats)

	for _, rec := range trace.Records {
		s, ok := stats[rec.Offset]
		if !ok {
			s = &regStats{values: make(map[uint32]bool)}
			stats[rec.Offset] = s
		}
		s.values[rec.Value] = true
		s.last = rec.Value

		if rec.Type == AccessRead {
			pattern.TotalReads++
			s.reads++
			s.readTimes = append(s.readTimes, rec.Timestamp)
		} else {
			pattern.TotalWrites++
			s.writes++
		}
	}

	// build hot-register list
	for off, s := range stats {
		vals := make([]uint32, 0, len(s.values))
		for v := range s.values {
			vals = append(vals, v)
		}
		sort.Slice(vals, func(i, j int) bool { return vals[i] < vals[j] })

		pattern.HotRegisters = append(pattern.HotRegisters, HotRegister{
			Offset:     off,
			ReadCount:  s.reads,
			WriteCount: s.writes,
			TotalCount: s.reads + s.writes,
			LastValue:  s.last,
			Values:     vals,
		})
	}
	sort.Slice(pattern.HotRegisters, func(i, j int) bool {
		return pattern.HotRegisters[i].TotalCount > pattern.HotRegisters[j].TotalCount
	})

	// polling = 3+ consecutive reads to the same offset
	for off, s := range stats {
		if len(s.readTimes) >= 3 {
			var totalInterval time.Duration
			for i := 1; i < len(s.readTimes); i++ {
				totalInterval += s.readTimes[i] - s.readTimes[i-1]
			}
			avgInterval := totalInterval / time.Duration(len(s.readTimes)-1)
			pattern.PollingLoops = append(pattern.PollingLoops, PollingLoop{
				Offset:   off,
				Count:    len(s.readTimes),
				Interval: avgInterval,
			})
		}
	}
	sort.Slice(pattern.PollingLoops, func(i, j int) bool {
		return pattern.PollingLoops[i].Count > pattern.PollingLoops[j].Count
	})

	// init sequence = first unique writes, in order
	seen := make(map[uint32]bool)
	for _, rec := range trace.Records {
		if rec.Type == AccessWrite && !seen[rec.Offset] {
			pattern.InitSequence = append(pattern.InitSequence, rec)
			seen[rec.Offset] = true
		}
		if len(pattern.InitSequence) >= 32 {
			break
		}
	}

	return pattern
}

// FormatReport prints a short text summary of the analysis.
func FormatReport(pattern *AccessPattern) string {
	if pattern == nil {
		return "No trace data.\n"
	}

	var sb strings.Builder
	fmt.Fprintf(&sb, "=== MMIO Trace Analysis ===\n")
	fmt.Fprintf(&sb, "Total accesses: %d (reads: %d, writes: %d)\n\n",
		pattern.TotalAccesses, pattern.TotalReads, pattern.TotalWrites)

	// hot registers (top 10)
	sb.WriteString("--- Hot Registers (top 10) ---\n")
	limit := 10
	if len(pattern.HotRegisters) < limit {
		limit = len(pattern.HotRegisters)
	}
	for i := 0; i < limit; i++ {
		hr := pattern.HotRegisters[i]
		fmt.Fprintf(&sb, "  0x%03X  R:%-4d W:%-4d  last=0x%08X  (%d unique values)\n",
			hr.Offset, hr.ReadCount, hr.WriteCount, hr.LastValue, len(hr.Values))
	}

	if len(pattern.PollingLoops) > 0 {
		sb.WriteString("\n--- Polling Patterns ---\n")
		for _, pl := range pattern.PollingLoops {
			fmt.Fprintf(&sb, "  0x%03X  %d polls, avg interval %v\n",
				pl.Offset, pl.Count, pl.Interval)
		}
	}

	if len(pattern.InitSequence) > 0 {
		sb.WriteString("\n--- Init Sequence (first writes) ---\n")
		for i, rec := range pattern.InitSequence {
			fmt.Fprintf(&sb, "  %2d. [%v] 0x%03X ← 0x%08X\n",
				i+1, rec.Timestamp, rec.Offset, rec.Value)
		}
	}

	return sb.String()
}
