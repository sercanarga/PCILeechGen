package scrub

import (
	"log/slog"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/firmware/overlay"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

type clearMiscPass struct{}

func (p *clearMiscPass) Name() string { return "clear misc registers" }
func (p *clearMiscPass) Apply(cs *pci.ConfigSpace, b *board.Board, om *overlay.Map, ctx *ScrubContext) {
	om.WriteU8(0x0F, 0x00, "clear BIST register")
	om.WriteU8(0x3C, 0x00, "clear Interrupt Line")
	om.WriteU8(0x0D, 0x00, "clear Latency Timer")
	om.WriteU8(0x0C, 0x00, "clear Cache Line Size")
}

type sanitizeCmdStatusPass struct{}

func (p *sanitizeCmdStatusPass) Name() string { return "sanitize Command/Status" }
func (p *sanitizeCmdStatusPass) Apply(cs *pci.ConfigSpace, b *board.Board, om *overlay.Map, ctx *ScrubContext) {
	om.WriteU16(0x04, (cs.Command()&cmdMask)|cmdForce, "sanitize Command register (force BME+MSE)")
	om.WriteU16(0x06, cs.Status()&statusMask, "sanitize Status register")
}

type scrubPCIeCapPass struct{}

func (p *scrubPCIeCapPass) Name() string { return "scrub PCIe capability" }
func (p *scrubPCIeCapPass) Apply(cs *pci.ConfigSpace, b *board.Board, om *overlay.Map, ctx *ScrubContext) {
	for _, cap := range ctx.Caps {
		if cap.ID == pci.CapIDPCIExpress && cap.Offset+10 < pci.ConfigSpaceLegacySize {
			om.WriteU16(cap.Offset+10, 0x0000, "clear PCIe Device Status")
			if cap.Offset+18 < pci.ConfigSpaceLegacySize {
				lstatus := cs.ReadU16(cap.Offset+18) & 0x3FFF
				om.WriteU16(cap.Offset+18, lstatus, "clear PCIe Link Status RW1C bits")
			}
		}
	}
}

type scrubPMCapPass struct{}

func (p *scrubPMCapPass) Name() string { return "scrub PM capability" }
func (p *scrubPMCapPass) Apply(cs *pci.ConfigSpace, b *board.Board, om *overlay.Map, ctx *ScrubContext) {
	for _, cap := range ctx.Caps {
		if cap.ID == pci.CapIDPowerManagement && cap.Offset+4 < pci.ConfigSpaceLegacySize {
			pmcsr := cs.ReadU16(cap.Offset + 4)
			pmcsr &= 0xFFFC // force D0
			pmcsr &= 0x7FFF // clear PME_Status
			pmcsr |= 0x0008 // NoSoftReset
			om.WriteU16(cap.Offset+4, pmcsr, "PM: D0, NoSoftReset, clear PME_Status")
		}
	}
}

type scrubAERPass struct{}

func (p *scrubAERPass) Name() string { return "scrub AER status" }
func (p *scrubAERPass) Apply(cs *pci.ConfigSpace, b *board.Board, om *overlay.Map, ctx *ScrubContext) {
	if cs.Size < pci.ConfigSpaceSize {
		return
	}
	for _, cap := range ctx.ExtCaps {
		if cap.ID == pci.ExtCapIDAER {
			if cap.Offset+8 <= pci.ConfigSpaceSize {
				om.WriteU32(cap.Offset+4, 0, "clear AER uncorrectable error status")
			}
			if cap.Offset+20 <= pci.ConfigSpaceSize {
				om.WriteU32(cap.Offset+16, 0, "clear AER correctable error status")
			}
			if cap.Offset+32 <= pci.ConfigSpaceSize {
				om.WriteU32(cap.Offset+28, 0, "clear AER root error status")
			}
		}
	}
}

type filterExtCapsPass struct{}

func (p *filterExtCapsPass) Name() string { return "filter ext capabilities" }
func (p *filterExtCapsPass) Apply(cs *pci.ConfigSpace, b *board.Board, om *overlay.Map, ctx *ScrubContext) {
	if cs.Size >= pci.ConfigSpaceSize {
		FilterExtCapabilities(cs)
	}
}

type clampBARsPass struct{}

func (p *clampBARsPass) Name() string { return "clamp BARs to FPGA" }
func (p *clampBARsPass) Apply(cs *pci.ConfigSpace, b *board.Board, om *overlay.Map, ctx *ScrubContext) {
	clampBARsToFPGA(cs, om)
}

type relocateMSIXPass struct{}

func (p *relocateMSIXPass) Name() string { return "relocate MSI-X to BRAM" }
func (p *relocateMSIXPass) Apply(cs *pci.ConfigSpace, b *board.Board, om *overlay.Map, ctx *ScrubContext) {
	relocateMSIXToBRAM(cs, om, ctx.Caps)
}

type clampLinkPass struct{}

func (p *clampLinkPass) Name() string { return "clamp link capability" }
func (p *clampLinkPass) Apply(cs *pci.ConfigSpace, b *board.Board, om *overlay.Map, ctx *ScrubContext) {
	clampLinkCapability(cs, b, om, ctx.Caps)
}

type clampDeviceCapPass struct{}

func (p *clampDeviceCapPass) Name() string { return "clamp device capability" }
func (p *clampDeviceCapPass) Apply(cs *pci.ConfigSpace, b *board.Board, om *overlay.Map, ctx *ScrubContext) {
	clampDeviceCapability(cs, om, ctx.Caps)
}

type zeroVendorPass struct{}

func (p *zeroVendorPass) Name() string { return "zero vendor registers" }
func (p *zeroVendorPass) Apply(cs *pci.ConfigSpace, b *board.Board, om *overlay.Map, ctx *ScrubContext) {
	zeroVendorRegisters(cs, om, ctx.Caps)
}

type applyVendorQuirksPass struct{}

func (p *applyVendorQuirksPass) Name() string { return "vendor quirks" }
func (p *applyVendorQuirksPass) Apply(cs *pci.ConfigSpace, b *board.Board, om *overlay.Map, ctx *ScrubContext) {
	applyVendorQuirks(cs, om)
}

type pruneStdCapsPass struct{}

func (p *pruneStdCapsPass) Name() string { return "prune standard caps" }
func (p *pruneStdCapsPass) Apply(cs *pci.ConfigSpace, b *board.Board, om *overlay.Map, ctx *ScrubContext) {
	if pruned := PruneStandardCaps(cs, om); len(pruned) > 0 {
		for _, pr := range pruned {
			slog.Info("pruned standard cap", "cap", pr)
		}
	}
}

type validateCapChainPass struct{}

func (p *validateCapChainPass) Name() string { return "validate cap chain" }
func (p *validateCapChainPass) Apply(cs *pci.ConfigSpace, b *board.Board, om *overlay.Map, ctx *ScrubContext) {
	if err := ValidateCapChain(cs); err != nil {
		slog.Warn("capability chain issue", "error", err)
	}
}

type scrubASPMPass struct{}

func (p *scrubASPMPass) Name() string { return "scrub ASPM / L1PM" }
func (p *scrubASPMPass) Apply(cs *pci.ConfigSpace, b *board.Board, om *overlay.Map, ctx *ScrubContext) {
	for _, cap := range ctx.Caps {
		if cap.ID != pci.CapIDPCIExpress {
			continue
		}
		if cap.Offset+0x10+2 <= pci.ConfigSpaceLegacySize {
			linkCtl := cs.ReadU16(cap.Offset + 0x10)
			linkCtl &= 0xFFFC
			om.WriteU16(cap.Offset+0x10, linkCtl, "disable ASPM L0s/L1")
		}
		break
	}

	if cs.Size < pci.ConfigSpaceSize {
		return
	}
	for _, cap := range ctx.ExtCaps {
		if cap.ID != pci.ExtCapIDL1PMSubstates {
			continue
		}
		if cap.Offset+0x0C <= pci.ConfigSpaceSize {
			om.WriteU32(cap.Offset+0x08, 0, "clear L1PM Substates Control 1")
		}
		if cap.Offset+0x10 <= pci.ConfigSpaceSize {
			om.WriteU32(cap.Offset+0x0C, 0, "clear L1PM Substates Control 2")
		}
	}
}

type normalizeAERMasksPass struct{}

func (p *normalizeAERMasksPass) Name() string { return "normalize AER masks" }
func (p *normalizeAERMasksPass) Apply(cs *pci.ConfigSpace, b *board.Board, om *overlay.Map, ctx *ScrubContext) {
	if cs.Size < pci.ConfigSpaceSize {
		return
	}
	for _, cap := range ctx.ExtCaps {
		if cap.ID != pci.ExtCapIDAER {
			continue
		}
		if cap.Offset+0x0C <= pci.ConfigSpaceSize {
			om.WriteU32(cap.Offset+0x08, 0x00462030, "set AER uncorrectable mask (spec defaults)")
		}
		if cap.Offset+0x18 <= pci.ConfigSpaceSize {
			om.WriteU32(cap.Offset+0x14, 0x00002000, "set AER correctable mask (spec defaults)")
		}
		if cap.Offset+0x10 <= pci.ConfigSpaceSize {
			om.WriteU32(cap.Offset+0x0C, 0x00462011, "set AER uncorrectable severity (spec defaults)")
		}
	}
}
