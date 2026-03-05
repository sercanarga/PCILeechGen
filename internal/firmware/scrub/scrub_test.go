package scrub

import (
	"testing"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/firmware"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

func TestScrubConfigSpace(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize

	cs.WriteU16(0x00, 0x8086) // Vendor ID
	cs.WriteU16(0x02, 0x1533) // Device ID
	cs.WriteU16(0x04, 0x0507) // Command
	cs.WriteU16(0x06, 0xFBB0) // Status
	cs.WriteU8(0x08, 0x03)    // Revision ID
	cs.WriteU8(0x0C, 0x10)    // Cache Line Size
	cs.WriteU8(0x0D, 0x40)    // Latency Timer
	cs.WriteU8(0x0F, 0xC0)    // BIST
	cs.WriteU8(0x3C, 0x0B)    // Interrupt Line

	scrubbed := ScrubConfigSpace(cs, nil)

	if scrubbed.VendorID() != 0x8086 {
		t.Errorf("VendorID changed: 0x%04x", scrubbed.VendorID())
	}
	if scrubbed.DeviceID() != 0x1533 {
		t.Errorf("DeviceID changed: 0x%04x", scrubbed.DeviceID())
	}
	if scrubbed.RevisionID() != 0x03 {
		t.Errorf("RevisionID changed: 0x%02x", scrubbed.RevisionID())
	}
	if scrubbed.BIST() != 0x00 {
		t.Errorf("BIST not cleared: 0x%02x", scrubbed.BIST())
	}
	if scrubbed.InterruptLine() != 0x00 {
		t.Errorf("InterruptLine not cleared: 0x%02x", scrubbed.InterruptLine())
	}
	if scrubbed.LatencyTimer() != 0x00 {
		t.Errorf("LatencyTimer not cleared: 0x%02x", scrubbed.LatencyTimer())
	}
	if scrubbed.CacheLineSize() != 0x00 {
		t.Errorf("CacheLineSize not cleared: 0x%02x", scrubbed.CacheLineSize())
	}
	status := scrubbed.Status()
	if status&0x0010 == 0 {
		t.Error("Status capability bit should be preserved")
	}
	if status&0xF100 != 0 {
		t.Errorf("Status error bits not cleared: 0x%04x", status)
	}
	if cs.BIST() != 0xC0 {
		t.Error("Original BIST was modified!")
	}
}

func makeExtCapHeader(id uint16, version uint8, nextOffset int) uint32 {
	return uint32(id) | uint32(version)<<16 | uint32(nextOffset)<<20
}

func TestFilterExtCaps_RemoveMiddle(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize

	cs.WriteU32(0x100, makeExtCapHeader(pci.ExtCapIDAER, 1, 0x150))
	cs.WriteU32(0x150, makeExtCapHeader(pci.ExtCapIDSRIOV, 1, 0x200))
	cs.WriteU32(0x200, makeExtCapHeader(pci.ExtCapIDDeviceSerialNumber, 1, 0x250))
	cs.WriteU32(0x250, makeExtCapHeader(pci.ExtCapIDLTR, 1, 0))

	removed := FilterExtCapabilities(cs)

	if len(removed) != 1 {
		t.Fatalf("Expected 1 removed cap, got %d: %v", len(removed), removed)
	}
	if cs.ReadU32(0x150) != 0 {
		t.Errorf("SR-IOV header not zeroed: 0x%08x", cs.ReadU32(0x150))
	}

	caps := pci.ParseExtCapabilities(cs)
	if len(caps) != 3 {
		t.Errorf("Expected 3 remaining caps, got %d", len(caps))
	}
}

func TestFilterExtCaps_AllRemoved(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize

	cs.WriteU32(0x100, makeExtCapHeader(pci.ExtCapIDSRIOV, 1, 0x150))
	cs.WriteU32(0x150, makeExtCapHeader(pci.ExtCapIDResizableBAR, 1, 0))

	removed := FilterExtCapabilities(cs)
	if len(removed) != 2 {
		t.Fatalf("Expected 2 removed caps, got %d", len(removed))
	}
	if cs.ReadU32(0x100) != 0 {
		t.Errorf("First ext cap header should be zero: 0x%08x", cs.ReadU32(0x100))
	}
}

func TestScrubConfigSpace_ClampBAR0(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize

	cs.WriteU16(0x00, 0x144D)
	cs.WriteU16(0x06, 0x0010)
	cs.WriteU32(0x10, 0xFFFFC004) // mem64, 16 KB
	cs.WriteU32(0x14, 0xFFFFFFFF) // upper 32 bits
	cs.WriteU32(0x18, 0xFFFF0000) // mem32, 64 KB

	scrubbed := ScrubConfigSpace(cs, nil)

	bar0 := scrubbed.BAR(0)
	if bar0 != 0xFFFFF004 {
		t.Errorf("BAR0 should be clamped to 4 KB: got 0x%08x, want 0xFFFFF004", bar0)
	}
	bar1 := scrubbed.BAR(1)
	if bar1 != 0 {
		t.Errorf("BAR1 (upper 64-bit) should be zero: got 0x%08x", bar1)
	}
	bar2 := scrubbed.BAR(2)
	if bar2 != 0xFFFFF000 {
		t.Errorf("BAR2 should be clamped to 4 KB: got 0x%08x, want 0xFFFFF000", bar2)
	}
}

func TestIsUnsafeExtCap(t *testing.T) {
	if IsUnsafeExtCap(pci.ExtCapIDAER) {
		t.Error("AER should be safe")
	}
	if !IsUnsafeExtCap(pci.ExtCapIDSRIOV) {
		t.Error("SR-IOV should be unsafe")
	}
	if !IsUnsafeExtCap(pci.ExtCapIDResizableBAR) {
		t.Error("Resizable BAR should be unsafe")
	}
}

func TestScrubConfigSpace_ClampLinkCapability(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize

	cs.WriteU16(0x00, 0x144D)
	cs.WriteU16(0x06, 0x0010)
	cs.WriteU8(0x34, 0x70)

	cs.WriteU8(0x70, pci.CapIDPCIExpress)
	cs.WriteU8(0x71, 0x00)

	linkCap := uint32(4) | (uint32(4) << 4) // Gen4 x4
	cs.WriteU32(0x7C, linkCap)
	cs.WriteU16(0x82, uint16(4)|(uint16(4)<<4))
	cs.WriteU16(0xA0, 4)
	cs.WriteU32(0x9C, (0x1E)<<1)

	b := &board.Board{PCIeLanes: 1}
	scrubbed := ScrubConfigSpace(cs, b)

	scrubbedLinkCap := scrubbed.ReadU32(0x7C)
	speed := uint8(scrubbedLinkCap & 0x0F)
	width := uint8((scrubbedLinkCap >> 4) & 0x3F)
	if speed != firmware.LinkSpeedGen2 {
		t.Errorf("Link Cap Max Speed: got %d, want %d (Gen2)", speed, firmware.LinkSpeedGen2)
	}
	if width != 1 {
		t.Errorf("Link Cap Max Width: got %d, want 1", width)
	}
}

func TestScrubConfigSpace_LinkCtl2ZeroSpeed(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize

	cs.WriteU16(0x00, 0x1102)
	cs.WriteU16(0x06, 0x0010)
	cs.WriteU8(0x34, 0x70)

	cs.WriteU8(0x70, pci.CapIDPCIExpress)
	cs.WriteU8(0x71, 0x00)

	// LinkCap at 0x7C: Gen1 x1
	cs.WriteU32(0x7C, uint32(1)|(uint32(1)<<4))
	// LinkCtl2 at 0xA0: target speed = 0 (invalid)
	cs.WriteU16(0xA0, 0x0000)

	scrubbed := ScrubConfigSpace(cs, nil)

	lctl2 := scrubbed.ReadU16(0xA0)
	tgtSpeed := uint8(lctl2 & 0x0F)
	if tgtSpeed == 0 {
		t.Errorf("LinkCtl2 target speed should not be 0 after scrub, got 0x%04x", lctl2)
	}
	if tgtSpeed > firmware.LinkSpeedGen2 {
		t.Errorf("LinkCtl2 target speed should be <= Gen2, got %d", tgtSpeed)
	}
}

func TestScrubConfigSpace_LinkCap2VectorClamp(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize

	cs.WriteU16(0x00, 0x1102)
	cs.WriteU16(0x06, 0x0010)
	cs.WriteU8(0x34, 0x70)

	cs.WriteU8(0x70, pci.CapIDPCIExpress)
	cs.WriteU8(0x71, 0x00)

	// LinkCap at 0x7C: Gen1 x1
	cs.WriteU32(0x7C, uint32(1)|(uint32(1)<<4))
	// LinkCap2 at 0x9C: only Gen1 supported (bit 1 only)
	cs.WriteU32(0x9C, 0x02)

	b := &board.Board{PCIeLanes: 1}
	scrubbed := ScrubConfigSpace(cs, b)

	lc2 := scrubbed.ReadU32(0x9C)
	vec := lc2 & 0xFE
	// donor only had Gen1, board supports Gen2 — vector should remain Gen1 only
	if vec != 0x02 {
		t.Errorf("LinkCap2 vector should be 0x02 (Gen1 only), got 0x%02x", vec)
	}
}

func TestScrubConfigSpace_ClampDeviceCapability(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize

	cs.WriteU16(0x00, 0x144D)
	cs.WriteU16(0x06, 0x0010)
	cs.WriteU8(0x34, 0x70)

	cs.WriteU8(0x70, pci.CapIDPCIExpress)
	cs.WriteU8(0x71, 0x00)

	devCap := uint32(2) | (1 << 3) | (1 << 5) | (0x07 << 6)
	cs.WriteU32(0x74, devCap)
	devCtl := uint16(2<<5) | (1 << 8) | (1 << 9) | (5 << 12)
	cs.WriteU16(0x78, devCtl)
	devCap2 := uint32(1<<16) | (1 << 17) | (0x0F)
	cs.WriteU32(0x94, devCap2)

	scrubbed := ScrubConfigSpace(cs, nil)

	scrubbedDevCap := scrubbed.ReadU32(0x74)
	if scrubbedDevCap&0x07 != 0 {
		t.Errorf("Device Cap MPS should be 0 (128B), got %d", scrubbedDevCap&0x07)
	}
	if (scrubbedDevCap>>3)&0x03 != 0 {
		t.Error("Device Cap Phantom should be 0")
	}
	if (scrubbedDevCap>>5)&0x01 != 0 {
		t.Error("Device Cap ExtTag should be 0")
	}

	scrubbedDevCtl := scrubbed.ReadU16(0x78)
	ctlMRRS := (scrubbedDevCtl >> 12) & 0x07
	if ctlMRRS != 2 {
		t.Errorf("Device Control MRRS should be clamped to 2, got %d", ctlMRRS)
	}

	scrubbedDevCap2 := scrubbed.ReadU32(0x94)
	if scrubbedDevCap2&(1<<16) != 0 {
		t.Error("Device Cap 2: 10-bit Tag Completer should be 0")
	}
}

func TestFakeRenesasFirmwareReady(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize

	cs.WriteU16(0x00, 0x1912) // Renesas
	cs.WriteU16(0x02, 0x0014)
	cs.WriteU8(0x09, 0x30)
	cs.WriteU8(0x0A, 0x03)
	cs.WriteU8(0x0B, 0x0C)
	cs.WriteU16(0x06, 0x0010)
	cs.WriteU8(0xF4, 0x00)
	cs.WriteU16(0xF6, 0x8000)

	scrubbed := ScrubConfigSpace(cs, nil)

	fwStatus := scrubbed.ReadU8(0xF4)
	if fwStatus&0x10 == 0 {
		t.Errorf("Renesas FW Status SUCCESS bit should be set, got 0x%02x", fwStatus)
	}
	if fwStatus&0x80 == 0 {
		t.Errorf("Renesas FW Status LOCK bit should be set, got 0x%02x", fwStatus)
	}
	romStatus := scrubbed.ReadU16(0xF6)
	if romStatus&0x0010 == 0 {
		t.Errorf("Renesas ROM Status SUCCESS bit should be set, got 0x%04x", romStatus)
	}
}

func TestZeroVendorRegisters(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize

	cs.WriteU16(0x00, 0x8086)
	cs.WriteU16(0x06, 0x0010)
	cs.WriteU8(0x34, 0x40)

	cs.WriteU8(0x40, pci.CapIDPowerManagement)
	cs.WriteU8(0x41, 0x70)
	cs.WriteU16(0x42, 0xC9C3)
	cs.WriteU16(0x44, 0x0008)
	cs.WriteU8(0x70, pci.CapIDPCIExpress)
	cs.WriteU8(0x71, 0x00)

	cs.WriteU32(0xB0, 0xDEADBEEF)
	cs.WriteU32(0xF4, 0xCAFEBABE)

	scrubbed := ScrubConfigSpace(cs, nil)

	if scrubbed.ReadU32(0xB0) != 0 {
		t.Errorf("vendor register at 0xB0 should be zeroed, got 0x%08x", scrubbed.ReadU32(0xB0))
	}
	if scrubbed.ReadU32(0xF4) != 0 {
		t.Errorf("vendor register at 0xF4 should be zeroed, got 0x%08x", scrubbed.ReadU32(0xF4))
	}
	if scrubbed.ReadU8(0x40) != pci.CapIDPowerManagement {
		t.Errorf("PM cap ID should be preserved, got 0x%02x", scrubbed.ReadU8(0x40))
	}
	if scrubbed.ReadU8(0x70) != pci.CapIDPCIExpress {
		t.Errorf("PCIe cap ID should be preserved, got 0x%02x", scrubbed.ReadU8(0x70))
	}
}

func TestComputeMSISize(t *testing.T) {
	tests := []struct {
		name   string
		msgCtl uint16
		wantSz int
	}{
		{"32bit_no_masking", 0x0000, 10},
		{"64bit_no_masking", 0x0080, 14},
		{"32bit_masking", 0x0100, 20},
		{"64bit_masking", 0x0180, 24},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := pci.NewConfigSpace()
			cs.WriteU8(0x50, pci.CapIDMSI)
			cs.WriteU8(0x51, 0x00)
			cs.WriteU16(0x52, tt.msgCtl)

			got := computeMSISize(cs, 0x50)
			if got != tt.wantSz {
				t.Errorf("computeMSISize() = %d, want %d", got, tt.wantSz)
			}
		})
	}
}

func TestZeroVendorRegisters_PreservesMSI64BitMasking(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize

	cs.WriteU16(0x00, 0x1102) // Creative
	cs.WriteU16(0x06, 0x0010) // caps present
	cs.WriteU8(0x34, 0x40)    // cap ptr

	// PM at 0x40, next -> MSI at 0x50
	cs.WriteU8(0x40, pci.CapIDPowerManagement)
	cs.WriteU8(0x41, 0x50)
	cs.WriteU16(0x42, 0xC9C3)
	cs.WriteU16(0x44, 0x0008)

	// MSI at 0x50: 64-bit + masking, occupies 24 bytes (0x50-0x67)
	cs.WriteU8(0x50, pci.CapIDMSI)
	cs.WriteU8(0x51, 0x70)        // next -> PCIe
	cs.WriteU16(0x52, 0x0180)     // 64bit + masking
	cs.WriteU32(0x54, 0xFEE00000) // addr lower
	cs.WriteU32(0x58, 0x00000000) // addr upper
	cs.WriteU16(0x5C, 0x4021)     // data
	cs.WriteU16(0x5E, 0x0000)     // reserved
	cs.WriteU32(0x60, 0x00000001) // mask bits
	cs.WriteU32(0x64, 0x00000000) // pending bits

	// PCIe at 0x70
	cs.WriteU8(0x70, pci.CapIDPCIExpress)
	cs.WriteU8(0x71, 0x00)

	// Some garbage in uncovered vendor region
	cs.WriteU32(0xB0, 0xDEADBEEF)

	scrubbed := ScrubConfigSpace(cs, nil)

	// cap header intact
	if scrubbed.ReadU8(0x50) != pci.CapIDMSI {
		t.Errorf("MSI cap ID gone, got 0x%02x", scrubbed.ReadU8(0x50))
	}
	// next pointer still chained to PCIe
	if scrubbed.ReadU8(0x51) != 0x70 {
		t.Errorf("MSI next ptr should be 0x70, got 0x%02x", scrubbed.ReadU8(0x51))
	}
	// 64bit + masking flags survived
	msgCtl := scrubbed.ReadU16(0x52)
	if msgCtl&0x0180 != 0x0180 {
		t.Errorf("MSI msgctl 64bit+masking lost, got 0x%04x", msgCtl)
	}
	// address not wiped
	if scrubbed.ReadU32(0x54) != 0xFEE00000 {
		t.Errorf("MSI addr lower wiped, got 0x%08x", scrubbed.ReadU32(0x54))
	}
	// mask bits at 0x60 inside the 24-byte window
	if scrubbed.ReadU32(0x60) != 0x00000001 {
		t.Errorf("MSI mask bits wiped, got 0x%08x", scrubbed.ReadU32(0x60))
	}
	// PCIe still where it should be
	if scrubbed.ReadU8(0x70) != pci.CapIDPCIExpress {
		t.Errorf("PCIe cap ID gone at 0x70, got 0x%02x", scrubbed.ReadU8(0x70))
	}
	// uncovered region zeroed
	if scrubbed.ReadU32(0xB0) != 0 {
		t.Errorf("vendor register at 0xB0 not zeroed, got 0x%08x", scrubbed.ReadU32(0xB0))
	}
}

func TestRelocateMSIXToBRAM_TableOutside(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize

	cs.WriteU16(0x06, 0x0010)
	cs.WriteU8(0x34, 0x90)

	cs.WriteU8(0x90, pci.CapIDMSIX)
	cs.WriteU8(0x91, 0x00)
	cs.WriteU16(0x92, 0x8007)     // 8 vectors, enabled
	cs.WriteU32(0x94, 0x00002000) // table at BAR0+0x2000 (outside 4KB)
	cs.WriteU32(0x98, 0x00002080) // PBA at BAR0+0x2080
	cs.WriteU32(0x10, 0xFFFFF004)

	scrubbed := ScrubConfigSpace(cs, nil)

	// MSI-X should remain enabled after relocation
	msgCtl := scrubbed.ReadU16(0x92)
	if msgCtl&0x8000 == 0 {
		t.Errorf("MSI-X should stay enabled after relocation, got 0x%04x", msgCtl)
	}
	if msgCtl&0x4000 != 0 {
		t.Errorf("MSI-X Function Mask should be cleared, got 0x%04x", msgCtl)
	}

	// table offset should be relocated to 0x1000
	tableReg := scrubbed.ReadU32(0x94)
	tableOff := tableReg &^ 0x07
	if tableOff != 0x1000 {
		t.Errorf("MSI-X table offset should be relocated to 0x1000, got 0x%X", tableOff)
	}

	// PBA offset should follow table (8 vectors * 16 bytes = 128 = 0x80)
	pbaReg := scrubbed.ReadU32(0x98)
	pbaOff := pbaReg &^ 0x07
	if pbaOff != 0x1080 {
		t.Errorf("MSI-X PBA offset should be 0x1080, got 0x%X", pbaOff)
	}
}

func TestRelocateMSIXToBRAM_TableInside(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize

	cs.WriteU16(0x06, 0x0010)
	cs.WriteU8(0x34, 0x90)

	cs.WriteU8(0x90, pci.CapIDMSIX)
	cs.WriteU8(0x91, 0x00)
	cs.WriteU16(0x92, 0x8003)     // 4 vectors, enabled
	cs.WriteU32(0x94, 0x00000200) // table at BAR0+0x200 (inside 4KB)
	cs.WriteU32(0x98, 0x00000280) // PBA at BAR0+0x280

	scrubbed := ScrubConfigSpace(cs, nil)

	// MSI-X should remain enabled
	msgCtl := scrubbed.ReadU16(0x92)
	if msgCtl&0x8000 == 0 {
		t.Errorf("MSI-X should remain enabled, got 0x%04x", msgCtl)
	}

	// table still relocated to 0x1000 (consistent placement)
	tableReg := scrubbed.ReadU32(0x94)
	tableOff := tableReg &^ 0x07
	if tableOff != 0x1000 {
		t.Errorf("MSI-X table should be relocated to 0x1000, got 0x%X", tableOff)
	}
}
