// Package vfio reads PCI config space and BAR memory via Linux VFIO ioctls.
// Needs IOMMU + vfio-pci driver. Non-Linux builds compile but return errors.
package vfio

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	sysfsBase  = "/sys/bus/pci/devices" // per-device attributes
	pciBusPath = "/sys/bus/pci"         // drivers_probe and drivers/ live here
)

// sysfsProbeTimeout bounds a probe-triggering sysfs write. It cannot interrupt a
// D-state kernel hang (SIGKILL can't either) — it just fails fast instead of hanging.
const sysfsProbeTimeout = 60 * time.Second

const resetSettleDelay = 500 * time.Millisecond

// writeWithDeadline bounds a blocking write. It does NOT interrupt an in-kernel
// probe parked in D state — it only stops the CLI hanging silently.
func writeWithDeadline(write func() error, label string, timeout time.Duration) error {
	type writeResult struct{ err error }
	ch := make(chan writeResult, 1)
	go func() { ch <- writeResult{write()} }()
	select {
	case r := <-ch:
		return r.err
	case <-time.After(timeout):
		return fmt.Errorf("%s did not complete within %s — the in-kernel driver probe/remove is likely blocked in D state and cannot be interrupted from userspace; REBOOT the machine to recover the device", label, timeout)
	}
}

func writeSysfsWithDeadline(path string, data []byte) error {
	return writeWithDeadline(func() error { return os.WriteFile(path, data, 0200) }, path, sysfsProbeTimeout)
}

// ConfigSpaceReachable returns false when config offset 0 reads 0xFFFFFFFF
// (master abort): the function is unreachable and probing it risks a kernel hang.
func ConfigSpaceReachable(bdf string) (bool, error) {
	f, err := os.OpenFile(filepath.Join(sysfsBase, bdf, "config"), os.O_RDONLY, 0)
	if err != nil {
		return false, fmt.Errorf("cannot open config space for %s: %w", bdf, err)
	}
	defer f.Close()
	var vd [4]byte
	if _, err := f.ReadAt(vd[:], 0); err != nil {
		return false, nil
	}
	return vd != [4]byte{0xFF, 0xFF, 0xFF, 0xFF}, nil
}

// ResetDevice issues a Function Level Reset; must run while driverless (after unbind).
func ResetDevice(bdf string) error {
	return writeSysfsWithDeadline(filepath.Join(sysfsBase, bdf, "reset"), []byte("1"))
}

// removeVFIODynamicID clears the vfio-pci dynamic id so it can't race the native
// driver on reprobe. Best-effort: no-op if bound via driver_override only.
func removeVFIODynamicID(bdf string) {
	vendor, err := os.ReadFile(filepath.Join(sysfsBase, bdf, "vendor"))
	if err != nil {
		return
	}
	device, err := os.ReadFile(filepath.Join(sysfsBase, bdf, "device"))
	if err != nil {
		return
	}
	idStr := strings.TrimSpace(string(vendor)) + " " + strings.TrimSpace(string(device))
	_ = os.WriteFile(vfioPCIDriverDir()+"/remove_id", []byte(idStr), 0200)
}

// SetSysfsBase overrides the per-device sysfs path (for testing).
func SetSysfsBase(path string) { sysfsBase = path }

// SetPciBusPath overrides the PCI bus sysfs path for testing.
func SetPciBusPath(path string) { pciBusPath = path }

// driversProbePath returns the drivers_probe sysfs path.
func driversProbePath() string { return pciBusPath + "/drivers_probe" }

// vfioPCIDriverDir returns the path to the vfio-pci driver directory.
func vfioPCIDriverDir() string { return pciBusPath + "/drivers/vfio-pci" }

// ResetSysfsBase restores the default sysfs paths.
func ResetSysfsBase() {
	sysfsBase = "/sys/bus/pci/devices"
	pciBusPath = "/sys/bus/pci"
}

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

// BoundDriver returns the basename of the bound driver symlink (e.g. "vfio-pci",
// "nvme"), or "" if no driver is bound.
func BoundDriver(bdf string) string {
	target, err := os.Readlink(filepath.Join(sysfsBase, bdf, "driver"))
	if err != nil {
		return ""
	}
	return filepath.Base(target)
}

// WaitForNativeDriver polls until a driver other than vfio-pci binds; otherwise
// BAR reads return 0xFF. drivers_probe runs probe synchronously, so this returns
// at once on success and only waits long enough to surface a bind failure.
func WaitForNativeDriver(bdf string, timeout, interval time.Duration) (string, error) {
	deadline := time.Now().Add(timeout)
	var last string
	for {
		if last = BoundDriver(bdf); last != "" && last != "vfio-pci" {
			return last, nil
		}
		if time.Now().After(deadline) || interval <= 0 {
			break
		}
		time.Sleep(interval)
	}
	if last == "vfio-pci" {
		return "", fmt.Errorf("device %s still on vfio-pci after %s (driver_override may be pinned)", bdf, timeout)
	}
	return "", fmt.Errorf("no native driver bound to %s within %s (is the native module loaded? probe may have failed)", bdf, timeout)
}

// WaitForNVMeLive polls <bdf>/nvme/nvmeN/state until "live". The driver symlink
// appears before nvme_probe finishes; "live" means the controller is up and its
// Identify Controller data is available.
func WaitForNVMeLive(bdf string, timeout, interval time.Duration) error {
	nvmeDir := filepath.Join(sysfsBase, bdf, "nvme")
	deadline := time.Now().Add(timeout)
	var lastState string
	for {
		if state, ok := readNVMeControllerState(nvmeDir); ok {
			lastState = state
			switch state {
			case "live":
				return nil
			case "dead", "deleting", "deleting-noio":
				return fmt.Errorf("nvme controller for %s is %q", bdf, state)
			}
		}
		if time.Now().After(deadline) || interval <= 0 {
			break
		}
		time.Sleep(interval)
	}
	if lastState == "" {
		return fmt.Errorf("nvme controller for %s not found within %s", bdf, timeout)
	}
	return fmt.Errorf("nvme controller for %s did not reach live within %s (last state %q)", bdf, timeout, lastState)
}

// readNVMeControllerState reads state from the first nvmeN entry under nvmeDir.
// ok is false when the controller is not registered yet.
func readNVMeControllerState(nvmeDir string) (string, bool) {
	entries, err := os.ReadDir(nvmeDir)
	if err != nil {
		return "", false
	}
	for _, e := range entries {
		if !e.IsDir() || !strings.HasPrefix(e.Name(), "nvme") {
			continue
		}
		data, err := os.ReadFile(filepath.Join(nvmeDir, e.Name(), "state"))
		if err != nil {
			continue
		}
		return strings.TrimSpace(string(data)), true
	}
	return "", false
}

// ListVFIODevices enumerates BDFs currently on the vfio-pci driver.
func ListVFIODevices() ([]string, error) {
	entries, err := os.ReadDir(vfioPCIDriverDir())
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
	if err := CheckSafeToBind(bdf); err != nil {
		return fmt.Errorf("unsafe VFIO bind: %w", err)
	}

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
		slog.Info("vfio: unbinding from current driver", "bdf", bdf, "driver", filepath.Base(driverLink))
		unbindPath := filepath.Join(devPath, "driver", "unbind")
		if writeErr := writeSysfsWithDeadline(unbindPath, []byte(bdf)); writeErr != nil {
			return fmt.Errorf("failed to unbind from current driver: %w", writeErr)
		}
	}

	slog.Info("vfio: setting driver_override to vfio-pci", "bdf", bdf)
	overridePath := filepath.Join(devPath, "driver_override")
	if writeErr := os.WriteFile(overridePath, []byte("vfio-pci"), 0200); writeErr != nil {
		return fmt.Errorf("failed to set driver override: %w", writeErr)
	}

	idStr := fmt.Sprintf("%s %s", vendor, device)
	_ = os.WriteFile(vfioPCIDriverDir()+"/new_id", []byte(idStr), 0200)

	slog.Info("vfio: writing drivers_probe (vfio-pci probe runs synchronously in this write)", "bdf", bdf)
	if writeErr := writeSysfsWithDeadline(driversProbePath(), []byte(bdf)); writeErr != nil {
		return fmt.Errorf("failed to probe device: %w", writeErr)
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
	// Disable kernel runtime PM so it cannot put the device back to sleep.
	powerCtrl := filepath.Join(sysfsBase, bdf, "power", "control")
	if err := os.WriteFile(powerCtrl, []byte("on"), 0200); err != nil {
		return fmt.Errorf("cannot disable runtime power management for %s: %w", bdf, err)
	}

	ps, err := CheckPowerState(bdf)
	if err != nil {
		return err
	}
	if ps == "D0" {
		return nil
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

// UnbindFromVFIO detaches from vfio-pci, resets, and re-probes the native driver.
func UnbindFromVFIO(bdf string) error {
	devPath := filepath.Join(sysfsBase, bdf)

	slog.Info("vfio: clearing driver_override", "bdf", bdf)
	// Newline, not empty: kernfs ignores zero-length writes (echo "" > driver_override).
	if err := os.WriteFile(filepath.Join(devPath, "driver_override"), []byte("\n"), 0200); err != nil {
		return fmt.Errorf("failed to clear driver_override: %w", err)
	}

	slog.Info("vfio: unbinding from vfio-pci", "bdf", bdf)
	if err := writeSysfsWithDeadline(vfioPCIDriverDir()+"/unbind", []byte(bdf)); err != nil {
		return fmt.Errorf("failed to unbind from vfio-pci: %w", err)
	}

	removeVFIODynamicID(bdf)

	// FLR before reprobe so nvme_probe sees a clean controller and can't wedge.
	if err := ResetDevice(bdf); err != nil {
		slog.Warn("vfio: pre-probe reset failed; continuing (device may lack FLR)", "bdf", bdf, "error", err)
	} else {
		slog.Info("vfio: issued pre-probe function reset", "bdf", bdf)
		time.Sleep(resetSettleDelay)
	}

	slog.Info("vfio: writing drivers_probe (native probe runs synchronously)", "bdf", bdf)
	if err := writeSysfsWithDeadline(driversProbePath(), []byte(bdf)); err != nil {
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
