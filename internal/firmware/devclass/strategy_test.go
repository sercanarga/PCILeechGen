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

func TestStrategyForClass_GPU(t *testing.T) {
	s := StrategyForClass(0x030000)
	if s == nil {
		t.Fatal("expected GPU strategy, got nil")
	}
	if s.ClassName() != "GPU" {
		t.Errorf("expected GPU, got %s", s.ClassName())
	}
	if s.DeviceClass() != ClassGPU {
		t.Errorf("expected %s, got %s", ClassGPU, s.DeviceClass())
	}
}

func TestStrategyForClass_SATA(t *testing.T) {
	s := StrategyForClass(0x010601)
	if s == nil {
		t.Fatal("expected SATA strategy, got nil")
	}
	if s.ClassName() != "SATA AHCI" {
		t.Errorf("expected SATA AHCI, got %s", s.ClassName())
	}
	if s.DeviceClass() != ClassSATA {
		t.Errorf("expected %s, got %s", ClassSATA, s.DeviceClass())
	}
}

func TestStrategyForClass_WiFi(t *testing.T) {
	s := StrategyForClass(0x028000)
	if s == nil {
		t.Fatal("expected Wi-Fi strategy, got nil")
	}
	if s.ClassName() != "Wi-Fi" {
		t.Errorf("expected Wi-Fi, got %s", s.ClassName())
	}
	if s.DeviceClass() != ClassWiFi {
		t.Errorf("expected %s, got %s", ClassWiFi, s.DeviceClass())
	}
}

func TestStrategyForClass_Thunderbolt(t *testing.T) {
	s := StrategyForClass(0x0C8000)
	if s == nil {
		t.Fatal("expected Thunderbolt strategy, got nil")
	}
	if s.ClassName() != "Thunderbolt" {
		t.Errorf("expected Thunderbolt, got %s", s.ClassName())
	}
	if s.DeviceClass() != ClassThunderbolt {
		t.Errorf("expected %s, got %s", ClassThunderbolt, s.DeviceClass())
	}
}

func TestStrategyForClass_Generic(t *testing.T) {
	s := StrategyForClass(0xFF0000)
	if s == nil {
		t.Fatal("generic strategy should never be nil")
	}
	if s.ClassName() != "Generic" {
		t.Errorf("expected Generic, got %s", s.ClassName())
	}
	if s.DeviceClass() != ClassGeneric {
		t.Errorf("expected %s, got %s", ClassGeneric, s.DeviceClass())
	}
}

func TestDeviceClassConstants(t *testing.T) {
	expected := map[string]string{
		ClassNVMe:        "nvme",
		ClassXHCI:        "xhci",
		ClassEthernet:    "ethernet",
		ClassAudio:       "audio",
		ClassGPU:         "gpu",
		ClassSATA:        "sata",
		ClassWiFi:        "wifi",
		ClassThunderbolt: "thunderbolt",
		ClassGeneric:     "generic",
	}
	for got, want := range expected {
		if got != want {
			t.Errorf("constant %q != %q", got, want)
		}
	}
}

func TestDeviceClassUniqueness(t *testing.T) {
	codes := []uint32{
		0x010802, 0x0C0330, 0x020000, 0x040300,
		0x030000, 0x010601, 0x028000, 0x0C8000,
	}
	classes := map[string]bool{}
	for _, code := range codes {
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

// --- ScrubBAR tests ---

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
		t.Errorf("CSTS should be 0x01, got 0x%02X", data[0x1C])
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
	s.ScrubBAR(data) // must not panic
}

func TestXHCIStrategy_ScrubBAR(t *testing.T) {
	s := &xhciStrategy{}
	data := make([]byte, 0x30)
	// USBCMD=0, USBSTS=0x01 (halted), DBOFF=0xFFFF, RTSOFF=0xFFFF
	data[0x24] = 0x01
	data[0x14] = 0xFF
	data[0x15] = 0xFF
	data[0x18] = 0xFF
	data[0x19] = 0xFF

	s.ScrubBAR(data)

	if data[0x20]&0x01 != 0x01 {
		t.Error("USBCMD R/S should be set")
	}
	if data[0x24] != 0x00 {
		t.Errorf("USBSTS should be cleared, got 0x%02X", data[0x24])
	}
	// DBOFF should be clamped
	dboff := uint32(data[0x14]) | uint32(data[0x15])<<8
	if dboff > 0x800 {
		t.Errorf("DBOFF should be clamped, got 0x%04X", dboff)
	}
}

func TestXHCIStrategy_ScrubBAR_TooShort(t *testing.T) {
	s := &xhciStrategy{}
	data := make([]byte, 0x10)
	s.ScrubBAR(data) // must not panic
}

func TestEthernetStrategy_ScrubBAR(t *testing.T) {
	s := &ethernetStrategy{}
	data := make([]byte, 0x24)

	s.ScrubBAR(data)

	status := uint32(data[0x08]) | uint32(data[0x09])<<8 | uint32(data[0x0A])<<16 | uint32(data[0x0B])<<24
	if status != 0x00000082 {
		t.Errorf("STATUS should be 0x82, got 0x%08X", status)
	}

	eecd := uint32(data[0x10]) | uint32(data[0x11])<<8 | uint32(data[0x12])<<16 | uint32(data[0x13])<<24
	if eecd != 0x00000300 {
		t.Errorf("EECD should be 0x300, got 0x%08X", eecd)
	}
}

func TestEthernetStrategy_ScrubBAR_TooShort(t *testing.T) {
	s := &ethernetStrategy{}
	data := make([]byte, 0x08)
	s.ScrubBAR(data) // must not panic
}

func TestAudioStrategy_ScrubBAR(t *testing.T) {
	s := &audioStrategy{}
	data := make([]byte, 0x10)

	s.ScrubBAR(data)

	if data[0x08]&0x01 != 0x01 {
		t.Error("GCTL.CRST should be set")
	}
	// STATESTS codec 0 present — bit 16 of DWORD at 0x0C
	if data[0x0E]&0x01 != 0x01 {
		t.Error("STATESTS codec 0 present should be set")
	}
}

func TestAudioStrategy_ScrubBAR_TooShort(t *testing.T) {
	s := &audioStrategy{}
	data := make([]byte, 0x04)
	s.ScrubBAR(data)
}

func TestGPUStrategy_ScrubBAR(t *testing.T) {
	s := &gpuStrategy{}
	data := make([]byte, 0x204)

	s.ScrubBAR(data)

	pmcEnable := uint32(data[0x200]) | uint32(data[0x201])<<8 |
		uint32(data[0x202])<<16 | uint32(data[0x203])<<24
	if pmcEnable != 0xFFFFFFFF {
		t.Errorf("PMC_ENABLE should be 0xFFFFFFFF, got 0x%08X", pmcEnable)
	}
}

func TestGPUStrategy_ScrubBAR_TooShort(t *testing.T) {
	s := &gpuStrategy{}
	data := make([]byte, 0x100)
	s.ScrubBAR(data)
}

func TestSATAStrategy_ScrubBAR(t *testing.T) {
	s := &sataStrategy{}
	data := make([]byte, 0x12C)

	s.ScrubBAR(data)

	ghc := uint32(data[0x04]) | uint32(data[0x05])<<8 |
		uint32(data[0x06])<<16 | uint32(data[0x07])<<24
	if ghc&0x80000000 == 0 {
		t.Error("GHC.AE should be set")
	}
	if ghc&0x02 != 0 {
		t.Error("GHC.IE should be cleared")
	}

	ssts := uint32(data[0x128]) | uint32(data[0x129])<<8 |
		uint32(data[0x12A])<<16 | uint32(data[0x12B])<<24
	if ssts != 0x00000113 {
		t.Errorf("PxSSTS should be 0x113, got 0x%08X", ssts)
	}
}

func TestSATAStrategy_ScrubBAR_TooShort(t *testing.T) {
	s := &sataStrategy{}
	data := make([]byte, 0x10)
	s.ScrubBAR(data)
}

func TestWiFiStrategy_ScrubBAR(t *testing.T) {
	s := &wifiStrategy{}
	data := make([]byte, 0x58)

	s.ScrubBAR(data)

	gpCtl := uint32(data[0x24]) | uint32(data[0x25])<<8 |
		uint32(data[0x26])<<16 | uint32(data[0x27])<<24
	if gpCtl != 0x00000080 {
		t.Errorf("GP_CTL should be 0x80, got 0x%08X", gpCtl)
	}

	ucode := uint32(data[0x54]) | uint32(data[0x55])<<8 |
		uint32(data[0x56])<<16 | uint32(data[0x57])<<24
	if ucode != 0x00000001 {
		t.Errorf("UCODE_DRV_GP1 should be 0x01, got 0x%08X", ucode)
	}
}

func TestWiFiStrategy_ScrubBAR_TooShort(t *testing.T) {
	s := &wifiStrategy{}
	data := make([]byte, 0x10)
	s.ScrubBAR(data)
}

func TestThunderboltStrategy_ScrubBAR(t *testing.T) {
	s := &thunderboltStrategy{}
	data := make([]byte, 0x18)

	s.ScrubBAR(data)

	lcSts := uint32(data[0x08]) | uint32(data[0x09])<<8 |
		uint32(data[0x0A])<<16 | uint32(data[0x0B])<<24
	if lcSts != 0x00000001 {
		t.Errorf("LC_STS should be 0x01, got 0x%08X", lcSts)
	}

	secLvl := uint32(data[0x10]) | uint32(data[0x11])<<8 |
		uint32(data[0x12])<<16 | uint32(data[0x13])<<24
	if secLvl != 0x00000000 {
		t.Errorf("SECURITY_LEVEL should be 0x00, got 0x%08X", secLvl)
	}
}

func TestThunderboltStrategy_ScrubBAR_TooShort(t *testing.T) {
	s := &thunderboltStrategy{}
	data := make([]byte, 0x04)
	s.ScrubBAR(data)
}

func TestGenericStrategy_ScrubBAR_NoOp(t *testing.T) {
	s := &genericStrategy{}
	data := []byte{0xAA, 0xBB}
	s.ScrubBAR(data)
	if data[0] != 0xAA || data[1] != 0xBB {
		t.Error("generic ScrubBAR should be a no-op")
	}
}

// --- PostInitRegisters tests ---

func TestNVMeStrategy_PostInitRegisters(t *testing.T) {
	s := &nvmeStrategy{}
	var csts uint32 = 0x0C
	regs := map[uint32]*uint32{0x1C: &csts}
	s.PostInitRegisters(regs)
	if csts&0x01 == 0 {
		t.Error("CSTS.RDY should be set")
	}
	if csts&0x0C != 0 {
		t.Error("CSTS.SHST should be cleared")
	}
}

func TestNVMeStrategy_PostInitRegisters_NoCSTS(t *testing.T) {
	s := &nvmeStrategy{}
	regs := map[uint32]*uint32{}
	s.PostInitRegisters(regs) // must not panic
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
		t.Error("STATESTS codec 0 present should be set")
	}
}

func TestAudioStrategy_PostInitRegisters_Empty(t *testing.T) {
	s := &audioStrategy{}
	regs := map[uint32]*uint32{}
	s.PostInitRegisters(regs)
}

func TestGPUStrategy_PostInitRegisters(t *testing.T) {
	s := &gpuStrategy{}
	var pmcEnable uint32 = 0x00
	regs := map[uint32]*uint32{0x200: &pmcEnable}
	s.PostInitRegisters(regs)
	if pmcEnable != 0xFFFFFFFF {
		t.Errorf("PMC_ENABLE should be 0xFFFFFFFF, got 0x%08X", pmcEnable)
	}
}

func TestSATAStrategy_PostInitRegisters(t *testing.T) {
	s := &sataStrategy{}
	var ghc uint32 = 0x02
	regs := map[uint32]*uint32{0x04: &ghc}
	s.PostInitRegisters(regs)
	if ghc&0x80000000 == 0 {
		t.Error("GHC.AE should be set")
	}
	if ghc&0x02 != 0 {
		t.Error("GHC.IE should be cleared")
	}
}

func TestWiFiStrategy_PostInitRegisters(t *testing.T) {
	s := &wifiStrategy{}
	var gpCtl uint32 = 0x00
	regs := map[uint32]*uint32{0x24: &gpCtl}
	s.PostInitRegisters(regs)
	if gpCtl != 0x00000080 {
		t.Errorf("GP_CTL should be 0x80, got 0x%08X", gpCtl)
	}
}

func TestThunderboltStrategy_PostInitRegisters(t *testing.T) {
	s := &thunderboltStrategy{}
	var lcSts uint32 = 0x00
	regs := map[uint32]*uint32{0x08: &lcSts}
	s.PostInitRegisters(regs)
	if lcSts&0x01 == 0 {
		t.Error("LC_STS.READY should be set")
	}
}

func TestGenericStrategy_PostInitRegisters_NoOp(t *testing.T) {
	s := &genericStrategy{}
	var val uint32 = 0x42
	regs := map[uint32]*uint32{0x00: &val}
	s.PostInitRegisters(regs)
	if val != 0x42 {
		t.Error("generic PostInit should be a no-op")
	}
}

func TestAllStrategies_Profile(t *testing.T) {
	codes := []uint32{
		0x010802, 0x0C0330, 0x020000, 0x040300,
		0x030000, 0x010601, 0x028000, 0x0C8000,
		0xFF0000, // generic
	}
	for _, code := range codes {
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
			t.Errorf("empty profile className for %s", s.ClassName())
		}
	}
}
