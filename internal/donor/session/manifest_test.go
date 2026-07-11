package session

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSaveAndLoadManifestChecksumsArtifacts(t *testing.T) {
	dir := t.TempDir()
	manifest := Manifest{Version: 1, Scenario: ScenarioTrace, BDF: "0000:02:00.0"}
	files := map[string][]byte{"config-before.bin": {1, 2, 3}, "resources.txt": []byte("resource\n")}

	if err := Save(dir, &manifest, files); err != nil {
		t.Fatal(err)
	}
	loaded, err := Load(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(loaded.Artifacts) != 2 || loaded.Artifacts["config-before.bin"].SHA256 == "" {
		t.Fatalf("artifacts = %+v", loaded.Artifacts)
	}
	if err := Verify(dir, loaded); err != nil {
		t.Fatalf("Verify() error = %v", err)
	}
	if err := os.WriteFile(filepath.Join(dir, "config-before.bin"), []byte("changed"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := Verify(dir, loaded); err == nil {
		t.Fatal("Verify() accepted changed artifact")
	}
}

func TestSaveRejectsUnsafeArtifactName(t *testing.T) {
	err := Save(t.TempDir(), &Manifest{Version: 1}, map[string][]byte{"../escape": {1}})
	if err == nil {
		t.Fatal("Save() accepted path traversal")
	}
}

func TestVerifyRejectsUnsafeArtifactName(t *testing.T) {
	root := t.TempDir()
	dir := filepath.Join(root, "session")
	mustMkdir(t, dir)
	mustWrite(t, filepath.Join(root, "outside"), []byte("x"))
	manifest := &Manifest{Artifacts: map[string]Artifact{"../outside": {SHA256: "2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881", Size: 1}}}
	if err := Verify(dir, manifest); err == nil {
		t.Fatal("Verify() accepted path traversal")
	}
}
