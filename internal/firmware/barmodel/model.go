// Package barmodel builds BAR register maps for SV codegen.
package barmodel

import (
	"fmt"
	"log/slog"

	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/firmware/devclass"
	"github.com/sercanarga/pcileechgen/internal/util"
)

// BARRegister describes a single register inside a BAR.
type BARRegister struct {
	Offset uint32 // byte offset in BAR
	Width  int    // 1, 2, or 4 bytes
	Reset  uint32 // reset/initial value (from donor snapshot or spec default)
	RWMask uint32 // writable bits (1 = host can write, 0 = read-only)
	Name   string // human-readable register name
}

// BARModel is the complete register map that ends up in the SV template.
type BARModel struct {
	Size      int           // BAR size in bytes (typically 4096)
	Registers []BARRegister // ordered by Offset
}

// BuildBARModel returns a register map from probe data or spec tables.
// Returns nil for unknown device classes.
func BuildBARModel(barData []byte, classCode uint32, profile *donor.BARProfile) *BARModel {
	// Use probe data when available, but bail if VFIO reported
	// everything as writable (breaks CC→CSTS handshake etc).
	if profile != nil && len(profile.Probes) > 0 {
		if !isProbeDataReliable(profile) {
			slog.Warn("BAR probe data unreliable (all registers report fully writable), falling back to spec-based model",
				"probes", len(profile.Probes))
		} else {
			model := SynthesizeBARModel(profile, classCode)
			if model != nil {
				return model
			}
		}
	}

	// fall back to hardcoded spec tables
	baseClass := (classCode >> 16) & 0xFF
	subClass := (classCode >> 8) & 0xFF
	progIF := classCode & 0xFF

	var model *BARModel
	switch {
	case baseClass == 0x01 && subClass == 0x08 && progIF == 0x02:
		model = buildNVMeBARModel(barData)
	case baseClass == 0x0C && subClass == 0x03 && progIF == 0x30:
		model = buildXHCIBARModel(barData)
	case baseClass == 0x02:
		model = buildEthernetBARModel(barData)
	case baseClass == 0x04 && subClass == 0x03:
		model = buildAudioBARModel(barData)
	}

	if model != nil {
		validateModel(model)
	}
	return model
}

// validateModel catches misaligned or duplicate offsets at build time.
// Non-aligned offsets produce broken Verilog case labels → silent Code 10.
func validateModel(m *BARModel) {
	seen := make(map[uint32]string, len(m.Registers))
	for _, r := range m.Registers {
		if r.Offset%4 != 0 {
			panic(fmt.Sprintf("barmodel: register %s at offset 0x%X is not DWORD-aligned", r.Name, r.Offset))
		}
		if prev, ok := seen[r.Offset]; ok {
			panic(fmt.Sprintf("barmodel: registers %s and %s share offset 0x%X — SV case conflict", prev, r.Name, r.Offset))
		}
		seen[r.Offset] = r.Name
	}
}

// NVMe BAR0 (spec 1.4, offsets 0x00–0x34).
func buildNVMeBARModel(barData []byte) *BARModel {
	regs := []BARRegister{
		// CAP — 64-bit RO
		{Offset: 0x00, Width: 4, Name: "CAP_LO", RWMask: 0x00000000},
		{Offset: 0x04, Width: 4, Name: "CAP_HI", RWMask: 0x00000000},
		// VS
		{Offset: 0x08, Width: 4, Name: "VS", RWMask: 0x00000000},
		// INTMS / INTMC
		{Offset: 0x0C, Width: 4, Name: "INTMS", RWMask: 0xFFFFFFFF},

		{Offset: 0x10, Width: 4, Name: "INTMC", RWMask: 0xFFFFFFFF},
		// CC — driver writes EN, FSM watches
		{Offset: 0x14, Width: 4, Name: "CC", RWMask: 0x00FFFFF1},
		// CSTS — RO, FSM drives RDY
		{Offset: 0x1C, Width: 4, Name: "CSTS", RWMask: 0x00000000},
		// NSSR
		{Offset: 0x20, Width: 4, Name: "NSSR", RWMask: 0xFFFFFFFF},
		// AQA
		{Offset: 0x24, Width: 4, Name: "AQA", RWMask: 0x0FFF0FFF},
		// ASQ — 64-bit
		{Offset: 0x28, Width: 4, Name: "ASQ_LO", RWMask: 0xFFFFF000},
		{Offset: 0x2C, Width: 4, Name: "ASQ_HI", RWMask: 0xFFFFFFFF},
		// ACQ — 64-bit
		{Offset: 0x30, Width: 4, Name: "ACQ_LO", RWMask: 0xFFFFF000},
		{Offset: 0x34, Width: 4, Name: "ACQ_HI", RWMask: 0xFFFFFFFF},
	}

	populateResetValues(regs, barData)

	// stornvme needs RDY=1 at boot
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

// xHCI BAR0 (spec 1.2).
func buildXHCIBARModel(barData []byte) *BARModel {
	regs := []BARRegister{
		// capability regs (RO)
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

		// operational regs
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

	// driver expects a running controller on first probe
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

// populateResetValues fills in reset values from donor BAR memory.
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

// RTL8125 register map (r8169 driver offsets).
func buildEthernetBARModel(barData []byte) *BARModel {
	regs := []BARRegister{
		{Offset: 0x00, Width: 4, Name: "MAC0_3", RWMask: 0xFFFFFFFF},
		{Offset: 0x04, Width: 4, Name: "MAC4_5", RWMask: 0xFFFFFFFF},
		{Offset: 0x34, Width: 4, Name: "CHIPCMD_DW", RWMask: 0x00000000}, // byte 0x37 = ChipCmd
		{Offset: 0x3C, Width: 4, Name: "INTRMASK", RWMask: 0xFFFFFFFF},
		{Offset: 0x40, Width: 4, Name: "TXCONFIG", RWMask: 0x00FF0000},
		{Offset: 0x44, Width: 4, Name: "RXCONFIG", RWMask: 0xFFFF7FFF},
		{Offset: 0x48, Width: 4, Name: "TIMER", RWMask: 0xFFFFFFFF},
		{Offset: 0x50, Width: 4, Name: "RXMAXSIZE", RWMask: 0x00003FFF},
		{Offset: 0x58, Width: 4, Name: "CPLUSCMD", RWMask: 0x0000FFFF},
		{Offset: 0x6C, Width: 4, Name: "PHYSTATUS", RWMask: 0x00000000}, // RO
		{Offset: 0xDC, Width: 4, Name: "PHYAR", RWMask: 0xFFFFFFFF},     // bit31 = ready
		{Offset: 0xE0, Width: 4, Name: "ERIAR", RWMask: 0xFFFFFFFF},     // bit31 = done
		{Offset: 0xFC, Width: 4, Name: "RXMISSED", RWMask: 0x00000000},  // RO
	}

	populateResetValues(regs, barData)

	for i := range regs {
		switch regs[i].Offset {
		case 0x00:
			if regs[i].Reset == 0 {
				regs[i].Reset = 0xBEADDE02
			}
		case 0x04:
			if regs[i].Reset == 0 {
				regs[i].Reset = 0x000000EF
			}
		case 0x34:
			regs[i].Reset |= 0x0C000000 // RxEn | TxEn
		case 0x40:
			regs[i].Reset |= 0x2F000000 // 8125B revision
		case 0x44:
			regs[i].Reset |= 0x00000E00
		case 0x50:
			regs[i].Reset |= 0x00003FFF
		case 0x58:
			regs[i].Reset |= 0x00002060
		case 0x6C:
			regs[i].Reset |= 0x00003010 // link + 2.5G + FDX
		case 0xDC:
			regs[i].Reset |= 0x80000000
		case 0xE0:
			regs[i].Reset |= 0x80000000
		}
	}

	return &BARModel{
		Size:      4096,
		Registers: regs,
	}
}

// HD Audio BAR0. Sub-word regs packed into DWORDs for SV template.
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

	// defaults for codec discovery
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

// SynthesizeBARModel builds a model from probe data.
// Drops dead regs and treats RW1C as RO.
func SynthesizeBARModel(profile *donor.BARProfile, classCode uint32) *BARModel {
	if profile == nil || len(profile.Probes) == 0 {
		return nil
	}

	nameHints := classRegisterNames(classCode)

	var regs []BARRegister
	for _, probe := range profile.Probes {
		// skip dead regs
		if probe.Original == 0 && probe.RWMask == 0 {
			continue
		}

		name := fmt.Sprintf("REG_0x%03X", probe.Offset)
		if hint, ok := nameHints[probe.Offset]; ok {
			name = hint
		}

		rwMask := probe.RWMask
		if probe.MaybeRW1C {
			rwMask = 0 // RW1C → force RO
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

// isProbeDataReliable rejects VFIO dumps where 90%+ of regs report
// fully writable — usually means the probe couldn't actually write.
func isProbeDataReliable(profile *donor.BARProfile) bool {
	var nonZero, allRW int
	for _, p := range profile.Probes {
		if p.Original == 0 && p.RWMask == 0 {
			continue
		}
		nonZero++
		if p.RWMask == 0xFFFFFFFF {
			allRW++
		}
	}
	if nonZero < 4 {
		return true // too few probes to judge
	}
	return allRW*10 < nonZero*9 // allRW < 90%
}

// classRegisterNames returns offset→name hints from the device profile.
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
