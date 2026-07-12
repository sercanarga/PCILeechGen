package codegen

import (
	"path/filepath"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/firmware/scrub"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

func loadAndScrubDonor(t *testing.T, fixture string) *pci.ConfigSpace {
	t.Helper()
	ctx, err := donor.LoadContext(filepath.Join("..", "..", "..", "testdata", "donors", fixture))
	if err != nil {
		t.Fatalf("LoadContext %s: %v", fixture, err)
	}
	scrubbed, _ := scrub.ScrubConfigSpaceWithDonor(ctx.ConfigSpace, nil, ctx.BARs, ctx.MSIXData, 16384)
	return scrubbed
}

func TestPipeline_NVMeIdentity(t *testing.T) {
	cs := loadAndScrubDonor(t, "nvme.json")
	if vid := cs.VendorID(); vid != 0x144D {
		t.Errorf("VID: got 0x%04X, want 0x144D", vid)
	}
	if did := cs.DeviceID(); did != 0xA809 {
		t.Errorf("DID: got 0x%04X, want 0xA809", did)
	}
	sv := cs.SubsysVendorID()
	sd := cs.SubsysDeviceID()
	if sv == 0 {
		t.Error("subsys vendor ID = 0, should be non-zero (Code-10: Windows INF match)")
	}
	if sd == 0 {
		t.Error("subsys device ID = 0, should be non-zero (Code-10: Windows INF match)")
	}
	if class := cs.ReadU32(0x08) >> 8; class != 0x010802 {
		t.Errorf("class code: got 0x%06X, want 0x010802 (NVMe)", class)
	}
}

func TestPipeline_NVMeMSIXCapInjected(t *testing.T) {
	cs := loadAndScrubDonor(t, "nvme.json")
	info := pci.ParseMSIXCap(cs)
	if info == nil {
		t.Fatal("MSI-X capability not found in scrubbed config space")
	}
	if info.TableSize != 5 {
		t.Errorf("MSI-X table size: got %d, want 5", info.TableSize)
	}
	if info.TableOffset == 0 {
		t.Error("MSI-X table offset = 0, would overlap NVMe CAP registers")
	}
}

func TestPipeline_NVMeCapabilityChain(t *testing.T) {
	cs := loadAndScrubDonor(t, "nvme.json")
	caps := pci.ParseCapabilities(cs)
	if len(caps) < 3 {
		t.Fatalf("expected >=3 capabilities, got %d", len(caps))
	}
	hasPM, hasPCIe, hasMSIX := false, false, false
	for _, c := range caps {
		switch c.ID {
		case pci.CapIDPowerManagement:
			hasPM = true
		case pci.CapIDPCIExpress:
			hasPCIe = true
		case pci.CapIDMSIX:
			hasMSIX = true
		}
	}
	if !hasPM {
		t.Error("PM capability missing")
	}
	if !hasPCIe {
		t.Error("PCIe capability missing")
	}
	if !hasMSIX {
		t.Error("MSI-X capability missing (inject failed)")
	}
}

func TestPipeline_AllDonorsIdentity(t *testing.T) {
	donors := []struct {
		name  string
		vid   uint16
		did   uint16
		class uint32
	}{
		{"audio", 0x8086, 0x9D71, 0x040300},
		{"ethernet", 0x8086, 0x15B7, 0x020000},
		{"gpu", 0x10DE, 0x1B06, 0x030000},
		{"multibar", 0x1234, 0x5678, 0x000000},
		{"nvme", 0x144D, 0xA809, 0x010802},
		{"sata", 0x8086, 0x9D03, 0x010601},
		{"thunderbolt", 0x8086, 0x15D9, 0x080700},
		{"wifi", 0x8086, 0x24FD, 0x028000},
		{"xhci", 0x8086, 0x9D2F, 0x0C0330},
		{"generic", 0x1234, 0x5678, 0x000000},
	}
	for _, d := range donors {
		t.Run(d.name, func(t *testing.T) {
			cs := loadAndScrubDonor(t, d.name+".json")
			if vid := cs.VendorID(); vid != d.vid {
				t.Errorf("VID: got 0x%04X, want 0x%04X", vid, d.vid)
			}
			if did := cs.DeviceID(); did != d.did {
				t.Errorf("DID: got 0x%04X, want 0x%04X", did, d.did)
			}
			if sv := cs.SubsysVendorID(); sv == 0 {
				t.Error("subsys vendor ID = 0 after scrub")
			}
			if sd := cs.SubsysDeviceID(); sd == 0 {
				t.Error("subsys device ID = 0 after scrub")
			}
			if d.class > 0 {
				if class := cs.ReadU32(0x08) >> 8; class != d.class {
					t.Errorf("class: got 0x%06X, want 0x%06X", class, d.class)
				}
			}
			caps := pci.ParseCapabilities(cs)
			if len(caps) == 0 {
				t.Error("no capabilities in scrubbed config space")
			}
		})
	}
}

func TestPipeline_MultibarTopology(t *testing.T) {
	cs := loadAndScrubDonor(t, "multibar.json")
	for _, off := range []int{0x10, 0x14, 0x18, 0x1C, 0x20, 0x24} {
		bar := cs.ReadU32(off)
		if bar != 0 {
			t.Logf("BAR at 0x%02X = 0x%08X", off, bar)
		}
	}
	bar0 := cs.ReadU32(0x10)
	if bar0 == 0 {
		t.Error("BAR0 = 0, expected non-zero (multibar fixture should have BAR0)")
	}
}

func TestPipeline_EthernetNoMSIX(t *testing.T) {
	cs := loadAndScrubDonor(t, "ethernet.json")
	info := pci.ParseMSIXCap(cs)
	if info != nil {
		t.Error("ethernet donor should not have MSI-X capability injected (no msix_data in donor JSON)")
	}
	caps := pci.ParseCapabilities(cs)
	hasMSI := false
	for _, c := range caps {
		if c.ID == pci.CapIDMSI {
			hasMSI = true
		}
	}
	if !hasMSI {
		t.Error("ethernet donor should retain MSI capability")
	}
}
