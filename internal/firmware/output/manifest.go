package output

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/provenance"
)

type BuildManifest struct {
	GeneratedAt  time.Time                 `json:"generated_at"`
	ToolVersion  string                    `json:"tool_version"`
	Board        string                    `json:"board"`
	BoardProfile *BoardProfile             `json:"board_profile,omitempty"`
	Provenance   *provenance.BuildProvenance `json:"provenance,omitempty"`
	DeviceBDF    string                    `json:"device_bdf,omitempty"`
	VendorID     uint16                    `json:"vendor_id"`
	DeviceID     uint16                    `json:"device_id"`
	Files        []ManifestEntry           `json:"files"`
}

type BoardProfile struct {
	Name         string `json:"name"`
	FPGAPart     string `json:"fpga_part"`
	PCIeLanes    int    `json:"pcie_lanes"`
	MaxLinkSpeed uint8  `json:"max_link_speed"`
	BRAMSize     int    `json:"bram_size"`
	TopModule    string `json:"top_module"`
	ProjectDir   string `json:"project_dir"`
	SubDir       string `json:"sub_dir,omitempty"`
	TCLFile      string `json:"tcl_file"`
	BuildTCL     string `json:"build_tcl"`
}

type ManifestEntry struct {
	Name   string `json:"name"`
	Size   int64  `json:"size"`
	SHA256 string `json:"sha256"`
}

func GenerateManifest(outputDir, toolVersion, boardName string, vendorID, deviceID uint16) (*BuildManifest, error) {
	return generateManifest(outputDir, toolVersion, boardName, nil, nil, nil, vendorID, deviceID)
}

func GenerateManifestForBoard(outputDir, toolVersion string, b *board.Board, vendorID, deviceID uint16) (*BuildManifest, error) {
	return GenerateManifestForBuild(outputDir, toolVersion, b, nil, nil, vendorID, deviceID)
}

// GenerateManifestForBuild assembles a manifest with full provenance: donor
// snapshot hash, board profile hash, generator git metadata, and external
// intake references. ctx may be nil; intakeRefs may be nil. Provenance is
// always set, with unavailable sources degrading to Unknown/omitted fields.
func GenerateManifestForBuild(outputDir, toolVersion string, b *board.Board, ctx *donor.DeviceContext, intakeRefs []string, vendorID, deviceID uint16) (*BuildManifest, error) {
	boardName := ""
	var profile *BoardProfile
	if b != nil {
		boardName = b.Name
		profile = boardProfileFromBoard(b)
	}
	return generateManifest(outputDir, toolVersion, boardName, profile, ctx, intakeRefs, vendorID, deviceID)
}

func generateManifest(outputDir, toolVersion, boardName string, profile *BoardProfile, ctx *donor.DeviceContext, intakeRefs []string, vendorID, deviceID uint16) (*BuildManifest, error) {
	m := &BuildManifest{
		GeneratedAt:  time.Now(),
		ToolVersion:  toolVersion,
		Board:        boardName,
		BoardProfile: profile,
		Provenance:   collectProvenance(profile, ctx, intakeRefs),
		VendorID:     vendorID,
		DeviceID:     deviceID,
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

func boardProfileFromBoard(b *board.Board) *BoardProfile {
	return &BoardProfile{
		Name:         b.Name,
		FPGAPart:     b.FPGAPart,
		PCIeLanes:    b.PCIeLanes,
		MaxLinkSpeed: b.MaxLinkSpeedOrDefault(),
		BRAMSize:     b.BRAMSizeOrDefault(),
		TopModule:    b.TopModule,
		ProjectDir:   b.ProjectDir,
		SubDir:       b.SubDir,
		TCLFile:      b.TCLFile,
		BuildTCL:     b.BuildTCL,
	}
}

// collectProvenance builds the provenance section from the board profile and
// donor context. Both may be nil; the resulting fields degrade to omitted or
// Unknown rather than erroring.
func collectProvenance(profile *BoardProfile, ctx *donor.DeviceContext, intakeRefs []string) *provenance.BuildProvenance {
	var profileJSON, donorJSON []byte
	if profile != nil {
		// ponytail: ceiling — Marshal is deterministic for our struct (fixed
		// field order). Upgrade path: canonical JSON if non-deterministic
		// types are ever added to BoardProfile.
		if data, err := json.Marshal(profile); err == nil {
			profileJSON = data
		}
	}
	if ctx != nil {
		if data, err := ctx.ToJSON(); err == nil {
			donorJSON = data
		}
	}
	return provenance.Collect(donorJSON, profileJSON, intakeRefs)
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

// LoadManifest reads a build manifest from a JSON file.
func LoadManifest(path string) (*BuildManifest, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read manifest: %w", err)
	}
	var m BuildManifest
	if err := json.Unmarshal(data, &m); err != nil {
		return nil, fmt.Errorf("failed to parse manifest: %w", err)
	}
	return &m, nil
}

// ManifestVerification holds the result of verifying a build manifest.
type ManifestVerification struct {
	Passed  []string
	Failed  []string
	Missing []string
}

func (v *ManifestVerification) OK() bool {
	return len(v.Failed) == 0 && len(v.Missing) == 0
}

func (v *ManifestVerification) Summary() string {
	return fmt.Sprintf("%d passed, %d failed, %d missing",
		len(v.Passed), len(v.Failed), len(v.Missing))
}

// VerifyManifest checks that all files in the manifest exist and match their checksums.
func VerifyManifest(manifestPath, outputDir string) (*ManifestVerification, error) {
	m, err := LoadManifest(manifestPath)
	if err != nil {
		return nil, err
	}

	v := &ManifestVerification{}

	for _, entry := range m.Files {
		filePath := filepath.Join(outputDir, entry.Name)

		info, err := os.Stat(filePath)
		if err != nil {
			v.Missing = append(v.Missing, entry.Name)
			continue
		}

		if info.Size() != entry.Size {
			v.Failed = append(v.Failed,
				fmt.Sprintf("%s: size mismatch (got %d, expected %d)", entry.Name, info.Size(), entry.Size))
			continue
		}

		hash, err := fileHash(filePath)
		if err != nil {
			v.Failed = append(v.Failed, fmt.Sprintf("%s: hash error: %v", entry.Name, err))
			continue
		}

		if hash != entry.SHA256 {
			v.Failed = append(v.Failed,
				fmt.Sprintf("%s: SHA256 mismatch", entry.Name))
			continue
		}

		v.Passed = append(v.Passed, entry.Name)
	}

	return v, nil
}
