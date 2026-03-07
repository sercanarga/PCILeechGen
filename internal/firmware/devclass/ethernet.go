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
	// MAC0-3 (locally-administered: 02:DE:AD:BE)
	util.WriteLE32(data, 0x00, 0xBEADDE02)
	if len(data) >= 0x08 {
		// MAC4-5 + padding (EF:00)
		util.WriteLE32(data, 0x04, 0x000000EF)
	}
	// ChipCmd: RxEnable + TxEnable
	data[0x37] = 0x0C
	// IntrMask: disabled
	if len(data) >= 0x40 {
		util.WriteLE32(data, 0x3C, 0x00000000)
	}
	// TxConfig: RTL8125B chip version + IFG
	if len(data) >= 0x44 {
		util.WriteLE32(data, 0x40, 0x2F000000)
	}
	// RxConfig: FIFO threshold
	if len(data) >= 0x48 {
		util.WriteLE32(data, 0x44, 0x00000E00)
	}
	// PHYStatus: Link up + 2500Mbps + Full-duplex
	if len(data) >= 0x70 {
		util.WriteLE32(data, 0x6C, 0x00003010)
	}
	// PHYar: PHY access ready
	if len(data) >= 0xDE {
		util.WriteLE32(data, 0xDA, 0x80000000)
	}
}

func (s *ethernetStrategy) PostInitRegisters(regs map[uint32]*uint32) {
	// ChipCmd is a single byte at 0x37, but register map uses DWORD-aligned access.
	// Use DWORD at 0x34 — ChipCmd is byte [3] of that DWORD.
	if v, ok := regs[0x34]; ok {
		*v |= 0x0C000000 // ChipCmd.RxEn + TxEn in MSB
	}
	// PHYStatus: link up + 2500 + full-duplex
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
			// MAC0-3: locally-administered MAC (02:DE:AD:BE)
			{Offset: 0x00, Width: 4, Name: "MAC0_3", Reset: 0xBEADDE02, RWMask: 0xFFFFFFFF},
			// MAC4-5 + pad (EF:00:00:00)
			{Offset: 0x04, Width: 4, Name: "MAC4_5", Reset: 0x000000EF, RWMask: 0xFFFFFFFF},
			// ChipCmd DWORD (ChipCmd at byte 0x37 = MSB of DWORD 0x34)
			{Offset: 0x34, Width: 4, Name: "CHIPCMD_DW", Reset: 0x0C000000, RWMask: 0x00000000},
			// IntrMask
			{Offset: 0x3C, Width: 4, Name: "INTRMASK", Reset: 0x00000000, RWMask: 0xFFFFFFFF},
			// TxConfig: chip version + IFG bits
			{Offset: 0x40, Width: 4, Name: "TXCONFIG", Reset: 0x2F000000, RWMask: 0x00FF0000},
			// RxConfig: FIFO threshold
			{Offset: 0x44, Width: 4, Name: "RXCONFIG", Reset: 0x00000E00, RWMask: 0xFFFF7FFF},
			// PHYStatus: link=1, speed=2500, duplex=full
			{Offset: 0x6C, Width: 4, Name: "PHYSTATUS", Reset: 0x00003010, RWMask: 0x00000000},
		},

		Notes: "Realtek RTL8125 2.5GbE profile. BAR2 is the primary 64KB MMIO region. " +
			"MAC address at offset 0x00. ChipCmd at 0x37. PHYStatus returns link-up " +
			"at 2500Mbps full-duplex. TxConfig carries chip version ID.",
	}
}
