package scrub

import (
	"encoding/binary"
	"fmt"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/firmware"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

// Severity ranks a detectability finding by how readily a detector can use it.
type Severity int

const (
	SevLow    Severity = iota // valid for an FPGA card; weak or speculative tell
	SevMedium                 // detectable by comparing against a donor reference
	SevHigh                   // strong, often unfixable (board physics)
)

func (s Severity) String() string {
	switch s {
	case SevHigh:
		return "HIGH"
	case SevMedium:
		return "MEDIUM"
	default:
		return "LOW"
	}
}

// Finding is one residual way the emulated device can be told apart from the donor.
type Finding struct {
	Severity Severity
	Title    string
	Detail   string // what the tell is and how it is observed
	Hint     string // remediation or donor/board-selection guidance
}

// Audit reports the residual detectability tells that remain after scrubbing
// donor for emulation on board b. It does not mutate donor. The findings cover
// what enumeration and a light driver probe can still distinguish; deeper
// device-function and timing tells are out of scope here.
func Audit(donor *pci.ConfigSpace, b *board.Board) []Finding {
	if donor == nil {
		return nil
	}

	caps := pci.ParseCapabilities(donor)
	extCaps := pci.ParseExtCapabilities(donor)
	ids := firmware.ExtractDeviceIDs(donor, extCaps)

	var f []Finding

	// 1. Link width/speed vs board physical capability. Unfixable in firmware:
	//    LnkCap/LnkSta report the board's real trained link, not the donor's.
	if ids.HasPCIeCap && b != nil {
		if b.PCIeLanes > 0 && int(ids.LinkWidth) > b.PCIeLanes {
			f = append(f, Finding{
				Severity: SevHigh,
				Title:    "Link width mismatch",
				Detail: fmt.Sprintf("donor is x%d but board is x%d; LnkCap/LnkSta read x%d (board physical limit)",
					ids.LinkWidth, b.PCIeLanes, b.PCIeLanes),
				Hint: "pick a board with >= donor lanes, or a donor whose native width <= board",
			})
		}
		if maxSpeed := b.MaxLinkSpeedOrDefault(); ids.LinkSpeed > maxSpeed {
			f = append(f, Finding{
				Severity: SevHigh,
				Title:    "Link speed mismatch",
				Detail: fmt.Sprintf("donor is %s but board tops out at %s; clamped value is visible in LnkCap",
					firmware.LinkSpeedName(ids.LinkSpeed), firmware.LinkSpeedName(maxSpeed)),
				Hint: "choose a donor whose max link speed <= board, so no clamp is needed",
			})
		}
	}

	// 2. Standard capabilities the donor exposes but scrub must strip: each
	//    leaves a gap a VID/DID reference scan can XOR against the real donor.
	for _, c := range caps {
		if name, bad := unsafeStandardCaps[c.ID]; bad {
			f = append(f, Finding{
				Severity: SevMedium,
				Title:    "Capability stripped: " + name,
				Detail: fmt.Sprintf("donor exposes %s (0x%02X) but the FPGA can't emulate it; removed from the cap chain",
					name, c.ID),
				Hint: "prefer a donor of the same class that does not expose this capability",
			})
		}
	}

	// 3. Extended capabilities the donor exposes but scrub strips.
	for _, c := range extCaps {
		if IsUnsafeExtCap(c.ID) {
			name := UnsafeExtCapName(c.ID)
			f = append(f, Finding{
				Severity: SevMedium,
				Title:    "Extended capability stripped: " + name,
				Detail: fmt.Sprintf("donor exposes %s (0x%04X) but it requires hardware the FPGA lacks; removed (cap-chain XOR detects this)",
					name, c.ID),
				Hint: "pick a donor without " + name + ", or accept the chain gap",
			})
		}
	}

	// 3b. Expansion ROM the donor implements but the FPGA never serves: the
	//     ROM BAR (0x30) is preserved in config space, but reads to the ROM
	//     window return nothing, and option ROM serving is not implemented.
	if rom := donor.ReadU32(0x30); rom&0xFFFFF801 != 0 {
		f = append(f, Finding{
			Severity: SevMedium,
			Title:    "Expansion ROM present",
			Detail:   "donor implements an Expansion ROM BAR (0x30); reads to the ROM window must return the donor image or the absence is a tell",
			Hint:     "build with --option-rom to capture + serve it via BAR6 (enables the IP expansion-ROM BAR)",
		})
	}

	// 4. No DSN to clone: drivers/anti-cheat that key on a serial see a blank.
	if !ids.HasDSN {
		f = append(f, Finding{
			Severity: SevMedium,
			Title:    "No Device Serial Number",
			Detail:   "donor has no DSN capability, so the emulation cannot present a serial number",
			Hint:     "choose a donor that exposes a DSN extended capability",
		})
	}

	// 5. Advanced PCIe features the donor advertises but scrub clears for
	//    stability. The uniform "all advanced features off" profile is itself
	//    a soft signature; these are fixable only with deeper RTL support.
	for _, c := range caps {
		if c.ID != pci.CapIDPCIExpress || len(c.Data) < 16 {
			continue
		}
		devCap := binary.LittleEndian.Uint32(c.Data[4:8])
		if devCap&(1<<28) != 0 {
			f = append(f, Finding{
				Severity: SevMedium,
				Title:    "FLR advertised, will be cleared",
				Detail:   "donor advertises Function Level Reset; scrub clears it (FPGA can't honor FLR). Most modern devices are FLR-capable, so absence stands out",
				Hint:     "requires RTL FLR handling to keep; otherwise unavoidable",
			})
		}
		linkCap := binary.LittleEndian.Uint32(c.Data[12:16])
		if linkCap&0x0C00 != 0 {
			f = append(f, Finding{
				Severity: SevLow,
				Title:    "ASPM advertised, will be cleared",
				Detail:   "donor advertises ASPM L0s/L1; scrub clears it (FPGA has no link power management)",
				Hint:     "common on FPGA cards; low signal on its own",
			})
		}
		break
	}
	for _, c := range extCaps {
		if c.ID == pci.ExtCapIDL1PMSubstates {
			f = append(f, Finding{
				Severity: SevLow,
				Title:    "L1 PM Substates neutered",
				Detail:   "donor exposes L1 PM Substates; scrub keeps the cap in the chain but zeros its contents",
				Hint:     "low signal; contents-zeroed cap is mildly anomalous",
			})
			break
		}
	}

	// 6. MSI-X relocation into BRAM at a generated offset.
	for _, c := range caps {
		if c.ID == pci.CapIDMSIX {
			f = append(f, Finding{
				Severity: SevLow,
				Title:    "MSI-X table relocated to BRAM",
				Detail:   "MSI-X table/PBA move into FPGA BRAM at a generated offset, deterministic per class/size",
				Hint:     "could be randomized per build (needs BAR-controller RTL coordination)",
			})
			break
		}
	}

	return f
}

// AuditSummary counts findings by severity.
func AuditSummary(findings []Finding) (high, medium, low int) {
	for _, f := range findings {
		switch f.Severity {
		case SevHigh:
			high++
		case SevMedium:
			medium++
		default:
			low++
		}
	}
	return
}
