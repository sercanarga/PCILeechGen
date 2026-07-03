package devclass

import (
	"github.com/sercanarga/pcileechgen/internal/pci"
	"github.com/sercanarga/pcileechgen/internal/util"
)

type sdhciStrategy struct{ baseStrategy }

// presentStateNoCard is Present State (0x24) with no card in the slot: CINS/CSS/CDPL
// clear (bits 16-18), command/data lines idle (bits 0-2), DAT[3:0]/CMD signal levels
// pulled high (bus idle-high, bits 20-24), WP reads "not protected" (bit 19) since
// there's no card to protect. A host driver enumerates fine with no card present -
// card presence is a separate polled/interrupt event, not an enumeration blocker.
const presentStateNoCard = 0x01F80000

func (s *sdhciStrategy) ScrubBAR(data []byte) {
	if len(data) < 0x30 {
		return
	}
	util.WriteLE32(data, 0x24, presentStateNoCard)

	// Clock Control: force Internal Clock Stable (bit1) so a driver polling for clock
	// lock never spins, and clear the Software Reset bits (24-26) so a donor dump
	// captured mid-reset doesn't replay as "reset still in progress".
	cc := util.ReadLE32(data, 0x2C)
	cc |= 0x00000002
	cc &^= 0x07000000
	util.WriteLE32(data, 0x2C, cc)

	if len(data) >= 0x34 {
		// Normal/Error Interrupt Status - no pending interrupts at boot.
		util.WriteLE32(data, 0x30, 0x00000000)
	}
}

func (s *sdhciStrategy) PostInitRegisters(regs map[uint32]*uint32) {
	if v, ok := regs[0x24]; ok {
		*v = presentStateNoCard
	}
	if v, ok := regs[0x2C]; ok {
		*v |= 0x00000002
		*v &^= 0x07000000
	}
}

func sdhciProfile() *DeviceProfile {
	return &DeviceProfile{
		ClassName:         "SD Host Controller",
		PreferredBAR:      0, // sdhci-pci (Linux) ioremaps BAR0 for the standard register set
		MinBARSize:        4096,
		Uses64BitBAR:      false,
		BARIsPrefetchable: false,

		PrefersMSIX:    false,
		MinMSIXVectors: 0,

		ExpectedCaps: []uint8{
			pci.CapIDPowerManagement,
			pci.CapIDMSI,
			pci.CapIDPCIExpress,
		},
		ExpectedExtCaps: []uint16{
			pci.ExtCapIDAER,
		},

		SupportsPME:   true,
		MaxPowerState: 3,

		BARDefaults: []BARDefault{
			// SDMA System Address / Argument2 (for Auto CMD23)
			{Offset: 0x00, Width: 4, Name: "SDMASYSADDR_ARG2", Reset: 0x00000000, RWMask: 0xFFFFFFFF},
			// Block Size [15:0] (xfer size[11:0] + SDMA boundary[14:12]) + Block Count [31:16]
			{Offset: 0x04, Width: 4, Name: "BLOCKSIZE_BLOCKCOUNT", Reset: 0x00000000, RWMask: 0xFFFF7FFF},
			// Argument1 - command argument
			{Offset: 0x08, Width: 4, Name: "ARGUMENT1", Reset: 0x00000000, RWMask: 0xFFFFFFFF},
			// Transfer Mode [15:0] (bits 0-5 used) + Command [31:16] (bits 0-13 used)
			{Offset: 0x0C, Width: 4, Name: "XFERMODE_COMMAND", Reset: 0x00000000, RWMask: 0x3FFF003F},
			// Response 0-3 - card response, read-only (no card, no command ever completes)
			{Offset: 0x10, Width: 4, Name: "RESPONSE0", Reset: 0x00000000, RWMask: 0x00000000},
			{Offset: 0x14, Width: 4, Name: "RESPONSE1", Reset: 0x00000000, RWMask: 0x00000000},
			{Offset: 0x18, Width: 4, Name: "RESPONSE2", Reset: 0x00000000, RWMask: 0x00000000},
			{Offset: 0x1C, Width: 4, Name: "RESPONSE3", Reset: 0x00000000, RWMask: 0x00000000},
			// Buffer Data Port - PIO data window
			{Offset: 0x20, Width: 4, Name: "BUFFERDATAPORT", Reset: 0x00000000, RWMask: 0xFFFFFFFF},
			// Present State - no card inserted, bus idle, lines pulled high
			{Offset: 0x24, Width: 4, Name: "PRESENTSTATE", Reset: presentStateNoCard, RWMask: 0x00000000},
			// Host Control 1 [7:0] + Power Control [15:8] + Block Gap Control [23:16] + Wakeup Control [31:24]
			{Offset: 0x28, Width: 4, Name: "HOSTCTRL1_PWRCTRL_BGAP_WAKECTRL", Reset: 0x00000000, RWMask: 0x070F0FFF},
			// Clock Control [15:0] (ICS bit1 is RO, forced via Scrub/PostInit) + Timeout Control [23:16] + Software Reset [31:24]
			{Offset: 0x2C, Width: 4, Name: "CLOCKCTRL_TIMEOUTCTRL_SWRESET", Reset: 0x000E0000, RWMask: 0x070FFFCD},
			// Normal Interrupt Status [15:0] (bits 0-7 W1C, bit8 Card Interrupt is RO-live) + Error Interrupt Status [31:16] (bits 0-9 W1C)
			{Offset: 0x30, Width: 4, Name: "NORMALINTSTS_ERRORINTSTS", Reset: 0x00000000, RWMask: 0x00000000, W1CMask: 0x03FF00FF},
			// Normal Interrupt Status Enable [15:0] (bits 0-8) + Error Interrupt Status Enable [31:16] (bits 0-9)
			{Offset: 0x34, Width: 4, Name: "NORMALINTEN_ERRORINTEN", Reset: 0x00000000, RWMask: 0x03FF01FF},
			// Normal Interrupt Signal Enable [15:0] + Error Interrupt Signal Enable [31:16] - masked off,
			// matches this codebase's "interrupts off for FPGA" convention (see sata.go GHC.IE)
			{Offset: 0x38, Width: 4, Name: "NORMALSIGEN_ERRORSIGEN", Reset: 0x00000000, RWMask: 0x03FF01FF},
			// Capabilities - SDHCI 2.00 feature set: SDMA yes, no ADMA/high-speed/UHS, 3.3V only
			// ponytail: base/timeout clock frequencies (50MHz/50KHz) are plausible round numbers picked
			// for internal consistency, not lifted from a real silicon datasheet
			{Offset: 0x40, Width: 4, Name: "CAPABILITIES", Reset: 0x01403232, RWMask: 0x00000000},
			// Slot Interrupt Status [15:0] (no pending) + Host Controller Version [31:16].
			// Version field = 0x01 -> SD Host Spec Version 2.00 (register encodes 0x00=1.00, 0x01=2.00,
			// 0x02=3.00 per spec Table 2-16), consistent with the no-UHS/no-ADMA Capabilities above.
			{Offset: 0xFC, Width: 4, Name: "SLOTINTSTS_HCVERSION", Reset: 0x00010000, RWMask: 0x00000000},
		},

		Notes: "SDHCI (SD Host Controller Simplified Spec) 2.00 profile - vendor-agnostic per-spec register " +
			"set, the same layout Linux's sdhci.c targets regardless of silicon vendor. No card inserted: " +
			"Present State CINS/CSS/CDPL clear, host still enumerates normally (card detect is a separate " +
			"polled/interrupt path). SDMA supported, no ADMA2/high-speed/UHS, 3.3V only. Interrupt Signal " +
			"Enable masked off (interrupts off for FPGA).",
	}
}
