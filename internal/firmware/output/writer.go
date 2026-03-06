// Package output writes the final firmware artifacts (COE, TCL, SV, HEX).
package output

import (
	"fmt"
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

// WriteAll is the main entry point — generates everything.
func (ow *OutputWriter) WriteAll(ctx *donor.DeviceContext, b *board.Board) error {
	if err := os.MkdirAll(ow.OutputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	ids := firmware.ExtractDeviceIDs(ctx.ConfigSpace, ctx.ExtCapabilities)

	if err := ow.writeDeviceContext(ctx); err != nil {
		return err
	}

	scrubbedCS, entropy := ow.scrubAndVary(ctx, b, ids)

	if err := ow.writeConfigSpaceArtifacts(ctx, scrubbedCS); err != nil {
		return err
	}

	if err := ow.writeTCLScripts(ctx, b); err != nil {
		return err
	}

	if err := ow.patchSVSources(ctx, b, ids); err != nil {
		return fmt.Errorf("SV patching failed: %w", err)
	}

	if err := ow.writeSVModules(ctx, scrubbedCS, ids, entropy); err != nil {
		return fmt.Errorf("SV module generation failed: %w", err)
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
func (ow *OutputWriter) scrubAndVary(ctx *donor.DeviceContext, b *board.Board, ids firmware.DeviceIDs) (*pci.ConfigSpace, uint32) {
	scrubbedCS, overlayMap := scrub.ScrubConfigSpaceWithOverlay(ctx.ConfigSpace, b)
	if overlayMap.Count() > 0 {
		fmt.Printf("[firmware] Config space scrub: %d modifications\n", overlayMap.Count())
	}

	entropy := svgen.BuildEntropyFromTime()
	varSeed := variance.BuildVarianceSeed(ids.VendorID, ids.DeviceID, entropy)
	variance.Apply(scrubbedCS, nil, variance.DefaultConfig(varSeed))

	return scrubbedCS, entropy
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

	scrub.ScrubBarContent(ctx.BARContents, ctx.Device.ClassCode)
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
		fmt.Printf("[firmware] Warning: manifest generation failed: %v\n", err)
		return
	}
	manifestPath := ow.OutputDir + "/build_manifest.json"
	if err := manifest.WriteJSON(manifestPath); err != nil {
		fmt.Printf("[firmware] Warning: manifest write failed: %v\n", err)
	} else {
		fmt.Printf("[firmware] Build manifest: %d files recorded\n", len(manifest.Files))
	}
}

// patchSVSources copies the board's SV tree and patches IDs in.
func (ow *OutputWriter) patchSVSources(ctx *donor.DeviceContext, b *board.Board, ids firmware.DeviceIDs) error {
	srcDir := b.SrcPath(ow.LibDir)
	dstDir := filepath.Join(ow.OutputDir, "src")

	if _, err := os.Stat(srcDir); os.IsNotExist(err) {
		fmt.Printf("[firmware] Warning: board source dir not found: %s\n", srcDir)
		fmt.Println("[firmware] Run: git submodule update --init --recursive")
		return fmt.Errorf("board sources not found at %s (is the pcileech-fpga submodule initialized?)", srcDir)
	}

	if err := util.CopyDir(srcDir, dstDir); err != nil {
		return fmt.Errorf("failed to copy SV sources: %w", err)
	}

	patcher := svgen.NewSVPatcher(ids, dstDir)

	if err := patcher.PatchAll(); err != nil {
		return fmt.Errorf("failed to patch SV sources: %w", err)
	}

	if results := patcher.Results(); len(results) > 0 {
		fmt.Println("[firmware] SV patches applied:")
		fmt.Print(svgen.FormatPatchSummary(results))
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
		"tlp_latency_emulator.sv",
		"device_config.sv",
		"config_space_init.hex",
		"msix_table_init.hex",
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
	barData := firmware.LowestBarData(ctx.BARContents)
	var barProfile *donor.BARProfile
	if ctx.BARProfiles != nil {
		barProfile = firmware.LowestBarProfile(ctx.BARProfiles)
	}
	bm := barmodel.BuildBARModel(barData, ctx.Device.ClassCode, barProfile)

	strategy := devclass.StrategyForClass(ctx.Device.ClassCode)
	isNVMe := strategy != nil && strategy.IsNVMe()
	isXHCI := strategy != nil && strategy.IsXHCI()

	latCfg := svgen.DefaultLatencyConfig(ctx.Device.ClassCode)
	seeds := svgen.BuildPRNGSeeds(ids.VendorID, ids.DeviceID, entropy)

	cfg := &svgen.SVGeneratorConfig{
		DeviceIDs:     ids,
		BARModel:      bm,
		ClassCode:     ctx.Device.ClassCode,
		LatencyConfig: latCfg,
		HasMSIX:       bm != nil,
		BuildEntropy:  entropy,
		PRNGSeeds:     seeds,
		IsNVMe:        isNVMe,
		IsXHCI:        isXHCI,
	}

	// NVMe Identify data generation
	if isNVMe {
		cfg.NVMeIdentify = nvme.BuildIdentifyData(ids, barData)
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

	// Generate core SV artifacts via pipeline
	for _, art := range coreSVArtifacts {
		content, err := art.generate(cfg)
		if err != nil {
			return fmt.Errorf("generating %s: %w", art.filename, err)
		}
		if err := ow.writeFile(art.filename, content); err != nil {
			return err
		}
	}

	// config_space_init.hex
	hex := codegen.GenerateConfigSpaceHex(scrubbedCS)
	if err := ow.writeFile("config_space_init.hex", hex); err != nil {
		return err
	}

	// MSI-X artifacts (conditional)
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
		msixHex := codegen.GenerateMSIXTableHex(entries)
		if err := ow.writeFile("msix_table_init.hex", msixHex); err != nil {
			return err
		}
	}

	// NVMe Responder artifacts (conditional)
	if cfg.NVMeIdentify != nil {
		nvmeSV, err := svgen.GenerateNVMeResponderSV(cfg)
		if err != nil {
			return fmt.Errorf("generating pcileech_nvme_admin_responder.sv: %w", err)
		}
		if err := ow.writeFile("pcileech_nvme_admin_responder.sv", nvmeSV); err != nil {
			return err
		}

		idHex := nvme.IdentifyDataToHex(cfg.NVMeIdentify)
		if err := ow.writeFile("identify_init.hex", idHex); err != nil {
			return err
		}
	}

	features := []string{}
	if isNVMe {
		features = append(features, "NVMe FSM")
		if cfg.NVMeIdentify != nil {
			features = append(features, "NVMe Admin Responder")
		}
	}
	if isXHCI {
		features = append(features, "xHCI FSM")
	}
	if cfg.MSIXConfig != nil {
		features = append(features, fmt.Sprintf("MSI-X %d vectors", cfg.MSIXConfig.NumVectors))
	}
	if bm != nil {
		features = append(features, fmt.Sprintf("%d registers", len(bm.Registers)))
	} else {
		features = append(features, "BRAM fallback")
	}
	features = append(features, "latency emulator", "interrupt controller")
	fmt.Printf("[firmware] SV modules generated: %s\n", strings.Join(features, ", "))

	return nil
}
