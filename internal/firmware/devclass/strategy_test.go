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
	if s.DeviceClass() != ClassNVMe {
		t.Errorf("expected %s, got %s", ClassNVMe, s.DeviceClass())
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
	if s.DeviceClass() != ClassXHCI {
		t.Errorf("expected %s, got %s", ClassXHCI, s.DeviceClass())
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
	if s.DeviceClass() != ClassEthernet {
		t.Errorf("expected %s, got %s", ClassEthernet, s.DeviceClass())
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
	if s.DeviceClass() != ClassAudio {
		t.Errorf("expected %s, got %s", ClassAudio, s.DeviceClass())
	}
}

func TestStrategyForClass_Unknown(t *testing.T) {
	s := StrategyForClass(0xFF0000)
	if s != nil {
		t.Errorf("expected nil for unknown class, got %v", s.ClassName())
	}
}

func TestDeviceClassConstants(t *testing.T) {
	if ClassNVMe != "nvme" {
		t.Errorf("ClassNVMe = %q, want nvme", ClassNVMe)
	}
	if ClassXHCI != "xhci" {
		t.Errorf("ClassXHCI = %q, want xhci", ClassXHCI)
	}
	if ClassEthernet != "ethernet" {
		t.Errorf("ClassEthernet = %q, want ethernet", ClassEthernet)
	}
	if ClassAudio != "audio" {
		t.Errorf("ClassAudio = %q, want audio", ClassAudio)
	}
}

func TestDeviceClassUniqueness(t *testing.T) {
	classes := map[string]bool{}
	for _, code := range []uint32{0x010802, 0x0C0330, 0x020000, 0x040300} {
		s := StrategyForClass(code)
		if s == nil {
			continue
		}
		dc := s.DeviceClass()
		if classes[dc] {
			t.Errorf("duplicate DeviceClass: %s", dc)
		}
		classes[dc] = true
	}
}

func TestNVMeStrategy_ScrubBAR(t *testing.T) {
	s := &nvmeStrategy{}
	data := make([]byte, 0x38)

	data[0x14] = 0x00
	data[0x1C] = 0x1E

	for i := 0x24; i < 0x38; i++ {
		data[i] = 0xFF
	}

	s.ScrubBAR(data)

	if data[0x14]&0x01 != 0x01 {
		t.Error("CC.EN should be set after scrub")
	}
	if data[0x1C] != 0x01 {
		t.Errorf("CSTS should be 0x01 after scrub, got 0x%02X", data[0x1C])
	}
	for i := 0x24; i < 0x38; i++ {
		if data[i] != 0 {
			t.Errorf("offset 0x%02X should be 0, got 0x%02X", i, data[i])
		}
	}
}

func TestNVMeStrategy_ScrubBAR_TooShort(t *testing.T) {
	s := &nvmeStrategy{}
	data := make([]byte, 0x10)
	s.ScrubBAR(data) // should not panic
}

func TestXHCIStrategy_ScrubBAR_NoOp(t *testing.T) {
	s := &xhciStrategy{}
	data := []byte{0xAA, 0xBB}
	s.ScrubBAR(data)
	if data[0] != 0xAA || data[1] != 0xBB {
		t.Error("xHCI ScrubBAR should be a no-op")
	}
}

func TestEthernetStrategy_ScrubBAR_NoOp(t *testing.T) {
	s := &ethernetStrategy{}
	data := []byte{0x12, 0x34}
	s.ScrubBAR(data)
	if data[0] != 0x12 {
		t.Error("Ethernet ScrubBAR should be a no-op")
	}
}

func TestAudioStrategy_ScrubBAR_NoOp(t *testing.T) {
	s := &audioStrategy{}
	data := []byte{0xAB}
	s.ScrubBAR(data)
	if data[0] != 0xAB {
		t.Error("Audio ScrubBAR should be a no-op")
	}
}

func TestNVMeStrategy_PostInitRegisters(t *testing.T) {
	s := &nvmeStrategy{}
	var csts uint32 = 0x0C
	regs := map[uint32]*uint32{0x1C: &csts}
	s.PostInitRegisters(regs)
	if csts&0x01 == 0 {
		t.Error("CSTS.RDY should be set")
	}
	if csts&0x0C != 0 {
		t.Error("CSTS.SHST bits should be cleared")
	}
}

func TestNVMeStrategy_PostInitRegisters_NoCSTS(t *testing.T) {
	s := &nvmeStrategy{}
	regs := map[uint32]*uint32{}
	s.PostInitRegisters(regs) // should not panic
}

func TestXHCIStrategy_PostInitRegisters(t *testing.T) {
	s := &xhciStrategy{}
	var usbcmd uint32 = 0x00
	var usbsts uint32 = 0x01
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
	var eecd uint32 = 0x00
	regs := map[uint32]*uint32{0x08: &status, 0x10: &eecd}
	s.PostInitRegisters(regs)
	if status&0x02 == 0 {
		t.Error("STATUS.LU should be set")
	}
	if status&0x80 == 0 {
		t.Error("STATUS speed bits should be set")
	}
	if eecd&0x200 == 0 {
		t.Error("EECD Auto-Read Done should be set")
	}
	if eecd&0x100 == 0 {
		t.Error("EECD EEPROM Present should be set")
	}
}

func TestAudioStrategy_PostInitRegisters(t *testing.T) {
	s := &audioStrategy{}
	var gctl uint32 = 0x00
	var wakeenStatests uint32 = 0x00
	regs := map[uint32]*uint32{0x08: &gctl, 0x0C: &wakeenStatests}
	s.PostInitRegisters(regs)
	if gctl&0x01 == 0 {
		t.Error("GCTL.CRST should be set")
	}
	if wakeenStatests&0x10000 == 0 {
		t.Error("STATESTS codec 0 present should be set (bit 16)")
	}
}

func TestAudioStrategy_PostInitRegisters_Empty(t *testing.T) {
	s := &audioStrategy{}
	regs := map[uint32]*uint32{}
	s.PostInitRegisters(regs) // should not panic
}

func TestAllStrategies_Profile(t *testing.T) {
	for _, code := range []uint32{0x010802, 0x0C0330, 0x020000, 0x040300} {
		s := StrategyForClass(code)
		if s == nil {
			t.Errorf("nil strategy for 0x%06X", code)
			continue
		}
		p := s.Profile()
		if p == nil {
			t.Errorf("nil profile for %s", s.ClassName())
		}
		if p.ClassName == "" {
			t.Errorf("nil profile className for %s", s.ClassName())
		}
	}
}
