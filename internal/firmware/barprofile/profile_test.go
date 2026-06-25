package barprofile

import (
	"testing"

	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

func TestBuildProfile_classifiesBarsAndRegisters(t *testing.T) {
	ctx := &donor.DeviceContext{
		Device: pci.PCIDevice{ClassCode: 0x010802},
		BARs: []pci.BAR{
			{Index: 0, Size: 16, Type: pci.BARTypeMem64, Is64Bit: true},
			{Index: 2, Size: 8, Type: pci.BARTypeIO},
		},
		BARContents: map[int][]byte{
			0: {
				0x00, 0x00, 0x00, 0x00,
				0x34, 0x12, 0x00, 0x00,
				0xFF, 0xFF, 0xFF, 0xFF,
				0x00, 0x00, 0x00, 0x00,
			},
			2: {0xAA, 0x00, 0x00, 0x00},
		},
		BARProfiles: map[int]*donor.BARProfile{
			0: {
				BarIndex: 0,
				Size:     16,
				Probes: []donor.BARProbeResult{
					{Offset: 0x00, Original: 0x00000000, RWMask: 0x00000000},
					{Offset: 0x04, Original: 0x00001234, RWMask: 0x00000000},
					{Offset: 0x08, Original: 0xFFFFFFFF, RWMask: 0x00000000},
					{Offset: 0x0C, Original: 0x00000080, RWMask: 0x00000080, MaybeRW1C: true},
				},
			},
		},
	}

	profile := Build(ctx)

	if profile.ClassCode != 0x010802 {
		t.Fatalf("ClassCode = 0x%06X, want 0x010802", profile.ClassCode)
	}
	if profile.DeviceClass != "nvme" {
		t.Fatalf("DeviceClass = %q, want nvme", profile.DeviceClass)
	}
	if len(profile.BARs) != 2 {
		t.Fatalf("BAR count = %d, want 2", len(profile.BARs))
	}
	bar0 := profile.BARs[0]
	if bar0.EmulationHint != HintProbeModel {
		t.Fatalf("BAR0 hint = %q, want %q", bar0.EmulationHint, HintProbeModel)
	}
	if bar0.StaticRegisters != 1 || bar0.RW1CRegisters != 1 ||
		bar0.AllOnesRegisters != 1 || bar0.DeadRegisters != 1 {
		t.Fatalf("BAR0 summary mismatch: %+v", bar0)
	}
	if len(bar0.Registers) != 4 {
		t.Fatalf("BAR0 registers = %d, want 4", len(bar0.Registers))
	}
	if bar0.Registers[3].Kind != RegisterRW1C {
		t.Fatalf("BAR0 register 0x0C kind = %q, want %q", bar0.Registers[3].Kind, RegisterRW1C)
	}
	if profile.BARs[1].EmulationHint != HintIOSpaceUnsupported {
		t.Fatalf("BAR2 hint = %q, want %q", profile.BARs[1].EmulationHint, HintIOSpaceUnsupported)
	}
}
