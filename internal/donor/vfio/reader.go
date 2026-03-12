// Package vfio reads PCI config space and BAR memory from VFIO-bound devices.

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
	vfioSetIOMMU            = 0x3B66 // VFIO_SET_IOMMU
	vfioGroupSetContainer   = 0x3B68 // VFIO_GROUP_SET_CONTAINER
	vfioGroupGetDeviceFD    = 0x3B6A // VFIO_GROUP_GET_DEVICE_FD
	vfioDeviceGetRegionInfo = 0x3B6C // VFIO_DEVICE_GET_REGION_INFO

	vfioType1IOMMU = 1

	// PCI region indices
	vfioPCIBAR0   = 0
	vfioPCIConfig = 7

	// Region flags
	vfioRegionFlagRead = 1

	configSpaceMaxSize = 4096
	barDumpMaxSize     = 4096
)

// vfioRegionInfo matches struct vfio_region_info from linux/vfio.h.
type vfioRegionInfo struct {
	Argsz  uint32
	Flags  uint32
	Index  uint32
	Cap    uint32
	Size   uint64
	Offset uint64
}

// vfioSession holds open file descriptors for a VFIO session.
type vfioSession struct {
	containerFD int
	groupFD     int
	deviceFD    int
}

// close releases all open file descriptors.
func (s *vfioSession) close() {
	if s.deviceFD >= 0 {
		unix.Close(s.deviceFD)
	}
	if s.groupFD >= 0 {
		unix.Close(s.groupFD)
	}
	if s.containerFD >= 0 {
		unix.Close(s.containerFD)
	}
}

// openSession sets up VFIO container, group, and device for the given BDF.
func openSession(bdf string) (*vfioSession, error) {
	s := &vfioSession{containerFD: -1, groupFD: -1, deviceFD: -1}

	groupPath, err := ResolveIOMMUGroup(bdf)
	if err != nil {
		return nil, err
	}

	// open container
	s.containerFD, err = unix.Open("/dev/vfio/vfio", unix.O_RDWR, 0)
	if err != nil {
		return nil, fmt.Errorf("cannot open /dev/vfio/vfio: %w", err)
	}

	// check API version (VFIO_API_VERSION is 0)
	if err := vfioIoctl(s.containerFD, vfioGetAPIVersion, 0); err != nil {
		s.close()
		return nil, fmt.Errorf("VFIO API check failed: %w", err)
	}

	// open group
	s.groupFD, err = unix.Open(groupPath, unix.O_RDWR, 0)
	if err != nil {
		s.close()
		return nil, fmt.Errorf("cannot open VFIO group %s: %w", groupPath, err)
	}

	// attach group to container
	if err := vfioIoctl(s.groupFD, vfioGroupSetContainer, uintptr(unsafe.Pointer(&s.containerFD))); err != nil {
		s.close()
		return nil, fmt.Errorf("VFIO_GROUP_SET_CONTAINER: %w", err)
	}

	// set IOMMU type
	if err := vfioIoctl(s.containerFD, vfioSetIOMMU, vfioType1IOMMU); err != nil {
		s.close()
		return nil, fmt.Errorf("VFIO_SET_IOMMU: %w", err)
	}

	// get device fd
	bdfBytes := append([]byte(bdf), 0)
	fd, _, errno := unix.Syscall(unix.SYS_IOCTL, uintptr(s.groupFD),
		vfioGroupGetDeviceFD, uintptr(unsafe.Pointer(&bdfBytes[0])))
	if errno != 0 || int(fd) < 0 {
		s.close()
		return nil, fmt.Errorf("cannot get VFIO device FD for %s: %v", bdf, errno)
	}
	s.deviceFD = int(fd)

	return s, nil
}

// vfioIoctl is a thin wrapper for simple ioctl calls.
func vfioIoctl(fd int, request uintptr, arg uintptr) error {
	_, _, errno := unix.Syscall(unix.SYS_IOCTL, uintptr(fd), request, arg)
	if errno != 0 {
		return fmt.Errorf("ioctl 0x%x: %v", request, errno)
	}
	return nil
}

// Collect grabs config space and BAR memory from a VFIO-bound device.
func Collect(bdf string) (*DeviceDump, error) {
	session, err := openSession(bdf)
	if err != nil {
		return nil, err
	}
	defer session.close()

	dump := &DeviceDump{
		BDF:         bdf,
		BARContents: make(map[int][]byte),
	}

	// config space (pread only, mmap fallback not needed)
	cs, csSize, err := readRegionPread(session.deviceFD, vfioPCIConfig, configSpaceMaxSize)
	if err != nil {
		return nil, fmt.Errorf("config space read failed: %w", err)
	}
	dump.ConfigSpace = cs
	dump.ConfigSpaceSize = csSize

	// BAR regions
	for i := 0; i < 6; i++ {
		info, err := getRegionInfo(session.deviceFD, vfioPCIBAR0+i)
		if err != nil {
			dump.BARInfo = append(dump.BARInfo, BARRegion{Index: i})
			continue
		}
		dump.BARInfo = append(dump.BARInfo, BARRegion{
			Index: i,
			Size:  info.Size,
			Flags: info.Flags,
		})

		if info.Size > 0 && info.Flags&vfioRegionFlagRead != 0 {
			barData, _, err := readBARRegion(session.deviceFD, vfioPCIBAR0+i, barDumpMaxSize)
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

// regionSize caps the read size to maxSize.
func regionSize(info *vfioRegionInfo, maxSize int) int {
	s := int(info.Size)
	if s > maxSize {
		s = maxSize
	}
	return s
}

// readRegionPread reads a VFIO region via pread (config space).
func readRegionPread(deviceFD, index, maxSize int) ([]byte, int, error) {
	info, err := getRegionInfo(deviceFD, index)
	if err != nil {
		return nil, 0, err
	}
	if info.Size == 0 {
		return nil, 0, nil
	}

	readSize := regionSize(info, maxSize)
	buf := make([]byte, readSize)
	n, err := unix.Pread(deviceFD, buf, int64(info.Offset))
	if err != nil {
		return nil, 0, fmt.Errorf("pread region %d: %w", index, err)
	}
	return buf[:n], n, nil
}

// readBARRegion reads a BAR via pread, falls back to mmap on 0xFF.
func readBARRegion(deviceFD, index, maxSize int) ([]byte, int, error) {
	info, err := getRegionInfo(deviceFD, index)
	if err != nil {
		return nil, 0, err
	}
	if info.Size == 0 {
		return nil, 0, nil
	}

	readSize := regionSize(info, maxSize)
	buf := make([]byte, readSize)
	n, err := unix.Pread(deviceFD, buf, int64(info.Offset))
	if err != nil {
		return nil, 0, fmt.Errorf("pread region %d: %w", index, err)
	}

	if isAllFF(buf[:n]) {
		data, mmapErr := mmapRegion(deviceFD, info, readSize)
		if mmapErr == nil && !isAllFF(data) {
			return data, len(data), nil
		}
	}

	return buf[:n], n, nil
}

// mmapRegion reads a VFIO region via mmap.
func mmapRegion(deviceFD int, info *vfioRegionInfo, size int) ([]byte, error) {
	pageSize := unix.Getpagesize()
	mmapSize := ((size + pageSize - 1) / pageSize) * pageSize

	mapped, err := unix.Mmap(deviceFD, int64(info.Offset), mmapSize,
		unix.PROT_READ, unix.MAP_SHARED)
	if err != nil {
		return nil, fmt.Errorf("mmap region %d: %w", info.Index, err)
	}

	data := make([]byte, size)
	copy(data, mapped)
	_ = unix.Munmap(mapped)
	return data, nil
}

// isAllFF returns true when every byte is 0xFF.
func isAllFF(data []byte) bool {
	for _, b := range data {
		if b != 0xFF {
			return false
		}
	}
	return len(data) > 0
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
