package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/board"
)

func TestInspectBoardArtifacts(t *testing.T) {
	tmp := t.TempDir()
	b := &board.Board{
		Name:       "TestBoard",
		FPGAPart:   "xc7a200tfbg484-2",
		PCIeLanes:  1,
		TopModule:  "pcileech_test_top",
		ProjectDir: "test-board",
		TCLFile:    "vivado_generate_project.tcl",
		BuildTCL:   "vivado_build.tcl",
		SubDir:     "",
	}

	libDir := filepath.Join(tmp, "pcileech-fpga")
	src := filepath.Join(libDir, b.ProjectDir, "src")
	tcl := filepath.Join(libDir, b.ProjectDir, b.TCLFile)
	build := filepath.Join(libDir, b.ProjectDir, b.BuildTCL)
	if err := os.MkdirAll(src, 0o755); err != nil {
		t.Fatalf("mkdir src: %v", err)
	}
	if err := os.WriteFile(tcl, []byte("module none;"), 0o644); err != nil {
		t.Fatalf("write tcl: %v", err)
	}
	if err := os.WriteFile(build, []byte("puts hi"), 0o644); err != nil {
		t.Fatalf("write build tcl: %v", err)
	}

	items := inspectBoardArtifacts(b, libDir)
	if len(items) == 0 {
		t.Fatalf("expected check items, got 0")
	}

	status := map[string]bool{}
	for _, item := range items {
		status[item.Label] = item.Exists
	}
	if !status["src"] || !status["tcl"] || !status["build_tcl"] {
		t.Fatalf("expected required artifacts to exist")
	}
	if status["ip"] {
		t.Fatalf("ip should be missing in this fixture")
	}
}

func TestWriteBoardInspectResultJSON(t *testing.T) {
	res := &boardInspectResult{
		Name:   "Board",
		FPGA:   "xc7a200",
		Checks: []boardCheckItem{{Label: "src", Path: "/tmp/src", Exists: true, Required: true}},
	}
	var out bytes.Buffer
	if err := writeBoardInspectResult(&out, res, true); err != nil {
		t.Fatalf("encode: %v", err)
	}
	var decoded boardInspectResult
	if err := json.Unmarshal(out.Bytes(), &decoded); err != nil {
		t.Fatalf("json decode: %v", err)
	}
	if decoded.Name != "Board" {
		t.Fatalf("name = %q", decoded.Name)
	}
}

func TestBoardScaffoldEmitsJSON(t *testing.T) {
	var out bytes.Buffer
	boardScaffoldOpts.name = "AcmeBoard"
	cmd := *boardScaffoldCmd
	cmd.SetOut(&out)
	err := cmd.RunE(&cmd, []string{})
	if err != nil {
		t.Fatalf("run scaffold: %v", err)
	}
	if !bytes.Contains(out.Bytes(), []byte("AcmeBoard")) {
		t.Fatalf("scaffold output did not include board name: %s", out.String())
	}
}
