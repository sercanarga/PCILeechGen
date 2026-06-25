package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/sercanarga/pcileechgen/internal/color"
	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/donor/mmio"
	"github.com/spf13/cobra"
)

var (
	traceReportContext   string
	traceReportOutput    string
	traceReportOverwrite bool
	traceReportBarBase   string
	traceReportBarSize   string
	traceReportBarIndex  int
)

// traceReportCmd reads a simple mmiotrace-like fixture plus a BAR-bounds
// source (a donor device_context.json or explicit --bar-base/--bar-size
// flags) and writes a deterministic bar_model_report.json. It is a REPORT
// only: it never emits HDL, never modifies the donor context, and never
// generates behavior beyond a conservative tag on observed accesses.
var traceReportCmd = &cobra.Command{
	Use:   "trace-report <trace-file>",
	Short: "Summarize a trace fixture into bar_model_report.json",
	Long: `Parses an mmiotrace-like fixture (R/W <width> <ts> <addr> <val> lines),
maps physical addresses to BAR offsets, and emits a deterministic
bar_model_report.json with per-offset classification tags:

  static | rw | counter | polling | volatile | unknown

BAR bounds come from a donor device_context.json (--context) or from
explicit --bar-base/--bar-size flags. Records outside every BAR window
are quarantined to unknown_regions with a warning; the report is still
produced. This command emits a report only - it never generates HDL.

Examples:
  pcileechgen trace-report trace.log --context device_context.json -o report.json
  pcileechgen trace-report trace.log --bar-base 0x80000000 --bar-size 0x1000 -o report.json`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		tracePath := args[0]
		if traceReportOutput == "" {
			return fmt.Errorf("--output is required (path to bar_model_report.json)")
		}
		if _, err := os.Stat(traceReportOutput); err == nil && !traceReportOverwrite {
			return fmt.Errorf("refusing to overwrite %s without --overwrite", traceReportOutput)
		}

		records, err := mmio.ParseMMIOTraceFile(tracePath)
		if err != nil {
			return err
		}

		bounds, err := resolveBarBounds()
		if err != nil {
			return err
		}

	report, warnings, err := mmio.BuildReport(records, bounds)
		if err != nil {
			return err
		}
		if err := report.WriteJSON(traceReportOutput); err != nil {
			return err
		}

		fmt.Printf("Trace %s -> %s\n", tracePath, color.Bold(traceReportOutput))
		fmt.Printf("In-BAR: %d  Out-of-BAR: %d  Reads: %d  Writes: %d\n",
			report.TotalInBar, report.TotalOutOfBar, report.TotalReads, report.TotalWrites)
		for _, bar := range report.Bars {
			fmt.Printf("  BAR%d [0x%X, +0x%X): %d registers\n",
				bar.BAR, bar.Base, bar.Size, len(bar.Registers))
		}
		if len(warnings) > 0 {
			fmt.Println()
			for _, w := range warnings {
				fmt.Println(color.Warn(w.String()))
			}
		}
		return nil
	},
}

// resolveBarBounds builds the BAR region map from --context, or from explicit
// --bar-base/--bar-size flags. At least one source must be provided.
func resolveBarBounds() (map[int]mmio.BarRegion, error) {
	if traceReportContext != "" {
		ctx, err := donor.LoadContext(traceReportContext)
		if err != nil {
			return nil, fmt.Errorf("load context: %w", err)
		}
		bounds := make(map[int]mmio.BarRegion)
		for _, bar := range ctx.BARs {
			if bar.IsDisabled() || bar.Size == 0 {
				continue
			}
			bounds[bar.Index] = mmio.BarRegion{Base: bar.Address, Size: bar.Size}
		}
		if len(bounds) == 0 {
			return nil, fmt.Errorf("context %s has no enabled memory BARs", traceReportContext)
		}
		return bounds, nil
	}

	if traceReportBarBase == "" || traceReportBarSize == "" {
		return nil, fmt.Errorf("provide --context <device_context.json> or both --bar-base and --bar-size")
	}
	base, err := parseHex(traceReportBarBase)
	if err != nil {
		return nil, fmt.Errorf("--bar-base: %w", err)
	}
	size, err := parseHex(traceReportBarSize)
	if err != nil {
		return nil, fmt.Errorf("--bar-size: %w", err)
	}
	if size == 0 {
		return nil, fmt.Errorf("--bar-size must be non-zero")
	}
	return map[int]mmio.BarRegion{traceReportBarIndex: {Base: base, Size: size}}, nil
}

// parseHex parses a base-10 or 0x-prefixed hex integer.
func parseHex(s string) (uint64, error) {
	return strconv.ParseUint(s, 0, 64)
}

func init() {
	traceReportCmd.Flags().StringVar(&traceReportContext, "context", "", "donor device_context.json for BAR bounds")
	traceReportCmd.Flags().StringVar(&traceReportBarBase, "bar-base", "", "BAR base address (hex 0x.. or decimal), used without --context")
	traceReportCmd.Flags().StringVar(&traceReportBarSize, "bar-size", "", "BAR size in bytes (hex 0x.. or decimal), used without --context")
	traceReportCmd.Flags().IntVar(&traceReportBarIndex, "bar-index", 0, "BAR index when using --bar-base/--bar-size")
	traceReportCmd.Flags().StringVarP(&traceReportOutput, "output", "o", "", "path to write bar_model_report.json (required)")
	traceReportCmd.Flags().BoolVar(&traceReportOverwrite, "overwrite", false, "allow overwriting an existing output file")
	_ = traceReportCmd.MarkFlagRequired("output")
	rootCmd.AddCommand(traceReportCmd)
}
