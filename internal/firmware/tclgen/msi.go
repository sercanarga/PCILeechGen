package tclgen

import (
	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

// extractMSIVectors returns the max MSI vector count from the donor's
// MSI capability (1, 2, 4, 8, 16 or 32). Defaults to 1 if absent.
func extractMSIVectors(ctx *donor.DeviceContext) int {
	for _, cap := range ctx.Capabilities {
		if cap.ID != pci.CapIDMSI {
			continue
		}
		if len(cap.Data) < 4 {
			return 1
		}
		// Message Control register: bits [3:1] = Multiple Message Capable
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

// msiVectorsToTCL formats the vector count for the Xilinx PCIe IP property.
func msiVectorsToTCL(vectors int) string {
	switch {
	case vectors >= 32:
		return "5_vectors"
	case vectors >= 16:
		return "4_vectors"
	case vectors >= 8:
		return "3_vectors"
	case vectors >= 4:
		return "2_vectors"
	case vectors >= 2:
		return "1_vector"
	default:
		return "1_vector"
	}
}
