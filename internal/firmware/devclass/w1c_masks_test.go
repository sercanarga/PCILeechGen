package devclass

import "testing"

func TestAudioProfile_W1CMasks(t *testing.T) {
	p := audioProfile()
	got := map[uint32]uint32{}
	for _, d := range p.BARDefaults {
		got[d.Offset] = d.W1CMask
	}
	if got[0x4C] != 0x00000100 {
		t.Errorf("0x4C W1CMask: got 0x%08X, want 0x00000100", got[0x4C])
	}
	if got[0x5C] != 0x00000700 {
		t.Errorf("0x5C W1CMask: got 0x%08X, want 0x00000700", got[0x5C])
	}
	if got[0x60] != 0x00000001 {
		t.Errorf("0x60 W1CMask: got 0x%08X, want 0x00000001", got[0x60])
	}
}

func TestAudioProfile_MasksDisjoint(t *testing.T) {
	p := audioProfile()
	for _, d := range p.BARDefaults {
		if d.W1CMask&d.RWMask != 0 {
			t.Errorf("0x%X %s: W1CMask 0x%08X overlaps RWMask 0x%08X", d.Offset, d.Name, d.W1CMask, d.RWMask)
		}
	}
}

func TestXHCIProfile_USBSTS_W1CMask(t *testing.T) {
	p := xhciProfile()
	for _, d := range p.BARDefaults {
		if d.Name == "USBSTS" {
			if d.W1CMask != 0x0000041C {
				t.Errorf("USBSTS W1CMask: got 0x%08X, want 0x0000041C", d.W1CMask)
			}
			if d.RWMask != 0 {
				t.Errorf("USBSTS RWMask: got 0x%08X, want 0", d.RWMask)
			}
			return
		}
	}
	t.Error("USBSTS not found")
}
