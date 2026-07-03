package mmio

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

type TextTraceOptions struct {
	BDF      string
	BARIndex int
	BARSize  int
	BARBase  uint64
	Format   TraceFormat
}

type TraceFormat string

const (
	TraceFormatAuto         TraceFormat = "auto"
	TraceFormatLive         TraceFormat = "mmiotrace"
	TraceFormatMMIO2Verilog TraceFormat = "mmio2verilog"
)

func ParseTextTrace(r io.Reader, opts TextTraceOptions) (*TraceResult, error) {
	if r == nil {
		return nil, fmt.Errorf("trace reader is nil")
	}
	format := opts.Format
	if format == "" {
		format = TraceFormatAuto
	}
	if !format.Valid() {
		return nil, fmt.Errorf("unsupported trace format %q", format)
	}

	trace := &TraceResult{
		BDF:      opts.BDF,
		BARIndex: opts.BARIndex,
		BARSize:  opts.BARSize,
	}

	var first time.Duration
	var last time.Duration
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		rec, ok := parseTextTraceLine(scanner.Text(), opts.BARBase, format)
		if !ok {
			continue
		}
		if len(trace.Records) == 0 {
			first = rec.Timestamp
		}
		last = rec.Timestamp
		trace.Records = append(trace.Records, rec)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("read trace: %w", err)
	}
	if len(trace.Records) == 0 {
		return nil, fmt.Errorf("trace did not contain parseable MMIO records")
	}
	if last >= first {
		trace.Duration = last - first
	}

	return trace, nil
}

func (f TraceFormat) Valid() bool {
	return f == TraceFormatAuto || f == TraceFormatLive || f == TraceFormatMMIO2Verilog
}

func parseTextTraceLine(line string, barBase uint64, format TraceFormat) (AccessRecord, bool) {
	fields := strings.Fields(strings.TrimSpace(line))
	if len(fields) < 5 {
		return AccessRecord{}, false
	}

	var rec AccessRecord
	switch fields[0] {
	case "R":
		rec.Type = AccessRead
	case "W":
		rec.Type = AccessWrite
	default:
		return AccessRecord{}, false
	}

	addrField, valueField, ok := traceAddressValueFields(fields, format)
	if !ok {
		return AccessRecord{}, false
	}

	addr, err := parseHexUint64(addrField)
	if err != nil {
		return AccessRecord{}, false
	}
	value, err := parseHexUint32(valueField)
	if err != nil {
		return AccessRecord{}, false
	}

	rec.Offset = traceOffset(addr, barBase)
	rec.Value = value
	if ts, err := strconv.ParseFloat(fields[2], 64); err == nil {
		rec.Timestamp = time.Duration(ts * float64(time.Second))
	}

	return rec, true
}

func traceAddressValueFields(fields []string, format TraceFormat) (string, string, bool) {
	if format != TraceFormatMMIO2Verilog && strings.HasPrefix(fields[3], "0x") {
		return fields[3], fields[4], true
	}
	if format != TraceFormatLive && len(fields) >= 6 && strings.HasPrefix(fields[4], "0x") {
		return fields[4], fields[5], true
	}
	return "", "", false
}

func parseHexUint64(raw string) (uint64, error) {
	return strconv.ParseUint(strings.TrimPrefix(strings.TrimPrefix(raw, "0x"), "0X"), 16, 64)
}

func parseHexUint32(raw string) (uint32, error) {
	value, err := strconv.ParseUint(strings.TrimPrefix(strings.TrimPrefix(raw, "0x"), "0X"), 16, 32)
	return uint32(value), err
}

func traceOffset(addr uint64, barBase uint64) uint32 {
	if barBase != 0 && addr >= barBase {
		return uint32(addr - barBase)
	}
	return uint32(addr & 0xFFF)
}
