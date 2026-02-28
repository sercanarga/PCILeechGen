package board

import (
	"testing"
)

func TestFindBoard(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{"PCIeSquirrel", "PCIeSquirrel", false},
		{"pciesquirrel", "PCIeSquirrel", false},
		{"PCIESQUIRREL", "PCIeSquirrel", false},
		{"ScreamerM2", "ScreamerM2", false},
		{"ZDMA", "ZDMA", false},
		{"CaptainDMA_100T", "CaptainDMA_100T", false},
		{"captaindma_100t", "CaptainDMA_100T", false},
		{"CaptainDMA_M2_x1", "CaptainDMA_M2_x1", false},
		{"CaptainDMA_M2_x4", "CaptainDMA_M2_x4", false},
		{"CaptainDMA_35T", "CaptainDMA_35T", false},
		{"CaptainDMA_75T", "CaptainDMA_75T", false},
		{"EnigmaX1", "EnigmaX1", false},
		{"NeTV2_35T", "NeTV2_35T", false},
		{"NeTV2_100T", "NeTV2_100T", false},
		{"acorn", "acorn", false},
		{"litefury", "litefury", false},
		{"nonexistent", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := Find(tt.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Find(%q) error = %v, wantErr %v", tt.name, err, tt.wantErr)
				return
			}
			if !tt.wantErr && b.Name != tt.want {
				t.Errorf("Find(%q).Name = %q, want %q", tt.name, b.Name, tt.want)
			}
		})
	}
}

func TestBoardPaths(t *testing.T) {
	libDir := "/path/to/lib/pcileech-fpga"

	// Simple board (no SubDir)
	b, _ := Find("PCIeSquirrel")
	if b.LibPath(libDir) != libDir+"/PCIeSquirrel" {
		t.Errorf("LibPath() = %q", b.LibPath(libDir))
	}
	if b.TCLPath(libDir) != libDir+"/PCIeSquirrel/vivado_generate_project.tcl" {
		t.Errorf("TCLPath() = %q", b.TCLPath(libDir))
	}
	if b.SrcPath(libDir) != libDir+"/PCIeSquirrel/src" {
		t.Errorf("SrcPath() = %q", b.SrcPath(libDir))
	}

	// Board with SubDir (CaptainDMA)
	cdma, _ := Find("CaptainDMA_100T")
	if cdma.LibPath(libDir) != libDir+"/CaptainDMA/100t484-1" {
		t.Errorf("CaptainDMA_100T LibPath() = %q", cdma.LibPath(libDir))
	}
	if cdma.SrcPath(libDir) != libDir+"/CaptainDMA/100t484-1/src" {
		t.Errorf("CaptainDMA_100T SrcPath() = %q", cdma.SrcPath(libDir))
	}
	if cdma.TCLPath(libDir) != libDir+"/CaptainDMA/100t484-1/vivado_generate_project_captaindma_100t.tcl" {
		t.Errorf("CaptainDMA_100T TCLPath() = %q", cdma.TCLPath(libDir))
	}

	// Board with custom BuildTCL (ZDMA)
	zdma, _ := Find("ZDMA")
	if zdma.BuildTCLPath(libDir) != libDir+"/ZDMA/vivado_build_100t.tcl" {
		t.Errorf("ZDMA BuildTCLPath() = %q", zdma.BuildTCLPath(libDir))
	}
}

func TestBoardString(t *testing.T) {
	b, _ := Find("PCIeSquirrel")
	if b.String() != "PCIeSquirrel" {
		t.Errorf("String() = %q, want PCIeSquirrel", b.String())
	}
}

func TestBoardIPPath(t *testing.T) {
	libDir := "/path/to/lib/pcileech-fpga"

	b, _ := Find("PCIeSquirrel")
	if b.IPPath(libDir) != libDir+"/PCIeSquirrel/ip" {
		t.Errorf("IPPath() = %q", b.IPPath(libDir))
	}

	cdma, _ := Find("CaptainDMA_100T")
	if cdma.IPPath(libDir) != libDir+"/CaptainDMA/100t484-1/ip" {
		t.Errorf("CaptainDMA_100T IPPath() = %q", cdma.IPPath(libDir))
	}
}

func TestBoardFPGAParts(t *testing.T) {
	tests := []struct {
		name     string
		wantPart string
		wantLane int
	}{
		{"PCIeSquirrel", "xc7a35tfgg484-2", 1},
		{"CaptainDMA_100T", "xc7a100tfgg484-2", 1},
		{"CaptainDMA_75T", "xc7a75tfgg484-2", 1},
		{"CaptainDMA_M2_x4", "xc7a35tcsg325-2", 4},
		{"ZDMA", "xc7a100tfgg484-2", 4},
		{"acorn", "xc7a200tfbg484-3", 4},
		{"litefury", "xc7a100tfgg484-2", 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := Find(tt.name)
			if err != nil {
				t.Fatal(err)
			}
			if b.FPGAPart != tt.wantPart {
				t.Errorf("FPGAPart = %q, want %q", b.FPGAPart, tt.wantPart)
			}
			if b.PCIeLanes != tt.wantLane {
				t.Errorf("PCIeLanes = %d, want %d", b.PCIeLanes, tt.wantLane)
			}
		})
	}
}

func TestListNames(t *testing.T) {
	names := ListNames()
	if len(names) < 16 {
		t.Errorf("ListNames() returned %d names, want at least 16", len(names))
	}

	// Check for key boards
	found := make(map[string]bool)
	for _, n := range names {
		found[n] = true
	}

	required := []string{
		"PCIeSquirrel", "ZDMA", "EnigmaX1",
		"CaptainDMA_M2_x1", "CaptainDMA_M2_x4", "CaptainDMA_35T",
		"CaptainDMA_75T", "CaptainDMA_100T",
		"NeTV2_35T", "NeTV2_100T",
		"acorn", "litefury",
	}
	for _, req := range required {
		if !found[req] {
			t.Errorf("ListNames() missing %q", req)
		}
	}
}

func TestAllBoards(t *testing.T) {
	boards := All()
	if len(boards) == 0 {
		t.Error("All() returned empty list")
	}

	for _, b := range boards {
		if b.Name == "" {
			t.Error("Board with empty name found")
		}
		if b.FPGAPart == "" {
			t.Errorf("Board %q has empty FPGAPart", b.Name)
		}
		if b.ProjectDir == "" {
			t.Errorf("Board %q has empty ProjectDir", b.Name)
		}
		if b.TopModule == "" {
			t.Errorf("Board %q has empty TopModule", b.Name)
		}
		if b.TCLFile == "" {
			t.Errorf("Board %q has empty TCLFile", b.Name)
		}
		if b.PCIeLanes != 1 && b.PCIeLanes != 4 {
			t.Errorf("Board %q has invalid PCIeLanes: %d", b.Name, b.PCIeLanes)
		}
	}
}

func TestFindBoardErrorMessage(t *testing.T) {
	_, err := Find("nonexistent")
	if err == nil {
		t.Fatal("expected error for nonexistent board")
	}
	// Error should contain available boards
	errMsg := err.Error()
	if len(errMsg) < 100 {
		t.Errorf("Error message too short, should list available boards: %s", errMsg)
	}
}
