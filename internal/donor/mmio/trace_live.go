// Live MMIO tracer. Uses kernel mmiotrace via ftrace.
// Needs root, CONFIG_MMIOTRACE=y, and debugfs at /sys/kernel/debug.

package mmio

import (
	"bufio"
	"fmt"
	"log"
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
			log.Printf("[mmio] warning: failed to restore tracer: %v", err)
		}
	}()

	// switch to mmiotrace
	if err := os.WriteFile(currentTracer, []byte("mmiotrace"), 0644); err != nil {
		return nil, fmt.Errorf("cannot enable mmiotrace (CONFIG_MMIOTRACE=y needed): %w", err)
	}

	if err := os.WriteFile(tracingOnPath, []byte("1"), 0644); err != nil {
		log.Printf("[mmio] warning: failed to enable tracing: %v", err)
	}
	defer func() {
		if err := os.WriteFile(tracingOnPath, []byte("0"), 0644); err != nil {
			log.Printf("[mmio] warning: failed to disable tracing: %v", err)
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
			rec, ok := parseMMIOTraceLine(line)
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
		log.Printf("[mmio] warning: failed to stop tracing: %v", err)
	}

	return result, nil
}

// parseMMIOTraceLine parses one mmiotrace line (R/W <width> <ts> <addr> <val>).
func parseMMIOTraceLine(line string) (AccessRecord, bool) {
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

	// lower 12 bits = offset within 4K BAR page
	addr, err := strconv.ParseUint(strings.TrimPrefix(fields[3], "0x"), 16, 64)
	if err != nil {
		return AccessRecord{}, false
	}
	rec.Offset = uint32(addr & 0xFFF) // BAR offset within 4K page

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
