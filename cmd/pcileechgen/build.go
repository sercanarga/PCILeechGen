package main

import (
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/donor/behavior"
	"github.com/sercanarga/pcileechgen/internal/firmware"
	"github.com/sercanarga/pcileechgen/internal/firmware/fallback"
	"github.com/sercanarga/pcileechgen/internal/firmware/output"
	"github.com/sercanarga/pcileechgen/internal/pci"
	"github.com/sercanarga/pcileechgen/internal/vivado"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

//go:embed fallback_defaults.yaml
var fallbackDefaultsYAML []byte

// loadFallbackConfig loads --fallback-config if given, otherwise falls back
// to the built-in embedded defaults (fallback recovery is on by default).
func loadFallbackConfig(path string) (*fallback.Config, error) {
	if path != "" {
		return fallback.LoadConfig(path)
	}
	var cfg fallback.Config
	if err := yaml.Unmarshal(fallbackDefaultsYAML, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse embedded fallback defaults: %w", err)
	}
	return &cfg, nil
}

// buildFlags groups all build command flags.
type buildFlags struct {
	bdf            string
	board          string
	vivadoPath     string
	output         string
	skipVivado     bool
	jobs           int
	timeout        int
	libDir         string
	fromJSON       string
	stockBar       bool
	force          bool
	mmioTrace      string
	ila            bool
	ilaDepth       int
	probeBars      bool
	fallbackConfig string
}

var buildOpts buildFlags

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build firmware from a donor PCI device",
	Long: `Collects device data from a real donor PCI card and generates
PCILeech FPGA firmware artifacts. Optionally synthesizes the
bitstream using Xilinx Vivado.

Use --from-json to build from a previously saved device context
(enables offline builds without direct access to donor hardware).

Example:
  pcileechgen build --bdf 0000:03:00.0 --board PCIeSquirrel
  pcileechgen build --bdf 03:00.0 --board ZDMA --skip-vivado
  pcileechgen build --from-json device_context.json --board PCIeSquirrel
  pcileechgen build --bdf 0000:03:00.0 --board PCIeSquirrel --vivado-path /tools/Xilinx/Vivado/2022.2`,
	RunE: runBuild,
}

func runBuild(cmd *cobra.Command, args []string) error {
	b, err := board.Find(buildOpts.board)
	if err != nil {
		return err
	}

	ctx, err := loadDonorContext()
	if err != nil {
		return err
	}

	printBuildSummary(ctx, b)

	// Gate using donor demand (BARs sizes + BARContents lens + uncapped MSIX req)
	// rather than only ctx.BARs (which may be 0 if resource parse fallback).
	// This prevents silent 4K cap/override when loading 16KB+ donor ctx via --from-json
	// against 4K default BRAM board (e.g. most 35T boards). CappedBAR0Size is still
	// used downstream for actual scrub/COE/SV sizes.
	msixTableSize := 0
	if ctx.MSIXData != nil && ctx.MSIXData.TableSize > 0 {
		msixTableSize = ctx.MSIXData.TableSize
	}
	donorDemand := firmware.DonorBAR0Demand(ctx, b, msixTableSize)
	boardBRAM := b.BRAMSizeOrDefault()
	if donorDemand > boardBRAM {
		if !buildOpts.force {
			return fmt.Errorf("donor largest BAR exceeds board BRAM (%d > %d)", donorDemand, boardBRAM)
		}
		slog.Warn("donor largest BAR exceeds board BRAM (forced)", "donor_bar", donorDemand, "board_bram", boardBRAM)
	}

	bopts := vivado.BuildOptions{
		VivadoPath: buildOpts.vivadoPath,
		OutputDir:  buildOpts.output,
		LibDir:     buildOpts.libDir,
		Jobs:       buildOpts.jobs,
		Timeout:    buildOpts.timeout,
		SkipVivado: buildOpts.skipVivado,
		StockBar:   buildOpts.stockBar,
		Force:      buildOpts.force,
	}

	if buildOpts.mmioTrace != "" {
		h, err := donorTimingHistogram(buildOpts.mmioTrace, ctx)
		if err != nil {
			return fmt.Errorf("load donor timing trace: %w", err)
		}
		bopts.TimingHistogram = h
		slog.Info("donor latency profile loaded", "trace", buildOpts.mmioTrace, "samples", h.SampleCount)
	}

	if buildOpts.ila {
		depth := buildOpts.ilaDepth
		if depth <= 0 {
			depth = firmware.DefaultILADepth
		}
		bopts.ILADepth = depth
		slog.Info("ILA debug core enabled", "depth", depth, "probes", len(firmware.ILAProbes()))
	}

	fbCfg, err := loadFallbackConfig(buildOpts.fallbackConfig)
	if err != nil {
		return fmt.Errorf("load fallback config: %w", err)
	}
	output.DefaultFallbackConfig = fbCfg

	return vivado.NewBuilder(b, bopts).Build(ctx)
}

func donorTimingHistogram(traceFile string, ctx *donor.DeviceContext) (*behavior.TimingHistogram, error) {
	idx := firmware.LargestBarIndex(ctx.BARContents)
	var size int
	var base uint64
	for _, bar := range ctx.BARs {
		if bar.Index == idx {
			size = int(bar.Size)
			base = bar.Address
		}
	}
	trace, err := loadMMIOTrace(mmioTraceOptions{
		bdf:       ctx.Device.BDF.String(),
		barIndex:  idx,
		barSize:   size,
		traceFile: traceFile,
	}, base)
	if err != nil {
		return nil, err
	}
	return behavior.ExtractTimingHistogram(trace), nil
}

// loadDonorContext loads device context from JSON or live device.
func loadDonorContext() (*donor.DeviceContext, error) {
	if buildOpts.fromJSON != "" {
		slog.Info("loading device context", "file", buildOpts.fromJSON)
		return donor.LoadContext(buildOpts.fromJSON)
	}

	if buildOpts.bdf == "" {
		return nil, fmt.Errorf("either --bdf or --from-json is required")
	}

	bdf, err := pci.ParseBDF(buildOpts.bdf)
	if err != nil {
		return nil, fmt.Errorf("invalid BDF: %w", err)
	}

	slog.Info("target device", "bdf", bdf.String())
	slog.Info("collecting donor device data")

	collector := donor.NewCollector()
	collector.ProbeBARs = buildOpts.probeBars
	ctx, err := collector.Collect(bdf)
	if err != nil {
		return nil, fmt.Errorf("device data collection failed: %w", err)
	}
	return ctx, nil
}

func printBuildSummary(ctx *donor.DeviceContext, b *board.Board) {
	slog.Info("build target",
		"board", b.Name,
		"fpga", b.FPGAPart,
		"output", buildOpts.output,
	)
	slog.Info("donor device",
		"vendor", fmt.Sprintf("%04x", ctx.Device.VendorID),
		"device", fmt.Sprintf("%04x", ctx.Device.DeviceID),
		"class", ctx.Device.ClassDescription(),
		"revision", fmt.Sprintf("%02x", ctx.Device.RevisionID),
	)

	csSize := 0
	if ctx.ConfigSpace != nil {
		csSize = ctx.ConfigSpace.Size
	}
	slog.Info("config space",
		"bytes", csSize,
		"std_caps", len(ctx.Capabilities),
		"ext_caps", len(ctx.ExtCapabilities),
		"bars", len(ctx.BARs),
		"bars_with_content", len(ctx.BARContents),
	)
}

func init() {
	buildCmd.Flags().StringVar(&buildOpts.bdf, "bdf", "", "donor device BDF address (e.g. 0000:03:00.0)")
	buildCmd.Flags().StringVar(&buildOpts.board, "board", "", "target FPGA board name (required, e.g. PCIeSquirrel)")
	buildCmd.Flags().StringVar(&buildOpts.fromJSON, "from-json", "", "load donor device data from JSON file (offline build)")
	buildCmd.Flags().StringVar(&buildOpts.mmioTrace, "mmio-trace", "", "donor MMIO trace file; drives TLP latency emulation off the donor's real timing")
	buildCmd.Flags().StringVar(&buildOpts.vivadoPath, "vivado-path", "", "path to Vivado installation")
	buildCmd.Flags().StringVar(&buildOpts.output, "output", "pcileech_datastore", "output directory")
	buildCmd.Flags().BoolVar(&buildOpts.skipVivado, "skip-vivado", false, "skip Vivado synthesis (only generate artifacts)")
	buildCmd.Flags().IntVar(&buildOpts.jobs, "jobs", 4, "number of parallel Vivado jobs")
	buildCmd.Flags().IntVar(&buildOpts.timeout, "timeout", 3600, "Vivado synthesis timeout in seconds")
	buildCmd.Flags().StringVar(&buildOpts.libDir, "lib-dir", "lib/pcileech-fpga", "path to pcileech-fpga library")
	buildCmd.Flags().BoolVar(&buildOpts.stockBar, "stock-bar", false, "use stock bar controller (diagnostic: skip custom SV modules)")
	buildCmd.Flags().BoolVar(&buildOpts.force, "force", false, "ignore donor BAR > board BRAM check")
	buildCmd.Flags().BoolVar(&buildOpts.ila, "ila", false, "insert a Vivado ILA debug core probing BAR/TLP/interrupt signals")
	buildCmd.Flags().IntVar(&buildOpts.ilaDepth, "ila-depth", firmware.DefaultILADepth, "ILA capture depth in samples (with --ila)")
	buildCmd.Flags().BoolVar(&buildOpts.probeBars, "probe-bars", true, "write-probe BAR registers to map RW/W1C bits (always skipped for network cards; set false if a donor hangs on BAR)")
	buildCmd.Flags().StringVar(&buildOpts.fallbackConfig, "fallback-config", "", "path to a custom fallback defaults YAML (fills zeroed BAR registers by device class); empty uses the built-in embedded defaults")

	_ = buildCmd.MarkFlagRequired("board")

	rootCmd.AddCommand(buildCmd)
}
