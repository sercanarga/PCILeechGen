package output

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/board"
)

func TestGenerateManifest_Empty(t *testing.T) {
	tmpDir := t.TempDir()
	m, err := GenerateManifest(tmpDir, "1.0.0", "TestBoard", 0x8086, 0x1533)
	if err != nil {
		t.Fatalf("GenerateManifest error: %v", err)
	}
	if m.ToolVersion != "1.0.0" {
		t.Errorf("ToolVersion = %q, want '1.0.0'", m.ToolVersion)
	}
	if m.VendorID != 0x8086 {
		t.Errorf("VendorID = %04x, want 0x8086", m.VendorID)
	}
	if m.DeviceID != 0x1533 {
		t.Errorf("DeviceID = %04x, want 0x1533", m.DeviceID)
	}
	if len(m.Files) != 0 {
		t.Errorf("Files should be empty, got %d", len(m.Files))
	}
}

func TestGenerateManifest_WithFiles(t *testing.T) {
	tmpDir := t.TempDir()
	// Create some output files
	if err := os.WriteFile(filepath.Join(tmpDir, "pcileech_cfgspace.coe"), []byte("test content"), 0644); err != nil {
		t.Fatalf("write cfgspace fixture: %v", err)
	}
	if err := os.WriteFile(filepath.Join(tmpDir, "device_context.json"), []byte("{}"), 0644); err != nil {
		t.Fatalf("write context fixture: %v", err)
	}

	m, err := GenerateManifest(tmpDir, "1.0.0", "TestBoard", 0x8086, 0x1533)
	if err != nil {
		t.Fatalf("GenerateManifest error: %v", err)
	}
	if len(m.Files) < 2 {
		t.Errorf("Files should include created files, got %d", len(m.Files))
	}
	for _, f := range m.Files {
		if f.SHA256 == "" || f.SHA256 == "error" {
			t.Errorf("File %s has bad SHA256: %q", f.Name, f.SHA256)
		}
		if f.Size == 0 {
			t.Errorf("File %s has zero size", f.Name)
		}
	}
}

func TestGenerateManifestForBoard_recordsBoardProfile(t *testing.T) {
	tmpDir := t.TempDir()
	b := &board.Board{
		Name:         "CaptainDMA_100T",
		FPGAPart:     "xc7a100tfgg484-2",
		PCIeLanes:    1,
		MaxLinkSpeed: 2,
		BRAMSize:     32768,
		TopModule:    "pcileech_100t484_x1_top",
		ProjectDir:   "CaptainDMA",
		SubDir:       "100t484-1",
		TCLFile:      "vivado_generate_project_captaindma_100t.tcl",
	}

	m, err := GenerateManifestForBoard(tmpDir, "1.0.0", b, 0x8086, 0x1533)
	if err != nil {
		t.Fatalf("GenerateManifestForBoard error: %v", err)
	}
	if m.Board != "CaptainDMA_100T" {
		t.Errorf("Board = %q, want CaptainDMA_100T", m.Board)
	}
	if m.BoardProfile == nil {
		t.Fatal("BoardProfile should be recorded")
	}
	if m.BoardProfile.BRAMSize != 32768 {
		t.Errorf("BoardProfile.BRAMSize = %d, want 32768", m.BoardProfile.BRAMSize)
	}
	if m.BoardProfile.SubDir != "100t484-1" {
		t.Errorf("BoardProfile.SubDir = %q, want 100t484-1", m.BoardProfile.SubDir)
	}
	if m.Provenance == nil {
		t.Fatal("Provenance should be recorded")
	}
	if m.Provenance.BoardProfileHash == "" {
		t.Error("Provenance.BoardProfileHash should be populated when a board profile is present")
	}
	if m.Provenance.GeneratorGitCommit == "" {
		t.Error("Provenance.GeneratorGitCommit should not be empty")
	}
}

func TestGenerateManifestForBuild_recordsDonorAndIntakeProvenance(t *testing.T) {
	tmpDir := t.TempDir()
	b := &board.Board{
		Name:       "TestBoard",
		FPGAPart:   "xc7a35tfgg484-2",
		PCIeLanes:  1,
		TopModule:  "pcileech_test_top",
		ProjectDir: "TestBoard",
		TCLFile:    "vivado_generate_project.tcl",
	}
	ctx := makeDonorContext(0x8086, 0x1533, 0x020000)
	refs := []string{"https://example/intake/1"}

	m, err := GenerateManifestForBuild(tmpDir, "1.0.0", b, ctx, refs, 0x8086, 0x1533)
	if err != nil {
		t.Fatalf("GenerateManifestForBuild error: %v", err)
	}
	if m.Provenance == nil {
		t.Fatal("Provenance should be recorded")
	}
	if m.Provenance.DonorSnapshotHash == "" {
		t.Error("Provenance.DonorSnapshotHash should be populated when a donor context is present")
	}
	if m.Provenance.BoardProfileHash == "" {
		t.Error("Provenance.BoardProfileHash should be populated when a board profile is present")
	}
	if len(m.Provenance.ExternalIntakeRefs) != 1 || m.Provenance.ExternalIntakeRefs[0] != refs[0] {
		t.Errorf("ExternalIntakeRefs = %v, want %v", m.Provenance.ExternalIntakeRefs, refs)
	}
	// Vivado is almost certainly not installed in CI; verify it degrades to
	// either a real version or omitted (empty string via omitempty) — never a
	// non-version placeholder.
	if v := m.Provenance.VivadoVersion; v != "" && !strings.Contains(v, ".") {
		t.Errorf("VivadoVersion = %q, expected a dotted version or omitted", v)
	}
}

func TestGenerateManifestForBuild_provenanceIsDeterministicForFixedInputs(t *testing.T) {
	tmpDir := t.TempDir()
	b := &board.Board{
		Name:       "DetBoard",
		FPGAPart:   "xc7a35tfgg484-2",
		PCIeLanes:  1,
		TopModule:  "pcileech_top",
		ProjectDir: "DetBoard",
		TCLFile:    "vivado.tcl",
	}
	ctx := makeDonorContext(0x8086, 0x1533, 0x020000)

	m1, err := GenerateManifestForBuild(tmpDir, "1.0.0", b, ctx, nil, 0x8086, 0x1533)
	if err != nil {
		t.Fatalf("first GenerateManifestForBuild error: %v", err)
	}
	m2, err := GenerateManifestForBuild(tmpDir, "1.0.0", b, ctx, nil, 0x8086, 0x1533)
	if err != nil {
		t.Fatalf("second GenerateManifestForBuild error: %v", err)
	}
	if m1.Provenance.DonorSnapshotHash != m2.Provenance.DonorSnapshotHash {
		t.Errorf("DonorSnapshotHash not deterministic: %q vs %q",
			m1.Provenance.DonorSnapshotHash, m2.Provenance.DonorSnapshotHash)
	}
	if m1.Provenance.BoardProfileHash != m2.Provenance.BoardProfileHash {
		t.Errorf("BoardProfileHash not deterministic: %q vs %q",
			m1.Provenance.BoardProfileHash, m2.Provenance.BoardProfileHash)
	}
}

func TestGenerateManifestForBuild_degradesGracefullyWithoutGitOrVivado(t *testing.T) {
	// This test documents the graceful-degradation contract: provenance is
	// always set, and unavailable sources become Unknown/omitted WITHOUT
	// failing the build. We cannot easily uninstall git/vivado in the test
	// environment, so we assert the contract holds with whatever is present:
	// commit is either a real hash or Unknown (never empty), VivadoVersion is
	// either a version or omitted (never empty), and the manifest builds.
	tmpDir := t.TempDir()
	m, err := GenerateManifestForBuild(tmpDir, "1.0.0", nil, nil, nil, 0, 0)
	if err != nil {
		t.Fatalf("GenerateManifestForBuild should not fail when git/vivado unavailable: %v", err)
	}
	if m.Provenance == nil {
		t.Fatal("Provenance should be set even when sources are unavailable")
	}
	if m.Provenance.GeneratorGitCommit == "" {
		t.Error("GeneratorGitCommit should be Unknown or a hash, never empty")
	}
	if m.Provenance.DonorSnapshotHash != "" {
		t.Errorf("DonorSnapshotHash should be omitted when no donor, got %q", m.Provenance.DonorSnapshotHash)
	}
	if m.Provenance.BoardProfileHash != "" {
		t.Errorf("BoardProfileHash should be omitted when no board, got %q", m.Provenance.BoardProfileHash)
	}
	// VivadoVersion: empty (omitted) or a real version are both acceptable.
	if v := m.Provenance.VivadoVersion; v != "" && !strings.Contains(v, ".") {
		t.Errorf("VivadoVersion = %q, expected a dotted version or omitted", v)
	}
}

func TestWriteAll_recordsBoardProfileInManifest(t *testing.T) {
	tmpDir := t.TempDir()
	libDir := filepath.Join(tmpDir, "lib")
	srcDir := filepath.Join(libDir, "TestBoard", "src")
	if err := os.MkdirAll(srcDir, 0755); err != nil {
		t.Fatalf("create source dir: %v", err)
	}
	controller := `module pcileech_tlps128_bar_rdengine; endmodule
module pcileech_tlps128_bar_wrengine; endmodule
module pcileech_bar_impl_none; endmodule
module pcileech_bar_impl_loopaddr; endmodule
module pcileech_bar_impl_zerowrite4k; endmodule
`
	if err := os.WriteFile(filepath.Join(srcDir, "pcileech_tlps128_bar_controller.sv"), []byte(controller), 0644); err != nil {
		t.Fatalf("write board controller: %v", err)
	}
	fifo := `module pcileech_fifo(input clk);
    rw[143:128] <= 16'h0000; // CFG_SUBSYS_VEND_ID
    rw[159:144] <= 16'h0000; // CFG_SUBSYS_ID
    rw[175:160] <= 16'h0000; // CFG_VEND_ID
    rw[191:176] <= 16'h0000; // CFG_DEV_ID
    rw[199:192] <= 8'h00;    // CFG_REV_ID
    rw[203] <= 1'b1; // CFGTLP ZERO DATA
    rw[206] <= 1'b0; // CFGTLP PCIE WRITE ENABLE
endmodule`
	if err := os.WriteFile(filepath.Join(srcDir, "pcileech_fifo.sv"), []byte(fifo), 0644); err != nil {
		t.Fatalf("write fifo source: %v", err)
	}
	cfg := `module pcileech_pcie_cfg_a7(input clk);
    rw[127:64] <= 64'hFFFFFFFFFFFFFFFF; // cfg_dsn
    rw[211] <= 0; // cfg_pm_halt_aspm_l0s
    rw[212] <= 0; // cfg_pm_halt_aspm_l1
    rw[210] <= 0; // cfg_pm_force_state_en
    rw[21] <= 0; // CFGSPACE_COMMAND_REGISTER_AUTO_SET
    rw[20] <= 0; // CFGSPACE_STATUS_REGISTER_AUTO_CLEAR
endmodule`
	if err := os.WriteFile(filepath.Join(srcDir, "pcileech_pcie_cfg_a7.sv"), []byte(cfg), 0644); err != nil {
		t.Fatalf("write cfg source: %v", err)
	}

	b := &board.Board{
		Name:         "TestBoard",
		FPGAPart:     "xc7a35tfgg484-2",
		PCIeLanes:    1,
		MaxLinkSpeed: 2,
		BRAMSize:     4096,
		TopModule:    "pcileech_test_top",
		ProjectDir:   "TestBoard",
		TCLFile:      "vivado_generate_project.tcl",
	}
	outputDir := filepath.Join(tmpDir, "out")
	ow := NewOutputWriter(outputDir, libDir, 0, 0)
	ctx := makeDonorContext(0x8086, 0x1533, 0x020000)

	if err := ow.WriteAll(ctx, b); err != nil {
		t.Fatalf("WriteAll failed: %v", err)
	}

	m, err := LoadManifest(filepath.Join(outputDir, "build_manifest.json"))
	if err != nil {
		t.Fatalf("LoadManifest failed: %v", err)
	}
	if m.Board != "TestBoard" {
		t.Errorf("Board = %q, want TestBoard", m.Board)
	}
	if m.BoardProfile == nil {
		t.Fatal("BoardProfile should be recorded")
	}
	if m.BoardProfile.TopModule != "pcileech_test_top" {
		t.Errorf("TopModule = %q, want pcileech_test_top", m.BoardProfile.TopModule)
	}
}

func TestGenerateManifest_WithSrcDir(t *testing.T) {
	tmpDir := t.TempDir()
	srcDir := filepath.Join(tmpDir, "src")
	if err := os.MkdirAll(srcDir, 0755); err != nil {
		t.Fatalf("create src dir: %v", err)
	}
	if err := os.WriteFile(filepath.Join(srcDir, "test.sv"), []byte("module test; endmodule"), 0644); err != nil {
		t.Fatalf("write src fixture: %v", err)
	}

	m, err := GenerateManifest(tmpDir, "1.0.0", "", 0, 0)
	if err != nil {
		t.Fatalf("GenerateManifest error: %v", err)
	}

	found := false
	for _, f := range m.Files {
		if f.Name == filepath.Join("src", "test.sv") {
			found = true
			break
		}
	}
	if !found {
		t.Error("Manifest should include src/ files")
	}
}

func TestWriteJSON(t *testing.T) {
	tmpDir := t.TempDir()
	m := &BuildManifest{
		ToolVersion: "1.0.0",
		VendorID:    0x8086,
		DeviceID:    0x1533,
		Files: []ManifestEntry{
			{Name: "test.coe", Size: 100, SHA256: "abc123"},
		},
	}

	path := filepath.Join(tmpDir, "manifest.json")
	if err := m.WriteJSON(path); err != nil {
		t.Fatalf("WriteJSON error: %v", err)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile error: %v", err)
	}

	var loaded BuildManifest
	if err := json.Unmarshal(data, &loaded); err != nil {
		t.Fatalf("Unmarshal error: %v", err)
	}
	if loaded.VendorID != 0x8086 {
		t.Errorf("Loaded VendorID = %04x, want 0x8086", loaded.VendorID)
	}
	if len(loaded.Files) != 1 {
		t.Errorf("Loaded Files = %d, want 1", len(loaded.Files))
	}
}

func TestFileHash(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "test.txt")
	if err := os.WriteFile(path, []byte("hello"), 0644); err != nil {
		t.Fatalf("write hash fixture: %v", err)
	}

	hash, err := fileHash(path)
	if err != nil {
		t.Fatalf("fileHash error: %v", err)
	}
	if len(hash) != 64 {
		t.Errorf("SHA256 hash length = %d, want 64", len(hash))
	}

	// Same content -> same hash
	path2 := filepath.Join(tmpDir, "test2.txt")
	if writeErr := os.WriteFile(path2, []byte("hello"), 0644); writeErr != nil {
		t.Fatalf("write second hash fixture: %v", writeErr)
	}
	hash2, err := fileHash(path2)
	if err != nil {
		t.Fatalf("hash second fixture: %v", err)
	}
	if hash != hash2 {
		t.Error("Same content should produce same hash")
	}

	// Non-existent file -> error
	_, err = fileHash("/nonexistent")
	if err == nil {
		t.Error("fileHash should fail for non-existent file")
	}
}

func TestVerifyManifest_AllPass(t *testing.T) {
	tmpDir := t.TempDir()

	if err := os.WriteFile(filepath.Join(tmpDir, "test.coe"), []byte("content"), 0644); err != nil {
		t.Fatalf("write manifest fixture: %v", err)
	}

	hash, err := fileHash(filepath.Join(tmpDir, "test.coe"))
	if err != nil {
		t.Fatalf("hash fixture: %v", err)
	}
	m := &BuildManifest{
		Files: []ManifestEntry{
			{Name: "test.coe", Size: 7, SHA256: hash},
		},
	}
	manifestPath := filepath.Join(tmpDir, "manifest.json")
	if writeErr := m.WriteJSON(manifestPath); writeErr != nil {
		t.Fatalf("write manifest: %v", writeErr)
	}

	v, err := VerifyManifest(manifestPath, tmpDir)
	if err != nil {
		t.Fatalf("VerifyManifest error: %v", err)
	}
	if !v.OK() {
		t.Errorf("expected OK, got: %s", v.Summary())
	}
	if len(v.Passed) != 1 {
		t.Errorf("Passed = %d, want 1", len(v.Passed))
	}
}

func TestVerifyManifest_MissingFile(t *testing.T) {
	tmpDir := t.TempDir()
	m := &BuildManifest{
		Files: []ManifestEntry{
			{Name: "missing.coe", Size: 100, SHA256: "abc"},
		},
	}
	manifestPath := filepath.Join(tmpDir, "manifest.json")
	if err := m.WriteJSON(manifestPath); err != nil {
		t.Fatalf("write manifest: %v", err)
	}

	v, err := VerifyManifest(manifestPath, tmpDir)
	if err != nil {
		t.Fatalf("VerifyManifest error: %v", err)
	}
	if v.OK() {
		t.Error("expected failure for missing file")
	}
	if len(v.Missing) != 1 {
		t.Errorf("Missing = %d, want 1", len(v.Missing))
	}
}

func TestVerifyManifest_HashMismatch(t *testing.T) {
	tmpDir := t.TempDir()
	if err := os.WriteFile(filepath.Join(tmpDir, "test.coe"), []byte("content"), 0644); err != nil {
		t.Fatalf("write manifest fixture: %v", err)
	}

	m := &BuildManifest{
		Files: []ManifestEntry{
			{Name: "test.coe", Size: 7, SHA256: "wrong_hash"},
		},
	}
	manifestPath := filepath.Join(tmpDir, "manifest.json")
	if err := m.WriteJSON(manifestPath); err != nil {
		t.Fatalf("write manifest: %v", err)
	}

	v, err := VerifyManifest(manifestPath, tmpDir)
	if err != nil {
		t.Fatalf("VerifyManifest error: %v", err)
	}
	if v.OK() {
		t.Error("expected failure for hash mismatch")
	}
	if len(v.Failed) != 1 {
		t.Errorf("Failed = %d, want 1", len(v.Failed))
	}
}

func TestLoadManifest_Invalid(t *testing.T) {
	_, err := LoadManifest("/nonexistent")
	if err == nil {
		t.Error("LoadManifest should fail for non-existent file")
	}

	tmpDir := t.TempDir()
	if writeErr := os.WriteFile(filepath.Join(tmpDir, "bad.json"), []byte("not json"), 0644); writeErr != nil {
		t.Fatalf("write bad manifest fixture: %v", writeErr)
	}
	_, err = LoadManifest(filepath.Join(tmpDir, "bad.json"))
	if err == nil {
		t.Error("LoadManifest should fail for invalid JSON")
	}
}
