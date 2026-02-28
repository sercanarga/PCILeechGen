package donor

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

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

	var devices []pci.PCIDevice
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

	readSize := int(fi.Size())
	if readSize == 0 {
		return nil, fmt.Errorf("BAR%d resource file is empty", barIndex)
	}
	if readSize > maxSize {
		readSize = maxSize
	}

	data := make([]byte, readSize)
	n, err := f.Read(data)
	if err != nil {
		return nil, fmt.Errorf("failed to read BAR%d content: %w", barIndex, err)
	}

	return data[:n], nil
}

// readHex16 reads a hex value from a sysfs file and returns it as uint16.
func (sr *SysfsReader) readHex16(devPath, name string) (uint16, error) {
	data, err := os.ReadFile(filepath.Join(devPath, name))
	if err != nil {
		return 0, err
	}
	val, err := strconv.ParseUint(strings.TrimSpace(string(data)), 0, 16)
	if err != nil {
		return 0, err
	}
	return uint16(val), nil
}

// readHex32 reads a hex value from a sysfs file and returns it as uint32.
func (sr *SysfsReader) readHex32(devPath, name string) (uint32, error) {
	data, err := os.ReadFile(filepath.Join(devPath, name))
	if err != nil {
		return 0, err
	}
	val, err := strconv.ParseUint(strings.TrimSpace(string(data)), 0, 32)
	if err != nil {
		return 0, err
	}
	return uint32(val), nil
}

// readHex8 reads a hex value from a sysfs file and returns it as uint8.
func (sr *SysfsReader) readHex8(devPath, name string) (uint8, error) {
	data, err := os.ReadFile(filepath.Join(devPath, name))
	if err != nil {
		return 0, err
	}
	val, err := strconv.ParseUint(strings.TrimSpace(string(data)), 0, 8)
	if err != nil {
		return 0, err
	}
	return uint8(val), nil
}
