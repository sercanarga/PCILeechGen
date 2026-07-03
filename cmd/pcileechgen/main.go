package main

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/sercanarga/pcileechgen/internal/color"
	"github.com/spf13/cobra"
)

var (
	logLevel string
	logFile  string
	logJSON  bool

	logFileHandle *os.File
)

func humanWriter() io.Writer {
	if logFileHandle == nil {
		return os.Stdout
	}
	return io.MultiWriter(os.Stdout, logFileHandle)
}

var rootCmd = &cobra.Command{
	Use:   "pcileechgen",
	Short: "PCILeech FPGA firmware generator",
	Long: `PCILeechGen generates custom PCILeech FPGA firmware from real donor PCI/PCIe devices.

It reads the donor device's configuration via VFIO/sysfs, generates firmware artifacts
(.coe, .sv, .tcl), and optionally builds the bitstream using Xilinx Vivado.

This tool requires:
  - Linux with IOMMU/VFIO support (for device reading)
  - A real donor PCI/PCIe card
  - Xilinx Vivado (optional, for bitstream synthesis)`,
	PersistentPreRunE: setupLogging,
}

func init() {
	pf := rootCmd.PersistentFlags()
	pf.StringVar(&logLevel, "log-level", "info", "log verbosity: debug|info|warn|error")
	pf.StringVar(&logFile, "log-file", "", "also write a full log to this file (attach it when reporting Code 10)")
	pf.BoolVar(&logJSON, "json-logs", false, "emit structured JSON logs instead of text")
}

func setupLogging(_ *cobra.Command, _ []string) error {
	levels := map[string]slog.Level{
		"debug": slog.LevelDebug, "info": slog.LevelInfo,
		"warn": slog.LevelWarn, "error": slog.LevelError,
	}
	lvl, ok := levels[logLevel]
	if !ok {
		return fmt.Errorf("invalid --log-level %q (want debug|info|warn|error)", logLevel)
	}

	var w io.Writer = os.Stderr
	if logFile != "" {
		if dir := filepath.Dir(logFile); dir != "." && dir != "" {
			if err := os.MkdirAll(dir, 0755); err != nil {
				return fmt.Errorf("cannot create --log-file directory: %w", err)
			}
		}
		f, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			return fmt.Errorf("cannot open --log-file: %w", err)
		}
		logFileHandle = f
		w = io.MultiWriter(os.Stderr, f)
		color.Disable()
	}

	opts := &slog.HandlerOptions{Level: lvl}
	var h slog.Handler = slog.NewTextHandler(w, opts)
	if logJSON {
		h = slog.NewJSONHandler(w, opts)
	}
	slog.SetDefault(slog.New(h))
	return nil
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
