package scrub

import (
	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/firmware/overlay"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

// ScrubContext holds pre-parsed data shared across all pipeline passes.
type ScrubContext struct {
	Caps      []pci.Capability
	ExtCaps   []pci.ExtCapability
	ClassCode uint32
	VendorID  uint16
	DeviceID  uint16
	Bar0Size  int
	// EmulatePM keeps the donor's Power Management capability faithful (PME
	// support + D-state support advertised, PMCSR power-state writable) instead
	// of stripping it. The IP core is still held in D0 for DMA stability, so the
	// reported D-state is cosmetic - it satisfies config-space readback probes
	// without the device actually powering down.
	EmulatePM bool
}

// ScrubPass is one step in the config space scrubbing pipeline.
// Each pass reads and/or modifies the config space via the overlay map.
type ScrubPass interface {
	Name() string
	Apply(cs *pci.ConfigSpace, b *board.Board, om *overlay.Map, ctx *ScrubContext)
}

// defaultPipeline returns the ordered list of scrub passes.
func defaultPipeline() []ScrubPass {
	return []ScrubPass{
		&clearMiscPass{},
		&sanitizeCmdStatusPass{},
		&injectPCIeCapPass{},
		&scrubPCIeCapPass{},
		&scrubPMCapPass{},
		&scrubAERPass{},
		&normalizeAERMasksPass{},
		&filterExtCapsPass{},
		&clampBARsPass{},
		&relocateMSIXPass{},
		&clampLinkPass{},
		&scrubASPMPass{},
		&clampDeviceCapPass{},
		&applyVendorQuirksPass{},
		&pruneStdCapsPass{},
		&zeroVendorPass{},
		&validateCapChainPass{},
	}
}
