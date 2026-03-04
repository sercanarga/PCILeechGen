package devclass

import "github.com/sercanarga/pcileechgen/internal/pci"

func ethernetProfile() *DeviceProfile {
	return &DeviceProfile{
		ClassName:         "Ethernet",
		PreferredBAR:      0,
		MinBARSize:        4096,
		Uses64BitBAR:      false, // many NICs use 32-bit BAR0
		BARIsPrefetchable: false,

		PrefersMSIX:    true,
		MinMSIXVectors: 3, // admin + rx + tx queues

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
			// CTRL — Device Control Register
			{Offset: 0x00, Width: 4, Name: "CTRL", Reset: 0x00000000, RWMask: 0xFFFFFFFF},
			// STATUS — Device Status Register
			{Offset: 0x08, Width: 4, Name: "STATUS", Reset: 0x00000002, RWMask: 0x00000000},
			// CTRL_EXT — Extended Device Control
			{Offset: 0x18, Width: 4, Name: "CTRL_EXT", Reset: 0x00000000, RWMask: 0xFFFFFFFF},
			// EECD — EEPROM/Flash Control & Data
			{Offset: 0x10, Width: 4, Name: "EECD", Reset: 0x00000100, RWMask: 0x00000000},
			// ICR — Interrupt Cause Read
			{Offset: 0xC0, Width: 4, Name: "ICR", Reset: 0x00000000, RWMask: 0x00000000},
			// IMS — Interrupt Mask Set
			{Offset: 0xD0, Width: 4, Name: "IMS", Reset: 0x00000000, RWMask: 0xFFFFFFFF},
		},

		Notes: "Intel GbE (I210/I211/I350) style profile. e1000e/igb driver probes CTRL/STATUS " +
			"first, then checks EECD for NVM presence. MAC address is typically at BAR0+0x5400.",
	}
}
