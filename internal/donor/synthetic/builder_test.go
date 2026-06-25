package synthetic

import (
	"encoding/json"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/firmware/devclass"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

func TestBuild_RoundTrip(t *testing.T) {
	for _, class := range devclass.AllClasses() {
		t.Run(class, func(t *testing.T) {
			ctx := Build(class)
			if ctx == nil {
				t.Fatalf("Build(%q) returned nil", class)
			}
			data, err := json.Marshal(ctx)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			got, err := donor.FromJSON(data)
			if err != nil {
				t.Fatalf("FromJSON: %v", err)
			}
			if got.Device.ClassCode != ctx.Device.ClassCode {
				t.Errorf("class code lost in round-trip: got %#x want %#x",
					got.Device.ClassCode, ctx.Device.ClassCode)
			}
			if len(got.BARs) == 0 {
				t.Error("no BARs after round-trip; BAR0 needed for BAR model build")
			}
			if got.Device.VendorID != ctx.Device.VendorID {
				t.Errorf("vendor ID lost in round-trip: got %#x want %#x",
					got.Device.VendorID, ctx.Device.VendorID)
			}
			if len(got.BARs) > 0 && got.BARs[0].Size != ctx.BARs[0].Size {
				t.Errorf("BAR0 size lost in round-trip: got %d want %d",
					got.BARs[0].Size, ctx.BARs[0].Size)
			}
		})
	}
}

func TestBuild_UnknownClassReturnsNil(t *testing.T) {
	if Build("does-not-exist") != nil {
		t.Error("Build with unknown class should return nil")
	}
}

func TestBuild_HasPCIeCapability(t *testing.T) {
	for _, class := range devclass.AllClasses() {
		t.Run(class, func(t *testing.T) {
			ctx := Build(class)
			if ctx == nil {
				t.Fatalf("Build(%q) returned nil", class)
			}
			caps := pci.ParseCapabilities(ctx.ConfigSpace)
			found := false
			for _, c := range caps {
				if c.ID == pci.CapIDPCIExpress {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("%s: config space has no PCIe capability after parse", class)
			}
		})
	}
}
