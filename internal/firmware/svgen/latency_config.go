package svgen

import (
	"time"

	"github.com/sercanarga/pcileechgen/internal/donor/behavior"
)

// LatencyConfig holds TLP response timing parameters.
type LatencyConfig struct {
	MinCycles        int
	MaxCycles        int
	AvgCycles        int
	BurstCorrelation int       // 0-255
	ThermalPeriod    int       // cycles between drift adjustments
	Histogram        [16]uint8 // 16-bucket weights (0-255)
	CDF              [16]uint8 // cumulative distribution for SV lookup
	HasHistogram     bool      // true = donor profiled
}

// DefaultLatencyConfig returns class-appropriate latency defaults.
func DefaultLatencyConfig(classCode uint32) *LatencyConfig {
	baseClass := (classCode >> 16) & 0xFF
	subClass := (classCode >> 8) & 0xFF
	progIF := classCode & 0xFF

	switch {
	case baseClass == 0x01 && subClass == 0x08 && progIF == 0x02: // NVMe
		return withUniformCDF(&LatencyConfig{MinCycles: 3, MaxCycles: 12, AvgCycles: 6, BurstCorrelation: 160, ThermalPeriod: 49152})
	case baseClass == 0x0C && subClass == 0x03 && progIF == 0x30: // xHCI
		return withUniformCDF(&LatencyConfig{MinCycles: 4, MaxCycles: 18, AvgCycles: 8, BurstCorrelation: 140, ThermalPeriod: 40960})
	case baseClass == 0x02: // Ethernet
		return withUniformCDF(&LatencyConfig{MinCycles: 2, MaxCycles: 8, AvgCycles: 4, BurstCorrelation: 180, ThermalPeriod: 57344})
	default:
		return withUniformCDF(&LatencyConfig{MinCycles: 3, MaxCycles: 15, AvgCycles: 7, BurstCorrelation: 128, ThermalPeriod: 32768})
	}
}

// LatencyConfigFromHistogram builds config from a donor timing histogram.
func LatencyConfigFromHistogram(h *behavior.TimingHistogram, classCode uint32) *LatencyConfig {
	if h == nil || h.SampleCount == 0 {
		return DefaultLatencyConfig(classCode)
	}

	defCfg := DefaultLatencyConfig(classCode)
	return &LatencyConfig{
		MinCycles:        h.MinCycles,
		MaxCycles:        h.MaxCycles,
		AvgCycles:        h.MedianCycles,
		BurstCorrelation: defCfg.BurstCorrelation,
		ThermalPeriod:    defCfg.ThermalPeriod,
		Histogram:        h.Buckets,
		CDF:              h.CDF,
		HasHistogram:     true,
	}
}

func withUniformCDF(cfg *LatencyConfig) *LatencyConfig {
	for i := range cfg.CDF {
		cfg.CDF[i] = uint8(((i + 1) * 255) / 16)
	}
	for i := range cfg.Histogram {
		cfg.Histogram[i] = 16
	}
	return cfg
}

// BuildPRNGSeeds generates 4 deterministic PRNG seeds from device IDs + entropy.
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
