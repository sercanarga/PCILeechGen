package vivado

import (
	"fmt"
	"log/slog"
	"path/filepath"
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
	if err := ow.WriteAll(ctx, b.board); err != nil {
		return fmt.Errorf("artifact generation failed: %w", err)
	}

	slog.Info("artifacts written", "dir", b.opts.OutputDir)
	for _, f := range fwout.ListOutputFiles() {
		slog.Info("artifact", "file", f)
	}

	if b.opts.SkipVivado {
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

	slog.Info("build completed successfully")
	return nil
}
