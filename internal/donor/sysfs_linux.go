//go:build linux

package donor

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
	"unsafe"

	"github.com/sercanarga/pcileechgen/internal/pci"
)

func (sr *SysfsReader) readBARViaMmap(f *os.File, size int) ([]byte, error) {
	pageSize := os.Getpagesize()
	mmapSize := ((size + pageSize - 1) / pageSize) * pageSize

	mapped, err := syscall.Mmap(int(f.Fd()), 0, mmapSize, syscall.PROT_READ, syscall.MAP_SHARED)
	if err != nil {
		return nil, fmt.Errorf("mmap failed: %w", err)
	}

	data := make([]byte, size)
	copy(data, mapped)

	if err := syscall.Munmap(mapped); err != nil {
		slog.Warn("munmap failed", "error", err)
	}
	return data, nil
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

const nvmeIOCAdminCmd = uintptr(0xC0484E41)

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

	devPath := "/dev/" + ctrlName
	f, err := os.OpenFile(devPath, os.O_RDWR, 0)
	if err != nil {
		return nil, fmt.Errorf("open %s: %w", devPath, err)
	}
	defer f.Close()

	buf := make([]byte, 4096)
	cmd := nvmeAdminCmd{
		opcode:    0x06,
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
	runtime.KeepAlive(buf)
	if errno != 0 {
		return nil, fmt.Errorf("NVMe ioctl failed (cns=%d): %w", cns, errno)
	}

	return buf, nil
}
