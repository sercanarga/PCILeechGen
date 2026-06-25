package donor

import "testing"

func TestClassifyRegisterBits_RWOnly(t *testing.T) {
	rw, w1c := classifyRegisterBits(0xFF, 0x00, 0xFF, 0xFF)
	if rw != 0xFF {
		t.Errorf("rwMask: got 0x%X, want 0xFF", rw)
	}
	if w1c != 0 {
		t.Errorf("w1cMask: got 0x%X, want 0", w1c)
	}
}

func TestClassifyRegisterBits_DetectsW1C(t *testing.T) {
	rw, w1c := classifyRegisterBits(0x100, 0x00, 0x100, 0x00)
	if rw != 0x100 {
		t.Errorf("rwMask: got 0x%X, want 0x100", rw)
	}
	if w1c != 0x100 {
		t.Errorf("w1cMask: got 0x%X, want 0x100", w1c)
	}
}

func TestClassifyRegisterBits_MixedRWandW1C(t *testing.T) {
	rw, w1c := classifyRegisterBits(0x103, 0x00, 0x103, 0x003)
	if rw != 0x103 {
		t.Errorf("rwMask: got 0x%X, want 0x103", rw)
	}
	if w1c != 0x100 {
		t.Errorf("w1cMask: got 0x%X, want 0x100", w1c)
	}
}

func TestProfileBARFromBuffer_NoW1COnPlainBuffer(t *testing.T) {
	prof := ProfileBARFromBuffer(make([]byte, 16), 0)
	for _, p := range prof.Probes {
		if p.W1CMask != 0 {
			t.Errorf("reg 0x%X: W1CMask 0x%X, want 0", p.Offset, p.W1CMask)
		}
		if p.MaybeRW1C {
			t.Errorf("reg 0x%X: MaybeRW1C true, want false", p.Offset)
		}
	}
}
