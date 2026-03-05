package board

import (
	"testing"
)

func TestFind_CaseInsensitive(t *testing.T) {
	b, err := Find("captaindma_100t")
	if err != nil {
		t.Fatalf("Find failed: %v", err)
	}
	if b.Name != "CaptainDMA_100T" {
		t.Errorf("expected CaptainDMA_100T, got %s", b.Name)
	}
}

func TestFind_Unknown(t *testing.T) {
	_, err := Find("nonexistent_board")
	if err == nil {
		t.Error("expected error for unknown board")
	}
}

func TestFind_AllBoards(t *testing.T) {
	names := ListNames()
	if len(names) < 10 {
		t.Errorf("expected at least 10 boards, got %d", len(names))
	}
	for _, name := range names {
		b, err := Find(name)
		if err != nil {
			t.Errorf("Find(%q) failed: %v", name, err)
			continue
		}
		if b.Name != name {
			t.Errorf("expected name %q, got %q", name, b.Name)
		}
		if b.FPGAPart == "" {
			t.Errorf("board %q has empty FPGAPart", name)
		}
		if b.TopModule == "" {
			t.Errorf("board %q has empty TopModule", name)
		}
		if b.PCIeLanes != 1 && b.PCIeLanes != 4 {
			t.Errorf("board %q has invalid PCIeLanes: %d", name, b.PCIeLanes)
		}
	}
}

func TestAll_ReturnsIndependentCopy(t *testing.T) {
	all1 := All()
	all2 := All()
	if len(all1) != len(all2) {
		t.Fatalf("different lengths: %d vs %d", len(all1), len(all2))
	}
	// mutating one shouldn't affect the other
	all1[0].Name = "mutated"
	if all2[0].Name == "mutated" {
		t.Error("All() should return independent copies")
	}
}

func TestBoard_String(t *testing.T) {
	b := &Board{Name: "TestBoard"}
	if b.String() != "TestBoard" {
		t.Errorf("String() = %q", b.String())
	}
}

func TestBoard_MaxLinkSpeedOrDefault(t *testing.T) {
	b := &Board{}
	if b.MaxLinkSpeedOrDefault() != 2 {
		t.Errorf("default should be Gen2, got %d", b.MaxLinkSpeedOrDefault())
	}
	b.MaxLinkSpeed = 3
	if b.MaxLinkSpeedOrDefault() != 3 {
		t.Errorf("should return 3, got %d", b.MaxLinkSpeedOrDefault())
	}
}

func TestBoard_BRAMSizeOrDefault(t *testing.T) {
	b := &Board{}
	if b.BRAMSizeOrDefault() != 4096 {
		t.Errorf("default should be 4096, got %d", b.BRAMSizeOrDefault())
	}
	b.BRAMSize = 8192
	if b.BRAMSizeOrDefault() != 8192 {
		t.Errorf("should return 8192, got %d", b.BRAMSizeOrDefault())
	}
}

func TestBoard_Paths(t *testing.T) {
	b := &Board{
		ProjectDir: "CaptainDMA",
		SubDir:     "100t484-1",
		TCLFile:    "vivado_generate_project.tcl",
		BuildTCL:   "vivado_build.tcl",
	}

	src := b.SrcPath("/fpga")
	if src != "/fpga/CaptainDMA/100t484-1/src" {
		t.Errorf("SrcPath = %q", src)
	}
	ip := b.IPPath("/fpga")
	if ip != "/fpga/CaptainDMA/100t484-1/ip" {
		t.Errorf("IPPath = %q", ip)
	}
	tcl := b.TCLPath("/fpga")
	if tcl != "/fpga/CaptainDMA/100t484-1/vivado_generate_project.tcl" {
		t.Errorf("TCLPath = %q", tcl)
	}
	build := b.BuildTCLPath("/fpga")
	if build != "/fpga/CaptainDMA/100t484-1/vivado_build.tcl" {
		t.Errorf("BuildTCLPath = %q", build)
	}
	lib := b.LibPath("/fpga")
	if lib != "/fpga/CaptainDMA/100t484-1" {
		t.Errorf("LibPath = %q", lib)
	}
}

func TestBoard_PathsNoSubDir(t *testing.T) {
	b := &Board{
		ProjectDir: "PCIeSquirrel",
		TCLFile:    "vivado_generate_project.tcl",
	}

	src := b.SrcPath("/fpga")
	if src != "/fpga/PCIeSquirrel/src" {
		t.Errorf("SrcPath = %q", src)
	}
	build := b.BuildTCLPath("/fpga")
	if build != "/fpga/PCIeSquirrel/vivado_build.tcl" {
		t.Errorf("BuildTCLPath should default, got %q", build)
	}
}

func TestBoard_ZDMA_BuildTCL(t *testing.T) {
	b, err := Find("ZDMA")
	if err != nil {
		t.Fatal(err)
	}
	if b.BuildTCL != "vivado_build_100t.tcl" {
		t.Errorf("ZDMA BuildTCL = %q", b.BuildTCL)
	}
	if b.PCIeLanes != 4 {
		t.Errorf("ZDMA should have 4 lanes, got %d", b.PCIeLanes)
	}
}
