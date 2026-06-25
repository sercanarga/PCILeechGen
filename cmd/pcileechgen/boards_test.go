package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/board"
)

func TestWriteBoards_printsJSON_whenRequested(t *testing.T) {
	// Given
	boards := []board.Board{{
		Name:       "TestBoard",
		FPGAPart:   "xc7a35tfgg484-2",
		PCIeLanes:  1,
		TopModule:  "pcileech_test_top",
		ProjectDir: "TestBoard",
		TCLFile:    "vivado_generate_project.tcl",
		Flash: &board.Flash{
			Tool:                "openFPGALoader",
			OpenFPGALoaderBoard: "testboard",
			Target:              "flash",
		},
	}}
	var out bytes.Buffer

	// When
	err := writeBoards(&out, boards, true)

	// Then
	if err != nil {
		t.Fatalf("writeBoards returned error: %v", err)
	}
	var got []boardOutput
	if err := json.Unmarshal(out.Bytes(), &got); err != nil {
		t.Fatalf("JSON output did not parse: %v", err)
	}
	if len(got) != 1 {
		t.Fatalf("JSON board count = %d, want 1", len(got))
	}
	if got[0].Name != "TestBoard" {
		t.Errorf("Name = %q, want TestBoard", got[0].Name)
	}
	if got[0].BRAMSize != board.DefaultBRAMSize {
		t.Errorf("BRAMSize = %d, want default %d", got[0].BRAMSize, board.DefaultBRAMSize)
	}
	if got[0].MaxLinkSpeed != 2 {
		t.Errorf("MaxLinkSpeed = %d, want default Gen2", got[0].MaxLinkSpeed)
	}
	if got[0].Flash == nil || got[0].Flash.OpenFPGALoaderBoard != "testboard" {
		t.Errorf("Flash = %#v, want testboard metadata", got[0].Flash)
	}
}

func TestWriteBoards_printsTable_whenJSONDisabled(t *testing.T) {
	// Given
	boards := []board.Board{{
		Name:       "TestBoard",
		FPGAPart:   "xc7a35tfgg484-2",
		PCIeLanes:  4,
		BRAMSize:   32768,
		TopModule:  "pcileech_test_top",
		ProjectDir: "TestBoard",
		TCLFile:    "vivado_generate_project.tcl",
	}}
	var out bytes.Buffer

	// When
	err := writeBoards(&out, boards, false)

	// Then
	if err != nil {
		t.Fatalf("writeBoards returned error: %v", err)
	}
	text := out.String()
	if !strings.Contains(text, "NAME") || !strings.Contains(text, "TestBoard") {
		t.Fatalf("table output missing expected board: %q", text)
	}
	if !strings.Contains(text, "Total: 1 boards") {
		t.Errorf("table output missing total: %q", text)
	}
}
