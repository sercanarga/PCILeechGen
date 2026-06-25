package svgen_test

import (
	"strings"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/firmware"
	"github.com/sercanarga/pcileechgen/internal/firmware/barmodel"
	"github.com/sercanarga/pcileechgen/internal/firmware/svgen"
)

func TestGenericRW1CWriteHandling(t *testing.T) {
	cfg := &svgen.SVGeneratorConfig{
		DeviceIDs: firmware.DeviceIDs{
			VendorID: 0x1AF4,
			DeviceID: 0x1000,
		},
		BARModel: &barmodel.BARModel{
			Size: 4096,
			Registers: []barmodel.BARRegister{
				{
					Offset:   0x10,
					Width:    4,
					Name:     "MIXED_RW1C",
					Reset:    0x00000000,
					RWMask:   0x00000003,
					RW1CMask: 0x00000002,
				},
			},
		},
		PRNGSeeds: [4]uint32{1, 2, 3, 4},
	}

	sv, err := svgen.GenerateBarImplDeviceSV(cfg)
	if err != nil {
		t.Fatalf("GenerateBarImplDeviceSV: %v", err)
	}

	if !strings.Contains(sv, "32'h00000010: begin") {
		t.Fatalf("expected offset 0x10 in generic write case")
	}

	if !strings.Contains(sv, "reg_0x00000010[7:0] <= (((reg_0x00000010[7:0]") ||
		!strings.Contains(sv, "& ~8'h01") ||
		!strings.Contains(sv, "& ~(8'h02 & wr_data[7:0]") {
		t.Errorf("expected write-1-to-clear semantics in generated write path")
	}
}
