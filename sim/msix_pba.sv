`default_nettype none
module msix_pba #(parameter NVEC = 4) (
    input  wire                     clk,
    input  wire                     rst_n,
    input  wire [NVEC-1:0]          vector_mask,
    input  wire                     req_valid,
    input  wire [$clog2(NVEC)-1:0] req_vec,
    output reg  [NVEC-1:0]          pba,
    output reg                      deliver_valid,
    output reg  [$clog2(NVEC)-1:0] deliver_vec
);
    localparam VW = $clog2(NVEC);

    reg     found;
    integer i;

    always @(posedge clk or negedge rst_n) begin
        if (!rst_n) begin
            pba           <= {NVEC{1'b0}};
            deliver_valid <= 1'b0;
            deliver_vec   <= {VW{1'b0}};
        end else begin
            deliver_valid <= 1'b0;
            deliver_vec   <= {VW{1'b0}};

            if (req_valid && vector_mask[req_vec]) begin
                pba[req_vec] <= 1'b1;
            end else if (req_valid) begin
                deliver_valid <= 1'b1;
                deliver_vec   <= req_vec;
            end else begin
                found = 1'b0;
                for (i = 0; i < NVEC; i = i + 1) begin
                    if (!found && !vector_mask[i] && pba[i]) begin
                        deliver_valid <= 1'b1;
                        deliver_vec   <= i[VW-1:0];
                        pba[i]        <= 1'b0;
                        found          = 1'b1;
                    end
                end
            end
        end
    end
endmodule
`default_nettype wire
