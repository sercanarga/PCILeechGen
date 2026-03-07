package firmware

import (
	"testing"

	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

func TestLowestBar_ByteSlice_EmptyMap(t *testing.T) {
	result := LowestBar(map[int][]byte(nil))
	if result != nil {
		t.Error("nil map should return nil")
	}

	result = LowestBar(map[int][]byte{})
	if result != nil {
		t.Error("empty map should return nil")
	}
}

func TestLowestBar_ByteSlice_SingleEntry(t *testing.T) {
	data := []byte{0x01, 0x02, 0x03, 0x04}
	result := LowestBar(map[int][]byte{2: data})
	if result == nil || len(result) != 4 {
		t.Error("should return single entry data")
	}
}

func TestLowestBar_ByteSlice_MultipleEntries(t *testing.T) {
	bar0 := []byte{0xAA}
	bar2 := []byte{0xBB}
	bar4 := []byte{0xCC}

	result := LowestBar(map[int][]byte{4: bar4, 0: bar0, 2: bar2})
	if result == nil || result[0] != 0xAA {
		t.Errorf("should pick BAR0 (lowest index), got %v", result)
	}
}

func TestLowestBar_Profile_EmptyMap(t *testing.T) {
	result := LowestBar(map[int]*donor.BARProfile(nil))
	if result != nil {
		t.Error("nil map should return nil")
	}
}

func TestLowestBar_Profile_SingleEntry(t *testing.T) {
	p := &donor.BARProfile{}
	result := LowestBar(map[int]*donor.BARProfile{1: p})
	if result == nil {
		t.Error("should return the single profile")
	}
}

func TestLowestBar_Profile_PicksLowest(t *testing.T) {
	p0 := &donor.BARProfile{}
	p2 := &donor.BARProfile{}

	result := LowestBar(map[int]*donor.BARProfile{2: p2, 0: p0})
	if result != p0 {
		t.Error("should pick BAR0 (lowest index)")
	}
}

func TestLargestBar_EmptyMap(t *testing.T) {
	if LargestBar(nil) != nil {
		t.Error("nil map should return nil")
	}
	if LargestBar(map[int][]byte{}) != nil {
		t.Error("empty map should return nil")
	}
}

func TestLargestBar_SingleEntry(t *testing.T) {
	data := []byte{0x01, 0x02, 0x03}
	result := LargestBar(map[int][]byte{2: data})
	if len(result) != 3 {
		t.Errorf("should return single entry, got len=%d", len(result))
	}
}

func TestLargestBar_PicksLargest(t *testing.T) {
	bar0 := make([]byte, 256)   // IO BAR (small)
	bar2 := make([]byte, 65536) // MMIO BAR (large)
	bar0[0] = 0xAA
	bar2[0] = 0xBB

	result := LargestBar(map[int][]byte{0: bar0, 2: bar2})
	if len(result) != 65536 || result[0] != 0xBB {
		t.Errorf("should pick BAR2 (largest), got len=%d first=0x%02X", len(result), result[0])
	}
}

func TestExtractDeviceIDs_ClassCode(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	cs.WriteU16(0x00, 0x144D) // VendorID
	cs.WriteU16(0x02, 0xA808) // DeviceID
	cs.WriteU8(0x09, 0x02)    // ProgIF
	cs.WriteU8(0x0A, 0x08)    // SubClass
	cs.WriteU8(0x0B, 0x01)    // BaseClass

	ids := ExtractDeviceIDs(cs, nil)

	if ids.VendorID != 0x144D {
		t.Errorf("VendorID: got 0x%04X, want 0x144D", ids.VendorID)
	}
	if ids.DeviceID != 0xA808 {
		t.Errorf("DeviceID: got 0x%04X, want 0xA808", ids.DeviceID)
	}
	if ids.ClassCode != 0x010802 {
		t.Errorf("ClassCode: got 0x%06X, want 0x010802", ids.ClassCode)
	}
}

func TestExtractDeviceIDs_SubsystemIDs(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	cs.WriteU16(0x00, 0x8086)
	cs.WriteU16(0x02, 0x1533)
	cs.WriteU16(0x2C, 0x1234) // Subsystem Vendor
	cs.WriteU16(0x2E, 0x5678) // Subsystem Device

	ids := ExtractDeviceIDs(cs, nil)
	if ids.SubsysVendorID != 0x1234 {
		t.Errorf("SubsysVendorID: got 0x%04X, want 0x1234", ids.SubsysVendorID)
	}
	if ids.SubsysDeviceID != 0x5678 {
		t.Errorf("SubsysDeviceID: got 0x%04X, want 0x5678", ids.SubsysDeviceID)
	}
}
