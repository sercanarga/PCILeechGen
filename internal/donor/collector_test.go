package donor

import (
	"strings"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/pci"
)

func TestIsAllFF(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		want bool
	}{
		{"nil", nil, false},
		{"empty", []byte{}, false},
		{"single_ff", []byte{0xFF}, true},
		{"all_ff", []byte{0xFF, 0xFF, 0xFF, 0xFF}, true},
		{"mixed", []byte{0xFF, 0x00, 0xFF, 0xFF}, false},
		{"all_zero", []byte{0x00, 0x00, 0x00, 0x00}, false},
		{"first_byte_different", []byte{0x01, 0xFF, 0xFF, 0xFF}, false},
		{"last_byte_different", []byte{0xFF, 0xFF, 0xFF, 0xFE}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAllFF(tt.data); got != tt.want {
				t.Errorf("isAllFF() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateBARContents_AllFF_NVMe(t *testing.T) {
	c := &Collector{}
	ctx := &DeviceContext{
		Device: pci.PCIDevice{
			ClassCode: 0x010802,
			Driver:    "vfio-pci",
		},
		BARs: []pci.BAR{
			{Index: 0, Type: pci.BARTypeMem64, Size: 16384},
		},
		BARContents: map[int][]byte{
			0: makeAllFF(4096),
		},
	}

	err := c.validateBARContents(ctx)
	if err == nil {
		t.Fatal("expected error for all-0xFF BAR on NVMe, got nil")
	}
	if !strings.Contains(err.Error(), "all 0xFF") {
		t.Errorf("error should mention all 0xFF, got: %s", err.Error())
	}
	if !strings.Contains(err.Error(), "Code 10") {
		t.Errorf("error should mention Code 10, got: %s", err.Error())
	}
}

func TestValidateBARContents_AllFF_NonCriticalClass(t *testing.T) {
	c := &Collector{}
	ctx := &DeviceContext{
		Device: pci.PCIDevice{
			ClassCode: 0x030000, // Display/VGA — not BAR-critical
			Driver:    "vfio-pci",
		},
		BARs: []pci.BAR{
			{Index: 0, Type: pci.BARTypeMem64, Size: 16384},
		},
		BARContents: map[int][]byte{
			0: makeAllFF(4096),
		},
	}

	err := c.validateBARContents(ctx)
	if err != nil {
		t.Errorf("non-critical class should not error on all-0xFF, got: %v", err)
	}
}

func TestValidateBARContents_ValidData_NVMe(t *testing.T) {
	c := &Collector{}
	barData := make([]byte, 4096)
	barData[0] = 0x17 // CAP register low byte

	ctx := &DeviceContext{
		Device: pci.PCIDevice{
			ClassCode: 0x010802,
			Driver:    "vfio-pci",
		},
		BARs: []pci.BAR{
			{Index: 0, Type: pci.BARTypeMem64, Size: 16384},
		},
		BARContents: map[int][]byte{
			0: barData,
		},
	}

	err := c.validateBARContents(ctx)
	if err != nil {
		t.Errorf("valid BAR data should pass, got: %v", err)
	}
}

func TestValidateBARContents_EmptyBAR_NVMe(t *testing.T) {
	c := &Collector{}
	ctx := &DeviceContext{
		Device: pci.PCIDevice{
			ClassCode: 0x010802,
			Driver:    "vfio-pci",
		},
		BARs: []pci.BAR{
			{Index: 0, Type: pci.BARTypeMem64, Size: 16384},
		},
		BARContents: map[int][]byte{},
	}

	err := c.validateBARContents(ctx)
	if err == nil {
		t.Fatal("expected error for empty BAR on NVMe")
	}
	if !strings.Contains(err.Error(), "requires BAR data") {
		t.Errorf("error should mention BAR data requirement, got: %s", err.Error())
	}
}

func makeAllFF(size int) []byte {
	data := make([]byte, size)
	for i := range data {
		data[i] = 0xFF
	}
	return data
}
