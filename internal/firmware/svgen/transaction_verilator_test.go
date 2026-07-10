package svgen

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

const axis128Header = "`ifndef _pcileech_header_svh_\n`define _pcileech_header_svh_\ninterface IfAXIS128;\nwire [127:0] tdata;\nwire [3:0] tkeepdw;\nwire tvalid;\nwire tlast;\nwire [8:0] tuser;\nwire tready;\nwire has_data;\nmodport source(input tready, output tdata, tkeepdw, tvalid, tlast, tuser, has_data);\nmodport sink(output tready, input tdata, tkeepdw, tvalid, tlast, tuser, has_data);\nmodport source_lite(output tdata, tkeepdw, tvalid, tlast, tuser);\nmodport sink_lite(input tdata, tkeepdw, tvalid, tlast, tuser);\nendinterface\n`endif\n"

func runVerilatorBinary(t *testing.T, dut, bench string) {
	t.Helper()
	verilator, err := exec.LookPath("verilator")
	if err != nil {
		t.Skip("verilator not installed")
	}
	dir := t.TempDir()
	for name, content := range map[string]string{
		"pcileech_header.svh": axis128Header,
		"dut.sv":              dut,
		"tb.sv":               bench,
	} {
		if err := os.WriteFile(filepath.Join(dir, name), []byte(content), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	cmd := exec.Command(verilator, "--binary", "--timing", "-Wno-fatal", "--top-module", "tb", "--Mdir", "obj", "-I"+dir, "dut.sv", "tb.sv")
	cmd.Dir = dir
	if output, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("verilator build failed: %v\n%s", err, output)
	}
	binary := filepath.Join(dir, "obj", "Vtb")
	if output, err := exec.Command(binary).CombinedOutput(); err != nil {
		t.Fatalf("verilator simulation failed: %v\n%s", err, output)
	}
}

func TestVerilatorURCompletionStableWhileStalled(t *testing.T) {
	dut, err := GenerateURCompleterSV(testConfig())
	if err != nil {
		t.Fatalf("GenerateURCompleterSV() error = %v", err)
	}
	bench := "`timescale 1ns/1ps\n`include \"pcileech_header.svh\"\nmodule tb;\nbit clk = 0;\nalways #5 clk = ~clk;\nbit rst = 1;\nbit ready = 0;\nbit request_valid = 0;\nwire request_ready;\nIfAXIS128 out();\nassign out.tready = ready;\npcileech_tlp_ur_completer dut(.rst(rst), .clk(clk), .pcie_id(16'h1234), .request_valid(request_valid), .request_ready(request_ready), .requester_id(16'hA15E), .tag(8'hD3), .traffic_class(3'h5), .attributes(3'h3), .tlps_out(out.source));\nreg [142:0] held;\ninitial begin\nrepeat (2) @(posedge clk);\n@(negedge clk); rst = 0; request_valid = 1;\n@(posedge clk);\n@(negedge clk); request_valid = 0;\nwait (out.tvalid);\nheld = {out.has_data, out.tuser, out.tlast, out.tkeepdw, out.tdata};\nif (out.tdata[31:29] !== 3'b000) $fatal(1, \"fmt\");\nif (out.tdata[28:24] !== 5'b01010) $fatal(1, \"type\");\nif (out.tdata[22:20] !== 3'h5) $fatal(1, \"tc\");\nif ({out.tdata[18], out.tdata[13:12]} !== 3'h3) $fatal(1, \"attr\");\nif (out.tdata[47:45] !== 3'b001) $fatal(1, \"status\");\nif (out.tdata[95:80] !== 16'hA15E) $fatal(1, \"requester\");\nif (out.tdata[79:72] !== 8'hD3) $fatal(1, \"tag\");\nif (out.tdata[9:0] !== 10'h000) $fatal(1, \"length\");\nif (out.tkeepdw !== 4'b0111) $fatal(1, \"keep\");\nrepeat (4) begin\n@(posedge clk);\n#1;\nif ({out.has_data, out.tuser, out.tlast, out.tkeepdw, out.tdata} !== held) $fatal(1, \"stability\");\nend\n@(negedge clk); ready = 1;\n@(posedge clk);\n@(negedge clk);\nif (out.tvalid !== 1'b0) $fatal(1, \"drain\");\n$finish;\nend\nendmodule\n"
	runVerilatorBinary(t, dut, bench)
}

func TestVerilatorCompletionArbiterLocksStalledPacket(t *testing.T) {
	controller, err := GenerateBarControllerSV(testConfig())
	if err != nil {
		t.Fatalf("GenerateBarControllerSV() error = %v", err)
	}
	block := extractHDLBlock(t, controller, "    reg cpl_grant_active;", "    wire        wrengine_ready;")
	dut := "`timescale 1ns/1ps\n`include \"pcileech_header.svh\"\nmodule cpl_arb(input clk, input rst, IfAXIS128.sink tlps_ur, IfAXIS128.sink tlps_rdeng, IfAXIS128.source tlps_cpl);\n" + block + "endmodule\n"
	bench := `
module tb;
timeunit 1ns;
timeprecision 1ps;
bit clk = 0;
always #5 clk = ~clk;
bit rst = 1;
bit ready = 0;
bit ur_valid = 0;
bit ur_last = 1;
bit [127:0] ur_data = 128'hA1;
bit rd_valid = 0;
bit rd_last = 0;
bit [127:0] rd_data = 128'hB1;
IfAXIS128 ur();
IfAXIS128 rd();
IfAXIS128 cpl();
assign ur.tdata = ur_data;
assign ur.tkeepdw = 4'h7;
assign ur.tvalid = ur_valid;
assign ur.tlast = ur_last;
assign ur.tuser = 9'h003;
assign ur.has_data = ur_valid;
assign rd.tdata = rd_data;
assign rd.tkeepdw = 4'hF;
assign rd.tvalid = rd_valid;
assign rd.tlast = rd_last;
assign rd.tuser = {7'h00, rd_last, 1'b1};
assign rd.has_data = rd_valid;
assign cpl.tready = ready;
cpl_arb dut(.clk(clk), .rst(rst), .tlps_ur(ur.sink), .tlps_rdeng(rd.sink), .tlps_cpl(cpl.source));
reg [142:0] held;
initial begin
repeat (2) @(posedge clk);
@(negedge clk); rst = 0; ur_valid = 1; rd_valid = 1;
@(posedge clk);
@(negedge clk);
if (!dut.cpl_grant_active || dut.cpl_grant_ur) $fatal(1, "grant");
if (cpl.tdata !== 128'hB1) $fatal(1, "read_select");
held = {cpl.has_data, cpl.tuser, cpl.tlast, cpl.tkeepdw, cpl.tdata};
ur_data = 128'hA2;
repeat (3) begin
@(posedge clk);
#1;
if ({cpl.has_data, cpl.tuser, cpl.tlast, cpl.tkeepdw, cpl.tdata} !== held) $fatal(1, "first_stall");
end
@(negedge clk); ready = 1;
@(posedge clk);
@(negedge clk); ready = 0; rd_data = 128'hB2; rd_last = 1;
#1;
held = {cpl.has_data, cpl.tuser, cpl.tlast, cpl.tkeepdw, cpl.tdata};
ur_data = 128'hA3;
repeat (3) begin
@(posedge clk);
#1;
if ({cpl.has_data, cpl.tuser, cpl.tlast, cpl.tkeepdw, cpl.tdata} !== held) $fatal(1, "last_stall");
end
@(negedge clk); ready = 1;
@(posedge clk);
@(negedge clk); rd_valid = 0;
#1;
if (dut.cpl_grant_active) $fatal(1, "release");
if (!cpl.tvalid || cpl.tdata !== 128'hA3) $fatal(1, "ur_after_read");
$finish;
end
endmodule
`
	runVerilatorBinary(t, dut, bench)
}

func TestVerilatorCompletionArbiterServicesQueuedURUnderContinuousReads(t *testing.T) {
	controller, err := GenerateBarControllerSV(testConfig())
	if err != nil {
		t.Fatalf("GenerateBarControllerSV() error = %v", err)
	}
	block := extractHDLBlock(t, controller, "    reg cpl_grant_active;", "    wire        wrengine_ready;")
	dut := "`timescale 1ns/1ps\n`include \"pcileech_header.svh\"\nmodule cpl_arb(input clk, input rst, IfAXIS128.sink tlps_ur, IfAXIS128.sink tlps_rdeng, IfAXIS128.source tlps_cpl);\n" + block + "endmodule\n"
	bench := `
module tb;
timeunit 1ns;
timeprecision 1ps;
bit clk = 0;
always #5 clk = ~clk;
bit rst = 1;
bit ready = 0;
bit [15:0] lfsr = 16'hACE1;
bit ur_valid = 0;
bit ur_queued = 0;
bit ur_seen = 0;
integer rd_beat = 0;
integer rd_packet = 0;
integer rd_boundaries_after_ur = 0;
bit monitor_active = 0;
bit monitor_ur = 0;
bit stalled = 0;
reg [142:0] stalled_value;
IfAXIS128 ur();
IfAXIS128 rd();
IfAXIS128 out();
wire [127:0] rd_data = {8'hB0, rd_packet[7:0], rd_beat[7:0], 104'h0};
wire rd_last = (rd_beat == 2);
assign ur.tdata = 128'hA50000000000000000000000000000D3;
assign ur.tkeepdw = 4'h7;
assign ur.tvalid = ur_valid;
assign ur.tlast = 1'b1;
assign ur.tuser = 9'h003;
assign ur.has_data = ur_valid;
assign rd.tdata = rd_data;
assign rd.tkeepdw = 4'hF;
assign rd.tvalid = !rst;
assign rd.tlast = rd_last;
assign rd.tuser = {7'h00, rd_last, 1'b1};
assign rd.has_data = !rst;
assign out.tready = ready;
cpl_arb dut(.clk(clk), .rst(rst), .tlps_ur(ur.sink), .tlps_rdeng(rd.sink), .tlps_cpl(out.source));
always @(negedge clk) begin
if (rst) begin
lfsr = 16'hACE1;
ready = 0;
end else begin
lfsr = {lfsr[14:0], lfsr[15] ^ lfsr[13] ^ lfsr[12] ^ lfsr[10]};
ready = lfsr[0] | lfsr[3];
end
end
always @(posedge clk) begin
if (rst) begin
rd_beat <= 0;
rd_packet <= 0;
ur_valid <= 0;
ur_queued <= 0;
end else begin
if (rd.tvalid && rd.tready) begin
if (!ur_queued) begin
ur_valid <= 1;
ur_queued <= 1;
end
if (rd_last) begin
rd_beat <= 0;
rd_packet <= rd_packet + 1;
end else begin
rd_beat <= rd_beat + 1;
end
end
if (ur_valid && ur.tready) begin
ur_valid <= 0;
ur_seen <= 1;
end
end
end
always @(posedge clk) begin
bit source_ur;
if (!rst && out.tvalid && out.tready) begin
source_ur = (out.tdata[127:120] == 8'hA5);
if (monitor_active && source_ur != monitor_ur) $fatal(1, "interleave");
if (!monitor_active) monitor_ur <= source_ur;
monitor_active <= !out.tlast;
if (ur_queued && !ur_seen && !source_ur && out.tlast) begin
if (rd_boundaries_after_ur >= 1) $fatal(1, "starvation");
rd_boundaries_after_ur <= rd_boundaries_after_ur + 1;
end
end
if (!rst && stalled) begin
if ({out.has_data, out.tuser, out.tlast, out.tkeepdw, out.tdata} !== stalled_value) $fatal(1, "stall_stability");
if (out.tvalid && out.tready) stalled <= 0;
end else if (!rst && out.tvalid && !out.tready) begin
stalled <= 1;
stalled_value <= {out.has_data, out.tuser, out.tlast, out.tkeepdw, out.tdata};
end
if (ur_seen) begin
if (rd_boundaries_after_ur > 1) $fatal(1, "bound");
$finish;
end
end
initial begin
repeat (2) @(posedge clk);
@(negedge clk); rst = 0;
repeat (1000) @(posedge clk);
$fatal(1, "timeout");
end
endmodule
`
	runVerilatorBinary(t, dut, bench)
}

func TestVerilatorHDAArbiterLocksPacketsUnderBackpressure(t *testing.T) {
	controller, err := GenerateBarControllerSV(audioConfig())
	if err != nil {
		t.Fatalf("GenerateBarControllerSV() error = %v", err)
	}
	block := extractHDLThroughAlways(t, controller, "    reg hda_packet_active;", "    always @(posedge clk) begin")
	dut := "`timescale 1ns/1ps\n`include \"pcileech_header.svh\"\nmodule hda_arb(input clk, input rst, input [127:0] hda_tlp_tx_tdata, input [3:0] hda_tlp_tx_tkeepdw, input hda_tlp_tx_tvalid, input hda_tlp_tx_tlast, input [8:0] hda_tlp_tx_tuser, output hda_tlp_tx_tready, IfAXIS128.sink tlps_cpl, IfAXIS128.source tlps_out);\n" + block + "endmodule\n"
	bench := `
module tb;
timeunit 1ns;
timeprecision 1ps;
bit clk = 0;
always #5 clk = ~clk;
bit rst = 1;
bit ready = 0;
bit dma_valid = 0;
bit dma_last = 0;
bit [127:0] dma_data = 128'hD1;
bit cpl_valid = 0;
bit cpl_last = 0;
bit [127:0] cpl_data = 128'hC1;
wire dma_ready;
IfAXIS128 cpl();
IfAXIS128 out();
assign cpl.tdata = cpl_data;
assign cpl.tkeepdw = 4'hF;
assign cpl.tvalid = cpl_valid;
assign cpl.tlast = cpl_last;
assign cpl.tuser = {7'h00, cpl_last, 1'b1};
assign cpl.has_data = cpl_valid;
assign out.tready = ready;
hda_arb dut(.clk(clk), .rst(rst), .hda_tlp_tx_tdata(dma_data), .hda_tlp_tx_tkeepdw(4'hF), .hda_tlp_tx_tvalid(dma_valid), .hda_tlp_tx_tlast(dma_last), .hda_tlp_tx_tuser({7'h00, dma_last, 1'b1}), .hda_tlp_tx_tready(dma_ready), .tlps_cpl(cpl.sink), .tlps_out(out.source));
reg [142:0] held;
initial begin
repeat (2) @(posedge clk);
@(negedge clk); rst = 0; dma_valid = 1; cpl_valid = 1;
@(posedge clk);
@(negedge clk);
if (!dut.hda_packet_active || !dut.hda_packet_select_dma) $fatal(1, "dma_grant");
if (out.tdata !== 128'hD1 || dma_ready || cpl.tready) $fatal(1, "dma_stall");
held = {out.has_data, out.tuser, out.tlast, out.tkeepdw, out.tdata};
cpl_data = 128'hC9;
repeat (2) begin
@(posedge clk);
#1;
if ({out.has_data, out.tuser, out.tlast, out.tkeepdw, out.tdata} !== held) $fatal(1, "dma_stability");
end
@(negedge clk); ready = 1;
@(posedge clk);
@(negedge clk); ready = 0; dma_data = 128'hD2; dma_last = 1;
#1;
if (out.tdata !== 128'hD2 || dma_ready || cpl.tready) $fatal(1, "dma_last_stall");
@(negedge clk); ready = 1;
@(posedge clk);
@(negedge clk); ready = 0; dma_data = 128'hD3; dma_last = 1; cpl_data = 128'hC1; cpl_last = 0;
@(posedge clk);
@(negedge clk);
if (!dut.hda_packet_active || dut.hda_packet_select_dma) $fatal(1, "cpl_grant");
if (out.tdata !== 128'hC1 || dma_ready || cpl.tready) $fatal(1, "cpl_stall");
held = {out.has_data, out.tuser, out.tlast, out.tkeepdw, out.tdata};
dma_data = 128'hD9;
repeat (2) begin
@(posedge clk);
#1;
if ({out.has_data, out.tuser, out.tlast, out.tkeepdw, out.tdata} !== held) $fatal(1, "cpl_stability");
end
@(negedge clk); ready = 1;
@(posedge clk);
@(negedge clk); ready = 0; cpl_data = 128'hC2; cpl_last = 1;
#1;
if (out.tdata !== 128'hC2 || dma_ready || cpl.tready) $fatal(1, "cpl_last_stall");
@(negedge clk); ready = 1;
@(posedge clk);
@(negedge clk); cpl_valid = 0; ready = 0;
#1;
if (dut.hda_packet_active) $fatal(1, "release");
if (!out.tvalid || out.tdata !== 128'hD9) $fatal(1, "dma_after_cpl");
$finish;
end
endmodule
`
	runVerilatorBinary(t, dut, bench)
}

func TestVerilatorReadEngineDrivesCompleteStreamContract(t *testing.T) {
	engine, err := GenerateBarReadEngineSV(testConfig())
	if err != nil {
		t.Fatalf("GenerateBarReadEngineSV() error = %v", err)
	}
	stub := `
module fifo_134_134_clk1_bar_rdrsp(input srst, input clk, input [133:0] din, input wr_en, input rd_en, output [133:0] dout, output full, output empty, output prog_empty, output valid);
reg [133:0] value;
reg occupied;
assign dout = value;
assign full = 1'b0;
assign empty = !occupied;
assign prog_empty = 1'b1;
assign valid = occupied;
always @(posedge clk) begin
if (srst) begin
value <= 134'h0;
occupied <= 1'b0;
end else if (wr_en) begin
value <= din;
occupied <= 1'b1;
end else if (rd_en) begin
occupied <= 1'b0;
end
end
endmodule
`
	bench := `
module tb;
timeunit 1ns;
timeprecision 1ps;
bit clk = 0;
always #5 clk = ~clk;
bit rst = 1;
bit ready = 0;
bit rsp_valid = 0;
bit [87:0] rsp_ctx;
IfAXIS128 out();
assign out.tready = ready;
wire [87:0] req_ctx;
wire [6:0] req_bar;
wire [31:0] req_addr;
wire req_valid;
pcileech_tlps128_bar_rdengine dut(.rst(rst), .clk(clk), .pcie_id(16'h1234), .tlps_out(out.source), .tlps_in_valid(1'b0), .norm_address(64'h0), .norm_length_dw(11'd1), .norm_first_be(4'hF), .norm_last_be(4'h0), .norm_requester_id(16'h0), .norm_tag(8'h0), .norm_traffic_class(3'h0), .norm_attributes(3'h0), .norm_header_4dw(1'b0), .norm_bar_mask(7'h1), .norm_enabled_byte_count(13'd4), .norm_first_completion_byte_count(13'd4), .norm_first_completion_dw(11'd1), .rd_req_ctx(req_ctx), .rd_req_bar(req_bar), .rd_req_addr(req_addr), .rd_req_valid(req_valid), .rd_rsp_ctx(rsp_ctx), .rd_rsp_data(32'hDEADBEEF), .rd_rsp_valid(rsp_valid));
reg [142:0] held;
initial begin
rsp_ctx = {1'b1, 1'b1, 11'd1, 12'd4, 3'h5, 3'h3, 1'b0, 8'hD3, 16'hA15E, 32'h00000040};
repeat (2) @(posedge clk);
@(negedge clk); rst = 0; rsp_valid = 1;
@(posedge clk);
@(negedge clk); rsp_valid = 0;
wait (out.tvalid);
#1;
if (out.tuser !== 9'b000000011) $fatal(1, "tuser");
if (out.has_data !== out.tvalid) $fatal(1, "has_data");
held = {out.has_data, out.tuser, out.tlast, out.tkeepdw, out.tdata};
repeat (3) begin
@(posedge clk);
#1;
if ({out.has_data, out.tuser, out.tlast, out.tkeepdw, out.tdata} !== held) $fatal(1, "stall");
end
@(negedge clk); ready = 1;
@(posedge clk);
@(negedge clk);
if (out.tvalid) $fatal(1, "drain");
$finish;
end
endmodule
`
	runVerilatorBinary(t, engine+stub, bench)
}

func TestVerilatorLockedReadsEmitURFor3DWAnd4DW(t *testing.T) {
	normalizer, err := GenerateTransactionNormalizerSV(testConfig())
	if err != nil {
		t.Fatalf("GenerateTransactionNormalizerSV() error = %v", err)
	}
	completer, err := GenerateURCompleterSV(testConfig())
	if err != nil {
		t.Fatalf("GenerateURCompleterSV() error = %v", err)
	}
	bench := `
module tb;
timeunit 1ns;
timeprecision 1ps;
bit clk = 0;
always #5 clk = ~clk;
bit rst = 1;
bit ready = 0;
bit [127:0] tlp_data;
bit [8:0] tlp_user = 0;
wire request_present;
wire ur_required;
wire non_posted_request;
wire posted_write;
wire [15:0] requester_id;
wire [7:0] tag;
wire [2:0] traffic_class;
wire [2:0] attributes;
wire request_ready;
IfAXIS128 out();
assign out.tready = ready;
pcileech_tlp_normalizer norm(.tlp_data(tlp_data), .tlp_user(tlp_user), .tlp_last(1'b1), .request_present(request_present), .ur_required(ur_required), .non_posted_request(non_posted_request), .posted_write(posted_write), .requester_id(requester_id), .tag(tag), .traffic_class(traffic_class), .attributes(attributes));
pcileech_tlp_ur_completer ur(.rst(rst), .clk(clk), .pcie_id(16'h1234), .request_valid(request_present && ur_required && non_posted_request), .request_ready(request_ready), .requester_id(requester_id), .tag(tag), .traffic_class(traffic_class), .attributes(attributes), .tlps_out(out.source));
task automatic send_locked(input [2:0] fmt);
begin
@(negedge clk);
tlp_data = 128'h0;
tlp_data[31:29] = fmt;
tlp_data[28:24] = 5'b00001;
tlp_data[22:20] = 3'h5;
tlp_data[18] = 1'b0;
tlp_data[13:12] = 2'b11;
tlp_data[9:0] = 10'd1;
tlp_data[35:32] = 4'hF;
tlp_data[39:36] = 4'h0;
tlp_data[63:48] = 16'hA15E;
tlp_data[47:40] = 8'hD3;
tlp_user = 9'b000000101;
#1;
if (!ur_required || !non_posted_request || posted_write) $fatal(1, "classification");
@(posedge clk);
@(negedge clk);
tlp_user = 0;
wait (out.tvalid);
#1;
if (out.tdata[47:45] !== 3'b001) $fatal(1, "status");
if (out.tdata[95:80] !== 16'hA15E || out.tdata[79:72] !== 8'hD3) $fatal(1, "identity");
if (out.tdata[22:20] !== 3'h5 || {out.tdata[18], out.tdata[13:12]} !== 3'h3) $fatal(1, "metadata");
@(negedge clk);
ready = 1;
@(posedge clk);
@(negedge clk);
ready = 0;
end
endtask
initial begin
tlp_data = 0;
repeat (2) @(posedge clk);
@(negedge clk); rst = 0;
send_locked(3'b000);
send_locked(3'b001);
$finish;
end
endmodule
`
	runVerilatorBinary(t, normalizer+completer, bench)
}

func extractHDLThroughAlways(t *testing.T, source, start, always string) string {
	t.Helper()
	startIndex := strings.Index(source, start)
	if startIndex < 0 {
		t.Fatalf("start marker missing: %s", start)
	}
	alwaysIndex := strings.Index(source[startIndex:], always)
	if alwaysIndex < 0 {
		t.Fatalf("always marker missing: %s", always)
	}
	cursor := startIndex + alwaysIndex
	depth := 0
	seen := false
	for cursor < len(source) {
		lineEnd := strings.IndexByte(source[cursor:], '\n')
		if lineEnd < 0 {
			lineEnd = len(source) - cursor
		}
		lineEnd += cursor
		for _, token := range strings.Fields(source[cursor:lineEnd]) {
			token = strings.Trim(token, "()@:;")
			if token == "begin" {
				depth++
				seen = true
			}
			if token == "end" {
				depth--
			}
		}
		if seen && depth == 0 {
			return source[startIndex : lineEnd+1]
		}
		cursor = lineEnd + 1
	}
	t.Fatalf("balanced always block missing after: %s", always)
	return ""
}

func extractHDLBlock(t *testing.T, source, start, end string) string {
	t.Helper()
	startIndex := strings.Index(source, start)
	if startIndex < 0 {
		t.Fatalf("start marker missing: %s", start)
	}
	endIndex := strings.Index(source[startIndex:], end)
	if endIndex < 0 {
		t.Fatalf("end marker missing: %s\n%.1200s", end, source[startIndex:])
	}
	return source[startIndex : startIndex+endIndex]
}
