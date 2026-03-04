package tclgen

import (
	"strings"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/donor"
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

	tcl := GenerateProjectTCL(ctx, b, "/tmp/lib")

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
	scale, size := barSizeToTCL(4096)
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
