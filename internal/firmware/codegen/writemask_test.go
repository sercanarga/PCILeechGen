package codegen

import (
	"strings"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/pci"
)

func TestGenerateWritemaskCOE(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	cs.WriteU16(0x00, 0x8086)
	cs.WriteU32(0x10, 0xFFFFF000)

	wm := GenerateWritemaskCOE(cs)

	if !strings.Contains(wm, "memory_initialization_radix=16") {
		t.Error("writemask COE should contain radix")
	}

	lines := strings.Split(wm, "\n")
	dwords := make([]string, 0, len(lines))
	for _, l := range lines {
		l = strings.TrimSpace(l)
		if l == "" || strings.HasPrefix(l, ";") || strings.HasPrefix(l, "memory") {
			continue
		}
		l = strings.TrimSuffix(l, ",")
		l = strings.TrimSuffix(l, ";")
		dwords = append(dwords, l)
	}

	if len(dwords) != 1024 {
		t.Fatalf("writemask should have 1024 DWORDs, got %d", len(dwords))
	}

	if dwords[0] != "00000000" {
		t.Errorf("VID:DID mask should be 00000000, got %s", dwords[0])
	}

	if dwords[2] != "00000000" {
		t.Errorf("Rev:ClassCode mask should be 00000000, got %s", dwords[2])
	}

	if dwords[4] != "fffff000" {
		t.Errorf("BAR0 mask should be fffff000 (4KB size e.g.), got %s", dwords[4])
	}

	if dwords[5] != "00000000" {
		t.Errorf("BAR1 mask should be 00000000 (unused), got %s", dwords[5])
	}

	if dwords[11] != "00000000" {
		t.Errorf("SubsysIDs mask should be 00000000, got %s", dwords[11])
	}

	if dwords[15] != "000000ff" {
		t.Errorf("IntLine/IntPin mask should be 000000ff, got %s", dwords[15])
	}

	// capability region (0x40-0xFF) must be fully writable; locking desyncs shadow BRAM from the IP core.
	for i := 0x40 / 4; i < 0x100/4; i++ {
		if dwords[i] != "ffffffff" {
			t.Errorf("cap region DWORD[%d] should be ffffffff, got %s", i, dwords[i])
		}
	}

	for i := 0x100 / 4; i < len(dwords); i++ {
		if dwords[i] != "00000000" {
			t.Errorf("ext cap DWORD[%d] should be 00000000, got %s", i, dwords[i])
			break
		}
	}
}

func TestWritemask_ScrubbedBAR0FromZero(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	cs.WriteU16(0x00, 0x1102)
	cs.WriteU16(0x02, 0x0012)
	cs.WriteU32(0x10, 0xFFFFF000)

	wm := GenerateWritemaskCOE(cs)
	dwords := parseCOEDwords(t, wm)

	if dwords[4] != "fffff000" {
		t.Errorf("BAR0 mask = %s, want fffff000 (4KB size mask e.g.)", dwords[4])
	}
}

func TestWritemask_BARSizingSimulation(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	cs.WriteU32(0x10, 0xFFFFF000)

	wm := GenerateWritemaskCOE(cs)
	dwords := parseCOEDwords(t, wm)

	if dwords[4] != "fffff000" {
		t.Fatalf("BAR0 mask = %s, want fffff000", dwords[4])
	}

	hostWrite := uint32(0xFFFFFFFF)
	mask := uint32(0xFFFFF000)
	bramResult := hostWrite & mask
	if bramResult != 0xFFFFF000 {
		t.Errorf("BAR sizing result = 0x%08X, want 0xFFFFF000", bramResult)
	}
	sizeBits := bramResult & ^uint32(0x0F)
	barSize := ^sizeBits + 1
	if barSize != 4096 {
		t.Errorf("decoded BAR size = %d, want 4096 (for this 4k mask test case)", barSize)
	}
}

func TestWritemask_IOBAROverwrittenToMemory(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	cs.WriteU32(0x10, 0xFFFFF000)

	wm := GenerateWritemaskCOE(cs)
	dwords := parseCOEDwords(t, wm)

	if dwords[4] == "00000000" {
		t.Error("BAR0 mask should not be zero after scrubber creates memory BAR")
	}
	if dwords[4] != "fffff000" {
		t.Errorf("BAR0 mask = %s, want fffff000", dwords[4])
	}
}

func TestWritemask_IdentityRegistersReadOnly(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	cs.WriteU16(0x00, 0x1102)
	cs.WriteU16(0x02, 0x0012)
	cs.WriteU32(0x08, 0x04030000)
	cs.WriteU16(0x2C, 0x1102)
	cs.WriteU16(0x2E, 0x0012)
	cs.WriteU32(0x10, 0xFFFFF000)

	wm := GenerateWritemaskCOE(cs)
	dwords := parseCOEDwords(t, wm)

	roRegs := map[int]string{
		0:  "VID:DID",
		2:  "RevisionID:ClassCode",
		10: "CardBus CIS",
		11: "SubsysVID:SubsysDID",
		12: "Expansion ROM",
		13: "CapPtr",
		14: "Reserved",
	}
	for dw, name := range roRegs {
		if dwords[dw] != "00000000" {
			t.Errorf("%s (DWORD %d) mask = %s, want 00000000 (read-only)", name, dw, dwords[dw])
		}
	}
}

func TestWritemask_CommandStatusWritable(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	cs.WriteU32(0x10, 0xFFFFF000)

	wm := GenerateWritemaskCOE(cs)
	dwords := parseCOEDwords(t, wm)

	if dwords[1] != "ffffffff" {
		t.Errorf("Command:Status mask = %s, want ffffffff", dwords[1])
	}
}

func TestWritemask_PrefetchableBAR(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	cs.WriteU32(0x10, 0xFFFFF008)

	wm := GenerateWritemaskCOE(cs)
	dwords := parseCOEDwords(t, wm)

	if dwords[4] != "fffff000" {
		t.Errorf("prefetchable BAR mask = %s, want fffff000", dwords[4])
	}
}

func TestWritemask_MultipleBARs(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	cs.WriteU32(0x10, 0xFFFFF000)
	cs.WriteU32(0x14, 0x00000000)
	cs.WriteU32(0x18, 0xFFFF0000)

	wm := GenerateWritemaskCOE(cs)
	dwords := parseCOEDwords(t, wm)

	if dwords[4] != "fffff000" {
		t.Errorf("BAR0 mask = %s, want fffff000", dwords[4])
	}
	if dwords[5] != "00000000" {
		t.Errorf("BAR1 mask = %s, want 00000000 (unused)", dwords[5])
	}
	if dwords[6] != "ffff0000" {
		t.Errorf("BAR2 mask = %s, want ffff0000", dwords[6])
	}
}

func TestWritemask_64BitMemoryBARPair(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	// 64-bit non-prefetchable memory BAR: lower dword type bits [2:1]=10b (0x4)
	cs.WriteU32(0x10, 0xFFFFFFF4) // BAR0 lower dword (64-bit, address bits set)
	cs.WriteU32(0x14, 0x00000000) // BAR1 upper dword

	wm := GenerateWritemaskCOE(cs)
	dwords := parseCOEDwords(t, wm)

	if dwords[4] != "fffffff0" {
		t.Errorf("BAR0 lower dword mask = %s, want fffffff0 (type bits [3:0] read-only)", dwords[4])
	}
	if dwords[5] != "ffffffff" {
		t.Errorf("BAR1 upper dword mask = %s, want ffffffff (64-bit upper dword is pure address)", dwords[5])
	}
}

func TestWritemask_64BitPrefetchableMemoryBARPair(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	// 64-bit prefetchable memory BAR: type bits [3:0]=0xC (prefetch + 64-bit)
	cs.WriteU32(0x18, 0xFFFFFFFC) // BAR2 lower dword
	cs.WriteU32(0x1C, 0x12345678) // BAR3 upper dword (nonzero, must still be fully writable)

	wm := GenerateWritemaskCOE(cs)
	dwords := parseCOEDwords(t, wm)

	if dwords[6] != "fffffff0" {
		t.Errorf("BAR2 lower dword mask = %s, want fffffff0", dwords[6])
	}
	if dwords[7] != "ffffffff" {
		t.Errorf("BAR3 upper dword mask = %s, want ffffffff", dwords[7])
	}
}

func parseCOEDwords(t *testing.T, coe string) []string {
	t.Helper()
	lines := strings.Split(coe, "\n")
	dwords := make([]string, 0, len(lines))
	for _, l := range lines {
		l = strings.TrimSpace(l)
		if l == "" || strings.HasPrefix(l, ";") || strings.HasPrefix(l, "memory") {
			continue
		}
		l = strings.TrimSuffix(l, ",")
		l = strings.TrimSuffix(l, ";")
		dwords = append(dwords, l)
	}
	if len(dwords) != 1024 {
		t.Fatalf("expected 1024 DWORDs, got %d", len(dwords))
	}
	return dwords
}
