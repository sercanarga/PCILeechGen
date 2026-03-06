package devclass

import (
	"github.com/sercanarga/pcileechgen/internal/util"
)

const (
	ClassNVMe        = "nvme"
	ClassXHCI        = "xhci"
	ClassEthernet    = "ethernet"
	ClassAudio       = "audio"
	ClassGPU         = "gpu"
	ClassSATA        = "sata"
	ClassWiFi        = "wifi"
	ClassThunderbolt = "thunderbolt"
	ClassGeneric     = "generic"
)

// DeviceStrategy centralizes device-class-specific behavior.
type DeviceStrategy interface {
	ClassName() string
	DeviceClass() string
	Profile() *DeviceProfile
	ScrubBAR(data []byte)
	PostInitRegisters(regs map[uint32]*uint32)
}

type baseStrategy struct {
	className   string
	deviceClass string
	profileFn   func() *DeviceProfile
}

func (s *baseStrategy) ClassName() string       { return s.className }
func (s *baseStrategy) DeviceClass() string     { return s.deviceClass }
func (s *baseStrategy) Profile() *DeviceProfile { return s.profileFn() }

// StrategyForClass returns a strategy for the given PCI class code.
// Returns a generic fallback for unrecognized classes.
func StrategyForClass(classCode uint32) DeviceStrategy {
	baseClass := (classCode >> 16) & 0xFF
	subClass := (classCode >> 8) & 0xFF

	switch {
	case baseClass == 0x01 && subClass == 0x08:
		return &nvmeStrategy{baseStrategy{"NVMe", ClassNVMe, nvmeProfile}}
	case baseClass == 0x0C && subClass == 0x03:
		return &xhciStrategy{baseStrategy{"xHCI", ClassXHCI, xhciProfile}}
	case baseClass == 0x02 && subClass == 0x00:
		return &ethernetStrategy{baseStrategy{"Ethernet", ClassEthernet, ethernetProfile}}
	case baseClass == 0x04 && subClass == 0x03:
		return &audioStrategy{baseStrategy{"HD Audio", ClassAudio, audioProfile}}
	case baseClass == 0x03 && subClass == 0x00:
		return &gpuStrategy{baseStrategy{"GPU", ClassGPU, gpuProfile}}
	case baseClass == 0x01 && subClass == 0x06:
		return &sataStrategy{baseStrategy{"SATA AHCI", ClassSATA, sataProfile}}
	case baseClass == 0x02 && subClass == 0x80:
		return &wifiStrategy{baseStrategy{"Wi-Fi", ClassWiFi, wifiProfile}}
	case baseClass == 0x0C && subClass == 0x80:
		return &thunderboltStrategy{baseStrategy{"Thunderbolt", ClassThunderbolt, thunderboltProfile}}
	default:
		return &genericStrategy{baseStrategy{"Generic", ClassGeneric, genericProfile}}
	}
}

// --- NVMe ---

type nvmeStrategy struct{ baseStrategy }

func (s *nvmeStrategy) ScrubBAR(data []byte) {
	if len(data) < 0x38 {
		return
	}
	// CSTS.RDY=1, clear SHST/CFS/NSSRO
	csts := util.ReadLE32(data, 0x1C)
	csts |= 0x01
	csts &= ^uint32(0x1E)
	util.WriteLE32(data, 0x1C, csts)

	// CC.EN=1
	cc := util.ReadLE32(data, 0x14)
	cc |= 0x01
	util.WriteLE32(data, 0x14, cc)

	// zero interrupt masks & queue base addresses
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

// --- xHCI ---

type xhciStrategy struct{ baseStrategy }

func (s *xhciStrategy) ScrubBAR(data []byte) {
	if len(data) < 0x28 {
		return
	}
	// USBCMD: R/S=1
	usbcmd := util.ReadLE32(data, 0x20)
	usbcmd |= 0x01
	util.WriteLE32(data, 0x20, usbcmd)

	// USBSTS: HCH=0, clear event bits
	util.WriteLE32(data, 0x24, 0x00000000)

	// PAGESIZE: 4KB
	if len(data) > 0x2C {
		util.WriteLE32(data, 0x28, 0x00000001)
	}

	// Clamp DBOFF/RTSOFF within the 4K BRAM window
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

// --- Ethernet ---

type ethernetStrategy struct{ baseStrategy }

func (s *ethernetStrategy) ScrubBAR(data []byte) {
	if len(data) < 0x14 {
		return
	}
	// STATUS: Link Up + 1000Mb/s
	util.WriteLE32(data, 0x08, 0x00000082)

	// EECD: Auto-Read Done + EEPROM Present
	util.WriteLE32(data, 0x10, 0x00000300)

	// MDIC: Ready bit set so PHY reads complete instantly
	if len(data) >= 0x24 {
		util.WriteLE32(data, 0x20, 0x10000000)
	}
}

func (s *ethernetStrategy) PostInitRegisters(regs map[uint32]*uint32) {
	if v, ok := regs[0x08]; ok {
		*v |= 0x00000082
	}
	if v, ok := regs[0x10]; ok {
		*v |= 0x00000300
	}
}

// --- HD Audio ---

type audioStrategy struct{ baseStrategy }

func (s *audioStrategy) ScrubBAR(data []byte) {
	if len(data) < 0x10 {
		return
	}
	// GCTL: CRST=1 (controller not in reset)
	gctl := util.ReadLE32(data, 0x08)
	gctl |= 0x01
	util.WriteLE32(data, 0x08, gctl)

	// STATESTS: codec 0 present (bit 0 of upper 16 = bit 16 of DWORD)
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

// --- GPU ---

type gpuStrategy struct{ baseStrategy }

func (s *gpuStrategy) ScrubBAR(data []byte) {
	if len(data) < 0x204 {
		return
	}
	// PMC_ENABLE: all engines enabled
	util.WriteLE32(data, 0x200, 0xFFFFFFFF)
}

func (s *gpuStrategy) PostInitRegisters(regs map[uint32]*uint32) {
	if v, ok := regs[0x200]; ok {
		*v = 0xFFFFFFFF
	}
}

// --- SATA/AHCI ---

type sataStrategy struct{ baseStrategy }

func (s *sataStrategy) ScrubBAR(data []byte) {
	if len(data) < 0x20 {
		return
	}
	// GHC: AE=1 (AHCI enable), IE=0
	ghc := util.ReadLE32(data, 0x04)
	ghc |= 0x80000000 // AE
	ghc &^= 0x02      // IE off
	util.WriteLE32(data, 0x04, ghc)

	// IS: clear all pending interrupts
	util.WriteLE32(data, 0x08, 0x00000000)

	// Port 0 SSTS: device detected + phy ready
	if len(data) >= 0x12C {
		util.WriteLE32(data, 0x128, 0x00000113)
	}
}

func (s *sataStrategy) PostInitRegisters(regs map[uint32]*uint32) {
	if v, ok := regs[0x04]; ok {
		*v |= 0x80000000
		*v &^= 0x02
	}
}

// --- Wi-Fi ---

type wifiStrategy struct{ baseStrategy }

func (s *wifiStrategy) ScrubBAR(data []byte) {
	if len(data) < 0x28 {
		return
	}
	// GP_CTL: INIT_DONE=1
	util.WriteLE32(data, 0x24, 0x00000080)

	// UCODE_DRV_GP1: firmware loaded
	if len(data) >= 0x58 {
		util.WriteLE32(data, 0x54, 0x00000001)
	}
}

func (s *wifiStrategy) PostInitRegisters(regs map[uint32]*uint32) {
	if v, ok := regs[0x24]; ok {
		*v = 0x00000080
	}
}

// --- Thunderbolt ---

type thunderboltStrategy struct{ baseStrategy }

func (s *thunderboltStrategy) ScrubBAR(data []byte) {
	if len(data) < 0x14 {
		return
	}
	// LC_STS: READY=1
	util.WriteLE32(data, 0x08, 0x00000001)

	// SECURITY_LEVEL: none (no DMA protection)
	util.WriteLE32(data, 0x10, 0x00000000)
}

func (s *thunderboltStrategy) PostInitRegisters(regs map[uint32]*uint32) {
	if v, ok := regs[0x08]; ok {
		*v |= 0x00000001
	}
}

// --- Generic fallback ---

type genericStrategy struct{ baseStrategy }

func (s *genericStrategy) ScrubBAR(data []byte) {}

func (s *genericStrategy) PostInitRegisters(regs map[uint32]*uint32) {}
