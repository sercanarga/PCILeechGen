// Package synthetic builds donor device contexts for CI fixture generation.
package synthetic

import (
	"time"

	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/firmware/devclass"
	"github.com/sercanarga/pcileechgen/internal/pci"
	"github.com/sercanarga/pcileechgen/internal/version"
)

type classProfile struct {
	vendorID   uint16
	deviceID   uint16
	classCode  uint32
	headerType uint8
	bar0Size   uint64
}

var profiles = map[string]classProfile{
	devclass.ClassNVMe:        {0x144d, 0xa809, 0x010802, 0x00, 0x4000},
	devclass.ClassXHCI:        {0x8086, 0x9d2f, 0x0c0330, 0x00, 0x10000},
	devclass.ClassEthernet:    {0x8086, 0x15b7, 0x020000, 0x00, 0x20000},
	devclass.ClassAudio:       {0x8086, 0x9d71, 0x040300, 0x00, 0x4000},
	devclass.ClassGPU:         {0x10de, 0x1b06, 0x030000, 0x00, 0x4000},
	devclass.ClassSATA:        {0x8086, 0x9d03, 0x010601, 0x00, 0x2000},
	devclass.ClassWiFi:        {0x8086, 0x24fd, 0x028000, 0x00, 0x2000},
	devclass.ClassThunderbolt: {0x8086, 0x15d9, 0x080700, 0x00, 0x4000},
	devclass.ClassGeneric:     {0x1234, 0x5678, 0x000000, 0x00, 0x1000},
}

// Build returns a representative DeviceContext for class, or nil if unknown.
func Build(class string) *donor.DeviceContext {
	p, ok := profiles[class]
	if !ok {
		return nil
	}
	cs := buildConfigSpace(p)
	// Capabilities are populated by parsing the config space, mirroring how
	// donor.FromJSON / collector populate real donor contexts (the field is a
	// denormalized cache; the scrub pipeline re-parses the config space).
	return &donor.DeviceContext{
		CollectedAt: time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
		ToolVersion: version.Version,
		Hostname:    "ci-synthetic",
		Device: pci.PCIDevice{
			BDF:        pci.BDF{Bus: 0x03, Device: 0x00, Function: 0x0},
			VendorID:   p.vendorID,
			DeviceID:   p.deviceID,
			ClassCode:  p.classCode,
			HeaderType: p.headerType,
		},
		ConfigSpace:  cs,
		BARs:         buildBARs(p.bar0Size),
		Capabilities: pci.ParseCapabilities(cs),
	}
}

func buildConfigSpace(p classProfile) *pci.ConfigSpace {
	cs := &pci.ConfigSpace{Size: pci.ConfigSpaceLegacySize}
	cs.WriteU16(0x00, p.vendorID)
	cs.WriteU16(0x02, p.deviceID)
	cs.WriteU8(0x08, 0x01)
	cs.WriteU8(0x09, byte(p.classCode))
	cs.WriteU8(0x0A, byte(p.classCode>>8))
	cs.WriteU8(0x0B, byte(p.classCode>>16))
	cs.WriteU8(0x0E, p.headerType)

	cs.WriteU16(0x06, 0x0010)

	cs.WriteU8(0x34, 0x40)

	cs.WriteU8(0x40, pci.CapIDPowerManagement)
	cs.WriteU8(0x41, 0x48)
	cs.WriteU16(0x42, 0x0003)
	cs.WriteU16(0x44, 0x0008)
	cs.WriteU16(0x46, 0x0000)

	cs.WriteU8(0x48, pci.CapIDMSI)
	cs.WriteU8(0x49, 0x54)
	cs.WriteU16(0x4A, 0x0000)
	cs.WriteU32(0x4C, 0x00000000)
	cs.WriteU16(0x50, 0x0000)
	cs.WriteU16(0x52, 0x0000)

	cs.WriteU8(0x54, pci.CapIDPCIExpress)
	cs.WriteU8(0x55, 0x00)
	cs.WriteU16(0x56, 0x0002)

	return cs
}

func buildBARs(bar0Size uint64) []pci.BAR {
	raw := uint32(^(bar0Size - 1)) & 0xFFFFFFF0
	return []pci.BAR{
		{
			Index:    0,
			RawValue: raw,
			Address:  0,
			Size:     bar0Size,
			Type:     "mem32",
		},
	}
}
