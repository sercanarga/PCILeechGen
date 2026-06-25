package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"text/tabwriter"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/spf13/cobra"
)

type boardOutput struct {
	Name         string       `json:"name"`
	FPGAPart     string       `json:"fpga_part"`
	PCIeLanes    int          `json:"pcie_lanes"`
	MaxLinkSpeed uint8        `json:"max_link_speed"`
	BRAMSize     int          `json:"bram_size"`
	TopModule    string       `json:"top_module"`
	ProjectDir   string       `json:"project_dir"`
	SubDir       string       `json:"sub_dir,omitempty"`
	TCLFile      string       `json:"tcl_file"`
	BuildTCL     string       `json:"build_tcl"`
	Flash        *board.Flash `json:"flash,omitempty"`
}

var boardsJSONOutput bool

var boardsCmd = &cobra.Command{
	Use:   "boards",
	Short: "List all supported FPGA boards",
	Long:  "Displays all supported pcileech-fpga board variants with their FPGA part and PCIe lane configuration.",
	RunE: func(cmd *cobra.Command, args []string) error {
		return writeBoards(os.Stdout, board.All(), boardsJSONOutput)
	},
}

func writeBoards(w io.Writer, boards []board.Board, jsonOutput bool) error {
	if jsonOutput {
		enc := json.NewEncoder(w)
		enc.SetIndent("", "  ")
		return enc.Encode(boardOutputs(boards))
	}

	tw := tabwriter.NewWriter(w, 0, 0, 2, ' ', 0)
	fmt.Fprintln(tw, "NAME\tFPGA PART\tPCIe\tBRAM\tTOP MODULE")
	fmt.Fprintln(tw, "----\t---------\t----\t----\t----------")

	for _, b := range boards {
		fmt.Fprintf(tw, "%s\t%s\tx%d\t%d\t%s\n",
			b.Name, b.FPGAPart, b.PCIeLanes, b.BRAMSizeOrDefault(), b.TopModule)
	}
	if err := tw.Flush(); err != nil {
		return fmt.Errorf("flush board output: %w", err)
	}

	fmt.Fprintf(w, "\nTotal: %d boards\n", len(boards))
	return nil
}

func boardOutputs(boards []board.Board) []boardOutput {
	result := make([]boardOutput, len(boards))
	for i, b := range boards {
		result[i] = boardOutput{
			Name:         b.Name,
			FPGAPart:     b.FPGAPart,
			PCIeLanes:    b.PCIeLanes,
			MaxLinkSpeed: b.MaxLinkSpeedOrDefault(),
			BRAMSize:     b.BRAMSizeOrDefault(),
			TopModule:    b.TopModule,
			ProjectDir:   b.ProjectDir,
			SubDir:       b.SubDir,
			TCLFile:      b.TCLFile,
			BuildTCL:     b.BuildTCL,
			Flash:        b.Flash,
		}
	}
	return result
}

func init() {
	boardsCmd.Flags().BoolVar(&boardsJSONOutput, "json", false, "print board data as JSON")
	rootCmd.AddCommand(boardsCmd)
}
