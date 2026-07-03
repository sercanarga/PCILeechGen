package codegen

import (
	"strings"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/pci"
)

func TestW1CMask_StatusRegisterAlwaysMarked(t *testing.T) {
	cs := pci.NewConfigSpace()
	words := W1CMaskWords(cs)

	if len(words) != shadowCfgSpaceWords {
		t.Fatalf("expected %d words, got %d", shadowCfgSpaceWords, len(words))
	}
	if words[1] != 0xF9000000 {
		t.Errorf("status W1C mask: want 0xF9000000 at word 1, got 0x%08X", words[1])
	}
	if words[0] != 0 {
		t.Errorf("word 0 (id regs) must have no W1C bits, got 0x%08X", words[0])
	}
}

func TestSetW1C_Placement(t *testing.T) {
	cases := []struct {
		name  string
		off   int
		bits  uint32
		width int
		idx   int
		want  uint32
	}{
		{"16bit high half (status 0x06)", 0x06, 0xF900, 16, 1, 0xF9000000},
		{"16bit low half (pmcsr cap+0x04)", 0x44, 0x8000, 16, 0x11, 0x00008000},
		{"32bit aligned (AER status)", 0x104, 0xFFFFFFFF, 32, 0x41, 0xFFFFFFFF},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			words := make([]uint32, shadowCfgSpaceWords)
			setW1C(words, c.off, c.bits, c.width)
			if words[c.idx] != c.want {
				t.Errorf("word[%#x] = 0x%08X, want 0x%08X", c.idx, words[c.idx], c.want)
			}
		})
	}
}

func TestSetW1C_OutOfRangeNoPanic(t *testing.T) {
	words := make([]uint32, shadowCfgSpaceWords)
	setW1C(words, 0x4000, 0xFFFFFFFF, 32)
}

func TestGenerateW1CMaskCOE_Format(t *testing.T) {
	cs := pci.NewConfigSpace()
	coe := GenerateW1CMaskCOE(cs)
	if !strings.Contains(coe, "memory_initialization_radix=16;") {
		t.Error("COE missing radix header")
	}
	if n := strings.Count(coe, "\n"); n < shadowCfgSpaceWords {
		t.Errorf("COE has too few lines: %d", n)
	}
}
