//go:build linux

package vivado

import (
	"context"
	"os/exec"
	"syscall"
)

const probePrintfScript = `printf 'XILINXD_LICENSE_FILE=%s\nLM_LICENSE_FILE=%s\nPATH=%s\n' "$XILINXD_LICENSE_FILE" "$LM_LICENSE_FILE" "$PATH"`

// probeShellAsUser runs the invoking user's login shell as uid:gid (a child-only
// credential drop; the parent stays root) with a minimal env so the profile
// files re-export the license vars, and returns the printed values.
func probeShellAsUser(ctx context.Context, shell, home string, uid, gid uint32) ([]byte, error) {
	cmd := exec.CommandContext(ctx, shell, "-lic", probePrintfScript)
	cmd.Env = []string{"HOME=" + home, "TERM=dumb", "PATH=/usr/bin:/bin"}
	cmd.SysProcAttr = &syscall.SysProcAttr{Credential: &syscall.Credential{Uid: uid, Gid: gid}}
	return cmd.Output()
}
