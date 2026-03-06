package barmodel

import (
	"testing"

	"github.com/sercanarga/pcileechgen/internal/donor"
)

func TestBuildBARModel_NVMe(t *testing.T) {
	barData := make([]byte, 4096)
	// CAP register: MQES=63, CQR=1, AMS=0, TO=1, DSTRD=0, CSS=1
	barData[0] = 0x3F // MQES low
	barData[1] = 0x00
	barData[2] = 0x01 // CQR
	barData[3] = 0x00

	model := BuildBARModel(barData, 0x010802, nil)
	if model == nil {
		t.Fatal("NVMe BuildBARModel returned nil")
	}
	if model.Size != 4096 {
		t.Errorf("BAR size: got %d, want 4096", model.Size)
	}
	// Must have CSTS register with RDY=1
	foundCSTS := false
	for _, reg := range model.Registers {
		if reg.Name == "CSTS" {
			foundCSTS = true
			if reg.Reset&0x01 == 0 {
				t.Error("CSTS.RDY should be set to 1")
			}
		}
	}
	if !foundCSTS {
		t.Error("NVMe model should contain CSTS register")
	}
}

func TestBuildBARModel_XHCI(t *testing.T) {
	model := BuildBARModel(nil, 0x0C0330, nil)
	if model == nil {
		t.Fatal("xHCI BuildBARModel returned nil")
	}
	foundUSBCMD := false
	for _, reg := range model.Registers {
		if reg.Name == "USBCMD" {
			foundUSBCMD = true
			if reg.Reset&0x01 == 0 {
				t.Error("USBCMD.R/S should be 1")
			}
		}
	}
	if !foundUSBCMD {
		t.Error("xHCI model should contain USBCMD register")
	}
}

func TestBuildBARModel_Unknown(t *testing.T) {
	model := BuildBARModel(nil, 0xFF0000, nil)
	if model != nil {
		t.Error("unknown class without profile should return nil")
	}
}

func TestSynthesizeBARModel(t *testing.T) {
	profile := &donor.BARProfile{
		BarIndex: 0,
		Size:     4096,
		Probes: []donor.BARProbeResult{
			{Offset: 0x00, Original: 0x3F010040, RWMask: 0x00000000},
			{Offset: 0x14, Original: 0x00000001, RWMask: 0x00FFFFF1},
			{Offset: 0x1C, Original: 0x00000001, RWMask: 0x00000000},
			{Offset: 0x20, Original: 0x00000000, RWMask: 0x00000000}, // dead
		},
	}
	model := SynthesizeBARModel(profile, 0x010802)
	if model == nil {
		t.Fatal("SynthesizeBARModel returned nil")
	}
	// Dead register (original=0, rwmask=0) should be dropped
	if len(model.Registers) != 3 {
		t.Errorf("expected 3 registers (dead dropped), got %d", len(model.Registers))
	}
}

func TestSynthesizeBARModel_RW1C(t *testing.T) {
	profile := &donor.BARProfile{
		BarIndex: 0,
		Size:     4096,
		Probes: []donor.BARProbeResult{
			{Offset: 0x00, Original: 0xDEADBEEF, RWMask: 0xFF00FF00, MaybeRW1C: true},
		},
	}
	model := SynthesizeBARModel(profile, 0x010802)
	if model == nil {
		t.Fatal("SynthesizeBARModel returned nil")
	}
	if model.Registers[0].RWMask != 0 {
		t.Error("RW1C register should have RWMask=0 (conservative)")
	}
}


func TestBuildEthernetBARModel(t *testing.T) {
	barData := make([]byte, 4096)
	barData[0] = 0x00 // CTRL
	barData[8] = 0x02 // STATUS

	model := buildEthernetBARModel(barData)
	if model == nil {
		t.Fatal("buildEthernetBARModel returned nil")
	}
	if len(model.Registers) == 0 {
		t.Error("Ethernet BAR model should have registers")
	}
	// Check that CTRL register exists
	found := false
	for _, r := range model.Registers {
		if r.Name == "CTRL" && r.Offset == 0x0000 {
			found = true
			break
		}
	}
	if !found {
		t.Error("Should have CTRL register at offset 0x0000")
	}
}

func TestBuildEthernetBARModel_NilData(t *testing.T) {
	model := buildEthernetBARModel(nil)
	if model == nil {
		t.Fatal("Ethernet model should not be nil with nil data")
	}
	if len(model.Registers) == 0 {
		t.Error("Should still create register definitions without data")
	}
}

func TestBuildBARModel_Ethernet(t *testing.T) {
	barData := make([]byte, 4096)
	model := BuildBARModel(barData, 0x020000, nil)
	if model == nil {
		t.Fatal("BuildBARModel for Ethernet should not be nil")
	}
}

func TestBuildBARModel_WithProfile(t *testing.T) {
	profile := &donor.BARProfile{
		Size: 4096,
		Probes: []donor.BARProbeResult{
			{Offset: 0x00, Original: 0x12345678, RWMask: 0xFFFF0000},
			{Offset: 0x04, Original: 0x00, RWMask: 0x00},
		},
	}
	model := BuildBARModel(nil, 0x020000, profile)
	if model == nil {
		t.Fatal("BuildBARModel with profile should not be nil")
	}
}

func TestSynthesizeBARModel_WithProbe(t *testing.T) {
	profile := &donor.BARProfile{
		Size: 4096,
		Probes: []donor.BARProbeResult{
			{Offset: 0x00, Original: 0x12345678, RWMask: 0xFFFF0000},
			{Offset: 0x04, Original: 0x00, RWMask: 0x00},
			{Offset: 0x08, Original: 0xABCD, RWMask: 0xFF},
		},
	}
	model := SynthesizeBARModel(profile, 0x020000)
	if model == nil {
		t.Fatal("SynthesizeBARModel should not be nil")
	}
	// Dead registers (zero value + zero mask) should be dropped
	for _, r := range model.Registers {
		if r.Offset == 0x04 {
			t.Error("Dead register at 0x04 should be dropped")
		}
	}
}

func TestClassRegisterNames(t *testing.T) {
	names := classRegisterNames(0x020000) // Ethernet
	if names == nil {
		t.Skip("No class-specific names for Ethernet")
	}
	// Just verify it doesn't panic
}

func TestClassRegisterNames_NVMe(t *testing.T) {
	names := classRegisterNames(0x010802) // NVMe
	if names == nil {
		t.Skip("No class-specific names for NVMe")
	}
}

func TestPopulateResetValues(t *testing.T) {
	regs := []BARRegister{
		{Offset: 0x00, Width: 4, Reset: 0},
		{Offset: 0x04, Width: 4, Reset: 0},
	}
	barData := []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}
	populateResetValues(regs, barData)
	if regs[0].Reset != 0x04030201 {
		t.Errorf("Reset[0] = 0x%08x, want 0x04030201", regs[0].Reset)
	}
}

func TestPopulateResetValues_NilData(t *testing.T) {
	regs := []BARRegister{
		{Offset: 0x00, Width: 4, Reset: 0},
	}
	populateResetValues(regs, nil)
	// Should not panic, values stay zero
	if regs[0].Reset != 0 {
		t.Errorf("Reset = 0x%08x with nil data, want 0", regs[0].Reset)
	}
}
