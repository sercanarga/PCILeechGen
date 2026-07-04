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
	// SourceSubDir is mutually exclusive with SubDir; it only overrides source/IP paths.
	SourceSubDir string `json:"source_sub_dir"`
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

const DefaultBRAMSize = 4096

// BRAMSizeOrDefault returns the board's BAR BRAM size, defaulting to 4096.
func (b *Board) BRAMSizeOrDefault() int {
	if b != nil && b.BRAMSize > 0 {
		return b.BRAMSize
	}
	return DefaultBRAMSize
}

// BRAM36Capacity returns the FPGA's RAMB36 block count parsed from the part
// number, used to size the NVMe disk cache. Returns 0 for non-Artix-7 parts.
func (b *Board) BRAM36Capacity() int {
	if b == nil {
		return 0
	}
	p := strings.ToLower(b.FPGAPart)
	switch {
	case strings.Contains(p, "7a200t"):
		return 365
	case strings.Contains(p, "7a100t"):
		return 135
	case strings.Contains(p, "7a75t"):
		return 105
	case strings.Contains(p, "7a50t"):
		return 65
	case strings.Contains(p, "7a35t"):
		return 50
	case strings.Contains(p, "7a15t"):
		return 25
	default:
		return 0
	}
}

// SourceBasePath returns the directory that contains source and IP folders.
func (b *Board) SourceBasePath(libDir string) string {
	if b.SourceSubDir != "" {
		return filepath.Join(libDir, b.ProjectDir, b.SourceSubDir)
	}
	return b.LibPath(libDir)
}

// SrcPath returns the path to source files for this board.
func (b *Board) SrcPath(libDir string) string {
	return filepath.Join(b.SourceBasePath(libDir), "src")
}

// IPPath returns the path to IP cores for this board.
func (b *Board) IPPath(libDir string) string {
	return filepath.Join(b.SourceBasePath(libDir), "ip")
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
