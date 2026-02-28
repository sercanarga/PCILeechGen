package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"text/tabwriter"

	"github.com/sercanarga/pcileechgen/internal/donor"
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

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

		for _, dev := range devices {
			driver := dev.Driver
			if driver == "" {
				driver = "-"
			}

			// IOMMU group
			iommuStr := "-"
			iommuLink, err := os.Readlink(filepath.Join("/sys/bus/pci/devices", dev.BDF.String(), "iommu_group"))
			if err == nil {
				iommuStr = filepath.Base(iommuLink)
			}

			// VFIO status
			vfioStatus := vfioCheck(dev.Driver, iommuStr, dev.BDF.String())

			// Device description from pci.ids
			devName := db.DeviceName(dev.VendorID, dev.DeviceID)
			vendorName := db.VendorName(dev.VendorID)
			description := dev.ClassDescription()
			if vendorName != "" && devName != "" {
				description = fmt.Sprintf("%s %s", vendorName, devName)
			} else if vendorName != "" {
				description = fmt.Sprintf("%s [%04x:%04x]", vendorName, dev.VendorID, dev.DeviceID)
			}

			fmt.Fprintf(w, "%s %s [%04x]: %s [%04x:%04x] (%s)\t%s\tiommu=%s\tvfio=%s\n",
				dev.BDF.String(),
				dev.ClassDescription(),
				dev.ClassCode>>8,
				description,
				dev.VendorID,
				dev.DeviceID,
				driver,
				"",
				iommuStr,
				vfioStatus,
			)
		}
		w.Flush()

		fmt.Printf("\nTotal: %d devices\n", len(devices))
		return nil
	},
}

// vfioCheck returns a short VFIO compatibility label.
func vfioCheck(driver, iommuGroup, bdf string) string {
	if driver == "vfio-pci" {
		return "ready"
	}

	if iommuGroup == "-" {
		return "no-iommu"
	}

	// Check if there are other devices in the same IOMMU group
	groupPath := filepath.Join("/sys/kernel/iommu_groups", iommuGroup, "devices")
	entries, err := os.ReadDir(groupPath)
	if err != nil {
		return "ok"
	}

	peers := 0
	for _, e := range entries {
		if e.Name() != bdf {
			peers++
		}
	}

	if peers > 0 {
		return "group(" + strconv.Itoa(peers+1) + ")"
	}

	return "ok"
}

func init() {
	rootCmd.AddCommand(scanCmd)
}
