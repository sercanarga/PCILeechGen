package output

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/firmware/barprofile"
	"github.com/sercanarga/pcileechgen/internal/pci"
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

func TestPrepareOutputDirCreatesOwnedOutput(t *testing.T) {
	target := filepath.Join(t.TempDir(), "generated-output")
	ow := NewOutputWriter(target, "lib/pcileech-fpga", 4, 3600)

	if err := ow.prepareOutputDir(); err != nil {
		t.Fatalf("prepareOutputDir failed: %v", err)
	}
	marker, err := os.ReadFile(filepath.Join(ow.OutputDir, outputOwnershipMarker))
	if err != nil {
		t.Fatalf("read ownership marker: %v", err)
	}
	if string(marker) != outputOwnershipContent {
		t.Fatalf("ownership marker = %q, want %q", marker, outputOwnershipContent)
	}

	if err := ow.prepareOutputDir(); err != nil {
		t.Fatalf("prepareOutputDir should reopen its owned output: %v", err)
	}
}

func TestPrepareOutputDirRejectsUnownedExistingDirectory(t *testing.T) {
	target := filepath.Join(t.TempDir(), "existing-output")
	if err := os.Mkdir(target, 0755); err != nil {
		t.Fatal(err)
	}

	err := NewOutputWriter(target, "lib/pcileech-fpga", 4, 3600).prepareOutputDir()
	if err == nil || !strings.Contains(err.Error(), "unowned") {
		t.Fatalf("prepareOutputDir error = %v, want unowned-output rejection", err)
	}
}

func TestPrepareOutputDirRejectsSymlink(t *testing.T) {
	root := t.TempDir()
	outside := filepath.Join(root, "outside")
	if err := os.Mkdir(outside, 0755); err != nil {
		t.Fatal(err)
	}
	target := filepath.Join(root, "output")
	if err := os.Symlink(outside, target); err != nil {
		t.Skipf("symlink support unavailable: %v", err)
	}

	err := NewOutputWriter(target, "lib/pcileech-fpga", 4, 3600).prepareOutputDir()
	if err == nil || !strings.Contains(err.Error(), "real directory") {
		t.Fatalf("prepareOutputDir error = %v, want symlink rejection", err)
	}
}

func TestPublishOutputDirectoryReplacesWholeTree(t *testing.T) {
	root := t.TempDir()
	target := filepath.Join(root, "output")
	stage := filepath.Join(root, "stage")
	if err := os.Mkdir(target, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(target, "stale.sv"), []byte("stale"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.Mkdir(stage, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(stage, "current.sv"), []byte("current"), 0644); err != nil {
		t.Fatal(err)
	}

	if err := publishOutputDirectory(stage, target); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(target, "stale.sv")); !os.IsNotExist(err) {
		t.Fatalf("stale artifact survived replacement: %v", err)
	}
	if data, err := os.ReadFile(filepath.Join(target, "current.sv")); err != nil || string(data) != "current" {
		t.Fatalf("current artifact missing: %q, %v", data, err)
	}
}

func TestExtractSubModulesFailsWhenRequiredModuleMissing(t *testing.T) {
	src := filepath.Join(t.TempDir(), "controller.sv")
	if err := os.WriteFile(src, []byte("module present; endmodule\n"), 0644); err != nil {
		t.Fatal(err)
	}
	err := extractSubModules(src, t.TempDir(), []string{"missing"})
	if err == nil || !strings.Contains(err.Error(), "required sub-module missing") {
		t.Fatalf("extractSubModules error = %v", err)
	}
}

func TestWriteAllFailurePreservesPreviousOutput(t *testing.T) {
	target := filepath.Join(t.TempDir(), "output")
	ow := NewOutputWriter(target, filepath.Join(t.TempDir(), "missing-lib"), 1, 1)
	if err := ow.prepareOutputDir(); err != nil {
		t.Fatal(err)
	}
	sentinel := filepath.Join(target, "previous.txt")
	if err := os.WriteFile(sentinel, []byte("keep"), 0644); err != nil {
		t.Fatal(err)
	}

	err := ow.WriteAll(outputModelContext(), &board.Board{Name: "missing", ProjectDir: "missing"})
	if err == nil {
		t.Fatal("WriteAll unexpectedly succeeded")
	}
	data, readErr := os.ReadFile(sentinel)
	if readErr != nil || string(data) != "keep" {
		t.Fatalf("previous output changed after failed generation: %q, %v", data, readErr)
	}
}

func TestWriteAllCreatesNestedOutputParent(t *testing.T) {
	target := filepath.Join(t.TempDir(), "generated", "nvme", "board")
	ow := NewOutputWriter(target, filepath.Join(t.TempDir(), "missing-lib"), 1, 1)

	err := ow.WriteAll(outputModelContext(), &board.Board{Name: "missing", ProjectDir: "missing"})
	if err == nil {
		t.Fatal("WriteAll unexpectedly succeeded")
	}
	if strings.Contains(err.Error(), "create output staging directory") {
		t.Fatalf("nested output parent was not prepared: %v", err)
	}
}

func TestWriteFileRejectsSymlink(t *testing.T) {
	root := t.TempDir()
	outside := filepath.Join(root, "outside.txt")
	if err := os.WriteFile(outside, []byte("preserve"), 0644); err != nil {
		t.Fatal(err)
	}
	output := filepath.Join(root, "output")
	if err := os.Mkdir(output, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(outside, filepath.Join(output, "test.txt")); err != nil {
		t.Skipf("symlink support unavailable: %v", err)
	}

	err := NewOutputWriter(output, "lib/pcileech-fpga", 4, 3600).writeFile("test.txt", "replace")
	if err == nil || !strings.Contains(err.Error(), "regular file") {
		t.Fatalf("writeFile error = %v, want symlink rejection", err)
	}
	contents, readErr := os.ReadFile(outside)
	if readErr != nil {
		t.Fatal(readErr)
	}
	if string(contents) != "preserve" {
		t.Fatalf("outside file was modified: %q", contents)
	}
}

func TestClearGeneratedSourceTreeRejectsSymlink(t *testing.T) {
	root := t.TempDir()
	dst := filepath.Join(root, "src")
	if err := os.Mkdir(dst, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(root, filepath.Join(dst, "escape")); err != nil {
		t.Skipf("symlink support unavailable: %v", err)
	}

	err := clearGeneratedSourceTree(dst)
	if err == nil || !strings.Contains(err.Error(), "symlink") {
		t.Fatalf("clearGeneratedSourceTree error = %v, want symlink rejection", err)
	}
	if _, statErr := os.Lstat(filepath.Join(dst, "escape")); statErr != nil {
		t.Fatalf("unsafe source tree was modified: %v", statErr)
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

func TestWriteBARBehaviorProfile_writesProfileArtifact(t *testing.T) {
	tmpDir := t.TempDir()
	ow := NewOutputWriter(tmpDir, "lib/pcileech-fpga", 4, 3600)
	ctx := &donor.DeviceContext{
		Device: pci.PCIDevice{ClassCode: 0x020000},
		BARs: []pci.BAR{
			{Index: 2, Size: 4, Type: pci.BARTypeMem32},
		},
		BARContents: map[int][]byte{
			2: {0x11, 0x22, 0x33, 0x44},
		},
	}

	if err := ow.writeBARBehaviorProfile(ctx); err != nil {
		t.Fatalf("writeBARBehaviorProfile failed: %v", err)
	}

	data, err := os.ReadFile(filepath.Join(tmpDir, "bar_behavior_profile.json"))
	if err != nil {
		t.Fatalf("ReadFile failed: %v", err)
	}
	var profile barprofile.Profile
	if err := json.Unmarshal(data, &profile); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if profile.ClassCode != 0x020000 {
		t.Fatalf("ClassCode = 0x%06X, want 0x020000", profile.ClassCode)
	}
	if len(profile.BARs) != 1 || profile.BARs[0].Index != 2 {
		t.Fatalf("unexpected BAR profile: %+v", profile.BARs)
	}
}
