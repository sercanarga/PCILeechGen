package devclass

import "github.com/sercanarga/pcileechgen/internal/pci"

const (
	realtekVID      uint16 = 0x10EC
	rts522aDeviceID uint16 = 0x522A
)

type rts522aCardReaderStrategy struct{ baseStrategy }

func (s *rts522aCardReaderStrategy) ScrubBAR(data []byte) {}

func (s *rts522aCardReaderStrategy) PostInitRegisters(regs map[uint32]*uint32) {}

func rts522aProfile() *DeviceProfile {
	return &DeviceProfile{
		ClassName:         "Card Reader (RTS522A)",
		PreferredBAR:      0,
		MinBARSize:        4096,
		Uses64BitBAR:      false,
		BARIsPrefetchable: false,
		PrefersMSIX:       false,
		MinMSIXVectors:    0,
		PrefersINTx:       false,
		ExpectedCaps: []uint8{
			pci.CapIDPowerManagement,
			pci.CapIDMSI,
			pci.CapIDPCIExpress,
		},
		SupportsPME:   true,
		MaxPowerState: 3,
		BARDefaults: []BARDefault{
			{Offset: 0x00, Width: 4, Name: "RTSX_HCBAR", Reset: 0x00000000, RWMask: 0xFFFFFFFF},
			{Offset: 0x04, Width: 4, Name: "RTSX_HCBCTLR", Reset: 0x00000000, RWMask: 0xF0000000},
			{Offset: 0x08, Width: 4, Name: "RTSX_HDBAR", Reset: 0x00000000, RWMask: 0xFFFFFFFF},
			{Offset: 0x0C, Width: 4, Name: "RTSX_HDBCTLR", Reset: 0x00000000, RWMask: 0x8C000000},
			{Offset: 0x10, Width: 4, Name: "RTSX_HAIMR", Reset: 0x00000000, RWMask: 0xFFFFFFFF, IsFSMDriven: true},
			{Offset: 0x14, Width: 4, Name: "RTSX_BIPR", Reset: 0x00010000, RWMask: 0x00000000, IsFSMDriven: true},
			{Offset: 0x18, Width: 4, Name: "RTSX_BIER", Reset: 0x00000000, RWMask: 0xFFC00000},
			{Offset: 0x1C, Width: 4, Name: "RTSX_DUM_REG", Reset: 0x00000000, RWMask: 0x00000000},
		},
		Notes: "Realtek RTS522A PCIe card reader. Derived from Linux rtsx_pci register definitions and donor drvscan config-space. Uses BAR0 host-command/data window registers and simple interrupt status/enable registers.",
	}
}
