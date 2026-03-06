package vivado

import (
	"fmt"
	"log"
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
	opts   BuildOptions
	board  *board.Board
	logger *log.Logger
}

// NewBuilder creates a new Builder.
func NewBuilder(b *board.Board, opts BuildOptions) *Builder {
	opts = opts.WithDefaults()
	return &Builder{
		opts:   opts,
		board:  b,
		logger: log.Default(),
	}
}

// Build generates firmware artifacts and optionally runs Vivado.
func (b *Builder) Build(ctx *donor.DeviceContext) error {
	// Stage 2: Generate firmware artifacts
	b.logger.Println("[build] Stage 2: Generating firmware artifacts...")
	ow := fwout.NewOutputWriter(b.opts.OutputDir, b.opts.LibDir, b.opts.Jobs, b.opts.Timeout)
	if err := ow.WriteAll(ctx, b.board); err != nil {
		return fmt.Errorf("artifact generation failed: %w", err)
	}

	b.logger.Printf("[build] Artifacts written to: %s\n", b.opts.OutputDir)
	for _, f := range fwout.ListOutputFiles() {
		b.logger.Printf("  - %s\n", f)
	}

	if b.opts.SkipVivado {
		b.logger.Println("[build] Vivado synthesis skipped (--skip-vivado)")
		return nil
	}

	// Stage 3: Run Vivado synthesis
	b.logger.Println("[build] Stage 3: Running Vivado synthesis...")

	vivado, err := Find(b.opts.VivadoPath)
	if err != nil {
		return fmt.Errorf("Vivado not found: %w", err)
	}
	b.logger.Printf("[build] Using Vivado %s at %s\n", vivado.Version, vivado.Path)

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
		b.logger.Printf("[build] Bitstream: %s\n", f)
	}
	for _, f := range binFiles {
		b.logger.Printf("[build] Binary: %s\n", f)
	}

	for _, f := range append(bitFiles, binFiles...) {
		dst := filepath.Join(b.opts.OutputDir, filepath.Base(f))
		if err := util.CopyFile(f, dst); err != nil {
			b.logger.Printf("[build] Warning: failed to copy %s: %v\n", f, err)
		}
	}

	b.logger.Println("[build] Build completed successfully!")
	return nil
}
