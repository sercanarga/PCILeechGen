`timescale 1ns / 1ps
`include "pcileech_header.svh"
`default_nettype none
module tb_top;
    reg clk = 0;
    reg rst = 1;
    always #5 clk = ~clk;

    wire [15:0] cfg_command = 16'h0006;
    wire [1:0]  cfg_power_state = 2'b00;
    wire        cfg_flr_in_process = 1'b0;
    wire        cfg_to_turnoff = 1'b0;
    wire        cfg_link_up = 1'b1;
    wire        cfg_msi_enable = 1'b0;
    wire        cfg_msix_enable = 1'b1;
    wire        cfg_msix_function_mask = 1'b0;
    wire [15:0] pcie_id = 16'h0000;
    wire        bar_en = 1'b1;

    reg  [127:0] tlps_in_tdata = 0;
    reg  [3:0]   tlps_in_tkeepdw = 0;
    reg          tlps_in_tvalid = 0;
    reg          tlps_in_tlast = 0;
    reg  [8:0]   tlps_in_tuser = 0;

    wire [127:0] tlps_out_tdata;
    wire [3:0]   tlps_out_tkeepdw;
    wire         tlps_out_tvalid;
    wire         tlps_out_tlast;
    wire [8:0]   tlps_out_tuser;
    wire         tlps_out_has_data;

    wire [127:0] tlps_dma_out_tdata;
    wire         tlps_dma_out_tvalid;
    wire         tlps_dma_out_tlast;
    wire [8:0]   tlps_dma_out_tuser;
    wire         tlps_dma_out_has_data;
    wire         intr_req;

    reg [31:0] host_mem [0:65535];
    integer i;

    reg         host_poke_valid = 1'b0;
    reg  [15:0] host_poke_addr = 16'h0;
    reg  [31:0] host_poke_data = 32'h0;
    reg  [15:0] host_peek_addr = 16'h0;
    wire [31:0] host_peek_data = host_mem[host_peek_addr];

    IfAXIS128 tlps_in_if();
    IfAXIS128 tlps_out_if();
    IfAXIS128 tlps_dma_out_if();

    assign tlps_in_if.tdata   = cpld_in_tvalid ? cpld_in_tdata   : tlps_in_tdata;
    assign tlps_in_if.tkeepdw = cpld_in_tvalid ? cpld_in_tkeepdw : tlps_in_tkeepdw;
    assign tlps_in_if.tvalid  = cpld_in_tvalid ? 1'b1             : tlps_in_tvalid;
    assign tlps_in_if.tlast   = cpld_in_tvalid ? cpld_in_tlast    : tlps_in_tlast;
    assign tlps_in_if.tuser   = cpld_in_tvalid ? cpld_in_tuser    : tlps_in_tuser;

    assign tlps_out_tdata = tlps_out_if.tdata;
    assign tlps_out_tkeepdw = tlps_out_if.tkeepdw;
    assign tlps_out_tvalid = tlps_out_if.tvalid;
    assign tlps_out_tlast = tlps_out_if.tlast;
    assign tlps_out_tuser = tlps_out_if.tuser;
    assign tlps_out_has_data = tlps_out_if.has_data;

    assign tlps_dma_out_tdata = tlps_dma_out_if.tdata;
    assign tlps_dma_out_tvalid = tlps_dma_out_if.tvalid;
    assign tlps_dma_out_tlast = tlps_dma_out_if.tlast;
    assign tlps_dma_out_tuser = tlps_dma_out_if.tuser;
    assign tlps_dma_out_has_data = tlps_dma_out_if.has_data;

    assign tlps_out_if.tready = 1'b1;
    assign tlps_dma_out_if.tready = 1'b1;

    function [31:0] swap32(input [31:0] v);
        swap32 = {v[7:0], v[15:8], v[23:16], v[31:24]};
    endfunction

    reg [127:0] cpld_in_tdata = 0;
    reg [3:0]   cpld_in_tkeepdw = 0;
    reg         cpld_in_tvalid = 0;
    reg         cpld_in_tlast = 0;
    reg [8:0]   cpld_in_tuser = 0;

    localparam [1:0] LB_IDLE = 2'd0, LB_MWR_DATA = 2'd1;
    reg [1:0]  lb_state = LB_IDLE;
    reg [63:0] mwr_addr_q;

    wire        dma_first = tlps_dma_out_if.tvalid && tlps_dma_out_if.tuser[0];
    wire [2:0]  dma_fmt   = tlps_dma_out_if.tdata[31:29];

    always @(posedge clk) begin
        if (!rst && host_poke_valid)
            host_mem[host_poke_addr] <= host_poke_data;
    end

    always @(posedge clk) begin
        if (rst) begin
            lb_state      <= LB_IDLE;
            cpld_in_tvalid<= 1'b0;
            cpld_in_tlast <= 1'b0;
            cpld_in_tuser <= 9'h0;
            cpld_in_tkeepdw <= 4'h0;
            cpld_in_tdata  <= 128'h0;
        end else begin
            cpld_in_tvalid <= 1'b0;
            cpld_in_tlast  <= 1'b0;
            case (lb_state)
                LB_IDLE: begin
                    if (dma_first) begin
                        if (dma_fmt == 3'b001) begin
                            cpld_in_tdata[31:0]   <= (3'b010 << 29) | (5'b01010 << 24) | 10'd1;
                            cpld_in_tdata[63:32]  <= (16'h0000 << 16) | (3'b000 << 13) | 12'd4;
                            cpld_in_tdata[95:64]  <= {tlps_dma_out_if.tdata[63:48],
                                                       tlps_dma_out_if.tdata[79:72],
                                                       tlps_dma_out_if.tdata[102:96]};
                            cpld_in_tdata[127:96] <= swap32(host_mem[tlps_dma_out_if.tdata[111:98]]);
                            cpld_in_tuser   <= 9'h001;
                            cpld_in_tlast   <= 1'b1;
                            cpld_in_tkeepdw <= 4'hF;
                            cpld_in_tvalid  <= 1'b1;
                        end else if (dma_fmt == 3'b011) begin
                            mwr_addr_q <= {tlps_dma_out_if.tdata[95:64],
                                           tlps_dma_out_if.tdata[127:96] & 32'hFFFF_FFFC};
                            lb_state <= LB_MWR_DATA;
                        end
                    end
                end
                LB_MWR_DATA: begin
                    if (tlps_dma_out_if.tvalid) begin
                        host_mem[mwr_addr_q[15:2]] <= swap32(tlps_dma_out_if.tdata[127:96]);
                        lb_state <= LB_IDLE;
                    end
                end
                default: lb_state <= LB_IDLE;
            endcase
        end
    end

    initial begin
        for (i = 0; i < 65536; i = i + 1)
            host_mem[i] = 32'h0;
    end

    pcileech_tlps128_bar_controller i_bar(
        .rst(rst), .clk(clk), .bar_en(bar_en), .pcie_id(pcie_id),
        .cfg_command(cfg_command), .cfg_power_state(cfg_power_state),
        .cfg_flr_in_process(cfg_flr_in_process), .cfg_to_turnoff(cfg_to_turnoff),
        .cfg_link_up(cfg_link_up), .cfg_msi_enable(cfg_msi_enable),
        .cfg_msix_enable(cfg_msix_enable), .cfg_msix_function_mask(cfg_msix_function_mask),
        .flash_csn(), .flash_sdi_dq0(), .flash_sdo_dq1(1'b0),
        .flash_wpn_dq2(), .flash_hldn_dq3(),
        .tlps_in(tlps_in_if.sink_lite),
        .tlps_out(tlps_out_if.source),
        .tlps_dma_out(tlps_dma_out_if.source),
        .intr_req(intr_req)
    );
endmodule
