package devclass

import "github.com/sercanarga/pcileechgen/internal/pci"

func xhciProfile() *DeviceProfile {
	return &DeviceProfile{
		ClassName:         "xHCI USB",
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
		},

		SupportsPME:   true,
		MaxPowerState: 3,

		BARDefaults: []BARDefault{
			// CAPLENGTH + HCIVERSION
			{Offset: 0x00, Width: 4, Name: "CAPLENGTH_HCIVER", Reset: 0x01000020, RWMask: 0x00000000},
			// HCSPARAMS1: MaxSlots=32, MaxIntrs=1, MaxPorts=2
			{Offset: 0x04, Width: 4, Name: "HCSPARAMS1", Reset: 0x02000120, RWMask: 0x00000000},
			// HCSPARAMS2: scratchpad=0
			{Offset: 0x08, Width: 4, Name: "HCSPARAMS2", Reset: 0x00000000, RWMask: 0x00000000},
			// HCSPARAMS3
			{Offset: 0x0C, Width: 4, Name: "HCSPARAMS3", Reset: 0x00000000, RWMask: 0x00000000},
			// HCCPARAMS1: xECP=0 (cleared for FPGA)
			{Offset: 0x10, Width: 4, Name: "HCCPARAMS1", Reset: 0x00000000, RWMask: 0x00000000},
			// DBOFF — doorbell offset, clamped for BRAM
			{Offset: 0x14, Width: 4, Name: "DBOFF", Reset: 0x00000100, RWMask: 0x00000000},
			// RTSOFF — runtime register offset
			{Offset: 0x18, Width: 4, Name: "RTSOFF", Reset: 0x00000200, RWMask: 0x00000000},
			// USBCMD — run/stop=1
			{Offset: 0x20, Width: 4, Name: "USBCMD", Reset: 0x00000001, RWMask: 0x00001F0F},
			// USBSTS — HCH=0 (running)
			{Offset: 0x24, Width: 4, Name: "USBSTS", Reset: 0x00000000, RWMask: 0x0000041C},
		},

		Notes: "xHCI 1.1 profile. Registers clamped to 4KB BRAM. DBOFF/RTSOFF must " +
			"be within BRAM range. usbxhci driver reads CAPLENGTH then accesses operational regs.",
	}
}
