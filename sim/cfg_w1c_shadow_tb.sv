`default_nettype none
`timescale 1ns/1ps
module cfg_w1c_shadow_tb;
    reg  [31:0] cur, data, wmask, w1c;
    wire [31:0] nv;
    integer errors = 0;

    cfg_w1c_shadow dut (.cur_val(cur), .wr_data(data), .wr_mask(wmask), .w1c_mask(w1c), .new_val(nv));

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
        cur = 32'hF900_0000; data = 32'h0800_0000; wmask = 32'hFFFF_0006; w1c = 32'hF900_0000;
        #1; check("w1c clears written bit", nv, 32'hF100_0000);

        cur = 32'hF900_0000; data = 32'h0000_0000; wmask = 32'hFFFF_0006; w1c = 32'hF900_0000;
        #1; check("w1c keeps unwritten bits", nv, 32'hF900_0000);

        cur = 32'hF900_0000; data = 32'hF900_0000; wmask = 32'hFFFF_0006; w1c = 32'hF900_0000;
        #1; check("w1c full clear", nv, 32'h0000_0000);

        cur = 32'h0000_0000; data = 32'h0000_0002; wmask = 32'hFFFF_0006; w1c = 32'hF900_0000;
        #1; check("rw bit takes data", nv, 32'h0000_0002);

        cur = 32'h0000_0000; data = 32'h0000_0001; wmask = 32'hFFFF_0006; w1c = 32'hF900_0000;
        #1; check("ro bit ignored", nv, 32'h0000_0000);

        cur = 32'hF900_0000; data = 32'h0800_0002; wmask = 32'hFFFF_0006; w1c = 32'hF900_0000;
        #1; check("mixed w1c+rw", nv, 32'hF100_0002);

        if (errors == 0) $display("ALL TESTS PASSED");
        else $display("%0d FAILURES", errors);
        $finish;
    end
endmodule
`default_nettype wire
