package devclass

import (
	"github.com/sercanarga/pcileechgen/internal/pci"
	"github.com/sercanarga/pcileechgen/internal/util"
)

const mediatekVID uint16 = 0x14C3

// mt7922 register values matching real hardware (dmesg: ASIC revision 79220010)
const (
	mt7922ChipID    uint32 = 0x79220010 // TOP_HW_CHIPID - ASIC revision
	mt7922HWSwVer   uint32 = 0x8A108A10 // TOP_HW_SW_VER - hw/sw version
	mt7922TopMisc2  uint32 = 0x00000002 // TOP_MISC2 - strap/misc status
	mt7922FwDlReady uint32 = 0x00000000 // FWDL status - not downloading
)

type mediatekWifiStrategy struct{ baseStrategy }

func (s *mediatekWifiStrategy) ScrubBAR(data []byte) {
	if len(data) < 0x10 {
		return
	}
	// chip identification registers
	util.WriteLE32(data, 0x00, mt7922ChipID)
	util.WriteLE32(data, 0x04, mt7922HWSwVer)
	util.WriteLE32(data, 0x08, mt7922TopMisc2)
	util.WriteLE32(data, 0x0C, mt7922FwDlReady)
}

func (s *mediatekWifiStrategy) PostInitRegisters(regs map[uint32]*uint32) {
	if v, ok := regs[0x00]; ok {
		*v = mt7922ChipID
	}
	if v, ok := regs[0x04]; ok {
		*v = mt7922HWSwVer
	}
	if v, ok := regs[0x08]; ok {
		*v = mt7922TopMisc2
	}
	if v, ok := regs[0x0C]; ok {
		*v = mt7922FwDlReady
	}
}

func mediatekWifiProfile() *DeviceProfile {
	return &DeviceProfile{
		ClassName:         "Wi-Fi (MediaTek)",
		PreferredBAR:      0,
		MinBARSize:        4096,
		Uses64BitBAR:      true,
		BARIsPrefetchable: true,

		PrefersMSIX:    false,
		MinMSIXVectors: 0,

		ExpectedCaps: []uint8{
			pci.CapIDPowerManagement,
			pci.CapIDMSI,
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
			{Offset: 0x00, Width: 4, Name: "TOP_HW_CHIPID", Reset: mt7922ChipID, RWMask: 0x00000000},
			{Offset: 0x04, Width: 4, Name: "TOP_HW_SW_VER", Reset: mt7922HWSwVer, RWMask: 0x00000000},
			{Offset: 0x08, Width: 4, Name: "TOP_MISC2", Reset: mt7922TopMisc2, RWMask: 0x00000000},
			{Offset: 0x0C, Width: 4, Name: "FWDL_STATUS", Reset: mt7922FwDlReady, RWMask: 0x00000000},
		},

		Notes: "MediaTek MT7921/MT7922 Wi-Fi profile. " +
			"ChipID=0x79220010 matches real ASIC revision. " +
			"Driver maps BAR0 only (pcim_iomap_regions BIT(0)). " +
			"WFDMA engine registers (0xd0000+) are beyond 4KB BRAM; " +
			"driver will timeout during firmware download phase.",
	}
}
