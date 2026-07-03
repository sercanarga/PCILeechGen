package output

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/firmware"
	"github.com/sercanarga/pcileechgen/internal/firmware/devclass"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

func (ow *OutputWriter) code10Risks(ctx *donor.DeviceContext, b *board.Board, ids firmware.DeviceIDs) []string {
	msixTableSize := 0
	if ctx.MSIXData != nil && ctx.MSIXData.TableSize > 0 {
		msixTableSize = ctx.MSIXData.TableSize
	}
	demand := firmware.DonorBAR0Demand(ctx, b, msixTableSize)
	capped := firmware.CappedBAR0Size(ctx, b, msixTableSize)

	devClass := ""
	if s := devclass.StrategyForClassAndVendor(ctx.Device.ClassCode, ids.VendorID); s != nil {
		devClass = s.DeviceClass()
	}

	var risks []string

	if capped < demand {
		risks = append(risks, fmt.Sprintf(
			"BAR0 TRUNCATED %d -> %d bytes (board BRAM limit): driver reads above 0x%X return garbage. Use a board with larger bram_size or a smaller-BAR donor.",
			demand, capped, capped))
	}

	switch devClass {
	case devclass.ClassNVMe:
		risks = append(risks, "NVMe donor: admin + IO-queue FSM with a small BRAM sector cache (4 sectors) is implemented - IO reads/writes round-trip through the cache. Not yet silicon-validated; cache is small so large transfers may wrap. stornvme Event 11 possible if the driver probes beyond the cache window.")
	case devclass.ClassXHCI:
		risks = append(risks, "xHCI (USB) donor: Command Ring + Event Ring engine is implemented (answers No-Op and Enable Slot commands with completion events). No downstream USB device or transfer-ring DMA yet - USBXHCI enumerates but cannot transfer data. Not yet silicon-validated.")
	case devclass.ClassSATA:
		risks = append(risks, "SATA/AHCI donor: no command-list FSM - PxCI never clears, so storahci commands never complete (Code 10). This class is an identity clone only.")
	case devclass.ClassWiFi:
		risks = append(risks, "WiFi donor: no uCode-upload DMA handshake FSM, and the BAR model may fall back to a generic NIC layout - driver init typically fails (Code 10).")
	case devclass.ClassGPU:
		risks = append(risks, "GPU donor: only a 4K register window is emulated of a multi-MB VRAM/register space - structurally insufficient for any GPU driver (Code 10).")
	case devclass.ClassThunderbolt:
		risks = append(risks, "Thunderbolt donor: no link/mailbox/tunnelling FSM - static snapshot only, driver init will not complete (Code 10).")
	case devclass.ClassEthernet:
		risks = append(risks, "Ethernet donor: static identity clone with link-up state but no TX/RX descriptor-ring DMA - enumerates but does not pass traffic.")
	}

	if !ow.StockBar && isInterruptDrivenClass(devClass) {
		risks = append(risks, "interrupt-driven class with custom SV: VERIFY intr_req is wired to the PCIe core cfg_interrupt in your pcileech-fpga top module - this tool does not emit that connection, and if it is missing, interrupts never fire (Code 10).")
	}

	if barData, ok := ctx.BARContents[firmware.LargestBarIndex(ctx.BARContents)]; ok {
		switch {
		case len(barData) == 0:
			risks = append(risks, "selected BAR has no captured content (donor likely in D3/IOMMU off): BAR0 will read as zeros, init will fail. Re-run 'check' and confirm D0 + vfio-pci binding.")
		case allBytes(barData, 0xFF):
			risks = append(risks, "selected BAR content is all 0xFF (donor not readable at capture time): the clone's BAR0 will read FF and the driver will Code 10. Wake the donor to D0 and re-collect.")
		}
	}

	if msixTableSize == 0 && !hasMSICap(ctx.ConfigSpace) {
		risks = append(risks, "no MSI-X table and no MSI capability detected on the donor: interrupt-driven drivers may never complete init (Code 10).")
	}

	if ctx.MSIXData != nil && (ctx.MSIXData.TableBIR != 0 || ctx.MSIXData.PBABIR != 0) {
		risks = append(risks, fmt.Sprintf(
			"MSI-X relocated to BAR0: donor BIR was table=%d pba=%d but only BAR0 is emulated, so the clone reports BIR 0. Required for the table to work, but a fidelity/detection divergence from the donor.",
			ctx.MSIXData.TableBIR, ctx.MSIXData.PBABIR))
	}

	if !ids.HasDSN {
		risks = append(risks, "donor has no DSN: serial-number emulation disabled (usually harmless).")
	}
	if ids.HasPCIeCap && b.PCIeLanes > 0 && int(ids.LinkWidth) > b.PCIeLanes {
		risks = append(risks, fmt.Sprintf("donor link width x%d clamped to board x%d: fine for most devices, but bandwidth-sensitive drivers may complain.", ids.LinkWidth, b.PCIeLanes))
	}

	return risks
}

func isInterruptDrivenClass(devClass string) bool {
	switch devClass {
	case devclass.ClassNVMe, devclass.ClassXHCI, devclass.ClassAudio, devclass.ClassEthernet:
		return true
	}
	return false
}

func (ow *OutputWriter) writeCode10Report(ctx *donor.DeviceContext, b *board.Board, ids firmware.DeviceIDs) {
	risks := ow.code10Risks(ctx, b, ids)

	header := fmt.Sprintf("Code-10 risk report for %04x:%04x (%s) on %s",
		ids.VendorID, ids.DeviceID, ctx.Device.ClassDescription(), b.Name)

	if len(risks) == 0 {
		slog.Info("code10-risk: none detected", "device", header)
		_ = ow.writeFile("code10_report.txt", header+"\n\nNo obvious Code-10 risks detected at generation time.\n")
		return
	}

	for _, r := range risks {
		slog.Warn("code10-risk", "detail", r)
	}

	var sb strings.Builder
	sb.WriteString(header + "\n\n")
	for i, r := range risks {
		fmt.Fprintf(&sb, "%d. %s\n", i+1, r)
	}
	sb.WriteString("\nIf the flash still Code-10s after addressing these, run tools/cleanup_device_history.ps1 (Admin) and reboot.\n")
	_ = ow.writeFile("code10_report.txt", sb.String())
}

func allBytes(b []byte, v byte) bool {
	for _, x := range b {
		if x != v {
			return false
		}
	}
	return len(b) > 0
}

func hasMSICap(cs *pci.ConfigSpace) bool {
	if cs == nil {
		return false
	}
	for _, c := range pci.ParseCapabilities(cs) {
		if c.ID == pci.CapIDMSI {
			return true
		}
	}
	return false
}
