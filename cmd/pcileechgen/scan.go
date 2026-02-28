package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/sercanarga/pcileechgen/internal/color"
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

		for _, dev := range devices {
			// Device description from pci.ids
			devName := db.DeviceName(dev.VendorID, dev.DeviceID)
			vendorName := db.VendorName(dev.VendorID)
			description := dev.ClassDescription()
			if vendorName != "" && devName != "" {
				description = fmt.Sprintf("%s %s", vendorName, devName)
			} else if vendorName != "" {
				description = vendorName
			}

			// IOMMU group
			iommuStr := ""
			iommuLink, err := os.Readlink(filepath.Join("/sys/bus/pci/devices", dev.BDF.String(), "iommu_group"))
			if err == nil {
				iommuStr = filepath.Base(iommuLink)
			}

			// VFIO tag
			vfio := vfioCheck(dev.Driver, iommuStr, dev.BDF.String())
			vfioTag := ""
			switch vfio {
			case "ready":
				vfioTag = " " + color.OK("vfio")
			case "ok":
				vfioTag = ""
			case "no-iommu":
				vfioTag = " " + color.Warn("no-iommu")
			default:
				vfioTag = " " + color.Warn(vfio)
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

// vfioCheck returns a short VFIO compatibility label.
func vfioCheck(driver, iommuGroup, bdf string) string {
	if driver == "vfio-pci" {
		return "ready"
	}

	if iommuGroup == "" {
		return "no-iommu"
	}

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
