package behavior

import (
	"math"
	"sort"

	"github.com/sercanarga/pcileechgen/internal/donor/mmio"
)

// TimingHistogram is a 16-bucket latency distribution from donor MMIO traces.
type TimingHistogram struct {
	Buckets      [16]uint8 // normalized weights (0-255)
	CDF          [16]uint8 // cumulative distribution
	MinCycles    int
	MaxCycles    int
	MedianCycles int
	SampleCount  int
}

const fpgaClockPeriodNs = 8 // 125MHz

// ExtractTimingHistogram builds a 16-bucket histogram from inter-read intervals.
func ExtractTimingHistogram(trace *mmio.TraceResult) *TimingHistogram {
	if trace == nil || len(trace.Records) < 2 {
		return defaultHistogram()
	}

	var intervals []int64
	var lastReadNs int64 = -1

	for _, rec := range trace.Records {
		if rec.Type != mmio.AccessRead {
			continue
		}
		ns := rec.Timestamp.Nanoseconds()
		if lastReadNs >= 0 {
			delta := ns - lastReadNs
			if delta > 0 && delta < 100_000 { // skip driver pauses >100µs
				intervals = append(intervals, delta)
			}
		}
		lastReadNs = ns
	}

	if len(intervals) < 3 {
		return defaultHistogram()
	}

	sort.Slice(intervals, func(i, j int) bool { return intervals[i] < intervals[j] })

	minCycles := int(math.Max(1, float64(intervals[0])/fpgaClockPeriodNs))
	maxCycles := int(math.Max(float64(minCycles)+1, float64(intervals[len(intervals)-1])/fpgaClockPeriodNs))
	medianCycles := int(math.Max(float64(minCycles), float64(intervals[len(intervals)/2])/fpgaClockPeriodNs))

	if maxCycles > 255 {
		maxCycles = 255
	}
	if medianCycles > maxCycles {
		medianCycles = maxCycles
	}

	rangeSize := maxCycles - minCycles + 1
	bucketWidth := float64(rangeSize) / 16.0
	if bucketWidth < 1 {
		bucketWidth = 1
	}

	var rawBuckets [16]int
	for _, ns := range intervals {
		cycles := float64(ns) / fpgaClockPeriodNs
		idx := int((cycles - float64(minCycles)) / bucketWidth)
		if idx < 0 {
			idx = 0
		}
		if idx > 15 {
			idx = 15
		}
		rawBuckets[idx]++
	}

	maxCount := 0
	for _, c := range rawBuckets {
		if c > maxCount {
			maxCount = c
		}
	}

	var buckets [16]uint8
	if maxCount > 0 {
		for i, c := range rawBuckets {
			buckets[i] = uint8((c * 255) / maxCount)
			if c > 0 && buckets[i] == 0 {
				buckets[i] = 1
			}
		}
	}

	return &TimingHistogram{
		Buckets:      buckets,
		CDF:          buildCDF(buckets),
		MinCycles:    minCycles,
		MaxCycles:    maxCycles,
		MedianCycles: medianCycles,
		SampleCount:  len(intervals),
	}
}

// buildCDF converts weight histogram to cumulative distribution [0-255].
func buildCDF(buckets [16]uint8) [16]uint8 {
	var sum int
	for _, b := range buckets {
		sum += int(b)
	}
	if sum == 0 {
		var cdf [16]uint8
		for i := range cdf {
			cdf[i] = uint8(((i + 1) * 255) / 16)
		}
		return cdf
	}

	var cdf [16]uint8
	var cumulative int
	for i, b := range buckets {
		cumulative += int(b)
		cdf[i] = uint8((cumulative * 255) / sum)
	}
	cdf[15] = 255
	return cdf
}

func defaultHistogram() *TimingHistogram {
	var buckets [16]uint8
	for i := range buckets {
		buckets[i] = 16
	}
	return &TimingHistogram{
		Buckets:      buckets,
		CDF:          buildCDF(buckets),
		MinCycles:    3,
		MaxCycles:    15,
		MedianCycles: 7,
		SampleCount:  0,
	}
}
