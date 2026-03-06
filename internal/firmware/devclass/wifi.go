package devclass

import "github.com/sercanarga/pcileechgen/internal/pci"

func wifiProfile() *DeviceProfile {
	return &DeviceProfile{
		ClassName:         "Wi-Fi",
		PreferredBAR:      0,
		MinBARSize:        4096,
		Uses64BitBAR:      true,
		BARIsPrefetchable: false,

		PrefersMSIX:    true,
		MinMSIXVectors: 1,

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
		MaxPowerState: 3,

		BARDefaults: []BARDefault{
			// CSR — control/status (HW revision, device ready etc.)
			{Offset: 0x00, Width: 4, Name: "CSR", Reset: 0x00000000, RWMask: 0x00000000},
			// FH_RSCSR_CHNL0 — RX status
			{Offset: 0x20, Width: 4, Name: "FH_RSCSR_CHNL0", Reset: 0x00000000, RWMask: 0x00000000},
			// GP_CTL — firmware ready handshake
			{Offset: 0x24, Width: 4, Name: "GP_CTL", Reset: 0x00000080, RWMask: 0xFFFFFFFF},
			// UCODE_DRV_GP1 — uCode ready flag (set = f/w loaded)
			{Offset: 0x54, Width: 4, Name: "UCODE_DRV_GP1", Reset: 0x00000001, RWMask: 0xFFFFFFFF},
			// HW_REV — hardware revision
			{Offset: 0x28, Width: 4, Name: "HW_REV", Reset: 0x00000000, RWMask: 0x00000000},
			// RF_ID — radio identification
			{Offset: 0x9C, Width: 4, Name: "RF_ID", Reset: 0x00000000, RWMask: 0x00000000},
		},

		Notes: "Intel Wi-Fi / Broadcom profile. GP_CTL.INIT_DONE=1 and " +
			"UCODE_DRV_GP1 set to indicate firmware is loaded. " +
			"L1PM Substates expected for aggressive power management.",
	}
}
