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

func TestGenerateBarContentCOE_ZeroFilled(t *testing.T) {
	coe := GenerateBarContentCOE(nil)

	if !strings.Contains(coe, "memory_initialization_radix=16;") {
		t.Error("BAR zero COE missing radix")
	}

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

	scrubbed := ScrubConfigSpace(cs, nil)
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

func TestScrubBarContent_NVMe(t *testing.T) {
	// NVMe class code 01:08:02
	classCode := uint32(0x010802)

	// Create a BAR0 with NVMe registers
	bar0 := make([]byte, 4096)
	// CC at offset 0x14: EN=0 (disabled)
	bar0[0x14] = 0x00
	// CSTS at offset 0x1C: RDY=0 (not ready), CFS=1 (fatal), SHST=10 (shutdown)
	bar0[0x1C] = 0x0A // CFS=1, SHST=10 → binary 00001010

	contents := map[int][]byte{0: bar0}
	ScrubBarContent(contents, classCode)

	// CSTS.RDY should now be 1
	csts := uint32(bar0[0x1C]) | uint32(bar0[0x1D])<<8 |
		uint32(bar0[0x1E])<<16 | uint32(bar0[0x1F])<<24
	if csts&0x01 != 1 {
		t.Errorf("NVMe CSTS.RDY should be 1, got %d", csts&0x01)
	}
	// CSTS.CFS should be cleared
	if csts&0x02 != 0 {
		t.Errorf("NVMe CSTS.CFS should be 0, got %d", (csts>>1)&0x01)
	}
	// CSTS.SHST should be 00
	if csts&0x0C != 0 {
		t.Errorf("NVMe CSTS.SHST should be 00, got %d", (csts>>2)&0x03)
	}

	// CC.EN should be 1 (coherent with CSTS.RDY=1)
	cc := uint32(bar0[0x14]) | uint32(bar0[0x15])<<8 |
		uint32(bar0[0x16])<<16 | uint32(bar0[0x17])<<24
	if cc&0x01 != 1 {
		t.Errorf("NVMe CC.EN should be 1, got %d", cc&0x01)
	}
}

func TestScrubBarContent_NonNVMe(t *testing.T) {
	// Sound card class code (04:03:00 = Audio Device)
	classCode := uint32(0x040300)

	bar0 := make([]byte, 256)
	bar0[0x1C] = 0x00 // some register
	bar0[0x14] = 0x42 // some register

	original1C := bar0[0x1C]
	original14 := bar0[0x14]

	contents := map[int][]byte{0: bar0}
	ScrubBarContent(contents, classCode)

	// Nothing should change for non-NVMe devices
	if bar0[0x1C] != original1C {
		t.Errorf("Non-NVMe BAR0[0x1C] should not change: got 0x%02x, want 0x%02x", bar0[0x1C], original1C)
	}
	if bar0[0x14] != original14 {
		t.Errorf("Non-NVMe BAR0[0x14] should not change: got 0x%02x, want 0x%02x", bar0[0x14], original14)
	}
}

func TestScrubBarContent_Empty(t *testing.T) {
	// Empty contents — should not panic
	ScrubBarContent(nil, 0x010802)
	ScrubBarContent(map[int][]byte{}, 0x010802)
}

func TestScrubBarContent_SmallBAR(t *testing.T) {
	// BAR smaller than NVMe register area — should not panic
	classCode := uint32(0x010802)
	bar0 := make([]byte, 16) // too small for CSTS at 0x1C
	contents := map[int][]byte{0: bar0}
	ScrubBarContent(contents, classCode)
	// Should return without modifying anything (no panic)
}

func TestScrubBarContent_NVMe_HigherBARIndex(t *testing.T) {
	// BAR content at index 2 instead of 0
	classCode := uint32(0x010802)
	bar2 := make([]byte, 4096)
	bar2[0x1C] = 0x00 // CSTS.RDY=0
	bar2[0x14] = 0x00 // CC.EN=0

	contents := map[int][]byte{2: bar2}
	ScrubBarContent(contents, classCode)

	if bar2[0x1C]&0x01 != 1 {
		t.Error("CSTS.RDY should be set even when BAR is at index 2")
	}
	if bar2[0x14]&0x01 != 1 {
		t.Error("CC.EN should be set even when BAR is at index 2")
	}
}

func TestExtCapWritemask_AER(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize

	cs.WriteU16(0x06, 0x0010)

	// AER at 0x100
	cs.WriteU32(0x100, makeExtCapHeader(pci.ExtCapIDAER, 1, 0))

	masks := make([]uint32, 1024)
	applyExtCapabilityWritemasks(cs, masks)

	// AER uncorrectable error status (cap+4 = word 0x104/4 = 65)
	if masks[0x104/4] != 0xFFFFFFFF {
		t.Errorf("AER uncorrectable error status writemask: got 0x%08x, want 0xFFFFFFFF", masks[0x104/4])
	}
	// AER uncorrectable error mask (cap+8 = word 66)
	if masks[0x108/4] != 0xFFFFFFFF {
		t.Errorf("AER uncorrectable error mask writemask: got 0x%08x, want 0xFFFFFFFF", masks[0x108/4])
	}
	// AER uncorrectable error severity (cap+12 = word 67)
	if masks[0x10C/4] != 0xFFFFFFFF {
		t.Errorf("AER uncorrectable error severity writemask: got 0x%08x, want 0xFFFFFFFF", masks[0x10C/4])
	}
	// AER correctable error status (cap+16 = word 68)
	if masks[0x110/4] != 0xFFFFFFFF {
		t.Errorf("AER correctable error status writemask: got 0x%08x, want 0xFFFFFFFF", masks[0x110/4])
	}
	// AER correctable error mask (cap+20 = word 69)
	if masks[0x114/4] != 0xFFFFFFFF {
		t.Errorf("AER correctable error mask writemask: got 0x%08x, want 0xFFFFFFFF", masks[0x114/4])
	}
}

func TestExtCapWritemask_LTR(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize

	cs.WriteU16(0x06, 0x0010)

	// LTR at 0x100
	cs.WriteU32(0x100, makeExtCapHeader(pci.ExtCapIDLTR, 1, 0))

	masks := make([]uint32, 1024)
	applyExtCapabilityWritemasks(cs, masks)

	if masks[0x104/4] != 0xFFFFFFFF {
		t.Errorf("LTR writemask at cap+4: got 0x%08x, want 0xFFFFFFFF", masks[0x104/4])
	}
}

func TestExtCapWritemask_NoExtCaps(t *testing.T) {
	// Small config space (256 bytes) — no ext caps
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceLegacySize

	masks := make([]uint32, 1024)
	applyExtCapabilityWritemasks(cs, masks)

	// All masks beyond legacy should be 0
	for i := 64; i < 1024; i++ {
		if masks[i] != 0 {
			t.Errorf("mask[%d] should be 0 for legacy config space, got 0x%08x", i, masks[i])
		}
	}
}

func TestLinkSpeedToTCL_AllCases(t *testing.T) {
	tests := []struct {
		speed uint8
		want  string
	}{
		{LinkSpeedGen1, "2.5_GT/s"},
		{LinkSpeedGen2, "5.0_GT/s"},
		{LinkSpeedGen3, "8.0_GT/s"},
		{4, "5.0_GT/s"},  // Gen4 falls to default
		{0, "5.0_GT/s"},  // unknown
		{99, "5.0_GT/s"}, // garbage
	}
	for _, tt := range tests {
		got := linkSpeedToTCL(tt.speed)
		if got != tt.want {
			t.Errorf("linkSpeedToTCL(%d) = %q, want %q", tt.speed, got, tt.want)
		}
	}
}

func TestLinkSpeedToTrgt_AllCases(t *testing.T) {
	tests := []struct {
		speed uint8
		want  string
	}{
		{LinkSpeedGen1, "4'h1"},
		{LinkSpeedGen2, "4'h2"},
		{LinkSpeedGen3, "4'h3"},
		{4, "4'h2"}, // default
		{0, "4'h2"}, // default
	}
	for _, tt := range tests {
		got := linkSpeedToTrgt(tt.speed)
		if got != tt.want {
			t.Errorf("linkSpeedToTrgt(%d) = %q, want %q", tt.speed, got, tt.want)
		}
	}
}

func TestLinkWidthToTCL_AllCases(t *testing.T) {
	tests := []struct {
		width uint8
		want  string
	}{
		{1, "X1"},
		{2, "X2"},
		{4, "X4"},
		{8, "X8"},
		{0, "X1"},  // default
		{3, "X1"},  // odd value
		{16, "X1"}, // unsupported
	}
	for _, tt := range tests {
		got := linkWidthToTCL(tt.width)
		if got != tt.want {
			t.Errorf("linkWidthToTCL(%d) = %q, want %q", tt.width, got, tt.want)
		}
	}
}

func TestScrubBarContent_XHCI(t *testing.T) {
	// xHCI class code 0C:03:30
	classCode := uint32(0x0C0330)

	// Create a BAR0 with xHCI registers
	bar0 := make([]byte, 4096)
	// CAPLENGTH at 0x00: operational regs start at offset 0x20
	bar0[0x00] = 0x20
	// HCIVERSION at 0x02-0x03
	bar0[0x02] = 0x00
	bar0[0x03] = 0x01 // version 1.0

	// USBCMD at CAPLENGTH+0x00 = 0x20: HCRST=1, Run/Stop=0
	bar0[0x20] = 0x02 // HCRST=1

	// USBSTS at CAPLENGTH+0x04 = 0x24: HCHalted=1, HSE=1
	bar0[0x24] = 0x05 // HCHalted=1, HSE=1

	contents := map[int][]byte{0: bar0}
	ScrubBarContent(contents, classCode)

	// USBCMD: Run/Stop (bit 0) should be 1, HCRST (bit 1) should be cleared
	usbcmd := uint32(bar0[0x20]) | uint32(bar0[0x21])<<8 |
		uint32(bar0[0x22])<<16 | uint32(bar0[0x23])<<24
	if usbcmd&0x01 != 1 {
		t.Errorf("xHCI USBCMD.Run/Stop should be 1, got 0x%08x", usbcmd)
	}
	if usbcmd&0x02 != 0 {
		t.Errorf("xHCI USBCMD.HCRST should be 0, got 0x%08x", usbcmd)
	}

	// USBSTS: HCHalted (bit 0) should be 0, HSE (bit 2) should be 0
	usbsts := uint32(bar0[0x24]) | uint32(bar0[0x25])<<8 |
		uint32(bar0[0x26])<<16 | uint32(bar0[0x27])<<24
	if usbsts&0x01 != 0 {
		t.Errorf("xHCI USBSTS.HCHalted should be 0, got 0x%08x", usbsts)
	}
	if usbsts&0x04 != 0 {
		t.Errorf("xHCI USBSTS.HSE should be 0, got 0x%08x", usbsts)
	}
}

func TestScrubBarContent_XHCI_SmallBAR(t *testing.T) {
	// BAR smaller than register area — should not panic
	classCode := uint32(0x0C0330)
	bar0 := make([]byte, 16) // too small
	contents := map[int][]byte{0: bar0}
	ScrubBarContent(contents, classCode)
	// Should return without panic
}

func TestScrubBarContent_XHCI_ZeroCAPLEN(t *testing.T) {
	// CAPLENGTH=0 — should use safe default (0x20)
	classCode := uint32(0x0C0330)
	bar0 := make([]byte, 4096)
	bar0[0x00] = 0x00 // CAPLENGTH=0 (invalid)

	// USBSTS at default CAPLENGTH (0x20) + 4 = 0x24: HCHalted=1
	bar0[0x24] = 0x01

	contents := map[int][]byte{0: bar0}
	ScrubBarContent(contents, classCode)

	// Should use default CAPLENGTH=0x20
	usbsts := uint32(bar0[0x24])
	if usbsts&0x01 != 0 {
		t.Errorf("xHCI USBSTS.HCHalted should be cleared with default CAPLENGTH, got 0x%02x", usbsts)
	}
}

func TestScrubBarContent_XHCI_ClampDBOFF(t *testing.T) {
	classCode := uint32(0x0C0330)
	bar0 := make([]byte, 4096)

	bar0[0x00] = 0x20 // CAPLENGTH
	// HCSPARAMS1: MaxSlots=32, MaxIntrs=1, MaxPorts=2
	writeLE32(bar0, 0x04, 32|(1<<8)|(2<<24))
	writeLE32(bar0, 0x14, 0x2000) // DBOFF outside BRAM

	contents := map[int][]byte{0: bar0}
	ScrubBarContent(contents, classCode)

	dboff := readLE32(bar0, 0x14)
	maxSlots := int(readLE32(bar0, 0x04) & 0xFF)
	doorbellSize := (maxSlots + 1) * 4
	if int(dboff)+doorbellSize > bramSize {
		t.Errorf("DBOFF (0x%04x) + doorbell array (%d) exceeds BRAM", dboff, doorbellSize)
	}
	if dboff&0x1F != 0 {
		t.Errorf("DBOFF (0x%04x) is not 32-byte aligned", dboff)
	}
}

func TestScrubBarContent_XHCI_ClampRTSOFF(t *testing.T) {
	classCode := uint32(0x0C0330)
	bar0 := make([]byte, 4096)

	bar0[0x00] = 0x20
	writeLE32(bar0, 0x04, 4|(1<<8)|(2<<24))
	writeLE32(bar0, 0x18, 0x3000) // RTSOFF outside BRAM

	contents := map[int][]byte{0: bar0}
	ScrubBarContent(contents, classCode)

	rtsoff := readLE32(bar0, 0x18)
	if int(rtsoff)+0x40 > bramSize {
		t.Errorf("RTSOFF (0x%04x) + runtime regs exceeds BRAM", rtsoff)
	}
	if rtsoff&0x1F != 0 {
		t.Errorf("RTSOFF (0x%04x) is not 32-byte aligned", rtsoff)
	}
}

func TestScrubBarContent_XHCI_AlreadyInBRAM(t *testing.T) {
	classCode := uint32(0x0C0330)
	bar0 := make([]byte, 4096)

	bar0[0x00] = 0x20
	writeLE32(bar0, 0x04, 4|(1<<8)|(2<<24)) // MaxSlots=4, MaxIntrs=1, MaxPorts=2
	writeLE32(bar0, 0x14, 0x0800)           // DBOFF inside BRAM
	writeLE32(bar0, 0x18, 0x0400)           // RTSOFF inside BRAM

	contents := map[int][]byte{0: bar0}
	ScrubBarContent(contents, classCode)

	if readLE32(bar0, 0x14) != 0x0800 {
		t.Errorf("DBOFF should remain 0x0800, got 0x%04x", readLE32(bar0, 0x14))
	}
	if readLE32(bar0, 0x18) != 0x0400 {
		t.Errorf("RTSOFF should remain 0x0400, got 0x%04x", readLE32(bar0, 0x18))
	}
}

func TestScrubBarContent_XHCI_PageSize(t *testing.T) {
	classCode := uint32(0x0C0330)
	bar0 := make([]byte, 4096)

	bar0[0x00] = 0x20
	writeLE32(bar0, 0x04, 4|(1<<8)|(2<<24))
	writeLE32(bar0, 0x28, 0xFF) // PAGESIZE at capLen+0x08

	contents := map[int][]byte{0: bar0}
	ScrubBarContent(contents, classCode)

	if readLE32(bar0, 0x28) != 0x01 {
		t.Errorf("PAGESIZE should be 0x01 (4KB), got 0x%08x", readLE32(bar0, 0x28))
	}
}

func TestScrubBarContent_XHCI_BothOffsetsOutside(t *testing.T) {
	classCode := uint32(0x0C0330)
	bar0 := make([]byte, 4096)

	bar0[0x00] = 0x20
	writeLE32(bar0, 0x04, 16|(1<<8)|(2<<24))
	writeLE32(bar0, 0x14, 0x4000)
	writeLE32(bar0, 0x18, 0x8000)

	contents := map[int][]byte{0: bar0}
	ScrubBarContent(contents, classCode)

	dboff := readLE32(bar0, 0x14)
	rtsoff := readLE32(bar0, 0x18)
	maxSlots := int(readLE32(bar0, 0x04) & 0xFF)

	if int(dboff)+(maxSlots+1)*4 > bramSize {
		t.Errorf("DBOFF (0x%04x) + doorbell array exceeds BRAM", dboff)
	}
	if int(rtsoff)+0x40 > bramSize {
		t.Errorf("RTSOFF (0x%04x) + runtime regs exceeds BRAM", rtsoff)
	}
	if dboff&0x1F != 0 {
		t.Errorf("DBOFF (0x%04x) not 32-byte aligned", dboff)
	}
	if rtsoff&0x1F != 0 {
		t.Errorf("RTSOFF (0x%04x) not 32-byte aligned", rtsoff)
	}
}

func TestScrubBarContent_XHCI_ScratchpadZeroed(t *testing.T) {
	classCode := uint32(0x0C0330)
	bar0 := make([]byte, 4096)

	bar0[0x00] = 0x20
	writeLE32(bar0, 0x04, 4|(1<<8)|(2<<24))
	// HCSPARAMS2: Scratchpad Hi=3, SPR=1, Lo=5, plus some low bits
	writeLE32(bar0, 0x08, (3<<27)|(1<<26)|(5<<21)|0x000F)

	contents := map[int][]byte{0: bar0}
	ScrubBarContent(contents, classCode)

	hcsparams2 := readLE32(bar0, 0x08)
	scratchHi := (hcsparams2 >> 27) & 0x1F
	scratchLo := (hcsparams2 >> 21) & 0x1F
	spr := (hcsparams2 >> 26) & 0x01
	if scratchHi != 0 || scratchLo != 0 || spr != 0 {
		t.Errorf("Scratchpad should be zeroed: Hi=%d Lo=%d SPR=%d", scratchHi, scratchLo, spr)
	}
	if hcsparams2&0x0F != 0x0F {
		t.Errorf("HCSPARAMS2 lower bits should be preserved: 0x%08x", hcsparams2)
	}
}

func TestScrubBarContent_XHCI_XECPClamp(t *testing.T) {
	classCode := uint32(0x0C0330)
	bar0 := make([]byte, 4096)

	bar0[0x00] = 0x20
	writeLE32(bar0, 0x04, 4|(1<<8)|(2<<24))
	// xECP=0x500 DWORDs → byte offset 0x1400, outside BRAM
	writeLE32(bar0, 0x10, (0x500<<16)|0x05)

	contents := map[int][]byte{0: bar0}
	ScrubBarContent(contents, classCode)

	hccparams1 := readLE32(bar0, 0x10)
	xecp := (hccparams1 >> 16) & 0xFFFF
	if xecp != 0 {
		t.Errorf("xECP should be zeroed when outside BRAM, got 0x%04x", xecp)
	}
	if hccparams1&0xFFFF != 0x05 {
		t.Errorf("HCCPARAMS1 lower bits should be preserved: 0x%08x", hccparams1)
	}
}

func TestScrubBarContent_XHCI_XECPInsideBRAM(t *testing.T) {
	classCode := uint32(0x0C0330)
	bar0 := make([]byte, 4096)

	bar0[0x00] = 0x20
	writeLE32(bar0, 0x04, 4|(1<<8)|(2<<24))
	// xECP=0x100 DWORDs → byte offset 0x400, inside BRAM
	writeLE32(bar0, 0x10, (0x100<<16)|0x05)

	contents := map[int][]byte{0: bar0}
	ScrubBarContent(contents, classCode)

	xecp := (readLE32(bar0, 0x10) >> 16) & 0xFFFF
	if xecp != 0x100 {
		t.Errorf("xECP should remain 0x100 when inside BRAM, got 0x%04x", xecp)
	}
}

func TestScrubBarContent_XHCI_ClampMaxIntrs(t *testing.T) {
	classCode := uint32(0x0C0330)
	bar0 := make([]byte, 4096)

	bar0[0x00] = 0x20
	// MaxSlots=4, MaxIntrs=256, MaxPorts=2
	writeLE32(bar0, 0x04, 4|(256<<8)|(2<<24))
	writeLE32(bar0, 0x18, 0x200) // RTSOFF

	contents := map[int][]byte{0: bar0}
	ScrubBarContent(contents, classCode)

	hcsparams1 := readLE32(bar0, 0x04)
	maxIntrs := int((hcsparams1 >> 8) & 0x7FF)
	rtsoff := int(readLE32(bar0, 0x18))

	if rtsoff+0x20+maxIntrs*0x20 > bramSize {
		t.Errorf("MaxIntrs (%d) interrupter sets overflow BRAM from RTSOFF=0x%x", maxIntrs, rtsoff)
	}
	if maxIntrs < 1 {
		t.Error("MaxIntrs should be at least 1")
	}
}

func TestScrubBarContent_XHCI_ClampMaxPorts(t *testing.T) {
	classCode := uint32(0x0C0330)
	bar0 := make([]byte, 4096)

	bar0[0x00] = 0x20
	// MaxSlots=4, MaxIntrs=1, MaxPorts=255
	writeLE32(bar0, 0x04, 4|(1<<8)|(255<<24))

	contents := map[int][]byte{0: bar0}
	ScrubBarContent(contents, classCode)

	maxPorts := int((readLE32(bar0, 0x04) >> 24) & 0xFF)
	portBase := 0x20 + 0x400
	if portBase+maxPorts*0x10 > bramSize {
		t.Errorf("MaxPorts (%d) port regs overflow BRAM", maxPorts)
	}
	if maxPorts < 1 {
		t.Error("MaxPorts should be at least 1")
	}
}

func TestScrubBarContent_XHCI_ConfigMaxSlotsEn(t *testing.T) {
	classCode := uint32(0x0C0330)
	bar0 := make([]byte, 4096)

	bar0[0x00] = 0x20
	writeLE32(bar0, 0x04, 8|(1<<8)|(2<<24)) // MaxSlots=8
	writeLE32(bar0, 0x58, 0xFF00FF00)       // CONFIG at capLen+0x38=0x58

	contents := map[int][]byte{0: bar0}
	ScrubBarContent(contents, classCode)

	config := readLE32(bar0, 0x58)
	maxSlotsEn := config & 0xFF
	maxSlots := readLE32(bar0, 0x04) & 0xFF
	if maxSlotsEn != maxSlots {
		t.Errorf("CONFIG MaxSlotsEn (%d) != MaxSlots (%d)", maxSlotsEn, maxSlots)
	}
	if config&0xFFFFFF00 != 0xFF00FF00 {
		t.Errorf("CONFIG upper bits should be preserved: 0x%08x", config)
	}
}

func TestScrubBarContent_XHCI_ClearDNCTRL_CRCR(t *testing.T) {
	classCode := uint32(0x0C0330)
	bar0 := make([]byte, 4096)

	bar0[0x00] = 0x20
	writeLE32(bar0, 0x04, 4|(1<<8)|(2<<24))
	writeLE32(bar0, 0x34, 0xFFFF)     // DNCTRL at capLen+0x14
	writeLE32(bar0, 0x38, 0xDEADBEEF) // CRCR lo at capLen+0x18
	writeLE32(bar0, 0x3C, 0xCAFEBABE) // CRCR hi at capLen+0x1C

	contents := map[int][]byte{0: bar0}
	ScrubBarContent(contents, classCode)

	if readLE32(bar0, 0x34) != 0 {
		t.Errorf("DNCTRL should be cleared, got 0x%08x", readLE32(bar0, 0x34))
	}
	if readLE32(bar0, 0x38) != 0 {
		t.Errorf("CRCR lo should be cleared, got 0x%08x", readLE32(bar0, 0x38))
	}
	if readLE32(bar0, 0x3C) != 0 {
		t.Errorf("CRCR hi should be cleared, got 0x%08x", readLE32(bar0, 0x3C))
	}
}
