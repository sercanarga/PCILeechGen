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

// pciSubClassNames maps (base_class << 8 | sub_class) to human-readable names.
var pciSubClassNames = map[uint16]string{
	// Mass Storage
	0x0101: "IDE interface",
	0x0104: "RAID bus controller",
	0x0106: "SATA controller",
	0x0107: "Serial Attached SCSI controller",
	0x0108: "Non-Volatile memory controller",
	// Network
	0x0200: "Ethernet controller",
	0x0280: "Network controller",
	// Display
	0x0300: "VGA compatible controller",
	0x0302: "3D controller",
	// Multimedia
	0x0400: "Multimedia video controller",
	0x0401: "Multimedia audio controller",
	0x0403: "Audio device",
	// Memory
	0x0500: "RAM memory",
	0x0580: "Memory controller",
	// Bridge
	0x0600: "Host bridge",
	0x0601: "ISA bridge",
	0x0604: "PCI bridge",
	0x0680: "Bridge",
	// Communication
	0x0700: "Serial controller",
	0x0780: "Communication controller",
	// System Peripheral
	0x0800: "PIC",
	0x0880: "System peripheral",
	// Serial Bus
	0x0C03: "USB controller",
	0x0C05: "SMBus",
	// Wireless
	0x0D00: "IRDA controller",
	0x0D11: "Bluetooth",
	0x0D80: "Wireless controller",
	// Signal Processing
	0x1180: "Signal processing controller",
	// Processing Accelerator
	0x1200: "Processing accelerator",
}

// pciBaseClassNames maps base_class to a fallback human-readable name.
var pciBaseClassNames = map[uint8]string{
	0x00: "Unclassified device",
	0x01: "Mass storage controller",
	0x02: "Network controller",
	0x03: "Display controller",
	0x04: "Multimedia controller",
	0x05: "Memory controller",
	0x06: "Bridge",
	0x07: "Communication controller",
	0x08: "System peripheral",
	0x09: "Input device controller",
	0x0A: "Docking station",
	0x0B: "Processor",
	0x0C: "Serial bus controller",
	0x0D: "Wireless controller",
	0x0E: "Intelligent controller",
	0x0F: "Satellite communication controller",
	0x10: "Encryption controller",
	0x11: "Signal processing controller",
	0x12: "Processing accelerator",
	0xFF: "Unassigned class",
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
