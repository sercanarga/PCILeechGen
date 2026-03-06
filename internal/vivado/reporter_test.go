package vivado

import (
	"strings"
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

func TestParseOutput_SynthComplete(t *testing.T) {
	output := `INFO: [Synth 8-7080] Merged and optimized
WARNING: [Synth 8-3331] design has unconnected port
synth_design completed successfully`

	r := ParseOutput(output)
	if !r.SynthComplete {
		t.Error("SynthComplete should be true")
	}
	if r.ImplComplete {
		t.Error("ImplComplete should be false")
	}
	if r.Warnings != 1 {
		t.Errorf("Warnings = %d, want 1", r.Warnings)
	}
}

func TestParseOutput_FullBuild(t *testing.T) {
	output := `synth_design completed successfully
place_design completed successfully
route_design completed successfully
write_bitstream completed successfully`

	r := ParseOutput(output)
	if !r.SynthComplete {
		t.Error("SynthComplete should be true")
	}
	if !r.ImplComplete {
		t.Error("ImplComplete should be true")
	}
	if !r.BitstreamReady {
		t.Error("BitstreamReady should be true")
	}
}

func TestParseOutput_MultipleErrors(t *testing.T) {
	output := `ERROR: [DRC 23-20] Rule violation critical
ERROR: [Place 30-574] cannot place component
WARNING: [Vivado 12-584] input port not connected`

	r := ParseOutput(output)
	if r.Errors != 2 {
		t.Errorf("Errors = %d, want 2", r.Errors)
	}
	if r.Warnings != 1 {
		t.Errorf("Warnings = %d, want 1", r.Warnings)
	}
}

func TestParseOutput_CriticalWarnings(t *testing.T) {
	output := `CRITICAL WARNING: [Timing 38-282] The design failed to meet timing`
	r := ParseOutput(output)
	if r.CriticalWarns != 1 {
		t.Errorf("CriticalWarns = %d, want 1", r.CriticalWarns)
	}
}

func TestIsBenign_KnownCodes(t *testing.T) {
	benign := LogEntry{Severity: SeverityWarning, Code: "Synth 8-7080"}
	if !benign.IsBenign() {
		t.Error("Synth 8-7080 should be benign")
	}

	notBenign := LogEntry{Severity: SeverityWarning, Code: "Unknown-999"}
	if notBenign.IsBenign() {
		t.Error("Unknown-999 should not be benign")
	}
}

func TestActionableEntries_FiltersBenign(t *testing.T) {
	output := `INFO: [Common 17-206] data loaded
WARNING: [Synth 8-7080] benign warning
WARNING: [Place 30-123] real warning
ERROR: [DRC 23-20] violation`

	r := ParseOutput(output)
	actionable := r.ActionableEntries()

	// Should filter out INFO and benign
	if len(actionable) != 2 {
		t.Errorf("Actionable = %d, want 2", len(actionable))
	}
}

func TestSummary_SuccessBuild(t *testing.T) {
	output := `synth_design completed successfully
route_design completed successfully
write_bitstream completed successfully
WARNING: [Synth 8-7080] benign`

	r := ParseOutput(output)
	summary := r.Summary()
	if !strings.Contains(summary, "SUCCESS") {
		t.Error("Summary should say SUCCESS for bitstream ready")
	}
	if !strings.Contains(summary, "benign") {
		t.Error("Summary should mention filtered benign warnings")
	}
}

func TestSummary_FailedBuild(t *testing.T) {
	output := `ERROR: [Place 30-574] cannot place`
	r := ParseOutput(output)
	summary := r.Summary()
	if !strings.Contains(summary, "FAILED") {
		t.Error("Summary should say FAILED when errors exist")
	}
}

func TestSummary_Unknown(t *testing.T) {
	output := `some random output`
	r := ParseOutput(output)
	summary := r.Summary()
	if !strings.Contains(summary, "unknown") {
		t.Error("Summary should say unknown for no progress")
	}
}

func TestSummary_SynthOnly(t *testing.T) {
	output := `synth_design completed successfully`
	r := ParseOutput(output)
	summary := r.Summary()
	if !strings.Contains(summary, "synthesis complete") {
		t.Error("Summary should mention synthesis complete")
	}
}

func TestSummary_ImplOnly(t *testing.T) {
	output := `synth_design completed successfully
place_design completed successfully`
	r := ParseOutput(output)
	summary := r.Summary()
	if !strings.Contains(summary, "implementation complete") {
		t.Error("Summary should mention implementation complete")
	}
}

func TestParseLogFile_NonExistent(t *testing.T) {
	_, err := ParseLogFile("/nonexistent/log.txt")
	if err == nil {
		t.Error("ParseLogFile should fail for non-existent file")
	}
}
