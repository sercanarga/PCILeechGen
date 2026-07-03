package main

import (
	"encoding/json"
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/spf13/cobra"
)

var boardsJSON bool

var boardsCmd = &cobra.Command{
	Use:   "boards",
	Short: "List all supported FPGA boards",
	Long:  "Displays all supported pcileech-fpga board variants with their FPGA part and PCIe lane configuration.",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runBoards(cmd.OutOrStdout(), boardsJSON)
	},
}

type boardInfo struct {
	Name         string `json:"name"`
	FPGAPart     string `json:"fpga_part"`
	PCIeLanes    int    `json:"pcie_lanes"`
	MaxLinkSpeed uint8  `json:"max_link_speed"`
	BRAMSize     int    `json:"bram_size"`
	TopModule    string `json:"top_module"`
	ProjectDir   string `json:"project_dir"`
	SubDir       string `json:"sub_dir,omitempty"`
	SourceSubDir string `json:"source_sub_dir,omitempty"`
	TCLFile      string `json:"tcl_file"`
	BuildTCL     string `json:"build_tcl,omitempty"`
}

func runBoards(out io.Writer, jsonOutput bool) error {
	boards := board.All()
	if jsonOutput {
		enc := json.NewEncoder(out)
		enc.SetIndent("", "  ")
		if err := enc.Encode(boardInfos(boards)); err != nil {
			return fmt.Errorf("render boards JSON: %w", err)
		}
		return nil
	}

	w := tabwriter.NewWriter(out, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "NAME\tFPGA PART\tPCIe\tBRAM\tTOP MODULE")
	fmt.Fprintln(w, "----\t---------\t----\t----\t----------")

	for _, b := range boards {
		fmt.Fprintf(w, "%s\t%s\tx%d\t%d\t%s\n",
			b.Name, b.FPGAPart, b.PCIeLanes, b.BRAMSizeOrDefault(), b.TopModule)
	}
	if err := w.Flush(); err != nil {
		return fmt.Errorf("render boards table: %w", err)
	}

	fmt.Fprintf(out, "\nTotal: %d boards\n", len(boards))
	return nil
}

func boardInfos(boards []board.Board) []boardInfo {
	infos := make([]boardInfo, 0, len(boards))
	for _, b := range boards {
		infos = append(infos, boardInfo{
			Name:         b.Name,
			FPGAPart:     b.FPGAPart,
			PCIeLanes:    b.PCIeLanes,
			MaxLinkSpeed: b.MaxLinkSpeedOrDefault(),
			BRAMSize:     b.BRAMSizeOrDefault(),
			TopModule:    b.TopModule,
			ProjectDir:   b.ProjectDir,
			SubDir:       b.SubDir,
			SourceSubDir: b.SourceSubDir,
			TCLFile:      b.TCLFile,
			BuildTCL:     b.BuildTCL,
		})
	}
	return infos
}

func init() {
	boardsCmd.Flags().BoolVar(&boardsJSON, "json", false, "emit machine-readable board definitions")
	rootCmd.AddCommand(boardsCmd)
}
