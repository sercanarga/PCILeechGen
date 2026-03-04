package pci

import (
	"testing"
)

func TestParseBDF(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    BDF
		wantErr bool
	}{
		{
			name:  "full format",
			input: "0000:03:00.0",
			want:  BDF{Domain: 0, Bus: 3, Device: 0, Function: 0},
		},
		{
			name:  "full format with domain",
			input: "0001:0a:1f.2",
			want:  BDF{Domain: 1, Bus: 0x0a, Device: 0x1f, Function: 2},
		},
		{
			name:  "short format",
			input: "03:00.0",
			want:  BDF{Domain: 0, Bus: 3, Device: 0, Function: 0},
		},
		{
			name:  "with whitespace",
			input: "  0000:03:00.0  ",
			want:  BDF{Domain: 0, Bus: 3, Device: 0, Function: 0},
		},
		{
			name:    "invalid format",
			input:   "invalid",
			wantErr: true,
		},
		{
			name:    "empty string",
			input:   "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseBDF(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseBDF() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("ParseBDF() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestBDFString(t *testing.T) {
	bdf := BDF{Domain: 0, Bus: 3, Device: 0, Function: 0}
	if got := bdf.String(); got != "0000:03:00.0" {
		t.Errorf("BDF.String() = %q, want %q", got, "0000:03:00.0")
	}
}

func TestBDFShort(t *testing.T) {
	bdf := BDF{Domain: 0, Bus: 3, Device: 0, Function: 0}
	if got := bdf.Short(); got != "03:00.0" {
		t.Errorf("BDF.Short() = %q, want %q", got, "03:00.0")
	}
}

func TestBDFSysfsPath(t *testing.T) {
	bdf := BDF{Domain: 0, Bus: 3, Device: 0, Function: 0}
	want := "/sys/bus/pci/devices/0000:03:00.0"
	if got := bdf.SysfsPath(); got != want {
		t.Errorf("BDF.SysfsPath() = %q, want %q", got, want)
	}
}

func TestPCIDeviceClassDescription(t *testing.T) {
	tests := []struct {
		classCode uint32
		want      string
	}{
		{0x020000, "Ethernet controller"},
		{0x010600, "SATA controller"},
		{0x030000, "VGA compatible controller"},
		{0x040300, "Audio device"},
		{0x060000, "Host bridge"},
		{0x060400, "PCI bridge"},
		{0x0C0300, "USB controller"},
		{0xFF0000, "Unassigned class"},
	}

	for _, tt := range tests {
		dev := &PCIDevice{ClassCode: tt.classCode}
		if got := dev.ClassDescription(); got != tt.want {
			t.Errorf("ClassDescription() for class 0x%06x = %q, want %q", tt.classCode, got, tt.want)
		}
	}
}

func TestPCIDeviceSummary(t *testing.T) {
	dev := &PCIDevice{
		BDF:        BDF{Domain: 0, Bus: 3, Device: 0, Function: 0},
		VendorID:   0x8086,
		DeviceID:   0x1533,
		ClassCode:  0x020000,
		RevisionID: 0x03,
	}
	summary := dev.Summary()
	if summary == "" {
		t.Error("Summary() returned empty string")
	}
}

func TestPCIDeviceProgIF(t *testing.T) {
	dev := &PCIDevice{ClassCode: 0x0C0330}
	if dev.ProgIF() != 0x30 {
		t.Errorf("ProgIF() = 0x%02x, want 0x30", dev.ProgIF())
	}
	if dev.SubClass() != 0x03 {
		t.Errorf("SubClass() = 0x%02x, want 0x03", dev.SubClass())
	}
	if dev.BaseClass() != 0x0C {
		t.Errorf("BaseClass() = 0x%02x, want 0x0C", dev.BaseClass())
	}
}

func TestPCIDeviceClassDescriptionUnknown(t *testing.T) {
	// Unknown class code should fall back to "Class [XXYY]" format
	dev := &PCIDevice{ClassCode: 0xAB1200}
	desc := dev.ClassDescription()
	if desc == "" {
		t.Error("ClassDescription returned empty for unknown class")
	}

	// Known base class but unknown subclass should use base class name
	dev2 := &PCIDevice{ClassCode: 0x019900}
	desc2 := dev2.ClassDescription()
	if desc2 != "Mass storage controller" {
		t.Errorf("ClassDescription() = %q, want base class fallback", desc2)
	}
}
