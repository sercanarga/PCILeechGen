package pci

import "testing"

func TestParseMSIXCap(t *testing.T) {
	cs := NewConfigSpace()
	cs.Size = ConfigSpaceLegacySize

	// capability pointer
	cs.WriteU16(0x06, 0x0010)
	cs.WriteU8(0x34, 0x40)

	// MSI-X cap at 0x40: ID=0x11, next=0
	cs.WriteU8(0x40, CapIDMSIX)
	cs.WriteU8(0x41, 0x00)
	// Message Control: table size = 129 (0x0080), enabled
	cs.WriteU16(0x42, 0x8080)
	// Table offset register: BIR=0, offset=0x2000
	cs.WriteU32(0x44, 0x00002000)
	// PBA offset register: BIR=0, offset=0x3000
	cs.WriteU32(0x48, 0x00003000)

	info := ParseMSIXCap(cs)
	if info == nil {
		t.Fatal("expected MSI-X cap, got nil")
	}
	if info.TableSize != 0x81 {
		t.Errorf("TableSize = %d, want 129", info.TableSize)
	}
	if info.TableBIR != 0 {
		t.Errorf("TableBIR = %d, want 0", info.TableBIR)
	}
	if info.TableOffset != 0x2000 {
		t.Errorf("TableOffset = 0x%X, want 0x2000", info.TableOffset)
	}
	if info.PBAOffset != 0x3000 {
		t.Errorf("PBAOffset = 0x%X, want 0x3000", info.PBAOffset)
	}
	if !info.Enabled {
		t.Error("should be enabled")
	}
}

func TestParseMSIXCap_Absent(t *testing.T) {
	cs := NewConfigSpace()
	cs.Size = ConfigSpaceLegacySize

	cs.WriteU16(0x06, 0x0010)
	cs.WriteU8(0x34, 0x40)
	cs.WriteU8(0x40, CapIDPowerManagement)
	cs.WriteU8(0x41, 0x00)

	if ParseMSIXCap(cs) != nil {
		t.Error("should return nil when no MSI-X cap")
	}
}

func TestReadMSIXTable(t *testing.T) {
	info := &MSIXInfo{
		TableSize:   2,
		TableOffset: 0x100,
	}
	bar := make([]byte, 0x200)

	// first entry
	bar[0x100] = 0x10
	bar[0x104] = 0x20
	bar[0x108] = 0x42
	bar[0x10C] = 0x01 // masked

	// second entry
	bar[0x110] = 0xAA
	bar[0x114] = 0xBB

	entries := ReadMSIXTable(bar, info)
	if len(entries) != 2 {
		t.Fatalf("got %d entries, want 2", len(entries))
	}
	if entries[0].AddrLo != 0x10 {
		t.Errorf("entry 0 AddrLo = 0x%X", entries[0].AddrLo)
	}
	if entries[0].Control != 0x01 {
		t.Errorf("entry 0 Control = 0x%X, want 1 (masked)", entries[0].Control)
	}
}

func TestReadMSIXTable_Nil(t *testing.T) {
	if ReadMSIXTable(nil, nil) != nil {
		t.Error("nil inputs should return nil")
	}
	if ReadMSIXTable([]byte{0}, nil) != nil {
		t.Error("nil info should return nil")
	}
}

func TestReadMSIXTable_Truncated(t *testing.T) {
	info := &MSIXInfo{
		TableSize:   100,
		TableOffset: 0x00,
	}
	bar := make([]byte, 48) // only fits 3 entries

	entries := ReadMSIXTable(bar, info)
	if len(entries) != 3 {
		t.Errorf("got %d entries, want 3 (truncated)", len(entries))
	}
}
