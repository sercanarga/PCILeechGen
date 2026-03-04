// VFIO device reader — uses ioctl to read config space and BAR memory.

package vfio

import (
	"fmt"
	"os"
	"unsafe"

	"golang.org/x/sys/unix"
)

// ioctl numbers from linux/vfio.h
const (
	vfioGetAPIVersion       = 0x3B64 // VFIO_GET_API_VERSION
	vfioCheckExtension      = 0x3B65 // VFIO_CHECK_EXTENSION
	vfioSetIOMMU            = 0x3B66 // VFIO_SET_IOMMU
	vfioGroupGetStatus      = 0x3B67 // VFIO_GROUP_GET_STATUS
	vfioGroupSetContainer   = 0x3B68 // VFIO_GROUP_SET_CONTAINER
	vfioGroupGetDeviceFD    = 0x3B6A // VFIO_GROUP_GET_DEVICE_FD
	vfioDeviceGetRegionInfo = 0x3B6C // VFIO_DEVICE_GET_REGION_INFO

	vfioType1IOMMU = 1

	// PCI region indices
	vfioPCIBAR0   = 0
	vfioPCIConfig = 7

	// Region flags
	vfioRegionFlagRead  = 1
	vfioRegionFlagWrite = 2

	configSpaceMaxSize = 4096
	barDumpMaxSize     = 4096 // clamp to FPGA BRAM
)

// vfioRegionInfo matches struct vfio_region_info.
type vfioRegionInfo struct {
	Argsz  uint32
	Flags  uint32
	Index  uint32
	Cap    uint32
	Size   uint64
	Offset uint64
}

// Collect grabs config space + BAR memory from a VFIO-bound device.
func Collect(bdf string) (*DeviceDump, error) {
	// find IOMMU group
	groupPath, err := ResolveIOMMUGroup(bdf)
	if err != nil {
		return nil, err
	}

	// open container fd
	containerFD, err := unix.Open("/dev/vfio/vfio", unix.O_RDWR, 0)
	if err != nil {
		return nil, fmt.Errorf("cannot open /dev/vfio/vfio: %w", err)
	}
	defer unix.Close(containerFD)

	// sanity check API
	ver, _, errno := unix.Syscall(unix.SYS_IOCTL, uintptr(containerFD),
		vfioGetAPIVersion, 0)
	if errno != 0 || ver == 0 {
		return nil, fmt.Errorf("VFIO API version check failed: %v", errno)
	}

	// open group
	groupFD, err := unix.Open(groupPath, unix.O_RDWR, 0)
	if err != nil {
		return nil, fmt.Errorf("cannot open VFIO group %s: %w", groupPath, err)
	}
	defer unix.Close(groupFD)

	// attach group to container
	_, _, errno = unix.Syscall(unix.SYS_IOCTL, uintptr(groupFD),
		vfioGroupSetContainer, uintptr(unsafe.Pointer(&containerFD)))
	if errno != 0 {
		return nil, fmt.Errorf("VFIO_GROUP_SET_CONTAINER: %v", errno)
	}

	// set IOMMU type
	iommuType := vfioType1IOMMU
	_, _, errno = unix.Syscall(unix.SYS_IOCTL, uintptr(containerFD),
		vfioSetIOMMU, uintptr(iommuType))
	if errno != 0 {
		return nil, fmt.Errorf("VFIO_SET_IOMMU: %v", errno)
	}

	// get device fd
	bdfBytes := append([]byte(bdf), 0) // null-terminated
	deviceFD, _, errno := unix.Syscall(unix.SYS_IOCTL, uintptr(groupFD),
		vfioGroupGetDeviceFD, uintptr(unsafe.Pointer(&bdfBytes[0])))
	if errno != 0 || int(deviceFD) < 0 {
		return nil, fmt.Errorf("cannot get VFIO device FD for %s: %v", bdf, errno)
	}
	defer unix.Close(int(deviceFD))

	dump := &DeviceDump{
		BDF:         bdf,
		BARContents: make(map[int][]byte),
	}

	// config space
	cs, csSize, err := readRegion(int(deviceFD), vfioPCIConfig, configSpaceMaxSize)
	if err != nil {
		return nil, fmt.Errorf("config space read failed: %w", err)
	}
	dump.ConfigSpace = cs
	dump.ConfigSpaceSize = csSize

	// BAR regions
	for i := 0; i < 6; i++ {
		info, err := getRegionInfo(int(deviceFD), vfioPCIBAR0+i)
		if err != nil {
			dump.BARInfo = append(dump.BARInfo, BARRegion{Index: i})
			continue
		}
		dump.BARInfo = append(dump.BARInfo, BARRegion{
			Index: i,
			Size:  info.Size,
			Flags: info.Flags,
		})

		// read if readable and nonzero
		if info.Size > 0 && info.Flags&vfioRegionFlagRead != 0 {
			barData, _, err := readRegion(int(deviceFD), vfioPCIBAR0+i, barDumpMaxSize)
			if err == nil && len(barData) > 0 {
				dump.BARContents[i] = barData
			}
		}
	}

	return dump, nil
}

// getRegionInfo queries region metadata via VFIO ioctl.
func getRegionInfo(deviceFD, index int) (*vfioRegionInfo, error) {
	info := vfioRegionInfo{
		Argsz: uint32(unsafe.Sizeof(vfioRegionInfo{})),
		Index: uint32(index),
	}
	_, _, errno := unix.Syscall(unix.SYS_IOCTL, uintptr(deviceFD),
		vfioDeviceGetRegionInfo, uintptr(unsafe.Pointer(&info)))
	if errno != 0 {
		return nil, fmt.Errorf("VFIO_DEVICE_GET_REGION_INFO index %d: %v", index, errno)
	}
	return &info, nil
}

// readRegion does a pread on a VFIO region, capped at maxSize.
func readRegion(deviceFD, index, maxSize int) ([]byte, int, error) {
	info, err := getRegionInfo(deviceFD, index)
	if err != nil {
		return nil, 0, err
	}
	if info.Size == 0 {
		return nil, 0, nil
	}

	readSize := int(info.Size)
	if readSize > maxSize {
		readSize = maxSize
	}

	buf := make([]byte, readSize)
	n, err := unix.Pread(deviceFD, buf, int64(info.Offset))
	if err != nil {
		return nil, 0, fmt.Errorf("pread region %d: %w", index, err)
	}

	return buf[:n], n, nil
}

// CollectToFile dumps device data to a JSON file.
func CollectToFile(bdf, outputPath string) error {
	dump, err := Collect(bdf)
	if err != nil {
		return err
	}

	data, err := dump.ToJSON()
	if err != nil {
		return fmt.Errorf("JSON marshal: %w", err)
	}

	return os.WriteFile(outputPath, data, 0644)
}
