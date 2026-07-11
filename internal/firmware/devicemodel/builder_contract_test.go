package devicemodel

import (
	"testing"

	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

func TestBuildRepresentsUnknownBARSizeExplicitly(t *testing.T) {
	ctx := deterministicContext()
	ctx.BARs = []pci.BAR{{Index: 0, Type: pci.BARTypeMem32}}
	ctx.BARContents = nil
	ctx.BARProfiles = nil
	ctx.MSIXData = nil

	model, err := Build(ctx)
	if err != nil {
		t.Fatalf("Build: %v", err)
	}
	if len(model.BARs) != 1 {
		t.Fatalf("BAR count = %d, want 1", len(model.BARs))
	}
	if model.BARs[0].SizeKnown || model.BARs[0].Size != 0 {
		t.Fatalf("unknown BAR size represented as known: %+v", model.BARs[0])
	}
	if err := model.Validate(); err != nil {
		t.Fatalf("unknown-size BAR model failed validation: %v", err)
	}
}

func TestBuildMarksMeasuredBARSizeKnown(t *testing.T) {
	ctx := deterministicContext()
	ctx.BARs = []pci.BAR{{Index: 0, Type: pci.BARTypeMem32, Size: 0x100}}
	ctx.BARContents = nil
	ctx.BARProfiles = nil
	ctx.MSIXData = nil

	model, err := Build(ctx)
	if err != nil {
		t.Fatalf("Build: %v", err)
	}
	if len(model.BARs) != 1 || !model.BARs[0].SizeKnown || model.BARs[0].Size != 0x100 {
		t.Fatalf("known BAR size was not preserved: %+v", model.BARs)
	}
}

func TestBuildOmitsUnprobedBARRegisters(t *testing.T) {
	ctx := deterministicContext()
	ctx.BARs = []pci.BAR{{Index: 0, Type: pci.BARTypeMem32, Size: 0x100}}
	ctx.BARContents = map[int][]byte{0: make([]byte, 0x100)}
	ctx.BARProfiles = nil
	ctx.MSIXData = nil

	model, err := Build(ctx)
	if err != nil {
		t.Fatalf("Build unprobed model: %v", err)
	}
	for _, register := range model.Registers {
		if register.Space == SpaceBAR {
			t.Fatalf("unprobed BAR content became a behavioral register: %+v", register)
		}
	}

	ctx.BARProfiles = map[int]*donor.BARProfile{
		0: {
			BarIndex: 0,
			Size:     0x100,
			Probes: []donor.BARProbeResult{{
				Offset: 0x20, Original: 0xa5a5a5a5, RWMask: 0x0000ffff, W1CMask: 0x00ff0000,
			}},
		},
	}
	ctx.BARContents[0][0x20] = 0xa5
	ctx.BARContents[0][0x21] = 0xa5
	ctx.BARContents[0][0x22] = 0xa5
	ctx.BARContents[0][0x23] = 0xa5
	model, err = Build(ctx)
	if err != nil {
		t.Fatalf("Build probed model: %v", err)
	}
	var barRegisters []Register
	for _, register := range model.Registers {
		if register.Space == SpaceBAR {
			barRegisters = append(barRegisters, register)
		}
	}
	if len(barRegisters) != 1 || barRegisters[0].Offset != 0x20 {
		t.Fatalf("probed BAR registers = %+v, want one register at 0x20", barRegisters)
	}
}

func TestBuildUsesConfigSpaceAsIdentityAuthority(t *testing.T) {
	ctx := deterministicContext()
	ctx.Device.VendorID = 0xaaaa
	ctx.Device.DeviceID = 0xbbbb
	ctx.Device.SubsysVendorID = 0xcccc
	ctx.Device.SubsysDeviceID = 0xdddd
	ctx.Device.RevisionID = 0xee
	ctx.Device.ClassCode = 0xffffff
	ctx.Device.HeaderType = 0xff
	ctx.ConfigSpace.WriteU16(0x00, 0x1111)
	ctx.ConfigSpace.WriteU16(0x02, 0x2222)
	ctx.ConfigSpace.WriteU8(0x08, 0x33)
	ctx.ConfigSpace.WriteU8(0x09, 0x03)
	ctx.ConfigSpace.WriteU8(0x0a, 0x02)
	ctx.ConfigSpace.WriteU8(0x0b, 0x01)
	ctx.ConfigSpace.WriteU8(0x0e, 0x80)
	ctx.ConfigSpace.WriteU16(0x2c, 0x4444)
	ctx.ConfigSpace.WriteU16(0x2e, 0x5555)

	model, err := Build(ctx)
	if err != nil {
		t.Fatalf("Build: %v", err)
	}
	if model.Name != "pci-1111-2222" {
		t.Fatalf("model name = %q, want config-space identity", model.Name)
	}
	function := model.Functions[0]
	if function.VendorID != 0x1111 || function.DeviceID != 0x2222 ||
		function.SubsystemVendorID != 0x4444 || function.SubsystemDeviceID != 0x5555 ||
		function.RevisionID != 0x33 || function.ClassCode != 0x010203 || function.HeaderType != 0x80 {
		t.Fatalf("function identity did not come from config space: %+v", function)
	}
}

func TestBuildUsesPCICommandAndStatusAccessMasks(t *testing.T) {
	ctx := deterministicContext()
	ctx.ConfigSpace.WriteU16(0x04, 0xffff)
	ctx.ConfigSpace.WriteU16(0x06, 0xffff)

	model, err := Build(ctx)
	if err != nil {
		t.Fatalf("Build: %v", err)
	}

	configExpected := map[string]struct {
		mask   uint64
		access AccessPolicy
	}{
		"command_rw":       {mask: 0x077f, access: AccessRW},
		"command_reserved": {mask: 0xf880, access: AccessReserved},
		"status_rw1c":      {mask: 0xf900, access: AccessRW1C},
		"status_ro":        {mask: 0x06f8, access: AccessRO},
		"status_reserved":  {mask: 0x0007, access: AccessReserved},
	}
	for name, expected := range configExpected {
		found := false
		for _, field := range model.ConfigSpace.Fields {
			if field.Name == name {
				found = true
				if field.Mask != expected.mask || field.Access != expected.access {
					t.Fatalf("config field %s = mask %#x access %q, want %#x %q", name, field.Mask, field.Access, expected.mask, expected.access)
				}
			}
		}
		if !found {
			t.Fatalf("config field %s is missing", name)
		}
	}

	var commandStatus *Register
	for i := range model.Registers {
		if model.Registers[i].Name == "command_status" {
			commandStatus = &model.Registers[i]
			break
		}
	}
	if commandStatus == nil {
		t.Fatal("command_status register is missing")
	}
	registerExpected := map[string]struct {
		mask   uint64
		access AccessPolicy
	}{
		"command_rw":       {mask: 0x0000077f, access: AccessRW},
		"command_reserved": {mask: 0x0000f880, access: AccessReserved},
		"status_rw1c":      {mask: 0xf9000000, access: AccessRW1C},
		"status_ro":        {mask: 0x06f80000, access: AccessRO},
		"status_reserved":  {mask: 0x00070000, access: AccessReserved},
	}
	for name, expected := range registerExpected {
		found := false
		for _, field := range commandStatus.Fields {
			if field.Name == name {
				found = true
				if field.Mask != expected.mask || field.Access != expected.access {
					t.Fatalf("register field %s = mask %#x access %q, want %#x %q", name, field.Mask, field.Access, expected.mask, expected.access)
				}
			}
		}
		if !found {
			t.Fatalf("register field %s is missing", name)
		}
	}
}

func TestBuildCanonicalizesMem64BARPair(t *testing.T) {
	ctx := deterministicContext()
	ctx.BARs = []pci.BAR{
		{Index: 0, Type: pci.BARTypeMem64, Size: 0x1000, Is64Bit: true},
		{Index: 1, Type: pci.BARTypeDisabled},
		{Index: 2, Type: pci.BARTypeDisabled},
		{Index: 3, Type: pci.BARTypeDisabled},
		{Index: 4, Type: pci.BARTypeDisabled},
		{Index: 5, Type: pci.BARTypeDisabled},
	}
	ctx.BARContents = nil
	ctx.BARProfiles = nil
	ctx.MSIXData = nil

	model, err := Build(ctx)
	if err != nil {
		t.Fatalf("Build: %v", err)
	}
	if len(model.BARs) != 1 {
		t.Fatalf("canonical BAR count = %d, want 1: %+v", len(model.BARs), model.BARs)
	}
	bar := model.BARs[0]
	if bar.BIR != 0 || bar.Type != BARTypeMem64 || bar.PairBIR == nil || *bar.PairBIR != 1 {
		t.Fatalf("canonical 64-bit BAR pair = %+v, want BAR0 paired with BIR1", bar)
	}
}

func TestBuildNormalizesParsedDescendingCapabilityChains(t *testing.T) {
	config := pci.NewConfigSpace()
	config.WriteU16(0x00, 0x1234)
	config.WriteU16(0x02, 0x5678)
	config.WriteU16(0x06, 0x0010)
	config.WriteU8(0x34, 0x80)
	config.WriteU8(0x80, pci.CapIDPowerManagement)
	config.WriteU8(0x81, 0x50)
	config.WriteU8(0x50, pci.CapIDMSI)
	config.WriteU8(0x51, 0)
	config.WriteU32(0x100, uint32(pci.ExtCapIDAER)|uint32(1)<<16|uint32(0x300)<<20)
	config.WriteU32(0x300, uint32(pci.ExtCapIDACS)|uint32(1)<<16|uint32(0x200)<<20)
	config.WriteU32(0x200, uint32(pci.ExtCapIDDeviceSerialNumber)|uint32(1)<<16)

	ctx := deterministicContext()
	ctx.ConfigSpace = config
	ctx.Capabilities = pci.ParseCapabilities(config)
	ctx.ExtCapabilities = pci.ParseExtCapabilities(config)
	ctx.BARs = nil
	ctx.BARContents = nil
	ctx.BARProfiles = nil
	ctx.MSIXData = nil

	if len(ctx.Capabilities) != 2 || ctx.Capabilities[0].Offset != 0x80 || ctx.Capabilities[1].Offset != 0x50 {
		t.Fatalf("standard parser did not traverse descending chain: %+v", ctx.Capabilities)
	}
	if len(ctx.ExtCapabilities) != 3 ||
		ctx.ExtCapabilities[0].Offset != 0x100 ||
		ctx.ExtCapabilities[1].Offset != 0x300 ||
		ctx.ExtCapabilities[2].Offset != 0x200 {
		t.Fatalf("extended parser did not traverse descending chain: %+v", ctx.ExtCapabilities)
	}

	model, err := Build(ctx)
	if err != nil {
		t.Fatalf("Build descending capability chains: %v", err)
	}
	if err := model.Validate(); err != nil {
		t.Fatalf("normalized descending capability model failed validation: %v", err)
	}

	lengths := make(map[uint16]uint16)
	next := make(map[uint16]uint16)
	for _, capability := range model.Capabilities {
		lengths[capability.Offset] = capability.Length
		next[capability.Offset] = capability.NextOffset
	}
	if lengths[0x50] != 0x30 || next[0x50] != 0 || next[0x80] != 0x50 {
		t.Fatalf("standard capability normalization = lengths %#v next %#v", lengths, next)
	}
	if lengths[0x100] != 0x100 || lengths[0x200] != 0x100 ||
		next[0x100] != 0x300 || next[0x300] != 0x200 || next[0x200] != 0 {
		t.Fatalf("extended capability normalization = lengths %#v next %#v", lengths, next)
	}
}

func TestBuildUsesCapturedBARResetImageOverLaterProbeValue(t *testing.T) {
	ctx := deterministicContext()
	ctx.BARs = []pci.BAR{{Index: 0, Type: pci.BARTypeMem32, Size: 0x100}}
	ctx.BARContents = map[int][]byte{0: {0x44, 0x33, 0x22, 0x11}}
	ctx.BARProfiles = map[int]*donor.BARProfile{
		0: {
			BarIndex: 0,
			Size:     0x100,
			Probes: []donor.BARProbeResult{
				{Offset: 0, Original: 0xaabbccdd, RWMask: 0x0000ffff, W1CMask: 0x00ff0000},
				{Offset: 4, Original: 0x55667788, RWMask: 0xffffffff},
			},
		},
	}
	ctx.MSIXData = nil

	model, err := Build(ctx)
	if err != nil {
		t.Fatalf("Build changing BAR snapshot: %v", err)
	}
	registers := make(map[uint64]Register)
	for _, register := range model.Registers {
		if register.Space == SpaceBAR {
			registers[register.Offset] = register
		}
	}
	captured, ok := registers[0]
	if !ok {
		t.Fatal("captured BAR register is missing")
	}
	if captured.ResetValue != 0x11223344 || captured.Confidence != ConfidenceMeasured {
		t.Fatalf("captured register reset evidence = value %#x confidence %q", captured.ResetValue, captured.Confidence)
	}
	expectedFields := map[AccessPolicy]struct {
		mask  uint64
		reset uint64
	}{
		AccessRW:   {mask: 0x0000ffff, reset: 0x00003344},
		AccessRW1C: {mask: 0x00ff0000, reset: 0x00220000},
		AccessRO:   {mask: 0xff000000, reset: 0x11000000},
	}
	for _, field := range captured.Fields {
		expected, exists := expectedFields[field.Access]
		if !exists {
			t.Fatalf("unexpected captured field: %+v", field)
		}
		if field.Mask != expected.mask || field.ResetValue != expected.reset {
			t.Fatalf("captured field %q = mask %#x reset %#x, want %#x %#x", field.Access, field.Mask, field.ResetValue, expected.mask, expected.reset)
		}
		delete(expectedFields, field.Access)
	}
	if len(expectedFields) != 0 {
		t.Fatalf("captured register omitted access evidence: %#v", expectedFields)
	}
	uncaptured, ok := registers[4]
	if !ok {
		t.Fatal("uncaptured BAR register is missing")
	}
	if uncaptured.ResetValue != 0x55667788 {
		t.Fatalf("uncaptured register reset = %#x, want probe original %#x", uncaptured.ResetValue, uint64(0x55667788))
	}
}
