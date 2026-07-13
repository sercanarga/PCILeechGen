package synthetic

import "testing"

func TestBuildDemoProfiles(t *testing.T) {
	want := map[string]uint32{
		"DiskTest":       0x010802,
		"IntelI210":      0x020000,
		"IntelI219":      0x020000,
		"IntelI225":      0x020000,
		"NICv2":          0x020000,
		"NVMEv2":         0x010802,
		"RTL8125":        0x020000,
		"RealtekRTL8125": 0x020000,
	}
	names := DemoProfileNames()
	if len(names) != len(want) {
		t.Fatalf("DemoProfileNames returned %d names, want %d", len(names), len(want))
	}
	for _, name := range names {
		ctx := BuildDemoProfile(name)
		if ctx == nil {
			t.Fatalf("BuildDemoProfile(%q) returned nil", name)
		}
		if ctx.Device.ClassCode != want[name] {
			t.Fatalf("%s class code = 0x%06x, want 0x%06x", name, ctx.Device.ClassCode, want[name])
		}
	}
}
