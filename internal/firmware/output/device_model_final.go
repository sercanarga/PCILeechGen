package output

import (
	"encoding/binary"

	"github.com/sercanarga/pcileechgen/internal/firmware/barmodel"
	"github.com/sercanarga/pcileechgen/internal/firmware/devicemodel"
)

func applyGeneratedBARModels(model *devicemodel.Model, generated []*barmodel.BARModel) {
	registers := model.Registers[:0]
	for _, register := range model.Registers {
		if register.Space == devicemodel.SpaceConfig {
			registers = append(registers, register)
		}
	}
	model.Registers = registers

	for _, source := range generated {
		if source == nil || source.BIR < 0 || source.BIR > 5 || source.Size <= 0 {
			continue
		}
		bar := deviceModelBAR(model, source)
		bar.ResetImage = make([]byte, source.Size)
		for _, sourceRegister := range source.Registers {
			width := sourceRegister.Width
			if width != 1 && width != 2 && width != 4 || int(sourceRegister.Offset)+width > len(bar.ResetImage) {
				continue
			}
			writeResetValue(bar.ResetImage[sourceRegister.Offset:], width, sourceRegister.Reset)
			model.Registers = append(model.Registers, generatedRegister(source.BIR, sourceRegister))
		}
	}
}

func deviceModelBAR(model *devicemodel.Model, source *barmodel.BARModel) *devicemodel.BAR {
	for index := range model.BARs {
		if model.BARs[index].BIR == source.BIR {
			model.BARs[index].Size = uint64(source.Size)
			model.BARs[index].SizeKnown = true
			return &model.BARs[index]
		}
	}
	barType := devicemodel.BARTypeMem32
	addressWidth := uint8(32)
	if source.Is64Bit {
		barType = devicemodel.BARTypeMem64
		addressWidth = 64
	}
	model.BARs = append(model.BARs, devicemodel.BAR{
		BIR: source.BIR, Type: barType, Size: uint64(source.Size), SizeKnown: true,
		Prefetchable: source.Prefetchable, AddressWidth: addressWidth,
	})
	return &model.BARs[len(model.BARs)-1]
}

func writeResetValue(target []byte, width int, value uint32) {
	switch width {
	case 1:
		target[0] = byte(value)
	case 2:
		binary.LittleEndian.PutUint16(target, uint16(value))
	case 4:
		binary.LittleEndian.PutUint32(target, value)
	}
}

func generatedRegister(bir int, source barmodel.BARRegister) devicemodel.Register {
	widthMask := uint64(1)<<(source.Width*8) - 1
	reset := uint64(source.Reset) & widthMask
	rw := uint64(source.RWMask) & widthMask &^ uint64(source.W1CMask)
	rw1c := uint64(source.W1CMask) & widthMask
	ro := widthMask &^ (rw | rw1c)
	fields := make([]devicemodel.RegisterField, 0, 3)
	if rw != 0 {
		fields = append(fields, devicemodel.RegisterField{Name: "rw", Mask: rw, Access: devicemodel.AccessRW, ResetValue: reset & rw})
	}
	if rw1c != 0 {
		fields = append(fields, devicemodel.RegisterField{Name: "rw1c", Mask: rw1c, Access: devicemodel.AccessRW1C, ResetValue: reset & rw1c})
	}
	if ro != 0 {
		fields = append(fields, devicemodel.RegisterField{Name: "ro", Mask: ro, Access: devicemodel.AccessRO, ResetValue: reset & ro})
	}
	return devicemodel.Register{
		Name: source.Name, Space: devicemodel.SpaceBAR, BIR: bir,
		Offset: uint64(source.Offset), Width: uint8(source.Width),
		ResetDomain: devicemodel.ResetPowerOn, ResetValue: reset,
		Fields: fields, Confidence: devicemodel.ConfidenceSpecified,
	}
}
