package nvme

import "github.com/sercanarga/pcileechgen/internal/donor/mmio"

type TraceEvidence struct {
	HasAdminQueues    bool   `json:"has_admin_queues"`
	AQA               uint32 `json:"aqa"`
	ASQ               uint64 `json:"asq"`
	ACQ               uint64 `json:"acq"`
	HasDoorbellStride bool   `json:"has_doorbell_stride"`
	DoorbellStride    uint32 `json:"doorbell_stride"`
	SQ0DoorbellOffset uint32 `json:"sq0_doorbell_offset"`
	CQ0DoorbellOffset uint32 `json:"cq0_doorbell_offset"`
}

func BuildTraceEvidence(trace *mmio.TraceResult) *TraceEvidence {
	if trace == nil || len(trace.Records) == 0 {
		return nil
	}

	var ev TraceEvidence
	var hasAQA, hasASQLo, hasASQHi, hasACQLo, hasACQHi bool
	var asqLo, asqHi, acqLo, acqHi uint32
	var sawSQ0 bool
	var cq0Offset uint32

	for _, rec := range trace.Records {
		if rec.Type != mmio.AccessWrite {
			continue
		}
		switch rec.Offset {
		case 0x24:
			ev.AQA, hasAQA = rec.Value, true
		case 0x28:
			asqLo, hasASQLo = rec.Value, true
		case 0x2C:
			asqHi, hasASQHi = rec.Value, true
		case 0x30:
			acqLo, hasACQLo = rec.Value, true
		case 0x34:
			acqHi, hasACQHi = rec.Value, true
		case 0x1000:
			sawSQ0 = true
		default:
			if rec.Offset > 0x1000 && (cq0Offset == 0 || rec.Offset < cq0Offset) {
				cq0Offset = rec.Offset
			}
		}
	}

	if hasAQA && hasASQLo && hasASQHi && hasACQLo && hasACQHi {
		ev.HasAdminQueues = true
		ev.ASQ = uint64(asqHi)<<32 | uint64(asqLo)
		ev.ACQ = uint64(acqHi)<<32 | uint64(acqLo)
	}
	if sawSQ0 && cq0Offset > 0x1000 {
		if stride, ok := doorbellStrideFromDelta(cq0Offset - 0x1000); ok {
			ev.HasDoorbellStride = true
			ev.DoorbellStride = stride
			ev.SQ0DoorbellOffset = 0x1000
			ev.CQ0DoorbellOffset = cq0Offset
		}
	}
	if ev.HasAdminQueues || ev.HasDoorbellStride {
		return &ev
	}
	return nil
}

func doorbellStrideFromDelta(delta uint32) (uint32, bool) {
	if delta < 4 || delta&(delta-1) != 0 {
		return 0, false
	}
	for stride := uint32(0); stride < 16; stride++ {
		if uint32(4)<<stride == delta {
			return stride, true
		}
	}
	return 0, false
}
