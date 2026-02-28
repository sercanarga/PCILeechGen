package firmware

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/util"
)

// OutputWriter generates COE, TCL, and patched SV files.
type OutputWriter struct {
	OutputDir string
	LibDir    string
}

// NewOutputWriter creates a new OutputWriter.
func NewOutputWriter(outputDir, libDir string) *OutputWriter {
	return &OutputWriter{
		OutputDir: outputDir,
		LibDir:    libDir,
	}
}

// WriteAll writes COE, TCL, and patched SV files to the output directory.
func (ow *OutputWriter) WriteAll(ctx *donor.DeviceContext, b *board.Board) error {
	if err := os.MkdirAll(ow.OutputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Device context JSON
	data, err := ctx.ToJSON()
	if err != nil {
		return fmt.Errorf("failed to marshal device context: %w", err)
	}
	if err := ow.writeFile("device_context.json", string(data)); err != nil {
		return fmt.Errorf("failed to write device context: %w", err)
	}

	// Scrub config space before COE generation (clean dangerous/error registers)
	scrubbedCS := ScrubConfigSpace(ctx.ConfigSpace)

	// COE files (using scrubbed config space for shadow BRAM)
	if err := ow.writeFile("pcileech_cfgspace.coe",
		GenerateConfigSpaceCOE(scrubbedCS)); err != nil {
		return fmt.Errorf("failed to write cfgspace COE: %w", err)
	}
	if err := ow.writeFile("pcileech_cfgspace_writemask.coe",
		GenerateWritemaskCOE(scrubbedCS)); err != nil {
		return fmt.Errorf("failed to write writemask COE: %w", err)
	}
	if err := ow.writeFile("pcileech_bar_zero4k.coe",
		GenerateBarContentCOE(ctx.BARContents)); err != nil {
		return fmt.Errorf("failed to write bar zero COE: %w", err)
	}

	// TCL scripts
	if err := ow.writeFile("vivado_generate_project.tcl",
		GenerateProjectTCL(ctx, b, ow.LibDir)); err != nil {
		return fmt.Errorf("failed to write project TCL: %w", err)
	}
	if err := ow.writeFile("vivado_build.tcl",
		GenerateBuildTCL(b, 4, 3600)); err != nil {
		return fmt.Errorf("failed to write build TCL: %w", err)
	}

	// Copy board SV sources and apply donor patches
	if err := ow.patchSVSources(ctx, b); err != nil {
		return fmt.Errorf("SV patching failed: %w", err)
	}

	return nil
}

// patchSVSources copies board SV source files to output, then applies donor patches.
// Original pcileech-fpga files are never modified.
func (ow *OutputWriter) patchSVSources(ctx *donor.DeviceContext, b *board.Board) error {
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

	ids := ExtractDeviceIDs(ctx.ConfigSpace, ctx.ExtCapabilities)
	patcher := NewSVPatcher(ids, dstDir)

	if err := patcher.PatchAll(); err != nil {
		return fmt.Errorf("failed to patch SV sources: %w", err)
	}

	if results := patcher.Results(); len(results) > 0 {
		fmt.Println("[firmware] SV patches applied:")
		fmt.Print(FormatPatchSummary(results))
	}

	return nil
}

func (ow *OutputWriter) writeFile(name, content string) error {
	return os.WriteFile(filepath.Join(ow.OutputDir, name), []byte(content), 0644)
}

// ListOutputFiles returns a list of files that will be generated.
func ListOutputFiles() []string {
	return []string{
		"device_context.json",
		"pcileech_cfgspace.coe",
		"pcileech_cfgspace_writemask.coe",
		"pcileech_bar_zero4k.coe",
		"vivado_generate_project.tcl",
		"vivado_build.tcl",
		"src/",
	}
}
