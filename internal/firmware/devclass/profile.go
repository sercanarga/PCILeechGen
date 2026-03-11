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
// Delegates to StrategyForClass to avoid duplicating the class dispatch logic.
func ProfileForClass(classCode uint32) *DeviceProfile {
	return StrategyForClass(classCode).Profile()
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
		mediatekWifiProfile(),
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
