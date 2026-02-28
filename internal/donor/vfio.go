package donor

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// VFIOManager handles VFIO device binding and unbinding.
type VFIOManager struct{}

// NewVFIOManager creates a new VFIOManager.
func NewVFIOManager() *VFIOManager {
	return &VFIOManager{}
}

// CheckIOMMU checks if IOMMU is enabled on the system.
func (vm *VFIOManager) CheckIOMMU() error {
	// Check for IOMMU groups directory
	if _, err := os.Stat("/sys/kernel/iommu_groups"); os.IsNotExist(err) {
		return fmt.Errorf("IOMMU not enabled: /sys/kernel/iommu_groups does not exist. " +
			"Enable IOMMU in BIOS and add 'intel_iommu=on' or 'amd_iommu=on' to kernel parameters")
	}

	// Check if any IOMMU groups exist
	entries, err := os.ReadDir("/sys/kernel/iommu_groups")
	if err != nil {
		return fmt.Errorf("failed to read IOMMU groups: %w", err)
	}
	if len(entries) == 0 {
		return fmt.Errorf("no IOMMU groups found: IOMMU may not be properly configured")
	}

	return nil
}

// CheckVFIOModules checks if VFIO kernel modules are loaded.
func (vm *VFIOManager) CheckVFIOModules() error {
	modules := []string{"vfio", "vfio-pci"}
	for _, mod := range modules {
		// Check sysfs for loaded module
		modPath := filepath.Join("/sys/module", strings.ReplaceAll(mod, "-", "_"))
		if _, err := os.Stat(modPath); os.IsNotExist(err) {
			return fmt.Errorf("kernel module %q not loaded. Run: sudo modprobe %s", mod, mod)
		}
	}
	return nil
}

// BindToVFIO binds a PCI device to the vfio-pci driver.
func (vm *VFIOManager) BindToVFIO(bdf string) error {
	devPath := filepath.Join("/sys/bus/pci/devices", bdf)

	// Check if already bound to vfio-pci
	driverLink, err := os.Readlink(filepath.Join(devPath, "driver"))
	if err == nil && filepath.Base(driverLink) == "vfio-pci" {
		return nil // already bound
	}

	// Read vendor and device IDs
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

	// Unbind from current driver if any
	if driverLink != "" {
		unbindPath := filepath.Join(filepath.Dir(driverLink), "unbind")
		if err := os.WriteFile(unbindPath, []byte(bdf), 0200); err != nil {
			return fmt.Errorf("failed to unbind from current driver: %w", err)
		}
	}

	// Set driver override
	overridePath := filepath.Join(devPath, "driver_override")
	if err := os.WriteFile(overridePath, []byte("vfio-pci"), 0200); err != nil {
		return fmt.Errorf("failed to set driver override: %w", err)
	}

	// Write to vfio-pci new_id
	newIDPath := "/sys/bus/pci/drivers/vfio-pci/new_id"
	idStr := fmt.Sprintf("%s %s", vendor, device)
	// Ignore error here - it may already be registered
	_ = os.WriteFile(newIDPath, []byte(idStr), 0200)

	// Probe the device
	probePath := "/sys/bus/pci/drivers_probe"
	if err := os.WriteFile(probePath, []byte(bdf), 0200); err != nil {
		return fmt.Errorf("failed to probe device: %w", err)
	}

	// Verify binding
	driverLink, err = os.Readlink(filepath.Join(devPath, "driver"))
	if err != nil || filepath.Base(driverLink) != "vfio-pci" {
		return fmt.Errorf("device %s was not bound to vfio-pci", bdf)
	}

	return nil
}

// UnbindFromVFIO unbinds a PCI device from vfio-pci driver.
func (vm *VFIOManager) UnbindFromVFIO(bdf string) error {
	devPath := filepath.Join("/sys/bus/pci/devices", bdf)

	// Clear driver override
	overridePath := filepath.Join(devPath, "driver_override")
	_ = os.WriteFile(overridePath, []byte(""), 0200)

	// Unbind from vfio-pci
	unbindPath := "/sys/bus/pci/drivers/vfio-pci/unbind"
	if err := os.WriteFile(unbindPath, []byte(bdf), 0200); err != nil {
		return fmt.Errorf("failed to unbind from vfio-pci: %w", err)
	}

	// Reprobe for original driver
	probePath := "/sys/bus/pci/drivers_probe"
	if err := os.WriteFile(probePath, []byte(bdf), 0200); err != nil {
		return fmt.Errorf("failed to reprobe device: %w", err)
	}

	return nil
}

// GetIOMMUGroup returns the IOMMU group number for a device.
func (vm *VFIOManager) GetIOMMUGroup(bdf string) (int, error) {
	devPath := filepath.Join("/sys/bus/pci/devices", bdf)
	link, err := os.Readlink(filepath.Join(devPath, "iommu_group"))
	if err != nil {
		return -1, fmt.Errorf("failed to read IOMMU group: %w", err)
	}

	groupStr := filepath.Base(link)
	var group int
	_, err = fmt.Sscanf(groupStr, "%d", &group)
	if err != nil {
		return -1, fmt.Errorf("failed to parse IOMMU group number: %w", err)
	}

	return group, nil
}
