package firmware

import (
	"strings"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/pci"
)

func makeTestConfigSpace() *pci.ConfigSpace {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize

	cs.WriteU16(0x00, 0x8086)     // Vendor ID
	cs.WriteU16(0x02, 0x1533)     // Device ID
	cs.WriteU16(0x04, 0x0406)     // Command
	cs.WriteU16(0x06, 0x0010)     // Status
	cs.WriteU8(0x08, 0x03)        // Revision
	cs.WriteU8(0x0B, 0x02)        // Base class (Network)
	cs.WriteU32(0x10, 0xFE000000) // BAR0
	cs.WriteU16(0x2C, 0x8086)     // Subsys Vendor
	cs.WriteU16(0x2E, 0x0001)     // Subsys Device
	cs.WriteU8(0x34, 0x40)        // Cap pointer

	// PM cap at 0x40
	cs.WriteU8(0x40, pci.CapIDPowerManagement)
	cs.WriteU8(0x41, 0x50) // next

	// MSI-X cap at 0x50
	cs.WriteU8(0x50, pci.CapIDMSIX)
	cs.WriteU8(0x51, 0x70) // next

	// PCIe cap at 0x70
	cs.WriteU8(0x70, pci.CapIDPCIExpress)
	cs.WriteU8(0x71, 0x00) // end

	return cs
}

func TestGenerateConfigSpaceCOE(t *testing.T) {
	cs := makeTestConfigSpace()
	coe := GenerateConfigSpaceCOE(cs)

	// Check header
	if !strings.Contains(coe, "memory_initialization_radix=16;") {
		t.Error("COE missing radix declaration")
	}
	if !strings.Contains(coe, "memory_initialization_vector=") {
		t.Error("COE missing vector declaration")
	}

	// Check it contains the vendor/device ID word
	if !strings.Contains(coe, "15338086") {
		t.Error("COE missing vendor/device ID word")
	}

	// Check it ends with semicolon
	lines := strings.Split(strings.TrimSpace(coe), "\n")
	lastLine := lines[len(lines)-1]
	if !strings.HasSuffix(lastLine, ";") {
		t.Errorf("Last line should end with ';', got: %q", lastLine)
	}

	// Count data words (should be 1024 for 4KB)
	dataLines := 0
	for _, line := range lines {
		if strings.Contains(line, ",") || strings.HasSuffix(line, ";") {
			if !strings.HasPrefix(line, ";") && !strings.Contains(line, "memory_") {
				dataLines++
			}
		}
	}
	if dataLines != 1024 {
		t.Errorf("Expected 1024 data words, got %d", dataLines)
	}
}

func TestGenerateWritemaskCOE(t *testing.T) {
	cs := makeTestConfigSpace()
	coe := GenerateWritemaskCOE(cs)

	if !strings.Contains(coe, "memory_initialization_radix=16;") {
		t.Error("Writemask COE missing radix declaration")
	}

	// Verify the output ends properly
	if !strings.HasSuffix(strings.TrimSpace(coe), ";") {
		t.Error("Writemask COE should end with semicolon")
	}
}

func TestGenerateBarZeroCOE(t *testing.T) {
	// Deprecated wrapper should still produce valid zero-filled COE
	coe := GenerateBarZeroCOE()

	if !strings.Contains(coe, "memory_initialization_radix=16;") {
		t.Error("BAR zero COE missing radix")
	}

	// Count data words (should be 1024 for 4KB)
	count := strings.Count(coe, "00000000")
	if count != 1024 {
		t.Errorf("Expected 1024 zero words, got %d", count)
	}
}

func TestGenerateBarContentCOE_WithData(t *testing.T) {
	// Simulate donor BAR data: first 8 bytes are non-zero (like HDA GCAP + VMIN/VMAJ)
	barData := make([]byte, 256)
	barData[0] = 0x59 // GCAP low byte
	barData[1] = 0x00 // GCAP high byte
	barData[2] = 0x01 // VMIN
	barData[3] = 0x01 // VMAJ
	barData[4] = 0x00 // OUTPAY low
	barData[5] = 0x3C // OUTPAY high
	barData[6] = 0x00 // INPAY low
	barData[7] = 0x1C // INPAY high

	contents := map[int][]byte{0: barData}
	coe := GenerateBarContentCOE(contents)

	// Should contain donor data header
	if !strings.Contains(coe, "donor device BAR memory") {
		t.Error("COE should mention donor BAR data in header")
	}

	// First word should be 0x01010059 (GCAP + VMIN/VMAJ in little-endian)
	if !strings.Contains(coe, "01010059") {
		t.Error("First BAR word should contain GCAP data (01010059)")
	}

	// Second word should be 0x1C003C00 (OUTPAY + INPAY)
	if !strings.Contains(coe, "1c003c00") {
		t.Error("Second BAR word should contain payload data (1c003c00)")
	}

	// Should still have 1024 total words
	lines := strings.Split(strings.TrimSpace(coe), "\n")
	dataLines := 0
	for _, line := range lines {
		if strings.Contains(line, ",") || strings.HasSuffix(line, ";") {
			if !strings.HasPrefix(line, ";") && !strings.Contains(line, "memory_") {
				dataLines++
			}
		}
	}
	if dataLines != 1024 {
		t.Errorf("Expected 1024 data words, got %d", dataLines)
	}
}

func TestGenerateBarContentCOE_Empty(t *testing.T) {
	// No BAR contents — should produce zero-filled with appropriate header
	coe := GenerateBarContentCOE(nil)

	if !strings.Contains(coe, "Zero-filled") {
		t.Error("Empty BAR COE should note zero-filled in header")
	}

	count := strings.Count(coe, "00000000")
	if count != 1024 {
		t.Errorf("Expected 1024 zero words, got %d", count)
	}
}

func TestGenerateBarContentCOE_MultipleBAR(t *testing.T) {
	// Multiple BARs — should use lowest index
	bar0 := make([]byte, 8)
	bar0[0] = 0xAA
	bar2 := make([]byte, 8)
	bar2[0] = 0xBB

	contents := map[int][]byte{0: bar0, 2: bar2}
	coe := GenerateBarContentCOE(contents)

	// Should use BAR0 (lowest index) — first word should contain 0xAA
	if !strings.Contains(coe, "000000aa") {
		t.Error("Should use lowest BAR index (BAR0 with 0xAA)")
	}
}

func TestPMWritemask(t *testing.T) {
	cs := makeTestConfigSpace()
	coe := GenerateWritemaskCOE(cs)

	// PM cap is at 0x40, PMCSR at 0x44 (word index 0x44/4 = 0x11 = 17)
	// Expected writemask: 0x00008103
	lines := strings.Split(strings.TrimSpace(coe), "\n")
	dataIdx := 0
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, ";") ||
			strings.Contains(line, "memory_") {
			continue
		}
		if dataIdx == 0x44/4 {
			cleaned := strings.TrimSuffix(strings.TrimSuffix(line, ","), ";")
			if cleaned != "00008103" {
				t.Errorf("PM writemask at word %d = %s, want 00008103", dataIdx, cleaned)
			}
			return
		}
		dataIdx++
	}
	t.Error("Could not find PM writemask word")
}

func TestScrubPMNoSoftReset(t *testing.T) {
	cs := makeTestConfigSpace()
	// Set some PM state (D3hot, PME_Status)
	cs.WriteU16(0x44, 0x8003) // D3hot + PME_Status set

	scrubbed := ScrubConfigSpace(cs)
	pmcsr := scrubbed.ReadU16(0x44)

	// Should be D0 (bits 1:0 = 00), NoSoftReset set (bit 3), PME_Status cleared (bit 15)
	if pmcsr&0x0003 != 0 {
		t.Errorf("PowerState should be D0 (0), got %d", pmcsr&0x0003)
	}
	if pmcsr&0x0008 == 0 {
		t.Error("NoSoftReset bit should be set")
	}
	if pmcsr&0x8000 != 0 {
		t.Error("PME_Status should be cleared")
	}
}

func TestCOEFormatValidity(t *testing.T) {
	cs := makeTestConfigSpace()
	coe := GenerateConfigSpaceCOE(cs)

	// Parse and validate COE format
	lines := strings.Split(coe, "\n")
	foundRadix := false
	foundVector := false
	dataCount := 0

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, ";") {
			continue
		}
		if strings.Contains(line, "memory_initialization_radix") {
			foundRadix = true
			continue
		}
		if strings.Contains(line, "memory_initialization_vector") {
			foundVector = true
			continue
		}

		// Should be a hex value followed by comma or semicolon
		line = strings.TrimSuffix(line, ",")
		line = strings.TrimSuffix(line, ";")
		if len(line) != 8 {
			t.Errorf("Data word should be 8 hex chars, got %q (%d)", line, len(line))
		}
		dataCount++
	}

	if !foundRadix {
		t.Error("Missing radix declaration")
	}
	if !foundVector {
		t.Error("Missing vector declaration")
	}
	if dataCount == 0 {
		t.Error("No data words found")
	}
}
