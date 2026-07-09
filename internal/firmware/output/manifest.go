package output

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
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
	return generateManifest(outputDir, toolVersion, boardName, vendorID, deviceID, true)
}

func generateManifest(outputDir, toolVersion, boardName string, vendorID, deviceID uint16, includeSynthesized bool) (*BuildManifest, error) {
	m := &BuildManifest{
		GeneratedAt: time.Now(),
		ToolVersion: toolVersion,
		Board:       boardName,
		VendorID:    vendorID,
		DeviceID:    deviceID,
	}

	rootArtifacts, err := discoverRootArtifacts(outputDir, includeSynthesized)
	if err != nil {
		return nil, err
	}
	candidates := make(map[string]struct{}, len(rootArtifacts))
	for _, name := range rootArtifacts {
		candidates[name] = struct{}{}
	}

	srcDir := filepath.Join(outputDir, "src")
	if _, err := os.Stat(srcDir); err == nil {
		err = filepath.WalkDir(srcDir, func(filePath string, entry os.DirEntry, walkErr error) error {
			if walkErr != nil {
				return walkErr
			}
			if entry.IsDir() {
				return nil
			}
			if entry.Type()&os.ModeSymlink != 0 {
				return fmt.Errorf("refusing to include symlink %s", filePath)
			}
			rel, err := filepath.Rel(outputDir, filePath)
			if err != nil {
				return err
			}
			candidates[filepath.ToSlash(rel)] = struct{}{}
			return nil
		})
		if err != nil {
			return nil, fmt.Errorf("inventory source artifacts: %w", err)
		}
	} else if !os.IsNotExist(err) {
		return nil, fmt.Errorf("inspect source artifacts: %w", err)
	}

	names := make([]string, 0, len(candidates))
	for name := range candidates {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		filePath := filepath.Join(outputDir, filepath.FromSlash(name))
		info, err := os.Lstat(filePath)
		if os.IsNotExist(err) {
			continue
		}
		if err != nil {
			return nil, fmt.Errorf("stat artifact %q: %w", name, err)
		}
		if !info.Mode().IsRegular() {
			return nil, fmt.Errorf("artifact %q is not a regular file", name)
		}
		hash, err := fileHash(filePath)
		if err != nil {
			return nil, fmt.Errorf("hash artifact %q: %w", name, err)
		}
		m.Files = append(m.Files, ManifestEntry{Name: name, Size: info.Size(), SHA256: hash})
	}
	return m, nil
}

var manifestRootExtensions = map[string]struct{}{
	".bin":  {},
	".bit":  {},
	".coe":  {},
	".hex":  {},
	".json": {},
	".sv":   {},
	".svh":  {},
	".tcl":  {},
	".txt":  {},
	".v":    {},
	".vh":   {},
	".xdc":  {},
	".xci":  {},
}

func discoverRootArtifacts(outputDir string, includeSynthesized bool) ([]string, error) {
	entries, err := os.ReadDir(outputDir)
	if err != nil {
		return nil, fmt.Errorf("inventory output artifacts: %w", err)
	}

	artifacts := make([]string, 0, len(entries))
	for _, entry := range entries {
		if entry.IsDir() || entry.Name() == "build_manifest.json" {
			continue
		}
		ext := strings.ToLower(filepath.Ext(entry.Name()))
		if _, ok := manifestRootExtensions[ext]; !ok {
			continue
		}
		if !includeSynthesized && (ext == ".bit" || ext == ".bin") {
			continue
		}
		if entry.Type()&os.ModeSymlink != 0 {
			return nil, fmt.Errorf("refusing to include symlink %s", filepath.Join(outputDir, entry.Name()))
		}
		artifacts = append(artifacts, filepath.ToSlash(entry.Name()))
	}
	sort.Strings(artifacts)
	return artifacts, nil
}

func (m *BuildManifest) WriteJSON(filePath string) error {
	data, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal manifest: %w", err)
	}
	data = append(data, '\n')

	dir := filepath.Dir(filePath)
	tmp, err := os.CreateTemp(dir, ".build-manifest-*")
	if err != nil {
		return fmt.Errorf("create temporary manifest: %w", err)
	}
	tmpName := tmp.Name()
	defer os.Remove(tmpName)

	if _, err := tmp.Write(data); err != nil {
		tmp.Close()
		return fmt.Errorf("write temporary manifest: %w", err)
	}
	if err := tmp.Chmod(0644); err != nil {
		tmp.Close()
		return fmt.Errorf("set manifest permissions: %w", err)
	}
	if err := tmp.Sync(); err != nil {
		tmp.Close()
		return fmt.Errorf("sync temporary manifest: %w", err)
	}
	if err := tmp.Close(); err != nil {
		return fmt.Errorf("close temporary manifest: %w", err)
	}
	if err := os.Rename(tmpName, filePath); err != nil {
		return fmt.Errorf("replace manifest: %w", err)
	}
	return nil
}

func fileHash(filePath string) (string, error) {
	f, err := os.Open(filePath)
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

// LoadManifest reads a build manifest from a JSON file.
func LoadManifest(filePath string) (*BuildManifest, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read manifest: %w", err)
	}
	var m BuildManifest
	if err := json.Unmarshal(data, &m); err != nil {
		return nil, fmt.Errorf("failed to parse manifest: %w", err)
	}
	return &m, nil
}

type ManifestVerification struct {
	Passed  []string
	Failed  []string
	Missing []string
}

func (v *ManifestVerification) OK() bool {
	return len(v.Failed) == 0 && len(v.Missing) == 0
}

func (v *ManifestVerification) Summary() string {
	return fmt.Sprintf("%d passed, %d failed, %d missing", len(v.Passed), len(v.Failed), len(v.Missing))
}

func validateManifestName(name string) error {
	if name == "" || name == "." {
		return fmt.Errorf("invalid empty artifact path")
	}
	if strings.ContainsRune(name, '\x00') {
		return fmt.Errorf("artifact path %q contains a NUL byte", name)
	}
	if strings.Contains(name, "\\") {
		return fmt.Errorf("artifact path %q must use forward slashes", name)
	}
	if strings.HasPrefix(name, "/") || path.IsAbs(name) || path.Clean(name) != name || name == ".." || strings.HasPrefix(name, "../") {
		return fmt.Errorf("artifact path %q escapes the output directory", name)
	}
	return nil
}

func validSHA256(value string) bool {
	if len(value) != sha256.Size*2 {
		return false
	}
	_, err := hex.DecodeString(value)
	return err == nil
}

func VerifyManifest(manifestPath, outputDir string) (*ManifestVerification, error) {
	m, err := LoadManifest(manifestPath)
	if err != nil {
		return nil, err
	}

	root, err := filepath.Abs(outputDir)
	if err != nil {
		return nil, fmt.Errorf("resolve output directory: %w", err)
	}
	resolvedRoot, err := filepath.EvalSymlinks(root)
	if err != nil {
		return nil, fmt.Errorf("resolve output directory symlinks: %w", err)
	}
	seen := make(map[string]struct{}, len(m.Files))
	v := &ManifestVerification{}

	for i, entry := range m.Files {
		if err := validateManifestName(entry.Name); err != nil {
			return nil, fmt.Errorf("manifest entry %d: %w", i, err)
		}
		if _, exists := seen[entry.Name]; exists {
			return nil, fmt.Errorf("manifest entry %d: duplicate artifact path %q", i, entry.Name)
		}
		seen[entry.Name] = struct{}{}
		if entry.Size < 0 {
			return nil, fmt.Errorf("manifest entry %d (%q): negative size", i, entry.Name)
		}
		if !validSHA256(entry.SHA256) {
			return nil, fmt.Errorf("manifest entry %d (%q): invalid SHA-256", i, entry.Name)
		}

		filePath := filepath.Join(root, filepath.FromSlash(entry.Name))
		info, err := os.Lstat(filePath)
		if os.IsNotExist(err) {
			v.Missing = append(v.Missing, entry.Name)
			continue
		}
		if err != nil {
			v.Failed = append(v.Failed, fmt.Sprintf("%s: stat error: %v", entry.Name, err))
			continue
		}
		if !info.Mode().IsRegular() {
			v.Failed = append(v.Failed, fmt.Sprintf("%s: not a regular file", entry.Name))
			continue
		}
		resolvedFile, err := filepath.EvalSymlinks(filePath)
		if err != nil {
			v.Failed = append(v.Failed, fmt.Sprintf("%s: resolve path: %v", entry.Name, err))
			continue
		}
		rel, err := filepath.Rel(resolvedRoot, resolvedFile)
		if err != nil || rel == ".." || strings.HasPrefix(rel, ".."+string(filepath.Separator)) {
			v.Failed = append(v.Failed, fmt.Sprintf("%s: resolved path escapes output directory", entry.Name))
			continue
		}
		if info.Size() != entry.Size {
			v.Failed = append(v.Failed, fmt.Sprintf("%s: size mismatch (got %d, expected %d)", entry.Name, info.Size(), entry.Size))
			continue
		}
		hash, err := fileHash(filePath)
		if err != nil {
			v.Failed = append(v.Failed, fmt.Sprintf("%s: hash error: %v", entry.Name, err))
			continue
		}
		if hash != strings.ToLower(entry.SHA256) {
			v.Failed = append(v.Failed, fmt.Sprintf("%s: SHA256 mismatch", entry.Name))
			continue
		}
		v.Passed = append(v.Passed, entry.Name)
	}
	return v, nil
}
