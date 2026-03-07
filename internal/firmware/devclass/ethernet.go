package devclass

import (
	"github.com/sercanarga/pcileechgen/internal/pci"
	"github.com/sercanarga/pcileechgen/internal/util"
)

type ethernetStrategy struct{ baseStrategy }

func (s *ethernetStrategy) ScrubBAR(data []byte) {
	if len(data) < 0x40 {
		return
	}
	util.WriteLE32(data, 0x00, 0xBEADDE02) // MAC0-3
	if len(data) >= 0x08 {
		util.WriteLE32(data, 0x04, 0x000000EF) // MAC4-5
	}
	data[0x37] = 0x0C // RxEn | TxEn
	if len(data) >= 0x40 {
		util.WriteLE32(data, 0x3C, 0x00000000) // IntrMask off
	}
	if len(data) >= 0x44 {
		util.WriteLE32(data, 0x40, 0x2F000000) // TxConfig
	}
	if len(data) >= 0x48 {
		util.WriteLE32(data, 0x44, 0x00000E00) // RxConfig
	}
	if len(data) >= 0x54 {
		util.WriteLE32(data, 0x50, 0x00003FFF) // RxMaxSize
		util.WriteLE32(data, 0x48, 0x00000000) // Timer
	}
	if len(data) >= 0x5C {
		util.WriteLE32(data, 0x58, 0x00002060) // CPlusCmd
	}
	if len(data) >= 0x70 {
		util.WriteLE32(data, 0x6C, 0x00003010) // PHYStatus
	}
	if len(data) >= 0xE0 {
		util.WriteLE32(data, 0xDC, 0x80000000) // PHYAR
	}
	if len(data) >= 0xE4 {
		util.WriteLE32(data, 0xE0, 0x80000000) // ERIAR
	}
	if len(data) >= 0x100 {
		util.WriteLE32(data, 0xFC, 0x00000000) // RxMissed
	}
}

func (s *ethernetStrategy) PostInitRegisters(regs map[uint32]*uint32) {
	if v, ok := regs[0x34]; ok {
		*v |= 0x0C000000 // ChipCmd @ byte 0x37
	}
	if v, ok := regs[0x6C]; ok {
		*v |= 0x00003010
	}
}

func ethernetProfile() *DeviceProfile {
	return &DeviceProfile{
		ClassName:         "Ethernet",
		PreferredBAR:      2,
		MinBARSize:        65536,
		Uses64BitBAR:      true,
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
			{Offset: 0x00, Width: 4, Name: "MAC0_3", Reset: 0xBEADDE02, RWMask: 0xFFFFFFFF},
			{Offset: 0x04, Width: 4, Name: "MAC4_5", Reset: 0x000000EF, RWMask: 0xFFFFFFFF},
			{Offset: 0x34, Width: 4, Name: "CHIPCMD_DW", Reset: 0x0C000000, RWMask: 0x00000000},
			{Offset: 0x3C, Width: 4, Name: "INTRMASK", Reset: 0x00000000, RWMask: 0xFFFFFFFF},
			{Offset: 0x40, Width: 4, Name: "TXCONFIG", Reset: 0x2F000000, RWMask: 0x00FF0000},
			{Offset: 0x44, Width: 4, Name: "RXCONFIG", Reset: 0x00000E00, RWMask: 0xFFFF7FFF},
			{Offset: 0x48, Width: 4, Name: "TIMER", Reset: 0x00000000, RWMask: 0xFFFFFFFF},
			{Offset: 0x50, Width: 4, Name: "RXMAXSIZE", Reset: 0x00003FFF, RWMask: 0x00003FFF},
			{Offset: 0x58, Width: 4, Name: "CPLUSCMD", Reset: 0x00002060, RWMask: 0x0000FFFF},
			{Offset: 0x6C, Width: 4, Name: "PHYSTATUS", Reset: 0x00003010, RWMask: 0x00000000},
			{Offset: 0xDC, Width: 4, Name: "PHYAR", Reset: 0x80000000, RWMask: 0xFFFFFFFF},
			{Offset: 0xE0, Width: 4, Name: "ERIAR", Reset: 0x80000000, RWMask: 0xFFFFFFFF},
			{Offset: 0xFC, Width: 4, Name: "RXMISSED", Reset: 0x00000000, RWMask: 0x00000000},
		},

		Notes: "Realtek RTL8125 2.5GbE profile. BAR2 is the primary 64KB MMIO region. " +
			"MAC address at offset 0x00. ChipCmd at 0x37. PHYStatus returns link-up " +
			"at 2500Mbps full-duplex. TxConfig carries chip version ID.",
	}
}
