package devicemodel

import (
	"fmt"
	"sync"
)

type Oracle struct {
	mu        sync.Mutex
	registers []oracleRegister
}

type oracleRegister struct {
	definition Register
	value      uint64
}

func NewOracle(model *Model) (*Oracle, error) {
	if err := model.Validate(); err != nil {
		return nil, fmt.Errorf("create oracle: %w", err)
	}
	o := &Oracle{registers: make([]oracleRegister, len(model.Registers))}
	for i, reg := range model.Registers {
		reg.Fields = append([]RegisterField(nil), reg.Fields...)
		o.registers[i] = oracleRegister{definition: reg, value: reg.ResetValue & widthMask(reg.Width)}
	}
	return o, nil
}

func (o *Oracle) Read(bir int, offset uint64, width int) (uint64, error) {
	if o == nil {
		return 0, fmt.Errorf("read: nil oracle")
	}
	if !validTransactionWidth(width) {
		return 0, fmt.Errorf("read: width %d is not 1, 2, 4, or 8 bytes", width)
	}
	o.mu.Lock()
	defer o.mu.Unlock()
	reg, shift, err := o.findRegister(bir, offset, width)
	if err != nil {
		return 0, fmt.Errorf("read: %w", err)
	}
	transactionMask := maskForBytes(width) << shift
	visible := reg.value
	for _, field := range reg.definition.Fields {
		if field.Access == AccessReserved {
			visible &^= field.Mask
		}
	}
	result := (visible & transactionMask) >> shift
	for _, field := range reg.definition.Fields {
		if field.Access == AccessRC && field.Mask&transactionMask != 0 {
			reg.value &^= field.Mask & transactionMask
		}
	}
	return result, nil
}

func (o *Oracle) Write(bir int, offset uint64, width int, value uint64, byteEnable uint8) error {
	if o == nil {
		return fmt.Errorf("write: nil oracle")
	}
	if !validTransactionWidth(width) {
		return fmt.Errorf("write: width %d is not 1, 2, 4, or 8 bytes", width)
	}
	if width < 8 && byteEnable>>width != 0 {
		return fmt.Errorf("write: byte enable 0x%x selects lanes outside width %d", byteEnable, width)
	}
	o.mu.Lock()
	defer o.mu.Unlock()
	reg, shift, err := o.findRegister(bir, offset, width)
	if err != nil {
		return fmt.Errorf("write: %w", err)
	}
	laneMask := uint64(0)
	for lane := range width {
		if byteEnable&(1<<lane) != 0 {
			laneMask |= uint64(0xff) << (lane * 8)
		}
	}
	laneMask <<= shift
	writeValue := (value & maskForBytes(width)) << shift
	for _, field := range reg.definition.Fields {
		affected := field.Mask & laneMask
		if affected == 0 {
			continue
		}
		switch field.Access {
		case AccessRW:
			reg.value = (reg.value &^ affected) | (writeValue & affected)
		case AccessRW1C:
			reg.value &^= writeValue & affected
		case AccessW1S:
			reg.value |= writeValue & affected
		case AccessW0C:
			reg.value &^= (^writeValue) & affected
		case AccessRO, AccessRC, AccessReserved:
		}
	}
	reg.value &= widthMask(reg.definition.Width)
	return nil
}

func (o *Oracle) Reset(domain ResetDomain) error {
	if o == nil {
		return fmt.Errorf("reset: nil oracle")
	}
	if !validResetDomain(domain) {
		return fmt.Errorf("reset: invalid domain %q", domain)
	}
	o.mu.Lock()
	defer o.mu.Unlock()
	for i := range o.registers {
		if resetAffects(domain, o.registers[i].definition.ResetDomain) {
			o.registers[i].value = o.registers[i].definition.ResetValue & widthMask(o.registers[i].definition.Width)
		}
	}
	return nil
}

func (o *Oracle) findRegister(bir int, offset uint64, width int) (*oracleRegister, uint, error) {
	end := offset + uint64(width)
	if end < offset {
		return nil, 0, fmt.Errorf("transaction address overflows")
	}
	for i := range o.registers {
		reg := &o.registers[i]
		if reg.definition.BIR != bir {
			continue
		}
		regEnd := reg.definition.Offset + uint64(reg.definition.Width)
		if offset >= reg.definition.Offset && end <= regEnd {
			return reg, uint((offset - reg.definition.Offset) * 8), nil
		}
	}
	return nil, 0, fmt.Errorf("unmodeled BIR %d range 0x%x..0x%x", bir, offset, end)
}

func resetAffects(event, register ResetDomain) bool {
	return event == register
}

func validTransactionWidth(width int) bool {
	return width == 1 || width == 2 || width == 4 || width == 8
}

func maskForBytes(width int) uint64 {
	if width == 8 {
		return ^uint64(0)
	}
	return uint64(1)<<(width*8) - 1
}
