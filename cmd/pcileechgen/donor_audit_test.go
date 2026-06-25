package main

import (
	"bytes"
	"encoding/json"
	"path/filepath"
	"testing"
	"time"

	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/donor/mmio"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

func TestAssessDonorContext_HappyPath(t *testing.T) {
	// Given
	ctx := makeDonorContextFixture(t, nil)

	// When
	report := assessDonorContext(ctx, "device_context.json")

	// Then
	if report.Status != "READY" {
		t.Fatalf("Status = %q, want READY", report.Status)
	}
	if report.MemoryBars != 1 {
		t.Fatalf("MemoryBars = %d, want 1", report.MemoryBars)
	}
	if report.Score != 100 {
		t.Fatalf("Score = %d, want 100", report.Score)
	}
	if len(report.Blockers) != 0 {
		t.Fatalf("unexpected blockers: %#v", report.Blockers)
	}
}

func TestAssessDonorContext_MissingMemoryBAR(t *testing.T) {
	// Given
	ctx := makeDonorContextFixture(t, func(context *donor.DeviceContext) {
		context.BARs = []pci.BAR{{Index: 0, Type: pci.BARTypeIO, Size: 0x1000, Address: 0xC000}}
	})

	// When
	report := assessDonorContext(ctx, "device_context.json")

	// Then
	if report.Status != "BLOCKED" {
		t.Fatalf("Status = %q, want BLOCKED", report.Status)
	}
	if report.MemoryBars != 0 {
		t.Fatalf("MemoryBars = %d, want 0", report.MemoryBars)
	}
	if len(report.Blockers) == 0 {
		t.Fatal("expected at least one blocker")
	}
}

func TestDonorAuditCmd_JSONOutput(t *testing.T) {
	// Given
	path := writeDonorAuditFixture(t, nil)
	resetDonorAuditFlags()
	var buf bytes.Buffer

	// When
	rootCmd.SetArgs([]string{"donor-audit", "--context", path, "--json"})
	rootCmd.SetOut(&buf)
	rootCmd.SetErr(&buf)
	err := rootCmd.Execute()

	// Then
	if err != nil {
		t.Fatalf("donor-audit failed: %v\n%s", err, buf.String())
	}

	var report donorAuditReport
	if err := json.Unmarshal(buf.Bytes(), &report); err != nil {
		t.Fatalf("expected JSON output, failed to unmarshal: %v", err)
	}
	if report.Status != "READY" && report.Status != "AT_RISK" {
		t.Fatalf("Status = %q", report.Status)
	}
	if report.ContextPath != path {
		t.Fatalf("ContextPath = %q, want %q", report.ContextPath, path)
	}
}

func TestRunDonorAudit_ReportsMissingContextPath(t *testing.T) {
	resetDonorAuditFlags()
	if _, err := runDonorAudit(donorAuditOptions{}); err == nil {
		t.Fatal("expected error when context path is missing")
	}
}

func writeDonorAuditFixture(t *testing.T, configure func(context *donor.DeviceContext)) string {
	t.Helper()
	ctx := makeDonorContextFixture(t, configure)
	path := filepath.Join(t.TempDir(), "device_context.json")
	if err := donor.SaveContext(ctx, path); err != nil {
		t.Fatalf("write fixture: %v", err)
	}
	return path
}

func makeDonorContextFixture(t *testing.T, configure func(context *donor.DeviceContext)) *donor.DeviceContext {
	t.Helper()
	cs := pci.NewConfigSpace()
	cs.WriteU16(0x00, 0x10ec)
	cs.WriteU16(0x02, 0x8168)

	ctx := &donor.DeviceContext{
		Device: pci.PCIDevice{
			BDF:        pci.BDF{Bus: 3, Device: 0, Function: 0},
			VendorID:   0x10ec,
			DeviceID:   0x8168,
			ClassCode:  0x020000,
			RevisionID: 0x01,
		},
		ConfigSpace: cs,
		BARs: []pci.BAR{{
			Index:   0,
			Type:    pci.BARTypeMem32,
			Address: 0x80000000,
			Size:    4096,
		}},
		BARContents: map[int][]byte{
			0: make([]byte, 4096),
		},
		BARProfiles: map[int]*donor.BARProfile{
			0: {BarIndex: 0, Size: 4096, Probes: []donor.BARProbeResult{{Offset: 0, Original: 0, RWMask: 0xFFFFFFFF, MaybeRW1C: false}}},
		},
		MMIOTraces: map[int]*mmio.TraceResult{
			0: {
				BDF:       "0000:03:00.0",
				BARIndex:  0,
				BARSize:   4096,
				Duration:  time.Millisecond,
				Records:   []mmio.AccessRecord{{Offset: 0, Type: mmio.AccessRead, Value: 1, Timestamp: time.Millisecond}},
				StartTime: time.Now().UTC().Add(-time.Millisecond),
			},
		},
		Capabilities:    []pci.Capability{{ID: pci.CapIDMSIX, Offset: 0x50}},
		ExtCapabilities: []pci.ExtCapability{{ID: pci.ExtCapIDDeviceSerialNumber, Offset: 0x100}},
		MSIXData:        &donor.MSIXData{TableSize: 1, TableBIR: 1, TableOffset: 0x2000, PBABIR: 0, PBAOffset: 0x3000},
	}

	if configure != nil {
		configure(ctx)
	}
	return ctx
}

func resetDonorAuditFlags() {
	donorAuditOpts = donorAuditOptions{}
}
