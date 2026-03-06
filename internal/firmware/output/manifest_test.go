package output

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
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
	os.WriteFile(filepath.Join(tmpDir, "pcileech_cfgspace.coe"), []byte("test content"), 0644)
	os.WriteFile(filepath.Join(tmpDir, "device_context.json"), []byte("{}"), 0644)

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
	os.MkdirAll(srcDir, 0755)
	os.WriteFile(filepath.Join(srcDir, "test.sv"), []byte("module test; endmodule"), 0644)

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
	os.WriteFile(path, []byte("hello"), 0644)

	hash, err := fileHash(path)
	if err != nil {
		t.Fatalf("fileHash error: %v", err)
	}
	if len(hash) != 64 {
		t.Errorf("SHA256 hash length = %d, want 64", len(hash))
	}

	// Same content → same hash
	path2 := filepath.Join(tmpDir, "test2.txt")
	os.WriteFile(path2, []byte("hello"), 0644)
	hash2, _ := fileHash(path2)
	if hash != hash2 {
		t.Error("Same content should produce same hash")
	}

	// Non-existent file → error
	_, err = fileHash("/nonexistent")
	if err == nil {
		t.Error("fileHash should fail for non-existent file")
	}
}
