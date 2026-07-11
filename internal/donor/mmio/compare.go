package mmio

import "sort"

// RegisterClass describes behavior observed across MMIO traces.
type RegisterClass string

const (
	RegisterStable              RegisterClass = "stable constant"
	RegisterChanging            RegisterClass = "changing"
	RegisterReadAfterWrite      RegisterClass = "read after write"
	RegisterInitializationWrite RegisterClass = "initialization write"
)

// RegisterComparison summarizes one register across multiple traces.
type RegisterComparison struct {
	Offset         uint32        `json:"offset"`
	Reads          int           `json:"reads"`
	Writes         int           `json:"writes"`
	Values         []uint32      `json:"values"`
	Classification RegisterClass `json:"classification"`
}

// CompareTraces classifies register behavior observed across trace captures.
func CompareTraces(traces []*TraceResult) []RegisterComparison {
	type aggregate struct {
		reads          int
		writes         int
		readAfterWrite bool
		values         map[uint32]struct{}
	}
	byOffset := make(map[uint32]*aggregate)
	for _, trace := range traces {
		if trace == nil {
			continue
		}
		written := make(map[uint32]bool)
		for _, record := range trace.Records {
			a := byOffset[record.Offset]
			if a == nil {
				a = &aggregate{values: make(map[uint32]struct{})}
				byOffset[record.Offset] = a
			}
			a.values[record.Value] = struct{}{}
			if record.Type == AccessWrite {
				a.writes++
				written[record.Offset] = true
			} else {
				a.reads++
				a.readAfterWrite = a.readAfterWrite || written[record.Offset]
			}
		}
	}

	result := make([]RegisterComparison, 0, len(byOffset))
	for offset, a := range byOffset {
		values := make([]uint32, 0, len(a.values))
		for value := range a.values {
			values = append(values, value)
		}
		sort.Slice(values, func(i, j int) bool { return values[i] < values[j] })
		class := RegisterStable
		switch {
		case a.readAfterWrite:
			class = RegisterReadAfterWrite
		case a.writes > 0 && a.reads == 0:
			class = RegisterInitializationWrite
		case len(values) > 1:
			class = RegisterChanging
		}
		result = append(result, RegisterComparison{Offset: offset, Reads: a.reads, Writes: a.writes, Values: values, Classification: class})
	}
	sort.Slice(result, func(i, j int) bool { return result[i].Offset < result[j].Offset })
	return result
}
