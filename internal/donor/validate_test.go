package donor

import (
	"strings"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/pci"
)

func TestValidateDeviceLayoutReportsBARAndMSIXBounds(t *testing.T) {
	ctx := &DeviceContext{
		BARs: []pci.BAR{
			{Index: 0, Type: pci.BARTypeMem32, Size: 0x1000},
			{Index: 5, Type: pci.BARTypeMem64, Size: 0x1000, Is64Bit: true},
		},
		MSIXData: &MSIXData{TableSize: 4, TableBIR: 0, TableOffset: 0xff0, PBABIR: 3, PBAOffset: 0},
	}

	issues := ValidateDeviceLayout(ctx)
	joined := strings.Join(issues, "\n")
	for _, want := range []string{"64-bit BAR5 has no upper BAR", "MSI-X table exceeds BAR0", "MSI-X PBA references unavailable BAR3"} {
		if !strings.Contains(joined, want) {
			t.Fatalf("issues = %v, want %q", issues, want)
		}
	}
}

func TestValidateDeviceLayoutAcceptsValidMSIXPlacement(t *testing.T) {
	ctx := &DeviceContext{
		BARs:     []pci.BAR{{Index: 0, Type: pci.BARTypeMem64, Size: 0x2000, Is64Bit: true}},
		MSIXData: &MSIXData{TableSize: 4, TableBIR: 0, TableOffset: 0x1000, PBABIR: 0, PBAOffset: 0x1100},
	}
	if issues := ValidateDeviceLayout(ctx); len(issues) != 0 {
		t.Fatalf("ValidateDeviceLayout() = %v, want no issues", issues)
	}
}
