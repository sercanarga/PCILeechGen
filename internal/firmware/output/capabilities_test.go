package output

import (
	"testing"

	"github.com/sercanarga/pcileechgen/internal/pci"
)

func TestExtractDonorCapabilities_FromStandardAndExtendedCaps(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	cs.WriteU16(0x06, 0x0010)
	cs.WriteU8(0x34, 0x40)

	cs.WriteU8(0x40, pci.CapIDPowerManagement)
	cs.WriteU8(0x41, 0x50)
	cs.WriteU16(0x42, 0xA800)
	cs.WriteU16(0x44, 0x0100)

	cs.WriteU8(0x50, pci.CapIDMSI)
	cs.WriteU8(0x51, 0x60)
	cs.WriteU16(0x52, 0x00F0)

	cs.WriteU8(0x60, pci.CapIDMSIX)
	cs.WriteU8(0x61, 0x70)

	cs.WriteU8(0x70, pci.CapIDPCIExpress)
	cs.WriteU8(0x71, 0x00)
	cs.WriteU32(0x7C, 0x00000C43)
	cs.WriteU16(0x80, 0x0002)

	cs.WriteU32(0x100, makeExtCapHeader(pci.ExtCapIDAER, 0x140))
	cs.WriteU32(0x140, makeExtCapHeader(pci.ExtCapIDLTR, 0x180))
	cs.WriteU32(0x180, makeExtCapHeader(pci.ExtCapIDL1PMSubstates, 0x1C0))
	cs.WriteU32(0x1C0, makeExtCapHeader(pci.ExtCapIDDeviceSerialNumber, 0))

	caps := extractDonorCapabilities(cs)

	if !caps.HasPMCap {
		t.Fatal("PM capability should be detected")
	}
	if caps.PMESupportMask != 0x15 {
		t.Errorf("PMESupportMask = 0x%02x, want 0x15", caps.PMESupportMask)
	}
	if !caps.PMEDefault {
		t.Error("PMEDefault should be true")
	}
	if !caps.HasMSICap {
		t.Fatal("MSI capability should be detected")
	}
	if caps.MSIDisable64Bit {
		t.Error("MSIDisable64Bit should be false")
	}
	if caps.MSIMultipleMsg != 0x07 {
		t.Errorf("MSIMultipleMsg = 0x%02x, want 0x07", caps.MSIMultipleMsg)
	}
	if !caps.HasMSIXCap {
		t.Error("MSI-X capability should be detected")
	}
	if !caps.HasPCIeCap {
		t.Fatal("PCIe capability should be detected")
	}
	if caps.PCIELinkSpeed != 0x03 {
		t.Errorf("PCIELinkSpeed = 0x%02x, want 0x03", caps.PCIELinkSpeed)
	}
	if caps.PCIELinkWidth != 0x04 {
		t.Errorf("PCIELinkWidth = 0x%02x, want 0x04", caps.PCIELinkWidth)
	}
	if caps.PCIeASPMCap != 0x03 {
		t.Errorf("PCIeASPMCap = 0x%02x, want 0x03", caps.PCIeASPMCap)
	}
	if caps.PCIeASPMEnable != 0x02 {
		t.Errorf("PCIeASPMEnable = 0x%02x, want 0x02", caps.PCIeASPMEnable)
	}
	if !caps.HasAERCap || !caps.HasLTRCap || !caps.HasL1PMSubstates || !caps.HasDSNCap {
		t.Errorf("extended caps missing: %+v", caps)
	}
}

func makeExtCapHeader(id uint16, next int) uint32 {
	return uint32(id) | (uint32(1) << 16) | (uint32(next&0xFFC) << 20)
}
