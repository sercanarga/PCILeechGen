package vfio

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type BARStatus struct {
	Index      int
	Size       uint64
	Accessible bool
	Error      string
}

// CheckPowerState reads D-state from sysfs (D0/D1/D2/D3hot/D3cold).
func CheckPowerState(bdf string) (string, error) {
	path := filepath.Join("/sys/bus/pci/devices", bdf, "power_state")
	data, err := os.ReadFile(path)
	if err != nil {
		return "unknown", fmt.Errorf("cannot read power state: %w", err)
	}
	return strings.TrimSpace(string(data)), nil
}

// ListIOMMUGroupDevices returns all BDFs sharing the same IOMMU group.
func ListIOMMUGroupDevices(bdf string) ([]string, error) {
	groupLink := filepath.Join("/sys/bus/pci/devices", bdf, "iommu_group", "devices")
	entries, err := os.ReadDir(groupLink)
	if err != nil {
		return nil, fmt.Errorf("cannot list IOMMU group devices: %w", err)
	}

	var devices []string
	for _, e := range entries {
		devices = append(devices, e.Name())
	}
	return devices, nil
}

// CheckBARAccessibility tries to open each BAR resource file.
func CheckBARAccessibility(bdf string) []BARStatus {
	var results []BARStatus
	for i := 0; i < 6; i++ {
		resPath := filepath.Join("/sys/bus/pci/devices", bdf, fmt.Sprintf("resource%d", i))
		info, err := os.Stat(resPath)
		if err != nil {
			continue // BAR doesn't exist or is disabled
		}
		status := BARStatus{Index: i, Size: uint64(info.Size())}
		if info.Size() == 0 {
			continue // disabled BAR
		}

		f, err := os.Open(resPath)
		if err != nil {
			status.Accessible = false
			status.Error = err.Error()
		} else {
			status.Accessible = true
			f.Close()
		}
		results = append(results, status)
	}
	return results
}

type DiagnosticResult struct {
	Name    string
	Passed  bool
	Message string
}

// RunDiagnostics runs the full VFIO readiness check suite.
func RunDiagnostics(bdf string) []DiagnosticResult {
	var results []DiagnosticResult

	if err := CheckIOMMU(); err != nil {
		results = append(results, DiagnosticResult{"IOMMU", false, err.Error()})
	} else {
		results = append(results, DiagnosticResult{"IOMMU", true, "IOMMU is enabled"})
	}

	if err := CheckVFIOModules(); err != nil {
		results = append(results, DiagnosticResult{"VFIO Modules", false, err.Error()})
	} else {
		results = append(results, DiagnosticResult{"VFIO Modules", true, "vfio and vfio-pci loaded"})
	}

	group, err := GetIOMMUGroup(bdf)
	if err != nil {
		results = append(results, DiagnosticResult{"IOMMU Group", false, err.Error()})
	} else {
		results = append(results, DiagnosticResult{"IOMMU Group", true, fmt.Sprintf("group %d", group)})
	}

	groupDevs, err := ListIOMMUGroupDevices(bdf)
	if err == nil && len(groupDevs) > 1 {
		others := []string{}
		for _, d := range groupDevs {
			if d != bdf {
				others = append(others, d)
			}
		}
		if len(others) > 0 {
			results = append(results, DiagnosticResult{
				"Group Isolation", false,
				fmt.Sprintf("shared IOMMU group with %d other device(s): %s — all must be unbound or on vfio-pci",
					len(others), strings.Join(others, ", ")),
			})
		}
	} else if err == nil {
		results = append(results, DiagnosticResult{"Group Isolation", true, "device is alone in its IOMMU group"})
	}

	ps, err := CheckPowerState(bdf)
	if err == nil {
		if ps == "D0" {
			results = append(results, DiagnosticResult{"Power State", true, "D0 (active)"})
		} else {
			results = append(results, DiagnosticResult{
				"Power State", false,
				fmt.Sprintf("%s — device must be in D0 for reliable reads. Try: echo 0 | sudo tee /sys/bus/pci/devices/%s/d3cold_allowed", ps, bdf),
			})
		}
	}

	if IsBoundToVFIO(bdf) {
		results = append(results, DiagnosticResult{"Driver", true, "vfio-pci"})
	} else {
		driverLink, err := os.Readlink(filepath.Join("/sys/bus/pci/devices", bdf, "driver"))
		if err == nil {
			drv := filepath.Base(driverLink)
			results = append(results, DiagnosticResult{
				"Driver", false,
				fmt.Sprintf("currently bound to %q — run: pcileechgen build --bdf %s (auto-binds)", drv, bdf),
			})
		} else {
			results = append(results, DiagnosticResult{"Driver", true, "no driver bound (ready for vfio-pci)"})
		}
	}

	barStatuses := CheckBARAccessibility(bdf)
	for _, bs := range barStatuses {
		if bs.Accessible {
			results = append(results, DiagnosticResult{
				fmt.Sprintf("BAR%d", bs.Index), true,
				fmt.Sprintf("accessible (%d bytes)", bs.Size),
			})
		} else {
			results = append(results, DiagnosticResult{
				fmt.Sprintf("BAR%d", bs.Index), false,
				fmt.Sprintf("not accessible: %s", bs.Error),
			})
		}
	}

	return results
}
