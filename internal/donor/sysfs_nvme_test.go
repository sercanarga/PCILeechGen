package donor

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/pci"
)

func writeNVMeSysfs(t *testing.T, base, bdfStr, model, serial, fw string) {
	t.Helper()
	nvmeDir := filepath.Join(base, bdfStr, "nvme", "nvme0")
	if err := os.MkdirAll(nvmeDir, 0o755); err != nil {
		t.Fatal(err)
	}
	for name, val := range map[string]string{
		"model":        model,
		"serial":       serial,
		"firmware_rev": fw,
	} {
		if err := os.WriteFile(filepath.Join(nvmeDir, name), []byte(val+"\n"), 0o644); err != nil {
			t.Fatal(err)
		}
	}
}

func TestSysfsReader_ReadNVMeIdentity(t *testing.T) {
	tmp := t.TempDir()
	bdf, err := pci.ParseBDF("0000:01:00.0")
	if err != nil {
		t.Fatal(err)
	}
	writeNVMeSysfs(t, tmp, bdf.String(), "Samsung SSD 990 PRO 2TB", "S6PXNG0T12345678W", "4B2QJXD7")

	sr := NewSysfsReaderWithPath(tmp)
	id, err := sr.ReadNVMeIdentity(bdf)
	if err != nil {
		t.Fatalf("ReadNVMeIdentity failed: %v", err)
	}
	if id.Serial != "S6PXNG0T12345678W" {
		t.Errorf("Serial: got %q", id.Serial)
	}
	if id.Model != "Samsung SSD 990 PRO 2TB" {
		t.Errorf("Model: got %q", id.Model)
	}
	if id.FWRev != "4B2QJXD7" {
		t.Errorf("FWRev: got %q", id.FWRev)
	}
}

func TestSysfsReader_ReadNVMeIdentity_NoNVMeDriver(t *testing.T) {
	tmp := t.TempDir()
	bdf, err := pci.ParseBDF("0000:01:00.0")
	if err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(filepath.Join(tmp, bdf.String()), 0o755); err != nil {
		t.Fatal(err)
	}
	sr := NewSysfsReaderWithPath(tmp)
	if _, err := sr.ReadNVMeIdentity(bdf); err == nil {
		t.Fatal("expected error when nvme driver is not bound")
	}
}
