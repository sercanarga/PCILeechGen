package main

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/sercanarga/pcileechgen/internal/color"
	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/pci"
	"github.com/spf13/cobra"
)

type donorAuditOptions struct {
	contextPath string
	jsonOutput  bool
}

type donorAuditReport struct {
	DeviceName          string   `json:"device_name"`
	ContextPath         string   `json:"context_path"`
	VendorID            uint16   `json:"vendor_id"`
	DeviceID            uint16   `json:"device_id"`
	Score               int      `json:"score"`
	MaxScore            int      `json:"max_score"`
	Status              string   `json:"status"`
	MemoryBars          int      `json:"memory_bars"`
	BarsWithContents    int      `json:"bars_with_contents"`
	BarsWithProfiles    int      `json:"bars_with_profiles"`
	BarsWithTraceModels int      `json:"bars_with_trace_models"`
	Blockers            []string `json:"blockers"`
	Warnings            []string `json:"warnings"`
	Highlights          []string `json:"highlights"`
}

var donorAuditOpts donorAuditOptions

var donorAuditCmd = &cobra.Command{
	Use:   "donor-audit",
	Short: "Assess donor context fidelity for emulation",
	Long: `Evaluates a saved device context for emulation risk and blockers.

Example:
  pcileechgen donor-audit --context device_context.json
  pcileechgen donor-audit --context device_context.json --json`,
	RunE: func(cmd *cobra.Command, args []string) error {
		report, err := runDonorAudit(donorAuditOpts)
		if err != nil {
			return err
		}
		if donorAuditOpts.jsonOutput {
			enc := json.NewEncoder(cmd.OutOrStdout())
			enc.SetIndent("", "  ")
			return enc.Encode(report)
		}
		printDonorAuditReport(cmd.OutOrStdout(), report)
		return nil
	},
}

func assessDonorContext(ctx *donor.DeviceContext, contextPath string) donorAuditReport {
	report := donorAuditReport{
		DeviceName:  ctx.Device.Summary(),
		ContextPath: contextPath,
		VendorID:    ctx.Device.VendorID,
		DeviceID:    ctx.Device.DeviceID,
		Score:       100,
		MaxScore:    100,
		Status:      "READY",
		Highlights:  []string{},
	}

	if ctx.ConfigSpace == nil {
		report.Score = 0
		report.Status = "BLOCKED"
		report.Blockers = append(report.Blockers, "Missing config space (cannot evaluate identity, BAR aperture, or class capabilities)")
		return report
	}

	if ctx.ConfigSpace.Size < pci.ConfigSpaceLegacySize {
		report.Score -= 18
		report.Warnings = append(report.Warnings, "Config-space payload is smaller than legacy config space")
	}

	if ctx.ConfigSpace.Size < pci.ConfigSpaceSize {
		report.Score -= 10
		report.Warnings = append(report.Warnings, "No extended config-space capture: advanced PCIe features cannot be reconstructed")
	}

	if len(ctx.BARs) == 0 {
		report.Score -= 35
		report.Status = "BLOCKED"
		report.Blockers = append(report.Blockers, "No BAR metadata available")
		return report
	}

	for _, bar := range ctx.BARs {
		if !bar.IsMemory() || bar.IsDisabled() {
			continue
		}
		report.MemoryBars++

		if data, ok := ctx.BARContents[bar.Index]; !ok || len(data) == 0 {
			report.Score -= 7
			report.Warnings = append(report.Warnings,
				fmt.Sprintf("BAR%d memory image missing; static register emulation fallback will be used", bar.Index))
		} else {
			report.BarsWithContents++
			report.Highlights = append(report.Highlights,
				fmt.Sprintf("BAR%d has %d bytes of captured memory image", bar.Index, len(data)))
		}

		if profile, ok := ctx.BARProfiles[bar.Index]; !ok || profile == nil {
			report.Score -= 4
			report.Warnings = append(report.Warnings,
				fmt.Sprintf("BAR%d RW/RO probing not present", bar.Index))
		} else if len(profile.Probes) > 0 {
			report.BarsWithProfiles++
		}

		if trace, ok := ctx.MMIOTraces[bar.Index]; !ok || trace == nil || len(trace.Records) == 0 {
			report.Score -= 2
			report.Warnings = append(report.Warnings,
				fmt.Sprintf("BAR%d lacks MMIO trace evidence", bar.Index))
		} else {
			report.BarsWithTraceModels++
		}
	}

	if report.MemoryBars == 0 {
		report.Score -= 35
		report.Status = "BLOCKED"
		report.Blockers = append(report.Blockers, "No enabled memory BAR discovered")
	}

	if report.Score < 70 {
		report.Status = "AT_RISK"
	}
	if report.Score < 45 {
		report.Status = "BLOCKED"
	}

	if !hasDSN(ctx) {
		report.Score -= 9
		report.Warnings = append(report.Warnings, "No Device Serial Number capability: serial emulation cannot be exact")
	}

	if hasMSIX(ctx) && ctx.MSIXData == nil {
		report.Score -= 3
		report.Warnings = append(report.Warnings, "MSI-X capability is present but MSI-X table snapshot is missing")
	}

	if ctx.NVMeIdentify != nil && ctx.NVMeIdentify.HasController {
		report.Highlights = append(report.Highlights, "NVMe Identify controller payload captured")
	}
	if ctx.NVMeIdentify != nil && ctx.NVMeIdentify.HasNamespace {
		report.Highlights = append(report.Highlights, "NVMe Identify namespace payload captured")
	}

	if len(report.Blockers) > 0 {
		report.Status = "BLOCKED"
	}
	if report.Score < 0 {
		report.Score = 0
	}

	return report
}

func hasDSN(ctx *donor.DeviceContext) bool {
	for _, cap := range ctx.ExtCapabilities {
		if cap.ID == pci.ExtCapIDDeviceSerialNumber {
			return true
		}
	}
	return false
}

func hasMSIX(ctx *donor.DeviceContext) bool {
	for _, cap := range ctx.Capabilities {
		if cap.ID == pci.CapIDMSIX {
			return true
		}
	}
	return false
}

func runDonorAudit(opts donorAuditOptions) (*donorAuditReport, error) {
	if opts.contextPath == "" {
		return nil, fmt.Errorf("--context is required")
	}

	ctx, err := donor.LoadContext(opts.contextPath)
	if err != nil {
		return nil, fmt.Errorf("load donor context %q: %w", opts.contextPath, err)
	}
	report := assessDonorContext(ctx, opts.contextPath)
	return &report, nil
}

func printDonorAuditReport(out io.Writer, report *donorAuditReport) {
	if out == nil {
		return
	}
	_, _ = out.Write([]byte(color.Header("Donor emulation readiness")))
	_, _ = fmt.Fprintf(out, " for %s\n", report.DeviceName)
	_, _ = fmt.Fprintf(out, "Context: %s\n", report.ContextPath)
	_, _ = fmt.Fprintf(out, "Score: %d/%d\n", report.Score, report.MaxScore)
	_, _ = fmt.Fprintf(out, "Status: %s\n", report.Status)
	_, _ = fmt.Fprintf(out, "Memory BARs: %d (captured=%d, profiled=%d, traced=%d)\n",
		report.MemoryBars, report.BarsWithContents, report.BarsWithProfiles, report.BarsWithTraceModels)
	_, _ = fmt.Fprintln(out)

	if len(report.Blockers) > 0 {
		_, _ = fmt.Fprintln(out, color.Fail("Blockers:"))
		for _, b := range report.Blockers {
			_, _ = fmt.Fprintf(out, "  - %s\n", b)
		}
		_, _ = fmt.Fprintln(out)
	}

	if len(report.Warnings) > 0 {
		_, _ = fmt.Fprintln(out, color.Warn("Warnings:"))
		for _, w := range report.Warnings {
			_, _ = fmt.Fprintf(out, "  - %s\n", w)
		}
		_, _ = fmt.Fprintln(out)
	}

	if len(report.Highlights) > 0 {
		_, _ = fmt.Fprintln(out, color.OK("Highlights:"))
		for _, h := range report.Highlights {
			_, _ = fmt.Fprintf(out, "  - %s\n", h)
		}
	}
}

func init() {
	donorAuditCmd.Flags().StringVar(&donorAuditOpts.contextPath, "context", "", "path to donor device_context.json")
	donorAuditCmd.Flags().BoolVar(&donorAuditOpts.jsonOutput, "json", false, "emit JSON report")
	_ = donorAuditCmd.MarkFlagRequired("context")
	rootCmd.AddCommand(donorAuditCmd)
}
