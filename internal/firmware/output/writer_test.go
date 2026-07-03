package output

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

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

func TestExtraBARPresence(t *testing.T) {
	bars := []pci.BAR{
		{Index: 0, Size: 4096, Type: pci.BARTypeMem32},
		{Index: 3, Size: 65536, Type: pci.BARTypeMem32}, // real, populated
		{Index: 4, Size: 0, Type: pci.BARTypeDisabled},  // genuinely absent
		{Index: 5, Size: 256, Type: pci.BARTypeIO},      // real, populated
		// BAR6 not present at all in the donor's list.
	}

	present := extraBARPresence(bars)

	want := [4]bool{true, false, true, false} // BAR3, BAR4, BAR5, BAR6
	if present != want {
		t.Errorf("extraBARPresence() = %v, want %v", present, want)
	}
}
