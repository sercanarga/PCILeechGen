`timescale 1ns / 1ps
`include "pcileech_header.svh"

module pcileech_tlp_ur_completer #(
    parameter integer REQUEST_FIFO_DEPTH = 256
)(
    input                   rst,
    input                   clk,
    input [15:0]            pcie_id,
    input                   request_valid,
    output                  request_ready,
    input [15:0]            requester_id,
    input [7:0]             tag,
    input [2:0]             traffic_class,
    input [2:0]             attributes,
    IfAXIS128.source        tlps_out
);
    reg [15:0] requester_fifo [0:REQUEST_FIFO_DEPTH-1];
    reg [7:0]  tag_fifo       [0:REQUEST_FIFO_DEPTH-1];
    reg [2:0]  tc_fifo        [0:REQUEST_FIFO_DEPTH-1];
    reg [2:0]  attr_fifo      [0:REQUEST_FIFO_DEPTH-1];
    reg [7:0]  write_ptr;
    reg [7:0]  read_ptr;
    reg [8:0]  request_count;
    reg [127:0] completion_data;

    wire request_fifo_full = (request_count == REQUEST_FIFO_DEPTH);
    wire request_pop = tlps_out.tvalid && tlps_out.tready;
    wire request_can_accept = !request_fifo_full || request_pop;
    wire request_push = request_valid && request_can_accept;

    assign request_ready = request_can_accept;

    always @(posedge clk) begin
        if (rst) begin
            write_ptr <= 8'h00;
        end else if (request_push) begin
            requester_fifo[write_ptr] <= requester_id;
            tag_fifo[write_ptr]       <= tag;
            tc_fifo[write_ptr]        <= traffic_class;
            attr_fifo[write_ptr]      <= attributes;
            write_ptr                 <= write_ptr + 1'b1;
        end
    end

    always @(posedge clk) begin
        if (rst) begin
            read_ptr      <= 8'h00;
            request_count <= 9'h000;
        end else begin
            case ({request_push, request_pop})
                2'b10: request_count <= request_count + 1'b1;
                2'b01: request_count <= request_count - 1'b1;
                default: request_count <= request_count;
            endcase
            if (request_pop)
                read_ptr <= read_ptr + 1'b1;
        end
    end

    always @(*) begin
        completion_data = 128'h0;
        completion_data[31:29] = 3'b000;
        completion_data[28:24] = 5'b01010;
        completion_data[22:20] = tc_fifo[read_ptr];
        completion_data[18]    = attr_fifo[read_ptr][2];
        completion_data[13:12] = attr_fifo[read_ptr][1:0];
        completion_data[63:32] = {pcie_id[7:0], pcie_id[15:8], 3'b001, 1'b0, 12'h000};
        completion_data[95:64] = {requester_fifo[read_ptr], tag_fifo[read_ptr], 8'h00};
    end

    assign tlps_out.tdata    = completion_data;
    assign tlps_out.tkeepdw  = 4'b0111;
    assign tlps_out.tvalid   = (request_count != 0);
    assign tlps_out.tlast    = 1'b1;
    assign tlps_out.tuser    = 9'b000000011;
    assign tlps_out.has_data = (request_count != 0);
endmodule
