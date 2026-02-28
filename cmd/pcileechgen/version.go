package main

import (
	"fmt"

	"github.com/sercanarga/pcileechgen/internal/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("pcileechgen %s\n", version.Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
