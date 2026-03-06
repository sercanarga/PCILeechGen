// Package fallback fills in missing device data with class-specific defaults.
package fallback

import (
	"fmt"
	"log/slog"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Defaults      Defaults                `yaml:"defaults"`
	DeviceClasses map[string]*DeviceClass `yaml:"device_classes"`
}

type Defaults struct {
	BAR0Size        uint64 `yaml:"bar0_size"`
	PowerManagement bool   `yaml:"power_management"`
	MSICapable      bool   `yaml:"msi_capable"`
}

type DeviceClass struct {
	Description  string            `yaml:"description"`
	BAR0Size     uint64            `yaml:"bar0_size"`
	LinkSpeed    uint8             `yaml:"link_speed"`
	LinkWidth    uint8             `yaml:"link_width"`
	Capabilities []string          `yaml:"capabilities"`
	BAR0Defaults map[string]uint32 `yaml:"bar0_defaults"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read fallback config: %w", err)
	}
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse fallback config: %w", err)
	}
	return &cfg, nil
}

// ClassKey turns base+sub class into a lookup key, e.g. 0x02,0x00 → "0200".
func ClassKey(baseClass, subClass uint8) string {
	return fmt.Sprintf("%02X%02X", baseClass, subClass)
}

type ApplyResult struct {
	Field    string
	OldValue string
	NewValue string
}

// Apply fills zeroed BAR registers with class-specific defaults.
func Apply(cfg *Config, classCode uint32, barContents map[int][]byte) []ApplyResult {
	if cfg == nil {
		return nil
	}

	baseClass := uint8((classCode >> 16) & 0xFF)
	subClass := uint8((classCode >> 8) & 0xFF)
	key := ClassKey(baseClass, subClass)

	dc, ok := cfg.DeviceClasses[key]
	if !ok {
		slog.Info("no class-specific fallbacks", "class", key)
		return nil
	}

	var results []ApplyResult
	slog.Info("applying fallback defaults", "description", dc.Description, "class", key)

	bar0, hasBAR0 := barContents[0]
	if hasBAR0 && len(bar0) >= 4 && dc.BAR0Defaults != nil {
		for offsetStr, defaultVal := range dc.BAR0Defaults {
			var offset uint32
			if _, err := fmt.Sscanf(offsetStr, "0x%x", &offset); err != nil {
				continue
			}
			if int(offset)+4 > len(bar0) {
				continue
			}
			current := uint32(bar0[offset]) | uint32(bar0[offset+1])<<8 |
				uint32(bar0[offset+2])<<16 | uint32(bar0[offset+3])<<24
			if current == 0 && defaultVal != 0 {
				bar0[offset] = byte(defaultVal)
				bar0[offset+1] = byte(defaultVal >> 8)
				bar0[offset+2] = byte(defaultVal >> 16)
				bar0[offset+3] = byte(defaultVal >> 24)
				results = append(results, ApplyResult{
					Field:    fmt.Sprintf("BAR0[0x%04X]", offset),
					OldValue: "0x00000000",
					NewValue: fmt.Sprintf("0x%08X", defaultVal),
				})
			}
		}
	}

	return results
}
