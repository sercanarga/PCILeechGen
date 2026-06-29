package scrub

import (
	"strings"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

// auditDonor builds a donor that triggers one finding of every category:
// x16/Gen3 link + FLR + ASPM (PCIe cap), MSI-X, a prunable std cap (VPD),
// a prunable ext cap (SR-IOV), and no DSN.
func auditDonor() *pci.ConfigSpace {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	cs.WriteU16(0x06, 0x0010) // capabilities list present
	cs.WriteU8(0x34, 0x40)    // cap pointer

	// chain: MSI-X (0x40) -> VPD (0x50) -> PCIe (0x60, last)
	cs.WriteU8(0x40, pci.CapIDMSIX)
	cs.WriteU8(0x41, 0x50)
	cs.WriteU8(0x50, pci.CapIDVPD)
	cs.WriteU8(0x51, 0x60)
	cs.WriteU8(0x60, pci.CapIDPCIExpress)
	cs.WriteU8(0x61, 0x00) // last std cap -> Data extends to end (>=16 bytes)

	cs.WriteU32(0x64, 1<<28)      // DevCap: FLR capable
	cs.WriteU32(0x6C, 0x00000D03) // LinkCap: x16 (0x10<<4), Gen3 (0x3), ASPM L0s+L1 (0xC00)

	// ext cap chain: SR-IOV at 0x100, last; no DSN anywhere
	cs.WriteU32(0x100, uint32(pci.ExtCapIDSRIOV)|(1<<16))
	return cs
}

func TestAudit_AllCategories(t *testing.T) {
	donor := auditDonor()
	b := &board.Board{Name: "test_x1", PCIeLanes: 1, MaxLinkSpeed: 2} // x1, Gen2

	findings := Audit(donor, b)
	high, medium, low := AuditSummary(findings)

	if high != 2 {
		t.Errorf("HIGH findings = %d, want 2 (link width + speed); got: %s", high, titles(findings))
	}
	if medium != 4 {
		t.Errorf("MEDIUM findings = %d, want 4 (VPD, SR-IOV, no-DSN, FLR); got: %s", medium, titles(findings))
	}
	if low != 2 {
		t.Errorf("LOW findings = %d, want 2 (ASPM, MSI-X); got: %s", low, titles(findings))
	}

	want := []string{"Link width", "Link speed", "VPD", "SR-IOV", "No Device Serial Number", "FLR", "ASPM", "MSI-X"}
	all := titles(findings)
	for _, w := range want {
		if !strings.Contains(all, w) {
			t.Errorf("missing finding containing %q; got: %s", w, all)
		}
	}
}

// A well-matched donor (x1, Gen2, no prunable caps, has DSN) produces no
// HIGH or MEDIUM findings.
func TestAudit_WellMatchedDonor(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	cs.WriteU16(0x06, 0x0010)
	cs.WriteU8(0x34, 0x40)
	cs.WriteU8(0x40, pci.CapIDPCIExpress)
	cs.WriteU8(0x41, 0x00)
	cs.WriteU32(0x44, 0)          // DevCap: no FLR
	cs.WriteU32(0x4C, 0x00000011) // LinkCap: x1 (0x1<<4), Gen1, no ASPM
	// DSN present at 0x100
	cs.WriteU32(0x100, uint32(pci.ExtCapIDDeviceSerialNumber)|(1<<16))

	b := &board.Board{Name: "test_x1", PCIeLanes: 1, MaxLinkSpeed: 2}
	high, medium, _ := AuditSummary(Audit(cs, b))
	if high != 0 || medium != 0 {
		t.Errorf("well-matched donor: high=%d medium=%d, want 0/0", high, medium)
	}
}

// A donor with an Expansion ROM BAR must produce an option-ROM finding.
func TestAudit_OptionROM(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	cs.WriteU16(0x06, 0x0010)
	cs.WriteU8(0x34, 0x40)
	cs.WriteU8(0x40, pci.CapIDPCIExpress)
	cs.WriteU8(0x41, 0x00)
	cs.WriteU32(0x4C, 0x00000011) // x1 Gen1, no ASPM
	cs.WriteU32(0x100, uint32(pci.ExtCapIDDeviceSerialNumber)|(1<<16))
	cs.WriteU32(0x30, 0xFE000001) // Expansion ROM BAR: base + enable

	b := &board.Board{Name: "test_x1", PCIeLanes: 1, MaxLinkSpeed: 2}
	if !strings.Contains(titles(Audit(cs, b)), "Expansion ROM") {
		t.Errorf("expected option-ROM finding; got: %s", titles(Audit(cs, b)))
	}
}

func TestAudit_NilDonor(t *testing.T) {
	if f := Audit(nil, &board.Board{PCIeLanes: 1}); f != nil {
		t.Errorf("nil donor should return nil, got %d findings", len(f))
	}
}

func titles(f []Finding) string {
	var s []string
	for _, x := range f {
		s = append(s, x.Severity.String()+":"+x.Title)
	}
	return strings.Join(s, " | ")
}
