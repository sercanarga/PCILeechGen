package svgen

import (
	"os/exec"
	"testing"
)

func TestLatencySpillBufferNoDrop(t *testing.T) {
	if _, err := exec.LookPath("verilator"); err != nil {
		t.Skip("verilator not installed")
	}
	cfg := testConfig()
	cfg.LatencyConfig = DefaultLatencyConfig(cfg.ClassCode)
	controller, err := GenerateBarControllerSV(cfg)
	if err != nil {
		t.Fatalf("GenerateBarControllerSV: %v", err)
	}
	block1 := extractHDLBlock(t, controller, "    wire [87:0] bar0_raw_ctx;", "    wire [87:0] bar_base_ctx[7];")
	block2 := extractHDLThroughAlways(t, controller, "    // 16-deep circular buffer FSM:", "always @(posedge clk)")

	dut := "`timescale 1ns/1ps\n" + `module bar0_lat_buf(
    input            clk,
    input            device_reset,
    input            raw_valid_i,
    input [87:0]     raw_ctx_i,
    input [31:0]     raw_data_i,
    input            emu_busy_i,
    output [4:0]     buf_count_o,
    output           buf_empty_o,
    output           buf_almost_full_o,
    output           spill_valid_o,
    output [31:0]    head_data_o,
    output           buf_pop_o
);
    wire hda_crst_falling = 1'b0;
` + block1 + `
    assign bar0_raw_valid = raw_valid_i;
    assign bar0_raw_ctx   = raw_ctx_i;
    assign bar0_raw_data  = raw_data_i;
    assign bar0_emu_busy  = emu_busy_i;
` + block2 + `
    assign buf_count_o       = bar0_buf_count;
    assign buf_empty_o       = bar0_buf_empty;
    assign buf_almost_full_o = bar0_buf_almost_full;
    assign spill_valid_o     = bar0_spill_valid;
    assign head_data_o       = bar0_buf_data_w;
    assign buf_pop_o         = bar0_buf_pop;
endmodule
`
	bench := `
module tb;
timeunit 1ns;
timeprecision 1ps;
bit clk = 0;
always #5 clk = ~clk;
bit rst = 1;
bit raw_valid_i = 0;
bit [87:0] raw_ctx_i = 0;
bit [31:0] raw_data_i = 0;
bit emu_busy_i = 0;
bit producer_run = 0;
wire [4:0] buf_count;
wire buf_empty;
wire buf_almost_full;
wire spill_valid;
wire [31:0] head_data;
wire buf_pop;

bar0_lat_buf dut(
    .clk(clk), .device_reset(rst),
    .raw_valid_i(raw_valid_i), .raw_ctx_i(raw_ctx_i), .raw_data_i(raw_data_i),
    .emu_busy_i(emu_busy_i),
    .buf_count_o(buf_count), .buf_empty_o(buf_empty),
    .buf_almost_full_o(buf_almost_full), .spill_valid_o(spill_valid),
    .head_data_o(head_data), .buf_pop_o(buf_pop)
);

integer expected;
integer pushed;
integer popped;
always @(posedge clk) begin
    if (!rst && buf_pop) begin
        if (head_data !== expected)
            $fatal(1, "fifo loss/order: got %0d exp %0d", head_data, expected);
        expected = expected + 1;
        popped = popped + 1;
    end
end

localparam integer ACCEPT_PERIOD = 3;
integer cyc;

task automatic fill(input integer n);
    integer i;
begin
    emu_busy_i = 0;
    for (i = 0; i < n; i = i + 1) begin
        @(negedge clk);
        raw_valid_i = 1;
        raw_data_i  = pushed;
        raw_ctx_i   = {56'h0, pushed[31:0]};
        pushed      = pushed + 1;
        @(posedge clk);
    end
    @(negedge clk);
    raw_valid_i = 0;
end
endtask

task automatic drain_all;
    integer guard;
begin
    guard = 0;
    while ((!(buf_empty && !spill_valid)) && guard < 20000) begin
        @(negedge clk);
        emu_busy_i = (cyc % ACCEPT_PERIOD == 0);
        cyc = cyc + 1;
        @(posedge clk);
        guard = guard + 1;
    end
    @(negedge clk);
    emu_busy_i = 0;
    if (guard >= 20000) $fatal(2, "drain timeout");
end
endtask

integer iter;
initial begin
    expected = 0; pushed = 0; popped = 0; cyc = 0;
    repeat (3) @(posedge clk);
    @(negedge clk); rst = 0;

    for (iter = 0; iter < 8; iter = iter + 1) begin
        fill(17);
        #1;
        if (buf_count !== 5'd16) $fatal(3, "iter %0d count %0d after fill", iter, buf_count);
        if (!spill_valid)        $fatal(4, "iter %0d spill not engaged", iter);
        if (!buf_almost_full)    $fatal(5, "iter %0d almost_full not asserted", iter);
        drain_all;
        #1;
        if (!buf_empty || spill_valid) $fatal(6, "iter %0d not drained", iter);
    end

    producer_run = 1;
    for (iter = 0; iter < 1000; iter = iter + 1) begin
        @(negedge clk);
        raw_valid_i = producer_run && !buf_almost_full;
        if (raw_valid_i) begin
            raw_data_i = pushed;
            raw_ctx_i  = {56'h0, pushed[31:0]};
            pushed     = pushed + 1;
        end
        emu_busy_i = (cyc % ACCEPT_PERIOD == 0);
        cyc = cyc + 1;
        @(posedge clk);
    end
    @(negedge clk);
    producer_run = 0;
    raw_valid_i = 0;
    drain_all;
    #1;
    if (!buf_empty || spill_valid) $fatal(7, "phase 2 not drained");

    if (pushed !== popped)
        $fatal(8, "produced %0d != delivered %0d (silent drop)", pushed, popped);
    if (expected !== pushed)
        $fatal(9, "expected %0d != produced %0d", expected, pushed);

    $display("LATENCY_SPILL_NODROP_PASS");
    $finish;
end

initial begin
    repeat (500000) @(posedge clk);
    $fatal(10, "global timeout");
end
endmodule
`
	runVerilatorBinary(t, dut, bench)
}
