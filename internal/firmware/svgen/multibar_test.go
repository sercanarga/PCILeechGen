package svgen

import (
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/firmware/barmodel"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

func TestBARSlotsPreserveSparseTopologyAnd64BitUpperSlot(t *testing.T) {
	primary := &barmodel.BARModel{
		BIR: 0, Size: 0x1800, Aperture: 0x1800, Type: pci.BARTypeMem64,
		Is64Bit: true, UpperBIR: 1, ClassSpecific: true,
	}
	secondary := &barmodel.BARModel{
		BIR: 3, Size: 0x1000, Aperture: 0x1000, Type: pci.BARTypeMem32,
	}
	cfg := &SVGeneratorConfig{BARModels: []*barmodel.BARModel{primary, secondary}, BARModel: primary}

	slots := cfg.BARSlots()
	if len(slots) != 6 {
		t.Fatalf("BARSlots length = %d, want six physical BAR selectors", len(slots))
	}
	if slots[0].Model != primary || !slots[0].Primary || slots[0].ModuleName != "pcileech_bar_impl_device" {
		t.Fatalf("BAR0 slot = %+v, want primary 64-bit endpoint", slots[0])
	}
	if slots[1].Model != nil {
		t.Fatalf("BAR1 slot = %+v, want no endpoint because 64-bit BAR0 consumes it", slots[1])
	}
	if slots[3].Model != secondary || slots[3].Primary || slots[3].ModuleName != "pcileech_bar_impl_device_bar3" {
		t.Fatalf("BAR3 slot = %+v, want independent secondary endpoint", slots[3])
	}
	for _, bir := range []int{2, 4, 5} {
		if slots[bir].Model != nil {
			t.Errorf("absent BAR%d unexpectedly has endpoint %+v", bir, slots[bir])
		}
	}
}

func TestGenerateBarImplDeviceSVEmitsOneModulePerPresentBAR(t *testing.T) {
	primary := &barmodel.BARModel{
		BIR: 2, Size: 0x1000, Aperture: 0x1000, Type: pci.BARTypeMem32,
		ClassSpecific: true,
		Registers:     []barmodel.BARRegister{{Offset: 0x00, Width: 4, Name: "PRIMARY", Reset: 0x11223344}},
	}
	secondary := &barmodel.BARModel{
		BIR: 4, Size: 0x2000, Aperture: 0x2000, Type: pci.BARTypeMem32,
		Registers: []barmodel.BARRegister{{Offset: 0x40, Width: 4, Name: "SECONDARY", Reset: 0x55667788}},
	}
	cfg := &SVGeneratorConfig{
		BARModels:   []*barmodel.BARModel{primary, secondary},
		BARModel:    primary,
		DeviceClass: "ethernet",
	}

	sv, err := GenerateBarImplDeviceSV(cfg)
	if err != nil {
		t.Fatalf("GenerateBarImplDeviceSV: %v", err)
	}
	if strings.Count(sv, "module pcileech_bar_impl_device #(") != 1 {
		t.Error("generated SV must contain exactly one compatibility/primary BAR module")
	}
	if strings.Count(sv, "module pcileech_bar_impl_device_bar4 #(") != 1 {
		t.Error("generated SV must contain an independent module for secondary BAR4")
	}
	if !strings.Contains(sv, "reg_0x00000000") || !strings.Contains(sv, "reg_0x00000040") {
		t.Error("generated endpoint modules do not contain both BAR register maps")
	}
}

func TestGenerateBarImplDeviceSVUsesWritableStorageForUnmodeledBAR(t *testing.T) {
	primary := &barmodel.BARModel{
		BIR: 2, Size: 0x1000, Aperture: 0x1000, Type: pci.BARTypeMem32,
		ClassSpecific: true,
		Registers:     []barmodel.BARRegister{{Offset: 0, Width: 4, Name: "PRIMARY"}},
	}
	unmodeled := &barmodel.BARModel{
		BIR: 4, Size: 0x1800, Aperture: 0x1800, Type: pci.BARTypeMem32,
	}
	cfg := &SVGeneratorConfig{
		BARModels: []*barmodel.BARModel{primary, unmodeled},
		BARModel:  primary,
	}

	sv, err := GenerateBarImplDeviceSV(cfg)
	if err != nil {
		t.Fatalf("GenerateBarImplDeviceSV: %v", err)
	}
	start := strings.Index(sv, "module pcileech_bar_impl_device_bar4 #(")
	if start < 0 {
		t.Fatal("unmodeled BAR4 module missing")
	}
	bar4 := sv[start:]
	for _, want := range []string{
		"parameter BAR_SIZE = 6144",
		"localparam BAR_STORAGE_BYTES = (BAR_SIZE > 4096) ? 4096 : BAR_SIZE",
		"localparam BAR_WORD_ADDR_BITS = (BAR_WORDS > 1) ? $clog2(BAR_WORDS) : 1",
		"wire [BAR_WORD_ADDR_BITS-1:0] wr_word = wr_addr[BAR_WORD_ADDR_BITS+1:2]",
		"wire [BAR_WORD_ADDR_BITS-1:0] rd_word = rd_req_addr[BAR_WORD_ADDR_BITS+1:2]",
		"reg [31:0] bar_mem",
		"if (wr_valid && (wr_addr < BAR_STORAGE_BYTES))",
		"if (wr_be[0]) bar_mem[wr_word]",
		"reg [31:0] rd_data_d1",
		"rd_req_valid_d1 <= rd_req_valid && (rd_req_addr < BAR_SIZE)",
	} {
		if !strings.Contains(bar4, want) {
			t.Errorf("unmodeled BAR4 lacks narrowed writable pipelined aperture behavior %q", want)
		}
	}
	for _, pattern := range []string{
		`rd_data_d1\s*<=\s*bar_mem\[rd_word\]`,
		`rd_rsp_data\s*<=\s*rd_data_d1`,
	} {
		if !regexp.MustCompile(pattern).MatchString(bar4) {
			t.Errorf("unmodeled BAR4 lacks pipelined read behavior matching %q", pattern)
		}
	}
	for _, forbidden := range []string{"bar_mem[rd_req_addr[31:2]]", "bar_mem[wr_addr[31:2]]"} {
		if strings.Contains(bar4, forbidden) {
			t.Errorf("fallback BRAM uses over-wide address index %q", forbidden)
		}
	}
	if regexp.MustCompile(`rd_rsp_data\s*<=\s*doutb`).MatchString(bar4) {
		t.Fatal("fallback read response uses the live request address instead of pipelined request data")
	}
}

func TestGenerateBarControllerSVHasNoSyntheticLoopbackOrMSIEndpoints(t *testing.T) {
	model := &barmodel.BARModel{
		BIR: 3, Size: 0x1800, Aperture: 0x1800, Type: pci.BARTypeMem32,
		Registers: []barmodel.BARRegister{{Offset: 0, Width: 4, Name: "REG"}},
	}
	cfg := &SVGeneratorConfig{BARModels: []*barmodel.BARModel{model}, BARModel: model}

	sv, err := GenerateBarControllerSV(cfg)
	if err != nil {
		t.Fatalf("GenerateBarControllerSV: %v", err)
	}
	for _, forbidden := range []string{"pcileech_bar_impl_loopaddr", "pcileech_bar_impl_msi"} {
		if strings.Contains(sv, forbidden) {
			t.Errorf("generated controller contains unconditional synthetic endpoint %q", forbidden)
		}
	}
	if !strings.Contains(sv, "pcileech_bar_impl_device") || !strings.Contains(sv, "i_bar3") {
		t.Error("present donor BAR3 endpoint was not instantiated at BIR3")
	}
	for _, bir := range []string{"0", "1", "2", "4", "5"} {
		if !strings.Contains(sv, "pcileech_bar_impl_none i_bar"+bir) {
			t.Errorf("absent BAR%s is not explicitly tied to a no-response endpoint", bir)
		}
	}
	if !strings.Contains(sv, "assign intr_req = 1'b0") {
		t.Error("device without an interrupt endpoint must tie intr_req low")
	}
}

func TestGenerateBarControllerSVRejectsOutOfApertureAddressesWithoutWrapping(t *testing.T) {
	model := &barmodel.BARModel{
		BIR: 3, Size: 0x1800, Aperture: 0x1800, Type: pci.BARTypeMem32,
		Registers: []barmodel.BARRegister{{Offset: 0, Width: 4, Name: "REG"}},
	}
	cfg := &SVGeneratorConfig{BARModels: []*barmodel.BARModel{model}, BARModel: model}

	sv, err := GenerateBarControllerSV(cfg)
	if err != nil {
		t.Fatalf("GenerateBarControllerSV: %v", err)
	}
	if strings.Contains(sv, "ADDR_MASK") || regexp.MustCompile(`(?m)wr_addr\s*&\s*32'h`).MatchString(sv) || regexp.MustCompile(`(?m)rd_req_addr\s*&\s*32'h`).MatchString(sv) {
		t.Fatal("BAR address masking aliases out-of-range requests back into the aperture")
	}
	legalCompare := regexp.MustCompile(`(?m)(wr_addr|rd_req_addr)\s*<\s*(32'h00001800|32'd6144|6144)`)
	if len(legalCompare.FindAllString(sv, -1)) < 2 {
		t.Fatalf("generated controller lacks explicit read/write addr < 0x1800 aperture checks")
	}
}

func TestGenerateBarControllerSVSuppressesOnlyExactMSIXWriteRanges(t *testing.T) {
	bar2 := &barmodel.BARModel{BIR: 2, Size: 0x2000, Aperture: 0x2000, Type: pci.BARTypeMem32}
	cfg := &SVGeneratorConfig{
		BARModels: []*barmodel.BARModel{bar2},
		BARModel:  bar2,
		HasMSIX:   true,
		MSIXConfig: &MSIXConfig{
			NumVectors: 4,
			TableBIR:   2, TableOffset: 0x400,
			PBABIR: 2, PBAOffset: 0x800,
		},
	}

	sv, err := GenerateBarControllerSV(cfg)
	if err != nil {
		t.Fatalf("GenerateBarControllerSV: %v", err)
	}
	flat := strings.Join(strings.Fields(sv), " ")
	for _, want := range []string{
		"wire msix_wr_table_select = wr_valid && wr_bar[2] && (wr_addr >= 1024) && (wr_addr < 1088);",
		"wire msix_wr_pba_select = wr_valid && wr_bar[2] && (wr_addr >= 2048) && (wr_addr < 2056);",
		"wr_valid && wr_bar[2] && bar0_wr_hit && !msix_wr_table_select && !msix_wr_pba_select",
	} {
		if !strings.Contains(flat, want) {
			t.Errorf("generated controller lacks exact MSI-X write decode %q", want)
		}
	}
	baseWrite := func(addr uint32) bool {
		table := addr >= 0x400 && addr < 0x440
		pba := addr >= 0x800 && addr < 0x808
		return addr < 0x2000 && !table && !pba
	}
	for _, addr := range []uint32{0, 0x3FC, 0x440, 0x7FC, 0x808, 0x1FFC} {
		if !baseWrite(addr) {
			t.Errorf("ordinary BAR2 write at 0x%X was excluded", addr)
		}
	}
	for _, addr := range []uint32{0x400, 0x43C, 0x800, 0x804} {
		if baseWrite(addr) {
			t.Errorf("MSI-X write at 0x%X leaked into base endpoint", addr)
		}
	}
}

func TestGenerateBarControllerSVRoutesMSIXRegionsByIndependentBIR(t *testing.T) {
	bar2 := &barmodel.BARModel{BIR: 2, Size: 0x2000, Aperture: 0x2000, Type: pci.BARTypeMem32}
	bar4 := &barmodel.BARModel{BIR: 4, Size: 0x1000, Aperture: 0x1000, Type: pci.BARTypeMem32}
	cfg := &SVGeneratorConfig{
		BARModels: []*barmodel.BARModel{bar2, bar4},
		BARModel:  bar2,
		HasMSIX:   true,
		MSIXConfig: &MSIXConfig{
			NumVectors: 4,
			TableBIR:   2, TableOffset: 0x400,
			PBABIR: 4, PBAOffset: 0x800,
		},
	}

	sv, err := GenerateBarControllerSV(cfg)
	if err != nil {
		t.Fatalf("GenerateBarControllerSV: %v", err)
	}
	for _, want := range []string{
		"wire msix_wr_table_select = wr_valid && wr_bar[2]",
		"wire msix_wr_pba_select   = wr_valid && wr_bar[4]",
		"wire msix_rd_table_select = rd_req_valid && rd_req_bar[2]",
		"wire msix_rd_pba_select   = rd_req_valid && rd_req_bar[4]",
	} {
		if !strings.Contains(sv, want) {
			t.Errorf("generated controller lacks BIR-qualified MSI-X route %q", want)
		}
	}
	for _, want := range []string{
		"wr_valid && wr_bar[2] && bar0_wr_hit && !msix_wr_table_select && !msix_wr_pba_select",
		"wr_valid && wr_bar[4] && (wr_addr < 4096) && !msix_wr_table_select && !msix_wr_pba_select",
	} {
		if !strings.Contains(sv, want) {
			t.Errorf("generated controller endpoint write path lacks MSI-X exclusion %q", want)
		}
	}
	block := systemVerilogInstance(sv, "pcileech_msix_table #(")
	if block == "" {
		t.Fatal("MSI-X table instance missing")
	}
	for _, want := range []string{
		".TABLE_BIR", "( 2", "32'h00000400",
		".PBA_BIR", "( 4", "32'h00000800",
		".wr_table_select", ".wr_pba_select", ".rd_table_select", ".rd_pba_select",
	} {
		if !strings.Contains(block, want) {
			t.Errorf("MSI-X instance does not preserve/connect routing token %q", want)
		}
	}
}

func TestBarResponseArbiterRandomizedCollisionScoreboard(t *testing.T) {
	arbiter, err := GenerateBarRspArbiterSV(&SVGeneratorConfig{})
	if err != nil {
		t.Fatalf("GenerateBarRspArbiterSV: %v", err)
	}
	verilator, err := exec.LookPath("verilator")
	if err != nil {
		t.Skip("verilator not installed")
	}
	dir := t.TempDir()
	arbiterPath := filepath.Join(dir, "pcileech_bar_rsp_arbiter.sv")
	tbPath := filepath.Join(dir, "tb.sv")
	if err := os.WriteFile(arbiterPath, []byte(arbiter), 0600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(tbPath, []byte(barResponseArbiterTestbench), 0600); err != nil {
		t.Fatal(err)
	}
	objDir := filepath.Join(dir, "obj")
	cmd := exec.Command(verilator, "--binary", "--timing", "-Wno-fatal", "--top-module", "tb", "--Mdir", objDir, "-o", "sim", arbiterPath, tbPath)
	if output, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("verilator build: %v\n%s", err, output)
	}
	sim := filepath.Join(objDir, "sim")
	if output, err := exec.Command(sim).CombinedOutput(); err != nil {
		t.Fatalf("arbiter simulation: %v\n%s", err, output)
	}
}

const barResponseArbiterTestbench = "`timescale 1ns/1ps\n" +
	"module tb;\n" +
	"reg clk = 0;\n" +
	"reg rst = 1;\n" +
	"reg [6:0] in_valid = 0;\n" +
	"reg [87:0] in_ctx [0:6];\n" +
	"reg [31:0] in_data [0:6];\n" +
	"wire [6:0] in_ready;\n" +
	"wire out_valid;\n" +
	"wire [87:0] out_ctx;\n" +
	"wire [31:0] out_data;\n" +
	"reg expected_valid [0:8191];\n" +
	"reg seen [0:8191];\n" +
	"reg [31:0] expected_data [0:8191];\n" +
	"integer next_id = 0;\n" +
	"integer accepted = 0;\n" +
	"integer received = 0;\n" +
	"reg saw_backpressure = 0;\n" +
	"reg saw_collision = 0;\n" +
	"integer cycle;\n" +
	"integer source;\n" +
	"integer id;\n" +
	"reg [31:0] rng = 32'h51A7C3E9;\n" +
	"always #1 clk = ~clk;\n" +
	"pcileech_bar_rsp_arbiter dut(.clk(clk), .rst(rst), .in_valid(in_valid), .in_ctx(in_ctx), .in_data(in_data), .in_ready(in_ready), .out_valid(out_valid), .out_ctx(out_ctx), .out_data(out_data));\n" +
	"task sample;\n" +
	"integer s;\n" +
	"integer accepted_mask;\n" +
	"begin\n" +
	"accepted_mask = 0;\n" +
	"for (s = 0; s < 7; s = s + 1) begin\n" +
	"if (in_valid[s] && in_ready[s]) begin\n" +
	"id = in_ctx[s][31:0];\n" +
	"expected_valid[id] = 1;\n" +
	"expected_data[id] = in_data[s];\n" +
	"accepted = accepted + 1;\n" +
	"accepted_mask = accepted_mask | (1 << s);\n" +
	"end\n" +
	"if (in_valid[s] && !in_ready[s]) saw_backpressure = 1;\n" +
	"end\n" +
	"if ((accepted_mask & 1) != 0 && (accepted_mask & 8) != 0) saw_collision = 1;\n" +
	"if (out_valid) begin\n" +
	"id = out_ctx[31:0];\n" +
	"if (id < 0 || id >= 8192 || !expected_valid[id] || seen[id]) $fatal(1, \"unexpected or duplicate response id=%0d\", id);\n" +
	"if (out_data !== expected_data[id]) $fatal(1, \"data mismatch id=%0d got=%08x want=%08x\", id, out_data, expected_data[id]);\n" +
	"seen[id] = 1;\n" +
	"received = received + 1;\n" +
	"end\n" +
	"end\n" +
	"endtask\n" +
	"initial begin\n" +
	"for (source = 0; source < 7; source = source + 1) begin in_ctx[source] = 0; in_data[source] = 0; end\n" +
	"for (id = 0; id < 8192; id = id + 1) begin expected_valid[id] = 0; seen[id] = 0; expected_data[id] = 0; end\n" +
	"repeat (4) @(posedge clk);\n" +
	"rst = 0;\n" +
	"for (cycle = 0; cycle < 1200; cycle = cycle + 1) begin\n" +
	"@(negedge clk);\n" +
	"for (source = 0; source < 7; source = source + 1) begin\n" +
	"if (!(in_valid[source] && !in_ready[source])) begin\n" +
	"rng = {rng[30:0], rng[31] ^ rng[21] ^ rng[1] ^ rng[0]};\n" +
	"if (rng[2:0] != 0 && next_id < 8192) begin\n" +
	"in_valid[source] = 1;\n" +
	"in_ctx[source] = {56'd0, next_id[31:0]};\n" +
	"in_data[source] = 32'hA5000000 ^ next_id ^ (source << 20);\n" +
	"next_id = next_id + 1;\n" +
	"end else in_valid[source] = 0;\n" +
	"end\n" +
	"end\n" +
	"@(posedge clk); sample();\n" +
	"end\n" +
	"@(negedge clk); in_valid = 0;\n" +
	"for (cycle = 0; cycle < 10000 && received < accepted; cycle = cycle + 1) begin @(posedge clk); sample(); end\n" +
	"if (received != accepted) $fatal(1, \"lost responses accepted=%0d received=%0d\", accepted, received);\n" +
	"if (!saw_collision) $fatal(1, \"no primary-secondary collision exercised\");\n" +
	"if (!saw_backpressure) $fatal(1, \"no input backpressure exercised\");\n" +
	"$finish;\n" +
	"end\n" +
	"endmodule\n"

func systemVerilogInstance(sv, start string) string {
	begin := strings.Index(sv, start)
	if begin < 0 {
		return ""
	}
	end := strings.Index(sv[begin:], "\n    );")
	if end < 0 {
		return sv[begin:]
	}
	return sv[begin : begin+end+len("\n    );")]
}
