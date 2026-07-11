package svgen

import (
	"strings"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/firmware/barmodel"
	"github.com/sercanarga/pcileechgen/internal/firmware/nvme"
)

func TestGenerateSharedServiceSVBehavioralContracts(t *testing.T) {
	cfg := testConfig()

	lifecycleSV, err := GenerateLifecycleServiceSV(cfg)
	if err != nil {
		t.Fatalf("GenerateLifecycleServiceSV: %v", err)
	}
	requireSVContracts(t, lifecycleSV,
		"module pcileech_lifecycle_service",
		"assign device_reset = fundamental_reset || flr_in_process;",
		"(pm_dstate == 2'b00)",
		"active && memory_space_enable",
		"active && bus_master_enable",
		"if (device_reset && !reset_d)",
	)

	tagSV, err := GenerateDMATagServiceSV(cfg)
	if err != nil {
		t.Fatalf("GenerateDMATagServiceSV: %v", err)
	}
	requireSVContracts(t, tagSV,
		"module pcileech_dma_tag_service",
		"reg [INDEX_WIDTH-1:0] alloc_cursor",
		"scan_index = (int'(alloc_cursor) + i) % TAG_COUNT",
		"completion_valid && (completion_tag >= TAG_FIRST)",
		"terminal_pending[completed_index] <= 1'b1",
		"terminal_status[completed_index] <=",
		"completion_error ? OUTCOME_ERROR : OUTCOME_COMPLETED",
		"age[i] >= TIMEOUT_CYCLES - 1'b1",
		"terminal_pending[timeout_index] <= 1'b1",
		"terminal_status[timeout_index] <= OUTCOME_TIMEOUT",
		"cancel_report_pending[i] <= 1'b1",
		"if (outcome_valid && outcome_ready_i)",
		"if (active_tags[i] && cancelled[i] && cancel_reported[i] &&",
		"active_tags[i] <= 1'b0",
	)

	interruptSV, err := GenerateInterruptServiceSV(cfg)
	if err != nil {
		t.Fatalf("GenerateInterruptServiceSV: %v", err)
	}
	requireSVContracts(t, interruptSV,
		"module pcileech_interrupt_service",
		"pending[event_vector[INDEX_WIDTH-1:0]] <= 1'b1",
		"pending[vector_select[INDEX_WIDTH-1:0]] && function_enable",
		"!function_mask) begin",
		"if (delivery_valid) begin",
		"if (delivery_ready) begin",
		"if (!(event_valid && (event_vector == delivery_vector)))",
		"pending[delivery_vector[INDEX_WIDTH-1:0]] <= 1'b0",
		"if (!msix_mode)",
		"msi_pulse <= 1'b1",
		"else if (!vector_masked)",
		"pba_set_valid <= 1'b1",
		"pba_clear_valid <= 1'b1",
	)
}

func TestGenerateBarControllerWiresSharedLifecycle(t *testing.T) {
	cfg := testConfig()
	controllerSV, err := GenerateBarControllerSV(cfg)
	if err != nil {
		t.Fatalf("GenerateBarControllerSV: %v", err)
	}
	requireSVContracts(t, controllerSV,
		"input [15:0]            cfg_command",
		"input [1:0]             cfg_power_state",
		"input                   cfg_flr_in_process",
		"pcileech_lifecycle_service i_lifecycle_service",
		".memory_space_enable ( cfg_mse",
		".bus_master_enable   ( cfg_bme",
		"wire in_is_bar      = io_enabled && bar_en",
		"if ( device_reset ) begin",
	)
}

func TestGenerateNVMeWiresSharedDMAInterruptAndPBA(t *testing.T) {
	cfg := testConfig()
	cfg.BARModel = &barmodel.BARModel{Size: 16 * 1024}
	cfg.Bar0Size = 16 * 1024
	cfg.NVMeIdentify = &nvme.IdentifyData{}
	cfg.HasMSIX = true
	cfg.MSIXConfig = &MSIXConfig{NumVectors: 4, TableOffset: 0x2000, PBAOffset: 0x2100}
	cfg.DonorCapabilities.MSIXCapOffset = 0x70

	controllerSV, err := GenerateBarControllerSV(cfg)
	if err != nil {
		t.Fatalf("GenerateBarControllerSV: %v", err)
	}
	requireSVContracts(t, controllerSV,
		"i_nvme_interrupt_service",
		".DEFER_MSIX_CLEAR  ( 1 )",
		".event_valid     ( nvme_irq_event_valid",
		".delivery_valid  ( nvme_irq_delivery_valid",
		".delivery_done   ( nvme_irq_delivery_done",
		".delivery_vector ( nvme_irq_delivery_vector",
		".pba_set_valid   ( nvme_msix_pba_set_valid",
		".pba_clear_valid ( nvme_msix_pba_clear_valid",
		".pba_clear_vector( nvme_msix_pba_vector",
		".irq_delivery_valid ( nvme_irq_delivery_valid",
		".irq_delivery_done  ( nvme_irq_delivery_done",
		".pba_set_vector     ( nvme_irq_event_vector",
		"i_nvme_dma_bridge",
		".rst                ( device_reset",
		".dma_enabled        ( dma_enabled",
	)
	if strings.Contains(controllerSV, ".msix_trigger       ( nvme_msix_trigger") {
		t.Fatal("NVMe responder and shared interrupt service must not drive the same trigger net")
	}
	if strings.Contains(controllerSV, ".msix_vector_select (") {
		t.Fatal("NVMe responder must not expose a second MSI-X vector driver")
	}

	responderSV, err := GenerateNVMeResponderSV(cfg)
	if err != nil {
		t.Fatalf("GenerateNVMeResponderSV: %v", err)
	}
	requireSVContracts(t, responderSV,
		"output reg          irq_delivery_done",
		"pba_set_vector <= active_io ? io_vector : 16'h0000",
	)
	if strings.Contains(responderSV, "msix_vector_select") {
		t.Fatal("NVMe responder retains an obsolete externally-driven MSI-X selector")
	}

	bridgeSV, err := GenerateNVMeDMABridgeSV(cfg)
	if err != nil {
		t.Fatalf("GenerateNVMeDMABridgeSV: %v", err)
	}
	requireSVContracts(t, bridgeSV,
		"input               dma_enabled",
		"pcileech_dma_tag_service",
		".cancel_all         ( !dma_enabled",
		"if (rst) begin",
		"if (!dma_enabled) begin",
		"if (dma_enabled && (bstate != B_IDLE)) begin",
	)

	msixSV, err := GenerateMSIXTableSV(cfg)
	if err != nil {
		t.Fatalf("GenerateMSIXTableSV: %v", err)
	}
	requireSVContracts(t, msixSV,
		"input               pba_clear_valid",
		"input  [15:0]       pba_clear_vector",
		"if (pba_set_valid && (pba_set_vector < NUM_VECTORS))",
		"if (pba_clear_valid && (pba_clear_vector < NUM_VECTORS))",
	)
}

func TestGenerateHDAWiresSharedDMAAndMSI(t *testing.T) {
	cfg := testConfig()
	cfg.DeviceClass = "audio"
	cfg.ClassCode = 0x040300
	cfg.BARModel = &barmodel.BARModel{Size: 4096}
	cfg.NVMeIdentify = nil
	cfg.MSIXConfig = nil
	cfg.HasMSIX = false
	cfg.MSIConfig = &MSIConfig{Enabled: true}

	controllerSV, err := GenerateBarControllerSV(cfg)
	if err != nil {
		t.Fatalf("GenerateBarControllerSV: %v", err)
	}
	requireSVContracts(t, controllerSV,
		"i_hda_rirb_dma",
		".rst            ( device_reset",
		".dma_enabled    ( dma_enabled",
		"i_hda_interrupt_service",
		".rst             ( device_reset || hda_crst_falling )",
		".event_valid     ( hda_msi_trigger",
		".msi_pulse       ( hda_intr_req",
	)

	bridgeSV, err := GenerateHDARIRBDMASV(cfg)
	if err != nil {
		t.Fatalf("GenerateHDARIRBDMASV: %v", err)
	}
	requireSVContracts(t, bridgeSV,
		"input               dma_enabled",
		"if (rst) begin",
		"dma_req_ready <= dma_enabled",
		"if (dma_enabled && dma_req_valid)",
	)
}

func requireSVContracts(t *testing.T, generated string, contracts ...string) {
	t.Helper()
	for _, contract := range contracts {
		if !strings.Contains(generated, contract) {
			t.Errorf("generated SystemVerilog missing contract %q", contract)
		}
	}
}
