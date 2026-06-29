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

	// With no capabilities present (status cap-list bit unset), the whole
	// 0x40-0xFF window stays writable - there are no capability registers to
	// lock. Read-only field locking only applies to actual capabilities; see
	// TestWritemask_CapabilityReadOnlyFieldsLocked.
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

func TestWritemask_CommandWritableStatusReadOnly(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	cs.WriteU32(0x10, 0xFFFFF000)

	wm := GenerateWritemaskCOE(cs)
	dwords := parseCOEDwords(t, wm)

	// Command (lower 16) is RW; Status (upper 16) is RO/RW1C - writes ignored
	// like real hardware so a write-then-readback can't change it.
	if dwords[1] != "0000ffff" {
		t.Errorf("Command:Status mask = %s, want 0000ffff (Command RW, Status RO)", dwords[1])
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

// Capability read-only fields must be locked in the shadow so a detector can't
// write to a cap ID / PMC / PCIe Capabilities / DevCap / LinkCap / MSI-X offset
// register and read it back changed (real hardware ignores such writes).
func TestWritemask_CapabilityReadOnlyFieldsLocked(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	cs.WriteU16(0x06, 0x0010) // capabilities list present
	cs.WriteU8(0x34, 0x40)    // cap pointer
	// chain: PM (0x40) -> MSI-X (0x50) -> PCIe (0x60, last)
	cs.WriteU8(0x40, pci.CapIDPowerManagement)
	cs.WriteU8(0x41, 0x50)
	cs.WriteU8(0x50, pci.CapIDMSIX)
	cs.WriteU8(0x51, 0x60)
	cs.WriteU8(0x60, pci.CapIDPCIExpress)
	cs.WriteU8(0x61, 0x00)

	dwords := parseCOEDwords(t, GenerateWritemaskCOE(cs))

	checks := map[int]string{
		16: "00000000", // 0x40 PM header (ID+next+PMC all RO)
		17: "ffffffff", // 0x44 PMCSR stays writable (power state control)
		20: "c0000000", // 0x50 MSI-X header: only Enable+FuncMask writable
		21: "00000000", // 0x54 MSI-X Table Offset/BIR (RO)
		22: "00000000", // 0x58 MSI-X PBA Offset/BIR (RO)
		24: "00000000", // 0x60 PCIe header (ID+next+PCIe Caps reg RO)
		25: "00000000", // 0x64 Device Capabilities (RO)
		26: "0000ffff", // 0x68 Device Control (RW) / Device Status (RO/RW1C)
		27: "00000000", // 0x6C Link Capabilities (RO)
		28: "0000ffff", // 0x70 Link Control (RW) / Link Status (RO/RW1C)
		30: "0000ffff", // 0x78 Slot Control (RW) / Slot Status (RO/RW1C)
		31: "0000ffff", // 0x7C Root Control (RW) / Root Capabilities (RO)
		32: "00000000", // 0x80 Root Status (RO/RW1C)
		33: "00000000", // 0x84 Device Capabilities 2 (RO)
		35: "00000000", // 0x8C Link Capabilities 2 (RO)
	}
	for dw, want := range checks {
		if dwords[dw] != want {
			t.Errorf("DWORD[%d] (0x%02X) mask = %s, want %s", dw, dw*4, dwords[dw], want)
		}
	}
}

// MSI Message Control: only Enable (bit0) + Multiple Message Enable (bits 6:4)
// are writable; the capable/64-bit/PVM bits are RO. Address/Data stay writable.
func TestWritemask_MSIMessageControlLocked(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	cs.WriteU16(0x06, 0x0010) // capabilities list present
	cs.WriteU8(0x34, 0x40)    // cap pointer
	cs.WriteU8(0x40, pci.CapIDMSI)
	cs.WriteU8(0x41, 0x00)    // last cap
	cs.WriteU16(0x42, 0x0080) // Message Control: 64-bit capable

	dwords := parseCOEDwords(t, GenerateWritemaskCOE(cs))

	// DWORD 16 (0x40): [ID|next|MsgCtl]; only bits 16 (Enable) + 20-22 (MME) RW.
	if dwords[16] != "00710000" {
		t.Errorf("MSI header mask = %s, want 00710000 (Enable+MME RW, ID/next/MsgCtl-RO locked)", dwords[16])
	}
	// MSI Address (0x44) stays writable for the OS to program.
	if dwords[17] != "ffffffff" {
		t.Errorf("MSI Address mask = %s, want ffffffff (OS programs it)", dwords[17])
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
