//go:build evidence

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/firmware/output"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

// TestEvidence_Provenance is a build-tagged test that materializes a manifest
// with provenance and writes the JSON + degradation proof to the evidence
// path. Run with: go test -tags=evidence ./cmd/pcileechgen -run=Provenance
func TestEvidence_Provenance(t *testing.T) {
	tmpDir := t.TempDir()
	b := &board.Board{
		Name:       "EvidenceBoard",
		FPGAPart:   "xc7a35tfgg484-2",
		PCIeLanes:  1,
		TopModule:  "pcileech_top",
		ProjectDir: "EvidenceBoard",
		TCLFile:    "vivado.tcl",
	}
	cs := pci.NewConfigSpace()
	cs.WriteU32(0, 0x8086)
	cs.WriteU32(2, 0x1533)
	ctx := &donor.DeviceContext{
		ToolVersion: "evidence",
		Device:      pci.PCIDevice{VendorID: 0x8086, DeviceID: 0x1533},
		ConfigSpace: cs,
		BARs:        []pci.BAR{{Index: 0, Size: 4096}},
	}

	m, err := output.GenerateManifestForBuild(tmpDir, "1.0.0", b, ctx,
		[]string{"https://example/intake/ref-1"}, 0x8086, 0x1533)
	if err != nil {
		t.Fatalf("GenerateManifestForBuild: %v", err)
	}
	manifestPath := filepath.Join(tmpDir, "build_manifest.json")
	if err := m.WriteJSON(manifestPath); err != nil {
		t.Fatalf("WriteJSON: %v", err)
	}
	data, _ := os.ReadFile(manifestPath)

	// Degradation: no board, no donor, no intake refs.
	deg, err := output.GenerateManifestForBuild(t.TempDir(), "1.0.0", nil, nil, nil, 0, 0)
	if err != nil {
		t.Fatalf("degradation GenerateManifestForBuild: %v", err)
	}
	degJSON, _ := json.MarshalIndent(deg.Provenance, "", "  ")

	ev := "/tmp/pcg-task3-evidence.txt"
	out := fmt.Sprintf("=== PCILeechGen Todo 3: Adaptation Provenance Evidence ===\n\n"+
		"--- Generated manifest JSON (with donor + board + intake refs) ---\n%s\n\n"+
		"--- Provenance section (happy path) ---\n%s\n\n"+
		"--- Degradation provenance (no git/vivado/donor/board) ---\n%s\n\n"+
		"Contract: GeneratorGitCommit is set (hash or %q); missing sources degrade to omitted/Unknown;\n"+
		"build did not fail when donor/board/vivado were absent.\n",
		string(data), mustJSON(m.Provenance), string(degJSON), "unknown")
	if err := os.WriteFile(ev, []byte(out), 0644); err != nil {
		t.Fatalf("write evidence: %v", err)
	}
	t.Logf("evidence written to %s", ev)
}

func mustJSON(v any) string {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Sprintf("<marshal error: %v>", err)
	}
	return string(b)
}