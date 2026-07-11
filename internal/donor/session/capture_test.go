package session

import (
	"testing"
	"time"

	"github.com/sercanarga/pcileechgen/internal/donor/mmio"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

type fakeCaptureSource struct {
	configReads int
}

func (f *fakeCaptureSource) Device() (*pci.PCIDevice, error) {
	return &pci.PCIDevice{BDF: pci.BDF{Bus: 2}, VendorID: 0x10ec, DeviceID: 0x8168}, nil
}
func (f *fakeCaptureSource) Config() (*pci.ConfigSpace, error) {
	f.configReads++
	cs := pci.NewConfigSpace()
	cs.WriteU32(4, uint32(f.configReads))
	return cs, nil
}
func (f *fakeCaptureSource) BARs() ([]pci.BAR, error) {
	return []pci.BAR{{Index: 0, Type: pci.BARTypeMem32, Size: 4096}}, nil
}
func (f *fakeCaptureSource) Resources() ([]byte, error)  { return []byte("resources\n"), nil }
func (f *fakeCaptureSource) Interrupts() ([]byte, error) { return []byte("interrupts\n"), nil }
func (f *fakeCaptureSource) Trace() (*mmio.TraceResult, error) {
	return &mmio.TraceResult{BDF: "0000:02:00.0", Duration: time.Millisecond, Records: []mmio.AccessRecord{{Offset: 4, Width: 4, ByteEnable: 0xf, Type: mmio.AccessWrite, Value: 1}}}, nil
}
func (f *fakeCaptureSource) Driver() string { return "r8169" }
func (f *fakeCaptureSource) Kernel() string { return "test-kernel" }

func TestCaptureWritesBeforeAfterTraceAndMetadata(t *testing.T) {
	dir := t.TempDir()
	manifest, err := Capture(dir, ScenarioTrace, &fakeCaptureSource{})
	if err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"config-before.bin", "config-after.bin", "resources.txt", "interrupts-before.txt", "interrupts-after.txt", "trace.json"} {
		if _, ok := manifest.Artifacts[name]; !ok {
			t.Errorf("missing artifact %q", name)
		}
	}
	if manifest.Driver != "r8169" || manifest.Kernel != "test-kernel" || manifest.Device.VendorID != 0x10ec {
		t.Fatalf("manifest metadata = %+v", manifest)
	}
	if verifyErr := Verify(dir, manifest); verifyErr != nil {
		t.Fatal(verifyErr)
	}
	loaded, err := LoadCapture(dir)
	if err != nil {
		t.Fatal(err)
	}
	if loaded.Trace == nil || len(loaded.Trace.Records) != 1 || loaded.ConfigBefore.ReadU32(4) != 1 || loaded.ConfigAfter.ReadU32(4) != 2 {
		t.Fatalf("loaded capture = %+v", loaded)
	}
}
