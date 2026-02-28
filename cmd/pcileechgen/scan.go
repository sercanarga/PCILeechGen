package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan and list available PCI devices",
	Long:  "Scans /sys/bus/pci/devices/ and lists all PCI devices with their details.",
	RunE: func(cmd *cobra.Command, args []string) error {
		sr := donor.NewSysfsReader()
		devices, err := sr.ScanDevices()
		if err != nil {
			return fmt.Errorf("failed to scan devices: %w", err)
		}

		if len(devices) == 0 {
			fmt.Println("No PCI devices found.")
			return nil
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintln(w, "BDF\tVENDOR\tDEVICE\tCLASS\tDRIVER")
		fmt.Fprintln(w, "---\t------\t------\t-----\t------")

		for _, dev := range devices {
			fmt.Fprintf(w, "%s\t%04x\t%04x\t%s\t%s\n",
				dev.BDF.String(),
				dev.VendorID,
				dev.DeviceID,
				dev.ClassDescription(),
				dev.Driver,
			)
		}
		w.Flush()

		fmt.Printf("\nTotal: %d devices\n", len(devices))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
}
