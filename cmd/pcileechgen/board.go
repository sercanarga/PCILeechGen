package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/spf13/cobra"
)

type boardCheckItem struct {
	Label    string `json:"label"`
	Path     string `json:"path"`
	Exists   bool   `json:"exists"`
	Required bool   `json:"required"`
	Reason   string `json:"reason,omitempty"`
}

type boardInspectOptions struct {
	name   string
	libDir string
	json   bool
}

type boardInspectResult struct {
	Name      string           `json:"name"`
	FPGA      string           `json:"fpga_part"`
	PCIeLanes int              `json:"pcie_lanes"`
	Top       string           `json:"top_module"`
	Project   string           `json:"project_dir"`
	LibDir    string           `json:"lib_dir"`
	Checks    []boardCheckItem `json:"checks"`
	Passes    int              `json:"passes"`
	Warnings  int              `json:"warnings"`
}

var boardInspectOpts boardInspectOptions
var boardScaffoldOpts struct {
	name string
}

var boardCmd = &cobra.Command{
	Use:   "board",
	Short: "Inspect and generate board metadata helpers",
	Long:  "Inspect board metadata and required source artifacts, and scaffold board metadata blocks.",
}

var boardInspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "Inspect board metadata and required local artifacts",
	RunE: func(cmd *cobra.Command, args []string) error {
		b, err := board.Find(boardInspectOpts.name)
		if err != nil {
			return err
		}

		checks := inspectBoardArtifacts(b, boardInspectOpts.libDir)
		res := boardInspectResult{
			Name:      b.Name,
			FPGA:      b.FPGAPart,
			PCIeLanes: b.PCIeLanes,
			Top:       b.TopModule,
			Project:   b.ProjectDir,
			LibDir:    boardInspectOpts.libDir,
			Checks:    checks,
		}
		for _, item := range checks {
			if item.Exists {
				res.Passes++
			} else if item.Required {
				res.Warnings++
			}
		}

		if err := writeBoardInspectResult(cmd.OutOrStdout(), &res, boardInspectOpts.json); err != nil {
			return err
		}

		if res.Warnings > 0 {
			return fmt.Errorf("%d required artifact(s) missing for board %q", res.Warnings, b.Name)
		}
		return nil
	},
}

var boardScaffoldCmd = &cobra.Command{
	Use:   "scaffold",
	Short: "Print a board metadata scaffold",
	RunE: func(cmd *cobra.Command, args []string) error {
		if boardScaffoldOpts.name == "" {
			return fmt.Errorf("--name is required")
		}
		scaffold := boardOutput{
			Name:         boardScaffoldOpts.name,
			FPGAPart:     "xc7a200t...",
			PCIeLanes:    1,
			MaxLinkSpeed: 2,
			BRAMSize:     8192,
			TopModule:    "pcileech_<board>_top",
			ProjectDir:   boardScaffoldOpts.name,
			SubDir:       "",
			TCLFile:      "vivado_generate_project.tcl",
			BuildTCL:     "vivado_build.tcl",
		}
		enc := json.NewEncoder(cmd.OutOrStdout())
		enc.SetIndent("", "  ")
		return enc.Encode([]boardOutput{scaffold})
	},
}

func inspectBoardArtifacts(b *board.Board, libDir string) []boardCheckItem {
	checks := []boardCheckItem{}
	add := func(label, path string, required bool) {
		_, err := os.Stat(path)
		exists := err == nil
		reason := ""
		if !exists {
			reason = "missing"
		}
		checks = append(checks, boardCheckItem{
			Label:    label,
			Path:     path,
			Exists:   exists,
			Required: required,
			Reason:   reason,
		})
	}

	checkLibDir := b.LibPath(libDir)
	add("lib_dir", checkLibDir, true)
	add("src", b.SrcPath(libDir), true)
	add("tcl", b.TCLPath(libDir), true)
	add("build_tcl", b.BuildTCLPath(libDir), false)
	add("ip", b.IPPath(libDir), false)

	if b.Flash != nil {
		if b.Flash.Config != "" {
			add("flash_config", filepath.Join(checkLibDir, b.Flash.Config), false)
		}
		if b.Flash.Script != "" {
			add("flash_script", filepath.Join(checkLibDir, b.Flash.Script), false)
		}
	}

	sort.Slice(checks, func(i, j int) bool {
		return checks[i].Label < checks[j].Label
	})
	return checks
}

func writeBoardInspectResult(w io.Writer, res *boardInspectResult, asJSON bool) error {
	if asJSON {
		enc := json.NewEncoder(w)
		enc.SetIndent("", "  ")
		return enc.Encode(res)
	}

	fmt.Fprintf(w, "Board: %s\n", res.Name)
	fmt.Fprintf(w, "FPGA: %s (x%d)\n", res.FPGA, res.PCIeLanes)
	fmt.Fprintf(w, "Top: %s\n", res.Top)
	fmt.Fprintf(w, "Project: %s\n", res.Project)
	fmt.Fprintf(w, "Lib dir: %s\n\n", res.LibDir)
	var present, missing, optional int
	for _, item := range res.Checks {
		status := "WARN"
		if item.Exists {
			status = "OK"
			present++
		} else if !item.Required {
			status = "SKIP"
			optional++
		} else {
			missing++
		}
		fmt.Fprintf(w, "%-14s %s\n", status, item.Label)
		fmt.Fprintf(w, "             %s\n", item.Path)
		if item.Reason != "" {
			fmt.Fprintf(w, "             %s\n", item.Reason)
		}
		fmt.Fprint(w, "\n")
	}
	fmt.Fprintf(w, "Summary: present=%d missing=%d optional=%d\n", present, missing, optional)
	if missing > 0 {
		return fmt.Errorf("%d required artifact(s) missing", missing)
	}
	return nil
}

func init() {
	boardCmd.AddCommand(boardInspectCmd)
	boardCmd.AddCommand(boardScaffoldCmd)

	boardInspectCmd.Flags().StringVarP(&boardInspectOpts.name, "name", "n", "", "board name")
	boardInspectCmd.Flags().StringVarP(&boardInspectOpts.libDir, "lib-dir", "l", "pcileech-fpga", "path to pcileech-fpga checkout")
	boardInspectCmd.Flags().BoolVar(&boardInspectOpts.json, "json", false, "output JSON")
	_ = boardInspectCmd.MarkFlagRequired("name")

	boardScaffoldCmd.Flags().StringVarP(&boardScaffoldOpts.name, "name", "n", "", "board name")
	_ = boardScaffoldCmd.MarkFlagRequired("name")

	rootCmd.AddCommand(boardCmd)
}
