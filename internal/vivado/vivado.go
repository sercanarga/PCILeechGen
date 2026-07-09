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

var DefaultPaths = platformDefaultPaths()

// Vivado represents a Vivado installation.
type Vivado struct {
	Path       string
	Version    string
	binaryPath string
}

// Find searches for a Vivado installation.
// If customPath is provided, it's used directly. Otherwise, searches default locations.
func Find(customPath string) (*Vivado, error) {
	if customPath != "" {
		return validateVivado(customPath)
	}

	for _, name := range launcherNames() {
		vivadoExec, err := exec.LookPath(name)
		if err != nil {
			continue
		}
		realPath, evalErr := filepath.EvalSymlinks(vivadoExec)
		if evalErr != nil {
			realPath = vivadoExec
		}
		installDir := filepath.Dir(filepath.Dir(realPath))
		v, validateErr := validateVivado(installDir)
		if validateErr == nil {
			v.binaryPath = realPath
			return v, nil
		}
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
	info, err := os.Stat(path)
	if err == nil && !info.IsDir() {
		path = filepath.Dir(filepath.Dir(path))
	}

	for _, name := range launcherNames() {
		binary := filepath.Join(path, "bin", name)
		if info, statErr := os.Stat(binary); statErr == nil && !info.IsDir() {
			return &Vivado{
				Path:       path,
				Version:    filepath.Base(path),
				binaryPath: binary,
			}, nil
		}
	}

	return nil, fmt.Errorf("Vivado launcher not found under %s", filepath.Join(path, "bin"))
}

// BinaryPath returns the path to the Vivado binary.
func (v *Vivado) BinaryPath() string {
	if v.binaryPath != "" {
		return v.binaryPath
	}
	for _, name := range launcherNames() {
		candidate := filepath.Join(v.Path, "bin", name)
		if info, err := os.Stat(candidate); err == nil && !info.IsDir() {
			return candidate
		}
	}
	return filepath.Join(v.Path, "bin", launcherNames()[0])
}

// RunTCL executes a TCL script in Vivado batch mode with a timeout. Under sudo
// the invoking user's license env is recovered via envOverrides; otherwise it's
// a no-op.
func (v *Vivado) RunTCL(tclScript string, workDir string, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	args := []string{"-mode", "batch", "-notrace", "-source", tclScript}
	cmd := vivadoCommand(ctx, v.BinaryPath(), args)
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
	cmd.Env = append(env, cmd.Env...)

	slog.Info("running Vivado", "cmd", strings.Join(cmd.Args, " "), "dir", workDir, "timeout", timeout)

	if err := cmd.Run(); err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return fmt.Errorf("Vivado timed out after %s", timeout)
		}
		return fmt.Errorf("Vivado execution failed: %w", err)
	}

	if line, ok := vivadoStartupFailure(output.String()); ok {
		return fmt.Errorf("Vivado startup failed: %s", line)
	}

	return nil
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
