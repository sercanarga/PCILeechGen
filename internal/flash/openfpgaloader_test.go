package flash

import (
	"strings"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/board"
)

func TestOpenFPGALoaderCommand_returnsDryRunArgs_whenMetadataPresent(t *testing.T) {
	// Given
	b := &board.Board{
		Name: "TestBoard",
		Flash: &board.Flash{
			Tool:                "openFPGALoader",
			OpenFPGALoaderBoard: "testboard",
			Cable:               "ft2232",
			Target:              "flash",
		},
	}
	req := Request{Board: b, Bitstream: "pcileech_top.bit"}

	// When
	cmd, err := OpenFPGALoaderCommand(req)

	// Then
	if err != nil {
		t.Fatalf("OpenFPGALoaderCommand returned error: %v", err)
	}
	want := []string{"openFPGALoader", "-b", "testboard", "-c", "ft2232", "-f", "pcileech_top.bit"}
	if strings.Join(cmd.Args, "\x00") != strings.Join(want, "\x00") {
		t.Fatalf("Args = %#v, want %#v", cmd.Args, want)
	}
}

func TestOpenFPGALoaderCommand_returnsUnsupportedBoard_whenFlashMetadataMissing(t *testing.T) {
	// Given
	req := Request{Board: &board.Board{Name: "NoFlash"}, Bitstream: "pcileech_top.bit"}

	// When
	_, err := OpenFPGALoaderCommand(req)

	// Then
	if err == nil {
		t.Fatal("OpenFPGALoaderCommand should reject boards without flash metadata")
	}
	if !strings.Contains(err.Error(), "board NoFlash has no flash metadata") {
		t.Fatalf("error = %q", err)
	}
}

func TestOpenFPGALoaderCommand_returnsUnsupportedTool_whenBoardUsesDifferentProvider(t *testing.T) {
	// Given
	b := &board.Board{
		Name: "OpenOCDBoard",
		Flash: &board.Flash{
			Tool:   "OpenOCD",
			Target: "flash",
		},
	}
	req := Request{Board: b, Bitstream: "pcileech_top.bit"}

	// When
	_, err := OpenFPGALoaderCommand(req)

	// Then
	if err == nil {
		t.Fatal("OpenFPGALoaderCommand should reject non-openFPGALoader metadata")
	}
	if !strings.Contains(err.Error(), "flash tool \"OpenOCD\"") {
		t.Fatalf("error = %q", err)
	}
}
