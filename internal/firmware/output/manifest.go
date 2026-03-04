package output

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

type BuildManifest struct {
	GeneratedAt time.Time       `json:"generated_at"`
	ToolVersion string          `json:"tool_version"`
	Board       string          `json:"board"`
	DeviceBDF   string          `json:"device_bdf,omitempty"`
	VendorID    uint16          `json:"vendor_id"`
	DeviceID    uint16          `json:"device_id"`
	Files       []ManifestEntry `json:"files"`
}

type ManifestEntry struct {
	Name   string `json:"name"`
	Size   int64  `json:"size"`
	SHA256 string `json:"sha256"`
}

func GenerateManifest(outputDir, toolVersion, boardName string, vendorID, deviceID uint16) (*BuildManifest, error) {
	m := &BuildManifest{
		GeneratedAt: time.Now(),
		ToolVersion: toolVersion,
		Board:       boardName,
		VendorID:    vendorID,
		DeviceID:    deviceID,
	}

	expectedFiles := ListOutputFiles()
	for _, name := range expectedFiles {
		if name == "src/" {
			continue
		}
		path := filepath.Join(outputDir, name)
		info, err := os.Stat(path)
		if err != nil {
			continue
		}
		if info.IsDir() {
			continue
		}

		hash, err := fileHash(path)
		if err != nil {
			hash = "error"
		}

		m.Files = append(m.Files, ManifestEntry{
			Name:   name,
			Size:   info.Size(),
			SHA256: hash,
		})
	}

	srcDir := filepath.Join(outputDir, "src")
	if entries, err := os.ReadDir(srcDir); err == nil {
		for _, e := range entries {
			if e.IsDir() {
				continue
			}
			path := filepath.Join(srcDir, e.Name())
			info, err := e.Info()
			if err != nil {
				continue
			}
			hash, err := fileHash(path)
			if err != nil {
				hash = "error"
			}
			m.Files = append(m.Files, ManifestEntry{
				Name:   filepath.Join("src", e.Name()),
				Size:   info.Size(),
				SHA256: hash,
			})
		}
	}

	return m, nil
}

func (m *BuildManifest) WriteJSON(path string) error {
	data, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal manifest: %w", err)
	}
	return os.WriteFile(path, data, 0644)
}

func fileHash(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
