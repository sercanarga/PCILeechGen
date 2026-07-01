package devclass

import (
	"github.com/sercanarga/pcileechgen/internal/pci"
	"github.com/sercanarga/pcileechgen/internal/util"
)

type gpuStrategy struct{ baseStrategy }

func (s *gpuStrategy) ScrubBAR(data []byte) {
	if len(data) < 0x204 {
		return
	}
	util.WriteLE32(data, 0x200, 0xFFFFFFFF)
}

func (s *gpuStrategy) PostInitRegisters(regs map[uint32]*uint32) {
	if v, ok := regs[0x200]; ok {
		*v = 0xFFFFFFFF
	}
}

func gpuProfile() *DeviceProfile {
	return &DeviceProfile{
		ClassName:         "GPU",
		PreferredBAR:      0,
		MinBARSize:        4096, // real VRAM BAR is huge, FPGA maps only 4K window
		Uses64BitBAR:      true,
		BARIsPrefetchable: true,

		PrefersMSIX:    true,
		MinMSIXVectors: 1,

		ExpectedCaps: []uint8{
			pci.CapIDPowerManagement,
			pci.CapIDMSIX,
			pci.CapIDPCIExpress,
		},
		ExpectedExtCaps: []uint16{
			pci.ExtCapIDAER,
			pci.ExtCapIDLTR,
			pci.ExtCapIDDeviceSerialNumber,
		},

		SupportsPME:   true,
		MaxPowerState: 3,

		BARDefaults: []BARDefault{
			// NV_PMC - boot status, master control
			{Offset: 0x00, Width: 4, Name: "PMC_BOOT", Reset: 0x00000000, RWMask: 0x00000000},
			// NV_PMC_ENABLE - engine enable bitmask
			{Offset: 0x200, Width: 4, Name: "PMC_ENABLE", Reset: 0xFFFFFFFF, RWMask: 0xFFFFFFFF},
			// NV_PBUS_PCI_NV_0 - mirrors VID/DID inside BAR
			{Offset: 0x1800, Width: 4, Name: "PBUS_PCI_NV_0", Reset: 0x00000000, RWMask: 0x00000000},
			// NV_PBUS_PCI_NV_1 - mirrors config command/status
			{Offset: 0x1804, Width: 4, Name: "PBUS_PCI_NV_1", Reset: 0x00100006, RWMask: 0x00000000},
			// NV_PTIMER_TIME_0 - low 32 bits of GPU timer
			{Offset: 0x9400, Width: 4, Name: "PTIMER_TIME_0", Reset: 0x00000000, RWMask: 0x00000000},
			// NV_PTIMER_TIME_1 - high 32 bits
			{Offset: 0x9410, Width: 4, Name: "PTIMER_TIME_1", Reset: 0x00000000, RWMask: 0x00000000},
		},

		Notes: "NVIDIA GPU profile (NV_PMC/NV_PBUS/NV_PTIMER). Only 4K BRAM window visible - " +
			"driver sees PMC_ENABLE and PTIMER but can't touch VRAM.",
	}
}

// --- AMD GPU ---

type amdGPUStrategy struct{ baseStrategy }

// ponytail: no verified AMD MMIO register map (RCC_DEV0_EPF0_STRAP0 family /
// ASIC revision regs) grounded from a real donor - safe no-op passthrough
// until a real AMD donor sample is available to reverse the layout.
func (s *amdGPUStrategy) ScrubBAR(data []byte) {}

func (s *amdGPUStrategy) PostInitRegisters(regs map[uint32]*uint32) {}

func amdGPUProfile() *DeviceProfile {
	return &DeviceProfile{
		ClassName:         "GPU (AMD)",
		PreferredBAR:      0,
		MinBARSize:        4096,
		Uses64BitBAR:      true,
		BARIsPrefetchable: true,

		PrefersMSIX:    true,
		MinMSIXVectors: 1,

		ExpectedCaps: []uint8{
			pci.CapIDPowerManagement,
			pci.CapIDMSIX,
			pci.CapIDPCIExpress,
		},
		ExpectedExtCaps: []uint16{
			pci.ExtCapIDAER,
			pci.ExtCapIDLTR,
		},

		SupportsPME:   true,
		MaxPowerState: 3,

		// ponytail: no BARDefaults - AMDGPU register offsets not grounded,
		// injecting fabricated values would be worse than leaving the BAR blank.
		Notes: "AMD GPU profile. No verified amdgpu MMIO register map - " +
			"safe passthrough (no BAR register scrub/defaults) until a real donor sample is available.",
	}
}

// --- Intel GPU ---

type intelGPUStrategy struct{ baseStrategy }

// ponytail: Intel iGPU register semantics (GTTMMADR, FORCEWAKE/FORCEWAKE_ACK
// protocol) are documented but nontrivial and not grounded from a real donor
// here - safe no-op passthrough until a real Intel donor sample is available.
func (s *intelGPUStrategy) ScrubBAR(data []byte) {}

func (s *intelGPUStrategy) PostInitRegisters(regs map[uint32]*uint32) {}

func intelGPUProfile() *DeviceProfile {
	return &DeviceProfile{
		ClassName:         "GPU (Intel)",
		PreferredBAR:      0,
		MinBARSize:        4096,
		Uses64BitBAR:      true,
		BARIsPrefetchable: true,

		PrefersMSIX:    true,
		MinMSIXVectors: 1,

		ExpectedCaps: []uint8{
			pci.CapIDPowerManagement,
			pci.CapIDMSIX,
			pci.CapIDPCIExpress,
		},
		ExpectedExtCaps: []uint16{
			pci.ExtCapIDAER,
			pci.ExtCapIDLTR,
		},

		SupportsPME:   true,
		MaxPowerState: 3,

		// ponytail: no BARDefaults - GTTMMADR/forcewake register offsets not
		// grounded, injecting fabricated values would be worse than leaving the BAR blank.
		Notes: "Intel GPU profile. No verified i915 MMIO register map - " +
			"safe passthrough (no BAR register scrub/defaults) until a real donor sample is available.",
	}
}

// --- Generic/unknown vendor GPU ---

type genericGPUStrategy struct{ baseStrategy }

// ponytail: unrecognized GPU vendor - never guess a register map, always
// safe no-op passthrough.
func (s *genericGPUStrategy) ScrubBAR(data []byte) {}

func (s *genericGPUStrategy) PostInitRegisters(regs map[uint32]*uint32) {}

func genericGPUProfile() *DeviceProfile {
	return &DeviceProfile{
		ClassName:         "GPU (Generic)",
		PreferredBAR:      0,
		MinBARSize:        4096,
		Uses64BitBAR:      true,
		BARIsPrefetchable: true,

		PrefersMSIX:    true,
		MinMSIXVectors: 1,

		SupportsPME:   true,
		MaxPowerState: 3,

		Notes: "Unrecognized GPU vendor. No register map to ground against - " +
			"safe passthrough (no BAR register scrub/defaults).",
	}
}

// gpuStrategyForAMD returns the AMD GPU strategy.
func gpuStrategyForAMD() DeviceStrategy {
	return &amdGPUStrategy{baseStrategy{"GPU (AMD)", ClassGPU, amdGPUProfile}}
}

// gpuStrategyForIntel returns the Intel GPU strategy.
func gpuStrategyForIntel() DeviceStrategy {
	return &intelGPUStrategy{baseStrategy{"GPU (Intel)", ClassGPU, intelGPUProfile}}
}

// gpuStrategyGeneric returns the neutral fallback strategy for unrecognized GPU vendors.
func gpuStrategyGeneric() DeviceStrategy {
	return &genericGPUStrategy{baseStrategy{"GPU (Generic)", ClassGPU, genericGPUProfile}}
}
