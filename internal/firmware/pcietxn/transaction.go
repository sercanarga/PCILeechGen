package pcietxn

import (
	"errors"
	"fmt"
	"math/bits"
)

type Kind uint8

const (
	KindUnknown Kind = iota
	MemoryRead
	MemoryWrite
	IORead
	IOWrite
	Atomic
	MemoryReadLocked
	ConfigurationRead
	ConfigurationWrite
)

const (
	KindMemoryRead         = MemoryRead
	KindMemoryWrite        = MemoryWrite
	KindIORead             = IORead
	KindIOWrite            = IOWrite
	KindAtomic             = Atomic
	KindMemoryReadLocked   = MemoryReadLocked
	KindConfigurationRead  = ConfigurationRead
	KindConfigurationWrite = ConfigurationWrite
)

type HeaderFormat uint8

const (
	HeaderUnknown HeaderFormat = iota
	Header3DW
	Header4DW
)

const (
	Format3DW = Header3DW
	Format4DW = Header4DW
)

type Decision uint8

const (
	DecisionUnknown Decision = iota
	DecisionComplete
	DecisionPosted
	DecisionUnsupportedRequest
	DecisionMalformedRequest
)

const (
	DecisionUR          = DecisionUnsupportedRequest
	DecisionUnsupported = DecisionUnsupportedRequest
	DecisionMalformed   = DecisionMalformedRequest
)

type CompletionStatus uint8

const (
	CompletionSuccessful CompletionStatus = iota
	CompletionUnsupportedRequest
)

type Request struct {
	Kind         Kind
	Format       HeaderFormat
	Address      uint64
	LengthDW     uint16
	FirstBE      uint8
	LastBE       uint8
	RequesterID  uint16
	Tag          uint8
	TrafficClass uint8
	Attributes   uint8
	BIR          uint8
	Malformed    bool
}

type Limits struct {
	ReadCompletionBoundary uint16
	MaxPayloadBytes        uint16
}

type Completion struct {
	Status       CompletionStatus
	HasData      bool
	Address      uint64
	LengthDW     uint16
	ByteCount    uint16
	LowerAddress uint8
	RequesterID  uint16
	Tag          uint8
	TrafficClass uint8
	Attributes   uint8
}

type Plan struct {
	Decision         Decision
	Posted           bool
	EnabledByteCount uint16
	Completions      []Completion
	Reason           string
}

var ErrInvalidLimits = errors.New("invalid PCIe completion limits")

func PlanRequest(req Request, limits Limits) (Plan, error) {
	if err := validateLimits(limits); err != nil {
		return Plan{}, err
	}
	if reason := malformedReason(req); reason != "" {
		return Plan{Decision: DecisionMalformedRequest, Reason: reason}, nil
	}

	switch req.Kind {
	case MemoryWrite:
		return Plan{
			Decision:         DecisionPosted,
			Posted:           true,
			EnabledByteCount: enabledByteCount(req),
		}, nil
	case MemoryRead:
	case IORead, IOWrite, Atomic, MemoryReadLocked, ConfigurationRead, ConfigurationWrite:
		return unsupportedPlan(req, "non-posted request is unsupported"), nil
	default:
		return Plan{Decision: DecisionUnsupportedRequest, Reason: "request type is unsupported"}, nil
	}

	enabledBytes := enabledByteCount(req)
	remainingBytes := completionByteCount(req)
	plan := Plan{
		Decision:         DecisionComplete,
		EnabledByteCount: enabledBytes,
		Completions:      make([]Completion, 0, completionCapacity(req, limits)),
	}

	for dwIndex := uint16(0); dwIndex < req.LengthDW; {
		address := req.Address + uint64(dwIndex)*4
		rcbRemainingBytes := uint64(limits.ReadCompletionBoundary) - address%uint64(limits.ReadCompletionBoundary)
		rcbDW := uint16(rcbRemainingBytes / 4)
		mpsDW := limits.MaxPayloadBytes / 4
		chunkDW := min16(req.LengthDW-dwIndex, min16(rcbDW, mpsDW))
		if chunkDW == 0 {
			return Plan{}, fmt.Errorf("%w: zero-sized completion", ErrInvalidLimits)
		}

		startBE := uint8(0xF)
		if dwIndex == 0 {
			startBE = req.FirstBE
		} else if dwIndex == req.LengthDW-1 {
			startBE = req.LastBE
		}
		firstByteOffset := uint8(0)
		if startBE != 0 {
			firstByteOffset = uint8(bits.TrailingZeros8(startBE))
		}
		plan.Completions = append(plan.Completions, Completion{
			Status:       CompletionSuccessful,
			HasData:      true,
			Address:      address,
			LengthDW:     chunkDW,
			ByteCount:    remainingBytes,
			LowerAddress: uint8((address + uint64(firstByteOffset)) & 0x7F),
			RequesterID:  req.RequesterID,
			Tag:          req.Tag,
			TrafficClass: req.TrafficClass,
			Attributes:   req.Attributes,
		})

		remainingBytes -= enabledBytesInRange(req, dwIndex, chunkDW)
		dwIndex += chunkDW
	}

	return plan, nil
}

func PlanCompletions(req Request, limits Limits) (Plan, error) {
	return PlanRequest(req, limits)
}

func unsupportedPlan(req Request, reason string) Plan {
	return Plan{
		Decision: DecisionUnsupportedRequest,
		Reason:   reason,
		Completions: []Completion{{
			Status:       CompletionUnsupportedRequest,
			HasData:      false,
			RequesterID:  req.RequesterID,
			Tag:          req.Tag,
			TrafficClass: req.TrafficClass,
			Attributes:   req.Attributes,
		}},
	}
}

func validateLimits(limits Limits) error {
	if limits.ReadCompletionBoundary != 64 && limits.ReadCompletionBoundary != 128 {
		return fmt.Errorf("%w: RCB must be 64 or 128 bytes", ErrInvalidLimits)
	}
	mps := limits.MaxPayloadBytes
	if mps < 128 || mps > 4096 || mps&(mps-1) != 0 {
		return fmt.Errorf("%w: MPS must be a power of two from 128 through 4096 bytes", ErrInvalidLimits)
	}
	return nil
}

func malformedReason(req Request) string {
	if req.Malformed {
		return "request was marked malformed by the wire normalizer"
	}
	if req.Format != Header3DW && req.Format != Header4DW {
		return "header format is neither 3DW nor 4DW"
	}
	if req.Format == Header3DW && req.Address > 0xFFFFFFFF {
		return "3DW request carries an address above 4 GiB"
	}
	if req.Address&3 != 0 {
		return "request address is not DWORD aligned"
	}
	if req.LengthDW == 0 || req.LengthDW > 1024 {
		return "request length is outside 1..1024 DWORDs"
	}
	if (req.Address&0xFFF)+uint64(req.LengthDW)*4 > 4096 {
		return "request crosses a 4 KiB boundary"
	}
	if req.FirstBE&0xF != req.FirstBE || req.LastBE&0xF != req.LastBE {
		return "byte enable uses bits outside a DWORD"
	}
	if req.LengthDW == 1 {
		if req.LastBE != 0 {
			return "1DW request requires zero LastBE"
		}
	} else if req.LengthDW == 2 && req.Address&7 == 0 {
		if req.FirstBE == 0 || req.LastBE == 0 {
			return "QW-aligned 2DW request requires nonzero byte enables"
		}
	} else if !validFirstEdge(req.FirstBE) || !validLastEdge(req.LastBE) {
		return "multi-DW request byte enables do not describe contiguous edges"
	}
	if req.TrafficClass > 7 || req.Attributes > 7 {
		return "traffic class or attributes exceed their header width"
	}
	if req.BIR > 6 {
		return "BIR is outside BAR0..BAR6"
	}
	return ""
}

func validFirstEdge(be uint8) bool {
	return be == 0x8 || be == 0xC || be == 0xE || be == 0xF
}

func validLastEdge(be uint8) bool {
	return be == 0x1 || be == 0x3 || be == 0x7 || be == 0xF
}

func enabledByteCount(req Request) uint16 {
	if req.LengthDW == 1 {
		return uint16(bits.OnesCount8(req.FirstBE))
	}
	return uint16(bits.OnesCount8(req.FirstBE)+bits.OnesCount8(req.LastBE)) + (req.LengthDW-2)*4
}

func completionByteCount(req Request) uint16 {
	if req.Kind == MemoryRead && req.LengthDW == 1 && req.FirstBE == 0 {
		return 4
	}
	return enabledByteCount(req)
}

func enabledBytesInRange(req Request, first, count uint16) uint16 {
	var enabled uint16
	for i := first; i < first+count; i++ {
		switch {
		case i == 0:
			enabled += uint16(bits.OnesCount8(req.FirstBE))
		case i == req.LengthDW-1:
			enabled += uint16(bits.OnesCount8(req.LastBE))
		default:
			enabled += 4
		}
	}
	return enabled
}

func completionCapacity(req Request, limits Limits) int {
	minBoundary := min16(limits.ReadCompletionBoundary, limits.MaxPayloadBytes)
	return int((req.LengthDW*4 + minBoundary - 1) / minBoundary)
}

func min16(a, b uint16) uint16 {
	if a < b {
		return a
	}
	return b
}
