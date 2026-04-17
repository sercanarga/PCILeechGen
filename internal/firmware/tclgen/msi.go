package tclgen

import (
	"fmt"

	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

// extractMSIVectors reads the MSI capability from the donor and returns
// how many vectors it supports. Falls back to 1 if no MSI cap is found.
func extractMSIVectors(ctx *donor.DeviceContext) int {
	for _, cap := range ctx.Capabilities {
		if cap.ID != pci.CapIDMSI {
			continue
		}
		if len(cap.Data) < 4 {
			return 1
		}
		// Multiple Message Capable sits at bits [3:1] of Message Control
		msgCtrl := uint16(cap.Data[2]) | uint16(cap.Data[3])<<8
		mmc := (msgCtrl >> 1) & 0x07
		vectors := 1 << mmc
		if vectors > 32 {
			vectors = 32
		}
		return vectors
	}
	return 1
}

// msiVectorsToTCL maps a vector count to the Vivado PCIe IP dropdown value.
func msiVectorsToTCL(vectors int) string {
	switch {
	case vectors >= 32:
		return "32_vectors"
	case vectors >= 16:
		return "16_vectors"
	case vectors >= 8:
		return "8_vectors"
	case vectors >= 4:
		return "4_vectors"
	case vectors >= 2:
		return "2_vectors"
	default:
		return "1_vector"
	}
}

// barBIRToTCL formats a BAR index for Vivado. All BARs are 32-bit after scrub,
// so BIR 0 maps to "BAR_0", not "BAR_1:0" (which is the 64-bit format).
func barBIRToTCL(bir int) string {
	return fmt.Sprintf("BAR_%d", bir)
}
