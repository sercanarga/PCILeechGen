// Package barmodel builds BAR register maps for SV codegen.
package barmodel

import (
	"fmt"
	"log/slog"
	"sort"

	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/firmware/devclass"
	"github.com/sercanarga/pcileechgen/internal/pci"
	"github.com/sercanarga/pcileechgen/internal/util"
)

// BARRegister describes a single register inside a BAR.
type BARRegister struct {
	Offset      uint32 // byte offset in BAR
	Width       int    // 1, 2, or 4 bytes
	Reset       uint32 // reset/initial value (from donor snapshot or spec default)
	RWMask      uint32 // writable bits (1 = host can write, 0 = read-only)
	W1CMask     uint32 // write-1-to-clear bits (1 = host writes 1 to clear, 0 = untouched)
	Name        string // human-readable register name
	IsRW1C      bool   // true if any W1C bits are present
	IsFSMDriven bool   // true if driven by a dedicated FSM always block (excluded from generic reset/write)
}

// BARModel is the complete register map that ends up in the SV template.
type BARModel struct {
	BIR           int
	Size          int
	Aperture      uint64
	Type          string
	Prefetchable  bool
	Is64Bit       bool
	UpperBIR      int
	ClassSpecific bool
	Registers     []BARRegister // ordered by Offset
}

func BuildBARModels(
	bars []pci.BAR,
	contents map[int][]byte,
	profiles map[int]*donor.BARProfile,
	classCode uint32,
	preferredBIR int,
) ([]*BARModel, error) {
	byBIR := make(map[int]pci.BAR, len(bars))
	for _, bar := range bars {
		if bar.Index >= 0 && bar.Index < 6 {
			byBIR[bar.Index] = bar
		}
	}

	for bir, data := range contents {
		if bir < 0 || bir >= 6 {
			continue
		}
		if _, ok := byBIR[bir]; !ok && len(data) > 0 {
			byBIR[bir] = pci.BAR{Index: bir, Size: uint64(len(data)), Type: pci.BARTypeMem32}
		}
	}
	for bir, profile := range profiles {
		if bir < 0 || bir >= 6 || profile == nil {
			continue
		}
		if _, ok := byBIR[bir]; !ok && profile.Size > 0 {
			byBIR[bir] = pci.BAR{Index: bir, Size: uint64(profile.Size), Type: pci.BARTypeMem32}
		}
	}

	birs := make([]int, 0, len(byBIR))
	for bir := range byBIR {
		birs = append(birs, bir)
	}
	sort.Ints(birs)

	models := make([]*BARModel, 0, len(birs))
	consumed := make(map[int]bool)
	for _, bir := range birs {
		if consumed[bir] {
			continue
		}
		bar := byBIR[bir]
		legacyMemory := bar.Type == "" && (bar.Size > 0 || len(contents[bir]) > 0 || profiles[bir] != nil)
		if (!bar.IsMemory() && !legacyMemory) || bar.Size == 0 && len(contents[bir]) == 0 &&
			(profiles[bir] == nil || profiles[bir].Size == 0) {
			continue
		}

		if bar.Is64Bit || bar.Type == pci.BARTypeMem64 {
			if bir == 5 {
				return nil, fmt.Errorf("64-bit BAR5 has no upper BIR")
			}
			upper := byBIR[bir+1]
			upperPresent := upper.Size > 0 && !upper.IsDisabled()
			if upperPresent || len(contents[bir+1]) > 0 ||
				profiles[bir+1] != nil && profiles[bir+1].Size > 0 {
				return nil, fmt.Errorf("64-bit BAR%d consumes occupied BAR%d", bir, bir+1)
			}
		}
		upperBIR := -1
		if bar.Is64Bit || bar.Type == pci.BARTypeMem64 {
			upperBIR = bir + 1
			if upperBIR < 6 {
				consumed[upperBIR] = true
			}
		}

		data := contents[bir]
		profile := profiles[bir]
		var model *BARModel
		classSpecific := bir == preferredBIR
		if !classSpecific && preferredBIR >= 0 && preferredBIR < 6 {
			if _, preferredExists := byBIR[preferredBIR]; !preferredExists {
				classSpecific = bir == birs[0]
			}
		}
		if classSpecific {
			model = BuildBARModel(data, classCode, profile)
		} else if profile != nil && len(profile.Probes) > 0 && isProbeDataReliable(profile) {
			model = SynthesizeBARModel(profile, 0)
		}
		if model == nil {
			model = &BARModel{}
		}

		aperture := bar.Size
		if uint64(len(data)) > aperture {
			aperture = uint64(len(data))
		}
		if profile != nil && uint64(profile.Size) > aperture {
			aperture = uint64(profile.Size)
		}
		model.BIR = bir
		model.Aperture = aperture
		model.Size = int(aperture)
		model.Type = bar.Type
		if model.Type == "" {
			model.Type = pci.BARTypeMem32
		}
		model.Prefetchable = bar.Prefetchable
		model.Is64Bit = bar.Is64Bit || bar.Type == pci.BARTypeMem64
		model.UpperBIR = upperBIR
		model.ClassSpecific = classSpecific && len(model.Registers) > 0
		validateModel(model)
		models = append(models, model)
	}
	return models, nil
}

func ModelForBIR(models []*BARModel, bir int) *BARModel {
	for _, model := range models {
		if model != nil && model.BIR == bir {
			return model
		}
	}
	return nil
}

// BuildBARModel returns a register map from probe data or spec tables.
// Returns nil for unknown device classes.
func BuildBARModel(barData []byte, classCode uint32, profile *donor.BARProfile) *BARModel {
	// Use probe data when available, but bail if VFIO reported
	// everything as writable (breaks CC->CSTS handshake etc).
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
	model := specBARModelForClass(classCode, barData)
	if model != nil {
		validateModel(model)
	}
	return model
}

// specBARModelForClass returns the hardcoded spec model for a class, or nil.
func specBARModelForClass(classCode uint32, barData []byte) *BARModel {
	baseClass := (classCode >> 16) & 0xFF
	subClass := (classCode >> 8) & 0xFF
	progIF := classCode & 0xFF
	switch {
	case baseClass == 0x01 && subClass == 0x08 && progIF == 0x02:
		return buildNVMeBARModel(barData)
	case baseClass == 0x0C && subClass == 0x03 && progIF == 0x30:
		return buildXHCIBARModel(barData)
	case baseClass == 0x02 && subClass == 0x00:
		return buildEthernetBARModel(barData)
	case baseClass == 0x04 && subClass == 0x03:
		return buildAudioBARModel(barData)
	}

	profile := devclass.ProfileForClass(classCode)
	if profile == nil || len(profile.BARDefaults) == 0 {
		return nil
	}
	regs := make([]BARRegister, 0, len(profile.BARDefaults))
	for _, d := range profile.BARDefaults {
		regs = append(regs, BARRegister{
			Offset:      d.Offset,
			Width:       d.Width,
			Reset:       d.Reset,
			RWMask:      d.RWMask,
			W1CMask:     d.W1CMask,
			Name:        d.Name,
			IsRW1C:      d.IsRW1C,
			IsFSMDriven: d.IsFSMDriven,
		})
	}
	populateResetValues(regs, barData)
	if classCode&0xffff00 == 0x030000 {
		for i := range regs {
			switch regs[i].Offset {
			case 0x00:
				if regs[i].Reset == 0 {
					regs[i].Reset = 0x134000A1
				}
			case 0x1800:
				if regs[i].Reset == 0 {
					regs[i].Reset = 0x1B0610DE
				}
			}
		}
	}
	return &BARModel{Size: len(barData), Registers: regs}
}

// specRegAttr carries the W1CMask and IsFSMDriven flags the spec assigns to a
// known offset, so a probe-synthesized model stays consistent with the spec.
type specRegAttr struct {
	W1CMask     uint32
	IsFSMDriven bool
}

// specRegisterAttrs returns the spec attributes for every offset the spec model
// marks as W1C or device-driven.
func specRegisterAttrs(classCode uint32) map[uint32]specRegAttr {
	spec := specBARModelForClass(classCode, nil)
	if spec == nil {
		return nil
	}
	var out map[uint32]specRegAttr
	for _, r := range spec.Registers {
		if r.W1CMask != 0 || r.IsFSMDriven {
			if out == nil {
				out = make(map[uint32]specRegAttr)
			}
			out[r.Offset] = specRegAttr{W1CMask: r.W1CMask, IsFSMDriven: r.IsFSMDriven}
		}
	}
	return out
}

// validateModel checks for misaligned or duplicate offsets.
func validateModel(m *BARModel) {
	seen := make(map[uint32]string, len(m.Registers))
	for _, r := range m.Registers {
		if m.Size > 0 && int(r.Offset) >= m.Size {
			slog.Warn("barmodel: register offset beyond BAR size",
				"reg", r.Name, "offset", fmt.Sprintf("0x%X", r.Offset), "bar_size", m.Size)
		}
		if r.Offset%4 != 0 {
			panic(fmt.Sprintf("barmodel: register %s at offset 0x%X is not DWORD-aligned", r.Name, r.Offset))
		}
		if prev, ok := seen[r.Offset]; ok {
			panic(fmt.Sprintf("barmodel: %s and %s share offset 0x%X", prev, r.Name, r.Offset))
		}
		// RW and W1C masks must be disjoint: a bit can't be both plain-writable
		// and write-1-to-clear.
		if r.W1CMask&r.RWMask != 0 {
			panic(fmt.Sprintf("barmodel: register %s at offset 0x%X has overlapping W1C/RW masks (W1CMask=0x%08X & RWMask=0x%08X = 0x%08X) — they must be disjoint",
				r.Name, r.Offset, r.W1CMask, r.RWMask, r.W1CMask&r.RWMask))
		}
		if r.Width > 0 {
			var widthMask uint32 = 0xFFFFFFFF
			if bits := r.Width * 8; bits < 32 {
				widthMask = (1 << bits) - 1
			}
			if r.W1CMask&^widthMask != 0 || r.RWMask&^widthMask != 0 {
				panic(fmt.Sprintf("barmodel: register %s at offset 0x%X has mask bits beyond Width %d (W1CMask=0x%08X, RWMask=0x%08X, valid=0x%08X)",
					r.Name, r.Offset, r.Width, r.W1CMask, r.RWMask, widthMask))
			}
		}
		seen[r.Offset] = r.Name
	}
}

// NVMe BAR0 (spec 1.4, offsets 0x00–0x34).
func buildNVMeBARModel(barData []byte) *BARModel {
	regs := []BARRegister{
		// CAP - 64-bit RO
		{Offset: 0x00, Width: 4, Name: "CAP_LO", RWMask: 0x00000000},
		{Offset: 0x04, Width: 4, Name: "CAP_HI", RWMask: 0x00000000},
		// VS
		{Offset: 0x08, Width: 4, Name: "VS", RWMask: 0x00000000},
		// INTMS/INTMC: RO when MSI-X is active
		{Offset: 0x0C, Width: 4, Name: "INTMS", RWMask: 0x00000000},
		{Offset: 0x10, Width: 4, Name: "INTMC", RWMask: 0x00000000},
		// CC - driver writes EN, FSM watches
		{Offset: 0x14, Width: 4, Name: "CC", RWMask: 0x00FFFFF1},
		// CSTS - RO, FSM drives RDY
		{Offset: 0x1C, Width: 4, Name: "CSTS", RWMask: 0x00000000, IsFSMDriven: true},
		// NSSR
		{Offset: 0x20, Width: 4, Name: "NSSR", RWMask: 0xFFFFFFFF},
		// AQA
		{Offset: 0x24, Width: 4, Name: "AQA", RWMask: 0x0FFF0FFF},
		// ASQ - 64-bit
		{Offset: 0x28, Width: 4, Name: "ASQ_LO", RWMask: 0xFFFFF000},
		{Offset: 0x2C, Width: 4, Name: "ASQ_HI", RWMask: 0xFFFFFFFF},
		// ACQ - 64-bit
		{Offset: 0x30, Width: 4, Name: "ACQ_LO", RWMask: 0xFFFFF000},
		{Offset: 0x34, Width: 4, Name: "ACQ_HI", RWMask: 0xFFFFFFFF},
	}

	populateResetValues(regs, barData)

	// VS must match Identify Controller VER; stornvme raises Code 10 on mismatch.
	// Default to 1.4 like identify.go when the donor VS is absent/0.
	for i := range regs {
		switch regs[i].Offset {
		case 0x00:
			if regs[i].Reset == 0 {
				regs[i].Reset = 0x0040FF17
			}
		case 0x04:
			if regs[i].Reset == 0 {
				regs[i].Reset = 0x00000020
			}
		case 0x08:
			if regs[i].Reset == 0 {
				regs[i].Reset = 0x00010400
			}
		}
	}

	// stornvme needs RDY=1 at boot
	for i := range regs {
		if regs[i].Offset == 0x1C {
			regs[i].Reset |= 0x00000001  // RDY bit
			regs[i].Reset &^= 0x0000000C // clear SHST (shutdown status)
			break
		}
	}

	sz := len(barData)
	// Do not force 4096 here. When donor barData is empty (len==0), Size starts at 0
	// and is set to the final capped bar0Size by the caller (output/writer.go) if larger.
	// This ensures Size is never forced to 4096 when donor data or Bar0Size is larger.
	return &BARModel{
		Size:      sz,
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
		{Offset: 0x20, Width: 4, Name: "USBCMD", RWMask: 0x00002F0E, IsFSMDriven: true},
		// USBSTS (mostly RW1C)
		{Offset: 0x24, Width: 4, Name: "USBSTS", RWMask: 0x00000000, W1CMask: 0x0000041C, IsFSMDriven: true},
		// Page Size (read-only)
		{Offset: 0x28, Width: 4, Name: "PAGESIZE", RWMask: 0x00000000},
		// Device Notification Control
		{Offset: 0x34, Width: 4, Name: "DNCTRL", RWMask: 0x0000FFFF},
		// Command Ring Control - 64-bit
		{Offset: 0x38, Width: 4, Name: "CRCR_LO", RWMask: 0xFFFFFFF7},
		{Offset: 0x3C, Width: 4, Name: "CRCR_HI", RWMask: 0xFFFFFFFF},
		// Extended capability: Supported Protocol (USB 2.0, two ports).
		{Offset: 0x40, Width: 4, Name: "XECAP_SUPPORTED_PROTOCOL", Reset: 0x00000002, RWMask: 0x00000000},
		{Offset: 0x44, Width: 4, Name: "XECAP_PROTOCOL_NAME", Reset: 0x20425355, RWMask: 0x00000000},
		{Offset: 0x48, Width: 4, Name: "XECAP_PROTOCOL_PORTS", Reset: 0x00000201, RWMask: 0x00000000},
		{Offset: 0x4C, Width: 4, Name: "XECAP_PROTOCOL_SLOT", Reset: 0x00000000, RWMask: 0x00000000},
		// Device Context Base Address Array Pointer - 64-bit
		{Offset: 0x50, Width: 4, Name: "DCBAAP_LO", RWMask: 0xFFFFFFC0},
		{Offset: 0x54, Width: 4, Name: "DCBAAP_HI", RWMask: 0xFFFFFFFF},
		// Configure (CONFIG)
		{Offset: 0x58, Width: 4, Name: "CONFIG", RWMask: 0x000000FF},
		// PORTSC1/2: powered port (PP=bit9) so the host sees a usable port.
		{Offset: 0x420, Width: 4, Name: "PORTSC1", Reset: devclass.XHCIPortscReset, RWMask: devclass.XHCIPortscRWMask},
		{Offset: 0x430, Width: 4, Name: "PORTSC2", Reset: devclass.XHCIPortscReset, RWMask: devclass.XHCIPortscRWMask},
	}

	populateResetValues(regs, barData)
	if profile := devclass.ProfileForClass(0x0C0330); profile != nil {
		defaults := make(map[uint32]uint32, len(profile.BARDefaults))
		for _, value := range profile.BARDefaults {
			defaults[value.Offset] = value.Reset
		}
		for index := range regs {
			if regs[index].Reset == 0 && defaults[regs[index].Offset] != 0 {
				regs[index].Reset = defaults[regs[index].Offset]
			}
		}
	}

	// driver expects a running controller on first probe
	for i := range regs {
		switch regs[i].Offset {
		case 0x20: // USBCMD
			regs[i].Reset |= 0x00000001 // R/S bit
		case 0x24: // USBSTS
			regs[i].Reset &^= 0x00000001 // clear HCH (halted)
		case 0x420, 0x430: // populateResetValues clobbers from donor data; re-force the powered reset.
			regs[i].Reset = devclass.XHCIPortscReset
		}
	}

	sz := len(barData)
	// Do not force 4096 here. When donor barData is empty (len==0), Size starts at 0
	// and is set to the final capped bar0Size by the caller (output/writer.go) if larger.
	// This ensures Size is never forced to 4096 when donor data or Bar0Size is larger.
	return &BARModel{
		Size:      sz,
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
		{Offset: 0x34, Width: 4, Name: "CHIPCMD_DW", RWMask: 0xFF000000}, // byte 0x37 = ChipCmd (RxEn|TxEn writable)
		{Offset: 0x3C, Width: 4, Name: "INTRMASK", RWMask: 0xFFFFFFFF},
		{Offset: 0x40, Width: 4, Name: "TXCONFIG", RWMask: 0x00FF0000},
		{Offset: 0x44, Width: 4, Name: "RXCONFIG", RWMask: 0xFFFF7FFF},
		{Offset: 0x48, Width: 4, Name: "TIMER", RWMask: 0xFFFFFFFF},
		{Offset: 0x50, Width: 4, Name: "RXMAXSIZE", RWMask: 0x00003FFF},
		{Offset: 0x58, Width: 4, Name: "CPLUSCMD", RWMask: 0x0000FFFF},
		{Offset: 0x6C, Width: 4, Name: "PHYSTATUS", RWMask: 0x00000000},                // RO
		{Offset: 0xDC, Width: 4, Name: "PHYAR", RWMask: 0xFFFFFFFF, IsFSMDriven: true}, // bit31 = ready
		{Offset: 0xE0, Width: 4, Name: "ERIAR", RWMask: 0xFFFFFFFF, IsFSMDriven: true}, // bit31 = done
		{Offset: 0xFC, Width: 4, Name: "RXMISSED", RWMask: 0x00000000},                 // RO
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

	sz := len(barData)
	// Do not force 4096 here. When donor barData is empty (len==0), Size starts at 0
	// and is set to the final capped bar0Size by the caller (output/writer.go) if larger.
	// This ensures Size is never forced to 4096 when donor data or Bar0Size is larger.
	return &BARModel{
		Size:      sz,
		Registers: regs,
	}
}

// BuildIntelE1000BARModel returns the tested Intel I219-LM/e1000e register
// subset used by the descriptor DMA engine. It is intentionally separate from
// the class-wide Realtek model: PCI class 0x020000 alone does not identify a
// compatible descriptor or register layout.
func BuildIntelE1000BARModel(barData []byte) *BARModel {
	regs := []BARRegister{
		{Offset: 0x0000, Width: 4, Name: "CTRL", RWMask: 0xFFFFFFFF},
		{Offset: 0x0008, Width: 4, Name: "STATUS", Reset: 0x80080783},
		{Offset: 0x0014, Width: 4, Name: "EERD", RWMask: 0x00000001, IsFSMDriven: true},
		{Offset: 0x0020, Width: 4, Name: "MDIC", Reset: 0x08000000, RWMask: 0xFFFFFFFF, IsFSMDriven: true},
		{Offset: 0x00C0, Width: 4, Name: "ICR", IsFSMDriven: true},
		{Offset: 0x00C8, Width: 4, Name: "ICS", RWMask: 0xFFFFFFFF, IsFSMDriven: true},
		{Offset: 0x00D0, Width: 4, Name: "IMS", RWMask: 0xFFFFFFFF, IsFSMDriven: true},
		{Offset: 0x00D8, Width: 4, Name: "IMC", RWMask: 0xFFFFFFFF, IsFSMDriven: true},
		{Offset: 0x0100, Width: 4, Name: "RCTL", RWMask: 0xFFFFFFFF},
		{Offset: 0x0400, Width: 4, Name: "TCTL", RWMask: 0xFFFFFFFF},
		{Offset: 0x2800, Width: 4, Name: "RDBAL", RWMask: 0xFFFFFF80},
		{Offset: 0x2804, Width: 4, Name: "RDBAH", RWMask: 0xFFFFFFFF},
		{Offset: 0x2808, Width: 4, Name: "RDLEN", RWMask: 0x000FFF80},
		{Offset: 0x2810, Width: 4, Name: "RDH", IsFSMDriven: true},
		{Offset: 0x2818, Width: 4, Name: "RDT", RWMask: 0x0000FFFF},
		{Offset: 0x3800, Width: 4, Name: "TDBAL", RWMask: 0xFFFFFF80},
		{Offset: 0x3804, Width: 4, Name: "TDBAH", RWMask: 0xFFFFFFFF},
		{Offset: 0x3808, Width: 4, Name: "TDLEN", RWMask: 0x000FFF80},
		{Offset: 0x3810, Width: 4, Name: "TDH", IsFSMDriven: true},
		{Offset: 0x3818, Width: 4, Name: "TDT", RWMask: 0x0000FFFF},
		{Offset: 0x5400, Width: 4, Name: "RAL0", Reset: 0x49435002, RWMask: 0xFFFFFFFF},
		{Offset: 0x5404, Width: 4, Name: "RAH0", Reset: 0x8000454C, RWMask: 0x8000FFFF},
	}

	populateResetValues(regs, barData)
	for i := range regs {
		switch regs[i].Offset {
		case 0x0008:
			if regs[i].Reset == 0 {
				regs[i].Reset = 0x80080783
			}
		case 0x0020:
			if regs[i].Reset == 0 {
				regs[i].Reset = 0x08000000
			}
		case 0x5400:
			if regs[i].Reset == 0 {
				regs[i].Reset = 0x49435002
			}
		case 0x5404:
			if regs[i].Reset == 0 {
				regs[i].Reset = 0x8000454C
			}
		}
	}
	return &BARModel{Size: len(barData), Registers: regs}
}

// HD Audio BAR0. Sub-word regs packed into DWORDs for SV template.
// buildAudioBARModel builds the HD Audio BAR0 register map.
// When donor BAR data is all 0xFF (no codec connected), spec defaults are used.
func buildAudioBARModel(barData []byte) *BARModel {
	regs := []BARRegister{
		// GCAP(16) + VMIN(8) + VMAJ(8) packed into one DWORD
		{Offset: 0x00, Width: 4, Name: "GCAP_VMIN_VMAJ", Reset: 0x01006401, RWMask: 0x00000000},
		// OUTPAY (15:0) + INPAY (31:16) — read-only stream payload capabilities
		{Offset: 0x04, Width: 4, Name: "OUTPAY_INPAY", Reset: 0x00400040, RWMask: 0x00000000},
		// GCTL - global control, CRST (bit 0) is the key writable bit
		{Offset: 0x08, Width: 4, Name: "GCTL", RWMask: 0x00000103},
		// WAKEEN(16) + STATESTS(16) packed into one DWORD
		// WAKEEN [15:0] is writable; STATESTS [31:16] is read-only status
		{Offset: 0x0C, Width: 4, Name: "WAKEEN_STATESTS", RWMask: 0x0000FFFF},
		// INTCTL - interrupt control
		{Offset: 0x20, Width: 4, Name: "INTCTL", RWMask: 0xC00000FF},
		// INTSTS - interrupt status
		{Offset: 0x24, Width: 4, Name: "INTSTS", RWMask: 0x00000000},
		// WALCLK - 32-bit free-running wall clock counter (read-only)
		{Offset: 0x30, Width: 4, Name: "WALCLK", Reset: 0x00000000, RWMask: 0x00000000},
		// CORB lower base address
		{Offset: 0x40, Width: 4, Name: "CORBLBASE", RWMask: 0xFFFFFF80},
		// CORB upper base address
		{Offset: 0x44, Width: 4, Name: "CORBUBASE", RWMask: 0xFFFFFFFF},
		// CORBWP(16) + CORBRP(16) packed
		{Offset: 0x48, Width: 4, Name: "CORBWP_CORBRP", RWMask: 0x0000FFFF, IsFSMDriven: true},
		// CORBCTL(8) + CORBSTS(8) + CORBSIZE(8) packed
		// CORBCTL: bit 1 (CORBRUN), bit 0 (CMEIE) writable; CORBSTS: bit 0 -> DWORD bit 8 (RPWP) RW1C; CORBSIZE: RO
		{Offset: 0x4C, Width: 4, Name: "CORBCTL_STS_SIZE", RWMask: 0x00030003, W1CMask: 0x00000100, IsRW1C: true},
		// RIRB lower base address
		{Offset: 0x50, Width: 4, Name: "RIRBLBASE", RWMask: 0xFFFFFF80},
		// RIRB upper base address
		{Offset: 0x54, Width: 4, Name: "RIRBUBASE", RWMask: 0xFFFFFFFF},
		// RIRBWP(16) + RINTCNT(16)
		{Offset: 0x58, Width: 4, Name: "RIRBWP_RINTCNT", RWMask: 0x0000FFFF, IsFSMDriven: true},
		// RIRBCTL(8) + RIRBSTS(8) + RIRBSIZE(8)
		// RIRBCTL: bit 0 (RINTCTL), bit 1 (RIRBDMAEN), bit 2 (OIC) writable
		// RIRBSTS: bit 8 (RINTFL) RW1C, bit 9 (OIS) RW1C
		{Offset: 0x5C, Width: 4, Name: "RIRBCTL_STS_SIZE", RWMask: 0x00000007, W1CMask: 0x00000700, IsRW1C: true, IsFSMDriven: true},
		// RIRBINTSTS - RIRB interrupt status (RW1C: bit 0 INTFL)
		{Offset: 0x60, Width: 4, Name: "RIRBINTSTS", RWMask: 0x00000000, W1CMask: 0x00000001, IsRW1C: true, IsFSMDriven: true},
		// IC (Immediate Command) - driver writes codec command
		{Offset: 0x64, Width: 4, Name: "IC", RWMask: 0xFFFFFFFF},
		// IR (Immediate Response) - driver reads codec response (RO)
		{Offset: 0x68, Width: 4, Name: "IR", RWMask: 0x00000000},
		// RIRBRESP_LO - RIRB response data lower 32 bits (RO, DMA-served)
		{Offset: 0x70, Width: 4, Name: "RIRBRESP_LO", RWMask: 0x00000000, IsFSMDriven: true},
		// RIRBRESP_HI - RIRB response data upper 32 bits (RO, DMA-served)
		// Must be at 0x74 (immediately after 0x70) - the driver reads 8 bytes
		// from offset 0x70 as a single RIRB entry. A gap at 0x74 would cause
		// the upper 32 bits to read as zero.
		{Offset: 0x74, Width: 4, Name: "RIRBRESP_HI", RWMask: 0x00000000, IsFSMDriven: true},
	}

	// Check if donor BAR data is all 0xFF (no codec connected).
	// When the HD Audio codec BAR has no responding codec, reads return 0xFF.
	// In this case we skip populateResetValues and use spec defaults directly,
	// because 0xFF | default = 0xFF (bitwise OR with all-ones is a no-op).
	allFF := isBARDataAllFF(barData)
	if !allFF {
		populateResetValues(regs, barData)
		// Apply spec defaults only where donor data didn't cover
		for i := range regs {
			switch regs[i].Offset {
			case 0x00:
				if regs[i].Reset == 0 {
					regs[i].Reset = 0x01006401
				}
			case 0x08:
				regs[i].Reset |= 0x00000001
			case 0x0C:
				regs[i].Reset = (regs[i].Reset & 0x0000FFFF) | 0x00010000
			case 0x4C:
				if regs[i].Reset == 0 {
					regs[i].Reset = 0x00820000
				}
			case 0x5C:
				if regs[i].Reset == 0 {
					regs[i].Reset = 0x00820000
				}
			}
		}
	} else {
		// No valid donor data - use HD Audio spec defaults.
		regs[0].Reset = 0x01006401  // GCAP=6401h (2-in/2-out, 44.1kHz, B64OK), VMIN=0, VMAJ=1
		regs[1].Reset = 0x00400040  // OUTPAY=0x40 (64 bytes), INPAY=0x40 (64 bytes)
		regs[2].Reset = 0x00000001  // GCTL.CRST=1 (out of reset)
		regs[3].Reset = 0x00010000  // STATESTS: codec 0 present (upper 16 bits)
		regs[4].Reset = 0x00000000  // INTCTL: no state interrupts enabled
		regs[5].Reset = 0x00000000  // INTSTS: no interrupts pending
		regs[6].Reset = 0x00000000  // WALCLK: wall clock counter starts at 0
		regs[7].Reset = 0x00000000  // CORBLBASE: driver will program before use
		regs[8].Reset = 0x00000000  // CORBUBASE: upper 32 bits of base
		regs[9].Reset = 0x00000000  // CORBWP=0, CORBRP=0 (both at start)
		regs[10].Reset = 0x00820000 // CORBSIZE=0x82 (supports 256/16/2 entries)
		regs[11].Reset = 0x00000000 // RIRBLBASE: driver will program before use
		regs[12].Reset = 0x00000000 // RIRBUBASE: upper 32 bits of base
		regs[13].Reset = 0x00000000 // RIRBWP=0, RINTCNT=0
		regs[14].Reset = 0x00820000 // RIRBSIZE=0x82 (supports 256/16/2 entries)
		// regs[16] (IC at 0x64) and regs[17] (IR at 0x68) default to 0 - correct.
	}

	sz := len(barData)
	// Do not force 4096 here. When donor barData is empty (len==0), Size starts at 0
	// and is set to the final capped bar0Size by the caller (output/writer.go) if larger.
	// This ensures Size is never forced to 4096 when donor data or Bar0Size is larger.
	return &BARModel{
		Size:      sz,
		Registers: regs,
	}
}

// SynthesizeBARModel builds a model from probe data, reconciling it with the spec.
// Spec-known W1C bits win over the probe and are emitted as real W1C; bits the
// probe flagged as W1C at unknown offsets are clamped to read-only (the probe is
// not trustworthy enough to emit live W1C hardware for them).
func SynthesizeBARModel(profile *donor.BARProfile, classCode uint32) *BARModel {
	if profile == nil || len(profile.Probes) == 0 {
		return nil
	}

	nameHints := classRegisterNames(classCode)
	attrs := specRegisterAttrs(classCode)

	regs := make([]BARRegister, 0, len(profile.Probes))
	for _, probe := range profile.Probes {
		attr := attrs[probe.Offset]
		// Drop dead regs, but keep ones the spec knows about — a status register
		// legitimately reads 0 when idle.
		if probe.Original == 0 && probe.RWMask == 0 && attr.W1CMask == 0 && !attr.IsFSMDriven {
			continue
		}

		name := fmt.Sprintf("REG_0x%03X", probe.Offset)
		if hint, ok := nameHints[probe.Offset]; ok {
			name = hint
		}

		rwMask := probe.RWMask
		var w1cMask uint32
		isRW1C := false

		// Spec W1C bits are authoritative.
		if attr.W1CMask != 0 {
			w1cMask = attr.W1CMask
			isRW1C = true
			rwMask &^= attr.W1CMask
		}

		// Probe-suspected W1C bits are forced read-only, not emitted as W1C.
		if probe.W1CMask != 0 {
			rwMask &^= probe.W1CMask
		} else if probe.MaybeRW1C {
			rwMask = 0
		}

		regs = append(regs, BARRegister{
			Offset:      probe.Offset,
			Width:       4,
			Reset:       probe.Original,
			RWMask:      rwMask,
			W1CMask:     w1cMask,
			Name:        name,
			IsRW1C:      isRW1C,
			IsFSMDriven: attr.IsFSMDriven,
		})
	}

	if len(regs) == 0 {
		return nil
	}

	model := &BARModel{
		Size:      profile.Size,
		Registers: regs,
	}
	validateModel(model)
	return model
}

// isProbeDataReliable rejects VFIO dumps where 90%+ of regs report
// fully writable - usually means the probe couldn't actually write.
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

// isBARDataAllFF checks if BAR memory is entirely 0xFF (no responding device).
// This happens when there's no codec connected - reads return all-ones.
func isBARDataAllFF(data []byte) bool {
	if len(data) == 0 {
		return true
	}
	for _, b := range data {
		if b != 0xFF {
			return false
		}
	}
	return true
}

// classRegisterNames returns offset->name hints from the device profile.
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
