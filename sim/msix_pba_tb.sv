`default_nettype none
`timescale 1ns/1ps
module msix_pba_tb;
    localparam NVEC = 4;
    localparam VW   = $clog2(NVEC);

    reg             clk         = 0;
    reg             rst_n       = 0;
    reg [NVEC-1:0]  vector_mask = 0;
    reg             req_valid   = 0;
    reg [VW-1:0]    req_vec     = 0;
    wire [NVEC-1:0] pba;
    wire            deliver_valid;
    wire [VW-1:0]   deliver_vec;

    integer errors = 0;

    msix_pba #(.NVEC(NVEC)) dut (
        .clk(clk), .rst_n(rst_n),
        .vector_mask(vector_mask),
        .req_valid(req_valid), .req_vec(req_vec),
        .pba(pba), .deliver_valid(deliver_valid), .deliver_vec(deliver_vec)
    );

    always #5 clk = ~clk;

    task check(input [255:0] name, input [31:0] got, input [31:0] exp);
        begin
            if (got !== exp) begin
                $display("FAIL %0s: got %08x exp %08x", name, got, exp);
                errors = errors + 1;
            end else begin
                $display("ok   %0s = %08x", name, got);
            end
        end
    endtask

    initial begin
        rst_n = 0;
        repeat (2) @(posedge clk);
        rst_n = 1;
        @(posedge clk); #1;
        check("reset clears pba", {28'b0, pba}, 32'h0);

        vector_mask = 4'b0010;
        req_valid = 1; req_vec = 2'd1;
        @(posedge clk); #1;
        req_valid = 0;
        check("T1 pba[1] set on masked req",    {31'b0, pba[1]},        32'h1);
        check("T1 deliver_valid=0 (no deliver)",{31'b0, deliver_valid}, 32'h0);

        vector_mask = 4'b0000;
        req_valid = 1; req_vec = 2'd2;
        @(posedge clk); #1;
        req_valid = 0;
        check("T2 deliver_valid=1",             {31'b0, deliver_valid}, 32'h1);
        check("T2 deliver_vec=2",               {30'b0, deliver_vec},   32'h2);
        check("T2 pba[2] stays 0",              {31'b0, pba[2]},        32'h0);

        check("T3 pre: pba[1] still pending",   {31'b0, pba[1]},        32'h1);
        @(posedge clk); #1;
        check("T3 deliver_valid=1 on unmask",   {31'b0, deliver_valid}, 32'h1);
        check("T3 deliver_vec=1",               {30'b0, deliver_vec},   32'h1);
        check("T3 pba[1] cleared",              {31'b0, pba[1]},        32'h0);
        @(posedge clk); #1;
        check("T3 deliver_valid falls next cycle",{31'b0,deliver_valid},32'h0);

        vector_mask = 4'b1111;
        req_valid = 1; req_vec = 2'd0; @(posedge clk); #1; req_valid = 0;
        req_valid = 1; req_vec = 2'd3; @(posedge clk); #1; req_valid = 0;
        check("T4 pre: pba has pending bits",   {31'b0, |pba},          32'h1);
        rst_n = 0; @(posedge clk); #1;
        rst_n = 1; @(posedge clk); #1;
        check("T4 pba cleared by reset",        {28'b0, pba},           32'h0);

        if (errors == 0) $display("ALL TESTS PASSED");
        else $display("%0d FAILURES", errors);
        $finish;
    end
endmodule
`default_nettype wire
