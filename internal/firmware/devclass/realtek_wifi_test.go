package devclass

import (
	"testing"

	"github.com/sercanarga/pcileechgen/internal/pci"
)

func TestStrategyForDevice_RealtekRTL8188EE(t *testing.T) {
	s := StrategyForDevice(0x028000, realtekVID, realtekRTL8188EEDeviceID)
	if s == nil {
		t.Fatal("expected Realtek RTL8188EE strategy, got nil")
	}
	if s.ClassName() != "Wi-Fi (Realtek RTL8188EE)" {
		t.Fatalf("expected Wi-Fi (Realtek RTL8188EE), got %s", s.ClassName())
	}
	if s.DeviceClass() != ClassWiFi {
		t.Fatalf("expected %s, got %s", ClassWiFi, s.DeviceClass())
	}
}

func TestStrategyForDevice_UnknownRealtekWiFiFallsBack(t *testing.T) {
	s := StrategyForDevice(0x028000, realtekVID, 0xFFFF)
	if s.ClassName() == "Wi-Fi (Realtek RTL8188EE)" {
		t.Fatal("unknown Realtek Wi-Fi device should not use the RTL8188EE-specific profile")
	}
	if s.ClassName() != "Wi-Fi" {
		t.Fatalf("expected generic Wi-Fi fallback, got %s", s.ClassName())
	}
}

func TestRealtekWiFiProfile_UsesMSIAnd32BitBAR(t *testing.T) {
	p := StrategyForDevice(0x028000, realtekVID, realtekRTL8188EEDeviceID).Profile()
	if p.PrefersMSIX {
		t.Fatal("RTL8188EE donor profile should prefer MSI, not MSI-X")
	}
	if p.Uses64BitBAR {
		t.Fatal("RTL8188EE donor profile should use a 32-bit BAR0")
	}

	caps := map[uint8]bool{}
	for _, capID := range p.ExpectedCaps {
		caps[capID] = true
	}
	for _, capID := range []uint8{pci.CapIDPowerManagement, pci.CapIDMSI, pci.CapIDPCIExpress} {
		if !caps[capID] {
			t.Fatalf("RTL8188EE profile missing expected cap 0x%02X", capID)
		}
	}
}
