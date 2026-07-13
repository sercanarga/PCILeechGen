package barmodel

import "testing"

func TestEthernetBARModelHasDescriptorDMARegisters(t *testing.T) {
	model := buildEthernetBARModel(nil)
	want := []uint32{0x14, 0x20, 0xC0, 0xD0, 0x280, 0x284, 0x288,
		0x2810, 0x2818, 0x380, 0x384, 0x388, 0x3810, 0x3818}
	for _, offset := range want {
		found := false
		for _, reg := range model.Registers {
			if reg.Offset == offset {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Ethernet BAR missing register 0x%X", offset)
		}
	}
}
