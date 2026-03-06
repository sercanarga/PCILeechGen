package pci

import "testing"

func TestCapabilityName_AllKnown(t *testing.T) {
	known := map[uint8]string{
		CapIDPowerManagement: "Power Management",
		CapIDPCIExpress:      "PCI Express",
		CapIDMSIX:            "MSI-X",
		CapIDMSI:             "MSI",
		CapIDVendorSpecific:  "Vendor Specific",
	}
	for id, want := range known {
		got := CapabilityName(id)
		if got != want {
			t.Errorf("CapabilityName(0x%02x) = %q, want %q", id, got, want)
		}
	}
}

func TestCapabilityName_Unknown(t *testing.T) {
	got := CapabilityName(0xFF)
	if got != "Unknown" {
		t.Errorf("CapabilityName(0xFF) = %q, want 'Unknown'", got)
	}
}

func TestExtCapabilityName_AllKnown(t *testing.T) {
	known := map[uint16]string{
		ExtCapIDAER:                "Advanced Error Reporting",
		ExtCapIDDeviceSerialNumber: "Device Serial Number",
		ExtCapIDSRIOV:              "Single Root I/O Virtualization",
		ExtCapIDLTR:                "Latency Tolerance Reporting",
	}
	for id, want := range known {
		got := ExtCapabilityName(id)
		if got != want {
			t.Errorf("ExtCapabilityName(0x%04x) = %q, want %q", id, got, want)
		}
	}
}

func TestExtCapabilityName_Unknown(t *testing.T) {
	got := ExtCapabilityName(0xFFFF)
	if got != "Unknown" {
		t.Errorf("ExtCapabilityName(0xFFFF) = %q, want 'Unknown'", got)
	}
}

func TestClassDescription_SubClass(t *testing.T) {
	dev := &PCIDevice{ClassCode: 0x020000} // Ethernet
	desc := dev.ClassDescription()
	if desc != "Ethernet controller" {
		t.Errorf("ClassDescription = %q, want 'Ethernet controller'", desc)
	}
}

func TestClassDescription_BaseClass(t *testing.T) {
	dev := &PCIDevice{ClassCode: 0x020100} // unknown subclass of network
	desc := dev.ClassDescription()
	if desc != "Network controller" {
		t.Errorf("ClassDescription = %q, want 'Network controller' (base class fallback)", desc)
	}
}

func TestClassDescription_Unknown(t *testing.T) {
	dev := &PCIDevice{ClassCode: 0xFE0000} // completely unknown
	desc := dev.ClassDescription()
	if desc == "" {
		t.Error("ClassDescription should not be empty for unknown class")
	}
}

func TestPCIDeviceSummary_Format(t *testing.T) {
	dev := &PCIDevice{
		BDF:        BDF{Domain: 0, Bus: 3, Device: 0, Function: 0},
		VendorID:   0x8086,
		DeviceID:   0x1533,
		ClassCode:  0x020000,
		RevisionID: 0x03,
	}
	summary := dev.Summary()
	if summary == "" {
		t.Error("Summary should not be empty")
	}
}
