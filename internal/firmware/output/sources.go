package output

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/firmware"
	"github.com/sercanarga/pcileechgen/internal/firmware/svgen"
)

// svFilesReplacedByGenerator: board source files we replace with
// generated versions. These are excluded from the board source copy
// so Vivado only sees the generated versions.
var svFilesReplacedByGenerator = map[string]bool{
	"pcileech_tlps128_bar_controller.sv": true,
}

// barControllerSubModules: sub-module names extracted from the board's
// BAR controller file. The top-level module (pcileech_tlps128_bar_controller)
// is replaced by the generated version, but these shared modules are still
// needed by the generated controller.
var barControllerSubModules = []string{
	"pcileech_tlps128_bar_wrengine",
	"pcileech_bar_impl_none",
	"pcileech_bar_impl_zerowrite4k",
}

// patchSVSources copies the board's SV tree (excluding files that will
// be regenerated), and patches donor IDs into the remaining sources.
func (ow *OutputWriter) patchSVSources(b *board.Board, ids firmware.DeviceIDs) error {
	srcDir := b.SrcPath(ow.LibDir)
	dstDir := filepath.Join(ow.OutputDir, "src")

	srcInfo, err := os.Lstat(srcDir)
	if os.IsNotExist(err) {
		slog.Warn("board source dir not found", "path", srcDir)
		slog.Info("run: git submodule update --init --recursive")
		return fmt.Errorf("board sources not found at %s (is the pcileech-fpga submodule initialized?)", srcDir)
	} else if err != nil {
		return fmt.Errorf("inspect board source dir %s: %w", srcDir, err)
	} else if srcInfo.Mode()&os.ModeSymlink != 0 || !srcInfo.IsDir() {
		return fmt.Errorf("board source dir must be a real directory: %s", srcDir)
	}

	if err := clearGeneratedSourceTree(dstDir); err != nil {
		return fmt.Errorf("failed to clear stale src dir %s: %w", dstDir, err)
	}

	// Copy board sources to local src/ (excluding the top-level bar controller
	// when generating custom version). The generated vivado_generate_project.tcl
	// now globs from this local dir so the donor-custom controller is used
	// instead of stock, avoiding duplicate module imports.
	if err := copyDirExcluding(srcDir, dstDir, svFilesReplacedByGenerator); err != nil {
		return fmt.Errorf("failed to copy SV sources: %w", err)
	}

	// Extract sub-modules from the board's BAR controller file and write
	// them as separate .sv files. The generated BAR controller depends on
	// these but the board's top-level module would conflict with the
	// generated version, so we exclude the original file and split it.
	if !ow.StockBar {
		ctrlSrc := filepath.Join(srcDir, "pcileech_tlps128_bar_controller.sv")
		if _, err := os.Stat(ctrlSrc); err == nil {
			if eerr := extractSubModules(ctrlSrc, dstDir, barControllerSubModules); eerr != nil {
				return fmt.Errorf("failed to extract BAR controller sub-modules: %w", eerr)
			}
		} else if os.IsNotExist(err) {
			return fmt.Errorf("required BAR controller not found: %s", ctrlSrc)
		} else {
			return fmt.Errorf("failed to inspect BAR controller: %w", err)
		}
	} else {
		slog.Info("stock-bar mode: keeping stock bar controller")
		srcFile := filepath.Join(srcDir, "pcileech_tlps128_bar_controller.sv")
		dstFile := filepath.Join(dstDir, "pcileech_tlps128_bar_controller.sv")
		data, err := os.ReadFile(srcFile)
		if err != nil {
			return fmt.Errorf("failed to read required stock BAR controller: %w", err)
		}
		if err := writeRegularFile(dstFile, data); err != nil {
			return fmt.Errorf("failed to copy stock BAR controller: %w", err)
		}
	}

	patcher := svgen.NewSVPatcher(ids, dstDir)

	if err := patcher.PatchAll(); err != nil {
		return fmt.Errorf("failed to patch SV sources: %w", err)
	}

	if results := patcher.Results(); len(results) > 0 {
		slog.Info("SV patches applied", "summary", svgen.FormatPatchSummary(results))
	}

	return nil
}

func clearGeneratedSourceTree(dstDir string) error {
	info, err := os.Lstat(dstDir)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return err
	}
	if info.Mode()&os.ModeSymlink != 0 || !info.IsDir() {
		return fmt.Errorf("generated source directory must be a real directory")
	}
	if err := filepath.WalkDir(dstDir, func(path string, entry os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.Type()&os.ModeSymlink != 0 {
			return fmt.Errorf("generated source tree contains a symlink: %s", path)
		}
		return nil
	}); err != nil {
		return err
	}
	return os.RemoveAll(dstDir)
}

// copyDirExcluding copies a directory recursively but skips files whose
// names are in the exclude map. Used to prevent board source files from
// being imported alongside generated versions with the same module names.
func copyDirExcluding(src, dst string, exclude map[string]bool) error {
	if err := os.MkdirAll(dst, 0755); err != nil {
		return err
	}
	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}
	for _, e := range entries {
		srcPath := filepath.Join(src, e.Name())
		dstPath := filepath.Join(dst, e.Name())
		if e.Type()&os.ModeSymlink != 0 {
			return fmt.Errorf("refusing to copy symlinked board source: %s", srcPath)
		}
		if e.IsDir() {
			if err := copyDirExcluding(srcPath, dstPath, exclude); err != nil {
				return err
			}
		} else if !exclude[e.Name()] {
			data, err := os.ReadFile(srcPath)
			if err != nil {
				return err
			}
			if err := writeRegularFile(dstPath, data); err != nil {
				return err
			}
		}
	}
	return nil
}

// extractSubModules reads a board BAR controller source file and writes
// each named sub-module as a separate .sv file. This allows the generated
// top-level BAR controller to use shared sub-modules without conflicting
// with the board's original file.
func extractSubModules(srcPath string, dstDir string, subModules []string) error {
	data, err := os.ReadFile(srcPath)
	if err != nil {
		return err
	}

	content := string(data)
	for _, modName := range subModules {
		searchStr := "module " + modName
		start := strings.Index(content, searchStr)
		if start == -1 {
			return fmt.Errorf("required sub-module %s not found in %s", modName, srcPath)
		}

		// Find the end: next "\nmodule " after our match, or end of file.
		rest := content[start+len(searchStr):]
		nextIdx := strings.Index(rest, "\nmodule ")
		var modBody string
		if nextIdx == -1 {
			modBody = content[start:]
		} else {
			modBody = content[start : start+len(searchStr)+nextIdx+1]
		}

		dstFile := filepath.Join(dstDir, modName+".sv")
		if err := writeRegularFile(dstFile, []byte(modBody)); err != nil {
			return fmt.Errorf("write required sub-module %s: %w", modName, err)
		}
	}
	return nil
}

func (ow *OutputWriter) writeFile(name, content string) error {
	if name == "" || filepath.Base(name) != name {
		return fmt.Errorf("output filename must be a simple basename: %q", name)
	}
	if err := validateRealDirectory(ow.OutputDir, "output directory"); err != nil {
		return err
	}
	return writeRegularFile(filepath.Join(ow.OutputDir, name), []byte(content))
}

func writeRegularFile(path string, content []byte) error {
	info, err := os.Lstat(path)
	if err == nil {
		if info.Mode()&os.ModeSymlink != 0 || !info.Mode().IsRegular() {
			return fmt.Errorf("output file must be absent or a regular file: %s", path)
		}
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("inspect output file %s: %w", path, err)
	}
	if err := os.WriteFile(path, content, 0644); err != nil {
		return err
	}
	return nil
}
