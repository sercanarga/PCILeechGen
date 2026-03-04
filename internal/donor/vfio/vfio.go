// Package vfio reads PCI config space and BAR memory via Linux VFIO ioctls.
// Needs IOMMU + vfio-pci driver. Non-Linux builds compile but return errors.
package vfio

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// DeviceDump is everything we pulled from a single device.
type DeviceDump struct {
	BDF             string         `json:"bdf"`
	ConfigSpace     []byte         `json:"config_space"`
	ConfigSpaceSize int            `json:"config_space_size"`
	BARContents     map[int][]byte `json:"bar_contents,omitempty"`
	BARInfo         []BARRegion    `json:"bars"`
}

// BARRegion describes a single BAR region as reported by VFIO.
type BARRegion struct {
	Index int    `json:"index"`
	Size  uint64 `json:"size"`
	Flags uint32 `json:"flags"`
}

func (d *DeviceDump) ToJSON() ([]byte, error) {
	return json.MarshalIndent(d, "", "  ")
}

// ResolveIOMMUGroup returns the /dev/vfio/<N> path for a BDF.
func ResolveIOMMUGroup(bdf string) (string, error) {
	linkPath := filepath.Join("/sys/bus/pci/devices", bdf, "iommu_group")
	target, err := os.Readlink(linkPath)
	if err != nil {
		return "", fmt.Errorf("cannot resolve IOMMU group for %s: %w (is IOMMU enabled?)", bdf, err)
	}
	groupNum := filepath.Base(target)
	return filepath.Join("/dev/vfio", groupNum), nil
}

// IsBoundToVFIO returns true if the device's current driver is vfio-pci.
func IsBoundToVFIO(bdf string) bool {
	driverLink := filepath.Join("/sys/bus/pci/devices", bdf, "driver")
	target, err := os.Readlink(driverLink)
	if err != nil {
		return false
	}
	return filepath.Base(target) == "vfio-pci"
}

// ListVFIODevices enumerates BDFs currently on the vfio-pci driver.
func ListVFIODevices() ([]string, error) {
	entries, err := os.ReadDir("/sys/bus/pci/drivers/vfio-pci")
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil // vfio-pci not loaded
		}
		return nil, err
	}

	var devices []string
	for _, e := range entries {
		name := e.Name()
		// BDF format: NNNN:BB:DD.F
		if len(name) >= 7 && name[4] == ':' {
			devices = append(devices, name)
		}
	}
	return devices, nil
}

// CheckIOMMU verifies /sys/kernel/iommu_groups exists and is populated.
func CheckIOMMU() error {
	if _, err := os.Stat("/sys/kernel/iommu_groups"); os.IsNotExist(err) {
		return fmt.Errorf("IOMMU not enabled: /sys/kernel/iommu_groups does not exist. " +
			"Enable IOMMU in BIOS and add 'intel_iommu=on' or 'amd_iommu=on' to kernel parameters")
	}
	entries, err := os.ReadDir("/sys/kernel/iommu_groups")
	if err != nil {
		return fmt.Errorf("failed to read IOMMU groups: %w", err)
	}
	if len(entries) == 0 {
		return fmt.Errorf("no IOMMU groups found: IOMMU may not be properly configured")
	}
	return nil
}

// CheckVFIOModules makes sure vfio and vfio-pci are loaded.
func CheckVFIOModules() error {
	modules := []string{"vfio", "vfio-pci"}
	for _, mod := range modules {
		modPath := filepath.Join("/sys/module", strings.ReplaceAll(mod, "-", "_"))
		if _, err := os.Stat(modPath); os.IsNotExist(err) {
			return fmt.Errorf("kernel module %q not loaded. Run: sudo modprobe %s", mod, mod)
		}
	}
	return nil
}

// BindToVFIO unbinds from the current driver and rebinds to vfio-pci.
func BindToVFIO(bdf string) error {
	devPath := filepath.Join("/sys/bus/pci/devices", bdf)

	driverLink, err := os.Readlink(filepath.Join(devPath, "driver"))
	if err == nil && filepath.Base(driverLink) == "vfio-pci" {
		return nil // already bound
	}

	vendorData, err := os.ReadFile(filepath.Join(devPath, "vendor"))
	if err != nil {
		return fmt.Errorf("failed to read vendor: %w", err)
	}
	deviceData, err := os.ReadFile(filepath.Join(devPath, "device"))
	if err != nil {
		return fmt.Errorf("failed to read device: %w", err)
	}

	vendor := strings.TrimSpace(string(vendorData))
	device := strings.TrimSpace(string(deviceData))

	if driverLink != "" {
		unbindPath := filepath.Join(filepath.Dir(driverLink), "unbind")
		if err := os.WriteFile(unbindPath, []byte(bdf), 0200); err != nil {
			return fmt.Errorf("failed to unbind from current driver: %w", err)
		}
	}

	overridePath := filepath.Join(devPath, "driver_override")
	if err := os.WriteFile(overridePath, []byte("vfio-pci"), 0200); err != nil {
		return fmt.Errorf("failed to set driver override: %w", err)
	}

	newIDPath := "/sys/bus/pci/drivers/vfio-pci/new_id"
	idStr := fmt.Sprintf("%s %s", vendor, device)
	_ = os.WriteFile(newIDPath, []byte(idStr), 0200)

	probePath := "/sys/bus/pci/drivers_probe"
	if err := os.WriteFile(probePath, []byte(bdf), 0200); err != nil {
		return fmt.Errorf("failed to probe device: %w", err)
	}

	driverLink, err = os.Readlink(filepath.Join(devPath, "driver"))
	if err != nil || filepath.Base(driverLink) != "vfio-pci" {
		return fmt.Errorf("device %s was not bound to vfio-pci", bdf)
	}
	return nil
}

// UnbindFromVFIO detaches from vfio-pci and re-probes for the original driver.
func UnbindFromVFIO(bdf string) error {
	devPath := filepath.Join("/sys/bus/pci/devices", bdf)
	_ = os.WriteFile(filepath.Join(devPath, "driver_override"), []byte(""), 0200)

	if err := os.WriteFile("/sys/bus/pci/drivers/vfio-pci/unbind", []byte(bdf), 0200); err != nil {
		return fmt.Errorf("failed to unbind from vfio-pci: %w", err)
	}

	if err := os.WriteFile("/sys/bus/pci/drivers_probe", []byte(bdf), 0200); err != nil {
		return fmt.Errorf("failed to reprobe device: %w", err)
	}
	return nil
}

// GetIOMMUGroup reads the IOMMU group number from sysfs.
func GetIOMMUGroup(bdf string) (int, error) {
	link, err := os.Readlink(filepath.Join("/sys/bus/pci/devices", bdf, "iommu_group"))
	if err != nil {
		return -1, fmt.Errorf("failed to read IOMMU group: %w", err)
	}
	var group int
	_, err = fmt.Sscanf(filepath.Base(link), "%d", &group)
	if err != nil {
		return -1, fmt.Errorf("failed to parse IOMMU group number: %w", err)
	}
	return group, nil
}
