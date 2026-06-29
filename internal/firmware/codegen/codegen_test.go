package codegen

import (
	"strings"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/pci"
)

func TestGenerateConfigSpaceCOE(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	cs.WriteU16(0x00, 0x8086)
	cs.WriteU16(0x02, 0x1533)

	coe := GenerateConfigSpaceCOE(cs)

	if !strings.HasPrefix(coe, ";") {
		t.Error("COE should start with header comment")
	}
	if !strings.Contains(coe, "memory_initialization_radix=16") {
		t.Error("COE should contain radix declaration")
	}
	// word 0 = VendorID:DeviceID in LE -> 0x15338086
	if !strings.Contains(coe, "15338086") {
		t.Errorf("COE should contain device identity word, got first lines:\n%s", coe[:200])
	}
}

func TestGenerateBarContentCOE_Empty(t *testing.T) {
	coe := GenerateBarContentCOE(nil, 4096)
	if !strings.Contains(coe, "Zero-filled") {
		t.Error("empty BAR should produce zero-filled header")
	}
}

func TestGenerateBarContentCOE_WithData(t *testing.T) {
	barContents := map[int][]byte{
		0: {0x17, 0xFF, 0x40, 0x00, 0x01, 0x00, 0x00, 0x00},
	}
	coe := GenerateBarContentCOE(barContents, 4096)
	if !strings.Contains(coe, "Populated from donor") {
		t.Error("populated BAR should mention donor")
	}
	if !strings.Contains(coe, "0040ff17") {
		t.Error("COE should contain LE32 of first BAR word")
	}
}

func TestGenerateConfigSpaceHex(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	cs.WriteU16(0x00, 0x8086)
	cs.WriteU16(0x02, 0x1533)

	hex := GenerateConfigSpaceHex(cs)

	if !strings.Contains(hex, "15338086") {
		t.Error("HEX should contain device identity")
	}
	lines := strings.Split(strings.TrimSpace(hex), "\n")
	// 2 header lines + 1024 data lines
	if len(lines) != 1026 {
		t.Errorf("HEX should have 1026 lines, got %d", len(lines))
	}
}

func TestGenerateMSIXTableHex(t *testing.T) {
	entries := []pci.MSIXEntry{
		{AddrLo: 0xFEE00000, AddrHi: 0x00000000, Data: 0x00004021, Control: 0x00},
		{AddrLo: 0xFEE00000, AddrHi: 0x00000000, Data: 0x00004022, Control: 0x01},
	}

	hex := GenerateMSIXTableHex(entries)

	if !strings.Contains(hex, "MSI-X Table Init") {
		t.Error("HEX should contain MSI-X header")
	}
	if !strings.Contains(hex, "2 vectors") {
		t.Error("HEX should mention vector count")
	}
	// Each entry = 4 DWORDs, with Control masked (|0x01)
	if !strings.Contains(hex, "FEE00000") {
		t.Error("HEX should contain addr_lo")
	}
	if !strings.Contains(hex, "00004021") {
		t.Error("HEX should contain data for first vector")
	}
	// Control for entry 0: 0x00 | 0x01 = 0x01
	if !strings.Contains(hex, "00000001") {
		t.Error("HEX should contain masked control (0x01)")
	}

	lines := strings.Split(strings.TrimSpace(hex), "\n")
	// 2 header lines + 2*4 data lines = 10
	if len(lines) != 10 {
		t.Errorf("Expected 10 lines, got %d", len(lines))
	}
}

func TestGenerateMSIXTableHex_Empty(t *testing.T) {
	hex := GenerateMSIXTableHex(nil)
	if !strings.Contains(hex, "0 vectors") {
		t.Error("Empty MSI-X should show 0 vectors")
	}
}

// config space COE must contain correct DWORDs for a scrubbed config space.
func TestConfigSpaceCOE_ScrubbedDeviceIdentity(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	cs.WriteU16(0x00, 0x1102)     // VID
	cs.WriteU16(0x02, 0x0012)     // DID
	cs.WriteU32(0x08, 0x04030000) // ClassCode 04:03:00

	coe := GenerateConfigSpaceCOE(cs)

	// DWORD 0 = 0x00121102 (LE: VID=1102, DID=0012)
	if !strings.Contains(coe, "00121102") {
		t.Error("COE should contain VID:DID DWORD 00121102")
	}
	// DWORD 2 = ClassCode bytes
	if !strings.Contains(coe, "04030000") {
		t.Error("COE should contain ClassCode DWORD")
	}
}

// config space HEX format verification.
func TestConfigSpaceHex_DWORDFormat(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	cs.WriteU32(0x10, 0xFFFFF000) // BAR0

	hex := GenerateConfigSpaceHex(cs)
	// offset [010] line should have FFFFF000
	if !strings.Contains(hex, "FFFFF000 // [010]") {
		t.Error("HEX should contain BAR0 at offset [010]")
	}
}

// BAR snapshot init hex: donor values seeded, sized to BAR words, zero-padded.
func TestGenerateBarInitHex(t *testing.T) {
	barData := []byte{
		0x78, 0x56, 0x34, 0x12, // word0 = 0x12345678
		0x21, 0x43, 0x65, 0x87, // word1 = 0x87654321
	}
	hex := GenerateBarInitHex(barData, 4096)

	lines := strings.Split(strings.TrimSpace(hex), "\n")
	var words []string
	for _, l := range lines {
		if !strings.HasPrefix(l, "//") {
			words = append(words, l)
		}
	}
	if len(words) != 1024 {
		t.Fatalf("expected 4096/4=1024 words, got %d", len(words))
	}
	if words[0] != "12345678" {
		t.Errorf("word0 = %s, want 12345678", words[0])
	}
	if words[1] != "87654321" {
		t.Errorf("word1 = %s, want 87654321", words[1])
	}
	if words[2] != "00000000" {
		t.Errorf("word2 (past snapshot) = %s, want 00000000 (zero-padded)", words[2])
	}
}

func TestGenerateOptionROMHex(t *testing.T) {
	rom := []byte{0x55, 0xAA, 0x00, 0x00, 0x11, 0x22, 0x33, 0x44}
	var words []string
	for _, l := range strings.Split(strings.TrimSpace(GenerateOptionROMHex(rom, 2048)), "\n") {
		if !strings.HasPrefix(l, "//") {
			words = append(words, l)
		}
	}
	if len(words) != 512 {
		t.Fatalf("want 512 words, got %d", len(words))
	}
	if words[0] != "0000AA55" {
		t.Errorf("word0 = %s, want 0000AA55", words[0])
	}
	if words[1] != "44332211" {
		t.Errorf("word1 = %s, want 44332211", words[1])
	}
	if words[2] != "00000000" {
		t.Errorf("word2 = %s, want 00000000 (zero-padded)", words[2])
	}
}
