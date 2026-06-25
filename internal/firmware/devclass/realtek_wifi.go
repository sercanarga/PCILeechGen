package devclass

import "github.com/sercanarga/pcileechgen/internal/pci"

const realtekRTL8188EEDeviceID uint16 = 0x8179

type realtekWifiStrategy struct{ baseStrategy }

func (s *realtekWifiStrategy) ScrubBAR(data []byte) {}

func (s *realtekWifiStrategy) PostInitRegisters(regs map[uint32]*uint32) {}

func realtekRTL8188EEProfile() *DeviceProfile {
	return &DeviceProfile{
		ClassName:         "Wi-Fi (Realtek RTL8188EE)",
		PreferredBAR:      0,
		MinBARSize:        4096,
		Uses64BitBAR:      false,
		BARIsPrefetchable: false,
		PrefersMSIX:       false,
		MinMSIXVectors:    0,
		ExpectedCaps: []uint8{
			pci.CapIDPowerManagement,
			pci.CapIDMSI,
			pci.CapIDPCIExpress,
		},
		ExpectedExtCaps: []uint16{
			pci.ExtCapIDAER,
			pci.ExtCapIDDeviceSerialNumber,
		},
		SupportsPME:   true,
		MaxPowerState: 3,
		BARDefaults: []BARDefault{
			{Offset: 0x0000, Width: 4, Name: "SYS_ISO_CTRL_FUNC_EN", Reset: 0x00000000, RWMask: 0xFFFFFFFF},
			{Offset: 0x0004, Width: 4, Name: "APS_FSMCO", Reset: 0x00000000, RWMask: 0xFFFFFFFF},
			{Offset: 0x0080, Width: 4, Name: "MCUFWDL_WOL_EVENT", Reset: 0x00000000, RWMask: 0xFFFFFFFF},
			{Offset: 0x00B0, Width: 4, Name: "HIMR", Reset: 0x00000000, RWMask: 0xFFFFFFFF},
			{Offset: 0x00B4, Width: 4, Name: "HISR", Reset: 0x00000000, RWMask: 0x00000000, IsRW1C: true},
			{Offset: 0x0100, Width: 4, Name: "CR_PBP", Reset: 0x00000000, RWMask: 0xFFFFFFFF},
			{Offset: 0x0300, Width: 4, Name: "PCIE_CTRL_INT_MIG", Reset: 0x00000000, RWMask: 0xFFFFFFFF},
		},
		Notes: "Realtek RTL8188EE PCIe Wi-Fi profile. Based on the public RTL8188EE donor config from kilmu1337/pcileech-rtl8188ee-wifi-emul and Linux rtl8188ee register names: 32-bit BAR0, MSI capability, no MSI-X table.",
	}
}
