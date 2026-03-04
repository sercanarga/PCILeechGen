package fallback

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	yamlContent := `
defaults:
  bar0_size: 4096
  power_management: true
  msi_capable: true
device_classes:
  "0200":
    description: "Ethernet controller"
    bar0_size: 131072
    link_speed: 3
    link_width: 1
    bar0_defaults:
      "0x00000000": 0
`
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "fallbacks.yaml")
	if err := os.WriteFile(path, []byte(yamlContent), 0644); err != nil {
		t.Fatal(err)
	}

	cfg, err := LoadConfig(path)
	if err != nil {
		t.Fatalf("LoadConfig failed: %v", err)
	}

	if cfg.Defaults.BAR0Size != 4096 {
		t.Errorf("expected default bar0_size 4096, got %d", cfg.Defaults.BAR0Size)
	}
	if dc, ok := cfg.DeviceClasses["0200"]; !ok {
		t.Error("expected device class 0200")
	} else if dc.Description != "Ethernet controller" {
		t.Errorf("expected Ethernet controller, got %s", dc.Description)
	}
}

func TestClassKey(t *testing.T) {
	key := ClassKey(0x02, 0x00)
	if key != "0200" {
		t.Errorf("expected 0200, got %s", key)
	}
	key = ClassKey(0x01, 0x08)
	if key != "0108" {
		t.Errorf("expected 0108, got %s", key)
	}
}

func TestApply_FillsZeroedBAR(t *testing.T) {
	cfg := &Config{
		DeviceClasses: map[string]*DeviceClass{
			"0108": {
				Description: "NVMe",
				BAR0Defaults: map[string]uint32{
					"0x00000008": 0x00010301,
				},
			},
		},
	}

	barContents := map[int][]byte{
		0: make([]byte, 4096),
	}

	// class 01:08:02 → NVMe
	results := Apply(cfg, 0x010802, barContents)
	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}
	if results[0].Field != "BAR0[0x0008]" {
		t.Errorf("unexpected field: %s", results[0].Field)
	}

	// verify the value was written (little-endian)
	got := uint32(barContents[0][8]) | uint32(barContents[0][9])<<8 |
		uint32(barContents[0][10])<<16 | uint32(barContents[0][11])<<24
	if got != 0x00010301 {
		t.Errorf("expected 0x00010301 at offset 8, got 0x%08X", got)
	}
}

func TestApply_SkipsNonZero(t *testing.T) {
	cfg := &Config{
		DeviceClasses: map[string]*DeviceClass{
			"0108": {
				Description: "NVMe",
				BAR0Defaults: map[string]uint32{
					"0x00000000": 0xDEADBEEF,
				},
			},
		},
	}

	barContents := map[int][]byte{
		0: make([]byte, 4096),
	}
	// pre-fill offset 0 with a non-zero value
	barContents[0][0] = 0xFF

	results := Apply(cfg, 0x010802, barContents)
	if len(results) != 0 {
		t.Errorf("expected 0 results (non-zero register), got %d", len(results))
	}
}

func TestApply_UnknownClass(t *testing.T) {
	cfg := &Config{
		DeviceClasses: map[string]*DeviceClass{},
	}
	results := Apply(cfg, 0xFF0000, nil)
	if len(results) != 0 {
		t.Errorf("expected 0 results for unknown class, got %d", len(results))
	}
}
