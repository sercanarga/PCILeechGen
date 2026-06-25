package svgen

import (
	"regexp"
	"strings"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/donor/behavior"
	"github.com/sercanarga/pcileechgen/internal/firmware"
	"github.com/sercanarga/pcileechgen/internal/firmware/barmodel"
	"github.com/sercanarga/pcileechgen/internal/util"
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
	if !strings.Contains(result, "HAS_INTX_INT     = 0") {
		t.Error("NVMe (no INTxConfig) should set HAS_INTX_INT = 0")
	}
}

func TestGenerateINTxSV(t *testing.T) {
	cfg := testConfig()
	cfg.INTxConfig = &INTxConfig{Line: 0} // INTA
	out, err := GenerateINTxSV(cfg)
	if err != nil {
		t.Fatalf("GenerateINTxSV failed: %v", err)
	}
	for _, want := range []string{
		"module pcileech_intx_gen",
		"INTX_LINE",
		"cfg_interrupt",
		"cfg_interrupt_assert",
		"cfg_pciecap_interrupt_msgnum",
		"ST_ASSERT",
		"ST_DEASSERT",
		"pending",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("intx_gen output missing %q", want)
		}
	}
	if !strings.Contains(out, "if (intr_req && state != ST_IDLE)") {
		t.Error("intx_gen should latch re-triggers during assert/deassert")
	}
}

func TestGenerateDeviceConfigSV_INTx(t *testing.T) {
	cfg := testConfig()
	cfg.INTxConfig = &INTxConfig{Line: 2} // INTC
	result, err := GenerateDeviceConfigSV(cfg)
	if err != nil {
		t.Fatalf("GenerateDeviceConfigSV failed: %v", err)
	}
	if !strings.Contains(result, "HAS_INTX_INT     = 1") {
		t.Error("with INTxConfig, HAS_INTX_INT should be 1")
	}
	if !strings.Contains(result, "INTX_LINE        = 3'd2") {
		t.Error("INTX_LINE should reflect the configured line (INTC=2)")
	}
}

func TestGenerateBarImplDeviceSV_NilBARModel(t *testing.T) {
	cfg := testConfig()
	cfg.BARModel = nil
	result, err := GenerateBarImplDeviceSV(cfg)
	if err != nil {
		t.Fatalf("GenerateBarImplDeviceSV failed: %v", err)
	}
	if !strings.Contains(result, "bar_mem") {
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
	if !strings.Contains(result, "csts_state") {
		t.Error("NVMe output should contain CC->CSTS state machine")
	}
	if !strings.Contains(result, "CST_ENABLING") {
		t.Error("NVMe output should contain timed enable FSM")
	}
	if !strings.Contains(result, "ENABLING_TICKS") {
		t.Error("NVMe output should contain timed enable ticks")
	}
}

func TestGenerateBarImplDeviceSV_DeclaresEveryRegisterReference(t *testing.T) {
	cases := []struct {
		name string
		cfg  *SVGeneratorConfig
	}{
		{name: "nvme", cfg: func() *SVGeneratorConfig {
			cfg := testConfig()
			cfg.BARModel = barmodel.BuildBARModel(nil, 0x010802, nil)
			return cfg
		}()},
		{name: "xhci", cfg: func() *SVGeneratorConfig {
			cfg := xhciConfig()
			cfg.BARModel = barmodel.BuildBARModel(nil, 0x0C0330, nil)
			return cfg
		}()},
		{name: "audio", cfg: func() *SVGeneratorConfig {
			cfg := audioConfig()
			cfg.BARModel = barmodel.BuildBARModel(nil, 0x040300, nil)
			return cfg
		}()},
		{name: "ethernet", cfg: func() *SVGeneratorConfig {
			cfg := ethernetConfig()
			cfg.BARModel = barmodel.BuildBARModel(nil, 0x020000, nil)
			return cfg
		}()},
		{name: "cardreader", cfg: func() *SVGeneratorConfig {
			cfg := cardReaderConfig()
			cfg.BARModel = barmodel.BuildBARModelForDeviceIDWithOverlay(nil, 0xFF0000, 0x10EC, 0x522A, nil, nil)
			return cfg
		}()},
	}

	declRE := regexp.MustCompile(`(?m)^\\s*reg\\s+\\[[^\\]]+\\]\\s+(reg_0x[0-9A-F]{8})\\b`)
	refRE := regexp.MustCompile(`\\breg_0x[0-9A-F]{8}\\b`)

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			out, err := GenerateBarImplDeviceSV(tc.cfg)
			if err != nil {
				t.Fatal(err)
			}

			declared := map[string]bool{}
			for _, match := range declRE.FindAllStringSubmatch(out, -1) {
				declared[match[1]] = true
			}
			for _, ref := range refRE.FindAllString(out, -1) {
				if !declared[ref] {
					t.Fatalf("%s referenced but not declared", ref)
				}
			}
		})
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
	if !strings.Contains(result, "intr_req_pending") {
		t.Error("bar controller should latch interrupt pulses during startup gate")
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
				{Offset: 0x20, Width: 4, Name: "USBCMD", RWMask: 0x00002F0E, Reset: 0x00000001, IsFSMDriven: true},
				{Offset: 0x24, Width: 4, Name: "USBSTS", RWMask: 0x0000041C, Reset: 0x00000000, IsRW1C: true, IsFSMDriven: true},
				{Offset: 0x220, Width: 4, Name: "IMAN0", RWMask: 0x00000003, Reset: 0x00000000, IsRW1C: true, IsFSMDriven: true},
				{Offset: 0x224, Width: 4, Name: "IMOD0", RWMask: 0xFFFFFFFF, Reset: 0x00000000},
				{Offset: 0x228, Width: 4, Name: "ERSTSZ0", RWMask: 0x0000FFFF, Reset: 0x00000000},
				{Offset: 0x230, Width: 4, Name: "ERSTBA_LO", RWMask: 0xFFFFFFC0, Reset: 0x00000000},
				{Offset: 0x234, Width: 4, Name: "ERSTBA_HI", RWMask: 0xFFFFFFFF, Reset: 0x00000000},
				{Offset: 0x238, Width: 4, Name: "ERDP_LO", RWMask: 0xFFFFFFF8, Reset: 0x00000000},
				{Offset: 0x23C, Width: 4, Name: "ERDP_HI", RWMask: 0xFFFFFFFF, Reset: 0x00000000},
				{Offset: 0x420, Width: 4, Name: "PORTSC1", RWMask: 0x8EFFC3F2, Reset: 0x000002A0, IsFSMDriven: true},
				{Offset: 0x430, Width: 4, Name: "PORTSC2", RWMask: 0x8EFFC3F2, Reset: 0x000002A0, IsFSMDriven: true},
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
			VendorID:       0x10EC,
			DeviceID:       0x8125,
			SubsysVendorID: 0x10EC,
			SubsysDeviceID: 0x8125,
			RevisionID:     0x04,
			ClassCode:      0x020000,
			HasPCIeCap:     true,
			LinkSpeed:      3,
			LinkWidth:      1,
		},
		ClassCode:    0x020000,
		DeviceClass:  "ethernet",
		BuildEntropy: 0xFEEDFACE,
		PRNGSeeds:    BuildPRNGSeeds(0x10EC, 0x8125, 0xFEEDFACE),
		BARModel: &barmodel.BARModel{
			Size: 4096,
			Registers: []barmodel.BARRegister{
				{Offset: 0x00, Width: 4, Name: "MAC0_3", RWMask: 0xFFFFFFFF, Reset: 0xBEADDE02},
				{Offset: 0x6C, Width: 4, Name: "PHYSTATUS", RWMask: 0x00000000, Reset: 0x00003010},
			},
		},
	}
}

func cardReaderConfig() *SVGeneratorConfig {
	return &SVGeneratorConfig{
		DeviceIDs: firmware.DeviceIDs{
			VendorID:       0x10EC,
			DeviceID:       0x522A,
			SubsysVendorID: 0x103C,
			SubsysDeviceID: 0x838F,
			RevisionID:     0x01,
			ClassCode:      0xFF0000,
			HasPCIeCap:     true,
			LinkSpeed:      3,
			LinkWidth:      1,
		},
		ClassCode:    0xFF0000,
		DeviceClass:  "cardreader",
		BuildEntropy: 0xA522A522,
		PRNGSeeds:    BuildPRNGSeeds(0x10EC, 0x522A, 0xA522A522),
		BARModel: &barmodel.BARModel{
			Size: 4096,
			Registers: []barmodel.BARRegister{
				{Offset: 0x10, Width: 4, Name: "RTSX_HAIMR", RWMask: 0xFFFFFFFF, Reset: 0x00000000, IsFSMDriven: true},
				{Offset: 0x14, Width: 4, Name: "RTSX_BIPR", RWMask: 0x00000000, Reset: 0x00000000, IsFSMDriven: true},
				{Offset: 0x18, Width: 4, Name: "RTSX_BIER", RWMask: 0xFFC00000, Reset: 0x00000000},
			},
		},
	}
}

func TestGenerateBarImplDeviceSV_CardReader(t *testing.T) {
	cfg := cardReaderConfig()
	result, err := GenerateBarImplDeviceSV(cfg)
	if err != nil {
		t.Fatalf("card reader bar_impl_device failed: %v", err)
	}
	if !strings.Contains(result, "cardreader_pm_ctrl3") {
		t.Error("card reader output should declare internal PM_CTRL3 state")
	}
	if !strings.Contains(result, "14'h3F7E") || !strings.Contains(result, "14'h3F78") {
		t.Error("card reader output should decode RTS522A indirect register addresses")
	}
	if !strings.Contains(result, "reg_0x00000014[31] <= 1'b1") {
		t.Error("card reader HAIMR completion should raise CMD_DONE in RTSX_BIPR")
	}
	if !strings.Contains(result, "32'h00000014") {
		t.Error("card reader output should handle RTSX_BIPR host clears")
	}
	if !strings.Contains(result, "reg_0x00000014           <= 32'h00010000;") {
		t.Error("card reader reset should expose SD_EXIST in RTSX_BIPR")
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
	// USBSTS status bits must be W1C.
	if !strings.Contains(result, "USBSTS W1C") {
		t.Error("xHCI output should service USBSTS W1C clears")
	}

	if !strings.Contains(result, "32'h00000420") || !strings.Contains(result, "32'h00000430") {
		t.Error("xHCI output should contain PORTSC write handling for both ports")
	}
	if !strings.Contains(result, "reg_0x00000420 <= 32'h000002A0;") ||
		!strings.Contains(result, "reg_0x00000430 <= 32'h000002A0;") {
		t.Error("xHCI output should initialize FSM-owned PORTSC registers")
	}
	if !strings.Contains(result, "reg_0x00000420[23:16] & ~8'h01") ||
		!strings.Contains(result, "reg_0x00000430[23:16] & ~8'h01") {
		t.Error("xHCI PORTSC byte-2 writes should preserve W1C change bits")
	}
	if !strings.Contains(result, "reg_0x00000024[4]") {
		t.Error("xHCI port reset should raise USBSTS.PCD")
	}
	if !strings.Contains(result, "reg_0x00000420[21]") {
		t.Error("xHCI port reset should set PORTSC1 Port Reset Change")
	}
	if !strings.Contains(result, "32'h00000220") {
		t.Error("xHCI output should contain IMAN0 handling")
	}
	if !strings.Contains(result, "reg_0x00000220[0]") {
		t.Error("xHCI port reset should raise IMAN0 interrupt pending")
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
		t.Error("Ethernet output should contain MAC0_3 register")
	}
	// Ethernet should NOT have NVMe FSM
	if strings.Contains(result, "cc_en_prev") {
		t.Error("Ethernet should NOT contain NVMe FSM")
	}
}

func TestGenerateBarImplDeviceSV_SeqRead(t *testing.T) {
	cfg := testConfig()
	cfg.DeviceClass = "other"
	cfg.BARModel = &barmodel.BARModel{
		Size: 4096,
		Registers: []barmodel.BARRegister{
			{Offset: 0x00, Width: 4, Name: "ID", RWMask: 0x00000000, Reset: 0x12345678},
			{
				Offset:           0x10,
				Width:            4,
				Name:             "EEPROM_REQ",
				RWMask:           0x00000000,
				Reset:            0xDEADBEEF,
				SequentialRead:   true,
				SequentialValues: []uint32{0xDEADBEEF, 0x00000002, 0x00000001},
			},
		},
	}
	result, err := GenerateBarImplDeviceSV(cfg)
	if err != nil {
		t.Fatalf("SeqRead bar_impl_device failed: %v", err)
	}
	if !strings.Contains(result, "seqrom_0x00000010 [0:2]") {
		t.Error("output should declare a sequential-read ROM sized to the configured values")
	}
	if !strings.Contains(result, "initial seqrom_0x00000010[0] = 32'hDEADBEEF;") {
		t.Error("output should emit the first sequential-read ROM value")
	}
	if !strings.Contains(result, "initial seqrom_0x00000010[2] = 32'h00000001;") {
		t.Error("output should emit the last sequential-read ROM value")
	}
	if !strings.Contains(result, "rd_data_d1 <= seqrom_0x00000010[seqcnt_0x00000010];") {
		t.Error("output should read from the sequential-read ROM")
	}
}

func TestGenerateBarImplDeviceSV_StaticShadow(t *testing.T) {
	cfg := testConfig()
	cfg.DeviceClass = "other"
	cfg.BARModel = &barmodel.BARModel{
		Size: 4096,
		Registers: []barmodel.BARRegister{
			{Offset: 0x00, Width: 4, Name: "ID", RWMask: 0x00000000, Reset: 0x12345678},
		},
		StaticShadow: []barmodel.StaticWord{
			{Offset: 0x100, Value: 0xCAFEBABE},
			{Offset: 0x200, Value: 0x0000BEEF},
		},
	}
	result, err := GenerateBarImplDeviceSV(cfg)
	if err != nil {
		t.Fatalf("StaticShadow bar_impl_device failed: %v", err)
	}
	if !strings.Contains(result, "32'h00000100: rd_data_d1 <= 32'hCAFEBABE") {
		t.Error("static shadow should emit a read case for offset 0x100")
	}
	if !strings.Contains(result, "32'h00000200: rd_data_d1 <= 32'h0000BEEF") {
		t.Error("static shadow should emit a read case for offset 0x200")
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

func TestGenerateMSIXTableSV(t *testing.T) {
	cfg := testConfig()
	cfg.HasMSIX = true
	cfg.MSIXConfig = &MSIXConfig{
		NumVectors:  8,
		TableOffset: 0x3000,
		PBAOffset:   0x3100,
	}
	result, err := GenerateMSIXTableSV(cfg)
	if err != nil {
		t.Fatalf("GenerateMSIXTableSV failed: %v", err)
	}
	if !strings.Contains(result, "msix") && !strings.Contains(result, "MSIX") {
		t.Error("output should contain MSI-X references")
	}
}

func TestNVMeDoorbellOffsets(t *testing.T) {
	cfg := &SVGeneratorConfig{NVMeDoorbellStride: 0}
	sq0 := cfg.NVMeSQ0DoorbellOffset()
	cq0 := cfg.NVMeCQ0DoorbellOffset()

	d := uint32(0x1000)
	if sq0 != d {
		t.Errorf("SQ0 doorbell = 0x%X, want 0x1000", sq0)
	}
	if cq0 != 0x1004 {
		t.Errorf("CQ0 doorbell = 0x%X, want 0x1004", cq0)
	}

	// with stride=1 doorbells are 8 bytes apart
	cfg.NVMeDoorbellStride = 1
	sq0 = cfg.NVMeSQ0DoorbellOffset()
	cq0 = cfg.NVMeCQ0DoorbellOffset()
	d = uint32(0x1000)
	if sq0 != d {
		t.Errorf("SQ0 doorbell (stride=1) = 0x%X, want 0x1000", sq0)
	}
	if cq0 != 0x1008 {
		t.Errorf("CQ0 doorbell (stride=1) = 0x%X, want 0x1008", cq0)
	}
}

func TestGenerateNVMeResponderSV(t *testing.T) {
	cfg := testConfig()
	cfg.NVMeDoorbellStride = 0
	result, err := GenerateNVMeResponderSV(cfg)
	if err != nil {
		t.Fatalf("GenerateNVMeResponderSV failed: %v", err)
	}
	if len(result) == 0 {
		t.Error("NVMe responder SV should not be empty")
	}
}

func TestGenerateNVMeDMABridgeSV(t *testing.T) {
	cfg := testConfig()
	cfg.NVMeDoorbellStride = 0
	result, err := GenerateNVMeDMABridgeSV(cfg)
	if err != nil {
		t.Fatalf("GenerateNVMeDMABridgeSV failed: %v", err)
	}
	if len(result) == 0 {
		t.Error("NVMe DMA bridge SV should not be empty")
	}
	// Fixed read tags 0xE0-0xE3 instead of the old 0x80 increment.
	if !strings.Contains(result, "8'hE0") {
		t.Error("DMA bridge should use fixed read tags starting at 0xE0")
	}
	if !strings.Contains(result, "rd_tag_sent") {
		t.Error("DMA bridge should latch the sent tag for CplD matching")
	}
	// Backpressure hold FSM.
	if !strings.Contains(result, "tx_busy") {
		t.Error("DMA bridge should include backpressure hold (tx_busy)")
	}
	// The fragile "rd_tag - 8'h01" matcher must be gone.
	if strings.Contains(result, "rd_tag - 8'h01") {
		t.Error("DMA bridge must not use the fragile rd_tag-1 CplD matcher")
	}
}

func TestGenerateNVMeResponderSV_NumQueuesAndStatus(t *testing.T) {
	cfg := testConfig()
	cfg.NVMeDoorbellStride = 0
	cfg.NVMeNumIOQueues = 0x00070007
	result, err := GenerateNVMeResponderSV(cfg)
	if err != nil {
		t.Fatalf("GenerateNVMeResponderSV failed: %v", err)
	}
	if !strings.Contains(result, "NUM_IO_QUEUES") {
		t.Error("responder should expose NUM_IO_QUEUES parameter")
	}
	if !strings.Contains(result, "SC_INVALID_OPCODE") {
		t.Error("responder should define NVMe status localparams")
	}
	if !strings.Contains(result, "SC_INVALID_LOGPAGE") {
		t.Error("responder should dispatch GetLogPage by LID with InvalidLogPage fallback")
	}
	if !strings.Contains(result, "cqe[3] <= {(SC_INVALID_LOGPAGE | {15'h0, cq_phase}), cmd_cid};") {
		t.Error("responder should pack NVMe status into DW3 upper half and CID into lower half")
	}
}

func TestLatencyConfigFromHistogram(t *testing.T) {
	h := &behavior.TimingHistogram{
		SampleCount:  1000,
		MinCycles:    2,
		MaxCycles:    25,
		MedianCycles: 8,
		Buckets:      [16]uint8{10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 110, 120, 130, 140, 150, 160},
		CDF:          [16]uint8{16, 32, 48, 64, 80, 96, 112, 128, 144, 160, 176, 192, 208, 224, 240, 255},
	}
	cfg := LatencyConfigFromHistogram(h, 0x010802)

	if cfg.MinCycles != 2 {
		t.Errorf("MinCycles = %d, want 2", cfg.MinCycles)
	}
	if cfg.MaxCycles != 25 {
		t.Errorf("MaxCycles = %d, want 25", cfg.MaxCycles)
	}
	if !cfg.HasHistogram {
		t.Error("HasHistogram should be true")
	}
}

func TestLatencyConfigFromHistogram_Nil(t *testing.T) {
	cfg := LatencyConfigFromHistogram(nil, 0x010802)
	if cfg.MinCycles != 3 {
		t.Errorf("nil histogram should use NVMe defaults, got MinCycles=%d", cfg.MinCycles)
	}
}

func TestLatencyConfigFromHistogram_ZeroSamples(t *testing.T) {
	h := &behavior.TimingHistogram{SampleCount: 0}
	cfg := LatencyConfigFromHistogram(h, 0x020000)
	if cfg.MinCycles != 2 {
		t.Errorf("zero samples should use Ethernet defaults, got MinCycles=%d", cfg.MinCycles)
	}
}

func TestDefaultLatencyConfig_AllClasses(t *testing.T) {
	classes := []struct {
		code uint32
		min  int
	}{
		{0x010802, 3}, // NVMe
		{0x0C0330, 4}, // xHCI
		{0x020000, 2}, // Ethernet
		{0x030000, 5}, // GPU
		{0x010600, 3}, // SATA
		{0x040300, 2}, // HD Audio
		{0x028000, 3}, // Wi-Fi
		{0x0C8000, 4}, // Thunderbolt
		{0xFF0000, 3}, // Unknown
	}
	for _, tc := range classes {
		cfg := DefaultLatencyConfig(tc.code)
		if cfg.MinCycles != tc.min {
			t.Errorf("class 0x%06X: MinCycles=%d, want %d", tc.code, cfg.MinCycles, tc.min)
		}
		if cfg.CDF[15] != 255 {
			t.Errorf("class 0x%06X: CDF[15]=%d, want 255", tc.code, cfg.CDF[15])
		}
	}
}

func TestBuildEntropyFromTime(t *testing.T) {
	e1 := BuildEntropyFromTime()
	e2 := BuildEntropyFromTime()

	if e1 == 0 && e2 == 0 {
		t.Error("BuildEntropyFromTime should produce non-zero values")
	}
}

func TestGenerateBarImplDeviceSV_AR9287_WiredFromProfile(t *testing.T) {
	// End-to-end: donor barData + Atheros VID drives BuildBARModelForDevice to
	// mark the EEPROM handshake regs SequentialRead, and the generated SV
	// contains the advancing-read counter for both offsets.
	barData := make([]byte, 0x4080)
	util.WriteLE32(barData, 0x4010, 0xDEADBEEF)
	util.WriteLE32(barData, 0x407C, 0x0000BEEF)

	cfg := testConfig()
	cfg.DeviceClass = "wifi"
	cfg.BARModel = barmodel.BuildBARModelForDeviceIDWithOverlay(barData, 0x028000, 0x168C, 0x002E, nil, nil)
	if cfg.BARModel == nil {
		t.Fatal("BuildBARModelForDeviceIDWithOverlay returned nil for AR9287")
	}
	out, err := GenerateBarImplDeviceSV(cfg)
	if err != nil {
		t.Fatalf("generate failed: %v", err)
	}
	for _, want := range []string{"seqcnt_0x00004010", "seqcnt_0x0000407C"} {
		if !strings.Contains(out, want) {
			t.Errorf("output missing %s (EEPROM SequentialRead counter not wired)", want)
		}
	}
}
