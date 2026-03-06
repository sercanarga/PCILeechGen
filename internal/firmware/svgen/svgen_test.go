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
		DeviceClass:  "nvme",
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

func xhciConfig() *SVGeneratorConfig {
	return &SVGeneratorConfig{
		DeviceIDs: firmware.DeviceIDs{
			VendorID:       0x8086,
			DeviceID:       0xA36D,
			SubsysVendorID: 0x8086,
			SubsysDeviceID: 0x7270,
			RevisionID:     0x10,
			ClassCode:      0x0C0330,
			HasPCIeCap:     true,
			LinkSpeed:      3,
			LinkWidth:      1,
		},
		ClassCode:    0x0C0330,
		DeviceClass:  "xhci",
		BuildEntropy: 0xCAFEBABE,
		PRNGSeeds:    BuildPRNGSeeds(0x8086, 0xA36D, 0xCAFEBABE),
		BARModel: &barmodel.BARModel{
			Size: 4096,
			Registers: []barmodel.BARRegister{
				{Offset: 0x00, Width: 4, Name: "CAPLENGTH_HCIVERSION", RWMask: 0x00000000, Reset: 0x01100020},
				{Offset: 0x20, Width: 4, Name: "USBCMD", RWMask: 0x00002F0E, Reset: 0x00000001},
				{Offset: 0x24, Width: 4, Name: "USBSTS", RWMask: 0x00000000, Reset: 0x00000000},
			},
		},
	}
}

func audioConfig() *SVGeneratorConfig {
	return &SVGeneratorConfig{
		DeviceIDs: firmware.DeviceIDs{
			VendorID:       0x8086,
			DeviceID:       0xA171,
			SubsysVendorID: 0x8086,
			SubsysDeviceID: 0x7270,
			RevisionID:     0x21,
			ClassCode:      0x040300,
			HasPCIeCap:     true,
			LinkSpeed:      3,
			LinkWidth:      1,
		},
		ClassCode:    0x040300,
		DeviceClass:  "audio",
		BuildEntropy: 0x12345678,
		PRNGSeeds:    BuildPRNGSeeds(0x8086, 0xA171, 0x12345678),
		BARModel: &barmodel.BARModel{
			Size: 4096,
			Registers: []barmodel.BARRegister{
				{Offset: 0x00, Width: 4, Name: "GCAP_VMIN_VMAJ", RWMask: 0x00000000, Reset: 0x01004401},
				{Offset: 0x08, Width: 4, Name: "GCTL", RWMask: 0x00000103, Reset: 0x00000001},
				{Offset: 0x0C, Width: 4, Name: "WAKEEN_STATESTS", RWMask: 0x7FFFFFFF, Reset: 0x00010000},
			},
		},
	}
}

func ethernetConfig() *SVGeneratorConfig {
	return &SVGeneratorConfig{
		DeviceIDs: firmware.DeviceIDs{
			VendorID:       0x8086,
			DeviceID:       0x15B7,
			SubsysVendorID: 0x8086,
			SubsysDeviceID: 0x0000,
			RevisionID:     0x02,
			ClassCode:      0x020000,
			HasPCIeCap:     true,
			LinkSpeed:      3,
			LinkWidth:      1,
		},
		ClassCode:    0x020000,
		DeviceClass:  "ethernet",
		BuildEntropy: 0xFEEDFACE,
		PRNGSeeds:    BuildPRNGSeeds(0x8086, 0x15B7, 0xFEEDFACE),
		BARModel: &barmodel.BARModel{
			Size: 32768,
			Registers: []barmodel.BARRegister{
				{Offset: 0x00, Width: 4, Name: "CTRL", RWMask: 0xFFFFFFFF, Reset: 0x00000000},
				{Offset: 0x08, Width: 4, Name: "STATUS", RWMask: 0x00000000, Reset: 0x00000082},
			},
		},
	}
}

func TestGenerateBarImplDeviceSV_XHCI(t *testing.T) {
	cfg := xhciConfig()
	result, err := GenerateBarImplDeviceSV(cfg)
	if err != nil {
		t.Fatalf("xHCI bar_impl_device failed: %v", err)
	}
	if !strings.Contains(result, "hcrst_cnt") {
		t.Error("xHCI output should contain HCRST auto-clear FSM")
	}
	if !strings.Contains(result, "reg_0x00000020") {
		t.Error("xHCI output should contain USBCMD register")
	}
}

func TestGenerateBarImplDeviceSV_Audio(t *testing.T) {
	cfg := audioConfig()
	result, err := GenerateBarImplDeviceSV(cfg)
	if err != nil {
		t.Fatalf("Audio bar_impl_device failed: %v", err)
	}
	if !strings.Contains(result, "crst_prev") {
		t.Error("Audio output should contain GCTL.CRST handshake FSM")
	}
	if !strings.Contains(result, "reg_0x0000000C") {
		t.Error("Audio output should contain WAKEEN_STATESTS register")
	}
}

func TestGenerateBarImplDeviceSV_Ethernet(t *testing.T) {
	cfg := ethernetConfig()
	result, err := GenerateBarImplDeviceSV(cfg)
	if err != nil {
		t.Fatalf("Ethernet bar_impl_device failed: %v", err)
	}
	// Ethernet has no FSM, but should have registers
	if !strings.Contains(result, "reg_0x00000000") {
		t.Error("Ethernet output should contain CTRL register")
	}
	// Ethernet should NOT have NVMe FSM
	if strings.Contains(result, "cc_en_prev") {
		t.Error("Ethernet should NOT contain NVMe FSM")
	}
}

func TestGenerateDeviceConfigSV_XHCI(t *testing.T) {
	cfg := xhciConfig()
	result, err := GenerateDeviceConfigSV(cfg)
	if err != nil {
		t.Fatalf("xHCI device_config failed: %v", err)
	}
	if !strings.Contains(result, "HAS_XHCI_FSM") {
		t.Error("xHCI device_config should contain HAS_XHCI_FSM")
	}
	if !strings.Contains(result, "8086") {
		t.Error("should contain vendor ID")
	}
}

func TestGenerateDeviceConfigSV_Audio(t *testing.T) {
	cfg := audioConfig()
	result, err := GenerateDeviceConfigSV(cfg)
	if err != nil {
		t.Fatalf("Audio device_config failed: %v", err)
	}
	if !strings.Contains(result, "HAS_AUDIO_FSM") {
		t.Error("Audio device_config should contain HAS_AUDIO_FSM")
	}
}

func TestGenerateDeviceConfigSV_NVMeFSMFlag(t *testing.T) {
	cfg := testConfig()
	result, err := GenerateDeviceConfigSV(cfg)
	if err != nil {
		t.Fatal(err)
	}
	// NVMe config should set HAS_NVME_FSM = 1
	if !strings.Contains(result, "HAS_NVME_FSM     = 1") {
		t.Error("NVMe HAS_NVME_FSM should be 1")
	}
	// should NOT set HAS_XHCI_FSM = 1
	if strings.Contains(result, "HAS_XHCI_FSM     = 1") {
		t.Error("NVMe config should NOT have HAS_XHCI_FSM = 1")
	}
}

func TestGenerateDeviceConfigSV_XHCIFSMFlag(t *testing.T) {
	cfg := xhciConfig()
	result, err := GenerateDeviceConfigSV(cfg)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(result, "HAS_XHCI_FSM     = 1") {
		t.Error("xHCI HAS_XHCI_FSM should be 1")
	}
	if strings.Contains(result, "HAS_NVME_FSM     = 1") {
		t.Error("xHCI should NOT have HAS_NVME_FSM = 1")
	}
}

func TestGenerateBarControllerSV_XHCI(t *testing.T) {
	cfg := xhciConfig()
	cfg.LatencyConfig = DefaultLatencyConfig(cfg.ClassCode)
	_, err := GenerateBarControllerSV(cfg)
	if err != nil {
		t.Fatalf("xHCI bar_controller failed: %v", err)
	}
}

func TestGenerateBarControllerSV_Audio(t *testing.T) {
	cfg := audioConfig()
	cfg.LatencyConfig = DefaultLatencyConfig(cfg.ClassCode)
	_, err := GenerateBarControllerSV(cfg)
	if err != nil {
		t.Fatalf("Audio bar_controller failed: %v", err)
	}
}

func TestDeviceClassEmptyString(t *testing.T) {
	// empty DeviceClass should not enable any FSM
	cfg := testConfig()
	cfg.DeviceClass = ""
	cfg.BARModel = &barmodel.BARModel{
		Size: 4096,
		Registers: []barmodel.BARRegister{
			{Offset: 0x00, Width: 4, Name: "REG0", RWMask: 0x00000000, Reset: 0x00},
		},
	}
	result, err := GenerateBarImplDeviceSV(cfg)
	if err != nil {
		t.Fatalf("empty DeviceClass failed: %v", err)
	}
	if strings.Contains(result, "cc_en_prev") {
		t.Error("empty DeviceClass should not have NVMe FSM")
	}
	if strings.Contains(result, "hcrst_cnt") {
		t.Error("empty DeviceClass should not have xHCI FSM")
	}
	if strings.Contains(result, "crst_prev") {
		t.Error("empty DeviceClass should not have Audio FSM")
	}
}
