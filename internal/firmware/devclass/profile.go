// Package devclass holds per-class hardware profiles (NVMe, xHCI, Ethernet, etc.).
package devclass

// DeviceProfile describes what a given device class normally looks like.
type DeviceProfile struct {
	ClassName string // human-readable name

	// BAR layout expectations
	PreferredBAR      int  // which BAR index the driver expects (typically 0)
	MinBARSize        int  // minimum BAR size the driver requires
	Uses64BitBAR      bool // whether the BAR is 64-bit addressable
	BARIsPrefetchable bool // memory BAR should be prefetchable

	// Interrupt model
	PrefersMSIX    bool // driver prefers MSI-X over MSI
	MinMSIXVectors int  // minimum MSI-X table entries the driver expects

	// Expected PCI capabilities (legacy, 0x00-0xFF)
	ExpectedCaps []uint8

	// Expected PCIe extended capabilities (0x100+)
	ExpectedExtCaps []uint16

	// Power management
	SupportsPME   bool  // device should advertise PME support
	MaxPowerState uint8 // deepest D-state supported (0=D0 only, 3=D3hot)

	// BAR register defaults — class-specific reset values
	BARDefaults []BARDefault

	// Device-class-specific notes
	Notes string
}

// BARDefault is one register the driver expects in the BAR.
type BARDefault struct {
	Offset uint32 // byte offset within the BAR
	Width  int    // register width in bytes (1, 2, or 4)
	Name   string // human-readable register name
	Reset  uint32 // expected reset (power-on) value
	RWMask uint32 // writable bits (0 = read-only)
}

// ProfileForClass returns a profile for the class code, or nil if unknown.
func ProfileForClass(classCode uint32) *DeviceProfile {
	baseClass := (classCode >> 16) & 0xFF
	subClass := (classCode >> 8) & 0xFF

	switch {
	case baseClass == 0x01 && subClass == 0x08: // NVMe
		return nvmeProfile()
	case baseClass == 0x0C && subClass == 0x03: // USB (xHCI, EHCI, etc.)
		return xhciProfile()
	case baseClass == 0x02 && subClass == 0x00: // Ethernet
		return ethernetProfile()
	case baseClass == 0x04 && subClass == 0x03: // HD Audio
		return audioProfile()
	default:
		return nil
	}
}

// AllProfiles lists every known profile.
func AllProfiles() []*DeviceProfile {
	return []*DeviceProfile{
		nvmeProfile(),
		xhciProfile(),
		ethernetProfile(),
		audioProfile(),
	}
}
