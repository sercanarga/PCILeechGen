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

// buildBAR0Config returns a fixed 4 KB BAR config for the Xilinx IP core.
// The FPGA's config space BRAM only serves 4 KB regions regardless of the
// donor's original BAR size. Matching the IP size to the actual served region
// prevents the driver from mapping more MMIO than the FPGA can back.
// BAR type is always 32-bit: the scrubber (clampBARsToFPGA) forces 32-bit
// type in the shadow config space. The IP core must match to avoid BAR
// sizing inconsistency that causes Code 10 on Windows.
func buildBAR0Config(ctx *donor.DeviceContext, b *board.Board) bar0Config {
	// Check if any memory BAR is enabled in the donor config
	for i := range ctx.BARs {
		bar := &ctx.BARs[i]
		if bar.Size > 0 && (bar.Type == pci.BARTypeMem32 || bar.Type == pci.BARTypeMem64) {
			return bar0Config{
				Enabled: true,
				Scale:   "Kilobytes",
				Size:    "4",
				Is64bit: false, // must match scrubber: always 32-bit
			}
		}
	}
	// No memory BAR found
	return bar0Config{Scale: "Kilobytes", Size: "4"}
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

	bar0 := buildBAR0Config(ctx, b)

	// Extract MSI vector count from donor capabilities
	msiVectors := extractMSIVectors(ctx)

	// Resolve to absolute paths so TCL works from any working directory
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
		data.MSIXTableSize = ctx.MSIXData.TableSize - 1 // Vivado expects N-1
		data.MSIXTableBIR = barBIRToTCL(ctx.MSIXData.TableBIR)
		data.MSIXTableOffset = fmt.Sprintf("%08X", ctx.MSIXData.TableOffset)
		data.MSIXPBABIR = barBIRToTCL(ctx.MSIXData.PBABIR)
		data.MSIXPBAOffset = fmt.Sprintf("%08X", ctx.MSIXData.PBAOffset)
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
