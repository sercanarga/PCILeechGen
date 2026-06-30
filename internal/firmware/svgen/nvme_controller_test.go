package svgen

import (
	"strings"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/firmware/nvme"
)

// TestGenerateBarControllerSV_WiresNVMeDoorbellsAndDisk verifies the bar
// controller folds the four decoded doorbell strobes into the responder's
// generalized doorbell interface and instantiates the BRAM disk cache wired to
// the responder's disk_req_* path with the generator-configured pinned windows.
func TestGenerateBarControllerSV_WiresNVMeDoorbellsAndDisk(t *testing.T) {
	cfg := testConfig()
	cfg.NVMeIdentify = &nvme.IdentifyData{}
	cfg.NVMeDoorbellStride = 0

	result, err := GenerateBarControllerSV(cfg)
	if err != nil {
		t.Fatalf("GenerateBarControllerSV failed: %v", err)
	}

	// The four per-queue doorbell address decoders are still emitted...
	for _, want := range []string{
		"wire nvme_sq0_db_wr",
		"wire nvme_cq0_db_wr",
		"wire nvme_sq1_db_wr",
		"wire nvme_cq1_db_wr",
		// ...folded into the generalized doorbell interface:
		"wire        nvme_db_wr",
		"wire        nvme_db_is_cq",
		"wire [15:0] nvme_db_qid",
		".doorbell_wr",
		".doorbell_is_cq",
		".doorbell_qid",
		// ...and the BRAM disk cache is instantiated + wired to the responder.
		"pcileech_bram_disk i_nvme_bram_disk",
		".disk_req_valid",
		".disk_req_lba",
		".req_done       ( nvme_disk_req_done",
	} {
		if !strings.Contains(result, want) {
			t.Fatalf("NVMe bar controller should contain %q", want)
		}
	}
}
