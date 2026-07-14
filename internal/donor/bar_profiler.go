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
	W1CMask   uint32 `json:"w1c_mask"`   // bits that self-cleared on write-of-1
	MaybeRW1C bool   `json:"maybe_rw1c"` // true if W1CMask != 0
}

// BARProfile is the full probe output for one BAR.
type BARProfile struct {
	BarIndex int              `json:"bar_index"`
	Size     int              `json:"size"`
	Probes   []BARProbeResult `json:"probes"`
}

// BARProfiler snapshots BAR registers by default. Active write probing is
// intentionally opt-in because writes to unknown MMIO registers can start DMA,
// reset hardware, acknowledge interrupts, or otherwise change device state.
type BARProfiler struct {
	active bool
}

func NewBARProfiler() *BARProfiler { return &BARProfiler{} }

// NewActiveBARProfiler enables the legacy destructive write/readback probe.
// Callers must only use it for a device and register range known to be safe.
func NewActiveBARProfiler() *BARProfiler { return &BARProfiler{active: true} }

// ProfileBAR snapshots each DWORD. It only performs write/readback probing when
// the profiler was explicitly created with NewActiveBARProfiler.
func (p *BARProfiler) ProfileBAR(resourcePath string, barIndex, maxSize int) (*BARProfile, error) {
	flags := os.O_RDONLY
	prot := syscall.PROT_READ
	mode := "read-only"
	if p != nil && p.active {
		flags = os.O_RDWR
		prot |= syscall.PROT_WRITE
		mode = "active"
	}
	f, err := os.OpenFile(resourcePath, flags, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to open BAR%d for %s profiling: %w", barIndex, mode, err)
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

	mapped, err := syscall.Mmap(int(f.Fd()), 0, mmapSize, prot, syscall.MAP_SHARED)
	if err != nil {
		return nil, fmt.Errorf("mmap %s failed for BAR%d: %w", mode, barIndex, err)
	}
	defer func() {
		if unmapErr := syscall.Munmap(mapped); unmapErr != nil {
			_ = unmapErr
		}
	}()

	profile := &BARProfile{
		BarIndex: barIndex,
		Size:     size,
	}

	if p != nil && p.active {
		profile.Probes = probeRegisters(mapped, size)
	} else {
		profile.Probes = snapshotRegisters(mapped, size)
	}

	return profile, nil
}

func snapshotRegisters(mem []byte, size int) []BARProbeResult {
	numRegs := size / 4
	probes := make([]BARProbeResult, 0, numRegs)
	for off := 0; off < numRegs*4; off += 4 {
		probes = append(probes, BARProbeResult{
			Offset:   uint32(off),
			Original: binary.LittleEndian.Uint32(mem[off : off+4]),
		})
	}
	return probes
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

// classifyRegisterBits derives the RW mask and the W1C-suspect mask from probe
// readbacks. Split out from probeOneRegister so it can be tested without a device.
func classifyRegisterBits(allOnes, allZeros, testVal, afterWrite uint32) (rwMask, w1cMask uint32) {
	rwMask = allOnes ^ allZeros
	if rwMask != 0 {
		w1cMask = testVal & ^afterWrite & rwMask
	}
	return
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

	// W1C probe: write 1s to writable bits and read back, to see which self-clear.
	rwMask := allOnes ^ allZeros
	var testVal, afterWrite uint32
	if rwMask != 0 {
		testVal = original | rwMask
		binary.LittleEndian.PutUint32(mem[off:off+4], testVal)
		afterWrite = binary.LittleEndian.Uint32(mem[off : off+4])
	}

	// put it back
	binary.LittleEndian.PutUint32(mem[off:off+4], original)

	_, w1cMask := classifyRegisterBits(allOnes, allZeros, testVal, afterWrite)

	return BARProbeResult{
		Offset:    offset,
		Original:  original,
		RWMask:    rwMask,
		W1CMask:   w1cMask,
		MaybeRW1C: w1cMask != 0,
	}
}
