package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/sercanarga/pcileechgen/internal/color"
	"github.com/sercanarga/pcileechgen/internal/donor/behavior"
	"github.com/sercanarga/pcileechgen/internal/donor/mmio"
	"github.com/sercanarga/pcileechgen/internal/pci"
	"github.com/spf13/cobra"
)

type mmioTraceOptions struct {
	bdf        string
	duration   time.Duration
	barIndex   int
	barSize    int
	barBase    string
	classCode  string
	jsonOutput bool
	outputFile string
	traceFile  string
}

var mmioTraceOpts mmioTraceOptions

var mmioTraceCmd = &cobra.Command{
	Use:   "mmio-trace",
	Short: "Capture live donor MMIO accesses with mmiotrace",
	Long: `Captures MMIO BAR accesses for a short duration using the kernel mmiotrace tracer.

Example:
  pcileechgen mmio-trace --bdf 0000:03:00.0 --duration 5s
  pcileechgen mmio-trace --bdf 03:00.0 --bar-size 4096 --class-code 0x010802
  pcileechgen mmio-trace --trace-file mmiotrace.txt --bar-base 0xf7800000 --bar-index 2 --json`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if mmioTraceOpts.traceFile == "" {
			if _, err := pci.ParseBDF(mmioTraceOpts.bdf); err != nil {
				return fmt.Errorf("invalid BDF %q: %w", mmioTraceOpts.bdf, err)
			}
		}
		if mmioTraceOpts.barIndex < 0 || mmioTraceOpts.barIndex > 5 {
			return fmt.Errorf("--bar-index must be between 0 and 5")
		}
		if mmioTraceOpts.barSize <= 0 || mmioTraceOpts.barSize > 16*1024*1024 {
			return fmt.Errorf("--bar-size must be between 1 and 16MB")
		}
		if mmioTraceOpts.duration <= 0 {
			return fmt.Errorf("--duration must be greater than zero")
		}

		classCode, err := parseTraceClassCode(mmioTraceOpts.classCode)
		if err != nil {
			return err
		}

		barBase, err := parseTraceBARBase(mmioTraceOpts.barBase)
		if err != nil {
			return err
		}

		trace, err := loadMMIOTrace(mmioTraceOpts, barBase)
		if err != nil {
			return err
		}

		pattern := mmio.Analyze(trace)
		profile := behavior.FromMMIOTrace(trace, classCode)
		timing := behavior.ExtractTimingHistogram(trace)

		if mmioTraceOpts.jsonOutput {
			report := map[string]any{
				"trace":            trace,
				"analysis":         pattern,
				"behavior_profile": profile,
				"timing_histogram": timing,
			}
			enc := json.NewEncoder(cmd.OutOrStdout())
			enc.SetIndent("", "  ")
			if err := enc.Encode(report); err != nil {
				return fmt.Errorf("render JSON report: %w", err)
			}
		} else {
			printMMIOTraceReport(cmd.OutOrStdout(), trace, pattern, profile, timing)
		}

		if mmioTraceOpts.outputFile != "" {
			traceData, err := json.MarshalIndent(trace, "", "  ")
			if err != nil {
				return fmt.Errorf("marshal trace: %w", err)
			}
			if err := os.WriteFile(mmioTraceOpts.outputFile, traceData, 0o644); err != nil {
				return fmt.Errorf("write %q: %w", mmioTraceOpts.outputFile, err)
			}
			fmt.Fprintln(cmd.ErrOrStderr(), color.OK(fmt.Sprintf("saved trace to %s", mmioTraceOpts.outputFile)))
		}

		return nil
	},
}

func loadMMIOTrace(opts mmioTraceOptions, barBase uint64) (*mmio.TraceResult, error) {
	if opts.traceFile != "" {
		f, err := os.Open(opts.traceFile)
		if err != nil {
			return nil, fmt.Errorf("open trace file %q: %w", opts.traceFile, err)
		}
		defer f.Close()

		trace, err := mmio.ParseTextTrace(f, mmio.TextTraceOptions{
			BDF:      opts.bdf,
			BARIndex: opts.barIndex,
			BARSize:  opts.barSize,
			BARBase:  barBase,
		})
		if err != nil {
			return nil, fmt.Errorf("parse trace file %q: %w", opts.traceFile, err)
		}
		trace.StartTime = time.Now().Add(-trace.Duration)
		return trace, nil
	}

	start := time.Now()
	trace, err := mmio.LiveTrace(opts.bdf, opts.duration)
	if err != nil {
		return nil, fmt.Errorf("trace capture failed: %w", err)
	}
	trace.BARIndex = opts.barIndex
	trace.BARSize = opts.barSize
	trace.Duration = opts.duration
	trace.StartTime = start
	return trace, nil
}

func parseTraceClassCode(raw string) (uint32, error) {
	if strings.TrimSpace(raw) == "" {
		return 0, nil
	}

	r := strings.TrimSpace(raw)
	if !strings.HasPrefix(r, "0x") && !strings.HasPrefix(r, "0X") {
		r = "0x" + r
	}

	v, err := strconv.ParseUint(r, 0, 24)
	if err != nil {
		return 0, fmt.Errorf("invalid --class-code %q: %w", raw, err)
	}

	return uint32(v), nil
}

func parseTraceBARBase(raw string) (uint64, error) {
	if strings.TrimSpace(raw) == "" {
		return 0, nil
	}

	r := strings.TrimSpace(raw)
	if !strings.HasPrefix(r, "0x") && !strings.HasPrefix(r, "0X") {
		r = "0x" + r
	}

	v, err := strconv.ParseUint(r, 0, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid --bar-base %q: %w", raw, err)
	}

	return v, nil
}

func printMMIOTraceReport(out io.Writer, trace *mmio.TraceResult, pattern *mmio.AccessPattern, profile *behavior.Profile, timing *behavior.TimingHistogram) {
	if out == nil {
		return
	}

	fmt.Fprintln(out, color.Header("MMIO Trace Capture"))
	fmt.Fprintf(out, "BDF: %s\n", trace.BDF)
	fmt.Fprintf(out, "BAR index: %d | bar size: %d\n", trace.BARIndex, trace.BARSize)
	fmt.Fprintf(out, "Duration: %s | records: %d\n\n", trace.Duration, len(trace.Records))

	if len(trace.Records) == 0 {
		fmt.Fprintln(out, color.Warn("No MMIO records were captured.\n"))
		return
	}

	fmt.Fprintln(out, mmio.FormatReport(pattern))

	fmt.Fprintln(out, color.Header("Behavior Profile"))
	fmt.Fprintf(out, "%s\n", behavior.FormatReport(profile))

	fmt.Fprintln(out, color.Header("Timing Histogram"))
	fmt.Fprintf(out, "Samples: %d\n", timing.SampleCount)
	fmt.Fprintf(out, "Min cycles: %d\n", timing.MinCycles)
	fmt.Fprintf(out, "Median cycles: %d\n", timing.MedianCycles)
	fmt.Fprintf(out, "Max cycles: %d\n", timing.MaxCycles)
}

func init() {
	mmioTraceCmd.Flags().StringVar(&mmioTraceOpts.bdf, "bdf", "", "device BDF address (e.g. 0000:03:00.0)")
	mmioTraceCmd.Flags().DurationVar(&mmioTraceOpts.duration, "duration", 5*time.Second, "trace length (e.g. 5s, 1m)")
	mmioTraceCmd.Flags().IntVar(&mmioTraceOpts.barIndex, "bar-index", 0, "BAR index that was targeted for this trace")
	mmioTraceCmd.Flags().IntVar(&mmioTraceOpts.barSize, "bar-size", 4096, "expected BAR size hint in bytes")
	mmioTraceCmd.Flags().StringVar(&mmioTraceOpts.barBase, "bar-base", "", "absolute BAR base for offline trace address normalization")
	mmioTraceCmd.Flags().StringVar(&mmioTraceOpts.classCode, "class-code", "", "PCI class code in hex (e.g. 0x010802)")
	mmioTraceCmd.Flags().StringVar(&mmioTraceOpts.outputFile, "output", "", "save raw trace JSON to file")
	mmioTraceCmd.Flags().StringVar(&mmioTraceOpts.traceFile, "trace-file", "", "analyze an existing mmiotrace text file instead of capturing live")
	mmioTraceCmd.Flags().BoolVar(&mmioTraceOpts.jsonOutput, "json", false, "emit machine-readable report")
	rootCmd.AddCommand(mmioTraceCmd)
}
