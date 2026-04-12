package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

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

		c := &checker{bdf: bdf, sysfs: donor.NewSysfsReader(), w: os.Stdout}
		return c.run()
	},
}

// checker bundles state for a single device compatibility check.
type checker struct {
	bdf    pci.BDF
	sysfs  *donor.SysfsReader
	w      io.Writer
	dev    *pci.PCIDevice
	cs     *pci.ConfigSpace
	issues int
}

func (c *checker) run() error {
	fmt.Fprintf(c.w, "Checking device %s...\n\n", color.Bold(c.bdf.String()))

	c.checkDeviceInfo()
	c.checkConfigSpace()
	c.checkVFIO()
	c.checkIOMMUGroup()
	c.checkPowerState()
	c.checkDriver()
	c.showCapabilities()
	c.showBARs()
	c.showBoardCompatibility()

	fmt.Fprintf(c.w, "\n%s\n", color.Header("Summary"))
	if c.issues == 0 {
		fmt.Fprintln(c.w, color.OK("Device is ready for firmware generation"))
	} else {
		fmt.Fprintln(c.w, color.Failf("%d issue(s) found - see above for details", c.issues))
	}
	return nil
}

func (c *checker) checkDeviceInfo() {
	dev, err := c.sysfs.ReadDeviceInfo(c.bdf)
	if err != nil {
		fmt.Fprintln(c.w, color.Failf("Cannot read device info: %v", err))
		c.issues++
		return
	}
	c.dev = dev
	fmt.Fprintln(c.w, color.Okf("Device found: %04x:%04x %s", dev.VendorID, dev.DeviceID, dev.ClassDescription()))
}

func (c *checker) checkConfigSpace() {
	cs, err := c.sysfs.ReadConfigSpace(c.bdf)
	if err != nil {
		fmt.Fprintln(c.w, color.Failf("Cannot read config space: %v", err))
		c.issues++
		return
	}
	c.cs = cs
	fmt.Fprintln(c.w, color.Okf("Config space readable: %d bytes", cs.Size))
}

func (c *checker) checkVFIO() {
	if err := vfio.CheckIOMMU(); err != nil {
		fmt.Fprintln(c.w, color.Failf("IOMMU: %v", err))
		c.issues++
	} else {
		fmt.Fprintln(c.w, color.OK("IOMMU is enabled"))
	}

	if err := vfio.CheckVFIOModules(); err != nil {
		fmt.Fprintln(c.w, color.Failf("VFIO modules: %v", err))
		c.issues++
	} else {
		fmt.Fprintln(c.w, color.OK("VFIO modules loaded"))
	}
}

func (c *checker) checkIOMMUGroup() {
	group, err := vfio.GetIOMMUGroup(c.bdf.String())
	if err != nil {
		fmt.Fprintln(c.w, color.Warnf("IOMMU group: %v", err))
		return
	}
	fmt.Fprintln(c.w, color.Okf("IOMMU group: %d", group))

	groupDevs, err := vfio.ListIOMMUGroupDevices(c.bdf.String())
	if err != nil || len(groupDevs) <= 1 {
		if err == nil {
			fmt.Fprintln(c.w, color.OK("Device is alone in its IOMMU group"))
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
		fmt.Fprintln(c.w, color.Warnf("IOMMU group shared with %d device(s): %s",
			len(others), strings.Join(others, ", ")))
		fmt.Fprintln(c.w, color.Dim("  All devices in the group must be unbound or on vfio-pci"))
	}
}

func (c *checker) checkPowerState() {
	ps, err := vfio.CheckPowerState(c.bdf.String())
	if err != nil {
		return
	}
	if ps == "D0" {
		fmt.Fprintln(c.w, color.Okf("Power state: %s (active)", ps))
		return
	}

	// attempt auto-wake
	fmt.Fprintf(c.w, color.Dim("Power state: %s - attempting D0 wake...\n"), ps)
	if err := vfio.WakeToD0(c.bdf.String()); err != nil {
		fmt.Fprintln(c.w, color.Failf("Power state: %s - failed to wake device: %v", ps, err))
		c.issues++
		return
	}

	// settle and re-check
	time.Sleep(150 * time.Millisecond)
	ps2, err := vfio.CheckPowerState(c.bdf.String())
	if err == nil && ps2 == "D0" {
		fmt.Fprintln(c.w, color.OK("Power state: D0 (auto-woken)"))
	} else {
		fmt.Fprintln(c.w, color.Failf("Power state: %s (still not D0 after wake attempt)", ps2))
		c.issues++
	}
}

func (c *checker) checkDriver() {
	if c.dev == nil {
		return
	}
	if c.dev.Driver == "" {
		fmt.Fprintln(c.w, color.OK("No driver bound"))
	} else if c.dev.Driver == "vfio-pci" {
		fmt.Fprintln(c.w, color.OK("Already bound to vfio-pci"))
	} else {
		fmt.Fprintln(c.w, color.Warnf("Currently bound to %q (will need unbinding)", c.dev.Driver))
	}
}

func (c *checker) showCapabilities() {
	if c.cs == nil {
		return
	}

	caps := pci.ParseCapabilities(c.cs)
	fmt.Fprintf(c.w, "\nCapabilities (%d):\n", len(caps))
	for _, cap := range caps {
		fmt.Fprintf(c.w, "  [%02x] %s at offset 0x%02x\n",
			cap.ID, pci.CapabilityName(cap.ID), cap.Offset)
	}

	extCaps := pci.ParseExtCapabilities(c.cs)
	if len(extCaps) > 0 {
		fmt.Fprintf(c.w, "\nExtended Capabilities (%d):\n", len(extCaps))
		for _, cap := range extCaps {
			fmt.Fprintf(c.w, "  [%04x] %s at offset 0x%03x\n",
				cap.ID, pci.ExtCapabilityName(cap.ID), cap.Offset)
		}
	}
}

func (c *checker) showBARs() {
	bars, err := c.sysfs.ReadResourceFile(c.bdf)
	if err != nil {
		return
	}

	fmt.Fprintf(c.w, "\nBARs:\n")
	for _, bar := range bars {
		if !bar.IsDisabled() {
			fmt.Fprintf(c.w, "  %s\n", bar.String())
		}
	}

	barStatuses := vfio.CheckBARAccessibility(c.bdf.String())
	for _, bs := range barStatuses {
		if !bs.Accessible {
			fmt.Fprintln(c.w, color.Warnf("  BAR%d: not accessible (%s)", bs.Index, bs.Error))
		}
	}
}

func (c *checker) showBoardCompatibility() {
	if c.cs == nil {
		return
	}

	fmt.Fprintf(c.w, "\n%s\n", color.Header("Board Compatibility"))
	ids := firmware.ExtractDeviceIDs(c.cs, pci.ParseExtCapabilities(c.cs))

	if ids.HasPCIeCap {
		fmt.Fprintf(c.w, "Donor Link: %s x%d\n",
			color.Bold(firmware.LinkSpeedName(ids.LinkSpeed)), ids.LinkWidth)
	}
	if ids.HasDSN {
		fmt.Fprintf(c.w, "Donor DSN:  0x%s\n", firmware.DSNToSVHex(ids.DSN))
	} else {
		fmt.Fprintln(c.w, color.Warn("Donor has no DSN capability (serial number won't be emulated)"))
	}

	allBoards := board.All()
	fmt.Fprintf(c.w, "\nCompatible boards:\n")
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
		fmt.Fprintf(c.w, "  %s%s\n", label, note)
		compatible++
	}
	fmt.Fprintf(c.w, "\nTotal: %d boards\n", compatible)
}

func init() {
	checkCmd.Flags().StringVar(&checkDevice, "bdf", "", "device BDF address to check (required)")
	_ = checkCmd.MarkFlagRequired("bdf")
	rootCmd.AddCommand(checkCmd)
}
