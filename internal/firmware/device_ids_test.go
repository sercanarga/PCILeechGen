package firmware

import (
	"encoding/binary"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/pci"
)

func TestExtractDeviceIDs(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize

	cs.WriteU16(0x00, 0x8086)
	cs.WriteU16(0x02, 0x1533)
	cs.WriteU16(0x2C, 0x8086)
	cs.WriteU16(0x2E, 0x0001)
	cs.WriteU8(0x08, 0x03)
	cs.WriteU8(0x09, 0x00) // ProgIF
	cs.WriteU8(0x0A, 0x00) // SubClass
	cs.WriteU8(0x0B, 0x02) // BaseClass

	ids := ExtractDeviceIDs(cs, nil)

	if ids.VendorID != 0x8086 {
		t.Errorf("VendorID = 0x%04x, want 0x8086", ids.VendorID)
	}
	if ids.DeviceID != 0x1533 {
		t.Errorf("DeviceID = 0x%04x, want 0x1533", ids.DeviceID)
	}
	if ids.SubsysVendorID != 0x8086 {
		t.Errorf("SubsysVendorID = 0x%04x, want 0x8086", ids.SubsysVendorID)
	}
	if ids.RevisionID != 0x03 {
		t.Errorf("RevisionID = 0x%02x, want 0x03", ids.RevisionID)
	}
	if ids.ClassCode != 0x020000 {
		t.Errorf("ClassCode = 0x%06x, want 0x020000", ids.ClassCode)
	}
	if ids.HasDSN {
		t.Error("HasDSN should be false when no DSN capability")
	}
	if ids.HasPCIeCap {
		t.Error("HasPCIeCap should be false when no PCIe capability")
	}
}

func TestExtractDeviceIDsWithPCIeCap(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize

	cs.WriteU16(0x00, 0x8086)
	cs.WriteU16(0x02, 0x1533)
	// Set capability pointer and status
	cs.WriteU8(0x34, 0x40)    // Capabilities pointer -> 0x40
	cs.WriteU16(0x06, 0x0010) // Status: capabilities list present

	// PCIe capability at offset 0x40
	cs.WriteU8(0x40, pci.CapIDPCIExpress) // Capability ID
	cs.WriteU8(0x41, 0x00)                // Next pointer
	// PCIe Capabilities Register (cap+2): Device/Port Type = 0 (Endpoint)
	cs.WriteU16(0x42, 0x0002) // Version=2, Type=0 (EP)
	// Device Capabilities (cap+4..7)
	cs.WriteU32(0x44, 0x00000000)
	// Device Control (cap+8..9)
	cs.WriteU16(0x48, 0x0000)
	// Device Status (cap+10..11)
	cs.WriteU16(0x4A, 0x0000)
	// Link Capabilities (cap+12..15): Speed=Gen2(2), Width=x4(4)
	// Bits [3:0] = Max Link Speed, Bits [9:4] = Max Link Width
	linkCap := uint32(0x02) | (uint32(0x04) << 4) // Gen2, x4
	cs.WriteU32(0x4C, linkCap)

	ids := ExtractDeviceIDs(cs, nil)

	if !ids.HasPCIeCap {
		t.Fatal("HasPCIeCap should be true")
	}
	if ids.LinkSpeed != LinkSpeedGen2 {
		t.Errorf("LinkSpeed = %d, want %d (Gen2)", ids.LinkSpeed, LinkSpeedGen2)
	}
	if ids.LinkWidth != 4 {
		t.Errorf("LinkWidth = %d, want 4", ids.LinkWidth)
	}
	if ids.PCIeDevType != 0 {
		t.Errorf("PCIeDevType = %d, want 0 (Endpoint)", ids.PCIeDevType)
	}
}

func TestLinkSpeedName(t *testing.T) {
	tests := []struct {
		speed uint8
		want  string
	}{
		{LinkSpeedGen1, "Gen1 (2.5 GT/s)"},
		{LinkSpeedGen2, "Gen2 (5.0 GT/s)"},
		{LinkSpeedGen3, "Gen3 (8.0 GT/s)"},
		{0, "Unknown (0)"},
	}
	for _, tt := range tests {
		got := LinkSpeedName(tt.speed)
		if got != tt.want {
			t.Errorf("LinkSpeedName(%d) = %q, want %q", tt.speed, got, tt.want)
		}
	}
}

func TestExtractDeviceIDsWithDSN(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	cs.WriteU16(0x00, 0x8086)
	cs.WriteU16(0x02, 0x1533)

	// Create a fake DSN extended capability
	dsn := uint64(0xDEADBEEF12345678)
	dsnData := make([]byte, 12)
	// Header (4 bytes)
	binary.LittleEndian.PutUint32(dsnData[0:4], uint32(pci.ExtCapIDDeviceSerialNumber)|0x00010000)
	// DSN value (8 bytes)
	binary.LittleEndian.PutUint64(dsnData[4:12], dsn)

	extCaps := []pci.ExtCapability{
		{ID: pci.ExtCapIDDeviceSerialNumber, Offset: 0x100, Data: dsnData},
	}

	ids := ExtractDeviceIDs(cs, extCaps)

	if !ids.HasDSN {
		t.Error("HasDSN should be true")
	}
	if ids.DSN != dsn {
		t.Errorf("DSN = 0x%016x, want 0x%016x", ids.DSN, dsn)
	}
}

func TestExtractDeviceIDsNoDSNCapTooShort(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize

	// DSN cap with insufficient data
	extCaps := []pci.ExtCapability{
		{ID: pci.ExtCapIDDeviceSerialNumber, Offset: 0x100, Data: make([]byte, 8)},
	}

	ids := ExtractDeviceIDs(cs, extCaps)
	if ids.HasDSN {
		t.Error("HasDSN should be false for truncated DSN data")
	}
}

func TestDSNToSVHex(t *testing.T) {
	tests := []struct {
		dsn  uint64
		want string
	}{
		{0x0000000101000A35, "0000000101000A35"},
		{0xDEADBEEF12345678, "DEADBEEF12345678"},
		{0, "0000000000000000"},
	}

	for _, tt := range tests {
		got := DSNToSVHex(tt.dsn)
		if got != tt.want {
			t.Errorf("DSNToSVHex(0x%016x) = %q, want %q", tt.dsn, got, tt.want)
		}
	}
}
