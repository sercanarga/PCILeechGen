`timescale 1ns / 1ps
`include "pcileech_header.svh"

module pcileech_tlps128_bar_rdengine #(
    parameter integer READ_COMPLETION_BOUNDARY = 64,
    parameter integer MAX_PAYLOAD_BYTES         = 128
)(
    input                   rst,
    input                   clk,
    input [15:0]            pcie_id,
    input                   stall,     // downstream completion path saturated; hold issue
    IfAXIS128.source        tlps_out,

    input                   tlps_in_valid,
    input [63:0]            norm_address,
    input [10:0]            norm_length_dw,
    input [3:0]             norm_first_be,
    input [3:0]             norm_last_be,
    input [15:0]            norm_requester_id,
    input [7:0]             norm_tag,
    input [2:0]             norm_traffic_class,
    input [2:0]             norm_attributes,
    input                   norm_header_4dw,
    input [6:0]             norm_bar_mask,
    input [12:0]            norm_enabled_byte_count,
    input [12:0]            norm_first_completion_byte_count,
    input [10:0]            norm_first_completion_dw,

    output [87:0]           rd_req_ctx,
    output [6:0]            rd_req_bar,
    output [31:0]           rd_req_addr,
    output                  rd_req_valid,
    input  [87:0]           rd_rsp_ctx,
    input  [31:0]           rd_rsp_data,
    input                   rd_rsp_valid
);
    localparam integer REQUEST_FIFO_DEPTH = 256;
    localparam ST_IDLE  = 1'b0;
    localparam ST_ISSUE = 1'b1;

    reg [63:0] req_addr_fifo [0:REQUEST_FIFO_DEPTH-1];
    reg [10:0] req_len_fifo  [0:REQUEST_FIFO_DEPTH-1];
    reg [3:0]  req_fbe_fifo  [0:REQUEST_FIFO_DEPTH-1];
    reg [3:0]  req_lbe_fifo  [0:REQUEST_FIFO_DEPTH-1];
    reg [15:0] req_id_fifo   [0:REQUEST_FIFO_DEPTH-1];
    reg [7:0]  req_tag_fifo  [0:REQUEST_FIFO_DEPTH-1];
    reg [2:0]  req_tc_fifo   [0:REQUEST_FIFO_DEPTH-1];
    reg [2:0]  req_attr_fifo [0:REQUEST_FIFO_DEPTH-1];
    reg        req_4dw_fifo  [0:REQUEST_FIFO_DEPTH-1];
    reg [6:0]  req_bar_fifo  [0:REQUEST_FIFO_DEPTH-1];
    reg [12:0] req_bc_fifo   [0:REQUEST_FIFO_DEPTH-1];
    reg [10:0] req_cpl_fifo  [0:REQUEST_FIFO_DEPTH-1];
    reg [7:0]  req_wr_ptr;
    reg [7:0]  req_rd_ptr;
    reg [8:0]  req_count;

    reg        state;
    reg [63:0] current_addr;
    reg [10:0] request_length_dw;
    reg [10:0] remaining_dw;
    reg [10:0] packet_length_dw;
    reg [10:0] packet_remaining_dw;
    reg [12:0] remaining_byte_count;
    reg [3:0]  request_first_be;
    reg [3:0]  request_last_be;
    reg [15:0] request_id;
    reg [7:0]  request_tag;
    reg [2:0]  request_tc;
    reg [2:0]  request_attr;
    reg        request_4dw;
    reg [6:0]  request_bar;
    reg        request_first_packet;
    reg        packet_first;

    wire request_fifo_full = (req_count == 9'd256);
    wire request_pop = (state == ST_IDLE) && (req_count != 0);
    wire request_can_accept = !request_fifo_full || request_pop;
    wire request_push = tlps_in_valid && request_can_accept;

    function automatic [2:0] popcount4(input [3:0] value);
        begin
            popcount4 = {2'b00, value[0]} + {2'b00, value[1]} +
                        {2'b00, value[2]} + {2'b00, value[3]};
        end
    endfunction

    function automatic [1:0] first_enabled_offset(input [3:0] value);
        begin
            if (value[0])      first_enabled_offset = 2'd0;
            else if (value[1]) first_enabled_offset = 2'd1;
            else if (value[2]) first_enabled_offset = 2'd2;
            else if (value[3]) first_enabled_offset = 2'd3;
            else               first_enabled_offset = 2'd0;
        end
    endfunction

    function automatic [10:0] next_completion_dw(
        input [63:0] address_in,
        input [10:0] remaining_in
    );
        integer rcb_dw;
        integer mps_dw;
        integer selected_dw;
        begin
            rcb_dw = (READ_COMPLETION_BOUNDARY -
                      (address_in & (READ_COMPLETION_BOUNDARY - 1))) / 4;
            mps_dw = MAX_PAYLOAD_BYTES / 4;
            selected_dw = remaining_in;
            if (selected_dw > rcb_dw)
                selected_dw = rcb_dw;
            if (selected_dw > mps_dw)
                selected_dw = mps_dw;
            next_completion_dw = selected_dw[10:0];
        end
    endfunction

    function automatic [12:0] completed_enabled_bytes(
        input [10:0] total_length,
        input [10:0] remaining_before,
        input [10:0] completion_length,
        input [3:0] first_enable,
        input [3:0] last_enable
    );
        reg [12:0] value;
        begin
            if (total_length == 11'd1) begin
                value = popcount4(first_enable);
            end else begin
                value = {completion_length, 2'b00};
                if (remaining_before == total_length)
                    value = value - (13'd4 - popcount4(first_enable));
                if (completion_length == remaining_before)
                    value = value - (13'd4 - popcount4(last_enable));
            end
            completed_enabled_bytes = value;
        end
    endfunction

    always @(posedge clk) begin
        if (rst) begin
            req_wr_ptr <= 8'h00;
        end else if (request_push) begin
            req_addr_fifo[req_wr_ptr] <= norm_address;
            req_len_fifo[req_wr_ptr]  <= norm_length_dw;
            req_fbe_fifo[req_wr_ptr]  <= norm_first_be;
            req_lbe_fifo[req_wr_ptr]  <= norm_last_be;
            req_id_fifo[req_wr_ptr]   <= norm_requester_id;
            req_tag_fifo[req_wr_ptr]  <= norm_tag;
            req_tc_fifo[req_wr_ptr]   <= norm_traffic_class;
            req_attr_fifo[req_wr_ptr] <= norm_attributes;
            req_4dw_fifo[req_wr_ptr]  <= norm_header_4dw;
            req_bar_fifo[req_wr_ptr]  <= norm_bar_mask;
            req_bc_fifo[req_wr_ptr]   <= norm_first_completion_byte_count;
            req_cpl_fifo[req_wr_ptr]  <= norm_first_completion_dw;
            req_wr_ptr <= req_wr_ptr + 1'b1;
        end
    end

    always @(posedge clk) begin
        if (rst) begin
            req_rd_ptr <= 8'h00;
            req_count  <= 9'h000;
        end else begin
            case ({request_push, request_pop})
                2'b10: req_count <= req_count + 1'b1;
                2'b01: req_count <= req_count - 1'b1;
                default: req_count <= req_count;
            endcase
            if (request_pop)
                req_rd_ptr <= req_rd_ptr + 1'b1;
        end
    end

    wire rd3_enable;
    wire issue_dw = (state == ST_ISSUE) && rd3_enable && !stall;
    wire packet_last = (packet_remaining_dw == 11'd1);
    wire request_last_packet = (packet_length_dw == remaining_dw);
    wire [3:0] packet_start_be = request_first_packet ? request_first_be :
                                 ((remaining_dw == 11'd1) ? request_last_be : 4'hF);
    wire [1:0] packet_first_offset = first_enabled_offset(packet_start_be);
    wire [31:0] context_address = packet_first
                                ? {current_addr[31:2], packet_first_offset}
                                : current_addr[31:0];

    assign rd_req_ctx = {
        packet_first,
        packet_last,
        packet_length_dw,
        remaining_byte_count[11:0],
        request_tc,
        request_attr,
        request_4dw,
        request_tag,
        request_id,
        context_address
    };
    assign rd_req_bar   = request_bar;
    assign rd_req_addr  = current_addr[31:0];
    assign rd_req_valid = issue_dw;

    always @(posedge clk) begin
        if (rst) begin
            state                 <= ST_IDLE;
            current_addr          <= 64'h0;
            request_length_dw     <= 11'h0;
            remaining_dw          <= 11'h0;
            packet_length_dw      <= 11'h0;
            packet_remaining_dw   <= 11'h0;
            remaining_byte_count  <= 13'h0;
            request_first_be      <= 4'h0;
            request_last_be       <= 4'h0;
            request_id            <= 16'h0;
            request_tag           <= 8'h0;
            request_tc            <= 3'h0;
            request_attr          <= 3'h0;
            request_4dw           <= 1'b0;
            request_bar           <= 7'h0;
            request_first_packet  <= 1'b0;
            packet_first          <= 1'b0;
        end else if (state == ST_IDLE) begin
            if (request_pop) begin
                current_addr         <= req_addr_fifo[req_rd_ptr];
                request_length_dw    <= req_len_fifo[req_rd_ptr];
                remaining_dw         <= req_len_fifo[req_rd_ptr];
                packet_length_dw     <= req_cpl_fifo[req_rd_ptr];
                packet_remaining_dw  <= req_cpl_fifo[req_rd_ptr];
                remaining_byte_count <= req_bc_fifo[req_rd_ptr];
                request_first_be     <= req_fbe_fifo[req_rd_ptr];
                request_last_be      <= req_lbe_fifo[req_rd_ptr];
                request_id           <= req_id_fifo[req_rd_ptr];
                request_tag          <= req_tag_fifo[req_rd_ptr];
                request_tc           <= req_tc_fifo[req_rd_ptr];
                request_attr         <= req_attr_fifo[req_rd_ptr];
                request_4dw          <= req_4dw_fifo[req_rd_ptr];
                request_bar          <= req_bar_fifo[req_rd_ptr];
                request_first_packet <= 1'b1;
                packet_first         <= 1'b1;
                state                <= ST_ISSUE;
            end
        end else if (issue_dw) begin
            if (!packet_last) begin
                current_addr        <= current_addr + 64'd4;
                packet_remaining_dw <= packet_remaining_dw - 1'b1;
                packet_first        <= 1'b0;
            end else if (request_last_packet) begin
                state        <= ST_IDLE;
                packet_first <= 1'b0;
            end else begin
                current_addr <= current_addr + 64'd4;
                remaining_dw <= remaining_dw - packet_length_dw;
                remaining_byte_count <= remaining_byte_count - completed_enabled_bytes(
                    request_length_dw, remaining_dw, packet_length_dw,
                    request_first_be, request_last_be);
                packet_length_dw <= next_completion_dw(
                    current_addr + 64'd4, remaining_dw - packet_length_dw);
                packet_remaining_dw <= next_completion_dw(
                    current_addr + 64'd4, remaining_dw - packet_length_dw);
                request_first_packet <= 1'b0;
                packet_first         <= 1'b1;
            end
        end
    end

    wire        rd_rsp_first       = rd_rsp_ctx[87];
    wire        rd_rsp_last        = rd_rsp_ctx[86];
    wire [10:0] rd_rsp_dwlen       = rd_rsp_ctx[85:75];
    wire [11:0] rd_rsp_byte_count  = rd_rsp_ctx[74:63];
    wire [2:0]  rd_rsp_tc          = rd_rsp_ctx[62:60];
    wire [2:0]  rd_rsp_attr        = rd_rsp_ctx[59:57];
    wire        rd_rsp_4dw         = rd_rsp_ctx[56];
    wire [7:0]  rd_rsp_tag         = rd_rsp_ctx[55:48];
    wire [15:0] rd_rsp_requester   = rd_rsp_ctx[47:32];
    wire [6:0]  rd_rsp_lower_addr  = rd_rsp_ctx[6:0];
    wire [31:0] rd_rsp_data_swapped = {
        rd_rsp_data[7:0], rd_rsp_data[15:8],
        rd_rsp_data[23:16], rd_rsp_data[31:24]
    };

    bit [127:0] tdata;
    bit [3:0]   tkeepdw;
    bit         tlast;
    bit         first;
    wire        tvalid = tlast || tkeepdw[3];

    always @(posedge clk) begin
        if (rst) begin
            tdata   <= 128'h0;
            tkeepdw <= 4'h0;
            tlast   <= 1'b0;
            first   <= 1'b0;
        end else if (rd_rsp_valid && rd_rsp_first) begin
            tdata[31:29]      <= 3'b010;
            tdata[28:24]      <= 5'b01010;
            tdata[23]         <= 1'b0;
            tdata[22:20]      <= rd_rsp_tc;
            tdata[19]         <= 1'b0;
            tdata[18]         <= rd_rsp_attr[2];
            tdata[17:14]      <= 4'b0000;
            tdata[13:12]      <= rd_rsp_attr[1:0];
            tdata[11:10]      <= 2'b00;
            tdata[9:0]        <= rd_rsp_dwlen[9:0];
            tdata[63:32]      <= {pcie_id[7:0], pcie_id[15:8], 4'b0000,
                                  rd_rsp_byte_count};
            tdata[95:64]      <= {rd_rsp_requester, rd_rsp_tag, 1'b0,
                                  rd_rsp_lower_addr};
            tdata[127:96]     <= rd_rsp_data_swapped;
            tkeepdw           <= 4'b1111;
            tlast             <= rd_rsp_last;
            first             <= 1'b1;
        end else begin
            tlast   <= rd_rsp_valid && rd_rsp_last;
            tkeepdw <= tvalid ? (rd_rsp_valid ? 4'b0001 : 4'b0000) :
                       (rd_rsp_valid ? ((tkeepdw << 1) | 1'b1) : tkeepdw);
            first <= 1'b0;
            if (rd_rsp_valid) begin
                if (tvalid || !tkeepdw[0]) tdata[31:0]   <= rd_rsp_data_swapped;
                if (!tkeepdw[1])           tdata[63:32]  <= rd_rsp_data_swapped;
                if (!tkeepdw[2])           tdata[95:64]  <= rd_rsp_data_swapped;
                if (!tkeepdw[3])           tdata[127:96] <= rd_rsp_data_swapped;
            end
        end
    end

    wire unused_rd_rsp_4dw = rd_rsp_4dw;

    fifo_134_134_clk1_bar_rdrsp i_fifo_134_134_clk1_bar_rdrsp(
        .srst       ( rst ),
        .clk        ( clk ),
        .din        ( {first, tlast, tkeepdw, tdata} ),
        .wr_en      ( tvalid ),
        .rd_en      ( tlps_out.tready ),
        .dout       ( {tlps_out.tuser[0], tlps_out.tlast,
                       tlps_out.tkeepdw, tlps_out.tdata} ),
        .full       ( ),
        .empty      ( ),
        .prog_empty ( rd3_enable ),
        .valid      ( tlps_out.tvalid )
    );
    assign tlps_out.tuser[8:2] = 7'h00;
    assign tlps_out.tuser[1]   = tlps_out.tlast;
    assign tlps_out.has_data   = tlps_out.tvalid;
endmodule
