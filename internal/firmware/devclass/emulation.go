package devclass

// EmulationLevel describes the deepest device-facing behavior implemented for a
// donor family. Levels are ordered from identity-only fallback to host-memory
// DMA and interrupt behavior.
type EmulationLevel string

const (
	EmulationIdentity  EmulationLevel = "identity"
	EmulationRegisters EmulationLevel = "registers"
	EmulationBehavior  EmulationLevel = "behavior"
	EmulationDMA       EmulationLevel = "dma"
)

// EmulationSupport is the machine-readable support contract for one donor.
// Validated means that the stated level has executable regression coverage; it
// does not imply that every command or proprietary firmware path is emulated.
type EmulationSupport struct {
	Family      string         `json:"family"`
	Level       EmulationLevel `json:"level"`
	Validated   bool           `json:"validated"`
	Limitations []string       `json:"limitations,omitempty"`
}

// Complete reports whether the device has validated DMA-level behavior without
// a known functional limitation. Most real hardware families intentionally do
// not meet this bar yet; callers must not infer completeness from PCI class.
func (s EmulationSupport) Complete() bool {
	return s.Level == EmulationDMA && s.Validated && len(s.Limitations) == 0
}

// AssessEmulation identifies the implemented family and current support level.
// Device-family dispatch is deliberately vendor/device aware where register or
// descriptor formats differ within one PCI class.
func AssessEmulation(classCode uint32, vendorID, deviceID uint16, hasMSIX bool) EmulationSupport {
	baseClass := byte(classCode >> 16)
	subClass := byte(classCode >> 8)
	progIF := byte(classCode)

	switch {
	case baseClass == 0x01 && subClass == 0x08 && progIF == 0x02:
		return EmulationSupport{
			Family:    "nvme",
			Level:     EmulationDMA,
			Validated: true,
			Limitations: []string{
				"admin and bounded BRAM-backed I/O paths only; arbitrary vendor commands are not emulated",
			},
		}
	case baseClass == 0x02 && subClass == 0x00 && vendorID == 0x8086 && deviceID == 0x15B7:
		if hasMSIX {
			return EmulationSupport{
				Family:      "intel-e1000e-i219",
				Level:       EmulationRegisters,
				Validated:   false,
				Limitations: []string{"descriptor DMA is disabled because the Ethernet MSI-X delivery path is not implemented"},
			}
		}
		return EmulationSupport{
			Family:    "intel-e1000e-i219",
			Level:     EmulationDMA,
			Validated: true,
			Limitations: []string{
				"single legacy descriptor queue, MSI-only interrupts, and 2048-byte receive buffers",
			},
		}
	case baseClass == 0x02 && subClass == 0x00 && vendorID == 0x10EC && deviceID == 0x8125:
		return EmulationSupport{
			Family:      "realtek-rtl8125",
			Level:       EmulationRegisters,
			Validated:   true,
			Limitations: []string{"link/register startup only; Realtek descriptor DMA and firmware paths are not emulated"},
		}
	case baseClass == 0x02 && subClass == 0x00:
		return EmulationSupport{
			Family:      "ethernet-generic",
			Level:       EmulationIdentity,
			Validated:   false,
			Limitations: []string{"unknown Ethernet register and descriptor family; donor-backed static fallback only"},
		}
	case baseClass == 0x0C && subClass == 0x03 && progIF == 0x30:
		return EmulationSupport{
			Family:      "xhci-generic",
			Level:       EmulationRegisters,
			Validated:   false,
			Limitations: []string{"command ring, transfer rings, event ring DMA, slots, endpoints, and interrupts are not emulated"},
		}
	case baseClass == 0x04 && subClass == 0x03:
		return EmulationSupport{
			Family:    "hda-generic",
			Level:     EmulationDMA,
			Validated: true,
			Limitations: []string{
				"basic codec verbs and RIRB DMA only; stream DMA and vendor codec extensions are not emulated",
			},
		}
	case baseClass == 0x01 && subClass == 0x06:
		return EmulationSupport{
			Family:      "ahci-generic",
			Level:       EmulationRegisters,
			Validated:   false,
			Limitations: []string{"command-list/FIS DMA and ATA command execution are not emulated"},
		}
	case baseClass == 0x02 && subClass == 0x80:
		family := "wifi-generic"
		if vendorID == 0x14C3 {
			family = "wifi-mediatek"
		}
		return EmulationSupport{
			Family:      family,
			Level:       EmulationRegisters,
			Validated:   false,
			Limitations: []string{"firmware upload, mailbox, DMA queues, radio, and interrupt behavior are not emulated"},
		}
	case baseClass == 0x03:
		return EmulationSupport{
			Family:      "gpu-generic",
			Level:       EmulationRegisters,
			Validated:   false,
			Limitations: []string{"VBIOS, engines, memory controller, display, command processors, and interrupts are not emulated"},
		}
	case baseClass == 0x0C && subClass == 0x80:
		return EmulationSupport{
			Family:      "thunderbolt-nhi",
			Level:       EmulationRegisters,
			Validated:   false,
			Limitations: []string{"NHI rings, mailboxes, topology discovery, DMA, and interrupts are not emulated"},
		}
	default:
		return EmulationSupport{Family: "generic", Level: EmulationIdentity, Validated: true}
	}
}
