package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

func TestBuildDonorReportIncludesIdentityCapabilitiesBARsAndMasks(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.WriteU16(0x06, 0x0010)
	cs.WriteU8(0x34, 0x40)
	cs.WriteU8(0x40, pci.CapIDPowerManagement)
	ctx := &donor.DeviceContext{
		Device:       pci.PCIDevice{BDF: pci.BDF{Bus: 2}, VendorID: 0x10ec, DeviceID: 0x8168, RevisionID: 0x15, ClassCode: 0x020000},
		ConfigSpace:  cs,
		Capabilities: pci.ParseCapabilities(cs),
		BARs:         []pci.BAR{{Index: 0, Type: pci.BARTypeMem64, Size: 4096, Is64Bit: true}},
	}

	report := buildDonorReport(ctx)
	if report.Identity.VendorID != "10ec" || report.Identity.ClassCode != "020000" {
		t.Fatalf("identity = %+v", report.Identity)
	}
	if len(report.Capabilities) != 1 || report.Capabilities[0].Name != "Power Management" {
		t.Fatalf("capabilities = %+v", report.Capabilities)
	}
	if len(report.BARs) != 1 || report.BARs[0].Size != 4096 {
		t.Fatalf("bars = %+v", report.BARs)
	}
	if report.ConfigSpace.Bytes != pci.ConfigSpaceSize || report.ConfigSpace.WritableBits == 0 {
		t.Fatalf("config space = %+v", report.ConfigSpace)
	}
}

func TestPrintDonorReportShowsCapabilityIssues(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.WriteU16(0x06, 0x0010)
	cs.WriteU8(0x34, 0x40)
	cs.WriteU8(0x40, pci.CapIDPowerManagement)
	cs.WriteU8(0x41, 0x40)
	ctx := &donor.DeviceContext{ConfigSpace: cs}

	var out bytes.Buffer
	printDonorReport(&out, buildDonorReport(ctx))
	text := out.String()
	for _, want := range []string{"Donor analysis", "Writable bits:", "standard capability loop at 0x040"} {
		if !strings.Contains(text, want) {
			t.Fatalf("report missing %q:\n%s", want, text)
		}
	}
}

func TestBuildDonorReportLimitsMasksToCapturedConfigSpace(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceLegacySize

	report := buildDonorReport(&donor.DeviceContext{ConfigSpace: cs})

	if report.ConfigSpace.WritableBits+report.ConfigSpace.ReadOnlyBits != pci.ConfigSpaceLegacySize*8 {
		t.Fatalf("mask bits = %+v, want %d total", report.ConfigSpace, pci.ConfigSpaceLegacySize*8)
	}
}
