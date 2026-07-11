package mmio

import (
	"bufio"
	"encoding/json"
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
}

type TraceTarget struct {
	BDF      string
	BARIndex int
	BARBase  uint64
	BARSize  int
}

func ParseTextTrace(r io.Reader, opts TextTraceOptions) (*TraceResult, error) {
	if r == nil {
		return nil, fmt.Errorf("trace reader is nil")
	}

	target := TraceTarget{
		BDF:      opts.BDF,
		BARIndex: opts.BARIndex,
		BARBase:  opts.BARBase,
		BARSize:  opts.BARSize,
	}
	if err := validateTraceTarget(target); err != nil {
		return nil, err
	}

	trace := &TraceResult{
		SchemaVersion: TraceSchemaVersion,
		BDF:           target.BDF,
		BARIndex:      target.BARIndex,
		BARBase:       target.BARBase,
		BARSize:       target.BARSize,
	}

	var first time.Duration
	var last time.Duration
	lineNumber := 0
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lineNumber++
		rec, isMMIO, err := parseTextTraceRecord(scanner.Text(), target)
		if err != nil {
			return nil, fmt.Errorf("parse MMIO trace line %d: %w", lineNumber, err)
		}
		if !isMMIO {
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

func ParseJSONTrace(r io.Reader) (*TraceResult, error) {
	if r == nil {
		return nil, fmt.Errorf("trace reader is nil")
	}
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("read MMIO trace JSON: %w", err)
	}
	var envelope map[string]json.RawMessage
	if err := json.Unmarshal(data, &envelope); err != nil {
		return nil, fmt.Errorf("decode MMIO trace JSON: %w", err)
	}
	if raw, ok := envelope["trace"]; ok {
		data = raw
	}
	var trace TraceResult
	if err := json.Unmarshal(data, &trace); err != nil {
		return nil, fmt.Errorf("decode MMIO trace JSON: %w", err)
	}
	return &trace, nil
}

func parseTextTraceRecord(line string, target TraceTarget) (AccessRecord, bool, error) {
	rec, isMMIO, err := parseRawMMIOTraceLine(line)
	if err != nil || !isMMIO {
		return AccessRecord{}, isMMIO, err
	}
	if err := applyTraceTarget(&rec, target); err != nil {
		return AccessRecord{}, true, err
	}
	return rec, true, nil
}

func parseRawMMIOTraceLine(line string) (AccessRecord, bool, error) {
	fields := strings.Fields(strings.TrimSpace(line))
	if len(fields) == 0 || (fields[0] != "R" && fields[0] != "W") {
		return AccessRecord{}, false, nil
	}
	if len(fields) < 5 {
		return AccessRecord{}, true, fmt.Errorf("incomplete MMIO record")
	}

	var rec AccessRecord
	if fields[0] == "R" {
		rec.Type = AccessRead
	} else {
		rec.Type = AccessWrite
	}

	width, err := strconv.ParseUint(fields[1], 10, 8)
	if err != nil {
		return AccessRecord{}, true, fmt.Errorf("invalid access width %q: %w", fields[1], err)
	}
	rec.Width = uint8(width)
	if !validAccessWidth(rec.Width) {
		return AccessRecord{}, true, fmt.Errorf("unsupported access width %d", rec.Width)
	}

	timestamp, err := strconv.ParseFloat(fields[2], 64)
	if err != nil {
		return AccessRecord{}, true, fmt.Errorf("invalid timestamp %q: %w", fields[2], err)
	}
	rec.Timestamp = time.Duration(timestamp * float64(time.Second))

	addrField, valueField, ok := traceAddressValueFields(fields)
	if !ok {
		return AccessRecord{}, true, fmt.Errorf("missing MMIO address or value")
	}
	rec.Address, err = parseHexUint64(addrField)
	if err != nil {
		return AccessRecord{}, true, fmt.Errorf("invalid address %q: %w", addrField, err)
	}
	rec.Value, err = parseHexUint64(valueField)
	if err != nil {
		return AccessRecord{}, true, fmt.Errorf("invalid value %q: %w", valueField, err)
	}
	if rec.Width < 8 && rec.Value >= uint64(1)<<(rec.Width*8) {
		return AccessRecord{}, true, fmt.Errorf("value %#x does not fit %d-byte access", rec.Value, rec.Width)
	}

	return rec, true, nil
}

func parseTextTraceLine(line string, barBase uint64) (AccessRecord, bool) {
	rec, isMMIO, err := parseRawMMIOTraceLine(line)
	if err != nil || !isMMIO || rec.Address < barBase {
		return AccessRecord{}, false
	}
	offset := rec.Address - barBase
	if offset > uint64(^uint32(0)) {
		return AccessRecord{}, false
	}
	rec.Offset = uint32(offset)
	return rec, true
}

func traceAddressValueFields(fields []string) (string, string, bool) {
	if len(fields) >= 5 && hasHexPrefix(fields[3]) {
		return fields[3], fields[4], true
	}
	if len(fields) >= 6 && hasHexPrefix(fields[4]) {
		return fields[4], fields[5], true
	}
	return "", "", false
}

func hasHexPrefix(raw string) bool {
	return strings.HasPrefix(raw, "0x") || strings.HasPrefix(raw, "0X")
}

func parseHexUint64(raw string) (uint64, error) {
	raw = strings.TrimSpace(raw)
	if !hasHexPrefix(raw) {
		return 0, fmt.Errorf("missing hexadecimal prefix")
	}
	return strconv.ParseUint(raw[2:], 16, 64)
}

func validateTraceTarget(target TraceTarget) error {
	if target.BARIndex < 0 || target.BARIndex > 5 {
		return fmt.Errorf("invalid BAR index %d", target.BARIndex)
	}
	if target.BARSize <= 0 {
		return fmt.Errorf("BAR size must be positive")
	}
	size := uint64(target.BARSize)
	if target.BARBase > ^uint64(0)-size {
		return fmt.Errorf("BAR aperture overflows physical address space: base=%#x size=%#x", target.BARBase, size)
	}
	return nil
}

func applyTraceTarget(rec *AccessRecord, target TraceTarget) error {
	if err := validateTraceTarget(target); err != nil {
		return err
	}
	if rec.Address < target.BARBase {
		return fmt.Errorf("address %#x is below BAR base %#x", rec.Address, target.BARBase)
	}
	offset := rec.Address - target.BARBase
	size := uint64(target.BARSize)
	if offset >= size {
		return fmt.Errorf("address %#x is outside BAR aperture [%#x, %#x)", rec.Address, target.BARBase, target.BARBase+size)
	}
	width := uint64(rec.Width)
	if width > size-offset {
		return fmt.Errorf("%d-byte access at %#x crosses BAR aperture end %#x", rec.Width, rec.Address, target.BARBase+size)
	}
	if rec.Address%width != 0 {
		return fmt.Errorf("%d-byte access at %#x is misaligned", rec.Width, rec.Address)
	}
	if offset > uint64(^uint32(0)) {
		return fmt.Errorf("BAR offset %#x exceeds uint32 range", offset)
	}

	rec.BDF = target.BDF
	rec.BARIndex = target.BARIndex
	rec.Offset = uint32(offset)
	return nil
}
