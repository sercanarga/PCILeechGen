`default_nettype none
`timescale 1ns/1ps
module nvme_store_tb;
    localparam NSECT   = 4;
    localparam SECT_DW = 8;

    reg clk = 0;
    always #5 clk = ~clk;

    reg         wr_en     = 0;
    reg  [1:0]  wr_sector = 0;
    reg  [2:0]  wr_dw     = 0;
    reg  [31:0] wr_data   = 0;
    reg         rd_en     = 0;
    reg  [1:0]  rd_sector = 0;
    reg  [2:0]  rd_dw     = 0;
    wire [31:0] rd_data;
    wire        rd_valid;

    nvme_store #(.NSECT(NSECT), .SECT_DW(SECT_DW)) dut (
        .clk(clk),
        .wr_en(wr_en), .wr_sector(wr_sector), .wr_dw(wr_dw), .wr_data(wr_data),
        .rd_en(rd_en), .rd_sector(rd_sector), .rd_dw(rd_dw),
        .rd_data(rd_data), .rd_valid(rd_valid)
    );

    integer errors = 0;

    task check;
        input [255:0] name;
        input [31:0]  got;
        input [31:0]  exp;
        begin
            if (got !== exp) begin
                $display("FAIL %0s: got %08x exp %08x", name, got, exp);
                errors = errors + 1;
            end else begin
                $display("ok   %0s = %08x", name, got);
            end
        end
    endtask

    task do_write;
        input [1:0]  sect;
        input [2:0]  dw;
        input [31:0] data;
        begin
            wr_en = 1; wr_sector = sect; wr_dw = dw; wr_data = data;
            @(posedge clk); #1;
            wr_en = 0;
        end
    endtask

    task do_read;
        input [1:0] sect;
        input [2:0] dw;
        begin
            rd_en = 1; rd_sector = sect; rd_dw = dw;
            @(posedge clk); #1;
            rd_en = 0;
        end
    endtask

    initial begin
        @(posedge clk); #1;

        do_write(0, 0, 32'hDEAD_BEEF);
        do_read(0, 0);
        check("write-readback",   rd_data,           32'hDEAD_BEEF);
        check("rd_valid-high",    {31'd0, rd_valid},  32'h0000_0001);

        do_read(2, 5);
        check("unwritten-zero",   rd_data,            32'h0000_0000);

        do_write(0, 0, 32'hCAFE_F00D);
        do_read(1, 0);
        check("no-alias-sector",  rd_data,            32'h0000_0000);

        do_write(0, 0, 32'h1234_5678);
        do_write(0, 0, 32'hABCD_EF01);
        do_read(0, 0);
        check("overwrite-new",    rd_data,            32'hABCD_EF01);

        if (errors == 0) $display("ALL TESTS PASSED");
        else $display("%0d FAILURES", errors);
        $finish;
    end
endmodule
`default_nettype wire
