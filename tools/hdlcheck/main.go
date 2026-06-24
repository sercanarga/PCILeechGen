package main

import (
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
)

type fileSet struct {
	HDL []string
	TCL []string
}

type commandPlan struct {
	Label string
	Args  []string
}

type plan struct {
	Commands []commandPlan
	Skipped  []string
}

type lookPathFunc func(string) (string, error)

func main() {
	dir := flag.String("dir", "pcileech_datastore", "directory containing generated SV/V/TCL artifacts")
	dryRun := flag.Bool("dry-run", false, "print checks without running external tools")
	flag.Parse()

	files, err := discoverFiles(*dir)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	p := buildPlan(files, exec.LookPath)
	for _, skipped := range p.Skipped {
		fmt.Println("SKIP", skipped)
	}
	if len(p.Commands) == 0 {
		fmt.Println("No HDL/TCL checker available; install verilator or vivado for syntax checks.")
		return
	}
	for _, cmd := range p.Commands {
		fmt.Println(strings.Join(cmd.Args, " "))
		if *dryRun {
			continue
		}
		proc := exec.Command(cmd.Args[0], cmd.Args[1:]...)
		proc.Stdout = os.Stdout
		proc.Stderr = os.Stderr
		if err := proc.Run(); err != nil {
			fmt.Fprintln(os.Stderr, cmd.Label, "failed:", err)
			os.Exit(1)
		}
	}
}

func discoverFiles(root string) (fileSet, error) {
	var files fileSet
	if root == "" {
		return files, errors.New("--dir must not be empty")
	}
	if err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		switch strings.ToLower(filepath.Ext(path)) {
		case ".sv", ".v":
			files.HDL = append(files.HDL, path)
		case ".tcl":
			files.TCL = append(files.TCL, path)
		}
		return nil
	}); err != nil {
		return files, fmt.Errorf("scan %s: %w", root, err)
	}
	sort.Strings(files.HDL)
	sort.Strings(files.TCL)
	return files, nil
}

func buildPlan(files fileSet, lookPath lookPathFunc) plan {
	var p plan
	if len(files.HDL) > 0 {
		if bin, err := lookPath("verilator"); err == nil {
			args := append([]string{bin, "--lint-only"}, files.HDL...)
			p.Commands = append(p.Commands, commandPlan{Label: "verilator", Args: args})
		} else {
			p.Skipped = append(p.Skipped, "verilator not found for HDL lint")
		}
	}
	if len(files.TCL) > 0 {
		p.Skipped = append(p.Skipped, "TCL files found; Vivado parse/build scripts are not run automatically")
	}
	return p
}
