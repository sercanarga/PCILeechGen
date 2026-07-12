package output

import (
	"encoding/binary"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/firmware/barmodel"
	"github.com/sercanarga/pcileechgen/internal/firmware/devicemodel"
)

func TestApplyGeneratedBARModelsCopiesFirmwareResetState(t *testing.T) {
	model := &devicemodel.Model{
		BARs:      []devicemodel.BAR{{BIR: 0, Type: devicemodel.BARTypeMem32, Size: 0x100}},
		Registers: []devicemodel.Register{{Name: "config", Space: devicemodel.SpaceConfig, BIR: -1}},
	}
	generated := []*barmodel.BARModel{{
		BIR:  0,
		Size: 0x100,
		Registers: []barmodel.BARRegister{{
			Offset:  0x08,
			Width:   4,
			Reset:   0x11223344,
			RWMask:  0x0000ffff,
			W1CMask: 0x00ff0000,
			Name:    "STATUS",
		}},
	}}

	applyGeneratedBARModels(model, generated)

	if got := binary.LittleEndian.Uint32(model.BARs[0].ResetImage[8:12]); got != 0x11223344 {
		t.Fatalf("BAR reset value = %#x", got)
	}
	if len(model.Registers) != 2 {
		t.Fatalf("register count = %d, want config plus generated BAR register", len(model.Registers))
	}
	register := model.Registers[1]
	if register.Space != devicemodel.SpaceBAR || register.BIR != 0 || register.Name != "STATUS" {
		t.Fatalf("generated register = %+v", register)
	}
	if len(register.Fields) != 3 {
		t.Fatalf("field count = %d, want RW, RW1C, and RO", len(register.Fields))
	}
}
