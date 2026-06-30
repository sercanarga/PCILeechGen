package tclgen

import (
	"strings"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

func ilaTestCtx() (*donor.DeviceContext, *board.Board) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	cs.WriteU16(0x00, 0x8086)
	cs.WriteU16(0x02, 0x1533)
	return &donor.DeviceContext{
		Device:      pci.PCIDevice{VendorID: 0x8086, DeviceID: 0x1533, ClassCode: 0x020000},
		ConfigSpace: cs,
		BARs:        []pci.BAR{},
	}, &board.Board{Name: "TestBoard", FPGAPart: "xc7a35tfgg484-2", PCIeLanes: 1, TopModule: "test_top"}
}

func TestGenerateProjectTCL_ILAEnabled(t *testing.T) {
	ctx, b := ilaTestCtx()
	tcl := GenerateProjectTCL(ctx, b, "/tmp/lib", false, 2048)
	for _, want := range []string{"create_ip -name ila", "-module_name ila_0", "CONFIG.C_DATA_DEPTH {2048}", "generate_target all [get_ips ila_0]"} {
		if !strings.Contains(tcl, want) {
			t.Errorf("ILA-enabled TCL missing %q", want)
		}
	}
}

func TestGenerateProjectTCL_ILADisabled(t *testing.T) {
	ctx, b := ilaTestCtx()
	tcl := GenerateProjectTCL(ctx, b, "/tmp/lib", false, 0)
	if strings.Contains(tcl, "create_ip -name ila") {
		t.Error("ILA must not be emitted when depth is 0")
	}
}
