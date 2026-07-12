package vivado

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"maps"
	"math"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type probeFunc func(ctx context.Context) ([]byte, error)

type envOverride map[string]string

// errHomeUnresolved is returned when the invoking user's home can't be found
// under sudo. A wrong HOME would send Vivado to the wrong ~/.Xilinx.
var errHomeUnresolved = errors.New("could not determine invoking user home under sudo; use 'sudo -E' or export XILINXD_LICENSE_FILE in your shell profile")

// licenseProbeVars are the env vars recovered from the invoking user's login shell.
var licenseProbeVars = map[string]struct{}{
	"XILINXD_LICENSE_FILE": {},
	"LM_LICENSE_FILE":      {},
	"PATH":                 {},
}

// parseProbeOutput extracts the license vars from the probe shell's stdout.
// Unknown lines are ignored, the first match per var wins, and empty values are
// skipped.
func parseProbeOutput(out []byte) envOverride {
	res := envOverride{}
	for _, line := range strings.Split(string(out), "\n") {
		key, val, ok := strings.Cut(line, "=")
		if !ok {
			continue
		}
		if _, want := licenseProbeVars[key]; !want {
			continue
		}
		if _, dup := res[key]; dup || val == "" {
			continue
		}
		res[key] = val
	}
	return res
}

// setEnv sets key in a KEY=VALUE slice, replacing an existing entry in place or
// appending. The slice may be mutated.
func setEnv(env []string, key, val string) []string {
	for i, e := range env {
		if k, _, ok := strings.Cut(e, "="); ok && k == key {
			env[i] = key + "=" + val
			return env
		}
	}
	return append(env, key+"="+val)
}

func applyOverrides(env []string, ov envOverride) []string {
	for k, v := range ov {
		env = setEnv(env, k, v)
	}
	return env
}

// parseSudoID parses a sudo-exported numeric id, rejecting empty, non-numeric,
// zero, and overflowing values.
func parseSudoID(s string) (uint32, bool) {
	n, err := strconv.ParseUint(s, 10, 32)
	if err != nil || n == 0 {
		return 0, false
	}
	return uint32(n), true
}

func isUnderSudo() bool {
	return os.Geteuid() == 0 && os.Getenv("SUDO_USER") != "" && os.Getenv("SUDO_UID") != ""
}

// resolveEnvOverrides restores HOME and, when a probe is supplied, merges in the
// license vars it returns. A nil or failing probe leaves HOME set alone.
func resolveEnvOverrides(home string, probe probeFunc) envOverride {
	ov := envOverride{"HOME": home}
	if probe == nil {
		return ov
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	out, err := probe(ctx)
	if err != nil {
		slog.Warn("license env probe failed; using HOME only", "error", err)
		return ov
	}
	maps.Copy(ov, parseProbeOutput(out))
	return ov
}

// buildSudoOverrides enforces that home is known and wires the real probe when
// the caller (or a test) does not provide one.
func buildSudoOverrides(uid, gid uint32, home, shell string, probe probeFunc) (envOverride, error) {
	if home == "" {
		return nil, errHomeUnresolved
	}
	if probe == nil {
		probe = func(ctx context.Context) ([]byte, error) {
			return probeShellAsUser(ctx, shell, home, uid, gid)
		}
	}
	return resolveEnvOverrides(home, probe), nil
}

// envOverrides rebuilds the Vivado child environment under sudo: HOME from the
// invoking user's passwd entry, plus license vars probed from their login shell.
// It is a no-op when not under sudo.
func envOverrides() (envOverride, error) {
	if !isUnderSudo() {
		return nil, nil
	}
	uid, ok := parseSudoID(os.Getenv("SUDO_UID"))
	if !ok {
		return nil, nil
	}
	gid, _ := parseSudoID(os.Getenv("SUDO_GID"))
	home, shell, err := resolveSudoUser(os.Getenv("SUDO_UID"), os.Getenv("SUDO_USER"))
	if err != nil {
		return nil, fmt.Errorf("resolve invoking user under sudo: %w", err)
	}
	return buildSudoOverrides(uid, gid, home, shell, nil)
}

// resolveSudoUser looks up the invoking user by uid, falling back to the sudo
// username, and returns its home and login shell.
func resolveSudoUser(uidStr, sudoUser string) (string, string, error) {
	u, err := user.LookupId(uidStr)
	if err != nil && sudoUser != "" {
		u, err = user.Lookup(sudoUser)
	}
	if err != nil {
		return "", "", fmt.Errorf("user not found (uid=%s user=%s): %w", uidStr, sudoUser, err)
	}
	if sudoUser != "" && u.Username != sudoUser {
		slog.Warn("SUDO_USER does not match passwd entry for SUDO_UID", "sudo_user", sudoUser, "passwd_user", u.Username)
	}
	return u.HomeDir, lookupShell(u.Username, u.Uid), nil
}

// lookupShell reads the login shell from /etc/passwd, defaulting to bash.
func lookupShell(username, uid string) string {
	data, err := os.ReadFile("/etc/passwd")
	if err != nil {
		return "/bin/bash"
	}
	for _, line := range strings.Split(string(data), "\n") {
		f := strings.Split(line, ":")
		if len(f) >= 7 && (f[0] == username || f[2] == uid) && f[6] != "" {
			return f[6]
		}
	}
	return "/bin/bash"
}

// chownOutputs hands the copied *.bit/*.bin to the invoking user after a sudo
// build, so the firmware doesn't come out root-owned. No-op otherwise and never
// fatal.
func chownOutputs(outputDir string) error {
	if !isUnderSudo() {
		return nil
	}
	uid, ok := parseSudoID(os.Getenv("SUDO_UID"))
	if !ok {
		return nil
	}
	gid, _ := parseSudoID(os.Getenv("SUDO_GID"))
	if uid > math.MaxInt32 || gid > math.MaxInt32 {
		return nil
	}
	var firstErr error
	for _, pattern := range []string{filepath.Join(outputDir, "*.bit"), filepath.Join(outputDir, "*.bin")} {
		matches, _ := filepath.Glob(pattern)
		for _, m := range matches {
			if err := os.Lchown(m, int(uid), int(gid)); err != nil {
				slog.Warn("chown output file", "file", m, "error", err)
				if firstErr == nil {
					firstErr = err
				}
			}
		}
	}
	return firstErr
}
