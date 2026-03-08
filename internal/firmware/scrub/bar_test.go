package scrub

import (
	"testing"

	"github.com/sercanarga/pcileechgen/internal/util"
)

func TestScrubBarContent_NVMe(t *testing.T) {
	barData := make([]byte, 4096)
	barContents := map[int][]byte{0: barData}
	ScrubBarContent(barContents, 0x010802) // NVMe

	// CSTS.RDY should be 1
	csts := util.ReadLE32(barData, 0x1C)
	if csts&0x01 == 0 {
		t.Error("NVMe CSTS.RDY should be set")
	}
}

func TestScrubBarContent_XHCI(t *testing.T) {
	barData := make([]byte, 4096)
	barData[0x00] = 0x20 // CAPLENGTH = 0x20
	barData[0x02] = 0x00
	barData[0x03] = 0x01 // HCIVERSION = 0x0100
	// HCSPARAMS1: 32 slots, 8 intrs, 4 ports
	util.WriteLE32(barData, 0x04, 0x04000820)
	// DBOFF, RTSOFF
	util.WriteLE32(barData, 0x14, 0x00000400)
	util.WriteLE32(barData, 0x18, 0x00000600)

	barContents := map[int][]byte{0: barData}
	ScrubBarContent(barContents, 0x0C0330) // xHCI

	// USBCMD R/S should be set
	capLen := int(barData[0x00])
	usbcmd := util.ReadLE32(barData, capLen)
	if usbcmd&0x01 == 0 {
		t.Error("xHCI USBCMD R/S should be set")
	}
}

func TestScrubBarContent_NoData(t *testing.T) {
	ScrubBarContent(nil, 0x010802)
	ScrubBarContent(map[int][]byte{}, 0x010802)
}

func TestScrubBarContent_Unknown(t *testing.T) {
	barData := make([]byte, 256)
	barContents := map[int][]byte{0: barData}
	ScrubBarContent(barContents, 0xFF0000) // unknown — should be no-op
}

func TestScrubXHCIBar0_TooSmall(t *testing.T) {
	data := make([]byte, 10)
	scrubXHCIBar0(data) // should not panic
}

func TestXhciFixCapLength(t *testing.T) {
	data := make([]byte, 256)

	// Zero caplen → should be fixed to 0x20
	data[0x00] = 0x00
	cl := xhciFixCapLength(data)
	if cl != 0x20 {
		t.Errorf("caplen = %d, want 0x20", cl)
	}

	// > 0x40 → should be fixed to 0x20
	data[0x00] = 0x50
	cl = xhciFixCapLength(data)
	if cl != 0x20 {
		t.Errorf("caplen = %d, want 0x20", cl)
	}

	// Valid value
	data[0x00] = 0x30
	cl = xhciFixCapLength(data)
	if cl != 0x30 {
		t.Errorf("caplen = %d, want 0x30", cl)
	}
}

func TestXhciFixHCIVersion(t *testing.T) {
	data := make([]byte, 256)

	// Bad version
	data[0x02] = 0x00
	data[0x03] = 0x00
	xhciFixHCIVersion(data)
	hci := uint16(data[0x02]) | uint16(data[0x03])<<8
	if hci < 0x0100 {
		t.Errorf("HCIVERSION = 0x%04x, should be >= 0x0100", hci)
	}

	// Already fine
	data[0x02] = 0x10
	data[0x03] = 0x01 // 0x0110
	xhciFixHCIVersion(data)
	hci = uint16(data[0x02]) | uint16(data[0x03])<<8
	if hci != 0x0110 {
		t.Errorf("HCIVERSION should stay 0x0110, got 0x%04x", hci)
	}
}

func TestXhciReadStructParams(t *testing.T) {
	data := make([]byte, 256)
	util.WriteLE32(data, 0x04, 0x04000820) // 32 slots, 8 intrs, 4 ports
	slots, intrs, ports := xhciReadStructParams(data)
	if slots != 32 {
		t.Errorf("MaxSlots = %d, want 32", slots)
	}
	if intrs != 8 {
		t.Errorf("MaxIntrs = %d, want 8", intrs)
	}
	if ports != 4 {
		t.Errorf("MaxPorts = %d, want 4", ports)
	}
}

func TestXhciReadStructParams_ZeroSlots(t *testing.T) {
	data := make([]byte, 256)
	util.WriteLE32(data, 0x04, 0x04000800) // 0 slots → default 32
	slots, _, _ := xhciReadStructParams(data)
	if slots != 32 {
		t.Errorf("Zero slots should default to 32, got %d", slots)
	}
}

func TestXhciClampScratchpads(t *testing.T) {
	data := make([]byte, 256)
	// 0xFFE12345: hi[31:27]=0x1F, bit26=1 (SPR), lo[25:21]=0x0F, rest=0x12345
	util.WriteLE32(data, 0x08, 0xFFE12345)
	xhciClampScratchpads(data)
	hcsparams2 := util.ReadLE32(data, 0x08)
	// hi[31:27] and lo[25:21] should be zeroed, bit 26 (SPR) preserved
	if hcsparams2&0xFBE00000 != 0 {
		t.Errorf("Scratchpad count bits should be zeroed, got 0x%08x", hcsparams2)
	}
	// bit 26 and lower bits should be preserved
	if hcsparams2 != 0x04012345 {
		t.Errorf("Non-scratchpad bits should be preserved, got 0x%08x, want 0x04012345", hcsparams2)
	}
}

func TestXhciClampXECP(t *testing.T) {
	data := make([]byte, 256)
	// xECP pointing outside BRAM
	util.WriteLE32(data, 0x10, 0x04000005) // xECP = 0x400*4 = 0x1000 > BRAMSize
	xhciClampXECP(data)
	hccparams1 := util.ReadLE32(data, 0x10)
	if hccparams1&0xFFFF0000 != 0 {
		t.Errorf("xECP should be zeroed, got 0x%08x", hccparams1)
	}
}

func TestXhciClampPorts(t *testing.T) {
	data := make([]byte, 256)
	maxPorts := xhciClampPorts(data, 0x20, 100)
	if maxPorts > 100 {
		t.Errorf("Clamped ports = %d, should be <= 100", maxPorts)
	}
	if maxPorts < 1 {
		t.Errorf("Clamped ports = %d, should be >= 1", maxPorts)
	}
}

func TestXhciSetOperationalState(t *testing.T) {
	data := make([]byte, 256)
	capLen := 0x20
	xhciSetOperationalState(data, capLen, 16)

	// PAGESIZE should be 1 (4KB)
	pageSize := util.ReadLE32(data, capLen+0x08)
	if pageSize != 1 {
		t.Errorf("PAGESIZE = %d, want 1", pageSize)
	}

	// USBCMD R/S should be set
	usbcmd := util.ReadLE32(data, capLen)
	if usbcmd&0x01 == 0 {
		t.Error("USBCMD R/S should be set")
	}

	// USBCMD HCRST should be clear
	if usbcmd&0x02 != 0 {
		t.Error("USBCMD HCRST should be cleared")
	}
}
