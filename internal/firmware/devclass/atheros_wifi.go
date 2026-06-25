package devclass

import (
	"github.com/sercanarga/pcileechgen/internal/pci"
	"github.com/sercanarga/pcileechgen/internal/util"
)

const atherosVID uint16 = 0x168C

// AR9287 EEPROM handshake constants (ath9k driver).
//
// ath9k reads the EEPROM in two phases:
//  1. Write the EEPROM word index to the EEPROM request register, then read
//     it back: the read returns a sentinel (0xDEADBEEF style) while the
//     request is pending, then a "done" value when the field is available.
//  2. Read the EEPROM data register to collect the word.
//
// We model the two registers as SequentialRead. When explicit
// SequentialValues are present, the generated SV returns the recorded sequence
// verbatim instead of the older Reset+counter fallback.
var (
	ar9287EepromReqSequence  = []uint32{0xDEADBEEF, 0x00000002, 0x00000001}
	ar9287EepromDataSequence = []uint32{0x0000A55A, 0x00000004, 0x0000FFFB, 0x0000E00E}
)

const (
	ar9287EepromReqSentinel uint32 = 0xDEADBEEF
	ar9287EepromDataBase    uint32 = 0x0000A55A
	ar9287DeviceID          uint16 = 0x002E
)

type ar9287WifiStrategy struct{ baseStrategy }

func (s *ar9287WifiStrategy) ScrubBAR(data []byte) {
	if len(data) < 0x28 {
		return
	}
	// Keep the driver-visible EEPROM handshake registers seeded with their
	// first-phase values. The generated SV read FSM advances them per read.
	util.WriteLE32(data, 0x4010, ar9287EepromReqSentinel)
	if len(data) >= 0x4080 {
		util.WriteLE32(data, 0x407C, ar9287EepromDataBase)
	}
}

func (s *ar9287WifiStrategy) PostInitRegisters(regs map[uint32]*uint32) {
	if v, ok := regs[0x4010]; ok {
		*v = ar9287EepromReqSentinel
	}
	if v, ok := regs[0x407C]; ok {
		*v = ar9287EepromDataBase
	}
}

func ar9287Profile() *DeviceProfile {
	return &DeviceProfile{
		ClassName:         "Wi-Fi (Atheros AR9287)",
		PreferredBAR:      0,
		MinBARSize:        8192,
		Uses64BitBAR:      true,
		BARIsPrefetchable: false,

		// ath9k binds to legacy INTx line interrupts on AR9287; it does not
		// use MSI/MSI-X for the EEPROM load path. Advertising MSI-X here would
		// route the IRQ through the wrong vector and the driver's INTx handler
		// would never fire.
		PrefersMSIX:    false,
		MinMSIXVectors: 0,
		PrefersINTx:    true,

		ExpectedCaps: []uint8{
			pci.CapIDPowerManagement,
			pci.CapIDMSI,
			pci.CapIDPCIExpress,
		},
		ExpectedExtCaps: []uint16{
			pci.ExtCapIDAER,
		},

		SupportsPME:   true,
		MaxPowerState: 3,

		BARDefaults: []BARDefault{
			// EEPROM request register: read returns a pending sentinel, then a
			// short done sequence. Modeled as SequentialRead with explicit values.
			{Offset: 0x4010, Width: 4, Name: "EEPROM_REQ", Reset: ar9287EepromReqSentinel, RWMask: 0x00000000, SequentialRead: true, SequentialValues: ar9287EepromReqSequence},
			// EEPROM data register: successive reads return a small donor-like
			// EEPROM word sequence instead of a synthetic +1 counter.
			{Offset: 0x407C, Width: 4, Name: "EEPROM_DATA", Reset: ar9287EepromDataBase, RWMask: 0x00000000, SequentialRead: true, SequentialValues: ar9287EepromDataSequence},
			// MAC address lower 32 bits. OUI (Atheros 00:13:74 / 00:1B:11) stays
			// fixed; the NIC-specific low bytes are randomized per-build via
			// the variance layer (free-running counter keeps OUI stable).
			{Offset: 0x4000, Width: 4, Name: "MAC_LO", Reset: 0x74001300, RWMask: 0x00000000},
			{Offset: 0x4004, Width: 4, Name: "MAC_HI", Reset: 0x00000000, RWMask: 0x00000000},
		},

		Notes: "Atheros AR9287 (ath9k) Wi-Fi profile. EEPROM load uses a 2-phase " +
			"request/response handshake: write index to EEPROM_REQ, poll for " +
			"done, read EEPROM_DATA. Both are modeled as SequentialRead so each " +
			"read advances the value (a static read makes ath9k time out). " +
			"Driver uses legacy INTx, not MSI-X. MAC OUI kept fixed; low bytes " +
			"randomized via the variance layer.",
	}
}
