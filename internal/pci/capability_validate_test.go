package pci

import (
	"strings"
	"testing"
)

func TestValidateCapabilityChainsAcceptsValidChains(t *testing.T) {
	cs := NewConfigSpace()
	cs.WriteU16(0x06, 0x0010)
	cs.WriteU8(0x34, 0x40)
	cs.WriteU8(0x40, CapIDPowerManagement)
	cs.WriteU8(0x41, 0x50)
	cs.WriteU8(0x50, CapIDPCIExpress)
	cs.WriteU32(0x100, uint32(ExtCapIDAER)|(1<<16)|(0x140<<20))
	cs.WriteU32(0x140, uint32(ExtCapIDDeviceSerialNumber)|(1<<16))

	if issues := ValidateCapabilityChains(cs); len(issues) != 0 {
		t.Fatalf("ValidateCapabilityChains() = %v, want no issues", issues)
	}
}

func TestValidateCapabilityChainsReportsLoopsAndInvalidPointers(t *testing.T) {
	tests := []struct {
		name string
		make func(*ConfigSpace)
		want string
	}{
		{
			name: "standard loop",
			make: func(cs *ConfigSpace) {
				cs.WriteU16(0x06, 0x0010)
				cs.WriteU8(0x34, 0x40)
				cs.WriteU8(0x40, CapIDPowerManagement)
				cs.WriteU8(0x41, 0x40)
			},
			want: "standard capability loop at 0x040",
		},
		{
			name: "standard pointer below capability area",
			make: func(cs *ConfigSpace) {
				cs.WriteU16(0x06, 0x0010)
				cs.WriteU8(0x34, 0x20)
			},
			want: "standard capability pointer 0x020 outside 0x040-0x0fc",
		},
		{
			name: "extended loop",
			make: func(cs *ConfigSpace) {
				cs.WriteU32(0x100, uint32(ExtCapIDAER)|(1<<16)|(0x140<<20))
				cs.WriteU32(0x140, uint32(ExtCapIDDeviceSerialNumber)|(1<<16)|(0x100<<20))
			},
			want: "extended capability loop at 0x100",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := NewConfigSpace()
			tt.make(cs)
			issues := ValidateCapabilityChains(cs)
			if len(issues) != 1 || !strings.Contains(issues[0], tt.want) {
				t.Fatalf("ValidateCapabilityChains() = %v, want %q", issues, tt.want)
			}
		})
	}
}
