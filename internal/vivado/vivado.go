// Package vivado wraps Xilinx Vivado CLI for synthesis.
package vivado

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

// DefaultPaths contains common Vivado installation paths.
var DefaultPaths = []string{
	"/tools/Xilinx/Vivado",
	"/opt/Xilinx/Vivado",
	"/usr/local/Xilinx/Vivado",
}

// Vivado represents a Vivado installation.
type Vivado struct {
	Path    string // path to Vivado installation (e.g., /tools/Xilinx/Vivado/2023.2)
	Version string // Vivado version string
}

// Find searches for a Vivado installation.
// If customPath is provided, it's used directly. Otherwise, searches default locations.
func Find(customPath string) (*Vivado, error) {
	if customPath != "" {
		return validateVivado(customPath)
	}

	// Check PATH first
	vivadoExec, err := exec.LookPath("vivado")
	if err == nil {
		// Resolve to installation directory
		realPath, _ := filepath.EvalSymlinks(vivadoExec)
		installDir := filepath.Dir(filepath.Dir(realPath))
		return validateVivado(installDir)
	}

	// Search default paths
	for _, basePath := range DefaultPaths {
		entries, err := os.ReadDir(basePath)
		if err != nil {
			continue
		}

		// Find the latest version by sorting directory names
		var versions []string
		for _, e := range entries {
			if e.IsDir() {
				versions = append(versions, e.Name())
			}
		}
		sort.Strings(versions)
		var latestVersion string
		if len(versions) > 0 {
			latestVersion = versions[len(versions)-1]
		}

		if latestVersion != "" {
			fullPath := filepath.Join(basePath, latestVersion)
			v, err := validateVivado(fullPath)
			if err == nil {
				return v, nil
			}
		}
	}

	return nil, fmt.Errorf("Vivado not found. Install Vivado and either add it to PATH " +
		"or specify the path with --vivado-path")
}

// validateVivado checks if a path contains a valid Vivado installation.
func validateVivado(path string) (*Vivado, error) {
	binary := filepath.Join(path, "bin", "vivado")
	if _, err := os.Stat(binary); os.IsNotExist(err) {
		return nil, fmt.Errorf("Vivado binary not found at %s", binary)
	}

	version := filepath.Base(path)

	return &Vivado{
		Path:    path,
		Version: version,
	}, nil
}

// BinaryPath returns the path to the Vivado binary.
func (v *Vivado) BinaryPath() string {
	return filepath.Join(v.Path, "bin", "vivado")
}

// RunTCL executes a TCL script in Vivado batch mode with a timeout. Under sudo
// the invoking user's license env is recovered via envOverrides; otherwise it's
// a no-op.
func (v *Vivado) RunTCL(tclScript string, workDir string, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, v.BinaryPath(), "-mode", "batch", "-notrace", "-source", tclScript)
	cmd.Dir = workDir
	var output bytes.Buffer
	cmd.Stdout = io.MultiWriter(os.Stdout, &output)
	cmd.Stderr = io.MultiWriter(os.Stderr, &output)

	env := os.Environ()
	ov, err := envOverrides()
	if err != nil {
		return err
	}
	env = applyOverrides(env, ov)
	env = setEnv(env, "XILINX_VIVADO", v.Path)
	cmd.Env = env

	slog.Info("running Vivado", "cmd", strings.Join(cmd.Args, " "), "dir", workDir, "timeout", timeout)

	runErr := cmd.Run()
	summary := summarizeRun(tclScript, output.String(), runErr)
	if err := writeRunSummary(workDir, tclScript, summary); err != nil {
		slog.Warn("write Vivado summary", "error", err)
	}
	if runErr != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return fmt.Errorf("Vivado timed out after %s\n%s", timeout, summary)
		}
		return fmt.Errorf("Vivado execution failed: %w\n%s", runErr, summary)
	}

	if line, ok := vivadoStartupFailure(output.String()); ok {
		return fmt.Errorf("Vivado startup failed: %s\n%s", line, summary)
	}

	slog.Info("Vivado summary written", "file", filepath.Join(workDir, summaryFileName(tclScript)))
	return nil
}

func summaryFileName(tclScript string) string {
	base := filepath.Base(tclScript)
	base = strings.TrimSuffix(base, filepath.Ext(base))
	if base == "" || base == "." {
		base = "vivado"
	}
	return base + "_summary.txt"
}

// summarizeRun persists the existing structured Vivado report alongside a
// concise diagnosis for common infrastructure failures. The result deliberately
// includes only the report's bounded actionable entries, never the full log.
func summarizeRun(tclScript, output string, runErr error) string {
	var summary strings.Builder
	summary.WriteString("Vivado summary for ")
	summary.WriteString(filepath.Base(tclScript))
	summary.WriteString("\n")
	summary.WriteString(ParseOutput(output).Summary())
	if runErr != nil {
		fmt.Fprintf(&summary, "process_error=%v\n", runErr)
	}
	if diagnosis := classifyRunFailure(output, runErr); diagnosis != "" {
		fmt.Fprintf(&summary, "diagnosis=%s\n", diagnosis)
	}
	return summary.String()
}

func classifyRunFailure(output string, runErr error) string {
	text := strings.ToLower(output)
	if runErr != nil {
		text += "\n" + strings.ToLower(runErr.Error())
	}
	switch {
	case strings.Contains(text, "out of memory"), strings.Contains(text, "oom-kill"),
		strings.Contains(text, "killed process"), strings.Contains(text, "signal: killed"):
		return "Vivado was killed by the OS, usually due to RAM pressure. Reduce --jobs or use a larger build host."
	case strings.Contains(text, "no space left on device"):
		return "The build host ran out of disk space. Clean old builds or enlarge the build volume."
	case strings.Contains(text, "license"):
		return "Vivado license or environment problem. Check XILINXD_LICENSE_FILE/LM_LICENSE_FILE."
	case strings.Contains(text, "invalid xml"), strings.Contains(text, "part0_pins.xml"):
		return "The Vivado installation or device-part database looks damaged; repair or reinstall Vivado."
	case strings.Contains(text, "timing constraints are not met"), strings.Contains(text, "timing not met"):
		return "Implementation completed but timing failed; review constraints and target clock settings."
	case runErr != nil:
		return "Vivado returned a non-zero exit status; review the actionable issues above."
	}
	return ""
}

func writeRunSummary(workDir, tclScript, summary string) error {
	info, err := os.Lstat(workDir)
	if err != nil {
		return fmt.Errorf("inspect Vivado work directory: %w", err)
	}
	if info.Mode()&os.ModeSymlink != 0 || !info.IsDir() {
		return fmt.Errorf("Vivado work directory must be a real directory: %s", workDir)
	}
	path := filepath.Join(workDir, summaryFileName(tclScript))
	if info, err := os.Lstat(path); err == nil {
		if info.Mode()&os.ModeSymlink != 0 || !info.Mode().IsRegular() {
			return fmt.Errorf("Vivado summary path must be absent or a regular file: %s", path)
		}
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("inspect Vivado summary path: %w", err)
	}
	return os.WriteFile(path, []byte(summary), 0o644)
}

func vivadoStartupFailure(output string) (string, bool) {
	for _, line := range strings.Split(output, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if strings.Contains(line, "application-specific initialization failed") {
			return line, true
		}
	}
	return "", false
}
