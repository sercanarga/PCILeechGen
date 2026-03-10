package nvme

import (
	"encoding/binary"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/firmware"
)

func sampleIDs() firmware.DeviceIDs {
	return firmware.DeviceIDs{
		VendorID:       0x144D,
		DeviceID:       0xA808,
		SubsysVendorID: 0x144D,
		SubsysDeviceID: 0xA801,
		RevisionID:     0x00,
		ClassCode:      0x010802,
	}
}

func TestBuildIdentifyData_NotNil(t *testing.T) {
	id := BuildIdentifyData(sampleIDs(), nil)
	if id == nil {
		t.Fatal("BuildIdentifyData returned nil")
	}
}

func TestIdentifyController_VID(t *testing.T) {
	id := BuildIdentifyData(sampleIDs(), nil)
	vid := binary.LittleEndian.Uint16(id.Controller[0x000:])
	if vid != 0x144D {
		t.Errorf("VID: got 0x%04X, want 0x144D", vid)
	}
}

func TestIdentifyController_SSVID(t *testing.T) {
	id := BuildIdentifyData(sampleIDs(), nil)
	ssvid := binary.LittleEndian.Uint16(id.Controller[0x002:])
	if ssvid != 0x144D {
		t.Errorf("SSVID: got 0x%04X, want 0x144D", ssvid)
	}
}

func TestIdentifyController_SN_NonEmpty(t *testing.T) {
	id := BuildIdentifyData(sampleIDs(), nil)
	sn := id.Controller[0x004:0x018]
	allSpaces := true
	for _, b := range sn {
		if b != ' ' && b != 0 {
			allSpaces = false
			break
		}
	}
	if allSpaces {
		t.Error("SN should not be all spaces/zeros")
	}
}

func TestIdentifyController_MN_NonEmpty(t *testing.T) {
	id := BuildIdentifyData(sampleIDs(), nil)
	mn := id.Controller[0x018:0x040]
	allSpaces := true
	for _, b := range mn {
		if b != ' ' && b != 0 {
			allSpaces = false
			break
		}
	}
	if allSpaces {
		t.Error("MN should not be all spaces/zeros")
	}
}

func TestIdentifyController_SQES(t *testing.T) {
	id := BuildIdentifyData(sampleIDs(), nil)
	sqes := id.Controller[0x200]
	if sqes != 0x66 {
		t.Errorf("SQES: got 0x%02X, want 0x66", sqes)
	}
}

func TestIdentifyController_CQES(t *testing.T) {
	id := BuildIdentifyData(sampleIDs(), nil)
	cqes := id.Controller[0x201]
	if cqes != 0x44 {
		t.Errorf("CQES: got 0x%02X, want 0x44", cqes)
	}
}

func TestIdentifyController_NN(t *testing.T) {
	id := BuildIdentifyData(sampleIDs(), nil)
	nn := binary.LittleEndian.Uint32(id.Controller[0x204:])
	if nn < 1 {
		t.Errorf("NN: got %d, want >= 1", nn)
	}
}

func TestIdentifyController_Version(t *testing.T) {
	// nil BAR data -> default NVMe 1.4
	id := BuildIdentifyData(sampleIDs(), nil)
	ver := binary.LittleEndian.Uint32(id.Controller[0x050:])
	if ver != 0x00010400 {
		t.Errorf("VER (nil BAR): got 0x%08X, want 0x00010400", ver)
	}
}

func TestIdentifyController_Version_FromBAR(t *testing.T) {
	// BAR data with VS=0x00010300 (NVMe 1.3) at offset 0x08
	barData := make([]byte, 64)
	binary.LittleEndian.PutUint32(barData[0x08:], 0x00010300)

	id := BuildIdentifyData(sampleIDs(), barData)
	ver := binary.LittleEndian.Uint32(id.Controller[0x050:])
	if ver != 0x00010300 {
		t.Errorf("VER (BAR NVMe 1.3): got 0x%08X, want 0x00010300", ver)
	}
}

func TestIdentifyController_Version_ShortBAR(t *testing.T) {
	// BAR data shorter than 0x0C -> falls back to default
	barData := make([]byte, 8)
	id := BuildIdentifyData(sampleIDs(), barData)
	ver := binary.LittleEndian.Uint32(id.Controller[0x050:])
	if ver != 0x00010400 {
		t.Errorf("VER (short BAR): got 0x%08X, want 0x00010400", ver)
	}
}

func TestIdentifyController_Version_ZeroVS(t *testing.T) {
	// BAR data with VS=0 -> falls back to default
	barData := make([]byte, 64)
	id := BuildIdentifyData(sampleIDs(), barData)
	ver := binary.LittleEndian.Uint32(id.Controller[0x050:])
	if ver != 0x00010400 {
		t.Errorf("VER (zero VS): got 0x%08X, want 0x00010400", ver)
	}
}

func TestIdentifyController_MDTS_Default(t *testing.T) {
	id := BuildIdentifyData(sampleIDs(), nil)
	mdts := id.Controller[0x04D]
	if mdts != 5 {
		t.Errorf("MDTS (nil BAR): got %d, want 5", mdts)
	}
}

func TestIdentifyController_MDTS_WithBAR(t *testing.T) {
	// MPSMIN=0 (4KB pages) -> MDTS stays 5
	barData := make([]byte, 64)
	binary.LittleEndian.PutUint32(barData[0x04:], 0x00000030) // CAP_HI with MPSMIN=0
	id := BuildIdentifyData(sampleIDs(), barData)
	mdts := id.Controller[0x04D]
	if mdts != 5 {
		t.Errorf("MDTS (MPSMIN=0): got %d, want 5", mdts)
	}
}

func TestIdentifyNamespace_NSZE(t *testing.T) {
	id := BuildIdentifyData(sampleIDs(), nil)
	nsze := binary.LittleEndian.Uint64(id.Namespace[0x000:])
	if nsze == 0 {
		t.Error("NSZE should be > 0")
	}
}

func TestIdentifyNamespace_NCAP(t *testing.T) {
	id := BuildIdentifyData(sampleIDs(), nil)
	ncap := binary.LittleEndian.Uint64(id.Namespace[0x008:])
	nsze := binary.LittleEndian.Uint64(id.Namespace[0x000:])
	if ncap != nsze {
		t.Errorf("NCAP should equal NSZE: got %d, want %d", ncap, nsze)
	}
}

func TestIdentifyNamespace_LBAF0(t *testing.T) {
	id := BuildIdentifyData(sampleIDs(), nil)
	lbaf0 := binary.LittleEndian.Uint32(id.Namespace[0x0C0:])
	// LBADS should be 9 (512B sectors): bits [23:16] = 0x09
	lbads := (lbaf0 >> 16) & 0xFF
	if lbads != 9 {
		t.Errorf("LBAF0 LBADS: got %d, want 9 (512B sectors)", lbads)
	}
}

func TestIdentifyDataToHex_Size(t *testing.T) {
	id := BuildIdentifyData(sampleIDs(), nil)
	hex := IdentifyDataToHex(id)
	if len(hex) == 0 {
		t.Fatal("IdentifyDataToHex returned empty string")
	}
	// Should contain 2048 data lines (8KB / 4 bytes per line)
	// plus 2 header comment lines
	lines := 0
	for _, c := range hex {
		if c == '\n' {
			lines++
		}
	}
	// 2 comment lines + 2048 data lines = 2050
	if lines != 2050 {
		t.Errorf("expected 2050 lines, got %d", lines)
	}
}

func TestPadASCII(t *testing.T) {
	result := padASCII("ABC", 8)
	if len(result) != 8 {
		t.Errorf("padASCII length: got %d, want 8", len(result))
	}
	if string(result[:3]) != "ABC" {
		t.Errorf("padASCII content: got %q", result[:3])
	}
	for i := 3; i < 8; i++ {
		if result[i] != ' ' {
			t.Errorf("padASCII padding at %d: got 0x%02X, want space", i, result[i])
		}
	}
}

func TestIdentifyController_Samsung_MN(t *testing.T) {
	ids := sampleIDs()
	ids.VendorID = 0x144D
	id := BuildIdentifyData(ids, nil)
	mn := string(id.Controller[0x018:0x040])
	if mn[:7] != "Samsung" {
		t.Errorf("Samsung VID should produce Samsung MN, got: %q", mn)
	}
}

func TestIdentifyController_DifferentVendors(t *testing.T) {
	vendors := []uint16{0x144D, 0x8086, 0x15B7, 0x1C5C, 0x1179, 0x1987, 0x1234}
	for _, vid := range vendors {
		ids := sampleIDs()
		ids.VendorID = vid
		ids.SubsysVendorID = vid
		id := BuildIdentifyData(ids, nil)
		gotVID := binary.LittleEndian.Uint16(id.Controller[0x000:])
		if gotVID != vid {
			t.Errorf("VID 0x%04X: identify VID mismatch, got 0x%04X", vid, gotVID)
		}
	}
}
