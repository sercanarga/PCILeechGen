package svgen

import (
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/firmware"
)

func TestPatchAllAcceptsNeTV2AcornAndPCIeScreamerLegacyWrappers(t *testing.T) {
	dir := t.TempDir()
	writeLegacyWrapperFixture(t, dir)

	ids := firmware.DeviceIDs{
		VendorID:       0x144D,
		DeviceID:       0xA808,
		SubsysVendorID: 0x144D,
		SubsysDeviceID: 0xA801,
		RevisionID:     0x01,
	}
	if err := NewSVPatcher(ids, dir).PatchAll(); err != nil {
		t.Fatalf("legacy 64-bit wrapper generation failed: %v", err)
	}
}

func TestPatchedCfgInterruptRequestPersistsUntilReadyAndAcknowledgesOnce(t *testing.T) {
	dir := t.TempDir()
	writeModernInterruptWrapperFixture(t, dir)
	patcher := NewSVPatcher(firmware.DeviceIDs{}, dir)
	if err := patcher.PatchAll(); err != nil {
		t.Fatalf("PatchAll: %v", err)
	}
	cfg, err := os.ReadFile(filepath.Join(dir, "pcileech_pcie_cfg_a7.sv"))
	if err != nil {
		t.Fatal(err)
	}
	testbench := "`timescale 1ns/1ps\n" + `module tb;
reg clk = 0;
always #1 clk = ~clk;
reg rst = 1;
reg generated_bar_intr_req = 0;
integer accepted = 0;
IfPCIeSignals ctx();

pcileech_pcie_cfg_a7 dut (
    .generated_bar_intr_req(generated_bar_intr_req),
    .rst(rst),
    .clk(clk),
    .clk_100(clk),
    .clk_pcie(clk),
    .ctx(ctx)
);

always @(posedge clk) begin
    if (!rst && ctx.cfg_interrupt && ctx.cfg_interrupt_rdy)
        accepted <= accepted + 1;
end

task cycle;
begin
    @(posedge clk);
    #1ps;
end
endtask

initial begin
    ctx.cfg_interrupt_rdy = 0;
    cycle();
    cycle();
    @(negedge clk);
    rst = 0;
    generated_bar_intr_req = 1;
    cycle();
    @(negedge clk);
    generated_bar_intr_req = 0;
    cycle();
    if (!ctx.cfg_interrupt || accepted !== 0) $fatal(1, "request was lost while ready was low");
    cycle();
    cycle();
    if (!ctx.cfg_interrupt || accepted !== 0) $fatal(1, "request did not remain pending");
    @(negedge clk);
    ctx.cfg_interrupt_rdy = 1;
    cycle();
    if (ctx.cfg_interrupt || accepted !== 1) $fatal(1, "acknowledgment did not consume exactly one request");
    cycle();
    cycle();
    if (ctx.cfg_interrupt || accepted !== 1) $fatal(1, "consumed request was emitted more than once");
    @(negedge clk);
    ctx.cfg_interrupt_rdy = 0;
    generated_bar_intr_req = 1;
    cycle();
    @(negedge clk);
    generated_bar_intr_req = 0;
    cycle();
    if (!ctx.cfg_interrupt) $fatal(1, "second request was not retained");
    @(negedge clk);
    ctx.cfg_interrupt_rdy = 1;
    generated_bar_intr_req = 1;
    cycle();
    if (!ctx.cfg_interrupt || accepted !== 2) $fatal(1, "request arriving with acknowledgment was lost");
    @(negedge clk);
    generated_bar_intr_req = 0;
    cycle();
    if (ctx.cfg_interrupt || accepted !== 3) $fatal(1, "overlapping requests did not complete exactly once");
    cycle();
    if (ctx.cfg_interrupt || accepted !== 3) $fatal(1, "overlapping request was emitted more than once");
    @(negedge clk);
    ctx.cfg_interrupt_rdy = 0;
    generated_bar_intr_req = 1;
    cycle();
    @(negedge clk);
    generated_bar_intr_req = 0;
    cycle();
    if (!ctx.cfg_interrupt) $fatal(1, "reset test request was not retained");
    @(negedge clk);
    rst = 1;
    cycle();
    if (ctx.cfg_interrupt || accepted !== 3) $fatal(1, "reset did not cancel pending request");
    @(negedge clk);
    rst = 0;
    ctx.cfg_interrupt_rdy = 1;
    cycle();
    if (ctx.cfg_interrupt || accepted !== 3) $fatal(1, "cancelled request reappeared after reset");
    $finish;
end
endmodule
`
	runVerilatorSimulation(t, map[string]string{
		"pcileech_pcie_cfg_a7.sv": string(cfg),
		"tb.sv":                   testbench,
	})
}

func TestGeneratedInterruptServiceHoldsValidAndVectorUntilReady(t *testing.T) {
	sv, err := GenerateInterruptServiceSV(&SVGeneratorConfig{})
	if err != nil {
		t.Fatalf("GenerateInterruptServiceSV: %v", err)
	}
	testbench := "`timescale 1ns/1ps\n" + `module tb;
reg clk = 0;
always #1 clk = ~clk;
reg rst = 1;
reg quiesce = 0;
reg msix_mode = 1;
reg function_enable = 1;
reg function_mask = 0;
reg event_valid = 0;
reg [15:0] event_vector = 0;
wire [15:0] vector_select;
reg vector_masked = 0;
reg delivery_ready = 0;
wire delivery_valid;
wire [15:0] delivery_vector;
wire msi_pulse;
wire pba_set_valid;
wire pba_clear_valid;
wire [15:0] pba_vector;
integer accepted = 0;
integer msi_accepted = 0;
integer i;
reg found = 0;

pcileech_interrupt_service #(
    .NUM_VECTORS(4)
) dut (
    .clk(clk),
    .rst(rst),
    .quiesce(quiesce),
    .msix_mode(msix_mode),
    .function_enable(function_enable),
    .function_mask(function_mask),
    .event_valid(event_valid),
    .event_vector(event_vector),
    .vector_select(vector_select),
    .vector_masked(vector_masked),
    .delivery_ready(delivery_ready),
    .delivery_valid(delivery_valid),
    .delivery_vector(delivery_vector),
    .msi_pulse(msi_pulse),
    .pba_set_valid(pba_set_valid),
    .pba_clear_valid(pba_clear_valid),
    .pba_vector(pba_vector)
);

always @(posedge clk) begin
    if (!rst && delivery_valid && delivery_ready)
        accepted <= accepted + 1;
    if (!rst && msi_pulse && delivery_ready)
        msi_accepted <= msi_accepted + 1;
end

task cycle;
begin
    @(posedge clk);
    #1ps;
end
endtask

task wait_for_delivery;
begin
    found = 0;
    for (i = 0; i < 32; i = i + 1) begin
        cycle();
        if (delivery_valid) begin
            found = 1;
            i = 32;
        end
    end
    if (!found) $fatal(1, "interrupt never asserted valid under backpressure");
end
endtask

initial begin
    cycle();
    cycle();
    @(negedge clk);
    rst = 0;
    event_valid = 1;
    event_vector = 16'd2;
    cycle();
    @(negedge clk);
    event_valid = 0;
    wait_for_delivery();
    if (delivery_vector !== 16'd2 || accepted !== 0) $fatal(1, "first delivery payload mismatch");
    for (i = 0; i < 4; i = i + 1) begin
        cycle();
        if (!delivery_valid || delivery_vector !== 16'd2 || accepted !== 0)
            $fatal(1, "valid or payload changed while ready was low");
    end

    @(negedge clk);
    delivery_ready = 1;
    event_valid = 1;
    event_vector = 16'd2;
    cycle();
    if (accepted !== 1 || !pba_set_valid || pba_vector !== 16'd2)
        $fatal(1, "acknowledgment collision did not preserve the new event");
    @(negedge clk);
    delivery_ready = 0;
    event_valid = 0;
    wait_for_delivery();
    if (delivery_vector !== 16'd2 || accepted !== 1) $fatal(1, "replacement delivery payload mismatch");
    for (i = 0; i < 3; i = i + 1) begin
        cycle();
        if (!delivery_valid || delivery_vector !== 16'd2 || accepted !== 1)
            $fatal(1, "replacement delivery did not remain stable");
    end
    @(negedge clk);
    delivery_ready = 1;
    cycle();
    if (accepted !== 2) $fatal(1, "replacement event was not delivered exactly once");
    @(negedge clk);
    delivery_ready = 0;
    cycle();
    cycle();
    if (delivery_valid || accepted !== 2) $fatal(1, "acknowledged event was redelivered");

    @(negedge clk);
    event_valid = 1;
    event_vector = 16'd1;
    cycle();
    @(negedge clk);
    event_valid = 0;
    wait_for_delivery();
    @(negedge clk);
    rst = 1;
    cycle();
    if (delivery_valid || accepted !== 2) $fatal(1, "reset did not cancel backpressured delivery");
    @(negedge clk);
    rst = 0;
    cycle();
    if (delivery_valid || accepted !== 2) $fatal(1, "cancelled delivery reappeared after reset");
    @(negedge clk);
    event_valid = 1;
    event_vector = 16'd3;
    cycle();
    @(negedge clk);
    event_valid = 0;
    wait_for_delivery();
    @(negedge clk);
    quiesce = 1;
    cycle();
    if (delivery_valid || accepted !== 2) $fatal(1, "quiesce exposed held MSI-X delivery");
    cycle();
    if (delivery_valid || accepted !== 2) $fatal(1, "quiesced MSI-X request was accepted");
    @(negedge clk);
    quiesce = 0;
    wait_for_delivery();
    if (delivery_vector !== 16'd3 || accepted !== 2) $fatal(1, "quiesced MSI-X request was not reissued");
    @(negedge clk);
    delivery_ready = 1;
    cycle();
    if (delivery_valid || accepted !== 3) $fatal(1, "resumed MSI-X request was not accepted exactly once");
    @(negedge clk);
    delivery_ready = 0;
    cycle();
    if (delivery_valid || accepted !== 3) $fatal(1, "resumed MSI-X request was redelivered");

    @(negedge clk);
    msix_mode = 0;
    event_valid = 1;
    event_vector = 0;
    cycle();
    @(negedge clk);
    event_valid = 0;
    found = 0;
    for (i = 0; i < 32; i = i + 1) begin
        cycle();
        if (msi_pulse) begin
            found = 1;
            i = 32;
        end
    end
    if (!found) $fatal(1, "MSI request never asserted under backpressure");
    for (i = 0; i < 3; i = i + 1) begin
        cycle();
        if (!msi_pulse || msi_accepted !== 0) $fatal(1, "MSI request did not remain held");
    end
    @(negedge clk);
    quiesce = 1;
    cycle();
    if (msi_pulse || msi_accepted !== 0) $fatal(1, "quiesce exposed held MSI request");
    cycle();
    if (msi_pulse || msi_accepted !== 0) $fatal(1, "quiesced MSI request was accepted");
    @(negedge clk);
    quiesce = 0;
    found = 0;
    for (i = 0; i < 32; i = i + 1) begin
        cycle();
        if (msi_pulse) begin
            found = 1;
            i = 32;
        end
    end
    if (!found || msi_accepted !== 0) $fatal(1, "quiesced MSI request was not reissued");
    @(negedge clk);
    delivery_ready = 1;
    cycle();
    if (msi_pulse || msi_accepted !== 1) $fatal(1, "resumed MSI request was not consumed exactly once");
    cycle();
    if (msi_pulse || msi_accepted !== 1) $fatal(1, "consumed MSI request was redelivered");
    @(negedge clk);
    delivery_ready = 0;
    event_valid = 1;
    cycle();
    @(negedge clk);
    event_valid = 0;
    found = 0;
    for (i = 0; i < 32; i = i + 1) begin
        cycle();
        if (msi_pulse) begin
            found = 1;
            i = 32;
        end
    end
    if (!found) $fatal(1, "reset test MSI was not pending");
    @(negedge clk);
    rst = 1;
    cycle();
    if (msi_pulse || msi_accepted !== 1) $fatal(1, "device reset did not clear held MSI");
    @(negedge clk);
    rst = 0;
    cycle();
    if (msi_pulse || msi_accepted !== 1) $fatal(1, "reset MSI reappeared");
    $finish;
end
endmodule
`
	runVerilatorSimulation(t, map[string]string{
		"pcileech_interrupt_service.sv": sv,
		"tb.sv":                         testbench,
	})
}

func TestGeneratedNVMeDMACancelledCompletionDoesNotPublishData(t *testing.T) {
	bridgeSV, err := GenerateNVMeDMABridgeSV(&SVGeneratorConfig{})
	if err != nil {
		t.Fatalf("GenerateNVMeDMABridgeSV: %v", err)
	}
	tagSV, err := GenerateDMATagServiceSV(&SVGeneratorConfig{})
	if err != nil {
		t.Fatalf("GenerateDMATagServiceSV: %v", err)
	}
	testbench := "`timescale 1ns/1ps\n" + `module tb;
reg clk = 0;
always #1 clk = ~clk;
reg rst = 1;
reg dma_enabled = 1;
reg [15:0] pcie_id = 0;
reg dma_wr_req = 0;
reg [63:0] dma_wr_addr = 0;
reg [31:0] dma_wr_data = 0;
reg [3:0] dma_wr_be = 0;
reg dma_wr_valid = 0;
wire dma_wr_done;
reg dma_rd_req = 0;
reg [63:0] dma_rd_addr = 64'h1000;
reg [9:0] dma_rd_len = 1;
wire dma_rd_valid;
wire [31:0] dma_rd_data;
wire dma_rd_done;
wire [127:0] tlp_tx_tdata;
wire [3:0] tlp_tx_tkeepdw;
wire tlp_tx_tvalid;
wire tlp_tx_tlast;
wire [8:0] tlp_tx_tuser;
reg tlp_tx_tready = 1;
reg [127:0] tlp_rx_tdata = 0;
reg [3:0] tlp_rx_tkeepdw = 0;
reg tlp_rx_tvalid = 0;
reg [8:0] tlp_rx_tuser = 0;
wire [31:0] dbg_status;
integer i;
reg cancelled = 0;

pcileech_nvme_dma_bridge dut (
    .rst(rst),
    .clk(clk),
    .dma_enabled(dma_enabled),
    .pcie_id(pcie_id),
    .dma_wr_req(dma_wr_req),
    .dma_wr_addr(dma_wr_addr),
    .dma_wr_data(dma_wr_data),
    .dma_wr_be(dma_wr_be),
    .dma_wr_valid(dma_wr_valid),
    .dma_wr_done(dma_wr_done),
    .dma_rd_req(dma_rd_req),
    .dma_rd_addr(dma_rd_addr),
    .dma_rd_len(dma_rd_len),
    .dma_rd_valid(dma_rd_valid),
    .dma_rd_data(dma_rd_data),
    .dma_rd_done(dma_rd_done),
    .tlp_tx_tdata(tlp_tx_tdata),
    .tlp_tx_tkeepdw(tlp_tx_tkeepdw),
    .tlp_tx_tvalid(tlp_tx_tvalid),
    .tlp_tx_tlast(tlp_tx_tlast),
    .tlp_tx_tuser(tlp_tx_tuser),
    .tlp_tx_tready(tlp_tx_tready),
    .tlp_rx_tdata(tlp_rx_tdata),
    .tlp_rx_tkeepdw(tlp_rx_tkeepdw),
    .tlp_rx_tvalid(tlp_rx_tvalid),
    .tlp_rx_tuser(tlp_rx_tuser),
    .dbg_status(dbg_status)
);

task cycle;
begin
    @(posedge clk);
    #1ps;
end
endtask

initial begin
    cycle();
    cycle();
    @(negedge clk);
    rst = 0;
    dma_rd_req = 1;
    cycle();
    if (!tlp_tx_tvalid) $fatal(1, "DMA read did not issue a PCIe request");
    @(negedge clk);
    dma_rd_req = 0;
    cycle();
    @(negedge clk);
    dma_enabled = 0;
    cycle();
    for (i = 0; i < 12; i = i + 1) begin
        cycle();
        if (dma_rd_valid) $fatal(1, "cancel published user data");
        if (dma_rd_done) begin
            cancelled = 1;
            i = 12;
        end
    end
    if (!cancelled) $fatal(1, "cancel did not notify the DMA owner");

    @(negedge clk);
    tlp_rx_tdata = 0;
    tlp_rx_tdata[31:25] = 7'b0100101;
    tlp_rx_tdata[47:45] = 3'b000;
    tlp_rx_tdata[79:72] = 8'h10;
    tlp_rx_tdata[127:96] = 32'hA5C35A7E;
    tlp_rx_tuser[0] = 1;
    tlp_rx_tkeepdw = 4'hF;
    tlp_rx_tvalid = 1;
    cycle();
    if (dma_rd_valid) $fatal(1, "tombstoned completion published user data");
    @(negedge clk);
    tlp_rx_tvalid = 0;
    tlp_rx_tuser = 0;
    for (i = 0; i < 4; i = i + 1) begin
        cycle();
        if (dma_rd_valid) $fatal(1, "retired tombstone published delayed user data");
    end
    $finish;
end
endmodule
`
	runVerilatorSimulation(t, map[string]string{
		"pcileech_dma_tag_service.sv": tagSV,
		"pcileech_nvme_dma_bridge.sv": bridgeSV,
		"tb.sv":                       testbench,
	})
}

func TestGeneratedDMATagServiceCancelRetiresTagsExactlyOnce(t *testing.T) {
	sv, err := GenerateDMATagServiceSV(&SVGeneratorConfig{})
	if err != nil {
		t.Fatalf("GenerateDMATagServiceSV: %v", err)
	}
	testbench := "`timescale 1ns/1ps\n" + `module tb;
reg clk = 0;
always #1 clk = ~clk;
reg rst = 1;
reg alloc_valid = 0;
wire alloc_ready;
wire [7:0] alloc_tag;
reg completion_valid = 0;
reg [7:0] completion_tag = 0;
reg completion_error = 0;
reg cancel_all = 0;
wire outcome_valid;
wire [7:0] outcome_tag;
wire [1:0] outcome_status;
reg outcome_ready = 1;
wire [1:0] active_tags;
wire [1:0] outstanding_count;
integer cancel_count = 0;
reg [1:0] cancel_seen = 0;
integer i;
reg retired = 0;
reg [7:0] held_outcome_tag = 0;
reg [1:0] held_outcome_status = 0;

pcileech_dma_tag_service #(
    .TAG_FIRST(8'h10),
    .TAG_COUNT(2),
    .TIMEOUT_WIDTH(8),
    .TIMEOUT_CYCLES(32)
) dut (
    .clk(clk),
    .rst(rst),
    .alloc_valid(alloc_valid),
    .alloc_ready(alloc_ready),
    .alloc_tag(alloc_tag),
    .completion_valid(completion_valid),
    .completion_tag(completion_tag),
    .completion_error(completion_error),
    .cancel_all(cancel_all),
    .outcome_valid(outcome_valid),
    .outcome_tag(outcome_tag),
    .outcome_status(outcome_status),
    .outcome_ready(outcome_ready),
    .active_tags(active_tags),
    .outstanding_count(outstanding_count)
);

task cycle;
begin
    @(posedge clk);
    #1ps;
end
endtask

always @(posedge clk) begin
    if (!rst && outcome_valid && outcome_ready) begin
        if (outcome_status !== 2'd3) $fatal(1, "cancel surfaced a non-cancelled owner status");
        if (outcome_tag === 8'h10) begin
            if (cancel_seen[0]) $fatal(1, "tag 10 cancellation repeated");
            cancel_seen[0] = 1;
        end else if (outcome_tag === 8'h11) begin
            if (cancel_seen[1]) $fatal(1, "tag 11 cancellation repeated");
            cancel_seen[1] = 1;
        end else begin
            $fatal(1, "cancellation reported an unowned tag");
        end
        cancel_count = cancel_count + 1;
    end
end

initial begin
    cycle();
    cycle();
    @(negedge clk);
    rst = 0;
    alloc_valid = 1;
    cycle();
    if (active_tags !== 2'b01 || outstanding_count !== 1) $fatal(1, "first allocation failed");
    @(negedge clk);
    alloc_valid = 0;
    cycle();
    @(negedge clk);
    alloc_valid = 1;
    cycle();
    if (active_tags !== 2'b11 || outstanding_count !== 2) $fatal(1, "second allocation failed");
    @(negedge clk);
    alloc_valid = 0;
    outcome_ready = 0;
    cancel_all = 1;
    cycle();
    @(negedge clk);
    cancel_all = 0;
    for (i = 0; i < 4; i = i + 1) begin
        cycle();
        if (outcome_valid)
            i = 4;
    end
    if (!outcome_valid || outcome_status !== 2'd3) $fatal(1, "cancel status was not presented under backpressure");
    held_outcome_tag = outcome_tag;
    held_outcome_status = outcome_status;
    for (i = 0; i < 4; i = i + 1) begin
        cycle();
        if (!outcome_valid || outcome_tag !== held_outcome_tag || outcome_status !== held_outcome_status)
            $fatal(1, "cancel status changed while outcome_ready was low");
        if (active_tags !== 2'b11 || outstanding_count !== 2 || alloc_ready)
            $fatal(1, "cancel freed an issued PCIe tag before retirement");
    end
    @(negedge clk);
    outcome_ready = 1;
    for (i = 0; i < 8; i = i + 1) begin
        cycle();
        if (active_tags !== 2'b11 || outstanding_count !== 2 || alloc_ready)
            $fatal(1, "accepted cancel status freed an issued PCIe tag");
        if (cancel_count == 2)
            i = 8;
    end
    if (cancel_count !== 2 || cancel_seen !== 2'b11) $fatal(1, "cancel did not report each owner exactly once");
    @(negedge clk);
    outcome_ready = 0;
    cycle();
    if (outcome_valid) $fatal(1, "cancel owner status repeated after drain");

    @(negedge clk);
    completion_valid = 1;
    completion_tag = 8'h10;
    cycle();
    if (outcome_valid) $fatal(1, "late completion emitted a second owner status");
    if (active_tags !== 2'b10 || outstanding_count !== 1) $fatal(1, "matching late completion did not retire only its tombstone");
    if (!alloc_ready || alloc_tag !== 8'h10) $fatal(1, "retired numeric tag did not return to the allocator");
    @(negedge clk);
    completion_valid = 0;

    for (i = 0; i < 40; i = i + 1) begin
        cycle();
        if (outcome_valid) $fatal(1, "cancelled timeout emitted a second owner status");
        if (!active_tags[1]) begin
            retired = 1;
            i = 40;
        end
    end
    if (!retired || active_tags !== 2'b00 || outstanding_count !== 0) $fatal(1, "timeout did not retire the remaining tombstone");
    if (!alloc_ready) $fatal(1, "timeout-retired tag remained unavailable");
    @(negedge clk);
    completion_valid = 1;
    completion_tag = 8'h11;
    cycle();
    if (outcome_valid || active_tags !== 2'b00) $fatal(1, "completion after timeout changed retired state");
    @(negedge clk);
    completion_valid = 0;
    alloc_valid = 1;
    cycle();
    if (active_tags !== 2'b01) $fatal(1, "post-retirement allocation failed");
    @(negedge clk);
    alloc_valid = 0;
    cancel_count = 0;
    cancel_seen = 0;
    outcome_ready = 1;
    cancel_all = 1;
    cycle();
    @(negedge clk);
    cancel_all = 0;
    for (i = 0; i < 6; i = i + 1) begin
        cycle();
        if (cancel_count == 1)
            i = 6;
    end
    if (cancel_count !== 1 || cancel_seen !== 2'b01 || active_tags !== 2'b01)
        $fatal(1, "single-tag cancel did not leave one tombstone");
    @(negedge clk);
    rst = 1;
    cycle();
    if (outcome_valid || active_tags !== 2'b00 || outstanding_count !== 0) $fatal(1, "FLR/reset retirement left stale owner state");
    @(negedge clk);
    rst = 0;
    cycle();
    if (outcome_valid || active_tags !== 2'b00 || outstanding_count !== 0) $fatal(1, "FLR/reset emitted a second owner status");
    $finish;
end
endmodule
`
	runVerilatorSimulation(t, map[string]string{
		"pcileech_dma_tag_service.sv": sv,
		"tb.sv":                       testbench,
	})
}

func TestGeneratedMSIXPBANewEventDominatesSameCycleDispatchClear(t *testing.T) {
	cfg := &SVGeneratorConfig{
		Bar0Size: 4096,
		HasMSIX:  true,
		MSIXConfig: &MSIXConfig{
			NumVectors:  4,
			TableOffset: 0x100,
			PBAOffset:   0x200,
		},
	}
	sv, err := GenerateMSIXTableSV(cfg)
	if err != nil {
		t.Fatalf("GenerateMSIXTableSV: %v", err)
	}
	testbench := "`timescale 1ns/1ps\n" + `module tb;
reg clk = 0;
always #1 clk = ~clk;
reg rst = 1;
reg [31:0] wr_addr = 0;
reg [31:0] wr_data = 0;
reg [3:0] wr_be = 0;
reg wr_valid = 0;
reg [87:0] rd_req_ctx = 0;
reg [31:0] rd_req_addr = 0;
reg rd_req_valid = 0;
wire [87:0] rd_rsp_ctx;
wire [31:0] rd_rsp_data;
wire rd_rsp_valid;
reg [15:0] vector_select = 0;
wire [63:0] vector_addr;
wire [31:0] vector_data;
wire vector_masked;
reg pba_set_valid = 0;
reg [15:0] pba_set_vector = 0;
reg pba_clear_valid = 0;
reg [15:0] pba_clear_vector = 0;
wire addr_hit;

pcileech_msix_table dut (
    .rst(rst),
    .clk(clk),
    .wr_addr(wr_addr),
    .wr_data(wr_data),
    .wr_be(wr_be),
    .wr_valid(wr_valid),
    .wr_table_select(1'b0),
    .wr_pba_select(1'b0),
    .rd_req_ctx(rd_req_ctx),
    .rd_req_addr(rd_req_addr),
    .rd_req_valid(rd_req_valid),
    .rd_table_select(1'b0),
    .rd_pba_select(1'b1),
    .rd_rsp_ctx(rd_rsp_ctx),
    .rd_rsp_data(rd_rsp_data),
    .rd_rsp_valid(rd_rsp_valid),
    .vector_select(vector_select),
    .vector_addr(vector_addr),
    .vector_data(vector_data),
    .vector_masked(vector_masked),
    .pba_set_valid(pba_set_valid),
    .pba_set_vector(pba_set_vector),
    .pba_clear_valid(pba_clear_valid),
    .pba_clear_vector(pba_clear_vector),
    .addr_hit(addr_hit)
);

task cycle;
begin
    @(posedge clk);
    #1ps;
end
endtask

task start_pba_read;
begin
    @(negedge clk);
    rd_req_addr = 32'h00000200;
    rd_req_valid = 1;
    cycle();
    @(negedge clk);
    rd_req_valid = 0;
    cycle();
    if (!rd_rsp_valid) $fatal(1, "PBA read response was not produced");
end
endtask

initial begin
    cycle();
    cycle();
    @(negedge clk);
    rst = 0;
    pba_set_valid = 1;
    pba_set_vector = 16'd2;
    pba_clear_valid = 1;
    pba_clear_vector = 16'd2;
    cycle();
    @(negedge clk);
    pba_set_valid = 0;
    pba_clear_valid = 0;
    start_pba_read();
    if (rd_rsp_data[2] !== 1'b1) $fatal(1, "same-cycle clear erased the newly pending event");

    @(negedge clk);
    pba_clear_valid = 1;
    pba_clear_vector = 16'd2;
    cycle();
    @(negedge clk);
    pba_clear_valid = 0;
    start_pba_read();
    if (rd_rsp_data[2] !== 1'b0) $fatal(1, "standalone dispatch clear did not clear PBA");
    $finish;
end
endmodule
`
	runVerilatorSimulation(t, map[string]string{
		"msix_table_init.hex":    strings.Repeat("00000000\n", 16),
		"pcileech_msix_table.sv": sv,
		"tb.sv":                  testbench,
	})
}

func TestGeneratedMSIXMaskedPendingSurvivesQuiesceAndClearsAfterResume(t *testing.T) {
	cfg := &SVGeneratorConfig{
		Bar0Size: 4096,
		HasMSIX:  true,
		MSIXConfig: &MSIXConfig{
			NumVectors:  4,
			TableOffset: 0x100,
			PBAOffset:   0x200,
		},
	}
	tableSV, err := GenerateMSIXTableSV(cfg)
	if err != nil {
		t.Fatalf("GenerateMSIXTableSV: %v", err)
	}
	interruptSV, err := GenerateInterruptServiceSV(cfg)
	if err != nil {
		t.Fatalf("GenerateInterruptServiceSV: %v", err)
	}
	testbench := "`timescale 1ns/1ps\n" + `module tb;
reg clk = 0;
always #1 clk = ~clk;
reg rst = 1;
reg quiesce = 0;
reg event_valid = 0;
reg [15:0] event_vector = 0;
wire [15:0] vector_select;
reg vector_masked = 1;
reg delivery_ready = 0;
wire delivery_valid;
wire [15:0] delivery_vector;
wire msi_pulse;
wire pba_set_valid;
wire pba_clear_valid;
wire [15:0] pba_vector;
reg [31:0] rd_req_addr = 0;
reg rd_req_valid = 0;
wire [31:0] rd_rsp_data;
wire rd_rsp_valid;
wire [63:0] vector_addr;
wire [31:0] vector_data;
wire table_vector_masked;
wire addr_hit;
integer accepted = 0;
integer i;
reg found = 0;

pcileech_interrupt_service #(
    .NUM_VECTORS(4)
) interrupts (
    .clk(clk),
    .rst(rst),
    .quiesce(quiesce),
    .msix_mode(1'b1),
    .function_enable(1'b1),
    .function_mask(1'b0),
    .event_valid(event_valid),
    .event_vector(event_vector),
    .vector_select(vector_select),
    .vector_masked(vector_masked),
    .delivery_ready(delivery_ready),
    .delivery_valid(delivery_valid),
    .delivery_vector(delivery_vector),
    .msi_pulse(msi_pulse),
    .pba_set_valid(pba_set_valid),
    .pba_clear_valid(pba_clear_valid),
    .pba_vector(pba_vector)
);

pcileech_msix_table msix_table_i (
    .rst(rst),
    .clk(clk),
    .wr_addr(32'd0),
    .wr_data(32'd0),
    .wr_be(4'd0),
    .wr_valid(1'b0),
    .wr_table_select(1'b0),
    .wr_pba_select(1'b0),
    .rd_req_ctx(88'd0),
    .rd_req_addr(rd_req_addr),
    .rd_req_valid(rd_req_valid),
    .rd_table_select(1'b0),
    .rd_pba_select(1'b1),
    .rd_rsp_ctx(),
    .rd_rsp_data(rd_rsp_data),
    .rd_rsp_valid(rd_rsp_valid),
    .vector_select(vector_select),
    .vector_addr(vector_addr),
    .vector_data(vector_data),
    .vector_masked(table_vector_masked),
    .pba_set_valid(pba_set_valid),
    .pba_set_vector(pba_vector),
    .pba_clear_valid(pba_clear_valid),
    .pba_clear_vector(pba_vector),
    .addr_hit(addr_hit)
);

always @(posedge clk) begin
    if (!rst && !quiesce && delivery_valid && delivery_ready)
        accepted <= accepted + 1;
end

task cycle;
begin
    @(posedge clk);
    #1ps;
end
endtask

task read_pba;
begin
    @(negedge clk);
    rd_req_addr = 32'h00000200;
    rd_req_valid = 1;
    cycle();
    @(negedge clk);
    rd_req_valid = 0;
    cycle();
    if (!rd_rsp_valid) $fatal(1, "PBA read response was not produced");
end
endtask

task wait_for_delivery;
begin
    found = 0;
    for (i = 0; i < 32; i = i + 1) begin
        cycle();
        if (delivery_valid) begin
            found = 1;
            i = 32;
        end
    end
    if (!found) $fatal(1, "queued MSI-X event did not resume");
end
endtask

initial begin
    cycle();
    cycle();
    @(negedge clk);
    rst = 0;
    event_valid = 1;
    event_vector = 16'd2;
    cycle();
    @(negedge clk);
    event_valid = 0;
    cycle();
    read_pba();
    if (rd_rsp_data[2] !== 1'b1 || accepted !== 0) $fatal(1, "masked event did not set PBA");

    @(negedge clk);
    quiesce = 1;
    for (i = 0; i < 4; i = i + 1) begin
        cycle();
        if (delivery_valid || accepted !== 0) $fatal(1, "quiesced masked event was delivered");
    end
    read_pba();
    if (rd_rsp_data[2] !== 1'b1) $fatal(1, "transient quiesce orphaned or cleared PBA");

    @(negedge clk);
    quiesce = 0;
    vector_masked = 0;
    wait_for_delivery();
    if (delivery_vector !== 16'd2 || accepted !== 0) $fatal(1, "resumed MSI-X payload mismatch");
    read_pba();
    if (rd_rsp_data[2] !== 1'b1) $fatal(1, "PBA cleared before resumed delivery was accepted");
    @(negedge clk);
    delivery_ready = 1;
    cycle();
    if (accepted !== 1) $fatal(1, "resumed MSI-X event was not accepted exactly once");
    @(negedge clk);
    delivery_ready = 0;
    cycle();
    read_pba();
    if (rd_rsp_data[2] !== 1'b0) $fatal(1, "accepted resumed delivery did not clear PBA");
    for (i = 0; i < 8; i = i + 1) begin
        cycle();
        if (delivery_valid || accepted !== 1) $fatal(1, "resumed MSI-X event was redelivered");
    end

    @(negedge clk);
    vector_masked = 1;
    event_valid = 1;
    event_vector = 16'd1;
    cycle();
    @(negedge clk);
    event_valid = 0;
    cycle();
    read_pba();
    if (rd_rsp_data[1] !== 1'b1) $fatal(1, "device-reset test event did not set PBA");
    @(negedge clk);
    rst = 1;
    cycle();
    @(negedge clk);
    rst = 0;
    vector_masked = 0;
    cycle();
    read_pba();
    if (rd_rsp_data[1] !== 1'b0) $fatal(1, "device reset did not clear PBA");
    for (i = 0; i < 12; i = i + 1) begin
        cycle();
        if (delivery_valid || accepted !== 1) $fatal(1, "device-reset event reappeared without PBA");
    end
    $finish;
end
endmodule
`
	runVerilatorSimulation(t, map[string]string{
		"msix_table_init.hex":           strings.Repeat("00000000\n", 16),
		"pcileech_interrupt_service.sv": interruptSV,
		"pcileech_msix_table.sv":        tableSV,
		"tb.sv":                         testbench,
	})
}

func TestGeneratedNVMeMSIXReadyQuiesceEdgeReissuesWithoutResponderHang(t *testing.T) {
	cfg := testConfig()
	cfg.Bar0Size = 4096
	cfg.HasMSIX = true
	cfg.MSIXConfig = &MSIXConfig{NumVectors: 4, TableOffset: 0x100, PBAOffset: 0x200}
	responderSV, err := GenerateNVMeResponderSV(cfg)
	if err != nil {
		t.Fatalf("GenerateNVMeResponderSV: %v", err)
	}
	bridgeSV, err := GenerateNVMeDMABridgeSV(cfg)
	if err != nil {
		t.Fatalf("GenerateNVMeDMABridgeSV: %v", err)
	}
	tagSV, err := GenerateDMATagServiceSV(cfg)
	if err != nil {
		t.Fatalf("GenerateDMATagServiceSV: %v", err)
	}
	interruptSV, err := GenerateInterruptServiceSV(cfg)
	if err != nil {
		t.Fatalf("GenerateInterruptServiceSV: %v", err)
	}
	tableSV, err := GenerateMSIXTableSV(cfg)
	if err != nil {
		t.Fatalf("GenerateMSIXTableSV: %v", err)
	}
	testbench := "`timescale 1ns/1ps\n" + `module tb;
reg clk = 0;
always #1 clk = ~clk;
reg rst = 1;
reg dma_enabled = 1;
reg quiesce = 0;
reg event_valid = 0;
reg [15:0] event_vector = 0;
wire [15:0] vector_select;
wire irq_delivery_valid;
wire [15:0] irq_delivery_vector;
wire irq_delivery_ready;
wire irq_delivery_done;
wire msi_pulse;
wire pba_set_valid;
wire pba_clear_valid;
wire [15:0] pba_vector;
reg [31:0] rd_req_addr = 0;
reg rd_req_valid = 0;
wire [31:0] rd_rsp_data;
wire rd_rsp_valid;
wire [63:0] table_vector_addr;
wire [31:0] table_vector_data;
wire table_vector_masked;
wire dma_rd_req;
wire [63:0] dma_rd_addr;
wire [9:0] dma_rd_len;
wire dma_rd_valid;
wire [31:0] dma_rd_data;
wire dma_rd_done;
wire dma_wr_req;
wire [63:0] dma_wr_addr;
wire [31:0] dma_wr_data;
wire [3:0] dma_wr_be;
wire dma_wr_valid;
wire dma_wr_done;
wire [127:0] tlp_tx_tdata;
wire [3:0] tlp_tx_tkeepdw;
wire tlp_tx_tvalid;
wire tlp_tx_tlast;
wire [8:0] tlp_tx_tuser;
wire [7:0] responder_state;
integer msix_writes = 0;
integer i;
reg found = 0;

pcileech_interrupt_service #(
    .NUM_VECTORS(4),
    .DEFER_MSIX_CLEAR(1)
) interrupts (
    .clk(clk),
    .rst(rst),
    .quiesce(quiesce),
    .msix_mode(1'b1),
    .function_enable(1'b1),
    .function_mask(1'b0),
    .event_valid(event_valid),
    .event_vector(event_vector),
    .vector_select(vector_select),
    .vector_masked(1'b0),
    .delivery_ready(irq_delivery_ready),
    .delivery_done(irq_delivery_done),
    .delivery_valid(irq_delivery_valid),
    .delivery_vector(irq_delivery_vector),
    .msi_pulse(msi_pulse),
    .pba_set_valid(pba_set_valid),
    .pba_clear_valid(pba_clear_valid),
    .pba_vector(pba_vector)
);

pcileech_msix_table msix_table_i (
    .rst(rst),
    .clk(clk),
    .wr_addr(32'd0),
    .wr_data(32'd0),
    .wr_be(4'd0),
    .wr_valid(1'b0),
    .wr_table_select(1'b0),
    .wr_pba_select(1'b0),
    .rd_req_ctx(88'd0),
    .rd_req_addr(rd_req_addr),
    .rd_req_valid(rd_req_valid),
    .rd_table_select(1'b0),
    .rd_pba_select(1'b1),
    .rd_rsp_ctx(),
    .rd_rsp_data(rd_rsp_data),
    .rd_rsp_valid(rd_rsp_valid),
    .vector_select(vector_select),
    .vector_addr(table_vector_addr),
    .vector_data(table_vector_data),
    .vector_masked(table_vector_masked),
    .pba_set_valid(pba_set_valid),
    .pba_set_vector(pba_vector),
    .pba_clear_valid(pba_clear_valid),
    .pba_clear_vector(pba_vector),
    .addr_hit()
);

pcileech_nvme_admin_responder responder (
    .rst(rst),
    .clk(clk),
    .dma_enabled(dma_enabled),
    .cc_en(1'b1),
    .cc_enable_wr(1'b0),
    .cc_disable_wr(1'b0),
    .asq_lo(32'd0),
    .asq_hi(32'd0),
    .acq_lo(32'd0),
    .acq_hi(32'd0),
    .aqa(32'd0),
    .doorbell_wr(1'b0),
    .doorbell_is_cq(1'b0),
    .doorbell_qid(16'd0),
    .doorbell_val(16'd0),
    .msix_vector_addr(64'h00000000FEE00000),
    .msix_vector_data(32'h00000044),
    .irq_delivery_valid(irq_delivery_valid),
    .irq_delivery_ready(irq_delivery_ready),
    .irq_delivery_done(irq_delivery_done),
    .dma_rd_req(dma_rd_req),
    .dma_rd_addr(dma_rd_addr),
    .dma_rd_len(dma_rd_len),
    .dma_rd_valid(dma_rd_valid),
    .dma_rd_data(dma_rd_data),
    .dma_rd_done(dma_rd_done),
    .dma_wr_req(dma_wr_req),
    .dma_wr_addr(dma_wr_addr),
    .dma_wr_data(dma_wr_data),
    .dma_wr_be(dma_wr_be),
    .dma_wr_valid(dma_wr_valid),
    .dma_wr_done(dma_wr_done),
    .disk_req_valid(),
    .disk_req_write(),
    .disk_req_flush(),
    .disk_req_lba(),
    .disk_req_word(),
    .disk_req_wdata(),
    .disk_req_done(1'b0),
    .disk_req_rdata(32'd0),
    .disk_req_hit(1'b0),
    .disk_busy(1'b0),
    .disk_error(1'b0),
    .msix_trigger(),
    .pba_set_valid(),
    .pba_set_vector(),
    .id_rom_addr(),
    .id_rom_data(32'd0),
    .dbg_state(responder_state),
    .dbg_active_qid(),
    .dbg_opcode(),
    .dbg_admin_queues(),
    .dbg_cmd_info()
);

pcileech_nvme_dma_bridge bridge (
    .rst(rst),
    .clk(clk),
    .dma_enabled(dma_enabled),
    .pcie_id(16'd0),
    .dma_wr_req(dma_wr_req),
    .dma_wr_addr(dma_wr_addr),
    .dma_wr_data(dma_wr_data),
    .dma_wr_be(dma_wr_be),
    .dma_wr_valid(dma_wr_valid),
    .dma_wr_done(dma_wr_done),
    .dma_rd_req(dma_rd_req),
    .dma_rd_addr(dma_rd_addr),
    .dma_rd_len(dma_rd_len),
    .dma_rd_valid(dma_rd_valid),
    .dma_rd_data(dma_rd_data),
    .dma_rd_done(dma_rd_done),
    .tlp_tx_tdata(tlp_tx_tdata),
    .tlp_tx_tkeepdw(tlp_tx_tkeepdw),
    .tlp_tx_tvalid(tlp_tx_tvalid),
    .tlp_tx_tlast(tlp_tx_tlast),
    .tlp_tx_tuser(tlp_tx_tuser),
    .tlp_tx_tready(1'b1),
    .tlp_rx_tdata(128'd0),
    .tlp_rx_tkeepdw(4'd0),
    .tlp_rx_tvalid(1'b0),
    .tlp_rx_tuser(9'd0),
    .dbg_status()
);

always @(posedge clk) begin
    if (!rst && tlp_tx_tvalid && tlp_tx_tlast)
        msix_writes <= msix_writes + 1;
end

task cycle;
begin
    @(posedge clk);
    #1ps;
end
endtask

task read_pba;
begin
    @(negedge clk);
    rd_req_addr = 32'h00000200;
    rd_req_valid = 1;
    cycle();
    @(negedge clk);
    rd_req_valid = 0;
    cycle();
    if (!rd_rsp_valid) $fatal(1, "PBA read response was not produced");
end
endtask

initial begin
    cycle();
    cycle();
    @(negedge clk);
    rst = 0;
    event_valid = 1;
    event_vector = 0;
    cycle();
    @(negedge clk);
    event_valid = 0;
    found = 0;
    for (i = 0; i < 32; i = i + 1) begin
        cycle();
        if (irq_delivery_valid && irq_delivery_ready) begin
            found = 1;
            i = 32;
        end
    end
    if (!found || responder_state !== 8'd0 || msix_writes !== 0)
        $fatal(1, "test did not reach ready-high held delivery");

    cycle();
    if (responder_state === 8'd0 || !dma_wr_valid)
        $fatal(1, "ready delivery was not accepted by the responder");
    quiesce = 1;
    dma_enabled = 0;
    cycle();
    if (irq_delivery_valid || irq_delivery_ready || dma_wr_valid || responder_state !== 8'd0 || msix_writes !== 0)
        $fatal(1, "BME drop failed to reset the accepted delivery");
    for (i = 0; i < 3; i = i + 1) begin
        cycle();
        if (dma_wr_valid || responder_state !== 8'd0 || msix_writes !== 0)
            $fatal(1, "disabled DMA bridge/responder did not remain idle");
    end
    read_pba();
    if (rd_rsp_data[0] !== 1'b1) $fatal(1, "quiesce edge lost pending PBA");

    @(negedge clk);
    quiesce = 0;
    dma_enabled = 1;
    found = 0;
    for (i = 0; i < 64; i = i + 1) begin
        cycle();
        if (msix_writes == 1) begin
            found = 1;
            i = 64;
        end
    end
    if (!found) $fatal(1, "resumed event did not complete an MSI-X write");
    for (i = 0; i < 16; i = i + 1) begin
        cycle();
        if (responder_state === 8'd0 && !dma_wr_valid)
            i = 16;
    end
    if (responder_state !== 8'd0 || dma_wr_valid) $fatal(1, "responder remained stuck waiting for DMA write");
    cycle();
    cycle();
    read_pba();
    if (rd_rsp_data[0] !== 1'b0) $fatal(1, "completed resumed MSI-X write did not clear PBA");
    for (i = 0; i < 16; i = i + 1) begin
        cycle();
        if (msix_writes !== 1 || irq_delivery_valid) $fatal(1, "resumed event duplicated or remained pending");
    end
    $finish;
end
endmodule
`
	runVerilatorSimulation(t, map[string]string{
		"msix_table_init.hex":              strings.Repeat("00000000\n", 16),
		"pcileech_dma_tag_service.sv":      tagSV,
		"pcileech_interrupt_service.sv":    interruptSV,
		"pcileech_msix_table.sv":           tableSV,
		"pcileech_nvme_admin_responder.sv": responderSV,
		"pcileech_nvme_dma_bridge.sv":      bridgeSV,
		"tb.sv":                            testbench,
	})
}

func runVerilatorSimulation(t *testing.T, sources map[string]string) {
	t.Helper()
	verilator, err := exec.LookPath("verilator")
	if err != nil {
		t.Skip("verilator not installed")
	}
	dir := t.TempDir()
	var names []string
	for name := range sources {
		names = append(names, name)
	}
	sort.Strings(names)
	var paths []string
	for _, name := range names {
		path := filepath.Join(dir, name)
		if err := os.WriteFile(path, []byte(sources[name]), 0644); err != nil {
			t.Fatalf("writing %s: %v", name, err)
		}
		if strings.HasSuffix(name, ".sv") {
			paths = append(paths, path)
		}
	}
	objDir := filepath.Join(dir, "obj")
	args := []string{"--binary", "--timing", "-Wno-fatal", "--top-module", "tb", "-Mdir", objDir, "-o", "sim"}
	args = append(args, paths...)
	build := exec.Command(verilator, args...)
	build.Dir = dir
	if output, err := build.CombinedOutput(); err != nil {
		t.Fatalf("verilator build failed: %v\n%s", err, output)
	}
	sim := exec.Command(filepath.Join(objDir, "sim"))
	sim.Dir = dir
	if output, err := sim.CombinedOutput(); err != nil {
		t.Fatalf("RTL simulation failed: %v\n%s", err, output)
	}
}

func writeLegacyWrapperFixture(t *testing.T, dir string) {
	t.Helper()
	fixtures := map[string]string{
		"pcileech_pcie_cfg_a7.sv": `module pcileech_pcie_cfg_a7(
    input clk
);
assign ctx.cfg_interrupt = rw[206];
endmodule`,
		"pcileech_pcie_tlp_a7.sv": `module pcileech_pcie_tlp_a7(
    input rst,
    input clk_100,
    input clk_pcie,
    IfPCIeFifoTlp.mp_pcie dfifo,
    IfTlp64.sink tlp_static
);
endmodule`,
		"pcileech_pcie_a7.sv": `module pcileech_pcie_a7();
IfPCIeSignals ctx();
pcileech_pcie_cfg_a7 i_cfg(
    .ctx(ctx)
);
pcileech_pcie_tlp_a7 i_tlp(
    .clk_pcie(clk_pcie)
);
endmodule`,
		"pcileech_fifo.sv": `module pcileech_fifo(input clk);
    rw[143:128] <= 16'h0000; // CFG_SUBSYS_VEND_ID
    rw[159:144] <= 16'h0000; // CFG_SUBSYS_ID
    rw[175:160] <= 16'h0000; // CFG_VEND_ID
    rw[191:176] <= 16'h0000; // CFG_DEV_ID
    rw[199:192] <= 8'h00; // CFG_REV_ID
    rw[203] <= 1'b1; // CFGTLP ZERO DATA
endmodule`,
	}
	for name, fixture := range fixtures {
		if err := os.WriteFile(filepath.Join(dir, name), []byte(fixture), 0644); err != nil {
			t.Fatalf("writing %s: %v", name, err)
		}
	}
}

func writeModernInterruptWrapperFixture(t *testing.T, dir string) {
	t.Helper()
	fixtures := map[string]string{
		"pcileech_pcie_cfg_a7.sv": `interface IfPCIeSignals;
logic cfg_interrupt;
logic cfg_interrupt_rdy;
endinterface

module pcileech_pcie_cfg_a7(
    input rst,
    input clk,
    input clk_100,
    input clk_pcie,
    IfPCIeSignals ctx
);
reg [255:0] rw = 0;
assign ctx.cfg_interrupt = rw[206];
endmodule`,
		"pcileech_pcie_tlp_a7.sv": `module pcileech_pcie_tlp_a7(
    input clk
);
pcileech_tlps128_bar_controller i_bar(
    .rst(rst)
);
endmodule`,
		"pcileech_pcie_a7.sv": `module pcileech_pcie_a7();
IfPCIeSignals ctx();
pcileech_pcie_cfg_a7 i_cfg(
    .ctx(ctx)
);
pcileech_pcie_tlp_a7 i_tlp(
    .clk(clk)
);
endmodule`,
		"pcileech_fifo.sv": `module pcileech_fifo(input clk);
    rw[143:128] <= 16'h0000; // CFG_SUBSYS_VEND_ID
    rw[159:144] <= 16'h0000; // CFG_SUBSYS_ID
    rw[175:160] <= 16'h0000; // CFG_VEND_ID
    rw[191:176] <= 16'h0000; // CFG_DEV_ID
    rw[199:192] <= 8'h00; // CFG_REV_ID
    rw[203] <= 1'b1; // CFGTLP ZERO DATA
endmodule`,
	}
	for name, fixture := range fixtures {
		if err := os.WriteFile(filepath.Join(dir, name), []byte(fixture), 0644); err != nil {
			t.Fatalf("writing %s: %v", name, err)
		}
	}
}
