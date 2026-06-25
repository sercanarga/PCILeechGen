package provenance

import (
	"strings"
	"testing"
)

func TestGit_returnsUnknownOrCommit(t *testing.T) {
	g := Git()
	// In a git repo we expect a real short commit; outside one, Unknown. Either
	// is acceptable as long as it's not empty.
	if g.Commit == "" {
		t.Error("Git().Commit should never be empty; expected a short hash or Unknown")
	}
	if g.Commit != Unknown && len(g.Commit) > 12 {
		t.Errorf("Git().Commit = %q, want a short hash (<=12 chars) or Unknown", g.Commit)
	}
}

func TestVivadoVersion_neverErrors(t *testing.T) {
	v := VivadoVersion()
	// Either a real version like "2023.2" or Unknown. Empty is a bug.
	if v == "" {
		t.Error("VivadoVersion() should return Unknown or a version, never empty")
	}
	if v != Unknown && !strings.Contains(v, ".") {
		t.Errorf("VivadoVersion() = %q, expected a dotted version or Unknown", v)
	}
}

func TestParseVivadoVersionLine(t *testing.T) {
	cases := []struct {
		line string
		want string
	}{
		{"Vivado v2023.2 (x86_64)", "2023.2"},
		{"vivado v2024.1.2 build something", "2024.1.2"},
		{"Vivado", ""},
		{"some other tool v1.2", ""},
		{"Vivado v2023.2.64.1", "2023.2.64.1"},
	}
	for _, c := range cases {
		if got := parseVivadoVersionLine(c.line); got != c.want {
			t.Errorf("parseVivadoVersionLine(%q) = %q, want %q", c.line, got, c.want)
		}
	}
}

func TestHashJSON_deterministic(t *testing.T) {
	type s struct {
		A string `json:"a"`
		B int    `json:"b"`
	}
	h1 := HashJSON(s{A: "x", B: 1})
	h2 := HashJSON(s{A: "x", B: 1})
	if h1 != h2 {
		t.Errorf("HashJSON not deterministic: %q vs %q", h1, h2)
	}
	if len(h1) != 64 {
		t.Errorf("HashJSON length = %d, want 64", len(h1))
	}
	h3 := HashJSON(s{A: "y", B: 1})
	if h1 == h3 {
		t.Error("HashJSON should change when input changes")
	}
}

func TestHashJSON_marshalErrorReturnsUnknown(t *testing.T) {
	// A channel cannot be marshaled to JSON.
	if got := HashJSON(make(chan int)); got != Unknown {
		t.Errorf("HashJSON on unmarshalable value = %q, want %q", got, Unknown)
	}
}

func TestCollect_populatesHashes(t *testing.T) {
	donorJSON := []byte(`{"collected_at":"0001-01-01T00:00:00Z","tool_version":"test","device":{"vendor_id":4096}}`)
	profileJSON := []byte(`{"name":"TestBoard","fpga_part":"xc7a35t"}`)
	refs := []string{"https://example/intake/1"}

	p := Collect(donorJSON, profileJSON, refs)
	if p.DonorSnapshotHash == "" || p.DonorSnapshotHash == Unknown {
		t.Errorf("DonorSnapshotHash = %q, want a real hash", p.DonorSnapshotHash)
	}
	if p.BoardProfileHash == "" || p.BoardProfileHash == Unknown {
		t.Errorf("BoardProfileHash = %q, want a real hash", p.BoardProfileHash)
	}
	if len(p.ExternalIntakeRefs) != 1 || p.ExternalIntakeRefs[0] != refs[0] {
		t.Errorf("ExternalIntakeRefs = %v, want %v", p.ExternalIntakeRefs, refs)
	}
	if p.GeneratorGitCommit == "" {
		t.Error("GeneratorGitCommit should not be empty")
	}
}

func TestCollect_omitsEmptyHashes(t *testing.T) {
	p := Collect(nil, nil, nil)
	if p.DonorSnapshotHash != "" {
		t.Errorf("DonorSnapshotHash should be omitted, got %q", p.DonorSnapshotHash)
	}
	if p.BoardProfileHash != "" {
		t.Errorf("BoardProfileHash should be omitted, got %q", p.BoardProfileHash)
	}
	if p.ExternalIntakeRefs != nil {
		t.Errorf("ExternalIntakeRefs should be nil, got %v", p.ExternalIntakeRefs)
	}
}

func TestCollect_isDeterministicForFixedGitState(t *testing.T) {
	donorJSON := []byte(`{"x":1}`)
	profileJSON := []byte(`{"y":2}`)
	// Collect twice with identical inputs: hashes and commit must match.
	// (Git state is read live, but within a clean repo it is stable across two
	// immediate calls; if the repo is dirty both calls still agree.)
	p1 := Collect(donorJSON, profileJSON, []string{"a"})
	p2 := Collect(donorJSON, profileJSON, []string{"a"})
	if p1.DonorSnapshotHash != p2.DonorSnapshotHash {
		t.Error("DonorSnapshotHash should be deterministic")
	}
	if p1.BoardProfileHash != p2.BoardProfileHash {
		t.Error("BoardProfileHash should be deterministic")
	}
	if p1.GeneratorGitCommit != p2.GeneratorGitCommit {
		t.Errorf("GeneratorGitCommit should be stable across two immediate calls: %q vs %q",
			p1.GeneratorGitCommit, p2.GeneratorGitCommit)
	}
}