package tclgen

import (
	"bytes"
	"fmt"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/firmware"
	"github.com/sercanarga/pcileechgen/internal/firmware/devclass"
	"github.com/sercanarga/pcileechgen/internal/firmware/svgen"
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
	ImportVFiles bool
	BARs         []barTCLConfig

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

type barTCLConfig struct {
	Index        int
	Enabled      bool
	Scale        string
	Size         string
	Is64bit      bool
	Prefetchable bool
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
				p := devclass.ProfileForClass(ctx.Device.ClassCode)
				if p != nil {
					is64 = p.Uses64BitBAR
				}
			}
		} else if ctx.ConfigSpace != nil {
			is64 = (ctx.ConfigSpace.BAR(0) & 0x06) == 0x04
		} else {
			p := devclass.ProfileForClass(ctx.Device.ClassCode)
			if p != nil {
				is64 = p.Uses64BitBAR
			}
		}
	}
	return bar0Config{Enabled: true, Scale: scale, Size: size, Is64bit: is64}
}

func GenerateProjectTCL(ctx *donor.DeviceContext, b *board.Board, libDir string, stockBar bool) string {
	return generateProjectTCL(ctx, b, libDir, stockBar, nil)
}

func GenerateProjectTCLWithConfig(ctx *donor.DeviceContext, b *board.Board, libDir string, stockBar bool, cfg *svgen.SVGeneratorConfig) string {
	return generateProjectTCL(ctx, b, libDir, stockBar, cfg)
}

func configBARIs64(cfg *svgen.SVGeneratorConfig, bir int) bool {
	if cfg == nil {
		return false
	}
	for _, model := range cfg.BARModels {
		if model != nil && model.BIR == bir {
			return model.Is64Bit
		}
	}
	return false
}

func generateProjectTCL(ctx *donor.DeviceContext, b *board.Board, libDir string, stockBar bool, cfg *svgen.SVGeneratorConfig) string {
	ids := firmware.ExtractDeviceIDs(ctx.ConfigSpace, ctx.ExtCapabilities)
	linkWidth := clampLinkWidth(ids.LinkWidth, b.PCIeLanes)
	linkSpeed := b.MaxLinkSpeedOrDefault()
	msixTableSize := 0
	if ctx.MSIXData != nil && ctx.MSIXData.TableSize > 0 {
		msixTableSize = ctx.MSIXData.TableSize
	}
	bar0Size := firmware.DonorBAR0Demand(ctx, b, msixTableSize)
	bar0 := buildBAR0Config(bar0Size, ctx)
	bars := make([]barTCLConfig, 6)
	for i := range bars {
		bars[i].Index = i
	}
	if cfg != nil && cfg.DonorBARTopology {
		for _, model := range cfg.BARModels {
			if model == nil || model.BIR < 0 || model.BIR >= len(bars) {
				continue
			}
			scale, size := barSizeToTCL(uint64(model.Size))
			bars[model.BIR] = barTCLConfig{
				Index: model.BIR, Enabled: true, Scale: scale, Size: size,
				Is64bit: model.Is64Bit, Prefetchable: model.Prefetchable,
			}
			if model.BIR == 0 {
				bar0Size = model.Size
				bar0 = buildBAR0Config(bar0Size, ctx)
				bar0.Is64bit = model.Is64Bit
			}
		}
	} else {
		bars[0] = barTCLConfig{Index: 0, Enabled: bar0.Enabled, Scale: bar0.Scale, Size: bar0.Size, Is64bit: bar0.Is64bit}
	}
	srcAbs, _ := filepath.Abs(b.SrcPath(libDir))
	ipAbs, _ := filepath.Abs(b.IPPath(libDir))
	srcAbs = tclPath(srcAbs)
	ipAbs = tclPath(ipAbs)
	data := projectTCLData{
		BoardName: b.Name, FPGAPart: b.FPGAPart, SrcPath: srcAbs, IPPath: ipAbs,
		TopModule: b.TopModule, DeviceID: fmt.Sprintf("%04X", ctx.Device.DeviceID),
		VendorID:       fmt.Sprintf("%04X", ctx.Device.VendorID),
		RevisionID:     fmt.Sprintf("%02X", ctx.Device.RevisionID),
		SubsysVendorID: fmt.Sprintf("%04X", ctx.Device.SubsysVendorID),
		SubsysDeviceID: fmt.Sprintf("%04X", ctx.Device.SubsysDeviceID),
		ClassCodeBase:  fmt.Sprintf("%02X", (ctx.Device.ClassCode>>16)&0xFF),
		ClassCodeSub:   fmt.Sprintf("%02X", (ctx.Device.ClassCode>>8)&0xFF),
		ClassCodeIntf:  fmt.Sprintf("%02X", ctx.Device.ClassCode&0xFF),
		LinkSpeed:      linkSpeedToTCL(linkSpeed), LinkWidth: linkWidthToTCL(linkWidth),
		TrgtLinkSpeed: linkSpeedToTrgt(linkSpeed), Bar0Enabled: bar0.Enabled,
		Bar0Size: bar0.Size, Bar0Scale: bar0.Scale, Bar064bit: bar0.Is64bit,
		Bar0ByteSize: bar0Size, StockBar: stockBar, ImportVFiles: b.ImportVFiles,
		BARs: bars, DSNEnabled: ids.HasDSN,
		MSICapVectorsStr: msiVectorsToTCL(extractMSIVectors(ctx)),
	}
	if cfg != nil && cfg.MSIXConfig != nil {
		data.MSIXEnabled = true
		data.MSIXTableSize = cfg.MSIXConfig.NumVectors - 1
		data.MSIXTableBIR = barBIRToTCL(cfg.MSIXConfig.TableBIR, configBARIs64(cfg, cfg.MSIXConfig.TableBIR))
		data.MSIXTableOffset = fmt.Sprintf("%08X", cfg.MSIXConfig.TableOffset)
		data.MSIXPBABIR = barBIRToTCL(cfg.MSIXConfig.PBABIR, configBARIs64(cfg, cfg.MSIXConfig.PBABIR))
		data.MSIXPBAOffset = fmt.Sprintf("%08X", cfg.MSIXConfig.PBAOffset)
	} else if ctx.MSIXData != nil && ctx.MSIXData.TableSize > 0 {
		data.MSIXEnabled = true
		data.MSIXTableSize = ctx.MSIXData.TableSize - 1
		data.MSIXTableBIR = barBIRToTCL(ctx.MSIXData.TableBIR, bar0.Is64bit && ctx.MSIXData.TableBIR == 0)
		data.MSIXTableOffset = fmt.Sprintf("%08X", ctx.MSIXData.TableOffset)
		data.MSIXPBABIR = barBIRToTCL(ctx.MSIXData.PBABIR, bar0.Is64bit && ctx.MSIXData.PBABIR == 0)
		data.MSIXPBAOffset = fmt.Sprintf("%08X", ctx.MSIXData.PBAOffset)
	}
	var buf bytes.Buffer
	if err := projectTCLTmpl.ExecuteTemplate(&buf, "project.tcl.tmpl", data); err != nil {
		panic(fmt.Sprintf("project TCL template error: %v", err))
	}
	return buf.String()
}

func tclPath(path string) string {
	return strings.ReplaceAll(path, `\`, "/")
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
