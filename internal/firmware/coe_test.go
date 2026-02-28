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
