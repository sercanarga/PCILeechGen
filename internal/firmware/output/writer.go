// Package output writes the final firmware artifacts (COE, TCL, SV, HEX).
package output

import (
	"encoding/binary"
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
	"github.com/sercanarga/pcileechgen/internal/firmware/devicemodel"
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
	if ids.SubsysVendorID == 0 {
		ids.SubsysVendorID = scrubbedCS.SubsysVendorID()
		ctx.Device.SubsysVendorID = ids.SubsysVendorID
	}
	if ids.SubsysDeviceID == 0 {
		ids.SubsysDeviceID = scrubbedCS.SubsysDeviceID()
		ctx.Device.SubsysDeviceID = ids.SubsysDeviceID
	}
	ctx.Device.RevisionID = scrubbedCS.RevisionID()
	ctx.Device.VendorID = scrubbedCS.VendorID()
	ctx.Device.DeviceID = scrubbedCS.DeviceID()
	ctx.Device.ClassCode = scrubbedCS.ReadU32(0x08) >> 8
	var svCfg *svgen.SVGeneratorConfig
	if !ow.StockBar {
		var err error
		svCfg, err = ow.buildSVConfig(ctx, scrubbedCS, ids, entropy, b)
		if err != nil {
			return fmt.Errorf("SV module generation failed: %w", err)
		}
	}
	if err := ow.writeDeviceModel(ctx, scrubbedCS, svCfg); err != nil {
		return err
	}

	if err := ow.writeConfigSpaceArtifacts(ctx, scrubbedCS, b); err != nil {
		return err
	}

	if err := ow.writeBARBehaviorProfile(ctx); err != nil {
		return err
	}

	if err := ow.writeTCLScripts(ctx, b, svCfg); err != nil {
		return err
	}

	if err := ow.patchSVSources(b, ids); err != nil {
		return fmt.Errorf("SV patching failed: %w", err)
	}

	if !ow.StockBar {
		if err := ow.writeSVModules(ctx, scrubbedCS, svCfg); err != nil {
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

	if err := writeBuildManifest(ow.OutputDir, ctx, b, false); err != nil {
		return fmt.Errorf("failed to write build manifest: %w", err)
	}
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

func (ow *OutputWriter) writeDeviceModel(ctx *donor.DeviceContext, scrubbedCS *pci.ConfigSpace, cfg *svgen.SVGeneratorConfig) error {
	modelContext := ctx
	if scrubbedCS != nil {
		finalContext := *ctx
		finalContext.ConfigSpace = scrubbedCS
		finalContext.Capabilities = pci.ParseCapabilities(scrubbedCS)
		finalContext.ExtCapabilities = pci.ParseExtCapabilities(scrubbedCS)
		modelContext = &finalContext
	}
	model, err := devicemodel.Build(modelContext)
	if err != nil {
		return fmt.Errorf("failed to build device model: %w", err)
	}
	if scrubbedCS != nil {
		size := scrubbedCS.Size
		if size > len(scrubbedCS.Data) {
			size = len(scrubbedCS.Data)
		}
		model.ConfigSpace.Size = uint32(size)
		model.ConfigSpace.ResetImage = append(model.ConfigSpace.ResetImage[:0], scrubbedCS.Data[:size]...)
		for index := range model.ConfigSpace.Fields {
			field := &model.ConfigSpace.Fields[index]
			end := int(field.Offset) + int(field.Width)
			if end <= len(model.ConfigSpace.ResetImage) {
				field.ResetValue = readModelField(model.ConfigSpace.ResetImage[field.Offset:end]) & field.Mask
			}
		}
		for index := range model.Registers {
			register := &model.Registers[index]
			end := int(register.Offset) + int(register.Width)
			if register.Space != devicemodel.SpaceConfig || end > len(model.ConfigSpace.ResetImage) {
				continue
			}
			register.ResetValue = readModelField(model.ConfigSpace.ResetImage[register.Offset:end])
			for fieldIndex := range register.Fields {
				field := &register.Fields[fieldIndex]
				field.ResetValue = register.ResetValue & field.Mask
			}
		}
		if len(model.Functions) == 1 {
			model.Functions[0].VendorID = scrubbedCS.VendorID()
			model.Functions[0].DeviceID = scrubbedCS.DeviceID()
			model.Functions[0].SubsystemVendorID = scrubbedCS.SubsysVendorID()
			model.Functions[0].SubsystemDeviceID = scrubbedCS.SubsysDeviceID()
			model.Functions[0].RevisionID = scrubbedCS.RevisionID()
			model.Functions[0].ClassCode = scrubbedCS.ClassCode()
		}
	}
	if cfg != nil {
		applyGeneratedBARModels(model, cfg.BARModels)
		if cfg.NVMeIdentify != nil && ctx.Device.ClassCode == 0x010802 &&
			ctx.NVMeIdentity != nil &&
			len(ctx.NVMeIdentity.RawControllerIdent) == 4096 &&
			len(ctx.NVMeIdentity.RawNamespaceIdent) == 4096 {
			model.NVMeIdentify = &devicemodel.NVMeIdentify{
				Controller: append([]byte(nil), cfg.NVMeIdentify.Controller[:]...),
				Namespace:  append([]byte(nil), cfg.NVMeIdentify.Namespace[:]...),
			}
		}
	}
	data, err := model.ToJSON()
	if err != nil {
		return fmt.Errorf("failed to marshal device model: %w", err)
	}
	if err := ow.writeFile("device_model.json", string(data)); err != nil {
		return fmt.Errorf("failed to write device model: %w", err)
	}
	return nil
}

func readModelField(data []byte) uint64 {
	var value uint64
	for index := len(data) - 1; index >= 0; index-- {
		value = value<<8 | uint64(data[index])
	}
	return value
}

// scrubAndVary runs config space scrubbing and per-build variance.
func (ow *OutputWriter) scrubAndVary(ctx *donor.DeviceContext, b *board.Board, ids firmware.DeviceIDs) (*pci.ConfigSpace, uint32, *overlay.Map) {
	msixTableSize := 0
	if ctx.MSIXData != nil && ctx.MSIXData.TableSize > 0 {
		msixTableSize = ctx.MSIXData.TableSize
	}
	bar0Size := firmware.CappedBAR0Size(ctx, b, msixTableSize)
	scrubbedCS, overlayMap := scrub.ScrubConfigSpaceWithDonor(ctx.ConfigSpace, b, ctx.BARs, ctx.MSIXData, bar0Size)
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
	if err := ow.writeFile("pcileech_bar_zero4k.coe",
		codegen.GenerateBarContentCOE(ctx.BARContents, firmware.CappedBAR0Size(ctx, b, msixTableSize))); err != nil {
		return fmt.Errorf("failed to write bar zero COE: %w", err)
	}
	return nil
}

// writeTCLScripts generates Vivado project and build TCL scripts.
func (ow *OutputWriter) writeTCLScripts(ctx *donor.DeviceContext, b *board.Board, configs ...*svgen.SVGeneratorConfig) error {
	var cfg *svgen.SVGeneratorConfig
	if len(configs) > 0 {
		cfg = configs[0]
	}
	if err := ow.writeFile("vivado_generate_project.tcl",
		tclgen.GenerateProjectTCLWithConfig(ctx, b, ow.LibDir, ow.StockBar, cfg)); err != nil {
		return fmt.Errorf("failed to write project TCL: %w", err)
	}
	if err := ow.writeFile("vivado_build.tcl",
		tclgen.GenerateBuildTCL(b, ow.Jobs, ow.Timeout)); err != nil {
		return fmt.Errorf("failed to write build TCL: %w", err)
	}
	return nil
}

func WriteBuildManifest(outputDir string, ctx *donor.DeviceContext, b *board.Board) error {
	return writeBuildManifest(outputDir, ctx, b, true)
}

func writeBuildManifest(outputDir string, ctx *donor.DeviceContext, b *board.Board, includeSynthesized bool) error {
	ids := firmware.ExtractDeviceIDs(ctx.ConfigSpace, ctx.ExtCapabilities)
	manifest, err := generateManifest(outputDir, ctx.ToolVersion, b.Name, ids.VendorID, ids.DeviceID, includeSynthesized)
	if err != nil {
		return err
	}
	manifest.DeviceBDF = ctx.Device.BDF.String()
	manifestPath := filepath.Join(outputDir, "build_manifest.json")
	if err := manifest.WriteJSON(manifestPath); err != nil {
		return err
	}
	slog.Info("build manifest written", "files", len(manifest.Files))
	return nil
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
	"pcileech_tlps128_bar_wrengine",
	"pcileech_bar_impl_none",
	"pcileech_bar_impl_zerowrite4k",
}

// patchSVSources copies the board's SV tree (excluding files that will
// be regenerated), and patches donor IDs into the remaining sources.
func (ow *OutputWriter) patchSVSources(b *board.Board, ids firmware.DeviceIDs) error {
	srcDir := b.SrcPath(ow.LibDir)
	dstDir := filepath.Join(ow.OutputDir, "src")

	if _, err := os.Stat(srcDir); os.IsNotExist(err) {
		slog.Warn("board source dir not found", "path", srcDir)
		slog.Info("run: git submodule update --init --recursive")
		return fmt.Errorf("board sources not found at %s (is the pcileech-fpga submodule initialized?)", srcDir)
	}

	if err := os.RemoveAll(dstDir); err != nil {
		return fmt.Errorf("failed to clear stale src dir %s: %w", dstDir, err)
	}

	// Copy board sources to local src/ (excluding the top-level bar controller
	// when generating custom version). The generated vivado_generate_project.tcl
	// now globs from this local dir so the donor-custom controller is used
	// instead of stock, avoiding duplicate module imports.
	if err := copyDirExcluding(srcDir, dstDir, svFilesReplacedByGenerator); err != nil {
		return fmt.Errorf("failed to copy SV sources: %w", err)
	}

	// Extract sub-modules from the board's BAR controller file and write
	// them as separate .sv files. The generated BAR controller depends on
	// these but the board's top-level module would conflict with the
	// generated version, so we exclude the original file and split it.
	if !ow.StockBar {
		ctrlSrc := filepath.Join(srcDir, "pcileech_tlps128_bar_controller.sv")
		if _, err := os.Stat(ctrlSrc); err == nil {
			if eerr := extractSubModules(ctrlSrc, dstDir, barControllerSubModules); eerr != nil {
				return fmt.Errorf("failed to extract BAR controller sub-modules: %w", eerr)
			}
		} else if !os.IsNotExist(err) {
			return fmt.Errorf("failed to inspect BAR controller: %w", err)
		}
	} else {
		slog.Info("stock-bar mode: keeping stock bar controller")
		// In stock mode, re-copy the controller since copyDirExcluding skipped it.
		srcFile := filepath.Join(srcDir, "pcileech_tlps128_bar_controller.sv")
		dstFile := filepath.Join(dstDir, "pcileech_tlps128_bar_controller.sv")
		if data, err := os.ReadFile(srcFile); err == nil {
			if writeErr := os.WriteFile(dstFile, data, 0644); writeErr != nil {
				return fmt.Errorf("failed to copy stock BAR controller: %w", writeErr)
			}
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
		"device_model.json",
		"pcileech_cfgspace.coe",
		"pcileech_cfgspace_writemask.coe",
		"pcileech_bar_zero4k.coe",
		"bar_behavior_profile.json",
		"vivado_generate_project.tcl",
		"vivado_build.tcl",
		"src/",
		"pcileech_lifecycle_service.sv",
		"pcileech_dma_tag_service.sv",
		"pcileech_interrupt_service.sv",
		"pcileech_bar_impl_device.sv",
		"pcileech_tlps128_bar_controller.sv",
		"pcileech_tlp_normalizer.sv",
		"pcileech_tlps128_bar_rdengine.sv",
		"pcileech_tlp_ur_completer.sv",
		"pcileech_bar_impl_msi.sv",
		"pcileech_bar_rsp_arbiter.sv",
		"pcileech_msix_table.sv",
		"pcileech_nvme_admin_responder.sv",
		"pcileech_nvme_dma_bridge.sv",
		"pcileech_ethernet_dma_bridge.sv",
		"pcileech_ethernet_dma_engine.sv",
		"pcileech_bram_disk.sv",
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
	{"pcileech_lifecycle_service.sv", svgen.GenerateLifecycleServiceSV},
	{"pcileech_dma_tag_service.sv", svgen.GenerateDMATagServiceSV},
	{"pcileech_interrupt_service.sv", svgen.GenerateInterruptServiceSV},
	{"pcileech_bar_impl_device.sv", svgen.GenerateBarImplDeviceSV},
	{"pcileech_tlps128_bar_controller.sv", svgen.GenerateBarControllerSV},
	{"pcileech_tlp_normalizer.sv", svgen.GenerateTransactionNormalizerSV},
	{"pcileech_tlps128_bar_rdengine.sv", svgen.GenerateBarReadEngineSV},
	{"pcileech_tlp_ur_completer.sv", svgen.GenerateURCompleterSV},
	{"pcileech_bar_rsp_arbiter.sv", svgen.GenerateBarRspArbiterSV},
	{"tlp_latency_emulator.sv", svgen.GenerateLatencyEmulatorSV},
	{"device_config.sv", svgen.GenerateDeviceConfigSV},
}

func (ow *OutputWriter) writeSVModules(ctx *donor.DeviceContext, scrubbedCS *pci.ConfigSpace, cfg *svgen.SVGeneratorConfig) error {
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
			return &svgen.MSIConfig{
				Enabled: (msgCtl & (1 << 0)) != 0,
			}
		}

		ptr = nextPtr
	}
	return nil
}

type barInterval struct {
	start uint64
	end   uint64
}

func alignUp(value, alignment uint64) uint64 {
	if alignment <= 1 {
		return value
	}
	return (value + alignment - 1) &^ (alignment - 1)
}

func placeBARRegion(model *barmodel.BARModel, preferred, length, alignment uint64, reserved []barInterval) (uint32, error) {
	if model == nil || model.Size <= 0 {
		return 0, fmt.Errorf("BAR endpoint is absent")
	}
	limit := uint64(model.Size)
	occupied := append([]barInterval(nil), reserved...)
	for _, reg := range model.Registers {
		width := reg.Width
		if width <= 0 {
			width = 4
		}
		occupied = append(occupied, barInterval{start: uint64(reg.Offset), end: uint64(reg.Offset) + uint64(width)})
	}
	fits := func(start uint64) bool {
		if start%alignment != 0 || start > limit || length > limit-start {
			return false
		}
		end := start + length
		for _, used := range occupied {
			if start < used.end && used.start < end {
				return false
			}
		}
		return start <= uint64(^uint32(0)) && end <= uint64(^uint32(0))+1
	}

	candidates := []uint64{preferred, 0}
	for _, used := range occupied {
		candidates = append(candidates, alignUp(used.end, alignment))
	}
	for _, candidate := range candidates {
		if fits(candidate) {
			return uint32(candidate), nil
		}
	}
	return 0, fmt.Errorf("BAR%d aperture %#x has no %#x-byte aligned gap for MSI-X region",
		model.BIR, limit, length)
}

func (ow *OutputWriter) buildSVConfig(ctx *donor.DeviceContext, scrubbedCS *pci.ConfigSpace, ids firmware.DeviceIDs, entropy uint32, b *board.Board) (*svgen.SVGeneratorConfig, error) {
	strategy := devclass.StrategyForClassAndVendor(ctx.Device.ClassCode, ids.VendorID)
	devClass := ""
	preferredBIR := 0
	if strategy != nil {
		devClass = strategy.DeviceClass()
		if profile := strategy.Profile(); profile != nil {
			preferredBIR = profile.PreferredBAR
		}
	}

	models, err := barmodel.BuildBARModels(ctx.BARs, ctx.BARContents, ctx.BARProfiles,
		ctx.Device.ClassCode, preferredBIR)
	if err != nil {
		return nil, fmt.Errorf("invalid donor BAR topology: %w", err)
	}
	primary := barmodel.ModelForBIR(models, preferredBIR)
	if primary == nil && len(models) > 0 {
		primary = models[0]
	}

	msixTableSize := 0
	if ctx.MSIXData != nil && ctx.MSIXData.TableSize > 0 {
		msixTableSize = ctx.MSIXData.TableSize
	}
	donorBar := firmware.DonorBAR0Demand(ctx, b, msixTableSize)
	bar0Size := firmware.CappedBAR0Size(ctx, b, msixTableSize)
	if bar0 := barmodel.ModelForBIR(models, 0); bar0 != nil && bar0.Size < bar0Size {
		bar0.Size = bar0Size
	}

	barData := []byte(nil)
	if primary != nil {
		barData = ctx.BARContents[primary.BIR]
	}
	slog.Info("BAR topology for SV codegen",
		"models", len(models),
		"preferred_bir", preferredBIR,
		"primary_bir", func() int {
			if primary != nil {
				return primary.BIR
			}
			return -1
		}(),
		"class_specific", primary != nil && primary.ClassSpecific,
		"class_code", fmt.Sprintf("0x%06X", ctx.Device.ClassCode),
	)

	var compiledRules *svgen.CompiledBehavior
	if ctx.BehaviorRules != nil {
		if ownershipErr := svgen.ValidateBehaviorRuleOwnership(ctx.BehaviorRules, devClass); ownershipErr != nil {
			return nil, ownershipErr
		}
		ruleModel := barmodel.ModelForBIR(models, ctx.BehaviorRules.BARIndex)
		if ruleModel == nil {
			return nil, fmt.Errorf("behavior rules target BAR%d but no generated endpoint exists at that BIR", ctx.BehaviorRules.BARIndex)
		}
		if ctx.BehaviorRules.BARSize > ruleModel.Size {
			return nil, fmt.Errorf("behavior BAR size %d exceeds generated BAR%d size %d", ctx.BehaviorRules.BARSize, ruleModel.BIR, ruleModel.Size)
		}
		applied, applyErr := barmodel.ApplyBehaviorRules(ruleModel, ctx.BehaviorRules)
		if applyErr != nil {
			return nil, fmt.Errorf("apply behavior rules: %w", applyErr)
		}
		for i := range models {
			if models[i].BIR == applied.BIR {
				models[i] = applied
				break
			}
		}
		if primary != nil && primary.BIR == applied.BIR {
			primary = applied
		}
		var compileErr error
		compiledRules, compileErr = svgen.CompileBehaviorRules(ctx.BehaviorRules, applied)
		if compileErr != nil {
			return nil, fmt.Errorf("compile behavior rules: %w", compileErr)
		}
	}

	cfg := &svgen.SVGeneratorConfig{
		DeviceIDs:                   ids,
		DonorCapabilities:           extractDonorCapabilities(scrubbedCS),
		BARModels:                   models,
		DonorBARTopology:            true,
		BARModel:                    primary,
		ClassCode:                   ctx.Device.ClassCode,
		LatencyConfig:               svgen.DefaultLatencyConfig(ctx.Device.ClassCode),
		HasMSIX:                     ctx.MSIXData != nil && ctx.MSIXData.TableSize > 0,
		BuildEntropy:                entropy,
		PRNGSeeds:                   svgen.BuildPRNGSeeds(ids.VendorID, ids.DeviceID, entropy),
		DeviceClass:                 devClass,
		Bar0Size:                    bar0Size,
		ReadCompletionBoundaryBytes: 64,
		MaxPayloadBytes:             128,
		BehaviorRules:               ctx.BehaviorRules,
		CompiledBehavior:            compiledRules,
	}
	if primary == nil {
		cfg.LatencyConfig = nil
	}

	if devClass == devclass.ClassNVMe && cfg.HasClassEndpoint() {
		var identity *nvme.ControllerIdentity
		if ctx.NVMeIdentity != nil {
			identity = &nvme.ControllerIdentity{
				Serial: ctx.NVMeIdentity.Serial,
				Model:  ctx.NVMeIdentity.Model,
				FWRev:  ctx.NVMeIdentity.FWRev,

				RawControllerIdent: ctx.NVMeIdentity.RawControllerIdent,
				RawNamespaceIdent:  ctx.NVMeIdentity.RawNamespaceIdent,
			}
		}
		cfg.NVMeIdentify = nvme.BuildIdentifyData(ids, barData, identity)
		cfg.NVMeSMART = nvme.BuildSMART()

		// Normalize donor NSZE to 512B units: firmware always serves LBADS=9.
		if nsze := binary.LittleEndian.Uint64(cfg.NVMeIdentify.Namespace[0x000:0x008]); nsze > 0 {
			lbaf0 := binary.LittleEndian.Uint32(cfg.NVMeIdentify.Namespace[0x0C0:0x0C4])
			if donorLBADS := uint((lbaf0 >> 16) & 0xFF); donorLBADS >= 9 {
				nsze <<= (donorLBADS - 9)
			}
			cfg.NVMeAdvertisedLBAs = nsze
		}
		if len(barData) >= 0x08 {
			cfg.NVMeDoorbellStride = nvme.DoorbellStrideFromCAP(util.ReadLE32(barData, 0x04))
		}
		// refuse early if board can't fit a cache
		cfg.NVMeDiskWords = svgen.NVMeDiskWordsForBRAM36(b.BRAM36Capacity())
		if cfg.NVMeDiskWords == 0 {
			return nil, fmt.Errorf(
				"board %q (%s) has insufficient block RAM for NVMe disk emulation; "+
					"use an Artix-7 75T/100T/200T board (35T/50T lack the block RAM), "+
					"200T for the full disk cache",
				b.Name, b.FPGAPart)
		}
	}

	if ctx.MSIXData != nil && ctx.MSIXData.TableSize > 0 {
		tableModel := barmodel.ModelForBIR(models, ctx.MSIXData.TableBIR)
		if tableModel == nil {
			return nil, fmt.Errorf("MSI-X table advertises BIR %d, but that BIR has no generated memory endpoint",
				ctx.MSIXData.TableBIR)
		}
		pbaModel := barmodel.ModelForBIR(models, ctx.MSIXData.PBABIR)
		if pbaModel == nil {
			return nil, fmt.Errorf("MSI-X PBA advertises BIR %d, but that BIR has no generated memory endpoint",
				ctx.MSIXData.PBABIR)
		}

		var tableReserved []barInterval
		if devClass == devclass.ClassNVMe && tableModel.BIR == preferredBIR {
			stride := uint64(4) << cfg.NVMeDoorbellStride
			tableReserved = append(tableReserved, barInterval{
				start: uint64(board.DefaultBRAMSize),
				end:   uint64(board.DefaultBRAMSize) + 4*stride,
			})
		}
		tableBytes := uint64(ctx.MSIXData.TableSize * 16)
		tableOff, err := placeBARRegion(tableModel, uint64(ctx.MSIXData.TableOffset),
			tableBytes, 16, tableReserved)
		if err != nil {
			return nil, fmt.Errorf("placing MSI-X table: %w", err)
		}

		pbaBytes := uint64((ctx.MSIXData.TableSize + 63) / 64 * 8)
		if pbaBytes < 8 {
			pbaBytes = 8
		}
		var pbaReserved []barInterval
		if pbaModel.BIR == tableModel.BIR {
			pbaReserved = append(pbaReserved, barInterval{
				start: uint64(tableOff),
				end:   uint64(tableOff) + tableBytes,
			})
		}
		if devClass == devclass.ClassNVMe && pbaModel.BIR == preferredBIR {
			stride := uint64(4) << cfg.NVMeDoorbellStride
			pbaReserved = append(pbaReserved, barInterval{
				start: uint64(board.DefaultBRAMSize),
				end:   uint64(board.DefaultBRAMSize) + 4*stride,
			})
		}
		pbaOff, err := placeBARRegion(pbaModel, uint64(ctx.MSIXData.PBAOffset),
			pbaBytes, 8, pbaReserved)
		if err != nil {
			return nil, fmt.Errorf("placing MSI-X PBA: %w", err)
		}

		cfg.MSIXConfig = &svgen.MSIXConfig{
			NumVectors:  ctx.MSIXData.TableSize,
			TableBIR:    tableModel.BIR,
			TableOffset: tableOff,
			PBABIR:      pbaModel.BIR,
			PBAOffset:   pbaOff,
		}
		cfg.HasMSIX = true
		if m := pci.ParseMSIXCap(scrubbedCS); m != nil {
			scrubbedCS.WriteU32(m.CapOffset+4, (tableOff&0xFFFFFFF8)|uint32(tableModel.BIR))
			scrubbedCS.WriteU32(m.CapOffset+8, (pbaOff&0xFFFFFFF8)|uint32(pbaModel.BIR))
		}
	}

	// Extract MSI capability for interrupt generation.
	// MSI is the primary interrupt mechanism for HDA devices.
	if msiInfo := extractMSIInfo(scrubbedCS); msiInfo != nil {
		cfg.MSIConfig = msiInfo
	}

	bram := b.BRAMSizeOrDefault()
	// Validate *donor demand* (may exceed) against board BRAM; error unless --force.
	// (bar0Size is the Capped value actually used for artifacts/scrub/SV.)
	if issues := ValidateBARSize(donorBar, bram, 0); len(issues) > 0 {
		if !ow.Force {
			return nil, fmt.Errorf("%s", issues[0])
		}
	}

	return cfg, nil
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

	if cfg.MSIConfig != nil {
		msiEpSV, err := svgen.GenerateBarImplMSISV(cfg)
		if err != nil {
			return fmt.Errorf("generating pcileech_bar_impl_msi.sv: %w", err)
		}
		if err := ow.writeFile("pcileech_bar_impl_msi.sv", msiEpSV); err != nil {
			return err
		}
	}

	if cfg.NVMeIdentify != nil {
		nvmeSV, err := svgen.GenerateNVMeResponderSV(cfg)
		if err != nil {
			return fmt.Errorf("generating pcileech_nvme_admin_responder.sv: %w", err)
		}
		if writeErr := ow.writeFile("pcileech_nvme_admin_responder.sv", nvmeSV); writeErr != nil {
			return writeErr
		}

		bridgeSV, err := svgen.GenerateNVMeDMABridgeSV(cfg)
		if err != nil {
			return fmt.Errorf("generating pcileech_nvme_dma_bridge.sv: %w", err)
		}
		if werr := ow.writeFile("pcileech_nvme_dma_bridge.sv", bridgeSV); werr != nil {
			return werr
		}

		diskSV, err := svgen.GenerateNVMeBRAMDiskSV(cfg)
		if err != nil {
			return fmt.Errorf("generating pcileech_bram_disk.sv: %w", err)
		}
		if err := ow.writeFile("pcileech_bram_disk.sv", diskSV); err != nil {
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
		if writeErr := ow.writeFile("pcileech_hda_rirb_dma.sv", hdaSV); writeErr != nil {
			return writeErr
		}
	}

	if cfg.DeviceClass == devclass.ClassEthernet && cfg.BARModel != nil {
		engineSV, err := svgen.GenerateEthernetDMAEngineSV(cfg)
		if err != nil {
			return fmt.Errorf("generating pcileech_ethernet_dma_engine.sv: %w", err)
		}
		if writeErr := ow.writeFile("pcileech_ethernet_dma_engine.sv", engineSV); writeErr != nil {
			return writeErr
		}

		bridgeSV, err := svgen.GenerateEthernetDMABridgeSV(cfg)
		if err != nil {
			return fmt.Errorf("generating pcileech_ethernet_dma_bridge.sv: %w", err)
		}
		if writeErr := ow.writeFile("pcileech_ethernet_dma_bridge.sv", bridgeSV); writeErr != nil {
			return writeErr
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
	case devclass.ClassEthernet:
		features = append(features, "Ethernet descriptor DMA", "Ethernet packet engine")
	}
	if cfg.MSIXConfig != nil {
		features = append(features, fmt.Sprintf("MSI-X %d vectors", cfg.MSIXConfig.NumVectors))
	}
	if cfg.BARModel != nil {
		features = append(features, fmt.Sprintf("%d registers", len(cfg.BARModel.Registers)))
	} else {
		features = append(features, "BRAM fallback")
	}
	if cfg.DonorCapabilities.HasPMCap {
		features = append(features, "donor PM cap")
	}
	if cfg.DonorCapabilities.HasMSICap {
		features = append(features, "donor MSI cap")
	}
	if cfg.DonorCapabilities.HasMSIXCap {
		features = append(features, "donor MSI-X cap")
	}
	if cfg.DonorCapabilities.HasPCIeCap {
		features = append(features, "donor PCIe cap")
	}
	features = append(features, "latency emulator", "interrupt controller")
	slog.Info("SV modules generated", "features", strings.Join(features, ", "))
}
