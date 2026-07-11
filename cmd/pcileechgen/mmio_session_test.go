package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/donor/session"
)

func TestPrintInitializationAnalysisShowsEvidenceAndCoverage(t *testing.T) {
	analysis := session.Analysis{
		SessionCount: 2,
		Phases: []session.Phase{{Session: 0, Index: 0, Kind: "command-poll", Records: 5}},
		Registers: []session.RegisterEvidence{{Offset: 0x37, Reads: 1, Writes: 1, Classification: "read-write", Confidence: session.ConfidenceObserved, WriteEffect: session.WriteEffectSelfClearing}},
		Dependencies: []session.Dependency{{WriteOffset: 0x37, ReadOffset: 0x3e, Occurrences: 2}},
	}
	var out bytes.Buffer

	printInitializationAnalysis(&out, analysis)

	text := out.String()
	for _, want := range []string{"Initialization Session Analysis", "Sessions: 2", "0x00000037", "self-clearing", "0x00000037 -> 0x0000003e"} {
		if !strings.Contains(text, want) {
			t.Fatalf("output missing %q:\n%s", want, text)
		}
	}
}
