// Package output writes the final firmware artifacts (COE, TCL, SV, HEX).
package output

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"sort"
	"strings"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/donor/mmio"
	"github.com/sercanarga/pcileechgen/internal/firmware"
	"github.com/sercanarga/pcileechgen/internal/firmware/codegen"
	"github.com/sercanarga/pcileechgen/internal/firmware/overlay"
	"github.com/sercanarga/pcileechgen/internal/firmware/scrub"
	"github.com/sercanarga/pcileechgen/internal/firmware/svgen"
	"github.com/sercanarga/pcileechgen/internal/firmware/tclgen"
	"github.com/sercanarga/pcileechgen/internal/firmware/variance"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

// OutputWriter drops all generated files into OutputDir.
type OutputWriter struct {
	OutputDir string
	LibDir    string
	Jobs      int
	Timeout   int
	StockBar  bool
	Force     bool
}

func NewOutputWriter(outputDir, libDir string, jobs, timeout int) *OutputWriter {
	if jobs <= 0 {
		jobs = 4
	}
	if timeout <= 0 {
		timeout = 3600
	}
	return &OutputWriter{
		OutputDir: outputDir,
		LibDir:    libDir,
		Jobs:      jobs,
		Timeout:   timeout,
	}
}

// WriteAll is the main entry point - generates everything.
func (ow *OutputWriter) WriteAll(ctx *donor.DeviceContext, b *board.Board) error {
	if err := os.MkdirAll(ow.OutputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	ids := firmware.ExtractDeviceIDs(ctx.ConfigSpace, ctx.ExtCapabilities)

	if err := ow.writeDeviceContext(ctx); err != nil {
		return err
	}

	scrubbedCS, entropy, overlayMap := ow.scrubAndVary(ctx, b, ids)

	if err := ow.writeConfigSpaceArtifacts(ctx, scrubbedCS, b); err != nil {
		return err
	}

	if err := ow.writeTraceReportArtifact(ctx); err != nil {
		return err
	}

	if err := ow.writeTCLScripts(ctx, b); err != nil {
		return err
	}

	if err := ow.patchSVSources(b, ids); err != nil {
		return fmt.Errorf("SV patching failed: %w", err)
	}

	if !ow.StockBar {
		if err := ow.writeSVModules(ctx, scrubbedCS, ids, entropy, b); err != nil {
			return fmt.Errorf("SV module generation failed: %w", err)
		}
	} else {
		slog.Info("stock-bar mode: skipping custom SV modules")
	}

	// write diff report
	if overlayMap.Count() > 0 {
		if err := ow.writeFile("scrub_diff_report.txt", overlayMap.FormatDiff()); err != nil {
			slog.Warn("failed to write diff report", "error", err)
		}
	}

	if err := ow.writeWindowsLabChecklist(ctx, b); err != nil {
		return err
	}

	ow.writeManifest(ctx, ids, b)
	return nil
}

// writeWindowsLabChecklist emits a class-aware Windows lab validation
// checklist so operators have a concrete verification script per build.
func (ow *OutputWriter) writeWindowsLabChecklist(ctx *donor.DeviceContext, b *board.Board) error {
	boardName := "unspecified"
	fpgaPart := "unspecified"
	if b != nil {
		boardName = b.Name
		fpgaPart = b.FPGAPart
	}
	lines := []string{
		"Windows lab checklist",
		"",
		fmt.Sprintf("Board: %s", boardName),
		fmt.Sprintf("FPGA: %s", fpgaPart),
		fmt.Sprintf("Device: %04X:%04X class 0x%06X", ctx.Device.VendorID, ctx.Device.DeviceID, ctx.Device.ClassCode),
		"",
		"Windows 10",
		"- Flash the generated firmware to the target FPGA board.",
		"- Boot with only the test board attached to the Windows host.",
		"- Verify Device Manager shows the device without Code 10 or Code 43.",
		"- Capture setupapi.dev.log and Device Manager hardware IDs.",
		"",
		"Windows 11",
		"- Repeat enumeration and driver binding checks.",
		"- Verify modern standby or resume does not wedge the device.",
		"- Re-test after cold boot and warm reboot.",
		"",
		"Class-specific checks",
	}
	switch uint16(ctx.Device.BaseClass())<<8 | uint16(ctx.Device.SubClass()) {
	case 0x0108:
		lines = append(lines,
			"- NVMe: verify stornvme loads, namespaces appear, and no Code 10 is reported.",
		)
	case 0x0C03:
		lines = append(lines,
			"- xHCI: verify USB root hub enumerates, hot-plug a device, and confirm no disconnect storms.",
		)
	case 0x0280:
		lines = append(lines,
			"- Wi-Fi: verify the radio enumerates without dropping the host network during install.",
		)
	case 0xFF00:
		lines = append(lines,
			"- Card reader: insert media and verify card-detect and interrupt paths behave as expected.",
		)
	}
	lines = append(lines,
		"",
		"Evidence",
		"- Save screenshots, event logs, and exact board name/FPGA part with each run.",
	)
	return ow.writeFile("windows_lab_checklist.txt", strings.Join(lines, "\n"))
}

// writeDeviceContext saves the raw device context JSON.
func (ow *OutputWriter) writeDeviceContext(ctx *donor.DeviceContext) error {
	data, err := ctx.ToJSON()
	if err != nil {
		return fmt.Errorf("failed to marshal device context: %w", err)
	}
	if err := ow.writeFile("device_context.json", string(data)); err != nil {
		return fmt.Errorf("failed to write device context: %w", err)
	}
	return nil
}

// scrubAndVary runs config space scrubbing and per-build variance.
func (ow *OutputWriter) scrubAndVary(ctx *donor.DeviceContext, b *board.Board, ids firmware.DeviceIDs) (*pci.ConfigSpace, uint32, *overlay.Map) {
	msixTableSize := 0
	if ctx.MSIXData != nil && ctx.MSIXData.TableSize > 0 {
		msixTableSize = ctx.MSIXData.TableSize
	}
	bar0Size := firmware.CappedBAR0Size(ctx, b, msixTableSize)
	scrubbedCS, overlayMap := scrub.ScrubConfigSpaceWithOverlay(ctx.ConfigSpace, b, bar0Size)
	if overlayMap.Count() > 0 {
		slog.Info("config space scrubbed", "modifications", overlayMap.Count())
	}

	entropy := svgen.BuildEntropyFromTime()
	varSeed := variance.BuildVarianceSeed(ids.VendorID, ids.DeviceID, entropy)
	varCfg := variance.DefaultConfig(varSeed)
	varCfg.DonorHasDSN = ids.HasDSN
	variance.Apply(scrubbedCS, nil, varCfg)

	return scrubbedCS, entropy, overlayMap
}

func (ow *OutputWriter) writeConfigSpaceArtifacts(ctx *donor.DeviceContext, scrubbedCS *pci.ConfigSpace, b *board.Board) error {
	if err := ow.writeFile("pcileech_cfgspace.coe",
		codegen.GenerateConfigSpaceCOE(scrubbedCS)); err != nil {
		return fmt.Errorf("failed to write cfgspace COE: %w", err)
	}
	if err := ow.writeFile("pcileech_cfgspace_writemask.coe",
		codegen.GenerateWritemaskCOE(scrubbedCS)); err != nil {
		return fmt.Errorf("failed to write writemask COE: %w", err)
	}

	msixTableSize := 0
	if ctx.MSIXData != nil && ctx.MSIXData.TableSize > 0 {
		msixTableSize = ctx.MSIXData.TableSize
	}
	bar0Size := firmware.CappedBAR0Size(ctx, b, msixTableSize)

	scrub.ScrubBarContent(ctx.BARContents, ctx.Device.ClassCode, ctx.Device.VendorID, bar0Size)
	if err := ow.writeBARContentArtifacts(ctx, bar0Size); err != nil {
		return err
	}
	return nil
}

func (ow *OutputWriter) writeBARContentArtifacts(ctx *donor.DeviceContext, fallbackSize int) error {
	if ctx != nil && len(ctx.BARContents) > 0 {
		indices := make([]int, 0, len(ctx.BARContents))
		for idx := range ctx.BARContents {
			indices = append(indices, idx)
		}
		sort.Ints(indices)

		for _, idx := range indices {
			data := ctx.BARContents[idx]
			size := len(data)
			if size == 0 {
				size = fallbackSize
			}
			name := fmt.Sprintf("pcileech_bar%d.coe", idx)
			if err := ow.writeFile(name, codegen.GenerateSingleBarContentCOE(idx, data, size)); err != nil {
				return fmt.Errorf("failed to write BAR%d COE: %w", idx, err)
			}
		}
	}

	var barContents map[int][]byte
	if ctx != nil {
		barContents = ctx.BARContents
	}
	if err := ow.writeFile("pcileech_bar_zero4k.coe",
		codegen.GenerateBarContentCOE(barContents, fallbackSize)); err != nil {
		return fmt.Errorf("failed to write legacy BAR COE: %w", err)
	}
	return nil
}

type traceReportArtifact struct {
	Reports       []*mmio.TraceReport           `json:"reports"`
	TraceOverlays map[int]*mmio.TraceBAROverlay `json:"trace_overlays,omitempty"`
}

func (ow *OutputWriter) writeTraceReportArtifact(ctx *donor.DeviceContext) error {
	if ctx == nil || len(ctx.MMIOTraces) == 0 {
		return nil
	}
	indices := make([]int, 0, len(ctx.MMIOTraces))
	for idx := range ctx.MMIOTraces {
		indices = append(indices, idx)
	}
	sort.Ints(indices)

	artifact := traceReportArtifact{
		Reports: make([]*mmio.TraceReport, 0, len(indices)),
	}
	if len(ctx.BARTraceOverlays) > 0 {
		artifact.TraceOverlays = make(map[int]*mmio.TraceBAROverlay, len(ctx.BARTraceOverlays))
		for idx, overlay := range ctx.BARTraceOverlays {
			artifact.TraceOverlays[idx] = overlay
		}
	}

	for _, idx := range indices {
		artifact.Reports = append(artifact.Reports, mmio.BuildLegacyTraceReport(ctx.MMIOTraces[idx]))
	}

	data, err := json.MarshalIndent(artifact, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal BAR model report: %w", err)
	}
	if err := ow.writeFile("bar_model_report.json", string(data)+"\n"); err != nil {
		return fmt.Errorf("failed to write BAR model report: %w", err)
	}
	return nil
}

// writeTCLScripts generates Vivado project and build TCL scripts.
func (ow *OutputWriter) writeTCLScripts(ctx *donor.DeviceContext, b *board.Board) error {
	if err := ow.writeFile("vivado_generate_project.tcl",
		tclgen.GenerateProjectTCL(ctx, b, ow.LibDir, ow.StockBar)); err != nil {
		return fmt.Errorf("failed to write project TCL: %w", err)
	}
	if err := ow.writeFile("vivado_build.tcl",
		tclgen.GenerateBuildTCL(b, ow.Jobs, ow.Timeout)); err != nil {
		return fmt.Errorf("failed to write build TCL: %w", err)
	}
	return nil
}

// writeManifest generates the build manifest with file checksums.
func (ow *OutputWriter) writeManifest(ctx *donor.DeviceContext, ids firmware.DeviceIDs, b *board.Board) {
	manifest, err := GenerateManifestForBuild(ow.OutputDir, ctx.ToolVersion, b, ctx, nil, ids.VendorID, ids.DeviceID)
	if err != nil {
		slog.Warn("manifest generation failed", "error", err)
		return
	}
	manifestPath := ow.OutputDir + "/build_manifest.json"
	if err := manifest.WriteJSON(manifestPath); err != nil {
		slog.Warn("manifest write failed", "error", err)
	} else {
		slog.Info("build manifest written", "files", len(manifest.Files))
	}
}
