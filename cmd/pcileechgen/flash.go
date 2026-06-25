package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/flash"
	"github.com/spf13/cobra"
)

type flashOptions struct {
	board     string
	bitstream string
	dryRun    bool
}

var flashOpts flashOptions

var flashCmd = &cobra.Command{
	Use:           "flash",
	Short:         "Program a generated bitstream onto a supported FPGA board",
	SilenceUsage:  true,
	SilenceErrors: true,
	Long: `Programs a generated bitstream with the board's configured flash provider.

Use --dry-run first to print the exact command without touching hardware.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		b, err := board.Find(flashOpts.board)
		if err != nil {
			return err
		}
		return runFlash(os.Stdout, b, flashOpts)
	},
}

func runFlash(out io.Writer, b *board.Board, opts flashOptions) error {
	cmd, err := flash.OpenFPGALoaderCommand(flash.Request{
		Board:     b,
		Bitstream: opts.bitstream,
	})
	if err != nil {
		return err
	}

	if opts.dryRun {
		fmt.Fprintf(out, "Dry run: %s\n", cmd.String())
		return nil
	}
	if _, err := os.Stat(opts.bitstream); err != nil {
		return fmt.Errorf("check bitstream %q: %w", opts.bitstream, err)
	}
	if _, err := exec.LookPath(cmd.Args[0]); err != nil {
		return fmt.Errorf("%s not found in PATH: %w", cmd.Args[0], err)
	}

	proc := exec.Command(cmd.Args[0], cmd.Args[1:]...)
	proc.Stdout = out
	proc.Stderr = os.Stderr
	if err := proc.Run(); err != nil {
		return fmt.Errorf("run %s: %w", cmd.Args[0], err)
	}
	return nil
}

func init() {
	flashCmd.Flags().StringVar(&flashOpts.board, "board", "", "target FPGA board name (required)")
	flashCmd.Flags().StringVar(&flashOpts.bitstream, "bitstream", "", "path to generated bitstream (.bit)")
	flashCmd.Flags().BoolVar(&flashOpts.dryRun, "dry-run", false, "print the flash command without running it")

	_ = flashCmd.MarkFlagRequired("board")
	_ = flashCmd.MarkFlagRequired("bitstream")
	rootCmd.AddCommand(flashCmd)
}
