package devclass

import (
	"github.com/sercanarga/pcileechgen/internal/util"
)

// device class constants — used by strategy, templates, and barmodel
const (
	ClassNVMe     = "nvme"
	ClassXHCI     = "xhci"
	ClassEthernet = "ethernet"
	ClassAudio    = "audio"
)

// DeviceStrategy centralizes device-class-specific behavior.
type DeviceStrategy interface {
	ClassName() string
	DeviceClass() string
	Profile() *DeviceProfile
	ScrubBAR(data []byte)
	PostInitRegisters(regs map[uint32]*uint32)
}

// StrategyForClass returns a strategy for the given PCI class code, or nil.
func StrategyForClass(classCode uint32) DeviceStrategy {
	baseClass := (classCode >> 16) & 0xFF
	subClass := (classCode >> 8) & 0xFF

	switch {
	case baseClass == 0x01 && subClass == 0x08:
		return &nvmeStrategy{}
	case baseClass == 0x0C && subClass == 0x03:
		return &xhciStrategy{}
	case baseClass == 0x02 && subClass == 0x00:
		return &ethernetStrategy{}
	case baseClass == 0x04 && subClass == 0x03:
		return &audioStrategy{}
	default:
		return nil
	}
}

// --- NVMe ---

type nvmeStrategy struct{}

func (s *nvmeStrategy) ClassName() string       { return "NVMe" }
func (s *nvmeStrategy) DeviceClass() string     { return ClassNVMe }
func (s *nvmeStrategy) Profile() *DeviceProfile { return nvmeProfile() }

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

// --- xHCI ---

type xhciStrategy struct{}

func (s *xhciStrategy) ClassName() string       { return "xHCI" }
func (s *xhciStrategy) DeviceClass() string     { return ClassXHCI }
func (s *xhciStrategy) Profile() *DeviceProfile { return xhciProfile() }
func (s *xhciStrategy) ScrubBAR(data []byte)    {}

func (s *xhciStrategy) PostInitRegisters(regs map[uint32]*uint32) {
	if v, ok := regs[0x20]; ok {
		*v |= 0x00000001
	}
	if v, ok := regs[0x24]; ok {
		*v &^= 0x00000001
	}
}

// --- Ethernet ---

type ethernetStrategy struct{}

func (s *ethernetStrategy) ClassName() string       { return "Ethernet" }
func (s *ethernetStrategy) DeviceClass() string     { return ClassEthernet }
func (s *ethernetStrategy) Profile() *DeviceProfile { return ethernetProfile() }
func (s *ethernetStrategy) ScrubBAR(data []byte)    {}

func (s *ethernetStrategy) PostInitRegisters(regs map[uint32]*uint32) {
	if v, ok := regs[0x08]; ok {
		*v |= 0x00000082
	}
	if v, ok := regs[0x10]; ok {
		*v |= 0x00000300
	}
}

// --- HD Audio ---

type audioStrategy struct{}

func (s *audioStrategy) ClassName() string       { return "HD Audio" }
func (s *audioStrategy) DeviceClass() string     { return ClassAudio }
func (s *audioStrategy) Profile() *DeviceProfile { return audioProfile() }
func (s *audioStrategy) ScrubBAR(data []byte)    {}

func (s *audioStrategy) PostInitRegisters(regs map[uint32]*uint32) {
	if v, ok := regs[0x08]; ok {
		*v |= 0x00000001
	}
	if v, ok := regs[0x0C]; ok {
		*v |= 0x00010000
	}
}
