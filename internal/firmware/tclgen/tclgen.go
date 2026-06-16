package tclgen

import (
	"bytes"
	"fmt"
	"path/filepath"
	"text/template"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/firmware"
	"github.com/sercanarga/pcileechgen/internal/firmware/scrub"
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

	DSNEnabled       bool
	MSICapVectorsStr string

	// MSI-X
	MSIXEnabled     bool
	MSIXTableSize   int
	MSIXTableBIR    string
	MSIXTableOffset string
	MSIXPBABIR      string
	MSIXPBAOffset   string
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

func buildBAR0Config(bar0Size int) bar0Config {
	scale, size := barSizeToTCL(uint64(bar0Size))
	return bar0Config{
		Enabled: true,
		Scale:   scale,
		Size:    size,
		Is64bit: false,
	}
}

// GenerateProjectTCL generates the Vivado project creation TCL script.
func GenerateProjectTCL(ctx *donor.DeviceContext, b *board.Board, libDir string) string {
	ids := firmware.ExtractDeviceIDs(ctx.ConfigSpace, ctx.ExtCapabilities)

	// Use the board's max link speed for the Xilinx IP core.
	// The donor's current speed depends on the slot it was captured in,
	// not the device's capability. The FPGA should advertise its full
	// capability so the root complex negotiates the highest common speed.
	linkWidth := clampLinkWidth(ids.LinkWidth, b.PCIeLanes)
	linkSpeed := b.MaxLinkSpeedOrDefault()

	bar0Size := 4096
	if ctx.MSIXData != nil && ctx.MSIXData.TableSize > 0 {
		bar0Size = scrub.ComputeBAR0Size(ctx.MSIXData.TableSize, b.BRAMSizeOrDefault())
	}
	bar0 := buildBAR0Config(bar0Size)

	msiVectors := extractMSIVectors(ctx)

	srcAbs, _ := filepath.Abs(b.SrcPath(libDir))
	ipAbs, _ := filepath.Abs(b.IPPath(libDir))

	data := projectTCLData{
		BoardName:        b.Name,
		FPGAPart:         b.FPGAPart,
		SrcPath:          srcAbs,
		IPPath:           ipAbs,
		TopModule:        b.TopModule,
		DeviceID:         fmt.Sprintf("%04X", ctx.Device.DeviceID),
		VendorID:         fmt.Sprintf("%04X", ctx.Device.VendorID),
		RevisionID:       fmt.Sprintf("%02X", ctx.Device.RevisionID),
		SubsysVendorID:   fmt.Sprintf("%04X", ctx.Device.SubsysVendorID),
		SubsysDeviceID:   fmt.Sprintf("%04X", ctx.Device.SubsysDeviceID),
		ClassCodeBase:    fmt.Sprintf("%02X", (ctx.Device.ClassCode>>16)&0xFF),
		ClassCodeSub:     fmt.Sprintf("%02X", (ctx.Device.ClassCode>>8)&0xFF),
		ClassCodeIntf:    fmt.Sprintf("%02X", ctx.Device.ClassCode&0xFF),
		LinkSpeed:        linkSpeedToTCL(linkSpeed),
		LinkWidth:        linkWidthToTCL(linkWidth),
		TrgtLinkSpeed:    linkSpeedToTrgt(linkSpeed),
		Bar0Enabled:      bar0.Enabled,
		Bar0Size:         bar0.Size,
		Bar0Scale:        bar0.Scale,
		Bar064bit:        bar0.Is64bit,
		DSNEnabled:       ids.HasDSN,
		MSICapVectorsStr: msiVectorsToTCL(msiVectors),
	}

	if ctx.MSIXData != nil && ctx.MSIXData.TableSize > 0 {
		data.MSIXEnabled = true
		data.MSIXTableSize = ctx.MSIXData.TableSize - 1
		data.MSIXTableBIR = barBIRToTCL(0)
		tableSize := ctx.MSIXData.TableSize * 16
		tableOff := uint32(0x1000)
		if ctx.Device.ClassCode>>8 == 0x0108 && bar0Size > 0 {
			tableOff = uint32(bar0Size/2) &^ 0xF
			if tableOff < 0x2000 {
				tableOff = 0x2000
			}
			if tableOff >= 0x1000 && tableOff < 0x1000+uint32(tableSize) {
				tableOff = 0x2000
			}
			if tableOff < 0x40 {
				tableOff = 0x1000
			}
			if tableOff+uint32(tableSize)+16 > uint32(bar0Size) {
				tableOff = uint32(bar0Size) - uint32(tableSize) - 16
				tableOff &^= 0xF
				if tableOff < 0x1000 {
					tableOff = 0x1000
				}
			}
		}
		pbaOffset := tableOff + uint32(tableSize)
		pbaOffset = (pbaOffset + 7) &^ 7
		data.MSIXTableOffset = fmt.Sprintf("%08X", tableOff)
		data.MSIXPBABIR = barBIRToTCL(0)
		data.MSIXPBAOffset = fmt.Sprintf("%08X", pbaOffset)
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
