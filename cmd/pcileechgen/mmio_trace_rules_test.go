package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/sercanarga/pcileechgen/internal/donor/behavior"
	"github.com/sercanarga/pcileechgen/internal/donor/mmio"
)

func TestMMIOTraceCLI_PersistsRulesWithoutChangingRawTraceOutput(t *testing.T) {
	dir := t.TempDir()
	textPath := filepath.Join(dir, "mmiotrace.txt")
	rawPath := filepath.Join(dir, "trace.json")
	rulesPath := filepath.Join(dir, "rules.json")
	input := []byte(
		"W 4 1.000 0xf7800020 0x1\n" +
			"R 4 1.001 0xf7800024 0x0\n" +
			"R 4 1.002 0xf7800024 0x1\n" +
			"W 4 1.010 0xf7800020 0x1\n" +
			"R 4 1.011 0xf7800024 0x0\n" +
			"R 4 1.012 0xf7800024 0x1\n")
	if err := os.WriteFile(textPath, input, 0o644); err != nil {
		t.Fatalf("write trace fixture: %v", err)
	}

	previous := mmioTraceOpts
	t.Cleanup(func() { mmioTraceOpts = previous })
	mmioTraceOpts = mmioTraceOptions{
		bdf: "0000:03:00.0", duration: time.Second, barIndex: 2, barSize: 0x1000,
		barBase: "0xf7800000", traceFile: textPath, outputFile: rawPath, rulesOutput: rulesPath,
	}
	var stdout, stderr bytes.Buffer
	mmioTraceCmd.SetOut(&stdout)
	mmioTraceCmd.SetErr(&stderr)
	t.Cleanup(func() {
		mmioTraceCmd.SetOut(nil)
		mmioTraceCmd.SetErr(nil)
	})
	if err := mmioTraceCmd.RunE(mmioTraceCmd, nil); err != nil {
		t.Fatalf("mmio-trace RunE: %v", err)
	}

	rawJSON, err := os.ReadFile(rawPath)
	if err != nil {
		t.Fatalf("read raw trace output: %v", err)
	}
	var trace mmio.TraceResult
	if perr := json.Unmarshal(rawJSON, &trace); perr != nil {
		t.Fatalf("raw --output is no longer a trace artifact: %v", perr)
	}
	if len(trace.Records) != 6 {
		t.Fatalf("raw trace records = %d, want 6", len(trace.Records))
	}

	rulesJSON, err := os.ReadFile(rulesPath)
	if err != nil {
		t.Fatalf("read rules output: %v", err)
	}
	var rules behavior.RuleSet
	if err := json.Unmarshal(rulesJSON, &rules); err != nil {
		t.Fatalf("decode rules output: %v", err)
	}
	if err := behavior.Validate(&rules); err != nil {
		t.Fatalf("persisted rules are invalid: %v", err)
	}
	if rules.Version != behavior.RuleSchemaVersion || len(rules.Rules) == 0 {
		t.Fatalf("persisted rule set = %+v", rules)
	}
}
