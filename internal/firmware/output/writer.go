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

	if err := ow.writeConfigSpaceArtifacts(ctx, scrubbedCS, b); err != nil {
		return err
	}

	if err := ow.writeBARBehaviorProfile(ctx); err != nil {
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
	if err := ow.writeFile("pcileech_bar_zero4k.coe",
		codegen.GenerateBarContentCOE(ctx.BARContents, firmware.CappedBAR0Size(ctx, b, msixTableSize))); err != nil {
		return fmt.Errorf("failed to write bar zero COE: %w", err)
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
		if err := extractSubModules(ctrlSrc, dstDir, barControllerSubModules); err != nil {
			slog.Warn("could not extract BAR controller sub-modules, board source may be incompatible", "error", err)
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
		"pcileech_cfgspace.coe",
		"pcileech_cfgspace_writemask.coe",
		"pcileech_bar_zero4k.coe",
		"bar_behavior_profile.json",
		"vivado_generate_project.tcl",
		"vivado_build.tcl",
		"src/",
		"pcileech_bar_impl_device.sv",
		"pcileech_tlps128_bar_controller.sv",
		"pcileech_bar_impl_msi.sv",
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
func (ow *OutputWriter) writeSVModules(ctx *donor.DeviceContext, scrubbedCS *pci.ConfigSpace, ids firmware.DeviceIDs, entropy uint32, b *board.Board) error {
	cfg, err := ow.buildSVConfig(ctx, scrubbedCS, ids, entropy, b)
	if err != nil {
		return err
	}

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
func (ow *OutputWriter) buildSVConfig(ctx *donor.DeviceContext, scrubbedCS *pci.ConfigSpace, ids firmware.DeviceIDs, entropy uint32, b *board.Board) (*svgen.SVGeneratorConfig, error) {
	// Use the same BAR index for content data and probe profile to avoid
	// mismatched register maps (e.g. IO BAR0 profile + MMIO BAR2 data).
	barIdx := firmware.LargestBarIndex(ctx.BARContents)
	// NVMe CAP/VS/CC/CSTS live in BAR0; force it even when BAR0 isn't the largest.
	if ctx.Device.ClassCode == devclass.ClassCodeNVMe {
		if _, ok := ctx.BARContents[0]; ok {
			barIdx = 0
		}
	}
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
		"reg_count", func() int {
			if bm != nil {
				return len(bm.Registers)
			}
			return 0
		}(),
	)

	strategy := devclass.StrategyForClassAndVendor(ctx.Device.ClassCode, ids.VendorID)
	devClass := ""
	if strategy != nil {
		devClass = strategy.DeviceClass()
	}
	slog.Info("device class resolution", "class", devClass)

	msixTableSize := 0
	if ctx.MSIXData != nil && ctx.MSIXData.TableSize > 0 {
		msixTableSize = ctx.MSIXData.TableSize
	}
	donorBar := firmware.DonorBAR0Demand(ctx, b, msixTableSize)
	bar0Size := firmware.CappedBAR0Size(ctx, b, msixTableSize)
	if bm != nil && bm.Size < bar0Size {
		bm.Size = bar0Size
	}
	if bm == nil && bar0Size > board.DefaultBRAMSize {
		bm = &barmodel.BARModel{Size: bar0Size}
	}

	cfg := &svgen.SVGeneratorConfig{
		DeviceIDs:         ids,
		DonorCapabilities: extractDonorCapabilities(scrubbedCS),
		BARModel:          bm,
		ClassCode:         ctx.Device.ClassCode,
		LatencyConfig:     svgen.DefaultLatencyConfig(ctx.Device.ClassCode),
		HasMSIX:           bm != nil,
		BuildEntropy:      entropy,
		PRNGSeeds:         svgen.BuildPRNGSeeds(ids.VendorID, ids.DeviceID, entropy),
		DeviceClass:       devClass,
		Bar0Size:          bar0Size,
	}

	if devClass == devclass.ClassNVMe {
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
		// Pass the full class code: MSIXPlacement detects NVMe via
		// (class>>8)==0x0108, so the truncated 0x0108 would skip the
		// doorbell-avoidance guard.
		tableOff, pbaOffset, _ := firmware.MSIXPlacement(cfg.Bar0Size, ctx.MSIXData.TableSize, ctx.Device.ClassCode, cfg.NVMeDoorbellStride)
		cfg.MSIXConfig = &svgen.MSIXConfig{
			NumVectors:  ctx.MSIXData.TableSize,
			TableOffset: tableOff,
			PBAOffset:   pbaOffset,
		}
		cfg.HasMSIX = true
		if m := pci.ParseMSIXCap(scrubbedCS); m != nil {
			scrubbedCS.WriteU32(m.CapOffset+4, tableOff&0xFFFFFFF8)
			scrubbedCS.WriteU32(m.CapOffset+8, pbaOffset&0xFFFFFFF8)
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
	if cfg.MSIXConfig != nil {
		if issues := ValidateBARSize(donorBar, bram, cfg.MSIXConfig.TableOffset); len(issues) > 0 {
			if !ow.Force {
				return nil, fmt.Errorf("%s", issues[0])
			}
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
		if err := ow.writeFile("pcileech_nvme_dma_bridge.sv", bridgeSV); err != nil {
			return err
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
