package devclass

import (
	"github.com/sercanarga/pcileechgen/internal/util"
)

// DeviceStrategy centralizes device-class-specific behavior.
// Replaces duplicated class code checks in barmodel, scrub, and svgen.
type DeviceStrategy interface {
	ClassName() string
	Profile() *DeviceProfile
	ScrubBAR(data []byte)
	PostInitRegisters(regs map[uint32]*uint32)
	IsNVMe() bool
	IsXHCI() bool
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

// NVMe

type nvmeStrategy struct{}

func (s *nvmeStrategy) ClassName() string       { return "NVMe" }
func (s *nvmeStrategy) Profile() *DeviceProfile { return nvmeProfile() }
func (s *nvmeStrategy) IsNVMe() bool            { return true }
func (s *nvmeStrategy) IsXHCI() bool            { return false }

func (s *nvmeStrategy) ScrubBAR(data []byte) {
	if len(data) < 0x38 {
		return
	}
	// CSTS.RDY=1, clear CFS/SHST/NSSRO
	csts := util.ReadLE32(data, 0x1C)
	csts |= 0x01
	csts &= ^uint32(0x1E)
	util.WriteLE32(data, 0x1C, csts)

	// CC.EN=1
	cc := util.ReadLE32(data, 0x14)
	cc |= 0x01
	util.WriteLE32(data, 0x14, cc)

	// zero interrupt masks, subsystem reset, queue config
	util.WriteLE32(data, 0x0C, 0x00)
	util.WriteLE32(data, 0x10, 0x00)
	util.WriteLE32(data, 0x20, 0x00)
	util.WriteLE32(data, 0x24, 0x00)
	util.WriteLE32(data, 0x28, 0x00)
	util.WriteLE32(data, 0x2C, 0x00)
	util.WriteLE32(data, 0x30, 0x00)
	util.WriteLE32(data, 0x34, 0x00)
}

func (s *nvmeStrategy) PostInitRegisters(regs map[uint32]*uint32) {
	if v, ok := regs[0x1C]; ok {
		*v |= 0x00000001  // CSTS.RDY=1
		*v &^= 0x0000000C // clear SHST
	}
}

// xHCI

type xhciStrategy struct{}

func (s *xhciStrategy) ClassName() string       { return "xHCI" }
func (s *xhciStrategy) Profile() *DeviceProfile { return xhciProfile() }
func (s *xhciStrategy) IsNVMe() bool            { return false }
func (s *xhciStrategy) IsXHCI() bool            { return true }

// ScrubBAR is intentionally empty for xHCI — BRAM-aware scrubbing
// requires BRAMSize which lives in the scrub package. The scrub package
// calls scrubXHCIBar0 directly after running strategy.ScrubBAR.
func (s *xhciStrategy) ScrubBAR(data []byte) {}

func (s *xhciStrategy) PostInitRegisters(regs map[uint32]*uint32) {
	if v, ok := regs[0x20]; ok {
		*v |= 0x00000001 // USBCMD R/S=1
	}
	if v, ok := regs[0x24]; ok {
		*v &^= 0x00000001 // USBSTS HCH=0
	}
}

// Ethernet

type ethernetStrategy struct{}

func (s *ethernetStrategy) ClassName() string       { return "Ethernet" }
func (s *ethernetStrategy) Profile() *DeviceProfile { return ethernetProfile() }
func (s *ethernetStrategy) IsNVMe() bool            { return false }
func (s *ethernetStrategy) IsXHCI() bool            { return false }
func (s *ethernetStrategy) ScrubBAR(data []byte)    {}

func (s *ethernetStrategy) PostInitRegisters(regs map[uint32]*uint32) {
	if v, ok := regs[0x08]; ok {
		*v |= 0x00000002 // STATUS.LU=1 (link up)
	}
}

// HD Audio

type audioStrategy struct{}

func (s *audioStrategy) ClassName() string                      { return "HD Audio" }
func (s *audioStrategy) Profile() *DeviceProfile                { return audioProfile() }
func (s *audioStrategy) IsNVMe() bool                           { return false }
func (s *audioStrategy) IsXHCI() bool                           { return false }
func (s *audioStrategy) ScrubBAR(data []byte)                   {}
func (s *audioStrategy) PostInitRegisters(_ map[uint32]*uint32) {}
