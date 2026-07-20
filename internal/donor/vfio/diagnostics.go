package vfio

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

var (
	initMountNamespacePath = "/proc/1/ns/mnt"
	mountInfoPath          = "/proc/self/mountinfo"
	procCmdlinePath        = "/proc/cmdline"
	selfMountNamespacePath = "/proc/self/ns/mnt"
	sysDevBlockBase        = "/sys/dev/block"
)

type BARStatus struct {
	Index      int
	Size       uint64
	Accessible bool
	Error      string
}

// CheckPowerState reads D-state from sysfs (D0/D1/D2/D3hot/D3cold).
func CheckPowerState(bdf string) (string, error) {
	path := filepath.Join(sysfsBase, bdf, "power_state")
	data, err := os.ReadFile(path)
	if err != nil {
		return "unknown", fmt.Errorf("cannot read power state: %w", err)
	}
	return strings.TrimSpace(string(data)), nil
}

// ListIOMMUGroupDevices returns all BDFs sharing the same IOMMU group.
func ListIOMMUGroupDevices(bdf string) ([]string, error) {
	groupLink := filepath.Join(sysfsBase, bdf, "iommu_group", "devices")
	entries, err := os.ReadDir(groupLink)
	if err != nil {
		return nil, fmt.Errorf("cannot list IOMMU group devices: %w", err)
	}

	devices := make([]string, 0, len(entries))
	for _, e := range entries {
		devices = append(devices, e.Name())
	}
	return devices, nil
}

// CheckIOMMUGroupSafe rejects groups with peers still owned by native drivers.
func CheckIOMMUGroupSafe(bdf string) error {
	devices, err := ListIOMMUGroupDevices(bdf)
	if err != nil {
		return fmt.Errorf("cannot verify IOMMU group for %s: %w", bdf, err)
	}

	var conflicts []string
	for _, device := range devices {
		if device == bdf {
			continue
		}
		driver, err := checkedBoundDriver(device)
		if err != nil {
			return err
		}
		if driver != "" && driver != "vfio-pci" {
			conflicts = append(conflicts, fmt.Sprintf("%s (%s)", device, driver))
		}
	}
	if len(conflicts) == 0 {
		return nil
	}
	sort.Strings(conflicts)
	return fmt.Errorf("IOMMU group contains device(s) on native drivers: %s; bind the whole group to vfio-pci or isolate the donor",
		strings.Join(conflicts, ", "))
}

func checkedBoundDriver(bdf string) (string, error) {
	target, err := os.Readlink(filepath.Join(sysfsBase, bdf, "driver"))
	if os.IsNotExist(err) {
		return "", nil
	}
	if err != nil {
		return "", fmt.Errorf("cannot read driver for IOMMU peer %s: %w", bdf, err)
	}
	return filepath.Base(target), nil
}

// CheckSafeToBind rejects a PCI device that backs a mounted filesystem or has
// IOMMU-group peers still owned by native drivers.
func CheckSafeToBind(bdf string) error {
	if err := CheckMountedDeviceSafe(bdf); err != nil {
		return err
	}
	return CheckIOMMUGroupSafe(bdf)
}

// CheckMountedDeviceSafe rejects a PCI device that backs a mounted filesystem.
func CheckMountedDeviceSafe(bdf string) error {
	if err := checkInitialPIDNamespace(); err != nil {
		return err
	}
	if err := checkInitialMountNamespace(selfMountNamespacePath, initMountNamespacePath); err != nil {
		return err
	}

	target, err := filepath.EvalSymlinks(filepath.Join(sysfsBase, bdf))
	if err != nil {
		return fmt.Errorf("cannot resolve PCI device %s: %w", bdf, err)
	}

	data, err := os.ReadFile(mountInfoPath)
	if err != nil {
		return fmt.Errorf("cannot verify mounted filesystems: %w", err)
	}
	rootSeen := false
	for lineNo, line := range strings.Split(string(data), "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) < 5 {
			return fmt.Errorf("cannot parse mountinfo line %d", lineNo+1)
		}
		major, _, err := parseDeviceNumber(fields[2])
		if err != nil {
			return fmt.Errorf("cannot parse mountinfo line %d: %w", lineNo+1, err)
		}
		mountPoint := unescapeMountInfo(fields[4])
		if mountPoint == "/" {
			rootSeen = true
		}
		if major == 0 {
			fsType := mountFilesystemType(fields)
			if isKernelPseudoFilesystem(fsType) {
				continue
			}
			return fmt.Errorf("cannot verify mount %s: filesystem %q uses major 0 without block-device topology; refusing VFIO bind",
				mountPoint, fsType)
		}
		backed, err := blockBackedByTarget(filepath.Join(sysDevBlockBase, fields[2]), target, map[string]bool{})
		if err != nil {
			return fmt.Errorf("cannot inspect mounted device %s at %s: %w", fields[2], mountPoint, err)
		}
		if backed {
			return fmt.Errorf("PCI device %s backs a filesystem mounted at %s; refusing to unbind mounted storage",
				bdf, mountPoint)
		}
	}
	if !rootSeen {
		return fmt.Errorf("cannot verify mounted filesystems: root mount is absent from %s", mountInfoPath)
	}
	return nil
}

func checkInitialMountNamespace(selfPath, initPath string) error {
	self, err := os.Readlink(selfPath)
	if err != nil {
		return fmt.Errorf("cannot verify current mount namespace: %w", err)
	}
	init, err := os.Readlink(initPath)
	if err != nil {
		return fmt.Errorf("cannot verify initial mount namespace: %w", err)
	}
	if self != init {
		return fmt.Errorf("current process is in a private mount namespace; refusing VFIO bind against incomplete host mount data")
	}
	return nil
}

func mountFilesystemType(fields []string) string {
	for i, field := range fields {
		if field == "-" && i+1 < len(fields) {
			return fields[i+1]
		}
	}
	return ""
}
func isKernelPseudoFilesystem(fsType string) bool {
	switch fsType {
	case "autofs", "binfmt_misc", "bpf", "cgroup", "cgroup2", "configfs",
		"debugfs", "devpts", "devtmpfs", "efivarfs", "fusectl", "hugetlbfs",
		"mqueue", "nsfs", "proc", "pstore", "ramfs", "rpc_pipefs", "securityfs",
		"selinuxfs", "sysfs", "tmpfs", "tracefs":
		return true
	default:
		return false
	}
}

func parseDeviceNumber(value string) (int, int, error) {
	majorText, minorText, ok := strings.Cut(value, ":")
	if !ok {
		return 0, 0, fmt.Errorf("invalid device number %q", value)
	}
	major, err := strconv.Atoi(majorText)
	if err != nil || major < 0 {
		return 0, 0, fmt.Errorf("invalid device number %q", value)
	}
	minor, err := strconv.Atoi(minorText)
	if err != nil || minor < 0 {
		return 0, 0, fmt.Errorf("invalid device number %q", value)
	}
	return major, minor, nil
}

func blockBackedByTarget(blockPath, target string, seen map[string]bool) (bool, error) {
	resolved, err := filepath.EvalSymlinks(blockPath)
	if err != nil {
		return false, err
	}
	if seen[resolved] {
		return false, nil
	}
	seen[resolved] = true

	rel, err := filepath.Rel(target, resolved)
	if err == nil && (rel == "." || (rel != ".." && !strings.HasPrefix(rel, ".."+string(filepath.Separator)))) {
		return true, nil
	}

	if _, err := os.Stat(filepath.Join(resolved, "partition")); err == nil {
		return blockBackedByTarget(filepath.Dir(resolved), target, seen)
	} else if !os.IsNotExist(err) {
		return false, err
	}

	if subsystem, ok := nvmeSubsystemDir(resolved); ok {
		multipath := filepath.Join(resolved, "multipath")
		paths, err := os.ReadDir(multipath)
		if err == nil && len(paths) > 0 {
			return blockLinksBackedByTarget(multipath, paths, target, seen)
		}
		if err != nil && !os.IsNotExist(err) {
			return false, fmt.Errorf("cannot resolve NVMe multipath head %s: %w", resolved, err)
		}

		entries, fallbackErr := os.ReadDir(subsystem)
		if fallbackErr != nil {
			return false, fmt.Errorf("cannot resolve NVMe multipath head %s: %w", resolved, fallbackErr)
		}
		var controllers []os.DirEntry
		for _, entry := range entries {
			if isNVMeControllerName(entry.Name()) {
				controllers = append(controllers, entry)
			}
		}
		if len(controllers) == 0 {
			return false, fmt.Errorf("cannot resolve NVMe multipath head %s: no controller paths", resolved)
		}
		return blockLinksBackedByTarget(subsystem, controllers, target, seen)
	}

	slaves, err := os.ReadDir(filepath.Join(resolved, "slaves"))
	if os.IsNotExist(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	for _, slave := range slaves {
		backed, err := blockBackedByTarget(filepath.Join(resolved, "slaves", slave.Name()), target, seen)
		if err != nil {
			return false, err
		}
		if backed {
			return true, nil
		}
	}
	return false, nil
}

func blockLinksBackedByTarget(dir string, entries []os.DirEntry, target string, seen map[string]bool) (bool, error) {
	for _, entry := range entries {
		link := filepath.Join(dir, entry.Name())
		info, err := os.Lstat(link)
		if err != nil {
			return false, err
		}
		if info.Mode()&os.ModeSymlink == 0 {
			return false, fmt.Errorf("block dependency entry is not a symlink: %s", link)
		}
		backed, err := blockBackedByTarget(link, target, seen)
		if err != nil {
			return false, err
		}
		if backed {
			return true, nil
		}
	}
	return false, nil
}

func isNVMeControllerName(name string) bool {
	if !strings.HasPrefix(name, "nvme") {
		return false
	}
	_, err := strconv.Atoi(strings.TrimPrefix(name, "nvme"))
	return err == nil
}

func nvmeSubsystemDir(path string) (string, bool) {
	for parent := filepath.Clean(path); ; parent = filepath.Dir(parent) {
		if strings.HasPrefix(filepath.Base(parent), "nvme-subsys") {
			return parent, true
		}
		next := filepath.Dir(parent)
		if next == parent {
			return "", false
		}
	}
}

func unescapeMountInfo(value string) string {
	replacer := strings.NewReplacer(`\040`, " ", `\011`, "\t", `\134`, `\`)
	return replacer.Replace(value)
}

// CheckLiveEnvironment detects common live-media boot markers without treating
// container overlay filesystems as live USB installations.
func CheckLiveEnvironment() (bool, string, error) {
	cmdline, err := os.ReadFile(procCmdlinePath)
	if err != nil {
		return false, "", fmt.Errorf("cannot read kernel command line: %w", err)
	}
	for _, field := range strings.Fields(string(cmdline)) {
		if field == "boot=casper" || field == "boot=live" {
			return true, field, nil
		}
	}

	data, err := os.ReadFile(mountInfoPath)
	if err != nil {
		return false, "", fmt.Errorf("cannot read mount information: %w", err)
	}
	for _, line := range strings.Split(string(data), "\n") {
		fields := strings.Fields(line)
		if len(fields) < 7 {
			continue
		}
		mountPoint := unescapeMountInfo(fields[4])
		if mountPoint != "/cdrom" && mountPoint != "/run/live/medium" {
			continue
		}
		for i, field := range fields {
			if field == "-" && i+1 < len(fields) &&
				(fields[i+1] == "iso9660" || fields[i+1] == "squashfs") {
				return true, fmt.Sprintf("%s mounted at %s", fields[i+1], mountPoint), nil
			}
		}
	}
	return false, "", nil
}

// CheckBARAccessibility tries to open each BAR resource file.
func CheckBARAccessibility(bdf string) []BARStatus {
	var results []BARStatus
	for i := 0; i < 6; i++ {
		resPath := filepath.Join(sysfsBase, bdf, fmt.Sprintf("resource%d", i))
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
				fmt.Sprintf("shared IOMMU group with %d other device(s): %s - all must be unbound or on vfio-pci",
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
				fmt.Sprintf("%s - device must be in D0 for reliable reads. Try: echo on | sudo tee /sys/bus/pci/devices/%s/power/control", ps, bdf),
			})
		}
	}

	if IsBoundToVFIO(bdf) {
		results = append(results, DiagnosticResult{"Driver", true, "vfio-pci"})
	} else {
		driverLink, err := os.Readlink(filepath.Join(sysfsBase, bdf, "driver"))
		if err == nil {
			drv := filepath.Base(driverLink)
			results = append(results, DiagnosticResult{
				"Driver", false,
				fmt.Sprintf("currently bound to %q - run: pcileechgen build --bdf %s (auto-binds)", drv, bdf),
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
