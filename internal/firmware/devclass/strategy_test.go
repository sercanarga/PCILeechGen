package devclass

import (
	"testing"
)

func TestStrategyForClass_NVMe(t *testing.T) {
	s := StrategyForClass(0x010802)
	if s == nil {
		t.Fatal("expected NVMe strategy, got nil")
	}
	if s.ClassName() != "NVMe" {
		t.Errorf("expected NVMe, got %s", s.ClassName())
	}
	if !s.IsNVMe() {
		t.Error("expected IsNVMe=true")
	}
	if s.IsXHCI() {
		t.Error("expected IsXHCI=false")
	}
	if s.Profile() == nil {
		t.Error("expected non-nil profile")
	}
}

func TestStrategyForClass_xHCI(t *testing.T) {
	s := StrategyForClass(0x0C0330)
	if s == nil {
		t.Fatal("expected xHCI strategy, got nil")
	}
	if s.ClassName() != "xHCI" {
		t.Errorf("expected xHCI, got %s", s.ClassName())
	}
	if s.IsNVMe() {
		t.Error("expected IsNVMe=false")
	}
	if !s.IsXHCI() {
		t.Error("expected IsXHCI=true")
	}
}

func TestStrategyForClass_Ethernet(t *testing.T) {
	s := StrategyForClass(0x020000)
	if s == nil {
		t.Fatal("expected Ethernet strategy, got nil")
	}
	if s.ClassName() != "Ethernet" {
		t.Errorf("expected Ethernet, got %s", s.ClassName())
	}
	if s.IsNVMe() || s.IsXHCI() {
		t.Error("Ethernet should not be NVMe or xHCI")
	}
}

func TestStrategyForClass_Audio(t *testing.T) {
	s := StrategyForClass(0x040300)
	if s == nil {
		t.Fatal("expected HD Audio strategy, got nil")
	}
	if s.ClassName() != "HD Audio" {
		t.Errorf("expected HD Audio, got %s", s.ClassName())
	}
}

func TestStrategyForClass_Unknown(t *testing.T) {
	s := StrategyForClass(0xFF0000)
	if s != nil {
		t.Errorf("expected nil for unknown class, got %v", s.ClassName())
	}
}

func TestNVMeStrategy_ScrubBAR(t *testing.T) {
	s := &nvmeStrategy{}
	data := make([]byte, 0x38)

	// set CC.EN=0, CSTS with dirty bits
	data[0x14] = 0x00 // CC.EN=0
	data[0x1C] = 0x1E // CSTS = CFS|SHST|NSSRO, RDY=0

	// fill queue config with garbage
	for i := 0x24; i < 0x38; i++ {
		data[i] = 0xFF
	}

	s.ScrubBAR(data)

	// CC.EN should be 1
	if data[0x14]&0x01 != 0x01 {
		t.Error("CC.EN should be set after scrub")
	}
	// CSTS.RDY should be 1, bits 4:1 cleared
	if data[0x1C] != 0x01 {
		t.Errorf("CSTS should be 0x01 after scrub, got 0x%02X", data[0x1C])
	}
	// queue config should be zeroed
	for i := 0x24; i < 0x38; i++ {
		if data[i] != 0 {
			t.Errorf("offset 0x%02X should be 0, got 0x%02X", i, data[i])
		}
	}
}

func TestNVMeStrategy_ScrubBAR_TooShort(t *testing.T) {
	s := &nvmeStrategy{}
	data := make([]byte, 0x10) // too short, should not panic
	s.ScrubBAR(data)
}

func TestXHCIStrategy_ScrubBAR_NoOp(t *testing.T) {
	s := &xhciStrategy{}
	data := []byte{0xAA, 0xBB}
	s.ScrubBAR(data)
	if data[0] != 0xAA || data[1] != 0xBB {
		t.Error("xHCI ScrubBAR should be a no-op")
	}
}

func TestNVMeStrategy_PostInitRegisters(t *testing.T) {
	s := &nvmeStrategy{}
	var csts uint32 = 0x0C // SHST bits set
	regs := map[uint32]*uint32{0x1C: &csts}
	s.PostInitRegisters(regs)
	if csts&0x01 == 0 {
		t.Error("CSTS.RDY should be set")
	}
	if csts&0x0C != 0 {
		t.Error("CSTS.SHST bits should be cleared")
	}
}

func TestXHCIStrategy_PostInitRegisters(t *testing.T) {
	s := &xhciStrategy{}
	var usbcmd uint32 = 0x00
	var usbsts uint32 = 0x01 // HCH=1
	regs := map[uint32]*uint32{0x20: &usbcmd, 0x24: &usbsts}
	s.PostInitRegisters(regs)
	if usbcmd&0x01 == 0 {
		t.Error("USBCMD R/S should be set")
	}
	if usbsts&0x01 != 0 {
		t.Error("USBSTS HCH should be cleared")
	}
}

func TestEthernetStrategy_PostInitRegisters(t *testing.T) {
	s := &ethernetStrategy{}
	var status uint32 = 0x00
	regs := map[uint32]*uint32{0x08: &status}
	s.PostInitRegisters(regs)
	if status&0x02 == 0 {
		t.Error("STATUS.LU should be set")
	}
}

func TestAudioStrategy_PostInitRegisters(t *testing.T) {
	s := &audioStrategy{}
	regs := map[uint32]*uint32{}
	s.PostInitRegisters(regs) // should not panic
}
