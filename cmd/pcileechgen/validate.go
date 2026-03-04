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
		// Load donor context
		ctx, err := donor.LoadContext(validateJSONPath)
		if err != nil {
			return fmt.Errorf("failed to load device context: %w", err)
		}
		fmt.Printf("Loaded donor context: %s\n\n",
			color.Bold(fmt.Sprintf("%04x:%04x (rev %02x)", ctx.Device.VendorID, ctx.Device.DeviceID, ctx.Device.RevisionID)))

		// Resolve board (optional)
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

		passed := 0
		failed := 0

		// Scrub once — reuse for all checks
		scrubbedCS := scrub.ScrubConfigSpace(ctx.ConfigSpace, b)

		// Validate config space COE
		coePath := filepath.Join(validateOutputDir, "pcileech_cfgspace.coe")
		if _, err := os.Stat(coePath); err == nil {
			coeData, err := os.ReadFile(coePath)
			if err != nil {
				return fmt.Errorf("failed to read COE file: %w", err)
			}

			expectedCOE := codegen.GenerateConfigSpaceCOE(scrubbedCS)

			if string(coeData) == expectedCOE {
				fmt.Println(color.OK("pcileech_cfgspace.coe matches donor config space (scrubbed)"))
				passed++
			} else {
				fmt.Println(color.Fail("pcileech_cfgspace.coe MISMATCH"))
				reportCOEDiff(string(coeData), expectedCOE)
				failed++
			}
		} else {
			fmt.Println(color.Warn("pcileech_cfgspace.coe not found"))
			failed++
		}

		// Validate writemask COE
		wmPath := filepath.Join(validateOutputDir, "pcileech_cfgspace_writemask.coe")
		if _, err := os.Stat(wmPath); err == nil {
			wmData, err := os.ReadFile(wmPath)
			if err != nil {
				return fmt.Errorf("failed to read writemask COE: %w", err)
			}

			expectedWM := codegen.GenerateWritemaskCOE(scrubbedCS)

			if string(wmData) == expectedWM {
				fmt.Println(color.OK("pcileech_cfgspace_writemask.coe matches expected writemask"))
				passed++
			} else {
				fmt.Println(color.Fail("pcileech_cfgspace_writemask.coe MISMATCH"))
				failed++
			}
		} else {
			fmt.Println(color.Warn("pcileech_cfgspace_writemask.coe not found"))
			failed++
		}

		// Validate critical fields in COE match donor identity
		if ctx.ConfigSpace != nil {
			fmt.Printf("\n%s\n", color.Header("Identity Verification"))

			coePath := filepath.Join(validateOutputDir, "pcileech_cfgspace.coe")
			if coeData, err := os.ReadFile(coePath); err == nil {
				coeStr := string(coeData)

				// Extract first word (VendorID:DeviceID)
				expectedWord0 := fmt.Sprintf("%08x", scrubbedCS.ReadU32(0))
				if strings.Contains(coeStr, expectedWord0) {
					fmt.Println(color.Okf("VendorID:DeviceID = %04X:%04X present in COE",
						ctx.Device.VendorID, ctx.Device.DeviceID))
					passed++
				} else {
					fmt.Println(color.Failf("VendorID:DeviceID = %04X:%04X NOT in COE",
						ctx.Device.VendorID, ctx.Device.DeviceID))
					failed++
				}

				// Check subsystem IDs (offset 0x2C)
				expectedSubsys := fmt.Sprintf("%08x", scrubbedCS.ReadU32(0x2C))
				if strings.Contains(coeStr, expectedSubsys) {
					fmt.Println(color.Okf("SubsysVendorID:SubsysDeviceID = %04X:%04X present in COE",
						ctx.Device.SubsysVendorID, ctx.Device.SubsysDeviceID))
					passed++
				} else {
					fmt.Println(color.Fail("Subsystem IDs NOT in COE"))
					failed++
				}
			}

			// Check DSN
			ids := firmware.ExtractDeviceIDs(ctx.ConfigSpace, ctx.ExtCapabilities)
			if ids.HasDSN {
				fmt.Println(color.Okf("DSN = 0x%s (will be patched into SV)", firmware.DSNToSVHex(ids.DSN)))
				passed++
			} else {
				fmt.Println(color.Warn("No DSN found in donor (serial number emulation disabled)"))
			}

			// Check extended config space coverage
			if ctx.ConfigSpace.Size >= pci.ConfigSpaceSize {
				extCaps := pci.ParseExtCapabilities(ctx.ConfigSpace)
				fmt.Println(color.Okf("Extended config space: %d capabilities covered", len(extCaps)))
				passed++
			} else {
				fmt.Println(color.Warnf("Only legacy config space (%d bytes) -- extended caps not populated", ctx.ConfigSpace.Size))
			}
		}

		// Validate all output files exist
		fmt.Printf("\n%s\n", color.Header("Output File Check"))
		vr := output.ValidateOutputDir(validateOutputDir)
		for _, p := range vr.Passed {
			fmt.Println(color.OK(p))
			passed++
		}
		for _, f := range vr.Failed {
			fmt.Println(color.Fail(f))
			failed++
		}

		// Validate device_config.sv has correct IDs
		devCfgPath := filepath.Join(validateOutputDir, "device_config.sv")
		if svData, err := os.ReadFile(devCfgPath); err == nil {
			ids := firmware.ExtractDeviceIDs(ctx.ConfigSpace, ctx.ExtCapabilities)
			if issues := output.ValidateSVIDs(string(svData), ids); len(issues) > 0 {
				for _, issue := range issues {
					fmt.Println(color.Fail(issue))
					failed++
				}
			} else {
				fmt.Println(color.OK("device_config.sv contains correct VendorID and DeviceID"))
				passed++
			}
		}

		// Validate config_space_init.hex format
		hexPath := filepath.Join(validateOutputDir, "config_space_init.hex")
		if hexData, err := os.ReadFile(hexPath); err == nil {
			if err := output.ValidateHexFile(string(hexData), 1024); err != nil {
				fmt.Println(color.Failf("config_space_init.hex: %v", err))
				failed++
			} else {
				fmt.Println(color.OK("config_space_init.hex: format valid (1024 words)"))
				passed++
			}
		}

		// Validate COE file format
		for _, coeName := range []string{"pcileech_cfgspace.coe", "pcileech_cfgspace_writemask.coe", "pcileech_bar_zero4k.coe"} {
			cPath := filepath.Join(validateOutputDir, coeName)
			if coeData, err := os.ReadFile(cPath); err == nil {
				if err := output.ValidateCOEFile(string(coeData)); err != nil {
					fmt.Println(color.Failf("%s: %v", coeName, err))
					failed++
				} else {
					fmt.Println(color.Okf("%s: format valid", coeName))
					passed++
				}
			}
		}

		fmt.Printf("\n%s\n", color.Header(fmt.Sprintf("Validation complete: %d passed, %d failed", passed, failed)))
		if failed > 0 {
			return fmt.Errorf("%d validation(s) failed", failed)
		}
		return nil
	},
}

// reportCOEDiff reports first differing line between two COE files.
func reportCOEDiff(got, expected string) {
	gotLines := strings.Split(got, "\n")
	expLines := strings.Split(expected, "\n")

	diffCount := 0
	maxDiffs := 5
	for i := 0; i < len(gotLines) && i < len(expLines); i++ {
		if gotLines[i] != expLines[i] {
			if diffCount < maxDiffs {
				fmt.Printf("  line %d: got=%q expected=%q\n", i+1, gotLines[i], expLines[i])
			}
			diffCount++
		}
	}
	if diffCount > maxDiffs {
		fmt.Printf("  ... and %d more differences\n", diffCount-maxDiffs)
	}
}

func init() {
	validateCmd.Flags().StringVar(&validateJSONPath, "json", "", "path to device_context.json (required)")
	validateCmd.Flags().StringVar(&validateOutputDir, "output-dir", ".", "path to firmware output directory")
	validateCmd.Flags().StringVar(&validateBoard, "board", "", "target FPGA board (for exact build-matching validation)")
	_ = validateCmd.MarkFlagRequired("json")
	rootCmd.AddCommand(validateCmd)
}
