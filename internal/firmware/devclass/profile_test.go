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
}

func TestProfileForClass_XHCI(t *testing.T) {
	p := ProfileForClass(0x0C0330)
	if p == nil {
		t.Fatal("xHCI profile should not be nil")
	}
	if p.ClassName != "xHCI USB" {
		t.Errorf("got %q, want xHCI USB", p.ClassName)
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
}

func TestProfileForClass_Audio(t *testing.T) {
	p := ProfileForClass(0x040300)
	if p == nil {
		t.Fatal("HD Audio profile should not be nil")
	}
	if p.PrefersMSIX {
		t.Error("HD Audio should prefer MSI (not MSI-X)")
	}
}

func TestProfileForClass_Unknown(t *testing.T) {
	p := ProfileForClass(0xFF0000)
	if p != nil {
		t.Error("unknown class should return nil")
	}
}

func TestProfileForClass_ProgIFAgnostic(t *testing.T) {
	// NVMe with non-standard progIF should still match
	p := ProfileForClass(0x010801) // NVMe Fabrics (progIF=01)
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
