package devclass

const (
	ClassNVMe        = "nvme"
	ClassXHCI        = "xhci"
	ClassEthernet    = "ethernet"
	ClassAudio       = "audio"
	ClassGPU         = "gpu"
	ClassSATA        = "sata"
	ClassWiFi        = "wifi"
	ClassCardReader  = "cardreader"
	ClassThunderbolt = "thunderbolt"
	ClassGeneric     = "generic"
)

// AllClasses returns all supported device class strings.
func AllClasses() []string {
	return []string{
		ClassNVMe, ClassXHCI, ClassEthernet, ClassAudio, ClassGPU,
		ClassSATA, ClassWiFi, ClassThunderbolt, ClassGeneric,
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
// vendor-specific dispatch is used only for devices where vendor identity is
// enough to select a safe profile.
func StrategyForClassAndVendor(classCode uint32, vendorID uint16) DeviceStrategy {
	return strategyForDevice(classCode, vendorID, 0)
}

// StrategyForDevice returns a strategy based on class, vendor and device ID.
func StrategyForDevice(classCode uint32, vendorID, deviceID uint16) DeviceStrategy {
	return strategyForDevice(classCode, vendorID, deviceID)
}

func strategyForDevice(classCode uint32, vendorID, deviceID uint16) DeviceStrategy {
	baseClass := (classCode >> 16) & 0xFF
	subClass := (classCode >> 8) & 0xFF

	switch {
	case baseClass == 0x01 && subClass == 0x08:
		return &nvmeStrategy{baseStrategy{"NVMe", ClassNVMe, nvmeProfile}}
	case baseClass == 0x0C && subClass == 0x03:
		return &xhciStrategy{baseStrategy{"xHCI", ClassXHCI, xhciProfile}}
	case baseClass == 0x02 && subClass == 0x00:
		return &ethernetStrategy{baseStrategy{"Ethernet", ClassEthernet, ethernetProfile}}
	case baseClass == 0x04 && subClass == 0x03:
		return &audioStrategy{baseStrategy{"HD Audio", ClassAudio, audioProfile}}
	case baseClass == 0x03 && subClass == 0x00:
		return &gpuStrategy{baseStrategy{"GPU", ClassGPU, gpuProfile}}
	case baseClass == 0x01 && subClass == 0x06:
		return &sataStrategy{baseStrategy{"SATA AHCI", ClassSATA, sataProfile}}
	case baseClass == 0x02 && subClass == 0x80:
		return wifiStrategyForDevice(vendorID, deviceID)
	case baseClass == 0xFF && vendorID == realtekVID && deviceID == rts522aDeviceID:
		return &rts522aCardReaderStrategy{baseStrategy{"Card Reader (RTS522A)", ClassCardReader, rts522aProfile}}
	case baseClass == 0x0C && subClass == 0x80:
		return &thunderboltStrategy{baseStrategy{"Thunderbolt", ClassThunderbolt, thunderboltProfile}}
	default:
		return &genericStrategy{baseStrategy{"Generic", ClassGeneric, genericProfile}}
	}
}

// wifiStrategyForDevice selects between mediatek, atheros, intel, exact
// Realtek-device and generic wifi profiles.
func wifiStrategyForDevice(vendorID, deviceID uint16) DeviceStrategy {
	if vendorID == mediatekVID {
		return &mediatekWifiStrategy{baseStrategy{"Wi-Fi (MediaTek)", ClassWiFi, mediatekWifiProfile}}
	}
	if vendorID == atherosVID && deviceID == ar9287DeviceID {
		return &ar9287WifiStrategy{baseStrategy{"Wi-Fi (Atheros AR9287)", ClassWiFi, ar9287Profile}}
	}
	if vendorID == intelVID && deviceID == intelBE200DeviceID {
		return &intelWifiStrategy{baseStrategy{"Wi-Fi (Intel BE200 family)", ClassWiFi, intelWiFiProfile}}
	}
	if vendorID == realtekVID && deviceID == realtekRTL8188EEDeviceID {
		return &realtekWifiStrategy{baseStrategy{"Wi-Fi (Realtek RTL8188EE)", ClassWiFi, realtekRTL8188EEProfile}}
	}
	return &wifiStrategy{baseStrategy{"Wi-Fi", ClassWiFi, wifiProfile}}
}

// --- Generic fallback ---

type genericStrategy struct{ baseStrategy }

func (s *genericStrategy) ScrubBAR(data []byte) {}

func (s *genericStrategy) PostInitRegisters(regs map[uint32]*uint32) {}
