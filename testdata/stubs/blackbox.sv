// stub xilinx modules so verilator can resolve svgen port refs
`timescale 1ns/1ps
`default_nettype none

module pcileech_pcie_cfg_a7 (
    input  wire        clk,
    input  wire        rst
);
endmodule

module pcileech_fifo (
    input  wire        clk,
    input  wire        rst
);
endmodule

module fifo_134_134_clk1_bar_rdrsp (
    input  wire         srst,
    input  wire         clk,
    input  wire [133:0] din,
    input  wire         wr_en,
    input  wire         rd_en,
    output wire [133:0] dout,
    output wire         full,
    output wire         empty,
    output wire         prog_empty,
    output wire         valid
);
endmodule

module fifo_141_141_clk1_bar_wr (
    input  wire         srst,
    input  wire         clk,
    input  wire [140:0] din,
    input  wire         wr_en,
    input  wire         rd_en,
    output wire [140:0] dout,
    output wire         full,
    output wire         empty,
    output wire         prog_empty,
    output wire         valid
);
endmodule

module fifo_74_74_clk1_bar_rd1 (
    input  wire        srst,
    input  wire        clk,
    input  wire [73:0] din,
    input  wire        wr_en,
    input  wire        rd_en,
    output wire [73:0] dout,
    output wire        full,
    output wire        empty,
    output wire        valid
);
endmodule
