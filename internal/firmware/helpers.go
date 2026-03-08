package firmware

// LowestBar picks the value with the smallest map key.
// Returns the zero value of V when the map is nil or empty.
func LowestBar[V any](m map[int]V) V {
	bestIdx := -1
	for idx := range m {
		if bestIdx == -1 || idx < bestIdx {
			bestIdx = idx
		}
	}
	if bestIdx < 0 {
		var zero V
		return zero
	}
	return m[bestIdx]
}

// LargestBar returns the longest byte slice in the map.
// Needed when the main MMIO BAR isn't at index 0.
func LargestBar(m map[int][]byte) []byte {
	var best []byte
	for _, v := range m {
		if len(v) > len(best) {
			best = v
		}
	}
	return best
}

// LargestBarIndex returns the index of the longest byte slice in the map.
// Used when both BAR content and probe profile must reference the same BAR.
func LargestBarIndex(m map[int][]byte) int {
	bestIdx := 0
	bestLen := 0
	for idx, v := range m {
		if len(v) > bestLen {
			bestLen = len(v)
			bestIdx = idx
		}
	}
	return bestIdx
}
