// Package svgen generates SystemVerilog source files for PCILeech FPGA firmware.
package svgen

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/sercanarga/pcileechgen/internal/firmware"
	"github.com/sercanarga/pcileechgen/internal/firmware/barmodel"
)

// SVGeneratorConfig is the input data for all SV template renders.
type SVGeneratorConfig struct {
	DeviceIDs     firmware.DeviceIDs
	BARModel      *barmodel.BARModel // nil = generic fallback (uses BRAM-based zerowrite4k)
	ClassCode     uint32
	LatencyConfig *LatencyConfig // TLP response timing (nil = no latency emulator)
	HasMSIX       bool           // generate MSI-X interrupt controller logic
	BuildEntropy  uint32         // seed for PRNG uniqueness per build
	PRNGSeeds     [4]uint32      // computed PRNG seeds for latency emulator
	IsNVMe        bool           // enable NVMe CC→CSTS state machine
	IsXHCI        bool           // enable xHCI state machine
}

func renderTemplate(name, tmplStr string, data any) (string, error) {
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
	return renderTemplate("bar_impl_device", barImplDeviceTmpl, cfg)
}

func GenerateBarControllerSV(cfg *SVGeneratorConfig) (string, error) {
	return renderTemplate("bar_controller", barControllerTmpl, cfg)
}

func GenerateDeviceConfigSV(cfg *SVGeneratorConfig) (string, error) {
	return renderTemplate("device_config", deviceConfigTmpl, cfg)
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
	}
}
