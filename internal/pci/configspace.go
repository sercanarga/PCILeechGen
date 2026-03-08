package pci

import (
	"encoding/binary"
	"fmt"
	"strings"
)

// ConfigSpaceSize is the full PCIe extended config space size (4KB).
const ConfigSpaceSize = 4096

// ConfigSpaceLegacySize is the legacy PCI config space size (256 bytes).
const ConfigSpaceLegacySize = 256

type ConfigSpace struct {
	Data [ConfigSpaceSize]byte
	Size int // actual bytes read (256 or 4096)
}

func NewConfigSpace() *ConfigSpace {
	return &ConfigSpace{Size: ConfigSpaceSize}
}

func NewConfigSpaceFromBytes(data []byte) *ConfigSpace {
	cs := &ConfigSpace{Size: len(data)}
	copy(cs.Data[:], data)
	return cs
}

func (cs *ConfigSpace) VendorID() uint16     { return binary.LittleEndian.Uint16(cs.Data[0x00:0x02]) } // 0x00
func (cs *ConfigSpace) DeviceID() uint16     { return binary.LittleEndian.Uint16(cs.Data[0x02:0x04]) } // 0x02
func (cs *ConfigSpace) Command() uint16      { return binary.LittleEndian.Uint16(cs.Data[0x04:0x06]) } // 0x04
func (cs *ConfigSpace) Status() uint16       { return binary.LittleEndian.Uint16(cs.Data[0x06:0x08]) } // 0x06
func (cs *ConfigSpace) RevisionID() uint8    { return cs.Data[0x08] }
func (cs *ConfigSpace) ProgIF() uint8        { return cs.Data[0x09] }
func (cs *ConfigSpace) SubClass() uint8      { return cs.Data[0x0A] }
func (cs *ConfigSpace) BaseClass() uint8     { return cs.Data[0x0B] }
func (cs *ConfigSpace) CacheLineSize() uint8 { return cs.Data[0x0C] }
func (cs *ConfigSpace) LatencyTimer() uint8  { return cs.Data[0x0D] }
func (cs *ConfigSpace) HeaderType() uint8    { return cs.Data[0x0E] }
func (cs *ConfigSpace) BIST() uint8          { return cs.Data[0x0F] }

func (cs *ConfigSpace) ClassCode() uint32 {
	return uint32(cs.BaseClass())<<16 | uint32(cs.SubClass())<<8 | uint32(cs.ProgIF())
}

func (cs *ConfigSpace) IsMultiFunction() bool { return (cs.HeaderType() & 0x80) != 0 }
func (cs *ConfigSpace) HeaderLayout() uint8   { return cs.HeaderType() & 0x7F }

func (cs *ConfigSpace) BAR(index int) uint32 {
	if index < 0 || index > 5 {
		return 0
	}
	offset := 0x10 + (index * 4)
	return binary.LittleEndian.Uint32(cs.Data[offset : offset+4])
}

func (cs *ConfigSpace) SubsysVendorID() uint16 { return binary.LittleEndian.Uint16(cs.Data[0x2C:0x2E]) } // 0x2C
func (cs *ConfigSpace) SubsysDeviceID() uint16 { return binary.LittleEndian.Uint16(cs.Data[0x2E:0x30]) } // 0x2E
func (cs *ConfigSpace) ExpansionROMBase() uint32 {
	return binary.LittleEndian.Uint32(cs.Data[0x30:0x34])
}                                                // 0x30
func (cs *ConfigSpace) CapabilityPointer() uint8 { return cs.Data[0x34] }
func (cs *ConfigSpace) InterruptLine() uint8     { return cs.Data[0x3C] }
func (cs *ConfigSpace) InterruptPin() uint8      { return cs.Data[0x3D] }
func (cs *ConfigSpace) MinGrant() uint8          { return cs.Data[0x3E] }
func (cs *ConfigSpace) MaxLatency() uint8        { return cs.Data[0x3F] }

func (cs *ConfigSpace) HasCapabilities() bool { return (cs.Status() & 0x0010) != 0 }

func (cs *ConfigSpace) ReadU8(offset int) uint8 {
	if offset < 0 || offset >= ConfigSpaceSize {
		return 0
	}
	return cs.Data[offset]
}

func (cs *ConfigSpace) ReadU16(offset int) uint16 {
	if offset < 0 || offset+2 > ConfigSpaceSize {
		return 0
	}
	return binary.LittleEndian.Uint16(cs.Data[offset : offset+2])
}

func (cs *ConfigSpace) ReadU32(offset int) uint32 {
	if offset < 0 || offset+4 > ConfigSpaceSize {
		return 0
	}
	return binary.LittleEndian.Uint32(cs.Data[offset : offset+4])
}

func (cs *ConfigSpace) WriteU8(offset int, val uint8) {
	if offset >= 0 && offset < ConfigSpaceSize {
		cs.Data[offset] = val
	}
}

func (cs *ConfigSpace) WriteU16(offset int, val uint16) {
	if offset >= 0 && offset+1 < ConfigSpaceSize {
		binary.LittleEndian.PutUint16(cs.Data[offset:offset+2], val)
	}
}

func (cs *ConfigSpace) WriteU32(offset int, val uint32) {
	if offset >= 0 && offset+3 < ConfigSpaceSize {
		binary.LittleEndian.PutUint32(cs.Data[offset:offset+4], val)
	}
}

func (cs *ConfigSpace) Clone() *ConfigSpace {
	clone := &ConfigSpace{Size: cs.Size}
	copy(clone.Data[:], cs.Data[:])
	return clone
}

func (cs *ConfigSpace) Bytes() []byte { return cs.Data[:cs.Size] }

func (cs *ConfigSpace) HexDump(maxBytes int) string {
	if maxBytes <= 0 || maxBytes > cs.Size {
		maxBytes = cs.Size
	}

	var sb strings.Builder
	for i := 0; i < maxBytes; i += 16 {
		sb.WriteString(fmt.Sprintf("%03x: ", i))
		for j := 0; j < 16 && i+j < maxBytes; j++ {
			sb.WriteString(fmt.Sprintf("%02x ", cs.Data[i+j]))
			if j == 7 {
				sb.WriteString(" ")
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}
