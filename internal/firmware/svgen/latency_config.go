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
	WrHasHistogram   bool
	WrHistogram      [16]uint8
	WrCDF            [16]uint8

	// Write TLP timing - determines how long write acknowledges are delayed.
	// Posted writes (MWr) don't need completions per PCIe spec, but the
	// internal pipeline still has a write-accept latency visible in timing analysis.
	WrMinCycles int
	WrMaxCycles int
	WrHotEnable int
	WrHotOffset uint32
	WrHotMask   uint32
	WrHotMin    int
	WrHotMax    int

	// Completion timeout - cycles before an unserviced read returns UR.
	// 0 = disabled (no timeout). Reasonable default: 65536 (~0.5ms @125MHz).
	CplTimeoutCycles int

	HotReadEnable  int
	HotReadOffset  uint32
	HotReadMask    uint32
	HotReadMin     int
	HotReadMax     int

	PollReadEnable int
	PollReadOffset uint32
	PollReadMask   uint32
	PollReadMin    int
	PollReadMax    int
}

// DefaultLatencyConfig returns class-appropriate latency defaults.
func DefaultLatencyConfig(classCode uint32) *LatencyConfig {
	baseClass := (classCode >> 16) & 0xFF
	subClass := (classCode >> 8) & 0xFF
	progIF := classCode & 0xFF

	switch {
	case baseClass == 0x01 && subClass == 0x08 && progIF == 0x02: // NVMe
		return withUniformCDF(&LatencyConfig{
			MinCycles: 3, MaxCycles: 12, AvgCycles: 6,
			BurstCorrelation: 160, ThermalPeriod: 49152,
			WrMinCycles: 2, WrMaxCycles: 6, CplTimeoutCycles: 65536,
		})
	case baseClass == 0x0C && subClass == 0x03 && progIF == 0x30: // xHCI
		return withUniformCDF(&LatencyConfig{
			MinCycles: 4, MaxCycles: 18, AvgCycles: 8,
			BurstCorrelation: 140, ThermalPeriod: 40960,
			WrMinCycles: 3, WrMaxCycles: 10, CplTimeoutCycles: 65536,
		})
	case baseClass == 0x02 && subClass == 0x00: // Ethernet
		return withUniformCDF(&LatencyConfig{
			MinCycles: 2, MaxCycles: 8, AvgCycles: 4,
			BurstCorrelation: 180, ThermalPeriod: 57344,
			WrMinCycles: 1, WrMaxCycles: 4, CplTimeoutCycles: 65536,
		})
	case baseClass == 0x03 && subClass == 0x00: // GPU
		return withUniformCDF(&LatencyConfig{
			MinCycles: 5, MaxCycles: 25, AvgCycles: 12,
			BurstCorrelation: 200, ThermalPeriod: 65536,
			WrMinCycles: 3, WrMaxCycles: 10, CplTimeoutCycles: 131072,
		})
	case baseClass == 0x01 && subClass == 0x06: // SATA/AHCI
		return withUniformCDF(&LatencyConfig{
			MinCycles: 3, MaxCycles: 14, AvgCycles: 7,
			BurstCorrelation: 150, ThermalPeriod: 45056,
			WrMinCycles: 2, WrMaxCycles: 8, CplTimeoutCycles: 65536,
		})
	case baseClass == 0x04 && subClass == 0x03: // HD Audio
		return withUniformCDF(&LatencyConfig{
			MinCycles: 2, MaxCycles: 10, AvgCycles: 5,
			BurstCorrelation: 120, ThermalPeriod: 36864,
			WrMinCycles: 1, WrMaxCycles: 5, CplTimeoutCycles: 262144,
		})
	case baseClass == 0x02 && subClass == 0x80: // Wi-Fi
		return withUniformCDF(&LatencyConfig{
			MinCycles: 3, MaxCycles: 16, AvgCycles: 8,
			BurstCorrelation: 160, ThermalPeriod: 53248,
			WrMinCycles: 2, WrMaxCycles: 8, CplTimeoutCycles: 65536,
		})
	case baseClass == 0x0C && subClass == 0x80: // Thunderbolt
		return withUniformCDF(&LatencyConfig{
			MinCycles: 4, MaxCycles: 20, AvgCycles: 10,
			BurstCorrelation: 170, ThermalPeriod: 49152,
			WrMinCycles: 3, WrMaxCycles: 12, CplTimeoutCycles: 65536,
		})
	default:
		return withUniformCDF(&LatencyConfig{
			MinCycles: 3, MaxCycles: 15, AvgCycles: 7,
			BurstCorrelation: 128, ThermalPeriod: 32768,
			WrMinCycles: 2, WrMaxCycles: 8, CplTimeoutCycles: 65536,
		})
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
		WrHasHistogram:   false,
		WrHistogram:      defCfg.WrHistogram,
		WrCDF:            defCfg.WrCDF,
		WrMinCycles:      defCfg.WrMinCycles,
		WrMaxCycles:      defCfg.WrMaxCycles,
		CplTimeoutCycles: defCfg.CplTimeoutCycles,
	}
}

func LatencyConfigFromProfile(profile *behavior.Profile, classCode uint32) *LatencyConfig {
	if profile == nil {
		return DefaultLatencyConfig(classCode)
	}

	cfg := LatencyConfigFromHistogram(profile.ReadLatency, classCode)
	cfg = applyWriteLatency(cfg, profile.WriteLatency, classCode)
	cfg = applyWriteHotTuning(cfg, profile)
	cfg = applyHotReadTuning(cfg, profile)
	cfg = applyPollReadTuning(cfg, profile)
	return cfg
}

func applyWriteLatency(cfg *LatencyConfig, h *behavior.TimingHistogram, classCode uint32) *LatencyConfig {
	if cfg == nil {
		return DefaultLatencyConfig(classCode)
	}
	if h == nil || h.SampleCount == 0 {
		return cfg
	}
	cfg.WrHasHistogram = true
	cfg.WrHistogram = h.Buckets
	cfg.WrCDF = h.CDF
	cfg.WrMinCycles = h.MinCycles
	cfg.WrMaxCycles = h.MaxCycles
	cfg.WrMinCycles = clampLatency(cfg.WrMinCycles, 1, cfg.WrMaxCycles)
	cfg.WrMaxCycles = clampLatency(cfg.WrMaxCycles, cfg.WrMinCycles, 255)
	return cfg
}

func applyWriteHotTuning(cfg *LatencyConfig, profile *behavior.Profile) *LatencyConfig {
	if cfg == nil || profile == nil || len(profile.AccessStats.HotWrites) == 0 {
		return cfg
	}
	hot := profile.AccessStats.HotWrites[0]
	if !isSignificantHotSpot(hot.Count, profile.AccessStats.TotalWrites) {
		return cfg
	}
	cfg.WrHotEnable = 1
	cfg.WrHotOffset = hot.Offset & 0xFFF
	cfg.WrHotMask = 0x00000FFF
	cfg.WrHotMin = clampLatency(cfg.WrMinCycles, 1, cfg.WrMaxCycles)
	cfg.WrHotMax = clampLatency(cfg.WrMaxCycles, cfg.WrMinCycles, 255)
	return cfg
}

func applyHotReadTuning(cfg *LatencyConfig, profile *behavior.Profile) *LatencyConfig {
	if cfg == nil || profile == nil || len(profile.AccessStats.HotReads) == 0 {
		return cfg
	}
	hot := profile.AccessStats.HotReads[0]
	if !isSignificantHotSpot(hot.Count, profile.AccessStats.TotalReads) {
		return cfg
	}
	cfg.HotReadEnable = 1
	cfg.HotReadOffset = hot.Offset & 0xFFF
	cfg.HotReadMask = 0x00000FFF
	cfg.HotReadMin = clampLatency(cfg.MinCycles-1, 1, cfg.MaxCycles)
	cfg.HotReadMax = clampLatency(cfg.MinCycles, cfg.HotReadMin, cfg.MaxCycles)
	return cfg
}

func applyPollReadTuning(cfg *LatencyConfig, profile *behavior.Profile) *LatencyConfig {
	if cfg == nil || profile == nil {
		return cfg
	}
	if len(profile.PollingLoops) == 0 {
		return cfg
	}
	poll := profile.PollingLoops[0]
	if poll.Count < 4 || poll.IntervalNs <= 0 {
		return cfg
	}
	if profile.AccessStats.TotalReads == 0 {
		return cfg
	}
	cfg.PollReadEnable = 1
	cfg.PollReadOffset = poll.Offset & 0xFFF
	cfg.PollReadMask = 0x00000FFF
	baseMin := cfg.MinCycles
	if cfg.HotReadEnable != 0 && cfg.HotReadMin > 0 {
		baseMin = cfg.HotReadMin
	}
	cfg.PollReadMin = clampLatency(baseMin-1, 1, cfg.MaxCycles)
	cfg.PollReadMax = clampLatency(baseMin, cfg.PollReadMin, cfg.MaxCycles)
	return cfg
}

func isSignificantHotSpot(count, total int) bool {
	if total <= 0 {
		return false
	}
	if count < 2 {
		return false
	}
	return count*10 >= total
}

func clampLatency(v, min, max int) int {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}

func withUniformCDF(cfg *LatencyConfig) *LatencyConfig {
	for i := range cfg.CDF {
		cfg.CDF[i] = uint8(((i + 1) * 255) / 16)
	}
	for i := range cfg.Histogram {
		cfg.Histogram[i] = 16
	}
	for i := range cfg.WrCDF {
		cfg.WrCDF[i] = uint8(((i + 1) * 255) / 16)
	}
	for i := range cfg.WrHistogram {
		cfg.WrHistogram[i] = 16
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
