package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/sercanarga/pcileechgen/internal/donor/mmio"
	"github.com/spf13/cobra"
)

var traceImportOpts struct {
	input   string
	output  string
	bdf     string
	bar     int
	barSize int
	barBase uint64
}

var traceCmd = &cobra.Command{
	Use:   "trace",
	Short: "Import and summarize MMIO traces",
}

var traceImportCmd = &cobra.Command{
	Use:   "import",
	Short: "Import an mmiotrace log and write bar_model_report.json",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runTraceImport(traceImportOpts.input, traceImportOpts.output, traceImportOpts.bdf, traceImportOpts.bar, traceImportOpts.barSize, traceImportOpts.barBase)
	},
}

type traceImportReport struct {
	Reports []*mmio.TraceReport `json:"reports"`
}

func runTraceImport(inputPath, outputPath, bdf string, barIndex, barSize int, barBase uint64) error {
	f, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("open trace input: %w", err)
	}
	defer f.Close()

	trace, err := mmio.ParseTraceReader(f, mmio.TraceImportOptions{
		BDF:      bdf,
		BARIndex: barIndex,
		BARSize:  barSize,
		BARBase:  barBase,
	})
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(traceImportReport{Reports: []*mmio.TraceReport{mmio.BuildReport(trace)}}, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal trace report: %w", err)
	}

	if dir := filepath.Dir(outputPath); dir != "." && dir != "" {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("create output directory: %w", err)
		}
	}
	if err := os.WriteFile(outputPath, append(data, '\n'), 0644); err != nil {
		return fmt.Errorf("write trace report: %w", err)
	}
	return nil
}

func init() {
	traceImportCmd.Flags().StringVar(&traceImportOpts.input, "input", "", "mmiotrace log path")
	traceImportCmd.Flags().StringVar(&traceImportOpts.output, "output", "bar_model_report.json", "output report path")
	traceImportCmd.Flags().StringVar(&traceImportOpts.bdf, "bdf", "", "donor device BDF for report metadata")
	traceImportCmd.Flags().IntVar(&traceImportOpts.bar, "bar", 0, "BAR index for report metadata")
	traceImportCmd.Flags().IntVar(&traceImportOpts.barSize, "bar-size", 4096, "BAR size for report metadata")
	traceImportCmd.Flags().Uint64Var(&traceImportOpts.barBase, "bar-base", 0, "BAR physical base address for offset calculation")
	_ = traceImportCmd.MarkFlagRequired("input")

	traceCmd.AddCommand(traceImportCmd)
	rootCmd.AddCommand(traceCmd)
}
