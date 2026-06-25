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

var svFilesReplacedByGenerator = map[string]bool{
	"pcileech_tlps128_bar_controller.sv": true,
}

var barControllerSubModules = []string{
	"pcileech_tlps128_bar_rdengine",
	"pcileech_tlps128_bar_wrengine",
	"pcileech_bar_impl_none",
	"pcileech_bar_impl_loopaddr",
	"pcileech_bar_impl_zerowrite4k",
}

func (ow *OutputWriter) patchSVSources(b *board.Board, ids firmware.DeviceIDs) error {
	srcDir := b.SrcPath(ow.LibDir)
	dstDir := filepath.Join(ow.OutputDir, "src")

	if _, err := os.Stat(srcDir); os.IsNotExist(err) {
		slog.Warn("board source dir not found", "path", srcDir)
		slog.Info("run: git submodule update --init --recursive")
		return fmt.Errorf("board sources not found at %s (is the pcileech-fpga submodule initialized?)", srcDir)
	}

	if err := copyDirExcluding(srcDir, dstDir, svFilesReplacedByGenerator); err != nil {
		return fmt.Errorf("failed to copy SV sources: %w", err)
	}

	if !ow.StockBar {
		ctrlSrc := filepath.Join(srcDir, "pcileech_tlps128_bar_controller.sv")
		if err := extractSubModules(ctrlSrc, dstDir, barControllerSubModules); err != nil {
			slog.Warn("could not extract BAR controller sub-modules, board source may be incompatible", "error", err)
		}
	} else {
		slog.Info("stock-bar mode: keeping stock bar controller")
		srcFile := filepath.Join(srcDir, "pcileech_tlps128_bar_controller.sv")
		dstFile := filepath.Join(dstDir, "pcileech_tlps128_bar_controller.sv")
		if data, err := os.ReadFile(srcFile); err == nil {
			if err := os.WriteFile(dstFile, data, 0644); err != nil {
				return fmt.Errorf("copy stock BAR controller: %w", err)
			}
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
		if e.IsDir() {
			if err := copyDirExcluding(srcPath, dstPath, exclude); err != nil {
				return err
			}
		} else if !exclude[e.Name()] {
			data, err := os.ReadFile(srcPath)
			if err != nil {
				return err
			}
			if err := os.WriteFile(dstPath, data, 0644); err != nil {
				return err
			}
		}
	}
	return nil
}

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
			slog.Warn("sub-module not found in board controller", "module", modName)
			continue
		}

		rest := content[start+len(searchStr):]
		nextIdx := strings.Index(rest, "\nmodule ")
		var modBody string
		if nextIdx == -1 {
			modBody = content[start:]
		} else {
			modBody = content[start : start+len(searchStr)+nextIdx+1]
		}

		dstFile := filepath.Join(dstDir, modName+".sv")
		if err := os.WriteFile(dstFile, []byte(modBody), 0644); err != nil {
			slog.Warn("failed to write sub-module file", "module", modName, "error", err)
		}
	}
	return nil
}

func (ow *OutputWriter) writeFile(name, content string) error {
	return os.WriteFile(filepath.Join(ow.OutputDir, name), []byte(content), 0644)
}

func ListOutputFiles() []string {
	return []string{
		"device_context.json",
		"pcileech_cfgspace.coe",
		"pcileech_cfgspace_writemask.coe",
		"pcileech_bar_zero4k.coe",
		"vivado_generate_project.tcl",
		"vivado_build.tcl",
		"src/",
		"pcileech_bar_impl_device.sv",
		"pcileech_tlps128_bar_controller.sv",
		"pcileech_bar_impl_msi.sv",
		"pcileech_msix_table.sv",
		"pcileech_nvme_admin_responder.sv",
		"pcileech_nvme_dma_bridge.sv",
		"tlp_latency_emulator.sv",
		"device_config.sv",
		"config_space_init.hex",
		"msix_table_init.hex",
		"identify_init.hex",
		"windows_lab_checklist.txt",
		"scrub_diff_report.txt",
		"build_manifest.json",
	}
}
