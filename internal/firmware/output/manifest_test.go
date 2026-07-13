package output

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/pci"
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
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(tmpDir, "device_context.json"), []byte("{}"), 0644); err != nil {
		t.Fatal(err)
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

func TestGenerateManifest_WithSrcDir(t *testing.T) {
	tmpDir := t.TempDir()
	srcDir := filepath.Join(tmpDir, "src")
	if err := os.MkdirAll(srcDir, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(srcDir, "test.sv"), []byte("module test; endmodule"), 0644); err != nil {
		t.Fatal(err)
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
		t.Fatal(err)
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
		t.Fatal(writeErr)
	}
	hash2, _ := fileHash(path2)
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
		t.Fatal(err)
	}

	hash, _ := fileHash(filepath.Join(tmpDir, "test.coe"))
	m := &BuildManifest{
		Files: []ManifestEntry{
			{Name: "test.coe", Size: 7, SHA256: hash},
		},
	}
	manifestPath := filepath.Join(tmpDir, "manifest.json")
	if err := m.WriteJSON(manifestPath); err != nil {
		t.Fatal(err)
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
			{Name: "missing.coe", Size: 100, SHA256: strings.Repeat("0", 64)},
		},
	}
	manifestPath := filepath.Join(tmpDir, "manifest.json")
	if err := m.WriteJSON(manifestPath); err != nil {
		t.Fatal(err)
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
		t.Fatal(err)
	}

	m := &BuildManifest{
		Files: []ManifestEntry{
			{Name: "test.coe", Size: 7, SHA256: strings.Repeat("0", 64)},
		},
	}
	manifestPath := filepath.Join(tmpDir, "manifest.json")
	if err := m.WriteJSON(manifestPath); err != nil {
		t.Fatal(err)
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
		t.Fatal(writeErr)
	}
	_, err = LoadManifest(filepath.Join(tmpDir, "bad.json"))
	if err == nil {
		t.Error("LoadManifest should fail for invalid JSON")
	}
}

func TestGenerateManifest_DiscoversRootDeliverablesByExtension(t *testing.T) {
	tmpDir := t.TempDir()
	files := map[string]string{
		"src/nested/core.sv":          "module core; endmodule",
		"a-output.bit":                "bit",
		"z-output.bin":                "bin",
		"identify_init.hex":           "identify",
		"pcileech_bar_impl_device.sv": "module pcileech_bar_impl_device; endmodule",
		"pcileech_hda_rirb_dma.sv":    "module pcileech_hda_rirb_dma; endmodule",
		"build_manifest.json":         `{"files":[]}`,
		"vivado.log":                  "build log",
		"vivado.jou":                  "build journal",
		"post-synthesis-scratch.tmp":  "temporary",
	}
	for name, content := range files {
		filePath := filepath.Join(tmpDir, filepath.FromSlash(name))
		if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
			t.Fatal(err)
		}
	}

	m, err := GenerateManifest(tmpDir, "1.2.3", "CaptainDMA_35T", 0x1234, 0x5678)
	if err != nil {
		t.Fatalf("GenerateManifest error: %v", err)
	}
	names := make([]string, 0, len(m.Files))
	for _, entry := range m.Files {
		names = append(names, entry.Name)
	}
	want := []string{
		"a-output.bit",
		"identify_init.hex",
		"pcileech_bar_impl_device.sv",
		"pcileech_hda_rirb_dma.sv",
		"src/nested/core.sv",
		"z-output.bin",
	}
	if !reflect.DeepEqual(names, want) {
		t.Fatalf("manifest names = %v, want only generated deliverables %v", names, want)
	}
}

func TestGenerateManifest_RejectsSourceSymlink(t *testing.T) {
	tmpDir := t.TempDir()
	if err := os.Mkdir(filepath.Join(tmpDir, "src"), 0755); err != nil {
		t.Fatal(err)
	}
	outside := filepath.Join(t.TempDir(), "outside.sv")
	if err := os.WriteFile(outside, []byte("secret"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(outside, filepath.Join(tmpDir, "src", "linked.sv")); err != nil {
		t.Skipf("symlink unavailable: %v", err)
	}
	if _, err := GenerateManifest(tmpDir, "dev", "board", 0, 0); err == nil {
		t.Fatal("GenerateManifest should reject symlinked artifacts")
	}
}

func TestGenerateManifest_RejectsRootDeliverableSymlink(t *testing.T) {
	tmpDir := t.TempDir()
	outside := filepath.Join(t.TempDir(), "outside.sv")
	if err := os.WriteFile(outside, []byte("module outside; endmodule"), 0644); err != nil {
		t.Fatal(err)
	}
	linkPath := filepath.Join(tmpDir, "pcileech_bar_impl_device.sv")
	if err := os.Symlink(outside, linkPath); err != nil {
		t.Skipf("symlink unavailable: %v", err)
	}

	_, err := GenerateManifest(tmpDir, "dev", "board", 0, 0)
	if err == nil {
		t.Fatal("GenerateManifest accepted a root deliverable symlink")
	}
	want := "refusing to include symlink " + linkPath
	if err.Error() != want {
		t.Fatalf("GenerateManifest error = %q, want %q", err, want)
	}
}

func TestVerifyManifest_RejectsUnsafeAndMalformedEntries(t *testing.T) {
	tests := []struct {
		name  string
		entry ManifestEntry
	}{
		{"parent traversal", ManifestEntry{Name: "../secret", Size: 0, SHA256: strings.Repeat("0", 64)}},
		{"absolute path", ManifestEntry{Name: "/etc/passwd", Size: 0, SHA256: strings.Repeat("0", 64)}},
		{"unclean path", ManifestEntry{Name: "src/../file", Size: 0, SHA256: strings.Repeat("0", 64)}},
		{"backslash path", ManifestEntry{Name: `src\\file`, Size: 0, SHA256: strings.Repeat("0", 64)}},
		{"negative size", ManifestEntry{Name: "file", Size: -1, SHA256: strings.Repeat("0", 64)}},
		{"invalid digest", ManifestEntry{Name: "file", Size: 0, SHA256: "not-a-digest"}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tmpDir := t.TempDir()
			m := &BuildManifest{Files: []ManifestEntry{tc.entry}}
			manifestPath := filepath.Join(tmpDir, "manifest.json")
			if err := m.WriteJSON(manifestPath); err != nil {
				t.Fatal(err)
			}
			if _, err := VerifyManifest(manifestPath, tmpDir); err == nil {
				t.Fatal("VerifyManifest should reject malformed entry")
			}
		})
	}
}

func TestVerifyManifest_ReportsNULBytePath(t *testing.T) {
	tmpDir := t.TempDir()
	manifestPath := filepath.Join(tmpDir, "manifest.json")
	m := &BuildManifest{
		Files: []ManifestEntry{{
			Name:   "artifact\x00.sv",
			Size:   1,
			SHA256: strings.Repeat("0", 64),
		}},
	}
	if err := m.WriteJSON(manifestPath); err != nil {
		t.Fatal(err)
	}

	_, err := VerifyManifest(manifestPath, tmpDir)
	if err == nil {
		t.Fatal("VerifyManifest accepted an artifact path containing a NUL byte")
	}
	if !strings.Contains(err.Error(), "contains a NUL byte") {
		t.Fatalf("VerifyManifest error = %q, want a distinct NUL-byte error", err)
	}
}

func TestVerifyManifest_RejectsDuplicateAndSymlinkEntries(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "artifact.bin")
	if err := os.WriteFile(filePath, []byte("content"), 0644); err != nil {
		t.Fatal(err)
	}
	hash, err := fileHash(filePath)
	if err != nil {
		t.Fatal(err)
	}
	entry := ManifestEntry{Name: "artifact.bin", Size: 7, SHA256: hash}
	m := &BuildManifest{Files: []ManifestEntry{entry, entry}}
	manifestPath := filepath.Join(tmpDir, "manifest.json")
	if werr := m.WriteJSON(manifestPath); werr != nil {
		t.Fatal(werr)
	}
	if _, werr := VerifyManifest(manifestPath, tmpDir); werr == nil {
		t.Fatal("VerifyManifest should reject duplicate entries")
	}

	linkPath := filepath.Join(tmpDir, "artifact-link.bin")
	if werr := os.Symlink(filePath, linkPath); werr != nil {
		t.Skipf("symlink unavailable: %v", werr)
	}
	m.Files = []ManifestEntry{{Name: "artifact-link.bin", Size: 7, SHA256: hash}}
	if werr := m.WriteJSON(manifestPath); werr != nil {
		t.Fatal(werr)
	}
	verification, err := VerifyManifest(manifestPath, tmpDir)
	if err != nil {
		t.Fatal(err)
	}
	if verification.OK() || len(verification.Failed) != 1 {
		t.Fatalf("symlink verification = %+v, want one failure", verification)
	}

	outsideDir := t.TempDir()
	outsideFile := filepath.Join(outsideDir, "outside.bin")
	if werr := os.WriteFile(outsideFile, []byte("content"), 0644); werr != nil {
		t.Fatal(werr)
	}
	if werr := os.Symlink(outsideDir, filepath.Join(tmpDir, "linked-dir")); werr != nil {
		t.Skipf("directory symlink unavailable: %v", werr)
	}
	m.Files = []ManifestEntry{{Name: "linked-dir/outside.bin", Size: 7, SHA256: hash}}
	if werr := m.WriteJSON(manifestPath); werr != nil {
		t.Fatal(werr)
	}
	verification, err = VerifyManifest(manifestPath, tmpDir)
	if err != nil {
		t.Fatal(err)
	}
	if verification.OK() || len(verification.Failed) != 1 || !strings.Contains(verification.Failed[0], "symlink") {
		t.Fatalf("intermediate symlink verification = %+v, want symlink failure", verification)
	}
}

func TestWriteBuildManifest_PopulatesMetadata(t *testing.T) {
	tmpDir := t.TempDir()
	cs := pci.NewConfigSpace()
	cs.WriteU16(0x00, 0x8086)
	cs.WriteU16(0x02, 0x15b7)
	ctx := &donor.DeviceContext{
		ToolVersion: "v1.2.3",
		Device: pci.PCIDevice{
			BDF: pci.BDF{Domain: 0, Bus: 3, Device: 0, Function: 1},
		},
		ConfigSpace: cs,
	}
	b := &board.Board{Name: "CaptainDMA_35T"}
	if err := WriteBuildManifest(tmpDir, ctx, b); err != nil {
		t.Fatalf("WriteBuildManifest error: %v", err)
	}
	m, err := LoadManifest(filepath.Join(tmpDir, "build_manifest.json"))
	if err != nil {
		t.Fatal(err)
	}
	if m.Board != b.Name || m.DeviceBDF != "0000:03:00.1" || m.ToolVersion != "v1.2.3" {
		t.Fatalf("manifest metadata = board %q, BDF %q, version %q", m.Board, m.DeviceBDF, m.ToolVersion)
	}
	if m.VendorID != 0x8086 || m.DeviceID != 0x15b7 {
		t.Fatalf("manifest IDs = %04x:%04x", m.VendorID, m.DeviceID)
	}
}

func TestGeneratePreSynthesisManifest_ExcludesStaleDeliverables(t *testing.T) {
	tmpDir := t.TempDir()
	for _, name := range []string{"stale.bit", "stale.bin"} {
		if err := os.WriteFile(filepath.Join(tmpDir, name), []byte("stale"), 0644); err != nil {
			t.Fatal(err)
		}
	}
	m, err := generateManifest(tmpDir, "dev", "board", 0, 0, false)
	if err != nil {
		t.Fatal(err)
	}
	for _, entry := range m.Files {
		if strings.HasSuffix(entry.Name, ".bit") || strings.HasSuffix(entry.Name, ".bin") {
			t.Fatalf("pre-synthesis manifest included stale deliverable %q", entry.Name)
		}
	}
}
