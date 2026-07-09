package vivado

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"
)

func TestFindValidation(t *testing.T) {
	_, err := Find("/nonexistent/path/to/vivado")
	if err == nil {
		t.Error("Find should fail for non-existent custom path")
	}
}

func TestFindNoArgs(t *testing.T) {
	_, err := Find("")
	if err == nil {
		return
	}
	if err.Error() == "" {
		t.Error("Error message should not be empty")
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

	if err := os.WriteFile(logPath, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

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

func TestRunTCL_ReturnsError_whenVivadoStartupFails(t *testing.T) {
	installDir := t.TempDir()
	binDir := filepath.Join(installDir, "bin")
	if err := os.MkdirAll(binDir, 0755); err != nil {
		t.Fatal(err)
	}
	vivadoPath := filepath.Join(binDir, "vivado")
	script := "#!/bin/sh\necho 'application-specific initialization failed: couldn'\\''t load file \"librdi_commontasks.so\": libtinfo.so.5: cannot open shared object file: No such file or directory' >&2\nexit 0\n"
	if runtime.GOOS == "windows" {
		vivadoPath += ".bat"
		script = "@echo off\r\necho application-specific initialization failed: couldn't load file \"librdi_commontasks.so\": libtinfo.so.5: cannot open shared object file: No such file or directory 1>&2\r\nexit /b 0\r\n"
	}
	if err := os.WriteFile(vivadoPath, []byte(script), 0755); err != nil {
		t.Fatal(err)
	}

	v := &Vivado{Path: installDir, Version: filepath.Base(installDir)}
	err := v.RunTCL("vivado_generate_project.tcl", t.TempDir(), 5*time.Second)

	if err == nil {
		t.Fatal("RunTCL should fail when Vivado reports a startup failure")
	}
	if !strings.Contains(err.Error(), "Vivado startup failed") {
		t.Fatalf("RunTCL error = %q, want startup failure", err.Error())
	}
	if !strings.Contains(err.Error(), "libtinfo.so.5") {
		t.Fatalf("RunTCL error = %q, want loader detail", err.Error())
	}
}
