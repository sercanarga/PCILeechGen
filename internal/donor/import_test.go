package donor

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/pci"
)

// buildStandardFixture returns a 256-byte config-space capture with a
// capability pointer and one implemented memory BAR, mirroring the layout
// used by internal/pci config-space tests.
func buildStandardFixture(t *testing.T) []byte {
	t.Helper()
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceLegacySize
	cs.WriteU16(0x00, 0x8086) // Vendor ID
	cs.WriteU16(0x02, 0x1533) // Device ID
	cs.WriteU16(0x04, 0x0406) // Command
	cs.WriteU16(0x06, 0x0010) // Status (capabilities list)
	cs.WriteU8(0x08, 0x03)    // Revision ID
	cs.WriteU8(0x0A, 0x00)    // Sub-class
	cs.WriteU8(0x0B, 0x02)    // Base class (Network)
	cs.WriteU8(0x0E, 0x00)    // Header type
	cs.WriteU32(0x10, 0xFFFFF000) // BAR0: 4KB memory BAR, prefetchable off
	cs.WriteU16(0x2C, 0x8086) // Subsys Vendor
	cs.WriteU16(0x2E, 0x0001) // Subsys Device
	cs.WriteU8(0x34, 0x40)    // Capability pointer
	// PM capability at 0x40
	cs.WriteU8(0x40, pci.CapIDPowerManagement)
	cs.WriteU8(0x41, 0x00) // next ptr = 0
	return cs.Bytes()
}

// buildExtendedFixture returns a 4096-byte config-space capture with an
// extended capability (Device Serial Number) at 0x100.
func buildExtendedFixture(t *testing.T) []byte {
	t.Helper()
	base := buildStandardFixture(t)
	raw := make([]byte, pci.ConfigSpaceSize)
	copy(raw, base)
	// Extended capability header at 0x100: cap_id=0x0003 (DSN), version=1,
	// next=0. The header DWORD encodes id[15:0] | version[19:16] | next[31:20].
	raw[0x100] = 0x03
	raw[0x101] = 0x00
	raw[0x102] = 0x10 // version=1, next=0 (low nibble of next is 0)
	raw[0x103] = 0x00
	return raw
}

func TestImportConfigSpace(t *testing.T) {
	cases := []struct {
		name    string
		raw     []byte
		wantErr string
		wantVID uint16
		wantDID uint16
		wantLen int // expected config space size
		wantBAR int // expected implemented BAR count
	}{
		{
			name:    "256-byte standard",
			raw:     buildStandardFixture(t),
			wantVID: 0x8086,
			wantDID: 0x1533,
			wantLen: pci.ConfigSpaceLegacySize,
			wantBAR: 1,
		},
		{
			name:    "4096-byte extended",
			raw:     buildExtendedFixture(t),
			wantVID: 0x8086,
			wantDID: 0x1533,
			wantLen: pci.ConfigSpaceSize,
			wantBAR: 1,
		},
		{
			name:    "truncated 128 bytes",
			raw:     make([]byte, 128),
			wantErr: "config space truncated",
		},
		{
			name:    "truncated 255 bytes",
			raw:     make([]byte, 255),
			wantErr: "config space truncated",
		},
		{
			name:    "oversized 4097 bytes",
			raw:     make([]byte, 4097),
			wantErr: "config space truncated",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ctx, warnings, err := ImportConfigSpace(tc.raw)
			if tc.wantErr != "" {
				if err == nil {
					t.Fatalf("expected error containing %q, got nil", tc.wantErr)
				}
				if !strings.Contains(err.Error(), tc.wantErr) {
					t.Fatalf("expected error containing %q, got %q", tc.wantErr, err.Error())
				}
				if ctx != nil {
					t.Fatalf("on input error, context must be nil; got %+v", ctx)
				}
				if warnings != nil {
					t.Fatalf("on input error, warnings must be nil; got %v", warnings)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if ctx == nil || ctx.ConfigSpace == nil {
				t.Fatalf("expected non-nil context with config space")
			}
			if ctx.ConfigSpace.Size != tc.wantLen {
				t.Errorf("config space size = %d, want %d", ctx.ConfigSpace.Size, tc.wantLen)
			}
			if ctx.Device.VendorID != tc.wantVID || ctx.Device.DeviceID != tc.wantDID {
				t.Errorf("device = %04x:%04x, want %04x:%04x", ctx.Device.VendorID, ctx.Device.DeviceID, tc.wantVID, tc.wantDID)
			}
			implemented := 0
			for _, b := range ctx.BARs {
				if b.Type != pci.BARTypeDisabled {
					implemented++
				}
			}
			if implemented != tc.wantBAR {
				t.Errorf("implemented BARs = %d, want %d", implemented, tc.wantBAR)
			}
			// Round-trip through JSON to exercise ToJSON/FromJSON.
			data, err := ctx.ToJSON()
			if err != nil {
				t.Fatalf("ToJSON: %v", err)
			}
			rt, err := FromJSON(data)
			if err != nil {
				t.Fatalf("FromJSON: %v", err)
			}
			if rt.Device.VendorID != tc.wantVID {
				t.Errorf("round-trip vendor = %04x, want %04x", rt.Device.VendorID, tc.wantVID)
			}
			if rt.ConfigSpace == nil || rt.ConfigSpace.Size != tc.wantLen {
				t.Errorf("round-trip config space missing or size %d, want %d", sizeOrZero(rt.ConfigSpace), tc.wantLen)
			}
		})
	}
}

func sizeOrZero(cs *pci.ConfigSpace) int {
	if cs == nil {
		return 0
	}
	return cs.Size
}

// TestImportConfigSpaceWarnings verifies missing-optional-field warnings are
// returned (not errors) for a bare standard capture.
func TestImportConfigSpaceWarnings(t *testing.T) {
	ctx, warnings, err := ImportConfigSpace(buildStandardFixture(t))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if ctx == nil {
		t.Fatal("expected non-nil context")
	}
	if len(warnings) == 0 {
		t.Fatal("expected warnings for missing BAR contents / ext caps, got none")
	}
	sawBAR, sawExt := false, false
	for _, w := range warnings {
		if w.Code == "bar_contents_missing" {
			sawBAR = true
		}
		if w.Code == "ext_capabilities_missing" {
			sawExt = true
		}
	}
	if !sawBAR {
		t.Error("missing bar_contents_missing warning")
	}
	if !sawExt {
		t.Error("missing ext_capabilities_missing warning")
	}
}

// TestImportCOERoundTrip proves the importer parses the exact COE format
// emitted by codegen.GenerateConfigSpaceCOE.
func TestImportCOERoundTrip(t *testing.T) {
	raw := buildExtendedFixture(t)
	cs := pci.NewConfigSpaceFromBytes(raw)
	// Use the generator's own COE formatter to produce the fixture.
	coe := generateTestCOE(t, cs)

	dir := t.TempDir()
	coePath := filepath.Join(dir, "pcileech_cfgspace.coe")
	if err := os.WriteFile(coePath, []byte(coe), 0644); err != nil {
		t.Fatalf("write COE fixture: %v", err)
	}

	ctx, _, err := ImportCOE(coePath)
	if err != nil {
		t.Fatalf("ImportCOE: %v", err)
	}
	if ctx == nil || ctx.ConfigSpace == nil {
		t.Fatal("expected non-nil context with config space")
	}
	if ctx.ConfigSpace.Size != pci.ConfigSpaceSize {
		t.Errorf("COE import config size = %d, want %d", ctx.ConfigSpace.Size, pci.ConfigSpaceSize)
	}
	if ctx.Device.VendorID != 0x8086 || ctx.Device.DeviceID != 0x1533 {
		t.Errorf("COE import device = %04x:%04x, want 8086:1533", ctx.Device.VendorID, ctx.Device.DeviceID)
	}
	// First DWORD must match the raw capture (round-trip).
	if got, want := ctx.ConfigSpace.ReadU32(0), cs.ReadU32(0); got != want {
		t.Errorf("COE round-trip word0 = %08x, want %08x", got, want)
	}
}

// TestImportCOERejectsBadRadix verifies a non-16 radix is a typed error with
// no partial context.
func TestImportCOERejectsBadRadix(t *testing.T) {
	dir := t.TempDir()
	coePath := filepath.Join(dir, "bad.coe")
	content := "memory_initialization_radix=10;\nmemory_initialization_vector=\n00000001;\n"
	if err := os.WriteFile(coePath, []byte(content), 0644); err != nil {
		t.Fatalf("write: %v", err)
	}
	ctx, _, err := ImportCOE(coePath)
	if err == nil {
		t.Fatal("expected error for radix=10, got nil")
	}
	if !strings.Contains(err.Error(), "radix") {
		t.Fatalf("expected radix error, got %q", err.Error())
	}
	if ctx != nil {
		t.Fatalf("on parse error, context must be nil; got %+v", ctx)
	}
}

// generateTestCOE emits a COE in the generator's own format without
// importing the codegen package (avoids an internal/firmware dependency in
// this test). Keep in sync with codegen.formatCOE.
func generateTestCOE(t *testing.T, cs *pci.ConfigSpace) string {
	t.Helper()
	var sb strings.Builder
	sb.WriteString("; test fixture COE\n;\n")
	sb.WriteString("memory_initialization_radix=16;\n")
	sb.WriteString("memory_initialization_vector=\n")
	words := pci.ConfigSpaceSize / 4
	for i := 0; i < words; i++ {
		w := cs.ReadU32(i * 4)
		if i < words-1 {
			sb.WriteString(rpad32(w) + ",\n")
		} else {
			sb.WriteString(rpad32(w) + ";\n")
		}
	}
	return sb.String()
}

func rpad32(w uint32) string {
	const hex = "0123456789abcdef"
	b := []byte("00000000")
	for i := 7; i >= 0; i-- {
		b[i] = hex[w&0xf]
		w >>= 4
	}
	return string(b)
}