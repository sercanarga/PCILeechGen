package tclgen

import (
	"strings"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

func TestExtractMSIVectors(t *testing.T) {
	tests := []struct {
		name    string
		msgCtrl uint16
		want    int
	}{
		{"32 vectors (MMC=5)", 0x008A, 32},
		{"16 vectors (MMC=4)", 0x0088, 16},
		{"8 vectors (MMC=3)", 0x0086, 8},
		{"4 vectors (MMC=2)", 0x0084, 4},
		{"2 vectors (MMC=1)", 0x0082, 2},
		{"1 vector (MMC=0)", 0x0080, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			capData := make([]byte, 24)
			capData[0] = pci.CapIDMSI
			capData[2] = byte(tt.msgCtrl & 0xFF)
			capData[3] = byte(tt.msgCtrl >> 8)

			ctx := &donor.DeviceContext{
				Capabilities: []pci.Capability{
					{ID: pci.CapIDMSI, Offset: 0x50, Data: capData},
				},
			}
			if got := extractMSIVectors(ctx); got != tt.want {
				t.Errorf("extractMSIVectors() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestExtractMSIVectors_NoMSI(t *testing.T) {
	ctx := &donor.DeviceContext{
		Capabilities: []pci.Capability{
			{ID: pci.CapIDPowerManagement, Offset: 0x40, Data: make([]byte, 8)},
		},
	}
	if got := extractMSIVectors(ctx); got != 1 {
		t.Errorf("extractMSIVectors() = %d, want 1 (no MSI cap)", got)
	}
}

func TestExtractMSIVectors_ShortData(t *testing.T) {
	capData := make([]byte, 2)
	capData[0] = pci.CapIDMSI

	ctx := &donor.DeviceContext{
		Capabilities: []pci.Capability{
			{ID: pci.CapIDMSI, Offset: 0x50, Data: capData},
		},
	}
	if got := extractMSIVectors(ctx); got != 1 {
		t.Errorf("extractMSIVectors() = %d, want 1 (short data)", got)
	}
}

func TestMSIVectorsToTCL(t *testing.T) {
	tests := []struct {
		vectors int
		want    string
	}{
		{1, "1_vector"},
		{2, "2_vectors"},
		{4, "4_vectors"},
		{8, "8_vectors"},
		{16, "16_vectors"},
		{32, "32_vectors"},
	}
	for _, tt := range tests {
		if got := msiVectorsToTCL(tt.vectors); got != tt.want {
			t.Errorf("msiVectorsToTCL(%d) = %q, want %q", tt.vectors, got, tt.want)
		}
	}
}

func TestBuildBAR0Config_PreservesOriginalSize(t *testing.T) {
	// 16K BAR0 mem64 with 130-vector MSI-X — BAR must stay >= 16K
	ctx := &donor.DeviceContext{
		BARs: []pci.BAR{
			{Index: 0, Size: 16384, Type: pci.BARTypeMem64},
		},
		MSIXData: &donor.MSIXData{TableSize: 130},
	}

	cfg := buildBAR0Config(ctx, &board.Board{})
	if !cfg.Enabled {
		t.Fatal("BAR0 should be enabled")
	}
	if cfg.Size != "16" || cfg.Scale != "Kilobytes" {
		t.Errorf("BAR0 size = %s %s, want 16 Kilobytes", cfg.Size, cfg.Scale)
	}
	if !cfg.Is64bit {
		t.Error("BAR0 should be 64-bit")
	}
}

func TestBuildBAR0Config_FallbackAlsoProtected(t *testing.T) {
	// IO BAR0 with a 32K size — fallback path should also preserve size
	ctx := &donor.DeviceContext{
		BARs: []pci.BAR{
			{Index: 0, Size: 32768, Type: pci.BARTypeIO},
		},
	}

	cfg := buildBAR0Config(ctx, &board.Board{})
	if !cfg.Enabled {
		t.Fatal("BAR0 should be enabled (fallback)")
	}
	if cfg.Size != "32" || cfg.Scale != "Kilobytes" {
		t.Errorf("BAR0 size = %s %s, want 32 Kilobytes", cfg.Size, cfg.Scale)
	}
}

func TestGenerateProjectTCL_MSIXConfig(t *testing.T) {
	b := &board.Board{
		Name:      "TestBoard",
		FPGAPart:  "xc7a35tfgg484-2",
		PCIeLanes: 1,
		TopModule: "test_top",
	}
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	cs.WriteU16(0x00, 0x144D)
	cs.WriteU16(0x02, 0xA808)

	ctx := &donor.DeviceContext{
		Device:      pci.PCIDevice{VendorID: 0x144D, DeviceID: 0xA808, ClassCode: 0x010802},
		ConfigSpace: cs,
		BARs:        []pci.BAR{{Index: 0, Size: 16384, Type: pci.BARTypeMem64}},
		MSIXData: &donor.MSIXData{
			TableSize: 130, TableBIR: 0, TableOffset: 0x3000,
			PBABIR: 0, PBAOffset: 0x2000,
		},
	}

	tcl := GenerateProjectTCL(ctx, b, "/tmp/lib")

	for _, want := range []string{
		"MSIx_Table_Size",
		"MSIx_Enabled",
		"MSIx_Table_BIR",
		"BAR_1:0",
		"MSIx_PBA_BIR",
		"129vec",
	} {
		if !strings.Contains(tcl, want) {
			t.Errorf("TCL output missing %q", want)
		}
	}
}

func TestGenerateProjectTCL_NoMSIX(t *testing.T) {
	b := &board.Board{
		Name:      "TestBoard",
		FPGAPart:  "xc7a35tfgg484-2",
		PCIeLanes: 1,
		TopModule: "test_top",
	}
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	cs.WriteU16(0x00, 0x8086)
	cs.WriteU16(0x02, 0x1533)

	ctx := &donor.DeviceContext{
		Device:      pci.PCIDevice{VendorID: 0x8086, DeviceID: 0x1533, ClassCode: 0x020000},
		ConfigSpace: cs,
		BARs:        []pci.BAR{},
	}

	tcl := GenerateProjectTCL(ctx, b, "/tmp/lib")

	// MSI-X should NOT be configured when donor has no MSI-X
	if strings.Contains(tcl, "MSIx_Table_Size") {
		t.Error("TCL should NOT contain MSIx_Table_Size when donor lacks MSI-X")
	}
}
