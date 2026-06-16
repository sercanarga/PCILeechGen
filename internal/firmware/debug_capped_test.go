package firmware

import (
	"fmt"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

func TestDebugCappedForMSIXTestCase(t *testing.T) {
	// mimic TestWriteConditionalArtifacts_MSIXDonor
	cs := pci.NewConfigSpace()
	ctx := &donor.DeviceContext{
		Device: pci.PCIDevice{
			VendorID:  0x8086,
			DeviceID:  0x15B7,
			ClassCode: 0x020000,
		},
		ConfigSpace: cs,
		BARs:        []pci.BAR{{Index: 0, Size: 4096}},
		BARContents: map[int][]byte{0: make([]byte, 4096)},
	}
	ctx.MSIXData = &donor.MSIXData{TableSize: 4}
	b := &board.Board{BRAMSize: 8192}
	msix := ctx.MSIXData.TableSize
	got := CappedBAR0Size(ctx, b, msix)
	fmt.Printf("DEBUG: Capped for msix4 donor4k board8k = %d\n", got)
	demand := DonorBAR0Demand(ctx, b, msix)
	fmt.Printf("DEBUG: Demand= %d  (bram in b=%d)\n", demand, b.BRAMSizeOrDefault())
	if got > 8192 {
		t.Errorf("capped gave %d >8192", got)
	}
}
