package main

import (
	"fmt"

	"github.com/sercanarga/pcileechgen/internal/color"
	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/donor/vfio"
	"github.com/sercanarga/pcileechgen/internal/pci"
	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan and list available PCI devices",
	Long:  "Scans /sys/bus/pci/devices/ and lists all PCI devices with their VFIO compatibility status.",
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

		db := pci.LoadPCIDB()

		for _, dev := range devices {
			devName := db.DeviceName(dev.VendorID, dev.DeviceID)
			vendorName := db.VendorName(dev.VendorID)
			description := dev.ClassDescription()
			if vendorName != "" && devName != "" {
				description = fmt.Sprintf("%s %s", vendorName, devName)
			} else if vendorName != "" {
				description = vendorName
			}

			status := vfio.QuickStatus(dev.BDF.String())
			vfioTag := ""
			switch status {
			case "ready":
				vfioTag = " " + color.OK("vfio")
			case "ok":
				// nothing
			case "no-iommu":
				vfioTag = " " + color.Warn("no-iommu")
			default:
				vfioTag = " " + color.Warn(status)
			}

			fmt.Printf("%s %s [%04x]: %s [%04x:%04x]%s\n",
				dev.BDF.String(),
				dev.ClassDescription(),
				dev.ClassCode>>8,
				description,
				dev.VendorID,
				dev.DeviceID,
				vfioTag,
			)
		}

		fmt.Printf("\nTotal: %d devices\n", len(devices))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
}
