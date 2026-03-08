// Package variance tweaks config space per-build so each bitstream is unique.
package variance

import (
	"encoding/binary"
	"hash/fnv"
	"math"

	"github.com/sercanarga/pcileechgen/internal/firmware/svgen"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

// Config controls which variations are applied.
type Config struct {
	Seed            uint32  // entropy seed — different seed ⇒ different build
	TimingJitter    float64 // fraction of latency range to jitter (0.0–0.15)
	RegisterNoise   bool    // ±1-2 LSB noise on non-critical reset values
	MutateDSN       bool    // generate a realistic DSN from OUI + seed
	SubsysOffset    bool    // apply tiny offset to subsystem ID
	TimingJitterPct float64 // (alias kept for readability)
}

// DefaultConfig returns sensible defaults.
func DefaultConfig(seed uint32) Config {
	return Config{
		Seed:          seed,
		TimingJitter:  0.08, // ±8 %
		RegisterNoise: true,
		MutateDSN:     true,
		SubsysOffset:  false, // disabled by default — changing subsys ID can confuse drivers
	}
}

// --- core ---

// Apply mutates config space + latency in-place. Same seed → same result.
func Apply(cs *pci.ConfigSpace, latCfg *svgen.LatencyConfig, cfg Config) {
	rng := newSplitMix64(uint64(cfg.Seed))

	if cfg.MutateDSN {
		applyDSNVariance(cs, rng)
	}

	if cfg.SubsysOffset {
		applySubsysOffset(cs, rng)
	}

	if cfg.RegisterNoise {
		applyRegisterNoise(cs, rng)
	}

	if latCfg != nil && cfg.TimingJitter > 0 {
		applyTimingJitter(latCfg, rng, cfg.TimingJitter)
	}

	applyPMTimingVariance(cs, rng)

	embedVSECEntropy(cs, cfg.Seed)
}

// --- DSN variance ---

func applyDSNVariance(cs *pci.ConfigSpace, rng *splitMix64) {
	// Find DSN extended capability (ID = 0x0003)
	extCaps := pci.ParseExtCapabilities(cs)
	for _, cap := range extCaps {
		if cap.ID != pci.ExtCapIDDeviceSerialNumber {
			continue
		}
		// DSN is 8 bytes at cap.Offset+4
		dsnOff := cap.Offset + 4
		if dsnOff+8 > cs.Size {
			continue
		}
		// Keep OUI (bytes 5-7 of the DSN — upper 3 bytes of the second DWORD)
		// and randomize the serial portion (lower 5 bytes)
		origHi := cs.ReadU32(dsnOff + 4)
		oui := origHi & 0xFFFFFF00 // preserve vendor OUI in upper 24 bits

		serial := uint32(rng.next() & 0xFFFFFFFF)
		loWord := serial
		hiWord := oui | uint32(rng.next()&0xFF) // one randomized byte under OUI

		cs.WriteU32(dsnOff, loWord)
		cs.WriteU32(dsnOff+4, hiWord)
		return
	}
}

// --- subsystem ID offset ---

func applySubsysOffset(cs *pci.ConfigSpace, rng *splitMix64) {
	subsysID := cs.ReadU16(0x2E)
	if subsysID == 0 {
		return
	}
	// ±1 offset — tiny enough to not break driver matching on most devices
	offset := int16(rng.next()%3) - 1 // -1, 0, or +1
	newID := uint16(int16(subsysID) + offset)
	if newID == 0 {
		newID = subsysID // don't zero it out
	}
	cs.WriteU16(0x2E, newID)
}

// --- register noise ---

// offsets where ±1 LSB noise is safe (informational fields).
var safeNoiseOffsets = []int{
	0x08, // Revision ID (lower 8 bits — but we only touch if already non-zero)
}

func applyRegisterNoise(cs *pci.ConfigSpace, rng *splitMix64) {
	for _, off := range safeNoiseOffsets {
		if off >= cs.Size {
			continue
		}
		val := cs.ReadU8(off)
		if val == 0 {
			continue // don't noise a zero field
		}
		noise := int8(rng.next()%3) - 1 // -1, 0, or +1
		newVal := int16(val) + int16(noise)
		if newVal < 1 {
			newVal = 1
		}
		if newVal > 255 {
			newVal = 255
		}
		cs.WriteU8(off, uint8(newVal))
	}
}

// --- timing jitter ---

func applyTimingJitter(latCfg *svgen.LatencyConfig, rng *splitMix64, jitterFrac float64) {
	// Jitter min/max cycles within ±jitterFrac
	latCfg.MinCycles = jitterInt(latCfg.MinCycles, jitterFrac, rng, 1)
	latCfg.MaxCycles = jitterInt(latCfg.MaxCycles, jitterFrac, rng, latCfg.MinCycles+1)

	// Jitter thermal drift period
	if latCfg.ThermalPeriod > 0 {
		latCfg.ThermalPeriod = jitterInt(latCfg.ThermalPeriod, jitterFrac, rng, 100)
	}

	// Jitter burst correlation
	if latCfg.BurstCorrelation > 0 {
		latCfg.BurstCorrelation = jitterInt(latCfg.BurstCorrelation, jitterFrac, rng, 1)
	}
}

func jitterInt(val int, frac float64, rng *splitMix64, minVal int) int {
	if val == 0 {
		return 0
	}
	spread := int(math.Ceil(float64(val) * frac))
	if spread == 0 {
		spread = 1
	}
	delta := int(rng.next()%uint64(2*spread+1)) - spread
	result := val + delta
	if result < minVal {
		result = minVal
	}
	return result
}

// --- VSEC entropy embed ---

// embedVSECEntropy writes a VSEC ext cap with build-unique payload.
func embedVSECEntropy(cs *pci.ConfigSpace, seed uint32) {
	if cs.Size < pci.ConfigSpaceSize {
		return
	}

	// find the last ext cap and the first free offset after it
	extCaps := pci.ParseExtCapabilities(cs)
	lastEnd := 0x100 // start of ext config
	var lastCapOff int
	for _, cap := range extCaps {
		capEnd := cap.Offset + len(cap.Data)
		if capEnd > lastEnd {
			lastEnd = capEnd
			lastCapOff = cap.Offset
		}
	}

	// VSEC needs 16 bytes: 4 header + 4 VSEC header + 8 payload
	vsecSize := 16
	vsecOff := (lastEnd + 3) &^ 3 // DWORD-align
	if vsecOff+vsecSize > pci.ConfigSpaceSize {
		return // no room
	}

	// build payload: FNV-1a hash of seed as 2 DWORDs
	h := fnv.New64a()
	var buf [4]byte
	binary.LittleEndian.PutUint32(buf[:], seed)
	h.Write(buf[:])
	entropy := h.Sum64()

	// ext cap header: ID=0x000B (VSEC), version=1, next=0
	vsecHeader := uint32(pci.ExtCapIDVendorSpecific) | (1 << 16)
	cs.WriteU32(vsecOff, vsecHeader)

	// VSEC header: VSEC ID=0xFC (private), rev=1, length=16
	vsecHdr2 := uint32(0xFC) | (1 << 16) | (uint32(vsecSize) << 20)
	cs.WriteU32(vsecOff+4, vsecHdr2)

	// payload
	cs.WriteU32(vsecOff+8, uint32(entropy))
	cs.WriteU32(vsecOff+12, uint32(entropy>>32))

	// chain: update last cap's next pointer
	if lastCapOff > 0 && lastCapOff >= 0x100 {
		oldHeader := cs.ReadU32(lastCapOff)
		newHeader := (oldHeader & 0x000FFFFF) | (uint32(vsecOff) << 20)
		cs.WriteU32(lastCapOff, newHeader)
	} else if len(extCaps) == 0 {
		// no ext caps at all — write at 0x100
		cs.WriteU32(0x100, vsecHeader)
	}
}

// applyPMTimingVariance jitters the PMC Data Scale field.
func applyPMTimingVariance(cs *pci.ConfigSpace, rng *splitMix64) {
	caps := pci.ParseCapabilities(cs)
	for _, cap := range caps {
		if cap.ID != pci.CapIDPowerManagement {
			continue
		}
		// PMC at cap+0x02, Data Scale = bits [14:13]
		if cap.Offset+4 >= pci.ConfigSpaceLegacySize {
			continue
		}
		pmc := cs.ReadU16(cap.Offset + 2)

		scale := (pmc >> 13) & 0x03
		noise := int16(rng.next()%3) - 1 // -1, 0, +1
		newScale := int16(scale) + noise
		if newScale < 0 {
			newScale = 0
		}
		if newScale > 3 {
			newScale = 3
		}
		pmc = (pmc & 0x9FFF) | (uint16(newScale) << 13)
		cs.WriteU16(cap.Offset+2, pmc)
		break
	}
}

// --- splitmix64 PRNG ---

type splitMix64 struct {
	state uint64
}

func newSplitMix64(seed uint64) *splitMix64 {
	if seed == 0 {
		seed = 0x9E3779B97F4A7C15
	}
	return &splitMix64{state: seed}
}

func (s *splitMix64) next() uint64 {
	s.state += 0x9E3779B97F4A7C15
	z := s.state
	z = (z ^ (z >> 30)) * 0xBF58476D1CE4E5B9
	z = (z ^ (z >> 27)) * 0x94D049BB133111EB
	return z ^ (z >> 31)
}

// --- seed helper ---

// BuildVarianceSeed derives a deterministic seed from device IDs + entropy.
func BuildVarianceSeed(vendorID, deviceID uint16, buildEntropy uint32) uint32 {
	h := fnv.New32a()
	var buf [8]byte
	binary.LittleEndian.PutUint16(buf[0:2], vendorID)
	binary.LittleEndian.PutUint16(buf[2:4], deviceID)
	binary.LittleEndian.PutUint32(buf[4:8], buildEntropy)
	h.Write(buf[:])
	return h.Sum32()
}
