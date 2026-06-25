package svgen

import "testing"

func TestNVMeDoorbellOffsets_QID1(t *testing.T) {
	cfg := &SVGeneratorConfig{NVMeDoorbellStride: 0}

	if got := cfg.NVMeSQ1DoorbellOffset(); got != 0x1008 {
		t.Errorf("SQ1 doorbell = 0x%X, want 0x1008", got)
	}
	if got := cfg.NVMeCQ1DoorbellOffset(); got != 0x100C {
		t.Errorf("CQ1 doorbell = 0x%X, want 0x100C", got)
	}

	cfg.NVMeDoorbellStride = 1
	if got := cfg.NVMeSQ1DoorbellOffset(); got != 0x1010 {
		t.Errorf("SQ1 doorbell (stride=1) = 0x%X, want 0x1010", got)
	}
	if got := cfg.NVMeCQ1DoorbellOffset(); got != 0x1018 {
		t.Errorf("CQ1 doorbell (stride=1) = 0x%X, want 0x1018", got)
	}
}
