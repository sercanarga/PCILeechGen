package devclass

import "testing"

func TestProfileForClass_NVMe(t *testing.T) {
	p := ProfileForClass(0x010802)
	if p == nil {
		t.Fatal("NVMe profile should not be nil")
	}
	if p.ClassName != "NVMe" {
		t.Errorf("got %q, want NVMe", p.ClassName)
	}
	if !p.PrefersMSIX {
		t.Error("NVMe should prefer MSI-X")
	}
	if len(p.BARDefaults) == 0 {
		t.Error("NVMe should have BAR defaults")
	}
	if !p.Uses64BitBAR {
		t.Error("NVMe should use 64-bit BAR")
	}
}

func TestProfileForClass_XHCI(t *testing.T) {
	p := ProfileForClass(0x0C0330)
	if p == nil {
		t.Fatal("xHCI profile should not be nil")
	}
	if p.ClassName != "xHCI USB" {
		t.Errorf("got %q, want xHCI USB", p.ClassName)
	}
	if !p.PrefersMSIX {
		t.Error("xHCI should prefer MSI-X")
	}
}

func TestProfileForClass_Ethernet(t *testing.T) {
	p := ProfileForClass(0x020000)
	if p == nil {
		t.Fatal("Ethernet profile should not be nil")
	}
	if p.MinMSIXVectors < 3 {
		t.Error("Ethernet should need at least 3 MSI-X vectors")
	}
	if !p.Uses64BitBAR {
		t.Error("Ethernet profile should use 64-bit BAR (RTL8125)")
	}
	if p.PreferredBAR != 2 {
		t.Errorf("Ethernet PreferredBAR should be 2, got %d", p.PreferredBAR)
	}
	if p.MinBARSize != 65536 {
		t.Errorf("Ethernet MinBARSize should be 65536, got %d", p.MinBARSize)
	}
}

func TestProfileForClass_Audio(t *testing.T) {
	p := ProfileForClass(0x040300)
	if p == nil {
		t.Fatal("HD Audio profile should not be nil")
	}
	if p.PrefersMSIX {
		t.Error("HD Audio should prefer MSI (not MSI-X)")
	}
	if p.Uses64BitBAR != false {
		t.Error("HD Audio should NOT use 64-bit BAR (scrub pipeline forces 32-bit)")
	}
}

func TestProfileForClass_GPU(t *testing.T) {
	p := ProfileForClass(0x030000)
	if p == nil {
		t.Fatal("GPU profile should not be nil")
	}
	if p.ClassName != "GPU" {
		t.Errorf("got %q, want GPU", p.ClassName)
	}
	if !p.BARIsPrefetchable {
		t.Error("GPU BAR should be prefetchable")
	}
	if !p.Uses64BitBAR {
		t.Error("GPU should use 64-bit BAR")
	}
}

func TestProfileForClass_SATA(t *testing.T) {
	p := ProfileForClass(0x010601)
	if p == nil {
		t.Fatal("SATA profile should not be nil")
	}
	if p.ClassName != "SATA AHCI" {
		t.Errorf("got %q, want SATA AHCI", p.ClassName)
	}
	if p.PreferredBAR != 5 {
		t.Errorf("SATA ABAR should be BAR5, got %d", p.PreferredBAR)
	}
}

func TestProfileForClass_WiFi(t *testing.T) {
	p := ProfileForClass(0x028000)
	if p == nil {
		t.Fatal("Wi-Fi profile should not be nil")
	}
	if p.ClassName != "Wi-Fi" {
		t.Errorf("got %q, want Wi-Fi", p.ClassName)
	}
	if !p.PrefersMSIX {
		t.Error("Wi-Fi should prefer MSI-X")
	}
}

func TestProfileForClass_Thunderbolt(t *testing.T) {
	p := ProfileForClass(0x0C8000)
	if p == nil {
		t.Fatal("Thunderbolt profile should not be nil")
	}
	if p.ClassName != "Thunderbolt" {
		t.Errorf("got %q, want Thunderbolt", p.ClassName)
	}
}

func TestProfileForClass_Generic(t *testing.T) {
	p := ProfileForClass(0xFF0000)
	if p == nil {
		t.Fatal("generic fallback should not be nil")
	}
	if p.ClassName != "Generic" {
		t.Errorf("got %q, want Generic", p.ClassName)
	}
}

func TestProfileForClass_ProgIFAgnostic(t *testing.T) {
	p := ProfileForClass(0x010801)
	if p == nil {
		t.Fatal("NVMe Fabrics should still match NVMe profile")
	}
}

func TestAllProfiles(t *testing.T) {
	profiles := AllProfiles()
	if len(profiles) != 14 {
		t.Errorf("expected 14 profiles, got %d", len(profiles))
	}
	names := make(map[string]bool)
	for _, p := range profiles {
		if p.ClassName == "" {
			t.Error("profile has empty ClassName")
		}
		names[p.ClassName] = true
	}
	for _, expected := range []string{
		"NVMe", "xHCI USB", "Ethernet", "HD Audio",
		"GPU", "SATA AHCI", "Wi-Fi", "Wi-Fi (MediaTek)",
		"Wi-Fi (Intel BE200 family)", "Wi-Fi (Atheros AR9287)",
		"Wi-Fi (Realtek RTL8188EE)", "Card Reader (RTS522A)", "Thunderbolt", "Generic",
	} {
		if !names[expected] {
			t.Errorf("missing profile: %s", expected)
		}
	}
}

func TestAllProfiles_BARDefaults(t *testing.T) {
	for _, p := range AllProfiles() {
		// Generic has no BAR defaults, skip
		if p.ClassName == "Generic" {
			continue
		}
		if len(p.BARDefaults) == 0 {
			t.Errorf("%s profile has no BAR defaults", p.ClassName)
		}
		for _, d := range p.BARDefaults {
			if d.Name == "" {
				t.Errorf("%s profile has BAR default at 0x%X with empty name", p.ClassName, d.Offset)
			}
			if d.Width != 4 {
				t.Errorf("%s: %s has Width=%d, all should be 4 (DWORD-aligned)", p.ClassName, d.Name, d.Width)
			}
		}
	}
}

func TestAllProfiles_ExpectedCaps(t *testing.T) {
	for _, p := range AllProfiles() {
		if p.ClassName == "Generic" {
			continue
		}
		if len(p.ExpectedCaps) == 0 {
			t.Errorf("%s profile has no expected capabilities", p.ClassName)
		}
	}
}

func TestNVMeProfile_CSTSReady(t *testing.T) {
	p := nvmeProfile()
	for _, d := range p.BARDefaults {
		if d.Name == "CSTS" {
			if d.Reset&0x01 == 0 {
				t.Error("CSTS.RDY should be 1 in profile")
			}
			return
		}
	}
	t.Error("CSTS not found in NVMe profile")
}

func TestXHCIProfile_PORTSC(t *testing.T) {
	p := xhciProfile()
	portCount := 0
	for _, d := range p.BARDefaults {
		if d.Name == "PORTSC1" || d.Name == "PORTSC2" {
			portCount++
			if d.Reset == 0 {
				t.Errorf("%s should have non-zero reset (PP set)", d.Name)
			}
		}
	}
	if portCount < 2 {
		t.Errorf("xHCI profile should have at least 2 PORTSC registers, found %d", portCount)
	}
}

func TestXHCIProfile_InterrupterRegs(t *testing.T) {
	p := xhciProfile()
	found := map[string]bool{}
	for _, d := range p.BARDefaults {
		switch d.Name {
		case "IMAN0", "ERSTSZ0", "ERSTBA_LO", "ERDP_LO":
			found[d.Name] = true
		}
	}
	for _, name := range []string{"IMAN0", "ERSTSZ0", "ERSTBA_LO", "ERDP_LO"} {
		if !found[name] {
			t.Errorf("xHCI profile missing interrupter register: %s", name)
		}
	}
}

func TestXHCIProfile_HCCPARAMS1_AC64(t *testing.T) {
	p := xhciProfile()
	for _, d := range p.BARDefaults {
		if d.Name == "HCCPARAMS1" {
			if d.Reset&0x01 == 0 {
				t.Error("HCCPARAMS1.AC64 should be set (64-bit capable)")
			}
			return
		}
	}
	t.Error("HCCPARAMS1 not found")
}

func TestIntelWiFiProfile_Basic(t *testing.T) {
	p := intelWiFiProfile()
	if p.ClassName != "Wi-Fi (Intel BE200 family)" {
		t.Fatalf("got %q, want Wi-Fi (Intel BE200 family)", p.ClassName)
	}
	if p.MinBARSize != 16384 {
		t.Errorf("MinBARSize = %d, want 16384", p.MinBARSize)
	}
	found := map[string]bool{}
	for _, d := range p.BARDefaults {
		switch d.Name {
		case "GP_CTL", "UCODE_DRV_GP1", "RF_ID":
			found[d.Name] = true
		}
	}
	for _, name := range []string{"GP_CTL", "UCODE_DRV_GP1", "RF_ID"} {
		if !found[name] {
			t.Errorf("Intel Wi-Fi profile missing BAR default %s", name)
		}
	}
}

func TestRTS522AProfile_Basic(t *testing.T) {
	p := rts522aProfile()
	if p.ClassName != "Card Reader (RTS522A)" {
		t.Fatalf("got %q, want Card Reader (RTS522A)", p.ClassName)
	}
	if p.PreferredBAR != 0 {
		t.Errorf("PreferredBAR = %d, want 0", p.PreferredBAR)
	}
	if p.Uses64BitBAR {
		t.Error("RTS522A should not use 64-bit BAR")
	}
	found := map[string]bool{}
	for _, d := range p.BARDefaults {
		switch d.Name {
		case "RTSX_HCBAR", "RTSX_BIPR", "RTSX_BIER":
			found[d.Name] = true
		}
	}
	for _, name := range []string{"RTSX_HCBAR", "RTSX_BIPR", "RTSX_BIER"} {
		if !found[name] {
			t.Errorf("RTS522A profile missing BAR default %s", name)
		}
	}
}

func TestRTS522AProfile_HAIMRFSM(t *testing.T) {
	p := rts522aProfile()
	for _, d := range p.BARDefaults {
		if d.Name == "RTSX_HAIMR" {
			if !d.IsFSMDriven {
				t.Fatal("RTSX_HAIMR should be FSM-driven for indirect register emulation")
			}
			return
		}
	}
	t.Fatal("RTSX_HAIMR not found")
}

func TestRTS522AProfile_BIPRFSM(t *testing.T) {
	p := rts522aProfile()
	for _, d := range p.BARDefaults {
		if d.Name == "RTSX_BIPR" {
			if !d.IsFSMDriven {
				t.Fatal("RTSX_BIPR should be FSM-driven for interrupt status semantics")
			}
			return
		}
	}
	t.Fatal("RTSX_BIPR not found")
}

func TestRTS522AProfile_BIPRCardPresent(t *testing.T) {
	p := rts522aProfile()
	for _, d := range p.BARDefaults {
		if d.Name == "RTSX_BIPR" {
			if d.Reset&0x00010000 == 0 {
				t.Fatal("RTSX_BIPR should report SD_EXIST on reset")
			}
			return
		}
	}
	t.Fatal("RTSX_BIPR not found")
}

func TestAudioProfile_DWORDPacked(t *testing.T) {
	p := audioProfile()
	for _, d := range p.BARDefaults {
		if d.Width != 4 {
			t.Errorf("Audio %s has Width=%d, must be 4", d.Name, d.Width)
		}
	}
}

func TestAudioProfile_STATESTS(t *testing.T) {
	p := audioProfile()
	for _, d := range p.BARDefaults {
		if d.Name == "WAKEEN_STATESTS" {
			if d.Reset&0x10000 == 0 {
				t.Error("STATESTS should have codec 0 present (bit 16)")
			}
			return
		}
	}
	t.Error("WAKEEN_STATESTS not found")
}

func TestEthernetProfile_MAC(t *testing.T) {
	p := ethernetProfile()
	var mac03, mac45 bool
	for _, d := range p.BARDefaults {
		if d.Name == "MAC0_3" {
			mac03 = true
			if d.Reset == 0 {
				t.Error("MAC0_3 should have default MAC")
			}
		}
		if d.Name == "MAC4_5" {
			mac45 = true
			if d.Reset == 0 {
				t.Error("MAC4_5 should have default MAC")
			}
		}
	}
	if !mac03 || !mac45 {
		t.Error("Ethernet profile should have MAC0_3 and MAC4_5")
	}
}

func TestGPUProfile_PMCEnable(t *testing.T) {
	p := gpuProfile()
	found := false
	for _, d := range p.BARDefaults {
		if d.Name == "PMC_ENABLE" {
			found = true
			if d.Reset != 0xFFFFFFFF {
				t.Errorf("PMC_ENABLE default should be 0xFFFFFFFF, got 0x%08X", d.Reset)
			}
		}
	}
	if !found {
		t.Error("PMC_ENABLE not found in GPU profile")
	}
}

func TestSATAProfile_GHC(t *testing.T) {
	p := sataProfile()
	for _, d := range p.BARDefaults {
		if d.Name == "GHC" {
			if d.Reset&0x80000000 == 0 {
				t.Error("GHC.AE should be set")
			}
			return
		}
	}
	t.Error("GHC not found in SATA profile")
}
