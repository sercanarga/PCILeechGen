// Package svgen generates SystemVerilog source files for PCILeech FPGA firmware.
package svgen

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/donor/behavior"
	"github.com/sercanarga/pcileechgen/internal/firmware"
	"github.com/sercanarga/pcileechgen/internal/firmware/barmodel"
	"github.com/sercanarga/pcileechgen/internal/firmware/nvme"
)

// MSIConfig describes the MSI capability programmed into config space.
type MSIConfig struct {
	Enabled bool   // MSI enabled in Message Control
	AddrLo  uint32 // MSI Message Address (lower 32 bits)
	Data    uint16 // MSI Message Data
}

type BARSlot struct {
	BIR        int
	Model      *barmodel.BARModel
	ModuleName string
	Primary    bool
}

// SVGeneratorConfig is the input data for all SV template renders.
type SVGeneratorConfig struct {
	DeviceIDs                   firmware.DeviceIDs
	DonorCapabilities           DonorCapabilities // donor capability summary for donor-emulation visibility
	BARModels                   []*barmodel.BARModel
	DonorBARTopology            bool
	BARModel                    *barmodel.BARModel // nil = generic fallback (uses BRAM-based zerowrite4k)
	BARModuleName               string
	ClassCode                   uint32
	LatencyConfig               *LatencyConfig     // TLP response timing (nil = no latency emulator)
	HasMSIX                     bool               // generate MSI-X interrupt controller logic
	BuildEntropy                uint32             // seed for PRNG uniqueness per build
	PRNGSeeds                   [4]uint32          // computed PRNG seeds for latency emulator
	DeviceClass                 string             // "nvme", "xhci", "audio", "ethernet", or ""
	MSIXConfig                  *MSIXConfig        // MSI-X table replication (nil = no MSI-X table)
	MSIConfig                   *MSIConfig         // MSI capability info (nil = no MSI cap or disabled)
	NVMeIdentify                *nvme.IdentifyData // NVMe Identify Controller/Namespace data (nil = no responder)
	NVMeSMART                   *nvme.SMART        // donor-plausible SMART/Health wear seeds (nil = zero wear)
	NVMeDoorbellStride          uint32             // CAP.DSTRD - doorbell stride (0 = 4B, default)
	NVMeDiskWords               int                // NVMe disk-cache depth (words), board-scaled
	NVMeAdvertisedLBAs          uint64             // actual NSZE from donor (0 = use default 2000409264)
	Bar0Size                    int
	ReadCompletionBoundaryBytes int
	MaxPayloadBytes             int
	BehaviorRules               *behavior.RuleSet
	CompiledBehavior            *CompiledBehavior
}

func (c *SVGeneratorConfig) ResolvedReadCompletionBoundaryBytes() int {
	if c != nil && (c.ReadCompletionBoundaryBytes == 64 || c.ReadCompletionBoundaryBytes == 128) {
		return c.ReadCompletionBoundaryBytes
	}
	return 64
}

func (c *SVGeneratorConfig) ResolvedMaxPayloadBytes() int {
	if c != nil && c.MaxPayloadBytes >= 128 && c.MaxPayloadBytes <= 4096 &&
		c.MaxPayloadBytes&(c.MaxPayloadBytes-1) == 0 {
		return c.MaxPayloadBytes
	}
	return 128
}

func (c *SVGeneratorConfig) HasClassEndpoint() bool {
	if len(c.BARModels) == 0 {
		return c.BARModel != nil
	}
	return c.BARModel != nil && c.BARModel.ClassSpecific
}

func (c *SVGeneratorConfig) PrimaryBIR() int {
	if c.BARModel != nil {
		return c.BARModel.BIR
	}
	return 0
}

func (c *SVGeneratorConfig) BARSlots() []BARSlot {
	slots := make([]BARSlot, 6)
	for bir := range slots {
		slots[bir] = BARSlot{BIR: bir}
	}
	if c.BARModel == nil && !c.DonorBARTopology {
		slots[0].Primary = true
	}
	models := c.BARModels
	if len(models) == 0 && c.BARModel != nil {
		models = []*barmodel.BARModel{c.BARModel}
	}
	for _, model := range models {
		if model == nil || model.BIR < 0 || model.BIR >= len(slots) {
			continue
		}
		name := fmt.Sprintf("pcileech_bar_impl_device_bar%d", model.BIR)
		primary := model == c.BARModel
		if primary {
			name = "pcileech_bar_impl_device"
		}
		slots[model.BIR] = BARSlot{BIR: model.BIR, Model: model, ModuleName: name, Primary: primary}
	}
	return slots
}

func (c *SVGeneratorConfig) BARImplementationModuleName() string {
	if c.BARModuleName != "" {
		return c.BARModuleName
	}
	return "pcileech_bar_impl_device"
}

// NVMeDiskWordsForBRAM36 returns the NVMe disk-cache depth in 32-bit words for
// an FPGA with bram36 RAMB36 blocks, or 0 if the board cannot fit a useful
// cache. Pinned per part to stay within BRAM budget.
func NVMeDiskWordsForBRAM36(bram36 int) int {
	switch {
	case bram36 >= 365: // 200T -> 32 KiB
		return 32768
	case bram36 >= 105: // 75T/100T -> 8 KiB
		return 8192
	default: // 35T/50T: too small for a disk cache
		return 0
	}
}

// DonorCapabilities summarizes parsed capabilities from donor config space.
// Values are best-effort snapshots used by generated SV for optional emulation
// behavior and debugging visibility.
type DonorCapabilities struct {
	HasPMCap         bool
	HasMSICap        bool
	HasMSIXCap       bool
	HasPCIeCap       bool
	PMCapOffset      uint16
	MSICapOffset     uint16
	MSIXCapOffset    uint16
	PCIeCapOffset    uint16
	PMESupportMask   uint8
	PMEDefault       bool
	MSIDisable64Bit  bool
	MSIMultipleMsg   uint8
	PCIELinkSpeed    uint8
	PCIELinkWidth    uint8
	PCIeASPMCap      uint8
	PCIeASPMEnable   uint8
	HasLTRCap        bool
	HasL1PMSubstates bool
	HasAERCap        bool
	HasDSNCap        bool
	AERCapOffset     uint16
	LTRCapOffset     uint16
	L1PMCapOffset    uint16
	DSNCapOffset     uint16
}

// NVMeSQ0DoorbellOffset returns the byte offset of the SQ0 tail doorbell.
// Doorbell base uses board.DefaultBRAMSize (0x1000 classic) so it stays correct for
// variable/large BAR0 (16k+) + post-doorbell MSIX placement on boards like CaptainDMA_75T.
func (c *SVGeneratorConfig) NVMeSQ0DoorbellOffset() uint32 {
	stride := uint32(4) << c.NVMeDoorbellStride
	dbBase := uint32(board.DefaultBRAMSize)
	return dbBase + 0*stride
}

// NVMeCQ0DoorbellOffset returns the byte offset of the CQ0 head doorbell.
// See NVMeSQ0DoorbellOffset for variable BAR rationale.
func (c *SVGeneratorConfig) NVMeCQ0DoorbellOffset() uint32 {
	stride := uint32(4) << c.NVMeDoorbellStride
	dbBase := uint32(board.DefaultBRAMSize)
	return dbBase + 1*stride
}

func (c *SVGeneratorConfig) NVMeSQ1DoorbellOffset() uint32 {
	stride := uint32(4) << c.NVMeDoorbellStride
	dbBase := uint32(board.DefaultBRAMSize)
	return dbBase + 2*stride
}

func (c *SVGeneratorConfig) NVMeCQ1DoorbellOffset() uint32 {
	stride := uint32(4) << c.NVMeDoorbellStride
	dbBase := uint32(board.DefaultBRAMSize)
	return dbBase + 3*stride
}

func renderTemplate(name string, data any) (string, error) {
	tmplStr := mustReadTemplate(name + ".sv.tmpl")
	tmpl, err := template.New(name).Funcs(svFuncMap()).Parse(tmplStr)
	if err != nil {
		return "", fmt.Errorf("parsing %s template: %w", name, err)
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("executing %s template: %w", name, err)
	}
	return buf.String(), nil
}

// renderTemplateDelim is renderTemplate with custom delimiters.
func renderTemplateDelim(name, leftDelim, rightDelim string, data any) (string, error) {
	tmplStr := mustReadTemplate(name + ".sv.tmpl")
	tmpl, err := template.New(name).Delims(leftDelim, rightDelim).Funcs(svFuncMap()).Parse(tmplStr)
	if err != nil {
		return "", fmt.Errorf("parsing %s template: %w", name, err)
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("executing %s template: %w", name, err)
	}
	return buf.String(), nil
}

func GenerateBarImplDeviceSV(cfg *SVGeneratorConfig) (string, error) {
	if len(cfg.BARModels) == 0 {
		prepared, err := prepareBehaviorConfig(cfg)
		if err != nil {
			return "", err
		}
		return renderTemplate("bar_impl_device", prepared)
	}

	var out bytes.Buffer
	for i, model := range cfg.BARModels {
		if model == nil {
			continue
		}
		endpoint := *cfg
		endpoint.BARModels = nil
		endpoint.BARModel = model
		if model == cfg.BARModel {
			endpoint.BARModuleName = "pcileech_bar_impl_device"
		} else {
			endpoint.BARModuleName = fmt.Sprintf("pcileech_bar_impl_device_bar%d", model.BIR)
		}
		if !model.ClassSpecific {
			endpoint.DeviceClass = ""
		}
		prepared, err := prepareBehaviorConfig(&endpoint)
		if err != nil {
			return "", fmt.Errorf("preparing BAR%d endpoint: %w", model.BIR, err)
		}
		rendered, err := renderTemplate("bar_impl_device", prepared)
		if err != nil {
			return "", fmt.Errorf("rendering BAR%d endpoint: %w", model.BIR, err)
		}
		if i > 0 {
			out.WriteString("\n")
		}
		out.WriteString(rendered)
	}
	return out.String(), nil
}

func GenerateBarControllerSV(cfg *SVGeneratorConfig) (string, error) {
	return renderTemplate("bar_controller", cfg)
}

func GenerateTransactionNormalizerSV(cfg *SVGeneratorConfig) (string, error) {
	return renderTemplate("transaction_normalizer", cfg)
}

func GenerateBarReadEngineSV(cfg *SVGeneratorConfig) (string, error) {
	return renderTemplate("bar_read_engine", cfg)
}

func GenerateURCompleterSV(cfg *SVGeneratorConfig) (string, error) {
	return renderTemplate("transaction_ur_completer", cfg)
}

func GenerateBarRspArbiterSV(cfg *SVGeneratorConfig) (string, error) {
	return renderTemplate("bar_rsp_arbiter", cfg)
}

func GenerateDeviceConfigSV(cfg *SVGeneratorConfig) (string, error) {
	return renderTemplate("device_config", cfg)
}

func GenerateMSIXTableSV(cfg *SVGeneratorConfig) (string, error) {
	return renderTemplate("msix_table", cfg)
}

func GenerateLatencyEmulatorSV(cfg *SVGeneratorConfig) (string, error) {
	return renderTemplate("latency_emulator", cfg)
}

func GenerateNVMeResponderSV(cfg *SVGeneratorConfig) (string, error) {
	return renderTemplate("nvme_admin_responder", cfg)
}

// GenerateNVMeDMABridgeSV renders the NVMe DMA TLP bridge module.
func GenerateNVMeDMABridgeSV(cfg *SVGeneratorConfig) (string, error) {
	return renderTemplate("nvme_dma_bridge", cfg)
}

// GenerateNVMeBRAMDiskSV renders the disk cache. Uses [[ ]] delimiters because
// the body's Verilog {{N{...}}} replication clashes with {{ }}.
func GenerateNVMeBRAMDiskSV(cfg *SVGeneratorConfig) (string, error) {
	return renderTemplateDelim("nvme_bram_disk", "[[", "]]", cfg)
}

// GenerateHDARIRBDMASV renders the HDA RIRB DMA bridge module.
func GenerateHDARIRBDMASV(cfg *SVGeneratorConfig) (string, error) {
	return renderTemplate("hda_rirb_dma", cfg)
}

// GenerateHDAMSISV renders the HDA MSI interrupt TLP generator module.
func GenerateHDAMSISV(cfg *SVGeneratorConfig) (string, error) {
	return renderTemplate("hda_msi", cfg)
}

// GenerateBarImplMSISV renders the MSI doorbell BAR implementation.
func GenerateBarImplMSISV(cfg *SVGeneratorConfig) (string, error) {
	return renderTemplate("bar_impl_msi", cfg)
}

// svFuncMap provides hex formatting and arithmetic helpers for templates.
func svFuncMap() template.FuncMap {
	return template.FuncMap{
		"hex08":         func(v uint32) string { return fmt.Sprintf("%08X", v) },
		"hex16":         func(v uint64) string { return fmt.Sprintf("%016X", v) },
		"hex04":         func(v uint16) string { return fmt.Sprintf("%04X", v) },
		"hex02":         func(v uint8) string { return fmt.Sprintf("%02X", v) },
		"sub":           func(a, b int) int { return a - b },
		"mul":           func(a, b int) int { return a * b },
		"alignedOffset": func(off uint32) uint32 { return (off / 4) * 4 },
		"classBase":     func(cc uint32) uint8 { return uint8((cc >> 16) & 0xFF) },
		"classSub":      func(cc uint32) uint8 { return uint8((cc >> 8) & 0xFF) },
		"classProgIF":   func(cc uint32) uint8 { return uint8(cc & 0xFF) },
		"rwMaskBytes": func(mask uint32) [4]uint8 {
			return [4]uint8{
				uint8(mask & 0xFF),
				uint8((mask >> 8) & 0xFF),
				uint8((mask >> 16) & 0xFF),
				uint8((mask >> 24) & 0xFF),
			}
		},
		"w1cMaskBytes": func(mask uint32) [4]uint8 {
			return [4]uint8{
				uint8(mask & 0xFF),
				uint8((mask >> 8) & 0xFF),
				uint8((mask >> 16) & 0xFF),
				uint8((mask >> 24) & 0xFF),
			}
		},
		"cdfVal": func(cdf [16]uint8, i int) uint8 {
			if i < 0 || i > 15 {
				return 0
			}
			return cdf[i]
		},
	}
}
