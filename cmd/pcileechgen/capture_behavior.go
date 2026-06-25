package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/donor/behavior"
	"github.com/sercanarga/pcileechgen/internal/donor/mmio"
	"github.com/sercanarga/pcileechgen/internal/pci"
	"github.com/spf13/cobra"
)

type captureBehaviorOptions struct {
	bdf       string
	fromJSON  string
	outPath   string
	duration  int
}

var captureBehaviorOpts captureBehaviorOptions

var captureBehaviorCmd = &cobra.Command{
	Use:   "capture-behavior",
	Short: "Capture donor MMIO behavior profile",
	Long: `Collect BAR MMIO access data from a live donor device and persist
behavior profile (read/write timing + hot/poll metadata) into a context JSON.

Examples:
  pcileechgen capture-behavior --bdf 0000:03:00.0 --seconds 20 --out donor_behavior.json
  pcileechgen capture-behavior --from-json device_context.json --seconds 20 --out donor_context_with_behavior.json`,
	RunE: runCaptureBehavior,
}

func runCaptureBehavior(cmd *cobra.Command, args []string) error {
	if captureBehaviorOpts.duration <= 0 {
		return fmt.Errorf("duration must be positive (seconds)")
	}

	ctx, err := loadOrCollectContextForBehavior(captureBehaviorOpts.bdf, captureBehaviorOpts.fromJSON)
	if err != nil {
		return err
	}

	if ctx.Device.BDF == "" {
		return fmt.Errorf("loaded context is missing donor BDF")
	}

	trace, err := mmio.LiveTrace(ctx.Device.BDF.String(), time.Duration(captureBehaviorOpts.duration)*time.Second)
	if err != nil {
		return fmt.Errorf("capture MMIO behavior: %w", err)
	}

	profile := behavior.FromMMIOTrace(trace, ctx.Device.ClassCode)
	ctx.BehaviorProfile = profile
	ctx.CollectedAt = time.Now()

	if captureBehaviorOpts.outPath == "" {
		captureBehaviorOpts.outPath = "donor_context_with_behavior.json"
	}

	data, err := json.MarshalIndent(ctx, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal behavior context: %w", err)
	}
	if err := os.WriteFile(captureBehaviorOpts.outPath, data, 0o644); err != nil {
		return fmt.Errorf("write %q: %w", captureBehaviorOpts.outPath, err)
	}

	slog.Info("captured behavior profile",
		"out", captureBehaviorOpts.outPath,
		"records", len(trace.Records),
		"duration", trace.Duration,
		"class", fmt.Sprintf("0x%06X", ctx.Device.ClassCode),
	)
	fmt.Println(behavior.FormatReport(profile))
	return nil
}

func loadOrCollectContextForBehavior(bdfArg, fromJSON string) (*donor.DeviceContext, error) {
	if fromJSON != "" {
		return donor.LoadContext(fromJSON)
	}

	if bdfArg == "" {
		return nil, fmt.Errorf("either --bdf or --from-json is required")
	}

	bdf, err := pci.ParseBDF(bdfArg)
	if err != nil {
		return nil, fmt.Errorf("invalid BDF: %w", err)
	}

	collector := donor.NewCollector()
	return collector.Collect(bdf)
}

func init() {
	captureBehaviorCmd.Flags().StringVar(&captureBehaviorOpts.bdf, "bdf", "", "donor device BDF address (e.g. 0000:03:00.0)")
	captureBehaviorCmd.Flags().StringVar(&captureBehaviorOpts.fromJSON, "from-json", "", "base donor context JSON (for offline behavior tagging)")
	captureBehaviorCmd.Flags().IntVar(&captureBehaviorOpts.duration, "seconds", 20, "trace duration in seconds")
	captureBehaviorCmd.Flags().StringVar(&captureBehaviorOpts.outPath, "out", "donor_context_with_behavior.json", "path for output context JSON")
	rootCmd.AddCommand(captureBehaviorCmd)
}
