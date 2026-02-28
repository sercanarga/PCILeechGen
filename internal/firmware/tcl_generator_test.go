package firmware

import (
	"strings"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

func makeTestContext() *donor.DeviceContext {
	cs := makeTestConfigSpace()

	return &donor.DeviceContext{
		Device: pci.PCIDevice{
			BDF:            pci.BDF{Domain: 0, Bus: 3, Device: 0, Function: 0},
			VendorID:       0x8086,
			DeviceID:       0x1533,
			SubsysVendorID: 0x8086,
			SubsysDeviceID: 0x0001,
			RevisionID:     0x03,
			ClassCode:      0x020000,
		},
		ConfigSpace: cs,
		BARs: []pci.BAR{
			{Index: 0, Type: pci.BARTypeMem32, Address: 0xFE000000, Size: 1048576},
		},
	}
}

func TestGenerateProjectTCL(t *testing.T) {
	ctx := makeTestContext()
	b, _ := board.Find("PCIeSquirrel")
	libDir := "/path/to/lib/pcileech-fpga"

	tcl := GenerateProjectTCL(ctx, b, libDir)

	// Should contain FPGA part
	if !strings.Contains(tcl, b.FPGAPart) {
		t.Errorf("TCL should contain FPGA part %q", b.FPGAPart)
	}

	// Should contain board name
	if !strings.Contains(tcl, b.Name) {
		t.Errorf("TCL should contain board name %q", b.Name)
	}

	// Should contain donor device info in comments
	if !strings.Contains(tcl, "8086:1533") {
		t.Error("TCL should contain donor device IDs in comments")
	}

	// Should contain top module
	if !strings.Contains(tcl, b.TopModule) {
		t.Errorf("TCL should contain top module %q", b.TopModule)
	}

	// Should contain COE file references
	if !strings.Contains(tcl, "pcileech_cfgspace.coe") {
		t.Error("TCL should reference pcileech_cfgspace.coe")
	}
}

func TestGenerateBuildTCL(t *testing.T) {
	b, _ := board.Find("PCIeSquirrel")

	tcl := GenerateBuildTCL(b, 8, 7200)

	if !strings.Contains(tcl, "-jobs 8") {
		t.Error("Build TCL should contain -jobs 8")
	}
	if !strings.Contains(tcl, "-timeout 7200") {
		t.Error("Build TCL should contain timeout")
	}
	if !strings.Contains(tcl, "write_cfgmem") {
		t.Error("Build TCL should contain write_cfgmem for .bin generation")
	}
	if !strings.Contains(tcl, "exit 0") {
		t.Error("Build TCL should exit cleanly on success")
	}
}

func TestGenerateBuildTCLDefaults(t *testing.T) {
	b, _ := board.Find("PCIeSquirrel")

	tcl := GenerateBuildTCL(b, 0, 0)

	if !strings.Contains(tcl, "-jobs 4") {
		t.Error("Build TCL should use default 4 jobs")
	}
}

func TestProjectTCLIPCoreProperties(t *testing.T) {
	ctx := makeTestContext()
	b, _ := board.Find("PCIeSquirrel")
	libDir := "/path/to/lib/pcileech-fpga"

	tcl := GenerateProjectTCL(ctx, b, libDir)

	// Should contain IP core property patching
	checks := []struct {
		name    string
		content string
	}{
		{"Device_ID", "CONFIG.Device_ID            1533"},
		{"Vendor_Id", "CONFIG.Vendor_Id            8086"},
		{"Revision_ID", "CONFIG.Revision_ID          03"},
		{"Subsystem_Vendor_ID", "CONFIG.Subsystem_Vendor_ID  8086"},
		{"Subsystem_ID", "CONFIG.Subsystem_ID         0001"},
		{"Class_Code_Base", "CONFIG.Class_Code_Base      02"},
		{"Class_Code_Sub", "CONFIG.Class_Code_Sub       00"},
		{"Link_Speed", "CONFIG.Link_Speed"},
		{"Maximum_Link_Width", "CONFIG.Maximum_Link_Width"},
		{"BAR0 size", "CONFIG.Bar0_Scale"},
		{"pcie_7x_0", "get_ips -quiet pcie_7x_0"},
	}

	for _, c := range checks {
		if !strings.Contains(tcl, c.content) {
			t.Errorf("TCL should contain %s (%q)", c.name, c.content)
		}
	}
}

func TestClampLinkWidth(t *testing.T) {
	tests := []struct {
		name       string
		donorWidth uint8
		boardLanes int
		want       uint8
	}{
		{"donor x4 on x1 board", 4, 1, 1},
		{"donor x1 on x4 board", 1, 4, 1},
		{"donor x4 on x4 board", 4, 4, 4},
		{"donor x8 on x4 board", 8, 4, 4},
		{"donor 0 (unknown)", 0, 4, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := clampLinkWidth(tt.donorWidth, tt.boardLanes)
			if got != tt.want {
				t.Errorf("clampLinkWidth(%d, %d) = %d, want %d",
					tt.donorWidth, tt.boardLanes, got, tt.want)
			}
		})
	}
}

func TestBarSizeToTCL(t *testing.T) {
	tests := []struct {
		size      uint64
		wantScale string
		wantSize  string
	}{
		{0, "Kilobytes", "4"},
		{4096, "Kilobytes", "4"},
		{65536, "Kilobytes", "64"},
		{1048576, "Megabytes", "1"},
		{16777216, "Megabytes", "16"},
		{1024, "Kilobytes", "4"}, // min 4KB
	}
	for _, tt := range tests {
		scale, size := barSizeToTCL(tt.size)
		if scale != tt.wantScale || size != tt.wantSize {
			t.Errorf("barSizeToTCL(%d) = (%q, %q), want (%q, %q)",
				tt.size, scale, size, tt.wantScale, tt.wantSize)
		}
	}
}
