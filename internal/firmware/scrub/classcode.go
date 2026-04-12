package scrub

import (
	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/firmware/overlay"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

// sanitizeClassCodePass overrides the class code for device classes where
// the Windows driver stalls during initialization (e.g., HD Audio waiting
// for CORB/RIRB responses from non-existent codecs), causing slow DMA.
// Changing to a generic class code prevents the specific driver from loading
// while keeping the device functional for DMA purposes.
type sanitizeClassCodePass struct{}

func (p *sanitizeClassCodePass) Name() string { return "sanitize class code" }
func (p *sanitizeClassCodePass) Apply(cs *pci.ConfigSpace, b *board.Board, om *overlay.Map, ctx *ScrubContext) {
	baseClass := (ctx.ClassCode >> 16) & 0xFF
	subClass := (ctx.ClassCode >> 8) & 0xFF

	switch {
	// HD Audio (0x040300) -> Multimedia controller (0x048000)
	// hdaudio.sys stalls on CORB/RIRB init when no real codec responds,
	// causing heavy PCIe polling traffic that congests the FPGA fabric
	// and slows DMA to ~80 KB/s.
	case baseClass == 0x04 && subClass == 0x03:
		om.WriteU8(0x09, 0x04, "class base: Multimedia device")
		om.WriteU8(0x0A, 0x80, "class sub: Multimedia controller")
		om.WriteU8(0x0B, 0x00, "prog-if: 0x00")
	}
}
