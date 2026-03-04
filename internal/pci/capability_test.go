package pci

import (
	"testing"
)

func TestParseCapabilities(t *testing.T) {
	cs := NewConfigSpace()

	// Set capabilities bit in status register
	cs.WriteU16(0x06, 0x0010)
	// Set capability pointer
	cs.WriteU8(0x34, 0x40)

	// First capability: PM at 0x40, next at 0x50
	cs.WriteU8(0x40, CapIDPowerManagement)
	cs.WriteU8(0x41, 0x50) // next pointer

	// Second capability: MSI-X at 0x50, next at 0x70
	cs.WriteU8(0x50, CapIDMSIX)
	cs.WriteU8(0x51, 0x70)

	// Third capability: PCIe at 0x70, no next
	cs.WriteU8(0x70, CapIDPCIExpress)
	cs.WriteU8(0x71, 0x00) // end of list

	caps := ParseCapabilities(cs)

	if len(caps) != 3 {
		t.Fatalf("ParseCapabilities() returned %d caps, want 3", len(caps))
	}

	if caps[0].ID != CapIDPowerManagement {
		t.Errorf("caps[0].ID = 0x%02x, want 0x%02x", caps[0].ID, CapIDPowerManagement)
	}
	if caps[0].Offset != 0x40 {
		t.Errorf("caps[0].Offset = 0x%02x, want 0x40", caps[0].Offset)
	}
	if caps[1].ID != CapIDMSIX {
		t.Errorf("caps[1].ID = 0x%02x, want 0x%02x", caps[1].ID, CapIDMSIX)
	}
	if caps[2].ID != CapIDPCIExpress {
		t.Errorf("caps[2].ID = 0x%02x, want 0x%02x", caps[2].ID, CapIDPCIExpress)
	}
}

func TestParseCapabilitiesNoCaps(t *testing.T) {
	cs := NewConfigSpace()
	// Status register without capabilities bit
	cs.WriteU16(0x06, 0x0000)

	caps := ParseCapabilities(cs)
	if caps != nil {
		t.Errorf("ParseCapabilities() returned %d caps for device without capabilities", len(caps))
	}
}

func TestParseCapabilitiesCircularProtection(t *testing.T) {
	cs := NewConfigSpace()
	cs.WriteU16(0x06, 0x0010) // caps bit set
	cs.WriteU8(0x34, 0x40)

	// Create a circular chain
	cs.WriteU8(0x40, CapIDPowerManagement)
	cs.WriteU8(0x41, 0x40) // points back to itself

	caps := ParseCapabilities(cs)
	if len(caps) != 1 {
		t.Errorf("Circular chain should return 1 cap, got %d", len(caps))
	}
}

func TestParseExtCapabilities(t *testing.T) {
	cs := NewConfigSpace()
	cs.Size = ConfigSpaceSize

	// Extended capability at 0x100: AER, version 1, next at 0x140
	header := uint32(ExtCapIDAER) | (uint32(1) << 16) | (uint32(0x140) << 20)
	cs.WriteU32(0x100, header)

	// Extended capability at 0x140: Device Serial Number, version 1, no next
	header2 := uint32(ExtCapIDDeviceSerialNumber) | (uint32(1) << 16) | (uint32(0) << 20)
	cs.WriteU32(0x140, header2)

	caps := ParseExtCapabilities(cs)

	if len(caps) != 2 {
		t.Fatalf("ParseExtCapabilities() returned %d caps, want 2", len(caps))
	}

	if caps[0].ID != ExtCapIDAER {
		t.Errorf("caps[0].ID = 0x%04x, want 0x%04x", caps[0].ID, ExtCapIDAER)
	}
	if caps[1].ID != ExtCapIDDeviceSerialNumber {
		t.Errorf("caps[1].ID = 0x%04x, want 0x%04x", caps[1].ID, ExtCapIDDeviceSerialNumber)
	}
}

func TestParseExtCapabilitiesSmallConfigSpace(t *testing.T) {
	cs := NewConfigSpace()
	cs.Size = ConfigSpaceLegacySize // Only 256 bytes

	caps := ParseExtCapabilities(cs)
	if caps != nil {
		t.Error("ParseExtCapabilities should return nil for legacy config space")
	}
}

func TestCapabilityNames(t *testing.T) {
	if CapabilityName(CapIDPCIExpress) != "PCI Express" {
		t.Error("CapabilityName for PCIe is wrong")
	}
	if CapabilityName(CapIDMSIX) != "MSI-X" {
		t.Error("CapabilityName for MSI-X is wrong")
	}
	if ExtCapabilityName(ExtCapIDAER) != "Advanced Error Reporting" {
		t.Error("ExtCapabilityName for AER is wrong")
	}
}

func TestCapabilityNameAll(t *testing.T) {
	// Test every standard capability ID returns a non-"Unknown" name
	knownCaps := []uint8{
		CapIDPowerManagement, CapIDAGP, CapIDVPD, CapIDSlotID,
		CapIDMSI, CapIDCompactPCIHotSwap, CapIDPCIX, CapIDHyperTransport,
		CapIDVendorSpecific, CapIDDebugPort, CapIDCompactPCI, CapIDPCIHotPlug,
		CapIDBridgeSubsysVID, CapIDAGP8x, CapIDSecureDevice, CapIDPCIExpress,
		CapIDMSIX, CapIDSATADataIndex, CapIDAdvancedFeatures,
		CapIDEnhancedAlloc, CapIDFlatteningPortal,
	}
	for _, id := range knownCaps {
		name := CapabilityName(id)
		if name == "Unknown" {
			t.Errorf("CapabilityName(0x%02x) = Unknown", id)
		}
		if name == "" {
			t.Errorf("CapabilityName(0x%02x) = empty", id)
		}
	}

	// Unknown ID should return "Unknown"
	if CapabilityName(0xFF) != "Unknown" {
		t.Error("CapabilityName(0xFF) should be Unknown")
	}
}

func TestExtCapabilityNameAll(t *testing.T) {
	knownExtCaps := []uint16{
		ExtCapIDAER, ExtCapIDVCNoMFVC, ExtCapIDDeviceSerialNumber,
		ExtCapIDPowerBudgeting, ExtCapIDRCLinkDeclaration,
		ExtCapIDVendorSpecific, ExtCapIDACS, ExtCapIDARI,
		ExtCapIDATS, ExtCapIDSRIOV, ExtCapIDResizableBAR,
		ExtCapIDLTR, ExtCapIDSecondaryPCIe, ExtCapIDL1PMSubstates,
		ExtCapIDPTM, ExtCapIDDPC, ExtCapIDPASID,
	}
	for _, id := range knownExtCaps {
		name := ExtCapabilityName(id)
		if name == "Unknown" {
			t.Errorf("ExtCapabilityName(0x%04x) = Unknown", id)
		}
		if name == "" {
			t.Errorf("ExtCapabilityName(0x%04x) = empty", id)
		}
	}

	// Unknown ID
	if ExtCapabilityName(0xFFFF) != "Unknown" {
		t.Error("ExtCapabilityName(0xFFFF) should be Unknown")
	}
}

func TestParseExtCapabilitiesInvalidHeaders(t *testing.T) {
	cs := NewConfigSpace()
	cs.Size = ConfigSpaceSize

	// All-0xFF header should be skipped
	cs.WriteU32(0x100, 0xFFFFFFFF)
	caps := ParseExtCapabilities(cs)
	if len(caps) != 0 {
		t.Errorf("Expected 0 caps for 0xFFFFFFFF header, got %d", len(caps))
	}

	// All-zero header should be skipped
	cs.WriteU32(0x100, 0x00000000)
	caps = ParseExtCapabilities(cs)
	if len(caps) != 0 {
		t.Errorf("Expected 0 caps for 0x00000000 header, got %d", len(caps))
	}
}
