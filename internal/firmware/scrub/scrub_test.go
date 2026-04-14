package scrub

import (
	"testing"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/firmware"
	"github.com/sercanarga/pcileechgen/internal/firmware/overlay"
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
	if scrubbed.InterruptLine() != 0x0B {
		t.Errorf("InterruptLine should be preserved (not cleared): got 0x%02x, want 0x0B", scrubbed.InterruptLine())
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
	if bar0 != 0xFFFFF000 {
		t.Errorf("BAR0 should be clamped to 4 KB (type=32-bit): got 0x%08x, want 0xFFFFF000", bar0)
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
	// donor only had Gen1, board supports Gen2 - vector should remain Gen1 only
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

	cs.WriteU16(0x00, 0x1234) // unknown vendor - no whitelist
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

func TestZeroVendorRegisters_WhitelistPreserves(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize

	cs.WriteU16(0x00, 0x8086) // Intel - 0xE0-0xFF whitelisted
	cs.WriteU16(0x06, 0x0010)
	cs.WriteU8(0x34, 0x40)
	cs.WriteU8(0x40, pci.CapIDPCIExpress)
	cs.WriteU8(0x41, 0x00)

	cs.WriteU32(0xE4, 0xDEADBEEF) // in Intel whitelist range
	cs.WriteU32(0xB0, 0xCAFEBABE) // outside any cap or whitelist

	om := overlay.NewMap(cs)
	zeroVendorRegisters(cs, om, pci.ParseCapabilities(cs))

	if cs.ReadU32(0xE4) != 0xDEADBEEF {
		t.Errorf("Intel whitelisted register at 0xE4 should be preserved, got 0x%08x", cs.ReadU32(0xE4))
	}
	if cs.ReadU32(0xB0) != 0 {
		t.Errorf("non-whitelisted register at 0xB0 should be zeroed, got 0x%08x", cs.ReadU32(0xB0))
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
		{"32bit_masking", 0x0100, 18},
		{"64bit_masking", 0x0180, 22},
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

func makeTestCS() *pci.ConfigSpace {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	cs.WriteU16(0x00, 0x8086) // vendor
	cs.WriteU16(0x02, 0x1533) // device
	cs.WriteU16(0x06, 0x0010) // status: cap list
	cs.WriteU8(0x34, 0x40)    // cap pointer

	// PCIe cap at 0x40
	cs.WriteU8(0x40, pci.CapIDPCIExpress)
	cs.WriteU8(0x41, 0x00) // no next
	// DevCap at 0x44
	cs.WriteU32(0x44, 0x000128FF) // MPS=128, phantoms, ext tag
	// DevCtl at 0x48
	cs.WriteU16(0x48, 0x3BEF) // various bits set
	// LinkCap at 0x4C: Gen3, x4
	cs.WriteU32(0x4C, 0x00000043) // speed=3, width=4
	// LinkStatus at 0x52: Gen3, x4
	cs.WriteU16(0x52, 0x0043)
	// LinkCtl2 at 0x70: target speed Gen3
	cs.WriteU16(0x70, 0x0003)
	// LinkCap2 at 0x6C: speed vector
	cs.WriteU32(0x6C, 0x0000000E) // Gen1,2,3
	// DevCap2 at 0x64
	cs.WriteU32(0x64, 0x00030000) // 10-bit tags

	return cs
}

func TestClampLinkCapability(t *testing.T) {
	cs := makeTestCS()
	b := &board.Board{PCIeLanes: 1, MaxLinkSpeed: 2} // Gen2 x1
	om := overlay.NewMap(cs)

	clampLinkCapability(cs, b, om, pci.ParseCapabilities(cs))

	// Link Capabilities speed should be 2 (Gen2)
	linkCap := cs.ReadU32(0x4C)
	speed := uint8(linkCap & 0x0F)
	if speed != 2 {
		t.Errorf("Link speed = %d, want 2 (Gen2)", speed)
	}
	width := uint8((linkCap >> 4) & 0x3F)
	if width != 1 {
		t.Errorf("Link width = %d, want 1", width)
	}

	// Link Status
	ls := cs.ReadU16(0x52)
	lsSpeed := uint8(ls & 0x0F)
	if lsSpeed != 2 {
		t.Errorf("Link Status speed = %d, want 2", lsSpeed)
	}

	// LinkCtl2 target speed
	lc2 := cs.ReadU16(0x70)
	if uint8(lc2&0x0F) != 2 {
		t.Errorf("LinkCtl2 speed = %d, want 2", lc2&0x0F)
	}
}

func TestClampLinkCapability_NilBoard(t *testing.T) {
	cs := makeTestCS()
	om := overlay.NewMap(cs)
	clampLinkCapability(cs, nil, om, pci.ParseCapabilities(cs)) // should default to Gen2
	linkCap := cs.ReadU32(0x4C)
	speed := uint8(linkCap & 0x0F)
	if speed != 2 {
		t.Errorf("Nil board link speed = %d, want 2 (Gen2 default)", speed)
	}
}

func TestClampDeviceCapability(t *testing.T) {
	cs := makeTestCS()
	om := overlay.NewMap(cs)

	clampDeviceCapability(cs, om, pci.ParseCapabilities(cs))

	devCap := cs.ReadU32(0x44)
	if devCap&0x07 != 0 {
		t.Error("MPS should be cleared to 128B")
	}
	if devCap&0x18 != 0 {
		t.Error("Phantom functions should be disabled")
	}
	if devCap&0x20 != 0 {
		t.Error("Extended tag should be disabled")
	}

	devCap2 := cs.ReadU32(0x64)
	if devCap2&(1<<16) != 0 {
		t.Error("10-bit tag completer should be off")
	}
	if devCap2&(1<<17) != 0 {
		t.Error("10-bit tag requester should be off")
	}
}

func TestClampBARsToFPGA(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	// BAR0: memory, 1MB
	cs.WriteU32(0x10, 0xFFF00004) // 64-bit memory BAR
	cs.WriteU32(0x14, 0x00000001) // upper 32 bits
	// BAR2: IO BAR
	cs.WriteU32(0x18, 0x0000FF01) // IO bar

	om := overlay.NewMap(cs)
	clampBARsToFPGA(cs, om)

	// BAR0 should be clamped to 4KB
	bar0 := cs.ReadU32(0x10)
	if bar0&bar0SizeMask != bar0SizeMask {
		t.Errorf("BAR0 should be clamped to 4KB, got 0x%08x", bar0)
	}
	// BAR1 (upper 32 bits) should be zeroed
	if cs.ReadU32(0x14) != 0 {
		t.Errorf("BAR1 upper bits should be zeroed, got 0x%08x", cs.ReadU32(0x14))
	}
	// BAR2 (IO) should be clamped to 256 bytes
	ioBar := cs.ReadU32(0x18)
	if ioBar != 0xFFFFFF01 {
		t.Errorf("IO BAR should be clamped to 256 bytes: got 0x%08x, want 0xFFFFFF01", ioBar)
	}
}

func TestFilterExtCapabilities(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize

	// Write 3 ext caps: DSN (safe) -> SR-IOV (unsafe) -> AER (safe)
	// DSN at 0x100
	cs.WriteU32(0x100, uint32(pci.ExtCapIDDeviceSerialNumber)|(1<<16)|(0x110<<20))
	cs.WriteU32(0x104, 0x11223344)
	cs.WriteU32(0x108, 0x55667788)
	// SR-IOV at 0x110
	cs.WriteU32(0x110, uint32(pci.ExtCapIDSRIOV)|(1<<16)|(0x150<<20))
	// AER at 0x150
	cs.WriteU32(0x150, uint32(pci.ExtCapIDAER)|(1<<16))

	removed := FilterExtCapabilities(cs)

	if len(removed) != 1 {
		t.Fatalf("Expected 1 removed cap, got %d: %v", len(removed), removed)
	}
	if !contains(removed[0], "SR-IOV") {
		t.Errorf("Expected SR-IOV to be removed, got %q", removed[0])
	}
}

func TestFilterExtCapabilities_NoCaps(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	removed := FilterExtCapabilities(cs)
	if removed != nil {
		t.Errorf("Expected nil, got %v", removed)
	}
}

func TestFilterExtCapabilities_AllSafe(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	cs.WriteU32(0x100, uint32(pci.ExtCapIDAER)|(1<<16))
	removed := FilterExtCapabilities(cs)
	if removed != nil {
		t.Errorf("Expected nil (all safe), got %v", removed)
	}
}

func TestIsUnsafeExtCap_Known(t *testing.T) {
	if !IsUnsafeExtCap(pci.ExtCapIDSRIOV) {
		t.Error("SR-IOV should be unsafe")
	}
	if IsUnsafeExtCap(pci.ExtCapIDAER) {
		t.Error("AER should be safe")
	}
}

func TestUnsafeExtCapName(t *testing.T) {
	if name := UnsafeExtCapName(pci.ExtCapIDSRIOV); name != "SR-IOV" {
		t.Errorf("SR-IOV name = %q", name)
	}
	if name := UnsafeExtCapName(0xFFFF); name != "" {
		t.Errorf("Unknown cap name = %q, want empty", name)
	}
}

func TestRelocateMSIXToBRAM(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	cs.WriteU16(0x06, 0x0010) // status: cap list
	cs.WriteU8(0x34, 0x80)    // cap pointer

	// MSI-X at 0x80
	cs.WriteU8(0x80, pci.CapIDMSIX)
	cs.WriteU8(0x81, 0x00)
	cs.WriteU16(0x82, 0x0003)     // 4 vectors, function unmasked
	cs.WriteU32(0x84, 0x00000000) // table: BAR0, offset 0
	cs.WriteU32(0x88, 0x00000040) // PBA: BAR0, offset 0x40

	om := overlay.NewMap(cs)
	relocateMSIXToBRAM(cs, om, pci.ParseCapabilities(cs))

	// Table should be relocated to 0x1000
	tableReg := cs.ReadU32(0x84)
	tableOff := tableReg & 0xFFFFFFF8
	if tableOff != 0x1000 {
		t.Errorf("Table offset = 0x%x, want 0x1000", tableOff)
	}

	// MSI-X control should be enabled
	msgCtl := cs.ReadU16(0x82)
	if msgCtl&0x8000 == 0 {
		t.Error("MSI-X should be enabled")
	}
}

func TestZeroVendorRegisters_ClearsUncovered(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	cs.WriteU16(0x06, 0x0010) // has caps
	cs.WriteU8(0x34, 0x40)    // cap at 0x40
	cs.WriteU8(0x40, pci.CapIDPCIExpress)
	cs.WriteU8(0x41, 0x00)

	// Write junk to uncovered vendor area
	cs.WriteU8(0x90, 0xFF)
	cs.WriteU8(0x91, 0xAA)

	om := overlay.NewMap(cs)
	zeroVendorRegisters(cs, om, pci.ParseCapabilities(cs))

	if cs.ReadU8(0x90) != 0 {
		t.Error("Uncovered vendor register at 0x90 should be zeroed")
	}
	if cs.ReadU8(0x91) != 0 {
		t.Error("Uncovered vendor register at 0x91 should be zeroed")
	}
}

func TestSecondaryPCIeNotFiltered(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize

	// AER -> SecondaryPCIe -> LTR
	cs.WriteU32(0x100, makeExtCapHeader(pci.ExtCapIDAER, 1, 0x150))
	cs.WriteU32(0x150, makeExtCapHeader(pci.ExtCapIDSecondaryPCIe, 1, 0x200))
	cs.WriteU32(0x200, makeExtCapHeader(pci.ExtCapIDLTR, 1, 0))

	removed := FilterExtCapabilities(cs)
	if len(removed) != 0 {
		t.Errorf("SecondaryPCIe should NOT be filtered, but removed: %v", removed)
	}

	// verify SecondaryPCIe still in chain
	caps := pci.ParseExtCapabilities(cs)
	found := false
	for _, c := range caps {
		if c.ID == pci.ExtCapIDSecondaryPCIe {
			found = true
			break
		}
	}
	if !found {
		t.Error("SecondaryPCIe should remain in the ext cap chain")
	}
}

func TestSamsungVendorWhitelist(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize

	cs.WriteU16(0x00, 0x144D) // Samsung
	cs.WriteU16(0x06, 0x0010) // caps present
	cs.WriteU8(0x34, 0x40)    // cap ptr
	cs.WriteU8(0x40, pci.CapIDPowerManagement)
	cs.WriteU8(0x41, 0x00)
	cs.WriteU16(0x42, 0xC9C3)
	cs.WriteU16(0x44, 0x0008)

	// write vendor data in Samsung whitelist range (0x40-0x50)
	cs.WriteU32(0x48, 0xDEADBEEF)
	cs.WriteU32(0x4C, 0xCAFEBABE)
	// write vendor data outside whitelist
	cs.WriteU32(0xB0, 0x12345678)

	scrubbed := ScrubConfigSpace(cs, nil)

	// Samsung whitelist range should be preserved
	if scrubbed.ReadU32(0x48) != 0xDEADBEEF {
		t.Errorf("Samsung vendor region at 0x48 should be preserved, got 0x%08x", scrubbed.ReadU32(0x48))
	}
	if scrubbed.ReadU32(0x4C) != 0xCAFEBABE {
		t.Errorf("Samsung vendor region at 0x4C should be preserved, got 0x%08x", scrubbed.ReadU32(0x4C))
	}
	// outside whitelist should be zeroed
	if scrubbed.ReadU32(0xB0) != 0 {
		t.Errorf("non-whitelisted register at 0xB0 should be zeroed, got 0x%08x", scrubbed.ReadU32(0xB0))
	}
}

func TestL1PMSubstatesNotFiltered(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize

	// AER -> L1PM -> LTR
	cs.WriteU32(0x100, makeExtCapHeader(pci.ExtCapIDAER, 1, 0x150))
	cs.WriteU32(0x150, makeExtCapHeader(pci.ExtCapIDL1PMSubstates, 1, 0x200))
	cs.WriteU32(0x200, makeExtCapHeader(pci.ExtCapIDLTR, 1, 0))

	removed := FilterExtCapabilities(cs)
	if len(removed) != 0 {
		t.Errorf("L1PM Substates should NOT be filtered (handled by scrubASPMPass), but removed: %v", removed)
	}

	caps := pci.ParseExtCapabilities(cs)
	found := false
	for _, c := range caps {
		if c.ID == pci.ExtCapIDL1PMSubstates {
			found = true
			break
		}
	}
	if !found {
		t.Error("L1PM Substates should remain in the ext cap chain")
	}
}

func TestExtCapFilterReasons(t *testing.T) {
	// every entry in unsafeExtCaps must have both Name and Reason
	for id, f := range unsafeExtCaps {
		if f.Name == "" {
			t.Errorf("ext cap 0x%04x has empty Name", id)
		}
		if f.Reason == "" {
			t.Errorf("ext cap 0x%04x (%s) has empty Reason", id, f.Name)
		}
	}
	// verify convenience accessors
	if UnsafeExtCapReason(pci.ExtCapIDSRIOV) == "" {
		t.Error("SR-IOV should have a non-empty reason")
	}
	if UnsafeExtCapReason(0xFFFF) != "" {
		t.Error("unknown cap should return empty reason")
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > 0 && containsSubstring(s, substr))
}

func containsSubstring(s, sub string) bool {
	for i := 0; i+len(sub) <= len(s); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}

// TestASPMFullyDisabledAfterScrub verifies that the full pipeline
// clears all ASPM/Clock PM/L1PM bits and the writemask blocks re-enablement.
func TestASPMFullyDisabledAfterScrub(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	cs.WriteU16(0x06, 0x0010) // caps present
	cs.WriteU8(0x34, 0x40)

	// PCIe capability at 0x40
	cs.WriteU8(0x40, pci.CapIDPCIExpress)
	cs.WriteU8(0x41, 0x00)

	// Link Capabilities at cap+0x0C = 0x4C
	// set ASPM support (bits 11:10 = 11 = L0s+L1) and Clock PM (bit 18)
	linkCap := uint32(1) | (uint32(1) << 4) | (0x0C00) | (1 << 18)
	cs.WriteU32(0x4C, linkCap)

	// Link Control at cap+0x10 = 0x50
	// set ASPM enable (bits 1:0 = 11) and Clock PM enable (bit 8)
	cs.WriteU16(0x50, 0x0103)

	// Device Control 2 at cap+0x28 = 0x68
	// set LTR Mechanism Enable (bit 10)
	cs.WriteU16(0x68, 0x0400)

	// L1PM Substates ext cap at 0x200
	cs.WriteU32(0x100, makeExtCapHeader(pci.ExtCapIDAER, 1, 0x200))
	cs.WriteU32(0x200, makeExtCapHeader(pci.ExtCapIDL1PMSubstates, 1, 0))
	cs.WriteU32(0x204, 0x0000001F) // L1PM Capabilities: L1.1 + L1.2 support
	cs.WriteU32(0x208, 0x0000000A) // L1PM Control 1: L1.1 + L1.2 enabled
	cs.WriteU32(0x20C, 0x00000032) // L1PM Control 2

	b := &board.Board{PCIeLanes: 1}
	scrubbed := ScrubConfigSpace(cs, b)

	// check Link Capabilities: ASPM support and Clock PM must be cleared
	scrubbedLinkCap := scrubbed.ReadU32(0x4C)
	if scrubbedLinkCap&0x0C00 != 0 {
		t.Errorf("Link Cap ASPM Support bits should be 0, got 0x%08x", scrubbedLinkCap)
	}
	if scrubbedLinkCap&(1<<18) != 0 {
		t.Errorf("Link Cap Clock PM bit should be 0, got 0x%08x", scrubbedLinkCap)
	}

	// check Link Control: ASPM enable and Clock PM enable must be cleared
	scrubbedLinkCtl := scrubbed.ReadU16(0x50)
	if scrubbedLinkCtl&0x03 != 0 {
		t.Errorf("Link Ctl ASPM Enable bits should be 0, got 0x%04x", scrubbedLinkCtl)
	}
	if scrubbedLinkCtl&(1<<8) != 0 {
		t.Errorf("Link Ctl Clock PM Enable should be 0, got 0x%04x", scrubbedLinkCtl)
	}

	// check Device Control 2: LTR Mechanism Enable must be cleared
	scrubbedDevCtl2 := scrubbed.ReadU16(0x68)
	if scrubbedDevCtl2&(1<<10) != 0 {
		t.Errorf("DevCtl2 LTR Mechanism Enable should be 0, got 0x%04x", scrubbedDevCtl2)
	}

	// check L1PM Substates: caps and controls must be zeroed
	if scrubbed.ReadU32(0x204) != 0 {
		t.Errorf("L1PM Capabilities should be 0, got 0x%08x", scrubbed.ReadU32(0x204))
	}
	if scrubbed.ReadU32(0x208) != 0 {
		t.Errorf("L1PM Control 1 should be 0, got 0x%08x", scrubbed.ReadU32(0x208))
	}
	if scrubbed.ReadU32(0x20C) != 0 {
		t.Errorf("L1PM Control 2 should be 0, got 0x%08x", scrubbed.ReadU32(0x20C))
	}
}
