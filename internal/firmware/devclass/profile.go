package devclass

// DeviceProfile describes what a given device class normally looks like.
type DeviceProfile struct {
	ClassName string

	PreferredBAR      int
	MinBARSize        int
	Uses64BitBAR      bool
	BARIsPrefetchable bool

	PrefersMSIX    bool
	MinMSIXVectors int

	ExpectedCaps    []uint8
	ExpectedExtCaps []uint16

	SupportsPME   bool
	MaxPowerState uint8

	BARDefaults []BARDefault

	Notes string
}

// BARDefault is one register the driver expects in the BAR.
type BARDefault struct {
	Offset uint32
	Width  int
	Name   string
	Reset  uint32
	RWMask uint32
}

// ProfileForClass returns a profile for the class code, or a generic fallback.
func ProfileForClass(classCode uint32) *DeviceProfile {
	baseClass := (classCode >> 16) & 0xFF
	subClass := (classCode >> 8) & 0xFF

	switch {
	case baseClass == 0x01 && subClass == 0x08:
		return nvmeProfile()
	case baseClass == 0x0C && subClass == 0x03:
		return xhciProfile()
	case baseClass == 0x02 && subClass == 0x00:
		return ethernetProfile()
	case baseClass == 0x04 && subClass == 0x03:
		return audioProfile()
	case baseClass == 0x03 && subClass == 0x00:
		return gpuProfile()
	case baseClass == 0x01 && subClass == 0x06:
		return sataProfile()
	case baseClass == 0x02 && subClass == 0x80:
		return wifiProfile()
	case baseClass == 0x0C && subClass == 0x80:
		return thunderboltProfile()
	default:
		return genericProfile()
	}
}

// AllProfiles lists every known device profile.
func AllProfiles() []*DeviceProfile {
	return []*DeviceProfile{
		nvmeProfile(),
		xhciProfile(),
		ethernetProfile(),
		audioProfile(),
		gpuProfile(),
		sataProfile(),
		wifiProfile(),
		thunderboltProfile(),
		genericProfile(),
	}
}

func genericProfile() *DeviceProfile {
	return &DeviceProfile{
		ClassName:         "Generic",
		PreferredBAR:      0,
		MinBARSize:        4096,
		Uses64BitBAR:      false,
		BARIsPrefetchable: false,
		PrefersMSIX:       false,
		MinMSIXVectors:    0,
		SupportsPME:       true,
		MaxPowerState:     3,
		Notes:             "Fallback profile for unrecognized device classes.",
	}
}
