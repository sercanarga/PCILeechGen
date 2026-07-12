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
    wire        cfg_msix_enable = 1'b0;
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
    wire         intr_req;

    IfAXIS128 tlps_in_if();
    IfAXIS128 tlps_out_if();
    IfAXIS128 tlps_dma_out_if();

    assign tlps_in_if.tdata   = tlps_in_tdata;
    assign tlps_in_if.tkeepdw = tlps_in_tkeepdw;
    assign tlps_in_if.tvalid  = tlps_in_tvalid;
    assign tlps_in_if.tlast   = tlps_in_tlast;
    assign tlps_in_if.tuser   = tlps_in_tuser;

    assign tlps_out_tdata = tlps_out_if.tdata;
    assign tlps_out_tkeepdw = tlps_out_if.tkeepdw;
    assign tlps_out_tvalid = tlps_out_if.tvalid;
    assign tlps_out_tlast = tlps_out_if.tlast;
    assign tlps_out_tuser = tlps_out_if.tuser;
    assign tlps_out_has_data = tlps_out_if.has_data;

    assign tlps_out_if.tready = 1'b1;
    assign tlps_dma_out_if.tready = 1'b1;

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
