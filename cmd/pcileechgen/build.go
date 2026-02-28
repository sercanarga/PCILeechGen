package main

import (
	"fmt"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/pci"
	"github.com/sercanarga/pcileechgen/internal/vivado"
	"github.com/spf13/cobra"
)

var (
	buildBDF        string
	buildBoard      string
	buildVivadoPath string
	buildOutput     string
	buildSkipVivado bool
	buildJobs       int
	buildTimeout    int
	buildLibDir     string
	buildFromJSON   string
)

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
	RunE: func(cmd *cobra.Command, args []string) error {
		// Find board
		b, err := board.Find(buildBoard)
		if err != nil {
			return err
		}

		var ctx *donor.DeviceContext

		if buildFromJSON != "" {
			// Offline mode: load from JSON
			fmt.Printf("[pcileechgen] Loading device context from: %s\n", buildFromJSON)
			ctx, err = donor.LoadContext(buildFromJSON)
			if err != nil {
				return fmt.Errorf("failed to load device context: %w", err)
			}
		} else {
			// Live mode: read from donor device
			if buildBDF == "" {
				return fmt.Errorf("either --bdf or --from-json is required")
			}

			bdf, err := pci.ParseBDF(buildBDF)
			if err != nil {
				return fmt.Errorf("invalid BDF: %w", err)
			}

			fmt.Printf("[pcileechgen] Target device: %s\n", bdf.String())
			fmt.Println("[pcileechgen] Stage 1: Collecting donor device data...")

			collector := donor.NewCollector()
			ctx, err = collector.Collect(bdf)
			if err != nil {
				return fmt.Errorf("device data collection failed: %w", err)
			}
		}

		fmt.Printf("[pcileechgen] Target board: %s (%s)\n", b.Name, b.FPGAPart)
		fmt.Printf("[pcileechgen] Output: %s\n", buildOutput)
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

		// Stage 2 & 3: Build
		builder := vivado.NewBuilder(b, vivado.BuildOptions{
			VivadoPath: buildVivadoPath,
			OutputDir:  buildOutput,
			LibDir:     buildLibDir,
			Jobs:       buildJobs,
			Timeout:    buildTimeout,
			SkipVivado: buildSkipVivado,
		})

		return builder.Build(ctx)
	},
}

func init() {
	buildCmd.Flags().StringVar(&buildBDF, "bdf", "", "donor device BDF address (e.g. 0000:03:00.0)")
	buildCmd.Flags().StringVar(&buildBoard, "board", "", "target FPGA board name (required, e.g. PCIeSquirrel)")
	buildCmd.Flags().StringVar(&buildFromJSON, "from-json", "", "load donor device data from JSON file (offline build)")
	buildCmd.Flags().StringVar(&buildVivadoPath, "vivado-path", "", "path to Vivado installation")
	buildCmd.Flags().StringVar(&buildOutput, "output", "pcileech_datastore", "output directory")
	buildCmd.Flags().BoolVar(&buildSkipVivado, "skip-vivado", false, "skip Vivado synthesis (only generate artifacts)")
	buildCmd.Flags().IntVar(&buildJobs, "jobs", 4, "number of parallel Vivado jobs")
	buildCmd.Flags().IntVar(&buildTimeout, "timeout", 3600, "Vivado synthesis timeout in seconds")
	buildCmd.Flags().StringVar(&buildLibDir, "lib-dir", "lib/pcileech-fpga", "path to pcileech-fpga library")

	_ = buildCmd.MarkFlagRequired("board")

	rootCmd.AddCommand(buildCmd)
}
