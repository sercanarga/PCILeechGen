package pci

// MSIXInfo describes MSI-X capability layout from config space.
type MSIXInfo struct {
	CapOffset   int    // offset of MSI-X capability in config space
	TableSize   int    // number of vectors (Message Control[10:0] + 1)
	TableBIR    int    // BAR Indicator Register for table (0-5)
	TableOffset uint32 // byte offset of table within BAR
	PBABIR      int    // BAR Indicator Register for PBA
	PBAOffset   uint32 // byte offset of PBA within BAR
	Enabled     bool   // MSI-X Enable bit
	FuncMask    bool   // Function Mask bit
}

// MSIXEntry is a single MSI-X table row (16 bytes).
type MSIXEntry struct {
	AddrLo  uint32 `json:"addr_lo"`
	AddrHi  uint32 `json:"addr_hi"`
	Data    uint32 `json:"data"`
	Control uint32 `json:"control"` // bit 0 = mask
}

// ParseMSIXCap locates and parses the MSI-X capability. Nil if absent.
func ParseMSIXCap(cs *ConfigSpace) *MSIXInfo {
	caps := ParseCapabilities(cs)
	for _, cap := range caps {
		if cap.ID != CapIDMSIX {
			continue
		}
		if cap.Offset+12 > ConfigSpaceLegacySize {
			return nil
		}

		msgCtl := cs.ReadU16(cap.Offset + 2)
		tableOffsetReg := cs.ReadU32(cap.Offset + 4)
		pbaOffsetReg := cs.ReadU32(cap.Offset + 8)

		return &MSIXInfo{
			CapOffset:   cap.Offset,
			TableSize:   int(msgCtl&0x07FF) + 1,
			Enabled:     msgCtl&0x8000 != 0,
			FuncMask:    msgCtl&0x4000 != 0,
			TableBIR:    int(tableOffsetReg & 0x07),
			TableOffset: tableOffsetReg & 0xFFFFFFF8,
			PBABIR:      int(pbaOffsetReg & 0x07),
			PBAOffset:   pbaOffsetReg & 0xFFFFFFF8,
		}
	}
	return nil
}

// ReadMSIXTable reads MSI-X entries from BAR memory.
func ReadMSIXTable(barData []byte, info *MSIXInfo) []MSIXEntry {
	if info == nil || barData == nil {
		return nil
	}

	tableStart := int(info.TableOffset)
	tableEnd := tableStart + info.TableSize*16

	if tableEnd > len(barData) {
		available := (len(barData) - tableStart) / 16
		if available <= 0 {
			return nil
		}
		info = &MSIXInfo{TableSize: available, TableOffset: info.TableOffset}
	}

	entries := make([]MSIXEntry, 0, info.TableSize)
	for i := 0; i < info.TableSize; i++ {
		off := tableStart + i*16
		if off+16 > len(barData) {
			break
		}
		entries = append(entries, MSIXEntry{
			AddrLo:  leU32(barData, off),
			AddrHi:  leU32(barData, off+4),
			Data:    leU32(barData, off+8),
			Control: leU32(barData, off+12),
		})
	}
	return entries
}

func leU32(b []byte, off int) uint32 {
	return uint32(b[off]) | uint32(b[off+1])<<8 | uint32(b[off+2])<<16 | uint32(b[off+3])<<24
}
