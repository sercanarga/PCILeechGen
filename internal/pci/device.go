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
	base := d.BaseClass()
	sub := d.SubClass()
	key := uint16(base)<<8 | uint16(sub)

	// Sub-class specific descriptions (matches lspci)
	switch key {
	// 0x01 Mass Storage
	case 0x0101:
		return "IDE interface"
	case 0x0104:
		return "RAID bus controller"
	case 0x0106:
		return "SATA controller"
	case 0x0107:
		return "Serial Attached SCSI controller"
	case 0x0108:
		return "Non-Volatile memory controller"
	// 0x02 Network
	case 0x0200:
		return "Ethernet controller"
	case 0x0280:
		return "Network controller"
	// 0x03 Display
	case 0x0300:
		return "VGA compatible controller"
	case 0x0302:
		return "3D controller"
	// 0x04 Multimedia
	case 0x0400:
		return "Multimedia video controller"
	case 0x0401:
		return "Multimedia audio controller"
	case 0x0403:
		return "Audio device"
	// 0x05 Memory
	case 0x0500:
		return "RAM memory"
	case 0x0580:
		return "Memory controller"
	// 0x06 Bridge
	case 0x0600:
		return "Host bridge"
	case 0x0601:
		return "ISA bridge"
	case 0x0604:
		return "PCI bridge"
	case 0x0680:
		return "Bridge"
	// 0x07 Communication
	case 0x0700:
		return "Serial controller"
	case 0x0780:
		return "Communication controller"
	// 0x08 System Peripheral
	case 0x0800:
		return "PIC"
	case 0x0880:
		return "System peripheral"
	// 0x0C Serial Bus
	case 0x0C03:
		return "USB controller"
	case 0x0C05:
		return "SMBus"
	// 0x0D Wireless
	case 0x0D00:
		return "IRDA controller"
	case 0x0D11:
		return "Bluetooth"
	case 0x0D80:
		return "Wireless controller"
	// 0x11 Signal Processing
	case 0x1180:
		return "Signal processing controller"
	// 0x12 Processing Accelerator
	case 0x1200:
		return "Processing accelerator"
	}

	// Fall back to base class
	switch base {
	case 0x00:
		return "Unclassified device"
	case 0x01:
		return "Mass storage controller"
	case 0x02:
		return "Network controller"
	case 0x03:
		return "Display controller"
	case 0x04:
		return "Multimedia controller"
	case 0x05:
		return "Memory controller"
	case 0x06:
		return "Bridge"
	case 0x07:
		return "Communication controller"
	case 0x08:
		return "System peripheral"
	case 0x09:
		return "Input device controller"
	case 0x0A:
		return "Docking station"
	case 0x0B:
		return "Processor"
	case 0x0C:
		return "Serial bus controller"
	case 0x0D:
		return "Wireless controller"
	case 0x0E:
		return "Intelligent controller"
	case 0x0F:
		return "Satellite communication controller"
	case 0x10:
		return "Encryption controller"
	case 0x11:
		return "Signal processing controller"
	case 0x12:
		return "Processing accelerator"
	case 0xFF:
		return "Unassigned class"
	default:
		return fmt.Sprintf("Class [%02x%02x]", base, sub)
	}
}

// Summary returns a short summary line for display.
func (d *PCIDevice) Summary() string {
	return fmt.Sprintf("%s %04x:%04x [%s] (rev %02x)",
		d.BDF.String(), d.VendorID, d.DeviceID, d.ClassDescription(), d.RevisionID)
}
