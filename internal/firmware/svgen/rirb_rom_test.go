package svgen_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/firmware"
	"github.com/sercanarga/pcileechgen/internal/firmware/barmodel"
	"github.com/sercanarga/pcileechgen/internal/firmware/svgen"
)

func TestAudioRIRBResponseROM(t *testing.T) {
	cfg := &svgen.SVGeneratorConfig{
		DeviceIDs: firmware.DeviceIDs{
			VendorID:   0x1102,
			DeviceID:   0x0012,
			RevisionID: 0x03,
		},
		BARModel: &barmodel.BARModel{
			Size: 4096,
			Registers: []barmodel.BARRegister{
				{Offset: 0x00, Width: 4, Reset: 0x00, RWMask: 0x0},
				{Offset: 0x08, Width: 4, Reset: 0x00, RWMask: 0xFFFFFFFF},
				{Offset: 0x0C, Width: 4, Reset: 0x00, RWMask: 0x0},
				{Offset: 0x48, Width: 4, Reset: 0x00, RWMask: 0x0},
				{Offset: 0x4C, Width: 4, Reset: 0x00, RWMask: 0x0},
				{Offset: 0x58, Width: 4, Reset: 0x00, RWMask: 0x0},
				{Offset: 0x5C, Width: 4, Reset: 0x00, RWMask: 0x0},
				{Offset: 0x70, Width: 4, Reset: 0x00, RWMask: 0x0},
				{Offset: 0x78, Width: 4, Reset: 0x00, RWMask: 0x0},
			},
		},
		DeviceClass: "audio",
		PRNGSeeds:   [4]uint32{0x12345678, 0x9ABCDEF0, 0xFEDCBA98, 0x76543210},
	}

	sv, err := svgen.GenerateBarImplDeviceSV(cfg)
	if err != nil {
		t.Fatalf("generate SV: %v", err)
	}

	// Verify ROM function exists
	if !strings.Contains(sv, "rirb_rom_response") {
		t.Fatal("rirb_rom_response function not found in generated SV")
	}

	// Verify ROM entries exist (6'd0 through 6'd19)
	for i := 0; i <= 19; i++ {
		pattern := fmt.Sprintf("6'd%d:", i)
		if !strings.Contains(sv, pattern) {
			t.Fatalf("ROM entry %d (%q) not found", i, pattern)
		}
	}

	// Verify key response values
	keyValues := []string{
		"00A00001", // AFG parameters (response 0)
		"11020001", // Creative subsystem ID (response 1)
		"CA0132",   // Codec reference in comment
	}
	for _, v := range keyValues {
		if !strings.Contains(sv, v) {
			t.Fatalf("expected %q in generated SV", v)
		}
	}

	// Verify response index counter and ROM lookup
	required := []string{
		"rirb_response_idx",
		"rirb_rom_response(rirb_response_idx)",
	}
	for _, r := range required {
		if !strings.Contains(sv, r) {
			t.Fatalf("expected %q in generated SV", r)
		}
	}

	// Verify response index reset on CRST
	if !strings.Contains(sv, "rirb_response_idx  <= 6'd0") {
		t.Fatal("response index reset not found on CRST")
	}

	t.Logf("Generated %d bytes with response ROM", len(sv))
}
