package svgen

import (
	"strings"
	"testing"
)

// TestGenerateNVMeResponderSV_UsesPRP2ForPageCrossingAdminData verifies admin
// data (Identify / log pages / CQE) addresses route through a PRP helper that
// falls back to PRP2 once the first PRP page is exhausted.
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

// TestGenerateNVMeResponderSV_NoHardcodedGigabyteStrings verifies SN/MN/FR are
// not overridden with hardcoded literals and fall through to the donor ROM.
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

// TestGenerateNVMeResponderSV_HandlesIOAndFormatPath verifies the responder
// exposes the generalized doorbell interface, admin/I/O queue bookkeeping, the
// Format NVM clear path, and a real disk-backed read path (read-after-write).
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
		"8'h80: begin",          // Format NVM
		"S_FORMAT_CLEAR",        // ...actually clears the disk
		"S_IO_READ_DISK_REQ",    // real disk-backed read path
		"S_IO_WRITE_DMA_REQ",    // real disk-backed write path
		"disk_req_lba",          // disk backend port
	} {
		if !strings.Contains(result, want) {
			t.Fatalf("NVMe responder should contain %q", want)
		}
	}
}
