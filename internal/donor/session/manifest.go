// Package session stores reproducible donor initialization captures.
package session

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/sercanarga/pcileechgen/internal/pci"
)

const manifestName = "capture.json"

// Scenario names one bounded initialization capture workflow.
type Scenario string

const (
	ScenarioTrace       Scenario = "trace"
	ScenarioUnbindBind  Scenario = "unbind-bind"
	ScenarioInterfaceUp Scenario = "interface-up"
)

// Artifact records the integrity and size of one capture file.
type Artifact struct {
	SHA256 string `json:"sha256"`
	Size   int    `json:"size"`
}

// Manifest describes one initialization capture.
type Manifest struct {
	Version   int                 `json:"version"`
	Scenario  Scenario            `json:"scenario"`
	BDF       string              `json:"bdf"`
	Driver    string              `json:"driver,omitempty"`
	Kernel    string              `json:"kernel,omitempty"`
	StartedAt time.Time           `json:"started_at"`
	Duration  time.Duration       `json:"duration"`
	Device    pci.PCIDevice       `json:"device"`
	BARs      []pci.BAR           `json:"bars,omitempty"`
	Artifacts map[string]Artifact `json:"artifacts"`
}

// Save writes capture artifacts and their manifest to dir.
func Save(dir string, manifest *Manifest, files map[string][]byte) error {
	if manifest == nil {
		return fmt.Errorf("manifest is nil")
	}
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("create session directory: %w", err)
	}
	manifest.Artifacts = make(map[string]Artifact, len(files))
	for name, data := range files {
		if name == "" || filepath.Base(name) != name || name == manifestName {
			return fmt.Errorf("unsafe artifact name %q", name)
		}
		if err := os.WriteFile(filepath.Join(dir, name), data, 0o644); err != nil {
			return fmt.Errorf("write artifact %q: %w", name, err)
		}
		sum := sha256.Sum256(data)
		manifest.Artifacts[name] = Artifact{SHA256: hex.EncodeToString(sum[:]), Size: len(data)}
	}
	data, err := json.MarshalIndent(manifest, "", "  ")
	if err != nil {
		return fmt.Errorf("encode manifest: %w", err)
	}
	if err := os.WriteFile(filepath.Join(dir, manifestName), append(data, '\n'), 0o644); err != nil {
		return fmt.Errorf("write manifest: %w", err)
	}
	return nil
}

// Load reads a capture manifest from dir.
func Load(dir string) (*Manifest, error) {
	data, err := os.ReadFile(filepath.Join(dir, manifestName))
	if err != nil {
		return nil, fmt.Errorf("read manifest: %w", err)
	}
	var manifest Manifest
	if err := json.Unmarshal(data, &manifest); err != nil {
		return nil, fmt.Errorf("decode manifest: %w", err)
	}
	return &manifest, nil
}

// Verify checks every capture artifact against its recorded checksum and size.
func Verify(dir string, manifest *Manifest) error {
	if manifest == nil {
		return fmt.Errorf("manifest is nil")
	}
	for name, artifact := range manifest.Artifacts {
		if name == "" || filepath.Base(name) != name || name == manifestName {
			return fmt.Errorf("unsafe artifact name %q", name)
		}
		data, err := os.ReadFile(filepath.Join(dir, name))
		if err != nil {
			return fmt.Errorf("read artifact %q: %w", name, err)
		}
		sum := sha256.Sum256(data)
		if len(data) != artifact.Size || hex.EncodeToString(sum[:]) != artifact.SHA256 {
			return fmt.Errorf("artifact %q failed integrity check", name)
		}
	}
	return nil
}
