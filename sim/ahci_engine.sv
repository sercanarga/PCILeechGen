`default_nettype none
// AHCI port command engine (slot 0). On PxCI bit0 with PxCMD.ST+FRE set it
// walks the command list -> command table (H2D FIS) -> PRDT data base, then:
//   IDENTIFY (0xEC)      -> DMA identify block to DBA
//   READ DMA EXT (0x25)  -> DMA one sector from backing store to DBA
//   WRITE DMA EXT (0x35) -> DMA one sector from DBA into backing store
// then writes a D2H Register FIS to PxFB+0x40, sets PxTFD ready, PxIS.DHRS,
// raises intr, and clears PxCI -- the handshake storahci waits on at init.
// (identify payload here is a placeholder pattern; the real donor-derived
// block is loaded from the codegen-emitted ahci_identify_init.hex.)
module ahci_engine #(
    parameter SECT_DW = 128,  // 512 bytes / sector
    parameter NSECT   = 4
)(
    input  wire        clk,
    input  wire        rst_n,
    input  wire        reg_wr,
    input  wire [11:0] reg_addr,
    input  wire [31:0] reg_wdata,
    input  wire [11:0] reg_raddr,
    output reg  [31:0] reg_rdata,
    output reg         mem_rd,
    output reg  [31:0] mem_addr,
    input  wire [31:0] mem_rdata,
    input  wire        mem_rvalid,
    output reg         mem_wr,
    output reg  [31:0] mem_waddr,
    output reg  [31:0] mem_wdata,
    output reg         intr
);
    localparam [11:0] A_PxCLB=12'h100, A_PxFB=12'h108, A_PxIS=12'h110,
                      A_PxIE=12'h114, A_PxCMD=12'h118, A_PxTFD=12'h120,
                      A_PxSIG=12'h124, A_PxSSTS=12'h128, A_PxCI=12'h138;

    reg [31:0] pxclb, pxfb, pxis, pxie, pxcmd, pxtfd, pxci;
    wire [31:0] pxsig  = 32'h00000101;
    wire [31:0] pxssts = 32'h00000113;

    reg [31:0] store [0:NSECT*SECT_DW-1];

    localparam S_IDLE=0,S_HDR=1,S_HDRW=2,S_F0=3,S_F0W=4,S_F1=5,S_F1W=6,
               S_PRDT=7,S_PRDTW=8,S_WDATA=9,S_RD=10,S_RDW=11,S_FIS=12,S_DONE=13;
    reg [3:0]  st;
    reg [31:0] ctba, dba, cmd, lba;
    reg [8:0]  cnt;
    wire [13:0] sidx = lba[1:0]*SECT_DW + cnt;

    always @(*) begin
        case (reg_raddr)
            A_PxCLB: reg_rdata=pxclb; A_PxFB: reg_rdata=pxfb;
            A_PxIS:  reg_rdata=pxis;  A_PxIE: reg_rdata=pxie;
            A_PxCMD: reg_rdata=pxcmd; A_PxTFD:reg_rdata=pxtfd;
            A_PxSIG: reg_rdata=pxsig; A_PxSSTS:reg_rdata=pxssts;
            A_PxCI:  reg_rdata=pxci;  default:reg_rdata=32'h0;
        endcase
    end

    always @(posedge clk) begin
        if (!rst_n) begin
            pxclb<=0;pxfb<=0;pxis<=0;pxie<=0;pxcmd<=0;pxtfd<=32'h50;pxci<=0;
            st<=S_IDLE;mem_rd<=0;mem_wr<=0;intr<=0;cnt<=0;
        end else begin
            mem_rd<=0; mem_wr<=0; intr<=0;
            if (reg_wr) case (reg_addr)
                A_PxCLB: pxclb<=reg_wdata;  A_PxFB: pxfb<=reg_wdata;
                A_PxIE:  pxie<=reg_wdata;   A_PxCMD:pxcmd<=reg_wdata;
                A_PxIS:  pxis<=pxis & ~reg_wdata;
                A_PxCI:  pxci<=pxci | reg_wdata;
                default: ;
            endcase

            case (st)
                S_IDLE: if (pxci[0]&&pxcmd[0]&&pxcmd[4]) begin
                            pxtfd<=32'h80; mem_rd<=1; mem_addr<=pxclb+32'h8; st<=S_HDR;
                        end
                S_HDR:  st<=S_HDRW;
                S_HDRW: if (mem_rvalid) begin ctba<=mem_rdata; mem_rd<=1; mem_addr<=mem_rdata; st<=S_F0; end
                S_F0:   st<=S_F0W;
                S_F0W:  if (mem_rvalid) begin cmd<=mem_rdata; mem_rd<=1; mem_addr<=ctba+32'h4; st<=S_F1; end
                S_F1:   st<=S_F1W;
                S_F1W:  if (mem_rvalid) begin lba<=mem_rdata; mem_rd<=1; mem_addr<=ctba+32'h80; st<=S_PRDT; end
                S_PRDT: st<=S_PRDTW;
                S_PRDTW:if (mem_rvalid) begin
                            dba<=mem_rdata; cnt<=0;
                            st<=(cmd[23:16]==8'h35)?S_RD:S_WDATA; // 0x35 write reads host first
                        end
                // IDENTIFY / READ: produce data to DBA
                S_WDATA: if (cnt<SECT_DW) begin
                            mem_wr<=1; mem_waddr<=dba+{cnt,2'b00};
                            mem_wdata<=(cmd[23:16]==8'hEC)?(32'hA5000000|cnt):store[sidx];
                            cnt<=cnt+1;
                        end else st<=S_FIS;
                // WRITE: pull data from DBA into backing store
                S_RD:   if (cnt<SECT_DW) begin mem_rd<=1; mem_addr<=dba+{cnt,2'b00}; st<=S_RDW; end
                        else st<=S_FIS;
                S_RDW:  if (mem_rvalid) begin store[sidx]<=mem_rdata; cnt<=cnt+1; st<=S_RD; end
                S_FIS:  begin mem_wr<=1; mem_waddr<=pxfb+32'h40; mem_wdata<=32'h00500034; st<=S_DONE; end
                S_DONE: begin
                            pxci<=pxci & ~32'h1; pxtfd<=32'h50; pxis<=pxis|32'h1;
                            if (pxie[0]) intr<=1; st<=S_IDLE;
                        end
                default: st<=S_IDLE;
            endcase
        end
    end
endmodule
`default_nettype wire
