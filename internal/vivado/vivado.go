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
	summary := summarizeVivadoOutput(tclScript, output.String(), runErr)
	if writeErr := os.WriteFile(filepath.Join(workDir, summaryFileName(tclScript)), []byte(summary), 0o644); writeErr != nil {
		slog.Warn("write Vivado summary", "error", writeErr)
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

	slog.Info("Vivado summary", "script", tclScript, "summary_file", summaryFileName(tclScript))
	return nil
}

func summaryFileName(tclScript string) string {
	base := filepath.Base(tclScript)
	base = strings.TrimSuffix(base, filepath.Ext(base))
	if base == "" {
		base = "vivado"
	}
	return base + "_summary.txt"
}

func summarizeVivadoOutput(script string, output string, runErr error) string {
	var b strings.Builder
	var errors, criticalWarnings, warnings int
	var keyLines []string
	diagnosis := classifyVivadoFailure(output, runErr)

	for _, raw := range strings.Split(output, "\n") {
		line := strings.TrimSpace(raw)
		if line == "" {
			continue
		}
		upper := strings.ToUpper(line)
		switch {
		case strings.Contains(upper, "ERROR:"):
			errors++
			keyLines = appendLimited(keyLines, line, 30)
		case strings.Contains(upper, "CRITICAL WARNING:"):
			criticalWarnings++
			keyLines = appendLimited(keyLines, line, 30)
		case strings.Contains(upper, "WARNING:"):
			warnings++
		case strings.Contains(strings.ToLower(line), "out of memory"),
			strings.Contains(strings.ToLower(line), "killed process"),
			strings.Contains(strings.ToLower(line), "no space left"),
			strings.Contains(strings.ToLower(line), "license"):
			keyLines = appendLimited(keyLines, line, 30)
		}
	}

	fmt.Fprintf(&b, "Vivado summary for %s\n", script)
	fmt.Fprintf(&b, "errors=%d critical_warnings=%d warnings=%d\n", errors, criticalWarnings, warnings)
	if runErr != nil {
		fmt.Fprintf(&b, "process_error=%v\n", runErr)
	}
	if diagnosis != "" {
		fmt.Fprintf(&b, "diagnosis=%s\n", diagnosis)
	}
	if len(keyLines) > 0 {
		b.WriteString("key_log_lines:\n")
		for _, line := range keyLines {
			fmt.Fprintf(&b, "- %s\n", line)
		}
	}
	return b.String()
}

func appendLimited(lines []string, line string, limit int) []string {
	if len(lines) >= limit {
		return lines
	}
	return append(lines, line)
}

func classifyVivadoFailure(output string, runErr error) string {
	text := strings.ToLower(output)
	if runErr != nil {
		text += "\n" + strings.ToLower(runErr.Error())
	}
	switch {
	case strings.Contains(text, "out of memory"),
		strings.Contains(text, "oom-kill"),
		strings.Contains(text, "killed process"),
		strings.Contains(text, "signal: killed"),
		strings.Contains(text, "signal: terminated"):
		return "Vivado was killed by the OS, usually RAM/OOM. Reduce --jobs, add swap, or build on a larger host."
	case strings.Contains(text, "no space left on device"):
		return "Build host ran out of disk space. Clean old jobs or enlarge the build volume."
	case strings.Contains(text, "license"):
		return "Vivado license/environment problem. Check XILINXD_LICENSE_FILE/LM_LICENSE_FILE and license availability."
	case strings.Contains(text, "invalid xml"), strings.Contains(text, "part0_pins.xml"):
		return "Vivado installation or device-part database looks broken. Repair/reinstall Vivado."
	case strings.Contains(text, "timing constraints are not met"), strings.Contains(text, "timing not met"):
		return "Implementation finished but timing failed. Lower clocks, reduce logic, or use a larger/faster FPGA target."
	case strings.Contains(text, "synthesis failed"), strings.Contains(text, "place_design"), strings.Contains(text, "route_design"):
		return "Vivado implementation failed. See key_log_lines and the generated runme.log files."
	default:
		if runErr != nil {
			return "Vivado returned a non-zero exit status. See key_log_lines and *_summary.txt."
		}
	}
	return ""
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
