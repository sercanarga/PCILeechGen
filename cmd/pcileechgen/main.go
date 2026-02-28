package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pcileechgen",
	Short: "PCILeech FPGA firmware generator",
	Long: `PCILeechGen generates custom PCILeech FPGA firmware from real donor PCI/PCIe devices.

It reads the donor device's configuration via VFIO/sysfs, generates firmware artifacts
(.coe, .sv, .tcl), and optionally builds the bitstream using Xilinx Vivado.

This tool requires:
  - Linux with IOMMU/VFIO support (for device reading)
  - A real donor PCI/PCIe card
  - Xilinx Vivado (optional, for bitstream synthesis)`,
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
