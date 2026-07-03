`default_nettype none
module cfg_w1c_shadow (
    input  wire [31:0] cur_val,
    input  wire [31:0] wr_data,
    input  wire [31:0] wr_mask,
    input  wire [31:0] w1c_mask,
    output wire [31:0] new_val
);
    wire [31:0] ro_keep   = cur_val & ~wr_mask & ~w1c_mask;
    wire [31:0] rw_take   = wr_data & wr_mask & ~w1c_mask;
    wire [31:0] w1c_keep  = cur_val & w1c_mask & ~wr_data;
    assign new_val = ro_keep | rw_take | w1c_keep;
endmodule
`default_nettype wire
