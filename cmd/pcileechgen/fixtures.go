package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/sercanarga/pcileechgen/internal/donor/synthetic"
	"github.com/sercanarga/pcileechgen/internal/firmware/devclass"
	"github.com/spf13/cobra"
)

var fixturesOpts struct {
	out string
}

var fixturesCmd = &cobra.Command{
	Use:   "fixtures",
	Short: "Generate synthetic donor fixtures for CI HDL verification",
	Long:  "Writes one representative device_context.json per device class.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := os.MkdirAll(fixturesOpts.out, 0o755); err != nil {
			return fmt.Errorf("create output dir: %w", err)
		}
		for _, class := range devclass.AllClasses() {
			ctx := synthetic.Build(class)
			if ctx == nil {
				return fmt.Errorf("no synthetic builder for class %q", class)
			}
			data, err := json.MarshalIndent(ctx, "", "  ")
			if err != nil {
				return fmt.Errorf("marshal %s: %w", class, err)
			}
			path := filepath.Join(fixturesOpts.out, class+".json")
			if err := os.WriteFile(path, data, 0o644); err != nil {
				return fmt.Errorf("write %s: %w", path, err)
			}
			fmt.Println("wrote", path)
		}
		return nil
	},
}

func init() {
	fixturesCmd.Flags().StringVar(&fixturesOpts.out, "out", "testdata/donors",
		"output directory for fixture JSON files")
	rootCmd.AddCommand(fixturesCmd)
}
