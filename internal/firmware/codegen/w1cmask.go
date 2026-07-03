package codegen

import "github.com/sercanarga/pcileechgen/internal/pci"

const (
	statusW1C        = 0xF900
	pmcsrW1C         = 0x8000
	pcieDevStatusW1C = 0x000F
	aerRootStatusW1C = 0x0000007F
)

func setW1C(words []uint32, byteOff int, bits uint32, width int) {
	idx := byteOff / 4
	if idx < 0 || idx >= len(words) {
		return
	}
	if width == 32 {
		words[idx] |= bits
		return
	}
	if byteOff%4 == 2 {
		words[idx] |= bits << 16
	} else {
		words[idx] |= bits & 0xFFFF
	}
}

func W1CMaskWords(cs *pci.ConfigSpace) []uint32 {
	words := make([]uint32, shadowCfgSpaceWords)

	setW1C(words, 0x06, statusW1C, 16)

	for _, c := range pci.ParseCapabilities(cs) {
		switch c.ID {
		case pci.CapIDPowerManagement:
			setW1C(words, c.Offset+0x04, pmcsrW1C, 16)
		case pci.CapIDPCIExpress:
			setW1C(words, c.Offset+0x0A, pcieDevStatusW1C, 16)
		}
	}

	for _, e := range pci.ParseExtCapabilities(cs) {
		if e.ID == pci.ExtCapIDAER {
			setW1C(words, e.Offset+0x04, 0xFFFFFFFF, 32)
			setW1C(words, e.Offset+0x10, 0xFFFFFFFF, 32)
			setW1C(words, e.Offset+0x30, aerRootStatusW1C, 32)
		}
	}

	return words
}

func GenerateW1CMaskCOE(cs *pci.ConfigSpace) string {
	return formatCOE(
		"; PCILeech config-space write-1-to-clear mask (1=bit clears on host write-1)\n",
		W1CMaskWords(cs),
	)
}
