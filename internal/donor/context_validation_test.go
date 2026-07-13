package donor

import (
	"bytes"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/pci"
)

func validContextJSON(t *testing.T) []byte {
	t.Helper()
	data, err := (&DeviceContext{ConfigSpace: pci.NewConfigSpace()}).ToJSON()
	if err != nil {
		t.Fatal(err)
	}
	return data
}

func TestFromJSONRejectsUnknownFields(t *testing.T) {
	data := bytes.TrimSuffix(validContextJSON(t), []byte("}"))
	data = append(data, []byte(`,"unknown_field":true}`)...)
	if _, err := FromJSON(data); err == nil {
		t.Fatal("unknown field was accepted")
	}
}

func TestFromJSONRejectsInvalidConfigSize(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = 64
	data, err := (&DeviceContext{ConfigSpace: cs}).ToJSON()
	if err != nil {
		t.Fatal(err)
	}
	if _, err := FromJSON(data); err == nil {
		t.Fatal("invalid config-space size was accepted")
	}
}

func TestFromJSONRejectsDuplicateBARs(t *testing.T) {
	ctx := &DeviceContext{
		ConfigSpace: pci.NewConfigSpace(),
		BARs: []pci.BAR{
			{Index: 0, Type: pci.BARTypeMem32, Size: 4096},
			{Index: 0, Type: pci.BARTypeMem32, Size: 4096},
		},
	}
	data, err := ctx.ToJSON()
	if err != nil {
		t.Fatal(err)
	}
	if _, err := FromJSON(data); err == nil {
		t.Fatal("duplicate BAR index was accepted")
	}
}

func TestFromJSONRejectsInvalidRawNVMeIdentityLength(t *testing.T) {
	ctx := &DeviceContext{
		ConfigSpace:  pci.NewConfigSpace(),
		NVMeIdentity: &NVMeIdentity{RawControllerIdent: []byte{1}},
	}
	data, err := ctx.ToJSON()
	if err != nil {
		t.Fatal(err)
	}
	if _, err := FromJSON(data); err == nil {
		t.Fatal("invalid raw Identify length was accepted")
	}
}

func FuzzFromJSONContext(f *testing.F) {
	seed, err := (&DeviceContext{ConfigSpace: pci.NewConfigSpace()}).ToJSON()
	if err != nil {
		f.Fatal(err)
	}
	f.Add(seed)
	f.Fuzz(func(t *testing.T, data []byte) {
		if len(data) > 1<<20 {
			t.Skip()
		}
		_, _ = FromJSON(data)
	})
}
