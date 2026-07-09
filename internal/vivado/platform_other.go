//go:build !windows

package vivado

import (
	"context"
	"os/exec"
)

func platformDefaultPaths() []string {
	return []string{"/tools/Xilinx/Vivado", "/opt/Xilinx/Vivado", "/usr/local/Xilinx/Vivado"}
}

func launcherNames() []string {
	return []string{"vivado"}
}

func vivadoCommand(ctx context.Context, binary string, args []string) *exec.Cmd {
	return exec.CommandContext(ctx, binary, args...)
}
