package svgen

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/firmware"
)

func TestPatchAll_ErrorOnMissingPatches(t *testing.T) {
	// Create temp dir with empty SV files - patches won't match
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
	err := patcher.PatchAll()
	if err == nil {
		t.Fatal("PatchAll should return error when critical patches are missing")
	}
	if !contains(err.Error(), "minimum patches applied") {
		t.Errorf("error should mention missing patches, got: %v", err)
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

func TestPatchCfgSVPreservesD0AndCommandRecovery(t *testing.T) {
	dir := t.TempDir()
	cfgPath := filepath.Join(dir, "pcileech_pcie_cfg_a7.sv")
	cfg := `module pcileech_pcie_cfg_a7;
    rw[210] <= 0; // cfg_pm_force_state_en
    rw[21] <= 0; // CFGSPACE_COMMAND_REGISTER_AUTO_SET
endmodule`
	if err := os.WriteFile(cfgPath, []byte(cfg), 0644); err != nil {
		t.Fatal(err)
	}

	if err := NewSVPatcher(firmware.DeviceIDs{}, dir).patchCfgSV(); err != nil {
		t.Fatalf("patchCfgSV: %v", err)
	}
	patched, err := os.ReadFile(cfgPath)
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{
		"rw[210] <= 1; // cfg_pm_force_state_en",
		"rw[21] <= 1; // CFGSPACE_COMMAND_REGISTER_AUTO_SET",
	} {
		if !contains(string(patched), want) {
			t.Errorf("patched cfg wrapper missing %q", want)
		}
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

func TestSVPatcherWiresSharedServicesThroughXilinxWrappers(t *testing.T) {
	dir := t.TempDir()
	fixtures := map[string]string{
		"pcileech_pcie_cfg_a7.sv": `module pcileech_pcie_cfg_a7(
    input clk
);
assign ctx.cfg_interrupt = rw[206];
endmodule`,
		"pcileech_pcie_tlp_a7.sv": `module pcileech_pcie_tlp_a7(
    input clk
);
pcileech_tlps128_bar_controller i_bar(
    .rst(rst)
);
endmodule`,
		"pcileech_pcie_a7.sv": `module pcileech_pcie_a7();
IfPCIeSignals ctx();
pcileech_pcie_cfg_a7 i_cfg(
    .ctx(ctx)
);
pcileech_pcie_tlp_a7 i_tlp(
    .clk(clk)
);
endmodule`,
	}
	for name, fixture := range fixtures {
		if err := os.WriteFile(filepath.Join(dir, name), []byte(fixture), 0644); err != nil {
			t.Fatalf("writing %s: %v", name, err)
		}
	}

	patcher := NewSVPatcher(firmware.DeviceIDs{}, dir)
	if err := patcher.patchCfgSV(); err != nil {
		t.Fatalf("patchCfgSV: %v", err)
	}
	if err := patcher.patchCfgInterruptHandshake("generated_bar_intr_req"); err != nil {
		t.Fatalf("patchCfgInterruptHandshake: %v", err)
	}
	if err := patcher.patchTLPServiceWiring(); err != nil {
		t.Fatalf("patchTLPServiceWiring: %v", err)
	}
	if err := patcher.patchPCIeWrapperServiceWiring(); err != nil {
		t.Fatalf("patchPCIeWrapperServiceWiring: %v", err)
	}

	cfg, err := os.ReadFile(filepath.Join(dir, "pcileech_pcie_cfg_a7.sv"))
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{
		"input                   generated_bar_intr_req",
		"reg intr_req_pending = 1'b0;",
		"else if (generated_bar_intr_req)",
		"else if (intr_req_pending && ctx.cfg_interrupt_rdy)",
		"assign ctx.cfg_interrupt = rw[206] | intr_req_pending;",
	} {
		if !contains(string(cfg), want) {
			t.Errorf("patched cfg wrapper missing %q", want)
		}
	}

	tlp, err := os.ReadFile(filepath.Join(dir, "pcileech_pcie_tlp_a7.sv"))
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{
		"IfPCIeSignals           ctx",
		"output wire             generated_bar_intr_req",
		".cfg_command     ( ctx.cfg_command",
		".cfg_power_state ( ctx.cfg_pmcsr_powerstate",
		".cfg_flr_in_process( ctx.cfg_received_func_lvl_rst",
		".cfg_to_turnoff  ( ctx.cfg_to_turnoff",
		".cfg_link_up     ( ctx.pl_phy_lnk_up",
		".intr_req        ( generated_bar_intr_req",
	} {
		if !contains(string(tlp), want) {
			t.Errorf("patched TLP wrapper missing %q", want)
		}
	}

	wrapper, err := os.ReadFile(filepath.Join(dir, "pcileech_pcie_a7.sv"))
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{
		"wire                    generated_bar_intr_req;",
		".generated_bar_intr_req   ( generated_bar_intr_req",
		".ctx                        ( ctx",
	} {
		if !contains(string(wrapper), want) {
			t.Errorf("patched PCIe wrapper missing %q", want)
		}
	}
}

func TestPatchAllUsesTLPDeclaredInterruptPort(t *testing.T) {
	dir := t.TempDir()
	fixtures := map[string]string{
		"pcileech_fifo.sv": `module pcileech_fifo;
    rw[175:160] <= 16'h0000; // CFG_VEND_ID
    rw[191:176] <= 16'h0000; // CFG_DEV_ID
endmodule`,
		"pcileech_pcie_cfg_a7.sv": `module pcileech_pcie_cfg_a7(
    input clk
);
assign ctx.cfg_interrupt = rw[206];
assign ctx.cfg_interrupt_assert = rw[205];
endmodule`,
		"pcileech_pcie_tlp_a7.sv": `module pcileech_pcie_tlp_a7(
    output intr_req,
    input clk
);
pcileech_tlps128_bar_controller i_bar(
    .rst(rst)
);
endmodule`,
		"pcileech_pcie_a7.sv": `module pcileech_pcie_a7();
IfPCIeSignals ctx();
pcileech_pcie_cfg_a7 i_cfg(
    .ctx(ctx)
);
pcileech_pcie_tlp_a7 i_tlp(
    .clk(clk)
);
endmodule`,
	}
	for name, fixture := range fixtures {
		if err := os.WriteFile(filepath.Join(dir, name), []byte(fixture), 0644); err != nil {
			t.Fatalf("writing %s: %v", name, err)
		}
	}

	if err := NewSVPatcher(firmware.DeviceIDs{VendorID: 0xBEEF, DeviceID: 0xCAFE}, dir).PatchAll(); err != nil {
		t.Fatalf("PatchAll: %v", err)
	}
	cfg, err := os.ReadFile(filepath.Join(dir, "pcileech_pcie_cfg_a7.sv"))
	if err != nil {
		t.Fatal(err)
	}
	if !contains(string(cfg), "input                   intr_req") {
		t.Fatalf("cfg wrapper did not use TLP interrupt port:\n%s", cfg)
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
