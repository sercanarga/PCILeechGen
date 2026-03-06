package output

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/firmware"
	"github.com/sercanarga/pcileechgen/internal/firmware/barmodel"
	"github.com/sercanarga/pcileechgen/internal/firmware/codegen"
	"github.com/sercanarga/pcileechgen/internal/firmware/devclass"
	"github.com/sercanarga/pcileechgen/internal/firmware/scrub"
	"github.com/sercanarga/pcileechgen/internal/firmware/svgen"
	"github.com/sercanarga/pcileechgen/internal/firmware/variance"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

// fakeConfigSpace builds a minimal 256-byte config space with the given IDs.
func fakeConfigSpace(vid, did uint16, classCode uint32) *pci.ConfigSpace {
	data := make([]byte, 256)
	data[0] = byte(vid)
	data[1] = byte(vid >> 8)
	data[2] = byte(did)
	data[3] = byte(did >> 8)
	data[0x09] = byte(classCode)
	data[0x0A] = byte(classCode >> 8)
	data[0x0B] = byte(classCode >> 16)
	// command: IO+Memory+BusMaster
	data[0x04] = 0x07
	return pci.NewConfigSpaceFromBytes(data)
}

func TestIntegration_ScrubVariancePipeline(t *testing.T) {
	cs := fakeConfigSpace(0x144D, 0xA808, 0x010802)
	scrubbed, _ := scrub.ScrubConfigSpaceWithOverlay(cs, nil)
	if scrubbed == nil {
		t.Fatal("scrub returned nil")
	}

	entropy := uint32(0xDEADBEEF)
	seed := variance.BuildVarianceSeed(0x144D, 0xA808, entropy)
	variance.Apply(scrubbed, nil, variance.DefaultConfig(seed))

	coe := codegen.GenerateConfigSpaceCOE(scrubbed)
	if !strings.Contains(coe, "memory_initialization_radix") {
		t.Error("COE should contain radix declaration")
	}
	if !strings.Contains(coe, "memory_initialization_vector") {
		t.Error("COE should have init vector")
	}

	hex := codegen.GenerateConfigSpaceHex(scrubbed)
	if len(hex) == 0 {
		t.Error("HEX output should not be empty")
	}
}

func TestIntegration_SVTemplateRender_AllClasses(t *testing.T) {
	cases := []struct {
		name      string
		vid, did  uint16
		classCode uint32
		devClass  string
	}{
		{"NVMe", 0x144D, 0xA808, 0x010802, devclass.ClassNVMe},
		{"xHCI", 0x8086, 0xA36D, 0x0C0330, devclass.ClassXHCI},
		{"Ethernet", 0x8086, 0x15B7, 0x020000, devclass.ClassEthernet},
		{"Audio", 0x8086, 0xA171, 0x040300, devclass.ClassAudio},
		{"GPU", 0x10DE, 0x2204, 0x030000, devclass.ClassGPU},
		{"SATA", 0x8086, 0xA102, 0x010601, devclass.ClassSATA},
		{"WiFi", 0x8086, 0x2725, 0x028000, devclass.ClassWiFi},
		{"Thunderbolt", 0x8086, 0x15EF, 0x0C8000, devclass.ClassThunderbolt},
		{"Generic", 0x1234, 0x5678, 0xFF0000, devclass.ClassGeneric},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			latCfg := svgen.DefaultLatencyConfig(tc.classCode)
			if latCfg == nil {
				t.Fatal("nil latency config")
			}
			if latCfg.WrMinCycles == 0 {
				t.Error("WrMinCycles should not be 0")
			}
			if latCfg.CplTimeoutCycles == 0 {
				t.Error("CplTimeoutCycles should not be 0")
			}

			entropy := uint32(0x42)
			seeds := svgen.BuildPRNGSeeds(tc.vid, tc.did, entropy)

			cfg := &svgen.SVGeneratorConfig{
				DeviceIDs: firmware.DeviceIDs{
					VendorID:       tc.vid,
					DeviceID:       tc.did,
					SubsysVendorID: tc.vid,
					SubsysDeviceID: tc.did,
					ClassCode:      tc.classCode,
					HasPCIeCap:     true,
					LinkSpeed:      3,
					LinkWidth:      1,
				},
				ClassCode:     tc.classCode,
				DeviceClass:   tc.devClass,
				BuildEntropy:  entropy,
				PRNGSeeds:     seeds,
				LatencyConfig: latCfg,
				BARModel: &barmodel.BARModel{
					Size: 4096,
					Registers: []barmodel.BARRegister{
						{Offset: 0x00, Width: 4, Name: "REG0", RWMask: 0, Reset: 0},
					},
				},
			}

			sv, err := svgen.GenerateBarImplDeviceSV(cfg)
			if err != nil {
				t.Fatalf("bar_impl_device render failed: %v", err)
			}
			if len(sv) == 0 {
				t.Error("bar_impl_device output empty")
			}

			ctrl, err := svgen.GenerateBarControllerSV(cfg)
			if err != nil {
				t.Fatalf("bar_controller render failed: %v", err)
			}
			if !strings.Contains(ctrl, "tlp_latency_emulator") {
				t.Error("bar_controller should contain latency emulator")
			}
			if !strings.Contains(ctrl, "wr_ack") {
				t.Error("bar_controller should contain wr_ack")
			}

			emu, err := svgen.GenerateLatencyEmulatorSV(cfg)
			if err != nil {
				t.Fatalf("latency_emulator render failed: %v", err)
			}
			if !strings.Contains(emu, "WR_MIN_LATENCY") {
				t.Error("latency emulator should have WR_MIN_LATENCY")
			}
			if !strings.Contains(emu, "CPL_TIMEOUT_CYCLES") {
				t.Error("latency emulator should have CPL_TIMEOUT_CYCLES")
			}
			if !strings.Contains(emu, "wr_ack") {
				t.Error("latency emulator should have wr_ack port")
			}

			dcfg, err := svgen.GenerateDeviceConfigSV(cfg)
			if err != nil {
				t.Fatalf("device_config render failed: %v", err)
			}
			if len(dcfg) == 0 {
				t.Error("device_config output empty")
			}
		})
	}
}

func TestIntegration_LatencyConfigWriteFields(t *testing.T) {
	classes := []struct {
		name string
		code uint32
	}{
		{"NVMe", 0x010802},
		{"xHCI", 0x0C0330},
		{"Ethernet", 0x020000},
		{"GPU", 0x030000},
		{"SATA", 0x010601},
		{"Audio", 0x040300},
		{"WiFi", 0x028000},
		{"Thunderbolt", 0x0C8000},
		{"Generic", 0xFF0000},
	}
	for _, tc := range classes {
		t.Run(tc.name, func(t *testing.T) {
			cfg := svgen.DefaultLatencyConfig(tc.code)
			if cfg.WrMinCycles >= cfg.WrMaxCycles {
				t.Errorf("WrMin (%d) >= WrMax (%d)", cfg.WrMinCycles, cfg.WrMaxCycles)
			}
			if cfg.MinCycles >= cfg.MaxCycles {
				t.Errorf("Min (%d) >= Max (%d)", cfg.MinCycles, cfg.MaxCycles)
			}
			if cfg.CplTimeoutCycles < 32768 {
				t.Errorf("CplTimeout too low: %d", cfg.CplTimeoutCycles)
			}
		})
	}
}

func TestIntegration_OutputManifest(t *testing.T) {
	dir := t.TempDir()
	testFile := filepath.Join(dir, "pcileech_cfgspace.coe")
	os.WriteFile(testFile, []byte("memory_initialization_radix=16;\nmemory_initialization_vector=00;"), 0644)

	manifest, err := GenerateManifest(dir, "test-version", "", 0x144D, 0xA808)
	if err != nil {
		t.Fatalf("manifest generation failed: %v", err)
	}
	if manifest == nil {
		t.Fatal("manifest should not be nil")
	}
}

func TestIntegration_OutputValidator(t *testing.T) {
	dir := t.TempDir()
	coe := "memory_initialization_radix=16;\nmemory_initialization_vector=\n00000000\n00000000\n00000000\n00000000;"
	os.WriteFile(filepath.Join(dir, "pcileech_cfgspace.coe"), []byte(coe), 0644)

	result := ValidateOutputDir(dir)
	// not all files exist so some will fail, but should not panic
	_ = result

	if err := ValidateCOEFile(coe); err != nil {
		t.Errorf("ValidateCOEFile failed: %v", err)
	}
}
