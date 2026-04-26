package svgen_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/firmware"
	"github.com/sercanarga/pcileechgen/internal/firmware/barmodel"
	"github.com/sercanarga/pcileechgen/internal/firmware/svgen"
)

// audioBARModel returns the HD Audio BAR register map for CA0132.
func audioBARModel() *barmodel.BARModel {
	return &barmodel.BARModel{
		Size: 4096,
		Registers: []barmodel.BARRegister{
			{Offset: 0x00, Width: 4, Name: "GCAP_VMIN_VMAJ", Reset: 0x01004401, RWMask: 0x00000000},
			{Offset: 0x08, Width: 4, Name: "GCTL", Reset: 0x00000001, RWMask: 0x00000103},
			{Offset: 0x0C, Width: 4, Name: "WAKEEN_STATESTS", Reset: 0x00010000, RWMask: 0x0000FFFF},
			{Offset: 0x20, Width: 4, Name: "INTCTL", Reset: 0x00000000, RWMask: 0xC00000FF},
			{Offset: 0x24, Width: 4, Name: "INTSTS", Reset: 0x00000000, RWMask: 0x00000000},
			{Offset: 0x40, Width: 4, Name: "CORBLBASE", Reset: 0x00000000, RWMask: 0xFFFFFF80},
			{Offset: 0x44, Width: 4, Name: "CORBUBASE", Reset: 0x00000000, RWMask: 0xFFFFFFFF},
			{Offset: 0x48, Width: 4, Name: "CORBWP_CORBRP", Reset: 0x00000000, RWMask: 0x80FF00FF},
			{Offset: 0x4C, Width: 4, Name: "CORBCTL_STS_SIZE", Reset: 0x00820000, RWMask: 0x00000082, IsRW1C: true},
			{Offset: 0x50, Width: 4, Name: "RIRBLBASE", Reset: 0x00000000, RWMask: 0xFFFFFF80},
			{Offset: 0x54, Width: 4, Name: "RIRBUBASE", Reset: 0x00000000, RWMask: 0xFFFFFFFF},
			{Offset: 0x58, Width: 4, Name: "RIRBWP_RINTCNT", Reset: 0x00000000, RWMask: 0x800000FF},
			{Offset: 0x5C, Width: 4, Name: "RIRBCTL_STS_SIZE", Reset: 0x00820000, RWMask: 0x00000307, IsRW1C: true},
			{Offset: 0x60, Width: 4, Name: "RIRBINTSTS", Reset: 0x00000000, RWMask: 0x00000001},
			{Offset: 0x64, Width: 4, Name: "IC", Reset: 0x00000000, RWMask: 0xFFFFFFFF},
			{Offset: 0x68, Width: 4, Name: "IR", Reset: 0x00000000, RWMask: 0x00000000},
			{Offset: 0x70, Width: 4, Name: "RIRBRESP_LO", Reset: 0x00000000, RWMask: 0x00000000},
			{Offset: 0x74, Width: 4, Name: "RIRBRESP_HI", Reset: 0x00000000, RWMask: 0x00000000},
		},
	}
}

// TestAudioFullGeneration verifies the complete audio SV generation pipeline
// including the RIRB response ROM and RIRBRESP registers.
func TestAudioFullGeneration(t *testing.T) {
	barModel := audioBARModel()

	cfg := &svgen.SVGeneratorConfig{
		DeviceIDs: firmware.DeviceIDs{
			VendorID:   0x1102,
			DeviceID:   0x0012,
			RevisionID: 0x03,
		},
		BARModel:    barModel,
		DeviceClass: "audio",
		ClassCode:   0x040300, // HD Audio
		PRNGSeeds:   [4]uint32{0x12345678, 0x9ABCDEF0, 0xFEDCBA98, 0x76543210},
	}

	sv, err := svgen.GenerateBarImplDeviceSV(cfg)
	if err != nil {
		t.Fatalf("GenerateBarImplDeviceSV: %v", err)
	}

	// Check RIRB ROM entries (discovery phase: 0-19)
	for i := 0; i <= 19; i++ {
		pattern := fmt.Sprintf("6'd%d:", i)
		if !strings.Contains(sv, pattern) {
			t.Errorf("ROM entry %d (%q) not found in generated SV", i, pattern)
		}
	}

	// Check runtime phase entries (20-31)
	for i := 20; i <= 31; i++ {
		pattern := fmt.Sprintf("6'd%d:", i)
		if !strings.Contains(sv, pattern) {
			t.Errorf("ROM runtime entry %d (%q) not found in generated SV", i, pattern)
		}
	}

	// Check extended runtime phase entries (32-63)
	for i := 32; i <= 63; i++ {
		pattern := fmt.Sprintf("6'd%d:", i)
		if !strings.Contains(sv, pattern) {
			t.Errorf("ROM extended runtime entry %d (%q) not found in generated SV", i, pattern)
		}
	}

	// Check ROM lookup call
	if !strings.Contains(sv, "rirb_rom_response(rirb_response_idx)") {
		t.Error("ROM lookup call not found")
	}

	// Check RIRBRESP register declarations exist (from audio profile)
	if !strings.Contains(sv, "reg_0x00000070") {
		t.Error("reg_0x00000070 (RIRBRESP_LO) declaration not found")
	}
	if !strings.Contains(sv, "reg_0x00000074") {
		t.Error("reg_0x00000074 (RIRBRESP_HI) declaration not found")
	}

	// Check response index reset on CRST
	if !strings.Contains(sv, "rirb_response_idx  <= 6'd0") {
		t.Error("response index reset on CRST not found")
	}

	// Check key CA0132 response values
	keyResponses := []string{
		"01010001", // AFG parameters (NumSubNodes=16, StartNode=1, FuncType=AFG)
		"11020010", // Creative subsystem ID (SB Audigy FX)
		"01014010", // Line-out pin default config
		"01A19020", // Mic-in pin default config
	}
	for _, resp := range keyResponses {
		if !strings.Contains(sv, resp) {
			t.Errorf("response value %q not found in generated SV", resp)
		}
	}

	// Verify RIRBRESP are in the read case statement (driver can read responses)
	if !strings.Contains(sv, "32'h00000070: rd_data_d1 <= reg_0x00000070") {
		t.Error("RIRBRESP_LO not in read case statement")
	}
	if !strings.Contains(sv, "32'h00000074: rd_data_d1 <= reg_0x00000074") {
		t.Error("RIRBRESP_HI not in read case statement")
	}

	t.Logf("Generated %d bytes of audio BAR SystemVerilog", len(sv))
}

// TestBarControllerAudioWithLatency verifies the full controller template
// renders correctly with audio + latency emulation.
func TestBarControllerAudioWithLatency(t *testing.T) {
	barModel := audioBARModel()
	latCfg := svgen.DefaultLatencyConfig(0x040300)

	cfg := &svgen.SVGeneratorConfig{
		DeviceIDs: firmware.DeviceIDs{
			VendorID:   0x1102,
			DeviceID:   0x0012,
			RevisionID: 0x03,
		},
		BARModel:      barModel,
		DeviceClass:   "audio",
		ClassCode:     0x040300,
		LatencyConfig: latCfg,
		PRNGSeeds:     [4]uint32{0x12345678, 0x9ABCDEF0, 0xFEDCBA98, 0x76543210},
	}

	sv, err := svgen.GenerateBarControllerSV(cfg)
	if err != nil {
		t.Fatalf("GenerateBarControllerSV: %v", err)
	}

	// Check latency emulator is instantiated
	if !strings.Contains(sv, "tlp_latency_emulator") {
		t.Error("latency emulator not instantiated")
	}

	// Check circular buffer exists
	if !strings.Contains(sv, "bar0_buf_ctx") {
		t.Error("circular buffer not found")
	}

	// Check write backpressure
	if !strings.Contains(sv, "bar0_wr_ack") {
		t.Error("write backpressure signal not found")
	}

	// Check address reconstruction from ctx
	if !strings.Contains(sv, "bar0_req_addr_from_ctx") {
		t.Error("address reconstruction from ctx not found")
	}

	// Check buffer is 16-deep
	if !strings.Contains(sv, "5'd16") {
		t.Error("buffer full threshold (16) not found")
	}

	t.Logf("Generated %d bytes of BAR controller SystemVerilog", len(sv))
}

// TestAudioDeviceConfig verifies the device config package renders correctly.
func TestAudioDeviceConfig(t *testing.T) {
	barModel := audioBARModel()
	latCfg := svgen.DefaultLatencyConfig(0x040300)

	cfg := &svgen.SVGeneratorConfig{
		DeviceIDs: firmware.DeviceIDs{
			VendorID:   0x1102,
			DeviceID:   0x0012,
			RevisionID: 0x03,
		},
		BARModel:      barModel,
		DeviceClass:   "audio",
		ClassCode:     0x040300,
		LatencyConfig: latCfg,
		PRNGSeeds:     [4]uint32{0x12345678, 0x9ABCDEF0, 0xFEDCBA98, 0x76543210},
	}

	sv, err := svgen.GenerateDeviceConfigSV(cfg)
	if err != nil {
		t.Fatalf("GenerateDeviceConfigSV: %v", err)
	}

	required := []string{
		"VENDOR_ID",
		"DEVICE_ID",
		"CLASS_BASE",
		"CLASS_SUB",
		"HAS_AUDIO_FSM",
		"HAS_LATENCY_EMU",
		"1102", // vendor ID hex
		"0012", // device ID hex
	}
	for _, r := range required {
		if !strings.Contains(sv, r) {
			t.Errorf("device config missing %q", r)
		}
	}

	t.Logf("Generated %d bytes of device config SystemVerilog", len(sv))
}

// TestLatencyEmulatorAudio verifies the latency emulator renders for audio.
func TestLatencyEmulatorAudio(t *testing.T) {
	latCfg := svgen.DefaultLatencyConfig(0x040300)

	cfg := &svgen.SVGeneratorConfig{
		DeviceIDs: firmware.DeviceIDs{
			VendorID:   0x1102,
			DeviceID:   0x0012,
			RevisionID: 0x03,
		},
		LatencyConfig: latCfg,
		PRNGSeeds:     [4]uint32{0x12345678, 0x9ABCDEF0, 0xFEDCBA98, 0x76543210},
	}

	sv, err := svgen.GenerateLatencyEmulatorSV(cfg)
	if err != nil {
		t.Fatalf("GenerateLatencyEmulatorSV: %v", err)
	}

	required := []string{
		"prng_s0",
		"prng_s1",
		"prng_s2",
		"prng_s3",
		"prng_next_s0",
		"prng_next_s3",
		"thermal_counter",
		"thermal_offset",
		"computed_latency",
		"wr_latency",
		"req_ready",
		"rsp_valid",
	}
	for _, r := range required {
		if !strings.Contains(sv, r) {
			t.Errorf("latency emulator missing %q", r)
		}
	}

	// Verify PRNG advances every cycle (not gated on req_valid)
	if !strings.Contains(sv, "prng_s0 <= prng_s1") {
		t.Error("PRNG should advance every cycle, not gated on requests")
	}

	t.Logf("Generated %d bytes of latency emulator SystemVerilog", len(sv))
}
