package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseTraceBARBase(t *testing.T) {
	got, err := parseTraceBARBase("f7800000")
	if err != nil {
		t.Fatalf("parseTraceBARBase returned error: %v", err)
	}
	if got != 0xf7800000 {
		t.Fatalf("bar base = 0x%X, want 0xF7800000", got)
	}
}

func TestLoadMMIOTrace_FromTraceFile(t *testing.T) {
	tracePath := filepath.Join(t.TempDir(), "mmiotrace.txt")
	input := []byte("R 4 2456.105919 2 0xf780010c 0x4c02 0x0 0\n")
	if err := os.WriteFile(tracePath, input, 0o644); err != nil {
		t.Fatalf("write trace fixture: %v", err)
	}

	trace, err := loadMMIOTrace(mmioTraceOptions{
		bdf:       "0000:03:00.0",
		barIndex:  2,
		barSize:   4096,
		traceFile: tracePath,
	}, 0xf7800000)

	if err != nil {
		t.Fatalf("loadMMIOTrace returned error: %v", err)
	}
	if trace.BDF != "0000:03:00.0" || trace.BARIndex != 2 || trace.BARSize != 4096 {
		t.Fatalf("trace metadata = %+v", trace)
	}
	if len(trace.Records) != 1 || trace.Records[0].Offset != 0x10c || trace.Records[0].Value != 0x4c02 {
		t.Fatalf("records = %+v", trace.Records)
	}
}
