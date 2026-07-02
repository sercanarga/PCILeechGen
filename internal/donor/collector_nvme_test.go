package donor

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/pci"
)

func TestCollector_collectNVMeIdentity_DirectRead(t *testing.T) {
	tmp := t.TempDir()
	bdf, err := pci.ParseBDF("0000:01:00.0")
	if err != nil {
		t.Fatal(err)
	}
	writeNVMeSysfs(t, tmp, bdf.String(), "Samsung SSD 990 PRO 2TB", "S6PXNG0T12345678W", "4B2QJXD7")

	c := NewCollectorWithSysfs(NewSysfsReaderWithPath(tmp))
	id := c.collectNVMeIdentity(bdf, 0x010802, nil) // base=01 sub=08 (NVMe)
	if id == nil {
		t.Fatal("expected captured identity for NVMe device")
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

func TestCollector_collectNVMeIdentity_NonNVMeReturnsNil(t *testing.T) {
	c := NewCollector()
	bdf, _ := pci.ParseBDF("0000:01:00.0")
	if id := c.collectNVMeIdentity(bdf, 0x010601, nil); id != nil { // SATA AHCI
		t.Errorf("expected nil identity for non-NVMe class, got %+v", id)
	}
}

func TestCollector_collectNVMeIdentity_NoNVMeDriverReturnsNil(t *testing.T) {
	tmp := t.TempDir()
	bdf, err := pci.ParseBDF("0000:01:00.0")
	if err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(filepath.Join(tmp, bdf.String()), 0o755); err != nil {
		t.Fatal(err)
	}
	c := NewCollectorWithSysfs(NewSysfsReaderWithPath(tmp))
	if id := c.collectNVMeIdentity(bdf, 0x010802, nil); id != nil {
		t.Errorf("expected nil when nvme driver absent, got %+v", id)
	}
}
