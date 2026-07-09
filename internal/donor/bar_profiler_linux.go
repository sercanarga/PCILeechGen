//go:build linux

package donor

import (
	"fmt"
	"os"
	"syscall"
)

func (p *BARProfiler) ProfileBAR(resourcePath string, barIndex, maxSize int) (*BARProfile, error) {
	f, err := os.OpenFile(resourcePath, os.O_RDWR, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to open BAR%d for R/W: %w", barIndex, err)
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		return nil, fmt.Errorf("failed to stat BAR%d: %w", barIndex, err)
	}

	size := int(fi.Size())
	if size == 0 {
		return nil, fmt.Errorf("BAR%d resource file is empty", barIndex)
	}
	if size > maxSize {
		size = maxSize
	}
	pageSize := os.Getpagesize()
	mmapSize := ((size + pageSize - 1) / pageSize) * pageSize

	mapped, err := syscall.Mmap(int(f.Fd()), 0, mmapSize,
		syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
	if err != nil {
		return nil, fmt.Errorf("mmap R/W failed for BAR%d: %w", barIndex, err)
	}
	defer func() {
		_ = syscall.Munmap(mapped)
	}()

	profile := &BARProfile{
		BarIndex: barIndex,
		Size:     size,
	}
	profile.Probes = probeRegisters(mapped, size)
	return profile, nil
}
