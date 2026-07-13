package main

import (
	"fmt"
	"log/slog"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/donor/behavior"
	"github.com/sercanarga/pcileechgen/internal/firmware"
	"github.com/sercanarga/pcileechgen/internal/pci"
	"github.com/sercanarga/pcileechgen/internal/vivado"
	"github.com/spf13/cobra"
)

// buildFlags groups all build command flags.
type buildFlags struct {
	bdf               string
	board             string
	vivadoPath        string
	output            string
	skipVivado        bool
	jobs              int
	timeout           int
	libDir            string
	fromJSON          string
	stockBar          bool
	behaviorRules     string
	force             bool
	allowStateChanges bool
	profileBARs       bool
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

	if buildOpts.stockBar && ctx.BehaviorRules != nil {
		return fmt.Errorf("behavior rules require generated BAR RTL; --stock-bar is unsupported")
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
		VivadoPath: buildOpts.vivadoPath,
		OutputDir:  buildOpts.output,
		LibDir:     buildOpts.libDir,
		Jobs:       buildOpts.jobs,
		Timeout:    buildOpts.timeout,
		SkipVivado: buildOpts.skipVivado,
		StockBar:   buildOpts.stockBar,
		Force:      buildOpts.force,
	})

	return builder.Build(ctx)
}

// loadDonorContext loads device context from JSON or live device.
func loadDonorContext() (*donor.DeviceContext, error) {
	var ctx *donor.DeviceContext
	var err error
	if buildOpts.fromJSON != "" && buildOpts.bdf != "" {
		return nil, fmt.Errorf("--bdf and --from-json are mutually exclusive")
	}
	if buildOpts.profileBARs && !buildOpts.allowStateChanges {
		return nil, fmt.Errorf("--profile-bars requires --allow-device-state-changes")
	}
	if buildOpts.fromJSON != "" {
		slog.Info("loading device context", "file", buildOpts.fromJSON)
		ctx, err = donor.LoadContext(buildOpts.fromJSON)
		if err != nil {
			return nil, err
		}
	} else {
		if buildOpts.bdf == "" {
			return nil, fmt.Errorf("either --bdf or --from-json is required")
		}
		bdf, parseErr := pci.ParseBDF(buildOpts.bdf)
		if parseErr != nil {
			return nil, fmt.Errorf("invalid BDF: %w", parseErr)
		}
		slog.Info("target device", "bdf", bdf.String())
		slog.Info("collecting donor device data")
		if buildOpts.allowStateChanges {
			slog.Warn("live donor state changes enabled; collection may bind drivers, change power/command registers, or reset the device")
		}
		if buildOpts.profileBARs {
			slog.Warn("destructive BAR profiling enabled; probe writes may trigger irreversible device side effects")
		}
		collector := donor.NewCollectorWithOptions(donor.CollectorOptions{
			AllowStateChanges: buildOpts.allowStateChanges,
			ProfileBARs:       buildOpts.profileBARs,
		})
		ctx, err = collector.Collect(bdf)
		if err != nil {
			return nil, fmt.Errorf("device data collection failed: %w", err)
		}
	}
	if buildOpts.behaviorRules != "" {
		rules, loadErr := behavior.LoadRuleSet(buildOpts.behaviorRules)
		if loadErr != nil {
			return nil, loadErr
		}
		ctx.BehaviorRules = rules
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
	buildCmd.Flags().StringVar(&buildOpts.behaviorRules, "behavior-rules", "", "attach a validated behavior-rule JSON artifact")
	buildCmd.Flags().StringVar(&buildOpts.vivadoPath, "vivado-path", "", "path to Vivado installation")
	buildCmd.Flags().StringVar(&buildOpts.output, "output", "pcileech_datastore", "output directory")
	buildCmd.Flags().BoolVar(&buildOpts.skipVivado, "skip-vivado", false, "skip Vivado synthesis (only generate artifacts)")
	buildCmd.Flags().IntVar(&buildOpts.jobs, "jobs", 4, "number of parallel Vivado jobs")
	buildCmd.Flags().IntVar(&buildOpts.timeout, "timeout", 3600, "Vivado synthesis timeout in seconds")
	buildCmd.Flags().StringVar(&buildOpts.libDir, "lib-dir", "lib/pcileech-fpga", "path to pcileech-fpga library")
	buildCmd.Flags().BoolVar(&buildOpts.stockBar, "stock-bar", false, "use stock bar controller (diagnostic: skip custom SV modules)")
	buildCmd.Flags().BoolVar(&buildOpts.force, "force", false, "ignore donor BAR > board BRAM check")
	buildCmd.Flags().BoolVar(&buildOpts.allowStateChanges, "allow-device-state-changes", false, "allow live collection to change power state, PCI command, driver binding, or reset the donor")
	buildCmd.Flags().BoolVar(&buildOpts.profileBARs, "profile-bars", false, "destructively probe live BAR registers (requires --allow-device-state-changes)")

	_ = buildCmd.MarkFlagRequired("board")

	rootCmd.AddCommand(buildCmd)
}
