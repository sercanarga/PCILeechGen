package firmware

import (
	"github.com/sercanarga/pcileechgen/internal/pci"
)

// ScrubConfigSpace cleans potentially dangerous or detection-revealing registers
// from the config space before writing to COE. This prevents leaking donor-specific
// debug/diagnostic data that the DMA card cannot actually implement.
func ScrubConfigSpace(cs *pci.ConfigSpace) *pci.ConfigSpace {
	scrubbed := cs.Clone()

	// Clear BIST register (offset 0x0F) — DMA card cannot run self-test
	scrubbed.WriteU8(0x0F, 0x00)

	// Clear Interrupt Line (offset 0x3C) — will be assigned by host at runtime
	scrubbed.WriteU8(0x3C, 0x00)

	// Clear Latency Timer (offset 0x0D) — not meaningful for PCIe devices
	scrubbed.WriteU8(0x0D, 0x00)

	// Clear Cache Line Size (offset 0x0C) — set by OS
	scrubbed.WriteU8(0x0C, 0x00)

	// Sanitize Command register (offset 0x04) — reset to sane defaults
	// Keep only: IO Space(0), Memory Space(1), Bus Master(2), Parity Error Response(6)
	cmd := scrubbed.Command() & 0x0547
	scrubbed.WriteU16(0x04, cmd)

	// Clear Status register write-1-to-clear bits (offset 0x06)
	// Keep capability list bit (4) and speed bits, clear error bits
	status := scrubbed.Status() & 0x06F0
	scrubbed.WriteU16(0x06, status)

	// Sanitize PCIe Device Status (clear all error/transaction bits)
	caps := pci.ParseCapabilities(scrubbed)
	for _, cap := range caps {
		if cap.ID == pci.CapIDPCIExpress && cap.Offset+10 < pci.ConfigSpaceLegacySize {
			// Device Status at cap+10: clear all RW1C bits (correctable, non-fatal, fatal, unsupported req)
			scrubbed.WriteU16(cap.Offset+10, 0x0000)

			// Link Status at cap+18: read-only from hardware, clear training bits
			if cap.Offset+18 < pci.ConfigSpaceLegacySize {
				lstatus := scrubbed.ReadU16(cap.Offset + 18)
				lstatus &= 0x3FFF // Clear link training and link training error bits
				scrubbed.WriteU16(cap.Offset+18, lstatus)
			}
		}

		if cap.ID == pci.CapIDPowerManagement && cap.Offset+4 < pci.ConfigSpaceLegacySize {
			// PM Control/Status: set to D0 state (powered on), preserve NoSoftReset
			pmcsr := scrubbed.ReadU16(cap.Offset + 4)
			pmcsr &= 0xFFFC // Clear PowerState bits (set to D0)
			pmcsr &= 0x7FFF // Clear PME_Status
			pmcsr |= 0x0008 // Set NoSoftReset bit — FPGA preserves state across D3hot→D0
			scrubbed.WriteU16(cap.Offset+4, pmcsr)
		}
	}

	// Scrub extended config space error registers
	if scrubbed.Size >= pci.ConfigSpaceSize {
		extCaps := pci.ParseExtCapabilities(scrubbed)
		for _, cap := range extCaps {
			if cap.ID == pci.ExtCapIDAER {
				// Clear all AER error status registers (they are RW1C)
				if cap.Offset+4+4 <= pci.ConfigSpaceSize {
					scrubbed.WriteU32(cap.Offset+4, 0x00000000) // Uncorrectable Error Status
				}
				if cap.Offset+16+4 <= pci.ConfigSpaceSize {
					scrubbed.WriteU32(cap.Offset+16, 0x00000000) // Correctable Error Status
				}
				if cap.Offset+28+4 <= pci.ConfigSpaceSize {
					scrubbed.WriteU32(cap.Offset+28, 0x00000000) // Root Error Status
				}
			}
		}
	}

	return scrubbed
}
