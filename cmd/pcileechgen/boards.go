package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/spf13/cobra"
)

var boardsCmd = &cobra.Command{
	Use:   "boards",
	Short: "List all supported FPGA boards",
	Long:  "Displays all supported pcileech-fpga board variants with their FPGA part and PCIe lane configuration.",
	Run: func(cmd *cobra.Command, args []string) {
		boards := board.All()

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintln(w, "NAME\tFPGA PART\tPCIe\tTOP MODULE")
		fmt.Fprintln(w, "----\t---------\t----\t----------")

		for _, b := range boards {
			fmt.Fprintf(w, "%s\t%s\tx%d\t%s\n",
				b.Name, b.FPGAPart, b.PCIeLanes, b.TopModule)
		}
		w.Flush()

		fmt.Printf("\nTotal: %d boards\n", len(boards))
	},
}

func init() {
	rootCmd.AddCommand(boardsCmd)
}
