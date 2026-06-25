package donor

import (
	"encoding/binary"
	"fmt"
	"os"
	"syscall"
)

// BARProbeResult is one 4-byte register's probe output.
type BARProbeResult struct {
	Offset    uint32 `json:"offset"`
	Original  uint32 `json:"original"`
	RWMask    uint32 `json:"rw_mask"`    // 1 = writable bit
	RW1CMask  uint32 `json:"rw1c_mask"`  // bits that clear when written as 1
	RW0CMask  uint32 `json:"rw0c_mask"`
	MaybeRW1C bool   `json:"maybe_rw1c"` // write-1-to-clear suspect
	MaybeRW0C bool   `json:"maybe_rw0c"`
}

// BARProfile is the full probe output for one BAR.
type BARProfile struct {
	BarIndex int              `json:"bar_index"`
	Size     int              `json:"size"`
	Probes   []BARProbeResult `json:"probes"`
}

// BARProfiler probes BAR registers via mmap to find RW/RO/RW1C bits.
type BARProfiler struct{}

func NewBARProfiler() *BARProfiler { return &BARProfiler{} }

// ProfileBAR mmaps a sysfs resource file R/W, writes test patterns to
// each register, reads back, and restores the original value.
// Returns a per-register RW mask and RW1C flag.
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
	// mmap needs page-aligned size
	pageSize := os.Getpagesize()
	mmapSize := ((size + pageSize - 1) / pageSize) * pageSize

	mapped, err := syscall.Mmap(int(f.Fd()), 0, mmapSize,
		syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
	if err != nil {
		return nil, fmt.Errorf("mmap R/W failed for BAR%d: %w", barIndex, err)
	}
	defer syscall.Munmap(mapped)

	profile := &BARProfile{
		BarIndex: barIndex,
		Size:     size,
	}

	profile.Probes = probeRegisters(mapped, size)

	return profile, nil
}

// ProfileBARFromBuffer runs probing against an in-memory buffer (for tests).
func ProfileBARFromBuffer(buf []byte, barIndex int) *BARProfile {
	profile := &BARProfile{
		BarIndex: barIndex,
		Size:     len(buf),
	}
	profile.Probes = probeRegisters(buf, len(buf))
	return profile
}

// probeRegisters walks every DWORD in the region.
func probeRegisters(mem []byte, size int) []BARProbeResult {
	numRegs := size / 4
	probes := make([]BARProbeResult, 0, numRegs)

	for i := 0; i < numRegs; i++ {
		off := i * 4
		result := probeOneRegister(mem, uint32(off))
		probes = append(probes, result)
	}

	return probes
}

// probeOneRegister does the write-readback dance on one DWORD.
func probeOneRegister(mem []byte, offset uint32) BARProbeResult {
	off := int(offset)

	// snapshot
	original := binary.LittleEndian.Uint32(mem[off : off+4])

	// all-ones
	binary.LittleEndian.PutUint32(mem[off:off+4], 0xFFFFFFFF)
	allOnes := binary.LittleEndian.Uint32(mem[off : off+4])

	// all-zeros
	binary.LittleEndian.PutUint32(mem[off:off+4], 0x00000000)
	allZeros := binary.LittleEndian.Uint32(mem[off : off+4])

	// put it back
	binary.LittleEndian.PutUint32(mem[off:off+4], original)

	// RW mask: bits that flipped between all-ones and all-zeros
	rwMask := allOnes ^ allZeros
	rw1CMask := uint32(0)

	// RW1C check: write 1s to writable bits, see if they self-clear
	maybeRW1C := false
	maybeRW0C := false
	rw0CMask := uint32(0)
	if rwMask != 0 {
		testVal := original | rwMask
		binary.LittleEndian.PutUint32(mem[off:off+4], testVal)
		afterWrite := binary.LittleEndian.Uint32(mem[off : off+4])
		rw1CMask = testVal & ^afterWrite & rwMask
		if rw1CMask != 0 {
			maybeRW1C = true
		}

		binary.LittleEndian.PutUint32(mem[off:off+4], original)
		testVal = original &^ rwMask
		binary.LittleEndian.PutUint32(mem[off:off+4], testVal)
		afterWrite = binary.LittleEndian.Uint32(mem[off : off+4])
		rw0CMask = original &^ afterWrite & rwMask
		if rw0CMask != 0 {
			maybeRW0C = true
		}
		binary.LittleEndian.PutUint32(mem[off:off+4], original)
	}

	return BARProbeResult{
		Offset:    offset,
		Original:  original,
		RWMask:    rwMask,
		RW1CMask:  rw1CMask,
		RW0CMask:  rw0CMask,
		MaybeRW1C: maybeRW1C,
		MaybeRW0C: maybeRW0C,
	}
}
