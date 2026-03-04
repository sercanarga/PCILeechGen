package devclass

import "github.com/sercanarga/pcileechgen/internal/pci"

func nvmeProfile() *DeviceProfile {
	return &DeviceProfile{
		ClassName:         "NVMe",
		PreferredBAR:      0,
		MinBARSize:        4096,  // minimum MLBAR size for NVMe
		Uses64BitBAR:      true,  // NVMe spec mandates 64-bit BAR0
		BARIsPrefetchable: false, // MMIO registers are not prefetchable

		PrefersMSIX:    true,
		MinMSIXVectors: 2, // at least admin + 1 I/O queue

		ExpectedCaps: []uint8{
			pci.CapIDPowerManagement,
			pci.CapIDMSIX,
			pci.CapIDPCIExpress,
		},
		ExpectedExtCaps: []uint16{
			pci.ExtCapIDAER,
			pci.ExtCapIDLTR,
			pci.ExtCapIDL1PMSubstates,
		},

		SupportsPME:   true,
		MaxPowerState: 3, // D3hot

		BARDefaults: []BARDefault{
			// NVMe Controller Capabilities (CAP)
			{Offset: 0x00, Width: 4, Name: "CAP_LO", Reset: 0x0040FF17, RWMask: 0x00000000},
			{Offset: 0x04, Width: 4, Name: "CAP_HI", Reset: 0x00000020, RWMask: 0x00000000},
			// Version (VS) — NVMe 1.4
			{Offset: 0x08, Width: 4, Name: "VS", Reset: 0x00010400, RWMask: 0x00000000},
			// Controller Configuration (CC)
			{Offset: 0x14, Width: 4, Name: "CC", Reset: 0x00460001, RWMask: 0x00FFFFF1},
			// Controller Status (CSTS) — RDY=1
			{Offset: 0x1C, Width: 4, Name: "CSTS", Reset: 0x00000001, RWMask: 0x00000000},
			// Admin Queue Attributes (AQA)
			{Offset: 0x24, Width: 4, Name: "AQA", Reset: 0x001F001F, RWMask: 0x0FFF0FFF},
		},

		Notes: "NVMe 1.4 profile. CC.EN→CSTS.RDY handshake is implemented in SV FSM. " +
			"stornvme driver reads CAP first, then writes CC, polls CSTS.RDY.",
	}
}
