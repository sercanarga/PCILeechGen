package main

import (
	"fmt"
	"os"

	"github.com/sercanarga/pcileechgen/internal/color"
	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/spf13/cobra"
)

var (
	importFormat   string
	importOutput   string
	importOverwrite bool
)

// importCmd converts external donor captures (raw config-space dumps or COE
// write-mask files emitted by the generator) into a donor.DeviceContext JSON
// file. It is a parser boundary only: it never emits HDL, SV, or COE output,
// and never invokes the firmware output writer. Missing optional donor data
// (BAR contents, MSI-X entries) is surfaced as warnings, not errors; the
// existing `validate` command decides build-readiness.
var importCmd = &cobra.Command{
	Use:   "import <capture-file>",
	Short: "Convert an external donor capture to device_context.json",
	Long: `Converts an external donor capture into a donor.DeviceContext JSON file.

Supported formats:
  --format raw   raw little-endian PCI config space (256 or 4096 bytes)
  --format coe   COE vector emitted by codegen.GenerateConfigSpaceCOE (4KB)

The importer never emits HDL. Missing optional donor fields (BAR contents,
MSI-X table) are reported as warnings; pipe the result through
'pcileechgen validate' to decide whether the context is build-ready.

Examples:
  pcileechgen import dump.bin --format raw -o device_context.json
  pcileechgen import pcileech_cfgspace.coe --format coe -o device_context.json`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		capturePath := args[0]

		if importOutput == "" {
			return fmt.Errorf("--output is required (path to device_context.json)")
		}
		if _, err := os.Stat(importOutput); err == nil && !importOverwrite {
			return fmt.Errorf("refusing to overwrite %s without --overwrite", importOutput)
		}

		ctx, warnings, err := runImport(capturePath, importFormat)
		if err != nil {
			// No partial output: return before touching the output file.
			return err
		}

		data, err := ctx.ToJSON()
		if err != nil {
			return fmt.Errorf("failed to marshal device context: %w", err)
		}
		if err := os.WriteFile(importOutput, data, 0644); err != nil {
			return fmt.Errorf("failed to write %s: %w", importOutput, err)
		}

		fmt.Printf("Imported %s -> %s\n", capturePath, color.Bold(importOutput))
		fmt.Printf("Device: %s\n",
			color.Bold(fmt.Sprintf("%04x:%04x (rev %02x)", ctx.Device.VendorID, ctx.Device.DeviceID, ctx.Device.RevisionID)))
		if len(warnings) > 0 {
			fmt.Println()
			for _, w := range warnings {
				fmt.Println(color.Warn(w.String()))
			}
		}
		fmt.Printf("\n%s\n", color.Info("Run 'pcileechgen validate --json "+importOutput+"' to check build-readiness."))
		return nil
	},
}

// runImport dispatches on --format and returns the parsed DeviceContext plus
// structured warnings. It performs no I/O on the output path, so a failure
// here guarantees no partial output file.
func runImport(capturePath, format string) (*donor.DeviceContext, []donor.Warning, error) {
	switch format {
	case "raw":
		raw, err := os.ReadFile(capturePath)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to read capture: %w", err)
		}
		return donor.ImportConfigSpace(raw)
	case "coe":
		return donor.ImportCOE(capturePath)
	default:
		return nil, nil, fmt.Errorf("unsupported --format %q (want 'raw' or 'coe')", format)
	}
}

func init() {
	importCmd.Flags().StringVar(&importFormat, "format", "raw", "capture format: raw | coe")
	importCmd.Flags().StringVarP(&importOutput, "output", "o", "", "path to write device_context.json (required)")
	importCmd.Flags().BoolVar(&importOverwrite, "overwrite", false, "allow overwriting an existing output file")
	_ = importCmd.MarkFlagRequired("output")
	rootCmd.AddCommand(importCmd)
}