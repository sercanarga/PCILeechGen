package output

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/sercanarga/pcileechgen/internal/firmware"
)

type ValidationResult struct {
	Passed   []string
	Failed   []string
	Warnings []string
}

func (vr *ValidationResult) HasFailures() bool {
	return len(vr.Failed) > 0
}

func (vr *ValidationResult) Summary() string {
	return fmt.Sprintf("%d passed, %d failed, %d warnings",
		len(vr.Passed), len(vr.Failed), len(vr.Warnings))
}

func (vr *ValidationResult) pass(msg string) { vr.Passed = append(vr.Passed, msg) }
func (vr *ValidationResult) fail(msg string) { vr.Failed = append(vr.Failed, msg) }
func (vr *ValidationResult) warn(msg string) { vr.Warnings = append(vr.Warnings, msg) }

// ValidateOutputDir checks that all expected files exist and aren't empty.
func ValidateOutputDir(dir string) *ValidationResult {
	result := &ValidationResult{}

	for _, name := range ListOutputFiles() {
		path := filepath.Join(dir, name)

		if strings.HasSuffix(name, "/") {
			if info, err := os.Stat(path); err != nil || !info.IsDir() {
				result.fail(fmt.Sprintf("%s: directory missing", name))
			} else {
				result.pass(fmt.Sprintf("%s: directory exists", name))
			}
			continue
		}

		info, err := os.Stat(path)
		if err != nil {
			result.fail(fmt.Sprintf("%s: missing", name))
			continue
		}
		if info.Size() == 0 {
			result.fail(fmt.Sprintf("%s: empty file", name))
			continue
		}
		result.pass(fmt.Sprintf("%s: OK (%d bytes)", name, info.Size()))
	}
	return result
}

// ValidateSVIDs checks that vendor/device IDs appear in generated SV.
func ValidateSVIDs(svContent string, ids firmware.DeviceIDs) []string {
	var issues []string

	vendorHex := fmt.Sprintf("%04X", ids.VendorID)
	deviceHex := fmt.Sprintf("%04X", ids.DeviceID)

	upper := strings.ToUpper(svContent)

	if !strings.Contains(upper, vendorHex) {
		issues = append(issues, fmt.Sprintf("VendorID 0x%s not found in SV", vendorHex))
	}
	if !strings.Contains(upper, deviceHex) {
		issues = append(issues, fmt.Sprintf("DeviceID 0x%s not found in SV", deviceHex))
	}

	return issues
}

// ValidateHexFile checks line length and hex chars in a .hex file.
func ValidateHexFile(hexContent string, expectedWords int) error {
	lines := strings.Split(strings.TrimSpace(hexContent), "\n")
	if len(lines) == 0 {
		return fmt.Errorf("hex file is empty")
	}

	for i, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if len(line) != 8 {
			return fmt.Errorf("line %d: expected 8 hex chars, got %d (%q)", i+1, len(line), line)
		}
		for _, c := range line {
			if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')) {
				return fmt.Errorf("line %d: invalid hex character %q", i+1, string(c))
			}
		}
	}

	nonEmpty := 0
	for _, l := range lines {
		if strings.TrimSpace(l) != "" {
			nonEmpty++
		}
	}

	if expectedWords > 0 && nonEmpty != expectedWords {
		return fmt.Errorf("expected %d words, got %d", expectedWords, nonEmpty)
	}

	return nil
}

// ValidateCOEFile checks radix/vector directives and trailing semicolon.
func ValidateCOEFile(content string) error {
	if !strings.Contains(content, "memory_initialization_radix") {
		return fmt.Errorf("missing memory_initialization_radix directive")
	}
	if !strings.Contains(content, "memory_initialization_vector") {
		return fmt.Errorf("missing memory_initialization_vector directive")
	}
	if !strings.HasSuffix(strings.TrimSpace(content), ";") {
		return fmt.Errorf("COE file should end with semicolon")
	}
	return nil
}
