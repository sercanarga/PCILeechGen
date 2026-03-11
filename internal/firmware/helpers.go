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

// countNonZero returns the number of non-zero bytes in a slice.
func countNonZero(b []byte) int {
	n := 0
	for _, v := range b {
		if v != 0 {
			n++
		}
	}
	return n
}

// LargestBar returns the longest byte slice in the map.
// When sizes are equal, picks the BAR with the most non-zero bytes.
func LargestBar(m map[int][]byte) []byte {
	var best []byte
	bestNZ := 0
	for _, v := range m {
		nz := countNonZero(v)
		if len(v) > len(best) || (len(v) == len(best) && nz > bestNZ) {
			best = v
			bestNZ = nz
		}
	}
	return best
}

// LargestBarIndex returns the index of the longest byte slice in the map.
// When sizes are equal, picks the BAR with the most non-zero bytes.
func LargestBarIndex(m map[int][]byte) int {
	bestIdx := 0
	bestLen := 0
	bestNZ := 0
	for idx, v := range m {
		nz := countNonZero(v)
		if len(v) > bestLen || (len(v) == bestLen && nz > bestNZ) {
			bestLen = len(v)
			bestIdx = idx
			bestNZ = nz
		}
	}
	return bestIdx
}
