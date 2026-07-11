module fifo_134_134_clk1_bar_rdrsp(
    input clk, input srst,
    input [133:0] din, input wr_en,
    output [133:0] dout, output full, output empty,
    output prog_full, input rd_en,
    output [10:0] rd_data_count, output [10:0] wr_data_count
);
    reg [133:0] mem [0:1023];
    reg [10:0] wr_ptr = 0, rd_ptr = 0;
    reg [10:0] count = 0;
    assign full = (count == 11'd1024);
    assign empty = (count == 11'd0);
    assign prog_full = (count >= 11'd1008);
    assign dout = mem[rd_ptr];
    assign rd_data_count = count;
    assign wr_data_count = count;
    always @(posedge clk) begin
        if (srst) begin
            wr_ptr <= 0; rd_ptr <= 0; count <= 0;
        end else begin
            if (wr_en && !full) begin
                mem[wr_ptr] <= din;
                wr_ptr <= wr_ptr + 1;
                count <= count + (rd_en && !empty ? 0 : 1);
            end else if (rd_en && !empty) begin
                rd_ptr <= rd_ptr + 1;
                count <= count - 1;
            end
        end
    end
endmodule
