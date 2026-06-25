package output

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/donor/mmio"
	"github.com/sercanarga/pcileechgen/internal/firmware"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

func TestNewOutputWriter_Defaults(t *testing.T) {
	ow := NewOutputWriter("/tmp/test", "lib/pcileech-fpga", 0, 0)
	if ow.Jobs != 4 {
		t.Errorf("default Jobs = %d, want 4", ow.Jobs)
	}
	if ow.Timeout != 3600 {
		t.Errorf("default Timeout = %d, want 3600", ow.Timeout)
	}
}

func TestNewOutputWriter_CustomValues(t *testing.T) {
	ow := NewOutputWriter("/tmp/test", "lib/pcileech-fpga", 8, 7200)
	if ow.Jobs != 8 {
		t.Errorf("Jobs = %d, want 8", ow.Jobs)
	}
	if ow.Timeout != 7200 {
		t.Errorf("Timeout = %d, want 7200", ow.Timeout)
	}
}

func TestWriteFile(t *testing.T) {
	tmpDir := t.TempDir()
	ow := NewOutputWriter(tmpDir, "lib/pcileech-fpga", 4, 3600)

	content := "test content"
	if err := ow.writeFile("test.txt", content); err != nil {
		t.Fatalf("writeFile failed: %v", err)
	}

	data, err := os.ReadFile(filepath.Join(tmpDir, "test.txt"))
	if err != nil {
		t.Fatalf("ReadFile failed: %v", err)
	}
	if string(data) != content {
		t.Errorf("content = %q, want %q", string(data), content)
	}
}

func TestWriteBARContentArtifacts_WritesCapturedBARsAndLegacyAlias(t *testing.T) {
	tmpDir := t.TempDir()
	ow := NewOutputWriter(tmpDir, "lib/pcileech-fpga", 4, 3600)
	ctx := &donor.DeviceContext{
		BARContents: map[int][]byte{
			0: {0x17, 0xFF, 0x40, 0x00},
			2: {0xDD, 0xCC, 0xBB, 0xAA},
		},
	}

	if err := ow.writeBARContentArtifacts(ctx, 4096); err != nil {
		t.Fatalf("writeBARContentArtifacts failed: %v", err)
	}

	for _, name := range []string{"pcileech_bar0.coe", "pcileech_bar2.coe", "pcileech_bar_zero4k.coe"} {
		if _, err := os.Stat(filepath.Join(tmpDir, name)); err != nil {
			t.Fatalf("expected %s to be written: %v", name, err)
		}
	}

	bar2, err := os.ReadFile(filepath.Join(tmpDir, "pcileech_bar2.coe"))
	if err != nil {
		t.Fatalf("read BAR2 COE: %v", err)
	}
	if !strings.Contains(string(bar2), "aabbccdd") {
		t.Fatalf("BAR2 COE missing BAR2 data: %s", string(bar2[:min(len(bar2), 200)]))
	}
}

func TestWriteTraceReportArtifact_WritesBarModelReport(t *testing.T) {
	tmpDir := t.TempDir()
	ow := NewOutputWriter(tmpDir, "lib/pcileech-fpga", 4, 3600)
	ctx := &donor.DeviceContext{
		MMIOTraces: map[int]*mmio.TraceResult{
			0: {
				BDF:      "0000:03:00.0",
				BARIndex: 0,
				BARSize:  4096,
				Records: []mmio.AccessRecord{
					{Offset: 0x1C, Type: mmio.AccessRead, Value: 0},
					{Offset: 0x1C, Type: mmio.AccessRead, Value: 0},
					{Offset: 0x1C, Type: mmio.AccessRead, Value: 1},
				},
			},
		},
	}

	if err := ow.writeTraceReportArtifact(ctx); err != nil {
		t.Fatalf("writeTraceReportArtifact failed: %v", err)
	}

	data, err := os.ReadFile(filepath.Join(tmpDir, "bar_model_report.json"))
	if err != nil {
		t.Fatalf("read bar_model_report.json: %v", err)
	}
	var got struct {
		Reports []struct {
			BARIndex  int `json:"bar_index"`
			Registers []struct {
				Offset         uint32 `json:"offset"`
				Classification string `json:"classification"`
			} `json:"registers"`
		} `json:"reports"`
	}
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("report JSON did not parse: %v", err)
	}
	if len(got.Reports) != 1 || got.Reports[0].BARIndex != 0 {
		t.Fatalf("unexpected reports: %#v", got.Reports)
	}
	if len(got.Reports[0].Registers) != 1 || got.Reports[0].Registers[0].Classification != "polled" {
		t.Fatalf("unexpected register classification: %#v", got.Reports[0].Registers)
	}
}

func TestBuildSVConfig_UsesTraceReportForBARModel(t *testing.T) {
	ow := NewOutputWriter(t.TempDir(), "lib/pcileech-fpga", 4, 3600)
	ctx := &donor.DeviceContext{
		Device: pci.PCIDevice{
			BDF:       pci.BDF{Bus: 3},
			ClassCode: 0xFF0000,
		},
		BARs: []pci.BAR{{Index: 0, Type: pci.BARTypeMem32, Size: 4096}},
		MMIOTraces: map[int]*mmio.TraceResult{
			0: {
				BDF:      "0000:03:00.0",
				BARIndex: 0,
				BARSize:  4096,
				Records: []mmio.AccessRecord{
					{Offset: 0x80, Type: mmio.AccessRead, Value: 0xA5A50001},
					{Offset: 0x80, Type: mmio.AccessRead, Value: 0xA5A50001},
				},
			},
		},
	}

	cfg, err := ow.buildSVConfig(ctx, pci.NewConfigSpace(), firmware.DeviceIDs{VendorID: 0x1234, DeviceID: 0x5678, ClassCode: 0xFF0000}, 1, &board.Board{BRAMSize: 4096})
	if err != nil {
		t.Fatalf("buildSVConfig failed: %v", err)
	}
	if cfg.BARModel == nil || len(cfg.BARModel.Registers) != 1 {
		t.Fatalf("trace-derived BARModel missing: %#v", cfg.BARModel)
	}
	if reg := cfg.BARModel.Registers[0]; reg.Offset != 0x80 || reg.Reset != 0xA5A50001 || reg.RWMask != 0 {
		t.Fatalf("trace-derived register mismatch: %#v", reg)
	}
}

func TestListOutputFiles(t *testing.T) {
	files := ListOutputFiles()
	if len(files) == 0 {
		t.Error("ListOutputFiles returned empty list")
	}

	// Check that critical files are listed
	expected := []string{
		"pcileech_cfgspace.coe",
		"pcileech_cfgspace_writemask.coe",
		"pcileech_bar_zero4k.coe",
		"pcileech_bar<N>.coe",
		"vivado_generate_project.tcl",
		"vivado_build.tcl",
		"device_context.json",
		"bar_model_report.json",
	}
	for _, name := range expected {
		found := false
		for _, f := range files {
			if strings.Contains(f, name) {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("ListOutputFiles missing %q", name)
		}
	}
}
