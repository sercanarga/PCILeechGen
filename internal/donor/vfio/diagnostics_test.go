package vfio

import (
	"os"
	"path/filepath"
	"strconv"
	"testing"
)

// These tests exercise the diagnostics surface (diagnostics.go) using the
// package's own fake-sysfs seam (SetSysfsBase / ResetSysfsBase), the same
// pattern used by the neighbouring tests in vfio_test.go. No hardware is
// touched.
//
// Note on the two live-host checks: RunDiagnostics calls CheckIOMMU and
// CheckVFIOModules, which read the *real* /sys/kernel/iommu_groups and
// /sys/module/* paths (not sysfsBase), so their outcome depends on the host
// the tests run on. We therefore locate results by Name and only assert on
// the sysfsBase-controlled checks (IOMMU Group, Group Isolation, Power State,
// Driver, BARn).

// findDiag returns the DiagnosticResult with the given Name from results.
func findDiag(results []DiagnosticResult, name string) (DiagnosticResult, bool) {
	for _, r := range results {
		if r.Name == name {
			return r, true
		}
	}
	return DiagnosticResult{}, false
}

// mkFakeDev creates a fake device directory under base for bdf and returns
// its path. The returned dir is the handle the other helpers take.
func mkFakeDev(t *testing.T, base, bdf string) string {
	t.Helper()
	devDir := filepath.Join(base, bdf)
	if err := os.MkdirAll(devDir, 0755); err != nil {
		t.Fatalf("mkdir %s: %v", devDir, err)
	}
	return devDir
}

// symlinkDriver creates a fake driver dir named driver under base and points
// <devDir>/driver at it (so IsBoundToVFIO / BoundDriver resolve it). bdf is
// derived from devDir's basename.
func symlinkDriver(t *testing.T, base, devDir, driver string) {
	t.Helper()
	dir := filepath.Join(base, "drivers", driver)
	if err := os.MkdirAll(dir, 0755); err != nil {
		t.Fatalf("mkdir %s: %v", dir, err)
	}
	if err := os.Symlink(dir, filepath.Join(devDir, "driver")); err != nil {
		t.Fatalf("symlink driver: %v", err)
	}
}

// symlinkIOMMUGroup points <devDir>/iommu_group at <base>/iommu_groups/<group>,
// creates that group's devices/ dir, and adds the device (bdf, derived from
// devDir) plus each peer into devices/ so GetIOMMUGroup / ListIOMMUGroupDevices
// resolve correctly.
func symlinkIOMMUGroup(t *testing.T, base, devDir string, group int, peers ...string) {
	t.Helper()
	bdf := filepath.Base(devDir)
	groupDir := filepath.Join(base, "iommu_groups", strconv.Itoa(group))
	if err := os.MkdirAll(filepath.Join(groupDir, "devices"), 0755); err != nil {
		t.Fatalf("mkdir group devices: %v", err)
	}
	if err := os.Symlink(groupDir, filepath.Join(devDir, "iommu_group")); err != nil {
		t.Fatalf("symlink iommu_group: %v", err)
	}
	if err := os.Symlink(devDir, filepath.Join(groupDir, "devices", bdf)); err != nil {
		t.Fatalf("symlink self into group: %v", err)
	}
	for _, p := range peers {
		peerDir := filepath.Join(base, p)
		if err := os.MkdirAll(peerDir, 0755); err != nil {
			t.Fatalf("mkdir peer %s: %v", p, err)
		}
		if err := os.Symlink(peerDir, filepath.Join(groupDir, "devices", p)); err != nil {
			t.Fatalf("symlink peer %s into group: %v", p, err)
		}
	}
}

func TestListIOMMUGroupDevices_NoGroup(t *testing.T) {
	tmpDir := t.TempDir()
	SetSysfsBase(tmpDir)
	defer ResetSysfsBase()

	mkFakeDev(t, tmpDir, "0000:03:00.0")

	// No iommu_group symlink -> ListIOMMUGroupDevices must error.
	devs, err := ListIOMMUGroupDevices("0000:03:00.0")
	if err == nil {
		t.Fatal("expected error when iommu_group/devices is missing")
	}
	if devs != nil {
		t.Errorf("expected nil devices, got %v", devs)
	}
}

func TestListIOMMUGroupDevices_Populated(t *testing.T) {
	tmpDir := t.TempDir()
	SetSysfsBase(tmpDir)
	defer ResetSysfsBase()

	bdf := "0000:03:00.0"
	devDir := mkFakeDev(t, tmpDir, bdf)
	mkFakeDev(t, tmpDir, "0000:03:00.1") // peer exists as a symlink target
	symlinkIOMMUGroup(t, tmpDir, devDir, 42, "0000:03:00.1")

	devs, err := ListIOMMUGroupDevices(bdf)
	if err != nil {
		t.Fatalf("ListIOMMUGroupDevices: %v", err)
	}
	if len(devs) != 2 {
		t.Fatalf("expected 2 group devices, got %d: %v", len(devs), devs)
	}
	// Order is filesystem-dependent; check membership only.
	seen := map[string]bool{}
	for _, d := range devs {
		seen[d] = true
	}
	if !seen[bdf] || !seen["0000:03:00.1"] {
		t.Errorf("group devices = %v, want both %s and 0000:03:00.1", devs, bdf)
	}
}

func TestCheckBARAccessibility_SkipsDisabledAndMissing(t *testing.T) {
	tmpDir := t.TempDir()
	SetSysfsBase(tmpDir)
	defer ResetSysfsBase()

	bdf := "0000:03:00.0"
	devDir := mkFakeDev(t, tmpDir, bdf)

	// resource0: present and non-empty -> accessible.
	if err := os.WriteFile(filepath.Join(devDir, "resource0"), make([]byte, 4096), 0644); err != nil {
		t.Fatal(err)
	}
	// resource1: present but size 0 -> disabled, must be skipped.
	if err := os.WriteFile(filepath.Join(devDir, "resource1"), []byte{}, 0644); err != nil {
		t.Fatal(err)
	}
	// resource2..resource5: absent -> skipped via Stat error.

	results := CheckBARAccessibility(bdf)
	if len(results) != 1 {
		t.Fatalf("expected exactly 1 BAR result (only resource0 qualifies), got %d", len(results))
	}
	if results[0].Index != 0 {
		t.Errorf("first result Index = %d, want 0", results[0].Index)
	}
	if !results[0].Accessible {
		t.Errorf("BAR0 should be accessible, error: %s", results[0].Error)
	}
	if results[0].Size != 4096 {
		t.Errorf("BAR0 Size = %d, want 4096", results[0].Size)
	}
}

func TestRunDiagnostics_AloneD0BoundToVFIO(t *testing.T) {
	tmpDir := t.TempDir()
	SetSysfsBase(tmpDir)
	defer ResetSysfsBase()

	bdf := "0000:03:00.0"
	devDir := mkFakeDev(t, tmpDir, bdf)
	symlinkIOMMUGroup(t, tmpDir, devDir, 30) // alone in group 30
	symlinkDriver(t, tmpDir, devDir, "vfio-pci")
	if err := os.WriteFile(filepath.Join(devDir, "power_state"), []byte("D0\n"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(devDir, "resource0"), make([]byte, 4096), 0644); err != nil {
		t.Fatal(err)
	}

	results := RunDiagnostics(bdf)

	// IOMMU Group is sysfsBase-controlled here (fake iommu_group symlink).
	if r, ok := findDiag(results, "IOMMU Group"); !ok {
		t.Fatal("missing IOMMU Group result")
	} else if !r.Passed || r.Message != "group 30" {
		t.Errorf("IOMMU Group = %+v, want Passed 'group 30'", r)
	}

	if r, ok := findDiag(results, "Group Isolation"); !ok {
		t.Fatal("missing Group Isolation result")
	} else if !r.Passed || !containsStr(r.Message, "alone") {
		t.Errorf("Group Isolation = %+v, want Passed 'alone'", r)
	}

	if r, ok := findDiag(results, "Power State"); !ok {
		t.Fatal("missing Power State result")
	} else if !r.Passed || !containsStr(r.Message, "D0") {
		t.Errorf("Power State = %+v, want Passed D0", r)
	}

	if r, ok := findDiag(results, "Driver"); !ok {
		t.Fatal("missing Driver result")
	} else if !r.Passed || r.Message != "vfio-pci" {
		t.Errorf("Driver = %+v, want Passed 'vfio-pci'", r)
	}

	if r, ok := findDiag(results, "BAR0"); !ok {
		t.Fatal("missing BAR0 result")
	} else if !r.Passed || !containsStr(r.Message, "accessible") {
		t.Errorf("BAR0 = %+v, want Passed accessible", r)
	}
}

func TestRunDiagnostics_SharedGroup(t *testing.T) {
	tmpDir := t.TempDir()
	SetSysfsBase(tmpDir)
	defer ResetSysfsBase()

	bdf := "0000:03:00.0"
	peer := "0000:03:00.1"
	devDir := mkFakeDev(t, tmpDir, bdf)
	symlinkIOMMUGroup(t, tmpDir, devDir, 7, peer)

	results := RunDiagnostics(bdf)
	r, ok := findDiag(results, "Group Isolation")
	if !ok {
		t.Fatal("missing Group Isolation result")
	}
	if r.Passed {
		t.Errorf("Group Isolation should fail for a shared group, got %+v", r)
	}
	if !containsStr(r.Message, peer) {
		t.Errorf("Group Isolation message = %q, want it to name peer %s", r.Message, peer)
	}
	if !containsStr(r.Message, "shared") {
		t.Errorf("Group Isolation message = %q, want 'shared'", r.Message)
	}
}

func TestRunDiagnostics_NonD0Power(t *testing.T) {
	tmpDir := t.TempDir()
	SetSysfsBase(tmpDir)
	defer ResetSysfsBase()

	bdf := "0000:03:00.0"
	devDir := mkFakeDev(t, tmpDir, bdf)
	symlinkIOMMUGroup(t, tmpDir, devDir, 30)
	if err := os.WriteFile(filepath.Join(devDir, "power_state"), []byte("D3hot\n"), 0644); err != nil {
		t.Fatal(err)
	}

	results := RunDiagnostics(bdf)
	r, ok := findDiag(results, "Power State")
	if !ok {
		t.Fatal("missing Power State result")
	}
	if r.Passed {
		t.Errorf("Power State should fail for D3hot, got %+v", r)
	}
	if !containsStr(r.Message, "D3hot") {
		t.Errorf("Power State message = %q, want it to report D3hot", r.Message)
	}
	if !containsStr(r.Message, bdf) {
		t.Errorf("Power State message = %q, want it to include the BDF %s for the fix hint", r.Message, bdf)
	}
}

func TestRunDiagnostics_NativeDriverBound(t *testing.T) {
	tmpDir := t.TempDir()
	SetSysfsBase(tmpDir)
	defer ResetSysfsBase()

	bdf := "0000:03:00.0"
	devDir := mkFakeDev(t, tmpDir, bdf)
	symlinkIOMMUGroup(t, tmpDir, devDir, 30)
	symlinkDriver(t, tmpDir, devDir, "nvme") // native driver, not vfio-pci

	results := RunDiagnostics(bdf)
	r, ok := findDiag(results, "Driver")
	if !ok {
		t.Fatal("missing Driver result")
	}
	if r.Passed {
		t.Errorf("Driver should fail when a native driver is bound, got %+v", r)
	}
	if !containsStr(r.Message, "nvme") {
		t.Errorf("Driver message = %q, want it to name the bound driver nvme", r.Message)
	}
	if !containsStr(r.Message, "pcileechgen build") {
		t.Errorf("Driver message = %q, want it to suggest 'pcileechgen build'", r.Message)
	}
}

func TestRunDiagnostics_NoDriverBound(t *testing.T) {
	tmpDir := t.TempDir()
	SetSysfsBase(tmpDir)
	defer ResetSysfsBase()

	bdf := "0000:03:00.0"
	devDir := mkFakeDev(t, tmpDir, bdf)
	symlinkIOMMUGroup(t, tmpDir, devDir, 30)
	// no driver symlink -> ready for vfio-pci

	results := RunDiagnostics(bdf)
	r, ok := findDiag(results, "Driver")
	if !ok {
		t.Fatal("missing Driver result")
	}
	if !r.Passed {
		t.Errorf("Driver should pass when no driver is bound, got %+v", r)
	}
	if !containsStr(r.Message, "no driver bound") {
		t.Errorf("Driver message = %q, want 'no driver bound'", r.Message)
	}
}

func TestRunDiagnostics_MissingPowerStateOmitsEntry(t *testing.T) {
	tmpDir := t.TempDir()
	SetSysfsBase(tmpDir)
	defer ResetSysfsBase()

	bdf := "0000:03:00.0"
	devDir := mkFakeDev(t, tmpDir, bdf)
	symlinkIOMMUGroup(t, tmpDir, devDir, 30)
	// deliberately no power_state file

	results := RunDiagnostics(bdf)
	if _, ok := findDiag(results, "Power State"); ok {
		t.Error("Power State result should be omitted when power_state is unreadable")
	}
}
