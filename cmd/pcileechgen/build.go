package main

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/donor/mmio"
	"github.com/sercanarga/pcileechgen/internal/firmware"
	"github.com/sercanarga/pcileechgen/internal/firmware/variance"
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
	force             bool
	publicSafe        bool
	nvmeIdentifyCtrl  string
	nvmeIdentifyNS    string
	traces            []string
	mmiotraceBar      int
	mmiotraceDuration time.Duration
}

var buildOpts buildFlags

var liveTraceFn = mmio.LiveTrace

func unsafeLiveDonorReason(dev *pci.PCIDevice) string {
	if dev == nil || dev.Driver == "" || dev.Driver == "vfio-pci" {
		return ""
	}
	switch dev.Driver {
	case "iwlwifi":
		return "active Wi-Fi device bound to iwlwifi"
	case "nvme":
		return "active NVMe device bound to nvme"
	case "xhci_hcd":
		return "active USB controller bound to xhci_hcd"
	case "i915", "xe":
		return "active GPU/display device"
	case "thunderbolt":
		return "active Thunderbolt/USB4 controller"
	case "snd_hda_intel", "snd_sof_pci_intel_mtl":
		return "active HD Audio device"
	case "vmd":
		return "active VMD storage fabric"
	}
	switch {
	case dev.BaseClass() == 0x02:
		return "active network device"
	case dev.BaseClass() == 0x03:
		return "active display device"
	case dev.BaseClass() == 0x01:
		return "active storage device"
	case dev.BaseClass() == 0x0C && dev.SubClass() == 0x03:
		return "active USB-class controller"
	case dev.BaseClass() == 0x04 && dev.SubClass() == 0x03:
		return "active audio device"
	}
	return ""
}

func filterOutDSNExtCaps(extCaps []pci.ExtCapability) []pci.ExtCapability {
	if len(extCaps) == 0 {
		return nil
	}
	out := make([]pci.ExtCapability, 0, len(extCaps))
	for _, cap := range extCaps {
		if cap.ID != pci.ExtCapIDDeviceSerialNumber {
			out = append(out, cap)
		}
	}
	return out
}

func applyPublicSafeMode(ctx *donor.DeviceContext) {
	if ctx == nil {
		return
	}
	if ctx.ConfigSpace != nil {
		variance.StripDSNExtCap(ctx.ConfigSpace)
	}
	ctx.ExtCapabilities = filterOutDSNExtCaps(ctx.ExtCapabilities)
	ctx.Hostname = ""
	ctx.Device.BDF = pci.BDF{}
	if ctx.NVMeIdentify != nil {
		slog.Warn("public-safe mode: dropping donor NVMe identify capture")
		ctx.NVMeIdentify = nil
	}
}

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
	if buildOpts.fromJSON == "" {
		if err := maybeCaptureMMIOTrace(ctx, buildOpts.bdf, buildOpts.mmiotraceDuration); err != nil {
			return err
		}
	}

	if err := applyNVMeIdentifyFiles(ctx); err != nil {
		return err
	}
	if buildOpts.publicSafe {
		applyPublicSafeMode(ctx)
	}

	if err := applyBuildTraceSpecs(ctx, buildOpts.traces); err != nil {
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

	sr := donor.NewSysfsReader()
	if dev, err := sr.ReadDeviceInfo(bdf); err == nil && buildOpts.mmiotraceDuration > 0 {
		if reason := unsafeLiveDonorReason(dev); reason != "" {
			return nil, fmt.Errorf("refusing invasive live MMIO trace for %s (%04x:%04x, driver %q): %s. Use --from-json with an offline donor dump instead",
				bdf.String(), dev.VendorID, dev.DeviceID, dev.Driver, reason)
		}
	}

	slog.Info("target device", "bdf", bdf.String())
	slog.Info("collecting donor device data")

	collector := donor.NewCollector()
	ctx, err := collector.Collect(bdf)
	if err != nil {
		return nil, fmt.Errorf("device data collection failed: %w", err)
	}
	return ctx, nil
}

func parseBuildTraceSpec(spec string) (int, string, error) {
	barText, path, ok := strings.Cut(spec, "=")
	if !ok || barText == "" || path == "" {
		return 0, "", fmt.Errorf("invalid trace spec %q (want BAR=path, e.g. 2=trace.log)", spec)
	}
	bar, err := strconv.Atoi(barText)
	if err != nil || bar < 0 || bar > 5 {
		return 0, "", fmt.Errorf("invalid trace BAR %q in %q", barText, spec)
	}
	return bar, path, nil
}

func applyBuildTraceSpecs(ctx *donor.DeviceContext, specs []string) error {
	if len(specs) == 0 {
		return nil
	}
	if ctx.MMIOTraces == nil {
		ctx.MMIOTraces = make(map[int]*mmio.TraceResult)
	}
	for _, spec := range specs {
		bar, path, err := parseBuildTraceSpec(spec)
		if err != nil {
			return err
		}
		f, err := os.Open(path)
		if err != nil {
			return fmt.Errorf("open BAR%d trace %q: %w", bar, path, err)
		}
		barSize := buildTraceBARSize(ctx, bar)
		barBase := buildTraceBARBase(ctx, bar)
		trace, parseErr := mmio.ParseTraceReader(f, mmio.TraceImportOptions{
			BDF:      ctx.Device.BDF.String(),
			BARIndex: bar,
			BARSize:  barSize,
			BARBase:  barBase,
		})
		closeErr := f.Close()
		if parseErr != nil {
			return fmt.Errorf("parse BAR%d trace %q: %w", bar, path, parseErr)
		}
		if closeErr != nil {
			return fmt.Errorf("close BAR%d trace %q: %w", bar, path, closeErr)
		}
		trace, err = mapTraceToSelectedBAR(ctx, bar, trace)
		if err != nil {
			return fmt.Errorf("map BAR%d trace %q: %w", bar, path, err)
		}
		attachTraceEvidenceToContext(ctx, bar, trace)
	}
	return nil
}

func buildTraceBARSize(ctx *donor.DeviceContext, barIndex int) int {
	for _, bar := range ctx.BARs {
		if bar.Index == barIndex && bar.Size > 0 {
			return int(bar.Size)
		}
	}
	if data := ctx.BARContents[barIndex]; len(data) > 0 {
		return len(data)
	}
	return 4096
}

func buildTraceBARBase(ctx *donor.DeviceContext, barIndex int) uint64 {
	for _, bar := range ctx.BARs {
		if bar.Index == barIndex {
			return bar.Address
		}
	}
	return 0
}

func maybeCaptureMMIOTrace(ctx *donor.DeviceContext, bdf string, d time.Duration) error {
	if d <= 0 {
		return nil
	}
	trace, err := liveTraceFn(bdf, d)
	if err != nil {
		return err
	}
	barIdx := buildOpts.mmiotraceBar
	if barIdx < 0 {
		barIdx = firmware.LargestBarIndex(ctx.BARContents)
	}
	trace, err = mapTraceToSelectedBAR(ctx, barIdx, trace)
	if err != nil {
		return err
	}
	attachTraceEvidenceToContext(ctx, barIdx, trace)
	return nil
}

func attachTraceEvidenceToContext(ctx *donor.DeviceContext, barIdx int, trace *mmio.TraceResult) {
	if ctx == nil || trace == nil {
		return
	}
	if ctx.MMIOTraces == nil {
		ctx.MMIOTraces = make(map[int]*mmio.TraceResult)
	}
	ctx.MMIOTraces[barIdx] = trace

	overlay := mmio.DeriveTraceBAROverlay(trace)
	if ctx.BARTraceOverlays == nil {
		ctx.BARTraceOverlays = make(map[int]*mmio.TraceBAROverlay)
	}
	if overlay != nil {
		ctx.BARTraceOverlays[barIdx] = overlay
	}
}

func mapTraceToSelectedBAR(ctx *donor.DeviceContext, barIdx int, trace *mmio.TraceResult) (*mmio.TraceResult, error) {
	if ctx == nil || trace == nil {
		return trace, nil
	}
	for _, bar := range ctx.BARs {
		if bar.Index == barIdx && bar.Address != 0 && bar.Size != 0 {
			return mmio.RemapTraceToBAROffsets(trace, bar.Address, bar.Size)
		}
	}
	if data := ctx.BARContents[barIdx]; len(data) > 4096 {
		return nil, fmt.Errorf("mmiotrace for BAR%d needs a sysfs resource address when BAR size exceeds 4 KiB", barIdx)
	}
	return trace, nil
}

// applyNVMeIdentifyFiles loads optional donor-captured NVMe Identify
// Controller / Namespace pages from --nvme-identify-ctrl / --nvme-identify-ns
// binary files (each exactly 4096 bytes) into the device context so the
// generated firmware mirrors the real donor's model/serial/FRU fields.
func applyNVMeIdentifyFiles(ctx *donor.DeviceContext) error {
	if buildOpts.nvmeIdentifyCtrl == "" && buildOpts.nvmeIdentifyNS == "" {
		return nil
	}
	if ctx.NVMeIdentify == nil {
		ctx.NVMeIdentify = &donor.NVMeIdentifyCapture{}
	}
	if buildOpts.nvmeIdentifyCtrl != "" {
		b, err := os.ReadFile(buildOpts.nvmeIdentifyCtrl)
		if err != nil {
			return fmt.Errorf("read NVMe identify controller: %w", err)
		}
		if len(b) != 4096 {
			return fmt.Errorf("NVMe identify controller file must be exactly 4096 bytes, got %d", len(b))
		}
		copy(ctx.NVMeIdentify.Controller[:], b)
		ctx.NVMeIdentify.HasController = true
		slog.Info("loaded donor NVMe identify controller", "file", buildOpts.nvmeIdentifyCtrl)
	}
	if buildOpts.nvmeIdentifyNS != "" {
		b, err := os.ReadFile(buildOpts.nvmeIdentifyNS)
		if err != nil {
			return fmt.Errorf("read NVMe identify namespace: %w", err)
		}
		if len(b) != 4096 {
			return fmt.Errorf("NVMe identify namespace file must be exactly 4096 bytes, got %d", len(b))
		}
		copy(ctx.NVMeIdentify.Namespace[:], b)
		ctx.NVMeIdentify.HasNamespace = true
		slog.Info("loaded donor NVMe identify namespace", "file", buildOpts.nvmeIdentifyNS)
	}
	return nil
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
	buildCmd.Flags().StringVar(&buildOpts.nvmeIdentifyCtrl, "nvme-identify-ctrl", "", "path to donor-captured NVMe Identify Controller page (4096-byte binary)")
	buildCmd.Flags().StringVar(&buildOpts.nvmeIdentifyNS, "nvme-identify-ns", "", "path to donor-captured NVMe Identify Namespace page (4096-byte binary)")
	buildCmd.Flags().DurationVar(&buildOpts.mmiotraceDuration, "mmiotrace-duration", 0, "capture live mmiotrace for BAR sequencing overlays (e.g. 2s)")
	buildCmd.Flags().IntVar(&buildOpts.mmiotraceBar, "mmiotrace-bar", -1, "BAR index to associate with live mmiotrace overlays (-1 = auto select largest BAR)")

	buildCmd.Flags().StringArrayVar(&buildOpts.traces, "trace", nil, "import MMIO trace into build context as BAR=path (repeatable, e.g. --trace 2=trace.log)")
	buildCmd.Flags().BoolVar(&buildOpts.publicSafe, "public-safe", false, "strip donor-unique identifiers and captured response assets from outputs")
	_ = buildCmd.MarkFlagRequired("board")

	rootCmd.AddCommand(buildCmd)
}
