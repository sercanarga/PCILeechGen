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

// GenerateProjectTCL generates the Vivado project creation TCL script.
func GenerateProjectTCL(ctx *donor.DeviceContext, b *board.Board, libDir string) string {
	ids := firmware.ExtractDeviceIDs(ctx.ConfigSpace, ctx.ExtCapabilities)

	// Clamp link width to board physical lanes
	linkWidth := clampLinkWidth(ids.LinkWidth, b.PCIeLanes)
	linkSpeed := ids.LinkSpeed
	// Clamp link speed to board's physical capability
	if linkSpeed == 0 || linkSpeed > b.MaxLinkSpeedOrDefault() {
		linkSpeed = b.MaxLinkSpeedOrDefault()
	}

	// BAR0 configuration
	bar0Enabled := false
	bar0Scale := "Kilobytes"
	bar0Size := "4"
	bar064bit := false
	if len(ctx.BARs) > 0 {
		bar0 := ctx.BARs[0]
		if bar0.Size > 0 {
			bar0Enabled = true
			// clamp to FPGA BRAM
			bar0SizeClamped := bar0.Size
			bramSize := uint64(b.BRAMSizeOrDefault())
			if bar0SizeClamped > bramSize {
				bar0SizeClamped = bramSize
			}
			bar0Scale, bar0Size = barSizeToTCL(bar0SizeClamped)
			bar064bit = bar0.Type == pci.BARTypeMem64
		}
	}

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
		Bar0Enabled:    bar0Enabled,
		Bar0Size:       bar0Size,
		Bar0Scale:      bar0Scale,
		Bar064bit:      bar064bit,
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
