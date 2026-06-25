package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/board"
)

func TestRunFlash_printsDryRunCommand_whenOpenFPGALoaderBoardPresent(t *testing.T) {
	// Given
	b := &board.Board{
		Name: "TestBoard",
		Flash: &board.Flash{
			Tool:                "openFPGALoader",
			OpenFPGALoaderBoard: "testboard",
			Target:              "flash",
		},
	}
	opts := flashOptions{bitstream: "pcileech_top.bit", dryRun: true}
	var out bytes.Buffer

	// When
	err := runFlash(&out, b, opts)

	// Then
	if err != nil {
		t.Fatalf("runFlash returned error: %v", err)
	}
	got := out.String()
	if !strings.Contains(got, "openFPGALoader -b testboard -f pcileech_top.bit") {
		t.Fatalf("dry-run output = %q", got)
	}
}
