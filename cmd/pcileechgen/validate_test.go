package main

import (
	"os"
	"path/filepath"
	"slices"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/firmware/codegen"
	"github.com/sercanarga/pcileechgen/internal/firmware/output"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

func TestValidateCOEFiles_ConfigSpaceMismatchFails(t *testing.T) {
	dir := t.TempDir()
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	cs.WriteU32(0x10, 0xFFFFF000)

	if err := os.WriteFile(filepath.Join(dir, "pcileech_cfgspace.coe"), []byte("bad cfgspace\n"), 0o644); err != nil {
		t.Fatalf("write cfgspace fixture: %v", err)
	}
	if err := os.WriteFile(filepath.Join(dir, "pcileech_cfgspace_writemask.coe"), []byte(codegen.GenerateWritemaskCOE(cs)), 0o644); err != nil {
		t.Fatalf("write writemask fixture: %v", err)
	}

	v := &validator{
		outputDir: dir,
		scrubbed:  cs,
		result:    &output.ValidationResult{},
	}

	v.validateCOEFiles()

	if !slices.Contains(v.result.Failed, "pcileech_cfgspace.coe MISMATCH") {
		t.Fatalf("failed validations = %v, want cfgspace mismatch failure", v.result.Failed)
	}
	if len(v.result.Warnings) != 0 {
		t.Fatalf("warnings = %v, want none", v.result.Warnings)
	}
}

func TestValidateCapabilityChainsFailsMalformedDonor(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.WriteU16(0x06, 0x0010)
	cs.WriteU8(0x34, 0x40)
	cs.WriteU8(0x40, pci.CapIDPowerManagement)
	cs.WriteU8(0x41, 0x40)
	v := &validator{
		ctx:    &donor.DeviceContext{ConfigSpace: cs},
		result: &output.ValidationResult{},
	}

	v.validateCapabilityChains()

	if !slices.Contains(v.result.Failed, "standard capability loop at 0x040") {
		t.Fatalf("failed validations = %v", v.result.Failed)
	}
}
