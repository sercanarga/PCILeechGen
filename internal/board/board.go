// Package board provides PCILeech FPGA board definitions and discovery.
package board

import (
	"fmt"
	"path/filepath"
	"strings"
)

// Board represents a supported PCILeech FPGA board (or board variant).
type Board struct {
	Name       string `json:"name"`        // canonical board name (unique key)
	FPGAPart   string `json:"fpga_part"`   // Xilinx FPGA part number (e.g. xc7a35tfgg484-2)
	PCIeLanes  int    `json:"pcie_lanes"`  // number of PCIe lanes (1 or 4)
	TopModule  string `json:"top_module"`  // top-level SystemVerilog module name
	ProjectDir string `json:"project_dir"` // top-level directory in pcileech-fpga (e.g. "CaptainDMA")
	SubDir     string `json:"sub_dir"`     // optional subdirectory within ProjectDir (e.g. "100t484-1")
	TCLFile    string `json:"tcl_file"`    // TCL project generation script filename
	BuildTCL   string `json:"build_tcl"`   // TCL build script filename (defaults to "vivado_build.tcl")
}

// String returns the board name.
func (b *Board) String() string {
	return b.Name
}

// SrcPath returns the path to source files for this board.
func (b *Board) SrcPath(libDir string) string {
	if b.SubDir != "" {
		return filepath.Join(libDir, b.ProjectDir, b.SubDir, "src")
	}
	return filepath.Join(libDir, b.ProjectDir, "src")
}

// IPPath returns the path to IP cores for this board.
func (b *Board) IPPath(libDir string) string {
	if b.SubDir != "" {
		return filepath.Join(libDir, b.ProjectDir, b.SubDir, "ip")
	}
	return filepath.Join(libDir, b.ProjectDir, "ip")
}

// TCLPath returns the full path to the Vivado project generation TCL script.
func (b *Board) TCLPath(libDir string) string {
	if b.SubDir != "" {
		return filepath.Join(libDir, b.ProjectDir, b.SubDir, b.TCLFile)
	}
	return filepath.Join(libDir, b.ProjectDir, b.TCLFile)
}

// BuildTCLPath returns the full path to the Vivado build TCL script.
func (b *Board) BuildTCLPath(libDir string) string {
	buildFile := b.BuildTCL
	if buildFile == "" {
		buildFile = "vivado_build.tcl"
	}
	if b.SubDir != "" {
		return filepath.Join(libDir, b.ProjectDir, b.SubDir, buildFile)
	}
	return filepath.Join(libDir, b.ProjectDir, buildFile)
}

// LibPath returns the base path for this board variant within pcileech-fpga.
func (b *Board) LibPath(libDir string) string {
	if b.SubDir != "" {
		return filepath.Join(libDir, b.ProjectDir, b.SubDir)
	}
	return filepath.Join(libDir, b.ProjectDir)
}

// registry holds all supported boards and their variants.
// Data sourced directly from pcileech-fpga submodule TCL files.
var registry = []Board{
	// ─── PCIeSquirrel ───────────────────────────────────────────
	{
		Name:       "PCIeSquirrel",
		FPGAPart:   "xc7a35tfgg484-2",
		PCIeLanes:  1,
		TopModule:  "pcileech_squirrel_top",
		ProjectDir: "PCIeSquirrel",
		TCLFile:    "vivado_generate_project.tcl",
	},

	// ─── ScreamerM2 ────────────────────────────────────────────
	{
		Name:       "ScreamerM2",
		FPGAPart:   "xc7a35tcsg325-2",
		PCIeLanes:  1,
		TopModule:  "pcileech_screamer_m2_top",
		ProjectDir: "ScreamerM2",
		TCLFile:    "vivado_generate_project.tcl",
	},

	// ─── pciescreamer (original) ───────────────────────────────
	{
		Name:       "pciescreamer",
		FPGAPart:   "xc7a35tfgg484-2",
		PCIeLanes:  1,
		TopModule:  "pcileech_pciescreamer_top",
		ProjectDir: "pciescreamer",
		TCLFile:    "vivado_generate_project.tcl",
	},

	// ─── EnigmaX1 ──────────────────────────────────────────────
	{
		Name:       "EnigmaX1",
		FPGAPart:   "xc7a75tfgg484-2",
		PCIeLanes:  1,
		TopModule:  "pcileech_enigma_x1_top",
		ProjectDir: "EnigmaX1",
		TCLFile:    "vivado_generate_project.tcl",
	},

	// ─── CaptainDMA M2 x1 (35T, CSG325 package) ───────────────
	{
		Name:       "CaptainDMA_M2_x1",
		FPGAPart:   "xc7a35tcsg325-2",
		PCIeLanes:  1,
		TopModule:  "pcileech_35t325_x1_top",
		ProjectDir: "CaptainDMA",
		SubDir:     "35t325_x1",
		TCLFile:    "vivado_generate_project_captaindma_m2x1.tcl",
	},

	// ─── CaptainDMA M2 x4 (35T, CSG325 package) ───────────────
	{
		Name:       "CaptainDMA_M2_x4",
		FPGAPart:   "xc7a35tcsg325-2",
		PCIeLanes:  4,
		TopModule:  "pcileech_35t325_x4_top",
		ProjectDir: "CaptainDMA",
		SubDir:     "35t325_x4",
		TCLFile:    "vivado_generate_project_captaindma_m2x4.tcl",
	},

	// ─── CaptainDMA 4.1th (35T, FGG484 package) ───────────────
	{
		Name:       "CaptainDMA_35T",
		FPGAPart:   "xc7a35tfgg484-2",
		PCIeLanes:  1,
		TopModule:  "pcileech_35t484_x1_top",
		ProjectDir: "CaptainDMA",
		SubDir:     "35t484_x1",
		TCLFile:    "vivado_generate_project_captaindma_35t.tcl",
	},

	// ─── CaptainDMA 75T ────────────────────────────────────────
	{
		Name:       "CaptainDMA_75T",
		FPGAPart:   "xc7a75tfgg484-2",
		PCIeLanes:  1,
		TopModule:  "pcileech_75t484_x1_top",
		ProjectDir: "CaptainDMA",
		SubDir:     "75t484_x1",
		TCLFile:    "vivado_generate_project_captaindma_75t.tcl",
	},

	// ─── CaptainDMA 100T ───────────────────────────────────────
	{
		Name:       "CaptainDMA_100T",
		FPGAPart:   "xc7a100tfgg484-2",
		PCIeLanes:  1,
		TopModule:  "pcileech_100t484_x1_top",
		ProjectDir: "CaptainDMA",
		SubDir:     "100t484-1",
		TCLFile:    "vivado_generate_project_captaindma_100t.tcl",
	},

	// ─── ZDMA (LambdaConcept / LightingZDMA - 100T) ───────────
	{
		Name:       "ZDMA",
		FPGAPart:   "xc7a100tfgg484-2",
		PCIeLanes:  4,
		TopModule:  "pcileech_tbx4_100t_top",
		ProjectDir: "ZDMA",
		TCLFile:    "vivado_generate_project_100t.tcl",
		BuildTCL:   "vivado_build_100t.tcl",
	},

	// ─── GBOX ──────────────────────────────────────────────────
	{
		Name:       "GBOX",
		FPGAPart:   "xc7a35tfgg484-2",
		PCIeLanes:  1,
		TopModule:  "pcileech_gbox_top",
		ProjectDir: "GBOX",
		TCLFile:    "vivado_generate_project.tcl",
	},

	// ─── NeTV2 (35T variant) ───────────────────────────────────
	{
		Name:       "NeTV2_35T",
		FPGAPart:   "xc7a35tfgg484-2",
		PCIeLanes:  1,
		TopModule:  "pcileech_netv2_top",
		ProjectDir: "NeTV2",
		TCLFile:    "vivado_generate_project_35t.tcl",
	},

	// ─── NeTV2 (100T variant) ──────────────────────────────────
	{
		Name:       "NeTV2_100T",
		FPGAPart:   "xc7a100tfgg484-2",
		PCIeLanes:  1,
		TopModule:  "pcileech_netv2_top",
		ProjectDir: "NeTV2",
		TCLFile:    "vivado_generate_project_100t.tcl",
	},

	// ─── ac701_ft601 ───────────────────────────────────────────
	{
		Name:       "ac701_ft601",
		FPGAPart:   "xc7a200tfbg676-2",
		PCIeLanes:  4,
		TopModule:  "pcileech_ac701_ft601_top",
		ProjectDir: "ac701_ft601",
		TCLFile:    "vivado_generate_project.tcl",
	},

	// ─── Acorn (SQRL Acorn CLE-215+) ──────────────────────────
	{
		Name:       "acorn",
		FPGAPart:   "xc7a200tfbg484-3",
		PCIeLanes:  4,
		TopModule:  "pcileech_acorn_top",
		ProjectDir: "acorn_ft2232h",
		TCLFile:    "vivado_generate_project_acorn.tcl",
	},

	// ─── LiteFury (RHS Research LiteFury) ──────────────────────
	{
		Name:       "litefury",
		FPGAPart:   "xc7a100tfgg484-2",
		PCIeLanes:  4,
		TopModule:  "pcileech_acorn_top",
		ProjectDir: "acorn_ft2232h",
		TCLFile:    "vivado_generate_project_litefury.tcl",
	},

	// ─── sp605_ft601 (legacy) ──────────────────────────────────
	{
		Name:       "sp605_ft601",
		FPGAPart:   "xc6slx45tfgg484-2",
		PCIeLanes:  1,
		TopModule:  "pcileech_top",
		ProjectDir: "sp605_ft601",
		TCLFile:    "vivado_generate_project.tcl",
	},
}

// Find looks up a board by name (case-insensitive).
func Find(name string) (*Board, error) {
	lower := strings.ToLower(name)
	for i := range registry {
		if strings.ToLower(registry[i].Name) == lower {
			return &registry[i], nil
		}
	}
	return nil, fmt.Errorf("unknown board %q, available boards:\n%s",
		name, formatBoardList())
}

// formatBoardList returns a formatted list of available boards for error messages.
func formatBoardList() string {
	var sb strings.Builder
	for _, b := range registry {
		sb.WriteString(fmt.Sprintf("  %-25s %s (x%d)\n", b.Name, b.FPGAPart, b.PCIeLanes))
	}
	return sb.String()
}

// ListNames returns all available board names.
func ListNames() []string {
	names := make([]string, len(registry))
	for i, b := range registry {
		names[i] = b.Name
	}
	return names
}

// All returns all registered boards.
func All() []Board {
	result := make([]Board, len(registry))
	copy(result, registry)
	return result
}
