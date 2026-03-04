package output

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestNewOutputWriter_Defaults(t *testing.T) {
	ow := NewOutputWriter("/tmp/test", "lib/pcileech-fpga", 0, 0)
	if ow.Jobs != 4 {
		t.Errorf("default Jobs = %d, want 4", ow.Jobs)
	}
	if ow.Timeout != 3600 {
		t.Errorf("default Timeout = %d, want 3600", ow.Timeout)
	}
}

func TestNewOutputWriter_CustomValues(t *testing.T) {
	ow := NewOutputWriter("/tmp/test", "lib/pcileech-fpga", 8, 7200)
	if ow.Jobs != 8 {
		t.Errorf("Jobs = %d, want 8", ow.Jobs)
	}
	if ow.Timeout != 7200 {
		t.Errorf("Timeout = %d, want 7200", ow.Timeout)
	}
}

func TestWriteFile(t *testing.T) {
	tmpDir := t.TempDir()
	ow := NewOutputWriter(tmpDir, "lib/pcileech-fpga", 4, 3600)

	content := "test content"
	if err := ow.writeFile("test.txt", content); err != nil {
		t.Fatalf("writeFile failed: %v", err)
	}

	data, err := os.ReadFile(filepath.Join(tmpDir, "test.txt"))
	if err != nil {
		t.Fatalf("ReadFile failed: %v", err)
	}
	if string(data) != content {
		t.Errorf("content = %q, want %q", string(data), content)
	}
}

func TestListOutputFiles(t *testing.T) {
	files := ListOutputFiles()
	if len(files) == 0 {
		t.Error("ListOutputFiles returned empty list")
	}

	// Check that critical files are listed
	expected := []string{
		"pcileech_cfgspace.coe",
		"pcileech_cfgspace_writemask.coe",
		"pcileech_bar_zero4k.coe",
		"vivado_generate_project.tcl",
		"vivado_build.tcl",
		"device_context.json",
	}
	for _, name := range expected {
		found := false
		for _, f := range files {
			if strings.Contains(f, name) {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("ListOutputFiles missing %q", name)
		}
	}
}
