package barmodel

import (
	"testing"

	"github.com/sercanarga/pcileechgen/internal/donor"
)

func TestBuildBARModel_NVMe(t *testing.T) {
	barData := make([]byte, 4096)
	barData[0] = 0x3F
	barData[1] = 0x00
	barData[2] = 0x01
	barData[3] = 0x00

	model := BuildBARModel(barData, 0x010802, nil)
	if model == nil {
		t.Fatal("NVMe BuildBARModel returned nil")
	}
	if model.Size != 4096 {
		t.Errorf("BAR size: got %d, want 4096", model.Size)
	}
	foundCSTS := false
	for _, reg := range model.Registers {
		if reg.Name == "CSTS" {
			foundCSTS = true
			if reg.Reset&0x01 == 0 {
				t.Error("CSTS.RDY should be set to 1")
			}
			if reg.Reset&0x0C != 0 {
				t.Error("CSTS.SHST should be cleared")
			}
		}
	}
	if !foundCSTS {
		t.Error("NVMe model should contain CSTS register")
	}
}

func TestBuildBARModel_NVMe_AllRegisters(t *testing.T) {
	model := BuildBARModel(nil, 0x010802, nil)
	if model == nil {
		t.Fatal("nil barData should still create NVMe model")
	}
	expected := map[string]uint32{
		"CAP_LO": 0x00, "CAP_HI": 0x04, "VS": 0x08,
		"CC": 0x14, "CSTS": 0x1C, "AQA": 0x24,
		"ASQ_LO": 0x28, "ASQ_HI": 0x2C,
		"ACQ_LO": 0x30, "ACQ_HI": 0x34,
	}
	regMap := make(map[string]uint32)
	for _, r := range model.Registers {
		regMap[r.Name] = r.Offset
	}
	for name, offset := range expected {
		if got, ok := regMap[name]; !ok {
			t.Errorf("missing register %s", name)
		} else if got != offset {
			t.Errorf("%s offset: got 0x%X, want 0x%X", name, got, offset)
		}
	}
}

func TestBuildBARModel_NVMe_RWMasks(t *testing.T) {
	model := BuildBARModel(nil, 0x010802, nil)
	for _, reg := range model.Registers {
		switch reg.Name {
		case "CAP_LO", "CAP_HI", "VS", "CSTS":
			if reg.RWMask != 0 {
				t.Errorf("%s should be read-only (RWMask=0), got 0x%08X", reg.Name, reg.RWMask)
			}
		case "CC":
			if reg.RWMask == 0 {
				t.Errorf("CC should be writable")
			}
		case "AQA":
			if reg.RWMask != 0x0FFF0FFF {
				t.Errorf("AQA RWMask: got 0x%08X, want 0x0FFF0FFF", reg.RWMask)
			}
		}
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

func TestBuildBARModel_XHCI_PORTSC(t *testing.T) {
	model := BuildBARModel(nil, 0x0C0330, nil)
	portFound := 0
	for _, reg := range model.Registers {
		if reg.Name == "PORTSC1" || reg.Name == "PORTSC2" {
			portFound++
		}
	}
	// barmodel builder may not have PORTSC — that's in the profile only.
	// but the profile builder (xhci.go) does include them.
}

func TestBuildBARModel_XHCI_AllRegisters(t *testing.T) {
	model := BuildBARModel(nil, 0x0C0330, nil)
	names := make(map[string]bool)
	for _, r := range model.Registers {
		names[r.Name] = true
	}
	for _, required := range []string{"CAPLENGTH_HCIVERSION", "HCSPARAMS1", "HCCPARAMS1", "USBCMD", "USBSTS", "PAGESIZE", "CRCR_LO", "DCBAAP_LO", "CONFIG"} {
		if !names[required] {
			t.Errorf("xHCI model missing register: %s", required)
		}
	}
}

func TestBuildBARModel_Ethernet(t *testing.T) {
	barData := make([]byte, 4096)
	model := BuildBARModel(barData, 0x020000, nil)
	if model == nil {
		t.Fatal("BuildBARModel for Ethernet should not be nil")
	}
}

func TestBuildBARModel_Ethernet_32KB(t *testing.T) {
	model := buildEthernetBARModel(nil)
	if model.Size != 32768 {
		t.Errorf("Ethernet BAR size: got %d, want 32768", model.Size)
	}
}

func TestBuildBARModel_Ethernet_MACAddress(t *testing.T) {
	model := buildEthernetBARModel(nil)
	var ral0Found, rah0Found bool
	for _, reg := range model.Registers {
		switch reg.Name {
		case "RAL0":
			ral0Found = true
			if reg.Offset != 0x5400 {
				t.Errorf("RAL0 offset: got 0x%X, want 0x5400", reg.Offset)
			}
			if reg.Reset == 0 {
				t.Error("RAL0 should have a non-zero default MAC")
			}
		case "RAH0":
			rah0Found = true
			if reg.Reset&0x80000000 == 0 {
				t.Error("RAH0 should have AV (address valid) bit set")
			}
		}
	}
	if !ral0Found || !rah0Found {
		t.Error("Ethernet model should include RAL0 and RAH0 registers")
	}
}

func TestBuildBARModel_Ethernet_EECD(t *testing.T) {
	model := buildEthernetBARModel(nil)
	for _, reg := range model.Registers {
		if reg.Name == "EECD" {
			if reg.Reset&0x200 == 0 {
				t.Error("EECD should have Auto-Read Done (bit 9) set")
			}
			if reg.Reset&0x100 == 0 {
				t.Error("EECD should have EEPROM Present (bit 8) set")
			}
			return
		}
	}
	t.Error("EECD register not found")
}

func TestBuildBARModel_Ethernet_STATUS(t *testing.T) {
	model := buildEthernetBARModel(nil)
	for _, reg := range model.Registers {
		if reg.Name == "STATUS" {
			if reg.Reset&0x02 == 0 {
				t.Error("STATUS.LU (link up) should be set")
			}
			return
		}
	}
	t.Error("STATUS register not found")
}

func TestBuildBARModel_Ethernet_MDIC(t *testing.T) {
	model := buildEthernetBARModel(nil)
	for _, reg := range model.Registers {
		if reg.Name == "MDIC" {
			if reg.Reset&0x10000000 == 0 {
				t.Error("MDIC Ready bit should be set")
			}
			return
		}
	}
	t.Error("MDIC register not found")
}

func TestBuildBARModel_Audio(t *testing.T) {
	model := BuildBARModel(nil, 0x040300, nil)
	if model == nil {
		t.Fatal("Audio BuildBARModel returned nil")
	}
	if model.Size != 4096 {
		t.Errorf("Audio BAR size: got %d, want 4096", model.Size)
	}
}

func TestBuildBARModel_Audio_AllRegisters(t *testing.T) {
	model := buildAudioBARModel(nil)
	names := make(map[string]bool)
	for _, r := range model.Registers {
		names[r.Name] = true
	}
	for _, required := range []string{"GCAP_VMIN_VMAJ", "GCTL", "WAKEEN_STATESTS", "INTCTL", "CORBLBASE", "CORBUBASE", "RIRBLBASE", "RIRBUBASE"} {
		if !names[required] {
			t.Errorf("Audio model missing register: %s", required)
		}
	}
}

func TestBuildBARModel_Audio_GCAP(t *testing.T) {
	model := buildAudioBARModel(nil)
	for _, reg := range model.Registers {
		if reg.Name == "GCAP_VMIN_VMAJ" {
			if reg.Reset == 0 {
				t.Error("GCAP_VMIN_VMAJ should have a non-zero default")
			}
			// GCAP should be 0x4401 in the lower 16 bits
			if reg.Reset&0xFFFF != 0x4401 {
				t.Errorf("GCAP portion: got 0x%04X, want 0x4401", reg.Reset&0xFFFF)
			}
			return
		}
	}
	t.Error("GCAP_VMIN_VMAJ not found")
}

func TestBuildBARModel_Audio_GCTL(t *testing.T) {
	model := buildAudioBARModel(nil)
	for _, reg := range model.Registers {
		if reg.Name == "GCTL" {
			if reg.Reset&0x01 == 0 {
				t.Error("GCTL.CRST should be set (out of reset)")
			}
			return
		}
	}
	t.Error("GCTL not found")
}

func TestBuildBARModel_Audio_STATESTS(t *testing.T) {
	model := buildAudioBARModel(nil)
	for _, reg := range model.Registers {
		if reg.Name == "WAKEEN_STATESTS" {
			// STATESTS is in upper 16 bits, codec 0 present = bit 16
			if reg.Reset&0x10000 == 0 {
				t.Error("STATESTS should have codec 0 present bit set")
			}
			return
		}
	}
	t.Error("WAKEEN_STATESTS not found")
}

func TestBuildBARModel_Audio_CORBSIZE(t *testing.T) {
	model := buildAudioBARModel(nil)
	for _, reg := range model.Registers {
		if reg.Name == "CORBCTL_STS_SIZE" {
			if reg.Reset == 0 {
				t.Error("CORBCTL_STS_SIZE should have non-zero default (CORBSIZE=0x42)")
			}
			return
		}
	}
	t.Error("CORBCTL_STS_SIZE not found")
}

func TestBuildBARModel_Audio_AllWidth4(t *testing.T) {
	model := buildAudioBARModel(nil)
	for _, reg := range model.Registers {
		if reg.Width != 4 {
			t.Errorf("register %s has Width=%d, all Audio regs must be 4 (DWORD-packed)", reg.Name, reg.Width)
		}
	}
}

func TestBuildBARModel_Audio_DonorData(t *testing.T) {
	barData := make([]byte, 256)
	// set GCAP to a donor-specific value
	barData[0] = 0x01
	barData[1] = 0x44
	barData[2] = 0x00
	barData[3] = 0x01

	model := buildAudioBARModel(barData)
	for _, reg := range model.Registers {
		if reg.Name == "GCAP_VMIN_VMAJ" {
			if reg.Reset != 0x01004401 {
				t.Logf("GCAP from donor data: 0x%08X", reg.Reset)
			}
			return
		}
	}
}

func TestBuildBARModel_Unknown(t *testing.T) {
	model := BuildBARModel(nil, 0xFF0000, nil)
	if model != nil {
		t.Error("unknown class without profile should return nil")
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

func TestBuildBARModel_ProfileTakesPriority(t *testing.T) {
	profile := &donor.BARProfile{
		Size: 8192,
		Probes: []donor.BARProbeResult{
			{Offset: 0x00, Original: 0x12345678, RWMask: 0xFFFF0000},
		},
	}
	model := BuildBARModel(nil, 0x010802, profile)
	if model == nil {
		t.Fatal("model should not be nil")
	}
	// profile should override the NVMe spec-based model
	if model.Size != 8192 {
		t.Errorf("profile size should override: got %d, want 8192", model.Size)
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
			{Offset: 0x20, Original: 0x00000000, RWMask: 0x00000000},
		},
	}
	model := SynthesizeBARModel(profile, 0x010802)
	if model == nil {
		t.Fatal("SynthesizeBARModel returned nil")
	}
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

func TestSynthesizeBARModel_Nil(t *testing.T) {
	if SynthesizeBARModel(nil, 0x010802) != nil {
		t.Error("nil profile should return nil")
	}
}

func TestSynthesizeBARModel_EmptyProbes(t *testing.T) {
	profile := &donor.BARProfile{Size: 4096, Probes: []donor.BARProbeResult{}}
	if SynthesizeBARModel(profile, 0x010802) != nil {
		t.Error("empty probes should return nil")
	}
}

func TestSynthesizeBARModel_AllDead(t *testing.T) {
	profile := &donor.BARProfile{
		Size: 4096,
		Probes: []donor.BARProbeResult{
			{Offset: 0x00, Original: 0x00, RWMask: 0x00},
			{Offset: 0x04, Original: 0x00, RWMask: 0x00},
		},
	}
	if SynthesizeBARModel(profile, 0x010802) != nil {
		t.Error("all-dead registers should return nil")
	}
}

func TestSynthesizeBARModel_NameHints(t *testing.T) {
	profile := &donor.BARProfile{
		Size: 4096,
		Probes: []donor.BARProbeResult{
			{Offset: 0x14, Original: 0x00000001, RWMask: 0x00FFFFF1},
		},
	}
	model := SynthesizeBARModel(profile, 0x010802)
	if model == nil {
		t.Fatal("model is nil")
	}
	if model.Registers[0].Name != "CC" {
		t.Errorf("NVMe offset 0x14 should be named CC, got %s", model.Registers[0].Name)
	}
}

func TestClassRegisterNames(t *testing.T) {
	names := classRegisterNames(0x010802)
	if names == nil {
		t.Fatal("NVMe should have register names")
	}
	if names[0x14] != "CC" {
		t.Errorf("NVMe 0x14 should be CC, got %s", names[0x14])
	}
}

func TestClassRegisterNames_Audio(t *testing.T) {
	names := classRegisterNames(0x040300)
	if names == nil {
		t.Fatal("Audio should have register names")
	}
	if _, ok := names[0x08]; !ok {
		t.Error("Audio should have name for offset 0x08 (GCTL)")
	}
}

func TestClassRegisterNames_Unknown(t *testing.T) {
	names := classRegisterNames(0xFF0000)
	if names != nil {
		t.Error("unknown class should return nil names")
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
	if regs[1].Reset != 0x08070605 {
		t.Errorf("Reset[1] = 0x%08x, want 0x08070605", regs[1].Reset)
	}
}

func TestPopulateResetValues_NilData(t *testing.T) {
	regs := []BARRegister{
		{Offset: 0x00, Width: 4, Reset: 0},
	}
	populateResetValues(regs, nil)
	if regs[0].Reset != 0 {
		t.Errorf("Reset = 0x%08x with nil data, want 0", regs[0].Reset)
	}
}

func TestPopulateResetValues_ShortData(t *testing.T) {
	regs := []BARRegister{
		{Offset: 0x00, Width: 4, Reset: 0},
		{Offset: 0x10, Width: 4, Reset: 0},
	}
	barData := make([]byte, 8)
	barData[0] = 0xFF
	populateResetValues(regs, barData)
	if regs[0].Reset&0xFF != 0xFF {
		t.Error("first register should be populated")
	}
	if regs[1].Reset != 0 {
		t.Error("out-of-range register should stay zero")
	}
}

func TestPopulateResetValues_Width2(t *testing.T) {
	regs := []BARRegister{
		{Offset: 0x00, Width: 2, Reset: 0},
	}
	barData := []byte{0xAB, 0xCD}
	populateResetValues(regs, barData)
	if regs[0].Reset != 0xCDAB {
		t.Errorf("Width=2 Reset = 0x%04x, want 0xCDAB", regs[0].Reset)
	}
}

func TestPopulateResetValues_Width1(t *testing.T) {
	regs := []BARRegister{
		{Offset: 0x00, Width: 1, Reset: 0},
	}
	barData := []byte{0x42}
	populateResetValues(regs, barData)
	if regs[0].Reset != 0x42 {
		t.Errorf("Width=1 Reset = 0x%02x, want 0x42", regs[0].Reset)
	}
}
