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
	ClassGeneric     = "generic"
)

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
		return &wifiStrategy{baseStrategy{"Wi-Fi", ClassWiFi, wifiProfile}}
	case baseClass == 0x0C && subClass == 0x80:
		return &thunderboltStrategy{baseStrategy{"Thunderbolt", ClassThunderbolt, thunderboltProfile}}
	default:
		return &genericStrategy{baseStrategy{"Generic", ClassGeneric, genericProfile}}
	}
}

// --- Generic fallback ---

type genericStrategy struct{ baseStrategy }

func (s *genericStrategy) ScrubBAR(data []byte) {}

func (s *genericStrategy) PostInitRegisters(regs map[uint32]*uint32) {}
