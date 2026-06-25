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
