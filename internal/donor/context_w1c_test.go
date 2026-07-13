package donor

import (
	"strings"
	"testing"
	"time"

	"github.com/sercanarga/pcileechgen/internal/pci"
)

func TestDeviceContext_W1CMaskRoundTrip(t *testing.T) {
	ctx := &DeviceContext{
		CollectedAt: time.Now(),
		ConfigSpace: pci.NewConfigSpace(),
		BARs:        []pci.BAR{{Index: 0, Type: pci.BARTypeMem32, Size: 4096}},
		BARProfiles: map[int]*BARProfile{
			0: {BarIndex: 0, Size: 4096, Probes: []BARProbeResult{
				{Offset: 0x10, Original: 0x000000FF, RWMask: 0x000000FF, W1CMask: 0x00000010, MaybeRW1C: true},
				{Offset: 0x14, Original: 0x00000000, RWMask: 0x00000000},
			}},
		},
	}
	data, err := ctx.ToJSON()
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	if !strings.Contains(string(data), `"w1c_mask": 16`) {
		t.Errorf("W1CMask not serialized; output:\n%s", string(data))
	}
	ctx2, err := FromJSON(data)
	if err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	prof := ctx2.BARProfiles[0]
	if prof == nil {
		t.Fatal("BARProfiles[0] missing")
	}
	var w1c uint32
	for _, p := range prof.Probes {
		if p.Offset == 0x10 {
			w1c = p.W1CMask
		}
	}
	if w1c != 0x00000010 {
		t.Errorf("W1CMask round-trip: got 0x%08X, want 0x00000010", w1c)
	}
}
