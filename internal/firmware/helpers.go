package firmware

import (
	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/donor"
)

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

func ComputeBAR0Size(msixTableSize int, bramLimit int) int {
	if msixTableSize <= 0 {
		if bramLimit > 0 {
			return bramLimit
		}
		return board.DefaultBRAMSize
	}
	size := MSIXRequiredBAR0Size(msixTableSize)
	if bramLimit > 0 && size > bramLimit {
		return bramLimit
	}
	return size
}

// MSIXRequiredBAR0Size returns the BAR0 size required for the given MSI-X table
// (starting from 4K doubling up to fit 0x2000 + table + PBA), WITHOUT applying
// any board BRAM limit. Used to determine donor demand for pre-cap checks.
func MSIXRequiredBAR0Size(msixTableSize int) int {
	if msixTableSize <= 0 {
		return board.DefaultBRAMSize
	}
	tableBytes := msixTableSize * 16
	pbaBytes := (msixTableSize + 63) / 64 * 8
	if pbaBytes < 8 {
		pbaBytes = 8
	}
	required := 0x2000 + tableBytes + pbaBytes
	size := board.DefaultBRAMSize
	for size < required {
		size *= 2
	}
	return size
}

// DonorBAR0Demand computes the BAR0 size "demanded" by the donor context
// (max of: board default, BAR register sizes from ctx, actual BARContents lengths,
// and the uncapped MSIX required size). This value may exceed the board's BRAM.
// Callers use it for compatibility gates (error unless --force); the final
// size used for scrubbing/generation is always CappedBAR0Size (min(demand, BRAM)).
func DonorBAR0Demand(ctx *donor.DeviceContext, b *board.Board, msixTableSize int) int {
	bram := board.DefaultBRAMSize
	if b != nil {
		bram = b.BRAMSizeOrDefault()
	}
	demand := 0
	if msixTableSize > 0 {
		demand = MSIXRequiredBAR0Size(msixTableSize)
	}
	if ctx != nil {
		if d := LargestBar(ctx.BARContents); d != nil && len(d) > demand {
			demand = len(d)
		}
		// Also consider declared BAR sizes (from resource or parsed); catches
		// cases where contents may be partial but BAR register encoded a large size.
		for _, bar := range ctx.BARs {
			if !bar.IsDisabled() && int(bar.Size) > demand {
				demand = int(bar.Size)
			}
		}
	}
	if demand == 0 {
		demand = bram
	}
	return demand
}

func CappedBAR0Size(ctx *donor.DeviceContext, b *board.Board, msixTableSize int) int {
	bram := board.DefaultBRAMSize
	if b != nil {
		bram = b.BRAMSizeOrDefault()
	}
	demand := DonorBAR0Demand(ctx, b, msixTableSize)
	if demand > bram {
		return bram
	}
	return demand
}

func MSIXPlacement(bar0Size int, msixTableSize int, class uint32, dstrd uint32) (uint32, uint32, uint32) {
	tableBytes := msixTableSize * 16
	isNVMe := class>>8 == 0x0108
	// dbBase uses board.DefaultBRAMSize (classic 0x1000) so doorbells + post-doorbell MSIX
	// placement for variable BAR0 (16k NVMe 010802 etc) is consistent with Capped/Compute.
	dbBase := uint32(board.DefaultBRAMSize)
	tableOff := dbBase
	if bar0Size > 0 {
		tableOff = uint32(bar0Size/2) &^ 0xF
		if tableOff < 0x2000 {
			tableOff = 0x2000
		}
		if tableOff >= dbBase && tableOff < dbBase+uint32(tableBytes) {
			tableOff = 0x2000
		}
		if tableOff < 0x40 {
			tableOff = dbBase
		}
		if tableOff+uint32(tableBytes)+16 > uint32(bar0Size) {
			tableOff = uint32(bar0Size) - uint32(tableBytes) - 16
			tableOff &^= 0xF
			if tableOff < dbBase {
				tableOff = dbBase
			}
		}
	}
	if isNVMe {
		stride := uint32(4) << dstrd
		dbEnd := dbBase + 2*stride
		if tableOff < dbEnd {
			tableOff = dbEnd
		}
		tableOff = (tableOff + 0xF) &^ 0xF
		if tableOff+uint32(tableBytes)+16 > uint32(bar0Size) {
			tableOff = uint32(bar0Size) - uint32(tableBytes) - 16
			tableOff &^= 0xF
			if tableOff < dbEnd {
				tableOff = dbEnd
			}
			tableOff = (tableOff + 0xF) &^ 0xF
			if tableOff < dbEnd {
				tableOff = dbEnd
			}
		}
	}
	pbaOff := tableOff + uint32(tableBytes)
	pbaOff = (pbaOff + 7) &^ 7
	return tableOff, pbaOff, dbBase
}
