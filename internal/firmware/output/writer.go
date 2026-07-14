// Package output writes the final firmware artifacts (COE, TCL, SV, HEX).
package output

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/firmware"
	"github.com/sercanarga/pcileechgen/internal/firmware/codegen"
	"github.com/sercanarga/pcileechgen/internal/firmware/devicemodel"
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

const outputOwnershipMarker = ".pcileechgen-output-v1"
const outputOwnershipContent = "owned by PCILeechGen generated output\n"

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
	if ctx == nil {
		return fmt.Errorf("device context is nil")
	}
	data, err := ctx.ToJSON()
	if err != nil {
		return fmt.Errorf("clone device context: %w", err)
	}
	workCtx, err := donor.FromJSON(data)
	if err != nil {
		return fmt.Errorf("clone device context: %w", err)
	}

	target, err := ow.validateOutputTarget()
	if err != nil {
		return fmt.Errorf("prepare output directory: %w", err)
	}
	parent := filepath.Dir(target)
	if err := os.MkdirAll(parent, 0755); err != nil {
		return fmt.Errorf("create output parent: %w", err)
	}
	if err := validateRealDirectory(parent, "output parent"); err != nil {
		return err
	}
	stage, err := os.MkdirTemp(parent, "."+filepath.Base(target)+".tmp-")
	if err != nil {
		return fmt.Errorf("create output staging directory: %w", err)
	}
	defer os.RemoveAll(stage)
	if err := os.WriteFile(filepath.Join(stage, outputOwnershipMarker), []byte(outputOwnershipContent), 0644); err != nil {
		return fmt.Errorf("write staging ownership marker: %w", err)
	}

	worker := *ow
	worker.OutputDir = stage
	if err := worker.writeAllPrepared(workCtx, b); err != nil {
		return err
	}
	if err := publishOutputDirectory(stage, target); err != nil {
		return fmt.Errorf("publish generated output: %w", err)
	}
	ow.OutputDir = target
	return nil
}

func (ow *OutputWriter) writeAllPrepared(ctx *donor.DeviceContext, b *board.Board) error {

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

func (ow *OutputWriter) validateOutputTarget() (string, error) {
	if ow.OutputDir == "" {
		return "", fmt.Errorf("output directory is empty")
	}
	target, err := filepath.Abs(ow.OutputDir)
	if err != nil {
		return "", fmt.Errorf("resolve output directory: %w", err)
	}
	target = filepath.Clean(target)
	if target == filepath.Dir(target) {
		return "", fmt.Errorf("refusing filesystem root as output directory")
	}
	if err := validateExistingOutputAncestor(target); err != nil {
		return "", err
	}
	info, err := os.Lstat(target)
	if os.IsNotExist(err) {
		return target, nil
	}
	if err != nil {
		return "", fmt.Errorf("inspect output directory: %w", err)
	}
	if info.Mode()&os.ModeSymlink != 0 || !info.IsDir() {
		return "", fmt.Errorf("output directory must be a real directory: %s", target)
	}
	marker := filepath.Join(target, outputOwnershipMarker)
	markerInfo, err := os.Lstat(marker)
	if os.IsNotExist(err) {
		return "", fmt.Errorf("refusing unowned existing output directory: %s", target)
	}
	if err != nil {
		return "", fmt.Errorf("inspect output ownership marker: %w", err)
	}
	if markerInfo.Mode()&os.ModeSymlink != 0 || !markerInfo.Mode().IsRegular() {
		return "", fmt.Errorf("output ownership marker is not a regular file: %s", marker)
	}
	contents, err := os.ReadFile(marker)
	if err != nil {
		return "", fmt.Errorf("read output ownership marker: %w", err)
	}
	if string(contents) != outputOwnershipContent {
		return "", fmt.Errorf("output ownership marker has unexpected content: %s", marker)
	}
	return target, nil
}

func publishOutputDirectory(stage, target string) error {
	backup := target + ".previous"
	if _, err := os.Lstat(backup); err == nil {
		return fmt.Errorf("backup path already exists: %s", backup)
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("inspect backup path: %w", err)
	}
	hadTarget := false
	if _, err := os.Lstat(target); err == nil {
		hadTarget = true
		if err := os.Rename(target, backup); err != nil {
			return fmt.Errorf("move previous output aside: %w", err)
		}
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("inspect output target: %w", err)
	}
	if err := os.Rename(stage, target); err != nil {
		if hadTarget {
			_ = os.Rename(backup, target)
		}
		return fmt.Errorf("activate staged output: %w", err)
	}
	if hadTarget {
		if err := os.RemoveAll(backup); err != nil {
			return fmt.Errorf("remove previous output: %w", err)
		}
	}
	return nil
}

// prepareOutputDir creates a new owned output directory or reopens one that
// carries this generator's marker. Requiring ownership before a rerun avoids
// writing generated files into an arbitrary existing directory.
func (ow *OutputWriter) prepareOutputDir() error {
	if ow.OutputDir == "" {
		return fmt.Errorf("output directory is empty")
	}
	target, err := filepath.Abs(ow.OutputDir)
	if err != nil {
		return fmt.Errorf("resolve output directory: %w", err)
	}
	target = filepath.Clean(target)
	if target == filepath.Dir(target) {
		return fmt.Errorf("refusing filesystem root as output directory")
	}
	if err := validateExistingOutputAncestor(target); err != nil {
		return err
	}

	info, err := os.Lstat(target)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(target, 0755); err != nil {
			return fmt.Errorf("create output directory: %w", err)
		}
		if err := validateRealDirectory(target, "output directory"); err != nil {
			return err
		}
		marker := filepath.Join(target, outputOwnershipMarker)
		if _, markerErr := os.Lstat(marker); !os.IsNotExist(markerErr) {
			if markerErr != nil {
				return fmt.Errorf("inspect output ownership marker: %w", markerErr)
			}
			return fmt.Errorf("output ownership marker unexpectedly exists: %s", marker)
		}
		if err := os.WriteFile(marker, []byte(outputOwnershipContent), 0644); err != nil {
			return fmt.Errorf("write output ownership marker: %w", err)
		}
	} else if err != nil {
		return fmt.Errorf("inspect output directory: %w", err)
	} else {
		if info.Mode()&os.ModeSymlink != 0 || !info.IsDir() {
			return fmt.Errorf("output directory must be a real directory: %s", target)
		}
		if err := validateRealDirectory(target, "output directory"); err != nil {
			return err
		}
		marker := filepath.Join(target, outputOwnershipMarker)
		markerInfo, markerErr := os.Lstat(marker)
		if markerErr != nil {
			if os.IsNotExist(markerErr) {
				return fmt.Errorf("refusing unowned existing output directory: %s", target)
			}
			return fmt.Errorf("inspect output ownership marker: %w", markerErr)
		}
		if markerInfo.Mode()&os.ModeSymlink != 0 || !markerInfo.Mode().IsRegular() {
			return fmt.Errorf("output ownership marker is not a regular file: %s", marker)
		}
		contents, readErr := os.ReadFile(marker)
		if readErr != nil {
			return fmt.Errorf("read output ownership marker: %w", readErr)
		}
		if string(contents) != outputOwnershipContent {
			return fmt.Errorf("output ownership marker has unexpected content: %s", marker)
		}
	}

	ow.OutputDir = target
	return nil
}

func validateExistingOutputAncestor(target string) error {
	for current := target; ; current = filepath.Dir(current) {
		info, err := os.Lstat(current)
		if os.IsNotExist(err) {
			parent := filepath.Dir(current)
			if parent == current {
				return fmt.Errorf("no existing ancestor for output directory %s", target)
			}
			continue
		}
		if err != nil {
			return fmt.Errorf("inspect output ancestor %s: %w", current, err)
		}
		if info.Mode()&os.ModeSymlink != 0 || !info.IsDir() {
			return fmt.Errorf("output ancestor must be a real directory: %s", current)
		}
		return nil
	}
}

func validateRealDirectory(path, label string) error {
	info, err := os.Lstat(path)
	if err != nil {
		return fmt.Errorf("inspect %s: %w", label, err)
	}
	if info.Mode()&os.ModeSymlink != 0 || !info.IsDir() {
		return fmt.Errorf("%s must be a real directory: %s", label, path)
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
		"pcileech_bram_disk.sv",
		"tlp_latency_emulator.sv",
		"device_config.sv",
		"config_space_init.hex",
		"msix_table_init.hex",
		"scrub_diff_report.txt",
		"build_manifest.json",
	}
}
