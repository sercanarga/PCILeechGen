package scrub

import (
	"encoding/binary"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/firmware/overlay"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

// makeBoard returns a test board with given speed and lanes.
func makeBoard(speed uint8, lanes int) *board.Board {
	return &board.Board{
		Name:         "test",
		FPGAPart:     "xc7a35t",
		PCIeLanes:    lanes,
		MaxLinkSpeed: speed,
		TopModule:    "top",
	}
}

// makeCSWithPMAndMSI builds a config space with PM + MSI caps but no PCIe.
// This simulates a conventional PCI donor device.
func makeCSWithPMAndMSI() *pci.ConfigSpace {
	cs := pci.NewConfigSpace()
	// VID:DID
	binary.LittleEndian.PutUint16(cs.Data[0x00:], 0x1102) // Creative Labs
	binary.LittleEndian.PutUint16(cs.Data[0x02:], 0x0012) // some device
	// Status: capabilities list present (bit 4)
	binary.LittleEndian.PutUint16(cs.Data[0x06:], 0x0010)
	// Class code: multimedia / HDA / 0
	cs.Data[0x09] = 0x00 // ProgIF
	cs.Data[0x0A] = 0x03 // SubClass (HDA)
	cs.Data[0x0B] = 0x04 // BaseClass (Multimedia)
	// CapPtr -> 0x50
	cs.Data[0x34] = 0x50

	// PM capability at 0x50 (8 bytes)
	cs.Data[0x50] = pci.CapIDPowerManagement // cap ID
	cs.Data[0x51] = 0x60                     // next -> MSI at 0x60
	binary.LittleEndian.PutUint16(cs.Data[0x52:], 0xC803)
	binary.LittleEndian.PutUint16(cs.Data[0x54:], 0x0008)

	// MSI capability at 0x60 (10 bytes, 32-bit)
	cs.Data[0x60] = pci.CapIDMSI // cap ID
	cs.Data[0x61] = 0x00         // next = 0 (end of chain)
	binary.LittleEndian.PutUint16(cs.Data[0x62:], 0x0000)

	cs.Size = pci.ConfigSpaceSize
	return cs
}

// makeCSWithPCIe builds a config space that already has a PCIe capability.
func makeCSWithPCIe() *pci.ConfigSpace {
	cs := makeCSWithPMAndMSI()
	// extend MSI to point to PCIe
	cs.Data[0x61] = 0x70 // MSI next -> PCIe at 0x70
	// PCIe capability at 0x70
	cs.Data[0x70] = pci.CapIDPCIExpress
	cs.Data[0x71] = 0x00 // end of chain
	binary.LittleEndian.PutUint16(cs.Data[0x72:], 0x0002) // version 2, endpoint
	return cs
}

// makeCSNoCaps builds a config space with no capabilities at all.
func makeCSNoCaps() *pci.ConfigSpace {
	cs := pci.NewConfigSpace()
	binary.LittleEndian.PutUint16(cs.Data[0x00:], 0x1102)
	binary.LittleEndian.PutUint16(cs.Data[0x02:], 0x0012)
	// Status: no capabilities (bit 4 = 0)
	binary.LittleEndian.PutUint16(cs.Data[0x06:], 0x0000)
	cs.Data[0x09] = 0x00
	cs.Data[0x0A] = 0x03
	cs.Data[0x0B] = 0x04
	cs.Data[0x34] = 0x00 // no CapPtr
	cs.Size = pci.ConfigSpaceSize
	return cs
}

func TestInjectPCIeCap_NoPCIe(t *testing.T) {
	cs := makeCSWithPMAndMSI()
	b := makeBoard(2, 1) // Gen2 x1
	om := overlay.NewMap(cs)
	ctx := &ScrubContext{
		Caps: pci.ParseCapabilities(cs),
	}

	// verify no PCIe cap before injection
	if hasPCIeCap(ctx.Caps) {
		t.Fatal("expected no PCIe cap before injection")
	}

	injectPCIeCapIfMissing(cs, b, om, ctx)

	// verify PCIe cap now exists
	if !hasPCIeCap(ctx.Caps) {
		t.Fatal("expected PCIe cap after injection")
	}

	// find the injected PCIe cap
	var pcieCap *pci.Capability
	for i := range ctx.Caps {
		if ctx.Caps[i].ID == pci.CapIDPCIExpress {
			pcieCap = &ctx.Caps[i]
			break
		}
	}

	// verify it's at a valid DWORD-aligned offset
	if pcieCap.Offset < 0x40 || pcieCap.Offset >= pci.ConfigSpaceLegacySize {
		t.Fatalf("PCIe cap at invalid offset 0x%02X", pcieCap.Offset)
	}
	if pcieCap.Offset%4 != 0 {
		t.Fatalf("PCIe cap at non-DWORD-aligned offset 0x%02X", pcieCap.Offset)
	}

	// verify link capabilities match board config
	linkCap := binary.LittleEndian.Uint32(cs.Data[pcieCap.Offset+0x0C:])
	speed := linkCap & 0x0F
	width := (linkCap >> 4) & 0x3F
	if speed != 2 {
		t.Errorf("link speed = %d, want 2 (Gen2)", speed)
	}
	if width != 1 {
		t.Errorf("link width = %d, want 1 (x1)", width)
	}

	// verify link status matches
	linkStatus := binary.LittleEndian.Uint16(cs.Data[pcieCap.Offset+0x12:])
	if linkStatus&0x0F != 2 {
		t.Errorf("link status speed = %d, want 2", linkStatus&0x0F)
	}
	if (linkStatus>>4)&0x3F != 1 {
		t.Errorf("link status width = %d, want 1", (linkStatus>>4)&0x3F)
	}

	// verify PCIe cap version and device type
	pcieCapReg := binary.LittleEndian.Uint16(cs.Data[pcieCap.Offset+0x02:])
	version := pcieCapReg & 0x0F
	devType := (pcieCapReg >> 4) & 0x0F
	if version != 2 {
		t.Errorf("PCIe cap version = %d, want 2", version)
	}
	if devType != 0 {
		t.Errorf("PCIe device type = %d, want 0 (endpoint)", devType)
	}
}

func TestInjectPCIeCap_AlreadyHasPCIe(t *testing.T) {
	cs := makeCSWithPCIe()
	b := makeBoard(2, 1)
	om := overlay.NewMap(cs)
	ctx := &ScrubContext{
		Caps: pci.ParseCapabilities(cs),
	}

	// should already have PCIe
	if !hasPCIeCap(ctx.Caps) {
		t.Fatal("test setup error: expected PCIe cap")
	}
	capCountBefore := len(ctx.Caps)

	injectPCIeCapIfMissing(cs, b, om, ctx)

	// cap count should not change
	if len(ctx.Caps) != capCountBefore {
		t.Errorf("cap count changed from %d to %d, expected no change", capCountBefore, len(ctx.Caps))
	}

	// should still have exactly one PCIe cap
	count := 0
	for _, c := range ctx.Caps {
		if c.ID == pci.CapIDPCIExpress {
			count++
		}
	}
	if count != 1 {
		t.Errorf("found %d PCIe caps, want 1", count)
	}
}

func TestInjectPCIeCap_NoCapsAtAll(t *testing.T) {
	cs := makeCSNoCaps()
	b := makeBoard(2, 1)
	om := overlay.NewMap(cs)
	ctx := &ScrubContext{
		Caps: pci.ParseCapabilities(cs),
	}

	if len(ctx.Caps) != 0 {
		t.Fatal("test setup error: expected no caps")
	}

	injectPCIeCapIfMissing(cs, b, om, ctx)

	// should now have PM + MSI + PCIe
	if len(ctx.Caps) < 3 {
		t.Fatalf("expected at least 3 caps after full chain injection, got %d", len(ctx.Caps))
	}

	hasPM, hasMSI, hasPCIe := false, false, false
	for _, c := range ctx.Caps {
		switch c.ID {
		case pci.CapIDPowerManagement:
			hasPM = true
		case pci.CapIDMSI:
			hasMSI = true
		case pci.CapIDPCIExpress:
			hasPCIe = true
		}
	}
	if !hasPM {
		t.Error("missing PM capability after full chain injection")
	}
	if !hasMSI {
		t.Error("missing MSI capability after full chain injection")
	}
	if !hasPCIe {
		t.Error("missing PCIe capability after full chain injection")
	}

	// verify Status register has capabilities list bit set
	if cs.Status()&0x0010 == 0 {
		t.Error("Status register capabilities list bit not set")
	}

	// verify CapPtr is set
	if cs.CapabilityPointer() == 0 {
		t.Error("CapPtr is still 0 after injection")
	}
}

func TestInjectPCIeCap_BoardGen3x4(t *testing.T) {
	cs := makeCSWithPMAndMSI()
	b := makeBoard(3, 4) // Gen3 x4
	om := overlay.NewMap(cs)
	ctx := &ScrubContext{
		Caps: pci.ParseCapabilities(cs),
	}

	injectPCIeCapIfMissing(cs, b, om, ctx)

	var pcieCap *pci.Capability
	for i := range ctx.Caps {
		if ctx.Caps[i].ID == pci.CapIDPCIExpress {
			pcieCap = &ctx.Caps[i]
			break
		}
	}
	if pcieCap == nil {
		t.Fatal("PCIe cap not found after injection")
	}

	// link capabilities
	linkCap := binary.LittleEndian.Uint32(cs.Data[pcieCap.Offset+0x0C:])
	if linkCap&0x0F != 3 {
		t.Errorf("link cap speed = %d, want 3 (Gen3)", linkCap&0x0F)
	}
	if (linkCap>>4)&0x3F != 4 {
		t.Errorf("link cap width = %d, want 4 (x4)", (linkCap>>4)&0x3F)
	}

	// link capabilities 2 - speed vector
	linkCap2 := binary.LittleEndian.Uint32(cs.Data[pcieCap.Offset+0x2C:])
	// should have bits 1,2,3 set (Gen1, Gen2, Gen3)
	expectedVec := uint32(0x0E) // bits 1+2+3
	if linkCap2 != expectedVec {
		t.Errorf("link cap 2 speed vector = 0x%X, want 0x%X", linkCap2, expectedVec)
	}

	// link control 2 - target speed
	linkCtl2 := binary.LittleEndian.Uint16(cs.Data[pcieCap.Offset+0x30:])
	if linkCtl2&0x0F != 3 {
		t.Errorf("link ctl2 target speed = %d, want 3", linkCtl2&0x0F)
	}
}

func TestInjectPCIeCap_CapChainIntegrity(t *testing.T) {
	cs := makeCSWithPMAndMSI()
	b := makeBoard(2, 1)
	om := overlay.NewMap(cs)
	ctx := &ScrubContext{
		Caps: pci.ParseCapabilities(cs),
	}

	injectPCIeCapIfMissing(cs, b, om, ctx)

	// validate the capability chain has no loops or bad pointers
	if err := ValidateCapChain(cs); err != nil {
		t.Fatalf("cap chain validation failed after injection: %v", err)
	}
}

func TestFindFreeCapSpace_Full(t *testing.T) {
	cs := pci.NewConfigSpace()
	// fill 0x40-0xFF with caps (simulate full cap region)
	var fakeCaps []pci.Capability
	for off := 0x40; off < pci.ConfigSpaceLegacySize; off += 4 {
		cs.Data[off] = 0x09 // vendor specific
		cs.Data[off+1] = 0x00
		fakeCaps = append(fakeCaps, pci.Capability{ID: 0x09, Offset: off})
	}

	result := findFreeCapSpace(cs, fakeCaps, pcieCapSize)
	if result != -1 {
		t.Errorf("expected -1 (no space), got 0x%02X", result)
	}
}

func TestBuildPCIeCapData_StructSize(t *testing.T) {
	b := makeBoard(2, 1)
	data := buildPCIeCapData(b)

	if len(data) != pcieCapSize {
		t.Fatalf("cap data size = %d, want %d", len(data), pcieCapSize)
	}

	// verify cap ID
	if data[0] != pci.CapIDPCIExpress {
		t.Errorf("cap ID = 0x%02X, want 0x%02X", data[0], pci.CapIDPCIExpress)
	}

	// verify next pointer is 0 (end of chain)
	if data[1] != 0x00 {
		t.Errorf("next ptr = 0x%02X, want 0x00", data[1])
	}
}

func TestInjectPCIeCap_FullPipelineIntegration(t *testing.T) {
	cs := makeCSWithPMAndMSI()
	// set a BAR so the scrubber has something to clamp
	binary.LittleEndian.PutUint32(cs.Data[0x10:], 0xFFF00000) // BAR0: 1MB memory
	b := makeBoard(2, 1)

	// run full scrub pipeline
	scrubbed := ScrubConfigSpace(cs, b)

	// verify PCIe cap exists in the scrubbed output
	caps := pci.ParseCapabilities(scrubbed)
	found := false
	for _, c := range caps {
		if c.ID == pci.CapIDPCIExpress {
			found = true
			break
		}
	}
	if !found {
		t.Fatal("PCIe capability not found after full scrub pipeline")
	}

	// verify cap chain is valid
	if err := ValidateCapChain(scrubbed); err != nil {
		t.Fatalf("cap chain invalid after full pipeline: %v", err)
	}
}
