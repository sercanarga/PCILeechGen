`timescale 1ns / 1ps
`default_nettype none

module pcileech_pcie_cfg_a7 (input wire clk, input wire rst);
endmodule

module pcileech_fifo (input wire clk, input wire rst);
endmodule

module fifo_134_134_clk1_bar_rdrsp (
    input wire srst, input wire clk,
    input wire [133:0] din, input wire wr_en, input wire rd_en,
    output wire [133:0] dout,
    output wire full, output wire empty,
    output wire prog_empty,
    output wire [10:0] rd_data_count, output wire [10:0] wr_data_count,
    output wire prog_full,
    output wire valid
);
    // Standard FIFO (registered dout/valid, one-cycle read latency) to match the .xci.
    reg [133:0] mem [0:1023];
    reg [10:0] wr_ptr, rd_ptr;
    reg [10:0] count;
    reg [133:0] dout_r;
    reg         valid_r;
    wire do_wr = wr_en & ~full;
    wire do_rd = rd_en & ~empty;
    assign full = (count == 11'd1024);
    assign empty = (count == 11'd0);
    assign prog_full = 1'b0;
    assign prog_empty = 1'b1;
    assign dout = dout_r;
    assign valid = valid_r;
    assign rd_data_count = count;
    assign wr_data_count = count;
    always @(posedge clk) begin
        if (srst) begin
            wr_ptr <= 0; rd_ptr <= 0; count <= 0;
            dout_r <= 134'h0; valid_r <= 1'b0;
        end else begin
            if (do_wr) begin mem[wr_ptr] <= din; wr_ptr <= wr_ptr + 1; end
            if (do_rd) begin
                dout_r  <= mem[rd_ptr];
                rd_ptr  <= rd_ptr + 1;
                valid_r <= 1'b1;
            end else begin
                valid_r <= 1'b0;
            end
            case ({do_wr, do_rd})
                2'b10: count <= count + 1;
                2'b01: count <= count - 1;
                default: count <= count;
            endcase
        end
    end
endmodule

module fifo_134_134_clk2 (input wire srst, input wire clk, input wire [133:0] din, input wire wr_en, input wire rd_en, output wire [133:0] dout, output wire full, output wire empty, output wire prog_empty, output wire valid);
    assign dout = 0; assign full = 0; assign empty = 1; assign prog_empty = 1; assign valid = 0;
endmodule

module fifo_134_134_clk2_rxfifo (input wire srst, input wire clk, input wire [133:0] din, input wire wr_en, input wire rd_en, output wire [133:0] dout, output wire full, output wire empty, output wire prog_empty, output wire valid);
    assign dout = 0; assign full = 0; assign empty = 1; assign prog_empty = 1; assign valid = 0;
endmodule

module fifo_141_141_clk1_bar_wr (
    input wire srst, input wire clk,
    input wire [140:0] din, input wire wr_en, input wire rd_en,
    output wire [140:0] dout,
    output wire full, output wire empty,
    output wire prog_empty,
    output wire valid
);
    // Standard FIFO (registered dout/valid, one-cycle read latency) to match the .xci.
    reg [140:0] mem [0:1023];
    reg [10:0] wr_ptr, rd_ptr;
    reg [10:0] count;
    reg [140:0] dout_r;
    reg         valid_r;
    wire do_wr = wr_en & ~full;
    wire do_rd = rd_en & ~empty;
    assign full = (count == 11'd1024);
    assign empty = (count == 11'd0);
    assign prog_empty = 1'b1;
    assign dout = dout_r;
    assign valid = valid_r;
    always @(posedge clk) begin
        if (srst) begin
            wr_ptr <= 0; rd_ptr <= 0; count <= 0;
            dout_r <= 0; valid_r <= 1'b0;
        end else begin
            if (do_wr) begin mem[wr_ptr] <= din; wr_ptr <= wr_ptr + 1; end
            if (do_rd) begin
                dout_r  <= mem[rd_ptr];
                rd_ptr  <= rd_ptr + 1;
                valid_r <= 1'b1;
            end else begin
                valid_r <= 1'b0;
            end
            case ({do_wr, do_rd})
                2'b10: count <= count + 1;
                2'b01: count <= count - 1;
                default: count <= count;
            endcase
        end
    end
endmodule

module fifo_74_74_clk1_bar_rd1 (input wire srst, input wire clk, input wire [73:0] din, input wire wr_en, input wire rd_en, output wire [73:0] dout, output wire full, output wire empty, output wire valid);
    assign dout = 0; assign full = 0; assign empty = 1; assign valid = 0;
endmodule
