package vivado

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/donor"
	fwout "github.com/sercanarga/pcileechgen/internal/firmware/output"
	"github.com/sercanarga/pcileechgen/internal/util"
)

// BuildOptions holds build configuration.
type BuildOptions struct {
	VivadoPath string
	OutputDir  string
	LibDir     string
	Jobs       int
	Timeout    int
	SkipVivado bool
	StockBar   bool
	Force      bool
}

// WithDefaults returns a copy of opts with zero values replaced by sensible defaults.
func (opts BuildOptions) WithDefaults() BuildOptions {
	if opts.Jobs <= 0 {
		opts.Jobs = 4
	}
	if opts.Timeout <= 0 {
		opts.Timeout = 3600
	}
	if opts.OutputDir == "" {
		opts.OutputDir = "pcileech_datastore"
	}
	return opts
}

// Builder runs firmware generation and optional Vivado synthesis.
type Builder struct {
	opts  BuildOptions
	board *board.Board
}

// NewBuilder creates a new Builder.
func NewBuilder(b *board.Board, opts BuildOptions) *Builder {
	opts = opts.WithDefaults()
	return &Builder{
		opts:  opts,
		board: b,
	}
}

// Build generates firmware artifacts and optionally runs Vivado.
func (b *Builder) Build(ctx *donor.DeviceContext) error {
	// Stage 2: Generate firmware artifacts
	slog.Info("generating firmware artifacts")
	ow := fwout.NewOutputWriter(b.opts.OutputDir, b.opts.LibDir, b.opts.Jobs, b.opts.Timeout)
	ow.StockBar = b.opts.StockBar
	ow.Force = b.opts.Force
	if err := ow.WriteAll(ctx, b.board); err != nil {
		return fmt.Errorf("artifact generation failed: %w", err)
	}
	if _, err := fwout.VerifyManifest(filepath.Join(b.opts.OutputDir, "build_manifest.json"), b.opts.OutputDir); err != nil {
		return fmt.Errorf("artifact manifest verification failed: %w", err)
	}

	slog.Info("artifacts written", "dir", b.opts.OutputDir)
	for _, f := range fwout.ListOutputFiles() {
		slog.Info("artifact", "file", f)
	}

	if b.opts.SkipVivado {
		if err := writeBuildSummary(b.opts.OutputDir, ctx, b.board, nil, nil); err != nil {
			slog.Warn("write build summary", "error", err)
		}
		slog.Info("Vivado synthesis skipped")
		return nil
	}

	// Stage 3: Run Vivado synthesis
	slog.Info("running Vivado synthesis")

	vivado, err := Find(b.opts.VivadoPath)
	if err != nil {
		return fmt.Errorf("Vivado not found: %w", err)
	}
	slog.Info("Vivado found", "version", vivado.Version, "path", vivado.Path)

	timeout := time.Duration(b.opts.Timeout) * time.Second

	// Run project creation
	projectTCL := "vivado_generate_project.tcl"
	if err := vivado.RunTCL(projectTCL, b.opts.OutputDir, timeout); err != nil {
		return fmt.Errorf("project creation failed: %w", err)
	}

	// Run synthesis and implementation
	buildTCL := "vivado_build.tcl"
	if err := vivado.RunTCL(buildTCL, b.opts.OutputDir, timeout); err != nil {
		return fmt.Errorf("build failed: %w", err)
	}

	// Find and copy output files
	bitFiles, _ := filepath.Glob(filepath.Join(b.opts.OutputDir, b.board.Name, "*.runs", "impl_1", "*.bit"))
	binFiles, _ := filepath.Glob(filepath.Join(b.opts.OutputDir, "*.bin"))

	for _, f := range bitFiles {
		slog.Info("bitstream", "file", f)
	}
	for _, f := range binFiles {
		slog.Info("binary", "file", f)
	}

	for _, f := range append(bitFiles, binFiles...) {
		dst := filepath.Join(b.opts.OutputDir, filepath.Base(f))
		if err := util.CopyFile(f, dst); err != nil {
			slog.Warn("failed to copy output file", "file", f, "error", err)
		}
	}

	refreshBuildManifest(fwout.WriteBuildManifest, b.opts.OutputDir, ctx, b.board)
	if _, err := fwout.VerifyManifest(filepath.Join(b.opts.OutputDir, "build_manifest.json"), b.opts.OutputDir); err != nil {
		return fmt.Errorf("post-synthesis manifest verification failed: %w", err)
	}

	if err := chownOutputs(b.opts.OutputDir); err != nil {
		slog.Warn("chown output files", "error", err)
	}

	if err := writeBuildSummary(b.opts.OutputDir, ctx, b.board, bitFiles, binFiles); err != nil {
		slog.Warn("write build summary", "error", err)
	}
	slog.Info("build completed successfully")
	return nil
}

func writeBuildSummary(outputDir string, ctx *donor.DeviceContext, b *board.Board, bitFiles, binFiles []string) error {
	var sb strings.Builder
	fmt.Fprintf(&sb, "PCILeechGen build summary\n")
	fmt.Fprintf(&sb, "board=%s fpga=%s lanes=x%d\n", b.Name, b.FPGAPart, b.PCIeLanes)
	fmt.Fprintf(&sb, "device=%04x:%04x class=0x%06x revision=0x%02x\n",
		ctx.Device.VendorID, ctx.Device.DeviceID, ctx.Device.ClassCode, ctx.Device.RevisionID)
	fmt.Fprintf(&sb, "bars=%d capabilities=%d ext_capabilities=%d\n",
		len(ctx.BARs), len(ctx.Capabilities), len(ctx.ExtCapabilities))
	if ctx.MSIXData != nil {
		fmt.Fprintf(&sb, "msix_vectors=%d table_bir=%d table_offset=0x%x pba_bir=%d pba_offset=0x%x\n",
			ctx.MSIXData.TableSize, ctx.MSIXData.TableBIR, ctx.MSIXData.TableOffset,
			ctx.MSIXData.PBABIR, ctx.MSIXData.PBAOffset)
	}
	writeFileList(&sb, "bitstreams", bitFiles)
	writeFileList(&sb, "binaries", binFiles)
	return os.WriteFile(filepath.Join(outputDir, "build_summary.txt"), []byte(sb.String()), 0o644)
}

func writeFileList(sb *strings.Builder, title string, files []string) {
	fmt.Fprintf(sb, "%s=%d\n", title, len(files))
	for _, file := range files {
		info, err := os.Stat(file)
		if err != nil {
			fmt.Fprintf(sb, "- %s\n", file)
			continue
		}
		fmt.Fprintf(sb, "- %s (%d bytes)\n", file, info.Size())
	}
}

func refreshBuildManifest(
	writeManifest func(string, *donor.DeviceContext, *board.Board) error,
	outputDir string,
	ctx *donor.DeviceContext,
	b *board.Board,
) {
	if err := writeManifest(outputDir, ctx, b); err != nil {
		slog.Warn("failed to refresh post-synthesis build manifest; bitstream remains available", "error", err)
	}
}
