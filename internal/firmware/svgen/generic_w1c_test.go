package svgen_test

import (
	"strings"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/firmware"
	"github.com/sercanarga/pcileechgen/internal/firmware/barmodel"
	"github.com/sercanarga/pcileechgen/internal/firmware/svgen"
)

// A register carrying W1CMask (but not IsRW1C/IsFSMDriven) is emitted by the
// generic write case with a per-byte blend: RW bits take wr_data, W1C bits clear
// on write-of-1, RO bits hold.
func TestGenericW1CEmitter_BlendedExpression(t *testing.T) {
	cfg := &svgen.SVGeneratorConfig{
		DeviceIDs:   firmware.DeviceIDs{VendorID: 0x1234, DeviceID: 0x5678, RevisionID: 0x01},
		DeviceClass: "",
		PRNGSeeds:   [4]uint32{1, 2, 3, 4},
		BARModel: &barmodel.BARModel{
			Size: 4096,
			Registers: []barmodel.BARRegister{
				{Offset: 0x10, Width: 4, Name: "STATUS_CTL", Reset: 0, RWMask: 0x00000003, W1CMask: 0x00000100},
				{Offset: 0x14, Width: 4, Name: "CTL2", Reset: 0, RWMask: 0x000000FF},
			},
		},
	}
	sv, err := svgen.GenerateBarImplDeviceSV(cfg)
	if err != nil {
		t.Fatalf("generate SV: %v", err)
	}
	if !strings.Contains(sv, "32'h00000010: begin") {
		t.Error("W1C register 0x10 should be in the generic write case")
	}
	if !strings.Contains(sv, "reg_0x00000010[7:0]   <= ((reg_0x00000010[7:0]") {
		t.Error("W1C register byte0 should use the blended expression")
	}
	if !strings.Contains(sv, "~wr_data[15:8]") {
		t.Error("W1C byte1 should contain a clear-on-write-1 term")
	}
	if !strings.Contains(sv, "& ~wr_data[15:8]  & 8'h01") {
		t.Error("W1C byte1 clear term should mask with 8'h01")
	}
	if !strings.Contains(sv, "(wr_data[7:0]   & 8'h03)") {
		t.Error("W1C register byte0 RW bits should blend wr_data & 8'h03")
	}
	if !strings.Contains(sv, "32'h00000014: begin") {
		t.Error("plain RW register 0x14 should be in the generic write case")
	}
	if strings.Contains(sv, "reg_0x00000014[7:0]   <= ((reg_0x00000014[7:0]") {
		t.Error("plain RW register must not use the blended expression")
	}
}

// A synthesized device-driven register (xHCI USBCMD) stays out of the generic
// write case so it is not driven twice.
func TestSynthesizedXHCI_FSMRegExcludedFromGenericWriteCase(t *testing.T) {
	profile := &donor.BARProfile{Size: 4096, Probes: []donor.BARProbeResult{
		{Offset: 0x20, Original: 0x00000001, RWMask: 0x00002F0E},
	}}
	model := barmodel.SynthesizeBARModel(profile, 0x0C0330)
	if model == nil {
		t.Fatal("model nil")
	}
	cfg := &svgen.SVGeneratorConfig{
		DeviceIDs:   firmware.DeviceIDs{VendorID: 0x8086, DeviceID: 0x1E31, RevisionID: 0x04},
		DeviceClass: "xhci",
		PRNGSeeds:   [4]uint32{1, 2, 3, 4},
		BARModel:    model,
	}
	sv, err := svgen.GenerateBarImplDeviceSV(cfg)
	if err != nil {
		t.Fatalf("generate: %v", err)
	}
	if strings.Contains(sv, "32'h00000020: begin") {
		t.Error("USBCMD must not appear in the generic write case")
	}
}

func TestAudioHandCodedW1CHandlers_RenderFromRealModel(t *testing.T) {
	model := barmodel.BuildBARModel(make([]byte, 4096), 0x040300, nil, 0)
	if model == nil {
		t.Fatal("audio model nil")
	}
	cfg := &svgen.SVGeneratorConfig{
		DeviceIDs:   firmware.DeviceIDs{VendorID: 0x8086, DeviceID: 0x1A98, RevisionID: 0x00},
		DeviceClass: "audio",
		PRNGSeeds:   [4]uint32{1, 2, 3, 4},
		BARModel:    model,
	}
	sv, err := svgen.GenerateBarImplDeviceSV(cfg)
	if err != nil {
		t.Fatalf("generate: %v", err)
	}
	for _, pattern := range []string{
		"reg_0x0000000C[23:16] <= reg_0x0000000C[23:16] & ~wr_data[23:16]",
		"reg_0x0000000C[31:24] <= reg_0x0000000C[31:24] & ~wr_data[31:24]",
		"intfl_clear_req",
		"reg_0x00000060[0]",
	} {
		if !strings.Contains(sv, pattern) {
			t.Errorf("expected audio W1C pattern %q", pattern)
		}
	}
}
