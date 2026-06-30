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
}
