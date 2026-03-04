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

// GenerateBarImplDeviceSV renders pcileech_bar_impl_device.sv.
// Falls back to a BRAM wrapper when cfg.BARModel is nil.
func GenerateBarImplDeviceSV(cfg *SVGeneratorConfig) (string, error) {
	tmpl, err := template.New("bar_impl_device").Funcs(svFuncMap()).Parse(barImplDeviceTmpl)
	if err != nil {
		return "", fmt.Errorf("parsing bar_impl_device template: %w", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, cfg); err != nil {
		return "", fmt.Errorf("executing bar_impl_device template: %w", err)
	}
	return buf.String(), nil
}

// GenerateBarControllerSV renders pcileech_tlps128_bar_controller.sv.
func GenerateBarControllerSV(cfg *SVGeneratorConfig) (string, error) {
	tmpl, err := template.New("bar_controller").Funcs(svFuncMap()).Parse(barControllerTmpl)
	if err != nil {
		return "", fmt.Errorf("parsing bar_controller template: %w", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, cfg); err != nil {
		return "", fmt.Errorf("executing bar_controller template: %w", err)
	}
	return buf.String(), nil
}

// GenerateDeviceConfigSV renders device_config.sv (identity + feature flags).
func GenerateDeviceConfigSV(cfg *SVGeneratorConfig) (string, error) {
	tmpl, err := template.New("device_config").Funcs(svFuncMap()).Parse(deviceConfigTmpl)
	if err != nil {
		return "", fmt.Errorf("parsing device_config template: %w", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, cfg); err != nil {
		return "", fmt.Errorf("executing device_config template: %w", err)
	}
	return buf.String(), nil
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
