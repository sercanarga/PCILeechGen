package svgen

import (
	"strings"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/firmware/barmodel"
	"github.com/sercanarga/pcileechgen/internal/firmware/nvme"
)

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
		"nvme_db_base",
		"nvme_db_wr",
		"nvme_db_is_cq",
		"nvme_db_qid",
		"nvme_cc_enable_wr",
		".cc_enable_wr",
		".doorbell_wr",
		".doorbell_is_cq",
		".doorbell_qid",
		"i_nvme_bram_disk",
		".disk_req_valid",
		".disk_req_lba",
	} {
		if !strings.Contains(result, want) {
			t.Fatalf("NVMe bar controller should contain %q", want)
		}
	}

	for _, mustNot := range []string{
		".DISK_LBAS",
		"DISK_LBAS",
	} {
		if strings.Contains(result, mustNot) {
			t.Fatalf("NVMe bar controller should NOT pass removed bram_disk param %q", mustNot)
		}
	}

	declIdx := strings.Index(result, "wire        wrengine_ready;")
	instIdx := strings.Index(result, "i_pcileech_tlps128_bar_wrengine")
	if declIdx < 0 {
		t.Fatal("NVMe bar controller should declare wrengine_ready")
	}
	if instIdx < 0 || !(declIdx < instIdx) {
		t.Fatalf("wrengine_ready must be declared before the wrengine instance (decl=%d inst=%d)", declIdx, instIdx)
	}

	if !strings.Contains(result, "wire [31:0] nvme_id_rom_data;") {
		t.Fatal("NVMe bar controller should pre-declare wire [31:0] nvme_id_rom_data before its use")
	}

	if !strings.Contains(result, "`define NVME_DISK_WORDS 8192") {
		t.Fatal("NVMe bar controller should emit `define NVME_DISK_WORDS from cfg.NVMeDiskWords")
	}
}

func nvmeShutdownRTL(t *testing.T) (bar, responder, controller, bridge string) {
	t.Helper()
	cfg := testConfig()
	cfg.NVMeIdentify = &nvme.IdentifyData{}
	cfg.BARModel = &barmodel.BARModel{
		BIR:  0,
		Size: 0x4000,
		Registers: []barmodel.BARRegister{
			{Offset: 0x14, Width: 4, Name: "CC", RWMask: 0x00FFFFF1},
			{Offset: 0x1C, Width: 4, Name: "CSTS", IsFSMDriven: true},
			{Offset: 0x24, Width: 4, Name: "AQA", RWMask: 0x0FFF0FFF},
			{Offset: 0x28, Width: 4, Name: "ASQ_LO", RWMask: 0xFFFFF000},
			{Offset: 0x2C, Width: 4, Name: "ASQ_HI", RWMask: 0xFFFFFFFF},
			{Offset: 0x30, Width: 4, Name: "ACQ_LO", RWMask: 0xFFFFF000},
			{Offset: 0x34, Width: 4, Name: "ACQ_HI", RWMask: 0xFFFFFFFF},
		},
	}

	var err error
	bar, err = GenerateBarImplDeviceSV(cfg)
	if err != nil {
		t.Fatalf("GenerateBarImplDeviceSV failed: %v", err)
	}
	responder, err = GenerateNVMeResponderSV(cfg)
	if err != nil {
		t.Fatalf("GenerateNVMeResponderSV failed: %v", err)
	}
	controller, err = GenerateBarControllerSV(cfg)
	if err != nil {
		t.Fatalf("GenerateBarControllerSV failed: %v", err)
	}
	bridge, err = GenerateNVMeDMABridgeSV(cfg)
	if err != nil {
		t.Fatalf("GenerateNVMeDMABridgeSV failed: %v", err)
	}
	return bar, responder, controller, bridge
}

func TestNVMeNormalShutdownTransitionsProcessingToComplete(t *testing.T) {
	bar, responder, _, _ := nvmeShutdownRTL(t)
	for _, want := range []string{
		"if (reg_0x00000014[15:14] == 2'b00)",
		"else if (nvme_shutdown_complete)",
		"reg_0x0000001C[3:2] <= 2'b10",
	} {
		if !strings.Contains(bar, want) {
			t.Fatalf("CSTS shutdown transition missing %q", want)
		}
	}
	for _, want := range []string{
		"normal_shutdown_requested",
		"disk_req_flush <= 1'b1",
		"else if (disk_req_done)",
		"shutdown_complete <= 1'b1",
	} {
		if !strings.Contains(responder, want) {
			t.Fatalf("normal shutdown terminal path missing %q", want)
		}
	}
}

func TestNVMeNormalShutdownDrainsCurrentCommand(t *testing.T) {
	_, responder, _, _ := nvmeShutdownRTL(t)
	idle := strings.Index(responder, "S_IDLE: begin")
	if idle < 0 {
		t.Fatal("generated responder has no S_IDLE state")
	}
	shutdown := strings.Index(responder[idle:], "if (normal_shutdown_requested)")
	adminFetch := strings.Index(responder[idle:], "else if ((sq_tail != sq_head)")
	if shutdown < 0 || adminFetch < 0 || shutdown >= adminFetch {
		t.Fatalf("normal shutdown must take priority only after active work returns to IDLE (idle=%d shutdown=%d fetch=%d)", idle, shutdown, adminFetch)
	}
	if !strings.Contains(responder, "Active work reaches IDLE only after its CQE") {
		t.Fatal("generated responder does not document the drain boundary")
	}
}

func TestNVMeAbruptShutdownCancelsDMA(t *testing.T) {
	_, responder, controller, bridge := nvmeShutdownRTL(t)
	for source, wants := range map[string][]string{
		"responder": {
			"assign dma_cancel = abrupt_shutdown_requested",
			"else if (abrupt_shutdown_start)",
		},
		"controller": {
			".dma_cancel         ( nvme_dma_cancel",
			".shutdown_complete  ( nvme_shutdown_complete",
		},
		"bridge": {
			".cancel_all         ( !dma_enabled || dma_cancel )",
			"if (!dma_enabled || dma_cancel)",
		},
	} {
		text := map[string]string{"responder": responder, "controller": controller, "bridge": bridge}[source]
		for _, want := range wants {
			if !strings.Contains(text, want) {
				t.Fatalf("%s abrupt-cancel wiring missing %q", source, want)
			}
		}
	}
}

func TestNVMeSHSTClearsWhenSHNReturnsToZero(t *testing.T) {
	bar, responder, _, _ := nvmeShutdownRTL(t)
	if !strings.Contains(bar, "reg_0x0000001C[3:2] <= 2'b00") {
		t.Fatal("CSTS.SHST does not clear for CC.SHN=0")
	}
	for _, want := range []string{
		"if (cc_shn == 2'b00)",
		"shutdown_complete <= 1'b0",
		"shutdown_flush_issued <= 1'b0",
	} {
		if !strings.Contains(responder, want) {
			t.Fatalf("responder SHN clear path missing %q", want)
		}
	}
}

func TestNVMeDisableWithoutShutdownCountsUnsafeShutdown(t *testing.T) {
	_, responder, _, _ := nvmeShutdownRTL(t)
	if !strings.Contains(responder, "if (cc_stop_event && !normal_shutdown_complete_epoch)") {
		t.Fatal("disable must count unsafe shutdown only before completed normal shutdown")
	}
	if !strings.Contains(responder, "normal_shutdown_complete_epoch <= 1'b1") {
		t.Fatal("normal shutdown terminal path does not record its completed epoch")
	}
}

func TestNVMeControllerFatalIsStickyUntilDisable(t *testing.T) {
	bar, responder, controller, _ := nvmeShutdownRTL(t)
	for _, want := range []string{
		"if (cc_en_falling)",
		"reg_0x0000001C[1] <= 1'b0",
		"else if (controller_fatal)",
		"reg_0x0000001C[1] <= 1'b1",
	} {
		if !strings.Contains(bar, want) {
			t.Fatalf("sticky CFS implementation missing %q", want)
		}
	}
	for _, want := range []string{
		"wire controller_fatal_event =",
		"assign controller_fatal = controller_fatal_latched || controller_fatal_event",
		"else if (controller_fatal_latched || controller_fatal_event)",
		"if (cc_en && !controller_fatal && doorbell_wr)",
	} {
		if !strings.Contains(responder, want) {
			t.Fatalf("fatal responder quiesce path missing %q", want)
		}
	}
	if !strings.Contains(controller, ".controller_fatal ( nvme_controller_fatal") {
		t.Fatal("fatal transport signal is not wired from responder to CSTS owner")
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
