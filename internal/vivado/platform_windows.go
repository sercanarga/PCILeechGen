//go:build windows

package vivado

import (
	"context"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

func platformDefaultPaths() []string {
	return []string{`C:\Xilinx\Vivado`, `C:\Program Files\Xilinx\Vivado`}
}

func launcherNames() []string {
	return []string{"vivado.bat"}
}

func vivadoCommand(ctx context.Context, binary string, args []string) *exec.Cmd {
	lower := strings.ToLower(binary)
	if !strings.HasSuffix(lower, ".bat") && !strings.HasSuffix(lower, ".cmd") {
		return exec.CommandContext(ctx, binary, args...)
	}
	comspec := os.Getenv("COMSPEC")
	if comspec == "" {
		comspec = "cmd.exe"
	}
	parts := make([]string, 0, len(args)+1)
	commandEnv := make([]string, 0, len(args)+1)
	const launcherEnv = "PCILEECHGEN_VIVADO_LAUNCHER"
	parts = append(parts, `"%`+launcherEnv+`%"`)
	commandEnv = append(commandEnv, launcherEnv+"="+binary)
	for i, arg := range args {
		name := "PCILEECHGEN_VIVADO_ARG_" + strconv.Itoa(i)
		parts = append(parts, `"%`+name+`%"`)
		commandEnv = append(commandEnv, name+"="+arg)
	}
	cmd := exec.CommandContext(ctx, comspec)
	cmd.Env = commandEnv
	cmd.SysProcAttr = &syscall.SysProcAttr{
		CmdLine: `/d /s /v:off /c "` + strings.Join(parts, " ") + `"`,
	}
	return cmd
}

