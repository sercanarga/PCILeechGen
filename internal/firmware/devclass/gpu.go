package devclass

import (
	"github.com/sercanarga/pcileechgen/internal/pci"
	"github.com/sercanarga/pcileechgen/internal/util"
)

type gpuStrategy struct{ baseStrategy }

func (s *gpuStrategy) ScrubBAR(data []byte) {
	if len(data) < 0x204 {
		return
	}
	util.WriteLE32(data, 0x200, 0xFFFFFFFF)
}

func (s *gpuStrategy) PostInitRegisters(regs map[uint32]*uint32) {
	if v, ok := regs[0x200]; ok {
		*v = 0xFFFFFFFF
	}
}

func gpuProfile() *DeviceProfile {
	return &DeviceProfile{
		ClassName:         "GPU",
		PreferredBAR:      0,
		MinBARSize:        4096, // real VRAM BAR is huge, FPGA maps only 4K window
		Uses64BitBAR:      true,
		BARIsPrefetchable: true,

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
			pci.ExtCapIDDeviceSerialNumber,
		},

		SupportsPME:   true,
		MaxPowerState: 3,

		BARDefaults: []BARDefault{
			// NV_PMC — boot status, master control
			{Offset: 0x00, Width: 4, Name: "PMC_BOOT", Reset: 0x00000000, RWMask: 0x00000000},
			// NV_PMC_ENABLE — engine enable bitmask
			{Offset: 0x200, Width: 4, Name: "PMC_ENABLE", Reset: 0xFFFFFFFF, RWMask: 0xFFFFFFFF},
			// NV_PBUS_PCI_NV_0 — mirrors VID/DID inside BAR
			{Offset: 0x1800, Width: 4, Name: "PBUS_PCI_NV_0", Reset: 0x00000000, RWMask: 0x00000000},
			// NV_PBUS_PCI_NV_1 — mirrors config command/status
			{Offset: 0x1804, Width: 4, Name: "PBUS_PCI_NV_1", Reset: 0x00100006, RWMask: 0x00000000},
			// NV_PTIMER_TIME_0 — low 32 bits of GPU timer
			{Offset: 0x9400, Width: 4, Name: "PTIMER_TIME_0", Reset: 0x00000000, RWMask: 0x00000000},
			// NV_PTIMER_TIME_1 — high 32 bits
			{Offset: 0x9410, Width: 4, Name: "PTIMER_TIME_1", Reset: 0x00000000, RWMask: 0x00000000},
		},

		Notes: "NVIDIA/AMD GPU profile. Only 4K BRAM window visible — " +
			"driver sees PMC_ENABLE and PTIMER but can't touch VRAM.",
	}
}
