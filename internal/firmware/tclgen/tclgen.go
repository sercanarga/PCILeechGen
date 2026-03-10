package tclgen

import (
	"bytes"
	"fmt"
	"path/filepath"
	"text/template"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/firmware"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

// projectTCLData holds template data for Vivado project generation.
type projectTCLData struct {
	BoardName string
	FPGAPart  string
	SrcPath   string
	IPPath    string
	TopModule string

	DeviceID       string
	VendorID       string
	RevisionID     string
	SubsysVendorID string
	SubsysDeviceID string
	ClassCodeBase  string
	ClassCodeSub   string
	ClassCodeIntf  string

	LinkSpeed     string
	LinkWidth     string
	TrgtLinkSpeed string

	Bar0Enabled bool
	Bar0Size    string
	Bar0Scale   string
	Bar064bit   bool

	DSNEnabled bool
}

// buildTCLData holds template data for Vivado build script.
type buildTCLData struct {
	BoardName string
	FPGAPart  string
	Jobs      int
	Timeout   int
}

var projectTCLTmpl = template.Must(template.New("project").ParseFS(templateFS, "templates/project.tcl.tmpl"))

var buildTCLTmpl = template.Must(template.New("build").ParseFS(templateFS, "templates/build.tcl.tmpl"))

// bar0Config holds BAR0 parameters for the TCL project template.
type bar0Config struct {
	Enabled bool
	Scale   string
	Size    string
	Is64bit bool
}

// buildBAR0Config picks the largest MMIO BAR for the PCIe IP.
// Falls back to BAR0 when there's no MMIO BAR.
// When MSI-X is present, ensures the BAR is large enough to contain the
// relocated MSI-X table (0x1000+) and PBA.
func buildBAR0Config(ctx *donor.DeviceContext, b *board.Board) bar0Config {
	cfg := bar0Config{Scale: "Kilobytes", Size: "4"}
	if len(ctx.BARs) == 0 {
		return cfg
	}

	// Find the largest MMIO BAR - this is the primary BAR the driver will use
	var bestSize uint64
	var bestBAR *pci.BAR
	for i := range ctx.BARs {
		bar := &ctx.BARs[i]
		if bar.Size == 0 {
			continue
		}
		if bar.Type == pci.BARTypeMem32 || bar.Type == pci.BARTypeMem64 {
			if bar.Size > bestSize {
				bestSize = bar.Size
				bestBAR = bar
			}
		}
	}

	// Fall back to BAR0 if no MMIO BAR found
	if bestBAR == nil {
		bar0 := &ctx.BARs[0]
		if bar0.Size == 0 {
			return cfg
		}
		cfg.Enabled = true
		barSize := bar0.Size
		bramSize := uint64(b.BRAMSizeOrDefault())
		if barSize > bramSize {
			barSize = bramSize
		}
		barSize = enforceMinMSIXSize(barSize, ctx)
		cfg.Scale, cfg.Size = barSizeToTCL(barSize)
		cfg.Is64bit = bar0.Type == pci.BARTypeMem64
		return cfg
	}

	cfg.Enabled = true
	barSize := bestBAR.Size
	bramSize := uint64(b.BRAMSizeOrDefault())
	if barSize > bramSize {
		barSize = bramSize
	}
	barSize = enforceMinMSIXSize(barSize, ctx)
	cfg.Scale, cfg.Size = barSizeToTCL(barSize)
	cfg.Is64bit = bestBAR.Type == pci.BARTypeMem64
	return cfg
}

// enforceMinMSIXSize bumps the BAR size to fit the relocated MSI-X table
// (starting at offset 0x1000) plus its PBA. Returns the next power-of-two
// that covers everything.
func enforceMinMSIXSize(barSize uint64, ctx *donor.DeviceContext) uint64 {
	if ctx.MSIXData == nil || ctx.MSIXData.TableSize == 0 {
		return barSize
	}
	tableBytes := uint64(ctx.MSIXData.TableSize) * 16
	pbaBytes := ((uint64(ctx.MSIXData.TableSize) + 63) / 64) * 8
	if pbaBytes < 8 {
		pbaBytes = 8
	}
	// table starts at 0x1000, PBA right after (8-byte aligned)
	pbaStart := (0x1000 + tableBytes + 7) &^ 7
	minSize := pbaStart + pbaBytes

	// round up to next power of two
	if minSize > barSize {
		barSize = 1
		for barSize < minSize {
			barSize <<= 1
		}
	}
	return barSize
}

// GenerateProjectTCL generates the Vivado project creation TCL script.
func GenerateProjectTCL(ctx *donor.DeviceContext, b *board.Board, libDir string) string {
	ids := firmware.ExtractDeviceIDs(ctx.ConfigSpace, ctx.ExtCapabilities)

	// Clamp link width/speed to board physical capability
	linkWidth := clampLinkWidth(ids.LinkWidth, b.PCIeLanes)
	linkSpeed := ids.LinkSpeed
	if linkSpeed == 0 || linkSpeed > b.MaxLinkSpeedOrDefault() {
		linkSpeed = b.MaxLinkSpeedOrDefault()
	}

	bar0 := buildBAR0Config(ctx, b)

	// Resolve to absolute paths so TCL works from any working directory
	srcAbs, _ := filepath.Abs(b.SrcPath(libDir))
	ipAbs, _ := filepath.Abs(b.IPPath(libDir))

	data := projectTCLData{
		BoardName:      b.Name,
		FPGAPart:       b.FPGAPart,
		SrcPath:        srcAbs,
		IPPath:         ipAbs,
		TopModule:      b.TopModule,
		DeviceID:       fmt.Sprintf("%04X", ctx.Device.DeviceID),
		VendorID:       fmt.Sprintf("%04X", ctx.Device.VendorID),
		RevisionID:     fmt.Sprintf("%02X", ctx.Device.RevisionID),
		SubsysVendorID: fmt.Sprintf("%04X", ctx.Device.SubsysVendorID),
		SubsysDeviceID: fmt.Sprintf("%04X", ctx.Device.SubsysDeviceID),
		ClassCodeBase:  fmt.Sprintf("%02X", (ctx.Device.ClassCode>>16)&0xFF),
		ClassCodeSub:   fmt.Sprintf("%02X", (ctx.Device.ClassCode>>8)&0xFF),
		ClassCodeIntf:  fmt.Sprintf("%02X", ctx.Device.ClassCode&0xFF),
		LinkSpeed:      linkSpeedToTCL(linkSpeed),
		LinkWidth:      linkWidthToTCL(linkWidth),
		TrgtLinkSpeed:  linkSpeedToTrgt(linkSpeed),
		Bar0Enabled:    bar0.Enabled,
		Bar0Size:       bar0.Size,
		Bar0Scale:      bar0.Scale,
		Bar064bit:      bar0.Is64bit,
		DSNEnabled:     ids.HasDSN,
	}

	var buf bytes.Buffer
	if err := projectTCLTmpl.ExecuteTemplate(&buf, "project.tcl.tmpl", data); err != nil {
		panic(fmt.Sprintf("project TCL template error: %v", err))
	}
	return buf.String()
}

// GenerateBuildTCL generates the Vivado build/synthesis TCL script.
func GenerateBuildTCL(b *board.Board, jobs int, timeout int) string {
	if jobs <= 0 {
		jobs = 4
	}
	if timeout <= 0 {
		timeout = 3600
	}

	data := buildTCLData{
		BoardName: b.Name,
		FPGAPart:  b.FPGAPart,
		Jobs:      jobs,
		Timeout:   timeout,
	}

	var buf bytes.Buffer
	if err := buildTCLTmpl.ExecuteTemplate(&buf, "build.tcl.tmpl", data); err != nil {
		panic(fmt.Sprintf("build TCL template error: %v", err))
	}
	return buf.String()
}
