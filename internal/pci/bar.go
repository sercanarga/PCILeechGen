package pci

import "fmt"

// BAR type constants
const (
	BARTypeIO       = "io"
	BARTypeMem32    = "mem32"
	BARTypeMem64    = "mem64"
	BARTypeDisabled = "disabled"
)

// BAR represents a PCI Base Address Register.
type BAR struct {
	Index        int    `json:"index"`
	RawValue     uint32 `json:"raw_value"`
	Address      uint64 `json:"address"`
	Size         uint64 `json:"size"`
	Type         string `json:"type"` // "io", "mem32", "mem64", "disabled"
	Prefetchable bool   `json:"prefetchable"`
	Is64Bit      bool   `json:"is_64bit"`
}

// IsIO returns true if this is an I/O BAR.
func (b *BAR) IsIO() bool {
	return b.Type == BARTypeIO
}

// IsMemory returns true if this is a memory BAR.
func (b *BAR) IsMemory() bool {
	return b.Type == BARTypeMem32 || b.Type == BARTypeMem64
}

// IsDisabled returns true if this BAR is disabled (zero size or value).
func (b *BAR) IsDisabled() bool {
	return b.Type == BARTypeDisabled || b.Size == 0
}

// SizeHuman returns the BAR size in human-readable format.
func (b *BAR) SizeHuman() string {
	if b.Size == 0 {
		return "0"
	}
	if b.Size >= 1<<30 {
		return fmt.Sprintf("%d GB", b.Size>>30)
	}
	if b.Size >= 1<<20 {
		return fmt.Sprintf("%d MB", b.Size>>20)
	}
	if b.Size >= 1<<10 {
		return fmt.Sprintf("%d KB", b.Size>>10)
	}
	return fmt.Sprintf("%d B", b.Size)
}

// String returns a summary of the BAR for display.
func (b *BAR) String() string {
	if b.IsDisabled() {
		return fmt.Sprintf("BAR%d: [disabled]", b.Index)
	}
	pf := ""
	if b.Prefetchable {
		pf = " [prefetchable]"
	}
	return fmt.Sprintf("BAR%d: %s at 0x%x, size %s%s",
		b.Index, b.Type, b.Address, b.SizeHuman(), pf)
}

// ParseBARsFromConfigSpace extracts BAR information from a config space.
// Note: BAR sizes cannot be determined from config space alone without probing;
// this function only extracts the address and type from raw BAR values.
// For actual sizes, use sysfs resource file or VFIO probing.
func ParseBARsFromConfigSpace(cs *ConfigSpace) []BAR {
	var bars []BAR

	for i := 0; i < 6; i++ {
		rawValue := cs.BAR(i)

		bar := BAR{
			Index:    i,
			RawValue: rawValue,
		}

		if rawValue == 0 {
			bar.Type = BARTypeDisabled
			bars = append(bars, bar)
			continue
		}

		if rawValue&0x01 != 0 {
			// I/O BAR
			bar.Type = BARTypeIO
			bar.Address = uint64(rawValue & 0xFFFFFFFC)
		} else {
			// Memory BAR
			bar.Prefetchable = (rawValue & 0x08) != 0
			memType := (rawValue >> 1) & 0x03

			switch memType {
			case 0x00:
				// 32-bit memory
				bar.Type = BARTypeMem32
				bar.Address = uint64(rawValue & 0xFFFFFFF0)
			case 0x02:
				// 64-bit memory
				bar.Type = BARTypeMem64
				bar.Is64Bit = true
				bar.Address = uint64(rawValue&0xFFFFFFF0) | (uint64(cs.BAR(i+1)) << 32)
			default:
				bar.Type = BARTypeDisabled
			}
		}

		bars = append(bars, bar)

		// Skip upper 32 bits of 64-bit BAR
		if bar.Is64Bit {
			i++
		}
	}

	return bars
}

// ParseBARsFromSysfsResource parses BAR information from sysfs resource lines.
// Each line has format: "start end flags"
func ParseBARsFromSysfsResource(lines []string) []BAR {
	var bars []BAR

	for i := 0; i < 6 && i < len(lines); i++ {
		var start, end, flags uint64
		n, _ := fmt.Sscanf(lines[i], "0x%x 0x%x 0x%x", &start, &end, &flags)
		if n != 3 {
			// Try without 0x prefix
			n, _ = fmt.Sscanf(lines[i], "%x %x %x", &start, &end, &flags)
		}

		bar := BAR{Index: i}

		if start == 0 && end == 0 {
			bar.Type = BARTypeDisabled
		} else {
			bar.Address = start
			bar.Size = end - start + 1

			if flags&0x01 != 0 {
				bar.Type = BARTypeIO
			} else {
				bar.Prefetchable = (flags & 0x08) != 0
				if flags&0x04 != 0 {
					bar.Type = BARTypeMem64
					bar.Is64Bit = true
				} else {
					bar.Type = BARTypeMem32
				}
			}
		}

		bars = append(bars, bar)
	}

	return bars
}
