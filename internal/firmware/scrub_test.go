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
