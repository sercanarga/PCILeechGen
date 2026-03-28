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

var sysfsBase = "/sys/bus/pci/devices"

// SetSysfsBase overrides the sysfs path (for testing).
func SetSysfsBase(path string) { sysfsBase = path }

// ResetSysfsBase restores the default sysfs path.
func ResetSysfsBase() { sysfsBase = "/sys/bus/pci/devices" }

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
	linkPath := filepath.Join(sysfsBase, bdf, "iommu_group")
	target, err := os.Readlink(linkPath)
	if err != nil {
		return "", fmt.Errorf("cannot resolve IOMMU group for %s: %w (is IOMMU enabled?)", bdf, err)
	}
	groupNum := filepath.Base(target)
	return filepath.Join("/dev/vfio", groupNum), nil
}

// IsBoundToVFIO returns true if the device's current driver is vfio-pci.
func IsBoundToVFIO(bdf string) bool {
	driverLink := filepath.Join(sysfsBase, bdf, "driver")
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
	devPath := filepath.Join(sysfsBase, bdf)

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
		unbindPath := filepath.Join(devPath, "driver", "unbind")
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

// EnableMemorySpace sets Memory Space (bit 1) and Bus Master (bit 2) in the
// PCI Command Register via sysfs config. Needed after vfio-pci bind.
func EnableMemorySpace(bdf string) error {
	configPath := filepath.Join(sysfsBase, bdf, "config")

	f, err := os.OpenFile(configPath, os.O_RDWR, 0)
	if err != nil {
		return fmt.Errorf("cannot open config space for %s: %w", bdf, err)
	}
	defer f.Close()

	// PCI Command Register is at offset 0x04, 2 bytes wide
	var cmd [2]byte
	if _, err := f.ReadAt(cmd[:], 0x04); err != nil {
		return fmt.Errorf("cannot read PCI Command Register for %s: %w", bdf, err)
	}

	const (
		memorySpaceEnable = 0x02 // bit 1
		busMasterEnable   = 0x04 // bit 2
	)

	needed := byte(memorySpaceEnable | busMasterEnable)
	if cmd[0]&needed == needed {
		return nil // already enabled
	}

	cmd[0] |= needed
	if _, err := f.WriteAt(cmd[:], 0x04); err != nil {
		return fmt.Errorf("cannot write PCI Command Register for %s: %w", bdf, err)
	}

	return nil
}

// WakeToD0 forces the device from D3hot/D3cold to D0 by clearing the
// power state bits in the PM Control/Status register and disabling
// kernel runtime power management.
func WakeToD0(bdf string) error {
	// disable kernel runtime PM so it won't put the device back to sleep
	powerCtrl := filepath.Join(sysfsBase, bdf, "power", "control")
	_ = os.WriteFile(powerCtrl, []byte("on"), 0200)

	ps, err := CheckPowerState(bdf)
	if err != nil || ps == "D0" {
		return nil // already in D0 or can't check
	}

	// find PM capability offset by walking config space cap list
	configPath := filepath.Join(sysfsBase, bdf, "config")
	f, err := os.OpenFile(configPath, os.O_RDWR, 0)
	if err != nil {
		return fmt.Errorf("cannot open config space for %s: %w", bdf, err)
	}
	defer f.Close()

	// read Status register (0x06) to check capabilities list
	var status [2]byte
	if _, err := f.ReadAt(status[:], 0x06); err != nil {
		return fmt.Errorf("cannot read status register: %w", err)
	}
	if status[0]&0x10 == 0 {
		return fmt.Errorf("device %s has no capabilities list", bdf)
	}

	// capabilities pointer at 0x34
	var capPtr [1]byte
	if _, err := f.ReadAt(capPtr[:], 0x34); err != nil {
		return fmt.Errorf("cannot read cap pointer: %w", err)
	}

	offset := int(capPtr[0]) & 0xFC
	for offset >= 0x40 && offset < 0x100 {
		var capHdr [2]byte
		if _, err := f.ReadAt(capHdr[:], int64(offset)); err != nil {
			break
		}
		if capHdr[0] == 0x01 { // Power Management capability
			// PMCSR at cap+4, clear bits 1:0 to force D0
			pmcsrOff := int64(offset + 4)
			var pmcsr [2]byte
			if _, err := f.ReadAt(pmcsr[:], pmcsrOff); err != nil {
				return fmt.Errorf("cannot read PMCSR: %w", err)
			}
			pmcsr[0] &= 0xFC // clear PowerState bits (D0 = 00)
			if _, err := f.WriteAt(pmcsr[:], pmcsrOff); err != nil {
				return fmt.Errorf("cannot write PMCSR: %w", err)
			}
			return nil
		}
		offset = int(capHdr[1]) & 0xFC
		if offset == 0 {
			break
		}
	}

	return fmt.Errorf("PM capability not found for %s", bdf)
}

// UnbindFromVFIO detaches from vfio-pci and re-probes for the original driver.
func UnbindFromVFIO(bdf string) error {
	devPath := filepath.Join(sysfsBase, bdf)
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
	link, err := os.Readlink(filepath.Join(sysfsBase, bdf, "iommu_group"))
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

// QuickStatus returns a short VFIO compatibility label for a device:
// "ready" (bound to vfio-pci), "no-iommu", "group(N)" (shared group), or "ok".
func QuickStatus(bdf string) string {
	if IsBoundToVFIO(bdf) {
		return "ready"
	}

	_, err := GetIOMMUGroup(bdf)
	if err != nil {
		return "no-iommu"
	}

	groupDevs, err := ListIOMMUGroupDevices(bdf)
	if err != nil {
		return "ok"
	}

	peers := 0
	for _, d := range groupDevs {
		if d != bdf {
			peers++
		}
	}

	if peers > 0 {
		return fmt.Sprintf("group(%d)", peers+1)
	}

	return "ok"
}
