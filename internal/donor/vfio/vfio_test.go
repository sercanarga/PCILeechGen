package vfio

import (
	"os"
	"path/filepath"
	"testing"
	"time"
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
		t.Log("CheckIOMMU succeeded - IOMMU is available (expected on Linux with IOMMU)")
	} else {
		t.Logf("CheckIOMMU: %v (expected on non-Linux or no IOMMU)", err)
	}
}

func TestCheckVFIOModules_NonLinux(t *testing.T) {
	err := CheckVFIOModules()
	if err == nil {
		t.Log("CheckVFIOModules succeeded - VFIO modules loaded")
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

// driver_override must be cleared with "\n", not an empty write: kernfs ignores
// zero-length writes, so "" leaves it pinned to vfio-pci and drivers_probe
// re-binds vfio-pci instead of the native driver.
func TestUnbindFromVFIO_ClearsDriverOverrideWithNewline(t *testing.T) {
	tmpDir := t.TempDir()
	SetSysfsBase(tmpDir)
	defer ResetSysfsBase()

	bdf := "0000:03:00.0"
	devDir := filepath.Join(tmpDir, bdf)
	if err := os.MkdirAll(devDir, 0755); err != nil {
		t.Fatal(err)
	}
	overridePath := filepath.Join(devDir, "driver_override")
	// Simulate the stale override that pins the device to vfio-pci.
	if err := os.WriteFile(overridePath, []byte("vfio-pci\n"), 0644); err != nil {
		t.Fatal(err)
	}

	// UnbindFromVFIO clears the override first, then writes to the hardcoded
	// /sys/bus/pci/... paths (absent on the test host) and returns an error.
	// We only assert the override-clear step here.
	_ = UnbindFromVFIO(bdf)

	got, err := os.ReadFile(overridePath)
	if err != nil {
		t.Fatalf("driver_override was not written: %v", err)
	}
	if string(got) != "\n" {
		t.Errorf("driver_override = %q, want %q (a newline clears the override; "+
			"an empty/zero-byte write is a no-op on kernfs and leaves it pinned)", got, "\n")
	}
}

func TestIsBoundToVFIO_WithFakeSysfs(t *testing.T) {
	tmpDir := t.TempDir()
	SetSysfsBase(tmpDir)
	defer ResetSysfsBase()

	bdf := "0000:03:00.0"
	devDir := filepath.Join(tmpDir, bdf)
	if err := os.MkdirAll(devDir, 0755); err != nil {
		t.Fatal(err)
	}

	// No driver symlink -> not bound
	if IsBoundToVFIO(bdf) {
		t.Error("should not be bound when no driver symlink exists")
	}

	// Create a fake driver symlink pointing to vfio-pci
	fakeDriver := filepath.Join(tmpDir, "drivers", "vfio-pci")
	if err := os.MkdirAll(fakeDriver, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(fakeDriver, filepath.Join(devDir, "driver")); err != nil {
		t.Fatal(err)
	}

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
	if err := os.MkdirAll(devDir, 0755); err != nil {
		t.Fatal(err)
	}

	// No iommu group -> "no-iommu"
	status := QuickStatus(bdf)
	if status != "no-iommu" {
		t.Errorf("QuickStatus = %q, want %q", status, "no-iommu")
	}

	// Add vfio-pci driver -> "ready"
	fakeDriver := filepath.Join(tmpDir, "drivers", "vfio-pci")
	if err := os.MkdirAll(fakeDriver, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(fakeDriver, filepath.Join(devDir, "driver")); err != nil {
		t.Fatal(err)
	}

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
	if err := os.MkdirAll(devDir, 0755); err != nil {
		t.Fatal(err)
	}

	// No power_state file -> error
	_, err := CheckPowerState(bdf)
	if err == nil {
		t.Error("should fail when power_state file doesn't exist")
	}

	// Write D0
	if writeErr := os.WriteFile(filepath.Join(devDir, "power_state"), []byte("D0\n"), 0644); writeErr != nil {
		t.Fatal(writeErr)
	}
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
	if err := os.MkdirAll(devDir, 0755); err != nil {
		t.Fatal(err)
	}

	// No resource files -> empty
	results := CheckBARAccessibility(bdf)
	if len(results) != 0 {
		t.Errorf("expected 0 results with no resource files, got %d", len(results))
	}

	// Create a non-empty resource0 file
	if err := os.WriteFile(filepath.Join(devDir, "resource0"), make([]byte, 4096), 0644); err != nil {
		t.Fatal(err)
	}
	results = CheckBARAccessibility(bdf)
	if len(results) == 0 {
		t.Error("expected at least one BAR result")
	}
	if len(results) > 0 && !results[0].Accessible {
		t.Error("BAR0 should be accessible")
	}
}

func TestEnableMemorySpace_WithFakeSysfs(t *testing.T) {
	tmpDir := t.TempDir()
	SetSysfsBase(tmpDir)
	defer ResetSysfsBase()

	bdf := "0000:03:00.0"
	devDir := filepath.Join(tmpDir, bdf)
	if err := os.MkdirAll(devDir, 0755); err != nil {
		t.Fatal(err)
	}

	// Create a fake config file (4096 bytes, Command Register at 0x04 = 0x0000)
	config := make([]byte, 4096)
	configPath := filepath.Join(devDir, "config")
	if err := os.WriteFile(configPath, config, 0644); err != nil {
		t.Fatal(err)
	}

	// Enable memory space
	err := EnableMemorySpace(bdf)
	if err != nil {
		t.Fatalf("EnableMemorySpace failed: %v", err)
	}

	// Read back and verify bits are set
	data, err := os.ReadFile(configPath)
	if err != nil {
		t.Fatalf("failed to read config: %v", err)
	}

	cmd := data[0x04]
	if cmd&0x02 == 0 {
		t.Error("Memory Space Enable bit (bit 1) should be set")
	}
	if cmd&0x04 == 0 {
		t.Error("Bus Master Enable bit (bit 2) should be set")
	}
}

func TestEnableMemorySpace_AlreadyEnabled(t *testing.T) {
	tmpDir := t.TempDir()
	SetSysfsBase(tmpDir)
	defer ResetSysfsBase()

	bdf := "0000:03:00.0"
	devDir := filepath.Join(tmpDir, bdf)
	if err := os.MkdirAll(devDir, 0755); err != nil {
		t.Fatal(err)
	}

	// Config with command register already set to 0x06
	config := make([]byte, 4096)
	config[0x04] = 0x06
	configPath := filepath.Join(devDir, "config")
	if err := os.WriteFile(configPath, config, 0644); err != nil {
		t.Fatal(err)
	}

	// Should be a no-op
	err := EnableMemorySpace(bdf)
	if err != nil {
		t.Fatalf("EnableMemorySpace failed when already enabled: %v", err)
	}
}

func TestEnableMemorySpace_NoConfigFile(t *testing.T) {
	tmpDir := t.TempDir()
	SetSysfsBase(tmpDir)
	defer ResetSysfsBase()

	err := EnableMemorySpace("0000:99:99.9")
	if err == nil {
		t.Error("EnableMemorySpace should fail for non-existent device")
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

func TestBoundDriver(t *testing.T) {
	tmpDir := t.TempDir()
	SetSysfsBase(tmpDir)
	defer ResetSysfsBase()

	bdf := "0000:03:00.0"
	devDir := filepath.Join(tmpDir, bdf)
	if err := os.MkdirAll(devDir, 0755); err != nil {
		t.Fatal(err)
	}

	if got := BoundDriver(bdf); got != "" {
		t.Errorf("no driver symlink: BoundDriver = %q, want %q", got, "")
	}

	if err := os.Symlink("/sys/bus/pci/drivers/nvme", filepath.Join(devDir, "driver")); err != nil {
		t.Fatal(err)
	}
	if got := BoundDriver(bdf); got != "nvme" {
		t.Errorf("BoundDriver = %q, want %q", got, "nvme")
	}
}

func TestWaitForNativeDriver_AlreadyBound(t *testing.T) {
	tmpDir := t.TempDir()
	SetSysfsBase(tmpDir)
	defer ResetSysfsBase()

	bdf := "0000:03:00.0"
	devDir := filepath.Join(tmpDir, bdf)
	if err := os.MkdirAll(devDir, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink("/sys/bus/pci/drivers/nvme", filepath.Join(devDir, "driver")); err != nil {
		t.Fatal(err)
	}

	drv, err := WaitForNativeDriver(bdf, 50*time.Millisecond, 10*time.Millisecond)
	if err != nil {
		t.Fatalf("WaitForNativeDriver failed: %v", err)
	}
	if drv != "nvme" {
		t.Errorf("driver = %q, want %q", drv, "nvme")
	}
}

func TestWaitForNativeDriver_TimeoutNoDriver(t *testing.T) {
	tmpDir := t.TempDir()
	SetSysfsBase(tmpDir)
	defer ResetSysfsBase()

	bdf := "0000:03:00.0"
	if err := os.MkdirAll(filepath.Join(tmpDir, bdf), 0755); err != nil {
		t.Fatal(err)
	}

	if _, err := WaitForNativeDriver(bdf, 40*time.Millisecond, 10*time.Millisecond); err == nil {
		t.Fatal("expected timeout error when no driver binds")
	}
}

func TestWaitForNativeDriver_StillVFIO(t *testing.T) {
	tmpDir := t.TempDir()
	SetSysfsBase(tmpDir)
	defer ResetSysfsBase()

	bdf := "0000:03:00.0"
	devDir := filepath.Join(tmpDir, bdf)
	if err := os.MkdirAll(devDir, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink("/sys/bus/pci/drivers/vfio-pci", filepath.Join(devDir, "driver")); err != nil {
		t.Fatal(err)
	}

	_, err := WaitForNativeDriver(bdf, 40*time.Millisecond, 10*time.Millisecond)
	if err == nil {
		t.Fatal("expected error when device stays on vfio-pci")
	}
	if !containsStr(err.Error(), "vfio-pci") {
		t.Errorf("error should mention vfio-pci, got: %v", err)
	}
}

func TestWaitForNativeDriver_BindsAfterDelay(t *testing.T) {
	tmpDir := t.TempDir()
	SetSysfsBase(tmpDir)
	defer ResetSysfsBase()

	bdf := "0000:03:00.0"
	devDir := filepath.Join(tmpDir, bdf)
	if err := os.MkdirAll(devDir, 0755); err != nil {
		t.Fatal(err)
	}
	driverPath := filepath.Join(devDir, "driver")

	go func() {
		time.Sleep(60 * time.Millisecond)
		_ = os.Symlink("/sys/bus/pci/drivers/nvme", driverPath)
	}()

	drv, err := WaitForNativeDriver(bdf, 1*time.Second, 20*time.Millisecond)
	if err != nil {
		t.Fatalf("WaitForNativeDriver failed: %v", err)
	}
	if drv != "nvme" {
		t.Errorf("driver = %q, want %q", drv, "nvme")
	}
}

func TestWaitForNVMeLive_Live(t *testing.T) {
	tmpDir := t.TempDir()
	SetSysfsBase(tmpDir)
	defer ResetSysfsBase()

	bdf := "0000:03:00.0"
	nvmeDir := filepath.Join(tmpDir, bdf, "nvme", "nvme0")
	if err := os.MkdirAll(nvmeDir, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(nvmeDir, "state"), []byte("live\n"), 0644); err != nil {
		t.Fatal(err)
	}

	if err := WaitForNVMeLive(bdf, 50*time.Millisecond, 10*time.Millisecond); err != nil {
		t.Fatalf("WaitForNVMeLive failed: %v", err)
	}
}

func TestWaitForNVMeLive_Dead(t *testing.T) {
	tmpDir := t.TempDir()
	SetSysfsBase(tmpDir)
	defer ResetSysfsBase()

	bdf := "0000:03:00.0"
	nvmeDir := filepath.Join(tmpDir, bdf, "nvme", "nvme0")
	if err := os.MkdirAll(nvmeDir, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(nvmeDir, "state"), []byte("dead\n"), 0644); err != nil {
		t.Fatal(err)
	}

	err := WaitForNVMeLive(bdf, 50*time.Millisecond, 10*time.Millisecond)
	if err == nil {
		t.Fatal("expected error for dead controller")
	}
	if !containsStr(err.Error(), "dead") {
		t.Errorf("error should mention dead state, got: %v", err)
	}
}

func TestWaitForNVMeLive_TimeoutNotLive(t *testing.T) {
	tmpDir := t.TempDir()
	SetSysfsBase(tmpDir)
	defer ResetSysfsBase()

	bdf := "0000:03:00.0"
	nvmeDir := filepath.Join(tmpDir, bdf, "nvme", "nvme0")
	if err := os.MkdirAll(nvmeDir, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(nvmeDir, "state"), []byte("connecting\n"), 0644); err != nil {
		t.Fatal(err)
	}

	err := WaitForNVMeLive(bdf, 40*time.Millisecond, 10*time.Millisecond)
	if err == nil {
		t.Fatal("expected timeout error when controller never goes live")
	}
	if !containsStr(err.Error(), "live") {
		t.Errorf("error should mention live, got: %v", err)
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
