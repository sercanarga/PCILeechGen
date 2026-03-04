package vivado

import (
	"testing"
)

func TestParseOutput_Errors(t *testing.T) {
	output := `INFO: [Synth 8-7065] analyzing module top
WARNING: [Synth 8-7080] Parallel synthesis criteria is not met
ERROR: [DRC AVAL-46] No PCIe hard block found
CRITICAL WARNING: [Timing 38-282] Setup timing not met
INFO: [Synth 8-7066] synthesis complete
`
	r := ParseOutput(output)
	if r.Errors != 1 {
		t.Errorf("expected 1 error, got %d", r.Errors)
	}
	if r.CriticalWarns != 1 {
		t.Errorf("expected 1 critical warning, got %d", r.CriticalWarns)
	}
	if r.Warnings != 1 {
		t.Errorf("expected 1 warning, got %d", r.Warnings)
	}
	if len(r.Entries) != 5 {
		t.Errorf("expected 5 entries, got %d", len(r.Entries))
	}
}

func TestParseOutput_Completion(t *testing.T) {
	output := `synth_design completed successfully
route_design completed successfully
write_bitstream completed successfully`

	r := ParseOutput(output)
	if !r.SynthComplete {
		t.Error("expected synth complete")
	}
	if !r.ImplComplete {
		t.Error("expected impl complete")
	}
	if !r.BitstreamReady {
		t.Error("expected bitstream ready")
	}
}

func TestIsBenign(t *testing.T) {
	benign := LogEntry{Code: "Synth 8-7080"}
	if !benign.IsBenign() {
		t.Error("Synth 8-7080 should be benign")
	}

	notBenign := LogEntry{Code: "DRC CUSTOM-99"}
	if notBenign.IsBenign() {
		t.Error("DRC CUSTOM-99 should not be benign")
	}
}

func TestActionableEntries(t *testing.T) {
	output := `WARNING: [Synth 8-7080] known benign warning
WARNING: [Custom 99-1] real warning
ERROR: [Build 1-1] something failed
INFO: [Synth 8-1] just info`

	r := ParseOutput(output)
	actionable := r.ActionableEntries()

	// should have 2: the non-benign warning + the error
	if len(actionable) != 2 {
		t.Errorf("expected 2 actionable, got %d", len(actionable))
		for _, e := range actionable {
			t.Logf("  %s: %s: %s", e.Severity, e.Code, e.Message)
		}
	}
}

func TestSummary_Success(t *testing.T) {
	r := &Report{
		BitstreamReady: true,
		SynthComplete:  true,
		ImplComplete:   true,
		Errors:         0,
		Warnings:       3,
	}
	s := r.Summary()
	if !containsStr(s, "SUCCESS") {
		t.Errorf("expected SUCCESS in summary, got:\n%s", s)
	}
}

func TestSummary_Failed(t *testing.T) {
	r := &Report{Errors: 1}
	s := r.Summary()
	if !containsStr(s, "FAILED") {
		t.Errorf("expected FAILED in summary, got:\n%s", s)
	}
}

func containsStr(s, sub string) bool {
	for i := 0; i <= len(s)-len(sub); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}
