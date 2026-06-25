package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/sercanarga/pcileechgen/internal/color"
	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/pci"
	"github.com/spf13/cobra"
)

type contextDiffOptions struct {
	leftPath  string
	rightPath string
	jsonOutput bool
}

type contextDiffReport struct {
	Left       string   `json:"left_path"`
	Right      string   `json:"right_path"`
	Differences []string `json:"differences"`
	Matches    []string `json:"matches"`
	Equal      bool     `json:"equal"`
}

var contextDiffOpts contextDiffOptions

var contextDiffCmd = &cobra.Command{
	Use:   "context-diff",
	Short: "Compare two saved donor context JSON files",
	Long: `Compares identity, BAR topology, capabilities, and donor content metadata
between two donor snapshots.

Example:
  pcileechgen context-diff --left before/context.json --right after/context.json
  pcileechgen context-diff --left a.json --right b.json --json`,
	RunE: func(cmd *cobra.Command, args []string) error {
		leftCtx, err := donor.LoadContext(contextDiffOpts.leftPath)
		if err != nil {
			return fmt.Errorf("load left context %q: %w", contextDiffOpts.leftPath, err)
		}
		rightCtx, err := donor.LoadContext(contextDiffOpts.rightPath)
		if err != nil {
			return fmt.Errorf("load right context %q: %w", contextDiffOpts.rightPath, err)
		}

		report := compareContextDiff(leftCtx, rightCtx)
		report.Left = contextDiffOpts.leftPath
		report.Right = contextDiffOpts.rightPath

		if contextDiffOpts.jsonOutput {
			enc := json.NewEncoder(cmd.OutOrStdout())
			enc.SetIndent("", "  ")
			if err := enc.Encode(report); err != nil {
				return fmt.Errorf("render JSON report: %w", err)
			}
			return nil
		}

		printContextDiffReport(cmd.OutOrStdout(), report)
		return nil
	},
}

func compareContextDiff(left, right *donor.DeviceContext) *contextDiffReport {
	report := &contextDiffReport{Equal: true}

	if left == nil && right == nil {
		report.Differences = append(report.Differences, "both contexts are nil")
		report.Equal = false
		return report
	}
	if left == nil {
		report.Differences = append(report.Differences, "left context is nil")
		report.Equal = false
		return report
	}
	if right == nil {
		report.Differences = append(report.Differences, "right context is nil")
		report.Equal = false
		return report
	}

	if left.Device.VendorID != right.Device.VendorID {
		report.addDiff(fmt.Sprintf("VendorID: %04x vs %04x", left.Device.VendorID, right.Device.VendorID))
	} else {
		report.addMatch("VendorID matches")
	}
	if left.Device.DeviceID != right.Device.DeviceID {
		report.addDiff(fmt.Sprintf("DeviceID: %04x vs %04x", left.Device.DeviceID, right.DeviceID))
	} else {
		report.addMatch("DeviceID matches")
	}
	if left.Device.RevisionID != right.Device.RevisionID {
		report.addDiff(fmt.Sprintf("RevisionID: %02x vs %02x", left.Device.RevisionID, right.Device.RevisionID))
	} else {
		report.addMatch("RevisionID matches")
	}
	if left.Device.ClassCode != right.Device.ClassCode {
		report.addDiff(fmt.Sprintf("ClassCode: %06x vs %06x", left.Device.ClassCode, right.Device.ClassCode))
	} else {
		report.addMatch("ClassCode matches")
	}

	if left.Device.SubsysVendorID != right.Device.SubsysVendorID {
		report.addDiff(fmt.Sprintf("Subsystem vendor: %04x vs %04x", left.Device.SubsysVendorID, right.Device.SubsysVendorID))
	} else {
		report.addMatch("Subsystem vendor matches")
	}
	if left.Device.SubsysDeviceID != right.Device.SubsysDeviceID {
		report.addDiff(fmt.Sprintf("Subsystem device: %04x vs %04x", left.Device.SubsysDeviceID, right.Device.SubsysDeviceID))
	} else {
		report.addMatch("Subsystem device matches")
	}

	leftSize := len(left.BARs)
	rightSize := len(right.BARs)
	if leftSize != rightSize {
		report.addDiff(fmt.Sprintf("BAR count: %d vs %d", leftSize, rightSize))
	} else {
		report.addMatch(fmt.Sprintf("BAR count: %d", leftSize))
	}

	leftBars := barIndexMap(left.BARs)
	rightBars := barIndexMap(right.BARs)
	indices := mergeIndexes(leftBars, rightBars)
	for _, idx := range indices {
		lbar, lfound := leftBars[idx]
		rbar, rfound := rightBars[idx]
		if !lfound {
			report.addDiff(fmt.Sprintf("BAR%d missing in left", idx))
			continue
		}
		if !rfound {
			report.addDiff(fmt.Sprintf("BAR%d missing in right", idx))
			continue
		}
		if !barEquals(lbar, rbar) {
			report.addDiff(fmt.Sprintf("BAR%d mismatch: %s vs %s", idx, lbar.String(), rbar.String()))
		} else {
			report.addMatch(fmt.Sprintf("BAR%d matches (%s)", idx, lbar.String()))
		}

		lBytes := len(left.BARContents[idx])
		rBytes := len(right.BARContents[idx])
		if lBytes != rBytes {
			report.addDiff(fmt.Sprintf("BAR%d content bytes: %d vs %d", idx, lBytes, rBytes))
		} else {
			report.addMatch(fmt.Sprintf("BAR%d content bytes match (%d)", idx, lBytes))
		}
	}

	if left.ConfigSpace == nil && right.ConfigSpace == nil {
		report.addMatch("ConfigSpace presence matches")
	} else if left.ConfigSpace == nil || right.ConfigSpace == nil {
		report.addDiff("ConfigSpace presence mismatch")
	} else if left.ConfigSpace.Size != right.ConfigSpace.Size {
		report.addDiff(fmt.Sprintf("ConfigSpace size: %d vs %d", left.ConfigSpace.Size, right.ConfigSpace.Size))
	} else {
		report.addMatch(fmt.Sprintf("ConfigSpace size: %d", left.ConfigSpace.Size))
	}

	if len(left.Capabilities) != len(right.Capabilities) {
		report.addDiff(fmt.Sprintf("Capabilities: %d vs %d", len(left.Capabilities), len(right.Capabilities)))
	} else {
		report.addMatch(fmt.Sprintf("Capabilities count: %d", len(left.Capabilities)))
	}
	if len(left.ExtCapabilities) != len(right.ExtCapabilities) {
		report.addDiff(fmt.Sprintf("Extended capabilities: %d vs %d", len(left.ExtCapabilities), len(right.ExtCapabilities)))
	} else {
		report.addMatch(fmt.Sprintf("Extended capabilities count: %d", len(left.ExtCapabilities)))
	}

	if left.MSIXData == nil && right.MSIXData == nil {
		report.addMatch("MSI-X metadata: both absent")
	} else if left.MSIXData == nil || right.MSIXData == nil {
		report.addDiff("MSI-X metadata presence mismatch")
	} else if left.MSIXData.TableSize != right.MSIXData.TableSize {
		report.addDiff(fmt.Sprintf("MSI-X table size: %d vs %d", left.MSIXData.TableSize, right.MSIXData.TableSize))
	} else {
		report.addMatch(fmt.Sprintf("MSI-X table size: %d", left.MSIXData.TableSize))
	}

	report.Equal = len(report.Differences) == 0
	return report
}

func (r *contextDiffReport) addDiff(msg string) {
	r.Differences = append(r.Differences, msg)
	r.Equal = false
}

func (r *contextDiffReport) addMatch(msg string) {
	r.Matches = append(r.Matches, msg)
}

func barIndexMap(bars []pci.BAR) map[int]pci.BAR {
	byIndex := make(map[int]pci.BAR, len(bars))
	for _, bar := range bars {
		byIndex[bar.Index] = bar
	}
	return byIndex
}

func mergeIndexes(a, b map[int]pci.BAR) []int {
	seen := make(map[int]struct{})
	for idx := range a {
		seen[idx] = struct{}{}
	}
	for idx := range b {
		seen[idx] = struct{}{}
	}

	var indices []int
	for idx := range seen {
		indices = append(indices, idx)
	}
	sort.Ints(indices)
	return indices
}

func printContextDiffReport(out io.Writer, report *contextDiffReport) {
	if out == nil || report == nil {
		return
	}

	fmt.Fprintln(out, color.Header("Context Diff"))
	fmt.Fprintf(out, "%s\n", strings.Repeat("-", 72))
	fmt.Fprintf(out, "left:  %s\n", report.Left)
	fmt.Fprintf(out, "right: %s\n", report.Right)
	fprintf := fmt.Fprintf
	fprintf(out, "result: %s\n\n", map[bool]string{true: color.OK("equal"), false: color.Fail("different")}[report.Equal])

	if len(report.Matches) > 0 {
		fmt.Fprintln(out, color.OK("Matches:"))
		for _, line := range report.Matches {
			fmt.Fprintf(out, "  - %s\n", line)
		}
		fmt.Fprintln(out)
	}

	if len(report.Differences) > 0 {
		fmt.Fprintln(out, color.Warn("Differences:"))
		for _, line := range report.Differences {
			fmt.Fprintf(out, "  - %s\n", line)
		}
		fmt.Fprintln(out)
	}

	if report.Equal {
		fmt.Fprintln(out, color.OK("No significant differences found."))
	}
}

func barEquals(left, right pci.BAR) bool {
	return left.Index == right.Index &&
		left.RawValue == right.RawValue &&
		left.Address == right.Address &&
		left.Size == right.Size &&
		left.Type == right.Type &&
		left.Prefetchable == right.Prefetchable &&
		left.Is64Bit == right.Is64Bit
}

func init() {
	contextDiffCmd.Flags().StringVar(&contextDiffOpts.leftPath, "left", "", "left context JSON path (base line)")
	contextDiffCmd.Flags().StringVar(&contextDiffOpts.rightPath, "right", "", "right context JSON path (current)" )
	contextDiffCmd.Flags().BoolVar(&contextDiffOpts.jsonOutput, "json", false, "emit machine-readable diff")
	_ = contextDiffCmd.MarkFlagRequired("left")
	_ = contextDiffCmd.MarkFlagRequired("right")
	rootCmd.AddCommand(contextDiffCmd)
}

