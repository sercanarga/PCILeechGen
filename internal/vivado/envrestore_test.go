package vivado

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestParseProbeOutput_AllKeys(t *testing.T) {
	out := []byte("XILINXD_LICENSE_FILE=/opt/lic.lic\nLM_LICENSE_FILE=2100@host\nPATH=/usr/bin:/bin\n")
	got := parseProbeOutput(out)
	want := envOverride{
		"XILINXD_LICENSE_FILE": "/opt/lic.lic",
		"LM_LICENSE_FILE":      "2100@host",
		"PATH":                 "/usr/bin:/bin",
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("parseProbeOutput = %#v, want %#v", got, want)
	}
}

func TestParseProbeOutput_ValueWithEquals(t *testing.T) {
	out := []byte("XILINXD_LICENSE_FILE=/a/b=port\n")
	got := parseProbeOutput(out)
	if got["XILINXD_LICENSE_FILE"] != "/a/b=port" {
		t.Errorf("value with '=' = %q, want %q", got["XILINXD_LICENSE_FILE"], "/a/b=port")
	}
}

func TestParseProbeOutput_MissingKeyAbsent(t *testing.T) {
	out := []byte("XILINXD_LICENSE_FILE=/opt/lic.lic\n")
	got := parseProbeOutput(out)
	if _, ok := got["LM_LICENSE_FILE"]; ok {
		t.Error("LM_LICENSE_FILE should be absent when not in output")
	}
	if _, ok := got["PATH"]; ok {
		t.Error("PATH should be absent when not in output")
	}
}

func TestParseProbeOutput_IgnoresNoiseAndEmpty(t *testing.T) {
	out := []byte("some profile echo noise\nXILINXD_LICENSE_FILE=/x.lic\nRANDOM_THING=ignored\n\nLM_LICENSE_FILE=\n")
	got := parseProbeOutput(out)
	if got["XILINXD_LICENSE_FILE"] != "/x.lic" {
		t.Errorf("XILINXD_LICENSE_FILE = %q, want /x.lic", got["XILINXD_LICENSE_FILE"])
	}
	if _, ok := got["LM_LICENSE_FILE"]; ok {
		t.Error("empty LM_LICENSE_FILE value should be dropped")
	}
}

func TestSetEnv_ReplaceInPlace(t *testing.T) {
	env := []string{"HOME=/root", "PATH=/usr/bin"}
	got := setEnv(env, "HOME", "/home/user")
	want := []string{"HOME=/home/user", "PATH=/usr/bin"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("setEnv replace = %#v, want %#v", got, want)
	}
}

func TestSetEnv_AppendWhenAbsent(t *testing.T) {
	env := []string{"PATH=/usr/bin"}
	got := setEnv(env, "HOME", "/home/user")
	want := []string{"PATH=/usr/bin", "HOME=/home/user"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("setEnv append = %#v, want %#v", got, want)
	}
}

func TestApplyOverrides_KeepsOthersAndDoesNotTouchXilinx(t *testing.T) {
	env := []string{"XILINX_VIVADO=/tools/Xilinx/Vivado/2023.2", "PATH=/usr/bin"}
	ov := envOverride{"HOME": "/home/user", "XILINXD_LICENSE_FILE": "/x.lic"}
	got := applyOverrides(env, ov)

	if envGet(got, "XILINX_VIVADO") != "/tools/Xilinx/Vivado/2023.2" {
		t.Errorf("XILINX_VIVADO must be untouched, got %q", envGet(got, "XILINX_VIVADO"))
	}
	if envGet(got, "HOME") != "/home/user" {
		t.Errorf("HOME override not applied, got %q", envGet(got, "HOME"))
	}
	if envGet(got, "XILINXD_LICENSE_FILE") != "/x.lic" {
		t.Errorf("license override not applied, got %q", envGet(got, "XILINXD_LICENSE_FILE"))
	}
}

func TestParseSudoID(t *testing.T) {
	cases := []struct {
		in   string
		want uint32
		ok   bool
	}{
		{"", 0, false},
		{"0", 0, false},
		{"abc", 0, false},
		{"-1", 0, false},
		{"4294967296", 0, false}, // 2^32 overflow
		{"1000", 1000, true},
	}
	for _, c := range cases {
		got, ok := parseSudoID(c.in)
		if ok != c.ok || got != c.want {
			t.Errorf("parseSudoID(%q) = (%d,%v), want (%d,%v)", c.in, got, ok, c.want, c.ok)
		}
	}
}

// envGet returns the value for key in a KEY=VALUE slice, or "".
func envGet(env []string, key string) string {
	for _, e := range env {
		if k, v, ok := strings.Cut(e, "="); ok && k == key {
			return v
		}
	}
	return ""
}

// fakeProbe returns a probeFunc emitting the given stdout, or an error.
func fakeProbe(out string, err error) probeFunc {
	return func(context.Context) ([]byte, error) { return []byte(out), err }
}

func TestResolveEnvOverrides_ProbeMerged(t *testing.T) {
	ov := resolveEnvOverrides("/home/u", fakeProbe("XILINXD_LICENSE_FILE=/a.lic\nLM_LICENSE_FILE=2100@h\nPATH=/u/bin\n", nil))
	if ov["HOME"] != "/home/u" {
		t.Errorf("HOME = %q, want /home/u", ov["HOME"])
	}
	if ov["XILINXD_LICENSE_FILE"] != "/a.lic" {
		t.Errorf("XILINXD_LICENSE_FILE = %q, want /a.lic", ov["XILINXD_LICENSE_FILE"])
	}
	if ov["LM_LICENSE_FILE"] != "2100@h" {
		t.Errorf("LM_LICENSE_FILE = %q, want 2100@h", ov["LM_LICENSE_FILE"])
	}
	if ov["PATH"] != "/u/bin" {
		t.Errorf("PATH = %q, want /u/bin", ov["PATH"])
	}
}

func TestResolveEnvOverrides_ProbeError_HomeOnly(t *testing.T) {
	ov := resolveEnvOverrides("/home/u", fakeProbe("", errors.New("boom")))
	if len(ov) != 1 || ov["HOME"] != "/home/u" {
		t.Fatalf("want HOME-only override on probe error, got %#v", ov)
	}
}

func TestResolveEnvOverrides_NilProbe_HomeOnly(t *testing.T) {
	ov := resolveEnvOverrides("/home/u", nil)
	if len(ov) != 1 || ov["HOME"] != "/home/u" {
		t.Fatalf("want HOME-only override when probe is nil, got %#v", ov)
	}
}

func TestBuildSudoOverrides_HomeUnresolved_HardError(t *testing.T) {
	_, err := buildSudoOverrides(1000, 1000, "", "/bin/bash", nil)
	if !errors.Is(err, errHomeUnresolved) {
		t.Fatalf("want errHomeUnresolved, got %v", err)
	}
}

func TestBuildSudoOverrides_Delegates(t *testing.T) {
	ov, err := buildSudoOverrides(1000, 1000, "/home/u", "/bin/bash", fakeProbe("XILINXD_LICENSE_FILE=/a.lic\n", nil))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if ov["XILINXD_LICENSE_FILE"] != "/a.lic" {
		t.Errorf("probed license lost through buildSudoOverrides: %q", ov["XILINXD_LICENSE_FILE"])
	}
	if ov["HOME"] != "/home/u" {
		t.Errorf("HOME lost: %q", ov["HOME"])
	}
}

func TestEnvOverrides_NotUnderSudo_Noop(t *testing.T) {
	// Not root in tests, so this must be a no-op.
	ov, err := envOverrides()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if ov != nil {
		t.Errorf("want nil override when not under sudo, got %#v", ov)
	}
}

func TestChownOutputs_NotUnderSudo_Noop(t *testing.T) {
	// Not root -> chownOutputs must be a no-op and never touch files.
	dir := t.TempDir()
	bit := filepath.Join(dir, "out.bit")
	if err := os.WriteFile(bit, []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := chownOutputs(dir); err != nil {
		t.Fatalf("chownOutputs returned error when not under sudo: %v", err)
	}
	if _, err := os.Stat(bit); err != nil {
		t.Errorf("output file should be untouched, got %v", err)
	}
}

func TestChownOutputs_NoMatches_NoError(t *testing.T) {
	// Empty dir, no .bit/.bin -> no error, nothing to chown.
	if err := chownOutputs(t.TempDir()); err != nil {
		t.Fatalf("chownOutputs on empty dir returned error: %v", err)
	}
}

func TestChownOutputs_MissingDir_NoError(t *testing.T) {
	// Best-effort: a missing output dir must not produce an error.
	if err := chownOutputs(filepath.Join(t.TempDir(), "does-not-exist")); err != nil {
		t.Fatalf("chownOutputs on missing dir returned error: %v", err)
	}
}
