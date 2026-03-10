package devclass

import (
	"github.com/sercanarga/pcileechgen/internal/pci"
	"github.com/sercanarga/pcileechgen/internal/util"
)

type nvmeStrategy struct{ baseStrategy }

func (s *nvmeStrategy) ScrubBAR(data []byte) {
	if len(data) < 0x38 {
		return
	}
	csts := util.ReadLE32(data, 0x1C)
	csts |= 0x01
	csts &= ^uint32(0x1E)
	util.WriteLE32(data, 0x1C, csts)

	cc := util.ReadLE32(data, 0x14)
	cc |= 0x01
	util.WriteLE32(data, 0x14, cc)

	for _, off := range []int{0x0C, 0x10, 0x20, 0x24, 0x28, 0x2C, 0x30, 0x34} {
		util.WriteLE32(data, off, 0x00)
	}
}

func (s *nvmeStrategy) PostInitRegisters(regs map[uint32]*uint32) {
	if v, ok := regs[0x1C]; ok {
		*v |= 0x00000001
		*v &^= 0x0000000C
	}
}

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
			// Version (VS) - NVMe 1.4
			{Offset: 0x08, Width: 4, Name: "VS", Reset: 0x00010400, RWMask: 0x00000000},
			// Interrupt Mask Set/Clear (RO in emulation - MSI-X used)
			{Offset: 0x0C, Width: 4, Name: "INTMS", Reset: 0x00000000, RWMask: 0x00000000},
			{Offset: 0x10, Width: 4, Name: "INTMC", Reset: 0x00000000, RWMask: 0x00000000},
			// Controller Configuration (CC)
			{Offset: 0x14, Width: 4, Name: "CC", Reset: 0x00460001, RWMask: 0x00FFFFF1},
			// Controller Status (CSTS) - RDY=1
			{Offset: 0x1C, Width: 4, Name: "CSTS", Reset: 0x00000001, RWMask: 0x00000000},
			// NVM Subsystem Reset (NSSR)
			{Offset: 0x20, Width: 4, Name: "NSSR", Reset: 0x00000000, RWMask: 0xFFFFFFFF},
			// Admin Queue Attributes (AQA)
			{Offset: 0x24, Width: 4, Name: "AQA", Reset: 0x001F001F, RWMask: 0x0FFF0FFF},
			// Admin Submission Queue Base Address (ASQ)
			{Offset: 0x28, Width: 4, Name: "ASQ_LO", Reset: 0x00000000, RWMask: 0xFFFFF000},
			{Offset: 0x2C, Width: 4, Name: "ASQ_HI", Reset: 0x00000000, RWMask: 0xFFFFFFFF},
			// Admin Completion Queue Base Address (ACQ)
			{Offset: 0x30, Width: 4, Name: "ACQ_LO", Reset: 0x00000000, RWMask: 0xFFFFF000},
			{Offset: 0x34, Width: 4, Name: "ACQ_HI", Reset: 0x00000000, RWMask: 0xFFFFFFFF},
		},

		Notes: "NVMe 1.4 profile. CC.EN->CSTS.RDY handshake is implemented in SV FSM. " +
			"stornvme driver reads CAP first, then writes CC, polls CSTS.RDY.",
	}
}
