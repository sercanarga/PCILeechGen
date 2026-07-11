//
// BAR implementation for MSI doorbell interrupt.
// Write to offset 0x00 triggers a single-cycle MSI interrupt pulse.
// Read from offset 0x00 returns the last written value.
//
// This provides a software-triggerable MSI interrupt mechanism: any
// driver that writes to BAR2 offset 0x00 will cause the FPGA to
// generate an MSI interrupt to the host.
//

`timescale 1ns / 1ps

module pcileech_bar_impl_msi(
    input               rst,
    input               clk,

    input [31:0]        wr_addr,
    input [31:0]        wr_data,
    input [3:0]         wr_be,
    input               wr_valid,

    input [87:0]        rd_req_ctx,
    input [31:0]        rd_req_addr,
    input               rd_req_valid,

    output reg [87:0]   rd_rsp_ctx,
    output reg [31:0]   rd_rsp_data,
    output reg          rd_rsp_valid,

    output              intr_req
);

    reg [31:0]      doorbell;
    reg             intr_req_reg;
    reg [31:0]      doorbell_q;
    reg             rd_req_valid_q;

    // Read response: 2-cycle latency
    always @ ( posedge clk ) begin
        doorbell_q          <= doorbell;
        rd_req_valid_q      <= rd_req_valid;
        rd_rsp_ctx          <= rd_req_ctx;
        rd_rsp_data         <= doorbell_q;
        rd_rsp_valid        <= rd_req_valid_q;
    end

    // Write handling: capture doorbell value and pulse intr_req
    always @ ( posedge clk ) begin
        if ( rst ) begin
            doorbell    <= 32'h0;
            intr_req_reg <= 1'b0;
        end
        else if ( wr_valid ) begin
            if ( wr_be[0] ) doorbell[7:0]   <= wr_data[7:0];
            if ( wr_be[1] ) doorbell[15:8]  <= wr_data[15:8];
            if ( wr_be[2] ) doorbell[23:16] <= wr_data[23:16];
            if ( wr_be[3] ) doorbell[31:24] <= wr_data[31:24];
            intr_req_reg <= 1'b1;
        end
        else begin
            intr_req_reg <= 1'b0;
        end
    end

    assign intr_req = intr_req_reg;

endmodule
