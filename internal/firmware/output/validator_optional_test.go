package output

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestValidateHexFile_GeneratedComments(t *testing.T) {
	hex := "// generated header\nDEADBEEF // [000]\n; comment\n12345678\n"
	if err := ValidateHexFile(hex, 2); err != nil {
		t.Errorf("expected generated hex comments to be ignored, got: %v", err)
	}
}

func TestValidateOutputDir_SkipsDisabledOptionalArtifacts(t *testing.T) {
	tmpDir := t.TempDir()
	optional := map[string]bool{
		"pcileech_msix_table.sv":           true,
		"msix_table_init.hex":              true,
		"pcileech_nvme_admin_responder.sv": true,
		"pcileech_nvme_dma_bridge.sv":      true,
	}
	deviceConfig := "package device_config;\n" +
		"localparam HAS_NVME_RESP    = 0;\n" +
		"localparam HAS_MSIX_INT     = 0;\n" +
		"endpackage\n"

	writeOutputFixture(t, tmpDir, optional, deviceConfig)

	result := ValidateOutputDir(tmpDir)
	if result.HasFailures() {
		t.Errorf("disabled optional artifacts should not fail validation: %v", result.Failed)
	}
}

func TestValidateOutputDir_RequiresMSIXArtifactsWhenTableConfigured(t *testing.T) {
	tmpDir := t.TempDir()
	optional := map[string]bool{
		"pcileech_msix_table.sv":           true,
		"msix_table_init.hex":              true,
		"pcileech_nvme_admin_responder.sv": true,
		"pcileech_nvme_dma_bridge.sv":      true,
	}
	deviceConfig := "package device_config;\n" +
		"localparam MSIX_NUM_VECTORS = 4;\n" +
		"localparam HAS_NVME_RESP    = 0;\n" +
		"endpackage\n"

	writeOutputFixture(t, tmpDir, optional, deviceConfig)

	result := ValidateOutputDir(tmpDir)
	if len(result.Failed) != 2 {
		t.Fatalf("expected missing MSI-X table artifacts to fail, got %v", result.Failed)
	}
}

func writeOutputFixture(t *testing.T, tmpDir string, skip map[string]bool, deviceConfig string) {
	t.Helper()
	for _, name := range ListOutputFiles() {
		if skip[name] {
			continue
		}
		if strings.HasSuffix(name, "/") {
			if err := os.MkdirAll(filepath.Join(tmpDir, name), 0755); err != nil {
				t.Fatalf("MkdirAll(%s): %v", name, err)
			}
			continue
		}
		content := []byte("content")
		if name == "device_config.sv" {
			content = []byte(deviceConfig)
		}
		if name == "device_model.json" {
			content = validOutputDeviceModelJSON(t)
		}
		if name == "emulation_report.json" {
			content = []byte(`{"schema_version":1,"vendor_id":"1234","device_id":"5678","class_code":"ff0000","support":{"family":"generic","level":"identity","validated":true}}`)
		}
		if err := os.WriteFile(filepath.Join(tmpDir, name), content, 0644); err != nil {
			t.Fatalf("WriteFile(%s): %v", name, err)
		}
	}
}
