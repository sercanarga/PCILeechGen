package devclass

import "github.com/sercanarga/pcileechgen/internal/pci"

func audioProfile() *DeviceProfile {
	return &DeviceProfile{
		ClassName:         "HD Audio",
		PreferredBAR:      0,
		MinBARSize:        4096,
		Uses64BitBAR:      true,
		BARIsPrefetchable: false,

		PrefersMSIX:    false, // HDA typically uses MSI, not MSI-X
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
			// GCAP — Global Capabilities
			{Offset: 0x00, Width: 2, Name: "GCAP", Reset: 0x4401, RWMask: 0x0000},
			// VMIN — Minor Version
			{Offset: 0x02, Width: 1, Name: "VMIN", Reset: 0x00, RWMask: 0x00},
			// VMAJ — Major Version
			{Offset: 0x03, Width: 1, Name: "VMAJ", Reset: 0x01, RWMask: 0x00},
			// GCTL — Global Control
			{Offset: 0x08, Width: 4, Name: "GCTL", Reset: 0x00000000, RWMask: 0x00000103},
			// WAKEEN — Wake Enable
			{Offset: 0x0C, Width: 2, Name: "WAKEEN", Reset: 0x0000, RWMask: 0xFFFF},
			// STATESTS — State Change Status
			{Offset: 0x0E, Width: 2, Name: "STATESTS", Reset: 0x0000, RWMask: 0x7FFF},
			// INTCTL — Interrupt Control
			{Offset: 0x20, Width: 4, Name: "INTCTL", Reset: 0x00000000, RWMask: 0xC00000FF},
			// INTSTS — Interrupt Status
			{Offset: 0x24, Width: 4, Name: "INTSTS", Reset: 0x00000000, RWMask: 0x00000000},
			// CORBSIZE — CORB Size
			{Offset: 0x4E, Width: 1, Name: "CORBSIZE", Reset: 0x42, RWMask: 0x03},
		},

		Notes: "Intel HD Audio profile. hdaudio.sys reads GCAP first, then probes GCTL for " +
			"controller reset. CORB/RIRB setup follows. Codec discovery via STATESTS.",
	}
}
