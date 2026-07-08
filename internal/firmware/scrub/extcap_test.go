package scrub

import (
	"testing"

	"github.com/sercanarga/pcileechgen/internal/firmware/overlay"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

func TestFilterExtCaps_RelocatePreservesBody(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize

	// ATS (unsafe) at 0x100, size 8.
	cs.WriteU32(0x100, makeExtCapHeader(pci.ExtCapIDATS, 0x108))
	// DSN (safe) at 0x108, size 64, body 0xAB.
	cs.WriteU32(0x108, makeExtCapHeader(pci.ExtCapIDDeviceSerialNumber, 0x148))
	for off := 0x10C; off < 0x148; off++ {
		cs.WriteU8(off, 0xAB)
	}
	// Trailing safe cap so DSN parses as exactly 64 bytes.
	cs.WriteU32(0x148, makeExtCapHeader(pci.ExtCapIDLTR, 0))

	removed := FilterExtCapabilities(cs, overlay.NewMap(cs))
	if len(removed) != 1 {
		t.Fatalf("expected ATS only removed, got %v", removed)
	}

	// Relocated DSN body (header now at 0x100) must be fully intact.
	for off := 0x104; off < 0x140; off++ {
		if got := cs.ReadU8(off); got != 0xAB {
			t.Errorf("relocated DSN body corrupted at 0x%03X: got 0x%02X, want 0xAB", off, got)
		}
	}
	// Old overlap-zero window: dest=[0x100,0x140) vs zero=[0x108,0x148).
	for off := 0x108; off < 0x140; off++ {
		if got := cs.ReadU8(off); got != 0xAB {
			t.Fatalf("overlap-zero corruption at 0x%03X: got 0x%02X, want 0xAB", off, got)
		}
	}

	hdr := cs.ReadU32(0x100)
	if uint16(hdr&0xFFFF) != pci.ExtCapIDDeviceSerialNumber {
		t.Errorf("first cap id = 0x%04X, want DSN", hdr&0xFFFF)
	}
}
