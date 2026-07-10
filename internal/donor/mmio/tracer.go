// Package mmio traces BAR register accesses and extracts access patterns.
package mmio

import (
	"encoding/json"
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

const TraceSchemaVersion uint32 = 1

func (a AccessType) String() string {
	if a == AccessWrite {
		return "W"
	}
	return "R"
}

// AccessRecord is one captured BAR read or write.
type AccessRecord struct {
	BDF       string
	BARIndex  int
	Address   uint64
	Offset    uint32 // byte offset within the BAR
	Width     uint8
	Type      AccessType // read or write
	Value     uint64
	Timestamp time.Duration // time since trace started
}

// TraceResult is everything captured during one tracing session.
type TraceResult struct {
	SchemaVersion uint32
	BDF           string
	BARIndex      int
	BARBase       uint64
	BARSize       int
	Duration      time.Duration
	Records       []AccessRecord
	StartTime     time.Time
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

type accessRecordJSON struct {
	BDF         string `json:"bdf"`
	BARIndex    int    `json:"bar_index"`
	Address     uint64 `json:"address"`
	Offset      uint32 `json:"offset"`
	Width       uint8  `json:"width"`
	Operation   string `json:"operation"`
	Value       uint64 `json:"value"`
	TimestampNS int64  `json:"timestamp_ns"`
}

type traceResultJSON struct {
	SchemaVersion uint32             `json:"schema_version"`
	BDF           string             `json:"bdf"`
	BARIndex      int                `json:"bar_index"`
	BARBase       uint64             `json:"bar_base"`
	BARSize       int                `json:"bar_size"`
	StartedAt     time.Time          `json:"started_at"`
	DurationNS    int64              `json:"duration_ns"`
	Records       []accessRecordJSON `json:"records"`
}

func (r AccessRecord) MarshalJSON() ([]byte, error) {
	wire, err := r.canonicalJSON()
	if err != nil {
		return nil, err
	}
	return json.Marshal(wire)
}

func (r AccessRecord) canonicalJSON() (accessRecordJSON, error) {
	width := r.Width
	if width == 0 {
		width = 4
	}
	if !validAccessWidth(width) {
		return accessRecordJSON{}, fmt.Errorf("invalid MMIO width %d", width)
	}
	if !valueFitsWidth(r.Value, width) {
		return accessRecordJSON{}, fmt.Errorf("value %#x does not fit %d-byte access", r.Value, width)
	}

	var operation string
	switch r.Type {
	case AccessRead:
		operation = "read"
	case AccessWrite:
		operation = "write"
	default:
		return accessRecordJSON{}, fmt.Errorf("invalid MMIO access type %d", r.Type)
	}

	return accessRecordJSON{
		BDF:         r.BDF,
		BARIndex:    r.BARIndex,
		Address:     r.Address,
		Offset:      r.Offset,
		Width:       width,
		Operation:   operation,
		Value:       r.Value,
		TimestampNS: int64(r.Timestamp),
	}, nil
}

func (r *AccessRecord) UnmarshalJSON(data []byte) error {
	if r == nil {
		return fmt.Errorf("cannot unmarshal MMIO access into nil record")
	}

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var decoded AccessRecord
	if err := unmarshalJSONField(fields, &decoded.BDF, "bdf", "BDF"); err != nil {
		return fmt.Errorf("decode record bdf: %w", err)
	}
	if err := unmarshalJSONField(fields, &decoded.BARIndex, "bar_index", "BARIndex"); err != nil {
		return fmt.Errorf("decode record BAR index: %w", err)
	}
	if err := unmarshalJSONField(fields, &decoded.Address, "address", "Address"); err != nil {
		return fmt.Errorf("decode record address: %w", err)
	}
	if err := unmarshalJSONField(fields, &decoded.Offset, "offset", "Offset"); err != nil {
		return fmt.Errorf("decode record offset: %w", err)
	}

	decoded.Width = 4
	if raw, ok := firstJSONField(fields, "width", "Width"); ok {
		if err := json.Unmarshal(raw, &decoded.Width); err != nil {
			return fmt.Errorf("decode record width: %w", err)
		}
	}
	if !validAccessWidth(decoded.Width) {
		return fmt.Errorf("invalid MMIO width %d", decoded.Width)
	}

	if raw, ok := firstJSONField(fields, "operation"); ok {
		var operation string
		if err := json.Unmarshal(raw, &operation); err != nil {
			return fmt.Errorf("decode record operation: %w", err)
		}
		switch operation {
		case "read":
			decoded.Type = AccessRead
		case "write":
			decoded.Type = AccessWrite
		default:
			return fmt.Errorf("invalid MMIO operation %q", operation)
		}
	} else if err := unmarshalJSONField(fields, &decoded.Type, "type", "Type"); err != nil {
		return fmt.Errorf("decode record access type: %w", err)
	}
	if decoded.Type != AccessRead && decoded.Type != AccessWrite {
		return fmt.Errorf("invalid MMIO access type %d", decoded.Type)
	}

	if err := unmarshalJSONField(fields, &decoded.Value, "value", "Value"); err != nil {
		return fmt.Errorf("decode record value: %w", err)
	}
	if !valueFitsWidth(decoded.Value, decoded.Width) {
		return fmt.Errorf("value %#x does not fit %d-byte access", decoded.Value, decoded.Width)
	}
	var timestampNS int64
	if raw, ok := firstJSONField(fields, "timestamp_ns"); ok {
		if err := json.Unmarshal(raw, &timestampNS); err != nil {
			return fmt.Errorf("decode record timestamp: %w", err)
		}
		decoded.Timestamp = time.Duration(timestampNS)
	} else if err := unmarshalJSONField(fields, &decoded.Timestamp, "timestamp", "Timestamp"); err != nil {
		return fmt.Errorf("decode record timestamp: %w", err)
	}

	*r = decoded
	return nil
}

func (t TraceResult) MarshalJSON() ([]byte, error) {
	records := make([]accessRecordJSON, len(t.Records))
	normalized := make([]AccessRecord, len(t.Records))
	for i, record := range t.Records {
		if record.BDF == "" {
			record.BDF = t.BDF
		}
		if record.BARIndex != t.BARIndex {
			return nil, fmt.Errorf("record %d BAR%d does not match trace BAR%d", i, record.BARIndex, t.BARIndex)
		}
		if record.Address == 0 && t.BARBase != 0 && t.BARBase <= ^uint64(0)-uint64(record.Offset) {
			record.Address = t.BARBase + uint64(record.Offset)
		}
		wire, err := record.canonicalJSON()
		if err != nil {
			return nil, fmt.Errorf("record %d: %w", i, err)
		}
		records[i] = wire
		normalized[i] = record
	}
	validated := t
	validated.Records = normalized
	if err := validated.validateRecords(true); err != nil {
		return nil, err
	}
	return json.Marshal(traceResultJSON{
		SchemaVersion: TraceSchemaVersion,
		BDF:           t.BDF, BARIndex: t.BARIndex, BARBase: t.BARBase, BARSize: t.BARSize,
		StartedAt: t.StartTime, DurationNS: int64(t.Duration), Records: records,
	})
}

func (t *TraceResult) UnmarshalJSON(data []byte) error {
	if t == nil {
		return fmt.Errorf("cannot unmarshal MMIO trace into nil result")
	}

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	if raw, canonical := fields["schema_version"]; canonical {
		var version uint32
		if err := json.Unmarshal(raw, &version); err != nil {
			return fmt.Errorf("decode trace schema version: %w", err)
		}
		if version != TraceSchemaVersion {
			return fmt.Errorf("unsupported MMIO trace schema version %d", version)
		}
		var wire struct {
			BDF        string            `json:"bdf"`
			BARIndex   int               `json:"bar_index"`
			BARBase    uint64            `json:"bar_base"`
			BARSize    int               `json:"bar_size"`
			StartedAt  time.Time         `json:"started_at"`
			DurationNS int64             `json:"duration_ns"`
			Records    []json.RawMessage `json:"records"`
		}
		if err := json.Unmarshal(data, &wire); err != nil {
			return err
		}
		decoded := TraceResult{
			SchemaVersion: version, BDF: wire.BDF, BARIndex: wire.BARIndex, BARBase: wire.BARBase,
			BARSize: wire.BARSize, Duration: time.Duration(wire.DurationNS), StartTime: wire.StartedAt,
			Records: make([]AccessRecord, len(wire.Records)),
		}
		for i, rawRecord := range wire.Records {
			var recordFields map[string]json.RawMessage
			if err := json.Unmarshal(rawRecord, &recordFields); err != nil {
				return fmt.Errorf("record %d: %w", i, err)
			}
			if _, present := recordFields["bar_index"]; !present {
				return fmt.Errorf("record %d: canonical bar_index is required", i)
			}
			if err := json.Unmarshal(rawRecord, &decoded.Records[i]); err != nil {
				return fmt.Errorf("record %d: %w", i, err)
			}
		}
		if err := decoded.validateRecords(true); err != nil {
			return err
		}
		*t = decoded
		return nil
	}

	var legacy struct {
		BDF       string
		BARIndex  int
		BARBase   uint64
		BARSize   int
		Duration  time.Duration
		StartTime time.Time
		Records   []json.RawMessage
	}
	if err := unmarshalJSONField(fields, &legacy.BDF, "BDF", "bdf"); err != nil {
		return err
	}
	if err := unmarshalJSONField(fields, &legacy.BARIndex, "BARIndex", "bar_index"); err != nil {
		return err
	}
	if err := unmarshalJSONField(fields, &legacy.BARBase, "BARBase", "bar_base"); err != nil {
		return err
	}
	if err := unmarshalJSONField(fields, &legacy.BARSize, "BARSize", "bar_size"); err != nil {
		return err
	}
	if err := unmarshalJSONField(fields, &legacy.Duration, "Duration", "duration"); err != nil {
		return err
	}
	if err := unmarshalJSONField(fields, &legacy.StartTime, "StartTime", "started_at"); err != nil {
		return err
	}
	if err := unmarshalJSONField(fields, &legacy.Records, "Records", "records"); err != nil {
		return err
	}
	decoded := TraceResult{
		SchemaVersion: TraceSchemaVersion, BDF: legacy.BDF, BARIndex: legacy.BARIndex,
		BARBase: legacy.BARBase, BARSize: legacy.BARSize, Duration: legacy.Duration,
		StartTime: legacy.StartTime, Records: make([]AccessRecord, len(legacy.Records)),
	}
	for i, rawRecord := range legacy.Records {
		var recordFields map[string]json.RawMessage
		if err := json.Unmarshal(rawRecord, &recordFields); err != nil {
			return fmt.Errorf("record %d: %w", i, err)
		}
		_, hasBARIndex := firstJSONField(recordFields, "BARIndex", "bar_index")
		if err := json.Unmarshal(rawRecord, &decoded.Records[i]); err != nil {
			return fmt.Errorf("record %d: %w", i, err)
		}
		if !hasBARIndex {
			decoded.Records[i].BARIndex = decoded.BARIndex
		}
		if decoded.Records[i].BDF == "" {
			decoded.Records[i].BDF = decoded.BDF
		}
		if decoded.Records[i].Address == 0 && decoded.BARBase != 0 {
			if decoded.BARBase > ^uint64(0)-uint64(decoded.Records[i].Offset) {
				return fmt.Errorf("record %d address overflows", i)
			}
			decoded.Records[i].Address = decoded.BARBase + uint64(decoded.Records[i].Offset)
		}
	}
	if err := decoded.validateRecords(false); err != nil {
		return err
	}
	*t = decoded
	return nil
}

func (t *TraceResult) validateRecords(requireAddress bool) error {
	if t.BARIndex < 0 || t.BARIndex > 5 {
		return fmt.Errorf("invalid trace BAR index %d", t.BARIndex)
	}
	if t.BARSize <= 0 {
		return fmt.Errorf("trace BAR size must be positive")
	}
	size := uint64(t.BARSize)
	if t.BARBase > ^uint64(0)-size {
		return fmt.Errorf("trace BAR aperture overflows physical address space")
	}
	var previous time.Duration
	for i, record := range t.Records {
		if record.BDF != t.BDF {
			return fmt.Errorf("record %d BDF %q does not match trace BDF %q", i, record.BDF, t.BDF)
		}
		if record.BARIndex != t.BARIndex {
			return fmt.Errorf("record %d BAR%d does not match trace BAR%d", i, record.BARIndex, t.BARIndex)
		}
		if !validAccessWidth(record.Width) {
			return fmt.Errorf("record %d: invalid MMIO width %d", i, record.Width)
		}
		if uint64(record.Offset)+uint64(record.Width) > size {
			return fmt.Errorf("record %d access at %#x is outside BAR size %#x", i, record.Offset, t.BARSize)
		}
		if record.Offset%uint32(record.Width) != 0 {
			return fmt.Errorf("record %d access at %#x is misaligned", i, record.Offset)
		}
		expectedAddress := t.BARBase + uint64(record.Offset)
		if requireAddress || (t.BARBase != 0 && record.Address != 0) {
			if record.Address != expectedAddress {
				return fmt.Errorf("record %d address %#x does not match BAR offset address %#x", i, record.Address, expectedAddress)
			}
		}
		if record.Timestamp < 0 || record.Timestamp < previous {
			return fmt.Errorf("record %d timestamp is not monotonic", i)
		}
		previous = record.Timestamp
	}
	return nil
}

func firstJSONField(fields map[string]json.RawMessage, names ...string) (json.RawMessage, bool) {
	for _, name := range names {
		if raw, ok := fields[name]; ok {
			return raw, true
		}
	}
	return nil, false
}

func unmarshalJSONField(fields map[string]json.RawMessage, dst any, names ...string) error {
	raw, ok := firstJSONField(fields, names...)
	if !ok {
		return nil
	}
	return json.Unmarshal(raw, dst)
}

func validAccessWidth(width uint8) bool {
	return width == 1 || width == 2 || width == 4 || width == 8
}

func valueFitsWidth(value uint64, width uint8) bool {
	return width == 8 || (validAccessWidth(width) && value < uint64(1)<<(width*8))
}

// HotRegister is a register with high access count.
type HotRegister struct {
	Offset     uint32
	ReadCount  int
	WriteCount int
	TotalCount int
	LastValue  uint64
	Values     []uint64
}

// PollingLoop is a repeated-read pattern on one register.
type PollingLoop struct {
	Offset   uint32
	Count    int           // number of polls
	Interval time.Duration // avg interval between reads
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
		values    map[uint64]bool
		last      uint64
		readTimes []time.Duration
	}
	stats := make(map[uint32]*regStats)

	for _, rec := range trace.Records {
		s, ok := stats[rec.Offset]
		if !ok {
			s = &regStats{values: make(map[uint64]bool)}
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
		vals := make([]uint64, 0, len(s.values))
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
		fmt.Fprintf(&sb, "  0x%03X  R:%-4d W:%-4d  last=0x%016X  (%d unique values)\n",
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
			fmt.Fprintf(&sb, "  %2d. [%v] 0x%03X ← 0x%016X\n",
				i+1, rec.Timestamp, rec.Offset, rec.Value)
		}
	}

	return sb.String()
}
