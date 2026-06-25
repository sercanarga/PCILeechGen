package devclass

import (
	"testing"

	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

// TestValidate_ProfileWarnings is the one runnable check for the warning-only
// profile validator. It feeds a known-good donor context (zero warnings) and a
// mismatched donor context (class/BAR/cap/unsupported-behavior warnings) and
// asserts the validator returns warnings only — never an error, since the
// function signature carries no error return, proving non-blocking semantics.
func TestValidate_ProfileWarnings(t *testing.T) {
	tests := []struct {
		name    string
		ctx     *donor.DeviceContext
		wantCodes map[string]bool
	}{
		{
			name: "known-good NVMe donor yields zero warnings",
			ctx: &donor.DeviceContext{
				Device: pci.PCIDevice{
					VendorID:  0x8086,
					DeviceID:  0xABCD,
					ClassCode: 0x010802,
				},
				BARs: []pci.BAR{{
					Index: 0, Type: pci.BARTypeMem64, Size: 4096, Is64Bit: true,
				}},
				Capabilities: []pci.Capability{
					{ID: pci.CapIDPowerManagement},
					{ID: pci.CapIDMSIX},
					{ID: pci.CapIDPCIExpress},
				},
				MSIXData: &donor.MSIXData{TableSize: 2},
			},
			wantCodes: map[string]bool{},
		},
		{
			name: "mismatched NVMe donor yields warnings only",
			ctx: &donor.DeviceContext{
				Device: pci.PCIDevice{
					VendorID:  0x8086,
					DeviceID:  0xABCD,
					ClassCode: 0x010802,
				},
				// Wrong BAR type + undersized.
				BARs: []pci.BAR{{
					Index: 0, Type: pci.BARTypeMem32, Size: 512,
				}},
				// Missing MSI-X capability (PM + PCIe only).
				Capabilities: []pci.Capability{
					{ID: pci.CapIDPowerManagement},
					{ID: pci.CapIDPCIExpress},
				},
				// ATS advertised but unsupported by codegen.
				ExtCapabilities: []pci.ExtCapability{
					{ID: pci.ExtCapIDATS},
				},
			},
			wantCodes: map[string]bool{
				"profile.bar.type":             true,
				"profile.bar.size":             true,
				"profile.cap.missing":          true,
				"profile.msix.missing":         true,
				"profile.unsupported.behavior": true,
			},
		},
		{
			name: "unknown class skips profile validation",
			ctx: &donor.DeviceContext{
				Device: pci.PCIDevice{
					ClassCode: 0x7F0000, // unassigned class -> generic fallback
				},
				BARs: []pci.BAR{{Index: 0, Type: pci.BARTypeMem32, Size: 16}},
			},
			wantCodes: map[string]bool{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ws := Validate(tc.ctx)
			got := make(map[string]bool, len(ws))
			for _, w := range ws {
				if w.Severity != "warning" {
					t.Errorf("%s: warning %q has severity %q, want \"warning\"", tc.name, w.Code, w.Severity)
				}
				if w.Code == "" || w.Message == "" {
					t.Errorf("%s: warning has empty Code or Message: %+v", tc.name, w)
				}
				got[w.Code] = true
			}
			for code := range got {
				if !tc.wantCodes[code] {
					t.Errorf("%s: unexpected warning code %q", tc.name, code)
				}
			}
			for code := range tc.wantCodes {
				if !got[code] {
					t.Errorf("%s: missing expected warning code %q (got %v)", tc.name, code, codesOf(ws))
				}
			}
		})
	}
}

func codesOf(ws []Warning) []string {
	out := make([]string, 0, len(ws))
	for _, w := range ws {
		out = append(out, w.Code)
	}
	return out
}

// TestExpectationForDevice_NilForGeneric confirms unknown devices get no
// profile (skipped, not failed).
func TestExpectationForDevice_NilForGeneric(t *testing.T) {
	if e := ExpectationForDevice(0x7F0000, 0, 0); e != nil {
		t.Errorf("expected nil expectation for unknown class, got %+v", e)
	}
}

// TestExpectationForDevice_SchemaVersion confirms the version is stamped on
// every derived expectation so callers can detect schema drift.
func TestExpectationForDevice_SchemaVersion(t *testing.T) {
	e := ExpectationForDevice(0x010802, 0x8086, 0)
	if e == nil || e.SchemaVersion != SchemaVersion {
		t.Fatalf("expected SchemaVersion=%d, got %+v", SchemaVersion, e)
	}
}