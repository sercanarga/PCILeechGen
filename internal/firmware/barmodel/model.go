// Package barmodel builds BAR register maps for SV codegen.
package barmodel

import (
	"fmt"
	"log/slog"
	"sort"

	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/donor/mmio"
	"github.com/sercanarga/pcileechgen/internal/firmware/devclass"
	"github.com/sercanarga/pcileechgen/internal/util"
)

type RegisterAccessKind string

const (
	RegisterReadOnly  RegisterAccessKind = "read_only"
	RegisterReadWrite RegisterAccessKind = "read_write"
	RegisterRW1C      RegisterAccessKind = "rw1c"
	RegisterFSMDriven RegisterAccessKind = "fsm"
)

// BARRegister describes a single register inside a BAR.
type BARRegister struct {
	Offset      uint32 // byte offset in BAR
	Width       int    // 1, 2, or 4 bytes
	Reset       uint32 // reset/initial value (from donor snapshot or spec default)
	RWMask      uint32 // writable bits (1 = host can write, 0 = read-only)
	RW1CMask    uint32 // write-1-to-clear bits inside RWMask
	Name        string // human-readable register name
	IsRW1C      bool   // true if this register uses write-1-to-clear semantics
	IsFSMDriven bool   // true if driven by a dedicated FSM always block (excluded from generic reset/write)
	// SequentialRead marks a register whose read value advances through a
	// small ROM on each read to the same address (stateful reads). Models
	// devices whose drivers expect different data on repeated reads of one
	// offset (e.g. Atheros ath9k EEPROM request/response handshakes).
	SequentialRead bool
	// SequentialValues optionally provides the exact readback sequence for
	// SequentialRead registers. When empty, the SV template falls back to
	// Reset + a small wrapping counter.
	SequentialValues []uint32
}

// BARModel is the complete register map that ends up in the SV template.
type BARModel struct {
	Size      int
	Registers []BARRegister // ordered by Offset
	// StaticShadow holds donor-derived read-only words for 4-byte-aligned
	// offsets that are NOT covered by Registers (i.e. unmodeled static regs).
	// The SV template emits a read case returning these before the offset-echo
	// fallback, so the full BAR0 mirrors the donor for static offsets without
	// bloating the writable register file. Populated from donor barData.
	StaticShadow []StaticWord
}

// StaticWord is one donor-derived read-only word for the static shadow.
type StaticWord struct {
	Offset uint32
	Value  uint32
}

// BuildBARModel returns a register map from probe data or spec tables.
// Returns nil for unknown device classes.
//
// vendorID is only used to look up vendor-specific device profiles (e.g. the
// Atheros AR9287 EEPROM handshake registers); pass 0 to use class-only
// resolution. See BuildBARModelForDevice for the vendor-aware entry point.
func BuildBARModel(barData []byte, classCode uint32, profile *donor.BARProfile) *BARModel {
	return BuildBARModelForDeviceWithOverlay(barData, classCode, 0, profile, nil)
}

// BuildBARModelForDevice is the vendor-aware BAR model builder. It resolves the
// active device profile from the class code + vendor ID and, after building the
// probe- or spec-table-based model, applies the profile's stateful-read
// register defaults (BARDefault.SequentialRead) so generated read logic
// advances those offsets per read (e.g. Atheros ath9k EEPROM request/response
// handshakes that the donor probe / spec table cannot describe on their own).
func BuildBARModelForDevice(barData []byte, classCode uint32, vendorID uint16, profile *donor.BARProfile) *BARModel {
	return BuildBARModelForDeviceWithOverlay(barData, classCode, vendorID, profile, nil)
}

// BuildBARModelForDeviceWithOverlay applies an optional trace-derived overlay
// before final static-shadow population.
func BuildBARModelForDeviceWithOverlay(barData []byte, classCode uint32, vendorID uint16, profile *donor.BARProfile, overlay *mmio.TraceBAROverlay) *BARModel {
	return BuildBARModelForDeviceIDWithOverlay(barData, classCode, vendorID, 0, profile, overlay)
}

// BuildBARModelForDeviceIDWithOverlay applies an optional trace-derived overlay
// and allows device-specific profile resolution when a vendor ID alone is not
// narrow enough.
func BuildBARModelForDeviceIDWithOverlay(barData []byte, classCode uint32, vendorID, deviceID uint16, profile *donor.BARProfile, overlay *mmio.TraceBAROverlay) *BARModel {
	devProfile := (*devclass.DeviceProfile)(nil)
	if vendorID != 0 {
		devProfile = devclass.ProfileForDevice(classCode, vendorID, deviceID)
	}

	model := buildBARModelCore(barData, classCode, profile)
	if model == nil && devProfile != nil && devProfile.ClassName != "Generic" && len(devProfile.BARDefaults) > 0 {
		model = buildBARModelFromProfile(devProfile, barData)
	}
	if model == nil {
		return nil
	}
	validateModel(model)
	// Even when probe data was deemed unreliable and we fell back to the
	// spec table, overlay per-donor write masks / reset values / RW1C flags
	// from the probe results where they intersect the spec registers. This
	// keeps the dynamic FSM registers (IsFSMDriven) on spec semantics while
	// letting donor-observed writability refine the static ones.
	if profile != nil && len(profile.Probes) > 0 {
		applyDonorProbeOverlay(model, profile)
	}
	// Apply vendor-specific stateful-read register defaults (SequentialRead)
	// from the resolved device profile BEFORE the trace/static shadow so the
	// handshake offsets are treated as modeled (and excluded from later shadow
	// cases), avoiding duplicate read-case labels in the generated SV.
	if devProfile != nil {
		applyProfileSequentialRead(model, devProfile, barData)
	}
	applyTraceOverlay(model, overlay)
	// Seed a static read shadow from donor barData for offsets the model
	// does not cover, so the full BAR0 mirrors the donor for static regs.
	if len(barData) > 0 {
		populateStaticShadow(model, barData)
	}
	return model
}

// buildBARModelCore is the probe-then-spec-table dispatch without the
// post-processing overlays (kept here so BuildBARModelForDevice can layer them
// in a fixed order).
func buildBARModelCore(barData []byte, classCode uint32, profile *donor.BARProfile) *BARModel {
	var model *BARModel

	// Use probe data when available, but bail if VFIO reported
	// everything as writable (breaks CC->CSTS handshake etc).
	if profile != nil && len(profile.Probes) > 0 {
		if !isProbeDataReliable(profile) {
			slog.Warn("BAR probe data unreliable (all registers report fully writable), falling back to spec-based model",
				"probes", len(profile.Probes))
		} else {
			model = SynthesizeBARModel(profile, classCode)
		}
	}

	// fall back to hardcoded spec tables
	if model == nil {
		if isNVMeClass(classCode) {
			model = buildNVMeBARModel(barData)
		} else if isXHCIClass(classCode) {
			model = buildXHCIBARModel(barData)
		} else {
			baseClass := (classCode >> 16) & 0xFF
			subClass := (classCode >> 8) & 0xFF
			switch {
			case baseClass == 0x02 && subClass == 0x00:
				model = buildEthernetBARModel(barData)
			case baseClass == 0x04 && subClass == 0x03:
				model = buildAudioBARModel(barData)
			}
		}
	}
	if model != nil && isNVMeClass(classCode) {
		ensureNVMeFSMRegisters(model, barData)
	}
	if model != nil && isXHCIClass(classCode) {
		ensureXHCIFSMRegisters(model, barData)
	}
	return model
}

func buildBARModelFromProfile(p *devclass.DeviceProfile, barData []byte) *BARModel {
	if p == nil || len(p.BARDefaults) == 0 {
		return nil
	}
	regs := make([]BARRegister, 0, len(p.BARDefaults))
	for _, d := range p.BARDefaults {
		reset := d.Reset
		if len(barData) > 0 && int(d.Offset)+4 <= len(barData) {
			if v := util.ReadLE32(barData, int(d.Offset)); v != 0 {
				reset = v
			}
		}
		regs = append(regs, BARRegister{
			Offset:           d.Offset,
			Width:            d.Width,
			Reset:            reset,
			RWMask:           d.RWMask,
			Name:             d.Name,
			IsRW1C:           d.IsRW1C,
			IsFSMDriven:      d.IsFSMDriven,
			SequentialRead:   d.SequentialRead,
			SequentialValues: append([]uint32(nil), d.SequentialValues...),
		})
	}
	model := &BARModel{Size: len(barData), Registers: regs}
	sortRegistersByOffset(model)
	return model
}

// applyProfileSequentialRead ensures every device-profile register marked
// SequentialRead is present in the model with the SequentialRead flag set.
// The donor probe / spec table cannot describe stateful reads on their own, so
// the profile is the authority for these offsets:
//   - If a register already exists (probe or spec), only the flag is set; the
//     donor-observed reset / RWMask are kept (the donor snapshot already holds
//     the handshake sentinel for a real device).
//   - If no register exists, a read-only one is added, seeded from the profile
//     BARDefault reset (or the donor barData word when available) so the first
//     read returns the handshake base value.
func applyProfileSequentialRead(model *BARModel, devProfile *devclass.DeviceProfile, barData []byte) {
	if model == nil || devProfile == nil {
		return
	}
	byOff := make(map[uint32]int, len(model.Registers))
	for i, r := range model.Registers {
		byOff[r.Offset] = i
	}
	for _, d := range devProfile.BARDefaults {
		if !d.SequentialRead {
			continue
		}
		if idx, ok := byOff[d.Offset]; ok {
			model.Registers[idx].SequentialRead = true
			// Keep donor reset when present; fall back to the profile default.
			if model.Registers[idx].Reset == 0 {
				model.Registers[idx].Reset = d.Reset
			}
			if len(d.SequentialValues) > 0 {
				seq := append([]uint32(nil), d.SequentialValues...)
				seq[0] = model.Registers[idx].Reset
				model.Registers[idx].SequentialValues = seq
			}
			continue
		}
		reset := d.Reset
		if len(barData) > 0 && int(d.Offset)+4 <= len(barData) {
			if v := util.ReadLE32(barData, int(d.Offset)); v != 0 {
				reset = v
			}
		}
		reg := BARRegister{
			Offset:         d.Offset,
			Width:          d.Width,
			Reset:          reset,
			RWMask:         d.RWMask,
			Name:           d.Name,
			SequentialRead: true,
		}
		if len(d.SequentialValues) > 0 {
			reg.SequentialValues = append([]uint32(nil), d.SequentialValues...)
			reg.SequentialValues[0] = reset
		}
		model.Registers = append(model.Registers, reg)
		byOff[d.Offset] = len(model.Registers) - 1
	}
	// Re-sort by offset so the generated case statement stays ordered.
	sortRegistersByOffset(model)
}

func applyTraceOverlay(model *BARModel, overlay *mmio.TraceBAROverlay) {
	if model == nil || overlay == nil {
		return
	}

	byOff := make(map[uint32]int, len(model.Registers))
	for i, r := range model.Registers {
		byOff[r.Offset] = i
	}

	for off, seq := range overlay.Sequential {
		if len(seq) == 0 {
			continue
		}
		if idx, ok := byOff[off]; ok {
			if model.Registers[idx].IsFSMDriven {
				continue
			}
			model.Registers[idx].SequentialRead = true
			model.Registers[idx].SequentialValues = append([]uint32(nil), seq...)
			continue
		}
		model.Registers = append(model.Registers, BARRegister{
			Offset:           off,
			Width:            4,
			Reset:            seq[0],
			RWMask:           0,
			Name:             fmt.Sprintf("TRACE_SEQ_0x%08X", off),
			SequentialRead:   true,
			SequentialValues: append([]uint32(nil), seq...),
		})
		byOff[off] = len(model.Registers) - 1
	}

	sortRegistersByOffset(model)

	for off, writeMask := range overlay.WriteMask {
		if writeMask == 0 {
			continue
		}
		if idx, ok := byOff[off]; ok {
			if model.Registers[idx].IsFSMDriven {
				continue
			}
			if model.Registers[idx].RWMask == 0 {
				model.Registers[idx].RWMask = writeMask
			} else {
				model.Registers[idx].RWMask |= writeMask
			}
			continue
		}

		reg := BARRegister{
			Offset: off,
			Width:  4,
			Reset:  overlay.Static[off],
			RWMask: writeMask,
			Name:   fmt.Sprintf("TRACE_WR_0x%08X", off),
		}
		model.Registers = append(model.Registers, reg)
		byOff[off] = len(model.Registers) - 1
	}

	for off, rw1cMask := range overlay.RW1CMask {
		if rw1cMask == 0 {
			continue
		}
		if idx, ok := byOff[off]; ok {
			if model.Registers[idx].IsFSMDriven {
				continue
			}
			model.Registers[idx].IsRW1C = true
			model.Registers[idx].RW1CMask = rw1cMask
			if model.Registers[idx].RWMask == 0 {
				model.Registers[idx].RWMask = rw1cMask
			} else {
				model.Registers[idx].RWMask |= rw1cMask
			}
			continue
		}

		reg := BARRegister{
			Offset:   off,
			Width:    4,
			Reset:    overlay.Static[off],
			RWMask:   rw1cMask,
			RW1CMask: rw1cMask,
			Name:     fmt.Sprintf("TRACE_RW1C_0x%08X", off),
			IsRW1C:   true,
		}
		model.Registers = append(model.Registers, reg)
		byOff[off] = len(model.Registers) - 1
	}

	sortRegistersByOffset(model)

	existingShadow := make(map[uint32]bool, len(model.StaticShadow))
	for _, sw := range model.StaticShadow {
		existingShadow[sw.Offset] = true
	}
	for off, val := range overlay.Static {
		if _, modeled := byOff[off]; modeled || existingShadow[off] {
			continue
		}
		model.StaticShadow = append(model.StaticShadow, StaticWord{Offset: off, Value: val})
	}
	sort.Slice(model.StaticShadow, func(i, j int) bool {
		return model.StaticShadow[i].Offset < model.StaticShadow[j].Offset
	})
}

// sortRegistersByOffset sorts model.Registers by Offset (stable).
func sortRegistersByOffset(model *BARModel) {
	sort.SliceStable(model.Registers, func(i, j int) bool {
		return model.Registers[i].Offset < model.Registers[j].Offset
	})
}

// validateModel checks for misaligned or duplicate offsets.
func validateModel(m *BARModel) {
	seen := make(map[uint32]string, len(m.Registers))
	for _, r := range m.Registers {
		if int(r.Offset) >= m.Size {
			slog.Warn("barmodel: register offset beyond BAR size",
				"reg", r.Name, "offset", fmt.Sprintf("0x%X", r.Offset), "bar_size", m.Size)
		}
		if r.Offset%4 != 0 {
			panic(fmt.Sprintf("barmodel: register %s at offset 0x%X is not DWORD-aligned", r.Name, r.Offset))
		}
		if prev, ok := seen[r.Offset]; ok {
			panic(fmt.Sprintf("barmodel: %s and %s share offset 0x%X", prev, r.Name, r.Offset))
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
		// USBSTS: status bits are write-1-to-clear (RW1C). Bits: 0=HCH (FSM),
		// 2=HSE, 3=INTR, 4=PCD, 10=HCE. HCH stays FSM-driven; the rest are
		// cleared by the host writing 1. Kept IsFSMDriven so the xHCI FSM
		// (the sole always block driving this reg) can both set event bits
		// and service host W1C clears in one place. (Hardening: the
		// original ignored host USBSTS writes entirely.)
		{Offset: 0x24, Width: 4, Name: "USBSTS", RWMask: 0x0000041C, IsRW1C: true, IsFSMDriven: true},
		// Page Size (read-only)
		{Offset: 0x28, Width: 4, Name: "PAGESIZE", RWMask: 0x00000000},
		// Device Notification Control
		{Offset: 0x34, Width: 4, Name: "DNCTRL", RWMask: 0x0000FFFF},
		// Command Ring Control - 64-bit
		{Offset: 0x38, Width: 4, Name: "CRCR_LO", RWMask: 0xFFFFFFF7},
		{Offset: 0x3C, Width: 4, Name: "CRCR_HI", RWMask: 0xFFFFFFFF},
		// Runtime interrupter 0
		{Offset: 0x220, Width: 4, Name: "IMAN0", RWMask: 0x00000003, IsRW1C: true, IsFSMDriven: true},
		{Offset: 0x224, Width: 4, Name: "IMOD0", RWMask: 0xFFFFFFFF},
		{Offset: 0x228, Width: 4, Name: "ERSTSZ0", RWMask: 0x0000FFFF},
		{Offset: 0x230, Width: 4, Name: "ERSTBA_LO", RWMask: 0xFFFFFFC0},
		{Offset: 0x234, Width: 4, Name: "ERSTBA_HI", RWMask: 0xFFFFFFFF},
		{Offset: 0x238, Width: 4, Name: "ERDP_LO", RWMask: 0xFFFFFFF8},
		{Offset: 0x23C, Width: 4, Name: "ERDP_HI", RWMask: 0xFFFFFFFF},
		// Port Status and Control registers are driven by the xHCI FSM block.
		{Offset: 0x420, Width: 4, Name: "PORTSC1", Reset: 0x000002A0, RWMask: 0x8EFFC3F2, IsRW1C: true, IsFSMDriven: true},
		{Offset: 0x430, Width: 4, Name: "PORTSC2", Reset: 0x000002A0, RWMask: 0x8EFFC3F2, IsRW1C: true, IsFSMDriven: true},
		// Device Context Base Address Array Pointer - 64-bit
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

	sz := len(barData)
	// Do not force 4096 here. When donor barData is empty (len==0), Size starts at 0
	// and is set to the final capped bar0Size by the caller (output/writer.go) if larger.
	// This ensures Size is never forced to 4096 when donor data or Bar0Size is larger.
	return &BARModel{
		Size:      sz,
		Registers: regs,
	}
}

func isNVMeClass(classCode uint32) bool {
	baseClass := (classCode >> 16) & 0xFF
	subClass := (classCode >> 8) & 0xFF
	progIF := classCode & 0xFF
	return baseClass == 0x01 && subClass == 0x08 && progIF == 0x02
}

func isXHCIClass(classCode uint32) bool {
	baseClass := (classCode >> 16) & 0xFF
	subClass := (classCode >> 8) & 0xFF
	progIF := classCode & 0xFF
	return baseClass == 0x0C && subClass == 0x03 && progIF == 0x30
}

func ensureNVMeFSMRegisters(model *BARModel, barData []byte) {
	spec := buildNVMeBARModel(barData)
	byOff := make(map[uint32]int, len(model.Registers))
	for i, reg := range model.Registers {
		byOff[reg.Offset] = i
	}
	for _, reg := range spec.Registers {
		switch reg.Offset {
		case 0x14, 0x1C, 0x24, 0x28, 0x2C, 0x30, 0x34:
		default:
			continue
		}
		if idx, ok := byOff[reg.Offset]; ok {
			model.Registers[idx].Width = reg.Width
			model.Registers[idx].Name = reg.Name
			model.Registers[idx].RWMask = reg.RWMask
			model.Registers[idx].IsRW1C = reg.IsRW1C
			model.Registers[idx].IsFSMDriven = reg.IsFSMDriven
			if model.Registers[idx].Reset == 0 {
				model.Registers[idx].Reset = reg.Reset
			}
			continue
		}
		model.Registers = append(model.Registers, reg)
	}
	sortRegistersByOffset(model)
}

func ensureXHCIFSMRegisters(model *BARModel, barData []byte) {
	spec := buildXHCIBARModel(barData)
	byOff := make(map[uint32]int, len(model.Registers))
	for i, reg := range model.Registers {
		byOff[reg.Offset] = i
	}
	for _, reg := range spec.Registers {
		switch reg.Offset {
		case 0x20, 0x24, 0x220, 0x420, 0x430:
		default:
			continue
		}
		if idx, ok := byOff[reg.Offset]; ok {
			model.Registers[idx].Width = reg.Width
			model.Registers[idx].Name = reg.Name
			model.Registers[idx].RWMask = reg.RWMask
			model.Registers[idx].IsRW1C = reg.IsRW1C
			model.Registers[idx].IsFSMDriven = reg.IsFSMDriven
			if model.Registers[idx].Reset == 0 {
				model.Registers[idx].Reset = reg.Reset
			}
			continue
		}
		model.Registers = append(model.Registers, reg)
	}
	sortRegistersByOffset(model)
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
		{Offset: 0x4C, Width: 4, Name: "CORBCTL_STS_SIZE", RWMask: 0x00030003, IsRW1C: true},
		// RIRB lower base address
		{Offset: 0x50, Width: 4, Name: "RIRBLBASE", RWMask: 0xFFFFFF80},
		// RIRB upper base address
		{Offset: 0x54, Width: 4, Name: "RIRBUBASE", RWMask: 0xFFFFFFFF},
		// RIRBWP(16) + RINTCNT(16)
		{Offset: 0x58, Width: 4, Name: "RIRBWP_RINTCNT", RWMask: 0x0000FFFF, IsFSMDriven: true},
		// RIRBCTL(8) + RIRBSTS(8) + RIRBSIZE(8)
		// RIRBCTL: bit 0 (RINTCTL), bit 1 (RIRBDMAEN), bit 2 (OIC) writable
		// RIRBSTS: bit 8 (RINTFL) RW1C, bit 9 (OIS) RW1C
		{Offset: 0x5C, Width: 4, Name: "RIRBCTL_STS_SIZE", RWMask: 0x00000007, IsRW1C: true, IsFSMDriven: true},
		// RIRBINTSTS - RIRB interrupt status (RW1C: bit 0 INTFL)
		{Offset: 0x60, Width: 4, Name: "RIRBINTSTS", RWMask: 0x00000001, IsRW1C: true, IsFSMDriven: true},
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
			rwMask = 0 // RW1C -> force RO
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

// applyDonorProbeOverlay refines spec-table registers with donor-observed
// write masks, RW1C flags, and reset values. FSM-driven registers
// (IsFSMDriven) keep their spec semantics (the FSM is the authority for those
// offsets); everything else inherits the donor's per-bit writability so host
// writes to read-only donor bits are rejected exactly as on the real device.
func applyDonorProbeOverlay(model *BARModel, profile *donor.BARProfile) {
	if profile == nil || model == nil {
		return
	}
	byOff := make(map[uint32]donor.BARProbeResult, len(profile.Probes))
	for _, p := range profile.Probes {
		byOff[p.Offset] = p
	}
	for i := range model.Registers {
		r := &model.Registers[i]
		probe, ok := byOff[r.Offset]
		if !ok {
			continue
		}
		if r.IsFSMDriven {
			// FSM owns this offset; only let the donor refine the reset seed
			// for non-handshake bits the FSM doesn't touch. Keep it simple: do
			// not override FSM registers.
			continue
		}
		if probe.MaybeRW1C {
			r.IsRW1C = true
			r.RWMask = 0 // RW1C registers are modeled as read-only to the host
		} else {
			r.RWMask = probe.RWMask
		}
		r.Reset = probe.Original
	}
}

// populateStaticShadow collects donor 4-byte words at offsets NOT already
// covered by Registers, so the generated read case returns the donor value
// for those static offsets instead of the offset-echo fallback. Only non-zero
// words are emitted (zero reads are already the natural default for unmapped
// space). Capped to keep the generated case statement bounded.
func populateStaticShadow(model *BARModel, barData []byte) {
	if model == nil || len(barData) == 0 {
		return
	}
	const maxShadow = 256
	existing := make(map[uint32]bool, len(model.StaticShadow)+len(model.Registers))
	shadow := make([]StaticWord, 0, maxShadow)
	for _, r := range model.Registers {
		existing[r.Offset] = true
	}
	for _, sw := range model.StaticShadow {
		if existing[sw.Offset] || len(shadow) >= maxShadow {
			continue
		}
		existing[sw.Offset] = true
		shadow = append(shadow, sw)
	}
	for off := 0; off+4 <= len(barData) && len(shadow) < maxShadow; off += 4 {
		if existing[uint32(off)] {
			continue
		}
		val := util.ReadLE32(barData, off)
		if val == 0 {
			continue
		}
		existing[uint32(off)] = true
		shadow = append(shadow, StaticWord{Offset: uint32(off), Value: val})
	}
	sort.Slice(shadow, func(i, j int) bool {
		return shadow[i].Offset < shadow[j].Offset
	})
	model.StaticShadow = shadow
}
