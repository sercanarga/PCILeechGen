package donor

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"unsafe"

	"github.com/sercanarga/pcileechgen/internal/pci"
)

const sysfsBasePath = "/sys/bus/pci/devices"

// SysfsReader reads PCI device information from Linux sysfs.
type SysfsReader struct {
	basePath string
}

// NewSysfsReader creates a new SysfsReader with default sysfs path.
func NewSysfsReader() *SysfsReader {
	return &SysfsReader{basePath: sysfsBasePath}
}

// NewSysfsReaderWithPath creates a new SysfsReader with a custom base path (for testing).
func NewSysfsReaderWithPath(basePath string) *SysfsReader {
	return &SysfsReader{basePath: basePath}
}

// ScanDevices returns a list of all PCI devices found in sysfs.
func (sr *SysfsReader) ScanDevices() ([]pci.PCIDevice, error) {
	entries, err := os.ReadDir(sr.basePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read sysfs: %w", err)
	}

	devices := make([]pci.PCIDevice, 0, len(entries))
	for _, entry := range entries {
		// sysfs entries are symlinks, not plain directories
		name := entry.Name()
		fullPath := filepath.Join(sr.basePath, name)

		fi, err := os.Stat(fullPath) // follows symlinks
		if err != nil || !fi.IsDir() {
			continue
		}

		bdf, err := pci.ParseBDF(name)
		if err != nil {
			continue
		}

		dev, err := sr.ReadDeviceInfo(bdf)
		if err != nil {
			continue
		}
		devices = append(devices, *dev)
	}

	return devices, nil
}

// ReadDeviceInfo reads basic device information from sysfs.
func (sr *SysfsReader) ReadDeviceInfo(bdf pci.BDF) (*pci.PCIDevice, error) {
	devPath := filepath.Join(sr.basePath, bdf.String())

	dev := &pci.PCIDevice{BDF: bdf}

	var err error
	dev.VendorID, err = sr.readHex16(devPath, "vendor")
	if err != nil {
		return nil, fmt.Errorf("failed to read vendor ID: %w", err)
	}

	dev.DeviceID, err = sr.readHex16(devPath, "device")
	if err != nil {
		return nil, fmt.Errorf("failed to read device ID: %w", err)
	}

	dev.SubsysVendorID, _ = sr.readHex16(devPath, "subsystem_vendor")
	dev.SubsysDeviceID, _ = sr.readHex16(devPath, "subsystem_device")

	classCode, err := sr.readHex32(devPath, "class")
	if err == nil {
		dev.ClassCode = classCode & 0xFFFFFF
	}

	rev, _ := sr.readHex8(devPath, "revision")
	dev.RevisionID = rev

	// Read header type from config space
	configPath := filepath.Join(devPath, "config")
	configData, err := os.ReadFile(configPath)
	if err == nil && len(configData) > 0x0E {
		dev.HeaderType = configData[0x0E] & 0x7F // mask multi-function bit
	}

	// Read driver symlink
	driverLink, err := os.Readlink(filepath.Join(devPath, "driver"))
	if err == nil {
		dev.Driver = filepath.Base(driverLink)
	}

	// Read IOMMU group
	iommuLink, err := os.Readlink(filepath.Join(devPath, "iommu_group"))
	if err == nil {
		groupStr := filepath.Base(iommuLink)
		if g, err := strconv.Atoi(groupStr); err == nil {
			dev.IOMMUGroup = g
		}
	}

	return dev, nil
}

// ReadConfigSpace reads the full PCI config space from sysfs.
func (sr *SysfsReader) ReadConfigSpace(bdf pci.BDF) (*pci.ConfigSpace, error) {
	configPath := filepath.Join(sr.basePath, bdf.String(), "config")

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config space: %w", err)
	}

	cs := pci.NewConfigSpaceFromBytes(data)
	return cs, nil
}

// ReadResourceFile reads BAR information from the sysfs resource file.
func (sr *SysfsReader) ReadResourceFile(bdf pci.BDF) ([]pci.BAR, error) {
	resourcePath := filepath.Join(sr.basePath, bdf.String(), "resource")

	f, err := os.Open(resourcePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read resource file: %w", err)
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return pci.ParseBARsFromSysfsResource(lines), nil
}

// ReadBARContent reads the memory contents of a BAR from sysfs resource file.
// The resource{N} files in sysfs provide direct access to the BAR's memory region.
// maxSize limits the read to prevent exceeding FPGA BRAM capacity.
//
// When a device is bound to vfio-pci, direct read() on resource files fails with
// "input/output error". In this case, mmap() is used instead, which works because
// the kernel exposes BAR memory via mmap even under vfio-pci.
func (sr *SysfsReader) ReadBARContent(bdf pci.BDF, barIndex int, maxSize int) ([]byte, error) {
	resourcePath := filepath.Join(sr.basePath, bdf.String(), fmt.Sprintf("resource%d", barIndex))

	f, err := os.Open(resourcePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open BAR%d resource file: %w", barIndex, err)
	}
	defer f.Close()

	// Check file size
	fi, err := f.Stat()
	if err != nil {
		return nil, fmt.Errorf("failed to stat BAR%d resource file: %w", barIndex, err)
	}

	fileSize := int(fi.Size())
	if fileSize == 0 {
		return nil, fmt.Errorf("BAR%d resource file is empty", barIndex)
	}

	readSize := fileSize
	if readSize > maxSize {
		readSize = maxSize
	}

	// Hard safety cap: some devices report very large BAR sizes in sysfs (hundreds
	// of MB+). Reading the full region via MMIO is extremely slow and makes the
	// tool appear to freeze. 64KB is sufficient for the register snapshot used by
	// the emulation models.
	if readSize > 65536 {
		readSize = 65536
	}

	// Try mmap first - works with vfio-pci bound devices
	data, err := sr.readBARViaMmap(f, readSize)
	if err == nil {
		return data, nil
	}

	// Fallback to read() - works for regular files (e.g. in tests)
	return sr.readBARViaRead(f, barIndex, readSize)
}

// readBARViaMmap reads BAR contents via mmap (preferred for vfio-pci).
func (sr *SysfsReader) readBARViaMmap(f *os.File, size int) ([]byte, error) {
	// page-align for mmap
	pageSize := os.Getpagesize()
	mmapSize := ((size + pageSize - 1) / pageSize) * pageSize

	mapped, err := syscall.Mmap(int(f.Fd()), 0, mmapSize, syscall.PROT_READ, syscall.MAP_SHARED)
	if err != nil {
		return nil, fmt.Errorf("mmap failed: %w", err)
	}

	// copy requested size only
	data := make([]byte, size)
	copy(data, mapped)

	if err := syscall.Munmap(mapped); err != nil {
		slog.Warn("munmap failed", "error", err)
	}
	return data, nil
}

// readBARViaRead reads the BAR resource file using standard read() syscall.
func (sr *SysfsReader) readBARViaRead(f *os.File, barIndex int, size int) ([]byte, error) {
	data := make([]byte, size)
	n, err := f.Read(data)
	if err != nil {
		return nil, fmt.Errorf("failed to read BAR%d content: %w", barIndex, err)
	}
	return data[:n], nil
}

func (sr *SysfsReader) readSysfsHex(devPath, name string, bitSize int) (uint64, error) {
	data, err := os.ReadFile(filepath.Join(devPath, name))
	if err != nil {
		return 0, err
	}
	return strconv.ParseUint(strings.TrimSpace(string(data)), 0, bitSize)
}

func (sr *SysfsReader) readSysfsString(devPath, name string) (string, error) {
	data, err := os.ReadFile(filepath.Join(devPath, name))
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(data)), nil
}

func (sr *SysfsReader) readHex16(devPath, name string) (uint16, error) {
	v, err := sr.readSysfsHex(devPath, name, 16)
	return uint16(v), err
}

func (sr *SysfsReader) readHex32(devPath, name string) (uint32, error) {
	v, err := sr.readSysfsHex(devPath, name, 32)
	return uint32(v), err
}

func (sr *SysfsReader) readHex8(devPath, name string) (uint8, error) {
	v, err := sr.readSysfsHex(devPath, name, 8)
	return uint8(v), err
}

// ReadNVMeIdentity reads NVMe controller strings from the nvme subdir of the
// PCI device. The subdir exists only while the native nvme driver is bound;
// it is absent under vfio-pci.
func (sr *SysfsReader) ReadNVMeIdentity(bdf pci.BDF) (*NVMeIdentity, error) {
	nvmeDir := filepath.Join(sr.basePath, bdf.String(), "nvme")
	entries, err := os.ReadDir(nvmeDir)
	if err != nil {
		return nil, fmt.Errorf("nvme driver not bound (no %s): %w", nvmeDir, err)
	}
	ctrlPath := ""
	for _, e := range entries {
		if e.IsDir() && strings.HasPrefix(e.Name(), "nvme") {
			ctrlPath = filepath.Join(nvmeDir, e.Name())
			break
		}
	}
	if ctrlPath == "" {
		return nil, fmt.Errorf("no nvme controller under %s", nvmeDir)
	}
	model, err := sr.readSysfsString(ctrlPath, "model")
	if err != nil {
		return nil, fmt.Errorf("read nvme model: %w", err)
	}
	serial, err := sr.readSysfsString(ctrlPath, "serial")
	if err != nil {
		return nil, fmt.Errorf("read nvme serial: %w", err)
	}
	fw, err := sr.readSysfsString(ctrlPath, "firmware_rev")
	if err != nil {
		return nil, fmt.Errorf("read nvme firmware_rev: %w", err)
	}
	return &NVMeIdentity{Serial: serial, Model: model, FWRev: fw}, nil
}

type nvmeAdminCmd struct {
	opcode      uint8
	flags       uint8
	rsvd1       uint16
	nsid        uint32
	cdw2        uint32
	cdw3        uint32
	metadata    uint64
	addr        uint64
	metadataLen uint32
	dataLen     uint32
	cdw10       uint32
	cdw11       uint32
	cdw12       uint32
	cdw13       uint32
	cdw14       uint32
	cdw15       uint32
	timeoutMs   uint32
	result      uint32
}

// NVME_IOCTL_ADMIN_CMD = _IOWR('N', 0x41, struct nvme_admin_cmd) = 0xC0484E41.
const nvmeIOCAdminCmd = uintptr(0xC0484E41)

// ReadNVMeRawIdentify reads raw 4KB Identify (cns: 0=Namespace, 1=Controller) via /dev/nvme{N} ioctl.
func (sr *SysfsReader) ReadNVMeRawIdentify(bdf pci.BDF, cns uint32, nsid uint32) ([]byte, error) {
	nvmeDir := filepath.Join(sr.basePath, bdf.String(), "nvme")
	entries, err := os.ReadDir(nvmeDir)
	if err != nil {
		return nil, fmt.Errorf("nvme driver not bound: %w", err)
	}

	ctrlName := ""
	for _, e := range entries {
		if e.IsDir() && strings.HasPrefix(e.Name(), "nvme") {
			ctrlName = e.Name()
			break
		}
	}
	if ctrlName == "" {
		return nil, fmt.Errorf("no nvme controller under %s", nvmeDir)
	}

	// O_RDWR: ioctl is _IOWR.
	devPath := "/dev/" + ctrlName
	f, err := os.OpenFile(devPath, os.O_RDWR, 0)
	if err != nil {
		return nil, fmt.Errorf("open %s: %w", devPath, err)
	}
	defer f.Close()

	buf := make([]byte, 4096)
	cmd := nvmeAdminCmd{
		opcode:    0x06, // Admin Identify
		nsid:      nsid,
		addr:      uint64(uintptr(unsafe.Pointer(&buf[0]))),
		dataLen:   4096,
		cdw10:     cns,
		timeoutMs: 5000,
	}

	_, _, errno := syscall.Syscall(
		syscall.SYS_IOCTL,
		f.Fd(),
		nvmeIOCAdminCmd,
		uintptr(unsafe.Pointer(&cmd)),
	)
	// KeepAlive: cmd.addr holds buf as a uintptr (GC-untracked).
	runtime.KeepAlive(buf)
	if errno != 0 {
		return nil, fmt.Errorf("NVMe ioctl failed (cns=%d): %w", cns, errno)
	}

	return buf, nil
}
