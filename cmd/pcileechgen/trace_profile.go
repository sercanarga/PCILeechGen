package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/sercanarga/pcileechgen/internal/color"
	"github.com/sercanarga/pcileechgen/internal/donor/behavior"
	"github.com/sercanarga/pcileechgen/internal/donor/mmio"
	"github.com/spf13/cobra"
)

type traceProfileOptions struct {
	tracePath  string
	classCode  string
	jsonOutput bool
}

var traceProfileOpts traceProfileOptions

var traceProfileCmd = &cobra.Command{
	Use:   "trace-profile",
	Short: "Profile an MMIO trace file for behavior reconstruction",
	Long: `Loads an MMIO trace file and emits derived behavior and timing models.

Example:
  pcileechgen trace-profile --trace /tmp/trace.json
  pcileechgen trace-profile --trace /tmp/trace.json --class-code 0x010802`,
	RunE: func(cmd *cobra.Command, args []string) error {
		traceData, err := os.ReadFile(traceProfileOpts.tracePath)
		if err != nil {
			return fmt.Errorf("read trace file %q: %w", traceProfileOpts.tracePath, err)
		}

		var trace mmio.TraceResult
		if err := json.Unmarshal(traceData, &trace); err != nil {
			return fmt.Errorf("parse trace JSON: %w", err)
		}

		classCode, err := parseProfileClassCode(traceProfileOpts.classCode)
		if err != nil {
			return err
		}

		pattern := mmio.Analyze(&trace)
		profile := behavior.FromMMIOTrace(&trace, classCode)
		timing := behavior.ExtractTimingHistogram(&trace)

		if traceProfileOpts.jsonOutput {
			report := map[string]any{
				"trace":            &trace,
				"analysis":         pattern,
				"behavior_profile": profile,
				"timing_histogram": timing,
			}
			enc := json.NewEncoder(cmd.OutOrStdout())
			enc.SetIndent("", "  ")
			if err := enc.Encode(report); err != nil {
				return fmt.Errorf("render JSON report: %w", err)
			}
			return nil
		}

		printTraceProfileReport(cmd.OutOrStdout(), &trace, pattern, profile, timing)
		return nil
	},
}

func printTraceProfileReport(out io.Writer, trace *mmio.TraceResult, pattern *mmio.AccessPattern, profile *behavior.Profile, timing *behavior.TimingHistogram) {
	if out == nil {
		return
	}

	if trace == nil {
		fmt.Fprintln(out, color.Warn("No trace data"))
		return
	}

	fmt.Fprintln(out, color.Header("MMIO Trace Profile"))
	fmt.Fprintf(out, "BDF: %s\n", trace.BDF)
	fmt.Fprintf(out, "BAR index: %d | bar size: %d\n", trace.BARIndex, trace.BARSize)
	fmt.Fprintf(out, "Records: %d | duration: %s\n\n", len(trace.Records), trace.Duration)

	fmt.Fprintln(out, mmio.FormatReport(pattern))
	fmt.Fprintf(out, "\n%s\n", behavior.FormatReport(profile))
	fmt.Fprintln(out, color.Header("Timing Histogram"))
	fmt.Fprintf(out, "Samples: %d\n", timing.SampleCount)
	fmt.Fprintf(out, "Min cycles: %d\n", timing.MinCycles)
	fmt.Fprintf(out, "Median cycles: %d\n", timing.MedianCycles)
	fmt.Fprintf(out, "Max cycles: %d\n", timing.MaxCycles)
}

func parseProfileClassCode(raw string) (uint32, error) {
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

func init() {
	traceProfileCmd.Flags().StringVar(&traceProfileOpts.tracePath, "trace", "", "path to MMIO trace JSON file")
	traceProfileCmd.Flags().StringVar(&traceProfileOpts.classCode, "class-code", "", "PCI class code in hex (e.g. 0x010802)")
	traceProfileCmd.Flags().BoolVar(&traceProfileOpts.jsonOutput, "json", false, "emit machine-readable report")
	_ = traceProfileCmd.MarkFlagRequired("trace")
	rootCmd.AddCommand(traceProfileCmd)
}

