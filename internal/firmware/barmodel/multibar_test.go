package barmodel

import (
	"strings"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

func TestBuildBARModelsPreservesDonorBIRsAndBuildsEveryMemoryBAR(t *testing.T) {
	bars := []pci.BAR{
		{Index: 0, Size: 0x1000, Type: pci.BARTypeMem32},
		{Index: 2, Size: 0x2000, Type: pci.BARTypeMem32, Prefetchable: true},
		{Index: 5, Size: 0x4000, Type: pci.BARTypeMem32},
	}
	contents := map[int][]byte{
		0: make([]byte, 0x1000),
		2: make([]byte, 0x2000),
		5: make([]byte, 0x4000),
	}

	models := mustBuildBARModels(t, bars, contents, nil, 0xFF0000, 2)
	if len(models) != 3 {
		t.Fatalf("BuildBARModels returned %d models, want one for each of 3 donor memory BARs", len(models))
	}

	for i, want := range []struct {
		bir          int
		size         int
		prefetchable bool
	}{{0, 0x1000, false}, {2, 0x2000, true}, {5, 0x4000, false}} {
		got := models[i]
		if got.BIR != want.bir {
			t.Errorf("models[%d].BIR = %d, want donor BIR %d (models must be BIR ordered)", i, got.BIR, want.bir)
		}
		if got.Size != want.size {
			t.Errorf("BAR%d size = 0x%X, want donor aperture 0x%X", want.bir, got.Size, want.size)
		}
		if got.Type != pci.BARTypeMem32 {
			t.Errorf("BAR%d type = %q, want %q", want.bir, got.Type, pci.BARTypeMem32)
		}
		if got.Prefetchable != want.prefetchable {
			t.Errorf("BAR%d Prefetchable = %v, want %v", want.bir, got.Prefetchable, want.prefetchable)
		}
	}
}

func TestBuildBARModelsUsesPreferredClassBARNotLargestBAR(t *testing.T) {
	bars := []pci.BAR{
		{Index: 0, Size: 0x4000, Type: pci.BARTypeMem32},
		{Index: 2, Size: 0x1000, Type: pci.BARTypeMem32},
	}
	contents := map[int][]byte{
		0: make([]byte, 0x4000),
		2: make([]byte, 0x1000),
	}
	profiles := map[int]*donor.BARProfile{
		0: {
			BarIndex: 0,
			Size:     0x4000,
			Probes: []donor.BARProbeResult{
				{Offset: 0x100, Original: 0x12345678, RWMask: 0x0000FFFF},
			},
		},
	}

	models := mustBuildBARModels(t, bars, contents, profiles, 0x020000, 2)
	if len(models) != 2 {
		t.Fatalf("BuildBARModels returned %d models, want 2", len(models))
	}
	byBIR := modelsByBIR(models)
	if !byBIR[2].ClassSpecific {
		t.Fatal("preferred BAR2 should be marked class-specific")
	}
	if byBIR[0].ClassSpecific {
		t.Fatal("non-preferred BAR0 must not be marked class-specific")
	}
	if !modelHasRegister(byBIR[2], "MAC0_3") {
		t.Fatal("Ethernet class register model was not assigned to preferred BAR2")
	}
	if modelHasRegister(byBIR[0], "MAC0_3") {
		t.Fatal("Ethernet class register model incorrectly followed the largest BAR instead of preferred BAR2")
	}
	if !modelHasOffset(byBIR[0], 0x100) {
		t.Fatal("non-preferred BAR0 should retain its donor probe-derived register model")
	}
}

func TestBuildBARModelsAbsentPreferredBIRDoesNotMoveClassSemantics(t *testing.T) {
	bars := []pci.BAR{
		{Index: 0, Type: pci.BARTypeDisabled},
		{Index: 1, Size: 0x100, Type: pci.BARTypeIO},
		{Index: 3, Size: 0x2000, Type: pci.BARTypeMem32},
		{Index: 5, Size: 0x1000, Type: pci.BARTypeMem32},
	}
	contents := map[int][]byte{3: make([]byte, 0x2000), 5: make([]byte, 0x1000)}

	models := mustBuildBARModels(t, bars, contents, nil, 0x020000, 2)
	if len(models) != 2 {
		t.Fatalf("BuildBARModels returned %d models, want only BAR3 and BAR5", len(models))
	}
	if models[0].BIR != 3 || models[1].BIR != 5 {
		t.Fatalf("model BIRs = [%d %d], want deterministic [3 5]", models[0].BIR, models[1].BIR)
	}
	for _, model := range models {
		if model.ClassSpecific || modelHasRegister(model, "MAC0_3") {
			t.Fatalf("absent preferred BAR2 must not move Ethernet class semantics onto BAR%d", model.BIR)
		}
	}
}

func TestBuildBARModelsSkipsUpperHalfOf64BitBARPair(t *testing.T) {
	bars := []pci.BAR{
		{Index: 0, Size: 0x2000, Type: pci.BARTypeMem64, Is64Bit: true},
		{Index: 1, Type: pci.BARTypeDisabled},
		{Index: 2, Size: 0x1000, Type: pci.BARTypeMem32},
	}
	contents := map[int][]byte{
		0: make([]byte, 0x2000),
		2: make([]byte, 0x1000),
	}

	models := mustBuildBARModels(t, bars, contents, nil, 0xFF0000, 0)
	if len(models) != 2 {
		t.Fatalf("BuildBARModels returned %d models, want BAR0 and BAR2 only", len(models))
	}
	if models[0].BIR != 0 || !models[0].Is64Bit {
		t.Fatalf("first model = %+v, want 64-bit BAR0", models[0])
	}
	if models[1].BIR != 2 {
		t.Fatalf("second model BIR = %d, want 2; BIR1 is consumed by 64-bit BAR0", models[1].BIR)
	}
}

func TestBuildBARModelsNoPresentMemoryBARs(t *testing.T) {
	bars := []pci.BAR{
		{Index: 0, Type: pci.BARTypeDisabled},
		{Index: 2, Size: 0x100, Type: pci.BARTypeIO},
	}
	contents := map[int][]byte{0: make([]byte, 0x1000), 2: make([]byte, 0x100)}

	if models := mustBuildBARModels(t, bars, contents, nil, 0xFF0000, 0); len(models) != 0 {
		t.Fatalf("BuildBARModels returned %+v, want no endpoints for disabled/I/O BARs", models)
	}
	if models := mustBuildBARModels(t, nil, nil, nil, 0xFF0000, 0); len(models) != 0 {
		t.Fatalf("BuildBARModels with absent donor BARs returned %+v, want none", models)
	}
}

func TestBuildBARModelsRejects64BitBARAtBIR5(t *testing.T) {
	_, err := BuildBARModels(
		[]pci.BAR{{Index: 5, Size: 0x1000, Type: pci.BARTypeMem64, Is64Bit: true}},
		map[int][]byte{5: make([]byte, 0x1000)}, nil, 0xFF0000, 0,
	)
	if err == nil || !strings.Contains(err.Error(), "64-bit BAR5 has no upper BIR") {
		t.Fatalf("BuildBARModels error = %v, want invalid 64-bit BAR5 rejection", err)
	}
}

func TestBuildBARModelsRejectsOccupied64BitUpperBIR(t *testing.T) {
	_, err := BuildBARModels(
		[]pci.BAR{
			{Index: 0, Size: 0x2000, Type: pci.BARTypeMem64, Is64Bit: true},
			{Index: 1, Size: 0x1000, Type: pci.BARTypeMem32},
		},
		map[int][]byte{0: make([]byte, 0x2000), 1: make([]byte, 0x1000)}, nil, 0xFF0000, 0,
	)
	if err == nil || !strings.Contains(err.Error(), "64-bit BAR0 consumes occupied BAR1") {
		t.Fatalf("BuildBARModels error = %v, want occupied upper BIR rejection", err)
	}
}

func mustBuildBARModels(t *testing.T, bars []pci.BAR, contents map[int][]byte, profiles map[int]*donor.BARProfile, classCode uint32, preferredBIR int) []*BARModel {
	t.Helper()
	models, err := BuildBARModels(bars, contents, profiles, classCode, preferredBIR)
	if err != nil {
		t.Fatal(err)
	}
	return models
}

func modelsByBIR(models []*BARModel) map[int]*BARModel {
	out := make(map[int]*BARModel, len(models))
	for _, model := range models {
		out[model.BIR] = model
	}
	return out
}

func modelHasRegister(model *BARModel, name string) bool {
	if model == nil {
		return false
	}
	for _, reg := range model.Registers {
		if reg.Name == name {
			return true
		}
	}
	return false
}

func modelHasOffset(model *BARModel, offset uint32) bool {
	if model == nil {
		return false
	}
	for _, reg := range model.Registers {
		if reg.Offset == offset {
			return true
		}
	}
	return false
}
