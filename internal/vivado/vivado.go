// Package vivado wraps Xilinx Vivado CLI for synthesis.
package vivado

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
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

		// Find the latest version
		var latestVersion string
		for _, e := range entries {
			if e.IsDir() {
				latestVersion = e.Name()
			}
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

// RunTCL executes a TCL script in Vivado batch mode.
func (v *Vivado) RunTCL(tclScript string, workDir string) error {
	cmd := exec.Command(v.BinaryPath(), "-mode", "batch", "-notrace", "-source", tclScript)
	cmd.Dir = workDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Set up environment
	env := os.Environ()
	env = append(env, fmt.Sprintf("XILINX_VIVADO=%s", v.Path))
	cmd.Env = env

	fmt.Printf("[vivado] Running: %s\n", strings.Join(cmd.Args, " "))
	fmt.Printf("[vivado] Working directory: %s\n", workDir)

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("Vivado execution failed: %w", err)
	}

	return nil
}
