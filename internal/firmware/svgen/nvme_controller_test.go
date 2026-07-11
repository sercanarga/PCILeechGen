package svgen

import (
	"strings"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/firmware/nvme"
)

// TestGenerateBarControllerSV_NVMeLatencyBypassUnconditional guards the Code-10
// fix: NVMe BAR0 reads must bypass the latency emulator regardless of MSI-X,
// or a failed MSI-X capture routes CSTS reads through the emulator (stale
// CSTS=0 -> stornvme fails ControllerEnable -> Code 10).
func TestGenerateBarControllerSV_NVMeLatencyBypassUnconditional(t *testing.T) {
	for _, msix := range []bool{true, false} {
		cfg := testConfig()
		cfg.LatencyConfig = DefaultLatencyConfig(cfg.ClassCode)
		cfg.NVMeIdentify = &nvme.IdentifyData{}
		if msix {
			cfg.HasMSIX = true
			cfg.MSIXConfig = &MSIXConfig{NumVectors: 9, TableOffset: 0x2000, PBAOffset: 0x2100}
		}
		result, err := GenerateBarControllerSV(cfg)
		if err != nil {
			t.Fatalf("msix=%v: %v", msix, err)
		}
		if !strings.Contains(result, "assign bar0_base_ctx   = bar0_raw_ctx;") {
			t.Fatalf("msix=%v: NVMe BAR0 latency bypass must be present", msix)
		}
		if strings.Contains(result, "bar_rsp_valid[0]   = bar0_emu_valid") {
			t.Fatalf("msix=%v: NVMe bar_rsp[0] must not route through the latency emulator", msix)
		}
	}
}

// TestGenerateBarControllerSV_WiresNVMeDoorbellsAndDisk verifies the bar
// controller folds the four doorbell strobes into the responder's generalized
// doorbell interface and instantiates the BRAM disk cache on disk_req_*.
func TestGenerateBarControllerSV_WiresNVMeDoorbellsAndDisk(t *testing.T) {
	cfg := testConfig()
	cfg.NVMeIdentify = &nvme.IdentifyData{}
	cfg.NVMeDoorbellStride = 0
	cfg.NVMeDiskWords = 8192

	result, err := GenerateBarControllerSV(cfg)
	if err != nil {
		t.Fatalf("GenerateBarControllerSV failed: %v", err)
	}

	for _, want := range []string{
		// range-based doorbell decode
		"nvme_db_base",
		"nvme_db_wr",
		"nvme_db_is_cq",
		"nvme_db_qid",
		// CC.EN explicit pulses
		"nvme_cc_enable_wr",
		".cc_enable_wr",
		// generalized doorbell interface to responder
		".doorbell_wr",
		".doorbell_is_cq",
		".doorbell_qid",
		// BRAM disk cache
		"i_nvme_bram_disk",
		".disk_req_valid",
		".disk_req_lba",
	} {
		if !strings.Contains(result, want) {
			t.Fatalf("NVMe bar controller should contain %q", want)
		}
	}

	// Regression for Synth 8-7136: bram_disk dropped DISK_LBAS.
	for _, mustNot := range []string{
		".DISK_LBAS",
		"DISK_LBAS",
	} {
		if strings.Contains(result, mustNot) {
			t.Fatalf("NVMe bar controller should NOT pass removed bram_disk param %q", mustNot)
		}
	}

	// Regression for Synth 8-11241/8-8895: wrengine_ready must be declared before use.
	declIdx := strings.Index(result, "wire        wrengine_ready;")
	instIdx := strings.Index(result, "i_pcileech_tlps128_bar_wrengine")
	if declIdx < 0 {
		t.Fatal("NVMe bar controller should declare wrengine_ready")
	}
	if instIdx < 0 || !(declIdx < instIdx) {
		t.Fatalf("wrengine_ready must be declared before the wrengine instance (decl=%d inst=%d)", declIdx, instIdx)
	}

	// Regression for Synth 8-11241: nvme_id_rom_data must be pre-declared before use.
	if !strings.Contains(result, "wire [31:0] nvme_id_rom_data;") {
		t.Fatal("NVMe bar controller should pre-declare wire [31:0] nvme_id_rom_data before its use")
	}

	// Regression: cfg.NVMeDiskWords must reach the SV define (board-scaled cache).
	// Was a hardcoded 32768 default, blowing the 75T BRAM budget (127 > 105 RAMB36).
	if !strings.Contains(result, "`define NVME_DISK_WORDS 8192") {
		t.Fatal("NVMe bar controller should emit `define NVME_DISK_WORDS from cfg.NVMeDiskWords")
	}
}

func TestGenerateBarControllerSV_NVMeDoorbellStride(t *testing.T) {
	cfg := testConfig()
	cfg.NVMeIdentify = &nvme.IdentifyData{}
	cfg.NVMeDoorbellStride = 1

	controller, err := GenerateBarControllerSV(cfg)
	if err != nil {
		t.Fatalf("GenerateBarControllerSV failed: %v", err)
	}
	decode := extractHDLBlock(t, controller,
		"    wire [31:0] nvme_db_base",
		"    wire [15:0] nvme_db_val")
	dut := `module nvme_doorbell_decode(
    input [31:0] wr_addr_bar0,
    input wr_valid,
    input [5:0] wr_bar,
    input bar0_wr_hit,
    output doorbell_wr,
    output doorbell_is_cq,
    output [15:0] doorbell_qid
);
` + decode + `
assign doorbell_wr = nvme_db_wr;
assign doorbell_is_cq = nvme_db_is_cq;
assign doorbell_qid = nvme_db_qid;
endmodule
`
	testbench := `module tb;
reg [31:0] wr_addr_bar0 = 0;
reg wr_valid = 1;
reg [5:0] wr_bar = 1;
reg bar0_wr_hit = 1;
wire doorbell_wr;
wire doorbell_is_cq;
wire [15:0] doorbell_qid;

nvme_doorbell_decode dut(.*);

task check_doorbell(input [31:0] addr, input [15:0] qid, input is_cq);
begin
    wr_addr_bar0 = addr;
    #1;
    if (!doorbell_wr || doorbell_qid !== qid || doorbell_is_cq !== is_cq)
        $fatal(1, "doorbell 0x%0h decoded qid=%0d cq=%0b", addr, doorbell_qid, doorbell_is_cq);
end
endtask

initial begin
    check_doorbell(32'h1000, 16'd0, 1'b0);
    check_doorbell(32'h1008, 16'd0, 1'b1);
    check_doorbell(32'h1010, 16'd1, 1'b0);
    check_doorbell(32'h1018, 16'd1, 1'b1);
    $finish;
end
endmodule
`
	runVerilatorSimulation(t, map[string]string{
		"dut.sv": dut,
		"tb.sv":  testbench,
	})
}

func TestGenerateBarControllerSV_NVMeMSIWithoutMSIX(t *testing.T) {
	cfg := testConfig()
	cfg.NVMeIdentify = &nvme.IdentifyData{}

	controller, err := GenerateBarControllerSV(cfg)
	if err != nil {
		t.Fatalf("GenerateBarControllerSV failed: %v", err)
	}
	service, err := GenerateInterruptServiceSV(cfg)
	if err != nil {
		t.Fatalf("GenerateInterruptServiceSV failed: %v", err)
	}
	instance := extractHDLBlock(t, controller,
		"    pcileech_interrupt_service #(",
		"\n\n    pcileech_nvme_admin_responder")
	intrAssign := extractHDLBlock(t, controller,
		"    assign intr_req =",
		"\n\nendmodule")
	dut := `module nvme_msi_only(
    input clk,
    input device_reset,
    input lifecycle_quiesce,
    input cfg_msi_enable,
    input nvme_irq_event_valid,
    output intr_req
);
wire nvme_irq_delivery_valid;
wire nvme_irq_delivery_ready;
wire nvme_irq_delivery_done;
wire [15:0] nvme_irq_delivery_vector;
wire nvme_intr_req;
` + instance + "\n" + intrAssign + `
endmodule
`
	testbench := `module tb;
reg clk = 0;
always #1 clk = ~clk;
reg device_reset = 1;
reg lifecycle_quiesce = 0;
reg cfg_msi_enable = 0;
reg nvme_irq_event_valid = 0;
wire intr_req;
integer i;
reg found;

nvme_msi_only dut(.*);

task cycle;
begin
    @(posedge clk);
    #1ps;
end
endtask

task wait_for_pulse;
begin
    found = 0;
    for (i = 0; i < 16; i = i + 1) begin
        cycle();
        if (intr_req) begin
            found = 1;
            i = 16;
        end
    end
    if (!found) $fatal(1, "NVMe MSI event produced no interrupt pulse");
end
endtask

initial begin
    cycle();
    cycle();
    @(negedge clk);
    device_reset = 0;
    nvme_irq_event_valid = 1;
    cycle();
    @(negedge clk);
    nvme_irq_event_valid = 0;
    repeat (6) begin
        cycle();
        if (intr_req) $fatal(1, "MSI asserted while disabled");
    end

    @(negedge clk);
    cfg_msi_enable = 1;
    wait_for_pulse();
    cycle();
    if (intr_req) $fatal(1, "acknowledged MSI remained asserted");
    repeat (6) begin
        cycle();
        if (intr_req) $fatal(1, "acknowledged MSI was redelivered");
    end

    @(negedge clk);
    lifecycle_quiesce = 1;
    nvme_irq_event_valid = 1;
    cycle();
    @(negedge clk);
    nvme_irq_event_valid = 0;
    repeat (4) begin
        cycle();
        if (intr_req) $fatal(1, "MSI asserted while quiesced");
    end
    @(negedge clk);
    lifecycle_quiesce = 0;
    wait_for_pulse();
    cycle();
    if (intr_req) $fatal(1, "resumed MSI was not acknowledged");

    @(negedge clk);
    nvme_irq_event_valid = 1;
    cycle();
    @(negedge clk);
    nvme_irq_event_valid = 0;
    device_reset = 1;
    cycle();
    @(negedge clk);
    device_reset = 0;
    repeat (8) begin
        cycle();
        if (intr_req) $fatal(1, "reset-cancelled MSI was delivered");
    end
    $finish;
end
endmodule
`
	runVerilatorSimulation(t, map[string]string{
		"pcileech_interrupt_service.sv": service,
		"dut.sv":                        dut,
		"tb.sv":                         testbench,
	})
}
