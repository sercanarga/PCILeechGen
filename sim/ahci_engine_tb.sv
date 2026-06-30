`default_nettype none
`timescale 1ns/1ps
module ahci_engine_tb;
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

    ahci_engine dut(.clk(clk),.rst_n(rst_n),
        .reg_wr(reg_wr),.reg_addr(reg_addr),.reg_wdata(reg_wdata),
        .reg_raddr(reg_raddr),.reg_rdata(reg_rdata),
        .mem_rd(mem_rd),.mem_addr(mem_addr),.mem_rdata(mem_rdata),.mem_rvalid(mem_rvalid),
        .mem_wr(mem_wr),.mem_waddr(mem_waddr),.mem_wdata(mem_wdata),.intr(intr));

    always #5 clk=~clk;

    reg [31:0] hm [0:1023];
    always @(posedge clk) begin
        mem_rvalid <= mem_rd;
        if (mem_rd) mem_rdata <= hm[mem_addr[11:2]];
        if (mem_wr) hm[mem_waddr[11:2]] <= mem_wdata;
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
    // issue one command via slot 0, wait for PxCI to clear
    task issue(input [7:0] cmd, input [31:0] lba, input [31:0] dba);
        begin
            hm['h08>>2]  = 32'h00000200;                 // cmd header dw2 = CTBA
            hm['h200>>2] = {8'h00, cmd, 16'h0027};       // H2D FIS dword0
            hm['h204>>2] = lba;                          // FIS dword1 (LBA)
            hm['h280>>2] = dba;                          // PRDT[0] DBA
            intr_seen=0;
            wr(12'h138, 32'h1);                          // PxCI bit0
            reg_raddr<=12'h138; @(posedge clk);
            i=0; while (reg_rdata!==0 && i<500) begin @(posedge clk); i=i+1; end
        end
    endtask

    initial begin
        for (i=0;i<1024;i=i+1) hm[i]=0;
        @(posedge clk); rst_n<=1;
        wr(12'h100,32'h0);      // PxCLB
        wr(12'h108,32'h600);    // PxFB
        wr(12'h114,32'h1);      // PxIE
        wr(12'h118,32'h11);     // PxCMD ST+FRE

        // 1) IDENTIFY -> data at DBA 0x400
        issue(8'hEC, 0, 32'h400);
        reg_raddr<=12'h138; @(posedge clk); chk("identify: PxCI clear", reg_rdata, 0);
        chk("identify data[0]",   hm['h400>>2],        32'hA5000000);
        chk("identify data[127]", hm[('h400>>2)+127],  32'hA500007F);
        chk("identify intr",      {31'h0,intr_seen},   32'h1);

        // 2) WRITE sector 1 from host buffer 0x800
        for (i=0;i<128;i=i+1) hm[('h800>>2)+i] = 32'hBEEF0000 | i;
        issue(8'h35, 1, 32'h800);
        reg_raddr<=12'h138; @(posedge clk); chk("write: PxCI clear", reg_rdata, 0);

        // 3) READ sector 1 back to 0xA00 -> must equal what we wrote
        for (i=0;i<128;i=i+1) hm[('hA00>>2)+i] = 32'h0;
        issue(8'h25, 1, 32'hA00);
        reg_raddr<=12'h138; @(posedge clk); chk("read: PxCI clear", reg_rdata, 0);
        chk("read-back[0]",   hm['hA00>>2],        32'hBEEF0000);
        chk("read-back[63]",  hm[('hA00>>2)+63],   32'hBEEF003F);
        chk("read-back[127]", hm[('hA00>>2)+127],  32'hBEEF007F);
        chk("D2H FIS",        hm['h640>>2],        32'h00500034);

        if (errors==0) $display("ALL TESTS PASSED");
        else $display("%0d FAILURES", errors);
        $finish;
    end
endmodule
`default_nettype wire
