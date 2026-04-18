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

// TestFullPipeline_NoBARs_NoCaps simulates a worst-case donor like Creative
// Labs CA0132: no memory BARs, no capability chain. The scrubber must create
// BAR0 and a full PM+MSI+PCIe cap chain from scratch.
func TestFullPipeline_NoBARs_NoCaps(t *testing.T) {
	cs := makeCSNoCaps()
	// BAR0 = 0 (no BAR at all)
	b := makeBoard(2, 1)

	scrubbed := ScrubConfigSpace(cs, b)

	// BAR0 must be a 4KB 32-bit memory BAR
	bar0 := scrubbed.ReadU32(0x10)
	if bar0 == 0 {
		t.Fatal("BAR0 should not be 0 after scrub (scrubber must create one)")
	}
	if bar0&0x01 != 0 {
		t.Fatal("BAR0 should be memory type, not IO")
	}
	if bar0&0x06 != 0 {
		t.Fatal("BAR0 should be 32-bit (type bits [2:1] = 00)")
	}

	// PCIe cap must exist
	caps := pci.ParseCapabilities(scrubbed)
	hasPCIe := false
	for _, c := range caps {
		if c.ID == pci.CapIDPCIExpress {
			hasPCIe = true
		}
	}
	if !hasPCIe {
		t.Fatal("PCIe capability not found after scrub (inject should create it)")
	}

	// cap chain must be valid
	if err := ValidateCapChain(scrubbed); err != nil {
		t.Fatalf("cap chain invalid: %v", err)
	}
}

// TestFullPipeline_IOOnlyBAR simulates a donor with only I/O BARs.
// The scrubber must create a memory BAR0 for the FPGA.
func TestFullPipeline_IOOnlyBAR(t *testing.T) {
	cs := makeCSWithPMAndMSI()
	// BAR0 = IO BAR (bit 0 = 1)
	binary.LittleEndian.PutUint32(cs.Data[0x10:], 0xFFFFFF01)
	b := makeBoard(2, 1)

	scrubbed := ScrubConfigSpace(cs, b)

	// BAR0 must be a 4KB memory BAR (overwritten by scrubber)
	bar0 := scrubbed.ReadU32(0x10)
	if bar0&0x01 != 0 {
		t.Fatalf("BAR0 should be memory after scrub, got 0x%08X (IO bit set)", bar0)
	}
	if bar0 == 0 {
		t.Fatal("BAR0 should not be 0 after scrub")
	}
}

// verify injected PCIe cap fields are correct after full scrub pipeline.
func TestFullPipeline_PCIeCapDataIntegrity(t *testing.T) {
	cs := makeCSNoCaps()
	b := makeBoard(2, 1) // Gen2 x1

	scrubbed := ScrubConfigSpace(cs, b)

	caps := pci.ParseCapabilities(scrubbed)
	var pcieOff int
	found := false
	for _, c := range caps {
		if c.ID == pci.CapIDPCIExpress {
			pcieOff = c.Offset
			found = true
			break
		}
	}
	if !found {
		t.Fatal("PCIe cap not found after pipeline")
	}

	// cap ID
	if scrubbed.ReadU8(pcieOff) != pci.CapIDPCIExpress {
		t.Errorf("cap ID = 0x%02X, want 0x10", scrubbed.ReadU8(pcieOff))
	}

	// PCIe Capabilities Register (offset+2): version=2, device type=0 (endpoint)
	pcieCapReg := scrubbed.ReadU16(pcieOff + 2)
	version := pcieCapReg & 0x0F
	devType := (pcieCapReg >> 4) & 0x0F
	if version != 2 {
		t.Errorf("PCIe cap version = %d, want 2", version)
	}
	if devType != 0 {
		t.Errorf("PCIe device type = %d, want 0 (endpoint)", devType)
	}

	// Device Status (offset+10): must be cleared by scrubPCIeCapPass
	devStatus := scrubbed.ReadU16(pcieOff + 10)
	if devStatus != 0 {
		t.Errorf("Device Status = 0x%04X, want 0x0000 (cleared)", devStatus)
	}

	// Link Capabilities (offset+0x0C): speed and width
	linkCap := scrubbed.ReadU32(pcieOff + 0x0C)
	speed := linkCap & 0x0F
	width := (linkCap >> 4) & 0x3F
	if speed != 2 {
		t.Errorf("link speed = %d, want 2 (Gen2)", speed)
	}
	if width != 1 {
		t.Errorf("link width = %d, want 1 (x1)", width)
	}

	// ASPM must be cleared (bits 11:10 of Link Capabilities)
	aspmSupport := (linkCap >> 10) & 0x03
	if aspmSupport != 0 {
		t.Errorf("ASPM support = %d, want 0 (disabled)", aspmSupport)
	}

	// Clock PM must be cleared (bit 18 of Link Capabilities)
	clockPM := (linkCap >> 18) & 0x01
	if clockPM != 0 {
		t.Errorf("Clock PM = %d, want 0 (disabled)", clockPM)
	}
}

// PM capability must be D0 + NoSoftReset after scrub.
func TestFullPipeline_PMCapState(t *testing.T) {
	cs := makeCSWithPMAndMSI()
	// set PM to D3hot
	binary.LittleEndian.PutUint16(cs.Data[0x54:], 0x0003) // PMCSR = D3
	b := makeBoard(2, 1)

	scrubbed := ScrubConfigSpace(cs, b)

	caps := pci.ParseCapabilities(scrubbed)
	for _, c := range caps {
		if c.ID != pci.CapIDPowerManagement {
			continue
		}
		pmcsr := scrubbed.ReadU16(c.Offset + 4)
		powerState := pmcsr & 0x03
		noSoftReset := (pmcsr >> 3) & 0x01
		pmeStatus := (pmcsr >> 15) & 0x01

		if powerState != 0 {
			t.Errorf("PM power state = D%d, want D0", powerState)
		}
		if noSoftReset != 1 {
			t.Error("NoSoftReset should be set")
		}
		if pmeStatus != 0 {
			t.Error("PME_Status should be cleared")
		}
		return
	}
	t.Fatal("PM capability not found")
}

// command register must have BME + MSE set after scrub.
func TestFullPipeline_CommandRegister(t *testing.T) {
	cs := makeCSNoCaps()
	// command = 0 (nothing enabled)
	binary.LittleEndian.PutUint16(cs.Data[0x04:], 0x0000)
	b := makeBoard(2, 1)

	scrubbed := ScrubConfigSpace(cs, b)

	cmd := scrubbed.ReadU16(0x04)
	if cmd&0x02 == 0 {
		t.Error("Memory Space Enable (MSE) should be set")
	}
	if cmd&0x04 == 0 {
		t.Error("Bus Master Enable (BME) should be set")
	}
}

// 64-bit BAR must be forced to 32-bit by scrubber.
func TestFullPipeline_64bitBAR(t *testing.T) {
	cs := makeCSWithPMAndMSI()
	// BAR0 = 64-bit memory BAR (type bits [2:1] = 10)
	binary.LittleEndian.PutUint32(cs.Data[0x10:], 0xFFF00004) // 64-bit, 1MB
	binary.LittleEndian.PutUint32(cs.Data[0x14:], 0x00000001) // upper 32 bits
	b := makeBoard(2, 1)

	scrubbed := ScrubConfigSpace(cs, b)

	bar0 := scrubbed.ReadU32(0x10)
	// type bits [2:1] must be 00 (32-bit)
	if bar0&0x06 != 0 {
		t.Errorf("BAR0 type bits = %d, want 0 (32-bit), got 0x%08X", (bar0>>1)&0x03, bar0)
	}
	// BAR1 (upper 32 of old 64-bit) must be cleared
	bar1 := scrubbed.ReadU32(0x14)
	if bar1 != 0 {
		t.Errorf("BAR1 (upper) = 0x%08X, want 0 (cleared after 64->32 conversion)", bar1)
	}
	// BAR0 size must be 4KB
	sizeBits := bar0 & 0xFFFFF000
	if sizeBits != 0xFFFFF000 {
		t.Errorf("BAR0 size mask = 0x%08X, want 0xFFFFF000 (4KB)", sizeBits)
	}
}

// interrupt pin must be set to INTA# if donor has 0.
func TestFullPipeline_InterruptPin(t *testing.T) {
	cs := makeCSNoCaps()
	cs.Data[0x3D] = 0x00 // interrupt pin = 0
	b := makeBoard(2, 1)

	scrubbed := ScrubConfigSpace(cs, b)

	pin := scrubbed.ReadU8(0x3D)
	if pin != 0x01 {
		t.Errorf("interrupt pin = %d, want 1 (INTA#)", pin)
	}
}

// interrupt pin must be preserved if donor has valid value.
func TestFullPipeline_InterruptPinPreserved(t *testing.T) {
	cs := makeCSWithPMAndMSI()
	cs.Data[0x3D] = 0x02 // INTB#
	b := makeBoard(2, 1)

	scrubbed := ScrubConfigSpace(cs, b)

	pin := scrubbed.ReadU8(0x3D)
	if pin != 0x02 {
		t.Errorf("interrupt pin = %d, want 2 (INTB# preserved)", pin)
	}
}

// status register capabilities list bit must be set after injection.
func TestFullPipeline_StatusCapListBit(t *testing.T) {
	cs := makeCSNoCaps()
	// status has no capabilities list bit
	binary.LittleEndian.PutUint16(cs.Data[0x06:], 0x0000)
	b := makeBoard(2, 1)

	scrubbed := ScrubConfigSpace(cs, b)

	status := scrubbed.ReadU16(0x06)
	if status&0x0010 == 0 {
		t.Error("Status register Capabilities List bit should be set after injection")
	}
}

// CapPtr must point to a valid capability after full chain injection.
func TestFullPipeline_CapPtrValid(t *testing.T) {
	cs := makeCSNoCaps()
	b := makeBoard(2, 1)

	scrubbed := ScrubConfigSpace(cs, b)

	capPtr := scrubbed.ReadU8(0x34)
	if capPtr == 0 {
		t.Fatal("CapPtr should be non-zero after injection")
	}
	if capPtr < 0x40 {
		t.Errorf("CapPtr = 0x%02X, should be >= 0x40", capPtr)
	}
	if capPtr&0x03 != 0 {
		t.Errorf("CapPtr = 0x%02X, should be DWORD-aligned", capPtr)
	}
}

// full chain must have PM + MSI + PCIe when donor has no caps at all.
func TestFullPipeline_FullChainFromScratch(t *testing.T) {
	cs := makeCSNoCaps()
	b := makeBoard(3, 4) // Gen3 x4

	scrubbed := ScrubConfigSpace(cs, b)

	caps := pci.ParseCapabilities(scrubbed)

	hasPM, hasMSI, hasPCIe := false, false, false
	for _, c := range caps {
		switch c.ID {
		case pci.CapIDPowerManagement:
			hasPM = true
		case pci.CapIDMSI:
			hasMSI = true
		case pci.CapIDPCIExpress:
			hasPCIe = true
			// verify Gen3 x4
			linkCap := scrubbed.ReadU32(c.Offset + 0x0C)
			speed := linkCap & 0x0F
			width := (linkCap >> 4) & 0x3F
			if speed != 3 {
				t.Errorf("link speed = %d, want 3 (Gen3)", speed)
			}
			if width != 4 {
				t.Errorf("link width = %d, want 4 (x4)", width)
			}
		}
	}
	if !hasPM {
		t.Error("missing PM in injected chain")
	}
	if !hasMSI {
		t.Error("missing MSI in injected chain")
	}
	if !hasPCIe {
		t.Error("missing PCIe in injected chain")
	}

	if err := ValidateCapChain(scrubbed); err != nil {
		t.Fatalf("cap chain invalid: %v", err)
	}
}

// BAR0 must have correct 4KB size mask after scrub (for zero-BAR donor).
func TestFullPipeline_BAR0SizeMask(t *testing.T) {
	cs := makeCSNoCaps()
	b := makeBoard(2, 1)

	scrubbed := ScrubConfigSpace(cs, b)

	bar0 := scrubbed.ReadU32(0x10)
	expectedMask := uint32(0xFFFFF000)
	if bar0 != expectedMask {
		t.Errorf("BAR0 = 0x%08X, want 0x%08X (4KB size mask)", bar0, expectedMask)
	}
}

// BIST register must be cleared after scrub.
func TestFullPipeline_BISTCleared(t *testing.T) {
	cs := makeCSWithPMAndMSI()
	cs.Data[0x0F] = 0xFF // non-zero BIST
	b := makeBoard(2, 1)

	scrubbed := ScrubConfigSpace(cs, b)

	if scrubbed.ReadU8(0x0F) != 0 {
		t.Error("BIST register should be cleared")
	}
}

// latency timer and cache line size must be cleared.
func TestFullPipeline_LatencyCacheCleared(t *testing.T) {
	cs := makeCSWithPMAndMSI()
	cs.Data[0x0C] = 0x10 // cache line size
	cs.Data[0x0D] = 0x40 // latency timer
	b := makeBoard(2, 1)

	scrubbed := ScrubConfigSpace(cs, b)

	if scrubbed.ReadU8(0x0C) != 0 {
		t.Error("Cache Line Size should be cleared")
	}
	if scrubbed.ReadU8(0x0D) != 0 {
		t.Error("Latency Timer should be cleared")
	}
}

// BAR1-BAR5 must remain 0 if donor has no BARs.
func TestFullPipeline_OnlyBAR0Created(t *testing.T) {
	cs := makeCSNoCaps()
	b := makeBoard(2, 1)

	scrubbed := ScrubConfigSpace(cs, b)

	for i := 1; i < 6; i++ {
		barVal := scrubbed.ReadU32(0x10 + i*4)
		if barVal != 0 {
			t.Errorf("BAR%d = 0x%08X, want 0 (only BAR0 should be created)", i, barVal)
		}
	}
}

// VID/DID must be preserved through entire pipeline.
func TestFullPipeline_IdentityPreserved(t *testing.T) {
	cs := makeCSWithPMAndMSI()
	b := makeBoard(2, 1)

	scrubbed := ScrubConfigSpace(cs, b)

	if scrubbed.VendorID() != 0x1102 {
		t.Errorf("VID = 0x%04X, want 0x1102", scrubbed.VendorID())
	}
	if scrubbed.DeviceID() != 0x0012 {
		t.Errorf("DID = 0x%04X, want 0x0012", scrubbed.DeviceID())
	}
}

// donor with BAR0=memory and BAR2=IO: BAR0 stays memory, BAR2 gets clamped.
func TestFullPipeline_MixedBARTypes(t *testing.T) {
	cs := makeCSWithPMAndMSI()
	binary.LittleEndian.PutUint32(cs.Data[0x10:], 0xFFF00000) // BAR0: 1MB memory
	binary.LittleEndian.PutUint32(cs.Data[0x18:], 0x0000FF01) // BAR2: IO
	b := makeBoard(2, 1)

	scrubbed := ScrubConfigSpace(cs, b)

	bar0 := scrubbed.ReadU32(0x10)
	if bar0&0x01 != 0 {
		t.Error("BAR0 should be memory type")
	}
	if bar0&0xFFFFF000 != 0xFFFFF000 {
		t.Errorf("BAR0 size mask = 0x%08X, want 4KB", bar0&0xFFFFF000)
	}
}
