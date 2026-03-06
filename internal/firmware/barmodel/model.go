// Package barmodel provides BAR register modeling for FPGA firmware generation.
package barmodel

import (
	"fmt"

	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/firmware/devclass"
	"github.com/sercanarga/pcileechgen/internal/util"
)

// BARRegister is one register in the BAR map.
type BARRegister struct {
	Offset uint32 // byte offset in BAR
	Width  int    // 1, 2, or 4 bytes
	Reset  uint32 // reset/initial value (from donor snapshot or spec default)
	RWMask uint32 // writable bits (1 = host can write, 0 = read-only)
	Name   string // human-readable register name
}

// BARModel is the full BAR register map wired into the SV output.
type BARModel struct {
	Size      int           // BAR size in bytes (typically 4096)
	Registers []BARRegister // ordered by Offset
}

// BuildBARModel returns a register map for the given device class.
// If a donor profile is available it takes priority over the spec tables.
// Falls back to spec-based register maps for known device classes.
// Returns nil for unknown classes without a profile.
func BuildBARModel(barData []byte, classCode uint32, profile *donor.BARProfile) *BARModel {
	// If we have profiling data, use the learned model
	if profile != nil && len(profile.Probes) > 0 {
		model := SynthesizeBARModel(profile, classCode)
		if model != nil {
			return model
		}
	}

	// Spec-based fallback for known device classes
	baseClass := (classCode >> 16) & 0xFF
	subClass := (classCode >> 8) & 0xFF
	progIF := classCode & 0xFF

	switch {
	case baseClass == 0x01 && subClass == 0x08 && progIF == 0x02:
		return buildNVMeBARModel(barData)
	case baseClass == 0x0C && subClass == 0x03 && progIF == 0x30:
		return buildXHCIBARModel(barData)
	case baseClass == 0x02:
		return buildEthernetBARModel(barData)
	case baseClass == 0x04 && subClass == 0x03:
		return buildAudioBARModel(barData)
	default:
		return nil
	}
}

// buildNVMeBARModel creates the NVMe controller BAR0 register map.
// Based on NVMe 1.4 spec: controller registers at BAR0 offset 0x00-0x4F.
func buildNVMeBARModel(barData []byte) *BARModel {
	regs := []BARRegister{
		// Controller Capabilities (CAP) — 64-bit, read-only
		{Offset: 0x00, Width: 4, Name: "CAP_LO", RWMask: 0x00000000},
		{Offset: 0x04, Width: 4, Name: "CAP_HI", RWMask: 0x00000000},
		// Version (VS) — read-only
		{Offset: 0x08, Width: 4, Name: "VS", RWMask: 0x00000000},
		// Interrupt Mask Set (INTMS)
		{Offset: 0x0C, Width: 4, Name: "INTMS", RWMask: 0xFFFFFFFF},
		// Interrupt Mask Clear (INTMC)
		{Offset: 0x10, Width: 4, Name: "INTMC", RWMask: 0xFFFFFFFF},
		// Controller Configuration (CC) — host writes EN, CSS, MPS, etc.
		{Offset: 0x14, Width: 4, Name: "CC", RWMask: 0x00FFFFF1},
		// Controller Status (CSTS) — read-only for host, FPGA sets RDY
		{Offset: 0x1C, Width: 4, Name: "CSTS", RWMask: 0x00000000},
		// NVM Subsystem Reset (NSSR) — write-only
		{Offset: 0x20, Width: 4, Name: "NSSR", RWMask: 0xFFFFFFFF},
		// Admin Queue Attributes (AQA)
		{Offset: 0x24, Width: 4, Name: "AQA", RWMask: 0x0FFF0FFF},
		// Admin Submission Queue Base Address (ASQ) — 64-bit
		{Offset: 0x28, Width: 4, Name: "ASQ_LO", RWMask: 0xFFFFF000},
		{Offset: 0x2C, Width: 4, Name: "ASQ_HI", RWMask: 0xFFFFFFFF},
		// Admin Completion Queue Base Address (ACQ) — 64-bit
		{Offset: 0x30, Width: 4, Name: "ACQ_LO", RWMask: 0xFFFFF000},
		{Offset: 0x34, Width: 4, Name: "ACQ_HI", RWMask: 0xFFFFFFFF},
	}

	populateResetValues(regs, barData)

	// Force CSTS.RDY=1 so stornvme.sys sees a ready controller
	for i := range regs {
		if regs[i].Offset == 0x1C {
			regs[i].Reset |= 0x00000001  // RDY bit
			regs[i].Reset &^= 0x0000000C // clear SHST (shutdown status)
			break
		}
	}

	return &BARModel{
		Size:      4096,
		Registers: regs,
	}
}

// buildXHCIBARModel creates the xHCI USB controller BAR0 register map.
// Based on xHCI 1.2 spec: capability + operational registers.
func buildXHCIBARModel(barData []byte) *BARModel {
	regs := []BARRegister{
		// -- Capability Registers (read-only) --
		// CAPLENGTH (byte) + HCIVERSION (word) packed in DWORD
		{Offset: 0x00, Width: 4, Name: "CAPLENGTH_HCIVERSION", RWMask: 0x00000000},
		// Structural Parameters 1
		{Offset: 0x04, Width: 4, Name: "HCSPARAMS1", RWMask: 0x00000000},
		// Structural Parameters 2
		{Offset: 0x08, Width: 4, Name: "HCSPARAMS2", RWMask: 0x00000000},
		// Structural Parameters 3
		{Offset: 0x0C, Width: 4, Name: "HCSPARAMS3", RWMask: 0x00000000},
		// Capability Parameters 1
		{Offset: 0x10, Width: 4, Name: "HCCPARAMS1", RWMask: 0x00000000},
		// Doorbell Offset
		{Offset: 0x14, Width: 4, Name: "DBOFF", RWMask: 0x00000000},
		// Runtime Register Space Offset
		{Offset: 0x18, Width: 4, Name: "RTSOFF", RWMask: 0x00000000},
		// Capability Parameters 2
		{Offset: 0x1C, Width: 4, Name: "HCCPARAMS2", RWMask: 0x00000000},

		// -- Operational Registers (at CAPLENGTH offset, typically 0x20) --
		// USBCMD
		{Offset: 0x20, Width: 4, Name: "USBCMD", RWMask: 0x00002F0E},
		// USBSTS (mostly RW1C)
		{Offset: 0x24, Width: 4, Name: "USBSTS", RWMask: 0x00000000},
		// Page Size (read-only)
		{Offset: 0x28, Width: 4, Name: "PAGESIZE", RWMask: 0x00000000},
		// Device Notification Control
		{Offset: 0x34, Width: 4, Name: "DNCTRL", RWMask: 0x0000FFFF},
		// Command Ring Control — 64-bit
		{Offset: 0x38, Width: 4, Name: "CRCR_LO", RWMask: 0xFFFFFFF7},
		{Offset: 0x3C, Width: 4, Name: "CRCR_HI", RWMask: 0xFFFFFFFF},
		// Device Context Base Address Array Pointer — 64-bit
		{Offset: 0x50, Width: 4, Name: "DCBAAP_LO", RWMask: 0xFFFFFFC0},
		{Offset: 0x54, Width: 4, Name: "DCBAAP_HI", RWMask: 0xFFFFFFFF},
		// Configure (CONFIG)
		{Offset: 0x58, Width: 4, Name: "CONFIG", RWMask: 0x000000FF},
	}

	populateResetValues(regs, barData)

	// Set USBCMD.R/S=1 and ensure USBSTS.HCH=0 so driver sees a running controller
	for i := range regs {
		switch regs[i].Offset {
		case 0x20: // USBCMD
			regs[i].Reset |= 0x00000001 // R/S bit
		case 0x24: // USBSTS
			regs[i].Reset &^= 0x00000001 // clear HCH (halted)
		}
	}

	return &BARModel{
		Size:      4096,
		Registers: regs,
	}
}

// populateResetValues fills Reset fields from donor BAR bytes.
func populateResetValues(regs []BARRegister, barData []byte) {
	if len(barData) == 0 {
		return
	}
	for i := range regs {
		off := int(regs[i].Offset)
		w := regs[i].Width
		if off+w > len(barData) {
			continue
		}
		switch w {
		case 4:
			regs[i].Reset = util.ReadLE32(barData, off)
		case 2:
			regs[i].Reset = uint32(barData[off]) | uint32(barData[off+1])<<8
		case 1:
			regs[i].Reset = uint32(barData[off])
		}
	}
}

// buildEthernetBARModel creates a register map for Intel-style Ethernet controllers.
// Based on Intel I210/I350 datasheet register layout.
func buildEthernetBARModel(barData []byte) *BARModel {
	regs := []BARRegister{
		{Offset: 0x00, Width: 4, Name: "CTRL", RWMask: 0xFFFFFFFF},
		{Offset: 0x08, Width: 4, Name: "STATUS", RWMask: 0x00000000},
		{Offset: 0x10, Width: 4, Name: "EECD", RWMask: 0x000000FF},
		{Offset: 0x14, Width: 4, Name: "EERD", RWMask: 0x0000FFFF},
		{Offset: 0x18, Width: 4, Name: "CTRL_EXT", RWMask: 0xFFFFFFFF},
		{Offset: 0x20, Width: 4, Name: "MDIC", RWMask: 0x0FFFFFFF},
		{Offset: 0x28, Width: 4, Name: "FCAL", RWMask: 0xFFFFFFFF},
		{Offset: 0x2C, Width: 4, Name: "FCAH", RWMask: 0x0000FFFF},
		{Offset: 0xC0, Width: 4, Name: "ICR", RWMask: 0x00000000},
		{Offset: 0xC8, Width: 4, Name: "ICS", RWMask: 0xFFFFFFFF},
		{Offset: 0xD0, Width: 4, Name: "IMS", RWMask: 0xFFFFFFFF},
		{Offset: 0xD8, Width: 4, Name: "IMC", RWMask: 0xFFFFFFFF},
		{Offset: 0x100, Width: 4, Name: "RCTL", RWMask: 0xFFFFFFFF},
		{Offset: 0x400, Width: 4, Name: "TCTL", RWMask: 0xFFFFFFFF},
		// MAC address — needs 32KB BAR to be reachable
		{Offset: 0x5400, Width: 4, Name: "RAL0", RWMask: 0xFFFFFFFF},
		{Offset: 0x5404, Width: 4, Name: "RAH0", RWMask: 0xFFFFFFFF},
	}

	populateResetValues(regs, barData)

	for i := range regs {
		switch regs[i].Offset {
		case 0x08:
			regs[i].Reset |= 0x00000082 // STATUS: link up + 1000Mb
		case 0x10:
			regs[i].Reset |= 0x00000300 // EECD: Auto-Read Done + EEPROM Present
		case 0x20:
			regs[i].Reset |= 0x10000000 // MDIC: Ready bit
		case 0x5400:
			if regs[i].Reset == 0 {
				regs[i].Reset = 0xADDE0200 // locally-administered MAC low
			}
		case 0x5404:
			if regs[i].Reset == 0 {
				regs[i].Reset = 0x8000EFBE // MAC high + AV (address valid)
			}
		}
	}

	return &BARModel{
		Size:      32768, // 32KB to reach RAL0/RAH0 at 0x5400
		Registers: regs,
	}
}

// buildAudioBARModel creates the HD Audio controller BAR0 register map.
// All registers are DWORD-aligned — sub-word HDA registers are packed into
// DWORD slots so the template's 32-bit write logic works correctly.
func buildAudioBARModel(barData []byte) *BARModel {
	regs := []BARRegister{
		// GCAP(16) + VMIN(8) + VMAJ(8) packed into one DWORD
		{Offset: 0x00, Width: 4, Name: "GCAP_VMIN_VMAJ", RWMask: 0x00000000},
		// GCTL — global control, CRST (bit 0) is the key writable bit
		{Offset: 0x08, Width: 4, Name: "GCTL", RWMask: 0x00000103},
		// WAKEEN(16) + STATESTS(16) packed into one DWORD
		{Offset: 0x0C, Width: 4, Name: "WAKEEN_STATESTS", RWMask: 0x7FFFFFFF},
		// INTCTL — interrupt control
		{Offset: 0x20, Width: 4, Name: "INTCTL", RWMask: 0xC00000FF},
		// INTSTS — interrupt status
		{Offset: 0x24, Width: 4, Name: "INTSTS", RWMask: 0x00000000},
		// CORB lower base address
		{Offset: 0x40, Width: 4, Name: "CORBLBASE", RWMask: 0xFFFFFF80},
		// CORB upper base address
		{Offset: 0x44, Width: 4, Name: "CORBUBASE", RWMask: 0xFFFFFFFF},
		// CORBWP(16) + CORBRP(16) packed
		{Offset: 0x48, Width: 4, Name: "CORBWP_CORBRP", RWMask: 0x80FF00FF},
		// CORBCTL(8) + CORBSTS(8) + CORBSIZE(8) packed
		{Offset: 0x4C, Width: 4, Name: "CORBCTL_STS_SIZE", RWMask: 0x00030300},
		// RIRB lower base address
		{Offset: 0x50, Width: 4, Name: "RIRBLBASE", RWMask: 0xFFFFFF80},
		// RIRB upper base address
		{Offset: 0x54, Width: 4, Name: "RIRBUBASE", RWMask: 0xFFFFFFFF},
		// RIRBWP(16) + RINTCNT(16)
		{Offset: 0x58, Width: 4, Name: "RIRBWP_RINTCNT", RWMask: 0x800000FF},
		// RIRBCTL(8) + RIRBSTS(8) + RIRBSIZE(8)
		{Offset: 0x5C, Width: 4, Name: "RIRBCTL_STS_SIZE", RWMask: 0x00070700},
	}

	populateResetValues(regs, barData)

	// set up sensible defaults for codec discovery
	for i := range regs {
		switch regs[i].Offset {
		case 0x00:
			if regs[i].Reset == 0 {
				regs[i].Reset = 0x01004401 // GCAP=4401h, VMIN=0, VMAJ=1
			}
		case 0x08:
			regs[i].Reset |= 0x00000001 // GCTL.CRST=1 (out of reset)
		case 0x0C:
			regs[i].Reset |= 0x00010000 // STATESTS: codec 0 present (upper 16 bits)
		case 0x4C:
			if regs[i].Reset == 0 {
				regs[i].Reset = 0x00420000 // CORBSIZE=0x42 (supports 256/16/2)
			}
		case 0x5C:
			if regs[i].Reset == 0 {
				regs[i].Reset = 0x00420000 // RIRBSIZE=0x42
			}
		}
	}

	return &BARModel{
		Size:      4096,
		Registers: regs,
	}
}

// SynthesizeBARModel turns probe data into a BARModel.
// Dead registers (zero value + zero RW mask) are dropped.
// RW1C bits are conservatively treated as read-only.
func SynthesizeBARModel(profile *donor.BARProfile, classCode uint32) *BARModel {
	if profile == nil || len(profile.Probes) == 0 {
		return nil
	}

	nameHints := classRegisterNames(classCode)

	var regs []BARRegister
	for _, probe := range profile.Probes {
		// Skip dead registers (no value, no writable bits)
		if probe.Original == 0 && probe.RWMask == 0 {
			continue
		}

		name := fmt.Sprintf("REG_0x%03X", probe.Offset)
		if hint, ok := nameHints[probe.Offset]; ok {
			name = hint
		}

		rwMask := probe.RWMask
		// RW1C → treat as read-only in SV (writing 1 clears, not normal RW)
		if probe.MaybeRW1C {
			rwMask = 0 // conservative: mark entire register as RO if any RW1C detected
		}

		regs = append(regs, BARRegister{
			Offset: probe.Offset,
			Width:  4,
			Reset:  probe.Original,
			RWMask: rwMask,
			Name:   name,
		})
	}

	if len(regs) == 0 {
		return nil
	}

	return &BARModel{
		Size:      profile.Size,
		Registers: regs,
	}
}

// classRegisterNames returns offset→name hints for known device classes.
// Derives names from devclass.ProfileForClass to avoid duplication.
func classRegisterNames(classCode uint32) map[uint32]string {
	profile := devclass.ProfileForClass(classCode)
	if profile == nil {
		return nil
	}
	names := make(map[uint32]string, len(profile.BARDefaults))
	for _, d := range profile.BARDefaults {
		names[d.Offset] = d.Name
	}
	return names
}
