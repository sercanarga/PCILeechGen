package devicemodel

import "testing"

func TestValidateAcceptsDescendingStandardCapabilityLink(t *testing.T) {
	model := validTestModel()
	model.Capabilities = []Capability{
		{ID: 0x01, Name: "power_management", Offset: 0x80, NextOffset: 0x50, Length: 4, Data: []byte{0x01, 0x50, 0, 0}},
		{ID: 0x05, Name: "msi", Offset: 0x50, NextOffset: 0, Length: 4, Data: []byte{0x05, 0, 0, 0}},
	}
	if err := model.Validate(); err != nil {
		t.Fatalf("descending standard capability link was rejected: %v", err)
	}
}

func TestValidateRejectsStandardCapabilityCycle(t *testing.T) {
	model := validTestModel()
	model.Capabilities = []Capability{
		{ID: 0x01, Name: "power_management", Offset: 0x80, NextOffset: 0x50, Length: 4, Data: []byte{0x01, 0x50, 0, 0}},
		{ID: 0x05, Name: "msi", Offset: 0x50, NextOffset: 0x80, Length: 4, Data: []byte{0x05, 0x80, 0, 0}},
	}
	if err := model.Validate(); err == nil {
		t.Fatal("capability cycle was accepted")
	}
}

func TestValidateRejectsOverlappingCapabilityRanges(t *testing.T) {
	model := validTestModel()
	model.Capabilities = []Capability{
		{ID: 0x01, Name: "power_management", Offset: 0x50, Length: 8, Data: make([]byte, 8)},
		{ID: 0x05, Name: "msi", Offset: 0x54, Length: 4, Data: make([]byte, 4)},
	}
	if err := model.Validate(); err == nil {
		t.Fatal("overlapping capability byte ranges were accepted")
	}
}
