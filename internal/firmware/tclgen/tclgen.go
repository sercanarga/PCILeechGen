package tclgen

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"path/filepath"
	"text/template"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/firmware"
	"github.com/sercanarga/pcileechgen/internal/firmware/devclass"
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

	Bar0Enabled  bool
	Bar0Size     string
	Bar0Scale    string
	Bar064bit    bool
	Bar0ByteSize int
	StockBar     bool

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

func buildBAR0Config(bar0Size int, ctx *donor.DeviceContext) bar0Config {
	scale, size := barSizeToTCL(uint64(bar0Size))
	is64 := false
	if ctx != nil {
		if len(ctx.BARs) > 0 {
			raw := ctx.BARs[0].RawValue
			if raw != 0 {
				is64 = (raw & 0x06) == 0x04
			} else {
				p := devclass.ProfileForDevice(ctx.Device.ClassCode, ctx.Device.VendorID, ctx.Device.DeviceID)
				if p != nil {
					is64 = p.Uses64BitBAR
				}
			}
		} else if ctx.ConfigSpace != nil {
			raw := ctx.ConfigSpace.BAR(0)
			is64 = (raw & 0x06) == 0x04
		} else {
			p := devclass.ProfileForClass(ctx.Device.ClassCode)
			if p != nil {
				is64 = p.Uses64BitBAR
			}
		}
	}
	return bar0Config{
		Enabled: true,
		Scale:   scale,
		Size:    size,
		Is64bit: is64,
	}
}

// GenerateProjectTCL generates the Vivado project creation TCL script.
// stockBar: when true, forces the stock zerowrite4k BRAM COE path in the
// generated TCL (no donor content patch into bram_bar_zero4k), while still
// reporting the correct (donor-demanded) Bar0ByteSize for PCIe IP BAR sizing
// and other config. This matches --stock-bar CLI semantics.
func GenerateProjectTCL(ctx *donor.DeviceContext, b *board.Board, libDir string, stockBar bool) string {
	ids := firmware.ExtractDeviceIDs(ctx.ConfigSpace, ctx.ExtCapabilities)

	// Use the board's max link speed for the Xilinx IP core.
	// The donor's current speed depends on the slot it was captured in,
	// not the device's capability. The FPGA should advertise its full
	// capability so the root complex negotiates the highest common speed.
	linkWidth := clampLinkWidth(ids.LinkWidth, b.PCIeLanes)
	linkSpeed := b.MaxLinkSpeedOrDefault()

	msixTableSize := 0
	if ctx.MSIXData != nil && ctx.MSIXData.TableSize > 0 {
		msixTableSize = ctx.MSIXData.TableSize
	}
	// IMPORTANT: use *DonorBAR0Demand* (uncapped) here, not CappedBAR0Size.
	// This ensures the PCIe IP gets configured with the donor's actual BAR0
	// size (Bar0_Size/Scale), and Bar0ByteSize reports the correct value
	// (used for bram coe patch decision). --stock-bar (and force oversized
	// on small-BRAM boards) still report the donor size, but force the
	// zerowrite4k path for the bram IP.
	bar0Size := firmware.DonorBAR0Demand(ctx, b, msixTableSize)
	bar0 := buildBAR0Config(bar0Size, ctx)

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
		Bar0ByteSize:     bar0Size,
		StockBar:         stockBar,
		DSNEnabled:       ids.HasDSN,
		MSICapVectorsStr: msiVectorsToTCL(msiVectors),
	}

	if ctx.MSIXData != nil && ctx.MSIXData.TableSize > 0 {
		data.MSIXEnabled = true
		data.MSIXTableSize = ctx.MSIXData.TableSize - 1
		bir := ctx.MSIXData.TableBIR
		is64 := bar0.Is64bit && bir == 0
		data.MSIXTableBIR = barBIRToTCL(bir, is64)
		dstrd := uint32(0)
		if bar0d := firmware.LargestBar(ctx.BARContents); len(bar0d) >= 8 {
			dstrd = binary.LittleEndian.Uint32(bar0d[4:8]) & 0x0F
		}
		tableOff, pbaOffset, _ := firmware.MSIXPlacement(bar0Size, ctx.MSIXData.TableSize, ctx.Device.ClassCode, dstrd)
		data.MSIXTableOffset = fmt.Sprintf("%08X", tableOff)
		data.MSIXPBABIR = barBIRToTCL(bir, is64)
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
