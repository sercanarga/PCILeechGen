package devclass

import "github.com/sercanarga/pcileechgen/internal/pci"

func sataProfile() *DeviceProfile {
	return &DeviceProfile{
		ClassName:         "SATA AHCI",
		PreferredBAR:      5, // AHCI spec: ABAR is BAR5
		MinBARSize:        4096,
		Uses64BitBAR:      false,
		BARIsPrefetchable: false,

		PrefersMSIX:    false,
		MinMSIXVectors: 0,

		ExpectedCaps: []uint8{
			pci.CapIDPowerManagement,
			pci.CapIDMSI,
			pci.CapIDPCIExpress,
			pci.CapIDSATADataIndex,
		},
		ExpectedExtCaps: []uint16{
			pci.ExtCapIDAER,
		},

		SupportsPME:   true,
		MaxPowerState: 3,

		BARDefaults: []BARDefault{
			// Generic Host Control
			{Offset: 0x00, Width: 4, Name: "CAP", Reset: 0x40341F05, RWMask: 0x00000000},
			// GHC — AE=1 (AHCI mode), IE=0 (interrupts off for FPGA)
			{Offset: 0x04, Width: 4, Name: "GHC", Reset: 0x80000000, RWMask: 0x80000003},
			// IS — interrupt status (all clear)
			{Offset: 0x08, Width: 4, Name: "IS", Reset: 0x00000000, RWMask: 0xFFFFFFFF},
			// PI — ports implemented (port 0 only)
			{Offset: 0x0C, Width: 4, Name: "PI", Reset: 0x00000001, RWMask: 0x00000000},
			// VS — AHCI version 1.3.1
			{Offset: 0x10, Width: 4, Name: "VS", Reset: 0x00010301, RWMask: 0x00000000},
			// CAP2 — APST support
			{Offset: 0x24, Width: 4, Name: "CAP2", Reset: 0x00000004, RWMask: 0x00000000},
			// Port 0: PxCLB
			{Offset: 0x100, Width: 4, Name: "PxCLB", Reset: 0x00000000, RWMask: 0xFFFFFC00},
			// Port 0: PxFB
			{Offset: 0x108, Width: 4, Name: "PxFB", Reset: 0x00000000, RWMask: 0xFFFFFF00},
			// Port 0: PxIS
			{Offset: 0x110, Width: 4, Name: "PxIS", Reset: 0x00000000, RWMask: 0xFDC000AF},
			// Port 0: PxCI — no commands pending
			{Offset: 0x138, Width: 4, Name: "PxCI", Reset: 0x00000000, RWMask: 0xFFFFFFFF},
			// Port 0: PxSSTS — device detected, phy established
			{Offset: 0x128, Width: 4, Name: "PxSSTS", Reset: 0x00000113, RWMask: 0x00000000},
			// Port 0: PxSIG — ATA signature (disk)
			{Offset: 0x124, Width: 4, Name: "PxSIG", Reset: 0x00000101, RWMask: 0x00000000},
		},

		Notes: "AHCI 1.3.1 profile. GHC.AE=1 for AHCI mode. " +
			"Port 0 shows device detected + phy ready via PxSSTS. " +
			"ABAR is at BAR5 per AHCI spec.",
	}
}
