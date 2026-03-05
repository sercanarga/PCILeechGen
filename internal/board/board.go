// Package board provides PCILeech FPGA board definitions and discovery.
package board

import (
	"embed"
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"
)

// Board represents a supported PCILeech FPGA board (or board variant).
type Board struct {
	Name         string `json:"name"`
	FPGAPart     string `json:"fpga_part"`
	PCIeLanes    int    `json:"pcie_lanes"`
	MaxLinkSpeed uint8  `json:"max_link_speed"`
	BRAMSize     int    `json:"bram_size"`
	TopModule    string `json:"top_module"`
	ProjectDir   string `json:"project_dir"`
	SubDir       string `json:"sub_dir"`
	TCLFile      string `json:"tcl_file"`
	BuildTCL     string `json:"build_tcl"`
}

// String returns the board name.
func (b *Board) String() string { return b.Name }

// MaxLinkSpeedOrDefault returns the board's max link speed, defaulting to Gen2.
func (b *Board) MaxLinkSpeedOrDefault() uint8 {
	if b.MaxLinkSpeed > 0 {
		return b.MaxLinkSpeed
	}
	return 2
}

// BRAMSizeOrDefault returns the board's BAR BRAM size, defaulting to 4096.
func (b *Board) BRAMSizeOrDefault() int {
	if b.BRAMSize > 0 {
		return b.BRAMSize
	}
	return 4096
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

//go:embed boards.json
var boardsJSON embed.FS

// registry holds all supported boards, loaded from embedded JSON at init.
var registry []Board

func init() {
	data, err := boardsJSON.ReadFile("boards.json")
	if err != nil {
		panic("failed to read embedded boards.json: " + err.Error())
	}
	if err := json.Unmarshal(data, &registry); err != nil {
		panic("failed to parse embedded boards.json: " + err.Error())
	}
	// Validate: every board must have name, fpga_part, top_module
	for i, b := range registry {
		if b.Name == "" || b.FPGAPart == "" || b.TopModule == "" {
			panic(fmt.Sprintf("boards.json entry %d: missing required field (name=%q, fpga_part=%q, top_module=%q)",
				i, b.Name, b.FPGAPart, b.TopModule))
		}
		if b.PCIeLanes == 0 {
			registry[i].PCIeLanes = 1
		}
	}
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
