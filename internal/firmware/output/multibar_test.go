package output

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/firmware"
	"github.com/sercanarga/pcileechgen/internal/firmware/barmodel"
	"github.com/sercanarga/pcileechgen/internal/firmware/svgen"
	"github.com/sercanarga/pcileechgen/internal/firmware/tclgen"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

func TestBuildSVConfigPreservesAllDonorBARModelsAndPreferredBIR(t *testing.T) {
	ctx := makeDonorContext(0x10EC, 0x8125, 0x020000)
	ctx.BARs = []pci.BAR{
		{Index: 0, Size: 0x4000, Type: pci.BARTypeMem32},
		{Index: 2, Size: 0x1000, Type: pci.BARTypeMem64, Is64Bit: true},
		{Index: 4, Size: 0x2000, Type: pci.BARTypeMem32, Prefetchable: true},
	}
	ctx.BARContents = map[int][]byte{
		0: make([]byte, 0x4000),
		2: make([]byte, 0x1000),
		4: make([]byte, 0x2000),
	}
	ow := NewOutputWriter(t.TempDir(), "", 0, 0)
	ids := firmware.ExtractDeviceIDs(ctx.ConfigSpace, nil)

	cfg, err := ow.buildSVConfig(ctx, ctx.ConfigSpace.Clone(), ids, 0x1234, &board.Board{BRAMSize: 0x10000})
	if err != nil {
		t.Fatalf("buildSVConfig: %v", err)
	}
	if len(cfg.BARModels) != 3 {
		t.Fatalf("BARModels count = %d, want BAR0, BAR2, and BAR4", len(cfg.BARModels))
	}
	for i, wantBIR := range []int{0, 2, 4} {
		if cfg.BARModels[i].BIR != wantBIR {
			t.Errorf("BARModels[%d].BIR = %d, want %d", i, cfg.BARModels[i].BIR, wantBIR)
		}
	}
	if cfg.BARModel == nil || cfg.BARModel.BIR != 2 {
		t.Fatalf("legacy primary BARModel = %+v, want class-preferred Ethernet BAR2", cfg.BARModel)
	}
	if !cfg.BARModel.ClassSpecific {
		t.Fatal("preferred Ethernet BAR2 should own class-specific behavior")
	}
	if cfg.MSIXConfig != nil || cfg.HasMSIX {
		t.Fatalf("device without an MSI-X capability acquired MSI-X endpoint: config=%+v HasMSIX=%v", cfg.MSIXConfig, cfg.HasMSIX)
	}
}

func TestBuildSVConfigRejectsOccupied64BitUpperBIR(t *testing.T) {
	ctx := makeDonorContext(0x10EC, 0x8125, 0x020000)
	ctx.BARs = []pci.BAR{
		{Index: 2, Size: 0x1000, Type: pci.BARTypeMem64, Is64Bit: true},
		{Index: 3, Size: 0x1000, Type: pci.BARTypeMem32},
	}
	ctx.BARContents = map[int][]byte{2: make([]byte, 0x1000), 3: make([]byte, 0x1000)}
	ow := NewOutputWriter(t.TempDir(), "", 0, 0)
	ids := firmware.ExtractDeviceIDs(ctx.ConfigSpace, nil)

	_, err := ow.buildSVConfig(ctx, ctx.ConfigSpace.Clone(), ids, 0x1234, &board.Board{BRAMSize: 0x10000})
	if err == nil || !strings.Contains(err.Error(), "64-bit BAR2 consumes occupied BAR3") {
		t.Fatalf("buildSVConfig error = %v, want occupied upper BAR3 rejection", err)
	}
}

func TestBuildSVConfigPreservesMSIXBIRsAndOffsets(t *testing.T) {
	cs := configSpaceWithMSIX(0x1234, 0x5678, 0xFF0000, 2, 0x400, 4, 0x800, 4)
	ctx := &donor.DeviceContext{
		Device:      pci.PCIDevice{VendorID: 0x1234, DeviceID: 0x5678, ClassCode: 0xFF0000},
		ConfigSpace: cs,
		BARs: []pci.BAR{
			{Index: 2, Size: 0x2000, Type: pci.BARTypeMem32},
			{Index: 4, Size: 0x1000, Type: pci.BARTypeMem32},
		},
		BARContents: map[int][]byte{
			2: make([]byte, 0x2000),
			4: make([]byte, 0x1000),
		},
		MSIXData: &donor.MSIXData{
			TableSize: 4, TableBIR: 2, TableOffset: 0x400,
			PBABIR: 4, PBAOffset: 0x800,
		},
	}
	scrubbed := cs.Clone()
	ow := NewOutputWriter(t.TempDir(), "", 0, 0)
	ids := firmware.ExtractDeviceIDs(cs, nil)

	cfg, err := ow.buildSVConfig(ctx, scrubbed, ids, 0xCAFE, &board.Board{BRAMSize: 0x10000})
	if err != nil {
		t.Fatalf("buildSVConfig: %v", err)
	}
	if cfg.MSIXConfig == nil {
		t.Fatal("MSIXConfig is nil")
	}
	if got := cfg.MSIXConfig; got.TableBIR != 2 || got.TableOffset != 0x400 || got.PBABIR != 4 || got.PBAOffset != 0x800 {
		t.Fatalf("MSI-X topology changed: got table BAR%d+0x%X, PBA BAR%d+0x%X", got.TableBIR, got.TableOffset, got.PBABIR, got.PBAOffset)
	}
	if got := scrubbed.ReadU32(0x54); got != 0x00000402 {
		t.Errorf("scrubbed MSI-X table dword = 0x%08X, want offset 0x400 | BIR2", got)
	}
	if got := scrubbed.ReadU32(0x58); got != 0x00000804 {
		t.Errorf("scrubbed MSI-X PBA dword = 0x%08X, want offset 0x800 | BIR4", got)
	}
}

func TestBuildSVConfigRelocatesOverlappingMSIXRegionsWithoutChangingBIR(t *testing.T) {
	cs := configSpaceWithMSIX(0x10EC, 0x8125, 0x020000, 2, 0, 2, 0, 2)
	ctx := &donor.DeviceContext{
		Device:      pci.PCIDevice{VendorID: 0x10EC, DeviceID: 0x8125, ClassCode: 0x020000},
		ConfigSpace: cs,
		BARs: []pci.BAR{
			{Index: 2, Size: 0x1000, Type: pci.BARTypeMem32},
		},
		BARContents: map[int][]byte{2: make([]byte, 0x1000)},
		MSIXData: &donor.MSIXData{
			TableSize: 2, TableBIR: 2, TableOffset: 0,
			PBABIR: 2, PBAOffset: 0,
		},
	}
	scrubbed := cs.Clone()
	ow := NewOutputWriter(t.TempDir(), "", 0, 0)
	ids := firmware.ExtractDeviceIDs(cs, nil)

	cfg, err := ow.buildSVConfig(ctx, scrubbed, ids, 0xCAFE, &board.Board{BRAMSize: 0x10000})
	if err != nil {
		t.Fatalf("buildSVConfig: %v", err)
	}
	msix := cfg.MSIXConfig
	if msix == nil {
		t.Fatal("MSIXConfig is nil")
	}
	if msix.TableBIR != 2 || msix.PBABIR != 2 {
		t.Fatalf("MSI-X relocation changed advertised BIR: table=%d PBA=%d, want both 2", msix.TableBIR, msix.PBABIR)
	}
	if msix.TableOffset == 0 || msix.PBAOffset == 0 {
		t.Fatalf("overlapping donor offsets were not relocated: table=0x%X PBA=0x%X", msix.TableOffset, msix.PBAOffset)
	}
	tableEnd := msix.TableOffset + uint32(ctx.MSIXData.TableSize*16)
	pbaEnd := msix.PBAOffset + 8
	if msix.TableOffset < pbaEnd && msix.PBAOffset < tableEnd {
		t.Fatalf("relocated MSI-X table [0x%X,0x%X) overlaps PBA [0x%X,0x%X)",
			msix.TableOffset, tableEnd, msix.PBAOffset, pbaEnd)
	}
	if got, want := scrubbed.ReadU32(0x54), (msix.TableOffset&0xFFFFFFF8)|2; got != want {
		t.Errorf("table capability dword = 0x%08X, want routed value 0x%08X", got, want)
	}
	if got, want := scrubbed.ReadU32(0x58), (msix.PBAOffset&0xFFFFFFF8)|2; got != want {
		t.Errorf("PBA capability dword = 0x%08X, want routed value 0x%08X", got, want)
	}
}

func TestBuildSVConfigRejectsMSIXBIRWithoutMemoryEndpoint(t *testing.T) {
	cs := configSpaceWithMSIX(0x1234, 0x5678, 0xFF0000, 3, 0x100, 3, 0x200, 1)
	ctx := &donor.DeviceContext{
		Device:      pci.PCIDevice{VendorID: 0x1234, DeviceID: 0x5678, ClassCode: 0xFF0000},
		ConfigSpace: cs,
		BARs: []pci.BAR{
			{Index: 2, Size: 0x1000, Type: pci.BARTypeMem32},
			{Index: 3, Type: pci.BARTypeDisabled},
		},
		BARContents: map[int][]byte{2: make([]byte, 0x1000)},
		MSIXData: &donor.MSIXData{
			TableSize: 1, TableBIR: 3, TableOffset: 0x100,
			PBABIR: 3, PBAOffset: 0x200,
		},
	}
	ow := NewOutputWriter(t.TempDir(), "", 0, 0)
	ids := firmware.ExtractDeviceIDs(cs, nil)

	_, err := ow.buildSVConfig(ctx, cs.Clone(), ids, 0xCAFE, &board.Board{BRAMSize: 0x10000})
	if err == nil {
		t.Fatal("buildSVConfig accepted MSI-X regions in absent/disabled BAR3")
	}
	if msg := strings.ToLower(err.Error()); !strings.Contains(msg, "bar3") && !strings.Contains(msg, "bir 3") {
		t.Fatalf("error %q does not identify unsupported MSI-X BIR3", err)
	}
}

func TestCoreSVArtifactsHaveNoUnconditionalMSIEndpoint(t *testing.T) {
	for _, artifact := range coreSVArtifacts {
		if artifact.filename == "pcileech_bar_impl_msi.sv" {
			t.Fatal("pcileech_bar_impl_msi.sv must not be generated for every device")
		}
	}
}

func TestGenerateProjectTCLWithConfigPreservesSparseBARTopology(t *testing.T) {
	cs := fakeConfigSpace(0x1234, 0x5678, 0xFF0000)
	ctx := &donor.DeviceContext{
		Device:      pci.PCIDevice{VendorID: 0x1234, DeviceID: 0x5678, ClassCode: 0xFF0000},
		ConfigSpace: cs,
		BARs: []pci.BAR{
			{Index: 0, Size: 0x2000, Type: pci.BARTypeMem64, Is64Bit: true},
			{Index: 2, Size: 0x1000, Type: pci.BARTypeMem32, Prefetchable: true},
			{Index: 5, Size: 0x100000, Type: pci.BARTypeMem32},
		},
	}
	bar0 := &barmodel.BARModel{BIR: 0, Size: 0x2000, Aperture: 0x2000, Type: pci.BARTypeMem64, Is64Bit: true, UpperBIR: 1}
	bar2 := &barmodel.BARModel{BIR: 2, Size: 0x1000, Aperture: 0x1000, Type: pci.BARTypeMem32, Prefetchable: true}
	bar5 := &barmodel.BARModel{BIR: 5, Size: 0x100000, Aperture: 0x100000, Type: pci.BARTypeMem32}
	cfg := &svgen.SVGeneratorConfig{
		DonorBARTopology: true,
		BARModels:        []*barmodel.BARModel{bar0, bar2, bar5},
		BARModel:         bar0,
		MSIXConfig: &svgen.MSIXConfig{
			NumVectors: 8,
			TableBIR: 0, TableOffset: 0x400, TableBARSize: 0x2000,
			PBABIR: 5, PBAOffset: 0x800, PBABARSize: 0x100000,
		},
	}
	b := &board.Board{Name: "Test", FPGAPart: "xc7a35tfgg484-2", PCIeLanes: 1, TopModule: "test_top", BRAMSize: 0x200000}

	tcl := tclgen.GenerateProjectTCLWithConfig(ctx, b, "/tmp/lib", false, cfg)
	for bir, enabled := range []bool{true, false, true, false, false, true} {
		want := "CONFIG.Bar" + string(rune('0'+bir)) + "_Enabled "
		if enabled {
			want += "true"
		} else {
			want += "false"
		}
		if !strings.Contains(tcl, want) {
			t.Errorf("TCL missing %q", want)
		}
	}
	for _, want := range []string{
		"CONFIG.Bar0_Type            Memory",
		"CONFIG.Bar0_Scale           Kilobytes",
		"CONFIG.Bar0_Size            8",
		"CONFIG.Bar0_64bit           true",
		"CONFIG.Bar0_Prefetchable    false",
		"CONFIG.Bar2_Type            Memory",
		"CONFIG.Bar2_Scale           Kilobytes",
		"CONFIG.Bar2_Size            4",
		"CONFIG.Bar2_64bit           false",
		"CONFIG.Bar2_Prefetchable    true",
		"CONFIG.Bar5_Type            Memory",
		"CONFIG.Bar5_Scale           Megabytes",
		"CONFIG.Bar5_Size            1",
		"CONFIG.Bar5_Prefetchable    false",
		"CONFIG.MSIx_Table_BIR           BAR_1:0",
		"CONFIG.MSIx_Table_Offset        00000400",
		"CONFIG.MSIx_PBA_BIR             BAR_5",
		"CONFIG.MSIx_PBA_Offset          00000800",
	} {
		if !strings.Contains(tcl, want) {
			t.Errorf("TCL missing topology property %q", want)
		}
	}
	if strings.Contains(tcl, "CONFIG.Bar5_64bit") {
		t.Fatal("TCL emits unsupported Bar5_64bit property")
	}
}

func TestFinalizedMSIXValuesMatchCOETCLAndRTL(t *testing.T) {
	cs := configSpaceWithMSIX(0x1234, 0x5678, 0xFF0000, 2, 0x400, 4, 0x800, 4)
	ctx := &donor.DeviceContext{
		Device:      pci.PCIDevice{VendorID: 0x1234, DeviceID: 0x5678, ClassCode: 0xFF0000},
		ConfigSpace: cs,
		BARs: []pci.BAR{
			{Index: 2, Size: 0x2000, Type: pci.BARTypeMem32},
			{Index: 4, Size: 0x1000, Type: pci.BARTypeMem32},
		},
		BARContents: map[int][]byte{2: make([]byte, 0x2000), 4: make([]byte, 0x1000)},
		MSIXData: &donor.MSIXData{
			TableSize: 4, TableBIR: 2, TableOffset: 0x400,
			PBABIR: 4, PBAOffset: 0x800,
		},
	}
	b := &board.Board{Name: "Test", FPGAPart: "xc7a35tfgg484-2", PCIeLanes: 1, TopModule: "test_top", BRAMSize: 0x10000}
	dir := t.TempDir()
	ow := NewOutputWriter(dir, "/tmp/lib", 1, 1)
	scrubbed := cs.Clone()
	ids := firmware.ExtractDeviceIDs(cs, nil)
	cfg, err := ow.buildSVConfig(ctx, scrubbed, ids, 0x1234, b)
	if err != nil {
		t.Fatalf("buildSVConfig: %v", err)
	}
	if err := ow.writeConfigSpaceArtifacts(ctx, scrubbed, b); err != nil {
		t.Fatalf("writeConfigSpaceArtifacts: %v", err)
	}
	if err := ow.writeTCLScripts(ctx, b, cfg); err != nil {
		t.Fatalf("writeTCLScripts: %v", err)
	}
	if err := ow.writeCoreSVArtifacts(cfg, scrubbed); err != nil {
		t.Fatalf("writeCoreSVArtifacts: %v", err)
	}
	coe := readTestArtifact(t, filepath.Join(dir, "pcileech_cfgspace.coe"))
	tcl := readTestArtifact(t, filepath.Join(dir, "vivado_generate_project.tcl"))
	rtl := readTestArtifact(t, filepath.Join(dir, "pcileech_tlps128_bar_controller.sv"))
	if got := coeWord(t, coe, 0x54/4); got != "00000402" {
		t.Errorf("COE MSI-X table dword = %s, want 00000402", got)
	}
	if got := coeWord(t, coe, 0x58/4); got != "00000804" {
		t.Errorf("COE MSI-X PBA dword = %s, want 00000804", got)
	}
	for _, want := range []string{"CONFIG.MSIx_Table_BIR           BAR_2", "CONFIG.MSIx_Table_Offset        00000400", "CONFIG.MSIx_PBA_BIR             BAR_4", "CONFIG.MSIx_PBA_Offset          00000800"} {
		if !strings.Contains(tcl, want) {
			t.Errorf("TCL missing finalized MSI-X value %q", want)
		}
	}
	for _, want := range []string{".TABLE_BIR      ( 2", ".TABLE_OFFSET   ( 32'h00000400", ".PBA_BIR        ( 4", ".PBA_OFFSET     ( 32'h00000800"} {
		if !strings.Contains(rtl, want) {
			t.Errorf("RTL missing finalized MSI-X value %q", want)
		}
	}
}

func readTestArtifact(t *testing.T, path string) string {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return string(data)
}

func coeWord(t *testing.T, coe string, index int) string {
	t.Helper()
	parts := strings.SplitN(coe, "memory_initialization_vector=\n", 2)
	if len(parts) != 2 {
		t.Fatal("COE initialization vector missing")
	}
	lines := strings.Split(parts[1], "\n")
	if index < 0 || index >= len(lines) {
		t.Fatalf("COE word index %d out of range", index)
	}
	return strings.TrimSuffix(strings.TrimSuffix(strings.TrimSpace(lines[index]), ","), ";")
}

func TestHDLLintHarnessRequiresNamedModernMultibarPass(t *testing.T) {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("runtime.Caller failed")
	}
	script := readTestArtifact(t, filepath.Join(filepath.Dir(file), "..", "..", "..", "scripts", "hdl-lint.sh"))
	for _, want := range []string{
		"LEGACY_BOARDS=(pciescreamer NeTV2_35T NeTV2_100T acorn litefury sp605_ft601)",
		`case " ${LEGACY_BOARDS[*]} " in`,
		"SKIP  $cell (explicit legacy board allowlist)",
		"modern_multibar_pass=0",
		`if [ "$cell" = "multibar×PCIeSquirrel" ]; then`,
		"modern_multibar_pass=1",
		`if [ "$modern_multibar_pass" -ne 1 ]; then`,
		"FAIL  mandatory modern multibar×PCIeSquirrel cell did not pass",
		"pcileech_bar_rsp_arbiter.sv",
	} {
		if !strings.Contains(script, want) {
			t.Errorf("HDL harness missing policy token %q", want)
		}
	}
	for _, forbidden := range []string{
		"legacy board source lacks IfAXIS128 controller architecture",
		"legacy board source lacks generated-controller integration modules",
	} {
		if strings.Contains(script, forbidden) {
			t.Errorf("HDL harness retains content-based legacy skip %q", forbidden)
		}
	}
}


func configSpaceWithMSIX(vid, did uint16, classCode uint32, tableBIR int, tableOffset uint32, pbaBIR int, pbaOffset uint32, vectors int) *pci.ConfigSpace {
	cs := fakeConfigSpace(vid, did, classCode)
	cs.WriteU16(0x06, cs.Status()|0x0010)
	cs.WriteU8(0x34, 0x50)
	cs.WriteU8(0x50, pci.CapIDMSIX)
	cs.WriteU8(0x51, 0)
	cs.WriteU16(0x52, uint16(vectors-1)&0x07FF)
	cs.WriteU32(0x54, tableOffset&0xFFFFFFF8|uint32(tableBIR))
	cs.WriteU32(0x58, pbaOffset&0xFFFFFFF8|uint32(pbaBIR))
	return cs
}
