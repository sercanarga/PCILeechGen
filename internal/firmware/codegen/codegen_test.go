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

func TestGenerateWritemaskCOE(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	cs.WriteU16(0x00, 0x8086)
	cs.WriteU32(0x10, 0xFFFFF000) // BAR0: 4KB 32-bit memory BAR (scrubbed)

	wm := GenerateWritemaskCOE(cs)

	if !strings.Contains(wm, "memory_initialization_radix=16") {
		t.Error("writemask COE should contain radix")
	}

	// parse data lines into DWORD array
	lines := strings.Split(wm, "\n")
	var dwords []string
	for _, l := range lines {
		l = strings.TrimSpace(l)
		if l == "" || strings.HasPrefix(l, ";") || strings.HasPrefix(l, "memory") {
			continue
		}
		l = strings.TrimSuffix(l, ",")
		l = strings.TrimSuffix(l, ";")
		dwords = append(dwords, l)
	}

	if len(dwords) != 1024 {
		t.Fatalf("writemask should have 1024 DWORDs, got %d", len(dwords))
	}

	// DWORD 0 (VID:DID) must be read-only
	if dwords[0] != "00000000" {
		t.Errorf("VID:DID mask should be 00000000, got %s", dwords[0])
	}

	// DWORD 2 (Rev:ClassCode) must be read-only
	if dwords[2] != "00000000" {
		t.Errorf("Rev:ClassCode mask should be 00000000, got %s", dwords[2])
	}

	// DWORD 4 (BAR0) must match scrubbed BAR size mask (0xFFFFF000)
	if dwords[4] != "fffff000" {
		t.Errorf("BAR0 mask should be fffff000 (4KB size), got %s", dwords[4])
	}

	// DWORD 5 (BAR1) must be 0 (unused after scrubber clamp)
	if dwords[5] != "00000000" {
		t.Errorf("BAR1 mask should be 00000000 (unused), got %s", dwords[5])
	}

	// DWORD 11 (SubsysIDs) must be read-only
	if dwords[11] != "00000000" {
		t.Errorf("SubsysIDs mask should be 00000000, got %s", dwords[11])
	}

	// DWORD 15 (IntLine/IntPin) - only IntLine writable
	if dwords[15] != "000000ff" {
		t.Errorf("IntLine/IntPin mask should be 000000ff, got %s", dwords[15])
	}

	// capability region (0x40-0xFF) must be fully writable.
	// the writemask only controls shadow BRAM; the Xilinx IP core
	// processes config writes independently. locking here would create
	// a dangerous BRAM/IP-core state mismatch.
	for i := 0x40 / 4; i < 0x100/4; i++ {
		if dwords[i] != "ffffffff" {
			t.Errorf("cap region DWORD[%d] should be ffffffff, got %s", i, dwords[i])
		}
	}

	// extended config space (0x100+) should be read-only
	for i := 0x100 / 4; i < len(dwords); i++ {
		if dwords[i] != "00000000" {
			t.Errorf("ext cap DWORD[%d] should be 00000000, got %s", i, dwords[i])
			break
		}
	}
}

func TestGenerateBarContentCOE_Empty(t *testing.T) {
	coe := GenerateBarContentCOE(nil)
	if !strings.Contains(coe, "Zero-filled") {
		t.Error("empty BAR should produce zero-filled header")
	}
}

func TestGenerateBarContentCOE_WithData(t *testing.T) {
	barContents := map[int][]byte{
		0: {0x17, 0xFF, 0x40, 0x00, 0x01, 0x00, 0x00, 0x00},
	}
	coe := GenerateBarContentCOE(barContents)
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
