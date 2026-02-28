package firmware

import (
	"encoding/binary"
	"fmt"

	"github.com/sercanarga/pcileechgen/internal/pci"
)

// PCIe Link Speed constants
const (
	LinkSpeedGen1 uint8 = 1 // 2.5 GT/s
	LinkSpeedGen2 uint8 = 2 // 5.0 GT/s
	LinkSpeedGen3 uint8 = 3 // 8.0 GT/s
)

// DeviceIDs holds all device identification values needed for SV patching.
type DeviceIDs struct {
	VendorID       uint16
	DeviceID       uint16
	SubsysVendorID uint16
	SubsysDeviceID uint16
	RevisionID     uint8
	ClassCode      uint32 // 24-bit: base<<16 | sub<<8 | progif
	DSN            uint64 // Device Serial Number (0 if not present)
	HasDSN         bool

	// PCIe Link Capability fields
	LinkSpeed   uint8 // Supported Link Speed (1=Gen1, 2=Gen2, 3=Gen3)
	LinkWidth   uint8 // Maximum Link Width (1, 2, 4, 8, 16)
	HasPCIeCap  bool
	PCIeDevType uint8 // PCIe Device/Port Type (from PCIe Capabilities Register)
}

// ExtractDeviceIDs collects all device identification from a config space and capabilities.
func ExtractDeviceIDs(cs *pci.ConfigSpace, extCaps []pci.ExtCapability) DeviceIDs {
	ids := DeviceIDs{
		VendorID:       cs.VendorID(),
		DeviceID:       cs.DeviceID(),
		SubsysVendorID: cs.SubsysVendorID(),
		SubsysDeviceID: cs.SubsysDeviceID(),
		RevisionID:     cs.RevisionID(),
		ClassCode:      cs.ClassCode(),
	}

	// Extract PCIe capability (link speed/width)
	caps := pci.ParseCapabilities(cs)
	for _, cap := range caps {
		if cap.ID == pci.CapIDPCIExpress && len(cap.Data) >= 16 {
			ids.HasPCIeCap = true

			// PCIe Capabilities Register at cap+2 (offset in data: bytes 2-3)
			pcieCapReg := binary.LittleEndian.Uint16(cap.Data[2:4])
			ids.PCIeDevType = uint8((pcieCapReg >> 4) & 0x0F)

			// Link Capabilities Register at cap+12 (offset in data: bytes 12-15)
			linkCap := binary.LittleEndian.Uint32(cap.Data[12:16])
			ids.LinkSpeed = uint8(linkCap & 0x0F)        // Max Link Speed
			ids.LinkWidth = uint8((linkCap >> 4) & 0x3F) // Max Link Width
			break
		}
	}

	// Extract DSN from extended capabilities
	for _, cap := range extCaps {
		if cap.ID == pci.ExtCapIDDeviceSerialNumber && len(cap.Data) >= 12 {
			ids.DSN = binary.LittleEndian.Uint64(cap.Data[4:12])
			ids.HasDSN = true
			break
		}
	}

	return ids
}

// LinkSpeedName returns a human-readable name for PCIe link speed.
func LinkSpeedName(speed uint8) string {
	switch speed {
	case LinkSpeedGen1:
		return "Gen1 (2.5 GT/s)"
	case LinkSpeedGen2:
		return "Gen2 (5.0 GT/s)"
	case LinkSpeedGen3:
		return "Gen3 (8.0 GT/s)"
	default:
		return fmt.Sprintf("Unknown (%d)", speed)
	}
}

// DSNToSVHex formats a 64-bit DSN for SystemVerilog (16-char hex string).
func DSNToSVHex(dsn uint64) string {
	return fmt.Sprintf("%016X", dsn)
}
