package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestRunTraceImport_WritesBarModelReport(t *testing.T) {
	tmpDir := t.TempDir()
	input := filepath.Join(tmpDir, "trace.log")
	output := filepath.Join(tmpDir, "bar_model_report.json")
	if err := os.WriteFile(input, []byte("R 4 0.001 0xfebf001c 0x00000000\nR 4 0.002 0xfebf001c 0x00000001\nR 4 0.003 0xfebf001c 0x00000001\n"), 0644); err != nil {
		t.Fatalf("write trace fixture: %v", err)
	}

	if err := runTraceImport(input, output, "0000:03:00.0", 0, 4096, 0); err != nil {
		t.Fatalf("runTraceImport failed: %v", err)
	}

	data, err := os.ReadFile(output)
	if err != nil {
		t.Fatalf("read output: %v", err)
	}
	var got struct {
		Reports []struct {
			BDF       string `json:"bdf"`
			BARIndex  int    `json:"bar_index"`
			Registers []struct {
				Offset         uint32 `json:"offset"`
				Classification string `json:"classification"`
			} `json:"registers"`
		} `json:"reports"`
	}
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("output JSON did not parse: %v", err)
	}
	if len(got.Reports) != 1 || got.Reports[0].BDF != "0000:03:00.0" || got.Reports[0].BARIndex != 0 {
		t.Fatalf("unexpected report identity: %#v", got.Reports)
	}
	if len(got.Reports[0].Registers) != 1 || got.Reports[0].Registers[0].Classification != "polled" {
		t.Fatalf("unexpected register classification: %#v", got.Reports[0].Registers)
	}
}
