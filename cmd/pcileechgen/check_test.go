package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/pci"
)

func TestCheckDriverWarnsLiveMMIOTraceWithoutBlockingNormalBuild(t *testing.T) {
	var out bytes.Buffer
	c := &checker{
		w: &out,
		dev: &pci.PCIDevice{
			ClassCode: 0x010601,
			Driver:    "ahci",
		},
	}

	c.checkDriver()

	if c.issues != 0 {
		t.Fatalf("checkDriver issues = %d, want 0", c.issues)
	}
	text := out.String()
	if !strings.Contains(text, "Live mmiotrace profiling blocked") {
		t.Fatalf("warning missing from output: %q", text)
	}
	if strings.Contains(text, "Live donor profiling blocked") {
		t.Fatalf("stale blocking wording in output: %q", text)
	}
}

// TestCheckProfileExpectationWarnsButDoesNotBlock proves the warning-only
// profile validator surfaces mismatches via the check path WITHOUT
// incrementing c.issues (non-blocking, exit-0 semantics).
func TestCheckProfileExpectationWarnsButDoesNotBlock(t *testing.T) {
	var out bytes.Buffer
	c := &checker{
		w: &out,
		dev: &pci.PCIDevice{
			VendorID:  0x8086,
			DeviceID:  0xABCD,
			ClassCode: 0x010802, // NVMe
		},
		cs:   buildMinimalConfigSpace(t),
		bars: []pci.BAR{{Index: 0, Type: pci.BARTypeMem32, Size: 512}},
	}

	c.showProfileExpectation()

	if c.issues != 0 {
		t.Fatalf("showProfileExpectation incremented issues to %d; must stay 0 (non-blocking)", c.issues)
	}
	text := out.String()
	if !strings.Contains(text, "profile.bar.type") {
		t.Fatalf("expected profile.bar.type warning in output: %q", text)
	}
	if !strings.Contains(text, "Profile:") {
		t.Fatalf("expected Profile: prefix in output: %q", text)
	}
}

func buildMinimalConfigSpace(t *testing.T) *pci.ConfigSpace {
	t.Helper()
	cs := pci.NewConfigSpace()
	// PM cap at 0x40, MSI-X at 0x50 (omitted to force msix.missing), PCIe at 0x70.
	cs.WriteU32(0x40, 0x00017001) // PM cap, next=0x70 -> set later
	return cs
}
