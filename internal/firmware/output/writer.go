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
// generated versions. These are excluded from the board source copy
// so Vivado only sees the generated versions.
var svFilesReplacedByGenerator = map[string]bool{
	"pcileech_tlps128_bar_controller.sv": true,
}

// barControllerSubModules: sub-module names extracted from the board's
// BAR controller file. The top-level module (pcileech_tlps128_bar_controller)
// is replaced by the generated version, but these shared modules are still
// needed by the generated controller.
var barControllerSubModules = []string{
	"pcileech_tlps128_bar_rdengine",
	"pcileech_tlps128_bar_wrengine",
	"pcileech_bar_impl_none",
	"pcileech_bar_impl_loopaddr",
	"pcileech_bar_impl_zerowrite4k",
}

// patchSVSources copies the board's SV tree (excluding files that will
// be regenerated), and patches donor IDs into the remaining sources.
func (ow *OutputWriter) patchSVSources(ctx *donor.DeviceContext, b *board.Board, ids firmware.DeviceIDs) error {
	srcDir := b.SrcPath(ow.LibDir)
	dstDir := filepath.Join(ow.OutputDir, "src")

	if _, err := os.Stat(srcDir); os.IsNotExist(err) {
		slog.Warn("board source dir not found", "path", srcDir)
		slog.Info("run: git submodule update --init --recursive")
		return fmt.Errorf("board sources not found at %s (is the pcileech-fpga submodule initialized?)", srcDir)
	}

	// Copy board sources, excluding files replaced by the generator.
	// This prevents Vivado from importing duplicate module definitions
	// (same name from both board source and generated file).
	if err := copyDirExcluding(srcDir, dstDir, svFilesReplacedByGenerator); err != nil {
		return fmt.Errorf("failed to copy SV sources: %w", err)
	}

	// Extract sub-modules from the board's BAR controller file and write
	// them as separate .sv files. The generated BAR controller depends on
	// these but the board's top-level module would conflict with the
	// generated version, so we exclude the original file and split it.
	if !ow.StockBar {
		ctrlSrc := filepath.Join(srcDir, "pcileech_tlps128_bar_controller.sv")
		if err := extractSubModules(ctrlSrc, dstDir, barControllerSubModules); err != nil {
			slog.Warn("could not extract BAR controller sub-modules, board source may be incompatible", "error", err)
		}
	} else {
		slog.Info("stock-bar mode: keeping stock bar controller")
		// In stock mode, re-copy the controller since copyDirExcluding skipped it.
		srcFile := filepath.Join(srcDir, "pcileech_tlps128_bar_controller.sv")
		dstFile := filepath.Join(dstDir, "pcileech_tlps128_bar_controller.sv")
		if data, err := os.ReadFile(srcFile); err == nil {
			os.WriteFile(dstFile, data, 0644)
		}
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

// copyDirExcluding copies a directory recursively but skips files whose
// names are in the exclude map. Used to prevent board source files from
// being imported alongside generated versions with the same module names.
func copyDirExcluding(src, dst string, exclude map[string]bool) error {
	if err := os.MkdirAll(dst, 0755); err != nil {
		return err
	}
	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}
	for _, e := range entries {
		srcPath := filepath.Join(src, e.Name())
		dstPath := filepath.Join(dst, e.Name())
		if e.IsDir() {
			if err := copyDirExcluding(srcPath, dstPath, exclude); err != nil {
				return err
			}
		} else if !exclude[e.Name()] {
			data, err := os.ReadFile(srcPath)
			if err != nil {
				return err
			}
			if err := os.WriteFile(dstPath, data, 0644); err != nil {
				return err
			}
		}
	}
	return nil
}

// extractSubModules reads a board BAR controller source file and writes
// each named sub-module as a separate .sv file. This allows the generated
// top-level BAR controller to use shared sub-modules without conflicting
// with the board's original file.
func extractSubModules(srcPath string, dstDir string, subModules []string) error {
	data, err := os.ReadFile(srcPath)
	if err != nil {
		return err
	}

	content := string(data)
	for _, modName := range subModules {
		searchStr := "module " + modName
		start := strings.Index(content, searchStr)
		if start == -1 {
			slog.Warn("sub-module not found in board controller", "module", modName)
			continue
		}

		// Find the end: next "\nmodule " after our match, or end of file.
		rest := content[start+len(searchStr):]
		nextIdx := strings.Index(rest, "\nmodule ")
		var modBody string
		if nextIdx == -1 {
			modBody = content[start:]
		} else {
			modBody = content[start : start+len(searchStr)+nextIdx+1]
		}

		dstFile := filepath.Join(dstDir, modName+".sv")
		if err := os.WriteFile(dstFile, []byte(modBody), 0644); err != nil {
			slog.Warn("failed to write sub-module file", "module", modName, "error", err)
		}
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
		"pcileech_bar_impl_msi.sv",
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
	{"pcileech_bar_impl_msi.sv", svgen.GenerateBarImplMSISV},
	{"tlp_latency_emulator.sv", svgen.GenerateLatencyEmulatorSV},
	{"device_config.sv", svgen.GenerateDeviceConfigSV},
}

// writeSVModules generates device-specific SV and HEX init files.
func (ow *OutputWriter) writeSVModules(ctx *donor.DeviceContext, scrubbedCS *pci.ConfigSpace, ids firmware.DeviceIDs, entropy uint32) error {
	cfg := ow.buildSVConfig(ctx, scrubbedCS, ids, entropy)

	if err := ow.writeCoreSVArtifacts(cfg, scrubbedCS); err != nil {
		return err
	}

	if err := ow.writeConditionalArtifacts(cfg, ctx); err != nil {
		return err
	}

	ow.logSVSummary(cfg)
	return nil
}

// extractMSIInfo parses the MSI capability from scrubbed config space.
// Returns nil if MSI is absent or not yet programmed by the host.
func extractMSIInfo(cs *pci.ConfigSpace) *svgen.MSIConfig {
	// Walk the capability linked list
	ptr := int(cs.CapabilityPointer()) & 0xFC
	for ptr != 0 && ptr < 0x100 {
		capID := cs.ReadU8(ptr)
		nextPtr := int(cs.ReadU8(ptr+1)) & 0xFC

		if capID == 0x05 { // CapIDMSI
			msgCtl := cs.ReadU16(ptr + 2)
			is64bit := (msgCtl & (1 << 7)) != 0

			addrLo := cs.ReadU32(ptr + 4)
			var data uint16
			if is64bit {
				data = cs.ReadU16(ptr + 12)
			} else {
				data = cs.ReadU16(ptr + 8)
			}

			// Use standard defaults if not yet programmed by host (addr=0).
			// The APIC accepts MSI writes to 0xFEE00000.
			if addrLo == 0 {
				addrLo = 0xFEE00000
			}
			if data == 0 {
				data = 0x0000 // vector 0, edge, physical mode
			}

			return &svgen.MSIConfig{
				Enabled: (msgCtl & (1 << 0)) != 0,
				AddrLo:  addrLo,
				Data:    data,
			}
		}

		ptr = nextPtr
	}
	return nil
}

// buildSVConfig assembles the SVGeneratorConfig from donor context.
func (ow *OutputWriter) buildSVConfig(ctx *donor.DeviceContext, scrubbedCS *pci.ConfigSpace, ids firmware.DeviceIDs, entropy uint32) *svgen.SVGeneratorConfig {
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
		"has_profile", barProfile != nil,
		"class_code", fmt.Sprintf("0x%06X", ctx.Device.ClassCode),
	)
	bm := barmodel.BuildBARModel(barData, ctx.Device.ClassCode, barProfile)
	slog.Info("BAR model built",
		"model_nil", bm == nil,
		"reg_count", func() int { if bm != nil { return len(bm.Registers) }; return 0 }(),
	)

	strategy := devclass.StrategyForClassAndVendor(ctx.Device.ClassCode, ids.VendorID)
	devClass := ""
	if strategy != nil {
		devClass = strategy.DeviceClass()
	}
	slog.Info("device class resolution", "class", devClass)

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

	// Extract MSI capability for interrupt generation.
	// MSI is the primary interrupt mechanism for HDA devices.
	if msiInfo := extractMSIInfo(scrubbedCS); msiInfo != nil {
		cfg.MSIConfig = msiInfo
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

		// MSI interrupt generator for HDA - critical for driver completion.
		// always generated for audio devices: template instantiates unconditionally.
		hdaMSISV, err := svgen.GenerateHDAMSISV(cfg)
		if err != nil {
			return fmt.Errorf("generating pcileech_hda_msi.sv: %w", err)
		}
		if err := ow.writeFile("pcileech_hda_msi.sv", hdaMSISV); err != nil {
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
		if cfg.MSIConfig != nil {
			features = append(features, "MSI Interrupt Gen")
		}
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
