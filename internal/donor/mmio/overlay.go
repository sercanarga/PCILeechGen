package mmio

import "fmt"

// TraceBAROverlay is the conservative, codegen-friendly subset of a live BAR
// trace: exact repeated-read sequences for some offsets and stable static values
// for others.
type TraceBAROverlay struct {
	Sequential map[uint32][]uint32 `json:"sequential,omitempty"`
	Static     map[uint32]uint32   `json:"static,omitempty"`
	WriteMask  map[uint32]uint32   `json:"write_mask,omitempty"`
	RW1CMask   map[uint32]uint32   `json:"rw1c_mask,omitempty"`
}

// DeriveTraceBAROverlay converts a live trace into a conservative overlay.
//
// Rules:
//   - only 32-bit aligned offsets are considered
//   - identical repeated reads become Static
//   - changing repeated reads become Sequential when the sequence is bounded
func DeriveTraceBAROverlay(trace *TraceResult) *TraceBAROverlay {
	out := &TraceBAROverlay{
		Sequential: map[uint32][]uint32{},
		Static:     map[uint32]uint32{},
		WriteMask:  map[uint32]uint32{},
		RW1CMask:   map[uint32]uint32{},
	}
	if trace == nil || len(trace.Records) == 0 {
		return out
	}

	type bucket struct {
		reads      []uint32
		readValues map[uint32]struct{}
		writes     []uint32
	}
	byOffset := map[uint32]*bucket{}
	for _, rec := range trace.Records {
		if rec.Offset%4 != 0 {
			continue
		}
		b := byOffset[rec.Offset]
		if b == nil {
			b = &bucket{}
			byOffset[rec.Offset] = b
		}
		if rec.Type == AccessRead {
			b.reads = append(b.reads, rec.Value)
			if b.readValues == nil {
				b.readValues = map[uint32]struct{}{}
			}
			b.readValues[rec.Value] = struct{}{}
		} else {
			b.writes = append(b.writes, rec.Value)
		}
	}

	for off, b := range byOffset {
		if len(b.reads) >= 2 {
			allSame := true
			for i := 1; i < len(b.reads); i++ {
				if b.reads[i] != b.reads[0] {
					allSame = false
					break
				}
			}
			if allSame {
				out.Static[off] = b.reads[0]
			} else if len(b.reads) <= 16 {
				out.Sequential[off] = append([]uint32(nil), b.reads...)
			}
		}

		writeMask := inferWriteMask(b.writes)
		if writeMask != 0 {
			out.WriteMask[off] = writeMask
		}
		rw1cMask := inferRW1CMask(b.writes, b.readValues, writeMask)
		if rw1cMask != 0 {
			out.RW1CMask[off] = rw1cMask
		}
	}

	return out
}

func inferWriteMask(values []uint32) uint32 {
	if len(values) < 2 {
		return 0
	}
	var mask uint32
	prev := values[0]
	for _, v := range values[1:] {
		mask |= prev ^ v
		prev = v
	}
	return mask
}

func inferRW1CMask(writes []uint32, observedReads map[uint32]struct{}, writeMask uint32) uint32 {
	if len(writes) < 1 || writeMask == 0 {
		return 0
	}

	var mask uint32
	for bit := uint(0); bit < 32; bit++ {
		bitMask := uint32(1) << bit
		if writeMask&bitMask == 0 {
			continue
		}
		seenZero := false
		seenOne := false
		for _, write := range writes {
			seenZero = seenZero || (write&bitMask) == 0
			seenOne = seenOne || (write&bitMask) != 0
		}
		if !seenZero || seenOne {
			continue
		}
		readSet := false
		for readVal := range observedReads {
			if readVal&bitMask != 0 {
				readSet = true
				break
			}
		}
		if readSet {
			mask |= bitMask
		}
	}
	return mask
}

// RemapTraceToBAROffsets converts physical mmiotrace addresses to offsets
// inside the selected BAR. Records outside the BAR are ignored.
func RemapTraceToBAROffsets(trace *TraceResult, base, size uint64) (*TraceResult, error) {
	if trace == nil || size == 0 {
		return trace, nil
	}
	out := *trace
	out.BARSize = int(size)
	out.Records = make([]AccessRecord, 0, len(trace.Records))
	for _, rec := range trace.Records {
		mapped := rec
		if rec.Address != 0 {
			if rec.Address < base || rec.Address+4 > base+size {
				continue
			}
			mapped.Offset = uint32(rec.Address - base)
		} else if uint64(rec.Offset)+4 > size {
			return nil, fmt.Errorf("trace offset 0x%x outside BAR size 0x%x", rec.Offset, size)
		}
		out.Records = append(out.Records, mapped)
	}
	return &out, nil
}
