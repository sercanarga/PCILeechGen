package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

type hdlCheckOptions struct {
	dir    string
	dryRun bool
}

var hdlCheckOpts hdlCheckOptions

var hdlCheckCmd = &cobra.Command{
	Use:   "hdl-check",
	Short: "Run generated HDL/TCL lint checks",
	Long: `Runs Verilator/TCL validation on generated artifacts.

Set --output-dir to the firmware output directory that contains generated .sv/.v/.tcl files.
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runHDLCheck(cmd.OutOrStdout(), hdlCheckOpts)
	},
}

func buildHDLCheckCommand(opts hdlCheckOptions) []string {
	args := []string{"run", "./tools/hdlcheck", "--dir", opts.dir}
	if opts.dryRun {
		args = append(args, "--dry-run")
	}
	return args
}

func runHDLCheck(out io.Writer, opts hdlCheckOptions) error {
	if _, err := os.Stat(opts.dir); err != nil {
		return fmt.Errorf("check output directory %q: %w", opts.dir, err)
	}

	if _, err := exec.LookPath("go"); err != nil {
		return fmt.Errorf("go toolchain not found: %w", err)
	}

	args := buildHDLCheckCommand(opts)
	cmd := execCommand("go", args...)
	cmd.Stdout = out
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("hdl-check failed: %w", err)
	}
	return nil
}

var execCommand = exec.Command

func init() {
	hdlCheckCmd.Flags().StringVarP(&hdlCheckOpts.dir, "output-dir", "o", ".", "path to generated firmware output directory")
	hdlCheckCmd.Flags().BoolVar(&hdlCheckOpts.dryRun, "dry-run", false, "print checks without running them")
	rootCmd.AddCommand(hdlCheckCmd)
}
