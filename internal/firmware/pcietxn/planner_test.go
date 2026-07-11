package pcietxn_test

import (
	"fmt"
	"math/bits"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/firmware/pcietxn"
)

func limits(rcb, mps int) pcietxn.Limits {
	return pcietxn.Limits{
		ReadCompletionBoundary: uint16(rcb),
		MaxPayloadBytes:        uint16(mps),
	}
}

func fullRead(format pcietxn.HeaderFormat, address uint64, lengthDW int) pcietxn.Request {
	lastBE := uint8(0xF)
	if lengthDW == 1 {
		lastBE = 0
	}
	return pcietxn.Request{
		Kind:         pcietxn.MemoryRead,
		Format:       format,
		Address:      address,
		LengthDW:     uint16(lengthDW),
		FirstBE:      0xF,
		LastBE:       lastBE,
		RequesterID:  0xA15E,
		Tag:          0xD3,
		TrafficClass: 5,
		Attributes:   3,
		BIR:          0,
	}
}

func mustPlan(t *testing.T, req pcietxn.Request, lim pcietxn.Limits) pcietxn.Plan {
	t.Helper()
	plan, err := pcietxn.PlanRequest(req, lim)
	if err != nil {
		t.Fatalf("PlanRequest() error = %v", err)
	}
	return plan
}

func completionLengths(plan pcietxn.Plan) []int {
	lengths := make([]int, len(plan.Completions))
	for i, completion := range plan.Completions {
		lengths[i] = int(completion.LengthDW)
	}
	return lengths
}

func requireLengths(t *testing.T, plan pcietxn.Plan, want ...int) {
	t.Helper()
	got := completionLengths(plan)
	if fmt.Sprint(got) != fmt.Sprint(want) {
		t.Fatalf("completion lengths = %v, want %v", got, want)
	}
}

func TestPlanMemoryReadHeadersAndBoundaryLengths(t *testing.T) {
	for _, format := range []pcietxn.HeaderFormat{pcietxn.Header3DW, pcietxn.Header4DW} {
		for _, tc := range []struct {
			lengthDW int
			want     []int
		}{
			{lengthDW: 1, want: []int{1}},
			{lengthDW: 2, want: []int{2}},
			{lengthDW: 31, want: []int{31}},
			{lengthDW: 32, want: []int{32}},
			{lengthDW: 33, want: []int{32, 1}},
		} {
			t.Run(fmt.Sprintf("format_%d/length_%d", format, tc.lengthDW), func(t *testing.T) {
				address := uint64(0x1000)
				if format == pcietxn.Header4DW {
					address = 0x1_0000_1000
				}
				plan := mustPlan(t, fullRead(format, address, tc.lengthDW), limits(128, 128))
				if plan.Decision != pcietxn.DecisionComplete || plan.Posted {
					t.Fatalf("decision/posted = %v/%t, want Complete/false", plan.Decision, plan.Posted)
				}
				requireLengths(t, plan, tc.want...)
				for i, completion := range plan.Completions {
					if completion.Status != pcietxn.CompletionSuccessful || !completion.HasData {
						t.Errorf("completion[%d] status/data = %v/%t, want Successful/true", i, completion.Status, completion.HasData)
					}
				}
			})
		}
	}
}

func TestPlanSingleDWAllMeaningfulFirstByteEnables(t *testing.T) {
	for firstBE := uint8(1); firstBE <= 0xF; firstBE++ {
		t.Run(fmt.Sprintf("first_be_%X", firstBE), func(t *testing.T) {
			req := fullRead(pcietxn.Header3DW, 0x2240, 1)
			req.FirstBE = firstBE
			plan := mustPlan(t, req, limits(128, 128))
			requireLengths(t, plan, 1)
			completion := plan.Completions[0]
			if got, want := int(completion.ByteCount), bits.OnesCount8(firstBE); got != want {
				t.Errorf("ByteCount = %d, want popcount(%04b) = %d", got, firstBE, want)
			}
			if got, want := int(completion.LowerAddress), 0x40+bits.TrailingZeros8(firstBE); got != want {
				t.Errorf("LowerAddress = %#x, want %#x", got, want)
			}
		})
	}
}

func TestPlanZeroLengthMemoryReadReturnsRequiredOneDWCplD(t *testing.T) {
	req := fullRead(pcietxn.Header3DW, 0x22C0, 1)
	req.FirstBE = 0
	plan := mustPlan(t, req, limits(128, 128))
	if plan.Decision != pcietxn.DecisionComplete || plan.Posted {
		t.Fatalf("decision/posted = %v/%t, want Complete/false", plan.Decision, plan.Posted)
	}
	if plan.EnabledByteCount != 0 {
		t.Errorf("zero-length read EnabledByteCount = %d, want 0", plan.EnabledByteCount)
	}
	requireLengths(t, plan, 1)
	completion := plan.Completions[0]
	if got := int(completion.ByteCount); got != 4 {
		t.Errorf("zero-length read Completion ByteCount = %d, want 4", got)
	}
	if got := int(completion.LowerAddress); got != 0x40 {
		t.Errorf("zero-length read Completion LowerAddress = %#x, want request address %#x", got, 0x40)
	}
}

func TestPlanQWAlignedTwoDWAllNonzeroByteEnablePatterns(t *testing.T) {
	for firstBE := uint8(1); firstBE <= 0xF; firstBE++ {
		for lastBE := uint8(1); lastBE <= 0xF; lastBE++ {
			t.Run(fmt.Sprintf("first_%X/last_%X", firstBE, lastBE), func(t *testing.T) {
				req := fullRead(pcietxn.Header3DW, 0x2800, 2)
				req.FirstBE = firstBE
				req.LastBE = lastBE
				plan := mustPlan(t, req, limits(128, 128))
				requireLengths(t, plan, 2)
				completion := plan.Completions[0]
				wantBytes := bits.OnesCount8(firstBE) + bits.OnesCount8(lastBE)
				if got := int(completion.ByteCount); got != wantBytes {
					t.Errorf("ByteCount = %d, want popcount(FirstBE)+popcount(LastBE) = %d", got, wantBytes)
				}
				if got, want := int(completion.LowerAddress), bits.TrailingZeros8(firstBE); got != want {
					t.Errorf("LowerAddress = %#x, want %#x", got, want)
				}
			})
		}
	}
}

func TestPlanMultiDWAllMeaningfulEdgeByteEnables(t *testing.T) {
	firstPatterns := []uint8{0x8, 0xC, 0xE, 0xF}
	lastPatterns := []uint8{0x1, 0x3, 0x7, 0xF}
	for _, lengthDW := range []int{31, 32, 33} {
		for _, firstBE := range firstPatterns {
			for _, lastBE := range lastPatterns {
				t.Run(fmt.Sprintf("length_%d/first_%X/last_%X", lengthDW, firstBE, lastBE), func(t *testing.T) {
					req := fullRead(pcietxn.Header3DW, 0x3000, lengthDW)
					req.FirstBE = firstBE
					req.LastBE = lastBE
					plan := mustPlan(t, req, limits(128, 128))
					first := plan.Completions[0]
					wantBytes := bits.OnesCount8(firstBE) + 4*(lengthDW-2) + bits.OnesCount8(lastBE)
					if got := int(first.ByteCount); got != wantBytes {
						t.Errorf("first completion ByteCount = %d, want %d", got, wantBytes)
					}
					if got, want := int(first.LowerAddress), bits.TrailingZeros8(firstBE); got != want {
						t.Errorf("first completion LowerAddress = %#x, want %#x", got, want)
					}
					if lengthDW == 33 {
						last := plan.Completions[1]
						if got, want := int(last.ByteCount), bits.OnesCount8(lastBE); got != want {
							t.Errorf("last completion ByteCount = %d, want remaining enabled bytes %d", got, want)
						}
					}
				})
			}
		}
	}
}

func TestPlanReadCompletionBoundary64And128(t *testing.T) {
	for _, tc := range []struct {
		name     string
		rcb      int
		address  uint64
		wantLens []int
		wantBC   []int
		wantLA   []int
	}{
		{name: "64_byte_RCB", rcb: 64, address: 0x203C, wantLens: []int{1, 16, 16}, wantBC: []int{132, 128, 64}, wantLA: []int{0x3C, 0x40, 0}},
		{name: "128_byte_RCB", rcb: 128, address: 0x207C, wantLens: []int{1, 32}, wantBC: []int{132, 128}, wantLA: []int{0x7C, 0}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			plan := mustPlan(t, fullRead(pcietxn.Header3DW, tc.address, 33), limits(tc.rcb, 256))
			requireLengths(t, plan, tc.wantLens...)
			for i, completion := range plan.Completions {
				if got := int(completion.ByteCount); got != tc.wantBC[i] {
					t.Errorf("completion[%d].ByteCount = %d, want %d", i, got, tc.wantBC[i])
				}
				if got := int(completion.LowerAddress); got != tc.wantLA[i] {
					t.Errorf("completion[%d].LowerAddress = %#x, want %#x", i, got, tc.wantLA[i])
				}
			}
		})
	}
}

func TestPlanReadMPSBoundaryAndPartialEdges(t *testing.T) {
	req := fullRead(pcietxn.Header4DW, 0x1_0000_3000, 33)
	req.FirstBE = 0xC
	req.LastBE = 0x3
	plan := mustPlan(t, req, limits(128, 128))
	requireLengths(t, plan, 32, 1)

	if got, want := int(plan.Completions[0].ByteCount), 128; got != want {
		t.Errorf("first ByteCount = %d, want %d", got, want)
	}
	if got, want := int(plan.Completions[0].LowerAddress), 2; got != want {
		t.Errorf("first LowerAddress = %#x, want %#x", got, want)
	}
	if got, want := int(plan.Completions[1].ByteCount), 2; got != want {
		t.Errorf("second ByteCount = %d, want partial last-DW bytes %d", got, want)
	}
	if got := int(plan.Completions[1].LowerAddress); got != 0 {
		t.Errorf("second LowerAddress = %#x, want 0", got)
	}
}

func TestPlanPreservesRequesterTagTrafficClassAndAttributes(t *testing.T) {
	req := fullRead(pcietxn.Header4DW, 0x1_0000_407C, 33)
	plan := mustPlan(t, req, limits(128, 128))
	if len(plan.Completions) < 2 {
		t.Fatalf("got %d completions, want a split request", len(plan.Completions))
	}
	for i, completion := range plan.Completions {
		if completion.RequesterID != req.RequesterID || completion.Tag != req.Tag ||
			completion.TrafficClass != req.TrafficClass || completion.Attributes != req.Attributes {
			t.Errorf("completion[%d] metadata = requester %#x tag %#x TC %#x attr %#x; want %#x/%#x/%#x/%#x",
				i, completion.RequesterID, completion.Tag, completion.TrafficClass, completion.Attributes,
				req.RequesterID, req.Tag, req.TrafficClass, req.Attributes)
		}
	}
}

func TestPlanPostedMemoryWrites(t *testing.T) {
	for _, format := range []pcietxn.HeaderFormat{pcietxn.Header3DW, pcietxn.Header4DW} {
		for _, lengthDW := range []int{1, 2, 31, 32, 33} {
			t.Run(fmt.Sprintf("format_%d/length_%d", format, lengthDW), func(t *testing.T) {
				req := fullRead(format, 0x5000, lengthDW)
				req.Kind = pcietxn.MemoryWrite
				plan := mustPlan(t, req, limits(128, 128))
				if plan.Decision != pcietxn.DecisionPosted || !plan.Posted {
					t.Fatalf("decision/posted = %v/%t, want Posted/true", plan.Decision, plan.Posted)
				}
				if len(plan.Completions) != 0 {
					t.Fatalf("posted write produced %d completions, want none", len(plan.Completions))
				}
			})
		}
	}
}

func TestPlanZeroByteMemoryWriteRemainsPosted(t *testing.T) {
	req := fullRead(pcietxn.Header3DW, 0x5800, 1)
	req.Kind = pcietxn.MemoryWrite
	req.FirstBE = 0
	plan := mustPlan(t, req, limits(128, 128))
	if plan.Decision != pcietxn.DecisionPosted || !plan.Posted {
		t.Fatalf("decision/posted = %v/%t, want Posted/true", plan.Decision, plan.Posted)
	}
	if plan.EnabledByteCount != 0 {
		t.Errorf("zero-byte write EnabledByteCount = %d, want 0", plan.EnabledByteCount)
	}
	if len(plan.Completions) != 0 {
		t.Fatalf("zero-byte posted write produced %d completions, want none", len(plan.Completions))
	}
}

func TestPlanUnsupportedTransactions(t *testing.T) {
	for _, kind := range []pcietxn.Kind{pcietxn.IORead, pcietxn.IOWrite, pcietxn.Atomic} {
		t.Run(fmt.Sprintf("kind_%d", kind), func(t *testing.T) {
			req := fullRead(pcietxn.Header3DW, 0x6000, 1)
			req.Kind = kind
			plan := mustPlan(t, req, limits(128, 128))
			if plan.Decision != pcietxn.DecisionUnsupportedRequest || plan.Posted {
				t.Fatalf("decision/posted = %v/%t, want UnsupportedRequest/false", plan.Decision, plan.Posted)
			}
			if len(plan.Completions) != 1 {
				t.Fatalf("unsupported request produced %d completions, want one UR", len(plan.Completions))
			}
			completion := plan.Completions[0]
			if completion.Status != pcietxn.CompletionUnsupportedRequest || completion.HasData {
				t.Errorf("UR status/data = %v/%t, want UnsupportedRequest/false", completion.Status, completion.HasData)
			}
			if completion.RequesterID != req.RequesterID || completion.Tag != req.Tag ||
				completion.TrafficClass != req.TrafficClass || completion.Attributes != req.Attributes {
				t.Errorf("UR metadata = requester %#x tag %#x TC %#x attr %#x; want %#x/%#x/%#x/%#x",
					completion.RequesterID, completion.Tag, completion.TrafficClass, completion.Attributes,
					req.RequesterID, req.Tag, req.TrafficClass, req.Attributes)
			}
			if completion.LengthDW != 0 || completion.ByteCount != 0 || completion.LowerAddress != 0 {
				t.Errorf("UR data fields = length %d byte count %d lower %#x, want zero", completion.LengthDW, completion.ByteCount, completion.LowerAddress)
			}
		})
	}
}

func TestPlanLockedRead3DWAnd4DWEmitsUR(t *testing.T) {
	for _, format := range []pcietxn.HeaderFormat{pcietxn.Header3DW, pcietxn.Header4DW} {
		t.Run(fmt.Sprintf("format_%d", format), func(t *testing.T) {
			req := fullRead(format, 0x6800, 1)
			if format == pcietxn.Header4DW {
				req.Address = 0x1_0000_6800
			}
			req.Kind = pcietxn.MemoryReadLocked
			plan := mustPlan(t, req, limits(64, 128))
			if plan.Decision != pcietxn.DecisionUnsupportedRequest || plan.Posted {
				t.Fatalf("decision/posted = %v/%t, want UnsupportedRequest/false", plan.Decision, plan.Posted)
			}
			if len(plan.Completions) != 1 {
				t.Fatalf("locked read produced %d completions, want one UR", len(plan.Completions))
			}
			completion := plan.Completions[0]
			if completion.Status != pcietxn.CompletionUnsupportedRequest || completion.HasData {
				t.Errorf("locked-read UR status/data = %v/%t, want UnsupportedRequest/false", completion.Status, completion.HasData)
			}
			if completion.RequesterID != req.RequesterID || completion.Tag != req.Tag ||
				completion.TrafficClass != req.TrafficClass {
				t.Errorf("locked-read UR metadata = %#x/%#x/%#x, want %#x/%#x/%#x",
					completion.RequesterID, completion.Tag, completion.TrafficClass,
					req.RequesterID, req.Tag, req.TrafficClass)
			}
		})
	}
}

func TestPlanMalformedRequests(t *testing.T) {
	valid := fullRead(pcietxn.Header3DW, 0x7000, 2)
	cases := map[string]func(*pcietxn.Request){
		"decoder_marked_malformed": func(req *pcietxn.Request) { req.Malformed = true },
		"zero_length":              func(req *pcietxn.Request) { req.LengthDW = 0 },
		"length_over_1024":         func(req *pcietxn.Request) { req.LengthDW = 1025 },
		"unknown_header_format":    func(req *pcietxn.Request) { req.Format = pcietxn.HeaderUnknown },
		"3DW_address_above_4GiB":   func(req *pcietxn.Request) { req.Address = 0x1_0000_7000 },
		"BE_outside_DWORD":         func(req *pcietxn.Request) { req.FirstBE = 0x18 },
		"traffic_class_overflow":   func(req *pcietxn.Request) { req.TrafficClass = 8 },
		"attribute_overflow":       func(req *pcietxn.Request) { req.Attributes = 8 },
		"one_DW_last_BE":           func(req *pcietxn.Request) { req.LengthDW, req.LastBE = 1, 1 },
		"multi_DW_first_BE":        func(req *pcietxn.Request) { req.LengthDW, req.FirstBE = 3, 3 },
		"multi_DW_last_BE":         func(req *pcietxn.Request) { req.LengthDW, req.LastBE = 3, 0xC },
		"non_QW_2DW_discontinuous": func(req *pcietxn.Request) { req.Address, req.FirstBE = 0x7004, 3 },
		"unaligned_address":        func(req *pcietxn.Request) { req.Address++ },
		"crosses_4KiB":             func(req *pcietxn.Request) { req.Address, req.LengthDW = 0xFF0, 8 },
		"reserved_BIR":             func(req *pcietxn.Request) { req.BIR = 7 },
	}
	for name, mutate := range cases {
		t.Run(name, func(t *testing.T) {
			req := valid
			mutate(&req)
			plan := mustPlan(t, req, limits(128, 128))
			if plan.Decision != pcietxn.DecisionMalformedRequest || plan.Posted {
				t.Fatalf("decision/posted = %v/%t, want MalformedRequest/false", plan.Decision, plan.Posted)
			}
			if len(plan.Completions) != 0 {
				t.Fatalf("malformed request produced %d completions, want none", len(plan.Completions))
			}
		})
	}
}
