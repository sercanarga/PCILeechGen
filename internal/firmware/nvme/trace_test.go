package nvme

import (
	"testing"

	"github.com/sercanarga/pcileechgen/internal/donor/mmio"
)

func TestBuildTraceEvidence_InfersDoorbellStride(t *testing.T) {
	evidence := BuildTraceEvidence(&mmio.TraceResult{
		BARSize: 8192,
		Records: []mmio.AccessRecord{
			{Offset: 0x1000, Type: mmio.AccessWrite, Value: 1},
			{Offset: 0x1008, Type: mmio.AccessWrite, Value: 1},
		},
	})
	if evidence == nil || !evidence.HasDoorbellStride {
		t.Fatalf("expected doorbell stride evidence, got %#v", evidence)
	}
	if evidence.DoorbellStride != 1 || evidence.SQ0DoorbellOffset != 0x1000 || evidence.CQ0DoorbellOffset != 0x1008 {
		t.Fatalf("unexpected doorbell evidence: %#v", evidence)
	}
}

func TestBuildTraceEvidence_ExtractsAdminQueueWrites(t *testing.T) {
	evidence := BuildTraceEvidence(&mmio.TraceResult{
		BARSize: 8192,
		Records: []mmio.AccessRecord{
			{Offset: 0x24, Type: mmio.AccessWrite, Value: 0x001f001f},
			{Offset: 0x28, Type: mmio.AccessWrite, Value: 0x12345000},
			{Offset: 0x2c, Type: mmio.AccessWrite, Value: 0x00000001},
			{Offset: 0x30, Type: mmio.AccessWrite, Value: 0x45678000},
			{Offset: 0x34, Type: mmio.AccessWrite, Value: 0x00000002},
		},
	})
	if evidence == nil || !evidence.HasAdminQueues {
		t.Fatalf("expected admin queue evidence, got %#v", evidence)
	}
	if evidence.AQA != 0x001f001f || evidence.ASQ != 0x0000000112345000 || evidence.ACQ != 0x0000000245678000 {
		t.Fatalf("unexpected queue evidence: %#v", evidence)
	}
}
