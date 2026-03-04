package pci

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseHex4(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"0000", 0},
		{"FFFF", 0xFFFF},
		{"8086", 0x8086},
		{"1533", 0x1533},
		{"abcd", 0xABCD},
		{"AbCd", 0xABCD},
	}
	for _, tt := range tests {
		got := parseHex4(tt.input)
		if got != tt.want {
			t.Errorf("parseHex4(%q) = %d, want %d", tt.input, got, tt.want)
		}
	}
}

func TestParseHex4Invalid(t *testing.T) {
	invalids := []string{"", "123", "12345", "GHIJ", "----", "zzzz"}
	for _, s := range invalids {
		if got := parseHex4(s); got != -1 {
			t.Errorf("parseHex4(%q) = %d, want -1", s, got)
		}
	}
}

func TestParsePCIIDs(t *testing.T) {
	content := `# PCI IDs test file
8086  Intel Corporation
	1533  I210 Gigabit Network Connection
	10d3  82574L Gigabit Network Connection
1022  Advanced Micro Devices, Inc. [AMD]
	1480  Starship/Matisse Root Complex
`
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "pci.ids")
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	db, err := parsePCIIDs(path)
	if err != nil {
		t.Fatal(err)
	}

	// Vendor checks
	if name := db.VendorName(0x8086); name == "" {
		t.Error("VendorName(0x8086) returned empty")
	}
	if name := db.VendorName(0x1022); name == "" {
		t.Error("VendorName(0x1022) returned empty")
	}
	if name := db.VendorName(0x9999); name != "" {
		t.Errorf("VendorName(0x9999) = %q, want empty", name)
	}

	// Device checks
	if name := db.DeviceName(0x8086, 0x1533); name == "" {
		t.Error("DeviceName(0x8086, 0x1533) returned empty")
	}
	if name := db.DeviceName(0x1022, 0x1480); name == "" {
		t.Error("DeviceName(0x1022, 0x1480) returned empty")
	}
	if name := db.DeviceName(0x8086, 0x9999); name != "" {
		t.Errorf("DeviceName(0x8086, 0x9999) = %q, want empty", name)
	}
}

func TestParsePCIIDsMissingFile(t *testing.T) {
	_, err := parsePCIIDs("/nonexistent/path/pci.ids")
	if err == nil {
		t.Error("parsePCIIDs should fail for missing file")
	}
}

func TestParsePCIIDsClassSection(t *testing.T) {
	// Parser should stop at class definitions
	content := `8086  Intel Corporation
	1533  I210 Gigabit Network Connection
C 00  Unclassified device
	00  Non-VGA unclassified device
`
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "pci.ids")
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	db, err := parsePCIIDs(path)
	if err != nil {
		t.Fatal(err)
	}

	if len(db.Vendors) != 1 {
		t.Errorf("Expected 1 vendor, got %d", len(db.Vendors))
	}
}

func TestLoadPCIDB(t *testing.T) {
	// LoadPCIDB should always return a non-nil db (even when no pci.ids exists)
	db := LoadPCIDB()
	if db == nil {
		t.Fatal("LoadPCIDB returned nil")
	}
	if db.Vendors == nil {
		t.Error("Vendors map is nil")
	}
	if db.Devices == nil {
		t.Error("Devices map is nil")
	}
}
