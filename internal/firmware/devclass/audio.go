package devclass

import (
	"github.com/sercanarga/pcileechgen/internal/pci"
	"github.com/sercanarga/pcileechgen/internal/util"
)

type audioStrategy struct{ baseStrategy }

func (s *audioStrategy) ScrubBAR(data []byte) {
	if len(data) < 0x10 {
		return
	}
	gctl := util.ReadLE32(data, 0x08)
	gctl |= 0x01
	util.WriteLE32(data, 0x08, gctl)

	statests := util.ReadLE32(data, 0x0C)
	statests |= 0x00010000
	util.WriteLE32(data, 0x0C, statests)
}

func (s *audioStrategy) PostInitRegisters(regs map[uint32]*uint32) {
	if v, ok := regs[0x08]; ok {
		*v |= 0x00000001
	}
	if v, ok := regs[0x0C]; ok {
		*v |= 0x00010000
	}
}

func audioProfile() *DeviceProfile {
	return &DeviceProfile{
		ClassName:         "HD Audio",
		PreferredBAR:      0,
		MinBARSize:        4096,
		Uses64BitBAR:      true,
		BARIsPrefetchable: false,

		PrefersMSIX:    false,
		MinMSIXVectors: 0,

		ExpectedCaps: []uint8{
			pci.CapIDPowerManagement,
			pci.CapIDMSI,
			pci.CapIDPCIExpress,
		},
		ExpectedExtCaps: []uint16{
			pci.ExtCapIDAER,
		},

		SupportsPME:   true,
		MaxPowerState: 3,

		// DWORD-aligned to match barmodel and bar_impl_device template
		BARDefaults: []BARDefault{
			// GCAP + VMIN + VMAJ packed in one DWORD
			{Offset: 0x00, Width: 4, Name: "GCAP_VMIN_VMAJ", Reset: 0x01004401, RWMask: 0x00000000},
			// GCTL - CRST (bit 0) is the key for reset handshake
			{Offset: 0x08, Width: 4, Name: "GCTL", Reset: 0x00000001, RWMask: 0x00000103},
			// WAKEEN (lower 16) + STATESTS (upper 16) - codec 0 present
			{Offset: 0x0C, Width: 4, Name: "WAKEEN_STATESTS", Reset: 0x00010000, RWMask: 0x7FFFFFFF},
			// INTCTL
			{Offset: 0x20, Width: 4, Name: "INTCTL", Reset: 0x00000000, RWMask: 0xC00000FF},
			// INTSTS
			{Offset: 0x24, Width: 4, Name: "INTSTS", Reset: 0x00000000, RWMask: 0x00000000},
			// CORB base addresses and control
			{Offset: 0x40, Width: 4, Name: "CORBLBASE", Reset: 0x00000000, RWMask: 0xFFFFFF80},
			{Offset: 0x44, Width: 4, Name: "CORBUBASE", Reset: 0x00000000, RWMask: 0xFFFFFFFF},
			{Offset: 0x48, Width: 4, Name: "CORBWP_CORBRP", Reset: 0x00000000, RWMask: 0x80FF00FF},
			{Offset: 0x4C, Width: 4, Name: "CORBCTL_STS_SIZE", Reset: 0x00420000, RWMask: 0x00030300},
			// RIRB base addresses and control
			{Offset: 0x50, Width: 4, Name: "RIRBLBASE", Reset: 0x00000000, RWMask: 0xFFFFFF80},
			{Offset: 0x54, Width: 4, Name: "RIRBUBASE", Reset: 0x00000000, RWMask: 0xFFFFFFFF},
			{Offset: 0x58, Width: 4, Name: "RIRBWP_RINTCNT", Reset: 0x00000000, RWMask: 0x800000FF},
			{Offset: 0x5C, Width: 4, Name: "RIRBCTL_STS_SIZE", Reset: 0x00420000, RWMask: 0x00070700},
			// RIRB response registers — read by hdaudio.sys after RIRBWP advances
			{Offset: 0x70, Width: 4, Name: "RIRBRESP_LO", Reset: 0x00000000, RWMask: 0x00000000},
			{Offset: 0x78, Width: 4, Name: "RIRBRESP_HI", Reset: 0x00000000, RWMask: 0x00000000},
		},

		Notes: "Intel HDA profile. DWORD-packed register layout. GCAP+VMIN+VMAJ at 0x00, " +
			"WAKEEN+STATESTS at 0x0C. hdaudio.sys resets via GCTL.CRST then reads STATESTS.",
	}
}
