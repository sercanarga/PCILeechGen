package main

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/board"
)

func TestRunBoardsJSON_EmitsMachineReadableBoards(t *testing.T) {
	var out bytes.Buffer

	if err := runBoards(&out, true); err != nil {
		t.Fatalf("runBoards returned error: %v", err)
	}

	var got []boardInfo
	if err := json.Unmarshal(out.Bytes(), &got); err != nil {
		t.Fatalf("boards JSON did not decode: %v", err)
	}
	if len(got) != len(board.All()) {
		t.Fatalf("board count = %d, want %d", len(got), len(board.All()))
	}
	if got[0].Name == "" || got[0].FPGAPart == "" || got[0].TopModule == "" || got[0].BRAMSize <= 0 {
		t.Fatalf("first board missing required fields: %+v", got[0])
	}
}
