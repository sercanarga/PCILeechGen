package firmware

import (
	"fmt"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

// ext caps the FPGA can't emulate
var unsafeExtCaps = map[uint16]string{
	pci.ExtCapIDSRIOV:         "SR-IOV",
	pci.ExtCapIDMRIOV:         "MR-IOV",
	pci.ExtCapIDResizableBAR:  "Resizable BAR",
	pci.ExtCapIDATS:           "ATS",
	pci.ExtCapIDPageRequest:   "Page Request",
	pci.ExtCapIDPASID:         "PASID",
	pci.ExtCapIDL1PMSubstates: "L1 PM Substates",
	pci.ExtCapIDDPC:           "DPC",
	pci.ExtCapIDPTM:           "PTM",
	pci.ExtCapIDSecondaryPCIe: "Secondary PCIe",
	pci.ExtCapIDMulticast:     "Multicast",
}

func IsUnsafeExtCap(id uint16) bool {
	_, ok := unsafeExtCaps[id]
	return ok
}

func UnsafeExtCapName(id uint16) string {
	if name, ok := unsafeExtCaps[id]; ok {
		return name
	}
	return ""
}

// vendorQuirk ties a vendor+class combo to a fixup.
// ClassCode 0 = match any class.
type vendorQuirk struct {
	VendorID  uint16
	ClassCode uint32
	Name      string
	Apply     func(cs *pci.ConfigSpace)
}

var vendorQuirks = []vendorQuirk{
	{
		VendorID:  0x1912,
		ClassCode: 0x0C0330,
		Name:      "Renesas xHCI FW status",
		Apply:     fixRenesasFirmwareStatus,
	},
	// add new vendors here
}

func applyVendorQuirks(cs *pci.ConfigSpace) {
	vid := cs.VendorID()
	cc := cs.ClassCode()
	for _, q := range vendorQuirks {
		if q.VendorID != vid {
			continue
		}
		if q.ClassCode != 0 && q.ClassCode != cc {
			continue
		}
		q.Apply(cs)
	}
}

// Renesas uPD720201/202: mark FW as loaded.
// Without this the driver starts a FW download handshake → Code 10.
func fixRenesasFirmwareStatus(cs *pci.ConfigSpace) {
	const (
		fwStatus      = 0xF4
		romStatus     = 0xF6
		fwSuccess     = 0x10 // bit 4
		fwLock        = 0x80 // bit 7
		romResultMask = 0x0070
	)

	cs.WriteU8(fwStatus, fwSuccess|fwLock)

	rs := cs.ReadU16(romStatus)
	rs = (rs &^ uint16(romResultMask)) | uint16(fwSuccess)
	cs.WriteU16(romStatus, rs)
}

// min sizes (bytes) for standard PCI caps
var capMinSize = map[uint8]int{
	pci.CapIDPowerManagement: 8,  // PM: 2 header + 2 PMC + 2 PMCSR + 2 data
	pci.CapIDMSI:             10, // MSI: varies, 10-24 depending on bits
	pci.CapIDMSIX:            12, // MSI-X: 2 header + 2 ctl + 4 table + 4 PBA
	pci.CapIDPCIExpress:      60, // PCIe: 0x3C typical for v2
	pci.CapIDVendorSpecific:  3,  // at least header + length byte
}

// capSize returns byte size for a cap ID. Unknown caps default to 8.
func capSize(id uint8) int {
	if s, ok := capMinSize[id]; ok {
		return s
	}
	return 8
}

// zeroVendorRegisters clears 0x40-0xFF bytes not covered by any PCI cap.
func zeroVendorRegisters(cs *pci.ConfigSpace) {
	covered := make([]bool, pci.ConfigSpaceLegacySize)
	for i := 0; i < 0x40; i++ {
		covered[i] = true
	}

	caps := pci.ParseCapabilities(cs)
	for _, cap := range caps {
		size := capSize(cap.ID)
		for i := cap.Offset; i < cap.Offset+size && i < pci.ConfigSpaceLegacySize; i++ {
			covered[i] = true
		}
	}

	for i := 0x40; i < pci.ConfigSpaceLegacySize; i++ {
		if !covered[i] {
			cs.WriteU8(i, 0x00)
		}
	}
}

const fpgaBRAMSize = 4096
const fpgaBAR0SizeMask = 0xFFFFF000    // 4 KB aligned
const fpgaMaxLinkSpeed = LinkSpeedGen2 // 7-series can't go faster

// ScrubConfigSpace cleans donor config space for COE generation.
func ScrubConfigSpace(cs *pci.ConfigSpace, b *board.Board) *pci.ConfigSpace {
	scrubbed := cs.Clone()

	scrubbed.WriteU8(0x0F, 0x00) // BIST
	scrubbed.WriteU8(0x3C, 0x00) // Interrupt Line
	scrubbed.WriteU8(0x0D, 0x00) // Latency Timer
	scrubbed.WriteU8(0x0C, 0x00) // Cache Line Size

	cmd := scrubbed.Command() & 0x0547
	scrubbed.WriteU16(0x04, cmd)

	status := scrubbed.Status() & 0x06F0
	scrubbed.WriteU16(0x06, status)

	caps := pci.ParseCapabilities(scrubbed)
	for _, cap := range caps {
		if cap.ID == pci.CapIDPCIExpress && cap.Offset+10 < pci.ConfigSpaceLegacySize {
			scrubbed.WriteU16(cap.Offset+10, 0x0000) // device status
			if cap.Offset+18 < pci.ConfigSpaceLegacySize {
				lstatus := scrubbed.ReadU16(cap.Offset + 18)
				lstatus &= 0x3FFF
				scrubbed.WriteU16(cap.Offset+18, lstatus)
			}
		}

		if cap.ID == pci.CapIDPowerManagement && cap.Offset+4 < pci.ConfigSpaceLegacySize {
			pmcsr := scrubbed.ReadU16(cap.Offset + 4)
			pmcsr &= 0xFFFC // D0
			pmcsr &= 0x7FFF // clear PME_Status
			pmcsr |= 0x0008 // NoSoftReset
			scrubbed.WriteU16(cap.Offset+4, pmcsr)
		}
	}

	if scrubbed.Size >= pci.ConfigSpaceSize {
		extCaps := pci.ParseExtCapabilities(scrubbed)
		for _, cap := range extCaps {
			if cap.ID == pci.ExtCapIDAER {
				if cap.Offset+4+4 <= pci.ConfigSpaceSize {
					scrubbed.WriteU32(cap.Offset+4, 0) // uncorrectable error status
				}
				if cap.Offset+16+4 <= pci.ConfigSpaceSize {
					scrubbed.WriteU32(cap.Offset+16, 0) // correctable error status
				}
				if cap.Offset+28+4 <= pci.ConfigSpaceSize {
					scrubbed.WriteU32(cap.Offset+28, 0) // root error status
				}
			}
		}
		FilterExtCapabilities(scrubbed)
	}

	clampBARsToFPGA(scrubbed)
	disableMSIXIfOutOfBRAM(scrubbed)
	clampLinkCapability(scrubbed, b)
	clampDeviceCapability(scrubbed)

	// wipe vendor regs, then apply known quirks on top
	zeroVendorRegisters(scrubbed)
	applyVendorQuirks(scrubbed)

	return scrubbed
}

// clampBARsToFPGA shrinks all memory BARs to 4 KB.
func clampBARsToFPGA(cs *pci.ConfigSpace) {
	for i := 0; i < 6; i++ {
		barOffset := 0x10 + (i * 4)
		barVal := cs.BAR(i)
		if barVal == 0 {
			continue
		}

		if barVal&0x01 != 0 {
			continue
		}

		is64bit := (barVal & 0x06) == 0x04
		newBar := fpgaBAR0SizeMask | (barVal & 0x0F)
		cs.WriteU32(barOffset, newBar)

		if is64bit && i < 5 {
			cs.WriteU32(barOffset+4, 0x00000000)
			i++ // skip upper half
		}
	}
}

// disableMSIXIfOutOfBRAM kills MSI-X if table/PBA sits past 4 KB.
func disableMSIXIfOutOfBRAM(cs *pci.ConfigSpace) {
	caps := pci.ParseCapabilities(cs)
	for _, cap := range caps {
		if cap.ID != pci.CapIDMSIX {
			continue
		}
		if cap.Offset+8 > pci.ConfigSpaceLegacySize {
			continue
		}

		tableOff := int(cs.ReadU32(cap.Offset+4) &^ 0x07)
		pbaOff := int(cs.ReadU32(cap.Offset+8) &^ 0x07)

		if tableOff >= fpgaBRAMSize || pbaOff >= fpgaBRAMSize {
			msgCtl := cs.ReadU16(cap.Offset + 2)
			msgCtl &= 0x3FFF
			cs.WriteU16(cap.Offset+2, msgCtl)
		}
		break
	}
}

// clampLinkCapability caps link speed/width to match the FPGA + board.
func clampLinkCapability(cs *pci.ConfigSpace, b *board.Board) {
	caps := pci.ParseCapabilities(cs)
	for _, cap := range caps {
		if cap.ID != pci.CapIDPCIExpress {
			continue
		}

		maxSpeed := uint8(fpgaMaxLinkSpeed)
		maxWidth := uint8(0)
		if b != nil && b.PCIeLanes > 0 {
			maxWidth = uint8(b.PCIeLanes)
		}

		// Link Capabilities (cap+0x0C)
		if cap.Offset+0x0C+4 <= pci.ConfigSpaceLegacySize {
			linkCap := cs.ReadU32(cap.Offset + 0x0C)
			if speed := uint8(linkCap & 0x0F); speed > maxSpeed {
				linkCap = (linkCap & 0xFFFFFFF0) | uint32(maxSpeed)
			}
			if maxWidth > 0 {
				if width := uint8((linkCap >> 4) & 0x3F); width > maxWidth {
					linkCap = (linkCap & 0xFFFFFC0F) | (uint32(maxWidth) << 4)
				}
			}
			cs.WriteU32(cap.Offset+0x0C, linkCap)
		}

		// Link Status (cap+0x12)
		if cap.Offset+0x12+2 <= pci.ConfigSpaceLegacySize {
			ls := cs.ReadU16(cap.Offset + 0x12)
			if speed := uint8(ls & 0x0F); speed > maxSpeed {
				ls = (ls & 0xFFF0) | uint16(maxSpeed)
			}
			if maxWidth > 0 {
				if width := uint8((ls >> 4) & 0x3F); width > maxWidth {
					ls = (ls & 0xFC0F) | (uint16(maxWidth) << 4)
				}
			}
			cs.WriteU16(cap.Offset+0x12, ls)
		}

		// Link Control 2 (cap+0x30) — target speed
		if cap.Offset+0x30+2 <= pci.ConfigSpaceLegacySize {
			lc2 := cs.ReadU16(cap.Offset + 0x30)
			if speed := uint8(lc2 & 0x0F); speed > maxSpeed {
				lc2 = (lc2 & 0xFFF0) | uint16(maxSpeed)
			}
			cs.WriteU16(cap.Offset+0x30, lc2)
		}

		// Link Capabilities 2 (cap+0x2C) — supported speed vector
		if cap.Offset+0x2C+4 <= pci.ConfigSpaceLegacySize {
			lc2 := cs.ReadU32(cap.Offset + 0x2C)
			if lc2 != 0 {
				var vec uint32
				for s := uint8(1); s <= maxSpeed; s++ {
					vec |= 1 << s
				}
				cs.WriteU32(cap.Offset+0x2C, (lc2&0xFFFFFF01)|vec)
			}
		}

		break
	}
}

// clampDeviceCapability limits MPS, phantoms, ext tags to FPGA limits.
func clampDeviceCapability(cs *pci.ConfigSpace) {
	caps := pci.ParseCapabilities(cs)
	for _, cap := range caps {
		if cap.ID != pci.CapIDPCIExpress {
			continue
		}

		// DevCap (cap+0x04)
		if cap.Offset+0x04+4 <= pci.ConfigSpaceLegacySize {
			devCap := cs.ReadU32(cap.Offset + 0x04)
			devCap &= ^uint32(0x07) // MPS → 128B
			devCap &= ^uint32(0x18) // phantom functions off
			devCap &= ^uint32(0x20) // extended tag off
			cs.WriteU32(cap.Offset+0x04, devCap)
		}

		// DevCtl (cap+0x08)
		if cap.Offset+0x08+2 <= pci.ConfigSpaceLegacySize {
			devCtl := cs.ReadU16(cap.Offset + 0x08)
			devCtl &= ^uint16(0x00E0) // MPS → 128B
			devCtl &= ^uint16(0x0100) // ext tag off
			devCtl &= ^uint16(0x0200) // phantom off
			if mrrs := (devCtl >> 12) & 0x07; mrrs > 2 {
				devCtl = (devCtl & 0x8FFF) | (2 << 12)
			}
			cs.WriteU16(cap.Offset+0x08, devCtl)
		}

		// DevCap2 (cap+0x24)
		if cap.Offset+0x24+4 <= pci.ConfigSpaceLegacySize {
			devCap2 := cs.ReadU32(cap.Offset + 0x24)
			devCap2 &= ^uint32(1 << 16) // 10-bit tag completer off
			devCap2 &= ^uint32(1 << 17) // 10-bit tag requester off
			cs.WriteU32(cap.Offset+0x24, devCap2)
		}

		break
	}
}

// FilterExtCapabilities removes unsupported ext caps and relinks the chain.
func FilterExtCapabilities(cs *pci.ConfigSpace) []string {
	var removed []string

	type capEntry struct {
		offset     int
		id         uint16
		version    uint8
		nextOffset int
		size       int
	}

	var entries []capEntry
	visited := make(map[int]bool)
	offset := 0x100

	for offset >= 0x100 && offset < pci.ConfigSpaceSize && !visited[offset] {
		visited[offset] = true

		header := cs.ReadU32(offset)
		if header == 0 || header == 0xFFFFFFFF {
			break
		}

		capID := uint16(header & 0xFFFF)
		capVer := uint8((header >> 16) & 0xF)
		nextOff := int((header >> 20) & 0xFFC)

		size := 4
		if nextOff > offset {
			size = nextOff - offset
		} else if nextOff == 0 {
			size = pci.ConfigSpaceSize - offset
		}

		entries = append(entries, capEntry{
			offset:     offset,
			id:         capID,
			version:    capVer,
			nextOffset: nextOff,
			size:       size,
		})

		if nextOff == 0 {
			break
		}
		offset = nextOff
	}

	if len(entries) == 0 {
		return nil
	}

	removeSet := make(map[int]bool)
	for i, e := range entries {
		if IsUnsafeExtCap(e.id) {
			removeSet[i] = true
			name := UnsafeExtCapName(e.id)
			removed = append(removed, fmt.Sprintf("%s (0x%04x) at offset 0x%03x", name, e.id, e.offset))
		}
	}

	if len(removeSet) == 0 {
		return nil
	}

	firstSurvivor := -1
	for i := range entries {
		if !removeSet[i] {
			firstSurvivor = i
			break
		}
	}

	needsRelocate := removeSet[0] && firstSurvivor > 0

	// wipe removed regions
	for i := range entries {
		if !removeSet[i] {
			continue
		}
		e := entries[i]
		for b := 0; b < e.size && e.offset+b < pci.ConfigSpaceSize; b++ {
			cs.WriteU8(e.offset+b, 0x00)
		}
	}

	if firstSurvivor < 0 {
		cs.WriteU32(0x100, 0x00000000) // all gone
		return removed
	}

	if needsRelocate {
		surv := entries[firstSurvivor]

		for b := 0; b < surv.size && b < surv.offset && surv.offset+b < pci.ConfigSpaceSize; b++ {
			cs.WriteU8(0x100+b, cs.ReadU8(surv.offset+b))
		}
		for b := 0; b < surv.size && surv.offset+b < pci.ConfigSpaceSize; b++ {
			cs.WriteU8(surv.offset+b, 0x00)
		}

		newNext := 0
		for j := firstSurvivor + 1; j < len(entries); j++ {
			if !removeSet[j] {
				newNext = entries[j].offset
				break
			}
		}

		hdr := cs.ReadU32(0x100)
		hdr = (hdr & 0x000FFFFF) | (uint32(newNext) << 20)
		cs.WriteU32(0x100, hdr)

		entries[firstSurvivor].offset = 0x100
	}

	var survivors []int
	for i := range entries {
		if !removeSet[i] {
			survivors = append(survivors, i)
		}
	}

	for si := 0; si < len(survivors); si++ {
		idx := survivors[si]
		e := entries[idx]

		newNext := 0
		if si+1 < len(survivors) {
			newNext = entries[survivors[si+1]].offset
		}

		hdr := cs.ReadU32(e.offset)
		hdr = (hdr & 0x000FFFFF) | (uint32(newNext) << 20)
		cs.WriteU32(e.offset, hdr)
	}

	return removed
}
