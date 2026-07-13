package nvme

import (
	"encoding/binary"
	"strings"
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
	id := BuildIdentifyData(sampleIDs(), nil, nil)
	if id == nil {
		t.Fatal("BuildIdentifyData returned nil")
	}
}

func TestIdentifyController_VID(t *testing.T) {
	id := BuildIdentifyData(sampleIDs(), nil, nil)
	vid := binary.LittleEndian.Uint16(id.Controller[0x000:])
	if vid != 0x144D {
		t.Errorf("VID: got 0x%04X, want 0x144D", vid)
	}
}

func TestIdentifyController_SSVID(t *testing.T) {
	id := BuildIdentifyData(sampleIDs(), nil, nil)
	ssvid := binary.LittleEndian.Uint16(id.Controller[0x002:])
	if ssvid != 0x144D {
		t.Errorf("SSVID: got 0x%04X, want 0x144D", ssvid)
	}
}

func TestIdentifyController_SN_NonEmpty(t *testing.T) {
	id := BuildIdentifyData(sampleIDs(), nil, nil)
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
	id := BuildIdentifyData(sampleIDs(), nil, nil)
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
	id := BuildIdentifyData(sampleIDs(), nil, nil)
	sqes := id.Controller[0x200]
	if sqes != 0x66 {
		t.Errorf("SQES: got 0x%02X, want 0x66", sqes)
	}
}

func TestIdentifyController_MetadataConsistency(t *testing.T) {
	id := BuildIdentifyData(sampleIDs(), nil, nil)

	if got := binary.LittleEndian.Uint16(id.Controller[0x100:]); got != 0x0003 {
		t.Fatalf("OACS = 0x%04X, want 0x0003", got)
	}
	if got := id.Controller[0x104]; got != 0x02 {
		t.Fatalf("FRMW = 0x%02X, want 0x02", got)
	}
	if got := id.Controller[0x105]; got != 0x00 {
		t.Fatalf("LPA = 0x%02X, want 0x00", got)
	}
	if got := id.Controller[0x106]; got != 0x00 {
		t.Fatalf("ELPE = 0x%02X, want 0x00", got)
	}
	if got := id.Controller[0x108]; got != 0x00 {
		t.Fatalf("AVSCC = 0x%02X, want 0x00", got)
	}
	if got := binary.LittleEndian.Uint16(id.Controller[0x10C:]); got != 358 {
		t.Fatalf("CCTEMP = %d, want 358", got)
	}
	if got := id.Controller[0x06F]; got != 0x01 {
		t.Fatalf("CNTRLTYPE = 0x%02X, want 0x01", got)
	}
	if got := binary.LittleEndian.Uint16(id.Controller[0x208:]); got != 0x000C {
		t.Fatalf("ONCS = 0x%04X, want 0x000C", got)
	}
}

func TestIdentifyController_CQES(t *testing.T) {
	id := BuildIdentifyData(sampleIDs(), nil, nil)
	cqes := id.Controller[0x201]
	if cqes != 0x44 {
		t.Errorf("CQES: got 0x%02X, want 0x44", cqes)
	}
}

func TestIdentifyController_NN(t *testing.T) {
	id := BuildIdentifyData(sampleIDs(), nil, nil)
	nn := binary.LittleEndian.Uint32(id.Controller[0x204:])
	if nn < 1 {
		t.Errorf("NN: got %d, want >= 1", nn)
	}
}

func TestIdentifyController_CNTLID(t *testing.T) {
	id := BuildIdentifyData(sampleIDs(), nil, nil)
	cntlid := binary.LittleEndian.Uint16(id.Controller[0x04E:])
	if cntlid == 0 {
		t.Errorf("CNTLID @ 0x050 = 0, want nonzero")
	}
}

func TestIdentifyController_MDTS_Default(t *testing.T) {
	id := BuildIdentifyData(sampleIDs(), nil, nil)
	mdts := id.Controller[0x04D]
	if mdts != 3 {
		t.Errorf("MDTS (nil BAR): got %d, want 3 (32KB, matches SV MAX_XFER_DW)", mdts)
	}
}

func TestIdentifyController_MDTS_WithBAR(t *testing.T) {
	barData := make([]byte, 64)
	binary.LittleEndian.PutUint32(barData[0x04:], 0x00000030) // CAP_HI with MPSMIN=0
	id := BuildIdentifyData(sampleIDs(), barData, nil)
	mdts := id.Controller[0x04D]
	if mdts != 3 {
		t.Errorf("MDTS (MPSMIN=0): got %d, want 3", mdts)
	}
}

func TestIdentifyNamespace_NSZE(t *testing.T) {
	id := BuildIdentifyData(sampleIDs(), nil, nil)
	nsze := binary.LittleEndian.Uint64(id.Namespace[0x000:])
	if nsze == 0 {
		t.Error("NSZE should be > 0")
	}
}

func TestIdentifyNamespace_NCAP(t *testing.T) {
	id := BuildIdentifyData(sampleIDs(), nil, nil)
	ncap := binary.LittleEndian.Uint64(id.Namespace[0x008:])
	nsze := binary.LittleEndian.Uint64(id.Namespace[0x000:])
	if ncap != nsze {
		t.Errorf("NCAP should equal NSZE: got %d, want %d", ncap, nsze)
	}
}

func TestIdentifyNamespace_LBAF0(t *testing.T) {
	id := BuildIdentifyData(sampleIDs(), nil, nil)
	lbaf0 := binary.LittleEndian.Uint32(id.Namespace[0x0C0:])
	// LBADS should be 9 (512B sectors): bits [23:16] = 0x09
	lbads := (lbaf0 >> 16) & 0xFF
	if lbads != 9 {
		t.Errorf("LBAF0 LBADS: got %d, want 9 (512B sectors)", lbads)
	}
}

func TestIdentifyDataToHex_Size(t *testing.T) {
	id := BuildIdentifyData(sampleIDs(), nil, nil)
	hex := IdentifyDataToHex(id)
	if len(hex) == 0 {
		t.Fatal("IdentifyDataToHex returned empty string")
	}
	// 4 KB Identify Controller only (1024 data lines); namespace is runtime-generated.
	lines := 0
	for _, c := range hex {
		if c == '\n' {
			lines++
		}
	}
	if lines != 1024 {
		t.Errorf("expected 1024 lines, got %d", lines)
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

func TestIdentifyController_MissingIdentityIsNeutral(t *testing.T) {
	ids := sampleIDs()
	ids.VendorID = 0x144D
	id := BuildIdentifyData(ids, nil, nil)
	mn := string(id.Controller[0x018:0x040])
	if strings.Contains(mn, "Samsung") {
		t.Errorf("missing donor identity must not claim Samsung, got: %q", mn)
	}
	if !strings.HasPrefix(mn, "PCILeechGen NVMe 144D:") {
		t.Errorf("missing donor identity should use neutral model, got: %q", mn)
	}
	sn := string(id.Controller[0x004:0x018])
	if strings.HasPrefix(sn, "S6PXNG0T") {
		t.Errorf("missing donor identity must not use Samsung serial prefix, got: %q", sn)
	}
	if got := trimCtrlString(id.Controller[0x040:0x048]); got != "1.0" {
		t.Errorf("missing donor identity should use neutral firmware revision, got: %q", got)
	}
}

// TestBuildIdentifyData_UsesCapturedIdentity checks donor strings override synthesized defaults.
func TestBuildIdentifyData_UsesCapturedIdentity(t *testing.T) {
	ids := sampleIDs() // Samsung VID 0x144D
	captured := &ControllerIdentity{
		Serial: "S6PXNG0T12345678W",
		Model:  "Samsung SSD 990 PRO 2TB",
		FWRev:  "4B2QJXD7",
	}
	id := BuildIdentifyData(ids, nil, captured)

	sn := trimCtrlString(id.Controller[0x04:0x18])
	if sn != captured.Serial {
		t.Errorf("SN: got %q, want %q (donor-captured)", sn, captured.Serial)
	}
	mn := trimCtrlString(id.Controller[0x18:0x40])
	if mn != captured.Model {
		t.Errorf("MN: got %q, want %q (donor-captured)", mn, captured.Model)
	}
	fr := trimCtrlString(id.Controller[0x40:0x48])
	if fr != captured.FWRev {
		t.Errorf("FR: got %q, want %q (donor-captured)", fr, captured.FWRev)
	}
}

// Raw donor Controller blob: MDTS clamped, VER==BAR VS, VID overridden.
func TestBuildIdentifyData_RawControllerClampsMDTSAndAlignsVER(t *testing.T) {
	ids := sampleIDs()
	raw := make([]byte, 4096)
	raw[0x04D] = 6                                         // MDTS too large for backend
	binary.LittleEndian.PutUint32(raw[0x050:], 0x00010300) // donor VER 1.3 (would mismatch BAR)
	binary.LittleEndian.PutUint16(raw[0x000:], 0xBEEF)     // donor VID (must be overridden)
	captured := &ControllerIdentity{RawControllerIdent: raw}

	barData := make([]byte, 0x10)
	binary.LittleEndian.PutUint32(barData[0x08:], 0x00010400) // BAR VS 1.4

	id := BuildIdentifyData(ids, barData, captured)

	if got := id.Controller[0x04D]; got != 3 {
		t.Errorf("raw MDTS must be 3 (SV match), got %d", got)
	}
	if got := binary.LittleEndian.Uint16(id.Controller[0x000:]); got != ids.VendorID {
		t.Errorf("VID must be forced to ids (0x%04X), got 0x%04X", ids.VendorID, got)
	}
}

func TestBuildIdentifyData_NilIdentityFallsBackToSynthesis(t *testing.T) {
	ids := sampleIDs()
	ids.VendorID = 0x144D
	id := BuildIdentifyData(ids, nil, nil)

	mn := trimCtrlString(id.Controller[0x18:0x40])
	if strings.Contains(mn, "Samsung") {
		t.Errorf("nil identity must not fall back to Samsung MN, got %q", mn)
	}
}

func TestBuildIdentifyData_PartialCapturedIdentityFallsBackPerField(t *testing.T) {
	ids := sampleIDs()
	captured := &ControllerIdentity{
		Serial: "REALSERIAL",
		Model:  "", // empty -> synthesize
	}
	id := BuildIdentifyData(ids, nil, captured)

	sn := trimCtrlString(id.Controller[0x04:0x18])
	if sn != "REALSERIAL" {
		t.Errorf("SN: got %q, want REALSERIAL", sn)
	}
	mn := trimCtrlString(id.Controller[0x18:0x40])
	if !strings.HasPrefix(mn, "PCILeechGen NVMe 144D:") {
		t.Errorf("empty Model should fall back to neutral model, got %q", mn)
	}
}

func trimCtrlString(b []byte) string {
	return strings.TrimRight(string(b), " \x00")
}

func TestIdentifyController_DifferentVendors(t *testing.T) {
	vendors := []uint16{0x144D, 0x8086, 0x15B7, 0x1C5C, 0x1179, 0x1987, 0x1234}
	for _, vid := range vendors {
		ids := sampleIDs()
		ids.VendorID = vid
		ids.SubsysVendorID = vid
		id := BuildIdentifyData(ids, nil, nil)
		gotVID := binary.LittleEndian.Uint16(id.Controller[0x000:])
		if gotVID != vid {
			t.Errorf("VID 0x%04X: identify VID mismatch, got 0x%04X", vid, gotVID)
		}
	}
}

// CAP_HI at offset 0x04 carries MPSMIN in bits 19:16 of the dword.
func TestDeriveMDTS(t *testing.T) {
	t.Run("nil_bar", func(t *testing.T) {
		if got := deriveMDTS(nil); got != 5 {
			t.Errorf("nil barData: got %d, want 5", got)
		}
	})
	t.Run("MPSMIN=0", func(t *testing.T) {
		barData := make([]byte, 0x10)
		binary.LittleEndian.PutUint32(barData[0x04:], 0x00000000) // MPSMIN=0
		if got := deriveMDTS(barData); got != 5 {
			t.Errorf("MPSMIN=0: got %d, want 5", got)
		}
	})
	t.Run("MPSMIN=2", func(t *testing.T) {
		barData := make([]byte, 0x10)
		binary.LittleEndian.PutUint32(barData[0x04:], 0x00020000) // MPSMIN=2 in bits 19:16
		if got := deriveMDTS(barData); got != 3 {
			t.Errorf("MPSMIN=2: got %d, want 3", got)
		}
	})
	t.Run("MPSMIN=4", func(t *testing.T) {
		barData := make([]byte, 0x10)
		binary.LittleEndian.PutUint32(barData[0x04:], 0x00040000) // MPSMIN=4 in bits 19:16
		if got := deriveMDTS(barData); got != 1 {
			t.Errorf("MPSMIN=4: got %d, want 1", got)
		}
	})
}

func TestDeriveVER(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		if got := deriveVER(nil); got != 0x00010400 {
			t.Errorf("nil barData: got 0x%08X, want 0x00010400", got)
		}
	})
	t.Run("short", func(t *testing.T) {
		if got := deriveVER(make([]byte, 8)); got != 0x00010400 {
			t.Errorf("short barData: got 0x%08X, want 0x00010400", got)
		}
	})
	t.Run("zero_VS", func(t *testing.T) {
		barData := make([]byte, 0x10)
		binary.LittleEndian.PutUint32(barData[0x08:], 0x00000000)
		if got := deriveVER(barData); got != 0x00010400 {
			t.Errorf("zero VS: got 0x%08X, want 0x00010400", got)
		}
	})
	t.Run("donor_VS", func(t *testing.T) {
		barData := make([]byte, 0x10)
		binary.LittleEndian.PutUint32(barData[0x08:], 0x00010300)
		if got := deriveVER(barData); got != 0x00010300 {
			t.Errorf("donor VS: got 0x%08X, want 0x00010300", got)
		}
	})
}

func TestDoorbellStrideFromCAP(t *testing.T) {
	cases := []struct {
		name  string
		capHi uint32
		want  uint32
	}{
		{"DSTRD=0 stride 4B", 0x00000000, 0},
		{"DSTRD=2 stride 16B", 0x00000002, 2},
		{"DSTRD=5", 0x00000005, 5},
		{"isolated from MPSMIN/CSS high bits", 0x000F0002, 2},
		{"max DSTRD=15", 0x0000000F, 15},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := DoorbellStrideFromCAP(c.capHi); got != c.want {
				t.Errorf("DoorbellStrideFromCAP(0x%08X): got %d, want %d", c.capHi, got, c.want)
			}
		})
	}
}

// Raw donor Controller blob: SSVID, OAES, CTRATT also forced.
func TestBuildIdentifyData_RawControllerOverridesAllForced(t *testing.T) {
	ids := sampleIDs()
	raw := make([]byte, 4096)
	binary.LittleEndian.PutUint16(raw[0x000:], 0xBEEF)     // donor VID
	binary.LittleEndian.PutUint16(raw[0x002:], 0xDEAD)     // donor SSVID
	binary.LittleEndian.PutUint32(raw[0x05C:], 0xCAFEBABE) // donor OAES
	binary.LittleEndian.PutUint32(raw[0x060:], 0xFEEDFACE) // donor CTRATT
	captured := &ControllerIdentity{RawControllerIdent: raw}

	id := BuildIdentifyData(ids, nil, captured)

	if got := binary.LittleEndian.Uint16(id.Controller[0x000:]); got != ids.VendorID {
		t.Errorf("VID must be forced to ids (0x%04X), got 0x%04X", ids.VendorID, got)
	}
	if got := binary.LittleEndian.Uint16(id.Controller[0x002:]); got != ids.SubsysVendorID {
		t.Errorf("SSVID must be forced to ids (0x%04X), got 0x%04X", ids.SubsysVendorID, got)
	}
	if got := binary.LittleEndian.Uint32(id.Controller[0x05C:]); got != 0 {
		t.Errorf("OAES must be forced to 0, got 0x%08X", got)
	}
	if got := binary.LittleEndian.Uint32(id.Controller[0x060:]); got != 0 {
		t.Errorf("CTRATT must be forced to 0, got 0x%08X", got)
	}
}

func TestBuildIdentifyData_RawNamespaceNormalizesToFirmwareContract(t *testing.T) {
	rawNS := make([]byte, 4096)
	binary.LittleEndian.PutUint64(rawNS[0x000:], 250000)
	binary.LittleEndian.PutUint64(rawNS[0x008:], 250000)
	binary.LittleEndian.PutUint64(rawNS[0x010:], 120000)
	rawNS[0x01A] = 0x01
	binary.LittleEndian.PutUint32(rawNS[0x0C0:], 0x000C0000)

	id := BuildIdentifyData(sampleIDs(), nil, &ControllerIdentity{RawNamespaceIdent: rawNS})
	info := NamespaceInfoFromIdentify(id)

	if info.LBADataSizePower != 9 {
		t.Fatalf("LBADS normalized: got %d want 9", info.LBADataSizePower)
	}
	if info.ActiveFormat != 0 {
		t.Fatalf("active format normalized: got %d want 0", info.ActiveFormat)
	}
	if info.NSZE != 2000000 || info.NCAP != 2000000 || info.NUSE != 960000 {
		t.Fatalf("namespace size normalization mismatch: %+v", info)
	}
	if issues := ValidateNamespace(id.Namespace[:]); len(issues) != 0 {
		t.Fatalf("ValidateNamespace returned issues: %v", issues)
	}
}
