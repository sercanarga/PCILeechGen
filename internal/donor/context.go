// Package donor handles reading PCI device information from a donor device
// using Linux sysfs and VFIO.
package donor

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/sercanarga/pcileechgen/internal/pci"
)

// DeviceContext holds all collected information about a donor PCI device.
type DeviceContext struct {
	CollectedAt time.Time `json:"collected_at"`
	ToolVersion string    `json:"tool_version"`
	Hostname    string    `json:"hostname"`

	Device          pci.PCIDevice       `json:"device"`
	ConfigSpace     *pci.ConfigSpace    `json:"config_space"`
	BARs            []pci.BAR           `json:"bars"`
	BARContents     map[int][]byte      `json:"-"` // BAR memory contents, keyed by BAR index
	Capabilities    []pci.Capability    `json:"capabilities"`
	ExtCapabilities []pci.ExtCapability `json:"ext_capabilities,omitempty"`
}

// configSpaceJSON is used for JSON serialization of config space as hex words.
type deviceContextJSON struct {
	CollectedAt     time.Time           `json:"collected_at"`
	ToolVersion     string              `json:"tool_version"`
	Hostname        string              `json:"hostname"`
	Device          pci.PCIDevice       `json:"device"`
	ConfigSpaceHex  []string            `json:"config_space_hex"`
	ConfigSpaceSize int                 `json:"config_space_size"`
	BARs            []pci.BAR           `json:"bars"`
	BARContents     map[string]string   `json:"bar_contents,omitempty"` // key: BAR index, value: base64
	Capabilities    []pci.Capability    `json:"capabilities"`
	ExtCapabilities []pci.ExtCapability `json:"ext_capabilities,omitempty"`
}

// MarshalJSON implements custom JSON marshaling for DeviceContext.
func (dc *DeviceContext) MarshalJSON() ([]byte, error) {
	j := deviceContextJSON{
		CollectedAt:     dc.CollectedAt,
		ToolVersion:     dc.ToolVersion,
		Hostname:        dc.Hostname,
		Device:          dc.Device,
		BARs:            dc.BARs,
		Capabilities:    dc.Capabilities,
		ExtCapabilities: dc.ExtCapabilities,
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
	}

	// Reconstruct config space from hex words
	if len(j.ConfigSpaceHex) > 0 {
		dc.ConfigSpace = pci.NewConfigSpace()
		dc.ConfigSpace.Size = j.ConfigSpaceSize
		for i, hexWord := range j.ConfigSpaceHex {
			var word uint32
			fmt.Sscanf(hexWord, "%x", &word)
			dc.ConfigSpace.WriteU32(i*4, word)
		}
	}

	// Reconstruct BAR contents from base64
	if len(j.BARContents) > 0 {
		dc.BARContents = make(map[int][]byte)
		for idxStr, b64 := range j.BARContents {
			idx, err := strconv.Atoi(idxStr)
			if err != nil {
				continue
			}
			data, err := base64.StdEncoding.DecodeString(b64)
			if err != nil {
				continue
			}
			dc.BARContents[idx] = data
		}
	}

	return dc, nil
}
