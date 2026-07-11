package svgen

import (
	"os/exec"
	"testing"
)

// nvmeAdminBehaviorBench returns a Verilator bench exercising the NVMe admin responder.
func nvmeAdminBehaviorBench() string {
	return "`timescale 1ns/1ps\n" + `module tb;
reg clk = 0;
always #5 clk = ~clk;
reg rst = 1;
reg dma_enabled = 0;
reg cc_en = 0;
reg cc_enable_wr = 0;
reg cc_disable_wr = 0;
reg [31:0] asq_lo = 0;
reg [31:0] asq_hi = 0;
reg [31:0] acq_lo = 0;
reg [31:0] acq_hi = 0;
reg [31:0] aqa = 0;
reg doorbell_wr = 0;
reg doorbell_is_cq = 0;
reg [15:0] doorbell_qid = 0;
reg [15:0] doorbell_val = 0;
reg [63:0] msix_vector_addr = 0;
reg [31:0] msix_vector_data = 0;
reg irq_delivery_valid = 0;
wire irq_delivery_ready;
wire irq_delivery_done;

wire dma_rd_req;
wire [63:0] dma_rd_addr;
wire [9:0] dma_rd_len;
reg  dma_rd_valid = 0;
reg  [31:0] dma_rd_data = 0;
reg  dma_rd_done = 0;

wire dma_wr_req;
wire [63:0] dma_wr_addr;
wire [31:0] dma_wr_data;
wire [3:0] dma_wr_be;
wire dma_wr_valid;
wire dma_wr_done;

wire disk_req_valid;
wire disk_req_write;
wire disk_req_flush;
wire [63:0] disk_req_lba;
wire [6:0] disk_req_word;
wire [31:0] disk_req_wdata;
// Auto-ack disk_req one cycle after valid (arms the DSM Deallocate path).
reg disk_req_done;
reg disk_req_valid_prev;
always @(posedge clk) begin
    if (rst) begin
        disk_req_done <= 1'b0;
        disk_req_valid_prev <= 1'b0;
    end else begin
        disk_req_valid_prev <= disk_req_valid;
        disk_req_done <= disk_req_valid && !disk_req_valid_prev;
    end
end
wire [31:0] disk_req_rdata = 32'h0;
wire disk_req_hit = 1'b0;
wire disk_busy = 1'b0;
wire disk_error = 1'b0;

wire msix_trigger;
wire pba_set_valid;
wire [15:0] pba_set_vector;
wire [12:0] id_rom_addr;
reg  [31:0] id_rom_data = 32'h0;

wire [7:0] dbg_state;
wire [15:0] dbg_active_qid;
wire [7:0] dbg_opcode;
wire [31:0] dbg_admin_queues;
wire [31:0] dbg_cmd_info;

// Ack every DMA write combinationally (host always accepts CQE / data writes).
assign dma_wr_done = dma_wr_valid;

// Minimal host memory: 16K DWORDs = 64 KiB. DWORD-addressed.
reg [31:0] host_mem [0:16383];
integer i;

// DMA read model: stream len beats from host_mem on dma_rd_req.
reg [63:0] rd_addr_q;
reg [9:0]  rd_len_q;
reg [9:0]  rd_beat_q;
reg        rd_busy;
always @(posedge clk) begin
    if (rst) begin
        rd_busy <= 1'b0;
        dma_rd_valid <= 1'b0;
        dma_rd_done <= 1'b0;
        rd_beat_q <= 10'h0;
    end else begin
        dma_rd_valid <= 1'b0;
        dma_rd_done <= 1'b0;
        if (!rd_busy) begin
            if (dma_rd_req) begin
                rd_addr_q <= dma_rd_addr;
                rd_len_q  <= dma_rd_len;
                rd_beat_q <= 10'h0;
                rd_busy   <= 1'b1;
            end
        end else begin
            dma_rd_valid <= 1'b1;
            dma_rd_data  <= host_mem[rd_addr_q[15:2] + {4'h0, rd_beat_q}];
            if ((rd_beat_q + 10'd1) >= rd_len_q) begin
                dma_rd_done <= 1'b1;
                rd_busy     <= 1'b0;
            end else begin
                rd_beat_q <= rd_beat_q + 10'd1;
            end
        end
    end
end

// Capture CQE status DW3 (addr[3:2]==2'b11) on ACQ/IOCQ writes.
integer cqe_count = 0;
integer cqe_snapshot = 0;
reg [31:0] last_cqe_status = 32'h0;
always @(posedge clk) begin
    if (!rst && dma_wr_valid && dma_wr_addr[3:2] == 2'b11 &&
        (dma_wr_addr[15:12] == 4'h2 || dma_wr_addr[15:12] == 4'h4)) begin
        last_cqe_status <= dma_wr_data;
        cqe_count <= cqe_count + 1;
    end
end

// Write back DMA writes into host_mem for payload inspection.
always @(posedge clk) begin
    if (!rst && dma_wr_valid)
        host_mem[dma_wr_addr[15:2]] <= dma_wr_data;
end

pcileech_nvme_admin_responder responder (
    .rst(rst),
    .clk(clk),
    .dma_enabled(dma_enabled),
    .cc_en(cc_en),
    .cc_enable_wr(cc_enable_wr),
    .cc_disable_wr(cc_disable_wr),
    .asq_lo(asq_lo),
    .asq_hi(asq_hi),
    .acq_lo(acq_lo),
    .acq_hi(acq_hi),
    .aqa(aqa),
    .doorbell_wr(doorbell_wr),
    .doorbell_is_cq(doorbell_is_cq),
    .doorbell_qid(doorbell_qid),
    .doorbell_val(doorbell_val),
    .msix_vector_addr(msix_vector_addr),
    .msix_vector_data(msix_vector_data),
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
    .disk_req_valid(disk_req_valid),
    .disk_req_write(disk_req_write),
    .disk_req_flush(disk_req_flush),
    .disk_req_lba(disk_req_lba),
    .disk_req_word(disk_req_word),
    .disk_req_wdata(disk_req_wdata),
    .disk_req_done(disk_req_done),
    .disk_req_rdata(disk_req_rdata),
    .disk_req_hit(disk_req_hit),
    .disk_busy(disk_busy),
    .disk_error(disk_error),
    .msix_trigger(msix_trigger),
    .pba_set_valid(pba_set_valid),
    .pba_set_vector(pba_set_vector),
    .id_rom_addr(id_rom_addr),
    .id_rom_data(id_rom_data),
    .dbg_state(dbg_state),
    .dbg_active_qid(dbg_active_qid),
    .dbg_opcode(dbg_opcode),
    .dbg_admin_queues(dbg_admin_queues),
    .dbg_cmd_info(dbg_cmd_info)
);

// Load a 16-DWORD SQE into host_mem at DWORD index dwbase.
task poke_sqe(input integer dwbase, input [7:0] op, input [31:0] nsid,
              input [31:0] prp1lo, input [31:0] prp1hi,
              input [31:0] prp2lo, input [31:0] prp2hi,
              input [31:0] cdw10, input [31:0] cdw11, input [31:0] cdw12);
begin
    host_mem[dwbase+0]  = {16'h0001, 8'h00, op};
    host_mem[dwbase+1]  = nsid;
    host_mem[dwbase+2]  = 32'h0;
    host_mem[dwbase+3]  = 32'h0;
    host_mem[dwbase+4]  = 32'h0;
    host_mem[dwbase+5]  = 32'h0;
    host_mem[dwbase+6]  = prp1lo;
    host_mem[dwbase+7]  = prp1hi;
    host_mem[dwbase+8]  = prp2lo;
    host_mem[dwbase+9]  = prp2hi;
    host_mem[dwbase+10] = cdw10;
    host_mem[dwbase+11] = cdw11;
    host_mem[dwbase+12] = cdw12;
    host_mem[dwbase+13] = 32'h0;
    host_mem[dwbase+14] = 32'h0;
    host_mem[dwbase+15] = 32'h0;
end
endtask

// Ring an SQ doorbell (qid, new tail value).
task ring_sq(input [15:0] qid, input [15:0] val);
begin
    @(negedge clk);
    doorbell_wr = 1'b1;
    doorbell_is_cq = 1'b0;
    doorbell_qid = qid;
    doorbell_val = val;
    @(negedge clk);
    doorbell_wr = 1'b0;
end
endtask

// Wait for the next CQE status DWORD and check DW3[31:17] against exp.
task wait_cqe(input [14:0] exp);
integer target;
integer cyc;
begin
    target = cqe_count + 1;
    cyc = 0;
    while (cqe_count < target && cyc < 20000) begin
        @(posedge clk);
        cyc = cyc + 1;
    end
    if (cqe_count < target) $fatal(1, "cqe timeout");
    #1;
    if (last_cqe_status[31:17] !== exp) $fatal(2, "wrong cqe status");
end
endtask

initial begin
    for (i = 0; i < 16384; i = i + 1)
        host_mem[i] = 32'h0;

    // Admin queue layout: ASQ=0x1000, ACQ=0x2000, 16 entries each.
    asq_lo = 32'h00001000;
    asq_hi = 32'h0;
    acq_lo = 32'h00002000;
    acq_hi = 32'h0;
    aqa    = {4'h0, 12'd15, 4'h0, 12'd15}; // asqs=15, acqs=15
    dma_enabled = 1'b1;

    repeat (2) @(posedge clk);
    @(negedge clk);
    rst = 1'b0;
    @(negedge clk);
    // cc_en stays 0 so scenario (0) can prove doorbells are gated until enable.

    // (0) Controller Enable handshake: doorbells before CC.EN must produce no CQE.
    poke_sqe(16'h400, 8'h0A, 32'h0, 32'h0, 32'h0, 32'h0, 32'h0, 32'h000000FF, 32'h0, 32'h0);
    ring_sq(16'd0, 16'd1); // cc_en=0 -> responder must ignore this admin doorbell
    cqe_snapshot = cqe_count;
    repeat (2000) @(posedge clk);
    if (cqe_count !== cqe_snapshot) $fatal(20, "doorbell processed before CC.EN set");
    #1;
    if (dbg_state !== 8'd0) $fatal(21, "responder left idle while CC.EN=0");
    // Assert CC.EN and pulse cc_enable_wr one cycle to latch AQA/ASQ/ACQ.
    @(negedge clk);
    cc_en = 1'b1;
    cc_enable_wr = 1'b1;
    @(negedge clk);
    cc_enable_wr = 1'b0;
    repeat (8) @(posedge clk);
    #1;
    if (dbg_state !== 8'd0) $fatal(22, "responder not idle after CC.EN handshake");
    @(negedge clk);

    // (1) Get Features with unsupported FID=0xFF -> INVALID_FIELD (now CC.EN=1 -> CQE).
    poke_sqe(16'h400, 8'h0A, 32'h0,
             32'h0, 32'h0, 32'h0, 32'h0,
             32'h000000FF, 32'h0, 32'h0);
    ring_sq(16'd0, 16'd1);
    wait_cqe(15'h0002);

    // (2) Identify CNS=0x01 (controller) -> SUCCESS. Data to 0x5000.
    poke_sqe(16'h410, 8'h06, 32'h0,
             32'h00005000, 32'h0, 32'h0, 32'h0,
             32'h00000001, 32'h0, 32'h0);
    ring_sq(16'd0, 16'd2);
    wait_cqe(15'h0000);

    // (3) Create I/O Completion Queue qid=1 @0x4000 -> SUCCESS.
    poke_sqe(16'h420, 8'h05, 32'h0,
             32'h00004000, 32'h0, 32'h0, 32'h0,
             32'h00010001, 32'h00000001, 32'h0);
    ring_sq(16'd0, 16'd3);
    wait_cqe(15'h0000);

    // (4) Create I/O Submission Queue qid=1 @0x3000 -> SUCCESS.
    poke_sqe(16'h430, 8'h01, 32'h0,
             32'h00003000, 32'h0, 32'h0, 32'h0,
             32'h00010001, 32'h00010001, 32'h0);
    ring_sq(16'd0, 16'd4);
    wait_cqe(15'h0000);

    // (5) Oversized I/O read (nlb=64 -> 8320 DW > MAX_XFER_DW) -> INVALID_FIELD.
    poke_sqe(16'hC00, 8'h02, 32'h00000001,
             32'h0, 32'h0, 32'h0, 32'h0,
             32'h0, 32'h0, 32'h00000040);
    ring_sq(16'd1, 16'd1);
    wait_cqe(15'h0002);

    // (6) Get Log Page LID=0x02 (SMART/Health) -> SUCCESS.
    poke_sqe(16'h440, 8'h02, 32'h0,
             32'h00006000, 32'h0, 32'h0, 32'h0,
             32'h007F0002, 32'h0, 32'h0);
    ring_sq(16'd0, 16'd5);
    wait_cqe(15'h0000);
    #1;
    if (host_mem[16'h1800][31:24] !== 8'h64) $fatal(5, "smart log spare byte");
    if (host_mem[16'h1800][7:0] !== 8'h00) $fatal(6, "smart log warning byte");
    if (host_mem[16'h1801] !== 32'h0000000A) $fatal(7, "smart log spare threshold");
    if (host_mem[16'h1824] !== 32'h00000003) $fatal(8, "smart log unsafe shutdowns");

    // (7) AER (opc 0x0C) must not complete synchronously.
    poke_sqe(16'h450, 8'h0C, 32'h0,
             32'h0, 32'h0, 32'h0, 32'h0,
             32'h0, 32'h0, 32'h0);
    cqe_snapshot = cqe_count;
    ring_sq(16'd0, 16'd6);
    repeat (2000) @(posedge clk);
    if (cqe_count !== cqe_snapshot) $fatal(8, "AER posted a synchronous CQE");
    #1;
    if (dbg_state !== 8'd0) $fatal(9, "AER did not return to idle");

    // (8) DSM Deallocate (opc 0x09, I/O qid=1) -> SUCCESS.
    host_mem[16'h1C00] = 32'h00000000; // DW0 (unused)
    host_mem[16'h1C01] = 32'h00000008; // DW1 = NLB (non-zero -> disk invalidate path)
    host_mem[16'h1C02] = 32'h00000000; // DW2 = SLBA lo
    host_mem[16'h1C03] = 32'h00000000; // DW3 = SLBA hi
    poke_sqe(16'hC10, 8'h09, 32'h00000001,
             32'h00007000, 32'h0, 32'h0, 32'h0,
             32'h00000000, 32'h00000004, 32'h0);
    ring_sq(16'd1, 16'd2);
    wait_cqe(15'h0000);
    #1;
    if (dbg_state !== 8'd0) $fatal(10, "DSM did not return to idle");

    // Responder must return to S_IDLE (8'd0) after the flow.
    repeat (8) @(posedge clk);
    #1;
    if (dbg_state !== 8'd0) $fatal(3, "responder did not return to idle");

    $display("NVME_ADMIN_BEHAVIOR_PASS");
    $finish;
end

initial begin
    repeat (200000) @(posedge clk);
    $fatal(4, "global timeout");
end
endmodule
`
}

// TestNVMeAdminBehaviorScenarios drives a DMA-backed admin/I-O flow through Verilator.
func TestNVMeAdminBehaviorScenarios(t *testing.T) {
	if _, err := exec.LookPath("verilator"); err != nil {
		t.Skip("verilator not installed")
	}
	cfg := testConfig()
	responderSV, err := GenerateNVMeResponderSV(cfg)
	if err != nil {
		t.Fatalf("GenerateNVMeResponderSV: %v", err)
	}
	runVerilatorBinary(t, responderSV, nvmeAdminBehaviorBench())
}
