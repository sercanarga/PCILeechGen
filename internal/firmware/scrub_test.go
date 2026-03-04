package firmware

import (
	"testing"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

func TestScrubConfigSpace(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize

	// Set up a realistic config space
	cs.WriteU16(0x00, 0x8086) // Vendor ID
	cs.WriteU16(0x02, 0x1533) // Device ID
	cs.WriteU16(0x04, 0x0507) // Command: IO+Mem+BusMaster+extra bits
	cs.WriteU16(0x06, 0xFBB0) // Status: all error bits set + caps
	cs.WriteU8(0x08, 0x03)    // Revision ID
	cs.WriteU8(0x0C, 0x10)    // Cache Line Size
	cs.WriteU8(0x0D, 0x40)    // Latency Timer
	cs.WriteU8(0x0F, 0xC0)    // BIST: running
	cs.WriteU8(0x3C, 0x0B)    // Interrupt Line: IRQ 11

	scrubbed := ScrubConfigSpace(cs, nil)

	// Vendor/Device should be preserved
	if scrubbed.VendorID() != 0x8086 {
		t.Errorf("VendorID changed: 0x%04x", scrubbed.VendorID())
	}
	if scrubbed.DeviceID() != 0x1533 {
		t.Errorf("DeviceID changed: 0x%04x", scrubbed.DeviceID())
	}
	if scrubbed.RevisionID() != 0x03 {
		t.Errorf("RevisionID changed: 0x%02x", scrubbed.RevisionID())
	}

	// BIST should be cleared
	if scrubbed.BIST() != 0x00 {
		t.Errorf("BIST not cleared: 0x%02x", scrubbed.BIST())
	}

	// Interrupt line should be cleared
	if scrubbed.InterruptLine() != 0x00 {
		t.Errorf("InterruptLine not cleared: 0x%02x", scrubbed.InterruptLine())
	}

	// Latency timer should be cleared
	if scrubbed.LatencyTimer() != 0x00 {
		t.Errorf("LatencyTimer not cleared: 0x%02x", scrubbed.LatencyTimer())
	}

	// Cache line size should be cleared
	if scrubbed.CacheLineSize() != 0x00 {
		t.Errorf("CacheLineSize not cleared: 0x%02x", scrubbed.CacheLineSize())
	}

	// Status error bits should be cleared, cap bit preserved
	status := scrubbed.Status()
	if status&0x0010 == 0 {
		t.Error("Status capability bit should be preserved")
	}
	if status&0xF100 != 0 {
		t.Errorf("Status error bits not cleared: 0x%04x", status)
	}

	// Original should NOT be modified
	if cs.BIST() != 0xC0 {
		t.Error("Original BIST was modified!")
	}
}

func TestScrubConfigSpaceWithPCIeCap(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize

	cs.WriteU16(0x00, 0x8086)
	cs.WriteU16(0x06, 0x0010) // Status: caps
	cs.WriteU8(0x34, 0x40)    // Cap pointer

	// PCIe capability at 0x40
	cs.WriteU8(0x40, pci.CapIDPCIExpress)
	cs.WriteU8(0x41, 0x50)    // Next -> PM at 0x50
	cs.WriteU16(0x42, 0x0002) // PCIe Caps
	cs.WriteU16(0x4A, 0x000F) // Device Status: all errors set
	cs.WriteU16(0x52, 0xC001) // Link Status: training bits

	// PM capability at 0x50
	cs.WriteU8(0x50, pci.CapIDPowerManagement)
	cs.WriteU8(0x51, 0x00)    // Next = 0
	cs.WriteU16(0x52, 0x0003) // PM Caps
	cs.WriteU16(0x54, 0x8003) // PMCSR: D3 + PME_Status

	scrubbed := ScrubConfigSpace(cs, nil)

	// Device Status errors should be cleared
	devStatus := scrubbed.ReadU16(0x4A)
	if devStatus != 0x0000 {
		t.Errorf("PCIe Device Status not cleared: 0x%04x", devStatus)
	}

	// PM should be set to D0
	pmcsr := scrubbed.ReadU16(0x54)
	if pmcsr&0x0003 != 0x0000 {
		t.Errorf("PM not set to D0: 0x%04x", pmcsr)
	}
}

// makeExtCapHeader builds a 32-bit extended capability header.
func makeExtCapHeader(id uint16, version uint8, nextOffset int) uint32 {
	return uint32(id) | uint32(version)<<16 | uint32(nextOffset)<<20
}

func TestFilterExtCaps_RemoveMiddle(t *testing.T) {
	// Chain: AER(0x100) -> SR-IOV(0x150) -> DSN(0x200) -> LTR(0x250) -> end
	// SR-IOV should be removed. Chain should become AER -> DSN -> LTR -> end
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize

	cs.WriteU32(0x100, makeExtCapHeader(pci.ExtCapIDAER, 1, 0x150))
	cs.WriteU32(0x104, 0xDEADBEEF) // AER data

	cs.WriteU32(0x150, makeExtCapHeader(pci.ExtCapIDSRIOV, 1, 0x200))
	cs.WriteU32(0x154, 0xCAFEBABE) // SR-IOV data

	cs.WriteU32(0x200, makeExtCapHeader(pci.ExtCapIDDeviceSerialNumber, 1, 0x250))
	cs.WriteU32(0x204, 0x12345678) // DSN low
	cs.WriteU32(0x208, 0x9ABCDEF0) // DSN high

	cs.WriteU32(0x250, makeExtCapHeader(pci.ExtCapIDLTR, 1, 0))
	cs.WriteU32(0x254, 0x11223344) // LTR data

	removed := FilterExtCapabilities(cs)

	// Should have removed SR-IOV
	if len(removed) != 1 {
		t.Fatalf("Expected 1 removed cap, got %d: %v", len(removed), removed)
	}
	if !contains(removed[0], "SR-IOV") {
		t.Errorf("Expected SR-IOV in removed list, got: %s", removed[0])
	}

	// SR-IOV data region should be zeroed
	if cs.ReadU32(0x150) != 0 {
		t.Errorf("SR-IOV header not zeroed: 0x%08x", cs.ReadU32(0x150))
	}
	if cs.ReadU32(0x154) != 0 {
		t.Errorf("SR-IOV data not zeroed: 0x%08x", cs.ReadU32(0x154))
	}

	// AER should now point to DSN (0x200)
	aerHeader := cs.ReadU32(0x100)
	aerNext := int((aerHeader >> 20) & 0xFFC)
	if aerNext != 0x200 {
		t.Errorf("AER next-pointer should be 0x200, got 0x%03x", aerNext)
	}

	// AER data should be preserved
	if cs.ReadU32(0x104) != 0xDEADBEEF {
		t.Errorf("AER data corrupted: 0x%08x", cs.ReadU32(0x104))
	}

	// DSN data should be preserved
	if cs.ReadU32(0x204) != 0x12345678 {
		t.Errorf("DSN data corrupted: 0x%08x", cs.ReadU32(0x204))
	}

	// LTR data should be preserved
	if cs.ReadU32(0x254) != 0x11223344 {
		t.Errorf("LTR data corrupted: 0x%08x", cs.ReadU32(0x254))
	}

	// Verify chain walks correctly after filtering
	caps := pci.ParseExtCapabilities(cs)
	if len(caps) != 3 { // AER, DSN, LTR
		t.Errorf("Expected 3 remaining caps, got %d", len(caps))
	}
}

func TestFilterExtCaps_NoUnsafe(t *testing.T) {
	// Chain: AER(0x100) -> DSN(0x150) -> LTR(0x200) -> end
	// Nothing should be removed
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize

	cs.WriteU32(0x100, makeExtCapHeader(pci.ExtCapIDAER, 1, 0x150))
	cs.WriteU32(0x150, makeExtCapHeader(pci.ExtCapIDDeviceSerialNumber, 1, 0x200))
	cs.WriteU32(0x200, makeExtCapHeader(pci.ExtCapIDLTR, 1, 0))

	removed := FilterExtCapabilities(cs)

	if len(removed) != 0 {
		t.Errorf("Expected no removals, got %d: %v", len(removed), removed)
	}

	// Chain should be intact
	caps := pci.ParseExtCapabilities(cs)
	if len(caps) != 3 {
		t.Errorf("Expected 3 caps, got %d", len(caps))
	}
}

func TestFilterExtCaps_AllRemoved(t *testing.T) {
	// Chain: SR-IOV(0x100) -> ResizableBAR(0x150) -> end
	// Both should be removed
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize

	cs.WriteU32(0x100, makeExtCapHeader(pci.ExtCapIDSRIOV, 1, 0x150))
	cs.WriteU32(0x104, 0xAAAAAAAA)

	cs.WriteU32(0x150, makeExtCapHeader(pci.ExtCapIDResizableBAR, 1, 0))
	cs.WriteU32(0x154, 0xBBBBBBBB)

	removed := FilterExtCapabilities(cs)

	if len(removed) != 2 {
		t.Fatalf("Expected 2 removed caps, got %d: %v", len(removed), removed)
	}

	// 0x100 should be zeroed (end of list)
	if cs.ReadU32(0x100) != 0 {
		t.Errorf("First ext cap header should be zero: 0x%08x", cs.ReadU32(0x100))
	}

	// No caps should remain
	caps := pci.ParseExtCapabilities(cs)
	if len(caps) != 0 {
		t.Errorf("Expected 0 remaining caps, got %d", len(caps))
	}
}

func TestFilterExtCaps_RemoveLast(t *testing.T) {
	// Chain: AER(0x100) -> DSN(0x150) -> L1PM(0x200) -> end
	// L1PM should be removed. Chain: AER -> DSN -> end
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize

	cs.WriteU32(0x100, makeExtCapHeader(pci.ExtCapIDAER, 1, 0x150))
	cs.WriteU32(0x150, makeExtCapHeader(pci.ExtCapIDDeviceSerialNumber, 1, 0x200))
	cs.WriteU32(0x200, makeExtCapHeader(pci.ExtCapIDL1PMSubstates, 1, 0))
	cs.WriteU32(0x204, 0xCCCCCCCC)

	removed := FilterExtCapabilities(cs)

	if len(removed) != 1 {
		t.Fatalf("Expected 1 removed cap, got %d", len(removed))
	}

	// DSN should now point to 0 (end of list)
	dsnHeader := cs.ReadU32(0x150)
	dsnNext := int((dsnHeader >> 20) & 0xFFC)
	if dsnNext != 0 {
		t.Errorf("DSN next-pointer should be 0 (end), got 0x%03x", dsnNext)
	}

	// L1PM data should be zeroed
	if cs.ReadU32(0x200) != 0 || cs.ReadU32(0x204) != 0 {
		t.Error("L1PM data should be zeroed")
	}

	// 2 caps should remain
	caps := pci.ParseExtCapabilities(cs)
	if len(caps) != 2 {
		t.Errorf("Expected 2 remaining caps, got %d", len(caps))
	}
}

func TestScrubConfigSpace_FiltersExtCaps(t *testing.T) {
	// Integration test: ScrubConfigSpace should call FilterExtCapabilities
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize

	cs.WriteU16(0x00, 0x8086)
	cs.WriteU16(0x06, 0x0010)

	// Add SR-IOV in extended space — should be filtered by ScrubConfigSpace
	cs.WriteU32(0x100, makeExtCapHeader(pci.ExtCapIDSRIOV, 1, 0))
	cs.WriteU32(0x104, 0xFFFFFFFF)

	scrubbed := ScrubConfigSpace(cs, nil)

	// SR-IOV should be removed
	if scrubbed.ReadU32(0x100) != 0 {
		t.Errorf("SR-IOV should be filtered by ScrubConfigSpace: 0x%08x", scrubbed.ReadU32(0x100))
	}
}

func TestFilterExtCaps_RemoveFirst(t *testing.T) {
	// Chain: SR-IOV(0x100) -> AER(0x150) -> DSN(0x200) -> end
	// SR-IOV should be removed, AER relocated to 0x100.
	// Chain should become AER(0x100) -> DSN(0x200) -> end
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize

	cs.WriteU32(0x100, makeExtCapHeader(pci.ExtCapIDSRIOV, 1, 0x150))
	cs.WriteU32(0x104, 0xAAAAAAAA) // SR-IOV data
	cs.WriteU32(0x108, 0xBBBBBBBB) // SR-IOV data

	cs.WriteU32(0x150, makeExtCapHeader(pci.ExtCapIDAER, 2, 0x200))
	cs.WriteU32(0x154, 0xDEADBEEF) // AER uncorrectable error status
	cs.WriteU32(0x158, 0x11111111) // AER uncorrectable error mask
	cs.WriteU32(0x15C, 0x22222222) // AER data

	cs.WriteU32(0x200, makeExtCapHeader(pci.ExtCapIDDeviceSerialNumber, 1, 0))
	cs.WriteU32(0x204, 0x12345678) // DSN low
	cs.WriteU32(0x208, 0x9ABCDEF0) // DSN high

	removed := FilterExtCapabilities(cs)

	// Should have removed SR-IOV
	if len(removed) != 1 {
		t.Fatalf("Expected 1 removed cap, got %d: %v", len(removed), removed)
	}

	// 0x100 should now contain AER (cap ID 0x0001)
	hdr := cs.ReadU32(0x100)
	capID := uint16(hdr & 0xFFFF)
	if capID != pci.ExtCapIDAER {
		t.Errorf("0x100 should contain AER (0x0001), got 0x%04x", capID)
	}

	// AER version should be preserved
	version := uint8((hdr >> 16) & 0xF)
	if version != 2 {
		t.Errorf("AER version should be 2, got %d", version)
	}

	// AER data should be relocated to 0x100
	if cs.ReadU32(0x104) != 0xDEADBEEF {
		t.Errorf("AER data at 0x104 should be 0xDEADBEEF, got 0x%08x", cs.ReadU32(0x104))
	}
	if cs.ReadU32(0x108) != 0x11111111 {
		t.Errorf("AER data at 0x108 should be 0x11111111, got 0x%08x", cs.ReadU32(0x108))
	}

	// 0x100 next-pointer should point to DSN at 0x200
	nextOff := int((hdr >> 20) & 0xFFC)
	if nextOff != 0x200 {
		t.Errorf("AER at 0x100 should point to 0x200, got 0x%03x", nextOff)
	}

	// Original AER location (0x150) should be zeroed
	if cs.ReadU32(0x150) != 0 {
		t.Errorf("Original AER location (0x150) should be zeroed, got 0x%08x", cs.ReadU32(0x150))
	}

	// DSN data should be preserved
	if cs.ReadU32(0x204) != 0x12345678 {
		t.Errorf("DSN data corrupted: 0x%08x", cs.ReadU32(0x204))
	}

	// Verify chain walks correctly: AER(0x100) -> DSN(0x200) -> end
	caps := pci.ParseExtCapabilities(cs)
	if len(caps) != 2 {
		t.Fatalf("Expected 2 remaining caps, got %d", len(caps))
	}
	if caps[0].ID != pci.ExtCapIDAER || caps[0].Offset != 0x100 {
		t.Errorf("First cap should be AER at 0x100, got 0x%04x at 0x%03x", caps[0].ID, caps[0].Offset)
	}
	if caps[1].ID != pci.ExtCapIDDeviceSerialNumber || caps[1].Offset != 0x200 {
		t.Errorf("Second cap should be DSN at 0x200, got 0x%04x at 0x%03x", caps[1].ID, caps[1].Offset)
	}
}

func TestFilterExtCaps_RemoveFirstAndMiddle(t *testing.T) {
	// Chain: L1PM(0x100) -> AER(0x140) -> SecPCIe(0x180) -> DSN(0x200) -> end
	// L1PM and SecPCIe should be removed.
	// Chain should become AER(0x100) -> DSN(0x200) -> end
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize

	cs.WriteU32(0x100, makeExtCapHeader(pci.ExtCapIDL1PMSubstates, 1, 0x140))
	cs.WriteU32(0x104, 0xAAAAAAAA)

	cs.WriteU32(0x140, makeExtCapHeader(pci.ExtCapIDAER, 1, 0x180))
	cs.WriteU32(0x144, 0xDEADBEEF) // AER data

	cs.WriteU32(0x180, makeExtCapHeader(pci.ExtCapIDSecondaryPCIe, 1, 0x200))
	cs.WriteU32(0x184, 0xCCCCCCCC)

	cs.WriteU32(0x200, makeExtCapHeader(pci.ExtCapIDDeviceSerialNumber, 1, 0))
	cs.WriteU32(0x204, 0x12345678)

	removed := FilterExtCapabilities(cs)

	if len(removed) != 2 {
		t.Fatalf("Expected 2 removed caps, got %d: %v", len(removed), removed)
	}

	// 0x100 should now contain AER
	hdr := cs.ReadU32(0x100)
	capID := uint16(hdr & 0xFFFF)
	if capID != pci.ExtCapIDAER {
		t.Errorf("0x100 should contain AER, got 0x%04x", capID)
	}

	// AER data should be relocated
	if cs.ReadU32(0x104) != 0xDEADBEEF {
		t.Errorf("AER data should be relocated to 0x104, got 0x%08x", cs.ReadU32(0x104))
	}

	// Chain: AER(0x100) -> DSN(0x200) -> end
	caps := pci.ParseExtCapabilities(cs)
	if len(caps) != 2 {
		t.Fatalf("Expected 2 remaining caps, got %d", len(caps))
	}
	if caps[0].ID != pci.ExtCapIDAER {
		t.Errorf("First cap should be AER, got 0x%04x", caps[0].ID)
	}
	if caps[1].ID != pci.ExtCapIDDeviceSerialNumber {
		t.Errorf("Second cap should be DSN, got 0x%04x", caps[1].ID)
	}
}

func TestScrubConfigSpace_ClampBAR0(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize

	cs.WriteU16(0x00, 0x144D) // Samsung Vendor ID
	cs.WriteU16(0x02, 0xA80A) // Device ID
	cs.WriteU16(0x06, 0x0010) // Status: caps

	// BAR0: 64-bit memory, 16 KB size (mask = 0xFFFFC000)
	cs.WriteU32(0x10, 0xFFFFC004) // mem64, 16 KB
	cs.WriteU32(0x14, 0xFFFFFFFF) // upper 32 bits

	// BAR2: 32-bit memory, 64 KB size
	cs.WriteU32(0x18, 0xFFFF0000) // mem32, 64 KB

	scrubbed := ScrubConfigSpace(cs, nil)

	// BAR0 should be clamped to 4 KB (mask = 0xFFFFF000, type bits preserved)
	bar0 := scrubbed.BAR(0)
	expectedBar0 := uint32(0xFFFFF004) // 4 KB mask | mem64 type bits
	if bar0 != expectedBar0 {
		t.Errorf("BAR0 should be clamped to 4 KB: got 0x%08x, want 0x%08x", bar0, expectedBar0)
	}

	// BAR1 (upper half of 64-bit BAR0) should be zeroed
	bar1 := scrubbed.BAR(1)
	if bar1 != 0 {
		t.Errorf("BAR1 (upper 64-bit) should be zero: got 0x%08x", bar1)
	}

	// BAR2 should also be clamped to 4 KB
	bar2 := scrubbed.BAR(2)
	expectedBar2 := uint32(0xFFFFF000) // 4 KB mask | mem32 type bits
	if bar2 != expectedBar2 {
		t.Errorf("BAR2 should be clamped to 4 KB: got 0x%08x, want 0x%08x", bar2, expectedBar2)
	}
}

func TestIsUnsafeExtCap(t *testing.T) {
	// Safe caps
	if IsUnsafeExtCap(pci.ExtCapIDAER) {
		t.Error("AER should be safe")
	}
	if IsUnsafeExtCap(pci.ExtCapIDDeviceSerialNumber) {
		t.Error("DSN should be safe")
	}
	if IsUnsafeExtCap(pci.ExtCapIDLTR) {
		t.Error("LTR should be safe")
	}

	// Unsafe caps
	if !IsUnsafeExtCap(pci.ExtCapIDSRIOV) {
		t.Error("SR-IOV should be unsafe")
	}
	if !IsUnsafeExtCap(pci.ExtCapIDResizableBAR) {
		t.Error("Resizable BAR should be unsafe")
	}
	if !IsUnsafeExtCap(pci.ExtCapIDL1PMSubstates) {
		t.Error("L1PM Substates should be unsafe")
	}
	if !IsUnsafeExtCap(pci.ExtCapIDPTM) {
		t.Error("PTM should be unsafe")
	}
}

func TestScrubConfigSpace_ClampLinkCapability(t *testing.T) {
	// Simulate Gen4 x4 donor with CaptainDMA_75T (x1) board
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize

	cs.WriteU16(0x00, 0x144D) // Samsung
	cs.WriteU16(0x06, 0x0010) // Status: caps
	cs.WriteU8(0x34, 0x70)    // Cap pointer

	// PCIe capability at 0x70
	cs.WriteU8(0x70, pci.CapIDPCIExpress)
	cs.WriteU8(0x71, 0x00) // end of cap list

	// Link Capabilities Register (cap+0x0C = 0x7C): Gen4 x4
	// bits [3:0] = 4 (Gen4), bits [9:4] = 4 (x4)
	linkCap := uint32(4) | (uint32(4) << 4) // speed=4, width=4
	cs.WriteU32(0x7C, linkCap)

	// Link Status Register (cap+0x12 = 0x82): Gen4 x4 current
	linkStatus := uint16(4) | (uint16(4) << 4) // speed=4, width=4
	cs.WriteU16(0x82, linkStatus)

	// Link Control 2 Register (cap+0x30 = 0xA0): Target=Gen4
	cs.WriteU16(0xA0, 4) // target speed = Gen4

	// Link Capabilities 2 Register (cap+0x2C = 0x9C): Gen1-4 supported
	// bits [7:1] = speed vector: bit1=Gen1, bit2=Gen2, bit3=Gen3, bit4=Gen4
	cs.WriteU32(0x9C, (0x1E)<<1) // 0x1E = 11110b (Gen1-4)

	b := &board.Board{PCIeLanes: 1} // x1 board
	scrubbed := ScrubConfigSpace(cs, b)

	// Link Capabilities: should be Gen2 x1
	scrubbedLinkCap := scrubbed.ReadU32(0x7C)
	speed := uint8(scrubbedLinkCap & 0x0F)
	width := uint8((scrubbedLinkCap >> 4) & 0x3F)
	if speed != LinkSpeedGen2 {
		t.Errorf("Link Cap Max Speed: got %d, want %d (Gen2)", speed, LinkSpeedGen2)
	}
	if width != 1 {
		t.Errorf("Link Cap Max Width: got %d, want 1", width)
	}

	// Link Status: should be Gen2 x1
	scrubbedLinkStatus := scrubbed.ReadU16(0x82)
	curSpeed := uint8(scrubbedLinkStatus & 0x0F)
	curWidth := uint8((scrubbedLinkStatus >> 4) & 0x3F)
	if curSpeed != LinkSpeedGen2 {
		t.Errorf("Link Status Current Speed: got %d, want %d (Gen2)", curSpeed, LinkSpeedGen2)
	}
	if curWidth != 1 {
		t.Errorf("Link Status Negotiated Width: got %d, want 1", curWidth)
	}

	// Link Control 2: Target Link Speed should be Gen2
	targetSpeed := uint8(scrubbed.ReadU16(0xA0) & 0x0F)
	if targetSpeed != LinkSpeedGen2 {
		t.Errorf("Link Control 2 Target Speed: got %d, want %d (Gen2)", targetSpeed, LinkSpeedGen2)
	}

	// Link Capabilities 2: Supported speeds should be Gen1+Gen2 only
	linkCap2 := scrubbed.ReadU32(0x9C)
	speedVector := (linkCap2 >> 1) & 0x7F
	// Gen1=bit0, Gen2=bit1 → 0x03
	expectedVector := uint32(0x03) // Gen1 + Gen2
	if speedVector != expectedVector {
		t.Errorf("Link Cap 2 speed vector: got 0x%02x, want 0x%02x", speedVector, expectedVector)
	}
}

func TestScrubConfigSpace_ClampLinkCapability_NilBoard(t *testing.T) {
	// Without board info, only speed should be clamped (not width)
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize

	cs.WriteU16(0x00, 0x144D)
	cs.WriteU16(0x06, 0x0010)
	cs.WriteU8(0x34, 0x70)

	cs.WriteU8(0x70, pci.CapIDPCIExpress)
	cs.WriteU8(0x71, 0x00)

	// Gen4 x4
	linkCap := uint32(4) | (uint32(4) << 4)
	cs.WriteU32(0x7C, linkCap)

	scrubbed := ScrubConfigSpace(cs, nil) // nil board

	scrubbedLinkCap := scrubbed.ReadU32(0x7C)
	speed := uint8(scrubbedLinkCap & 0x0F)
	width := uint8((scrubbedLinkCap >> 4) & 0x3F)

	// Speed should be clamped to Gen2
	if speed != LinkSpeedGen2 {
		t.Errorf("Speed should be clamped to Gen2: got %d", speed)
	}
	// Width should remain unchanged (no board info)
	if width != 4 {
		t.Errorf("Width should remain x4 without board info: got %d", width)
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > 0 && containsHelper(s, substr))
}

func containsHelper(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

func TestScrubConfigSpace_ClampDeviceCapability(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize

	cs.WriteU16(0x00, 0x144D)
	cs.WriteU16(0x06, 0x0010) // Status: caps
	cs.WriteU8(0x34, 0x70)    // Cap pointer

	// PCIe capability at 0x70
	cs.WriteU8(0x70, pci.CapIDPCIExpress)
	cs.WriteU8(0x71, 0x00) // end of cap list

	// Device Capabilities (cap+0x04 = 0x74):
	// bits [2:0] = 2 (MPS 512B), bit 5 = 1 (ExtTag), bits [4:3] = 1 (Phantom)
	devCap := uint32(2) | (1 << 3) | (1 << 5) | (0x07 << 6) // MPS=512B, Phantom=1, ExtTag=1, L0s latency
	cs.WriteU32(0x74, devCap)

	// Device Control (cap+0x08 = 0x78):
	// bits [7:5] = 2 (MPS 512B), bit 8 = 1 (ExtTag), bit 9 = 1 (Phantom), bits [14:12] = 5 (MRRS 4KB)
	devCtl := uint16(2<<5) | (1 << 8) | (1 << 9) | (5 << 12)
	cs.WriteU16(0x78, devCtl)

	// Device Capabilities 2 (cap+0x24 = 0x94):
	// bit 16 = 1 (10-bit Tag Completer), bit 17 = 1 (10-bit Tag Requester)
	devCap2 := uint32(1<<16) | (1 << 17) | (0x0F) // some other bits
	cs.WriteU32(0x94, devCap2)

	scrubbed := ScrubConfigSpace(cs, nil)

	// Device Capabilities: MPS should be 0 (128B), ExtTag=0, Phantom=0
	scrubbedDevCap := scrubbed.ReadU32(0x74)
	mps := scrubbedDevCap & 0x07
	if mps != 0 {
		t.Errorf("Device Cap MPS Supported should be 0 (128B), got %d", mps)
	}
	phantom := (scrubbedDevCap >> 3) & 0x03
	if phantom != 0 {
		t.Errorf("Device Cap Phantom Functions should be 0, got %d", phantom)
	}
	extTag := (scrubbedDevCap >> 5) & 0x01
	if extTag != 0 {
		t.Errorf("Device Cap Extended Tag should be 0, got %d", extTag)
	}
	// L0s latency (bits [8:6]) should be preserved
	l0sLat := (scrubbedDevCap >> 6) & 0x07
	if l0sLat != 0x07 {
		t.Errorf("Device Cap L0s Latency should be preserved (7), got %d", l0sLat)
	}

	// Device Control: MPS=0, ExtTag=0, Phantom=0, MRRS clamped to 2 (512B)
	scrubbedDevCtl := scrubbed.ReadU16(0x78)
	ctlMPS := (scrubbedDevCtl >> 5) & 0x07
	if ctlMPS != 0 {
		t.Errorf("Device Control MPS should be 0 (128B), got %d", ctlMPS)
	}
	ctlExtTag := (scrubbedDevCtl >> 8) & 0x01
	if ctlExtTag != 0 {
		t.Errorf("Device Control Extended Tag Enable should be 0, got %d", ctlExtTag)
	}
	ctlPhantom := (scrubbedDevCtl >> 9) & 0x01
	if ctlPhantom != 0 {
		t.Errorf("Device Control Phantom Enable should be 0, got %d", ctlPhantom)
	}
	ctlMRRS := (scrubbedDevCtl >> 12) & 0x07
	if ctlMRRS != 2 {
		t.Errorf("Device Control MRRS should be clamped to 2 (512B), got %d", ctlMRRS)
	}

	// Device Capabilities 2: 10-bit tag fields should be 0
	scrubbedDevCap2 := scrubbed.ReadU32(0x94)
	if scrubbedDevCap2&(1<<16) != 0 {
		t.Error("Device Cap 2: 10-bit Tag Completer should be 0")
	}
	if scrubbedDevCap2&(1<<17) != 0 {
		t.Error("Device Cap 2: 10-bit Tag Requester should be 0")
	}
	// Other bits should be preserved
	if scrubbedDevCap2&0x0F != 0x0F {
		t.Errorf("Device Cap 2: other bits should be preserved, got 0x%08x", scrubbedDevCap2)
	}
}

func TestScrubConfigSpace_ClampDeviceControl_SmallMRRS(t *testing.T) {
	// If MRRS is already <= 512B, it should not be changed
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize

	cs.WriteU16(0x06, 0x0010)
	cs.WriteU8(0x34, 0x70)

	cs.WriteU8(0x70, pci.CapIDPCIExpress)
	cs.WriteU8(0x71, 0x00)

	// Device Control: MRRS = 1 (256B) — should stay
	devCtl := uint16(1 << 12)
	cs.WriteU16(0x78, devCtl)

	scrubbed := ScrubConfigSpace(cs, nil)

	scrubbedDevCtl := scrubbed.ReadU16(0x78)
	ctlMRRS := (scrubbedDevCtl >> 12) & 0x07
	if ctlMRRS != 1 {
		t.Errorf("Device Control MRRS should remain 1 (256B), got %d", ctlMRRS)
	}
}

func TestDisableMSIXIfOutOfBRAM_TableOutside(t *testing.T) {
	// MSI-X table at BAR0+0x1000 (>= 4KB BRAM) should be disabled
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize

	cs.WriteU16(0x06, 0x0010) // Status: caps
	cs.WriteU8(0x34, 0x90)    // Cap pointer

	// MSI-X cap at 0x90
	cs.WriteU8(0x90, pci.CapIDMSIX)
	cs.WriteU8(0x91, 0x00)    // end of cap list
	cs.WriteU16(0x92, 0x8007) // Message Control: Enable=1, table_size=8

	// Table: BIR=0, offset=0x1000 (outside 4KB BRAM)
	cs.WriteU32(0x94, 0x00001000)
	// PBA: BIR=0, offset=0x1080
	cs.WriteU32(0x98, 0x00001080)

	// BAR0: mem64
	cs.WriteU32(0x10, 0xFFFFF004)

	scrubbed := ScrubConfigSpace(cs, nil)

	// MSI-X Enable (bit 15) should be cleared
	msgCtl := scrubbed.ReadU16(0x92)
	if msgCtl&0x8000 != 0 {
		t.Errorf("MSI-X should be disabled when table is outside BRAM, got msgCtl=0x%04x", msgCtl)
	}
	// Function Mask (bit 14) should also be cleared
	if msgCtl&0x4000 != 0 {
		t.Errorf("MSI-X Function Mask should be cleared, got msgCtl=0x%04x", msgCtl)
	}
	// Table size bits should be preserved
	if msgCtl&0x07FF != 0x0007 {
		t.Errorf("MSI-X table size should be preserved, got 0x%04x", msgCtl&0x07FF)
	}
}

func TestDisableMSIXIfOutOfBRAM_TableInside(t *testing.T) {
	// MSI-X table at BAR0+0x0000 (inside 4KB BRAM) should NOT be disabled
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize

	cs.WriteU16(0x06, 0x0010)
	cs.WriteU8(0x34, 0x90)

	cs.WriteU8(0x90, pci.CapIDMSIX)
	cs.WriteU8(0x91, 0x00)
	cs.WriteU16(0x92, 0x8003) // Enable=1, table_size=4

	// Table: BIR=0, offset=0x0000 (inside BRAM)
	cs.WriteU32(0x94, 0x00000000)
	// PBA: BIR=0, offset=0x0100
	cs.WriteU32(0x98, 0x00000100)

	scrubbed := ScrubConfigSpace(cs, nil)

	// MSI-X Enable should remain set
	msgCtl := scrubbed.ReadU16(0x92)
	if msgCtl&0x8000 == 0 {
		t.Errorf("MSI-X should remain enabled when table is inside BRAM, got msgCtl=0x%04x", msgCtl)
	}
}

func TestDisableMSIXIfOutOfBRAM_PBAOutside(t *testing.T) {
	// Table inside but PBA outside — should still disable
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize

	cs.WriteU16(0x06, 0x0010)
	cs.WriteU8(0x34, 0x90)

	cs.WriteU8(0x90, pci.CapIDMSIX)
	cs.WriteU8(0x91, 0x00)
	cs.WriteU16(0x92, 0x8003) // Enable=1

	// Table inside BRAM
	cs.WriteU32(0x94, 0x00000800)
	// PBA at offset 0x2000 (outside BRAM)
	cs.WriteU32(0x98, 0x00002000)

	scrubbed := ScrubConfigSpace(cs, nil)

	msgCtl := scrubbed.ReadU16(0x92)
	if msgCtl&0x8000 != 0 {
		t.Errorf("MSI-X should be disabled when PBA is outside BRAM, got msgCtl=0x%04x", msgCtl)
	}
}

func TestFakeRenesasFirmwareReady(t *testing.T) {
	// Renesas xHCI device should get firmware status set to SUCCESS
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize

	cs.WriteU16(0x00, 0x1912) // Renesas Vendor ID
	cs.WriteU16(0x02, 0x0014) // uPD720201 Device ID
	cs.WriteU8(0x09, 0x30)    // ProgIF = xHCI
	cs.WriteU8(0x0A, 0x03)    // SubClass = USB
	cs.WriteU8(0x0B, 0x0C)    // BaseClass = Serial Bus
	cs.WriteU16(0x06, 0x0010) // Status: caps

	// Set ROM_EXISTS bit in ROM Status (word at 0xF6)
	// DWORD at 0xF4 = FW_STATUS(byte) | FW_STATUS_MSB(byte) | ROM_STATUS(word)
	cs.WriteU8(0xF4, 0x00)    // FW Status: no result (firmware not loaded)
	cs.WriteU16(0xF6, 0x8000) // ROM Status: ROM_EXISTS=1, Result=0 (no result)

	scrubbed := ScrubConfigSpace(cs, nil)

	// FW Status should be SUCCESS (0x10) | LOCK (0x80) = 0x90
	fwStatus := scrubbed.ReadU8(0xF4)
	if fwStatus&0x10 == 0 {
		t.Errorf("Renesas FW Status SUCCESS bit should be set, got 0x%02x", fwStatus)
	}
	if fwStatus&0x80 == 0 {
		t.Errorf("Renesas FW Status LOCK bit should be set, got 0x%02x", fwStatus)
	}

	// ROM Status should have SUCCESS (bit 4) set
	// (ROM_EXISTS may be cleared by vendor region wipe — that's fine,
	// the driver only cares about the result code)
	romStatus := scrubbed.ReadU16(0xF6)
	if romStatus&0x0010 == 0 {
		t.Errorf("Renesas ROM Status SUCCESS bit should be set, got 0x%04x", romStatus)
	}
}

func TestFakeRenesasFirmwareReady_NonRenesas(t *testing.T) {
	// Non-Renesas xHCI device should NOT get firmware status modified
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize

	cs.WriteU16(0x00, 0x8086) // Intel Vendor ID
	cs.WriteU16(0x02, 0xA36D) // Some Intel xHCI
	cs.WriteU8(0x09, 0x30)    // ProgIF = xHCI
	cs.WriteU8(0x0A, 0x03)    // SubClass = USB
	cs.WriteU8(0x0B, 0x0C)    // BaseClass = Serial Bus
	cs.WriteU16(0x06, 0x0010)

	cs.WriteU8(0xF4, 0x00) // Some value at the same offset

	scrubbed := ScrubConfigSpace(cs, nil)

	// Offset 0xF4 should remain 0x00 (not modified for Intel)
	fwStatus := scrubbed.ReadU8(0xF4)
	if fwStatus != 0x00 {
		t.Errorf("Non-Renesas device offset 0xF4 should not be changed, got 0x%02x", fwStatus)
	}
}

func TestZeroVendorRegisters(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize

	cs.WriteU16(0x00, 0x8086)
	cs.WriteU16(0x06, 0x0010) // caps present
	cs.WriteU8(0x34, 0x40)    // cap pointer

	// PM cap at 0x40, next → 0x70
	cs.WriteU8(0x40, pci.CapIDPowerManagement)
	cs.WriteU8(0x41, 0x70)
	cs.WriteU16(0x42, 0xC9C3) // PM caps
	cs.WriteU16(0x44, 0x0008) // PMCSR

	// PCIe cap at 0x70, end of chain
	cs.WriteU8(0x70, pci.CapIDPCIExpress)
	cs.WriteU8(0x71, 0x00)

	// vendor garbage at 0xB0 and 0xF4 (outside any cap)
	cs.WriteU32(0xB0, 0xDEADBEEF)
	cs.WriteU32(0xF4, 0xCAFEBABE)

	scrubbed := ScrubConfigSpace(cs, nil)

	// vendor bytes should be wiped
	if scrubbed.ReadU32(0xB0) != 0 {
		t.Errorf("vendor register at 0xB0 should be zeroed, got 0x%08x", scrubbed.ReadU32(0xB0))
	}
	if scrubbed.ReadU32(0xF4) != 0 {
		t.Errorf("vendor register at 0xF4 should be zeroed, got 0x%08x", scrubbed.ReadU32(0xF4))
	}

	// cap data should survive
	if scrubbed.ReadU8(0x40) != pci.CapIDPowerManagement {
		t.Errorf("PM cap ID at 0x40 should be preserved, got 0x%02x", scrubbed.ReadU8(0x40))
	}
	if scrubbed.ReadU8(0x70) != pci.CapIDPCIExpress {
		t.Errorf("PCIe cap ID at 0x70 should be preserved, got 0x%02x", scrubbed.ReadU8(0x70))
	}
}
