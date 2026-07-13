// Package donor collects PCI device info for firmware generation.
package donor

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/sercanarga/pcileechgen/internal/donor/behavior"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

const maxContextFileSize = 32 << 20

// MSIXData holds donor MSI-X table content.
type MSIXData struct {
	TableSize   int             `json:"table_size"`
	TableBIR    int             `json:"table_bir"`
	TableOffset uint32          `json:"table_offset"`
	PBABIR      int             `json:"pba_bir"`
	PBAOffset   uint32          `json:"pba_offset"`
	Entries     []pci.MSIXEntry `json:"entries"`
}

// NVMeIdentity holds NVMe controller strings captured from the donor device.
type NVMeIdentity struct {
	Serial string `json:"serial,omitempty"`
	Model  string `json:"model,omitempty"`
	FWRev  string `json:"firmware_rev,omitempty"`

	RawControllerIdent []byte `json:"raw_controller_ident,omitempty"`
	RawNamespaceIdent  []byte `json:"raw_namespace_ident,omitempty"`
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
	NVMeIdentity    *NVMeIdentity       `json:"nvme_identity,omitempty"`
	BehaviorRules   *behavior.RuleSet   `json:"behavior_rules,omitempty"`
}

// JSON wire format - config space as hex words, BARs as base64.
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
	NVMeIdentity    *NVMeIdentity          `json:"nvme_identity,omitempty"`
	BehaviorRules   *behavior.RuleSet      `json:"behavior_rules,omitempty"`
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
		NVMeIdentity:    dc.NVMeIdentity,
		BehaviorRules:   dc.BehaviorRules,
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

// UnmarshalJSON implements json.Unmarshaler for DeviceContext.
func (dc *DeviceContext) UnmarshalJSON(data []byte) error {
	var j deviceContextJSON
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&j); err != nil {
		return fmt.Errorf("failed to parse device context JSON: %w", err)
	}
	if err := decoder.Decode(&struct{}{}); err != io.EOF {
		return fmt.Errorf("failed to parse device context JSON: trailing data")
	}
	if j.ConfigSpaceSize != pci.ConfigSpaceLegacySize && j.ConfigSpaceSize != pci.ConfigSpaceSize {
		return fmt.Errorf("config_space_size must be %d or %d, got %d", pci.ConfigSpaceLegacySize, pci.ConfigSpaceSize, j.ConfigSpaceSize)
	}
	if len(j.ConfigSpaceHex) != j.ConfigSpaceSize/4 {
		return fmt.Errorf("config_space_hex has %d words, want %d", len(j.ConfigSpaceHex), j.ConfigSpaceSize/4)
	}

	dc.CollectedAt = j.CollectedAt
	dc.ToolVersion = j.ToolVersion
	dc.Hostname = j.Hostname
	dc.Device = j.Device
	dc.BARs = j.BARs
	dc.Capabilities = j.Capabilities
	dc.ExtCapabilities = j.ExtCapabilities
	dc.MSIXData = j.MSIXData
	dc.NVMeIdentity = j.NVMeIdentity
	dc.BehaviorRules = j.BehaviorRules

	if len(dc.BARs) > 6 {
		return fmt.Errorf("bars contains %d entries, maximum is 6", len(dc.BARs))
	}
	declaredBARs := make(map[int]pci.BAR, len(dc.BARs))
	for _, bar := range dc.BARs {
		if bar.Index < 0 || bar.Index > 5 {
			return fmt.Errorf("BAR index %d is outside 0..5", bar.Index)
		}
		if _, duplicate := declaredBARs[bar.Index]; duplicate {
			return fmt.Errorf("duplicate BAR index %d", bar.Index)
		}
		if !bar.IsDisabled() && (bar.Size == 0 || (bar.Size&(bar.Size-1)) != 0) {
			return fmt.Errorf("BAR%d size %#x is not a nonzero power of two", bar.Index, bar.Size)
		}
		declaredBARs[bar.Index] = bar
	}

	// Reconstruct config space from hex words
	if len(j.ConfigSpaceHex) > 0 {
		dc.ConfigSpace = pci.NewConfigSpace()
		dc.ConfigSpace.Size = j.ConfigSpaceSize
		for i, hexWord := range j.ConfigSpaceHex {
			if len(hexWord) != 8 {
				return fmt.Errorf("invalid config space hex word %d (%q): want exactly 8 hexadecimal digits", i, hexWord)
			}
			parsed, err := strconv.ParseUint(hexWord, 16, 32)
			if err != nil {
				return fmt.Errorf("invalid config space hex word %d (%q): %w", i, hexWord, err)
			}
			dc.ConfigSpace.WriteU32(i*4, uint32(parsed))
		}
	}

	// Reconstruct BAR contents from base64
	if len(j.BARContents) > 0 {
		dc.BARContents = make(map[int][]byte)
		for idxStr, b64 := range j.BARContents {
			idx, err := strconv.Atoi(idxStr)
			if err != nil || idx < 0 || idx > 5 {
				return fmt.Errorf("invalid BAR index %q", idxStr)
			}
			bar, declared := declaredBARs[idx]
			if !declared || bar.IsDisabled() || bar.IsIO() {
				return fmt.Errorf("BAR%d content references an undeclared or non-memory BAR", idx)
			}
			barData, err := base64.StdEncoding.DecodeString(b64)
			if err != nil {
				return fmt.Errorf("failed to decode BAR%d content: %w", idx, err)
			}
			if len(barData) > maxBARReadSize || uint64(len(barData)) > bar.Size {
				return fmt.Errorf("BAR%d content size %d exceeds capture/aperture limit", idx, len(barData))
			}
			dc.BARContents[idx] = barData
		}
	}

	// Reconstruct BAR profiles
	if len(j.BARProfiles) > 0 {
		dc.BARProfiles = make(map[int]*BARProfile)
		for idxStr, profile := range j.BARProfiles {
			idx, err := strconv.Atoi(idxStr)
			if err != nil || idx < 0 || idx > 5 {
				return fmt.Errorf("invalid BAR profile index %q", idxStr)
			}
			bar, declared := declaredBARs[idx]
			if !declared || bar.IsDisabled() || bar.IsIO() || profile == nil {
				return fmt.Errorf("BAR%d profile references an undeclared or non-memory BAR", idx)
			}
			if profile.BarIndex != idx || profile.Size < 0 || profile.Size > maxBARReadSize || uint64(profile.Size) > bar.Size {
				return fmt.Errorf("BAR%d profile has invalid index or size", idx)
			}
			seenOffsets := make(map[uint32]struct{}, len(profile.Probes))
			for _, probe := range profile.Probes {
				if probe.Offset%4 != 0 || uint64(probe.Offset)+4 > uint64(profile.Size) {
					return fmt.Errorf("BAR%d profile probe offset %#x is out of range or unaligned", idx, probe.Offset)
				}
				if _, duplicate := seenOffsets[probe.Offset]; duplicate {
					return fmt.Errorf("BAR%d profile contains duplicate probe offset %#x", idx, probe.Offset)
				}
				if probe.W1CMask&^probe.RWMask != 0 {
					return fmt.Errorf("BAR%d profile W1C mask is not a subset of RW mask at %#x", idx, probe.Offset)
				}
				seenOffsets[probe.Offset] = struct{}{}
			}
			dc.BARProfiles[idx] = profile
		}
	}

	if dc.NVMeIdentity != nil {
		for name, raw := range map[string][]byte{
			"controller": dc.NVMeIdentity.RawControllerIdent,
			"namespace":  dc.NVMeIdentity.RawNamespaceIdent,
		} {
			if len(raw) != 0 && len(raw) != 4096 {
				return fmt.Errorf("raw NVMe %s identify data must be exactly 4096 bytes", name)
			}
		}
	}
	if dc.MSIXData != nil {
		m := dc.MSIXData
		if m.TableSize < 1 || m.TableSize > 2048 || len(m.Entries) > m.TableSize {
			return fmt.Errorf("invalid MSI-X table size or entry count")
		}
		for name, bir := range map[string]int{"table": m.TableBIR, "PBA": m.PBABIR} {
			bar, ok := declaredBARs[bir]
			if !ok || bar.IsDisabled() || bar.IsIO() {
				return fmt.Errorf("MSI-X %s references invalid BAR%d", name, bir)
			}
		}
		if m.TableOffset%8 != 0 || m.PBAOffset%8 != 0 {
			return fmt.Errorf("MSI-X table and PBA offsets must be 8-byte aligned")
		}
		tableBytes := uint64(m.TableSize) * 16
		pbaBytes := uint64((m.TableSize+63)/64) * 8
		if uint64(m.TableOffset)+tableBytes > declaredBARs[m.TableBIR].Size ||
			uint64(m.PBAOffset)+pbaBytes > declaredBARs[m.PBABIR].Size {
			return fmt.Errorf("MSI-X table or PBA exceeds its BAR aperture")
		}
	}

	return nil
}

// FromJSON deserializes a DeviceContext from JSON.
func FromJSON(data []byte) (*DeviceContext, error) {
	if len(data) > maxContextFileSize {
		return nil, fmt.Errorf("device context exceeds %d-byte limit", maxContextFileSize)
	}
	dc := &DeviceContext{}
	if err := json.Unmarshal(data, dc); err != nil {
		return nil, err
	}

	if dc.ConfigSpace == nil {
		return nil, fmt.Errorf("config_space_hex not found in JSON, file may be from an unsupported tool")
	}

	// Capability caches are derived data; never trust serialized copies.
	dc.Capabilities = pci.ParseCapabilities(dc.ConfigSpace)
	dc.ExtCapabilities = pci.ParseExtCapabilities(dc.ConfigSpace)

	if dc.BehaviorRules != nil {
		if err := behavior.Validate(dc.BehaviorRules); err != nil {
			return nil, fmt.Errorf("invalid behavior_rules: %w", err)
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
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read device context file: %w", err)
	}
	defer f.Close()
	data, err := io.ReadAll(io.LimitReader(f, maxContextFileSize+1))
	if err != nil {
		return nil, fmt.Errorf("failed to read device context file: %w", err)
	}
	if len(data) > maxContextFileSize {
		return nil, fmt.Errorf("device context exceeds %d-byte limit", maxContextFileSize)
	}
	return FromJSON(data)
}
