package vivado

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestFindValidation(t *testing.T) {
	// Non-existent path should fail
	_, err := Find("/nonexistent/path/to/vivado")
	if err == nil {
		t.Error("Find should fail for non-existent custom path")
	}
}

func TestFindNoArgs(t *testing.T) {
	// Without vivado installed, Find("") should fail gracefully
	_, err := Find("")
	if err == nil {
		// Vivado is actually installed - still valid
		return
	}
	if err.Error() == "" {
		t.Error("Error message should not be empty")
	}
}

func TestVivadoBinaryPath(t *testing.T) {
	v := &Vivado{
		Path:    "/tools/Xilinx/Vivado/2022.2",
		Version: "2022.2",
	}

	path := v.BinaryPath()
	expected := "/tools/Xilinx/Vivado/2022.2/bin/vivado"
	if path != expected {
		t.Errorf("BinaryPath() = %q, want %q", path, expected)
	}
}

func TestBuilderDefaults(t *testing.T) {
	b, _ := Find("") // will fail but ok for testing builder creation
	_ = b

	opts := BuildOptions{}
	builder := NewBuilder(nil, opts)

	if builder.opts.Jobs != 4 {
		t.Errorf("Default jobs = %d, want 4", builder.opts.Jobs)
	}
	if builder.opts.Timeout != 3600 {
		t.Errorf("Default timeout = %d, want 3600", builder.opts.Timeout)
	}
	if builder.opts.OutputDir != "pcileech_datastore" {
		t.Errorf("Default output = %q, want 'pcileech_datastore'", builder.opts.OutputDir)
	}
}

func TestParseLogFile_Valid(t *testing.T) {
	tmpDir := t.TempDir()
	logPath := filepath.Join(tmpDir, "vivado.log")

	content := `INFO: [Common 17-206] data loaded
WARNING: [Synth 8-7080] benign merged
WARNING: [Place 30-123] real warning
ERROR: [DRC 23-20] rule violation
synth_design completed successfully
route_design completed successfully
write_bitstream completed successfully`

	os.WriteFile(logPath, []byte(content), 0644)

	report, err := ParseLogFile(logPath)
	if err != nil {
		t.Fatalf("ParseLogFile error: %v", err)
	}
	if !report.SynthComplete {
		t.Error("SynthComplete should be true")
	}
	if !report.BitstreamReady {
		t.Error("BitstreamReady should be true")
	}
	if report.Errors != 1 {
		t.Errorf("Errors = %d, want 1", report.Errors)
	}
	if report.Warnings != 2 {
		t.Errorf("Warnings = %d, want 2", report.Warnings)
	}
}

func TestSummary_ManyActionable(t *testing.T) {
	var lines []string
	for i := 0; i < 25; i++ {
		lines = append(lines, "WARNING: [Custom 1-1] warning message")
	}
	r := ParseOutput(strings.Join(lines, "\n"))
	summary := r.Summary()
	if !strings.Contains(summary, "and") {
		t.Error("Summary should truncate excessive actionable warnings")
	}
}

func TestRunTCL_NonExistent(t *testing.T) {
	v := &Vivado{Path: "/nonexistent", Version: "2023.2"}
	err := v.RunTCL("/tmp/test.tcl", "/tmp", 5*time.Second)
	if err == nil {
		t.Error("RunTCL should fail with non-existent vivado")
	}
}
