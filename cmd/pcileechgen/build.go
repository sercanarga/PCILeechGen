package main

import (
	"fmt"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/donor"
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

	builder := vivado.NewBuilder(b, vivado.BuildOptions{
		VivadoPath: buildOpts.vivadoPath,
		OutputDir:  buildOpts.output,
		LibDir:     buildOpts.libDir,
		Jobs:       buildOpts.jobs,
		Timeout:    buildOpts.timeout,
		SkipVivado: buildOpts.skipVivado,
	})

	return builder.Build(ctx)
}

// loadDonorContext loads device context from JSON or live device.
func loadDonorContext() (*donor.DeviceContext, error) {
	if buildOpts.fromJSON != "" {
		fmt.Printf("[pcileechgen] Loading device context from: %s\n", buildOpts.fromJSON)
		return donor.LoadContext(buildOpts.fromJSON)
	}

	if buildOpts.bdf == "" {
		return nil, fmt.Errorf("either --bdf or --from-json is required")
	}

	bdf, err := pci.ParseBDF(buildOpts.bdf)
	if err != nil {
		return nil, fmt.Errorf("invalid BDF: %w", err)
	}

	fmt.Printf("[pcileechgen] Target device: %s\n", bdf.String())
	fmt.Println("[pcileechgen] Stage 1: Collecting donor device data...")

	collector := donor.NewCollector()
	ctx, err := collector.Collect(bdf)
	if err != nil {
		return nil, fmt.Errorf("device data collection failed: %w", err)
	}
	return ctx, nil
}

func printBuildSummary(ctx *donor.DeviceContext, b *board.Board) {
	fmt.Printf("[pcileechgen] Target board: %s (%s)\n", b.Name, b.FPGAPart)
	fmt.Printf("[pcileechgen] Output: %s\n", buildOpts.output)
	fmt.Printf("[pcileechgen] Device: %04x:%04x %s (rev %02x)\n",
		ctx.Device.VendorID, ctx.Device.DeviceID,
		ctx.Device.ClassDescription(), ctx.Device.RevisionID)
	fmt.Printf("[pcileechgen] Config space: %d bytes\n", ctx.ConfigSpace.Size)
	fmt.Printf("[pcileechgen] Capabilities: %d standard, %d extended\n",
		len(ctx.Capabilities), len(ctx.ExtCapabilities))
	barContentCount := len(ctx.BARContents)
	if barContentCount > 0 {
		fmt.Printf("[pcileechgen] BARs: %d (%d with content)\n\n", len(ctx.BARs), barContentCount)
	} else {
		fmt.Printf("[pcileechgen] BARs: %d\n\n", len(ctx.BARs))
	}
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

	_ = buildCmd.MarkFlagRequired("board")

	rootCmd.AddCommand(buildCmd)
}
