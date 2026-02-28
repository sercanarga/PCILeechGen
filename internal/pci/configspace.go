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

// ConfigSpace represents a full PCI/PCIe configuration space (4096 bytes).
type ConfigSpace struct {
	Data [ConfigSpaceSize]byte
	Size int // actual bytes read (256 or 4096)
}

// NewConfigSpace creates an empty ConfigSpace.
func NewConfigSpace() *ConfigSpace {
	return &ConfigSpace{Size: ConfigSpaceSize}
}

// NewConfigSpaceFromBytes creates a ConfigSpace from a byte slice.
func NewConfigSpaceFromBytes(data []byte) *ConfigSpace {
	cs := &ConfigSpace{Size: len(data)}
	copy(cs.Data[:], data)
	return cs
}

// --- Standard PCI Header (Type 0) accessor methods ---

// VendorID returns the Vendor ID (offset 0x00).
func (cs *ConfigSpace) VendorID() uint16 {
	return binary.LittleEndian.Uint16(cs.Data[0x00:0x02])
}

// DeviceID returns the Device ID (offset 0x02).
func (cs *ConfigSpace) DeviceID() uint16 {
	return binary.LittleEndian.Uint16(cs.Data[0x02:0x04])
}

// Command returns the Command register (offset 0x04).
func (cs *ConfigSpace) Command() uint16 {
	return binary.LittleEndian.Uint16(cs.Data[0x04:0x06])
}

// Status returns the Status register (offset 0x06).
func (cs *ConfigSpace) Status() uint16 {
	return binary.LittleEndian.Uint16(cs.Data[0x06:0x08])
}

// RevisionID returns the Revision ID (offset 0x08).
func (cs *ConfigSpace) RevisionID() uint8 {
	return cs.Data[0x08]
}

// ProgIF returns the Programming Interface (offset 0x09).
func (cs *ConfigSpace) ProgIF() uint8 {
	return cs.Data[0x09]
}

// SubClass returns the Sub-Class code (offset 0x0A).
func (cs *ConfigSpace) SubClass() uint8 {
	return cs.Data[0x0A]
}

// BaseClass returns the Base Class code (offset 0x0B).
func (cs *ConfigSpace) BaseClass() uint8 {
	return cs.Data[0x0B]
}

// ClassCode returns the full 24-bit class code.
func (cs *ConfigSpace) ClassCode() uint32 {
	return uint32(cs.BaseClass())<<16 | uint32(cs.SubClass())<<8 | uint32(cs.ProgIF())
}

// CacheLineSize returns the Cache Line Size (offset 0x0C).
func (cs *ConfigSpace) CacheLineSize() uint8 {
	return cs.Data[0x0C]
}

// LatencyTimer returns the Latency Timer (offset 0x0D).
func (cs *ConfigSpace) LatencyTimer() uint8 {
	return cs.Data[0x0D]
}

// HeaderType returns the Header Type (offset 0x0E).
func (cs *ConfigSpace) HeaderType() uint8 {
	return cs.Data[0x0E]
}

// IsMultiFunction returns true if the device is multi-function.
func (cs *ConfigSpace) IsMultiFunction() bool {
	return (cs.HeaderType() & 0x80) != 0
}

// HeaderLayout returns the header layout type (0, 1, or 2).
func (cs *ConfigSpace) HeaderLayout() uint8 {
	return cs.HeaderType() & 0x7F
}

// BIST returns the Built-In Self Test register (offset 0x0F).
func (cs *ConfigSpace) BIST() uint8 {
	return cs.Data[0x0F]
}

// BAR returns the Base Address Register value at the given index (0-5).
func (cs *ConfigSpace) BAR(index int) uint32 {
	if index < 0 || index > 5 {
		return 0
	}
	offset := 0x10 + (index * 4)
	return binary.LittleEndian.Uint32(cs.Data[offset : offset+4])
}

// SubsysVendorID returns the Subsystem Vendor ID (offset 0x2C).
func (cs *ConfigSpace) SubsysVendorID() uint16 {
	return binary.LittleEndian.Uint16(cs.Data[0x2C:0x2E])
}

// SubsysDeviceID returns the Subsystem Device ID (offset 0x2E).
func (cs *ConfigSpace) SubsysDeviceID() uint16 {
	return binary.LittleEndian.Uint16(cs.Data[0x2E:0x30])
}

// ExpansionROMBase returns the Expansion ROM Base Address (offset 0x30).
func (cs *ConfigSpace) ExpansionROMBase() uint32 {
	return binary.LittleEndian.Uint32(cs.Data[0x30:0x34])
}

// CapabilityPointer returns the Capabilities Pointer (offset 0x34).
func (cs *ConfigSpace) CapabilityPointer() uint8 {
	return cs.Data[0x34]
}

// InterruptLine returns the Interrupt Line (offset 0x3C).
func (cs *ConfigSpace) InterruptLine() uint8 {
	return cs.Data[0x3C]
}

// InterruptPin returns the Interrupt Pin (offset 0x3D).
func (cs *ConfigSpace) InterruptPin() uint8 {
	return cs.Data[0x3D]
}

// MinGrant returns the Min Grant (offset 0x3E).
func (cs *ConfigSpace) MinGrant() uint8 {
	return cs.Data[0x3E]
}

// MaxLatency returns the Max Latency (offset 0x3F).
func (cs *ConfigSpace) MaxLatency() uint8 {
	return cs.Data[0x3F]
}

// HasCapabilities returns true if the device has capabilities (status bit 4).
func (cs *ConfigSpace) HasCapabilities() bool {
	return (cs.Status() & 0x0010) != 0
}

// ReadU8 reads a uint8 from the given offset.
func (cs *ConfigSpace) ReadU8(offset int) uint8 {
	if offset < 0 || offset >= ConfigSpaceSize {
		return 0
	}
	return cs.Data[offset]
}

// ReadU16 reads a little-endian uint16 from the given offset.
func (cs *ConfigSpace) ReadU16(offset int) uint16 {
	if offset < 0 || offset+1 >= ConfigSpaceSize {
		return 0
	}
	return binary.LittleEndian.Uint16(cs.Data[offset : offset+2])
}

// ReadU32 reads a little-endian uint32 from the given offset.
func (cs *ConfigSpace) ReadU32(offset int) uint32 {
	if offset < 0 || offset+3 >= ConfigSpaceSize {
		return 0
	}
	return binary.LittleEndian.Uint32(cs.Data[offset : offset+4])
}

// WriteU8 writes a uint8 at the given offset.
func (cs *ConfigSpace) WriteU8(offset int, val uint8) {
	if offset >= 0 && offset < ConfigSpaceSize {
		cs.Data[offset] = val
	}
}

// WriteU16 writes a little-endian uint16 at the given offset.
func (cs *ConfigSpace) WriteU16(offset int, val uint16) {
	if offset >= 0 && offset+1 < ConfigSpaceSize {
		binary.LittleEndian.PutUint16(cs.Data[offset:offset+2], val)
	}
}

// WriteU32 writes a little-endian uint32 at the given offset.
func (cs *ConfigSpace) WriteU32(offset int, val uint32) {
	if offset >= 0 && offset+3 < ConfigSpaceSize {
		binary.LittleEndian.PutUint32(cs.Data[offset:offset+4], val)
	}
}

// Clone creates a deep copy of the ConfigSpace.
func (cs *ConfigSpace) Clone() *ConfigSpace {
	clone := &ConfigSpace{Size: cs.Size}
	copy(clone.Data[:], cs.Data[:])
	return clone
}

// Bytes returns the actual config space data as a byte slice.
func (cs *ConfigSpace) Bytes() []byte {
	return cs.Data[:cs.Size]
}

// HexDump returns a hex dump of the config space for debugging.
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
