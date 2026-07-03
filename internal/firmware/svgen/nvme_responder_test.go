package svgen

import (
	"strings"
	"testing"
)

func TestGenerateNVMeResponderSV_UsesPRP2ForPageCrossingAdminData(t *testing.T) {
	cfg := testConfig()

	result, err := GenerateNVMeResponderSV(cfg)
	if err != nil {
		t.Fatalf("GenerateNVMeResponderSV failed: %v", err)
	}

	if !strings.Contains(result, "function [63:0] cmd_prp_addr") {
		t.Fatal("NVMe responder should compute admin data addresses through a PRP helper")
	}
	if !strings.Contains(result, "{cmd_prp2[63:12], 12'h000}") {
		t.Fatal("NVMe responder should use PRP2 for admin data crossing the first PRP page")
	}
	if strings.Contains(result, "dma_wr_addr  <= cmd_prp1 +") {
		t.Fatal("NVMe responder should not write all admin data relative to PRP1 only")
	}
}

func TestGenerateNVMeResponderSV_HandlesFormattingIOQueuePath(t *testing.T) {
	cfg := testConfig()

	result, err := GenerateNVMeResponderSV(cfg)
	if err != nil {
		t.Fatalf("GenerateNVMeResponderSV failed: %v", err)
	}

	for _, want := range []string{
		"input               sq1_doorbell_wr",
		"iosq_base",
		"iocq_base",
		"8'h80: begin",
		"8'h00, 8'h08, 8'h09: begin",
		"S_EXEC_READ_ZERO",
		"S_EXEC_WRITE_STORE",
		"nvme_store[io_store_addr]",
	} {
		if !strings.Contains(result, want) {
			t.Fatalf("NVMe responder should contain %q", want)
		}
	}
}

func TestGenerateNVMeResponderSV_HandlesCompatibilityLogsAndFeatures(t *testing.T) {
	cfg := testConfig()

	result, err := GenerateNVMeResponderSV(cfg)
	if err != nil {
		t.Fatalf("GenerateNVMeResponderSV failed: %v", err)
	}

	for _, want := range []string{
		"log_page_id",
		"log_dw_limit",
		"8'h01: dma_wr_data <= 32'h0;",
		"8'h03: dma_wr_data <= (data_dw_cnt == 11'h0) ? 32'h00000001 : 32'h0;",
		"8'h04: dma_wr_data <= 32'h0;",
		"volatile_write_cache_enabled",
		"async_event_config",
		"8'h06:   cqe[0] <= {31'h0, volatile_write_cache_enabled};",
		"8'h08:   cqe[0] <= interrupt_coalescing;",
		"8'h09:   cqe[0] <= interrupt_vector_config;",
		"8'h0B:   cqe[0] <= async_event_config;",
		"8'h0F:   cqe[0] <= keep_alive_timer;",
	} {
		if !strings.Contains(result, want) {
			t.Fatalf("NVMe responder should contain %q", want)
		}
	}
}

func TestGenerateNVMeResponderSV_TracksDebugState(t *testing.T) {
	cfg := testConfig()

	result, err := GenerateNVMeResponderSV(cfg)
	if err != nil {
		t.Fatalf("GenerateNVMeResponderSV failed: %v", err)
	}

	for _, want := range []string{
		"output reg [31:0]   debug_status",
		"output reg [31:0]   debug_last_cmd",
		"output reg [31:0]   debug_last_nsid",
		"output reg [31:0]   debug_last_cdw10",
		"output reg [31:0]   debug_queue_state",
		"debug_last_cmd <= {7'h0, active_io_cmd, cmd_opcode, cmd_cid};",
		"debug_last_nsid <= cmd_nsid;",
		"debug_last_cdw10 <= cmd_cdw10;",
	} {
		if !strings.Contains(result, want) {
			t.Fatalf("NVMe responder should contain debug state %q", want)
		}
	}
}
