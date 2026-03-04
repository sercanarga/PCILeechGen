package firmware

import (
	"github.com/sercanarga/pcileechgen/internal/donor"
)

// LowestBarData picks the smallest BAR index that has data.
func LowestBarData(barContents map[int][]byte) []byte {
	bestIdx := -1
	for idx := range barContents {
		if bestIdx == -1 || idx < bestIdx {
			bestIdx = idx
		}
	}
	if bestIdx < 0 {
		return nil
	}
	return barContents[bestIdx]
}

// LowestBarProfile picks the smallest BAR index that has a profile.
func LowestBarProfile(barProfiles map[int]*donor.BARProfile) *donor.BARProfile {
	bestIdx := -1
	for idx := range barProfiles {
		if bestIdx == -1 || idx < bestIdx {
			bestIdx = idx
		}
	}
	if bestIdx < 0 {
		return nil
	}
	return barProfiles[bestIdx]
}
