package mmio

import (
	"bytes"
	"encoding/json"
	"math"
	"strings"
	"testing"
	"time"
)

func TestParseTextTrace_PreservesWideOffsetsWidthsAndValues(t *testing.T) {
	input := strings.NewReader(strings.Join([]string{
		"R 1 1.000 0xf7801001 0x7f",
		"W 2 1.010 0xf7801002 0xbeef",
		"R 4 1.020 0xf7801004 0x89abcdef",
		"W 8 1.030 0xf7801008 0x1122334455667788",
	}, "\n"))

	trace, err := ParseTextTrace(input, TextTraceOptions{
		BDF:      "0000:03:00.0",
		BARIndex: 2,
		BARBase:  0xf7800000,
		BARSize:  0x4000,
	})
	if err != nil {
		t.Fatalf("ParseTextTrace returned error: %v", err)
	}

	want := []AccessRecord{
		{BDF: "0000:03:00.0", BARIndex: 2, Address: 0xf7801001, Offset: 0x1001, Width: 1, Type: AccessRead, Value: 0x7f},
		{BDF: "0000:03:00.0", BARIndex: 2, Address: 0xf7801002, Offset: 0x1002, Width: 2, Type: AccessWrite, Value: 0xbeef},
		{BDF: "0000:03:00.0", BARIndex: 2, Address: 0xf7801004, Offset: 0x1004, Width: 4, Type: AccessRead, Value: 0x89abcdef},
		{BDF: "0000:03:00.0", BARIndex: 2, Address: 0xf7801008, Offset: 0x1008, Width: 8, Type: AccessWrite, Value: 0x1122334455667788},
	}
	if len(trace.Records) != len(want) {
		t.Fatalf("records = %d, want %d: %+v", len(trace.Records), len(want), trace.Records)
	}
	for i := range want {
		got := trace.Records[i]
		if got.BDF != want[i].BDF || got.BARIndex != want[i].BARIndex || got.Address != want[i].Address ||
			got.Offset != want[i].Offset || got.Width != want[i].Width || got.Type != want[i].Type || got.Value != want[i].Value {
			t.Errorf("record %d = %+v, want %+v", i, got, want[i])
		}
	}
}

func TestParseTextTrace_RejectsAddressesOutsideBARWithoutAliasing(t *testing.T) {
	const base = uint64(0xf7800000)
	cases := map[string]string{
		"below base":       "R 4 1.000 0xf77ffffc 0xaaaaaaaa\n",
		"crossing end":     "R 8 1.020 0xf7800ffc 0xbbbbbbbb\n",
		"at end":           "R 4 1.030 0xf7801000 0xcccccccc\n",
		"unrelated region": "R 4 1.040 0xf790010c 0xdddddddd\n",
	}
	for name, input := range cases {
		t.Run(name, func(t *testing.T) {
			if _, err := ParseTextTrace(strings.NewReader(input), TextTraceOptions{BARBase: base, BARSize: 0x1000}); err == nil {
				t.Fatal("expected out-of-aperture record to be rejected instead of aliased")
			}
		})
	}
}

func TestParseTextTrace_RejectsOverflowingBARRange(t *testing.T) {
	input := strings.NewReader("R 4 1.000 0xfffffffffffffffc 0x1\n")
	_, err := ParseTextTrace(input, TextTraceOptions{
		BARBase: math.MaxUint64 - 3,
		BARSize: 8,
	})
	if err == nil {
		t.Fatal("expected overflowing BAR base + size to be rejected")
	}
}

func TestParseTextTrace_RejectsUnsupportedWidths(t *testing.T) {
	for _, width := range []string{"0", "3", "16"} {
		t.Run(width, func(t *testing.T) {
			input := strings.NewReader("R " + width + " 1.000 0xf7800000 0x1\n")
			if _, err := ParseTextTrace(input, TextTraceOptions{BARBase: 0xf7800000, BARSize: 0x1000}); err == nil {
				t.Fatalf("expected width %s to be rejected", width)
			}
		})
	}
}

func TestAccessRecord_LegacyJSONDefaultsToDWordWidth(t *testing.T) {
	legacy := []byte(`{"Offset":4660,"Type":0,"Value":2309737967,"Timestamp":1000}`)
	var record AccessRecord
	if err := json.Unmarshal(legacy, &record); err != nil {
		t.Fatalf("unmarshal legacy record: %v", err)
	}
	if record.Offset != 0x1234 || record.Width != 4 || record.Type != AccessRead || record.Value != 0x89abcdef {
		t.Fatalf("legacy record = %+v, want offset/value preserved and width defaulted to 4", record)
	}
}

func TestTraceResult_CanonicalV1JSONRoundTripIsDeterministic(t *testing.T) {
	start := time.Date(2026, time.July, 9, 12, 30, 0, 123, time.UTC)
	trace := &TraceResult{
		SchemaVersion: TraceSchemaVersion,
		BDF:           "0000:03:00.0",
		BARIndex:      2,
		BARBase:       0xf7800000,
		BARSize:       0x4000,
		Duration:      25 * time.Microsecond,
		StartTime:     start,
		Records: []AccessRecord{{
			BDF: "0000:03:00.0", BARIndex: 2, Address: 0xf7801008, Offset: 0x1008,
			Width: 8, Type: AccessWrite, Value: 0x1122334455667788, Timestamp: 25 * time.Microsecond,
		}},
	}
	first, err := json.Marshal(trace)
	if err != nil {
		t.Fatalf("marshal canonical trace: %v", err)
	}
	second, err := json.Marshal(trace)
	if err != nil {
		t.Fatalf("marshal canonical trace again: %v", err)
	}
	if !bytes.Equal(first, second) {
		t.Fatalf("canonical trace JSON is nondeterministic:\nfirst:  %s\nsecond: %s", first, second)
	}

	var schema map[string]any
	if err := json.Unmarshal(first, &schema); err != nil {
		t.Fatalf("decode canonical trace schema: %v", err)
	}
	for _, key := range []string{"schema_version", "bdf", "bar_index", "bar_base", "bar_size", "duration_ns", "started_at", "records"} {
		if _, ok := schema[key]; !ok {
			t.Errorf("canonical trace JSON missing field %q: %v", key, schema)
		}
	}

	roundTrip, err := ParseJSONTrace(bytes.NewReader(first))
	if err != nil {
		t.Fatalf("ParseJSONTrace canonical v1: %v", err)
	}
	if roundTrip.SchemaVersion != TraceSchemaVersion || roundTrip.BDF != trace.BDF ||
		roundTrip.BARIndex != trace.BARIndex || roundTrip.BARBase != trace.BARBase ||
		roundTrip.BARSize != trace.BARSize || roundTrip.Duration != trace.Duration ||
		!roundTrip.StartTime.Equal(trace.StartTime) || len(roundTrip.Records) != 1 {
		t.Fatalf("canonical round trip = %+v, want %+v", roundTrip, trace)
	}
	got := roundTrip.Records[0]
	want := trace.Records[0]
	if got.BDF != want.BDF || got.BARIndex != want.BARIndex || got.Address != want.Address ||
		got.Offset != want.Offset || got.Width != want.Width || got.Type != want.Type ||
		got.Value != want.Value || got.Timestamp != want.Timestamp {
		t.Fatalf("canonical record round trip = %+v, want %+v", got, want)
	}
}

func TestParseJSONTrace_LegacyPascalCaseDefaultsWidths(t *testing.T) {
	legacy := []byte(`{
		"BDF":"0000:03:00.0",
		"BARIndex":2,
		"BARSize":16384,
		"Duration":25000,
		"StartTime":"2026-07-09T12:30:00Z",
		"Records":[{
			"BDF":"0000:03:00.0",
			"BARIndex":2,
			"Address":4152360964,
			"Offset":4100,
			"Type":0,
			"Value":2309737967,
			"Timestamp":25000
		}]
	}`)
	trace, err := ParseJSONTrace(bytes.NewReader(legacy))
	if err != nil {
		t.Fatalf("ParseJSONTrace legacy: %v", err)
	}
	if trace.SchemaVersion != TraceSchemaVersion || trace.BDF != "0000:03:00.0" ||
		trace.BARIndex != 2 || trace.BARSize != 16384 || trace.Duration != 25*time.Microsecond ||
		len(trace.Records) != 1 {
		t.Fatalf("legacy trace = %+v", trace)
	}
	if got := trace.Records[0]; got.Width != 4 || got.Offset != 0x1004 ||
		got.Value != 0x89abcdef || got.BDF != trace.BDF || got.BARIndex != trace.BARIndex {
		t.Fatalf("legacy record = %+v, want inherited metadata and dword width", got)
	}
}
