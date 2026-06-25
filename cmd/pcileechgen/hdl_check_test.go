package main

import (
	"testing"
)

func TestBuildHDLCheckCommand(t *testing.T) {
	got := buildHDLCheckCommand(hdlCheckOptions{dir: "out", dryRun: true})
	want := []string{"run", "./tools/hdlcheck", "--dir", "out", "--dry-run"}
	if len(got) != len(want) {
		t.Fatalf("len(got) = %d, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("arg[%d] = %q, want %q", i, got[i], want[i])
		}
	}
}

func TestBuildHDLCheckCommandNonDryRun(t *testing.T) {
	got := buildHDLCheckCommand(hdlCheckOptions{dir: "out"})
	if len(got) != 4 {
		t.Fatalf("len(got) = %d, want 4", len(got))
	}
	if got[0] != "run" || got[1] != "./tools/hdlcheck" || got[2] != "--dir" || got[3] != "out" {
		t.Fatalf("unexpected args: %#v", got)
	}
}
