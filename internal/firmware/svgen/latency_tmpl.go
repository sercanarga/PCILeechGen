package svgen

import (
	"time"
)

// LatencyConfig holds TLP response timing parameters.
// Ideally these come from donor device profiling.
type LatencyConfig struct {
	MinCycles        int // minimum read response latency in clock cycles
	MaxCycles        int // maximum
	AvgCycles        int // average (used when jitter disabled)
	BurstCorrelation int // 0-255, higher = more correlated sequential timing
	ThermalPeriod    int // cycles between thermal drift baseline adjustments
}

// DefaultLatencyConfig returns device-class-appropriate latency defaults.
func DefaultLatencyConfig(classCode uint32) *LatencyConfig {
	baseClass := (classCode >> 16) & 0xFF
	subClass := (classCode >> 8) & 0xFF
	progIF := classCode & 0xFF

	switch {
	case baseClass == 0x01 && subClass == 0x08 && progIF == 0x02:
		// NVMe — SSDs are fast
		return &LatencyConfig{MinCycles: 3, MaxCycles: 12, AvgCycles: 6, BurstCorrelation: 160, ThermalPeriod: 49152}
	case baseClass == 0x0C && subClass == 0x03 && progIF == 0x30:
		// xHCI — USB controllers are moderate
		return &LatencyConfig{MinCycles: 4, MaxCycles: 18, AvgCycles: 8, BurstCorrelation: 140, ThermalPeriod: 40960}
	case baseClass == 0x02:
		// Ethernet — network controllers are fast
		return &LatencyConfig{MinCycles: 2, MaxCycles: 8, AvgCycles: 4, BurstCorrelation: 180, ThermalPeriod: 57344}
	default:
		return &LatencyConfig{MinCycles: 3, MaxCycles: 15, AvgCycles: 7, BurstCorrelation: 128, ThermalPeriod: 32768}
	}
}

// BuildPRNGSeeds generates 4 deterministic but unique PRNG seeds from device IDs + entropy.
func BuildPRNGSeeds(vid, did uint16, entropy uint32) [4]uint32 {
	v := uint32(vid)
	d := uint32(did)
	return [4]uint32{
		(v ^ d) ^ entropy,
		((v << 16) | d) ^ ((entropy * 31) & 0xFFFFFFFF),
		(d ^ uint32(vid)) ^ ((entropy * 127) & 0xFFFFFFFF),
		v ^ ((entropy * 8191) & 0xFFFFFFFF),
	}
}

// BuildEntropyFromTime generates a build entropy seed from current time.
func BuildEntropyFromTime() uint32 {
	return uint32(time.Now().UnixNano() & 0xFFFFFFFF)
}
