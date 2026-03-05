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
	// word 0 = VendorID:DeviceID in LE → 0x15338086
	if !strings.Contains(coe, "15338086") {
		t.Errorf("COE should contain device identity word, got first lines:\n%s", coe[:200])
	}
}

func TestGenerateWritemaskCOE(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	cs.WriteU16(0x00, 0x8086)
	cs.WriteU32(0x10, 0xFFFFF004) // BAR0 memory

	wm := GenerateWritemaskCOE(cs)

	if !strings.Contains(wm, "memory_initialization_radix=16") {
		t.Error("writemask COE should contain radix")
	}
	// Command register mask at word 1 (offset 0x04)
	lines := strings.Split(wm, "\n")
	found := false
	for _, l := range lines {
		if strings.Contains(l, "0000ffff") {
			found = true
			break
		}
	}
	if !found {
		t.Error("writemask should contain Command register mask 0x0000ffff")
	}
}

func TestGenerateWritemaskCOE_MSI64BitMasking(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	cs.WriteU16(0x06, 0x0010)
	cs.WriteU8(0x34, 0x50)

	// MSI at 0x50: 64-bit + per-vector masking
	cs.WriteU8(0x50, pci.CapIDMSI)
	cs.WriteU8(0x51, 0x00)
	cs.WriteU16(0x52, 0x0180) // 64bit + masking

	masks := make([]uint32, shadowCfgSpaceWords)
	applyCapabilityWritemasks(cs, masks)

	// Message Control (DWORD at 0x50): Enable + MultiMsg bits
	if masks[0x50/4]&0x00710000 == 0 {
		t.Error("MSI msg control writable bits missing")
	}
	// Addr Low (0x54): bits [31:2]
	if masks[0x54/4] != 0xFFFFFFFC {
		t.Errorf("MSI addr low mask: got 0x%08x, want 0xFFFFFFFC", masks[0x54/4])
	}
	// Addr High (0x58): fully writable
	if masks[0x58/4] != 0xFFFFFFFF {
		t.Errorf("MSI addr high mask: got 0x%08x, want 0xFFFFFFFF", masks[0x58/4])
	}
	// Data (0x5C): lower 16 bits
	if masks[0x5C/4] != 0x0000FFFF {
		t.Errorf("MSI data mask: got 0x%08x, want 0x0000FFFF", masks[0x5C/4])
	}
	// Mask Bits (0x60): fully writable
	if masks[0x60/4] != 0xFFFFFFFF {
		t.Errorf("MSI mask bits mask: got 0x%08x, want 0xFFFFFFFF", masks[0x60/4])
	}
}

func TestGenerateWritemaskCOE_MSI32BitNoMask(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	cs.WriteU16(0x06, 0x0010)
	cs.WriteU8(0x34, 0x50)

	// MSI at 0x50: 32-bit, no masking
	cs.WriteU8(0x50, pci.CapIDMSI)
	cs.WriteU8(0x51, 0x00)
	cs.WriteU16(0x52, 0x0000)

	masks := make([]uint32, shadowCfgSpaceWords)
	applyCapabilityWritemasks(cs, masks)

	// Addr Low (0x54)
	if masks[0x54/4] != 0xFFFFFFFC {
		t.Errorf("MSI addr low mask: got 0x%08x, want 0xFFFFFFFC", masks[0x54/4])
	}
	// Data at 0x58 (32-bit layout)
	if masks[0x58/4] != 0x0000FFFF {
		t.Errorf("MSI data mask (32-bit): got 0x%08x, want 0x0000FFFF", masks[0x58/4])
	}
	// 0x5C should be 0 (no addr_hi, no mask bits)
	if masks[0x5C/4] != 0 {
		t.Errorf("no mask bits expected at 0x5C for 32-bit no-mask MSI, got 0x%08x", masks[0x5C/4])
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
