package mmio

import (
	"strings"
	"testing"
	"time"
)

func TestParseTextTrace_Ret2cShapeWithBARBase(t *testing.T) {
	input := strings.NewReader(strings.Join([]string{
		"R 4 2456.105919 2 0xf780010c 0x4c02 0x0 0",
		"W 4 2456.130642 2 0xf7800114 0x1 0x0 0",
	}, "\n"))

	trace, err := ParseTextTrace(input, TextTraceOptions{
		BDF:      "0000:03:00.0",
		BARIndex: 2,
		BARSize:  4096,
		BARBase:  0xf7800000,
	})

	if err != nil {
		t.Fatalf("ParseTextTrace returned error: %v", err)
	}
	if len(trace.Records) != 2 {
		t.Fatalf("records = %d, want 2", len(trace.Records))
	}
	if trace.Records[0].Type != AccessRead || trace.Records[0].Offset != 0x10c || trace.Records[0].Value != 0x4c02 {
		t.Fatalf("first record = %+v", trace.Records[0])
	}
	if trace.Records[1].Type != AccessWrite || trace.Records[1].Offset != 0x114 || trace.Records[1].Value != 0x1 {
		t.Fatalf("second record = %+v", trace.Records[1])
	}
	if trace.Duration <= 0 {
		t.Fatal("duration should be derived from timestamps")
	}
}

func TestParseTextTrace_LiveTracePipeShape(t *testing.T) {
	input := strings.NewReader("R 4 1234567.890 0xfee00100 0x00000001 extra\n")

	trace, err := ParseTextTrace(input, TextTraceOptions{BARSize: 4096})

	if err != nil {
		t.Fatalf("ParseTextTrace returned error: %v", err)
	}
	if len(trace.Records) != 1 {
		t.Fatalf("records = %d, want 1", len(trace.Records))
	}
	if trace.Records[0].Offset != 0x100 || trace.Records[0].Value != 1 {
		t.Fatalf("record = %+v", trace.Records[0])
	}
}

func TestParseTextTrace_RejectsEmptyTrace(t *testing.T) {
	_, err := ParseTextTrace(strings.NewReader("not a trace line\n"), TextTraceOptions{})
	if err == nil {
		t.Fatal("expected error for empty trace")
	}
}

func TestParseTextTrace_DurationUsesLastTimestamp(t *testing.T) {
	input := strings.NewReader(strings.Join([]string{
		"R 4 1.000 0x1000 0x1",
		"R 4 1.250 0x1004 0x2",
	}, "\n"))

	trace, err := ParseTextTrace(input, TextTraceOptions{})

	if err != nil {
		t.Fatalf("ParseTextTrace returned error: %v", err)
	}
	if trace.Duration != 250*time.Millisecond {
		t.Fatalf("duration = %s, want 250ms", trace.Duration)
	}
}
