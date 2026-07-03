package svgen

import (
	"strings"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/firmware/nvme"
)

// TestGenerateBarControllerSV_NVMeLatencyBypassUnconditional guards the Code-10
// fix: NVMe BAR0 reads must bypass the latency emulator regardless of MSI-X,
// or a failed MSI-X capture routes CSTS reads through the emulator (stale
// CSTS=0 -> stornvme fails ControllerEnable -> Code 10).
func TestGenerateBarControllerSV_NVMeLatencyBypassUnconditional(t *testing.T) {
	for _, msix := range []bool{true, false} {
		cfg := testConfig()
		cfg.LatencyConfig = DefaultLatencyConfig(cfg.ClassCode)
		cfg.NVMeIdentify = &nvme.IdentifyData{}
		if msix {
			cfg.HasMSIX = true
			cfg.MSIXConfig = &MSIXConfig{NumVectors: 9, TableOffset: 0x2000, PBAOffset: 0x2100}
		}
		result, err := GenerateBarControllerSV(cfg)
		if err != nil {
			t.Fatalf("msix=%v: %v", msix, err)
		}
		if !strings.Contains(result, "assign bar0_base_ctx   = bar0_raw_ctx;") {
			t.Fatalf("msix=%v: NVMe BAR0 latency bypass must be present", msix)
		}
		if strings.Contains(result, "bar_rsp_valid[0]   = bar0_emu_valid") {
			t.Fatalf("msix=%v: NVMe bar_rsp[0] must not route through the latency emulator", msix)
		}
	}
}

// TestGenerateBarControllerSV_WiresNVMeDoorbellsAndDisk verifies the bar
// controller folds the four doorbell strobes into the responder's generalized
// doorbell interface and instantiates the BRAM disk cache on disk_req_*.
func TestGenerateBarControllerSV_WiresNVMeDoorbellsAndDisk(t *testing.T) {
	cfg := testConfig()
	cfg.NVMeIdentify = &nvme.IdentifyData{}
	cfg.NVMeDoorbellStride = 0
	cfg.NVMeDiskWords = 8192

	result, err := GenerateBarControllerSV(cfg)
	if err != nil {
		t.Fatalf("GenerateBarControllerSV failed: %v", err)
	}

	for _, want := range []string{
		// range-based doorbell decode
		"nvme_db_base",
		"nvme_db_wr",
		"nvme_db_is_cq",
		"nvme_db_qid",
		// CC.EN explicit pulses
		"nvme_cc_enable_wr",
		".cc_enable_wr",
		// generalized doorbell interface to responder
		".doorbell_wr",
		".doorbell_is_cq",
		".doorbell_qid",
		// BRAM disk cache
		"i_nvme_bram_disk",
		".disk_req_valid",
		".disk_req_lba",
	} {
		if !strings.Contains(result, want) {
			t.Fatalf("NVMe bar controller should contain %q", want)
		}
	}

	// Regression for Synth 8-7136: bram_disk dropped DISK_LBAS.
	for _, mustNot := range []string{
		".DISK_LBAS",
		"DISK_LBAS",
	} {
		if strings.Contains(result, mustNot) {
			t.Fatalf("NVMe bar controller should NOT pass removed bram_disk param %q", mustNot)
		}
	}

	// Regression for Synth 8-11241/8-8895: wrengine_ready must be declared before use.
	declIdx := strings.Index(result, "wire        wrengine_ready;")
	instIdx := strings.Index(result, "i_pcileech_tlps128_bar_wrengine")
	if declIdx < 0 {
		t.Fatal("NVMe bar controller should declare wrengine_ready")
	}
	if instIdx < 0 || !(declIdx < instIdx) {
		t.Fatalf("wrengine_ready must be declared before the wrengine instance (decl=%d inst=%d)", declIdx, instIdx)
	}

	// Regression for Synth 8-11241: nvme_id_rom_data must be pre-declared before use.
	if !strings.Contains(result, "wire [31:0] nvme_id_rom_data;") {
		t.Fatal("NVMe bar controller should pre-declare wire [31:0] nvme_id_rom_data before its use")
	}

	// Regression: cfg.NVMeDiskWords must reach the SV define (board-scaled cache).
	// Was a hardcoded 32768 default, blowing the 75T BRAM budget (127 > 105 RAMB36).
	if !strings.Contains(result, "`define NVME_DISK_WORDS 8192") {
		t.Fatal("NVMe bar controller should emit `define NVME_DISK_WORDS from cfg.NVMeDiskWords")
	}
}
