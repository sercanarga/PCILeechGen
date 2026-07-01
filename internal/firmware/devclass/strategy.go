package devclass

const (
	ClassNVMe        = "nvme"
	ClassXHCI        = "xhci"
	ClassEthernet    = "ethernet"
	ClassAudio       = "audio"
	ClassGPU         = "gpu"
	ClassSATA        = "sata"
	ClassWiFi        = "wifi"
	ClassThunderbolt = "thunderbolt"
	ClassSDHCI       = "sdhci"
	ClassGeneric     = "generic"
)

// AllClasses returns all supported device class strings.
func AllClasses() []string {
	return []string{
		ClassNVMe, ClassXHCI, ClassEthernet, ClassAudio, ClassGPU,
		ClassSATA, ClassWiFi, ClassThunderbolt, ClassSDHCI, ClassGeneric,
	}
}

// DeviceStrategy centralizes device-class-specific behavior.
type DeviceStrategy interface {
	ClassName() string
	DeviceClass() string
	Profile() *DeviceProfile
	ScrubBAR(data []byte)
	PostInitRegisters(regs map[uint32]*uint32)
}

type baseStrategy struct {
	className   string
	deviceClass string
	profileFn   func() *DeviceProfile
}

func (s *baseStrategy) ClassName() string       { return s.className }
func (s *baseStrategy) DeviceClass() string     { return s.deviceClass }
func (s *baseStrategy) Profile() *DeviceProfile { return s.profileFn() }

// StrategyForClass returns a strategy for the given PCI class code.
// Returns a generic fallback for unrecognized classes.
func StrategyForClass(classCode uint32) DeviceStrategy {
	return StrategyForClassAndVendor(classCode, 0)
}

// StrategyForClassAndVendor returns a strategy based on class code and vendor ID.
// vendor-specific dispatch is used for classes with multiple vendors (e.g. WiFi).
func StrategyForClassAndVendor(classCode uint32, vendorID uint16) DeviceStrategy {
	baseClass := (classCode >> 16) & 0xFF
	subClass := (classCode >> 8) & 0xFF
	progIF := classCode & 0xFF

	switch {
	case baseClass == 0x01 && subClass == 0x08:
		return &nvmeStrategy{baseStrategy{"NVMe", ClassNVMe, nvmeProfile}}
	case baseClass == 0x0C && subClass == 0x03 && progIF == 0x30:
		return &xhciStrategy{baseStrategy{"xHCI", ClassXHCI, xhciProfile}}
	case baseClass == 0x02 && subClass == 0x00:
		return ethernetStrategyForVendor(vendorID)
	case baseClass == 0x04 && subClass == 0x03:
		return &audioStrategy{baseStrategy{"HD Audio", ClassAudio, audioProfile}}
	case baseClass == 0x03 && subClass == 0x00,
		baseClass == 0x03 && subClass == 0x02 && progIF == 0x00:
		return gpuStrategyForVendor(vendorID)
	case baseClass == 0x01 && subClass == 0x06:
		return &sataStrategy{baseStrategy{"SATA AHCI", ClassSATA, sataProfile}}
	case baseClass == 0x02 && subClass == 0x80:
		return wifiStrategyForVendor(vendorID)
	case baseClass == 0x0C && subClass == 0x80:
		return &thunderboltStrategy{baseStrategy{"Thunderbolt", ClassThunderbolt, thunderboltProfile}}
	case baseClass == 0x08 && subClass == 0x05:
		return &sdhciStrategy{baseStrategy{"SD Host Controller", ClassSDHCI, sdhciProfile}}
	default:
		return &genericStrategy{baseStrategy{"Generic", ClassGeneric, genericProfile}}
	}
}

// wifiStrategyForVendor selects between mediatek and intel/broadcom wifi profiles.
// WWAN/cellular modem vendors under class 0x0280 don't share the iwlwifi-shaped
// register layout the wifi profiles assume, so they get routed to the safe
// generic passthrough instead of wrong register writes.
func wifiStrategyForVendor(vendorID uint16) DeviceStrategy {
	switch vendorID {
	case 0x1199, 0x17CB, 0x2CB7, 0x2C7C: // Sierra Wireless, Qualcomm, Fibocom, Quectel
		return &genericStrategy{baseStrategy{"Generic", ClassGeneric, genericProfile}}
	}
	if vendorID == mediatekVID {
		return &mediatekWifiStrategy{baseStrategy{"Wi-Fi (MediaTek)", ClassWiFi, mediatekWifiProfile}}
	}
	return &wifiStrategy{baseStrategy{"Wi-Fi", ClassWiFi, wifiProfile}}
}

// gpuStrategyForVendor selects a vendor-specific GPU strategy. NVIDIA is the
// only one with real register evidence today; AMD/Intel/generic bodies live
// in gpu.go.
func gpuStrategyForVendor(vendorID uint16) DeviceStrategy {
	switch vendorID {
	case 0, 0x10DE: // unknown (no vendor context) or NVIDIA
		return &gpuStrategy{baseStrategy{"GPU", ClassGPU, gpuProfile}}
	case 0x1002: // AMD
		return gpuStrategyForAMD()
	case 0x8086: // Intel
		return gpuStrategyForIntel()
	default:
		return gpuStrategyGeneric()
	}
}

// ethernetStrategyForVendor selects a vendor-specific ethernet strategy.
// Realtek is the one with real register evidence today; the generic body
// lives in ethernet.go.
func ethernetStrategyForVendor(vendorID uint16) DeviceStrategy {
	if vendorID == 0 || vendorID == 0x10EC { // unknown (no vendor context) or Realtek
		return &ethernetStrategy{baseStrategy{"Ethernet", ClassEthernet, ethernetProfile}}
	}
	return ethernetStrategyGeneric()
}

// --- Generic fallback ---

// genericStrategy is the deliberate default for PCI base classes with no
// PCI-SIG-standardized common register layout (e.g. 0x09 input controllers,
// 0x0A docking stations, 0x10 encryption accelerators, 0x11 signal
// processors). Donor silicon in these classes is vendor-proprietary and
// heterogeneous, unlike NVMe/xHCI/AHCI/SDHCI, so raw donor-snapshot
// passthrough is the correct, honest default rather than fabricating
// unverifiable fake registers. This is a scope boundary, not an oversight.
type genericStrategy struct{ baseStrategy }

func (s *genericStrategy) ScrubBAR(data []byte) {}

func (s *genericStrategy) PostInitRegisters(regs map[uint32]*uint32) {}
