package firmware

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

func TestOutputWriterWriteAll(t *testing.T) {
	outputDir := t.TempDir()
	libDir := "/fake/lib/pcileech-fpga"

	ctx := makeTestContext()
	b, _ := board.Find("PCIeSquirrel")

	ow := NewOutputWriter(outputDir, libDir)
	if err := ow.WriteAll(ctx, b); err != nil {
		t.Fatalf("WriteAll() error: %v", err)
	}

	// Verify all expected files were created
	expectedFiles := ListOutputFiles()
	for _, name := range expectedFiles {
		path := filepath.Join(outputDir, name)
		info, err := os.Stat(path)
		if err != nil {
			t.Errorf("Expected file %q not found: %v", name, err)
			continue
		}
		if info.Size() == 0 {
			t.Errorf("File %q is empty", name)
		}
	}

	// Verify device_context.json is valid JSON
	jsonData, err := os.ReadFile(filepath.Join(outputDir, "device_context.json"))
	if err != nil {
		t.Fatal(err)
	}
	loaded, err := donor.FromJSON(jsonData)
	if err != nil {
		t.Fatalf("device_context.json is not valid: %v", err)
	}
	if loaded.Device.VendorID != 0x8086 {
		t.Errorf("Loaded VendorID = 0x%04x, want 0x8086", loaded.Device.VendorID)
	}
}

func TestOutputWriterBadDir(t *testing.T) {
	// Try writing to an invalid path
	ow := NewOutputWriter("/dev/null/impossible/path", "/fake")

	ctx := makeTestContext()
	b, _ := board.Find("PCIeSquirrel")

	err := ow.WriteAll(ctx, b)
	if err == nil {
		t.Error("Expected error for invalid output directory")
	}
}

func makeTestContextFull() *donor.DeviceContext {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize

	cs.WriteU16(0x00, 0x8086)
	cs.WriteU16(0x02, 0x1533)
	cs.WriteU16(0x04, 0x0406)
	cs.WriteU16(0x06, 0x0010) // Status: caps
	cs.WriteU8(0x08, 0x03)
	cs.WriteU8(0x0B, 0x02)
	cs.WriteU32(0x10, 0xFE000000)
	cs.WriteU16(0x2C, 0x8086)
	cs.WriteU16(0x2E, 0x0001)

	return &donor.DeviceContext{
		Device: pci.PCIDevice{
			BDF:            pci.BDF{Domain: 0, Bus: 3, Device: 0, Function: 0},
			VendorID:       0x8086,
			DeviceID:       0x1533,
			SubsysVendorID: 0x8086,
			SubsysDeviceID: 0x0001,
			RevisionID:     0x03,
			ClassCode:      0x020000,
		},
		ConfigSpace: cs,
		BARs: []pci.BAR{
			{Index: 0, Type: pci.BARTypeMem32, Address: 0xFE000000, Size: 1048576},
		},
		Capabilities: pci.ParseCapabilities(cs),
	}
}
