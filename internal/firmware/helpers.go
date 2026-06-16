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
	if bramLimit > 0 && size > bramLimit {
		return bramLimit
	}
	return size
}

func CappedBAR0Size(ctx *donor.DeviceContext, b *board.Board, msixTableSize int) int {
	bram := board.DefaultBRAMSize
	if b != nil {
		bram = b.BRAMSizeOrDefault()
	}
	bar0Size := bram
	if msixTableSize > 0 {
		bar0Size = ComputeBAR0Size(msixTableSize, bram)
	}
	if ctx != nil {
		if d := LargestBar(ctx.BARContents); d != nil && len(d) > bar0Size {
			bar0Size = len(d)
		}
	}
	if bar0Size > bram {
		bar0Size = bram
	}
	return bar0Size
}

func MSIXPlacement(bar0Size int, msixTableSize int, class uint32, dstrd uint32) (uint32, uint32, uint32) {
	tableBytes := msixTableSize * 16
	isNVMe := class>>8 == 0x0108
	dbBase := uint32(board.DefaultBRAMSize)
	tableOff := uint32(board.DefaultBRAMSize)
	if bar0Size > 0 {
		tableOff = uint32(bar0Size/2) &^ 0xF
		if tableOff < 0x2000 {
			tableOff = 0x2000
		}
		if tableOff >= board.DefaultBRAMSize && tableOff < board.DefaultBRAMSize+uint32(tableBytes) {
			tableOff = 0x2000
		}
		if tableOff < 0x40 {
			tableOff = board.DefaultBRAMSize
		}
		if tableOff+uint32(tableBytes)+16 > uint32(bar0Size) {
			tableOff = uint32(bar0Size) - uint32(tableBytes) - 16
			tableOff &^= 0xF
			if tableOff < board.DefaultBRAMSize {
				tableOff = board.DefaultBRAMSize
			}
		}
	}
	if isNVMe {
		stride := uint32(4) << dstrd
		dbEnd := uint32(board.DefaultBRAMSize) + 2*stride
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
			if tableOff < board.DefaultBRAMSize {
				tableOff = board.DefaultBRAMSize
			}
		}
	}
	pbaOff := tableOff + uint32(tableBytes)
	pbaOff = (pbaOff + 7) &^ 7
	return tableOff, pbaOff, dbBase
}
