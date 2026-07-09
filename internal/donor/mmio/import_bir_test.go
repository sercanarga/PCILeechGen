package mmio

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestParseJSONTrace_CanonicalRecordBARIndexMustMatchEnvelope(t *testing.T) {
	canonical := []byte(`{
		"schema_version":1,
		"bdf":"0000:03:00.0",
		"bar_index":1,
		"bar_base":4096,
		"bar_size":4096,
		"started_at":"2026-07-09T12:30:00Z",
		"duration_ns":1,
		"records":[{
			"bdf":"0000:03:00.0",
			"bar_index":0,
			"address":4100,
			"offset":4,
			"width":4,
			"operation":"read",
			"value":0,
			"timestamp_ns":1
		}]
	}`)
	if _, err := ParseJSONTrace(bytes.NewReader(canonical)); err == nil {
		t.Fatal("canonical BAR0 record inside a BAR1 envelope should be rejected")
	}
}

func TestParseJSONTrace_LegacyMissingBARIndexInheritsEnvelope(t *testing.T) {
	legacy := []byte(`{
		"BDF":"0000:03:00.0",
		"BARIndex":1,
		"BARBase":4096,
		"BARSize":4096,
		"Duration":1,
		"Records":[{
			"BDF":"0000:03:00.0",
			"Address":4100,
			"Offset":4,
			"Width":4,
			"Type":0,
			"Value":0,
			"Timestamp":1
		}]
	}`)
	trace, err := ParseJSONTrace(bytes.NewReader(legacy))
	if err != nil {
		t.Fatalf("ParseJSONTrace legacy: %v", err)
	}
	if len(trace.Records) != 1 || trace.Records[0].BARIndex != 1 {
		t.Fatalf("legacy record target = %+v, want inherited BAR1", trace.Records)
	}
}

func TestParseJSONTrace_LegacyExplicitBARZeroDoesNotInheritBAROne(t *testing.T) {
	legacy := []byte(`{
		"BDF":"0000:03:00.0",
		"BARIndex":1,
		"BARBase":4096,
		"BARSize":4096,
		"Duration":1,
		"Records":[{
			"BDF":"0000:03:00.0",
			"BARIndex":0,
			"Address":4100,
			"Offset":4,
			"Width":4,
			"Type":0,
			"Value":0,
			"Timestamp":1
		}]
	}`)
	if _, err := ParseJSONTrace(bytes.NewReader(legacy)); err == nil {
		t.Fatal("schema-less legacy BAR0 record inside a BAR1 envelope should be rejected")
	}
}

func TestTraceResultMarshalRejectsRecordBARIndexMismatch(t *testing.T) {
	trace := TraceResult{
		SchemaVersion: TraceSchemaVersion,
		BDF:           "0000:03:00.0",
		BARIndex:      1,
		BARBase:       0x1000,
		BARSize:       0x1000,
		Records: []AccessRecord{{
			BDF: "0000:03:00.0", BARIndex: 0, Address: 0x1004,
			Offset: 4, Width: 4, Type: AccessRead,
		}},
	}
	if _, err := json.Marshal(trace); err == nil {
		t.Fatal("marshal should reject a BAR0 record instead of rewriting it into BAR1")
	}
}
