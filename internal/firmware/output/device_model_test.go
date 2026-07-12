package output

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/firmware/devicemodel"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

func TestOutputWriterDeviceModelArtifactAndManifest(t *testing.T) {
	ctx := outputModelContext()

	outputDir := t.TempDir()
	writer := NewOutputWriter(outputDir, "unused", 1, 1)
	if err := writer.writeDeviceModel(ctx, nil, nil); err != nil {
		t.Fatalf("writeDeviceModel: %v", err)
	}

	artifactPath := filepath.Join(outputDir, "device_model.json")
	data, err := os.ReadFile(artifactPath)
	if err != nil {
		t.Fatalf("device_model.json was not written: %v", err)
	}
	model, err := devicemodel.FromJSON(data)
	if err != nil {
		t.Fatalf("device_model.json is not a valid model artifact: %v", err)
	}
	if len(model.Functions) != 1 || model.Functions[0].VendorID != 0x1234 || model.Functions[0].DeviceID != 0x5678 {
		t.Fatalf("artifact lost donor identity: %+v", model.Functions)
	}

	listed := false
	for _, name := range ListOutputFiles() {
		if name == "device_model.json" {
			listed = true
			break
		}
	}
	if !listed {
		t.Fatal("ListOutputFiles does not include device_model.json")
	}

	manifest, err := GenerateManifest(outputDir, ctx.ToolVersion, "test-board", ctx.Device.VendorID, ctx.Device.DeviceID)
	if err != nil {
		t.Fatalf("GenerateManifest: %v", err)
	}
	for _, entry := range manifest.Files {
		if entry.Name == "device_model.json" {
			if entry.Size != int64(len(data)) {
				t.Fatalf("manifest size = %d, want %d", entry.Size, len(data))
			}
			if entry.SHA256 == "" || entry.SHA256 == "error" {
				t.Fatalf("manifest has invalid device model checksum %q", entry.SHA256)
			}
			return
		}
	}
	t.Fatal("build manifest does not include device_model.json")
}

func TestOutputWriterWriteAllWithRealisticDonor(t *testing.T) {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("runtime.Caller failed")
	}
	ctx, err := donor.LoadContext(filepath.Join(filepath.Dir(file), "..", "..", "..", "testdata", "donors", "nvme.json"))
	if err != nil {
		t.Fatalf("load realistic donor fixture: %v", err)
	}

	libDir := t.TempDir()
	srcDir := filepath.Join(libDir, "test-board", "src")
	if werr := os.MkdirAll(srcDir, 0755); werr != nil {
		t.Fatal(werr)
	}
	for name, content := range map[string]string{
		"pcileech_fifo.sv": `module pcileech_fifo;
    rw[175:160] <= 16'h0000; // CFG_VEND_ID
    rw[191:176] <= 16'h0000; // CFG_DEV_ID
endmodule`,
		"pcileech_pcie_cfg_a7.sv": `module pcileech_pcie_cfg_a7; endmodule`,
		"pcileech_pcie_a7.sv":     `module pcileech_pcie_a7; endmodule`,
		"pcileech_tlps128_bar_controller.sv": `module pcileech_tlps128_bar_controller;
endmodule`,
	} {
		if werr := os.WriteFile(filepath.Join(srcDir, name), []byte(content), 0644); werr != nil {
			t.Fatalf("write board source %s: %v", name, werr)
		}
	}

	outDir := t.TempDir()
	writer := NewOutputWriter(outDir, libDir, 1, 1)
	writer.StockBar = true
	b := &board.Board{
		Name: "test-board", ProjectDir: "test-board", FPGAPart: "xc7a35tfgg484-2",
		PCIeLanes: 1, BRAMSize: 0x10000, TopModule: "test_top",
	}
	if werr := writer.WriteAll(ctx, b); werr != nil {
		t.Fatalf("WriteAll realistic donor: %v", werr)
	}
	for _, name := range []string{
		"device_context.json",
		"device_model.json",
		"pcileech_cfgspace.coe",
		"vivado_generate_project.tcl",
		"build_manifest.json",
	} {
		if _, werr := os.Stat(filepath.Join(outDir, name)); werr != nil {
			t.Errorf("WriteAll did not produce %s: %v", name, werr)
		}
	}
	data, err := os.ReadFile(filepath.Join(outDir, "device_model.json"))
	if err != nil {
		t.Fatal(err)
	}
	model, err := devicemodel.FromJSON(data)
	if err != nil {
		t.Fatalf("WriteAll device model is invalid: %v", err)
	}
	if len(model.Functions) != 1 || model.Functions[0].VendorID != ctx.Device.VendorID {
		t.Fatalf("WriteAll device model lost donor identity: %+v", model.Functions)
	}
	for _, capability := range model.Capabilities {
		if capability.ID == 0x11 {
			return
		}
	}
	t.Fatal("WriteAll device model lost final MSI-X capability")
}

func outputModelContext() *donor.DeviceContext {
	config := pci.NewConfigSpaceFromBytes(make([]byte, pci.ConfigSpaceLegacySize))
	config.WriteU16(0x00, 0x1234)
	config.WriteU16(0x02, 0x5678)
	config.WriteU8(0x0b, 0x02)
	return &donor.DeviceContext{
		CollectedAt: time.Date(2026, 1, 2, 3, 4, 5, 0, time.UTC),
		ToolVersion: "test-version",
		Hostname:    "test-host",
		Device: pci.PCIDevice{
			BDF:       pci.BDF{Domain: 0, Bus: 3, Device: 0, Function: 0},
			VendorID:  0x1234,
			DeviceID:  0x5678,
			ClassCode: 0x020000,
		},
		ConfigSpace: config,
		BARs: []pci.BAR{{
			Index: 0,
			Type:  pci.BARTypeMem32,
			Size:  0x100,
		}},
		BARContents: map[int][]byte{0: make([]byte, 0x100)},
	}
}

func validOutputDeviceModelJSON(t *testing.T) []byte {
	t.Helper()
	model, err := devicemodel.Build(outputModelContext())
	if err != nil {
		t.Fatalf("Build device model fixture: %v", err)
	}
	data, err := model.ToJSON()
	if err != nil {
		t.Fatalf("ToJSON device model fixture: %v", err)
	}
	return data
}

func TestValidateOutputDirRejectsInvalidDeviceModel(t *testing.T) {
	optional := map[string]bool{
		"pcileech_msix_table.sv":           true,
		"msix_table_init.hex":              true,
		"pcileech_nvme_admin_responder.sv": true,
		"pcileech_nvme_dma_bridge.sv":      true,
	}
	deviceConfig := "package device_config;\n" +
		"localparam HAS_NVME_RESP = 0;\n" +
		"localparam HAS_MSIX_INT = 0;\n" +
		"endpackage\n"
	tests := []struct {
		name    string
		content string
	}{
		{name: "malformed JSON", content: "not-json"},
		{name: "invalid schema", content: `{"schema_version":999}`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir := t.TempDir()
			writeOutputFixture(t, dir, optional, deviceConfig)
			if err := os.WriteFile(filepath.Join(dir, "device_model.json"), []byte(tt.content), 0644); err != nil {
				t.Fatal(err)
			}
			result := ValidateOutputDir(dir)
			found := false
			for _, failure := range result.Failed {
				if strings.Contains(failure, "device_model.json") {
					found = true
					break
				}
			}
			if !found {
				t.Fatalf("invalid device_model.json was not rejected: %v", result.Failed)
			}
		})
	}
}
