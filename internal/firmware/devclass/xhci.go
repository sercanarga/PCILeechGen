package devclass

import (
	"github.com/sercanarga/pcileechgen/internal/pci"
	"github.com/sercanarga/pcileechgen/internal/util"
)

type xhciStrategy struct{ baseStrategy }

func (s *xhciStrategy) ScrubBAR(data []byte) {
	if len(data) < 0x28 {
		return
	}
	usbcmd := util.ReadLE32(data, 0x20)
	usbcmd |= 0x01
	util.WriteLE32(data, 0x20, usbcmd)

	util.WriteLE32(data, 0x24, 0x00000000)

	if len(data) > 0x2C {
		util.WriteLE32(data, 0x28, 0x00000001)
	}

	if len(data) > 0x1C {
		dboff := util.ReadLE32(data, 0x14)
		if dboff > 0x800 {
			util.WriteLE32(data, 0x14, 0x00000100)
		}
		rtsoff := util.ReadLE32(data, 0x18)
		if rtsoff > 0xC00 {
			util.WriteLE32(data, 0x18, 0x00000200)
		}
	}
}

func (s *xhciStrategy) PostInitRegisters(regs map[uint32]*uint32) {
	if v, ok := regs[0x20]; ok {
		*v |= 0x00000001
	}
	if v, ok := regs[0x24]; ok {
		*v &^= 0x00000001
	}
}

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
			// CAPLENGTH=0x20, HCIVERSION=1.10
			{Offset: 0x00, Width: 4, Name: "CAPLENGTH_HCIVER", Reset: 0x01100020, RWMask: 0x00000000},
			// HCSPARAMS1: MaxSlots=32, MaxIntrs=1, MaxPorts=2
			{Offset: 0x04, Width: 4, Name: "HCSPARAMS1", Reset: 0x02000120, RWMask: 0x00000000},
			// HCSPARAMS2: no scratchpad
			{Offset: 0x08, Width: 4, Name: "HCSPARAMS2", Reset: 0x00000000, RWMask: 0x00000000},
			// HCSPARAMS3: exit latencies
			{Offset: 0x0C, Width: 4, Name: "HCSPARAMS3", Reset: 0x00000000, RWMask: 0x00000000},
			// HCCPARAMS1: 64-bit capable, no xECP
			{Offset: 0x10, Width: 4, Name: "HCCPARAMS1", Reset: 0x00000001, RWMask: 0x00000000},
			// DBOFF - doorbell array offset
			{Offset: 0x14, Width: 4, Name: "DBOFF", Reset: 0x00000100, RWMask: 0x00000000},
			// RTSOFF - runtime register space offset
			{Offset: 0x18, Width: 4, Name: "RTSOFF", Reset: 0x00000200, RWMask: 0x00000000},
			// Operational registers (at CAPLENGTH offset 0x20)
			// USBCMD - R/S=1 (running)
			{Offset: 0x20, Width: 4, Name: "USBCMD", Reset: 0x00000001, RWMask: 0x00001F0F},
			// USBSTS - HCH=0 (not halted)
			{Offset: 0x24, Width: 4, Name: "USBSTS", Reset: 0x00000000, RWMask: 0x0000041C},
			// PAGESIZE - 4KB pages
			{Offset: 0x28, Width: 4, Name: "PAGESIZE", Reset: 0x00000001, RWMask: 0x00000000},
			// DNCTRL - device notification control
			{Offset: 0x34, Width: 4, Name: "DNCTRL", Reset: 0x00000000, RWMask: 0x0000FFFF},
			// CRCR - command ring control
			{Offset: 0x38, Width: 4, Name: "CRCR_LO", Reset: 0x00000000, RWMask: 0xFFFFFFF0},
			{Offset: 0x3C, Width: 4, Name: "CRCR_HI", Reset: 0x00000000, RWMask: 0xFFFFFFFF},
			// DCBAAP - device context base address
			{Offset: 0x50, Width: 4, Name: "DCBAAP_LO", Reset: 0x00000000, RWMask: 0xFFFFFFC0},
			{Offset: 0x54, Width: 4, Name: "DCBAAP_HI", Reset: 0x00000000, RWMask: 0xFFFFFFFF},
			// CONFIG - max device slots enabled
			{Offset: 0x58, Width: 4, Name: "CONFIG", Reset: 0x00000000, RWMask: 0x000000FF},
			// PORTSC1 - port 1 status/control (powered, no device)
			{Offset: 0x420, Width: 4, Name: "PORTSC1", Reset: 0x000002A0, RWMask: 0x8EFFC3F2},
			// PORTSC2 - port 2 status/control
			{Offset: 0x430, Width: 4, Name: "PORTSC2", Reset: 0x000002A0, RWMask: 0x8EFFC3F2},
		},

		Notes: "xHCI 1.1 profile. HCCPARAMS1 bit 0 = AC64 (64-bit capable). " +
			"PORTSC powered+PP with no device attached. USBCMD R/S=1, USBSTS HCH=0.",
	}
}
