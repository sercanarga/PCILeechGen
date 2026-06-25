package devclass

import "github.com/sercanarga/pcileechgen/internal/pci"

const (
	intelVID           uint16 = 0x8086
	intelBE200DeviceID uint16 = 0x272B
)

type intelWifiStrategy struct{ baseStrategy }

func (s *intelWifiStrategy) ScrubBAR(data []byte) {
	if len(data) < 0x28 {
		return
	}
	// Keep the same firmware-ready handshake as the generic Intel/Broadcom path.
	data[0x24] = 0x80
	if len(data) >= 0x58 {
		data[0x54] = 0x01
	}
}

func (s *intelWifiStrategy) PostInitRegisters(regs map[uint32]*uint32) {
	if v, ok := regs[0x24]; ok {
		*v = 0x00000080
	}
	if v, ok := regs[0x54]; ok {
		*v = 0x00000001
	}
}

func intelWiFiProfile() *DeviceProfile {
	return &DeviceProfile{
		ClassName:         "Wi-Fi (Intel BE200 family)",
		PreferredBAR:      0,
		MinBARSize:        16384,
		Uses64BitBAR:      true,
		BARIsPrefetchable: false,
		PrefersMSIX:       true,
		MinMSIXVectors:    1,
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
			{Offset: 0x00, Width: 4, Name: "CSR", Reset: 0x00000000, RWMask: 0x00000000},
			{Offset: 0x20, Width: 4, Name: "FH_RSCSR_CHNL0", Reset: 0x00000000, RWMask: 0x00000000},
			{Offset: 0x24, Width: 4, Name: "GP_CTL", Reset: 0x00000080, RWMask: 0xFFFFFFFF},
			{Offset: 0x54, Width: 4, Name: "UCODE_DRV_GP1", Reset: 0x00000001, RWMask: 0xFFFFFFFF},
			{Offset: 0x28, Width: 4, Name: "HW_REV", Reset: 0x00000000, RWMask: 0x00000000},
			{Offset: 0x9C, Width: 4, Name: "RF_ID", Reset: 0x00000000, RWMask: 0x00000000},
		},
		Notes: "Passive Intel Wi-Fi 7 / BE200-family profile derived from the local lspci footprint: 16KB BAR0, MSI-X-capable Wi-Fi endpoint, firmware-ready GP_CTL/UCODE handshake. This avoids live reprofiling of the active network device.",
	}
}
