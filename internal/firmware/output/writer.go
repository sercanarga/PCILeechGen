// Package output writes the final firmware artifacts (COE, TCL, SV, HEX).
package output

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/firmware"
	"github.com/sercanarga/pcileechgen/internal/firmware/barmodel"
	"github.com/sercanarga/pcileechgen/internal/firmware/codegen"
	"github.com/sercanarga/pcileechgen/internal/firmware/devclass"
	"github.com/sercanarga/pcileechgen/internal/firmware/nvme"
	"github.com/sercanarga/pcileechgen/internal/firmware/overlay"
	"github.com/sercanarga/pcileechgen/internal/firmware/scrub"
	"github.com/sercanarga/pcileechgen/internal/firmware/svgen"
	"github.com/sercanarga/pcileechgen/internal/firmware/tclgen"
	"github.com/sercanarga/pcileechgen/internal/firmware/variance"
	"github.com/sercanarga/pcileechgen/internal/pci"
	"github.com/sercanarga/pcileechgen/internal/util"
)

// OutputWriter drops all generated files into OutputDir.
type OutputWriter struct {
	OutputDir string
	LibDir    string
	Jobs      int
	Timeout   int
	StockBar  bool
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

	if err := ow.writeConfigSpaceArtifacts(ctx, scrubbedCS); err != nil {
		return err
	}

	if err := ow.writeTCLScripts(ctx, b); err != nil {
		return err
	}

	if err := ow.patchSVSources(ctx, b, ids); err != nil {
		return fmt.Errorf("SV patching failed: %w", err)
	}

	if !ow.StockBar {
		if err := ow.writeSVModules(ctx, scrubbedCS, ids, entropy); err != nil {
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

	ow.writeManifest(ctx, ids)
	return nil
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
	scrubbedCS, overlayMap := scrub.ScrubConfigSpaceWithOverlay(ctx.ConfigSpace, b)
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

// writeConfigSpaceArtifacts generates COE files for config space, writemask, and BAR content.
func (ow *OutputWriter) writeConfigSpaceArtifacts(ctx *donor.DeviceContext, scrubbedCS *pci.ConfigSpace) error {
	if err := ow.writeFile("pcileech_cfgspace.coe",
		codegen.GenerateConfigSpaceCOE(scrubbedCS)); err != nil {
		return fmt.Errorf("failed to write cfgspace COE: %w", err)
	}
	if err := ow.writeFile("pcileech_cfgspace_writemask.coe",
		codegen.GenerateWritemaskCOE(scrubbedCS)); err != nil {
		return fmt.Errorf("failed to write writemask COE: %w", err)
	}

	scrub.ScrubBarContent(ctx.BARContents, ctx.Device.ClassCode, ctx.Device.VendorID)
	if err := ow.writeFile("pcileech_bar_zero4k.coe",
		codegen.GenerateBarContentCOE(ctx.BARContents)); err != nil {
		return fmt.Errorf("failed to write bar zero COE: %w", err)
	}
	return nil
}

// writeTCLScripts generates Vivado project and build TCL scripts.
func (ow *OutputWriter) writeTCLScripts(ctx *donor.DeviceContext, b *board.Board) error {
	if err := ow.writeFile("vivado_generate_project.tcl",
		tclgen.GenerateProjectTCL(ctx, b, ow.LibDir)); err != nil {
		return fmt.Errorf("failed to write project TCL: %w", err)
	}
	if err := ow.writeFile("vivado_build.tcl",
		tclgen.GenerateBuildTCL(b, ow.Jobs, ow.Timeout)); err != nil {
		return fmt.Errorf("failed to write build TCL: %w", err)
	}
	return nil
}

// writeManifest generates the build manifest with file checksums.
func (ow *OutputWriter) writeManifest(ctx *donor.DeviceContext, ids firmware.DeviceIDs) {
	manifest, err := GenerateManifest(ow.OutputDir, ctx.ToolVersion, "", ids.VendorID, ids.DeviceID)
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

// svFilesReplacedByGenerator: board source files we replace with
// generated versions. Must be deleted from src/ to avoid duplicates.
var svFilesReplacedByGenerator = []string{
	"pcileech_tlps128_bar_controller.sv",
}

// patchSVSources copies the board's SV tree, removes files that will
// be regenerated, and patches donor IDs into the remaining sources.
func (ow *OutputWriter) patchSVSources(ctx *donor.DeviceContext, b *board.Board, ids firmware.DeviceIDs) error {
	srcDir := b.SrcPath(ow.LibDir)
	dstDir := filepath.Join(ow.OutputDir, "src")

	if _, err := os.Stat(srcDir); os.IsNotExist(err) {
		slog.Warn("board source dir not found", "path", srcDir)
		slog.Info("run: git submodule update --init --recursive")
		return fmt.Errorf("board sources not found at %s (is the pcileech-fpga submodule initialized?)", srcDir)
	}

	if err := util.CopyDir(srcDir, dstDir); err != nil {
		return fmt.Errorf("failed to copy SV sources: %w", err)
	}

	// Delete stock copies so Vivado only sees our generated versions.
	// Without this, Vivado uses the stock file and causes Code 10.
	if !ow.StockBar {
		for _, name := range svFilesReplacedByGenerator {
			if err := os.Remove(filepath.Join(dstDir, name)); err != nil && !os.IsNotExist(err) {
				slog.Warn("could not remove stock SV file", "file", name, "error", err)
			}
		}
	} else {
		slog.Info("stock-bar mode: keeping stock bar controller")
	}

	patcher := svgen.NewSVPatcher(ids, dstDir)

	if err := patcher.PatchAll(); err != nil {
		return fmt.Errorf("failed to patch SV sources: %w", err)
	}

	if results := patcher.Results(); len(results) > 0 {
		slog.Info("SV patches applied", "summary", svgen.FormatPatchSummary(results))
	}

	return nil
}

func (ow *OutputWriter) writeFile(name, content string) error {
	return os.WriteFile(filepath.Join(ow.OutputDir, name), []byte(content), 0644)
}

// ListOutputFiles is used by the CLI to show what gets generated.
func ListOutputFiles() []string {
	return []string{
		"device_context.json",
		"pcileech_cfgspace.coe",
		"pcileech_cfgspace_writemask.coe",
		"pcileech_bar_zero4k.coe",
		"vivado_generate_project.tcl",
		"vivado_build.tcl",
		"src/",
		"pcileech_bar_impl_device.sv",
		"pcileech_tlps128_bar_controller.sv",
		"pcileech_msix_table.sv",
		"pcileech_nvme_admin_responder.sv",
		"pcileech_nvme_dma_bridge.sv",
		"tlp_latency_emulator.sv",
		"device_config.sv",
		"config_space_init.hex",
		"msix_table_init.hex",
		"scrub_diff_report.txt",
		"build_manifest.json",
	}
}

// svArtifact describes one SV file to generate.
type svArtifact struct {
	filename string
	generate func(cfg *svgen.SVGeneratorConfig) (string, error)
}

// coreSVArtifacts are always generated.
var coreSVArtifacts = []svArtifact{
	{"pcileech_bar_impl_device.sv", svgen.GenerateBarImplDeviceSV},
	{"pcileech_tlps128_bar_controller.sv", svgen.GenerateBarControllerSV},
	{"tlp_latency_emulator.sv", svgen.GenerateLatencyEmulatorSV},
	{"device_config.sv", svgen.GenerateDeviceConfigSV},
}

// writeSVModules generates device-specific SV and HEX init files.
func (ow *OutputWriter) writeSVModules(ctx *donor.DeviceContext, scrubbedCS *pci.ConfigSpace, ids firmware.DeviceIDs, entropy uint32) error {
	cfg := ow.buildSVConfig(ctx, ids, entropy)

	if err := ow.writeCoreSVArtifacts(cfg, scrubbedCS); err != nil {
		return err
	}

	if err := ow.writeConditionalArtifacts(cfg, ctx); err != nil {
		return err
	}

	ow.logSVSummary(cfg)
	return nil
}

// buildSVConfig assembles the SVGeneratorConfig from donor context.
func (ow *OutputWriter) buildSVConfig(ctx *donor.DeviceContext, ids firmware.DeviceIDs, entropy uint32) *svgen.SVGeneratorConfig {
	// Use the same BAR index for content data and probe profile to avoid
	// mismatched register maps (e.g. IO BAR0 profile + MMIO BAR2 data).
	barIdx := firmware.LargestBarIndex(ctx.BARContents)
	barData := ctx.BARContents[barIdx]
	var barProfile *donor.BARProfile
	if ctx.BARProfiles != nil {
		if p, ok := ctx.BARProfiles[barIdx]; ok {
			barProfile = p
		}
	}
	slog.Info("BAR selection for SV codegen",
		"bar_index", barIdx,
		"bar_size", len(barData),
		"has_profile", barProfile != nil)
	bm := barmodel.BuildBARModel(barData, ctx.Device.ClassCode, barProfile)

	strategy := devclass.StrategyForClassAndVendor(ctx.Device.ClassCode, ids.VendorID)
	devClass := ""
	if strategy != nil {
		devClass = strategy.DeviceClass()
	}

	cfg := &svgen.SVGeneratorConfig{
		DeviceIDs:     ids,
		BARModel:      bm,
		ClassCode:     ctx.Device.ClassCode,
		LatencyConfig: svgen.DefaultLatencyConfig(ctx.Device.ClassCode),
		HasMSIX:       bm != nil,
		BuildEntropy:  entropy,
		PRNGSeeds:     svgen.BuildPRNGSeeds(ids.VendorID, ids.DeviceID, entropy),
		DeviceClass:   devClass,
	}

	if devClass == devclass.ClassNVMe {
		cfg.NVMeIdentify = nvme.BuildIdentifyData(ids, barData)
		// CAP_HI at BAR offset 0x04, DSTRD = bits [3:0]
		if len(barData) >= 0x08 {
			capHI := util.ReadLE32(barData, 0x04)
			cfg.NVMeDoorbellStride = capHI & 0x0F
		}
	}

	if ctx.MSIXData != nil && ctx.MSIXData.TableSize > 0 {
		tableSize := ctx.MSIXData.TableSize * 16
		pbaOffset := uint32(0x1000) + uint32(tableSize)
		pbaOffset = (pbaOffset + 7) &^ 7
		cfg.MSIXConfig = &svgen.MSIXConfig{
			NumVectors:  ctx.MSIXData.TableSize,
			TableOffset: 0x1000,
			PBAOffset:   pbaOffset,
		}
		cfg.HasMSIX = true
	}

	return cfg
}

// writeCoreSVArtifacts generates core SV modules and config space HEX.
func (ow *OutputWriter) writeCoreSVArtifacts(cfg *svgen.SVGeneratorConfig, scrubbedCS *pci.ConfigSpace) error {
	for _, art := range coreSVArtifacts {
		content, err := art.generate(cfg)
		if err != nil {
			return fmt.Errorf("generating %s: %w", art.filename, err)
		}
		if err := ow.writeFile(art.filename, content); err != nil {
			return err
		}
	}

	hex := codegen.GenerateConfigSpaceHex(scrubbedCS)
	return ow.writeFile("config_space_init.hex", hex)
}

// writeConditionalArtifacts generates MSI-X and NVMe artifacts when applicable.
func (ow *OutputWriter) writeConditionalArtifacts(cfg *svgen.SVGeneratorConfig, ctx *donor.DeviceContext) error {
	if cfg.MSIXConfig != nil {
		msixSV, err := svgen.GenerateMSIXTableSV(cfg)
		if err != nil {
			return fmt.Errorf("generating pcileech_msix_table.sv: %w", err)
		}
		if err := ow.writeFile("pcileech_msix_table.sv", msixSV); err != nil {
			return err
		}

		entries := ctx.MSIXData.Entries
		if entries == nil {
			entries = make([]pci.MSIXEntry, cfg.MSIXConfig.NumVectors)
			for i := range entries {
				entries[i].Control = 0x01 // masked
			}
		}
		if err := ow.writeFile("msix_table_init.hex", codegen.GenerateMSIXTableHex(entries)); err != nil {
			return err
		}
	}

	if cfg.NVMeIdentify != nil {
		nvmeSV, err := svgen.GenerateNVMeResponderSV(cfg)
		if err != nil {
			return fmt.Errorf("generating pcileech_nvme_admin_responder.sv: %w", err)
		}
		if err := ow.writeFile("pcileech_nvme_admin_responder.sv", nvmeSV); err != nil {
			return err
		}

		bridgeSV, err := svgen.GenerateNVMeDMABridgeSV(cfg)
		if err != nil {
			return fmt.Errorf("generating pcileech_nvme_dma_bridge.sv: %w", err)
		}
		if err := ow.writeFile("pcileech_nvme_dma_bridge.sv", bridgeSV); err != nil {
			return err
		}

		if err := ow.writeFile("identify_init.hex", nvme.IdentifyDataToHex(cfg.NVMeIdentify)); err != nil {
			return err
		}
	}

	if cfg.DeviceClass == devclass.ClassAudio && cfg.BARModel != nil {
		hdaSV, err := svgen.GenerateHDARIRBDMASV(cfg)
		if err != nil {
			return fmt.Errorf("generating pcileech_hda_rirb_dma.sv: %w", err)
		}
		if err := ow.writeFile("pcileech_hda_rirb_dma.sv", hdaSV); err != nil {
			return err
		}
	}

	return nil
}

// logSVSummary prints a summary of generated SV features.
func (ow *OutputWriter) logSVSummary(cfg *svgen.SVGeneratorConfig) {
	var features []string
	switch cfg.DeviceClass {
	case devclass.ClassNVMe:
		features = append(features, "NVMe FSM")
		if cfg.NVMeIdentify != nil {
			features = append(features, "NVMe Admin Responder", "NVMe DMA Bridge")
		}
	case devclass.ClassXHCI:
		features = append(features, "xHCI FSM")
	case devclass.ClassAudio:
		features = append(features, "HD Audio FSM", "RIRB DMA Bridge")
	}
	if cfg.MSIXConfig != nil {
		features = append(features, fmt.Sprintf("MSI-X %d vectors", cfg.MSIXConfig.NumVectors))
	}
	if cfg.BARModel != nil {
		features = append(features, fmt.Sprintf("%d registers", len(cfg.BARModel.Registers)))
	} else {
		features = append(features, "BRAM fallback")
	}
	features = append(features, "latency emulator", "interrupt controller")
	slog.Info("SV modules generated", "features", strings.Join(features, ", "))
}
