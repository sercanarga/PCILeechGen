package main

import (
	"fmt"
	"strings"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/color"
	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/donor/vfio"
	"github.com/sercanarga/pcileechgen/internal/firmware"
	"github.com/sercanarga/pcileechgen/internal/pci"
	"github.com/spf13/cobra"
)

var checkDevice string

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check VFIO compatibility for a PCI device",
	Long: `Runs diagnostic checks on a PCI device to verify it can be used
as a donor device with VFIO. Also shows board compatibility analysis.

Example:
  pcileechgen check --bdf 0000:03:00.0`,
	RunE: func(cmd *cobra.Command, args []string) error {
		bdf, err := pci.ParseBDF(checkDevice)
		if err != nil {
			return fmt.Errorf("invalid BDF: %w", err)
		}

		c := &checker{bdf: bdf, sysfs: donor.NewSysfsReader()}
		return c.run()
	},
}

// checker bundles state for a single device compatibility check.
type checker struct {
	bdf    pci.BDF
	sysfs  *donor.SysfsReader
	dev    *pci.PCIDevice
	cs     *pci.ConfigSpace
	issues int
}

func (c *checker) run() error {
	fmt.Printf("Checking device %s...\n\n", color.Bold(c.bdf.String()))

	c.checkDeviceInfo()
	c.checkConfigSpace()
	c.checkVFIO()
	c.checkIOMMUGroup()
	c.checkPowerState()
	c.checkDriver()
	c.showCapabilities()
	c.showBARs()
	c.showBoardCompatibility()

	fmt.Printf("\n%s\n", color.Header("Summary"))
	if c.issues == 0 {
		fmt.Println(color.OK("Device is ready for firmware generation"))
	} else {
		fmt.Println(color.Failf("%d issue(s) found — see above for details", c.issues))
	}
	return nil
}

func (c *checker) checkDeviceInfo() {
	dev, err := c.sysfs.ReadDeviceInfo(c.bdf)
	if err != nil {
		fmt.Println(color.Failf("Cannot read device info: %v", err))
		c.issues++
		return
	}
	c.dev = dev
	fmt.Println(color.Okf("Device found: %04x:%04x %s", dev.VendorID, dev.DeviceID, dev.ClassDescription()))
}

func (c *checker) checkConfigSpace() {
	cs, err := c.sysfs.ReadConfigSpace(c.bdf)
	if err != nil {
		fmt.Println(color.Failf("Cannot read config space: %v", err))
		c.issues++
		return
	}
	c.cs = cs
	fmt.Println(color.Okf("Config space readable: %d bytes", cs.Size))
}

func (c *checker) checkVFIO() {
	if err := vfio.CheckIOMMU(); err != nil {
		fmt.Println(color.Failf("IOMMU: %v", err))
		c.issues++
	} else {
		fmt.Println(color.OK("IOMMU is enabled"))
	}

	if err := vfio.CheckVFIOModules(); err != nil {
		fmt.Println(color.Failf("VFIO modules: %v", err))
		c.issues++
	} else {
		fmt.Println(color.OK("VFIO modules loaded"))
	}
}

func (c *checker) checkIOMMUGroup() {
	group, err := vfio.GetIOMMUGroup(c.bdf.String())
	if err != nil {
		fmt.Println(color.Warnf("IOMMU group: %v", err))
		return
	}
	fmt.Println(color.Okf("IOMMU group: %d", group))

	groupDevs, err := vfio.ListIOMMUGroupDevices(c.bdf.String())
	if err != nil || len(groupDevs) <= 1 {
		if err == nil {
			fmt.Println(color.OK("Device is alone in its IOMMU group"))
		}
		return
	}

	var others []string
	for _, d := range groupDevs {
		if d != c.bdf.String() {
			others = append(others, d)
		}
	}
	if len(others) > 0 {
		fmt.Println(color.Warnf("IOMMU group shared with %d device(s): %s",
			len(others), strings.Join(others, ", ")))
		fmt.Println(color.Dim("  All devices in the group must be unbound or on vfio-pci"))
	}
}

func (c *checker) checkPowerState() {
	ps, err := vfio.CheckPowerState(c.bdf.String())
	if err != nil {
		return
	}
	if ps == "D0" {
		fmt.Println(color.Okf("Power state: %s (active)", ps))
	} else {
		fmt.Println(color.Failf("Power state: %s — device should be in D0 for reliable reads", ps))
		fmt.Println(color.Dim(fmt.Sprintf("  Fix: echo 0 | sudo tee /sys/bus/pci/devices/%s/d3cold_allowed", c.bdf.String())))
		c.issues++
	}
}

func (c *checker) checkDriver() {
	if c.dev == nil {
		return
	}
	if c.dev.Driver == "" {
		fmt.Println(color.OK("No driver bound"))
	} else if c.dev.Driver == "vfio-pci" {
		fmt.Println(color.OK("Already bound to vfio-pci"))
	} else {
		fmt.Println(color.Warnf("Currently bound to %q (will need unbinding)", c.dev.Driver))
	}
}

func (c *checker) showCapabilities() {
	if c.cs == nil {
		return
	}

	caps := pci.ParseCapabilities(c.cs)
	fmt.Printf("\nCapabilities (%d):\n", len(caps))
	for _, cap := range caps {
		fmt.Printf("  [%02x] %s at offset 0x%02x\n",
			cap.ID, pci.CapabilityName(cap.ID), cap.Offset)
	}

	extCaps := pci.ParseExtCapabilities(c.cs)
	if len(extCaps) > 0 {
		fmt.Printf("\nExtended Capabilities (%d):\n", len(extCaps))
		for _, cap := range extCaps {
			fmt.Printf("  [%04x] %s at offset 0x%03x\n",
				cap.ID, pci.ExtCapabilityName(cap.ID), cap.Offset)
		}
	}
}

func (c *checker) showBARs() {
	bars, err := c.sysfs.ReadResourceFile(c.bdf)
	if err != nil {
		return
	}

	fmt.Printf("\nBARs:\n")
	for _, bar := range bars {
		if !bar.IsDisabled() {
			fmt.Printf("  %s\n", bar.String())
		}
	}

	barStatuses := vfio.CheckBARAccessibility(c.bdf.String())
	for _, bs := range barStatuses {
		if !bs.Accessible {
			fmt.Println(color.Warnf("  BAR%d: not accessible (%s)", bs.Index, bs.Error))
		}
	}
}

func (c *checker) showBoardCompatibility() {
	if c.cs == nil {
		return
	}

	fmt.Printf("\n%s\n", color.Header("Board Compatibility"))
	ids := firmware.ExtractDeviceIDs(c.cs, pci.ParseExtCapabilities(c.cs))

	if ids.HasPCIeCap {
		fmt.Printf("Donor Link: %s x%d\n",
			color.Bold(firmware.LinkSpeedName(ids.LinkSpeed)), ids.LinkWidth)
	}
	if ids.HasDSN {
		fmt.Printf("Donor DSN:  0x%s\n", firmware.DSNToSVHex(ids.DSN))
	} else {
		fmt.Println(color.Warn("Donor has no DSN capability (serial number won't be emulated)"))
	}

	allBoards := board.All()
	fmt.Printf("\nCompatible boards:\n")
	compatible := 0
	for _, b := range allBoards {
		label := color.Okf("%-22s %s x%d", b.Name, b.FPGAPart, b.PCIeLanes)
		note := ""
		if ids.HasPCIeCap && int(ids.LinkWidth) > b.PCIeLanes {
			label = color.Warnf("%-22s %s x%d", b.Name, b.FPGAPart, b.PCIeLanes)
			note = color.Dim(fmt.Sprintf(" (link clamped: x%d -> x%d)", ids.LinkWidth, b.PCIeLanes))
		}
		if ids.HasPCIeCap && int(ids.LinkWidth) == b.PCIeLanes {
			note = color.Dim(" (exact match)")
		}
		fmt.Printf("  %s%s\n", label, note)
		compatible++
	}
	fmt.Printf("\nTotal: %d boards\n", compatible)
}

func init() {
	checkCmd.Flags().StringVar(&checkDevice, "bdf", "", "device BDF address to check (required)")
	_ = checkCmd.MarkFlagRequired("bdf")
	rootCmd.AddCommand(checkCmd)
}
