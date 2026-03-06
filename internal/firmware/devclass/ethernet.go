package devclass

import "github.com/sercanarga/pcileechgen/internal/pci"

func ethernetProfile() *DeviceProfile {
	return &DeviceProfile{
		ClassName:         "Ethernet",
		PreferredBAR:      0,
		MinBARSize:        4096,
		Uses64BitBAR:      false,
		BARIsPrefetchable: false,

		PrefersMSIX:    true,
		MinMSIXVectors: 3,

		ExpectedCaps: []uint8{
			pci.CapIDPowerManagement,
			pci.CapIDMSIX,
			pci.CapIDPCIExpress,
		},
		ExpectedExtCaps: []uint16{
			pci.ExtCapIDAER,
			pci.ExtCapIDLTR,
			pci.ExtCapIDDeviceSerialNumber,
		},

		SupportsPME:   true,
		MaxPowerState: 3,

		BARDefaults: []BARDefault{
			// CTRL — device control
			{Offset: 0x00, Width: 4, Name: "CTRL", Reset: 0x00000000, RWMask: 0xFFFFFFFF},
			// STATUS — link up, speed 1000Mb
			{Offset: 0x08, Width: 4, Name: "STATUS", Reset: 0x00000082, RWMask: 0x00000000},
			// EECD — EEPROM control (Auto-Read Done + EEPROM Present)
			{Offset: 0x10, Width: 4, Name: "EECD", Reset: 0x00000300, RWMask: 0x00000000},
			// CTRL_EXT — extended control
			{Offset: 0x18, Width: 4, Name: "CTRL_EXT", Reset: 0x00000000, RWMask: 0xFFFFFFFF},
			// MDIC — PHY access, auto-complete reads
			{Offset: 0x20, Width: 4, Name: "MDIC", Reset: 0x10000000, RWMask: 0x0FFFFFFF},
			// ICR — interrupt cause (read-clear)
			{Offset: 0xC0, Width: 4, Name: "ICR", Reset: 0x00000000, RWMask: 0x00000000},
			// IMS — interrupt mask set
			{Offset: 0xD0, Width: 4, Name: "IMS", Reset: 0x00000000, RWMask: 0xFFFFFFFF},
			// IMC — interrupt mask clear
			{Offset: 0xD8, Width: 4, Name: "IMC", Reset: 0x00000000, RWMask: 0xFFFFFFFF},
			// RAL0 — receive address low (fake MAC: 02:DE:AD:BE:EF:00)
			{Offset: 0x5400, Width: 4, Name: "RAL0", Reset: 0xADDE0200, RWMask: 0xFFFFFFFF},
			// RAH0 — receive address high + AV (address valid)
			{Offset: 0x5404, Width: 4, Name: "RAH0", Reset: 0x8000EFBE, RWMask: 0xFFFFFFFF},
		},

		Notes: "Intel GbE profile. STATUS.LU=1 + speed bits. EECD has Auto-Read Done " +
			"and EEPROM Present. RAL0/RAH0 provide a locally-administered MAC address. " +
			"MDIC.Ready=1 so PHY reads complete instantly.",
	}
}
