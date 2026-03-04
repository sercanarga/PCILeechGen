package svgen

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/firmware"
)

func TestPatchAll_WarningOnMissingPatches(t *testing.T) {
	// Create temp dir with empty SV files — patches won't match
	dir := t.TempDir()

	// Write dummy fifo file with no matching patterns
	fifoContent := `module pcileech_fifo(input clk); endmodule`
	if err := os.WriteFile(filepath.Join(dir, "pcileech_fifo.sv"), []byte(fifoContent), 0644); err != nil {
		t.Fatal(err)
	}

	// Write dummy cfg file
	cfgContent := `module pcileech_pcie_cfg_a7(input clk); endmodule`
	if err := os.WriteFile(filepath.Join(dir, "pcileech_pcie_cfg_a7.sv"), []byte(cfgContent), 0644); err != nil {
		t.Fatal(err)
	}

	ids := firmware.DeviceIDs{
		VendorID:       0x144D,
		DeviceID:       0xA808,
		SubsysVendorID: 0x144D,
		SubsysDeviceID: 0xA801,
		RevisionID:     0x00,
		HasDSN:         true,
		DSN:            0x123456789ABCDEF0,
	}

	patcher := NewSVPatcher(ids, dir)
	if err := patcher.PatchAll(); err != nil {
		t.Fatalf("PatchAll error: %v", err)
	}

	// Should have warnings since patterns didn't match
	results := patcher.Results()
	hasWarning := false
	for _, r := range results {
		if len(r.Warnings) > 0 {
			hasWarning = true
			break
		}
	}
	if !hasWarning {
		t.Error("expected warnings when patches don't match, got none")
	}
}

func TestPatchAll_SuccessfulPatching(t *testing.T) {
	dir := t.TempDir()

	// Write fifo file with matching patterns
	fifoContent := `module pcileech_fifo(input clk);
    rw[143:128] <= 16'h0000; // CFG_SUBSYS_VEND_ID
    rw[159:144] <= 16'h0000; // CFG_SUBSYS_ID
    rw[175:160] <= 16'h0000; // CFG_VEND_ID
    rw[191:176] <= 16'h0000; // CFG_DEV_ID
    rw[199:192] <= 8'h00;    // CFG_REV_ID
    rw[203] <= 1'b1; // CFGTLP ZERO DATA
endmodule`
	if err := os.WriteFile(filepath.Join(dir, "pcileech_fifo.sv"), []byte(fifoContent), 0644); err != nil {
		t.Fatal(err)
	}

	cfgContent := `module pcileech_pcie_cfg_a7(input clk); endmodule`
	if err := os.WriteFile(filepath.Join(dir, "pcileech_pcie_cfg_a7.sv"), []byte(cfgContent), 0644); err != nil {
		t.Fatal(err)
	}

	ids := firmware.DeviceIDs{
		VendorID:       0xBEEF,
		DeviceID:       0xCAFE,
		SubsysVendorID: 0x1234,
		SubsysDeviceID: 0x5678,
		RevisionID:     0xAB,
	}

	patcher := NewSVPatcher(ids, dir)
	if err := patcher.PatchAll(); err != nil {
		t.Fatalf("PatchAll error: %v", err)
	}

	// Verify patches were applied
	results := patcher.Results()
	patchCount := 0
	for _, r := range results {
		patchCount += len(r.Patches)
	}
	if patchCount < 2 {
		t.Errorf("expected at least 2 patches, got %d", patchCount)
	}

	// Verify file content was modified
	modified, _ := os.ReadFile(filepath.Join(dir, "pcileech_fifo.sv"))
	content := string(modified)
	if !contains(content, "BEEF") {
		t.Error("patched fifo should contain VendorID BEEF")
	}
	if !contains(content, "CAFE") {
		t.Error("patched fifo should contain DeviceID CAFE")
	}
}

func TestFormatPatchSummary_Empty(t *testing.T) {
	result := FormatPatchSummary(nil)
	if result != "  (no patches applied)" {
		t.Errorf("expected empty summary, got %q", result)
	}
}

func TestFormatPatchSummary_WithWarnings(t *testing.T) {
	results := []PatchResult{
		{File: "test.sv", Warnings: []string{"warning1"}},
	}
	summary := FormatPatchSummary(results)
	if !contains(summary, "warning1") {
		t.Error("summary should contain warning")
	}
}

func contains(s, sub string) bool {
	return len(s) >= len(sub) && (s == sub || len(s) > 0 && containsStr(s, sub))
}

func containsStr(s, sub string) bool {
	for i := 0; i <= len(s)-len(sub); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}
