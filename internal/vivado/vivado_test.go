package vivado

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/donor"
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
	script := `#!/bin/sh
echo 'application-specific initialization failed: couldn'\''t load file "librdi_commontasks.so": libtinfo.so.5: cannot open shared object file: No such file or directory' >&2
exit 0
`
	if err := os.WriteFile(vivadoPath, []byte(script), 0755); err != nil {
		t.Fatal(err)
	}

	v := &Vivado{Path: installDir, Version: filepath.Base(installDir)}
	workDir := t.TempDir()
	err := v.RunTCL("vivado_generate_project.tcl", workDir, 5*time.Second)

	if err == nil {
		t.Fatal("RunTCL should fail when Vivado reports a startup failure")
	}
	if !strings.Contains(err.Error(), "Vivado startup failed") {
		t.Fatalf("RunTCL error = %q, want startup failure", err.Error())
	}
	if !strings.Contains(err.Error(), "libtinfo.so.5") {
		t.Fatalf("RunTCL error = %q, want loader detail", err.Error())
	}
	summary, readErr := os.ReadFile(filepath.Join(workDir, "vivado_generate_project_summary.txt"))
	if readErr != nil {
		t.Fatalf("read Vivado summary: %v", readErr)
	}
	if !strings.Contains(string(summary), "Build Status: unknown") {
		t.Fatalf("summary = %q, want structured report", summary)
	}
}

func TestSummarizeRunIncludesBoundedReportAndDiagnosis(t *testing.T) {
	output := strings.Join([]string{
		"WARNING: [Synth 8-7080] benign merged",
		"CRITICAL WARNING: [Place 30-123] congested",
		"ERROR: [Route 35-39] routing failed",
		"ERROR: [Common 17-69] out of memory while placing design",
	}, "\n")
	summary := summarizeRun("vivado_build.tcl", output, errors.New("exit status 1"))
	for _, want := range []string{
		"Vivado summary for vivado_build.tcl",
		"Build Status: FAILED",
		"Errors: 2, Critical Warnings: 1, Warnings: 1",
		"process_error=exit status 1",
		"diagnosis=Vivado was killed by the OS",
	} {
		if !strings.Contains(summary, want) {
			t.Fatalf("summary missing %q:\n%s", want, summary)
		}
	}
}

func TestWriteRunSummaryRejectsSymlink(t *testing.T) {
	root := t.TempDir()
	outside := filepath.Join(root, "outside.txt")
	if err := os.WriteFile(outside, []byte("preserve"), 0o644); err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(root, "vivado_build_summary.txt")
	if err := os.Symlink(outside, path); err != nil {
		t.Skipf("symlink support unavailable: %v", err)
	}
	if err := writeRunSummary(root, "vivado_build.tcl", "replacement"); err == nil {
		t.Fatal("writeRunSummary accepted a symlink")
	}
	contents, err := os.ReadFile(outside)
	if err != nil {
		t.Fatal(err)
	}
	if string(contents) != "preserve" {
		t.Fatalf("symlink target changed to %q", contents)
	}
}

func TestRefreshBuildManifest_ReturnsFailure(t *testing.T) {
	err := refreshBuildManifest(
		func(string, *donor.DeviceContext, *board.Board) error {
			return errors.New("manifest storage unavailable")
		},
		t.TempDir(),
		&donor.DeviceContext{},
		&board.Board{Name: "TestBoard"},
	)
	if err == nil || !strings.Contains(err.Error(), "manifest storage unavailable") {
		t.Fatalf("refreshBuildManifest error = %v, want storage failure", err)
	}
}
