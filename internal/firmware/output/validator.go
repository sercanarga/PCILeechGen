package output

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/sercanarga/pcileechgen/internal/firmware"
	"github.com/sercanarga/pcileechgen/internal/firmware/devicemodel"
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

// ValidateOutputDir checks that all expected files exist and aren't empty.
func ValidateOutputDir(dir string) *ValidationResult {
	result := &ValidationResult{}

	for _, name := range ListOutputFiles() {
		if !outputFileRequired(dir, name) {
			continue
		}

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
		if name == "device_model.json" {
			data, readErr := os.ReadFile(path)
			if readErr != nil {
				result.fail(fmt.Sprintf("%s: unreadable: %v", name, readErr))
				continue
			}
			if _, parseErr := devicemodel.FromJSON(data); parseErr != nil {
				result.fail(fmt.Sprintf("%s: invalid: %v", name, parseErr))
				continue
			}
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

	nonEmpty := 0
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "//") || strings.HasPrefix(line, ";") {
			continue
		}
		if beforeComment, _, ok := strings.Cut(line, "//"); ok {
			line = strings.TrimSpace(beforeComment)
			if line == "" {
				continue
			}
		}
		nonEmpty++
		if len(line) != 8 {
			return fmt.Errorf("line %d: expected 8 hex chars, got %d (%q)", i+1, len(line), line)
		}
		for _, c := range line {
			if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')) {
				return fmt.Errorf("line %d: invalid hex character %q", i+1, string(c))
			}
		}
	}

	if expectedWords > 0 && nonEmpty != expectedWords {
		return fmt.Errorf("expected %d words, got %d", expectedWords, nonEmpty)
	}

	return nil
}

func outputFileRequired(dir string, name string) bool {
	switch name {
	case "pcileech_msix_table.sv", "msix_table_init.hex":
		return generatedLocalparamPresent(dir, "MSIX_NUM_VECTORS")
	case "pcileech_nvme_admin_responder.sv", "pcileech_nvme_dma_bridge.sv":
		return generatedFeatureEnabled(dir, "HAS_NVME_RESP")
	default:
		return true
	}
}

func generatedFeatureEnabled(dir string, feature string) bool {
	data, err := os.ReadFile(filepath.Join(dir, "device_config.sv"))
	if err != nil {
		return true
	}
	for _, line := range strings.Split(string(data), "\n") {
		fields := strings.Fields(line)
		if len(fields) >= 4 && fields[1] == feature && fields[2] == "=" {
			return strings.TrimSuffix(fields[3], ";") == "1"
		}
	}
	return false
}

func generatedLocalparamPresent(dir string, name string) bool {
	data, err := os.ReadFile(filepath.Join(dir, "device_config.sv"))
	if err != nil {
		return true
	}
	for _, line := range strings.Split(string(data), "\n") {
		fields := strings.Fields(line)
		if len(fields) >= 2 && fields[1] == name {
			return true
		}
	}
	return false
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

func ValidateBARSize(bar0Size, boardBRAM int, tableOffset uint32) []string {
	issues := []string{}
	if bar0Size > boardBRAM {
		issues = append(issues, fmt.Sprintf("Bar0Size %d exceeds board BRAM %d", bar0Size, boardBRAM))
	}
	if tableOffset > 0 && tableOffset >= uint32(bar0Size) {
		issues = append(issues, fmt.Sprintf("MSIX table offset 0x%x not within Bar0Size %d", tableOffset, bar0Size))
	}
	return issues
}
