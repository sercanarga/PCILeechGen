package util

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCopyFile(t *testing.T) {
	tmpDir := t.TempDir()
	src := filepath.Join(tmpDir, "source.txt")
	dst := filepath.Join(tmpDir, "dest.txt")

	content := []byte("hello world")
	if err := os.WriteFile(src, content, 0644); err != nil {
		t.Fatal(err)
	}

	if err := CopyFile(src, dst); err != nil {
		t.Fatalf("CopyFile error: %v", err)
	}

	got, err := os.ReadFile(dst)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != string(content) {
		t.Errorf("Content = %q, want %q", got, content)
	}
}

func TestCopyFileSamePath(t *testing.T) {
	// Should be a no-op
	if err := CopyFile("/tmp/same", "/tmp/same"); err != nil {
		t.Fatalf("CopyFile same path error: %v", err)
	}
}

func TestCopyDir(t *testing.T) {
	srcDir := filepath.Join(t.TempDir(), "src")
	dstDir := filepath.Join(t.TempDir(), "dst")

	// Create nested structure
	os.MkdirAll(filepath.Join(srcDir, "sub"), 0755)
	os.WriteFile(filepath.Join(srcDir, "a.txt"), []byte("aaa"), 0644)
	os.WriteFile(filepath.Join(srcDir, "sub", "b.txt"), []byte("bbb"), 0644)

	if err := CopyDir(srcDir, dstDir); err != nil {
		t.Fatalf("CopyDir error: %v", err)
	}

	// Verify files exist
	got, _ := os.ReadFile(filepath.Join(dstDir, "a.txt"))
	if string(got) != "aaa" {
		t.Errorf("a.txt = %q, want 'aaa'", got)
	}
	got, _ = os.ReadFile(filepath.Join(dstDir, "sub", "b.txt"))
	if string(got) != "bbb" {
		t.Errorf("sub/b.txt = %q, want 'bbb'", got)
	}
}
