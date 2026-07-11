`timescale 1ns / 1ps
`default_nettype none

module pcileech_bar_rsp_arbiter (
    input  wire          clk,
    input  wire          rst,
    input  wire [6:0]    in_valid,
    input  wire [87:0]   in_ctx [0:6],
    input  wire [31:0]   in_data [0:6],
    output wire [6:0]    in_ready,
    output wire          out_valid,
    output wire [87:0]   out_ctx,
    output wire [31:0]   out_data
);
    localparam integer FIFO_DEPTH = 32;
    reg [87:0] ctx_mem [0:6][0:FIFO_DEPTH-1];
    reg [31:0] data_mem [0:6][0:FIFO_DEPTH-1];
    reg [4:0] wr_ptr [0:6];
    reg [4:0] rd_ptr [0:6];
    reg [5:0] count [0:6];
    reg [2:0] rr_ptr;
    reg [2:0] selected;
    reg selected_valid;
    wire [6:0] push = in_valid & in_ready;
    genvar g;
    generate
        for (g = 0; g < 7; g = g + 1) begin : g_ready
            assign in_ready[g] = count[g] < 6'd32;
        end
    endgenerate
    integer scan;
    reg [3:0] index;
    always @* begin
        selected = 3'd0;
        selected_valid = 1'b0;
        for (scan = 0; scan < 7; scan = scan + 1) begin
            index = {1'b0, rr_ptr} + scan[3:0];
            if (index >= 4'd7)
                index = index - 4'd7;
            if (!selected_valid && count[index[2:0]] != 0) begin
                selected = index[2:0];
                selected_valid = 1'b1;
            end
        end
    end
    assign out_valid = selected_valid;
    assign out_ctx = selected_valid ? ctx_mem[selected][rd_ptr[selected]] : 88'h0;
    assign out_data = selected_valid ? data_mem[selected][rd_ptr[selected]] : 32'h0;
    integer k;
    always @(posedge clk) begin
        if (rst) begin
            rr_ptr <= 3'd0;
            for (k = 0; k < 7; k = k + 1) begin
                wr_ptr[k] <= 5'd0;
                rd_ptr[k] <= 5'd0;
                count[k] <= 6'd0;
            end
        end else begin
            for (k = 0; k < 7; k = k + 1) begin
                if (push[k]) begin
                    ctx_mem[k][wr_ptr[k]] <= in_ctx[k];
                    data_mem[k][wr_ptr[k]] <= in_data[k];
                    wr_ptr[k] <= wr_ptr[k] + 5'd1;
                end
                if (selected_valid && selected == k[2:0])
                    rd_ptr[k] <= rd_ptr[k] + 5'd1;
                case ({push[k], selected_valid && selected == k[2:0]})
                    2'b10: count[k] <= count[k] + 6'd1;
                    2'b01: count[k] <= count[k] - 6'd1;
                    default: count[k] <= count[k];
                endcase
            end
            if (selected_valid)
                rr_ptr <= selected == 3'd6 ? 3'd0 : selected + 3'd1;
        end
    end
endmodule

`default_nettype wire
