package output

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
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

func TestValidateOutputDir_Empty(t *testing.T) {
	tmpDir := t.TempDir()
	result := ValidateOutputDir(tmpDir)
	if !result.HasFailures() {
		t.Error("Empty dir should have failures")
	}
	summary := result.Summary()
	if !strings.Contains(summary, "failed") {
		t.Error("Summary should mention failures")
	}
}

func TestValidateOutputDir_WithFiles(t *testing.T) {
	tmpDir := t.TempDir()
	for _, name := range ListOutputFiles() {
		if strings.HasSuffix(name, "/") {
			os.MkdirAll(filepath.Join(tmpDir, name), 0755)
		} else {
			os.WriteFile(filepath.Join(tmpDir, name), []byte("content"), 0644)
		}
	}

	result := ValidateOutputDir(tmpDir)
	if result.HasFailures() {
		t.Errorf("All files present, but got failures: %v", result.Failed)
	}
}

func TestValidateOutputDir_EmptyFile(t *testing.T) {
	tmpDir := t.TempDir()
	os.WriteFile(filepath.Join(tmpDir, "pcileech_cfgspace.coe"), []byte(""), 0644)

	result := ValidateOutputDir(tmpDir)
	foundEmpty := false
	for _, f := range result.Failed {
		if strings.Contains(f, "empty") {
			foundEmpty = true
		}
	}
	if !foundEmpty {
		t.Error("Should detect empty file")
	}
}

func TestValidateSVIDs(t *testing.T) {
	ids := firmware.DeviceIDs{VendorID: 0x8086, DeviceID: 0x1533}
	content := "16'h8086 // vendor\n16'h1533 // device"
	issues := ValidateSVIDs(content, ids)
	if len(issues) != 0 {
		t.Errorf("Expected no issues, got %v", issues)
	}
}

func TestValidateSVIDs_BothMissing(t *testing.T) {
	ids := firmware.DeviceIDs{VendorID: 0x8086, DeviceID: 0x1533}
	content := "some random content"
	issues := ValidateSVIDs(content, ids)
	if len(issues) != 2 {
		t.Errorf("Expected 2 issues (vendor+device), got %d", len(issues))
	}
}

func TestValidateHexFile_ValidFile(t *testing.T) {
	var lines []string
	for i := 0; i < 10; i++ {
		lines = append(lines, fmt.Sprintf("%08X", i))
	}
	hex := strings.Join(lines, "\n")
	if err := ValidateHexFile(hex, 10); err != nil {
		t.Errorf("ValidateHexFile error: %v", err)
	}
}

func TestValidateHexFile_Empty(t *testing.T) {
	// ValidateHexFile with truly empty content (after trim, no non-empty lines)
	// should pass if expectedWords=0, since there are no lines to validate
	if err := ValidateHexFile("", 10); err == nil {
		t.Error("Empty hex with expected words should fail")
	}
}

func TestValidateHexFile_BadChars(t *testing.T) {
	hex := "GGGGGGGG"
	if err := ValidateHexFile(hex, 0); err == nil {
		t.Error("Invalid hex chars should fail")
	}
}

func TestValidateHexFile_ShortLine(t *testing.T) {
	hex := "12345"
	if err := ValidateHexFile(hex, 0); err == nil {
		t.Error("Wrong line length should fail")
	}
}

func TestValidateHexFile_WrongWordCount(t *testing.T) {
	hex := "00000000\n11111111"
	if err := ValidateHexFile(hex, 5); err == nil {
		t.Error("Wrong word count should fail")
	}
}

func TestValidateCOEFile_ValidContent(t *testing.T) {
	coe := "; header\nmemory_initialization_radix=16;\nmemory_initialization_vector=\n00000000;"
	if err := ValidateCOEFile(coe); err != nil {
		t.Errorf("ValidateCOEFile error: %v", err)
	}
}

func TestValidateCOEFile_NoRadix(t *testing.T) {
	coe := "memory_initialization_vector=\n00000000;"
	if err := ValidateCOEFile(coe); err == nil {
		t.Error("Missing radix should fail")
	}
}

func TestValidateCOEFile_MissingVector(t *testing.T) {
	coe := "memory_initialization_radix=16;\n00000000;"
	if err := ValidateCOEFile(coe); err == nil {
		t.Error("Missing vector should fail")
	}
}

func TestValidateCOEFile_NoSemicolon(t *testing.T) {
	coe := "memory_initialization_radix=16;\nmemory_initialization_vector=\n00000000"
	if err := ValidateCOEFile(coe); err == nil {
		t.Error("Missing trailing semicolon should fail")
	}
}

func TestValidationResult_Summary(t *testing.T) {
	vr := &ValidationResult{
		Passed:   []string{"a", "b"},
		Failed:   []string{"c"},
		Warnings: []string{"d"},
	}
	summary := vr.Summary()
	if !strings.Contains(summary, "2 passed") {
		t.Error("Summary should show 2 passed")
	}
	if !strings.Contains(summary, "1 failed") {
		t.Error("Summary should show 1 failed")
	}
}
