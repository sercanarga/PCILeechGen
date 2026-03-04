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
