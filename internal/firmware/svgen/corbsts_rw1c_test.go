package svgen_test

import (
	"strings"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/firmware"
	"github.com/sercanarga/pcileechgen/internal/firmware/barmodel"
	"github.com/sercanarga/pcileechgen/internal/firmware/svgen"
)

// TestCORBSTSRW1CGeneration verifies that CORBSTS.RPWP (bit 8 of 0x4C)
// is handled as RW1C in the generated SystemVerilog, not as a generic
// writable bit. The HD Audio spec requires RPWP to be write-1-to-clear.
func TestCORBSTSRW1CGeneration(t *testing.T) {
	cfg := &svgen.SVGeneratorConfig{
		DeviceIDs: firmware.DeviceIDs{
			VendorID: 0x1102, DeviceID: 0x0012, RevisionID: 0x03,
		},
		BARModel: &barmodel.BARModel{
			Size: 4096,
			Registers: []barmodel.BARRegister{
				{Offset: 0x00, Width: 4, Name: "GCAP_VMIN_VMAJ", Reset: 0x01004401, RWMask: 0x00000000},
				{Offset: 0x08, Width: 4, Name: "GCTL", Reset: 0x00000001, RWMask: 0x00000103},
				{Offset: 0x0C, Width: 4, Name: "WAKEEN_STATESTS", Reset: 0x00010000, RWMask: 0x0000FFFF},
				{Offset: 0x4C, Width: 4, Name: "CORBCTL_STS_SIZE", Reset: 0x00420000, RWMask: 0x00000082, IsRW1C: true},
				{Offset: 0x5C, Width: 4, Name: "RIRBCTL_STS_SIZE", Reset: 0x00420000, RWMask: 0x00000307, IsRW1C: true},
				{Offset: 0x60, Width: 4, Name: "RIRBINTSTS", Reset: 0x00000000, RWMask: 0x00000001, IsRW1C: true},
			},
		},
		DeviceClass: "audio",
		PRNGSeeds:   [4]uint32{1, 2, 3, 4},
	}
	sv, err := svgen.GenerateBarImplDeviceSV(cfg)
	if err != nil {
		t.Fatalf("generate SV: %v", err)
	}

	// Verify dedicated RW1C handler exists using actual generated patterns.
	// Template uses corbsts_clear_bit wire (= wr_data[8]) for RW1C bit.
	for _, pattern := range []string{
		"corbsts_rw1c",
		"CORBCTL/CORBSTS",
		"corbsts_clear_bit",
		"reg_0x0000004C[8] <= reg_0x0000004C[8] & ~corbsts_clear_bit",
		"reg_0x0000004C[1] <= corbctl_wr_run_val",
	} {
		if !strings.Contains(sv, pattern) {
			t.Errorf("expected %q in generated SV", pattern)
		}
	}

	// Verify 0x4C is NOT in the generic write case block.
	// The generic write case has the pattern "32'hXXXX: begin" followed by
	// wr_be lines. The read case has "32'hXXXX: rd_data_d1 <=" - we only
	// want to ensure the write case is absent.
	// Search for the write-case pattern: "32'h0000004C: begin"
	if strings.Contains(sv, "32'h0000004C: begin") {
		t.Error("0x4C should NOT be in generic write case (IsRW1C=true)")
	}
}
