package devclass

import (
	"github.com/sercanarga/pcileechgen/internal/pci"
	"github.com/sercanarga/pcileechgen/internal/util"
)

// raidStrategy models an LSI/Broadcom MegaRAID SAS controller (the most common
// RAID donor, VID 0x1000). The register layout follows the MFI register set
// (struct megasas_register_set in the Linux megaraid_sas driver).
//
// Driver gate: megasas_transition_to_ready() polls outbound_scratch_pad for the
// firmware state in bits [31:28]. MFI_STATE_READY = 0xB. The register geometry
// differs by controller generation (see raidGenerationForDevice), so the
// strategy is constructed from the donor device ID via newRAIDStrategy.
const (
	// Donor vendor is LSI/Broadcom (VID 0x1000); see synthetic donor template.
	mfiStateReady uint32 = 0xB0000000 // MFI_STATE_READY in bits [31:28]
	mfiStateMask  uint32 = 0xF0000000
	raidScratchG3 uint32 = 0xA8 // outbound_scratch_pad_0 (Gen3+ Fusion)
	raidScratchG2 uint32 = 0xB0 // outbound_scratch_pad   (Gen2 legacy MFI)
)

// raidGeneration captures the per-generation register geometry that actually
// affects emulated output: register-BAR width/index and the primary scratch-pad
// offset the driver polls for firmware-ready.
type raidGeneration struct {
	label        string
	uses64BitBAR bool   // Gen2 MFI register BAR is 32-bit; Gen3+ Fusion is 64-bit
	preferredBAR int    // Gen2 exposes registers in BAR0; Gen3+ in BAR1
	scratchPad   uint32 // primary outbound_scratch_pad offset for the ready poll
}

var (
	raidGenFusion = raidGeneration{"Gen3+ Fusion", true, 1, raidScratchG3}
	raidGenMFI    = raidGeneration{"Gen2 MFI", false, 0, raidScratchG2}
)

// gen2DeviceIDs lists LSI/Broadcom SAS2208-class MegaRAID controllers (legacy
// MFI register layout: 32-bit BAR0, scratch pad at 0xB0). Everything else —
// SAS3108 Invader/Fury, SAS35xx Ventura, SAS39xx Aero — uses the unified Fusion
// layout (64-bit BAR1, scratch_pad_0 at 0xA8), the safe default.
var gen2DeviceIDs = map[uint16]bool{
	0x005b: true, // SAS2208 (MegaRAID 9265/9266/9270/9271/9285/9286)
	0x0073: true, // SAS2008 (MegaRAID 9240)
}

// raidGenerationForDevice resolves the MegaRAID generation from the donor's
// device ID. Defaults to the Fusion layout for unknown/3108-class IDs.
func raidGenerationForDevice(deviceID uint16) raidGeneration {
	if gen2DeviceIDs[deviceID] {
		return raidGenMFI
	}
	return raidGenFusion
}

type raidStrategy struct {
	baseStrategy
	gen raidGeneration
}

func newRAIDStrategy(deviceID uint16) *raidStrategy {
	gen := raidGenerationForDevice(deviceID)
	return &raidStrategy{
		baseStrategy: baseStrategy{"RAID Controller", ClassRAID, raidProfile},
		gen:          gen,
	}
}

// Profile overrides baseStrategy.Profile to fold in generation-specific geometry.
func (s *raidStrategy) Profile() *DeviceProfile {
	return raidProfileForGen(s.gen)
}

func (s *raidStrategy) ScrubBAR(data []byte) {
	if len(data) < 0xB4 {
		return
	}
	// Advertise firmware READY at both scratch-pad generations so the spoof
	// works regardless of which offset the donor's driver polls.
	setMFIState(data, int(raidScratchG3))
	setMFIState(data, int(raidScratchG2))

	// Clear interrupt status and outbound doorbell (no pending events).
	util.WriteLE32(data, 0x24, 0x00) // inbound_intr_status
	util.WriteLE32(data, 0x30, 0x00) // outbound_intr_status
	util.WriteLE32(data, 0x2C, 0x00) // outbound_doorbell
}

// setMFIState forces the firmware-state nibble of a scratch pad to READY while
// preserving the lower bits (queue-depth / capability hints the driver reads).
func setMFIState(data []byte, off int) {
	v := util.ReadLE32(data, off)
	v = (v &^ mfiStateMask) | mfiStateReady
	util.WriteLE32(data, off, v)
}

// PostInitRegisters is a no-op: ScrubBAR already sets MFI_STATE_READY in
// both scratch-pad offsets as part of the BAR data scrub pass.
func (s *raidStrategy) PostInitRegisters(regs map[uint32]*uint32) {}

// raidProfile is the no-arg profile used by the registry (AllProfiles) and the
// class-only ProfileForClass path. It reflects the default Fusion generation.
func raidProfile() *DeviceProfile {
	return raidProfileForGen(raidGenFusion)
}

func raidProfileForGen(gen raidGeneration) *DeviceProfile {
	return &DeviceProfile{
		ClassName:         "RAID Controller",
		PreferredBAR:      gen.preferredBAR,
		MinBARSize:        65536, // 64 KiB MMIO register window
		Uses64BitBAR:      gen.uses64BitBAR,
		BARIsPrefetchable: false, // control registers are not prefetchable

		PrefersMSIX: true,
		// MegaRAID exposes many reply-post vectors; the exact count is cloned
		// verbatim from the donor's MSI-X capability (ctx.MSIXData.TableSize).
		// This is only the validation floor: admin + at least one reply queue.
		MinMSIXVectors: 2,

		ExpectedCaps: []uint8{
			pci.CapIDPowerManagement,
			pci.CapIDMSI,
			pci.CapIDMSIX,
			pci.CapIDPCIExpress,
		},
		ExpectedExtCaps: []uint16{
			pci.ExtCapIDAER,
		},

		SupportsPME:   true,
		MaxPowerState: 3, // D3hot

		BARDefaults: []BARDefault{
			// Handshake / messaging registers (MFI register set).
			{Offset: 0x00, Width: 4, Name: "DOORBELL", Reset: 0x00000000, RWMask: 0x00000000},
			{Offset: 0x10, Width: 4, Name: "INBOUND_MSG_0", Reset: 0x00000000, RWMask: 0xFFFFFFFF},
			{Offset: 0x18, Width: 4, Name: "OUTBOUND_MSG_0", Reset: 0x00000000, RWMask: 0x00000000},
			{Offset: 0x20, Width: 4, Name: "INBOUND_DOORBELL", Reset: 0x00000000, RWMask: 0xFFFFFFFF},
			// Interrupt status registers are write-1-to-clear.
			{Offset: 0x24, Width: 4, Name: "INBOUND_INTR_STATUS", Reset: 0x00000000, W1CMask: 0xFFFFFFFF, IsRW1C: true},
			{Offset: 0x28, Width: 4, Name: "INBOUND_INTR_MASK", Reset: 0x00000000, RWMask: 0xFFFFFFFF},
			{Offset: 0x2C, Width: 4, Name: "OUTBOUND_DOORBELL", Reset: 0x00000000, W1CMask: 0xFFFFFFFF, IsRW1C: true},
			{Offset: 0x30, Width: 4, Name: "OUTBOUND_INTR_STATUS", Reset: 0x00000000, W1CMask: 0xFFFFFFFF, IsRW1C: true},
			{Offset: 0x34, Width: 4, Name: "OUTBOUND_INTR_MASK", Reset: 0x00000000, RWMask: 0xFFFFFFFF},
			{Offset: 0x38, Width: 4, Name: "INBOUND_QUEUE_PORT", Reset: 0x00000000, RWMask: 0xFFFFFFFF},
			{Offset: 0x3C, Width: 4, Name: "OUTBOUND_QUEUE_PORT", Reset: 0x00000000, RWMask: 0x00000000},
			// FW state machine: READY in bits [31:28] so the driver proceeds.
			// Both scratch-pad generations are presented (Gen3 0xA8 + Gen2 0xB0).
			{Offset: 0xA8, Width: 4, Name: "OUTBOUND_SCRATCH_PAD_0", Reset: 0xB0000000, RWMask: 0x00000000},
			{Offset: 0xAC, Width: 4, Name: "OUTBOUND_SCRATCH_PAD_1", Reset: 0x00000000, RWMask: 0x00000000},
			{Offset: 0xB0, Width: 4, Name: "OUTBOUND_SCRATCH_PAD_2", Reset: 0xB0000000, RWMask: 0x00000000},
			{Offset: 0xB4, Width: 4, Name: "OUTBOUND_SCRATCH_PAD_3", Reset: 0x00000000, RWMask: 0x00000000},
			// Fusion 64-bit inbound queue ports.
			{Offset: 0xC0, Width: 4, Name: "INBOUND_LOW_QUEUE_PORT", Reset: 0x00000000, RWMask: 0xFFFFFFFF},
			{Offset: 0xC4, Width: 4, Name: "INBOUND_HIGH_QUEUE_PORT", Reset: 0x00000000, RWMask: 0xFFFFFFFF},
		},

		Notes: "LSI/Broadcom MegaRAID (MFI register set), " + gen.label + ". megaraid_sas " +
			"gates init on outbound_scratch_pad firmware-state == READY (0xB in bits [31:28]); " +
			"presented at 0xA8 (Fusion) and 0xB0 (legacy MFI). MSI-X vector count and table/PBA " +
			"offsets are cloned verbatim from the donor capability. Register BAR / 64-bit width " +
			"are generation-derived from the donor device ID.",
	}
}
