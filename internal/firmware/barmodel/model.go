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
	Size      int
	Registers []BARRegister // ordered by Offset
}

// BuildBARModel returns a register map from probe data or spec tables.
// barSize is the final BAR0 aperture the caller intends to use (e.g. the
// board-BRAM-capped size from firmware.CappedBAR0Size). When it exceeds the
// data-derived size, the returned model's Size is raised to it before
// validation runs, so a donor loaded without raw BAR contents (e.g.
// --from-json without probed bytes) doesn't validate register offsets
// against a not-yet-finalized Size of 0.
// Returns nil for unknown device classes.
func BuildBARModel(barData []byte, classCode uint32, profile *donor.BARProfile, barSize int) *BARModel {
	// Use probe data when available, but bail if VFIO reported
	// everything as writable (breaks CC->CSTS handshake etc).
	if profile != nil && len(profile.Probes) > 0 {
		if !isProbeDataReliable(profile) {
			slog.Warn("BAR probe data unreliable (all registers report fully writable), falling back to spec-based model",
				"probes", len(profile.Probes))
		} else if model := SynthesizeBARModel(profile, classCode); model != nil {
			growToBARSize(model, barSize)
			return model
		}
	}

	// fall back to hardcoded spec tables
	model := specBARModelForClass(classCode, barData)
	if model != nil {
		growToBARSize(model, barSize)
		validateModel(model)
	}
	return model
}

// growToBARSize raises m.Size to barSize when barSize is the larger of the
// two, so validateModel (and downstream SV codegen) sees the real
// board-capped aperture instead of a data-derived size that may still be 0.
func growToBARSize(m *BARModel, barSize int) {
	if barSize > m.Size {
		m.Size = barSize
	}
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
	case baseClass == 0x02:
		return buildEthernetBARModel(barData)
	case baseClass == 0x04 && subClass == 0x03:
		return buildAudioBARModel(barData)
	}
	return nil
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
		// Command Ring Control - 64-bit. bits[2:0] = RCS/CS/CA (RW), bit3 = CRR
		// (RO - not modeled by the ring engine, see xhci_ring_engine.sv.tmpl),
		// bits[5:4] reserved, bits[31:6] = 64-byte-aligned pointer (RW).
		{Offset: 0x38, Width: 4, Name: "CRCR_LO", RWMask: 0xFFFFFFC7},
		{Offset: 0x3C, Width: 4, Name: "CRCR_HI", RWMask: 0xFFFFFFFF},
		// Device Context Base Address Array Pointer - 64-bit
		{Offset: 0x50, Width: 4, Name: "DCBAAP_LO", RWMask: 0xFFFFFFC0},
		{Offset: 0x54, Width: 4, Name: "DCBAAP_HI", RWMask: 0xFFFFFFFF},
		// Configure (CONFIG)
		{Offset: 0x58, Width: 4, Name: "CONFIG", RWMask: 0x000000FF},

		// Primary interrupter (Runtime Register Space, RTSOFF fixed at 0x200
		// per xhci.go's ScrubBAR/PostInitRegisters -- IR0 registers live at
		// RTSOFF+0x20). Consumed by xhci_ring_engine.sv.tmpl (Command
		// Completion Event delivery) and bar_impl_device.sv.tmpl.
		// IMAN: bit0=IP (RW1C, HW-set by the ring engine on event post),
		// bit1=IE (plain RW). IsFSMDriven because IP is hardware-settable,
		// not just software-clearable, so it needs the same hand-rolled
		// write path as USBCMD/USBSTS above.
		{Offset: 0x220, Width: 4, Name: "IMAN", RWMask: 0x00000002, W1CMask: 0x00000001, IsFSMDriven: true},
		// IMOD: interrupt moderation - accepted, not enforced (no rate limiting emulated).
		{Offset: 0x224, Width: 4, Name: "IMOD", RWMask: 0xFFFFFFFF},
		// ERSTSZ: number of segments in the Event Ring Segment Table.
		{Offset: 0x228, Width: 4, Name: "ERSTSZ", RWMask: 0x0000FFFF},
		// ERSTBA - 64-bit, 16-byte aligned pointer to the (single-segment) ERST.
		{Offset: 0x230, Width: 4, Name: "ERSTBA_LO", RWMask: 0xFFFFFFF0},
		{Offset: 0x234, Width: 4, Name: "ERSTBA_HI", RWMask: 0xFFFFFFFF},
		// ERDP - 64-bit. bit0 = EHB (Event Handler Busy); left as plain RW
		// (ponytail: the ring engine never throttles on EHB -- a handful of
		// init-time commands never overrun a driver that hasn't drained
		// events yet, so there's no real backpressure to model here).
		{Offset: 0x238, Width: 4, Name: "ERDP_LO", RWMask: 0xFFFFFFF7},
		{Offset: 0x23C, Width: 4, Name: "ERDP_HI", RWMask: 0xFFFFFFFF},
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
