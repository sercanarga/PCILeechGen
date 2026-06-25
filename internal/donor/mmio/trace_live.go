// Live MMIO tracer. Uses kernel mmiotrace via ftrace.
// Needs root, CONFIG_MMIOTRACE=y, and debugfs at /sys/kernel/debug.

package mmio

import (
	"bufio"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	tracingDir    = "/sys/kernel/debug/tracing"
	currentTracer = tracingDir + "/current_tracer"
	tracePipe     = tracingDir + "/trace_pipe"
	tracingOnPath = tracingDir + "/tracing_on"
)

type TraceImportOptions struct {
	BDF      string
	BARIndex int
	BARSize  int
	BARBase  uint64
}

func ParseTraceReader(r io.Reader, opts TraceImportOptions) (*TraceResult, error) {
	trace := &TraceResult{
		BDF:      opts.BDF,
		BARIndex: opts.BARIndex,
		BARSize:  opts.BARSize,
	}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		rec, ok := parseMMIOTraceLine(scanner.Text(), opts)
		if !ok {
			continue
		}
		trace.Records = append(trace.Records, rec)
		if rec.Timestamp > trace.Duration {
			trace.Duration = rec.Timestamp
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("read trace: %w", err)
	}
	return trace, nil
}

// LiveTrace enables mmiotrace, records BAR accesses for `duration`, then stops.
func LiveTrace(bdf string, duration time.Duration) (*TraceResult, error) {
	// debugfs present?
	if _, err := os.Stat(currentTracer); err != nil {
		return nil, fmt.Errorf("ftrace not available (debugfs mounted?): %w", err)
	}

	// save + restore previous tracer
	prevTracer, err := os.ReadFile(currentTracer)
	if err != nil {
		return nil, fmt.Errorf("cannot read current tracer: %w", err)
	}
	defer func() {
		if err := os.WriteFile(currentTracer, prevTracer, 0644); err != nil {
			slog.Warn("failed to restore tracer", "error", err)
		}
	}()

	// switch to mmiotrace
	if err := os.WriteFile(currentTracer, []byte("mmiotrace"), 0644); err != nil {
		return nil, fmt.Errorf("cannot enable mmiotrace (CONFIG_MMIOTRACE=y needed): %w", err)
	}

	if err := os.WriteFile(tracingOnPath, []byte("1"), 0644); err != nil {
		slog.Warn("failed to enable tracing", "error", err)
	}
	defer func() {
		if err := os.WriteFile(tracingOnPath, []byte("0"), 0644); err != nil {
			slog.Warn("failed to disable tracing", "error", err)
		}
	}()

	// read from the pipe
	pipe, err := os.Open(tracePipe)
	if err != nil {
		return nil, fmt.Errorf("cannot open trace_pipe: %w", err)
	}
	defer pipe.Close()

	result := &TraceResult{
		BDF:      bdf,
		Duration: duration,
	}

	// read records until timeout
	done := make(chan struct{})
	go func() {
		defer close(done)
		scanner := bufio.NewScanner(pipe)
		deadline := time.Now().Add(duration)

		for scanner.Scan() {
			if time.Now().After(deadline) {
				break
			}
			line := scanner.Text()
			rec, ok := parseMMIOTraceLine(line, TraceImportOptions{})
			if ok {
				result.Records = append(result.Records, rec)
			}
		}
	}()

	// wait or timeout
	select {
	case <-done:
	case <-time.After(duration + 500*time.Millisecond):
	}

	if err := os.WriteFile(tracingOnPath, []byte("0"), 0644); err != nil {
		slog.Warn("failed to stop tracing", "error", err)
	}

	return result, nil
}

// parseMMIOTraceLine parses one mmiotrace line (R/W <width> <ts> <addr> <val>).
func parseMMIOTraceLine(line string, opts TraceImportOptions) (AccessRecord, bool) {
	line = strings.TrimSpace(line)
	// mmiotrace lines typically look like:
	// R 4 1234567.890 0xfee00000 0x00000001 ...
	// W 4 1234567.890 0xfee00000 0x00000001 ...
	fields := strings.Fields(line)
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

	// Keep the full physical address so callers can map it back to the selected
	// BAR resource. Offset remains the legacy low-page offset until remapped.
	addr, err := strconv.ParseUint(strings.TrimPrefix(fields[3], "0x"), 16, 64)
	if err != nil {
		return AccessRecord{}, false
	}
	rec.Address = addr
	offset, ok := traceOffset(addr, opts)
	if !ok {
		return AccessRecord{}, false
	}
	rec.Offset = offset

	// Parse value
	val, err := strconv.ParseUint(strings.TrimPrefix(fields[4], "0x"), 16, 32)
	if err != nil {
		return AccessRecord{}, false
	}
	rec.Value = uint32(val)

	// Parse timestamp if available
	if ts, err := strconv.ParseFloat(fields[2], 64); err == nil {
		rec.Timestamp = time.Duration(ts * float64(time.Second))
	}

	return rec, true
}

func traceOffset(addr uint64, opts TraceImportOptions) (uint32, bool) {
	if opts.BARBase == 0 {
		return uint32(addr & 0xFFF), true
	}
	if addr < opts.BARBase {
		return 0, false
	}
	offset := addr - opts.BARBase
	if opts.BARSize > 0 && offset >= uint64(opts.BARSize) {
		return 0, false
	}
	return uint32(offset), true
}
