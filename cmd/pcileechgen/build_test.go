package main

import (
	"testing"
	"time"

	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/donor/mmio"
	"github.com/sercanarga/pcileechgen/internal/firmware"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

func TestMaybeCaptureMMIOTrace_AttachesOverlay(t *testing.T) {
	old := liveTraceFn
	defer func() { liveTraceFn = old }()
	liveTraceFn = func(string, time.Duration) (*mmio.TraceResult, error) {
		return &mmio.TraceResult{Records: []mmio.AccessRecord{
			{Offset: 0x10, Type: mmio.AccessRead, Value: 0xDEADBEEF},
			{Offset: 0x10, Type: mmio.AccessRead, Value: 0x00000002},
		}}, nil
	}

	ctx := &donor.DeviceContext{}
	if err := maybeCaptureMMIOTrace(ctx, "0000:03:00.0", 2*time.Second); err != nil {
		t.Fatal(err)
	}
	if ctx.BARTraceOverlays[0] == nil {
		t.Fatal("expected BAR0 trace overlay")
	}
}

func TestMaybeCaptureMMIOTrace_UsesLargestBARIndex(t *testing.T) {
	oldTrace := liveTraceFn
	oldBar := buildOpts.mmiotraceBar
	defer func() {
		liveTraceFn = oldTrace
		buildOpts.mmiotraceBar = oldBar
	}()
	liveTraceFn = func(string, time.Duration) (*mmio.TraceResult, error) {
		return &mmio.TraceResult{Records: []mmio.AccessRecord{
			{Offset: 0x10, Type: mmio.AccessRead, Value: 0xDEADBEEF},
			{Offset: 0x10, Type: mmio.AccessRead, Value: 0x00000002},
		}}, nil
	}
	buildOpts.mmiotraceBar = -1

	bar0 := make([]byte, 4096)
	bar2 := make([]byte, 4096)
	bar2[0] = 0x22
	ctx := &donor.DeviceContext{BARContents: map[int][]byte{0: bar0, 2: bar2}}
	if err := maybeCaptureMMIOTrace(ctx, "0000:03:00.0", 2*time.Second); err != nil {
		t.Fatal(err)
	}
	if ctx.BARTraceOverlays[firmware.LargestBarIndex(ctx.BARContents)] == nil {
		t.Fatal("expected overlay on largest BAR index")
	}
}

func TestMaybeCaptureMMIOTrace_MapsPhysicalAddressToBAROffset(t *testing.T) {
	oldTrace := liveTraceFn
	oldBar := buildOpts.mmiotraceBar
	defer func() {
		liveTraceFn = oldTrace
		buildOpts.mmiotraceBar = oldBar
	}()
	liveTraceFn = func(string, time.Duration) (*mmio.TraceResult, error) {
		return &mmio.TraceResult{Records: []mmio.AccessRecord{
			{Address: 0x80004010, Offset: 0x10, Type: mmio.AccessRead, Value: 0xDEADBEEF},
			{Address: 0x80004010, Offset: 0x10, Type: mmio.AccessRead, Value: 0x00000002},
		}}, nil
	}
	buildOpts.mmiotraceBar = 0

	ctx := &donor.DeviceContext{
		BARs:        []pci.BAR{{Index: 0, Address: 0x80000000, Size: 0x8000}},
		BARContents: map[int][]byte{0: make([]byte, 0x8000)},
	}
	if err := maybeCaptureMMIOTrace(ctx, "0000:03:00.0", 2*time.Second); err != nil {
		t.Fatal(err)
	}
	if got := ctx.BARTraceOverlays[0].Sequential[0x4010]; len(got) != 2 {
		t.Fatalf("expected trace at BAR offset 0x4010, got %#v", ctx.BARTraceOverlays[0].Sequential)
	}
}

func TestUnsafeLiveDonorReason_ActiveWiFi(t *testing.T) {
	dev := &pci.PCIDevice{ClassCode: 0x028000, Driver: "iwlwifi"}
	if reason := unsafeLiveDonorReason(dev); reason == "" {
		t.Fatal("expected active Wi-Fi donor to be blocked")
	}
}

func TestUnsafeLiveDonorReason_VFIOAllowed(t *testing.T) {
	dev := &pci.PCIDevice{ClassCode: 0x028000, Driver: "vfio-pci"}
	if reason := unsafeLiveDonorReason(dev); reason != "" {
		t.Fatalf("vfio-bound donor should be allowed, got %q", reason)
	}
}

func TestUnsafeLiveDonorReason_ActiveAHCI(t *testing.T) {
	dev := &pci.PCIDevice{ClassCode: 0x010601, Driver: "ahci"}
	if reason := unsafeLiveDonorReason(dev); reason == "" {
		t.Fatal("expected active AHCI donor to be blocked for invasive tracing")
	}
}

func TestApplyPublicSafeMode_StripsSensitiveIdentity(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	dsnHeader := uint32(pci.ExtCapIDDeviceSerialNumber) | (1 << 16)
	cs.WriteU32(0x100, dsnHeader)
	cs.WriteU32(0x104, 0x11223344)
	cs.WriteU32(0x108, 0xAABBCCDD)

	ctx := &donor.DeviceContext{
		Hostname:    "lab-host",
		Device:      pci.PCIDevice{BDF: pci.BDF{Domain: 1, Bus: 2, Device: 3, Function: 4}},
		ConfigSpace: cs,
		ExtCapabilities: []pci.ExtCapability{
			{ID: pci.ExtCapIDDeviceSerialNumber, Offset: 0x100, Data: make([]byte, 12)},
		},
		NVMeIdentify: &donor.NVMeIdentifyCapture{HasController: true},
	}

	applyPublicSafeMode(ctx)

	ids := firmware.ExtractDeviceIDs(ctx.ConfigSpace, ctx.ExtCapabilities)
	if ids.HasDSN {
		t.Fatal("public-safe mode should strip DSN")
	}
	if ctx.NVMeIdentify != nil {
		t.Fatal("public-safe mode should drop donor-captured NVMe identify")
	}
	if ctx.Hostname != "" || ctx.Device.BDF != (pci.BDF{}) {
		t.Fatal("public-safe mode should clear host-local identifiers")
	}
}
