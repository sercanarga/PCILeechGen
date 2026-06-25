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

func TestNVMeDiagnosticOffsets_FollowDoorbellStride(t *testing.T) {
	cfg := &SVGeneratorConfig{NVMeDoorbellStride: 0}

	if got := cfg.NVMeDiagBaseOffset(); got != 0x1010 {
		t.Errorf("diagnostic base = 0x%X, want 0x1010", got)
	}
	if got := cfg.NVMeDiagLastCommandOffset(); got != 0x1014 {
		t.Errorf("last command diagnostic = 0x%X, want 0x1014", got)
	}

	cfg.NVMeDoorbellStride = 1
	if got := cfg.NVMeDiagBaseOffset(); got != 0x1020 {
		t.Errorf("diagnostic base (stride=1) = 0x%X, want 0x1020", got)
	}
	if got := cfg.NVMeDiagLastCommandOffset(); got != 0x1024 {
		t.Errorf("last command diagnostic (stride=1) = 0x%X, want 0x1024", got)
	}
}
