package output

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/firmware"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

func makeDonorContext(vid, did uint16, classCode uint32) *donor.DeviceContext {
	cs := fakeConfigSpace(vid, did, classCode)
	return &donor.DeviceContext{
		CollectedAt: time.Now(),
		ToolVersion: "test",
		Device: pci.PCIDevice{
			VendorID:  vid,
			DeviceID:  did,
			ClassCode: classCode,
		},
		ConfigSpace:     cs,
		BARs:            []pci.BAR{{Index: 0, Size: 4096}},
		BARContents:     map[int][]byte{0: make([]byte, 4096)},
		Capabilities:    pci.ParseCapabilities(cs),
		ExtCapabilities: pci.ParseExtCapabilities(cs),
	}
}

func TestScrubAndVary(t *testing.T) {
	ctx := makeDonorContext(0x144D, 0xA808, 0x010802)
	ow := NewOutputWriter(t.TempDir(), "", 0, 0)
	ids := firmware.ExtractDeviceIDs(ctx.ConfigSpace, ctx.ExtCapabilities)

	scrubbed, entropy, _ := ow.scrubAndVary(ctx, nil, ids)
	if scrubbed == nil {
		t.Fatal("scrubbed config space should not be nil")
	}
	if entropy == 0 {
		t.Error("entropy should be non-zero")
	}
	if scrubbed.VendorID() != 0x144D {
		t.Errorf("VendorID = 0x%04x, want 0x144D", scrubbed.VendorID())
	}
}

func TestWriteConfigSpaceArtifacts(t *testing.T) {
	dir := t.TempDir()
	ctx := makeDonorContext(0x8086, 0x1533, 0x020000)
	ow := NewOutputWriter(dir, "", 0, 0)

	scrubbed := ctx.ConfigSpace.Clone()
	if err := ow.writeConfigSpaceArtifacts(ctx, scrubbed); err != nil {
		t.Fatalf("writeConfigSpaceArtifacts failed: %v", err)
	}

	for _, name := range []string{"pcileech_cfgspace.coe", "pcileech_cfgspace_writemask.coe", "pcileech_bar_zero4k.coe"} {
		path := filepath.Join(dir, name)
		data, err := os.ReadFile(path)
		if err != nil {
			t.Errorf("%s not created: %v", name, err)
			continue
		}
		if len(data) == 0 {
			t.Errorf("%s is empty", name)
		}
	}
}

func TestBuildSVConfig(t *testing.T) {
	classes := []struct {
		name      string
		vid, did  uint16
		classCode uint32
	}{
		{"NVMe", 0x144D, 0xA808, 0x010802},
		{"Ethernet", 0x8086, 0x15B7, 0x020000},
		{"xHCI", 0x8086, 0xA36D, 0x0C0330},
		{"Generic", 0x1234, 0x5678, 0xFF0000},
	}

	for _, tc := range classes {
		t.Run(tc.name, func(t *testing.T) {
			ctx := makeDonorContext(tc.vid, tc.did, tc.classCode)
			ow := NewOutputWriter(t.TempDir(), "", 0, 0)
			ids := firmware.ExtractDeviceIDs(ctx.ConfigSpace, ctx.ExtCapabilities)

			cfg := ow.buildSVConfig(ctx, ctx.ConfigSpace, ids, 0xDEAD, &board.Board{})
			if cfg == nil {
				t.Fatal("config should not be nil")
			}
			if cfg.LatencyConfig == nil {
				t.Error("LatencyConfig should not be nil")
			}
			if cfg.ClassCode != tc.classCode {
				t.Errorf("ClassCode = 0x%06x, want 0x%06x", cfg.ClassCode, tc.classCode)
			}
		})
	}
}

func TestBuildSVConfig_NVMeHasIdentify(t *testing.T) {
	ctx := makeDonorContext(0x144D, 0xA808, 0x010802)
	ow := NewOutputWriter(t.TempDir(), "", 0, 0)
	ids := firmware.ExtractDeviceIDs(ctx.ConfigSpace, ctx.ExtCapabilities)

	cfg := ow.buildSVConfig(ctx, ctx.ConfigSpace, ids, 0xBEEF, &board.Board{})
	if cfg.NVMeIdentify == nil {
		t.Error("NVMe class should have Identify data")
	}
}

func TestWriteDeviceContext(t *testing.T) {
	dir := t.TempDir()
	ctx := makeDonorContext(0x8086, 0x1533, 0x020000)
	ow := NewOutputWriter(dir, "", 0, 0)

	if err := ow.writeDeviceContext(ctx); err != nil {
		t.Fatalf("writeDeviceContext failed: %v", err)
	}

	path := filepath.Join(dir, "device_context.json")
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("device_context.json not created: %v", err)
	}
	if !strings.Contains(string(data), "8086") {
		t.Error("device context should contain vendor ID")
	}
}

func TestWriteCoreSVArtifacts(t *testing.T) {
	dir := t.TempDir()
	ctx := makeDonorContext(0x8086, 0x15B7, 0x020000)
	ow := NewOutputWriter(dir, "", 0, 0)
	ids := firmware.ExtractDeviceIDs(ctx.ConfigSpace, ctx.ExtCapabilities)

	cfg := ow.buildSVConfig(ctx, ctx.ConfigSpace, ids, 0x42, &board.Board{})
	scrubbed := ctx.ConfigSpace.Clone()

	if err := ow.writeCoreSVArtifacts(cfg, scrubbed); err != nil {
		t.Fatalf("writeCoreSVArtifacts failed: %v", err)
	}

	expected := []string{
		"pcileech_bar_impl_device.sv",
		"pcileech_tlps128_bar_controller.sv",
		"tlp_latency_emulator.sv",
		"device_config.sv",
		"config_space_init.hex",
	}
	for _, name := range expected {
		path := filepath.Join(dir, name)
		data, err := os.ReadFile(path)
		if err != nil {
			t.Errorf("%s not created: %v", name, err)
			continue
		}
		if len(data) == 0 {
			t.Errorf("%s is empty", name)
		}
	}
}

func TestWriteConditionalArtifacts_NVMe(t *testing.T) {
	dir := t.TempDir()
	ctx := makeDonorContext(0x144D, 0xA808, 0x010802)
	ow := NewOutputWriter(dir, "", 0, 0)
	ids := firmware.ExtractDeviceIDs(ctx.ConfigSpace, ctx.ExtCapabilities)

	cfg := ow.buildSVConfig(ctx, ctx.ConfigSpace, ids, 0x42, &board.Board{})

	if err := ow.writeConditionalArtifacts(cfg, ctx); err != nil {
		t.Fatalf("writeConditionalArtifacts failed: %v", err)
	}

	// NVMe should generate responder and bridge
	if cfg.NVMeIdentify != nil {
		for _, name := range []string{"pcileech_nvme_admin_responder.sv", "pcileech_nvme_dma_bridge.sv", "identify_init.hex"} {
			if _, err := os.ReadFile(filepath.Join(dir, name)); err != nil {
				t.Errorf("%s not created: %v", name, err)
			}
		}
	}
}

func TestWriteConditionalArtifacts_MSIXDonor(t *testing.T) {
	dir := t.TempDir()
	ctx := makeDonorContext(0x8086, 0x15B7, 0x020000)
	ctx.MSIXData = &donor.MSIXData{
		TableSize:   4,
		TableBIR:    0,
		TableOffset: 0x2000,
		PBABIR:      0,
		PBAOffset:   0x2040,
	}
	ow := NewOutputWriter(dir, "", 0, 0)
	ids := firmware.ExtractDeviceIDs(ctx.ConfigSpace, ctx.ExtCapabilities)

	cfg := ow.buildSVConfig(ctx, ctx.ConfigSpace, ids, 0x42, &board.Board{})

	if err := ow.writeConditionalArtifacts(cfg, ctx); err != nil {
		t.Fatalf("writeConditionalArtifacts failed: %v", err)
	}

	if cfg.MSIXConfig == nil {
		t.Fatal("MSIXConfig should not be nil with MSIXData present")
	}
	for _, name := range []string{"pcileech_msix_table.sv", "msix_table_init.hex"} {
		if _, err := os.ReadFile(filepath.Join(dir, name)); err != nil {
			t.Errorf("%s not created: %v", name, err)
		}
	}
}

func TestLogSVSummary(t *testing.T) {
	ctx := makeDonorContext(0x144D, 0xA808, 0x010802)
	ow := NewOutputWriter(t.TempDir(), "", 0, 0)
	ids := firmware.ExtractDeviceIDs(ctx.ConfigSpace, ctx.ExtCapabilities)

	cfg := ow.buildSVConfig(ctx, ctx.ConfigSpace, ids, 0x42, &board.Board{})
	// should not panic
	ow.logSVSummary(cfg)
}

func TestWriteTCLScripts(t *testing.T) {
	dir := t.TempDir()
	ctx := makeDonorContext(0x8086, 0x1533, 0x020000)
	b := &board.Board{Name: "75t"}
	ow := NewOutputWriter(dir, "/tmp/pcileech-fpga", 4, 3600)

	if err := ow.writeTCLScripts(ctx, b); err != nil {
		t.Fatalf("writeTCLScripts failed: %v", err)
	}

	for _, name := range []string{"vivado_generate_project.tcl", "vivado_build.tcl"} {
		data, err := os.ReadFile(filepath.Join(dir, name))
		if err != nil {
			t.Errorf("%s not created: %v", name, err)
			continue
		}
		if len(data) == 0 {
			t.Errorf("%s is empty", name)
		}
	}
}
