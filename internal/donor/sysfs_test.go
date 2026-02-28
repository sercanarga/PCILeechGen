package donor

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/pci"
)

// createMockSysfs creates a mock sysfs directory for testing.
func createMockSysfs(t *testing.T) string {
	t.Helper()
	base := t.TempDir()

	// Create a mock device: 0000:03:00.0
	devDir := filepath.Join(base, "0000:03:00.0")
	if err := os.MkdirAll(devDir, 0755); err != nil {
		t.Fatal(err)
	}

	// Write mock device files
	writeFile(t, devDir, "vendor", "0x8086\n")
	writeFile(t, devDir, "device", "0x1533\n")
	writeFile(t, devDir, "class", "0x020000\n")
	writeFile(t, devDir, "subsystem_vendor", "0x8086\n")
	writeFile(t, devDir, "subsystem_device", "0x0001\n")
	writeFile(t, devDir, "revision", "0x03\n")

	// Write mock config space (256 bytes)
	configData := make([]byte, 256)
	configData[0] = 0x86    // Vendor ID low
	configData[1] = 0x80    // Vendor ID high
	configData[2] = 0x33    // Device ID low
	configData[3] = 0x15    // Device ID high
	configData[6] = 0x10    // Status: capabilities list
	configData[8] = 0x03    // Revision ID
	configData[0x0B] = 0x02 // Base class (Network)
	if err := os.WriteFile(filepath.Join(devDir, "config"), configData, 0644); err != nil {
		t.Fatal(err)
	}

	// Write mock resource file
	resourceContent := `0x00000000fe000000 0x00000000fe0fffff 0x00040200
0x0000000000001000 0x000000000000103f 0x00040101
0x0000000000000000 0x0000000000000000 0x00000000
0x0000000000000000 0x0000000000000000 0x00000000
0x0000000000000000 0x0000000000000000 0x00000000
0x0000000000000000 0x0000000000000000 0x00000000
`
	writeFile(t, devDir, "resource", resourceContent)

	return base
}

func writeFile(t *testing.T, dir, name, content string) {
	t.Helper()
	if err := os.WriteFile(filepath.Join(dir, name), []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
}

func TestSysfsReaderScanDevices(t *testing.T) {
	base := createMockSysfs(t)
	sr := NewSysfsReaderWithPath(base)

	devices, err := sr.ScanDevices()
	if err != nil {
		t.Fatal(err)
	}

	if len(devices) != 1 {
		t.Fatalf("ScanDevices() returned %d devices, want 1", len(devices))
	}

	dev := devices[0]
	if dev.VendorID != 0x8086 {
		t.Errorf("VendorID = 0x%04x, want 0x8086", dev.VendorID)
	}
	if dev.DeviceID != 0x1533 {
		t.Errorf("DeviceID = 0x%04x, want 0x1533", dev.DeviceID)
	}
	if dev.ClassCode != 0x020000 {
		t.Errorf("ClassCode = 0x%06x, want 0x020000", dev.ClassCode)
	}
}

func TestSysfsReaderReadConfigSpace(t *testing.T) {
	base := createMockSysfs(t)
	sr := NewSysfsReaderWithPath(base)

	bdf := pci.BDF{Domain: 0, Bus: 3, Device: 0, Function: 0}
	cs, err := sr.ReadConfigSpace(bdf)
	if err != nil {
		t.Fatal(err)
	}

	if cs.VendorID() != 0x8086 {
		t.Errorf("VendorID = 0x%04x, want 0x8086", cs.VendorID())
	}
	if cs.DeviceID() != 0x1533 {
		t.Errorf("DeviceID = 0x%04x, want 0x1533", cs.DeviceID())
	}
}

func TestSysfsReaderReadResource(t *testing.T) {
	base := createMockSysfs(t)
	sr := NewSysfsReaderWithPath(base)

	bdf := pci.BDF{Domain: 0, Bus: 3, Device: 0, Function: 0}
	bars, err := sr.ReadResourceFile(bdf)
	if err != nil {
		t.Fatal(err)
	}

	if len(bars) < 2 {
		t.Fatalf("ReadResourceFile returned %d BARs, want at least 2", len(bars))
	}

	if bars[0].Address != 0xFE000000 {
		t.Errorf("BAR0 address = 0x%x, want 0xFE000000", bars[0].Address)
	}
	if bars[0].Size != 0x100000 {
		t.Errorf("BAR0 size = 0x%x, want 0x100000", bars[0].Size)
	}
}

func TestCollectorWithMockSysfs(t *testing.T) {
	base := createMockSysfs(t)
	sr := NewSysfsReaderWithPath(base)
	c := NewCollectorWithSysfs(sr)

	bdf := pci.BDF{Domain: 0, Bus: 3, Device: 0, Function: 0}
	ctx, err := c.Collect(bdf)
	if err != nil {
		t.Fatal(err)
	}

	if ctx.Device.VendorID != 0x8086 {
		t.Errorf("Device.VendorID = 0x%04x, want 0x8086", ctx.Device.VendorID)
	}
	if ctx.ConfigSpace == nil {
		t.Fatal("ConfigSpace is nil")
	}
	if ctx.Hostname == "" {
		t.Error("Hostname is empty")
	}
}

func TestDeviceContextJSONRoundtrip(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = 256
	cs.WriteU16(0x00, 0x8086)
	cs.WriteU16(0x02, 0x1533)

	ctx := &DeviceContext{
		Device: pci.PCIDevice{
			BDF:      pci.BDF{Domain: 0, Bus: 3, Device: 0, Function: 0},
			VendorID: 0x8086,
			DeviceID: 0x1533,
		},
		ConfigSpace: cs,
		BARs: []pci.BAR{
			{Index: 0, Type: pci.BARTypeMem32, Address: 0xFE000000, Size: 1048576},
		},
	}

	jsonData, err := ctx.ToJSON()
	if err != nil {
		t.Fatal(err)
	}

	loaded, err := FromJSON(jsonData)
	if err != nil {
		t.Fatal(err)
	}

	if loaded.Device.VendorID != 0x8086 {
		t.Errorf("roundtrip VendorID = 0x%04x, want 0x8086", loaded.Device.VendorID)
	}
	if loaded.ConfigSpace.VendorID() != 0x8086 {
		t.Errorf("roundtrip ConfigSpace.VendorID = 0x%04x, want 0x8086", loaded.ConfigSpace.VendorID())
	}
}
