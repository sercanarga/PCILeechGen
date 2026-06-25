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

	// PrefersINTx is set for legacy devices whose drivers bind to legacy INTx
	// line interrupts. Codegen does not auto-advertise INTx until board wiring
	// exists.
	PrefersINTx bool

	ExpectedCaps    []uint8
	ExpectedExtCaps []uint16

	SupportsPME   bool
	MaxPowerState uint8

	BARDefaults []BARDefault

	Notes string
}

// BARDefault is one register the driver expects in the BAR.
type BARDefault struct {
	Offset      uint32
	Width       int
	Name        string
	Reset       uint32
	RWMask      uint32
	IsRW1C      bool
	IsFSMDriven bool // driven by dedicated FSM, excluded from generic reset/write
	// SequentialRead marks a register whose read value advances through a
	// ROM/counter on each read to the same address (stateful reads). Used to
	// model devices whose drivers expect different data on repeated reads of
	// the same offset (e.g. Atheros EEPROM handshakes). When true the generated
	// read logic routes the offset through a sequential-read ROM+counter in
	// bar_impl_device instead of the static register file.
	SequentialRead bool
	// SequentialValues optionally provides the exact values returned on
	// successive reads before wrapping back to index 0. When empty, the
	// generated logic falls back to Reset + a small wrapping counter.
	SequentialValues []uint32
}

// ProfileForClass returns a profile for the class code, or a generic fallback.
// Delegates to StrategyForClass to avoid duplicating the class dispatch logic.
func ProfileForClass(classCode uint32) *DeviceProfile {
	return StrategyForClass(classCode).Profile()
}

// ProfileForClassAndVendor returns a profile using vendor-specific class
// dispatch where available.
func ProfileForClassAndVendor(classCode uint32, vendorID uint16) *DeviceProfile {
	return StrategyForClassAndVendor(classCode, vendorID).Profile()
}

// ProfileForDevice returns a profile using vendor/device-specific dispatch where
// available.
func ProfileForDevice(classCode uint32, vendorID, deviceID uint16) *DeviceProfile {
	return StrategyForDevice(classCode, vendorID, deviceID).Profile()
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
		intelWiFiProfile(),
		ar9287Profile(),
		realtekRTL8188EEProfile(),
		rts522aProfile(),
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
