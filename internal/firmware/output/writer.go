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

	// extract IDs once — reused by patching and codegen
	ids := firmware.ExtractDeviceIDs(ctx.ConfigSpace, ctx.ExtCapabilities)

	// save raw device context
	data, err := ctx.ToJSON()
	if err != nil {
		return fmt.Errorf("failed to marshal device context: %w", err)
	}
	if err := ow.writeFile("device_context.json", string(data)); err != nil {
		return fmt.Errorf("failed to write device context: %w", err)
	}

	// scrub + overlay audit
	scrubbedCS, overlayMap := scrub.ScrubConfigSpaceWithOverlay(ctx.ConfigSpace, b)
	if overlayMap.Count() > 0 {
		fmt.Printf("[firmware] Config space scrub: %d modifications\n", overlayMap.Count())
	}

	// per-build variance so each bitstream differs slightly
	entropy := svgen.BuildEntropyFromTime()
	varSeed := variance.BuildVarianceSeed(ids.VendorID, ids.DeviceID, entropy)
	variance.Apply(scrubbedCS, nil, variance.DefaultConfig(varSeed))

	// COE: config + writemask
	if err := ow.writeFile("pcileech_cfgspace.coe",
		codegen.GenerateConfigSpaceCOE(scrubbedCS)); err != nil {
		return fmt.Errorf("failed to write cfgspace COE: %w", err)
	}
	if err := ow.writeFile("pcileech_cfgspace_writemask.coe",
		codegen.GenerateWritemaskCOE(scrubbedCS)); err != nil {
		return fmt.Errorf("failed to write writemask COE: %w", err)
	}

	// fix device-class quirks in BAR content (e.g. NVMe CSTS.RDY)
	scrub.ScrubBarContent(ctx.BARContents, ctx.Device.ClassCode)

	if err := ow.writeFile("pcileech_bar_zero4k.coe",
		codegen.GenerateBarContentCOE(ctx.BARContents)); err != nil {
		return fmt.Errorf("failed to write bar zero COE: %w", err)
	}

	// TCL
	if err := ow.writeFile("vivado_generate_project.tcl",
		tclgen.GenerateProjectTCL(ctx, b, ow.LibDir)); err != nil {
		return fmt.Errorf("failed to write project TCL: %w", err)
	}
	if err := ow.writeFile("vivado_build.tcl",
		tclgen.GenerateBuildTCL(b, ow.Jobs, ow.Timeout)); err != nil {
		return fmt.Errorf("failed to write build TCL: %w", err)
	}

	// copy board SV and patch donor IDs in
	if err := ow.patchSVSources(ctx, b, ids); err != nil {
		return fmt.Errorf("SV patching failed: %w", err)
	}

	// generate per-device SV modules + HEX
	if err := ow.writeSVModules(ctx, scrubbedCS, ids, entropy); err != nil {
		return fmt.Errorf("SV module generation failed: %w", err)
	}

	// build manifest with checksums
	manifest, err := GenerateManifest(ow.OutputDir, ctx.ToolVersion, "", ids.VendorID, ids.DeviceID)
	if err != nil {
		fmt.Printf("[firmware] Warning: manifest generation failed: %v\n", err)
	} else {
		manifestPath := ow.OutputDir + "/build_manifest.json"
		if err := manifest.WriteJSON(manifestPath); err != nil {
			fmt.Printf("[firmware] Warning: manifest write failed: %v\n", err)
		} else {
			fmt.Printf("[firmware] Build manifest: %d files recorded\n", len(manifest.Files))
		}
	}

	return nil
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
		"tlp_latency_emulator.sv",
		"device_config.sv",
		"config_space_init.hex",
		"msix_table_init.hex",
	}
}

// writeSVModules generates device-specific SV and HEX init files.
func (ow *OutputWriter) writeSVModules(ctx *donor.DeviceContext, scrubbedCS *pci.ConfigSpace, ids firmware.DeviceIDs, entropy uint32) error {
	barData := firmware.LowestBarData(ctx.BARContents)
	var barProfile *donor.BARProfile
	if ctx.BARProfiles != nil {
		barProfile = firmware.LowestBarProfile(ctx.BARProfiles)
	}
	bm := barmodel.BuildBARModel(barData, ctx.Device.ClassCode, barProfile)

	baseClass := (ctx.Device.ClassCode >> 16) & 0xFF
	subClass := (ctx.Device.ClassCode >> 8) & 0xFF
	progIF := ctx.Device.ClassCode & 0xFF
	isNVMe := baseClass == 0x01 && subClass == 0x08 && progIF == 0x02
	isXHCI := baseClass == 0x0C && subClass == 0x03 && progIF == 0x30

	latCfg := svgen.DefaultLatencyConfig(ctx.Device.ClassCode)
	seeds := svgen.BuildPRNGSeeds(ids.VendorID, ids.DeviceID, entropy)

	cfg := &svgen.SVGeneratorConfig{
		DeviceIDs:     ids,
		BARModel:      bm,
		ClassCode:     ctx.Device.ClassCode,
		LatencyConfig: latCfg,
		HasMSIX:       bm != nil, // MSI-X for known device classes
		BuildEntropy:  entropy,
		PRNGSeeds:     seeds,
		IsNVMe:        isNVMe,
		IsXHCI:        isXHCI,
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

	// pcileech_bar_impl_device.sv
	barImpl, err := svgen.GenerateBarImplDeviceSV(cfg)
	if err != nil {
		return fmt.Errorf("generating bar_impl_device.sv: %w", err)
	}
	if err := ow.writeFile("pcileech_bar_impl_device.sv", barImpl); err != nil {
		return err
	}

	// pcileech_tlps128_bar_controller.sv
	barCtrl, err := svgen.GenerateBarControllerSV(cfg)
	if err != nil {
		return fmt.Errorf("generating bar_controller.sv: %w", err)
	}
	if err := ow.writeFile("pcileech_tlps128_bar_controller.sv", barCtrl); err != nil {
		return err
	}

	// tlp_latency_emulator.sv
	latEmu, err := svgen.GenerateLatencyEmulatorSV(cfg)
	if err != nil {
		return fmt.Errorf("generating tlp_latency_emulator.sv: %w", err)
	}
	if err := ow.writeFile("tlp_latency_emulator.sv", latEmu); err != nil {
		return err
	}

	// device_config.sv
	devCfg, err := svgen.GenerateDeviceConfigSV(cfg)
	if err != nil {
		return fmt.Errorf("generating device_config.sv: %w", err)
	}
	if err := ow.writeFile("device_config.sv", devCfg); err != nil {
		return err
	}

	// config_space_init.hex
	hex := codegen.GenerateConfigSpaceHex(scrubbedCS)
	if err := ow.writeFile("config_space_init.hex", hex); err != nil {
		return err
	}

	// pcileech_msix_table.sv + msix_table_init.hex
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

	features := []string{}
	if isNVMe {
		features = append(features, "NVMe FSM")
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
