package svgen_test

import (
	"strings"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/firmware"
	"github.com/sercanarga/pcileechgen/internal/firmware/svgen"
)

func TestGenerateBarControllerSV_ILAInstance(t *testing.T) {
	base := &svgen.SVGeneratorConfig{
		DeviceIDs:   firmware.DeviceIDs{VendorID: 0x1102, DeviceID: 0x0012},
		BARModel:    hdaTestBARModel(),
		DeviceClass: "audio",
		ClassCode:   0x040300,
		PRNGSeeds:   [4]uint32{1, 2, 3, 4},
	}

	without, err := svgen.GenerateBarControllerSV(base)
	if err != nil {
		t.Fatalf("render without ILA: %v", err)
	}
	if strings.Contains(without, "u_ila_dbg") {
		t.Error("controller should not contain ILA instance when ILAInstanceSV is empty")
	}

	base.ILAInstanceSV = firmware.ILAInstanceSV()
	with, err := svgen.GenerateBarControllerSV(base)
	if err != nil {
		t.Fatalf("render with ILA: %v", err)
	}
	if !strings.Contains(with, "ila_0 u_ila_dbg") || !strings.Contains(with, ".probe0(wr_addr)") {
		t.Errorf("controller missing ILA instance:\n%s", with[len(with)-400:])
	}
}
