package svgen

import (
	"strings"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/firmware/nvme"
)

func TestGenerateNVMeResponderSV_UsesPRP2ForPageCrossingAdminData(t *testing.T) {
	cfg := testConfig()

	result, err := GenerateNVMeResponderSV(cfg)
	if err != nil {
		t.Fatalf("GenerateNVMeResponderSV failed: %v", err)
	}

	if !strings.Contains(result, "function [63:0] prp_data_addr") {
		t.Fatal("NVMe responder should compute admin data addresses through a PRP helper")
	}
	if !strings.Contains(result, "cmd_prp2 + {32'h00000000, (byte_off - first_span)}") {
		t.Fatal("NVMe responder should use PRP2 for admin data crossing the first PRP page")
	}
	if strings.Contains(result, "dma_wr_addr  <= cmd_prp1 +") {
		t.Fatal("NVMe responder should not write all admin data relative to PRP1 only")
	}
}

func TestGenerateNVMeResponderSV_SMARTSeedsResetCounters(t *testing.T) {
	cfg := testConfig()

	result, err := GenerateNVMeResponderSV(cfg)
	if err != nil {
		t.Fatalf("GenerateNVMeResponderSV failed: %v", err)
	}

	for _, want := range []string{
		"power_on_hours          <= 32'h000004D2",         // 1234
		"power_cycle_count       <= 32'h00000237",         // 567
		"stat_unsafe_shutdowns   <= 32'h00000003",         // 3
		"stat_data_units_written <= 64'h00000000000F4240", // 1000000
	} {
		if !strings.Contains(result, want) {
			t.Errorf("reset block should seed from SMART; missing %q", want)
		}
	}
	for _, stale := range []string{
		"power_cycle_count       <= 32'h00000001",
		"stat_data_units_written <= 64'h0;",
		"power_on_hours          <= 32'h0;",
	} {
		if strings.Contains(result, stale) {
			t.Errorf("reset block should no longer use default %q", stale)
		}
	}
}

func TestGenerateNVMeResponderSV_NoHardcodedGigabyteStrings(t *testing.T) {
	cfg := testConfig()

	result, err := GenerateNVMeResponderSV(cfg)
	if err != nil {
		t.Fatalf("GenerateNVMeResponderSV failed: %v", err)
	}

	for _, bad := range []string{
		"32'h41474947", // "GIGA"
		"GIGABYTE",
	} {
		if strings.Contains(result, bad) {
			t.Fatalf("responder must not hardcode %q; donor ROM must supply SN/MN/FR", bad)
		}
	}

	if strings.Contains(result, "11'd1:  identify_data_word =") {
		t.Fatal("responder must not override SN word (dw_index 1); ROM must supply it")
	}
}

func TestGenerateNVMeResponderSV_ErrorLogPage(t *testing.T) {
	cfg := testConfig()

	result, err := GenerateNVMeResponderSV(cfg)
	if err != nil {
		t.Fatalf("GenerateNVMeResponderSV failed: %v", err)
	}

	for _, want := range []string{
		"8'd0: log_page_word = last_error_count[31:0]",
		"8'd2: log_page_word = {last_error_cid, last_error_sqid}",
		"8'd3: log_page_word = {last_error_peloc, last_error_status}",
		"8'd4: log_page_word = last_error_lba[31:0]",
		"last_error_lba <= (active_qid == 16'h0000) ? 64'h0000000000000000 : cmd_slba",
		"stat_error_log_entries <= stat_error_log_entries + 1'b1",
	} {
		if !strings.Contains(result, want) {
			t.Errorf("error log implementation missing %q", want)
		}
	}
}

func TestGenerateNVMeResponderSV_FirmwareSlotLogUsesIdentifyFirmware(t *testing.T) {
	cfg := testConfig()
	cfg.NVMeIdentify = nvme.BuildIdentifyData(cfg.DeviceIDs, nil, &nvme.ControllerIdentity{
		Serial: "TESTSERIAL0000000001",
		Model:  "Test NVMe Controller",
		FWRev:  "FW123456",
	})

	result, err := GenerateNVMeResponderSV(cfg)
	if err != nil {
		t.Fatalf("GenerateNVMeResponderSV failed: %v", err)
	}

	for _, want := range []string{
		"8'd2: log_page_word = 32'h32315746",
		"8'd3: log_page_word = 32'h36353433",
	} {
		if !strings.Contains(result, want) {
			t.Fatalf("firmware slot log should contain %q", want)
		}
	}
	if strings.Contains(result, "32'h31564E44") {
		t.Fatal("firmware slot log must not use a hardcoded revision")
	}
}

func TestGenerateNVMeResponderSV_VendorDiagnosticLog(t *testing.T) {
	cfg := testConfig()

	result, err := GenerateNVMeResponderSV(cfg)
	if err != nil {
		t.Fatalf("GenerateNVMeResponderSV failed: %v", err)
	}

	for _, want := range []string{
		"LOG_PAGE_VENDOR    = 8'hC0",
		"8'd48: log_page_word = 32'h00000001",
		"8'd0: log_page_word = 32'h444D564E",
		"8'd18: log_page_word = stat_flush_cmds[31:0]",
		"8'd45: log_page_word = {31'h0, feat_write_cache[0]}",
		"8'd8: log_page_word = stat_dma_mrd_tlps[31:0]",
		"8'd12: log_page_word = stat_prp_list_fetches[31:0]",
		"8'd32: log_page_word = stat_invalid_cmds",
		"8'd40: log_page_word = {30'h0, aer_event_pending, aer_registered}",
	} {
		if !strings.Contains(result, want) {
			t.Errorf("vendor diagnostic log should contain %q", want)
		}
	}
	if !strings.Contains(result, "LOG_PAGE_VENDOR: begin") {
		t.Error("Get Log Page should accept vendor diagnostic page")
	}
}

func TestGenerateNVMeResponderSV_QueueValidation(t *testing.T) {
	cfg := testConfig()

	result, err := GenerateNVMeResponderSV(cfg)
	if err != nil {
		t.Fatalf("GenerateNVMeResponderSV failed: %v", err)
	}

	for _, want := range []string{
		"wire        admin_aqa_valid",
		"wire        admin_queue_config_valid",
		"wire        cmd_prp1_page_valid",
		"if ((sq_tail != sq_head) && admin_queue_config_valid)",
		"io_cq_created &&",
		"cmd_prp1_page_valid) begin",
	} {
		if !strings.Contains(result, want) {
			t.Errorf("queue validation missing %q", want)
		}
	}
}

func TestGenerateNVMeResponderSV_HandlesIOAndFormatPath(t *testing.T) {
	cfg := testConfig()

	result, err := GenerateNVMeResponderSV(cfg)
	if err != nil {
		t.Fatalf("GenerateNVMeResponderSV failed: %v", err)
	}

	for _, want := range []string{
		"input               doorbell_wr",
		"input [15:0]        doorbell_qid",
		"io_sq_base",
		"io_cq_base",
		"8'h80: begin",       // Format NVM
		"S_FORMAT_CLEAR",     // ...actually clears the disk
		"S_IO_READ_DISK_REQ", // real disk-backed read path
		"S_IO_WRITE_DMA_REQ", // real disk-backed write path
		"disk_req_lba",       // disk backend port
	} {
		if !strings.Contains(result, want) {
			t.Fatalf("NVMe responder should contain %q", want)
		}
	}
}
