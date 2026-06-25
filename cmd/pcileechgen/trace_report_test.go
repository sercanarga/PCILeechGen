package main

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

// TestTraceReportCmd_HappyAndQuarantine drives the CLI end-to-end on a
// synthetic fixture containing every category plus an out-of-BAR record, and
// asserts the report is produced, deterministic, and carries the quarantine
// warning.
func TestTraceReportCmd_HappyAndQuarantine(t *testing.T) {
	resetTraceReportFlags()
	dir := t.TempDir()
	tracePath := filepath.Join(dir, "trace.log")
	const body = `# synthetic multi-category fixture
R 4 0.001 0x80000000 0x0040ff17
R 4 0.002 0x8000001c 0x00000000
R 4 0.003 0x8000001c 0x00000000
R 4 0.004 0x8000001c 0x00000000
R 4 0.005 0x8000001c 0x00000000
R 4 0.006 0x8000001c 0x00000000
R 4 0.007 0x80000030 0x00000001
R 4 0.008 0x80000030 0x00000002
R 4 0.009 0x80000030 0x00000003
R 4 0.010 0x80000030 0x00000004
W 4 0.011 0x80000014 0x00460001
R 4 0.012 0x80000014 0x00460001
R 4 0.013 0x80000040 0x000000aa
R 4 0.014 0x80000040 0x00000011
R 4 0.015 0x80000040 0x000000ee
R 4 0.016 0x90000000 0x0000dead
`
	if err := os.WriteFile(tracePath, []byte(body), 0644); err != nil {
		t.Fatal(err)
	}

	out1 := filepath.Join(dir, "report1.json")
	out2 := filepath.Join(dir, "report2.json")

	run := func(out string) {
		rootCmd.SetArgs([]string{"trace-report", tracePath,
			"--bar-base", "0x80000000", "--bar-size", "0x1000",
			"-o", out, "--overwrite"})
		var buf bytes.Buffer
		rootCmd.SetOut(&buf)
		rootCmd.SetErr(&buf)
		if err := rootCmd.Execute(); err != nil {
			t.Fatalf("trace-report failed: %v\n%s", err, buf.String())
		}
	}
	run(out1)
	run(out2)

	b1, err := os.ReadFile(out1)
	if err != nil {
		t.Fatal(err)
	}
	b2, err := os.ReadFile(out2)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(b1, b2) {
		t.Fatalf("non-deterministic CLI output:\n--- run 1 ---\n%s\n--- run 2 ---\n%s", b1, b2)
	}

	// Spot-check the schema fields are present.
	if !bytes.Contains(b1, []byte(`"tag": "static"`)) ||
		!bytes.Contains(b1, []byte(`"tag": "polling"`)) ||
		!bytes.Contains(b1, []byte(`"tag": "counter"`)) ||
		!bytes.Contains(b1, []byte(`"tag": "rw"`)) ||
		!bytes.Contains(b1, []byte(`"tag": "volatile"`)) ||
		!bytes.Contains(b1, []byte(`"unknown_regions"`)) {
		t.Fatalf("report missing expected tags or unknown_regions:\n%s", b1)
	}
}

// resetTraceReportFlags clears package-level flag vars so tests do not leak
// state into each other (cobra flags persist across Execute calls).
func resetTraceReportFlags() {
	traceReportContext = ""
	traceReportBarBase = ""
	traceReportBarSize = ""
	traceReportBarIndex = 0
	traceReportOutput = ""
	traceReportOverwrite = false
}

// TestTraceReportCmd_RequiresBounds verifies the command refuses to run
// without any BAR bounds source.
func TestTraceReportCmd_RequiresBounds(t *testing.T) {
	resetTraceReportFlags()
	dir := t.TempDir()
	tracePath := filepath.Join(dir, "trace.log")
	if err := os.WriteFile(tracePath, []byte("R 4 0.001 0x80000000 0x1\n"), 0644); err != nil {
		t.Fatal(err)
	}
	out := filepath.Join(dir, "out.json")
	rootCmd.SetArgs([]string{"trace-report", tracePath, "-o", out, "--overwrite"})
	var buf bytes.Buffer
	rootCmd.SetOut(&buf)
	rootCmd.SetErr(&buf)
	err := rootCmd.Execute()
	if err == nil {
		t.Fatal("expected error when no BAR bounds provided")
	}
}