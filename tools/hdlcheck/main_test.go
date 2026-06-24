package main

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestDiscoverFiles(t *testing.T) {
	dir := t.TempDir()
	writeTestFile(t, filepath.Join(dir, "a.sv"))
	writeTestFile(t, filepath.Join(dir, "b.v"))
	writeTestFile(t, filepath.Join(dir, "c.tcl"))
	writeTestFile(t, filepath.Join(dir, "ignore.txt"))

	files, err := discoverFiles(dir)
	if err != nil {
		t.Fatalf("discoverFiles failed: %v", err)
	}
	if len(files.HDL) != 2 || len(files.TCL) != 1 {
		t.Fatalf("discoverFiles = %#v, want 2 HDL and 1 TCL", files)
	}
}

func TestBuildPlanSkipsMissingTools(t *testing.T) {
	files := fileSet{HDL: []string{"a.sv"}, TCL: []string{"build.tcl"}}
	plan := buildPlan(files, func(string) (string, error) { return "", errors.New("missing") })

	if len(plan.Commands) != 0 {
		t.Fatalf("commands = %#v, want none", plan.Commands)
	}
	if len(plan.Skipped) != 2 {
		t.Fatalf("skipped = %#v, want verilator and vivado skips", plan.Skipped)
	}
}

func TestBuildPlanUsesAvailableVerilator(t *testing.T) {
	files := fileSet{HDL: []string{"a.sv", "b.v"}}
	plan := buildPlan(files, func(name string) (string, error) {
		if name == "verilator" {
			return "/usr/bin/verilator", nil
		}
		return "", errors.New("missing")
	})

	if len(plan.Commands) != 1 {
		t.Fatalf("commands = %#v, want one", plan.Commands)
	}
	cmd := plan.Commands[0]
	if cmd.Args[0] != "/usr/bin/verilator" || cmd.Args[1] != "--lint-only" {
		t.Fatalf("unexpected command: %#v", cmd.Args)
	}
}

func writeTestFile(t *testing.T, path string) {
	t.Helper()
	if err := os.WriteFile(path, []byte("module dummy; endmodule\n"), 0644); err != nil {
		t.Fatalf("write %s: %v", path, err)
	}
}
