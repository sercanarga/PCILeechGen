package devclass

import (
	"testing"

	"github.com/sercanarga/pcileechgen/internal/util"
)

func TestStrategyForClass_RAID(t *testing.T) {
	s := StrategyForClass(0x010400)
	if s == nil {
		t.Fatal("expected RAID strategy, got nil")
	}
	if s.ClassName() != "RAID Controller" {
		t.Errorf("expected RAID Controller, got %s", s.ClassName())
	}
	if s.DeviceClass() != ClassRAID {
		t.Errorf("expected %s, got %s", ClassRAID, s.DeviceClass())
	}
	if s.Profile() == nil {
		t.Error("expected non-nil profile")
	}
}

func TestProfileForClass_RAID(t *testing.T) {
	p := ProfileForClass(0x010400)
	if p == nil {
		t.Fatal("RAID profile should not be nil")
	}
	if p.ClassName != "RAID Controller" {
		t.Errorf("got %q, want RAID Controller", p.ClassName)
	}
	if !p.PrefersMSIX {
		t.Error("RAID should prefer MSI-X")
	}
	if !p.Uses64BitBAR {
		t.Error("RAID (Fusion) should use 64-bit BAR")
	}
	if p.PreferredBAR != 1 {
		t.Errorf("RAID PreferredBAR should be 1, got %d", p.PreferredBAR)
	}
	if p.MinBARSize != 65536 {
		t.Errorf("RAID MinBARSize should be 65536, got %d", p.MinBARSize)
	}
	if len(p.BARDefaults) == 0 {
		t.Error("RAID should have BAR defaults")
	}
}

func TestRAIDGeneration_Fusion(t *testing.T) {
	// SAS3108 (Invader) and unknown IDs default to the Fusion layout.
	for _, did := range []uint16{0x005d, 0x0017, 0x0000, 0xFFFF} {
		p := ProfileForDevice(0x010400, 0x1000, did)
		if !p.Uses64BitBAR {
			t.Errorf("device 0x%04X: Fusion should use 64-bit BAR", did)
		}
		if p.PreferredBAR != 1 {
			t.Errorf("device 0x%04X: Fusion register BAR should be 1, got %d", did, p.PreferredBAR)
		}
	}
}

func TestRAIDGeneration_MFI(t *testing.T) {
	// SAS2208 / SAS2008 use the legacy MFI layout: 32-bit BAR0.
	for _, did := range []uint16{0x005b, 0x0073} {
		p := ProfileForDevice(0x010400, 0x1000, did)
		if p.Uses64BitBAR {
			t.Errorf("device 0x%04X: Gen2 MFI BAR is 32-bit", did)
		}
		if p.PreferredBAR != 0 {
			t.Errorf("device 0x%04X: Gen2 register BAR should be 0, got %d", did, p.PreferredBAR)
		}
	}
}

func TestRAIDStrategyForDevice_Class(t *testing.T) {
	s := StrategyForDevice(0x010400, 0x1000, 0x005b)
	if s.DeviceClass() != ClassRAID {
		t.Errorf("expected %s, got %s", ClassRAID, s.DeviceClass())
	}
	if s.ClassName() != "RAID Controller" {
		t.Errorf("expected RAID Controller, got %s", s.ClassName())
	}
}

// TestRAIDProfile_ScratchPadReady verifies the firmware-state nibble of both
// scratch pads is READY (0xB) so megaraid_sas proceeds past the ready poll.
func TestRAIDProfile_ScratchPadReady(t *testing.T) {
	p := raidProfile()
	want := map[uint32]bool{0xA8: false, 0xB0: false}
	for _, d := range p.BARDefaults {
		if _, ok := want[d.Offset]; ok {
			if d.Reset&mfiStateMask != mfiStateReady {
				t.Errorf("scratch pad 0x%X FW-state should be READY, got reset 0x%08X", d.Offset, d.Reset)
			}
			want[d.Offset] = true
		}
	}
	for off, seen := range want {
		if !seen {
			t.Errorf("scratch pad 0x%X missing from BARDefaults", off)
		}
	}
}

func TestRAIDProfile_MasksDisjoint(t *testing.T) {
	p := raidProfile()
	for _, d := range p.BARDefaults {
		if d.W1CMask&d.RWMask != 0 {
			t.Errorf("0x%X %s: W1CMask 0x%08X overlaps RWMask 0x%08X", d.Offset, d.Name, d.W1CMask, d.RWMask)
		}
	}
}

func TestRAIDStrategy_ScrubBAR(t *testing.T) {
	s := &raidStrategy{}
	data := make([]byte, 0xC8)
	// Seed scratch pads with a non-ready state plus lower hint bits.
	util.WriteLE32(data, 0xA8, 0x000000AB) // lower bits must be preserved
	util.WriteLE32(data, 0xB0, 0x40000000) // wrong FW state

	s.ScrubBAR(data)

	sp0 := util.ReadLE32(data, 0xA8)
	if sp0&mfiStateMask != mfiStateReady {
		t.Errorf("scratch_pad_0 FW-state should be READY, got 0x%08X", sp0)
	}
	if sp0&0x0000FFFF != 0x000000AB {
		t.Errorf("scratch_pad_0 lower bits should be preserved, got 0x%08X", sp0)
	}
	sp2 := util.ReadLE32(data, 0xB0)
	if sp2&mfiStateMask != mfiStateReady {
		t.Errorf("scratch_pad_2 FW-state should be READY, got 0x%08X", sp2)
	}
}

func TestRAIDStrategy_ScrubBAR_TooShort(t *testing.T) {
	s := &raidStrategy{}
	data := make([]byte, 0x40)
	s.ScrubBAR(data) // must not panic
}

func TestRAIDStrategy_PostInitRegisters(t *testing.T) {
	s := &raidStrategy{}
	var sp0 uint32 = 0x40000000
	var sp2 uint32 = 0x00000000
	regs := map[uint32]*uint32{0xA8: &sp0, 0xB0: &sp2}
	s.PostInitRegisters(regs)
	if sp0&mfiStateMask != mfiStateReady {
		t.Errorf("scratch_pad_0 should be READY, got 0x%08X", sp0)
	}
	if sp2&mfiStateMask != mfiStateReady {
		t.Errorf("scratch_pad_2 should be READY, got 0x%08X", sp2)
	}
}
