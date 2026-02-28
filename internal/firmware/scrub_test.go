package firmware

import (
	"testing"

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

	scrubbed := ScrubConfigSpace(cs)

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

	scrubbed := ScrubConfigSpace(cs)

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

	scrubbed := ScrubConfigSpace(cs)

	// SR-IOV should be removed
	if scrubbed.ReadU32(0x100) != 0 {
		t.Errorf("SR-IOV should be filtered by ScrubConfigSpace: 0x%08x", scrubbed.ReadU32(0x100))
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
