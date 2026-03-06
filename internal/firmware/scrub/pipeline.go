package scrub

import (
	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/firmware/overlay"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

// ScrubContext holds pre-parsed data shared across all pipeline passes.
type ScrubContext struct {
	Caps    []pci.Capability
	ExtCaps []pci.ExtCapability
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
		&scrubPCIeCapPass{},
		&scrubPMCapPass{},
		&scrubAERPass{},
		&normalizeAERMasksPass{},
		&scrubASPMPass{},
		&filterExtCapsPass{},
		&clampBARsPass{},
		&relocateMSIXPass{},
		&clampLinkPass{},
		&clampDeviceCapPass{},
		&zeroVendorPass{},
		&applyVendorQuirksPass{},
		&pruneStdCapsPass{},
		&validateCapChainPass{},
	}
}
