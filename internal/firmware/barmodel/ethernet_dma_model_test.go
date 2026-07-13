package barmodel

import "testing"

func TestIntelE1000BARModelHasDescriptorDMARegisters(t *testing.T) {
	model := BuildIntelE1000BARModel(nil)
	want := []uint32{0x14, 0x20, 0xC0, 0xC8, 0xD0, 0xD8, 0x2800, 0x2804, 0x2808,
		0x2810, 0x2818, 0x3800, 0x3804, 0x3808, 0x3810, 0x3818}
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
