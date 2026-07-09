package tclgen

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

func TestGenerateProjectTCLImportsNativeXCIPaths(t *testing.T) {
	libDir := filepath.Join(t.TempDir(), "PCILeech FPGA library")
	b := &board.Board{
		Name:       "TestBoard",
		FPGAPart:   "xc7a35tfgg484-2",
		PCIeLanes:  1,
		TopModule:  "test_top",
		ProjectDir: "board files",
	}
	ctx := projectTCLTestContext()

	tcl := GenerateProjectTCL(ctx, b, libDir, false)
	ipPath, err := filepath.Abs(b.IPPath(libDir))
	if err != nil {
		t.Fatalf("resolve IP path: %v", err)
	}
	wantGlob := filepath.ToSlash(filepath.Join(ipPath, "*.xci"))
	if !strings.Contains(tcl, `glob -nocomplain "`+wantGlob+`"`) {
		t.Fatalf("generated Tcl does not use normalized XCI glob %q", wantGlob)
	}
	importIndex := strings.Index(tcl, "import_ip")
	if importIndex < 0 {
		t.Fatal("generated Tcl does not load XCI files with import_ip")
	}
	if strings.Contains(tcl, "import_files -fileset sources_1 $ip_files") {
		t.Fatal("generated Tcl imports XCI files as generic sources instead of native IP")
	}
	lookupIndex := strings.Index(tcl, "set all_ips [get_ips -quiet *]")
	if lookupIndex < 0 || importIndex > lookupIndex {
		t.Fatalf("native XCI loading must precede IP lookup: import_ip=%d get_ips=%d", importIndex, lookupIndex)
	}
}

func TestGenerateProjectTCLFailsWhenRequiredIPIsMissing(t *testing.T) {
	b := &board.Board{
		Name:       "TestBoard",
		FPGAPart:   "xc7a35tfgg484-2",
		PCIeLanes:  1,
		TopModule:  "test_top",
		ProjectDir: "TestBoard",
	}
	tcl := GenerateProjectTCL(projectTCLTestContext(), b, t.TempDir(), false)

	ipFilesIndex := strings.Index(tcl, "set ip_files")
	importIndex := strings.Index(tcl, "import_ip")
	if ipFilesIndex < 0 || importIndex < 0 {
		t.Fatalf("generated Tcl lacks XCI discovery or native loading: set=%d import=%d", ipFilesIndex, importIndex)
	}
	ipGuard := tcl[ipFilesIndex:importIndex]
	if !strings.Contains(ipGuard, "error ") {
		t.Fatalf("generated Tcl does not fail when the board supplies no XCI files: %q", ipGuard)
	}

	pcieIndex := strings.Index(tcl, "set pcie_ip [get_ips -quiet pcie_7x_0]")
	topIndex := strings.Index(tcl, "set_property -name \"top\"")
	if pcieIndex < 0 || topIndex < 0 || pcieIndex >= topIndex {
		t.Fatalf("generated Tcl lacks bounded PCIe IP configuration section: pcie=%d top=%d", pcieIndex, topIndex)
	}
	pcieSection := tcl[pcieIndex:topIndex]
	if !strings.Contains(pcieSection, "error ") {
		t.Fatal("generated Tcl does not fail when pcie_7x_0 is absent")
	}
	if strings.Contains(pcieSection, "WARNING:") || strings.Contains(pcieSection, "skipping donor identity configuration") {
		t.Fatal("generated Tcl silently accepts a missing pcie_7x_0 core")
	}

	generateIndex := strings.Index(tcl, "generate_target all")
	if generateIndex < 0 || generateIndex < pcieIndex {
		t.Fatalf("generated Tcl must generate configured IP targets after required-IP validation: generate=%d pcie=%d", generateIndex, pcieIndex)
	}
}

func projectTCLTestContext() *donor.DeviceContext {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	cs.WriteU16(0x00, 0x10ee)
	cs.WriteU16(0x02, 0x7024)
	return &donor.DeviceContext{
		Device: pci.PCIDevice{
			VendorID:  0x10ee,
			DeviceID:  0x7024,
			ClassCode: 0x058000,
		},
		ConfigSpace: cs,
	}
}
