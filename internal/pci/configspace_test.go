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

	// Test boundary writes are no-ops (no panic)
	cs.WriteU8(-1, 0xFF)
	cs.WriteU8(ConfigSpaceSize, 0xFF)
	cs.WriteU16(ConfigSpaceSize-1, 0xFFFF)
	cs.WriteU32(ConfigSpaceSize-3, 0xFFFFFFFF)

	// Valid boundary reads
	cs.WriteU8(ConfigSpaceSize-1, 0xAB)
	if cs.ReadU8(ConfigSpaceSize-1) != 0xAB {
		t.Error("ReadU8 at last byte should work")
	}

	cs.WriteU16(ConfigSpaceSize-2, 0xCDEF)
	if cs.ReadU16(ConfigSpaceSize-2) != 0xCDEF {
		t.Error("ReadU16 at last 2 bytes should work")
	}

	cs.WriteU32(ConfigSpaceSize-4, 0x12345678)
	if cs.ReadU32(ConfigSpaceSize-4) != 0x12345678 {
		t.Error("ReadU32 at last 4 bytes should work")
	}
}

func TestConfigSpaceAllAccessors(t *testing.T) {
	cs := NewConfigSpace()

	// Set up a full PCI Type 0 header
	cs.WriteU16(0x04, 0x0547)     // Command
	cs.WriteU8(0x0C, 0x10)        // Cache Line Size
	cs.WriteU8(0x0D, 0x40)        // Latency Timer
	cs.WriteU8(0x0E, 0x80)        // Header Type (multi-function)
	cs.WriteU8(0x0F, 0x00)        // BIST
	cs.WriteU32(0x30, 0xFFF00001) // Expansion ROM
	cs.WriteU8(0x3C, 0x0B)        // Interrupt Line
	cs.WriteU8(0x3D, 0x01)        // Interrupt Pin
	cs.WriteU8(0x3E, 0x00)        // Min Grant
	cs.WriteU8(0x3F, 0xFF)        // Max Latency

	if cs.Command() != 0x0547 {
		t.Errorf("Command() = 0x%04x, want 0x0547", cs.Command())
	}
	if cs.CacheLineSize() != 0x10 {
		t.Errorf("CacheLineSize() = 0x%02x, want 0x10", cs.CacheLineSize())
	}
	if cs.LatencyTimer() != 0x40 {
		t.Errorf("LatencyTimer() = 0x%02x, want 0x40", cs.LatencyTimer())
	}
	if cs.HeaderType() != 0x80 {
		t.Errorf("HeaderType() = 0x%02x, want 0x80", cs.HeaderType())
	}
	if !cs.IsMultiFunction() {
		t.Error("IsMultiFunction() = false, want true")
	}
	if cs.HeaderLayout() != 0x00 {
		t.Errorf("HeaderLayout() = 0x%02x, want 0x00", cs.HeaderLayout())
	}
	if cs.BIST() != 0x00 {
		t.Errorf("BIST() = 0x%02x, want 0x00", cs.BIST())
	}
	if cs.ExpansionROMBase() != 0xFFF00001 {
		t.Errorf("ExpansionROMBase() = 0x%08x, want 0xFFF00001", cs.ExpansionROMBase())
	}
	if cs.InterruptLine() != 0x0B {
		t.Errorf("InterruptLine() = 0x%02x, want 0x0B", cs.InterruptLine())
	}
	if cs.InterruptPin() != 0x01 {
		t.Errorf("InterruptPin() = 0x%02x, want 0x01", cs.InterruptPin())
	}
	if cs.MinGrant() != 0x00 {
		t.Errorf("MinGrant() = 0x%02x, want 0x00", cs.MinGrant())
	}
	if cs.MaxLatency() != 0xFF {
		t.Errorf("MaxLatency() = 0x%02x, want 0xFF", cs.MaxLatency())
	}

	// Test non-multifunction header
	cs.WriteU8(0x0E, 0x01) // Type 1 header, not multi-function
	if cs.IsMultiFunction() {
		t.Error("IsMultiFunction() = true, want false for single-function")
	}
	if cs.HeaderLayout() != 0x01 {
		t.Errorf("HeaderLayout() = 0x%02x, want 0x01", cs.HeaderLayout())
	}
}

func TestConfigSpaceBAROutOfRange(t *testing.T) {
	cs := NewConfigSpace()
	cs.WriteU32(0x10, 0xFE000000)

	// Valid BAR read
	if cs.BAR(0) != 0xFE000000 {
		t.Errorf("BAR(0) = 0x%08x, want 0xFE000000", cs.BAR(0))
	}

	// Out-of-range reads
	if cs.BAR(-1) != 0 {
		t.Error("BAR(-1) should return 0")
	}
	if cs.BAR(6) != 0 {
		t.Error("BAR(6) should return 0")
	}
}

func TestConfigSpaceHexDumpDefaults(t *testing.T) {
	cs := NewConfigSpace()
	cs.Size = 256
	cs.WriteU16(0x00, 0x8086)

	// Zero/negative maxBytes should dump full size
	dump := cs.HexDump(0)
	if len(dump) == 0 {
		t.Error("HexDump(0) should dump full config space")
	}

	dump = cs.HexDump(-1)
	if len(dump) == 0 {
		t.Error("HexDump(-1) should dump full config space")
	}
}
