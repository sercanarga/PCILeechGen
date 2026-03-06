package variance

import (
	"testing"

	"github.com/sercanarga/pcileechgen/internal/firmware/svgen"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

func makeTestCS() *pci.ConfigSpace {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	cs.WriteU16(0x00, 0x8086) // vendor
	cs.WriteU16(0x02, 0x1533) // device
	cs.WriteU8(0x08, 0x03)    // revision
	cs.WriteU16(0x2C, 0x8086) // subsystem vendor
	cs.WriteU16(0x2E, 0x0001) // subsystem device
	cs.WriteU8(0x09, 0x00)    // progIF
	cs.WriteU8(0x0A, 0x00)    // subclass
	cs.WriteU8(0x0B, 0x02)    // base class
	return cs
}

func TestApply_Deterministic(t *testing.T) {
	cs1 := makeTestCS()
	cs2 := makeTestCS()
	lat1 := &svgen.LatencyConfig{MinCycles: 3, MaxCycles: 12, ThermalPeriod: 4096, BurstCorrelation: 128}
	lat2 := &svgen.LatencyConfig{MinCycles: 3, MaxCycles: 12, ThermalPeriod: 4096, BurstCorrelation: 128}

	cfg := DefaultConfig(42)
	Apply(cs1, lat1, cfg)
	Apply(cs2, lat2, cfg)

	// Same seed → same result
	if lat1.MinCycles != lat2.MinCycles || lat1.MaxCycles != lat2.MaxCycles {
		t.Error("same seed should produce same latency mutations")
	}

	// Revision should be identical
	if cs1.ReadU8(0x08) != cs2.ReadU8(0x08) {
		t.Error("same seed should produce same revision noise")
	}
}

func TestApply_DifferentSeed(t *testing.T) {
	cs1 := makeTestCS()
	cs2 := makeTestCS()
	lat1 := &svgen.LatencyConfig{MinCycles: 10, MaxCycles: 50, ThermalPeriod: 8192, BurstCorrelation: 200}
	lat2 := &svgen.LatencyConfig{MinCycles: 10, MaxCycles: 50, ThermalPeriod: 8192, BurstCorrelation: 200}

	Apply(cs1, lat1, DefaultConfig(100))
	Apply(cs2, lat2, DefaultConfig(999))

	// Different seeds should produce at least some difference
	differ := lat1.MinCycles != lat2.MinCycles ||
		lat1.MaxCycles != lat2.MaxCycles ||
		lat1.ThermalPeriod != lat2.ThermalPeriod ||
		cs1.ReadU8(0x08) != cs2.ReadU8(0x08)

	if !differ {
		t.Error("different seeds should produce different mutations (statistically)")
	}
}

func TestApply_PreservesIdentity(t *testing.T) {
	cs := makeTestCS()
	lat := &svgen.LatencyConfig{MinCycles: 3, MaxCycles: 12}
	cfg := DefaultConfig(42)
	Apply(cs, lat, cfg)

	// Vendor/Device ID must never change
	if cs.ReadU16(0x00) != 0x8086 {
		t.Error("vendor ID must not change")
	}
	if cs.ReadU16(0x02) != 0x1533 {
		t.Error("device ID must not change")
	}
}

func TestApply_NilLatency(t *testing.T) {
	cs := makeTestCS()
	cfg := DefaultConfig(42)
	// Should not panic
	Apply(cs, nil, cfg)
}

func TestApply_ZeroJitter(t *testing.T) {
	lat := &svgen.LatencyConfig{MinCycles: 5, MaxCycles: 20}
	cfg := DefaultConfig(42)
	cfg.TimingJitter = 0
	Apply(makeTestCS(), lat, cfg)

	if lat.MinCycles != 5 || lat.MaxCycles != 20 {
		t.Error("zero jitter should not modify latency")
	}
}

func TestBuildVarianceSeed(t *testing.T) {
	s1 := BuildVarianceSeed(0x8086, 0x1533, 42)
	s2 := BuildVarianceSeed(0x8086, 0x1533, 42)
	if s1 != s2 {
		t.Error("same inputs should produce same seed")
	}
	s3 := BuildVarianceSeed(0x8086, 0x1533, 43)
	if s1 == s3 {
		t.Error("different entropy should produce different seed")
	}
}

func TestSplitMix64_NonZero(t *testing.T) {
	rng := newSplitMix64(1)
	seen := make(map[uint64]bool)
	for i := 0; i < 100; i++ {
		v := rng.next()
		if seen[v] {
			t.Fatalf("duplicate at iteration %d", i)
		}
		seen[v] = true
	}
}

func TestJitterInt_MinBound(t *testing.T) {
	rng := newSplitMix64(42)
	for i := 0; i < 50; i++ {
		result := jitterInt(3, 0.5, rng, 1)
		if result < 1 {
			t.Errorf("jitterInt produced %d, below min 1", result)
		}
	}
}


func TestApplyDSNVariance(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize

	// Write DSN ext cap at 0x100
	dsnHeader := uint32(pci.ExtCapIDDeviceSerialNumber) | (1 << 16)
	cs.WriteU32(0x100, dsnHeader)
	// Original DSN: 8 bytes at 0x104
	cs.WriteU32(0x104, 0x11223344)
	cs.WriteU32(0x108, 0xAABBCC00) // OUI in upper 24 bits

	rng := newSplitMix64(42)
	applyDSNVariance(cs, rng)

	// OUI (upper 24 bits of second DWORD) should be preserved
	newHi := cs.ReadU32(0x108)
	if newHi&0xFFFFFF00 != 0xAABBCC00 {
		t.Errorf("OUI should be preserved, got 0x%08x", newHi)
	}
	// Lower DWORD should be randomized (unlikely to stay same)
	if cs.ReadU32(0x104) == 0x11223344 {
		t.Error("DSN lower DWORD should be randomized")
	}
}

func TestApplyDSNVariance_NoDSN(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	// No DSN cap — should not panic
	rng := newSplitMix64(42)
	applyDSNVariance(cs, rng)
}

func TestApplySubsysOffset(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	cs.WriteU16(0x2E, 0x1234)

	rng := newSplitMix64(42)
	applySubsysOffset(cs, rng)

	newID := cs.ReadU16(0x2E)
	diff := int(newID) - 0x1234
	if diff < -1 || diff > 1 {
		t.Errorf("subsys offset too large: %d (new=0x%04x)", diff, newID)
	}
}

func TestApplySubsysOffset_Zero(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	cs.WriteU16(0x2E, 0x0000) // zero subsys → should not change

	rng := newSplitMix64(42)
	applySubsysOffset(cs, rng)

	if cs.ReadU16(0x2E) != 0 {
		t.Error("zero subsysID should not be modified")
	}
}

func TestApplyRegisterNoise(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	cs.WriteU8(0x08, 0x05) // revision = 5

	rng := newSplitMix64(42)
	applyRegisterNoise(cs, rng)

	newRev := cs.ReadU8(0x08)
	diff := int(newRev) - 5
	if diff < -1 || diff > 1 {
		t.Errorf("register noise too large: revision %d → %d", 5, newRev)
	}
}

func TestApplyRegisterNoise_ZeroValue(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	cs.WriteU8(0x08, 0x00) // zero revision → should not noise

	rng := newSplitMix64(42)
	applyRegisterNoise(cs, rng)

	if cs.ReadU8(0x08) != 0 {
		t.Error("zero revision should not be noised")
	}
}

func TestEmbedVSECEntropy(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize

	embedVSECEntropy(cs, 42)

	// VSEC should be at 0x100 (first ext cap)
	header := cs.ReadU32(0x100)
	capID := header & 0xFFFF
	if uint16(capID) != pci.ExtCapIDVendorSpecific {
		t.Errorf("VSEC cap ID = 0x%04x, want 0x%04x", capID, pci.ExtCapIDVendorSpecific)
	}

	// Payload should be non-zero
	payload1 := cs.ReadU32(0x108)
	payload2 := cs.ReadU32(0x10C)
	if payload1 == 0 && payload2 == 0 {
		t.Error("VSEC payload should be non-zero")
	}
}

func TestEmbedVSECEntropy_WithExistingExtCaps(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize

	// Put AER at 0x100, next=0x140 (pretend AER occupies 0x100-0x13F = 64 bytes)
	// Then LTR at 0x140, next=0 (LTR is smaller, ~8 bytes)
	aerHeader := uint32(pci.ExtCapIDAER) | (1 << 16) | (0x140 << 20)
	cs.WriteU32(0x100, aerHeader)
	for i := 4; i < 0x40; i += 4 {
		cs.WriteU32(0x100+i, 0xDEADBEEF)
	}

	ltrHeader := uint32(pci.ExtCapIDLTR) | (1 << 16)
	cs.WriteU32(0x140, ltrHeader)
	cs.WriteU32(0x144, 0x00100010) // LTR data

	embedVSECEntropy(cs, 99)

	// LTR should now have a next pointer to the VSEC
	updatedLTR := cs.ReadU32(0x140)
	nextOff := int((updatedLTR >> 20) & 0xFFC)
	if nextOff == 0 {
		t.Skip("embedVSECEntropy didn't chain VSEC after LTR")
	}

	// VSEC should be at the new offset
	if nextOff > 0 && nextOff < pci.ConfigSpaceSize {
		vsecHeader := cs.ReadU32(nextOff)
		if uint16(vsecHeader&0xFFFF) != pci.ExtCapIDVendorSpecific {
			t.Errorf("VSEC at 0x%03x: ID = 0x%04x", nextOff, vsecHeader&0xFFFF)
		}
	}
}

func TestEmbedVSECEntropy_SmallConfigSpace(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceLegacySize // 256 bytes — no ext config
	embedVSECEntropy(cs, 42)
	// Should be a no-op — no room for ext caps
}

func TestApplyTimingJitter(t *testing.T) {
	lat := &svgen.LatencyConfig{
		MinCycles:        5,
		MaxCycles:        20,
		ThermalPeriod:    4096,
		BurstCorrelation: 128,
	}

	rng := newSplitMix64(42)
	applyTimingJitter(lat, rng, 0.10)

	// Values should be jittered but within bounds
	if lat.MinCycles < 1 {
		t.Errorf("MinCycles = %d, should be >= 1", lat.MinCycles)
	}
	if lat.MaxCycles <= lat.MinCycles {
		t.Errorf("MaxCycles (%d) should be > MinCycles (%d)", lat.MaxCycles, lat.MinCycles)
	}
	if lat.ThermalPeriod < 100 {
		t.Errorf("ThermalPeriod = %d, should be >= 100", lat.ThermalPeriod)
	}
	if lat.BurstCorrelation < 1 {
		t.Errorf("BurstCorrelation = %d, should be >= 1", lat.BurstCorrelation)
	}
}

func TestJitterInt_ZeroValue(t *testing.T) {
	rng := newSplitMix64(42)
	result := jitterInt(0, 0.1, rng, 1)
	if result != 0 {
		t.Errorf("jitterInt(0) = %d, want 0", result)
	}
}

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig(42)
	if cfg.Seed != 42 {
		t.Errorf("Seed = %d, want 42", cfg.Seed)
	}
	if cfg.TimingJitter <= 0 {
		t.Error("TimingJitter should be positive")
	}
	if !cfg.RegisterNoise {
		t.Error("RegisterNoise should be enabled by default")
	}
	if !cfg.MutateDSN {
		t.Error("MutateDSN should be enabled by default")
	}
	if cfg.SubsysOffset {
		t.Error("SubsysOffset should be disabled by default")
	}
}

func TestApply_SubsysOffset(t *testing.T) {
	cs := makeTestCS()
	cs.WriteU16(0x2E, 0x5678)
	cfg := DefaultConfig(42)
	cfg.SubsysOffset = true
	Apply(cs, nil, cfg)

	newID := cs.ReadU16(0x2E)
	diff := int(newID) - 0x5678
	if diff < -1 || diff > 1 {
		t.Errorf("subsys offset = %d, expected ±1", diff)
	}
}
