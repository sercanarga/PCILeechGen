package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/color"
	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/firmware/devclass"
	"github.com/spf13/cobra"
)

var diagnoseOpts struct {
	jsonPath  string
	outputDir string
	boardName string
}

var diagnoseCmd = &cobra.Command{
	Use:   "diagnose",
	Short: "Diagnose build, BAR/MSI-X, and Windows Code 10 risks",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, err := donor.LoadContext(diagnoseOpts.jsonPath)
		if err != nil {
			return fmt.Errorf("load donor context: %w", err)
		}
		var b *board.Board
		if diagnoseOpts.boardName != "" {
			b, err = board.Find(diagnoseOpts.boardName)
			if err != nil {
				return err
			}
		}
		reportDiagnostics(ctx, b, diagnoseOpts.outputDir)
		return nil
	},
}

func reportDiagnostics(ctx *donor.DeviceContext, b *board.Board, outputDir string) {
	var pass, warn, fail []string
	strategy := devclass.StrategyForClassAndVendor(ctx.Device.ClassCode, ctx.Device.VendorID)
	profile := strategy.Profile()

	fmt.Printf("Device %04x:%04x class=0x%06x %s\n",
		ctx.Device.VendorID, ctx.Device.DeviceID, ctx.Device.ClassCode, ctx.Device.ClassDescription())
	if b != nil {
		fmt.Printf("Board %s %s x%d\n", b.Name, b.FPGAPart, b.PCIeLanes)
	}

	if ctx.ConfigSpace == nil || ctx.ConfigSpace.Size < 256 {
		fail = append(fail, "config space is missing or shorter than 256 bytes")
	} else {
		pass = append(pass, fmt.Sprintf("config space present (%d bytes)", ctx.ConfigSpace.Size))
	}
	if len(ctx.BARs) == 0 {
		fail = append(fail, "no BARs found in donor context")
	}
	for _, bar := range ctx.BARs {
		if bar.Size == 0 {
			fail = append(fail, fmt.Sprintf("BAR%d has zero size", bar.Index))
		}
		if profile != nil && bar.Index == profile.PreferredBAR && int(bar.Size) < profile.MinBARSize {
			fail = append(fail, fmt.Sprintf("BAR%d size %d is below %s minimum %d",
				bar.Index, bar.Size, profile.ClassName, profile.MinBARSize))
		}
	}
	if profile != nil {
		pass = append(pass, fmt.Sprintf("device profile=%s preferred_bar=%d min_bar=%d",
			profile.ClassName, profile.PreferredBAR, profile.MinBARSize))
	}
	checkMSIX(ctx, &pass, &warn, &fail)
	checkClassSpecific(ctx, &pass, &warn, &fail)
	checkOutputDiagnostics(outputDir, &pass, &warn, &fail)

	for _, item := range pass {
		fmt.Println(color.OK(item))
	}
	for _, item := range warn {
		fmt.Println(color.Warn(item))
	}
	for _, item := range fail {
		fmt.Println(color.Fail(item))
	}
	fmt.Printf("\n%s\n", color.Header(fmt.Sprintf("%d passed, %d warnings, %d failed", len(pass), len(warn), len(fail))))
	if len(fail) > 0 {
		fmt.Println(color.Fail("Code 10 risk is high until failed checks are fixed."))
	} else if len(warn) > 0 {
		fmt.Println(color.Warn("No hard blocker found, but warnings can still cause driver-specific failures."))
	} else {
		fmt.Println(color.OK("No obvious Code 10 blocker found in static artifacts."))
	}
}

func checkMSIX(ctx *donor.DeviceContext, pass, warn, fail *[]string) {
	if ctx.MSIXData == nil {
		if ctx.Device.BaseClass() == 0x01 || ctx.Device.BaseClass() == 0x02 {
			*warn = append(*warn, "MSI-X data missing; Windows storage/NIC drivers often expect MSI-X")
		}
		return
	}
	if ctx.MSIXData.TableSize <= 0 {
		*fail = append(*fail, "MSI-X table has no vectors")
		return
	}
	*pass = append(*pass, fmt.Sprintf("MSI-X table present (%d vectors)", ctx.MSIXData.TableSize))
	if !barContains(ctx, ctx.MSIXData.TableBIR, ctx.MSIXData.TableOffset) {
		*fail = append(*fail, fmt.Sprintf("MSI-X table offset 0x%x is outside BAR%d", ctx.MSIXData.TableOffset, ctx.MSIXData.TableBIR))
	}
	if !barContains(ctx, ctx.MSIXData.PBABIR, ctx.MSIXData.PBAOffset) {
		*fail = append(*fail, fmt.Sprintf("MSI-X PBA offset 0x%x is outside BAR%d", ctx.MSIXData.PBAOffset, ctx.MSIXData.PBABIR))
	}
}

func checkClassSpecific(ctx *donor.DeviceContext, pass, warn, fail *[]string) {
	switch devclass.StrategyForClassAndVendor(ctx.Device.ClassCode, ctx.Device.VendorID).DeviceClass() {
	case devclass.ClassNVMe:
		if ctx.Device.ClassCode != devclass.ClassCodeNVMe {
			*fail = append(*fail, fmt.Sprintf("NVMe class should be 0x010802, got 0x%06x", ctx.Device.ClassCode))
		} else {
			*pass = append(*pass, "NVMe class code is 0x010802")
		}
		if ctx.NVMeIdentity == nil {
			*warn = append(*warn, "NVMe identity strings missing; generated model/serial defaults will be used")
		} else {
			*pass = append(*pass, "NVMe identity strings present")
		}
	case devclass.ClassEthernet:
		if ctx.Device.BaseClass() == 0x02 && ctx.Device.SubClass() == 0x00 {
			*pass = append(*pass, "Ethernet class code is valid")
		}
		if ctx.MSIXData != nil && ctx.MSIXData.TableSize < 3 {
			*warn = append(*warn, "Ethernet MSI-X vector count is below profile preference of 3")
		}
	}
}

func checkOutputDiagnostics(outputDir string, pass, warn, fail *[]string) {
	if outputDir == "" {
		return
	}
	for _, name := range []string{"build_summary.txt", "vivado_generate_project_summary.txt", "vivado_build_summary.txt"} {
		path := filepath.Join(outputDir, name)
		data, err := os.ReadFile(path)
		if err != nil {
			continue
		}
		*pass = append(*pass, fmt.Sprintf("%s found", name))
		for _, line := range strings.Split(string(data), "\n") {
			if strings.HasPrefix(line, "diagnosis=") {
				*warn = append(*warn, strings.TrimPrefix(line, "diagnosis="))
			}
		}
	}
}

func barContains(ctx *donor.DeviceContext, index int, offset uint32) bool {
	for _, bar := range ctx.BARs {
		if bar.Index == index {
			return uint64(offset) < bar.Size
		}
	}
	return false
}

func init() {
	diagnoseCmd.Flags().StringVar(&diagnoseOpts.jsonPath, "json", "", "path to device_context.json")
	diagnoseCmd.Flags().StringVar(&diagnoseOpts.outputDir, "output-dir", "", "optional build output directory")
	diagnoseCmd.Flags().StringVar(&diagnoseOpts.boardName, "board", "", "optional target board name")
	_ = diagnoseCmd.MarkFlagRequired("json")
	rootCmd.AddCommand(diagnoseCmd)
}
