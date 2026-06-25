// Package provenance records stable, deterministic build provenance for
// generated firmware manifests: generator git commit, dirty flag, content
// hashes, and optionally-detected tool versions.
//
// All fields are content- or VCS-derived (never wall-clock), so a future
// deterministic build mode can reproduce them byte-for-byte. When a source
// is unavailable the helpers degrade to explicit "unknown"/omitted values
// rather than erroring — a missing Vivado install or git metadata must never
// fail a build.
package provenance

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"os/exec"
	"strings"
)

// Unknown is the sentinel value for unavailable provenance fields.
const Unknown = "unknown"

// GitInfo captures generator VCS state at manifest build time.
type GitInfo struct {
	Commit string // short commit hash, or Unknown
	Dirty  bool   // true if the working tree has uncommitted changes
}

// Git collects the generator's git commit and dirty flag by shelling out to
// git. On any error it returns Commit=Unknown and Dirty=true (conservative:
// treat unreadable state as dirty). It is called lazily by the manifest
// builder, never at package import.
func Git() GitInfo {
	commit, err := shortCommit()
	if err != nil {
		return GitInfo{Commit: Unknown, Dirty: true}
	}
	dirty, err := isDirty()
	if err != nil {
		// ponytail: ceiling — on git-status failure we conservatively mark
		// dirty=true so a broken git never silently claims cleanliness.
		// Upgrade path: surface the git error via slog when a logger is wired
		// through this package.
		dirty = true
	}
	return GitInfo{Commit: commit, Dirty: dirty}
}

func shortCommit() (string, error) {
	out, err := exec.Command("git", "rev-parse", "--short", "HEAD").Output()
	if err != nil {
		return "", err
	}
	c := strings.TrimSpace(string(out))
	if c == "" {
		return "", &exec.Error{Name: "git", Err: errEmptyCommit}
	}
	return c, nil
}

func isDirty() (bool, error) {
	out, err := exec.Command("git", "status", "--porcelain").Output()
	if err != nil {
		return false, err
	}
	return len(strings.TrimSpace(string(out))) > 0, nil
}

// VivadoVersion returns the detected Vivado version string, or Unknown when
// Vivado is not on PATH or its version cannot be parsed. Never errors.
//
// ponytail: ceiling — direct exec.LookPath + `vivado -version` parse instead
// of importing internal/vivado (which would create an import cycle via
// internal/firmware/output). Upgrade path: share a small version-detection
// helper in a leaf package if a second consumer needs it.
func VivadoVersion() string {
	path, err := exec.LookPath("vivado")
	if err != nil {
		return Unknown
	}
	out, err := exec.Command(path, "-version").Output()
	if err != nil {
		return Unknown
	}
	// First line typically: "Vivado v2023.2 (x86_64)". Extract the vX.Y token.
	for _, line := range strings.Split(string(out), "\n") {
		if v := parseVivadoVersionLine(line); v != "" {
			return v
		}
	}
	return Unknown
}

// parseVivadoVersionLine extracts a "vX.Y" style version token from a single
// `vivado -version` output line, returning the bare version (e.g. "2023.2")
// or "" if none is found.
func parseVivadoVersionLine(line string) string {
	low := strings.ToLower(line)
	if !strings.Contains(low, "vivado") {
		return ""
	}
	// Find the 'v' that immediately precedes the version digits, not the 'v'
	// in "vivado" itself: scan for a 'v'/'V' whose next char is a digit.
	idx := -1
	for i := 0; i < len(line); i++ {
		c := line[i]
		if (c == 'v' || c == 'V') && i+1 < len(line) && line[i+1] >= '0' && line[i+1] <= '9' {
			idx = i + 1
			break
		}
	}
	if idx < 0 {
		return ""
	}
	rest := line[idx:]
	var b strings.Builder
	for _, r := range rest {
		if (r >= '0' && r <= '9') || r == '.' {
			b.WriteRune(r)
		} else {
			break
		}
	}
	v := strings.TrimSpace(b.String())
	if v == "" || !strings.Contains(v, ".") {
		return ""
	}
	return v
}

// HashJSON returns the lowercase hex SHA-256 of the canonical JSON encoding
// of v. On marshal failure it returns Unknown so the field is explicit rather
// than empty. Deterministic for a fixed input.
func HashJSON(v any) string {
	data, err := json.Marshal(v)
	if err != nil {
		return Unknown
	}
	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:])
}

// BuildProvenance is the stable provenance section embedded in a build
// manifest. All fields are deterministic (VCS refs + content hashes) or
// explicitly Unknown/omitted when unavailable.
type BuildProvenance struct {
	GeneratorGitCommit string   `json:"generator_git_commit"`           // short commit or Unknown
	GeneratorDirty     bool     `json:"generator_dirty"`                // working tree had uncommitted changes
	DonorSnapshotHash  string   `json:"donor_snapshot_hash,omitempty"`  // SHA-256 of donor DeviceContext JSON, omitted when no donor
	BoardProfileHash   string   `json:"board_profile_hash,omitempty"`   // SHA-256 of board_profile JSON, omitted when no profile
	VivadoVersion      string   `json:"vivado_version,omitempty"`       // detected Vivado version, omitted when Unknown
	ExternalIntakeRefs []string `json:"external_intake_refs,omitempty"` // populated only when profile-derived behavior is used
}

// Collect assembles a BuildProvenance from the donor snapshot JSON, board
// profile JSON, and any external intake references. donorJSON and
// profileJSON may be nil/empty, in which case the corresponding hash is
// omitted. vivado detection is best-effort.
func Collect(donorJSON, profileJSON []byte, intakeRefs []string) *BuildProvenance {
	g := Git()
	p := &BuildProvenance{
		GeneratorGitCommit: g.Commit,
		GeneratorDirty:     g.Dirty,
		ExternalIntakeRefs: intakeRefs,
	}
	if len(donorJSON) > 0 {
		sum := sha256.Sum256(donorJSON)
		p.DonorSnapshotHash = hex.EncodeToString(sum[:])
	}
	if len(profileJSON) > 0 {
		sum := sha256.Sum256(profileJSON)
		p.BoardProfileHash = hex.EncodeToString(sum[:])
	}
	if vv := VivadoVersion(); vv != Unknown {
		p.VivadoVersion = vv
	}
	return p
}

// errEmptyCommit is a sentinel for an empty rev-parse result.
var errEmptyCommit = &exec.Error{Name: "git", Err: exec.ErrNotFound}