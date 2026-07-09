package devicemodel

import "testing"

func oracleContractModel() *Model {
	model := validTestModel()
	model.MSIX = nil
	model.BARs = []BAR{{
		BIR:          0,
		Type:         BARTypeMem32,
		Size:         0x100,
		SizeKnown:    true,
		AddressWidth: 32,
		ResetImage:   make([]byte, 0x100),
	}}
	model.Registers = []Register{
		{
			Name: "read_only", Space: SpaceBAR, BIR: 0, Offset: 0x00, Width: 4,
			ResetDomain: ResetPowerOn, ResetValue: 0xa5a5a5a5,
			Fields: []RegisterField{{Name: "value", Mask: 0xffffffff, Access: AccessRO, ResetValue: 0xa5a5a5a5}},
		},
		{
			Name: "byte_lanes", Space: SpaceBAR, BIR: 0, Offset: 0x04, Width: 4,
			ResetDomain: ResetPowerOn, ResetValue: 0x11223344,
			Fields: []RegisterField{{Name: "value", Mask: 0xffffffff, Access: AccessRW, ResetValue: 0x11223344}},
		},
		{
			Name: "mixed_rw_w1c", Space: SpaceBAR, BIR: 0, Offset: 0x08, Width: 4,
			ResetDomain: ResetFunction, ResetValue: 0x00f00012,
			Fields: []RegisterField{
				{Name: "control", Mask: 0x000000ff, Access: AccessRW, ResetValue: 0x12},
				{Name: "status", Mask: 0x00ff0000, Access: AccessW1C, ResetValue: 0x00f00000},
			},
		},
		{
			Name: "read_clear", Space: SpaceBAR, BIR: 0, Offset: 0x0c, Width: 2,
			ResetDomain: ResetFunction, ResetValue: 0xabcd,
			Fields: []RegisterField{
				{Name: "events", Mask: 0x00ff, Access: AccessRC, ResetValue: 0x00cd},
				{Name: "control", Mask: 0xff00, Access: AccessRW, ResetValue: 0xab00},
			},
		},
		{
			Name: "write_one_set", Space: SpaceBAR, BIR: 0, Offset: 0x10, Width: 1,
			ResetDomain: ResetSoftware, ResetValue: 0x0f,
			Fields: []RegisterField{{Name: "flags", Mask: 0xff, Access: AccessW1S, ResetValue: 0x0f}},
		},
		{
			Name: "write_zero_clear", Space: SpaceBAR, BIR: 0, Offset: 0x11, Width: 1,
			ResetDomain: ResetSoftware, ResetValue: 0xff,
			Fields: []RegisterField{{Name: "flags", Mask: 0xff, Access: AccessW0C, ResetValue: 0xff}},
		},
		{
			Name: "reserved", Space: SpaceBAR, BIR: 0, Offset: 0x12, Width: 1,
			ResetDomain: ResetPowerOn, ResetValue: 0x5a,
			Fields: []RegisterField{{Name: "reserved", Mask: 0xff, Access: AccessReserved, ResetValue: 0x5a}},
		},
		{
			Name: "fundamental_domain", Space: SpaceBAR, BIR: 0, Offset: 0x14, Width: 4,
			ResetDomain: ResetFundamental, ResetValue: 0x11111111,
			Fields: []RegisterField{{Name: "value", Mask: 0xffffffff, Access: AccessRW, ResetValue: 0x11111111}},
		},
		{
			Name: "function_domain", Space: SpaceBAR, BIR: 0, Offset: 0x18, Width: 4,
			ResetDomain: ResetFunction, ResetValue: 0x22222222,
			Fields: []RegisterField{{Name: "value", Mask: 0xffffffff, Access: AccessRW, ResetValue: 0x22222222}},
		},
	}
	for i := range model.Registers {
		register := &model.Registers[i]
		register.Confidence = ConfidenceSpecified
		for lane := range int(register.Width) {
			model.BARs[0].ResetImage[int(register.Offset)+lane] = byte(register.ResetValue >> (lane * 8))
		}
	}
	return model
}

func newContractOracle(t *testing.T) *Oracle {
	t.Helper()
	oracle, err := NewOracle(oracleContractModel())
	if err != nil {
		t.Fatalf("NewOracle: %v", err)
	}
	return oracle
}

func readOracle(t *testing.T, oracle *Oracle, offset uint64, width int) uint64 {
	t.Helper()
	value, err := oracle.Read(0, offset, width)
	if err != nil {
		t.Fatalf("Read(BAR0, %#x, %d): %v", offset, width, err)
	}
	return value
}

func writeOracle(t *testing.T, oracle *Oracle, offset uint64, width int, value uint64, byteEnable uint8) {
	t.Helper()
	if err := oracle.Write(0, offset, width, value, byteEnable); err != nil {
		t.Fatalf("Write(BAR0, %#x, %d, %#x, %#x): %v", offset, width, value, byteEnable, err)
	}
}

func TestOracleByteEnablesAreLittleEndianLaneMasks(t *testing.T) {
	oracle := newContractOracle(t)
	writeOracle(t, oracle, 0x04, 4, 0xaabbccdd, 0b0101)
	if got, want := readOracle(t, oracle, 0x04, 4), uint64(0x11bb33dd); got != want {
		t.Fatalf("partial write = %#08x, want %#08x", got, want)
	}
}

func TestOracleReadOnlyBitsRemainStable(t *testing.T) {
	oracle := newContractOracle(t)
	before := readOracle(t, oracle, 0x00, 4)
	writeOracle(t, oracle, 0x00, 4, 0x00000000, 0xf)
	writeOracle(t, oracle, 0x00, 4, 0xffffffff, 0xf)
	if after := readOracle(t, oracle, 0x00, 4); after != before {
		t.Fatalf("RO register changed from %#08x to %#08x", before, after)
	}
}

func TestOracleReservedBitsRemainStable(t *testing.T) {
	oracle := newContractOracle(t)
	writeOracle(t, oracle, 0x12, 1, 0, 0x1)
	if got, want := readOracle(t, oracle, 0x12, 1), uint64(0); got != want {
		t.Fatalf("reserved register read = %#02x, want hardwired %#02x", got, want)
	}
}

func TestOracleMixedRWAndW1CFields(t *testing.T) {
	oracle := newContractOracle(t)
	writeOracle(t, oracle, 0x08, 4, 0x00550034, 0xf)
	if got, want := readOracle(t, oracle, 0x08, 4), uint64(0x00a00034); got != want {
		t.Fatalf("mixed RW/W1C write = %#08x, want %#08x", got, want)
	}
}

func TestOracleReadClearReturnsOldValueThenClearsOnlyRCBits(t *testing.T) {
	oracle := newContractOracle(t)
	if got, want := readOracle(t, oracle, 0x0c, 2), uint64(0xabcd); got != want {
		t.Fatalf("first RC read = %#04x, want %#04x", got, want)
	}
	if got, want := readOracle(t, oracle, 0x0c, 2), uint64(0xab00); got != want {
		t.Fatalf("second RC read = %#04x, want preserved RW bits %#04x", got, want)
	}
}

func TestOracleWriteOneSetAndWriteZeroClear(t *testing.T) {
	oracle := newContractOracle(t)

	writeOracle(t, oracle, 0x10, 1, 0x30, 0x1)
	if got, want := readOracle(t, oracle, 0x10, 1), uint64(0x3f); got != want {
		t.Fatalf("W1S result = %#02x, want %#02x", got, want)
	}

	writeOracle(t, oracle, 0x11, 1, 0xf0, 0x1)
	if got, want := readOracle(t, oracle, 0x11, 1), uint64(0xf0); got != want {
		t.Fatalf("W0C result = %#02x, want %#02x", got, want)
	}
}

func TestOracleResetDomainsAreIndependent(t *testing.T) {
	oracle := newContractOracle(t)
	writeOracle(t, oracle, 0x14, 4, 0, 0xf)
	writeOracle(t, oracle, 0x18, 4, 0, 0xf)

	if err := oracle.Reset(ResetFunction); err != nil {
		t.Fatalf("Reset(function): %v", err)
	}
	if got := readOracle(t, oracle, 0x14, 4); got != 0 {
		t.Fatalf("function reset changed fundamental-domain register to %#x", got)
	}
	if got, want := readOracle(t, oracle, 0x18, 4), uint64(0x22222222); got != want {
		t.Fatalf("function-domain reset = %#x, want %#x", got, want)
	}

	if err := oracle.Reset(ResetFundamental); err != nil {
		t.Fatalf("Reset(fundamental): %v", err)
	}
	if got, want := readOracle(t, oracle, 0x14, 4), uint64(0x11111111); got != want {
		t.Fatalf("fundamental-domain reset = %#x, want %#x", got, want)
	}
	if got, want := readOracle(t, oracle, 0x18, 4), uint64(0x22222222); got != want {
		t.Fatalf("fundamental reset disturbed function-domain register: got %#x, want %#x", got, want)
	}
}
