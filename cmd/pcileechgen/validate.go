package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/color"
	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/firmware"
	"github.com/sercanarga/pcileechgen/internal/firmware/codegen"
	"github.com/sercanarga/pcileechgen/internal/firmware/output"
	"github.com/sercanarga/pcileechgen/internal/firmware/scrub"
	"github.com/sercanarga/pcileechgen/internal/pci"
	"github.com/spf13/cobra"
)

var validateJSONPath string
var validateOutputDir string
var validateBoard string

// validator bundles state for a single validation run.
type validator struct {
	ctx       *donor.DeviceContext
	b         *board.Board
	outputDir string
	scrubbed  *pci.ConfigSpace
	result    *output.ValidationResult
}

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate generated firmware artifacts against donor data",
	Long: `Validates generated firmware artifacts (COE files) against the
original donor device context JSON. Reports any mismatches that could
cause detection.

Use --board to match exact build conditions (link speed/width clamping).
Without --board, validation uses no board constraints.

Example:
  pcileechgen validate --json device_context.json --output-dir pcileech_datastore/
  pcileechgen validate --json device_context.json --board PCIeSquirrel`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, err := donor.LoadContext(validateJSONPath)
		if err != nil {
			return fmt.Errorf("failed to load device context: %w", err)
		}
		fmt.Printf("Loaded donor context: %s\n\n",
			color.Bold(fmt.Sprintf("%04x:%04x (rev %02x)", ctx.Device.VendorID, ctx.Device.DeviceID, ctx.Device.RevisionID)))

		var b *board.Board
		if validateBoard != "" {
			b, err = board.Find(validateBoard)
			if err != nil {
				return err
			}
			fmt.Printf("Board: %s (%s x%d)\n\n", b.Name, b.FPGAPart, b.PCIeLanes)
		} else {
			fmt.Println(color.Warn("No --board specified: link speed/width clamping not applied in validation"))
			fmt.Println()
		}

		v := &validator{
			ctx:       ctx,
			b:         b,
			outputDir: validateOutputDir,
			scrubbed:  scrub.ScrubConfigSpace(ctx.ConfigSpace, b),
			result:    &output.ValidationResult{},
		}

		v.validateCOEFiles()
		v.validateIdentity()
		v.validateOutputFiles()
		v.validateSVIDs()
		v.validateFormats()

		// Print results
		for _, p := range v.result.Passed {
			fmt.Println(color.OK(p))
		}
		for _, w := range v.result.Warnings {
			fmt.Println(color.Warn(w))
		}
		for _, f := range v.result.Failed {
			fmt.Println(color.Fail(f))
		}

		fmt.Printf("\n%s\n", color.Header(v.result.Summary()))
		if v.result.HasFailures() {
			return fmt.Errorf("%d validation(s) failed", len(v.result.Failed))
		}
		return nil
	},
}

// validateCOEFiles checks config space and writemask COE against expected output.
func (v *validator) validateCOEFiles() {
	coePath := filepath.Join(v.outputDir, "pcileech_cfgspace.coe")
	if _, err := os.Stat(coePath); err == nil {
		coeData, err := os.ReadFile(coePath)
		if err != nil {
			v.result.Failed = append(v.result.Failed, fmt.Sprintf("pcileech_cfgspace.coe read error: %v", err))
			return
		}

		expectedCOE := codegen.GenerateConfigSpaceCOE(v.scrubbed)
		if string(coeData) == expectedCOE {
			v.result.Passed = append(v.result.Passed, "pcileech_cfgspace.coe matches donor config space (scrubbed)")
		} else {
			v.result.Failed = append(v.result.Failed, "pcileech_cfgspace.coe MISMATCH")
		}
	} else {
		v.result.Failed = append(v.result.Failed, "pcileech_cfgspace.coe not found")
	}

	wmPath := filepath.Join(v.outputDir, "pcileech_cfgspace_writemask.coe")
	if _, err := os.Stat(wmPath); err == nil {
		wmData, err := os.ReadFile(wmPath)
		if err != nil {
			v.result.Failed = append(v.result.Failed, fmt.Sprintf("pcileech_cfgspace_writemask.coe read error: %v", err))
			return
		}

		expectedWM := codegen.GenerateWritemaskCOE(v.scrubbed)
		if string(wmData) == expectedWM {
			v.result.Passed = append(v.result.Passed, "pcileech_cfgspace_writemask.coe matches expected writemask")
		} else {
			v.result.Failed = append(v.result.Failed, "pcileech_cfgspace_writemask.coe MISMATCH")
		}
	} else {
		v.result.Failed = append(v.result.Failed, "pcileech_cfgspace_writemask.coe not found")
	}
}

// validateIdentity checks that critical IDs, DSN, and ext caps are correct.
func (v *validator) validateIdentity() {
	if v.ctx.ConfigSpace == nil {
		return
	}

	coePath := filepath.Join(v.outputDir, "pcileech_cfgspace.coe")
	if coeData, err := os.ReadFile(coePath); err == nil {
		coeStr := string(coeData)

		expectedWord0 := fmt.Sprintf("%08x", v.scrubbed.ReadU32(0))
		if strings.Contains(coeStr, expectedWord0) {
			v.result.Passed = append(v.result.Passed,
				fmt.Sprintf("VendorID:DeviceID = %04X:%04X present in COE", v.ctx.Device.VendorID, v.ctx.Device.DeviceID))
		} else {
			v.result.Failed = append(v.result.Failed,
				fmt.Sprintf("VendorID:DeviceID = %04X:%04X NOT in COE", v.ctx.Device.VendorID, v.ctx.Device.DeviceID))
		}

		expectedSubsys := fmt.Sprintf("%08x", v.scrubbed.ReadU32(0x2C))
		if strings.Contains(coeStr, expectedSubsys) {
			v.result.Passed = append(v.result.Passed,
				fmt.Sprintf("SubsysVendorID:SubsysDeviceID = %04X:%04X present in COE", v.ctx.Device.SubsysVendorID, v.ctx.Device.SubsysDeviceID))
		} else {
			v.result.Failed = append(v.result.Failed, "Subsystem IDs NOT in COE")
		}
	}

	ids := firmware.ExtractDeviceIDs(v.ctx.ConfigSpace, v.ctx.ExtCapabilities)
	if ids.HasDSN {
		v.result.Passed = append(v.result.Passed,
			fmt.Sprintf("DSN = 0x%s (will be patched into SV)", firmware.DSNToSVHex(ids.DSN)))
	} else {
		v.result.Warnings = append(v.result.Warnings, "No DSN found in donor (serial number emulation disabled)")
	}

	if v.ctx.ConfigSpace.Size >= pci.ConfigSpaceSize {
		extCaps := pci.ParseExtCapabilities(v.ctx.ConfigSpace)
		v.result.Passed = append(v.result.Passed,
			fmt.Sprintf("Extended config space: %d capabilities covered", len(extCaps)))
	} else {
		v.result.Warnings = append(v.result.Warnings,
			fmt.Sprintf("Only legacy config space (%d bytes) -- extended caps not populated", v.ctx.ConfigSpace.Size))
	}
}

// validateOutputFiles checks that all expected output files exist.
func (v *validator) validateOutputFiles() {
	vr := output.ValidateOutputDir(v.outputDir)
	v.result.Passed = append(v.result.Passed, vr.Passed...)
	v.result.Failed = append(v.result.Failed, vr.Failed...)
	v.result.Warnings = append(v.result.Warnings, vr.Warnings...)
}

// validateSVIDs checks device_config.sv and config_space_init.hex.
func (v *validator) validateSVIDs() {
	devCfgPath := filepath.Join(v.outputDir, "device_config.sv")
	if svData, err := os.ReadFile(devCfgPath); err == nil {
		ids := firmware.ExtractDeviceIDs(v.ctx.ConfigSpace, v.ctx.ExtCapabilities)
		if issues := output.ValidateSVIDs(string(svData), ids); len(issues) > 0 {
			v.result.Failed = append(v.result.Failed, issues...)
		} else {
			v.result.Passed = append(v.result.Passed, "device_config.sv contains correct VendorID and DeviceID")
		}
	}

	hexPath := filepath.Join(v.outputDir, "config_space_init.hex")
	if hexData, err := os.ReadFile(hexPath); err == nil {
		if err := output.ValidateHexFile(string(hexData), 1024); err != nil {
			v.result.Failed = append(v.result.Failed, fmt.Sprintf("config_space_init.hex: %v", err))
		} else {
			v.result.Passed = append(v.result.Passed, "config_space_init.hex: format valid (1024 words)")
		}
	}
}

// validateFormats checks COE file format validity.
func (v *validator) validateFormats() {
	for _, coeName := range []string{"pcileech_cfgspace.coe", "pcileech_cfgspace_writemask.coe", "pcileech_bar_zero4k.coe"} {
		cPath := filepath.Join(v.outputDir, coeName)
		if coeData, err := os.ReadFile(cPath); err == nil {
			if err := output.ValidateCOEFile(string(coeData)); err != nil {
				v.result.Failed = append(v.result.Failed, fmt.Sprintf("%s: %v", coeName, err))
			} else {
				v.result.Passed = append(v.result.Passed, fmt.Sprintf("%s: format valid", coeName))
			}
		}
	}
}

func init() {
	validateCmd.Flags().StringVar(&validateJSONPath, "json", "", "path to device_context.json (required)")
	validateCmd.Flags().StringVar(&validateOutputDir, "output-dir", ".", "path to firmware output directory")
	validateCmd.Flags().StringVar(&validateBoard, "board", "", "target FPGA board (for exact build-matching validation)")
	_ = validateCmd.MarkFlagRequired("json")
	rootCmd.AddCommand(validateCmd)
}
