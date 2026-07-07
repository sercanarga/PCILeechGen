package main

import (
	"encoding/json"
	"fmt"
	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/color"
	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/donor/vfio"
	"github.com/sercanarga/pcileechgen/internal/pci"
	"github.com/sercanarga/pcileechgen/internal/vivado"
	"github.com/spf13/cobra"
	"io"
	"strings"
)

type doctorOptions struct {
	bdf        string
	board      string
	vivadoPath string
	skipVivado bool
	jsonOutput bool
}

type doctorDeviceInfo struct {
	vendorID  uint16
	deviceID  uint16
	className string
}

type doctorBAR struct {
	index int
	size  uint64
}

type doctorCheck struct {
	name    string
	status  string
	message string
	fix     string
}

type doctorProbes struct {
	checkIOMMU     func() error
	checkVFIOMods  func() error
	readDeviceInfo func(string) (doctorDeviceInfo, error)
	readIOMMUGroup func(string) ([]string, error)
	readPowerState func(string) (string, error)
	readDriver     func(string) (string, error)
	readBARs       func(string) ([]doctorBAR, error)
	boardBRAMSize  func(string) (int, error)
	findVivado     func(string) (string, error)
}

var doctorOpts doctorOptions

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Run host and donor preflight checks",
	Long: `Checks the host environment, optional donor BDF, board fit, and Vivado
availability before running a firmware build.

Example:
  pcileechgen doctor
  pcileechgen doctor --bdf 0000:03:00.0 --board PCIeSquirrel
  pcileechgen doctor --bdf 0000:03:00.0 --board PCIeSquirrel --json`,
	RunE: func(cmd *cobra.Command, args []string) error {
		checks, err := runDoctor(doctorOpts, defaultDoctorProbes())
		if err != nil {
			return err
		}
		if doctorOpts.jsonOutput {
			data, err := json.MarshalIndent(checks, "", "  ")
			if err != nil {
				return err
			}
			fmt.Fprintln(cmd.OutOrStdout(), string(data))
			return doctorError(checks)
		}
		printDoctorChecks(cmd.OutOrStdout(), checks)
		return doctorError(checks)
	},
}

func init() {
	doctorCmd.Flags().StringVar(&doctorOpts.bdf, "bdf", "", "donor device BDF address")
	doctorCmd.Flags().StringVar(&doctorOpts.board, "board", "", "target board name")
	doctorCmd.Flags().StringVar(&doctorOpts.vivadoPath, "vivado-path", "", "Vivado installation path")
	doctorCmd.Flags().BoolVar(&doctorOpts.skipVivado, "skip-vivado", false, "skip Vivado availability check")
	doctorCmd.Flags().BoolVar(&doctorOpts.jsonOutput, "json", false, "output JSON")
	rootCmd.AddCommand(doctorCmd)
}

func runDoctor(opts doctorOptions, probes doctorProbes) ([]doctorCheck, error) {
	probes = fillDoctorProbes(probes)
	var checks []doctorCheck

	if err := probes.checkIOMMU(); err != nil {
		checks = append(checks, doctorFail("IOMMU", err.Error(), "enable IOMMU in BIOS and add intel_iommu=on or amd_iommu=on with iommu=pt, then reboot"))
	} else {
		checks = append(checks, doctorPass("IOMMU", "IOMMU groups are available"))
	}

	if err := probes.checkVFIOMods(); err != nil {
		checks = append(checks, doctorFail("VFIO modules", err.Error(), "run: sudo modprobe vfio vfio-pci"))
	} else {
		checks = append(checks, doctorPass("VFIO modules", "vfio and vfio-pci are loaded"))
	}

	if opts.bdf != "" {
		if _, err := pci.ParseBDF(opts.bdf); err != nil {
			return nil, fmt.Errorf("invalid BDF: %w", err)
		}
		checks = append(checks, runDoctorDeviceChecks(opts, probes)...)
	}

	checks = append(checks, runDoctorVivadoCheck(opts, probes))
	return checks, nil
}

func runDoctorDeviceChecks(opts doctorOptions, probes doctorProbes) []doctorCheck {
	checks := make([]doctorCheck, 0, 5)
	info, err := probes.readDeviceInfo(opts.bdf)
	if err != nil {
		checks = append(checks, doctorFail("Device", err.Error(), "check the BDF and run pcileechgen scan"))
	} else {
		checks = append(checks, doctorPass("Device", fmt.Sprintf("%04x:%04x %s", info.vendorID, info.deviceID, info.className)))
	}

	if group, err := probes.readIOMMUGroup(opts.bdf); err != nil {
		checks = append(checks, doctorFail("IOMMU group", err.Error(), "enable IOMMU and verify the device has an IOMMU group"))
	} else {
		checks = append(checks, doctorIOMMUGroupCheck(opts.bdf, group))
	}

	if state, err := probes.readPowerState(opts.bdf); err == nil {
		checks = append(checks, doctorPowerCheck(opts.bdf, state))
	} else {
		checks = append(checks, doctorWarn("Power state", err.Error(), "if reads fail, keep the donor in D0 and disable runtime PM"))
	}

	if driver, err := probes.readDriver(opts.bdf); err == nil {
		checks = append(checks, doctorDriverCheck(driver))
	} else {
		checks = append(checks, doctorWarn("Driver", err.Error(), "bind the donor to vfio-pci before build"))
	}

	if opts.board != "" {
		checks = append(checks, doctorBoardCheck(opts, probes))
	}
	return checks
}

func doctorIOMMUGroupCheck(bdf string, group []string) doctorCheck {
	peers := make([]string, 0, len(group))
	for _, dev := range group {
		if dev != bdf {
			peers = append(peers, dev)
		}
	}
	if len(peers) == 0 {
		return doctorPass("IOMMU group", "device is alone in its IOMMU group")
	}
	return doctorWarn("IOMMU group", fmt.Sprintf("shared with %s", strings.Join(peers, ", ")), "unbind every device in the group or bind the full group to vfio-pci")
}

func doctorPowerCheck(bdf, state string) doctorCheck {
	if state == "D0" {
		return doctorPass("Power state", "D0 active")
	}
	return doctorFail("Power state", fmt.Sprintf("%s is not D0", state), fmt.Sprintf("run: echo on | sudo tee /sys/bus/pci/devices/%s/power/control", bdf))
}

func doctorDriverCheck(driver string) doctorCheck {
	switch driver {
	case "vfio-pci":
		return doctorPass("Driver", "vfio-pci")
	case "":
		return doctorPass("Driver", "no driver bound")
	default:
		return doctorWarn("Driver", fmt.Sprintf("currently bound to %q", driver), "unbind the native driver or let build bind the donor to vfio-pci")
	}
}

func doctorBoardCheck(opts doctorOptions, probes doctorProbes) doctorCheck {
	bram, err := probes.boardBRAMSize(opts.board)
	if err != nil {
		return doctorFail("Board compatibility", err.Error(), "run pcileechgen boards and choose a supported board")
	}
	bars, err := probes.readBARs(opts.bdf)
	if err != nil {
		return doctorWarn("Board compatibility", err.Error(), "run pcileechgen check --bdf "+opts.bdf)
	}
	largest := doctorBAR{index: -1}
	for _, bar := range bars {
		if bar.size > largest.size {
			largest = bar
		}
	}
	if largest.size > uint64(bram) {
		return doctorFail("Board compatibility", fmt.Sprintf("BAR%d size %d exceeds %s BRAM %d", largest.index, largest.size, opts.board, bram), "choose a board with larger BRAM or a smaller-BAR donor")
	}
	return doctorPass("Board compatibility", fmt.Sprintf("largest BAR fits %s BRAM (%d <= %d)", opts.board, largest.size, bram))
}

func runDoctorVivadoCheck(opts doctorOptions, probes doctorProbes) doctorCheck {
	if opts.skipVivado {
		return doctorPass("Vivado", "skipped by --skip-vivado")
	}
	path, err := probes.findVivado(opts.vivadoPath)
	if err != nil {
		return doctorWarn("Vivado", err.Error(), "use --skip-vivado for artifacts only or pass --vivado-path")
	}
	return doctorPass("Vivado", path)
}

func fillDoctorProbes(probes doctorProbes) doctorProbes {
	defaults := defaultDoctorProbes()
	if probes.checkIOMMU == nil {
		probes.checkIOMMU = defaults.checkIOMMU
	}
	if probes.checkVFIOMods == nil {
		probes.checkVFIOMods = defaults.checkVFIOMods
	}
	if probes.readDeviceInfo == nil {
		probes.readDeviceInfo = defaults.readDeviceInfo
	}
	if probes.readIOMMUGroup == nil {
		probes.readIOMMUGroup = defaults.readIOMMUGroup
	}
	if probes.readPowerState == nil {
		probes.readPowerState = defaults.readPowerState
	}
	if probes.readDriver == nil {
		probes.readDriver = defaults.readDriver
	}
	if probes.readBARs == nil {
		probes.readBARs = defaults.readBARs
	}
	if probes.boardBRAMSize == nil {
		probes.boardBRAMSize = defaults.boardBRAMSize
	}
	if probes.findVivado == nil {
		probes.findVivado = defaults.findVivado
	}
	return probes
}

func defaultDoctorProbes() doctorProbes {
	sr := donor.NewSysfsReader()
	return doctorProbes{
		checkIOMMU:    vfio.CheckIOMMU,
		checkVFIOMods: vfio.CheckVFIOModules,
		readDeviceInfo: func(bdf string) (doctorDeviceInfo, error) {
			parsed, err := pci.ParseBDF(bdf)
			if err != nil {
				return doctorDeviceInfo{}, err
			}
			dev, err := sr.ReadDeviceInfo(parsed)
			if err != nil {
				return doctorDeviceInfo{}, err
			}
			return doctorDeviceInfo{vendorID: dev.VendorID, deviceID: dev.DeviceID, className: dev.ClassDescription()}, nil
		},
		readIOMMUGroup: vfio.ListIOMMUGroupDevices,
		readPowerState: vfio.CheckPowerState,
		readDriver: func(bdf string) (string, error) {
			return vfio.BoundDriver(bdf), nil
		},
		readBARs: func(bdf string) ([]doctorBAR, error) {
			parsed, err := pci.ParseBDF(bdf)
			if err != nil {
				return nil, err
			}
			bars, err := sr.ReadResourceFile(parsed)
			if err != nil {
				return nil, err
			}
			out := make([]doctorBAR, 0, len(bars))
			for _, bar := range bars {
				if !bar.IsDisabled() {
					out = append(out, doctorBAR{index: bar.Index, size: bar.Size})
				}
			}
			return out, nil
		},
		boardBRAMSize: func(name string) (int, error) {
			b, err := board.Find(name)
			if err != nil {
				return 0, err
			}
			return b.BRAMSizeOrDefault(), nil
		},
		findVivado: func(path string) (string, error) {
			v, err := vivado.Find(path)
			if err != nil {
				return "", err
			}
			return v.BinaryPath(), nil
		},
	}
}

func doctorPass(name, message string) doctorCheck {
	return doctorCheck{name: name, status: "pass", message: message}
}

func doctorWarn(name, message, fix string) doctorCheck {
	return doctorCheck{name: name, status: "warn", message: message, fix: fix}
}

func doctorFail(name, message, fix string) doctorCheck {
	return doctorCheck{name: name, status: "fail", message: message, fix: fix}
}

func doctorError(checks []doctorCheck) error {
	fails := 0
	for _, check := range checks {
		if check.status == "fail" {
			fails++
		}
	}
	if fails > 0 {
		return fmt.Errorf("%d doctor check(s) failed", fails)
	}
	return nil
}

func printDoctorChecks(w io.Writer, checks []doctorCheck) {
	for _, check := range checks {
		label := check.status
		switch check.status {
		case "pass":
			label = color.OK("PASS")
		case "warn":
			label = color.Warn("WARN")
		case "fail":
			label = color.Fail("FAIL")
		}
		fmt.Fprintf(w, "%-45s %s %s\n", check.name, label, check.message)
		if check.fix != "" {
			fmt.Fprintf(w, "  fix: %s\n", check.fix)
		}
	}
}

func (c doctorCheck) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Name    string `json:"name"`
		Status  string `json:"status"`
		Message string `json:"message"`
		Fix     string `json:"fix,omitempty"`
	}{
		Name:    c.name,
		Status:  c.status,
		Message: c.message,
		Fix:     c.fix,
	})
}
