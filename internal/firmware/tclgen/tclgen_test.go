package tclgen

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/firmware/barmodel"
	"github.com/sercanarga/pcileechgen/internal/firmware/svgen"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

func TestGenerateProjectTCL(t *testing.T) {
	b := &board.Board{
		Name:      "TestBoard",
		FPGAPart:  "xc7a35tfgg484-2",
		PCIeLanes: 1,
		TopModule: "test_top",
	}

	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	cs.WriteU16(0x00, 0x8086)
	cs.WriteU16(0x02, 0x1533)
	cs.WriteU8(0x08, 0x03)
	cs.WriteU8(0x09, 0x00)
	cs.WriteU8(0x0A, 0x00)
	cs.WriteU8(0x0B, 0x02)

	ctx := &donor.DeviceContext{
		Device: pci.PCIDevice{
			VendorID:  0x8086,
			DeviceID:  0x1533,
			ClassCode: 0x020000,
		},
		ConfigSpace:     cs,
		ExtCapabilities: nil,
		BARs:            []pci.BAR{},
	}

	tcl := GenerateProjectTCL(ctx, b, "/tmp/lib", false)

	if !strings.Contains(tcl, "8086") {
		t.Error("TCL should contain vendor ID")
	}
	if !strings.Contains(tcl, "1533") {
		t.Error("TCL should contain device ID")
	}
	if !strings.Contains(tcl, "xc7a35tfgg484-2") {
		t.Error("TCL should contain FPGA part")
	}
	if !strings.Contains(tcl, "test_top") {
		t.Error("TCL should contain top module name")
	}
	if strings.Contains(tcl, "SEED") {
		t.Error("TCL should not contain SEED property (not supported in Vivado 2023.2)")
	}
	if strings.Contains(tcl, "MORE OPTIONS") {
		t.Error("TCL should not use MORE OPTIONS for seed")
	}
	if strings.Contains(tcl, "src/*.v") {
		t.Error("TCL must NOT glob src/*.v for ImportVFiles=false (default); that risks binding a stale pcileech_com_e.v netlist over pcileech_com.sv")
	}
	if !strings.Contains(tcl, "set v_files") {
		t.Error("TCL must define `set v_files` even when empty (concat references it)")
	}
}

func TestTclPathNormalizesWindowsSeparators(t *testing.T) {
	if got, want := tclPath(`C:\Users\Admin\pcileech-fpga\board\ip`), `C:/Users/Admin/pcileech-fpga/board/ip`; got != want {
		t.Fatalf("tclPath = %q, want %q", got, want)
	}
}

func TestGeneratedMultiBARTCLSourcesInTclsh(t *testing.T) {
	tclsh, err := exec.LookPath("tclsh")
	if err != nil {
		t.Skip("tclsh is required to execute generated Tcl")
	}

	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	ctx := &donor.DeviceContext{
		Device:      pci.PCIDevice{VendorID: 0x1234, DeviceID: 0x5678, ClassCode: 0xFF0000},
		ConfigSpace: cs,
		BARs: []pci.BAR{
			{Index: 0, Size: 0x2000, Type: pci.BARTypeMem64, Is64Bit: true},
			{Index: 2, Size: 0x1000, Type: pci.BARTypeMem32, Prefetchable: true},
			{Index: 5, Size: 0x100000, Type: pci.BARTypeMem32},
		},
	}
	bar0 := &barmodel.BARModel{BIR: 0, Size: 0x2000, Type: pci.BARTypeMem64, Is64Bit: true}
	bar2 := &barmodel.BARModel{BIR: 2, Size: 0x1000, Type: pci.BARTypeMem32, Prefetchable: true}
	bar5 := &barmodel.BARModel{BIR: 5, Size: 0x100000, Type: pci.BARTypeMem32}
	cfg := &svgen.SVGeneratorConfig{
		DonorBARTopology: true,
		BARModels:        []*barmodel.BARModel{bar0, bar2, bar5},
		BARModel:         bar0,
	}
	b := &board.Board{Name: "TclSyntax", FPGAPart: "xc7a35tfgg484-2", PCIeLanes: 1, TopModule: "test_top"}
	generated := GenerateProjectTCLWithConfig(ctx, b, t.TempDir(), false, cfg)
	tclPath := filepath.Join(t.TempDir(), "vivado_generate_project.tcl")
	if werr := os.WriteFile(tclPath, []byte(generated), 0644); werr != nil {
		t.Fatal(werr)
	}
	harnessPath := filepath.Join(t.TempDir(), "source_generated.tcl")
	harness := `proc create_project args {}
proc get_property args { return "." }
proc current_project args { return "project" }
proc set_property args {}
proc get_filesets args { return "sources_1" }
proc create_fileset args {}
proc import_files args { return "" }
proc get_files args { return "" }
proc remove_files args {}
proc get_ips args { return "pcie_7x_0" }
proc upgrade_ip args {}
proc create_run args {}
proc get_runs args { return "run" }
proc current_run args {}
` + fmt.Sprintf("source {%s}\n", filepath.ToSlash(tclPath))
	if werr := os.WriteFile(harnessPath, []byte(harness), 0644); werr != nil {
		t.Fatal(werr)
	}
	output, err := exec.Command(tclsh, harnessPath).CombinedOutput()
	if err != nil {
		t.Fatalf("tclsh rejected generated multi-BAR Tcl: %v\n%s", err, output)
	}
}

func TestGenerateProjectTCL_ImportVFiles(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	ctx := &donor.DeviceContext{
		Device:      pci.PCIDevice{VendorID: 0x8086, DeviceID: 0x1533, ClassCode: 0x020000},
		ConfigSpace: cs,
	}

	// ImportVFiles=true ZDMA/GBOX
	bOn := &board.Board{
		Name:         "ZDMA",
		FPGAPart:     "xc7a100tfgg484-2",
		TopModule:    "pcileech_tbx4_100t_top",
		ImportVFiles: true,
	}
	tclOn := GenerateProjectTCL(ctx, bOn, "/tmp/lib", false)
	if !strings.Contains(tclOn, "src/*.v") {
		t.Error("ImportVFiles=true: TCL must glob src/*.v (board ships pcileech_com_e.v netlist)")
	}

	bOff := &board.Board{
		Name:      "CaptainDMA_75T",
		FPGAPart:  "xc7a75tfgg484-2",
		TopModule: "pcileech_75t484_x1_top",
	}
	tclOff := GenerateProjectTCL(ctx, bOff, "/tmp/lib", false)
	if strings.Contains(tclOff, "src/*.v") {
		t.Error("ImportVFiles=false: TCL must NOT glob src/*.v")
	}
}

func TestGenerateBuildTCL(t *testing.T) {
	b := &board.Board{
		Name:     "TestBoard",
		FPGAPart: "xc7a35tfgg484-2",
	}

	tcl := GenerateBuildTCL(b, 4, 3600)

	if !strings.Contains(tcl, "TestBoard") {
		t.Error("build TCL should contain board name")
	}
	if !strings.Contains(tcl, "-jobs 4") {
		t.Error("build TCL should contain jobs count")
	}
}

func TestLinkSpeedToTCL(t *testing.T) {
	tests := []struct {
		speed uint8
		want  string
	}{
		{1, "2.5_GT/s"},
		{2, "5.0_GT/s"},
		{3, "8.0_GT/s"},
		{0, "5.0_GT/s"},
	}
	for _, tt := range tests {
		got := linkSpeedToTCL(tt.speed)
		if got != tt.want {
			t.Errorf("linkSpeedToTCL(%d) = %q, want %q", tt.speed, got, tt.want)
		}
	}
}

func TestClampLinkWidth(t *testing.T) {
	if clampLinkWidth(4, 1) != 1 {
		t.Error("x4 donor should clamp to x1 board")
	}
	if clampLinkWidth(1, 4) != 1 {
		t.Error("x1 donor should stay x1 on x4 board")
	}
	if clampLinkWidth(0, 4) != 4 {
		t.Error("zero donor width should default to board lanes")
	}
}

func TestBarSizeToTCL(t *testing.T) {
	k4 := uint64(4096)
	scale, size := barSizeToTCL(k4)
	if scale != "Kilobytes" || size != "4" {
		t.Errorf("4KB: got %s/%s, want Kilobytes/4", scale, size)
	}
	scale, size = barSizeToTCL(1024 * 1024)
	if scale != "Megabytes" || size != "1" {
		t.Errorf("1MB: got %s/%s, want Megabytes/1", scale, size)
	}
	scale, size = barSizeToTCL(0)
	if scale != "Kilobytes" || size != "4" {
		t.Errorf("0: got %s/%s, want Kilobytes/4 (minimum)", scale, size)
	}
}

func TestGenerateProjectTCL_ValidatesCode10Parameters(t *testing.T) {
	b := &board.Board{
		Name:         "TestBoard",
		FPGAPart:     "xc7a35tfgg484-2",
		PCIeLanes:    1,
		TopModule:    "test_top",
		MaxLinkSpeed: 2,
	}

	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	cs.WriteU16(0x00, 0x144D)
	cs.WriteU16(0x02, 0xA809)
	cs.WriteU16(0x2C, 0x144D)
	cs.WriteU16(0x2E, 0xA809)
	cs.WriteU8(0x08, 0x01)
	cs.WriteU8(0x09, 0x08)
	cs.WriteU8(0x0A, 0x02)
	cs.WriteU8(0x0B, 0x01)

	ctx := &donor.DeviceContext{
		Device: pci.PCIDevice{
			VendorID:       0x144D,
			DeviceID:       0xA809,
			SubsysVendorID: 0x144D,
			SubsysDeviceID: 0xA809,
			ClassCode:      0x010802,
		},
		ConfigSpace: cs,
		MSIXData: &donor.MSIXData{
			TableSize:   5,
			TableBIR:    0,
			TableOffset: 0x2000,
			PBABIR:      0,
			PBAOffset:   0x3000,
		},
	}

	tcl := GenerateProjectTCL(ctx, b, "/tmp/lib", false)

	checks := []struct {
		param string
		want  string
	}{
		{"Max_Payload_Size", "128_bytes"},
		{"Extended_Tag_Field", "false"},
		{"Extended_Tag_Default", "false"},
		{"AER_Enabled", "true"},
		{"AER_Completion_Timeout", "true"},
		{"MSI_Enabled", "false"},
		{"MSIx_Enabled", "true"},
		{"Vendor_Id", "144D"},
		{"Device_ID", "A809"},
		{"Subsystem_Vendor_ID", "144D"},
		{"Subsystem_ID", "A809"},
	}

	for _, c := range checks {
		if !strings.Contains(tcl, c.param) || !strings.Contains(tcl, c.want) {
			t.Errorf("TCL missing %s = %s", c.param, c.want)
		}
	}
}
