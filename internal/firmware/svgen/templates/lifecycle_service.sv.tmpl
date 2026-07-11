`timescale 1ns / 1ps

module pcileech_lifecycle_service(
    input              clk,
    input              fundamental_reset,
    input              flr_in_process,
    input [1:0]        pm_dstate,
    input              memory_space_enable,
    input              bus_master_enable,
    input              turnoff_pending,
    input              link_up,
    output             device_reset,
    output             io_enabled,
    output             dma_enabled,
    output             quiesce,
    output reg [31:0]  generation
);
    wire active = !device_reset && (pm_dstate == 2'b00) &&
                  !turnoff_pending && link_up;

    assign device_reset = fundamental_reset || flr_in_process;
    assign io_enabled   = active && memory_space_enable;
    assign dma_enabled  = active && bus_master_enable;
    assign quiesce      = !dma_enabled;

    reg reset_d = 1'b0;
    initial generation = 32'd0;
    always @(posedge clk) begin
        reset_d <= device_reset;
        if (device_reset && !reset_d)
            generation <= generation + 1'b1;
    end
endmodule
