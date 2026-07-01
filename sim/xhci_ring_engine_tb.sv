`default_nettype none
`timescale 1ns/1ps
module xhci_ring_engine_tb;
    reg clk=0, rst_n=0;
    reg        reg_wr=0;
    reg [11:0] reg_addr=0, reg_raddr=0;
    reg [31:0] reg_wdata=0;
    wire [31:0] reg_rdata;
    wire        mem_rd, mem_wr;
    wire [31:0] mem_addr, mem_waddr, mem_wdata;
    reg  [31:0] mem_rdata; reg mem_rvalid;
    wire        intr;
    integer errors=0, i;

    xhci_ring_engine dut(.clk(clk),.rst_n(rst_n),
        .reg_wr(reg_wr),.reg_addr(reg_addr),.reg_wdata(reg_wdata),
        .reg_raddr(reg_raddr),.reg_rdata(reg_rdata),
        .mem_rd(mem_rd),.mem_addr(mem_addr),.mem_rdata(mem_rdata),.mem_rvalid(mem_rvalid),
        .mem_wr(mem_wr),.mem_waddr(mem_waddr),.mem_wdata(mem_wdata),.intr(intr));

    always #5 clk=~clk;

    reg [31:0] hm [0:16383];
    always @(posedge clk) begin
        mem_rvalid <= mem_rd;
        if (mem_rd) mem_rdata <= hm[mem_addr[15:2]];
        if (mem_wr) hm[mem_waddr[15:2]] <= mem_wdata;
    end

    reg intr_seen=0;
    always @(posedge clk) if (intr) intr_seen<=1;

    task wr(input [11:0] a, input [31:0] d);
        begin @(posedge clk); reg_wr<=1; reg_addr<=a; reg_wdata<=d;
              @(posedge clk); reg_wr<=0; end
    endtask
    task chk(input [255:0] n, input [31:0] g, input [31:0] e);
        begin if (g!==e) begin $display("FAIL %0s: got %08x exp %08x",n,g,e); errors=errors+1; end
              else $display("ok   %0s = %08x",n,g); end
    endtask
    // ring the command doorbell, wait for the completion interrupt, then
    // let the FSM's trailing "check for more queued TRBs" drain pass settle
    // back to idle before handing control back (it keeps walking the ring
    // for a few cycles after posting the event).
    task ring_doorbell;
        begin
            intr_seen = 0;
            wr(12'h100, 32'h0);
            i=0; while (!intr_seen && i<500) begin @(posedge clk); i=i+1; end
            repeat (20) @(posedge clk);
        end
    endtask

    initial begin
        for (i=0;i<16384;i=i+1) hm[i]=0;
        @(posedge clk); rst_n<=1;

        // Command Ring base 0x1000, RCS=1
        wr(12'h038, 32'h00001001);
        wr(12'h03C, 32'h00000000);

        // Event Ring Segment Table at 0x2000: {base=0x3000, hi=0, size=16}
        hm['h2000>>2] = 32'h00003000;
        hm[('h2000>>2)+1] = 32'h00000000;
        hm[('h2000>>2)+2] = 32'h00000010; // ERSTSZ = 16 TRBs
        wr(12'h230, 32'h00002000);
        wr(12'h234, 32'h00000000);

        // enable interrupts
        wr(12'h220, 32'h00000002); // IMAN.IE=1

        // ---- 1) No-Op Command at 0x1000 (type 23, cycle=1) ----
        hm['h1000>>2] = 32'h0;
        hm[('h1000>>2)+1] = 32'h0;
        hm[('h1000>>2)+2] = 32'h0;
        hm[('h1000>>2)+3] = 32'h00005C01; // type=23<<10=0x5C00, cycle=1

        ring_doorbell();
        chk("noop: intr fired", {31'h0, intr_seen}, 32'h1);
        chk("noop: evt[0] dw0 (trb ptr)", hm['h3000>>2],      32'h00001000);
        chk("noop: evt[0] dw1",          hm[('h3000>>2)+1],   32'h00000000);
        chk("noop: evt[0] dw2 (compl)",  hm[('h3000>>2)+2],   32'h01000000);
        chk("noop: evt[0] dw3 (type/slot/cycle)", hm[('h3000>>2)+3], 32'h00008401);

        reg_raddr<=12'h220; @(posedge clk); #1;
        chk("noop: IMAN.IP set", reg_rdata, 32'h00000003);

        // driver acks the interrupt (W1C bit0), IE stays set
        wr(12'h220, 32'h00000003);
        reg_raddr<=12'h220; @(posedge clk); #1;
        chk("noop: IMAN.IP cleared", reg_rdata, 32'h00000002);

        // ---- 2) Enable Slot Command at 0x1010 (type 9, cycle=1) ----
        hm['h1010>>2] = 32'h0;
        hm[('h1010>>2)+1] = 32'h0;
        hm[('h1010>>2)+2] = 32'h0;
        hm[('h1010>>2)+3] = 32'h00002401; // type=9<<10=0x2400, cycle=1

        ring_doorbell();
        chk("enslot: evt[1] dw0 (trb ptr)", hm[('h3010>>2)],   32'h00001010);
        chk("enslot: evt[1] dw3 (slot=1)",  hm[('h3010>>2)+3], 32'h01008401);

        // ---- 3) Link TRB (TC=1) back to 0x1000, then a new No-Op with
        //         cycle=0 (ccs will have toggled) -----------------------
        hm['h1020>>2] = 32'h00001000; // link target
        hm[('h1020>>2)+1] = 32'h0;
        hm[('h1020>>2)+2] = 32'h0;
        hm[('h1020>>2)+3] = 32'h00001803; // type=6<<10=0x1800, TC=1, cycle=1

        hm['h1000>>2] = 32'h0;
        hm[('h1000>>2)+1] = 32'h0;
        hm[('h1000>>2)+2] = 32'h0;
        hm[('h1000>>2)+3] = 32'h00005C00; // type=23<<10=0x5C00, cycle=0 (post-toggle)

        ring_doorbell();
        chk("link+noop: evt[2] dw0 (trb ptr)", hm[('h3020>>2)],   32'h00001000);
        chk("link+noop: evt[2] dw3",           hm[('h3020>>2)+3], 32'h00008401);

        // ---- register readback sanity ----
        reg_raddr<=12'h038; @(posedge clk); #1;
        chk("CRCR_LO readback", reg_rdata, 32'h00001001);
        reg_raddr<=12'h230; @(posedge clk); #1;
        chk("ERSTBA_LO readback", reg_rdata, 32'h00002000);

        if (errors==0) $display("ALL TESTS PASSED");
        else $display("%0d FAILURES", errors);
        $finish;
    end
endmodule
`default_nettype wire
