package svgen_test

import (
	"strings"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/firmware"
	"github.com/sercanarga/pcileechgen/internal/firmware/barmodel"
	"github.com/sercanarga/pcileechgen/internal/firmware/svgen"
)

func hdaTestBARModel() *barmodel.BARModel {
	return &barmodel.BARModel{
		Size: 4096,
		Registers: []barmodel.BARRegister{
			{Offset: 0x00, Width: 4, Name: "GCAP_VMIN_VMAJ", Reset: 0x01004401, RWMask: 0x00000000},
			{Offset: 0x08, Width: 4, Name: "GCTL", Reset: 0x00000001, RWMask: 0x00000103},
			{Offset: 0x0C, Width: 4, Name: "WAKEEN_STATESTS", Reset: 0x00010000, RWMask: 0x0000FFFF},
			{Offset: 0x20, Width: 4, Name: "INTCTL", Reset: 0x00000000, RWMask: 0xC00000FF},
			{Offset: 0x24, Width: 4, Name: "INTSTS", Reset: 0x00000000, RWMask: 0x00000000},
			{Offset: 0x40, Width: 4, Name: "CORBLBASE", Reset: 0x00000000, RWMask: 0xFFFFFF80},
			{Offset: 0x44, Width: 4, Name: "CORBUBASE", Reset: 0x00000000, RWMask: 0xFFFFFFFF},
			{Offset: 0x48, Width: 4, Name: "CORBWP_CORBRP", Reset: 0x00000000, RWMask: 0x80FF00FF},
			{Offset: 0x4C, Width: 4, Name: "CORBCTL_STS_SIZE", Reset: 0x00420000, RWMask: 0x00000082, IsRW1C: true},
			{Offset: 0x50, Width: 4, Name: "RIRBLBASE", Reset: 0x00000000, RWMask: 0xFFFFFF80},
			{Offset: 0x54, Width: 4, Name: "RIRBUBASE", Reset: 0x00000000, RWMask: 0xFFFFFFFF},
			{Offset: 0x58, Width: 4, Name: "RIRBWP_RINTCNT", Reset: 0x00000000, RWMask: 0x800000FF},
			{Offset: 0x5C, Width: 4, Name: "RIRBCTL_STS_SIZE", Reset: 0x00420000, RWMask: 0x00000307, IsRW1C: true},
			{Offset: 0x60, Width: 4, Name: "RIRBINTSTS", Reset: 0x00000000, RWMask: 0x00000001},
			{Offset: 0x64, Width: 4, Name: "IC", Reset: 0x00000000, RWMask: 0xFFFFFFFF},
			{Offset: 0x68, Width: 4, Name: "IR", Reset: 0x00000000, RWMask: 0x00000000},
			{Offset: 0x70, Width: 4, Name: "RIRBRESP_LO", Reset: 0x00000000, RWMask: 0x00000000},
			{Offset: 0x78, Width: 4, Name: "RIRBRESP_HI", Reset: 0x00000000, RWMask: 0x00000000},
		},
	}
}

func TestHDADMABridgeIntegration(t *testing.T) {
	barModel := hdaTestBARModel()

	cfg := &svgen.SVGeneratorConfig{
		DeviceIDs: firmware.DeviceIDs{
			VendorID:   0x1102,
			DeviceID:   0x0012,
			RevisionID: 0x03,
		},
		BARModel:    barModel,
		DeviceClass: "audio",
		ClassCode:   0x040300,
		PRNGSeeds:   [4]uint32{0x12345678, 0x9ABCDEF0, 0xFEDCBA98, 0x76543210},
	}

	// Check bar_impl_device for DMA interface
	barSV, err := svgen.GenerateBarImplDeviceSV(cfg)
	if err != nil {
		t.Fatalf("GenerateBarImplDeviceSV: %v", err)
	}

	for _, c := range []string{
		"dma_req_valid", "dma_resp_lo", "dma_resp_hi",
		"dma_resp_wp", "dma_rirb_lbase", "dma_rirb_ubase",
		"dma_req_ready", "dma_done", "dma_req_valid_r",
	} {
		if !strings.Contains(barSV, c) {
			t.Errorf("bar_impl_device missing: %s", c)
		}
	}

	// Check bar_controller for DMA bridge wiring
	ctrlSV, err := svgen.GenerateBarControllerSV(cfg)
	if err != nil {
		t.Fatalf("GenerateBarControllerSV: %v", err)
	}

	for _, c := range []string{
		"pcileech_hda_rirb_dma", "hda_dma_req_valid",
		"hda_tlp_tx_tvalid", "hda_tlp_tx_tdata", "i_hda_rirb_dma",
	} {
		if !strings.Contains(ctrlSV, c) {
			t.Errorf("bar_controller missing: %s", c)
		}
	}

	// Check HDA DMA bridge
	dmaSV, err := svgen.GenerateHDARIRBDMASV(cfg)
	if err != nil {
		t.Fatalf("GenerateHDARIRBDMASV: %v", err)
	}

	for _, c := range []string{
		"pcileech_hda_rirb_dma", "D_WR_TLP1", "D_WR_TLP2",
		"D_WR_TLP3", "D_DONE", "d_is_64bit",
	} {
		if !strings.Contains(dmaSV, c) {
			t.Errorf("hda_rirb_dma missing: %s", c)
		}
	}

	t.Logf("bar_impl_device: %d bytes, bar_controller: %d bytes, hda_rirb_dma: %d bytes",
		len(barSV), len(ctrlSV), len(dmaSV))
}
