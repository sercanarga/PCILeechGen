package svgen

import (
	"strings"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/firmware/nvme"
)

// TestGenerateBarControllerSV_WiresNVMeDoorbellsAndDisk verifies the bar
// controller folds the four doorbell strobes into the responder's generalized
// doorbell interface and instantiates the BRAM disk cache on disk_req_*.
func TestGenerateBarControllerSV_WiresNVMeDoorbellsAndDisk(t *testing.T) {
	cfg := testConfig()
	cfg.NVMeIdentify = &nvme.IdentifyData{}
	cfg.NVMeDoorbellStride = 0

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
}
