package devclass

import (
	"github.com/sercanarga/pcileechgen/internal/pci"
	"github.com/sercanarga/pcileechgen/internal/util"
)

type thunderboltStrategy struct{ baseStrategy }

func (s *thunderboltStrategy) ScrubBAR(data []byte) {
	if len(data) < 0x14 {
		return
	}
	util.WriteLE32(data, 0x08, 0x00000001)
	util.WriteLE32(data, 0x10, 0x00000000)
}

func (s *thunderboltStrategy) PostInitRegisters(regs map[uint32]*uint32) {
	if v, ok := regs[0x08]; ok {
		*v |= 0x00000001
	}
}

func thunderboltProfile() *DeviceProfile {
	return &DeviceProfile{
		ClassName:         "Thunderbolt",
		PreferredBAR:      0,
		MinBARSize:        262144,
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
			{Offset: 0x37800, Width: 4, Name: "RING_NOTIFY", Reset: 0x00000000, RWMask: 0xFFFFFFFF},
			{Offset: 0x38200, Width: 4, Name: "RING_INTERRUPT", Reset: 0x00000000, RWMask: 0xFFFFFFFF},
			{Offset: 0x39640, Width: 4, Name: "HOP_COUNT", Reset: 0x00000001, RWMask: 0x000007FF},
			{Offset: 0x39898, Width: 4, Name: "NHI_RESET", Reset: 0x00000000, RWMask: 0x00000001},
			{Offset: 0x39900, Width: 4, Name: "INMAIL_DATA", Reset: 0x00000000, RWMask: 0xFFFFFFFF},
			{Offset: 0x39904, Width: 4, Name: "INMAIL_CMD", Reset: 0x00000000, RWMask: 0xFFFFFFFF},
			{Offset: 0x3990C, Width: 4, Name: "OUTMAIL_CMD", Reset: 0x00000100, RWMask: 0x00000000},
			{Offset: 0x39944, Width: 4, Name: "FW_STATUS", Reset: 0x00000007, RWMask: 0x00000000},
			// LC_MAILBOX_IN - host-to-controller mailbox
			{Offset: 0x00, Width: 4, Name: "LC_MAILBOX_IN", Reset: 0x00000000, RWMask: 0xFFFFFFFF},
			// LC_MAILBOX_OUT - controller reply
			{Offset: 0x04, Width: 4, Name: "LC_MAILBOX_OUT", Reset: 0x00000000, RWMask: 0x00000000},
			// LC_STS - link controller status (ready + no error)
			{Offset: 0x08, Width: 4, Name: "LC_STS", Reset: 0x00000001, RWMask: 0x00000000},
			// SECURITY_LEVEL - none (no DMA protection active)
			{Offset: 0x10, Width: 4, Name: "SECURITY_LEVEL", Reset: 0x00000000, RWMask: 0x00000000},
			// THUNDERBOLT_CAP - capability flags
			{Offset: 0x14, Width: 4, Name: "THUNDERBOLT_CAP", Reset: 0x00000001, RWMask: 0x00000000},
		},

		Notes: "Intel Thunderbolt controller profile. " +
			"LC_STS.READY=1. SECURITY_LEVEL=0 (no security) for full DMA access.",
	}
}
