package vfio

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

// RunDiagnostics includes host-dependent checks, so tests assert by result name.
func findDiag(results []DiagnosticResult, name string) (DiagnosticResult, bool) {
	for _, r := range results {
		if r.Name == name {
			return r, true
		}
	}
	return DiagnosticResult{}, false
}

func mkFakeDev(t *testing.T, base, bdf string) string {
	t.Helper()
	devDir := filepath.Join(base, bdf)
	if err := os.MkdirAll(devDir, 0755); err != nil {
		t.Fatalf("mkdir %s: %v", devDir, err)
	}
	return devDir
}

func fakeInitialMountNamespace(t *testing.T) {
	t.Helper()
	dir := t.TempDir()
	self := filepath.Join(dir, "self")
	init := filepath.Join(dir, "init")
	if err := os.Symlink("mnt:[1]", self); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink("mnt:[1]", init); err != nil {
		t.Fatal(err)
	}
	oldSelf, oldInit := selfMountNamespacePath, initMountNamespacePath
	selfMountNamespacePath, initMountNamespacePath = self, init
	t.Cleanup(func() {
		selfMountNamespacePath, initMountNamespacePath = oldSelf, oldInit
	})
}

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
	peer := "0000:03:00.1"
	mkFakeDev(t, tmpDir, peer)
	symlinkIOMMUGroup(t, tmpDir, devDir, 42, peer)

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

	if err := os.WriteFile(filepath.Join(devDir, "resource0"), make([]byte, 4096), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(devDir, "resource1"), []byte{}, 0644); err != nil {
		t.Fatal(err)
	}

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
	symlinkIOMMUGroup(t, tmpDir, devDir, 30)
	symlinkDriver(t, tmpDir, devDir, "vfio-pci")
	if err := os.WriteFile(filepath.Join(devDir, "power_state"), []byte("D0\n"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(devDir, "resource0"), make([]byte, 4096), 0644); err != nil {
		t.Fatal(err)
	}

	results := RunDiagnostics(bdf)

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
	symlinkDriver(t, tmpDir, devDir, "nvme")

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

	results := RunDiagnostics(bdf)
	if _, ok := findDiag(results, "Power State"); ok {
		t.Error("Power State result should be omitted when power_state is unreadable")
	}
}

func TestCheckIOMMUGroupSafe(t *testing.T) {
	for _, tc := range []struct {
		name       string
		peerDriver string
		wantErr    bool
	}{
		{name: "alone"},
		{name: "peer unbound"},
		{name: "peer on vfio", peerDriver: "vfio-pci"},
		{name: "peer on native driver", peerDriver: "nvme", wantErr: true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			tmpDir := t.TempDir()
			SetSysfsBase(tmpDir)
			defer ResetSysfsBase()

			bdf := "0000:03:00.0"
			devDir := mkFakeDev(t, tmpDir, bdf)
			if tc.name == "alone" {
				symlinkIOMMUGroup(t, tmpDir, devDir, 42)
			} else {
				peer := "0000:03:00.1"
				peerDir := mkFakeDev(t, tmpDir, peer)
				symlinkIOMMUGroup(t, tmpDir, devDir, 42, peer)
				if tc.peerDriver != "" {
					symlinkDriver(t, tmpDir, peerDir, tc.peerDriver)
				}
			}

			err := CheckIOMMUGroupSafe(bdf)
			if tc.wantErr {
				if err == nil || !strings.Contains(err.Error(), "0000:03:00.1") ||
					!strings.Contains(err.Error(), "nvme") {
					t.Fatalf("CheckIOMMUGroupSafe() = %v, want peer BDF and driver", err)
				}
			} else if err != nil {
				t.Fatalf("CheckIOMMUGroupSafe() = %v, want nil", err)
			}
		})
	}
}

func TestCheckIOMMUGroupSafeRejectsUnknownPeerDriverState(t *testing.T) {
	tmpDir := t.TempDir()
	SetSysfsBase(tmpDir)
	defer ResetSysfsBase()

	bdf := "0000:03:00.0"
	devDir := mkFakeDev(t, tmpDir, bdf)
	peer := "0000:03:00.1"
	peerDir := mkFakeDev(t, tmpDir, peer)
	symlinkIOMMUGroup(t, tmpDir, devDir, 42, peer)
	if err := os.WriteFile(filepath.Join(peerDir, "driver"), []byte("not a symlink"), 0644); err != nil {
		t.Fatal(err)
	}

	err := CheckIOMMUGroupSafe(bdf)
	if err == nil || !strings.Contains(err.Error(), peer) || !strings.Contains(err.Error(), "cannot read driver") {
		t.Fatalf("CheckIOMMUGroupSafe() = %v, want fail-closed driver-state error", err)
	}
}

func TestCheckSafeToBindRejectsMountedSystemDevice(t *testing.T) {
	fakeInitialMountNamespace(t)
	tmpDir := t.TempDir()
	sysRoot := filepath.Join(tmpDir, "sys")
	SetSysfsBase(filepath.Join(sysRoot, "bus", "pci", "devices"))
	defer ResetSysfsBase()

	oldMountInfo, oldDevBlock := mountInfoPath, sysDevBlockBase
	mountInfoPath = filepath.Join(tmpDir, "mountinfo")
	sysDevBlockBase = filepath.Join(sysRoot, "dev", "block")
	defer func() {
		mountInfoPath, sysDevBlockBase = oldMountInfo, oldDevBlock
	}()

	bdf := "0000:03:00.0"
	devDir := mkFakeDev(t, sysfsBase, bdf)
	symlinkIOMMUGroup(t, sysfsBase, devDir, 42)
	blockDir := filepath.Join(devDir, "nvme", "nvme0", "nvme0n1")
	if err := os.MkdirAll(blockDir, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(sysDevBlockBase, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(blockDir, filepath.Join(sysDevBlockBase, "259:0")); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(mountInfoPath,
		[]byte("36 25 259:0 / / rw,relatime - ext4 /dev/nvme0n1 rw\n"), 0644); err != nil {
		t.Fatal(err)
	}

	err := CheckSafeToBind(bdf)
	if err == nil || !strings.Contains(err.Error(), "mounted at /") {
		t.Fatalf("CheckSafeToBind() = %v, want mounted root rejection", err)
	}
}

func TestCheckSafeToBindAllowsUnrelatedMountedDevice(t *testing.T) {
	fakeInitialMountNamespace(t)
	tmpDir := t.TempDir()
	sysRoot := filepath.Join(tmpDir, "sys")
	SetSysfsBase(filepath.Join(sysRoot, "bus", "pci", "devices"))
	defer ResetSysfsBase()

	oldMountInfo, oldDevBlock := mountInfoPath, sysDevBlockBase
	mountInfoPath = filepath.Join(tmpDir, "mountinfo")
	sysDevBlockBase = filepath.Join(sysRoot, "dev", "block")
	defer func() {
		mountInfoPath, sysDevBlockBase = oldMountInfo, oldDevBlock
	}()

	bdf := "0000:03:00.0"
	devDir := mkFakeDev(t, sysfsBase, bdf)
	symlinkIOMMUGroup(t, sysfsBase, devDir, 42)
	otherBlock := filepath.Join(sysRoot, "devices", "virtual", "block", "loop0")
	if err := os.MkdirAll(otherBlock, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(sysDevBlockBase, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(otherBlock, filepath.Join(sysDevBlockBase, "7:0")); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(mountInfoPath,
		[]byte("36 25 7:0 / / rw,relatime - ext4 /dev/loop0 rw\n"+
			"37 25 0:42 / /sys/fs/selinux rw - selinuxfs selinuxfs rw\n"+
			"38 25 0:43 / /var/lib/nfs/rpc_pipefs rw - rpc_pipefs rpc_pipefs rw\n"), 0644); err != nil {
		t.Fatal(err)
	}

	if err := CheckSafeToBind(bdf); err != nil {
		t.Fatalf("CheckSafeToBind() = %v, want unrelated device allowed", err)
	}
}

func TestCheckSafeToBindRejectsMissingBlockMapping(t *testing.T) {
	fakeInitialMountNamespace(t)
	tmpDir := t.TempDir()
	SetSysfsBase(filepath.Join(tmpDir, "sys", "bus", "pci", "devices"))
	defer ResetSysfsBase()

	oldMountInfo, oldDevBlock := mountInfoPath, sysDevBlockBase
	mountInfoPath = filepath.Join(tmpDir, "mountinfo")
	sysDevBlockBase = filepath.Join(tmpDir, "sys", "dev", "block")
	defer func() {
		mountInfoPath, sysDevBlockBase = oldMountInfo, oldDevBlock
	}()

	bdf := "0000:03:00.0"
	devDir := mkFakeDev(t, sysfsBase, bdf)
	symlinkIOMMUGroup(t, sysfsBase, devDir, 42)
	if err := os.WriteFile(mountInfoPath,
		[]byte("36 25 259:2 / / rw,relatime - ext4 /dev/nvme0n1p2 rw\n"), 0644); err != nil {
		t.Fatal(err)
	}

	err := CheckSafeToBind(bdf)
	if err == nil || !strings.Contains(err.Error(), "259:2") {
		t.Fatalf("CheckSafeToBind() = %v, want missing nonzero block mapping error", err)
	}
}

func TestCheckSafeToBindTraversesVirtualPartitionParent(t *testing.T) {
	fakeInitialMountNamespace(t)
	tmpDir := t.TempDir()
	sysRoot := filepath.Join(tmpDir, "sys")
	SetSysfsBase(filepath.Join(sysRoot, "bus", "pci", "devices"))
	defer ResetSysfsBase()

	oldMountInfo, oldDevBlock := mountInfoPath, sysDevBlockBase
	mountInfoPath = filepath.Join(tmpDir, "mountinfo")
	sysDevBlockBase = filepath.Join(sysRoot, "dev", "block")
	defer func() {
		mountInfoPath, sysDevBlockBase = oldMountInfo, oldDevBlock
	}()

	bdf := "0000:03:00.0"
	devDir := mkFakeDev(t, sysfsBase, bdf)
	symlinkIOMMUGroup(t, sysfsBase, devDir, 42)
	physical := filepath.Join(devDir, "nvme", "nvme0", "nvme0n1")
	virtualDisk := filepath.Join(sysRoot, "devices", "virtual", "block", "dm-0")
	partition := filepath.Join(virtualDisk, "dm-0p1")
	if err := os.MkdirAll(filepath.Join(virtualDisk, "slaves"), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(physical, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(partition, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(partition, "partition"), []byte("1\n"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(physical, filepath.Join(virtualDisk, "slaves", "nvme0n1")); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(sysDevBlockBase, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(partition, filepath.Join(sysDevBlockBase, "253:1")); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(mountInfoPath,
		[]byte("36 25 253:1 / / rw,relatime - ext4 /dev/dm-0p1 rw\n"), 0644); err != nil {
		t.Fatal(err)
	}

	err := CheckSafeToBind(bdf)
	if err == nil || !strings.Contains(err.Error(), "mounted at /") {
		t.Fatalf("CheckSafeToBind() = %v, want virtual partition rejection", err)
	}
}

func TestCheckSafeToBindTraversesNVMeMultipath(t *testing.T) {
	fakeInitialMountNamespace(t)
	tmpDir := t.TempDir()
	sysRoot := filepath.Join(tmpDir, "sys")
	SetSysfsBase(filepath.Join(sysRoot, "bus", "pci", "devices"))
	defer ResetSysfsBase()

	oldMountInfo, oldDevBlock := mountInfoPath, sysDevBlockBase
	mountInfoPath = filepath.Join(tmpDir, "mountinfo")
	sysDevBlockBase = filepath.Join(sysRoot, "dev", "block")
	defer func() {
		mountInfoPath, sysDevBlockBase = oldMountInfo, oldDevBlock
	}()

	bdf := "0000:03:00.0"
	devDir := mkFakeDev(t, sysfsBase, bdf)
	symlinkIOMMUGroup(t, sysfsBase, devDir, 42)
	pathDisk := filepath.Join(devDir, "nvme", "nvme0", "nvme0c0n1")
	head := filepath.Join(sysRoot, "devices", "virtual", "nvme-subsystem", "nvme-subsys0", "nvme0n1")
	if err := os.MkdirAll(filepath.Join(head, "multipath"), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(pathDisk, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(pathDisk, filepath.Join(head, "multipath", "nvme0c0n1")); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(sysDevBlockBase, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(head, filepath.Join(sysDevBlockBase, "259:0")); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(mountInfoPath,
		[]byte("36 25 259:0 / / rw,relatime - ext4 /dev/nvme0n1 rw\n"), 0644); err != nil {
		t.Fatal(err)
	}

	err := CheckSafeToBind(bdf)
	if err == nil || !strings.Contains(err.Error(), "mounted at /") {
		t.Fatalf("CheckSafeToBind() = %v, want NVMe multipath rejection", err)
	}
}

func TestCheckSafeToBindRejectsUnresolvedNVMeMultipath(t *testing.T) {
	fakeInitialMountNamespace(t)
	tmpDir := t.TempDir()
	sysRoot := filepath.Join(tmpDir, "sys")
	SetSysfsBase(filepath.Join(sysRoot, "bus", "pci", "devices"))
	defer ResetSysfsBase()

	oldMountInfo, oldDevBlock := mountInfoPath, sysDevBlockBase
	mountInfoPath = filepath.Join(tmpDir, "mountinfo")
	sysDevBlockBase = filepath.Join(sysRoot, "dev", "block")
	defer func() {
		mountInfoPath, sysDevBlockBase = oldMountInfo, oldDevBlock
	}()

	bdf := "0000:03:00.0"
	devDir := mkFakeDev(t, sysfsBase, bdf)
	symlinkIOMMUGroup(t, sysfsBase, devDir, 42)
	head := filepath.Join(sysRoot, "devices", "virtual", "nvme-subsystem", "nvme-subsys0", "nvme0n1")
	if err := os.MkdirAll(head, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(sysDevBlockBase, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(head, filepath.Join(sysDevBlockBase, "259:0")); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(mountInfoPath,
		[]byte("36 25 259:0 / / rw,relatime - ext4 /dev/nvme0n1 rw\n"), 0644); err != nil {
		t.Fatal(err)
	}

	err := CheckSafeToBind(bdf)
	if err == nil || !strings.Contains(err.Error(), "NVMe multipath") {
		t.Fatalf("CheckSafeToBind() = %v, want unresolved NVMe multipath error", err)
	}
}

func TestCheckSafeToBindUsesLegacyNVMeSubsystemLinks(t *testing.T) {
	fakeInitialMountNamespace(t)
	tmpDir := t.TempDir()
	sysRoot := filepath.Join(tmpDir, "sys")
	SetSysfsBase(filepath.Join(sysRoot, "bus", "pci", "devices"))
	defer ResetSysfsBase()

	oldMountInfo, oldDevBlock := mountInfoPath, sysDevBlockBase
	mountInfoPath = filepath.Join(tmpDir, "mountinfo")
	sysDevBlockBase = filepath.Join(sysRoot, "dev", "block")
	defer func() {
		mountInfoPath, sysDevBlockBase = oldMountInfo, oldDevBlock
	}()

	bdf := "0000:03:00.0"
	devDir := mkFakeDev(t, sysfsBase, bdf)
	symlinkIOMMUGroup(t, sysfsBase, devDir, 42)
	controller := filepath.Join(devDir, "nvme", "nvme0")
	subsystem := filepath.Join(sysRoot, "devices", "virtual", "nvme-subsystem", "nvme-subsys0")
	head := filepath.Join(subsystem, "nvme0n1")
	if err := os.MkdirAll(controller, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(head, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(controller, filepath.Join(subsystem, "nvme0")); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(sysDevBlockBase, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(head, filepath.Join(sysDevBlockBase, "259:0")); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(mountInfoPath,
		[]byte("36 25 259:0 / / rw,relatime - ext4 /dev/nvme0n1 rw\n"), 0644); err != nil {
		t.Fatal(err)
	}

	err := CheckSafeToBind(bdf)
	if err == nil || !strings.Contains(err.Error(), "mounted at /") {
		t.Fatalf("CheckSafeToBind() = %v, want legacy NVMe subsystem rejection", err)
	}
}

func TestCheckSafeToBindRejectsMajorZeroStorage(t *testing.T) {
	fakeInitialMountNamespace(t)
	tmpDir := t.TempDir()
	SetSysfsBase(filepath.Join(tmpDir, "sys", "bus", "pci", "devices"))
	defer ResetSysfsBase()

	oldMountInfo := mountInfoPath
	mountInfoPath = filepath.Join(tmpDir, "mountinfo")
	defer func() { mountInfoPath = oldMountInfo }()

	bdf := "0000:03:00.0"
	devDir := mkFakeDev(t, sysfsBase, bdf)
	symlinkIOMMUGroup(t, sysfsBase, devDir, 42)
	if err := os.WriteFile(mountInfoPath,
		[]byte("36 25 0:32 / / rw,relatime - overlay overlay rw\n"), 0644); err != nil {
		t.Fatal(err)
	}

	err := CheckSafeToBind(bdf)
	if err == nil || !strings.Contains(err.Error(), "major 0") {
		t.Fatalf("CheckSafeToBind() = %v, want unverifiable major-0 root rejection", err)
	}
}

func TestCheckInitialMountNamespace(t *testing.T) {
	dir := t.TempDir()
	self := filepath.Join(dir, "self")
	init := filepath.Join(dir, "init")
	if err := os.Symlink("mnt:[1]", self); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink("mnt:[2]", init); err != nil {
		t.Fatal(err)
	}
	if err := checkInitialMountNamespace(self, init); err == nil {
		t.Fatal("checkInitialMountNamespace() accepted a private mount namespace")
	}
}

func TestCheckLiveEnvironment(t *testing.T) {
	oldMountInfo, oldCmdline := mountInfoPath, procCmdlinePath
	defer func() {
		mountInfoPath, procCmdlinePath = oldMountInfo, oldCmdline
	}()

	for _, tc := range []struct {
		name      string
		mountInfo string
		cmdline   string
		wantLive  bool
	}{
		{
			name:      "Ubuntu casper",
			mountInfo: "36 25 0:32 / / rw - overlay overlay rw\n37 25 8:1 / /cdrom ro - iso9660 /dev/sda1 ro\n",
			cmdline:   "quiet splash boot=casper",
			wantLive:  true,
		},
		{
			name:      "Debian live media",
			mountInfo: "36 25 0:32 / / rw - overlay overlay rw\n37 25 8:1 / /run/live/medium ro - squashfs /dev/sda1 ro\n",
			wantLive:  true,
		},
		{
			name:      "container overlay only",
			mountInfo: "36 25 0:32 / / rw - overlay overlay rw\n",
		},
		{
			name:      "installed ext4",
			mountInfo: "36 25 259:2 / / rw - ext4 /dev/nvme0n1p2 rw\n",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			dir := t.TempDir()
			mountInfoPath = filepath.Join(dir, "mountinfo")
			procCmdlinePath = filepath.Join(dir, "cmdline")
			if err := os.WriteFile(mountInfoPath, []byte(tc.mountInfo), 0644); err != nil {
				t.Fatal(err)
			}
			if err := os.WriteFile(procCmdlinePath, []byte(tc.cmdline), 0644); err != nil {
				t.Fatal(err)
			}

			live, _, err := CheckLiveEnvironment()
			if err != nil {
				t.Fatalf("CheckLiveEnvironment() error = %v", err)
			}
			if live != tc.wantLive {
				t.Fatalf("CheckLiveEnvironment() live = %v, want %v", live, tc.wantLive)
			}
		})
	}
}
