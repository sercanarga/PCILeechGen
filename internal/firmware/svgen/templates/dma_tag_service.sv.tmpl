`timescale 1ns / 1ps

module pcileech_dma_tag_service #(
    parameter [7:0] TAG_FIRST = 8'h10,
    parameter integer TAG_COUNT = 16,
    parameter integer TIMEOUT_WIDTH = 24,
    parameter [TIMEOUT_WIDTH-1:0] TIMEOUT_CYCLES = {TIMEOUT_WIDTH{1'b1}}
)(
    input                         clk,
    input                         rst,
    input                         alloc_valid,
    output reg                    alloc_ready,
    output reg [7:0]              alloc_tag,
    input                         completion_valid,
    input [7:0]                   completion_tag,
    input                         completion_error,
    input                         cancel_all,
    output reg                    outcome_valid,
    input tri1                    outcome_ready,
    output reg [7:0]              outcome_tag,
    output reg [1:0]              outcome_status,
    output reg [TAG_COUNT-1:0]    active_tags,
    output reg [$clog2(TAG_COUNT+1)-1:0] outstanding_count
);
    localparam [1:0] OUTCOME_COMPLETED = 2'd0;
    localparam [1:0] OUTCOME_ERROR     = 2'd1;
    localparam [1:0] OUTCOME_TIMEOUT   = 2'd2;
    localparam [1:0] OUTCOME_CANCELLED = 2'd3;
    localparam integer INDEX_WIDTH = (TAG_COUNT <= 1) ? 1 : $clog2(TAG_COUNT);

    reg [TIMEOUT_WIDTH-1:0] age [0:TAG_COUNT-1];
    reg [1:0] terminal_status [0:TAG_COUNT-1];
    reg [TAG_COUNT-1:0] cancelled;
    reg [TAG_COUNT-1:0] cancel_reported;
    reg [TAG_COUNT-1:0] cancel_report_pending;
    reg [TAG_COUNT-1:0] terminal_pending;
    reg [INDEX_WIDTH-1:0] alloc_cursor;
    integer i;
    integer scan_index;
    integer selected;
    integer completed_index;
    integer outcome_index;
    reg timeout_found;
    reg [INDEX_WIDTH-1:0] timeout_index;
    reg cancel_found;
    reg [INDEX_WIDTH-1:0] cancel_index;
    reg terminal_found;
    reg [INDEX_WIDTH-1:0] terminal_index;
    wire outcome_ready_i = outcome_ready;

    always @(*) begin
        alloc_ready = 1'b0;
        alloc_tag = TAG_FIRST;
        selected = -1;
        scan_index = 0;
        completed_index = {24'd0, completion_tag} - {24'd0, TAG_FIRST};
        outcome_index = {24'd0, outcome_tag} - {24'd0, TAG_FIRST};
        outstanding_count = 0;
        timeout_found = 1'b0;
        timeout_index = 0;
        cancel_found = 1'b0;
        cancel_index = 0;
        terminal_found = 1'b0;
        terminal_index = 0;
        for (i = 0; i < TAG_COUNT; i = i + 1) begin
            if (active_tags[i]) begin
                outstanding_count = outstanding_count + 1'b1;
                if (!timeout_found && !terminal_pending[i] &&
                    (TIMEOUT_CYCLES != 0) &&
                    (age[i] >= TIMEOUT_CYCLES - 1'b1)) begin
                    timeout_found = 1'b1;
                    timeout_index = INDEX_WIDTH'(i);
                end
            end
            if (!cancel_found && cancel_report_pending[i]) begin
                cancel_found = 1'b1;
                cancel_index = INDEX_WIDTH'(i);
            end
            if (!terminal_found && terminal_pending[i] && !cancelled[i]) begin
                terminal_found = 1'b1;
                terminal_index = INDEX_WIDTH'(i);
            end
        end
        if (!(|cancel_report_pending) && !cancel_all) begin
            for (i = 0; i < TAG_COUNT; i = i + 1) begin
                scan_index = (int'(alloc_cursor) + i) % TAG_COUNT;
                if ((selected < 0) && !active_tags[scan_index]) begin
                    selected = scan_index;
                    alloc_ready = 1'b1;
                    alloc_tag = TAG_FIRST + 8'(scan_index);
                end
            end
        end
    end

    always @(posedge clk) begin
        if (rst) begin
            active_tags <= {TAG_COUNT{1'b0}};
            cancelled <= {TAG_COUNT{1'b0}};
            cancel_reported <= {TAG_COUNT{1'b0}};
            cancel_report_pending <= {TAG_COUNT{1'b0}};
            terminal_pending <= {TAG_COUNT{1'b0}};
            alloc_cursor <= {INDEX_WIDTH{1'b0}};
            outcome_valid <= 1'b0;
            outcome_tag <= TAG_FIRST;
            outcome_status <= OUTCOME_CANCELLED;
            for (i = 0; i < TAG_COUNT; i = i + 1) begin
                age[i] <= {TIMEOUT_WIDTH{1'b0}};
                terminal_status[i] <= OUTCOME_COMPLETED;
            end
        end else begin
            for (i = 0; i < TAG_COUNT; i = i + 1) begin
                if (active_tags[i] && !terminal_pending[i])
                    age[i] <= age[i] + 1'b1;
                if (cancel_all && active_tags[i] && !cancelled[i] &&
                    !terminal_pending[i]) begin
                    cancelled[i] <= 1'b1;
                    cancel_report_pending[i] <= 1'b1;
                end
                if (active_tags[i] && cancelled[i] && cancel_reported[i] &&
                    terminal_pending[i]) begin
                    active_tags[i] <= 1'b0;
                    cancelled[i] <= 1'b0;
                    cancel_reported[i] <= 1'b0;
                    terminal_pending[i] <= 1'b0;
                    age[i] <= {TIMEOUT_WIDTH{1'b0}};
                end
            end

            if (outcome_valid && outcome_ready_i) begin
                outcome_valid <= 1'b0;
                if (outcome_index >= 0 && outcome_index < TAG_COUNT) begin
                    if (outcome_status == OUTCOME_CANCELLED) begin
                        cancel_reported[outcome_index] <= 1'b1;
                    end else begin
                        active_tags[outcome_index] <= 1'b0;
                        terminal_pending[outcome_index] <= 1'b0;
                        age[outcome_index] <= {TIMEOUT_WIDTH{1'b0}};
                    end
                end
            end else if (!outcome_valid && cancel_found) begin
                cancel_report_pending[cancel_index] <= 1'b0;
                outcome_valid <= 1'b1;
                outcome_tag <= TAG_FIRST + 8'(cancel_index);
                outcome_status <= OUTCOME_CANCELLED;
            end else if (!outcome_valid && terminal_found) begin
                outcome_valid <= 1'b1;
                outcome_tag <= TAG_FIRST + 8'(terminal_index);
                outcome_status <= terminal_status[terminal_index];
            end

            if (completion_valid && (completion_tag >= TAG_FIRST) &&
                (completed_index < TAG_COUNT) && active_tags[completed_index] &&
                !terminal_pending[completed_index]) begin
                if (cancelled[completed_index] && cancel_reported[completed_index]) begin
                    active_tags[completed_index] <= 1'b0;
                    cancelled[completed_index] <= 1'b0;
                    cancel_reported[completed_index] <= 1'b0;
                    age[completed_index] <= {TIMEOUT_WIDTH{1'b0}};
                end else begin
                    terminal_pending[completed_index] <= 1'b1;
                    terminal_status[completed_index] <=
                        completion_error ? OUTCOME_ERROR : OUTCOME_COMPLETED;
                end
            end else if (timeout_found) begin
                if (cancelled[timeout_index] && cancel_reported[timeout_index]) begin
                    active_tags[timeout_index] <= 1'b0;
                    cancelled[timeout_index] <= 1'b0;
                    cancel_reported[timeout_index] <= 1'b0;
                    age[timeout_index] <= {TIMEOUT_WIDTH{1'b0}};
                end else begin
                    terminal_pending[timeout_index] <= 1'b1;
                    terminal_status[timeout_index] <= OUTCOME_TIMEOUT;
                end
            end

            if (alloc_valid && alloc_ready) begin
                active_tags[selected] <= 1'b1;
                cancelled[selected] <= 1'b0;
                cancel_reported[selected] <= 1'b0;
                cancel_report_pending[selected] <= 1'b0;
                terminal_pending[selected] <= 1'b0;
                terminal_status[selected] <= OUTCOME_COMPLETED;
                age[selected] <= {TIMEOUT_WIDTH{1'b0}};
                if (selected >= TAG_COUNT - 1)
                    alloc_cursor <= {INDEX_WIDTH{1'b0}};
                else
                    alloc_cursor <= INDEX_WIDTH'(selected + 1);
            end
        end
    end
endmodule
