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

func TestDeviceContext_RejectsBadConfigSpaceSize(t *testing.T) {
	// The --from-json path treats config_space_size as untrusted. A declared
	// size that disagrees with the hex data, is negative, or exceeds the 4096B
	// backing array must be rejected rather than yielding an out-of-bounds
	// ConfigSpace (Bytes()/Size-bounded consumers would over-read).
	oversized := `{"config_space_hex":[` +
		strings.TrimRight(strings.Repeat(`"00000000",`, 1025), ",") +
		`],"config_space_size":4100}` // 1025 words = 4100B > 4096
	cases := map[string]string{
		"mismatched": `{"config_space_hex":["12345678"],"config_space_size":256}`,
		"negative":   `{"config_space_hex":["12345678"],"config_space_size":-4}`,
		"oversized":  oversized,
	}
	for name, js := range cases {
		if _, err := FromJSON([]byte(js)); err == nil {
			t.Errorf("%s: expected error, got nil", name)
		}
	}

	// A consistent, in-range size still round-trips.
	if _, err := FromJSON([]byte(`{"config_space_hex":["12345678"],"config_space_size":4}`)); err != nil {
		t.Errorf("valid 4-byte config space rejected: %v", err)
	}
}
