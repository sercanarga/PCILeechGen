package barmodel

import (
	"reflect"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/donor/mmio"
)

func TestBARRegisterAccessKind(t *testing.T) {
	tests := []struct {
		name string
		reg  BARRegister
		want RegisterAccessKind
	}{
		{name: "read-only", reg: BARRegister{RWMask: 0}, want: RegisterReadOnly},
		{name: "read-write", reg: BARRegister{RWMask: 0xFFFFFFFF}, want: RegisterReadWrite},
		{name: "rw1c", reg: BARRegister{RWMask: 0x0000000F, IsRW1C: true}, want: RegisterRW1C},
		{name: "fsm", reg: BARRegister{RWMask: 0xFFFFFFFF, IsFSMDriven: true}, want: RegisterFSMDriven},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.reg.AccessKind(); got != tt.want {
				t.Fatalf("AccessKind() = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestApplyTraceOverlay_InfersWriteMaskForExistingRegister(t *testing.T) {
	model := &BARModel{
		Registers: []BARRegister{{
			Offset: 0x1C,
			Name:   "CSTS",
		}},
	}
	overlay := &mmio.TraceBAROverlay{
		WriteMask: map[uint32]uint32{0x1C: 0x0000000F},
		RW1CMask:  map[uint32]uint32{0x1C: 0x0000000F},
	}

	applyTraceOverlay(model, overlay)

	if len(model.Registers) != 1 {
		t.Fatalf("register count = %d, want 1", len(model.Registers))
	}
	reg := model.Registers[0]
	if reg.RWMask != 0x0000000F {
		t.Fatalf("RWMask = 0x%08X, want 0x0000000F", reg.RWMask)
	}
	if !reg.IsRW1C {
		t.Fatal("RW1C flag should be set from overlay")
	}
	if reg.RW1CMask != 0x0000000F {
		t.Fatalf("RW1CMask = 0x%08X, want 0x0000000F", reg.RW1CMask)
	}
}

func TestApplyTraceOverlay_AddsWritableOnlyRegister(t *testing.T) {
	model := &BARModel{}
	overlay := &mmio.TraceBAROverlay{WriteMask: map[uint32]uint32{0x100: 0x000000F0}}

	applyTraceOverlay(model, overlay)

	if len(model.Registers) != 1 {
		t.Fatalf("register count = %d, want 1", len(model.Registers))
	}
	reg := model.Registers[0]
	if reg.Offset != 0x100 {
		t.Fatalf("offset = 0x%X, want 0x100", reg.Offset)
	}
	if reg.RWMask != 0x000000F0 || reg.Name != "TRACE_WR_0x00000100" {
		t.Fatalf("write-only inferred register mismatch: %#v", reg)
	}
}

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

func TestBuildBARModel_NVMe_ProfiledMandatoryFSMRegs(t *testing.T) {
	profile := &donor.BARProfile{
		Size: 0x1000,
		Probes: []donor.BARProbeResult{
			{Offset: 0x14, Original: 0x00000001, RWMask: 0xFFFFFFFF},
			{Offset: 0x24, Original: 0x00000000, RWMask: 0xFFFFFFFF},
		},
	}

	model := BuildBARModelForDeviceIDWithOverlay(nil, 0x010802, 0, 0, profile, nil)
	if model == nil {
		t.Fatal("NVMe profiled model returned nil")
	}

	byOff := map[uint32]BARRegister{}
	for _, reg := range model.Registers {
		byOff[reg.Offset] = reg
	}
	for _, off := range []uint32{0x14, 0x1C, 0x24, 0x28, 0x2C, 0x30, 0x34} {
		if _, ok := byOff[off]; !ok {
			t.Fatalf("mandatory NVMe register 0x%X missing", off)
		}
	}
	if !byOff[0x1C].IsFSMDriven {
		t.Fatal("NVMe CSTS must stay FSM-driven")
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
			if !reg.IsFSMDriven {
				t.Fatalf("%s must be FSM-driven to avoid duplicate generic writers", reg.Name)
			}
			if reg.Reset != 0x000002A0 {
				t.Fatalf("%s reset = 0x%08X, want 0x000002A0", reg.Name, reg.Reset)
			}
		}
	}
	if portFound != 2 {
		t.Fatalf("xHCI model PORTSC count = %d, want 2", portFound)
	}
}

func TestBuildBARModel_XHCI_ProfiledMandatoryFSMRegs(t *testing.T) {
	profile := &donor.BARProfile{
		Size: 0x1000,
		Probes: []donor.BARProbeResult{
			{Offset: 0x20, Original: 0x00000001, RWMask: 0xFFFFFFFF},
			{Offset: 0x420, Original: 0x000002A0, RWMask: 0xFFFFFFFF},
		},
	}

	model := BuildBARModelForDeviceIDWithOverlay(nil, 0x0C0330, 0, 0, profile, nil)
	if model == nil {
		t.Fatal("xHCI profiled model returned nil")
	}

	byOff := map[uint32]BARRegister{}
	for _, reg := range model.Registers {
		byOff[reg.Offset] = reg
	}
	for _, off := range []uint32{0x20, 0x24, 0x220, 0x420, 0x430} {
		reg, ok := byOff[off]
		if !ok {
			t.Fatalf("mandatory xHCI FSM reg 0x%X missing", off)
		}
		if !reg.IsFSMDriven {
			t.Fatalf("mandatory xHCI FSM reg 0x%X not marked FSM-driven", off)
		}
	}
}

func TestBuildBARModel_XHCI_AllRegisters(t *testing.T) {
	model := BuildBARModel(nil, 0x0C0330, nil)
	names := make(map[string]bool)
	for _, r := range model.Registers {
		names[r.Name] = true
	}
	for _, required := range []string{
		"CAPLENGTH_HCIVERSION", "HCSPARAMS1", "HCCPARAMS1", "USBCMD", "USBSTS",
		"PAGESIZE", "CRCR_LO", "DCBAAP_LO", "CONFIG", "IMAN0", "ERDP_LO",
	} {
		if !names[required] {
			t.Errorf("xHCI model missing register: %s", required)
		}
	}
}

func TestBuildBARModelForDevice_RTS522A_FromProfileFallback(t *testing.T) {
	model := BuildBARModelForDeviceIDWithOverlay(nil, 0xFF0000, 0x10EC, 0x522A, nil, nil)
	if model == nil {
		t.Fatal("RTS522A BuildBARModelForDeviceIDWithOverlay returned nil")
	}
	names := map[string]bool{}
	for _, reg := range model.Registers {
		names[reg.Name] = true
	}
	for _, required := range []string{"RTSX_HCBAR", "RTSX_BIPR", "RTSX_BIER"} {
		if !names[required] {
			t.Errorf("RTS522A fallback model missing register: %s", required)
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

func TestBuildBARModel_Ethernet_4KB(t *testing.T) {
	model := buildEthernetBARModel(nil)
	if model.Size != 0 {
		t.Errorf("Ethernet BAR size: got %d, want 0 (no donor data, no longer forced 4096)", model.Size)
	}
}

func TestBuildBARModel_Ethernet_MACAddress(t *testing.T) {
	model := buildEthernetBARModel(nil)
	var mac03Found, mac45Found bool
	for _, reg := range model.Registers {
		switch reg.Name {
		case "MAC0_3":
			mac03Found = true
			if reg.Offset != 0x00 {
				t.Errorf("MAC0_3 offset: got 0x%X, want 0x00", reg.Offset)
			}
			if reg.Reset == 0 {
				t.Error("MAC0_3 should have a non-zero default MAC")
			}
		case "MAC4_5":
			mac45Found = true
			if reg.Offset != 0x04 {
				t.Errorf("MAC4_5 offset: got 0x%X, want 0x04", reg.Offset)
			}
			if reg.Reset == 0 {
				t.Error("MAC4_5 should have a non-zero default")
			}
		}
	}
	if !mac03Found || !mac45Found {
		t.Error("Ethernet model should include MAC0_3 and MAC4_5 registers")
	}
}

func TestBuildBARModel_Ethernet_ChipCmd(t *testing.T) {
	model := buildEthernetBARModel(nil)
	for _, reg := range model.Registers {
		if reg.Name == "CHIPCMD_DW" {
			if reg.Reset&0x0C000000 == 0 {
				t.Error("ChipCmd should have RxEnable + TxEnable set")
			}
			return
		}
	}
	t.Error("CHIPCMD_DW register not found")
}

func TestBuildBARModel_Ethernet_PHYStatus(t *testing.T) {
	model := buildEthernetBARModel(nil)
	for _, reg := range model.Registers {
		if reg.Name == "PHYSTATUS" {
			if reg.Reset&0x00003010 == 0 {
				t.Error("PHYStatus should have link up + 2500Mbps + full-duplex")
			}
			return
		}
	}
	t.Error("PHYSTATUS register not found")
}

func TestBuildBARModel_Ethernet_TxConfig(t *testing.T) {
	model := buildEthernetBARModel(nil)
	for _, reg := range model.Registers {
		if reg.Name == "TXCONFIG" {
			if reg.Reset&0x2F000000 == 0 {
				t.Error("TxConfig should have RTL8125B chip version")
			}
			return
		}
	}
	t.Error("TXCONFIG register not found")
}

func TestBuildBARModel_Ethernet_ERIAR(t *testing.T) {
	model := buildEthernetBARModel(nil)
	for _, reg := range model.Registers {
		if reg.Name == "ERIAR" {
			if reg.Reset&0x80000000 == 0 {
				t.Error("ERIAR should have completed bit set")
			}
			return
		}
	}
	t.Error("ERIAR register not found")
}

func TestBuildBARModel_Ethernet_PHYAR(t *testing.T) {
	model := buildEthernetBARModel(nil)
	for _, reg := range model.Registers {
		if reg.Name == "PHYAR" {
			if reg.Offset != 0xDC {
				t.Errorf("PHYAR offset: got 0x%X, want 0xDC", reg.Offset)
			}
			if reg.Reset&0x80000000 == 0 {
				t.Error("PHYAR should have ready bit set")
			}
			return
		}
	}
	t.Error("PHYAR register not found")
}

func TestBuildBARModel_Ethernet_DWORDAligned(t *testing.T) {
	model := buildEthernetBARModel(nil)
	for _, reg := range model.Registers {
		if reg.Offset%4 != 0 {
			t.Errorf("register %s at offset 0x%X is not DWORD-aligned (must be multiple of 4)", reg.Name, reg.Offset)
		}
	}
}

func TestBuildBARModel_Ethernet_NoDuplicateAlignedOffsets(t *testing.T) {
	model := buildEthernetBARModel(nil)
	seen := make(map[uint32]string)
	for _, reg := range model.Registers {
		aligned := (reg.Offset / 4) * 4
		if prev, ok := seen[aligned]; ok {
			t.Errorf("registers %s and %s both map to aligned offset 0x%X - SV case conflict", prev, reg.Name, aligned)
		}
		seen[aligned] = reg.Name
	}
}

func TestBuildBARModel_Audio(t *testing.T) {
	model := BuildBARModel(nil, 0x040300, nil)
	if model == nil {
		t.Fatal("Audio BuildBARModel returned nil")
	}
	if model.Size != 0 {
		t.Errorf("Audio BAR size: got %d, want 0 (no donor data, no longer forced 4096)", model.Size)
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
			// GCAP should be 0x6401 in the lower 16 bits (includes B64OK bit 13)
			if reg.Reset&0xFFFF != 0x6401 {
				t.Errorf("GCAP portion: got 0x%04X, want 0x6401", reg.Reset&0xFFFF)
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
				t.Error("CORBCTL_STS_SIZE should have non-zero default (CORBSIZE=0x82)")
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
			if reg.Reset != 0x01006401 {
				t.Logf("GCAP from donor data: 0x%08X", reg.Reset)
			}
			return
		}
	}
}

func TestBuildBARModel_Audio_AllFFDonor(t *testing.T) {
	// Simulate no codec connected - BAR reads as all 0xFF
	barData := make([]byte, 4096)
	for i := range barData {
		barData[i] = 0xFF
	}
	model := buildAudioBARModel(barData)
	if model == nil {
		t.Fatal("model should not be nil")
	}
	for _, reg := range model.Registers {
		if reg.Reset == 0xFFFFFFFF {
			t.Errorf("register %s @ 0x%X still has 0xFFFFFFFF reset value", reg.Name, reg.Offset)
		}
	}
	// Spot check key registers
	for _, reg := range model.Registers {
		switch reg.Offset {
		case 0x00:
			if reg.Reset != 0x01006401 {
				t.Errorf("GCAP_VMIN_VMAJ: expected 0x01006401, got 0x%08X", reg.Reset)
			}
		case 0x08:
			if reg.Reset != 0x00000001 {
				t.Errorf("GCTL: expected 0x00000001, got 0x%08X", reg.Reset)
			}
		case 0x0C:
			if reg.Reset != 0x00010000 {
				t.Errorf("WAKEEN_STATESTS: expected 0x00010000, got 0x%08X", reg.Reset)
			}
		case 0x4C:
			if reg.Reset != 0x00820000 {
				t.Errorf("CORBCTL_STS_SIZE: expected 0x00820000, got 0x%08X", reg.Reset)
			}
		case 0x5C:
			if reg.Reset != 0x00820000 {
				t.Errorf("RIRBCTL_STS_SIZE: expected 0x00820000, got 0x%08X", reg.Reset)
			}
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
	if len(names) != 0 {
		t.Error("unknown class should return empty names map")
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

func TestBuildBARModel_UnreliableProbe_FallsBackToSpec(t *testing.T) {
	// Simulate VFIO returning all-RW for every register (Samsung NVMe scenario)
	probes := make([]donor.BARProbeResult, 20)
	for i := range probes {
		probes[i] = donor.BARProbeResult{
			Offset:   uint32(i * 4),
			Original: uint32(0x28033FFF + i),
			RWMask:   0xFFFFFFFF, // all writable - unreliable!
		}
	}
	profile := &donor.BARProfile{Size: 4096, Probes: probes}

	barData := make([]byte, 4096)
	model := BuildBARModel(barData, 0x010802, profile) // NVMe class

	if model == nil {
		t.Fatal("should fall back to spec-based NVMe model, got nil")
	}

	// verify CSTS is read-only (CC->CSTS handshake FSM relies on this)
	for _, reg := range model.Registers {
		if reg.Name == "CSTS" {
			if reg.RWMask != 0 {
				t.Errorf("CSTS should be RO (RWMask=0) for CC->CSTS handshake, got 0x%08X", reg.RWMask)
			}
			return
		}
	}
	t.Error("spec-based NVMe model should contain CSTS register")
}

func TestIsProbeDataReliable_Mixed(t *testing.T) {
	profile := &donor.BARProfile{
		Size: 4096,
		Probes: []donor.BARProbeResult{
			{Offset: 0x00, Original: 0x12345678, RWMask: 0x00000000},
			{Offset: 0x04, Original: 0xDEADBEEF, RWMask: 0xFFFF0000},
			{Offset: 0x08, Original: 0x00000001, RWMask: 0xFFFFFFFF},
			{Offset: 0x0C, Original: 0x00000000, RWMask: 0x00000000}, // dead
		},
	}
	if !isProbeDataReliable(profile) {
		t.Error("mixed RW masks should be considered reliable")
	}
}

func TestIsProbeDataReliable_AllRW(t *testing.T) {
	probes := make([]donor.BARProbeResult, 10)
	for i := range probes {
		probes[i] = donor.BARProbeResult{
			Offset:   uint32(i * 4),
			Original: uint32(i + 1),
			RWMask:   0xFFFFFFFF,
		}
	}
	profile := &donor.BARProfile{Size: 4096, Probes: probes}
	if isProbeDataReliable(profile) {
		t.Error("all-RW probes should be considered unreliable")
	}
}

func TestApplyDonorProbeOverlay_Writes(t *testing.T) {
	// spec-table NVMe model; donor probe says CC has extra writable bits.
	barData := make([]byte, 4096)
	model := BuildBARModel(barData, 0x010802, nil)
	if model == nil {
		t.Fatal("nil model")
	}
	profile := &donor.BARProfile{BarIndex: 0, Size: 4096}
	// CC is at 0x14, spec RWMask 0x00FFFFF1; donor says 0x00FFFFF9 (extra bit 3)
	profile.Probes = append(profile.Probes, donor.BARProbeResult{Offset: 0x14, Original: 0x00460001, RWMask: 0x00FFFFF9})
	profile.Probes = append(profile.Probes, donor.BARProbeResult{Offset: 0x24, Original: 0x00000000, RWMask: 0x00000000})
	// CSTS at 0x1C is IsFSMDriven -> must NOT be overridden even if probe says writable
	profile.Probes = append(profile.Probes, donor.BARProbeResult{Offset: 0x1C, Original: 0x00000001, RWMask: 0xFFFFFFFF})
	applyDonorProbeOverlay(model, profile)
	for _, r := range model.Registers {
		switch r.Offset {
		case 0x14:
			if r.RWMask != 0x00FFFFF9 {
				t.Errorf("CC RWMask: got 0x%08X, want 0x00FFFFF9 (donor overlay)", r.RWMask)
			}
		case 0x24:
			if r.RWMask != 0 || r.Reset != 0 {
				t.Errorf("AQA should accept zero-valued donor overlay, got reset=0x%08X rw=0x%08X", r.Reset, r.RWMask)
			}
		case 0x1C:
			if r.RWMask != 0x00000000 {
				t.Errorf("CSTS is IsFSMDriven; RWMask must stay 0, got 0x%08X", r.RWMask)
			}
		}
	}
}

func TestApplyDonorProbeOverlay_RW1C(t *testing.T) {
	barData := make([]byte, 4096)
	model := BuildBARModel(barData, 0x0C0330, nil) // xHCI
	if model == nil {
		t.Fatal("nil model")
	}
	profile := &donor.BARProfile{BarIndex: 0, Size: 4096}
	// A non-FSM register (DBOFF 0x14) donor-detected RW1C
	profile.Probes = append(profile.Probes, donor.BARProbeResult{Offset: 0x14, Original: 0x00000020, RWMask: 0x000000FF, MaybeRW1C: true})
	applyDonorProbeOverlay(model, profile)
	for _, r := range model.Registers {
		if r.Offset == 0x14 {
			if !r.IsRW1C {
				t.Error("DBOFF should inherit IsRW1C from donor probe")
			}
			if r.RWMask != 0 {
				t.Error("RW1C register should be modeled read-only (RWMask=0)")
			}
		}
	}
}

func TestPopulateStaticShadow(t *testing.T) {
	barData := make([]byte, 4096)
	// NVMe model covers 0x00-0x34. Put a non-zero static word at 0x100 (unmodeled).
	utilWriteLE32(barData, 0x100, 0xCAFEBABE)
	// And a word inside the modeled range (0x00) that should NOT appear in shadow.
	utilWriteLE32(barData, 0x00, 0x11223344)
	model := BuildBARModel(barData, 0x010802, nil)
	if model == nil {
		t.Fatal("nil model")
	}
	found := false
	for _, w := range model.StaticShadow {
		if w.Offset == 0x100 && w.Value == 0xCAFEBABE {
			found = true
		}
		if w.Offset == 0x00 {
			t.Error("modeled offset 0x00 must not appear in static shadow")
		}
	}
	if !found {
		t.Error("static shadow should contain donor word at 0x100")
	}
}

func utilWriteLE32(b []byte, off int, v uint32) {
	b[off] = byte(v)
	b[off+1] = byte(v >> 8)
	b[off+2] = byte(v >> 16)
	b[off+3] = byte(v >> 24)
}

func TestBuildBARModelForDevice_Atheros_SequentialRead(t *testing.T) {
	// AR9287 is class 0x028000 (network, other), vendor 0x168C (Atheros).
	// Use a donor barData that has the EEPROM handshake sentinel at 0x4010.
	barData := make([]byte, 0x4080)
	utilWriteLE32(barData, 0x4010, 0xDEADBEEF)
	utilWriteLE32(barData, 0x407C, 0x0000BEEF)

	model := BuildBARModelForDeviceIDWithOverlay(barData, 0x028000, 0x168C, 0x002E, nil, nil)
	if model == nil {
		t.Fatal("AR9287 model should not be nil")
	}
	var req, data *BARRegister
	for i := range model.Registers {
		if model.Registers[i].Offset == 0x4010 {
			req = &model.Registers[i]
		}
		if model.Registers[i].Offset == 0x407C {
			data = &model.Registers[i]
		}
	}
	if req == nil {
		t.Fatal("EEPROM_REQ (0x4010) not in model")
	}
	if !req.SequentialRead {
		t.Error("EEPROM_REQ should be SequentialRead")
	}
	if len(req.SequentialValues) != 3 {
		t.Fatalf("EEPROM_REQ sequence len = %d, want 3", len(req.SequentialValues))
	}
	if req.SequentialValues[0] != 0xDEADBEEF || req.SequentialValues[2] != 0x00000001 {
		t.Fatalf("EEPROM_REQ sequence = %#v, want donor-like pending->done sequence", req.SequentialValues)
	}
	if data == nil {
		t.Fatal("EEPROM_DATA (0x407C) not in model")
	}
	if !data.SequentialRead {
		t.Error("EEPROM_DATA should be SequentialRead")
	}
	if len(data.SequentialValues) != 4 {
		t.Fatalf("EEPROM_DATA sequence len = %d, want 4", len(data.SequentialValues))
	}
	if data.SequentialValues[0] != 0x0000BEEF {
		t.Fatalf("EEPROM_DATA first sequence value = 0x%08X, want donor reset 0x0000BEEF", data.SequentialValues[0])
	}
	// Ensure no duplicate static-shadow entries at the handshake offsets
	// (would produce duplicate case labels in the generated SV).
	for _, sw := range model.StaticShadow {
		if sw.Offset == 0x4010 || sw.Offset == 0x407C {
			t.Errorf("static shadow should not include modeled handshake offset 0x%X", sw.Offset)
		}
	}
}

func TestApplyProfileSequentialRead_ExistingRegKeepsReset(t *testing.T) {
	// Probe path already discovered 0x4010 with the sentinel. The profile
	// overlay must set the flag but keep the donor-observed reset.
	profile := &donor.BARProfile{
		Size: 0x4080,
		Probes: []donor.BARProbeResult{
			{Offset: 0x4010, Original: 0xCAFEBABE, RWMask: 0x00000000},
		},
	}
	model := BuildBARModelForDeviceIDWithOverlay(nil, 0x028000, 0x168C, 0x002E, profile, nil)
	if model == nil {
		t.Fatal("model should not be nil")
	}
	var req *BARRegister
	for i := range model.Registers {
		if model.Registers[i].Offset == 0x4010 {
			req = &model.Registers[i]
		}
	}
	if req == nil {
		t.Fatal("EEPROM_REQ not in model")
	}
	if !req.SequentialRead {
		t.Error("EEPROM_REQ should be SequentialRead")
	}
	if req.Reset != 0xCAFEBABE {
		t.Errorf("donor reset should be kept: got 0x%08X, want 0xCAFEBABE", req.Reset)
	}
	if len(req.SequentialValues) == 0 || req.SequentialValues[0] != 0xCAFEBABE {
		t.Fatalf("donor reset should seed sequential ROM[0], got %#v", req.SequentialValues)
	}
}

func TestApplyTraceOverlay_RefinesExistingSequentialRegister(t *testing.T) {
	model := &BARModel{Registers: []BARRegister{{
		Offset: 0x10, Name: "EEPROM_REQ", Reset: 0xAAAAAAAA, SequentialRead: true,
	}}}
	overlay := &mmio.TraceBAROverlay{
		Sequential: map[uint32][]uint32{0x10: {0xDEADBEEF, 0x00000002, 0x00000001}},
	}

	applyTraceOverlay(model, overlay)

	got := model.Registers[0].SequentialValues
	want := []uint32{0xDEADBEEF, 0x00000002, 0x00000001}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("seq values = %#v, want %#v", got, want)
	}
}

func TestApplyTraceOverlay_AddsSequentialRegisterForUnmodeledOffset(t *testing.T) {
	model := &BARModel{Registers: []BARRegister{{Offset: 0x00, Name: "CAP"}}}
	overlay := &mmio.TraceBAROverlay{
		Sequential: map[uint32][]uint32{0x100: {0xAAAA0001, 0xAAAA0002}},
		Static:     map[uint32]uint32{0x100: 0xBBBB0000},
	}

	applyTraceOverlay(model, overlay)

	for _, reg := range model.Registers {
		if reg.Offset == 0x100 {
			if !reg.SequentialRead {
				t.Fatal("trace-only register should be SequentialRead")
			}
			if !reflect.DeepEqual(reg.SequentialValues, []uint32{0xAAAA0001, 0xAAAA0002}) {
				t.Fatalf("trace-only sequence = %#v", reg.SequentialValues)
			}
			if reg.Reset != 0xAAAA0001 {
				t.Fatalf("trace-only reset = 0x%08X", reg.Reset)
			}
			for _, sw := range model.StaticShadow {
				if sw.Offset == 0x100 {
					t.Fatal("trace-only sequential register should not also be static shadow")
				}
			}
			return
		}
	}
	t.Fatal("trace-only sequential register missing")
}

func TestApplyTraceOverlay_SkipsFSMDrivenSequentialRegister(t *testing.T) {
	model := &BARModel{Registers: []BARRegister{{
		Offset: 0x24, Name: "USBSTS", IsFSMDriven: true,
	}}}
	overlay := &mmio.TraceBAROverlay{
		Sequential: map[uint32][]uint32{0x24: {0x00000010, 0x00000018}},
	}

	applyTraceOverlay(model, overlay)

	if model.Registers[0].SequentialRead || len(model.Registers[0].SequentialValues) != 0 {
		t.Fatalf("FSM-driven register should not get trace sequence: %#v", model.Registers[0])
	}
}

func TestBuildBARModelForDevice_RealtekWiFiUsesProfile(t *testing.T) {
	model := BuildBARModelForDeviceIDWithOverlay(nil, 0x028000, 0x10EC, 0x8179, nil, nil)
	if model == nil {
		t.Fatal("model should not be nil")
	}
	names := map[string]bool{}
	for _, reg := range model.Registers {
		names[reg.Name] = true
	}
	for _, want := range []string{"HIMR", "HISR", "PCIE_CTRL_INT_MIG"} {
		if !names[want] {
			t.Fatalf("Realtek model missing %s; got %#v", want, names)
		}
	}
	if names["MAC0_3"] {
		t.Fatal("Realtek Wi-Fi should not fall back to Ethernet BAR model")
	}
}

func TestApplyTraceOverlay_AddsStaticShadowForUnmodeledOffset(t *testing.T) {
	model := &BARModel{Registers: []BARRegister{{Offset: 0x00, Name: "CAP"}}}
	overlay := &mmio.TraceBAROverlay{Static: map[uint32]uint32{0x100: 0xCAFEBABE}}

	applyTraceOverlay(model, overlay)

	if len(model.StaticShadow) != 1 || model.StaticShadow[0].Offset != 0x100 {
		t.Fatalf("static shadow = %#v", model.StaticShadow)
	}
}

func TestPopulateStaticShadowPreservesTraceStatic(t *testing.T) {
	barData := make([]byte, 0x200)
	utilWriteLE32(barData, 0x100, 0x11111111)
	model := &BARModel{
		Registers:    []BARRegister{{Offset: 0x00, Name: "CAP"}},
		StaticShadow: []StaticWord{{Offset: 0x180, Value: 0x22222222}},
	}

	populateStaticShadow(model, barData)

	found := map[uint32]uint32{}
	for _, sw := range model.StaticShadow {
		found[sw.Offset] = sw.Value
	}
	if found[0x180] != 0x22222222 {
		t.Fatalf("trace static shadow lost: %#v", model.StaticShadow)
	}
	if found[0x100] != 0x11111111 {
		t.Fatalf("donor static shadow missing: %#v", model.StaticShadow)
	}
}
