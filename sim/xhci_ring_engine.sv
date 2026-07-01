`default_nettype none
// xHCI Command Ring + Event Ring engine.
//
// Emulates the HOST CONTROLLER side of the xHCI init handshake with
// xhci.sys -- there is no downstream USB device attached (PORTSC CCS=0 on
// both ports, see xhci.go's xhciProfile), so this does not pretend to
// enumerate a real device. It only has to:
//   - latch the Command Ring pointer + cycle state from CRCR on the first
//     doorbell ring (CRCR is only consumed once, while the ring is stopped)
//   - walk the ring on each Command doorbell (DB0) write, following Link
//     TRBs (type 6) and consuming every TRB whose cycle bit matches the
//     ring's current consumer cycle state, until it catches up (cycle
//     mismatch = ring drained)
//   - answer No-Op Command (type 23) and Enable Slot Command (type 9) with
//     a Command Completion Event (type 33), and fall back to a generic
//     Success completion (echoing the TRB's own Slot ID field) for anything
//     else so a command never just hangs the driver
//   - post events through the one-segment Event Ring described by the ERST
//     entry at ERSTBA (fetched lazily, invalidated whenever ERSTBA is
//     rewritten), tracking its own enqueue index + producer cycle state
//   - set IMAN.IP and pulse an interrupt (gated on IMAN.IE) per completion
//
// ponytail: single fixed Slot ID (FIXED_SLOT_ID) for Enable Slot Command --
// with zero devices attached there is at most one plausible slot in flight,
// not a real free-list across HCSPARAMS1.MaxSlots. ERDP.EHB and CRCR.CRR
// are not modeled: neither is polled during a no-device init handshake,
// only real-hardware backpressure/abort bookkeeping the driver doesn't need
// to see here.
module xhci_ring_engine #(
    parameter [7:0] FIXED_SLOT_ID = 8'h01
)(
    input  wire        clk,
    input  wire        rst_n,

    input  wire        reg_wr,
    input  wire [11:0]  reg_addr,
    input  wire [31:0]  reg_wdata,
    input  wire [11:0]  reg_raddr,
    output reg  [31:0]  reg_rdata,

    output reg          mem_rd,
    output reg  [31:0]  mem_addr,
    input  wire [31:0]  mem_rdata,
    input  wire         mem_rvalid,
    output reg          mem_wr,
    output reg  [31:0]  mem_waddr,
    output reg  [31:0]  mem_wdata,

    output reg          intr
);
    localparam [11:0] A_DB0         = 12'h100; // Command doorbell (DBOFF fixed 0x100)
    localparam [11:0] A_CRCR_LO     = 12'h038;
    localparam [11:0] A_CRCR_HI     = 12'h03C;
    localparam [11:0] A_ERSTBA_LO   = 12'h230; // RTSOFF(0x200) + IR0 ERSTBA (0x30)
    localparam [11:0] A_ERSTBA_HI   = 12'h234;
    localparam [11:0] A_IMAN        = 12'h220; // RTSOFF(0x200) + IR0 IMAN (0x00... +0x20 IR0 base)

    localparam [5:0] TRB_LINK       = 6'd6;
    localparam [5:0] TRB_ENSLOT     = 6'd9;
    localparam [5:0] TRB_NOOP       = 6'd23;
    localparam [5:0] TRB_CMD_EVENT  = 6'd33;

    localparam [3:0] S_IDLE   = 4'h0;
    localparam [3:0] S_FETCH  = 4'h1;
    localparam [3:0] S_FETCHW = 4'h2;
    localparam [3:0] S_DECODE = 4'h3;
    localparam [3:0] S_ERST   = 4'h4;
    localparam [3:0] S_ERSTW  = 4'h5;
    localparam [3:0] S_EVT    = 4'h6;
    localparam [3:0] S_POST   = 4'h7;

    // driver-programmed registers
    reg [31:0] crcr_lo, crcr_hi;
    reg [31:0] erstba_lo, erstba_hi;
    reg [1:0]  iman; // bit0=IP (RW1C, HW-set), bit1=IE (RW)

    // command ring consumer state
    reg        cr_started;
    reg [31:0] crdp;    // command ring dequeue pointer (32-bit: sim host address space)
    reg        ccs;     // consumer cycle state

    // event ring producer state
    reg        erst_valid;
    reg [31:0] erst_base;
    reg [15:0] erst_size;
    reg [15:0] evt_idx;
    reg        ecs;     // producer cycle state (starts at 1 per spec)

    reg [3:0]  state;
    reg [1:0]  cnt;
    reg [31:0] trb0, trb1, trb2, trb3;
    reg [31:0] cmd_trb_addr;
    reg [7:0]  slot_id_out;

    wire [31:0] evt_addr = erst_base + {16'h0, evt_idx, 4'b0};

    always @(*) begin
        case (reg_raddr)
            A_CRCR_LO:   reg_rdata = crcr_lo;
            A_CRCR_HI:   reg_rdata = crcr_hi;
            A_ERSTBA_LO: reg_rdata = erstba_lo;
            A_ERSTBA_HI: reg_rdata = erstba_hi;
            A_IMAN:      reg_rdata = {30'h0, iman};
            default:     reg_rdata = 32'h0;
        endcase
    end

    always @(posedge clk) begin
        if (!rst_n) begin
            crcr_lo <= 32'h0; crcr_hi <= 32'h0;
            erstba_lo <= 32'h0; erstba_hi <= 32'h0;
            iman <= 2'b00;
            cr_started <= 1'b0; crdp <= 32'h0; ccs <= 1'b1;
            erst_valid <= 1'b0; erst_base <= 32'h0; erst_size <= 16'h0;
            evt_idx <= 16'h0; ecs <= 1'b1;
            state <= S_IDLE; cnt <= 2'h0;
            mem_rd <= 1'b0; mem_wr <= 1'b0; intr <= 1'b0;
        end else begin
            mem_rd <= 1'b0; mem_wr <= 1'b0; intr <= 1'b0;

            if (reg_wr) case (reg_addr)
                A_CRCR_LO: if (!cr_started) crcr_lo <= reg_wdata & 32'hFFFFFFC7;
                A_CRCR_HI: if (!cr_started) crcr_hi <= reg_wdata;
                A_ERSTBA_LO: begin erstba_lo <= reg_wdata & 32'hFFFFFFF0; erst_valid <= 1'b0; end
                A_ERSTBA_HI: begin erstba_hi <= reg_wdata; erst_valid <= 1'b0; end
                A_IMAN: begin
                    iman[1] <= reg_wdata[1];          // IE, plain RW
                    if (reg_wdata[0]) iman[0] <= 1'b0; // IP, write-1-to-clear
                end
                default: ;
            endcase

            case (state)
                // ---- wait for a Command doorbell ring ------------------
                S_IDLE: if (reg_wr && reg_addr == A_DB0) begin
                    if (!cr_started) begin
                        crdp       <= {crcr_lo[31:6], 6'b0};
                        ccs        <= crcr_lo[0];
                        cr_started <= 1'b1;
                    end
                    cnt   <= 2'h0;
                    state <= S_FETCH;
                end

                // ---- pull 16 bytes (one TRB) from the command ring -----
                S_FETCH: begin
                    mem_rd  <= 1'b1;
                    mem_addr <= crdp + {28'h0, cnt, 2'b00};
                    state   <= S_FETCHW;
                end
                S_FETCHW: if (mem_rvalid) begin
                    case (cnt)
                        2'd0: trb0 <= mem_rdata;
                        2'd1: trb1 <= mem_rdata;
                        2'd2: trb2 <= mem_rdata;
                        2'd3: trb3 <= mem_rdata;
                    endcase
                    if (cnt == 2'd3) begin cnt <= 2'h0; state <= S_DECODE; end
                    else begin cnt <= cnt + 2'd1; state <= S_FETCH; end
                end

                // ---- ring caught up, follow Link, or execute a command -
                S_DECODE: begin
                    if (trb3[0] != ccs) begin
                        // cycle bit doesn't match -- nothing new queued
                        state <= S_IDLE;
                    end else if (trb3[15:10] == TRB_LINK) begin
                        crdp  <= trb0 & 32'hFFFFFFF0;
                        if (trb3[1]) ccs <= ~ccs; // Toggle Cycle
                        cnt   <= 2'h0;
                        state <= S_FETCH;
                    end else begin
                        cmd_trb_addr <= crdp;
                        crdp         <= crdp + 32'd16;
                        if (trb3[15:10] == TRB_NOOP)
                            slot_id_out <= 8'h00;
                        else if (trb3[15:10] == TRB_ENSLOT)
                            slot_id_out <= FIXED_SLOT_ID;
                        else
                            slot_id_out <= trb3[31:24]; // everything else: echo, say yes
                        cnt   <= 2'h0;
                        state <= erst_valid ? S_EVT : S_ERST;
                    end
                end

                // ---- fetch the (single) Event Ring Segment Table entry -
                S_ERST: begin
                    mem_rd   <= 1'b1;
                    mem_addr <= erstba_lo + {28'h0, cnt, 2'b00};
                    state    <= S_ERSTW;
                end
                S_ERSTW: if (mem_rvalid) begin
                    case (cnt)
                        2'd0: erst_base <= mem_rdata;
                        2'd2: erst_size <= mem_rdata[15:0];
                        default: ;
                    endcase
                    if (cnt == 2'd3) begin
                        erst_valid <= 1'b1;
                        cnt        <= 2'h0;
                        state      <= S_EVT;
                    end else begin
                        cnt   <= cnt + 2'd1;
                        state <= S_ERST;
                    end
                end

                // ---- write the 16-byte Command Completion Event --------
                S_EVT: begin
                    mem_wr   <= 1'b1;
                    mem_waddr <= evt_addr + {28'h0, cnt, 2'b00};
                    case (cnt)
                        2'd0: mem_wdata <= cmd_trb_addr;
                        2'd1: mem_wdata <= 32'h0;
                        2'd2: mem_wdata <= {8'h01, 24'h0}; // Completion Code = Success
                        2'd3: mem_wdata <= {slot_id_out, 8'h00, TRB_CMD_EVENT, 9'h0, ecs};
                    endcase
                    if (cnt == 2'd3) begin cnt <= 2'h0; state <= S_POST; end
                    else cnt <= cnt + 2'd1;
                end

                // ---- advance the event ring, raise the interrupt, loop -
                S_POST: begin
                    iman[0] <= 1'b1; // IP
                    intr    <= iman[1]; // gated on IE
                    if (evt_idx + 16'd1 >= erst_size) begin
                        evt_idx <= 16'h0;
                        ecs     <= ~ecs;
                    end else begin
                        evt_idx <= evt_idx + 16'd1;
                    end
                    cnt   <= 2'h0;
                    state <= S_FETCH; // drain loop: check for more queued TRBs
                end

                default: state <= S_IDLE;
            endcase
        end
    end
endmodule
`default_nettype wire
