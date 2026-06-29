package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/donor/behavior"
	"github.com/sercanarga/pcileechgen/internal/donor/mmio"
	"github.com/sercanarga/pcileechgen/internal/firmware"
	"github.com/sercanarga/pcileechgen/internal/pci"
	"github.com/sercanarga/pcileechgen/internal/vivado"
	"github.com/spf13/cobra"
)

// buildFlags groups all build command flags.
type buildFlags struct {
	bdf        string
	board      string
	vivadoPath string
	output     string
	skipVivado bool
	jobs       int
	timeout    int
	libDir     string
	fromJSON   string
	stockBar   bool
	force      bool
	mmioTrace  string
	emulatePM  bool
	noVFIO     bool
	optionROM  bool
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

	timing, err := loadTimingHistogram(buildOpts.mmioTrace)
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

	builder := vivado.NewBuilder(b, vivado.BuildOptions{
		VivadoPath:      buildOpts.vivadoPath,
		OutputDir:       buildOpts.output,
		LibDir:          buildOpts.libDir,
		Jobs:            buildOpts.jobs,
		Timeout:         buildOpts.timeout,
		SkipVivado:      buildOpts.skipVivado,
		StockBar:        buildOpts.stockBar,
		Force:           buildOpts.force,
		EmulatePM:       buildOpts.emulatePM,
		TimingHistogram: timing,
	})

	return builder.Build(ctx)
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
	collector.NoVFIO = buildOpts.noVFIO
	collector.CaptureOptionROM = buildOpts.optionROM
	ctx, err := collector.Collect(bdf)
	if err != nil {
		return nil, fmt.Errorf("device data collection failed: %w", err)
	}
	return ctx, nil
}

// loadTimingHistogram reads a saved MMIO trace (mmio.TraceResult JSON, as
// emitted by `mmio-trace --output`) and derives the latency histogram used to
// drive the TLP latency emulator. Returns nil (synthetic defaults) when no
// trace was supplied.
func loadTimingHistogram(path string) (*behavior.TimingHistogram, error) {
	if path == "" {
		return nil, nil
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read mmio trace %q: %w", path, err)
	}
	var tr mmio.TraceResult
	if err := json.Unmarshal(data, &tr); err != nil {
		return nil, fmt.Errorf("parse mmio trace %q (expected JSON from `mmio-trace --output`): %w", path, err)
	}
	h := behavior.ExtractTimingHistogram(&tr)
	if h.SampleCount == 0 {
		slog.Warn("mmio trace yielded no usable read intervals; using synthetic timing", "file", path)
	} else {
		slog.Info("loaded donor timing trace", "file", path, "samples", h.SampleCount)
	}
	return h, nil
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
	buildCmd.Flags().StringVar(&buildOpts.vivadoPath, "vivado-path", "", "path to Vivado installation")
	buildCmd.Flags().StringVar(&buildOpts.output, "output", "pcileech_datastore", "output directory")
	buildCmd.Flags().BoolVar(&buildOpts.skipVivado, "skip-vivado", false, "skip Vivado synthesis (only generate artifacts)")
	buildCmd.Flags().IntVar(&buildOpts.jobs, "jobs", 4, "number of parallel Vivado jobs")
	buildCmd.Flags().IntVar(&buildOpts.timeout, "timeout", 3600, "Vivado synthesis timeout in seconds")
	buildCmd.Flags().StringVar(&buildOpts.libDir, "lib-dir", "lib/pcileech-fpga", "path to pcileech-fpga library")
	buildCmd.Flags().BoolVar(&buildOpts.stockBar, "stock-bar", false, "use stock bar controller (diagnostic: skip custom SV modules)")
	buildCmd.Flags().BoolVar(&buildOpts.force, "force", false, "ignore donor BAR > board BRAM check")
	buildCmd.Flags().StringVar(&buildOpts.mmioTrace, "mmio-trace", "", "saved donor MMIO trace JSON (from `mmio-trace --output`) for measured TLP latency")
	buildCmd.Flags().BoolVar(&buildOpts.emulatePM, "emulate-pm", false, "keep donor Power Management capability faithful (cosmetic D-state; IP core stays D0)")
	buildCmd.Flags().BoolVar(&buildOpts.noVFIO, "no-vfio", false, "sysfs-only donor capture, no vfio-pci bind (for hosts without IOMMU/VFIO)")
	buildCmd.Flags().BoolVar(&buildOpts.optionROM, "option-rom", false, "capture + serve the donor expansion ROM via BAR6 (needs IP expansion-ROM enable; see docs)")

	_ = buildCmd.MarkFlagRequired("board")

	rootCmd.AddCommand(buildCmd)
}
