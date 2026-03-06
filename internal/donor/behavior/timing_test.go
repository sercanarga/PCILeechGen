package behavior

import (
	"testing"
	"time"

	"github.com/sercanarga/pcileechgen/internal/donor/mmio"
)

func timingTrace() *mmio.TraceResult {
	return &mmio.TraceResult{
		BDF:      "0000:03:00.0",
		BARIndex: 0,
		BARSize:  4096,
		Duration: 100 * time.Millisecond,
		Records: []mmio.AccessRecord{
			// Fast reads (100-200ns apart → ~12-25 cycles at 125MHz)
			{Offset: 0x00, Type: mmio.AccessRead, Value: 0x0040FF17, Timestamp: 100 * time.Nanosecond},
			{Offset: 0x08, Type: mmio.AccessRead, Value: 0x00010400, Timestamp: 220 * time.Nanosecond},
			{Offset: 0x00, Type: mmio.AccessRead, Value: 0x0040FF17, Timestamp: 350 * time.Nanosecond},
			{Offset: 0x1C, Type: mmio.AccessRead, Value: 0x00000000, Timestamp: 500 * time.Nanosecond},
			{Offset: 0x1C, Type: mmio.AccessRead, Value: 0x00000000, Timestamp: 680 * time.Nanosecond},
			// Slightly slower reads
			{Offset: 0x1C, Type: mmio.AccessRead, Value: 0x00000001, Timestamp: 1100 * time.Nanosecond},
			{Offset: 0x00, Type: mmio.AccessRead, Value: 0x0040FF17, Timestamp: 1500 * time.Nanosecond},
			{Offset: 0x08, Type: mmio.AccessRead, Value: 0x00010400, Timestamp: 1700 * time.Nanosecond},
			// Interspersed write (should be ignored for timing)
			{Offset: 0x14, Type: mmio.AccessWrite, Value: 0x00460001, Timestamp: 1800 * time.Nanosecond},
			{Offset: 0x1C, Type: mmio.AccessRead, Value: 0x00000001, Timestamp: 1900 * time.Nanosecond},
		},
	}
}

func TestExtractTimingHistogram_Basic(t *testing.T) {
	h := ExtractTimingHistogram(timingTrace())
	if h == nil {
		t.Fatal("returned nil")
	}
	if h.SampleCount == 0 {
		t.Error("should have timing samples")
	}
	if h.MinCycles <= 0 {
		t.Errorf("MinCycles should be > 0, got %d", h.MinCycles)
	}
	if h.MaxCycles < h.MinCycles {
		t.Errorf("MaxCycles (%d) < MinCycles (%d)", h.MaxCycles, h.MinCycles)
	}
}

func TestExtractTimingHistogram_CDF(t *testing.T) {
	h := ExtractTimingHistogram(timingTrace())
	// CDF must be monotonically non-decreasing
	for i := 1; i < 16; i++ {
		if h.CDF[i] < h.CDF[i-1] {
			t.Errorf("CDF not monotonic at %d: %d < %d", i, h.CDF[i], h.CDF[i-1])
		}
	}
	// Last bucket must be 255
	if h.CDF[15] != 255 {
		t.Errorf("CDF[15] should be 255, got %d", h.CDF[15])
	}
}

func TestExtractTimingHistogram_Nil(t *testing.T) {
	h := ExtractTimingHistogram(nil)
	if h == nil {
		t.Fatal("nil trace should return default, not nil")
	}
	if h.SampleCount != 0 {
		t.Error("nil trace should have 0 samples")
	}
	if h.CDF[15] != 255 {
		t.Errorf("default CDF[15] should be 255, got %d", h.CDF[15])
	}
}

func TestExtractTimingHistogram_TooFewRecords(t *testing.T) {
	trace := &mmio.TraceResult{
		Records: []mmio.AccessRecord{
			{Offset: 0x00, Type: mmio.AccessRead, Timestamp: 100 * time.Nanosecond},
		},
	}
	h := ExtractTimingHistogram(trace)
	if h.SampleCount != 0 {
		t.Error("too few records should return default histogram")
	}
}

func TestExtractTimingHistogram_NonZeroBuckets(t *testing.T) {
	h := ExtractTimingHistogram(timingTrace())
	hasNonZero := false
	for _, b := range h.Buckets {
		if b > 0 {
			hasNonZero = true
			break
		}
	}
	if !hasNonZero {
		t.Error("histogram should have at least one non-zero bucket")
	}
}

func TestBuildCDF_Uniform(t *testing.T) {
	var buckets [16]uint8
	for i := range buckets {
		buckets[i] = 16
	}
	cdf := buildCDF(buckets)
	if cdf[15] != 255 {
		t.Errorf("CDF[15] should be 255, got %d", cdf[15])
	}
	// Should be roughly linear
	if cdf[7] < 100 || cdf[7] > 140 {
		t.Errorf("uniform CDF[7] should be ~127, got %d", cdf[7])
	}
}

func TestBuildCDF_Empty(t *testing.T) {
	var buckets [16]uint8
	cdf := buildCDF(buckets)
	// Empty → uniform fallback
	if cdf[15] != 255 {
		t.Errorf("empty CDF[15] should be 255, got %d", cdf[15])
	}
}

func TestDefaultHistogram(t *testing.T) {
	h := defaultHistogram()
	if h.MinCycles != 3 {
		t.Errorf("default MinCycles: got %d, want 3", h.MinCycles)
	}
	if h.MaxCycles != 15 {
		t.Errorf("default MaxCycles: got %d, want 15", h.MaxCycles)
	}
}
