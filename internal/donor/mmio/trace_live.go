// Live MMIO tracer. Uses kernel mmiotrace via ftrace.
// Needs root, CONFIG_MMIOTRACE=y, and debugfs at /sys/kernel/debug.

package mmio

import (
	"bufio"
	"fmt"
	"io"
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

func LiveTrace(bdf string, duration time.Duration) (*TraceResult, error) {
	return nil, fmt.Errorf("target BAR aperture is required for live MMIO capture")
}

func LiveTraceTarget(target TraceTarget, duration time.Duration) (*TraceResult, error) {
	if err := validateTraceTarget(target); err != nil {
		return nil, err
	}
	return captureLiveTrace(&target, target.BDF, duration)
}

type liveTraceOutcome struct {
	record    AccessRecord
	hasRecord bool
	err       error
	done      bool
}

func collectLiveTrace(reader io.ReadCloser, target TraceTarget, duration time.Duration) ([]AccessRecord, error) {
	if reader == nil {
		return nil, fmt.Errorf("trace reader is nil")
	}
	if err := validateTraceTarget(target); err != nil {
		return nil, err
	}
	if duration <= 0 {
		return nil, fmt.Errorf("trace duration must be positive")
	}
	events := make(chan liveTraceOutcome)
	stop := make(chan struct{})
	go func() {
		scanner := bufio.NewScanner(reader)
		var firstTimestamp time.Duration
		haveFirst := false
		for scanner.Scan() {
			rec, isMMIO, parseErr := parseRawMMIOTraceLine(scanner.Text())
			if parseErr != nil || !isMMIO {
				continue
			}
			if applyErr := applyTraceTarget(&rec, target); applyErr != nil {
				continue
			}
			if !haveFirst {
				firstTimestamp = rec.Timestamp
				haveFirst = true
			}
			rec.Timestamp -= firstTimestamp
			select {
			case events <- liveTraceOutcome{record: rec, hasRecord: true}:
			case <-stop:
				return
			}
		}
		select {
		case events <- liveTraceOutcome{err: scanner.Err(), done: true}:
		case <-stop:
		}
	}()
	timer := time.NewTimer(duration)
	defer timer.Stop()
	defer reader.Close()
	var records []AccessRecord
	for {
		select {
		case event := <-events:
			if event.hasRecord {
				records = append(records, event.record)
			}
			if event.done {
				if event.err != nil {
					return nil, fmt.Errorf("read trace pipe: %w", event.err)
				}
				return records, nil
			}
		case <-timer.C:
			close(stop)
			_ = reader.Close()
			grace := time.NewTimer(10 * time.Millisecond)
			select {
			case event := <-events:
				grace.Stop()
				if event.done && event.err != nil {
					return records, nil
				}
			case <-grace.C:
			}
			return records, nil
		}
	}
}

func captureLiveTrace(target *TraceTarget, bdf string, duration time.Duration) (*TraceResult, error) {
	if duration <= 0 {
		return nil, fmt.Errorf("trace duration must be positive")
	}
	if _, err := os.Stat(currentTracer); err != nil {
		return nil, fmt.Errorf("ftrace not available (debugfs mounted?): %w", err)
	}

	prevTracer, err := os.ReadFile(currentTracer)
	if err != nil {
		return nil, fmt.Errorf("cannot read current tracer: %w", err)
	}
	defer func() {
		if writeErr := os.WriteFile(currentTracer, prevTracer, 0644); writeErr != nil {
			slog.Warn("failed to restore tracer", "error", writeErr)
		}
	}()

	if writeErr := os.WriteFile(currentTracer, []byte("mmiotrace"), 0644); writeErr != nil {
		return nil, fmt.Errorf("cannot enable mmiotrace (CONFIG_MMIOTRACE=y needed): %w", writeErr)
	}
	if writeErr := os.WriteFile(tracingOnPath, []byte("1"), 0644); writeErr != nil {
		return nil, fmt.Errorf("cannot enable tracing: %w", writeErr)
	}
	defer func() {
		if writeErr := os.WriteFile(tracingOnPath, []byte("0"), 0644); writeErr != nil {
			slog.Warn("failed to disable tracing", "error", writeErr)
		}
	}()

	pipe, err := os.Open(tracePipe)
	if err != nil {
		return nil, fmt.Errorf("cannot open trace_pipe: %w", err)
	}
	defer pipe.Close()

	startedAt := time.Now()
	result := &TraceResult{
		SchemaVersion: TraceSchemaVersion,
		BDF:           bdf,
		Duration:      duration,
		StartTime:     startedAt,
	}
	if target != nil {
		result.BARIndex = target.BARIndex
		result.BARBase = target.BARBase
		result.BARSize = target.BARSize
	}

	if target == nil {
		return nil, fmt.Errorf("target BAR aperture is required")
	}
	records, collectErr := collectLiveTrace(pipe, *target, duration)
	if collectErr != nil {
		return nil, collectErr
	}
	result.Records = records
	result.Duration = time.Since(startedAt)
	return result, nil
}
