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
