package devicemodel

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

func validConfigResetImage() []byte {
	image := make([]byte, pci.ConfigSpaceLegacySize)
	image[0x00], image[0x01] = 0x34, 0x12
	image[0x02], image[0x03] = 0x78, 0x56
	image[0x08] = 0x01
	image[0x0b] = 0x02
	image[0x2c], image[0x2d] = 0x34, 0x12
	image[0x2e], image[0x2f] = 0xcd, 0xab
	return image
}

func validTestModel() *Model {
	pair := 3
	return &Model{
		SchemaVersion: CurrentSchemaVersion,
		Name:          "contract-device",
		Functions: []Function{{
			BDF:               "0000:03:00.0",
			VendorID:          0x1234,
			DeviceID:          0x5678,
			SubsystemVendorID: 0x1234,
			SubsystemDeviceID: 0xabcd,
			RevisionID:        1,
			ClassCode:         0x020000,
		}},
		ConfigSpace: ConfigSpace{
			Size:       pci.ConfigSpaceLegacySize,
			ResetImage: validConfigResetImage(),
		},
		BARs: []BAR{
			{BIR: 0, Type: BARTypeMem32, Size: 0x100, SizeKnown: true, AddressWidth: 32, ResetImage: make([]byte, 0x100)},
			{BIR: 2, Type: BARTypeMem64, Size: 0x1000, SizeKnown: true, Prefetchable: true, AddressWidth: 64, PairBIR: &pair, ResetImage: make([]byte, 0x1000)},
		},
		MSIX: &MSIXDescriptor{
			CapabilityOffset: 0,
			TableSize:        2,
			TableBIR:         0,
			TableOffset:      0x40,
			PBABIR:           0,
			PBAOffset:        0x80,
		},
		Confidence: Confidence{Overall: ConfidenceMeasured, Evidence: []string{"unit-test fixture"}},
		Provenance: Provenance{
			Source:      "test",
			ToolVersion: "test-version",
			CollectedAt: time.Date(2026, 1, 2, 3, 4, 5, 0, time.UTC),
			DonorBDF:    "0000:03:00.0",
		},
	}
}

func deterministicContext() *donor.DeviceContext {
	cs := pci.NewConfigSpaceFromBytes(make([]byte, pci.ConfigSpaceLegacySize))
	cs.WriteU16(0x00, 0x1234)
	cs.WriteU16(0x02, 0x5678)
	cs.WriteU32(0x10, 0x80000000)
	cs.WriteU32(0x18, 0x9000000c)
	cs.WriteU32(0x1c, 0x00000001)
	return &donor.DeviceContext{
		CollectedAt: time.Date(2026, 1, 2, 3, 4, 5, 0, time.UTC),
		ToolVersion: "test-version",
		Hostname:    "test-host",
		Device: pci.PCIDevice{
			BDF:            pci.BDF{Domain: 0, Bus: 3, Device: 0, Function: 0},
			VendorID:       0x1234,
			DeviceID:       0x5678,
			SubsysVendorID: 0x1234,
			SubsysDeviceID: 0xabcd,
			RevisionID:     1,
			ClassCode:      0x020000,
		},
		ConfigSpace: cs,
		BARs: []pci.BAR{
			{Index: 2, Type: pci.BARTypeMem64, Size: 0x1000, Is64Bit: true, Prefetchable: true},
			{Index: 0, Type: pci.BARTypeMem32, Size: 0x100},
		},
		BARContents: map[int][]byte{
			2: make([]byte, 0x1000),
			0: make([]byte, 0x100),
		},
		MSIXData: &donor.MSIXData{
			TableSize: 2, TableBIR: 0, TableOffset: 0x40,
			PBABIR: 0, PBAOffset: 0x80,
		},
	}
}

func TestBuildIsDeterministicAndCanonical(t *testing.T) {
	first, err := Build(deterministicContext())
	if err != nil {
		t.Fatalf("Build(first): %v", err)
	}
	second, err := Build(deterministicContext())
	if err != nil {
		t.Fatalf("Build(second): %v", err)
	}

	firstJSON, err := first.ToJSON()
	if err != nil {
		t.Fatalf("ToJSON(first): %v", err)
	}
	secondJSON, err := second.ToJSON()
	if err != nil {
		t.Fatalf("ToJSON(second): %v", err)
	}
	if string(firstJSON) != string(secondJSON) {
		t.Fatalf("identical donor snapshots produced different JSON\nfirst: %s\nsecond: %s", firstJSON, secondJSON)
	}
	if len(first.BARs) != 2 || first.BARs[0].BIR != 0 || first.BARs[1].BIR != 2 {
		t.Fatalf("BARs are not in canonical BIR order: %+v", first.BARs)
	}
	if first.BARs[1].PairBIR == nil || *first.BARs[1].PairBIR != 3 {
		t.Fatalf("64-bit BAR pair was not represented explicitly: %+v", first.BARs[1])
	}
}

func TestJSONRoundTripPreservesModel(t *testing.T) {
	model := validTestModel()
	data, err := model.ToJSON()
	if err != nil {
		t.Fatalf("ToJSON: %v", err)
	}
	got, err := FromJSON(data)
	if err != nil {
		t.Fatalf("FromJSON: %v", err)
	}
	if !reflect.DeepEqual(got, model) {
		t.Fatalf("round trip changed model\ngot:  %#v\nwant: %#v", got, model)
	}

	again, err := got.ToJSON()
	if err != nil {
		t.Fatalf("ToJSON(round-tripped): %v", err)
	}
	if string(again) != string(data) {
		t.Fatalf("JSON is not canonical after round trip\nfirst:  %s\nsecond: %s", data, again)
	}
}

func TestJSONRoundTripPreservesRawNVMeIdentify(t *testing.T) {
	model := validTestModel()
	model.NVMeIdentify = &NVMeIdentify{
		Controller: make([]byte, 4096),
		Namespace:  make([]byte, 4096),
	}
	model.NVMeIdentify.Controller[0x18] = 'D'
	model.NVMeIdentify.Namespace[0x19] = 1
	data, err := model.ToJSON()
	if err != nil {
		t.Fatal(err)
	}
	got, err := FromJSON(data)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got.NVMeIdentify, model.NVMeIdentify) {
		t.Fatalf("raw NVMe Identify changed across JSON round trip")
	}
}

func TestFromJSONRejectsUnsupportedSchemaAndUnknownFields(t *testing.T) {
	data, err := validTestModel().ToJSON()
	if err != nil {
		t.Fatal(err)
	}

	var document map[string]any
	if err := json.Unmarshal(data, &document); err != nil {
		t.Fatal(err)
	}

	t.Run("future schema", func(t *testing.T) {
		document["schema_version"] = float64(CurrentSchemaVersion + 1)
		bad, err := json.Marshal(document)
		if err != nil {
			t.Fatal(err)
		}
		if _, err := FromJSON(bad); err == nil {
			t.Fatal("FromJSON accepted an unsupported schema version")
		}
	})

	t.Run("unknown field", func(t *testing.T) {
		document["schema_version"] = float64(CurrentSchemaVersion)
		document["silently_ignored_contract_change"] = true
		bad, err := json.Marshal(document)
		if err != nil {
			t.Fatal(err)
		}
		if _, err := FromJSON(bad); err == nil {
			t.Fatal("FromJSON accepted an unknown top-level field")
		}
	})
}

func TestValidateBARPairInvariants(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*Model)
	}{
		{
			name: "64-bit BAR must name adjacent upper BIR",
			mutate: func(m *Model) {
				wrong := 4
				m.BARs[1].PairBIR = &wrong
			},
		},
		{
			name: "64-bit BAR cannot start at BIR5",
			mutate: func(m *Model) {
				pair := 6
				m.BARs[1].BIR = 5
				m.BARs[1].PairBIR = &pair
			},
		},
		{
			name: "32-bit BAR cannot consume a pair",
			mutate: func(m *Model) {
				pair := 1
				m.BARs[0].PairBIR = &pair
			},
		},
		{
			name: "paired upper BIR cannot have an independent descriptor",
			mutate: func(m *Model) {
				m.BARs = append(m.BARs, BAR{BIR: 3, Type: BARTypeDisabled})
			},
		},
		{
			name: "BAR aperture must be a power of two",
			mutate: func(m *Model) {
				m.BARs[0].Size = 0x180
				m.BARs[0].ResetImage = make([]byte, 0x180)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			model := validTestModel()
			tt.mutate(model)
			if err := model.Validate(); err == nil {
				t.Fatal("Validate accepted an unlawful BAR topology")
			}
		})
	}
}

func TestValidateMSIXInvariants(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*Model)
	}{
		{"table BIR must exist", func(m *Model) { m.MSIX.TableBIR = 1 }},
		{"PBA BIR must exist", func(m *Model) { m.MSIX.PBABIR = 1 }},
		{"table must fit BAR", func(m *Model) { m.MSIX.TableOffset = 0xf8 }},
		{"PBA must fit BAR", func(m *Model) { m.MSIX.PBAOffset = 0x100 }},
		{"table offset must be aligned", func(m *Model) { m.MSIX.TableOffset = 0x43 }},
		{"PBA offset must be aligned", func(m *Model) { m.MSIX.PBAOffset = 0x83 }},
		{"table must contain a vector", func(m *Model) { m.MSIX.TableSize = 0 }},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			model := validTestModel()
			tt.mutate(model)
			if err := model.Validate(); err == nil {
				t.Fatal("Validate accepted an invalid MSI-X layout")
			}
		})
	}
}

func TestValidateRejectsOverlappingRegisterFields(t *testing.T) {
	model := validTestModel()
	model.Registers = []Register{{
		Name:        "status_control",
		Space:       SpaceBAR,
		BIR:         0,
		Offset:      0,
		Width:       4,
		ResetDomain: ResetFunction,
		Confidence:  ConfidenceSpecified,
		Fields: []RegisterField{
			{Name: "control", Mask: 0x000000ff, Access: AccessRW},
			{Name: "status", Mask: 0x000000f0, Access: AccessW1C},
		},
	}}
	if err := model.Validate(); err == nil {
		t.Fatal("Validate accepted overlapping field masks")
	}
}

func TestValidateRejectsInvalidFieldPoliciesAndMasks(t *testing.T) {
	tests := []struct {
		name  string
		field RegisterField
	}{
		{
			name:  "unknown access policy",
			field: RegisterField{Name: "bad_access", Mask: 0xff, Access: AccessPolicy("write-anything")},
		},
		{
			name:  "mask outside register width",
			field: RegisterField{Name: "too_wide", Mask: 0x100, Access: AccessRW},
		},
		{
			name:  "empty field mask",
			field: RegisterField{Name: "empty", Mask: 0, Access: AccessRW},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			model := validTestModel()
			model.Registers = []Register{{
				Name:        "invalid",
				Space:       SpaceBAR,
				BIR:         0,
				Offset:      0,
				Width:       1,
				ResetDomain: ResetFunction,
				Confidence:  ConfidenceSpecified,
				Fields:      []RegisterField{tt.field},
			}}
			if err := model.Validate(); err == nil {
				t.Fatal("Validate accepted an invalid register field")
			}
		})
	}
}

func TestValidateRejectsRegisterFieldResetValueMismatch(t *testing.T) {
	model := validTestModel()
	model.Registers = []Register{{
		Name:        "reset_mismatch",
		Space:       SpaceBAR,
		BIR:         0,
		Offset:      0,
		Width:       1,
		ResetDomain: ResetFunction,
		ResetValue:  0xa5,
		Confidence:  ConfidenceSpecified,
		Fields: []RegisterField{{
			Name:       "value",
			Mask:       0xff,
			Access:     AccessRW,
			ResetValue: 0x5a,
		}},
	}}
	if err := model.Validate(); err == nil {
		t.Fatal("Validate accepted a register field reset value inconsistent with its register")
	}
}

func TestValidateConfigFieldResetValueUsesLittleEndianResetImage(t *testing.T) {
	model := validTestModel()
	model.ConfigSpace.ResetImage[0x20] = 0x34
	model.ConfigSpace.ResetImage[0x21] = 0x12
	model.ConfigSpace.Fields = []ConfigField{{
		Name:       "little_endian",
		Offset:     0x20,
		Width:      2,
		Mask:       0x0fff,
		Access:     AccessRO,
		ResetValue: 0x0234,
	}}
	if err := model.Validate(); err != nil {
		t.Fatalf("Validate rejected matching little-endian config reset value: %v", err)
	}

	model.ConfigSpace.Fields[0].ResetValue = 0x0235
	if err := model.Validate(); err == nil {
		t.Fatal("Validate accepted a config field reset value inconsistent with reset image bytes")
	}
}

func TestValidateRejectsEachFunctionIdentityMismatchWithConfigImage(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*Function)
	}{
		{name: "vendor ID", mutate: func(function *Function) { function.VendorID ^= 1 }},
		{name: "device ID", mutate: func(function *Function) { function.DeviceID ^= 1 }},
		{name: "subsystem vendor ID", mutate: func(function *Function) { function.SubsystemVendorID ^= 1 }},
		{name: "subsystem device ID", mutate: func(function *Function) { function.SubsystemDeviceID ^= 1 }},
		{name: "revision ID", mutate: func(function *Function) { function.RevisionID ^= 1 }},
		{name: "class code", mutate: func(function *Function) { function.ClassCode ^= 1 }},
		{name: "header type", mutate: func(function *Function) { function.HeaderType ^= 1 }},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			model := validTestModel()
			tt.mutate(&model.Functions[0])
			if err := model.Validate(); err == nil {
				t.Fatal("Validate accepted function identity inconsistent with config reset image")
			}
		})
	}
}

func TestValidateRejectsConfigRegisterResetValueMismatchWithImage(t *testing.T) {
	model := validTestModel()
	model.Registers = []Register{{
		Name:        "vendor_device",
		Space:       SpaceConfig,
		BIR:         ConfigBIR,
		Offset:      0,
		Width:       4,
		ResetDomain: ResetPowerOn,
		ResetValue:  0x56781235,
		Confidence:  ConfidenceSpecified,
	}}
	if err := model.Validate(); err == nil {
		t.Fatal("Validate accepted config register reset value inconsistent with config reset image")
	}
}

func TestValidateRejectsBARRegisterResetValueMismatchWithImage(t *testing.T) {
	model := validTestModel()
	model.Registers = []Register{{
		Name:        "bar_reset_mismatch",
		Space:       SpaceBAR,
		BIR:         0,
		Offset:      0,
		Width:       4,
		ResetDomain: ResetPowerOn,
		ResetValue:  1,
		Confidence:  ConfidenceSpecified,
	}}
	if err := model.Validate(); err == nil {
		t.Fatal("Validate accepted BAR register reset value inconsistent with BAR reset image")
	}
}

func TestValidateRejectsMissingProvenanceFields(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*Provenance)
	}{
		{name: "source", mutate: func(provenance *Provenance) { provenance.Source = "" }},
		{name: "tool version", mutate: func(provenance *Provenance) { provenance.ToolVersion = "" }},
		{name: "collection time", mutate: func(provenance *Provenance) { provenance.CollectedAt = time.Time{} }},
		{name: "donor BDF", mutate: func(provenance *Provenance) { provenance.DonorBDF = "" }},
		{name: "blank source", mutate: func(provenance *Provenance) { provenance.Source = " \t" }},
		{name: "blank tool version", mutate: func(provenance *Provenance) { provenance.ToolVersion = "\n" }},
		{name: "blank donor BDF", mutate: func(provenance *Provenance) { provenance.DonorBDF = " " }},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			model := validTestModel()
			tt.mutate(&model.Provenance)
			if err := model.Validate(); err == nil {
				t.Fatal("Validate accepted incomplete provenance")
			}
		})
	}
}

func TestValidateRequiresExactlyOneFunction(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*Model)
	}{
		{name: "missing", mutate: func(model *Model) { model.Functions = nil }},
		{
			name: "multiple",
			mutate: func(model *Model) {
				second := model.Functions[0]
				second.BDF = "0000:04:00.0"
				model.Functions = append(model.Functions, second)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			model := validTestModel()
			tt.mutate(model)
			if err := model.Validate(); err == nil {
				t.Fatal("Validate accepted a model without exactly one function")
			}
		})
	}
}
