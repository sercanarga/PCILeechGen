package main

import (
	"fmt"

	"github.com/sercanarga/pcileechgen/internal/color"
	"github.com/sercanarga/pcileechgen/internal/firmware/output"
	"github.com/spf13/cobra"
)

var verifyManifestPath string
var verifyOutputDir string

var verifyManifestCmd = &cobra.Command{
	Use:   "verify-manifest",
	Short: "Verify build integrity against manifest checksums",
	Long: `Reads a build_manifest.json and verifies that all listed files
exist and their SHA256 checksums match. Use after a build to confirm
artifacts haven't been corrupted or tampered with.

Example:
  pcileechgen verify-manifest --manifest pcileech_datastore/build_manifest.json`,
	RunE: func(cmd *cobra.Command, args []string) error {
		v, err := output.VerifyManifest(verifyManifestPath, verifyOutputDir)
		if err != nil {
			return fmt.Errorf("manifest verification failed: %w", err)
		}

		for _, p := range v.Passed {
			fmt.Println(color.Okf("  %-45s checksum OK", p))
		}
		for _, m := range v.Missing {
			fmt.Println(color.Failf("  %-45s MISSING", m))
		}
		for _, f := range v.Failed {
			fmt.Println(color.Fail("  " + f))
		}

		fmt.Printf("\n%s\n", color.Header(v.Summary()))
		if !v.OK() {
			return fmt.Errorf("manifest verification: %s", v.Summary())
		}
		fmt.Println(color.OK("All files match manifest checksums"))
		return nil
	},
}

func init() {
	verifyManifestCmd.Flags().StringVar(&verifyManifestPath, "manifest", "build_manifest.json", "path to build_manifest.json")
	verifyManifestCmd.Flags().StringVar(&verifyOutputDir, "output-dir", ".", "directory containing the build artifacts")
	rootCmd.AddCommand(verifyManifestCmd)
}
