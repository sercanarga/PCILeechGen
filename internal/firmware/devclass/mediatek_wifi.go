package devclass

import (
	"github.com/sercanarga/pcileechgen/internal/pci"
	"github.com/sercanarga/pcileechgen/internal/util"
)

const mediatekVID uint16 = 0x14C3

// mt7922 chip identification values
const (
	mt7922ChipID uint32 = 0x00007922
	mt7922HWRev  uint32 = 0x00000001
	mt7922HWVer  uint32 = 0x00000001
)

type mediatekWifiStrategy struct{ baseStrategy }

func (s *mediatekWifiStrategy) ScrubBAR(data []byte) {
	if len(data) < 0x0C {
		return
	}
	util.WriteLE32(data, 0x00, mt7922ChipID)
	util.WriteLE32(data, 0x04, mt7922HWRev)
	util.WriteLE32(data, 0x08, mt7922HWVer)
}

func (s *mediatekWifiStrategy) PostInitRegisters(regs map[uint32]*uint32) {
	if v, ok := regs[0x00]; ok {
		*v = mt7922ChipID
	}
	if v, ok := regs[0x04]; ok {
		*v = mt7922HWRev
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
			{Offset: 0x00, Width: 4, Name: "HW_CHIPID", Reset: mt7922ChipID, RWMask: 0x00000000},
			{Offset: 0x04, Width: 4, Name: "HW_REV", Reset: mt7922HWRev, RWMask: 0x00000000},
			{Offset: 0x08, Width: 4, Name: "HW_VER", Reset: mt7922HWVer, RWMask: 0x00000000},
		},

		Notes: "MediaTek MT7921/MT7922 Wi-Fi profile. " +
			"HW_CHIPID set to 0x7922 for driver probe. " +
			"WFDMA registers (0xd4000+) are beyond 4KB BRAM, " +
			"driver may timeout during firmware download.",
	}
}
