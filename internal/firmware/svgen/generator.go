// Package svgen generates SystemVerilog source files for PCILeech FPGA firmware.
package svgen

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/sercanarga/pcileechgen/internal/board"
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

// SVGeneratorConfig is the input data for all SV template renders.
type SVGeneratorConfig struct {
	DeviceIDs          firmware.DeviceIDs
	DonorCapabilities  DonorCapabilities  // donor capability summary for donor-emulation visibility
	BARModel           *barmodel.BARModel // nil = generic fallback (uses BRAM-based zerowrite4k)
	ClassCode          uint32
	LatencyConfig      *LatencyConfig     // TLP response timing (nil = no latency emulator)
	HasMSIX            bool               // generate MSI-X interrupt controller logic
	BuildEntropy       uint32             // seed for PRNG uniqueness per build
	PRNGSeeds          [4]uint32          // computed PRNG seeds for latency emulator
	DeviceClass        string             // "nvme", "xhci", "audio", "ethernet", or ""
	MSIXConfig         *MSIXConfig        // MSI-X table replication (nil = no MSI-X table)
	MSIConfig          *MSIConfig         // MSI capability info (nil = no MSI cap or disabled)
	NVMeIdentify       *nvme.IdentifyData // NVMe Identify Controller/Namespace data (nil = no responder)
	NVMeDoorbellStride uint32             // CAP.DSTRD - doorbell stride (0 = 4B, default)
	Bar0Size           int
	BARInitHexFile     string // generic BRAM fallback: $readmemh seed file ("" = zero-init only)
	OptionROMHexFile   string // BAR6 expansion ROM responder: $readmemh seed ("" = no ROM served)
	OptionROMSize      int    // expansion ROM aperture size in bytes (power of 2)
	ILAInstanceSV      string
	// ExtraBARPresent flags donor BAR3-6 presence: index 0=BAR3 ... 3=BAR6.
	// true = donor's real hardware has a populated (nonzero-size) BAR there,
	// so bar_controller.sv.tmpl presents a real (loopaddr) aperture instead
	// of pcileech_bar_impl_none. Zero value (all false) preserves the old
	// always-none behavior.
	ExtraBARPresent [4]bool
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

func (c *SVGeneratorConfig) NVMeDiagBaseOffset() uint32 {
	stride := uint32(4) << c.NVMeDoorbellStride
	dbBase := uint32(board.DefaultBRAMSize)
	return dbBase + 4*stride
}

func (c *SVGeneratorConfig) NVMeDiagLastCommandOffset() uint32 {
	return c.NVMeDiagBaseOffset() + 0x4
}

func (c *SVGeneratorConfig) NVMeDiagLastNSIDOffset() uint32 {
	return c.NVMeDiagBaseOffset() + 0x8
}

func (c *SVGeneratorConfig) NVMeDiagLastCDW10Offset() uint32 {
	return c.NVMeDiagBaseOffset() + 0xC
}

func (c *SVGeneratorConfig) NVMeDiagQueueStateOffset() uint32 {
	return c.NVMeDiagBaseOffset() + 0x10
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

func GenerateBarImplDeviceSV(cfg *SVGeneratorConfig) (string, error) {
	return renderTemplate("bar_impl_device", cfg)
}

func GenerateBarControllerSV(cfg *SVGeneratorConfig) (string, error) {
	return renderTemplate("bar_controller", cfg)
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

// GenerateXHCIRingEngineSV renders the xHCI Command/Event ring engine and its
// DMA bridge (both modules live in the same template file).
func GenerateXHCIRingEngineSV(cfg *SVGeneratorConfig) (string, error) {
	return renderTemplate("xhci_ring_engine", cfg)
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

// GenerateOptionROMSV renders the BAR6 expansion ROM responder.
func GenerateOptionROMSV(cfg *SVGeneratorConfig) (string, error) {
	return renderTemplate("bar_impl_optrom", cfg)
}

// svFuncMap provides hex formatting and arithmetic helpers for templates.
func svFuncMap() template.FuncMap {
	return template.FuncMap{
		"hex08":         func(v uint32) string { return fmt.Sprintf("%08X", v) },
		"hex04":         func(v uint16) string { return fmt.Sprintf("%04X", v) },
		"hex02":         func(v uint8) string { return fmt.Sprintf("%02X", v) },
		"sub":           func(a, b int) int { return a - b },
		"add":           func(a, b int) int { return a + b },
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
