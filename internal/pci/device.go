// Package pci defines PCI/PCIe device types and config space accessors.
package pci

import (
	"fmt"
	"strings"
)

// BDF represents a PCI Bus:Device.Function address.
type BDF struct {
	Domain   uint16
	Bus      uint8
	Device   uint8
	Function uint8
}

// ParseBDF parses a BDF string in the format "DDDD:BB:DD.F" or "BB:DD.F".
func ParseBDF(s string) (BDF, error) {
	s = strings.TrimSpace(s)
	var bdf BDF

	// Try full format: DDDD:BB:DD.F
	n, err := fmt.Sscanf(s, "%x:%x:%x.%x", &bdf.Domain, &bdf.Bus, &bdf.Device, &bdf.Function)
	if err == nil && n == 4 {
		return bdf, nil
	}

	// Try short format: BB:DD.F (domain defaults to 0)
	n, err = fmt.Sscanf(s, "%x:%x.%x", &bdf.Bus, &bdf.Device, &bdf.Function)
	if err == nil && n == 3 {
		bdf.Domain = 0
		return bdf, nil
	}

	return BDF{}, fmt.Errorf("invalid BDF format %q: expected DDDD:BB:DD.F or BB:DD.F", s)
}

// String returns the canonical BDF representation: "DDDD:BB:DD.F".
func (b BDF) String() string {
	return fmt.Sprintf("%04x:%02x:%02x.%x", b.Domain, b.Bus, b.Device, b.Function)
}

// Short returns the short BDF representation without domain: "BB:DD.F".
func (b BDF) Short() string {
	return fmt.Sprintf("%02x:%02x.%x", b.Bus, b.Device, b.Function)
}

// SysfsPath returns the sysfs path for this device.
func (b BDF) SysfsPath() string {
	return fmt.Sprintf("/sys/bus/pci/devices/%s", b.String())
}

// PCIDevice holds all discovered information about a PCI device.
type PCIDevice struct {
	BDF            BDF    `json:"bdf"`
	VendorID       uint16 `json:"vendor_id"`
	DeviceID       uint16 `json:"device_id"`
	SubsysVendorID uint16 `json:"subsys_vendor_id"`
	SubsysDeviceID uint16 `json:"subsys_device_id"`
	RevisionID     uint8  `json:"revision_id"`
	ClassCode      uint32 `json:"class_code"` // 24-bit: base_class << 16 | sub_class << 8 | prog_if
	HeaderType     uint8  `json:"header_type"`
	Driver         string `json:"driver,omitempty"`
	IOMMUGroup     int    `json:"iommu_group,omitempty"`
}

// BaseClass returns the PCI base class code.
func (d *PCIDevice) BaseClass() uint8 {
	return uint8((d.ClassCode >> 16) & 0xFF)
}

// SubClass returns the PCI sub-class code.
func (d *PCIDevice) SubClass() uint8 {
	return uint8((d.ClassCode >> 8) & 0xFF)
}

// ProgIF returns the PCI programming interface.
func (d *PCIDevice) ProgIF() uint8 {
	return uint8(d.ClassCode & 0xFF)
}

// ClassDescription returns a human-readable description matching lspci style.
func (d *PCIDevice) ClassDescription() string {
	key := uint16(d.BaseClass())<<8 | uint16(d.SubClass())
	if name, ok := pciSubClassNames[key]; ok {
		return name
	}
	if name, ok := pciBaseClassNames[d.BaseClass()]; ok {
		return name
	}
	return fmt.Sprintf("Class [%02x%02x]", d.BaseClass(), d.SubClass())
}

// Summary returns a short summary line for display.
func (d *PCIDevice) Summary() string {
	return fmt.Sprintf("%s %04x:%04x [%s] (rev %02x)",
		d.BDF.String(), d.VendorID, d.DeviceID, d.ClassDescription(), d.RevisionID)
}
