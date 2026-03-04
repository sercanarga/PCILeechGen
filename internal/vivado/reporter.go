package vivado

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const (
	SeverityInfo     = "INFO"
	SeverityWarning  = "WARNING"
	SeverityCritical = "CRITICAL WARNING"
	SeverityError    = "ERROR"
)

type LogEntry struct {
	Severity string
	Code     string
	Message  string
}

type Report struct {
	Entries        []LogEntry
	Errors         int
	CriticalWarns  int
	Warnings       int
	SynthComplete  bool
	ImplComplete   bool
	BitstreamReady bool
}

var logLineRe = regexp.MustCompile(`^(INFO|WARNING|CRITICAL WARNING|ERROR):\s*\[([^\]]+)\]\s*(.*)`)

// benign warnings we can safely ignore
var benignCodes = map[string]bool{
	"Synth 8-7080":        true,
	"Synth 8-3331":        true,
	"Synth 8-6014":        true,
	"Synth 8-3332":        true,
	"Synth 8-3295":        true,
	"Vivado 12-584":       true,
	"Power 33-332":        true,
	"Constraints 18-5210": true,
	"DRC AVAL-46":         true,
	"DRC REQP-1839":       true,
}

func ParseLogFile(path string) (*Report, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("cannot open Vivado log: %w", err)
	}
	defer f.Close()

	report := &Report{}
	scanner := bufio.NewScanner(f)
	buf := make([]byte, 0, 256*1024)
	scanner.Buffer(buf, 1024*1024)

	for scanner.Scan() {
		line := scanner.Text()
		report.parseLine(line)
	}
	return report, scanner.Err()
}

func ParseOutput(output string) *Report {
	report := &Report{}
	for _, line := range strings.Split(output, "\n") {
		report.parseLine(line)
	}
	return report
}

func (r *Report) parseLine(line string) {
	if strings.Contains(line, "synth_design completed successfully") {
		r.SynthComplete = true
	}
	if strings.Contains(line, "route_design completed successfully") ||
		strings.Contains(line, "place_design completed successfully") {
		r.ImplComplete = true
	}
	if strings.Contains(line, "write_bitstream completed successfully") {
		r.BitstreamReady = true
	}

	matches := logLineRe.FindStringSubmatch(line)
	if matches == nil {
		return
	}

	entry := LogEntry{
		Severity: matches[1],
		Code:     strings.TrimSpace(matches[2]),
		Message:  strings.TrimSpace(matches[3]),
	}
	r.Entries = append(r.Entries, entry)

	switch entry.Severity {
	case SeverityError:
		r.Errors++
	case SeverityCritical:
		r.CriticalWarns++
	case SeverityWarning:
		r.Warnings++
	}
}

func (e *LogEntry) IsBenign() bool {
	return benignCodes[e.Code]
}

// ActionableEntries filters out INFO and known-benign entries.
func (r *Report) ActionableEntries() []LogEntry {
	var result []LogEntry
	for _, e := range r.Entries {
		if e.Severity == SeverityInfo {
			continue
		}
		if e.IsBenign() {
			continue
		}
		result = append(result, e)
	}
	return result
}

func (r *Report) Summary() string {
	var sb strings.Builder

	if r.BitstreamReady {
		sb.WriteString("Build Status: SUCCESS (bitstream generated)\n")
	} else if r.ImplComplete {
		sb.WriteString("Build Status: implementation complete, bitstream pending\n")
	} else if r.SynthComplete {
		sb.WriteString("Build Status: synthesis complete, implementation pending\n")
	} else if r.Errors > 0 {
		sb.WriteString("Build Status: FAILED\n")
	} else {
		sb.WriteString("Build Status: unknown\n")
	}

	sb.WriteString(fmt.Sprintf("Errors: %d, Critical Warnings: %d, Warnings: %d\n",
		r.Errors, r.CriticalWarns, r.Warnings))

	actionable := r.ActionableEntries()
	benign := len(r.Entries) - len(actionable)

	if benign > 0 {
		sb.WriteString(fmt.Sprintf("(%d benign warnings filtered)\n", benign))
	}

	if len(actionable) > 0 {
		sb.WriteString("\nActionable issues:\n")
		shown := 0
		for _, e := range actionable {
			if shown >= 20 {
				sb.WriteString(fmt.Sprintf("  ... and %d more\n", len(actionable)-shown))
				break
			}
			sb.WriteString(fmt.Sprintf("  [%s] %s: %s\n", e.Severity, e.Code, e.Message))
			shown++
		}
	}

	return sb.String()
}
