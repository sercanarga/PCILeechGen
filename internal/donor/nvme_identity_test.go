package donor

import (
	"testing"

	"github.com/sercanarga/pcileechgen/internal/pci"
)

func TestDeviceContext_NVMeIdentity_RoundTrip(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = 64
	cs.WriteU32(0, 0x12345678)

	ctx := &DeviceContext{
		ConfigSpace: cs,
		NVMeIdentity: &NVMeIdentity{
			Serial: "S6PXNG0T12345678W",
			Model:  "Samsung SSD 990 PRO 2TB",
			FWRev:  "4B2QJXD7",
		},
	}

	data, err := ctx.ToJSON()
	if err != nil {
		t.Fatalf("ToJSON failed: %v", err)
	}

	loaded, err := FromJSON(data)
	if err != nil {
		t.Fatalf("FromJSON failed: %v", err)
	}
	if loaded.NVMeIdentity == nil {
		t.Fatal("NVMeIdentity lost across JSON round-trip")
	}
	if loaded.NVMeIdentity.Serial != "S6PXNG0T12345678W" {
		t.Errorf("Serial: got %q", loaded.NVMeIdentity.Serial)
	}
	if loaded.NVMeIdentity.Model != "Samsung SSD 990 PRO 2TB" {
		t.Errorf("Model: got %q", loaded.NVMeIdentity.Model)
	}
	if loaded.NVMeIdentity.FWRev != "4B2QJXD7" {
		t.Errorf("FWRev: got %q", loaded.NVMeIdentity.FWRev)
	}
}

func TestDeviceContext_NVMeIdentity_NilOmitted(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = 64
	cs.WriteU32(0, 0x12345678)

	ctx := &DeviceContext{ConfigSpace: cs}
	data, err := ctx.ToJSON()
	if err != nil {
		t.Fatalf("ToJSON failed: %v", err)
	}
	loaded, err := FromJSON(data)
	if err != nil {
		t.Fatalf("FromJSON failed: %v", err)
	}
	if loaded.NVMeIdentity != nil {
		t.Errorf("NVMeIdentity should be nil when absent, got %+v", loaded.NVMeIdentity)
	}
}
