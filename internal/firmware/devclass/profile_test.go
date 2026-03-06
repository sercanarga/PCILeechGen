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
	if p.Uses64BitBAR {
		t.Error("Ethernet profile should use 32-bit BAR")
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
	if p.Uses64BitBAR != true {
		t.Error("HD Audio should use 64-bit BAR")
	}
}

func TestProfileForClass_Unknown(t *testing.T) {
	p := ProfileForClass(0xFF0000)
	if p != nil {
		t.Error("unknown class should return nil")
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
	if len(profiles) != 4 {
		t.Errorf("expected 4 profiles, got %d", len(profiles))
	}
	names := make(map[string]bool)
	for _, p := range profiles {
		if p.ClassName == "" {
			t.Error("profile has empty ClassName")
		}
		names[p.ClassName] = true
	}
	for _, expected := range []string{"NVMe", "xHCI USB", "Ethernet", "HD Audio"} {
		if !names[expected] {
			t.Errorf("missing profile: %s", expected)
		}
	}
}

func TestAllProfiles_BARDefaults(t *testing.T) {
	for _, p := range AllProfiles() {
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
	var ral0, rah0 bool
	for _, d := range p.BARDefaults {
		if d.Name == "RAL0" {
			ral0 = true
			if d.Reset == 0 {
				t.Error("RAL0 should have default MAC")
			}
		}
		if d.Name == "RAH0" {
			rah0 = true
			if d.Reset&0x80000000 == 0 {
				t.Error("RAH0 should have AV bit")
			}
		}
	}
	if !ral0 || !rah0 {
		t.Error("Ethernet profile should have RAL0 and RAH0")
	}
}
