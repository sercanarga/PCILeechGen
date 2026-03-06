package vfio

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDeviceDump_ToJSON(t *testing.T) {
	dd := &DeviceDump{
		BDF:             "0000:03:00.0",
		ConfigSpace:     make([]byte, 256),
		ConfigSpaceSize: 256,
		BARContents:     map[int][]byte{0: {0x01, 0x02, 0x03, 0x04}},
		BARInfo: []BARRegion{
			{Index: 0, Size: 4096, Flags: 0x01},
		},
	}

	data, err := dd.ToJSON()
	if err != nil {
		t.Fatalf("ToJSON failed: %v", err)
	}
	s := string(data)
	if len(s) == 0 {
		t.Error("ToJSON should return non-empty data")
	}
	if !containsStr(s, "0000:03:00.0") {
		t.Error("JSON should contain BDF")
	}
}

func TestResolveIOMMUGroup_InvalidDevice(t *testing.T) {
	_, err := ResolveIOMMUGroup("9999:99:99.9")
	if err == nil {
		t.Error("ResolveIOMMUGroup should fail for invalid device")
	}
}

func TestGetIOMMUGroup_InvalidDevice(t *testing.T) {
	_, err := GetIOMMUGroup("9999:99:99.9")
	if err == nil {
		t.Error("GetIOMMUGroup should fail for invalid device")
	}
}

func TestCheckIOMMU_NonLinux(t *testing.T) {
	err := CheckIOMMU()
	if err == nil {
		t.Log("CheckIOMMU succeeded — IOMMU is available (expected on Linux with IOMMU)")
	} else {
		t.Logf("CheckIOMMU: %v (expected on non-Linux or no IOMMU)", err)
	}
}

func TestCheckVFIOModules_NonLinux(t *testing.T) {
	err := CheckVFIOModules()
	if err == nil {
		t.Log("CheckVFIOModules succeeded — VFIO modules loaded")
	} else {
		t.Logf("CheckVFIOModules: %v (expected on non-Linux)", err)
	}
}

func TestBindToVFIO_InvalidDevice(t *testing.T) {
	err := BindToVFIO("9999:99:99.9")
	if err == nil {
		t.Error("BindToVFIO should fail for non-existent device")
	}
}

func TestUnbindFromVFIO_InvalidDevice(t *testing.T) {
	err := UnbindFromVFIO("9999:99:99.9")
	if err == nil {
		t.Log("UnbindFromVFIO succeeded (unexpected)")
	}
}

func TestIsBoundToVFIO_WithFakeSysfs(t *testing.T) {
	tmpDir := t.TempDir()
	SetSysfsBase(tmpDir)
	defer ResetSysfsBase()

	bdf := "0000:03:00.0"
	devDir := filepath.Join(tmpDir, bdf)
	os.MkdirAll(devDir, 0755)

	// No driver symlink → not bound
	if IsBoundToVFIO(bdf) {
		t.Error("should not be bound when no driver symlink exists")
	}

	// Create a fake driver symlink pointing to vfio-pci
	fakeDriver := filepath.Join(tmpDir, "drivers", "vfio-pci")
	os.MkdirAll(fakeDriver, 0755)
	os.Symlink(fakeDriver, filepath.Join(devDir, "driver"))

	if !IsBoundToVFIO(bdf) {
		t.Error("should be bound when driver symlink points to vfio-pci")
	}
}

func TestQuickStatus_WithFakeSysfs(t *testing.T) {
	tmpDir := t.TempDir()
	SetSysfsBase(tmpDir)
	defer ResetSysfsBase()

	bdf := "0000:03:00.0"
	devDir := filepath.Join(tmpDir, bdf)
	os.MkdirAll(devDir, 0755)

	// No iommu group → "no-iommu"
	status := QuickStatus(bdf)
	if status != "no-iommu" {
		t.Errorf("QuickStatus = %q, want %q", status, "no-iommu")
	}

	// Add vfio-pci driver → "ready"
	fakeDriver := filepath.Join(tmpDir, "drivers", "vfio-pci")
	os.MkdirAll(fakeDriver, 0755)
	os.Symlink(fakeDriver, filepath.Join(devDir, "driver"))

	status = QuickStatus(bdf)
	if status != "ready" {
		t.Errorf("QuickStatus = %q, want %q", status, "ready")
	}
}

func TestCheckPowerState_WithFakeSysfs(t *testing.T) {
	tmpDir := t.TempDir()
	SetSysfsBase(tmpDir)
	defer ResetSysfsBase()

	bdf := "0000:03:00.0"
	devDir := filepath.Join(tmpDir, bdf)
	os.MkdirAll(devDir, 0755)

	// No power_state file → error
	_, err := CheckPowerState(bdf)
	if err == nil {
		t.Error("should fail when power_state file doesn't exist")
	}

	// Write D0
	os.WriteFile(filepath.Join(devDir, "power_state"), []byte("D0\n"), 0644)
	state, err := CheckPowerState(bdf)
	if err != nil {
		t.Fatalf("CheckPowerState failed: %v", err)
	}
	if state != "D0" {
		t.Errorf("power state = %q, want %q", state, "D0")
	}
}

func TestCheckBARAccessibility_WithFakeSysfs(t *testing.T) {
	tmpDir := t.TempDir()
	SetSysfsBase(tmpDir)
	defer ResetSysfsBase()

	bdf := "0000:03:00.0"
	devDir := filepath.Join(tmpDir, bdf)
	os.MkdirAll(devDir, 0755)

	// No resource files → empty
	results := CheckBARAccessibility(bdf)
	if len(results) != 0 {
		t.Errorf("expected 0 results with no resource files, got %d", len(results))
	}

	// Create a non-empty resource0 file
	os.WriteFile(filepath.Join(devDir, "resource0"), make([]byte, 4096), 0644)
	results = CheckBARAccessibility(bdf)
	if len(results) == 0 {
		t.Error("expected at least one BAR result")
	}
	if len(results) > 0 && !results[0].Accessible {
		t.Error("BAR0 should be accessible")
	}
}

func TestSetSysfsBase(t *testing.T) {
	original := sysfsBase
	defer func() { sysfsBase = original }()

	SetSysfsBase("/tmp/test")
	if sysfsBase != "/tmp/test" {
		t.Errorf("sysfsBase = %q, want %q", sysfsBase, "/tmp/test")
	}

	ResetSysfsBase()
	if sysfsBase != "/sys/bus/pci/devices" {
		t.Errorf("sysfsBase = %q, want default", sysfsBase)
	}
}

func containsStr(s, sub string) bool {
	for i := 0; i <= len(s)-len(sub); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}
