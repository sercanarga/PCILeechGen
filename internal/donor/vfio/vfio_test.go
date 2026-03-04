package vfio

import (
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
	// On macOS this should fail (no /sys/kernel/iommu_groups)
	// On Linux without IOMMU it should also fail
	// This test documents the expected behavior on non-Linux
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
	// UnbindFromVFIO writes to sysfs, should fail on invalid/non-Linux
	if err == nil {
		t.Log("UnbindFromVFIO succeeded (unexpected)")
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
