package svgen

import (
	"strings"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/firmware/barmodel"
	"github.com/sercanarga/pcileechgen/internal/firmware/nvme"
)

func TestGenerateBarControllerSV_WiresNVMeQID1Doorbells(t *testing.T) {
	cfg := testConfig()
	cfg.NVMeIdentify = &nvme.IdentifyData{}
	cfg.BARModel = &barmodel.BARModel{Size: 0x4000}
	cfg.NVMeDoorbellStride = 0

	result, err := GenerateBarControllerSV(cfg)
	if err != nil {
		t.Fatalf("GenerateBarControllerSV failed: %v", err)
	}

	for _, want := range []string{
		"wire nvme_sq1_db_wr",
		"wire nvme_cq1_db_wr",
		".sq1_doorbell_wr",
		".cq1_doorbell_wr",
	} {
		if !strings.Contains(result, want) {
			t.Fatalf("NVMe bar controller should contain %q", want)
		}
	}
}

func TestGenerateBarControllerSV_ExposesNVMeDiagnostics(t *testing.T) {
	cfg := testConfig()
	cfg.NVMeIdentify = &nvme.IdentifyData{}
	cfg.BARModel = &barmodel.BARModel{Size: 0x4000}
	cfg.NVMeDoorbellStride = 0

	result, err := GenerateBarControllerSV(cfg)
	if err != nil {
		t.Fatalf("GenerateBarControllerSV failed: %v", err)
	}

	for _, want := range []string{
		"nvme_diag_addr_hit",
		"nvme_diag_data",
		"32'h00001010",
		"32'h00001014",
		"nvme_debug_status",
		"nvme_debug_last_cmd",
		"nvme_debug_queue_state",
		".debug_last_cdw10",
	} {
		if !strings.Contains(result, want) {
			t.Fatalf("NVMe bar controller should contain diagnostic wiring %q", want)
		}
	}
}
