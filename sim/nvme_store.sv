`default_nettype none
module nvme_store #(
    parameter NSECT   = 4,
    parameter SECT_DW = 8
) (
    input  wire clk,
    input  wire [$clog2(NSECT)-1:0]   wr_sector,
    input  wire [$clog2(SECT_DW)-1:0] wr_dw,
    input  wire [31:0]                 wr_data,
    input  wire                        wr_en,
    input  wire [$clog2(NSECT)-1:0]   rd_sector,
    input  wire [$clog2(SECT_DW)-1:0] rd_dw,
    input  wire                        rd_en,
    output reg  [31:0]                 rd_data,
    output reg                         rd_valid
);
    localparam DEPTH = NSECT * SECT_DW;
    localparam AW    = $clog2(NSECT) + $clog2(SECT_DW);

    reg [31:0] mem [0:DEPTH-1];
    integer i;
    initial begin
        for (i = 0; i < DEPTH; i = i + 1)
            mem[i] = 32'h0000_0000;
    end

    wire [AW-1:0] wr_addr = {wr_sector, wr_dw};
    wire [AW-1:0] rd_addr = {rd_sector, rd_dw};

    always @(posedge clk) begin
        if (wr_en)
            mem[wr_addr] <= wr_data;
        rd_valid <= rd_en;
        if (rd_en)
            rd_data <= mem[rd_addr];
    end
endmodule
`default_nettype wire
