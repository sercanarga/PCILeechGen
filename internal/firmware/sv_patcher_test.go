package firmware

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func writeMockSV(t *testing.T, dir, filename, content string) {
	t.Helper()
	if err := os.MkdirAll(dir, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, filename), []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
}

func readSV(t *testing.T, dir, filename string) string {
	t.Helper()
	data, err := os.ReadFile(filepath.Join(dir, filename))
	if err != nil {
		t.Fatal(err)
	}
	return string(data)
}

const mockCfgSV = `// pcileech_pcie_cfg_a7.sv
task pcileech_pcie_cfg_a7_initialvalues;
    begin
        rw[127:64]  <= 64'h0000000101000A35;    // +008: cfg_dsn
    end
endtask

assign ctx.cfg_dsn = rw[127:64];
`

const mockFifoSV = `// pcileech_fifo.sv
reg     [79:0]      _pcie_core_config = { 4'hf, 1'b1, 1'b1, 1'b0, 1'b0, 8'h02, 16'h0666, 16'h10EE, 16'h0007, 16'h10EE };

task pcileech_fifo_ctl_initialvalues;
    begin
        rw[143:128] <= 16'h10EE;                    // +010: CFG_SUBSYS_VEND_ID (NOT IMPLEMENTED)
        rw[159:144] <= 16'h0007;                    // +012: CFG_SUBSYS_ID      (NOT IMPLEMENTED)
        rw[175:160] <= 16'h10EE;                    // +014: CFG_VEND_ID        (NOT IMPLEMENTED)
        rw[191:176] <= 16'h0666;                    // +016: CFG_DEV_ID         (NOT IMPLEMENTED)
        rw[199:192] <= 8'h02;                       // +018: CFG_REV_ID         (NOT IMPLEMENTED)
        rw[203]     <= 1'b1;                        //       CFGTLP ZERO DATA
    end
endtask
`

func TestSVPatcherCfgDSN(t *testing.T) {
	srcDir := filepath.Join(t.TempDir(), "src")
	writeMockSV(t, srcDir, "pcileech_pcie_cfg_a7.sv", mockCfgSV)

	ids := DeviceIDs{
		HasDSN: true,
		DSN:    0xDEADBEEF12345678,
	}

	patcher := NewSVPatcher(ids, srcDir)
	if err := patcher.patchCfgSV(); err != nil {
		t.Fatal(err)
	}

	content := readSV(t, srcDir, "pcileech_pcie_cfg_a7.sv")
	if !strings.Contains(content, "DEADBEEF12345678") {
		t.Error("DSN not patched in cfg SV")
	}
	if strings.Contains(content, "0000000101000A35") {
		t.Error("Old DSN still present")
	}

	if len(patcher.Results()) != 1 {
		t.Errorf("Expected 1 result, got %d", len(patcher.Results()))
	}
}

func TestSVPatcherCfgNoDSN(t *testing.T) {
	srcDir := filepath.Join(t.TempDir(), "src")
	writeMockSV(t, srcDir, "pcileech_pcie_cfg_a7.sv", mockCfgSV)

	ids := DeviceIDs{
		HasDSN: false,
	}

	patcher := NewSVPatcher(ids, srcDir)
	if err := patcher.patchCfgSV(); err != nil {
		t.Fatal(err)
	}

	content := readSV(t, srcDir, "pcileech_pcie_cfg_a7.sv")
	// DSN should NOT be changed when donor has no DSN
	if !strings.Contains(content, "0000000101000A35") {
		t.Error("DSN should remain unchanged when donor has no DSN")
	}
}

func TestSVPatcherFifo(t *testing.T) {
	srcDir := filepath.Join(t.TempDir(), "src")
	writeMockSV(t, srcDir, "pcileech_fifo.sv", mockFifoSV)

	ids := DeviceIDs{
		VendorID:       0x8086,
		DeviceID:       0x1533,
		SubsysVendorID: 0x8086,
		SubsysDeviceID: 0x0001,
		RevisionID:     0x03,
	}

	patcher := NewSVPatcher(ids, srcDir)
	if err := patcher.patchFifoSV(); err != nil {
		t.Fatal(err)
	}

	content := readSV(t, srcDir, "pcileech_fifo.sv")

	// Shadow config space should be enabled
	if !strings.Contains(content, "1'b0") || strings.Contains(content, "rw[203]     <= 1'b1") {
		t.Error("CFGTLP ZERO DATA not patched to 0")
	}

	// Vendor/Device IDs should be patched
	if !strings.Contains(content, "16'h8086") {
		t.Error("Vendor ID not patched")
	}
	if !strings.Contains(content, "16'h1533") {
		t.Error("Device ID not patched")
	}
	if !strings.Contains(content, "16'h0001") {
		t.Error("Subsys Device ID not patched")
	}
	if !strings.Contains(content, "8'h03") {
		t.Error("Revision ID not patched")
	}

	// _pcie_core_config should be updated
	if !strings.Contains(content, "8'h03, 16'h1533, 16'h8086, 16'h0001, 16'h8086") {
		t.Error("_pcie_core_config not patched correctly")
	}

	results := patcher.Results()
	if len(results) != 1 {
		t.Fatalf("Expected 1 result, got %d", len(results))
	}
	if len(results[0].Patches) < 6 {
		t.Errorf("Expected at least 6 patches, got %d", len(results[0].Patches))
	}
}

func TestSVPatcherPatchAll(t *testing.T) {
	srcDir := filepath.Join(t.TempDir(), "src")
	writeMockSV(t, srcDir, "pcileech_pcie_cfg_a7.sv", mockCfgSV)
	writeMockSV(t, srcDir, "pcileech_fifo.sv", mockFifoSV)

	ids := DeviceIDs{
		VendorID:       0x10EC,
		DeviceID:       0x8168,
		SubsysVendorID: 0x10EC,
		SubsysDeviceID: 0x0123,
		RevisionID:     0x15,
		HasDSN:         true,
		DSN:            0xABCDEF0123456789,
	}

	patcher := NewSVPatcher(ids, srcDir)
	if err := patcher.PatchAll(); err != nil {
		t.Fatal(err)
	}

	// Verify cfg file
	cfgContent := readSV(t, srcDir, "pcileech_pcie_cfg_a7.sv")
	if !strings.Contains(cfgContent, "ABCDEF0123456789") {
		t.Error("DSN not applied in PatchAll")
	}

	// Verify fifo file
	fifoContent := readSV(t, srcDir, "pcileech_fifo.sv")
	if !strings.Contains(fifoContent, "16'h10EC") {
		t.Error("Vendor ID not applied in PatchAll")
	}
	if !strings.Contains(fifoContent, "16'h8168") {
		t.Error("Device ID not applied in PatchAll")
	}

	summary := FormatPatchSummary(patcher.Results())
	if len(summary) < 50 {
		t.Errorf("Summary too short: %s", summary)
	}
}

func TestFormatPatchSummaryEmpty(t *testing.T) {
	s := FormatPatchSummary(nil)
	if !strings.Contains(s, "no patches") {
		t.Error("Empty summary should say no patches")
	}
}
