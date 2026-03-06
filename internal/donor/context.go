// Package donor collects PCI device info for firmware generation.
package donor

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/sercanarga/pcileechgen/internal/pci"
)

// MSIXData holds donor MSI-X table content.
type MSIXData struct {
	TableSize   int             `json:"table_size"`
	TableBIR    int             `json:"table_bir"`
	TableOffset uint32          `json:"table_offset"`
	PBABIR      int             `json:"pba_bir"`
	PBAOffset   uint32          `json:"pba_offset"`
	Entries     []pci.MSIXEntry `json:"entries"`
}

// DeviceContext is the full snapshot of a donor device.
type DeviceContext struct {
	CollectedAt time.Time `json:"collected_at"`
	ToolVersion string    `json:"tool_version"`
	Hostname    string    `json:"hostname"`

	Device          pci.PCIDevice       `json:"device"`
	ConfigSpace     *pci.ConfigSpace    `json:"config_space"`
	BARs            []pci.BAR           `json:"bars"`
	BARContents     map[int][]byte      `json:"-"` // BAR memory contents, keyed by BAR index
	BARProfiles     map[int]*BARProfile `json:"-"` // probing results, keyed by BAR index
	Capabilities    []pci.Capability    `json:"capabilities"`
	ExtCapabilities []pci.ExtCapability `json:"ext_capabilities,omitempty"`
	MSIXData        *MSIXData           `json:"msix_data,omitempty"`
}

// JSON wire format — config space as hex words, BARs as base64.
type deviceContextJSON struct {
	CollectedAt     time.Time              `json:"collected_at"`
	ToolVersion     string                 `json:"tool_version"`
	Hostname        string                 `json:"hostname"`
	Device          pci.PCIDevice          `json:"device"`
	ConfigSpaceHex  []string               `json:"config_space_hex"`
	ConfigSpaceSize int                    `json:"config_space_size"`
	BARs            []pci.BAR              `json:"bars"`
	BARContents     map[string]string      `json:"bar_contents,omitempty"` // key: BAR index, value: base64
	BARProfiles     map[string]*BARProfile `json:"bar_profiles,omitempty"`
	Capabilities    []pci.Capability       `json:"capabilities"`
	ExtCapabilities []pci.ExtCapability    `json:"ext_capabilities,omitempty"`
	MSIXData        *MSIXData              `json:"msix_data,omitempty"`
}

func (dc *DeviceContext) MarshalJSON() ([]byte, error) {
	j := deviceContextJSON{
		CollectedAt:     dc.CollectedAt,
		ToolVersion:     dc.ToolVersion,
		Hostname:        dc.Hostname,
		Device:          dc.Device,
		BARs:            dc.BARs,
		Capabilities:    dc.Capabilities,
		ExtCapabilities: dc.ExtCapabilities,
		MSIXData:        dc.MSIXData,
	}

	if dc.ConfigSpace != nil {
		j.ConfigSpaceSize = dc.ConfigSpace.Size
		for i := 0; i < dc.ConfigSpace.Size; i += 4 {
			word := dc.ConfigSpace.ReadU32(i)
			j.ConfigSpaceHex = append(j.ConfigSpaceHex, fmt.Sprintf("%08x", word))
		}
	}

	if len(dc.BARContents) > 0 {
		j.BARContents = make(map[string]string)
		for idx, data := range dc.BARContents {
			j.BARContents[strconv.Itoa(idx)] = base64.StdEncoding.EncodeToString(data)
		}
	}

	if len(dc.BARProfiles) > 0 {
		j.BARProfiles = make(map[string]*BARProfile)
		for idx, profile := range dc.BARProfiles {
			j.BARProfiles[strconv.Itoa(idx)] = profile
		}
	}

	return json.Marshal(j)
}

// ToJSON serializes the DeviceContext to indented JSON.
func (dc *DeviceContext) ToJSON() ([]byte, error) {
	return json.MarshalIndent(dc, "", "  ")
}

// FromJSON deserializes a DeviceContext from JSON.
func FromJSON(data []byte) (*DeviceContext, error) {
	var j deviceContextJSON
	if err := json.Unmarshal(data, &j); err != nil {
		return nil, fmt.Errorf("failed to parse device context JSON: %w", err)
	}

	dc := &DeviceContext{
		CollectedAt:     j.CollectedAt,
		ToolVersion:     j.ToolVersion,
		Hostname:        j.Hostname,
		Device:          j.Device,
		BARs:            j.BARs,
		Capabilities:    j.Capabilities,
		ExtCapabilities: j.ExtCapabilities,
		MSIXData:        j.MSIXData,
	}

	// Reconstruct config space from hex words
	if len(j.ConfigSpaceHex) > 0 {
		dc.ConfigSpace = pci.NewConfigSpace()
		dc.ConfigSpace.Size = j.ConfigSpaceSize
		for i, hexWord := range j.ConfigSpaceHex {
			var word uint32
			if _, err := fmt.Sscanf(hexWord, "%x", &word); err != nil {
				return nil, fmt.Errorf("invalid config space hex word %d (%q): %w", i, hexWord, err)
			}
			dc.ConfigSpace.WriteU32(i*4, word)
		}
	}

	// Reconstruct BAR contents from base64
	if len(j.BARContents) > 0 {
		dc.BARContents = make(map[int][]byte)
		for idxStr, b64 := range j.BARContents {
			idx, err := strconv.Atoi(idxStr)
			if err != nil {
				return nil, fmt.Errorf("invalid BAR index %q: %w", idxStr, err)
			}
			data, err := base64.StdEncoding.DecodeString(b64)
			if err != nil {
				return nil, fmt.Errorf("failed to decode BAR%d content: %w", idx, err)
			}
			dc.BARContents[idx] = data
		}
	}

	// Reconstruct BAR profiles
	if len(j.BARProfiles) > 0 {
		dc.BARProfiles = make(map[int]*BARProfile)
		for idxStr, profile := range j.BARProfiles {
			idx, err := strconv.Atoi(idxStr)
			if err != nil {
				return nil, fmt.Errorf("invalid BAR profile index %q: %w", idxStr, err)
			}
			dc.BARProfiles[idx] = profile
		}
	}

	return dc, nil
}

// SaveContext dumps a DeviceContext to JSON on disk.
func SaveContext(ctx *DeviceContext, path string) error {
	data, err := ctx.ToJSON()
	if err != nil {
		return fmt.Errorf("failed to marshal device context: %w", err)
	}
	return os.WriteFile(path, data, 0644)
}

// LoadContext restores a DeviceContext from a JSON file.
func LoadContext(path string) (*DeviceContext, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read device context file: %w", err)
	}
	return FromJSON(data)
}
