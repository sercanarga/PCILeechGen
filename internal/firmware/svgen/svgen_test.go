package svgen

import (
	"strings"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/firmware"
	"github.com/sercanarga/pcileechgen/internal/firmware/barmodel"
)

func testConfig() *SVGeneratorConfig {
	return &SVGeneratorConfig{
		DeviceIDs: firmware.DeviceIDs{
			VendorID:       0x144D,
			DeviceID:       0xA808,
			SubsysVendorID: 0x144D,
			SubsysDeviceID: 0xA801,
			RevisionID:     0x00,
			ClassCode:      0x010802,
			HasPCIeCap:     true,
			LinkSpeed:      3,
			LinkWidth:      4,
		},
		ClassCode:    0x010802,
		IsNVMe:       true,
		BuildEntropy: 0xDEADBEEF,
		PRNGSeeds:    BuildPRNGSeeds(0x144D, 0xA808, 0xDEADBEEF),
	}
}

func TestGenerateDeviceConfigSV(t *testing.T) {
	cfg := testConfig()
	result, err := GenerateDeviceConfigSV(cfg)
	if err != nil {
		t.Fatalf("GenerateDeviceConfigSV failed: %v", err)
	}
	if !strings.Contains(result, "144D") {
		t.Error("output should contain vendor ID 144D")
	}
	if !strings.Contains(result, "A808") {
		t.Error("output should contain device ID A808")
	}
	if !strings.Contains(result, "HAS_NVME_FSM") {
		t.Error("output should contain HAS_NVME_FSM")
	}
}

func TestGenerateBarImplDeviceSV_NilBARModel(t *testing.T) {
	cfg := testConfig()
	cfg.BARModel = nil
	result, err := GenerateBarImplDeviceSV(cfg)
	if err != nil {
		t.Fatalf("GenerateBarImplDeviceSV failed: %v", err)
	}
	if !strings.Contains(result, "bram_bar_zero4k") {
		t.Error("nil BARModel should fall back to BRAM-based implementation")
	}
}

func TestGenerateBarImplDeviceSV_WithBARModel(t *testing.T) {
	cfg := testConfig()
	cfg.BARModel = &barmodel.BARModel{
		Size: 4096,
		Registers: []barmodel.BARRegister{
			{Offset: 0x00, Width: 4, Name: "CAP_LO", RWMask: 0x00000000, Reset: 0x0040FF17},
			{Offset: 0x14, Width: 4, Name: "CC", RWMask: 0x00FFFFF1, Reset: 0x00000001},
			{Offset: 0x1C, Width: 4, Name: "CSTS", RWMask: 0x00000000, Reset: 0x00000001},
		},
	}
	result, err := GenerateBarImplDeviceSV(cfg)
	if err != nil {
		t.Fatalf("GenerateBarImplDeviceSV failed: %v", err)
	}
	if !strings.Contains(result, "reg_0x00000014") {
		t.Error("output should contain CC register")
	}
	if !strings.Contains(result, "cc_en_prev") {
		t.Error("NVMe output should contain CC→CSTS state machine")
	}
}

func TestGenerateBarControllerSV(t *testing.T) {
	cfg := testConfig()
	cfg.LatencyConfig = DefaultLatencyConfig(cfg.ClassCode)
	result, err := GenerateBarControllerSV(cfg)
	if err != nil {
		t.Fatalf("GenerateBarControllerSV failed: %v", err)
	}
	if !strings.Contains(result, "tlp_latency_emulator") {
		t.Error("output should contain latency emulator instantiation")
	}
}

func TestGenerateLatencyEmulatorSV(t *testing.T) {
	cfg := testConfig()
	cfg.LatencyConfig = DefaultLatencyConfig(cfg.ClassCode)
	result, err := GenerateLatencyEmulatorSV(cfg)
	if err != nil {
		t.Fatalf("GenerateLatencyEmulatorSV failed: %v", err)
	}
	if !strings.Contains(result, "module tlp_latency_emulator") {
		t.Error("output should contain tlp_latency_emulator module declaration")
	}
	if !strings.Contains(result, "xorshift128") {
		t.Error("output should contain PRNG description")
	}
}

func TestBuildPRNGSeeds_Deterministic(t *testing.T) {
	s1 := BuildPRNGSeeds(0x144D, 0xA808, 42)
	s2 := BuildPRNGSeeds(0x144D, 0xA808, 42)
	if s1 != s2 {
		t.Error("same inputs should produce same seeds")
	}
	s3 := BuildPRNGSeeds(0x144D, 0xA808, 43)
	if s1 == s3 {
		t.Error("different entropy should produce different seeds")
	}
}

func TestDefaultLatencyConfig(t *testing.T) {
	nvme := DefaultLatencyConfig(0x010802)
	if nvme.MinCycles != 3 || nvme.MaxCycles != 12 {
		t.Errorf("NVMe latency: got min=%d max=%d, want 3/12", nvme.MinCycles, nvme.MaxCycles)
	}
	xhci := DefaultLatencyConfig(0x0C0330)
	if xhci.MinCycles != 4 {
		t.Errorf("xHCI latency min: got %d, want 4", xhci.MinCycles)
	}
	eth := DefaultLatencyConfig(0x020000)
	if eth.MinCycles != 2 {
		t.Errorf("Ethernet latency min: got %d, want 2", eth.MinCycles)
	}
	generic := DefaultLatencyConfig(0xFF0000)
	if generic.MinCycles != 3 {
		t.Errorf("Generic latency min: got %d, want 3", generic.MinCycles)
	}
}
