package main

import (
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

func TestLoadDonorContextSupportsOfflineBuildOnWindows(t *testing.T) {
	if runtime.GOOS != "windows" {
		t.Skip("Windows offline build contract")
	}
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	cs.WriteU16(0x00, 0x10ee)
	cs.WriteU16(0x02, 0x7024)
	fixture := &donor.DeviceContext{
		Device: pci.PCIDevice{
			VendorID:  0x10ee,
			DeviceID:  0x7024,
			ClassCode: 0x058000,
		},
		ConfigSpace: cs,
	}
	path := filepath.Join(t.TempDir(), "offline donor.json")
	if err := donor.SaveContext(fixture, path); err != nil {
		t.Fatalf("save offline donor context: %v", err)
	}
	oldOpts := buildOpts
	t.Cleanup(func() { buildOpts = oldOpts })
	buildOpts = buildFlags{fromJSON: path, bdf: "not-a-bdf"}

	got, err := loadDonorContext()
	if err != nil {
		t.Fatalf("loadDonorContext returned error for --from-json on Windows: %v", err)
	}
	if got.Device.VendorID != fixture.Device.VendorID || got.Device.DeviceID != fixture.Device.DeviceID {
		t.Fatalf("loaded donor identity = %04x:%04x, want %04x:%04x", got.Device.VendorID, got.Device.DeviceID, fixture.Device.VendorID, fixture.Device.DeviceID)
	}
	if got.ConfigSpace.ReadU32(0) != fixture.ConfigSpace.ReadU32(0) {
		t.Fatalf("loaded config dword 0 = %08x, want %08x", got.ConfigSpace.ReadU32(0), fixture.ConfigSpace.ReadU32(0))
	}
}

func TestLoadDonorContextRejectsLiveCollectionOnWindows(t *testing.T) {
	if runtime.GOOS != "windows" {
		t.Skip("Windows unsupported live collection contract")
	}
	oldOpts := buildOpts
	t.Cleanup(func() { buildOpts = oldOpts })
	buildOpts = buildFlags{bdf: "not-a-bdf"}

	_, err := loadDonorContext()
	if err == nil {
		t.Fatal("loadDonorContext accepted live collection on Windows")
	}
	want := "live donor collection is unsupported on windows; use --from-json"
	if !strings.Contains(strings.ToLower(err.Error()), want) {
		t.Fatalf("loadDonorContext error = %q, want actionable platform error containing %q", err, want)
	}
	if strings.Contains(err.Error(), "invalid BDF") {
		t.Fatalf("loadDonorContext validated Linux-only BDF input before rejecting the unsupported platform: %v", err)
	}
}

func TestLoadMMIOTraceRejectsLiveCaptureOnWindows(t *testing.T) {
	if runtime.GOOS != "windows" {
		t.Skip("Windows unsupported live MMIO contract")
	}
	_, err := loadMMIOTrace(mmioTraceOptions{
		bdf:      "0000:03:00.0",
		duration: time.Second,
		barIndex: 0,
		barSize:  4096,
	}, 0)
	if err == nil {
		t.Fatal("loadMMIOTrace accepted live capture on Windows")
	}
	want := "live mmio tracing is unsupported on windows; use --trace-file"
	if !strings.Contains(strings.ToLower(err.Error()), want) {
		t.Fatalf("loadMMIOTrace error = %q, want actionable platform error containing %q", err, want)
	}
}
