package output

import (
	"testing"

	"github.com/sercanarga/pcileechgen/internal/firmware"
)

func TestValidateSVIDs_Found(t *testing.T) {
	sv := `
module device_config;
    parameter VENDOR_ID = 16'h8086;
    parameter DEVICE_ID = 16'h1533;
endmodule
`
	ids := firmware.DeviceIDs{VendorID: 0x8086, DeviceID: 0x1533}
	issues := ValidateSVIDs(sv, ids)
	if len(issues) != 0 {
		t.Errorf("expected no issues, got %v", issues)
	}
}

func TestValidateSVIDs_Missing(t *testing.T) {
	sv := `module device_config; endmodule`
	ids := firmware.DeviceIDs{VendorID: 0x8086, DeviceID: 0x1533}
	issues := ValidateSVIDs(sv, ids)
	if len(issues) != 2 {
		t.Errorf("expected 2 issues, got %d: %v", len(issues), issues)
	}
}

func TestValidateHexFile_Valid(t *testing.T) {
	hex := "deadbeef\n12345678\nabcd0000\n"
	if err := ValidateHexFile(hex, 3); err != nil {
		t.Errorf("expected valid hex, got: %v", err)
	}
}

func TestValidateHexFile_BadChar(t *testing.T) {
	hex := "deadbXef"
	if err := ValidateHexFile(hex, 0); err == nil {
		t.Error("expected error for bad hex char")
	}
}

func TestValidateHexFile_WrongLength(t *testing.T) {
	hex := "deadbeef\n"
	if err := ValidateHexFile(hex, 10); err == nil {
		t.Error("expected word count mismatch error")
	}
}

func TestValidateCOEFile_Valid(t *testing.T) {
	coe := `memory_initialization_radix=16;
memory_initialization_vector=
deadbeef,
12345678;`
	if err := ValidateCOEFile(coe); err != nil {
		t.Errorf("expected valid COE, got: %v", err)
	}
}

func TestValidateCOEFile_MissingRadix(t *testing.T) {
	coe := `memory_initialization_vector=deadbeef;`
	if err := ValidateCOEFile(coe); err == nil {
		t.Error("expected error for missing radix")
	}
}

func TestValidateOutputDir_NonExistent(t *testing.T) {
	vr := ValidateOutputDir("/tmp/pcileechgen_nonexistent_test_dir")
	if !vr.HasFailures() {
		t.Error("expected failures for non-existent dir")
	}
}

func TestValidationResultSummary(t *testing.T) {
	vr := &ValidationResult{
		Passed: []string{"a", "b"},
		Failed: []string{"c"},
	}
	s := vr.Summary()
	if s != "2 passed, 1 failed, 0 warnings" {
		t.Errorf("unexpected summary: %s", s)
	}
}
