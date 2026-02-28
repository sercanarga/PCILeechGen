package pci

import (
	"strings"
	"testing"
)

func TestConfigSpaceAccessors(t *testing.T) {
	cs := NewConfigSpace()

	// Set up a typical Intel NIC config space header
	cs.WriteU16(0x00, 0x8086) // Vendor ID
	cs.WriteU16(0x02, 0x1533) // Device ID
	cs.WriteU16(0x04, 0x0406) // Command
	cs.WriteU16(0x06, 0x0010) // Status (capabilities list)
	cs.WriteU8(0x08, 0x03)    // Revision ID
	cs.WriteU8(0x09, 0x00)    // Prog IF
	cs.WriteU8(0x0A, 0x00)    // Sub-class
	cs.WriteU8(0x0B, 0x02)    // Base class (Network)
	cs.WriteU8(0x0E, 0x00)    // Header type
	cs.WriteU16(0x2C, 0x8086) // Subsys Vendor
	cs.WriteU16(0x2E, 0x0001) // Subsys Device
	cs.WriteU8(0x34, 0x40)    // Capability pointer

	if cs.VendorID() != 0x8086 {
		t.Errorf("VendorID() = 0x%04x, want 0x8086", cs.VendorID())
	}
	if cs.DeviceID() != 0x1533 {
		t.Errorf("DeviceID() = 0x%04x, want 0x1533", cs.DeviceID())
	}
	if cs.RevisionID() != 0x03 {
		t.Errorf("RevisionID() = 0x%02x, want 0x03", cs.RevisionID())
	}
	if cs.BaseClass() != 0x02 {
		t.Errorf("BaseClass() = 0x%02x, want 0x02", cs.BaseClass())
	}
	if cs.ClassCode() != 0x020000 {
		t.Errorf("ClassCode() = 0x%06x, want 0x020000", cs.ClassCode())
	}
	if cs.SubsysVendorID() != 0x8086 {
		t.Errorf("SubsysVendorID() = 0x%04x, want 0x8086", cs.SubsysVendorID())
	}
	if cs.SubsysDeviceID() != 0x0001 {
		t.Errorf("SubsysDeviceID() = 0x%04x, want 0x0001", cs.SubsysDeviceID())
	}
	if !cs.HasCapabilities() {
		t.Error("HasCapabilities() = false, want true")
	}
	if cs.CapabilityPointer() != 0x40 {
		t.Errorf("CapabilityPointer() = 0x%02x, want 0x40", cs.CapabilityPointer())
	}
}

func TestConfigSpaceFromBytes(t *testing.T) {
	data := make([]byte, 256)
	data[0] = 0x86
	data[1] = 0x80

	cs := NewConfigSpaceFromBytes(data)
	if cs.VendorID() != 0x8086 {
		t.Errorf("VendorID() = 0x%04x, want 0x8086", cs.VendorID())
	}
	if cs.Size != 256 {
		t.Errorf("Size = %d, want 256", cs.Size)
	}
}

func TestConfigSpaceClone(t *testing.T) {
	cs := NewConfigSpace()
	cs.WriteU16(0x00, 0x8086)
	cs.WriteU16(0x02, 0x1533)

	clone := cs.Clone()
	if clone.VendorID() != 0x8086 {
		t.Errorf("Clone VendorID = 0x%04x, want 0x8086", clone.VendorID())
	}

	// Modify original, clone should be independent
	cs.WriteU16(0x00, 0xFFFF)
	if clone.VendorID() != 0x8086 {
		t.Error("Clone was affected by modifying original")
	}
}

func TestConfigSpaceBytes(t *testing.T) {
	cs := NewConfigSpace()
	cs.Size = 256
	cs.WriteU16(0x00, 0x8086)

	bytes := cs.Bytes()
	if len(bytes) != 256 {
		t.Errorf("Bytes() len = %d, want 256", len(bytes))
	}
	if bytes[0] != 0x86 || bytes[1] != 0x80 {
		t.Errorf("Bytes() content wrong: %02x %02x", bytes[0], bytes[1])
	}
}

func TestConfigSpaceHexDump(t *testing.T) {
	cs := NewConfigSpace()
	cs.WriteU16(0x00, 0x8086)

	dump := cs.HexDump(16)
	if !strings.Contains(dump, "86 80") {
		t.Errorf("HexDump missing expected bytes, got: %s", dump)
	}
}

func TestConfigSpaceReadWriteBoundary(t *testing.T) {
	cs := NewConfigSpace()

	// Test boundary reads return 0
	if cs.ReadU8(-1) != 0 {
		t.Error("ReadU8 at -1 should return 0")
	}
	if cs.ReadU8(ConfigSpaceSize) != 0 {
		t.Error("ReadU8 at ConfigSpaceSize should return 0")
	}
	if cs.ReadU16(ConfigSpaceSize-1) != 0 {
		t.Error("ReadU16 at boundary should return 0")
	}
	if cs.ReadU32(ConfigSpaceSize-3) != 0 {
		t.Error("ReadU32 at boundary should return 0")
	}
}
