package donor

import (
	"strings"
	"testing"
	"time"

	"github.com/sercanarga/pcileechgen/internal/donor/mmio"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

func TestDeviceContextJSON_RoundTripsMMIOTraces(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	ctx := &DeviceContext{
		ConfigSpace: cs,
		MMIOTraces: map[int]*mmio.TraceResult{
			2: {
				BDF:      "0000:03:00.0",
				BARIndex: 2,
				BARSize:  4096,
				Duration: 5 * time.Millisecond,
				Records:  []mmio.AccessRecord{{Offset: 0x1C, Type: mmio.AccessRead, Value: 1}},
			},
		},
	}

	data, err := ctx.ToJSON()
	if err != nil {
		t.Fatalf("ToJSON failed: %v", err)
	}
	if !strings.Contains(string(data), "mmio_traces") {
		t.Fatalf("serialized context missing mmio_traces: %s", data)
	}

	loaded, err := FromJSON(data)
	if err != nil {
		t.Fatalf("FromJSON failed: %v", err)
	}
	trace := loaded.MMIOTraces[2]
	if trace == nil {
		t.Fatal("BAR2 trace missing after round trip")
	}
	if trace.BARIndex != 2 || len(trace.Records) != 1 || trace.Records[0].Offset != 0x1C {
		t.Fatalf("trace did not round-trip: %#v", trace)
	}
}
