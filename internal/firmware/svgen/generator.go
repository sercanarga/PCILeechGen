// Package svgen generates SystemVerilog source files for PCILeech FPGA firmware.
package svgen

import (
	"bytes"
	"fmt"
	"text/template"

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
}

// NVMeSQ0DoorbellOffset returns the byte offset of the SQ0 tail doorbell.
func (c *SVGeneratorConfig) NVMeSQ0DoorbellOffset() uint32 {
	stride := uint32(4) << c.NVMeDoorbellStride
	return 0x1000 + 0*stride // SQ0 tail
}

// NVMeCQ0DoorbellOffset returns the byte offset of the CQ0 head doorbell.
func (c *SVGeneratorConfig) NVMeCQ0DoorbellOffset() uint32 {
	stride := uint32(4) << c.NVMeDoorbellStride
	return 0x1000 + 1*stride // CQ0 head
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

// GenerateHDARIRBDMASV renders the HDA RIRB DMA bridge module.
func GenerateHDARIRBDMASV(cfg *SVGeneratorConfig) (string, error) {
	return renderTemplate("hda_rirb_dma", cfg)
}

// GenerateHDAMSISV renders the HDA MSI interrupt TLP generator module.
func GenerateHDAMSISV(cfg *SVGeneratorConfig) (string, error) {
	return renderTemplate("hda_msi", cfg)
}

// svFuncMap provides hex formatting and arithmetic helpers for templates.
func svFuncMap() template.FuncMap {
	return template.FuncMap{
		"hex08":         func(v uint32) string { return fmt.Sprintf("%08X", v) },
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
		"cdfVal": func(cdf [16]uint8, i int) uint8 {
			if i < 0 || i > 15 {
				return 0
			}
			return cdf[i]
		},
	}
}
