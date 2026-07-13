package devclass

import "testing"

func TestAssessEmulationByFamily(t *testing.T) {
	tests := []struct {
		name      string
		classCode uint32
		vendorID  uint16
		deviceID  uint16
		hasMSIX   bool
		family    string
		level     EmulationLevel
		validated bool
	}{
		{"nvme", 0x010802, 0x144D, 0xA808, true, "nvme", EmulationDMA, true},
		{"intel-i219-msi", 0x020000, 0x8086, 0x15B7, false, "intel-e1000e-i219", EmulationDMA, true},
		{"intel-i219-msix-fallback", 0x020000, 0x8086, 0x15B7, true, "intel-e1000e-i219", EmulationRegisters, false},
		{"realtek-rtl8125", 0x020000, 0x10EC, 0x8125, false, "realtek-rtl8125", EmulationRegisters, true},
		{"unknown-ethernet", 0x020000, 0x14E4, 0x16B1, false, "ethernet-generic", EmulationIdentity, false},
		{"xhci", 0x0C0330, 0x8086, 0xA36D, true, "xhci-generic", EmulationRegisters, false},
		{"hda", 0x040300, 0x8086, 0xA348, false, "hda-generic", EmulationDMA, true},
		{"sata-ahci", 0x010601, 0x8086, 0xA352, true, "ahci-generic", EmulationRegisters, false},
		{"wifi", 0x028000, 0x14C3, 0x7922, true, "wifi-mediatek", EmulationRegisters, false},
		{"gpu", 0x030000, 0x10DE, 0x2684, true, "gpu-generic", EmulationRegisters, false},
		{"thunderbolt", 0x0C8000, 0x8086, 0x15EF, true, "thunderbolt-nhi", EmulationRegisters, false},
		{"unknown", 0xFF0000, 0x1234, 0x5678, false, "generic", EmulationIdentity, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			support := AssessEmulation(tt.classCode, tt.vendorID, tt.deviceID, tt.hasMSIX)
			if support.Family != tt.family {
				t.Fatalf("Family = %q, want %q", support.Family, tt.family)
			}
			if support.Level != tt.level {
				t.Fatalf("Level = %q, want %q", support.Level, tt.level)
			}
			if support.Validated != tt.validated {
				t.Fatalf("Validated = %v, want %v", support.Validated, tt.validated)
			}
			if support.Level != EmulationDMA && len(support.Limitations) == 0 && support.Family != "generic" {
				t.Fatal("partial family must report at least one limitation")
			}
		})
	}
}

func TestEmulationSupportCompleteOnlyWithoutLimitations(t *testing.T) {
	if (EmulationSupport{Level: EmulationDMA, Validated: true}).Complete() != true {
		t.Fatal("validated DMA support without limitations should be complete")
	}
	if (EmulationSupport{Level: EmulationDMA, Validated: true, Limitations: []string{"missing I/O queues"}}).Complete() {
		t.Fatal("support with limitations must not be complete")
	}
	if (EmulationSupport{Level: EmulationRegisters, Validated: true}).Complete() {
		t.Fatal("register-only support must not be complete")
	}
}
