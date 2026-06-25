// Live MMIO tracer. Uses kernel mmiotrace via ftrace.
// Needs root, CONFIG_MMIOTRACE=y, and debugfs at /sys/kernel/debug.

package mmio

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
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
		if writeErr := os.WriteFile(currentTracer, prevTracer, 0644); writeErr != nil {
			slog.Warn("failed to restore tracer", "error", writeErr)
		}
	}()

	// switch to mmiotrace
	if writeErr := os.WriteFile(currentTracer, []byte("mmiotrace"), 0644); writeErr != nil {
		return nil, fmt.Errorf("cannot enable mmiotrace (CONFIG_MMIOTRACE=y needed): %w", writeErr)
	}

	if writeErr := os.WriteFile(tracingOnPath, []byte("1"), 0644); writeErr != nil {
		slog.Warn("failed to enable tracing", "error", writeErr)
	}
	defer func() {
		if writeErr := os.WriteFile(tracingOnPath, []byte("0"), 0644); writeErr != nil {
			slog.Warn("failed to disable tracing", "error", writeErr)
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
		slog.Warn("failed to stop tracing", "error", err)
	}

	return result, nil
}

// parseMMIOTraceLine parses one mmiotrace line (R/W <width> <ts> <addr> <val>).
func parseMMIOTraceLine(line string) (AccessRecord, bool) {
	return parseTextTraceLine(line, 0)
}
