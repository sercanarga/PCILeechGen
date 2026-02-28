package pci

// Standard PCI Capability IDs
const (
	CapIDPowerManagement   uint8 = 0x01
	CapIDAGP               uint8 = 0x02
	CapIDVPD               uint8 = 0x03
	CapIDSlotID            uint8 = 0x04
	CapIDMSI               uint8 = 0x05
	CapIDCompactPCIHotSwap uint8 = 0x06
	CapIDPCIX              uint8 = 0x07
	CapIDHyperTransport    uint8 = 0x08
	CapIDVendorSpecific    uint8 = 0x09
	CapIDDebugPort         uint8 = 0x0A
	CapIDCompactPCI        uint8 = 0x0B
	CapIDPCIHotPlug        uint8 = 0x0C
	CapIDBridgeSubsysVID   uint8 = 0x0D
	CapIDAGP8x             uint8 = 0x0E
	CapIDSecureDevice      uint8 = 0x0F
	CapIDPCIExpress        uint8 = 0x10
	CapIDMSIX              uint8 = 0x11
	CapIDSATADataIndex     uint8 = 0x12
	CapIDAdvancedFeatures  uint8 = 0x13
	CapIDEnhancedAlloc     uint8 = 0x14
	CapIDFlatteningPortal  uint8 = 0x15
)

// Extended PCI Capability IDs (PCIe extended config space)
const (
	ExtCapIDAER                uint16 = 0x0001
	ExtCapIDVCNoMFVC           uint16 = 0x0002
	ExtCapIDDeviceSerialNumber uint16 = 0x0003
	ExtCapIDPowerBudgeting     uint16 = 0x0004
	ExtCapIDRCLinkDeclaration  uint16 = 0x0005
	ExtCapIDRCInternalLinkCtl  uint16 = 0x0006
	ExtCapIDRCEventCollector   uint16 = 0x0007
	ExtCapIDMFVC               uint16 = 0x0008
	ExtCapIDVC                 uint16 = 0x0009
	ExtCapIDRCRB               uint16 = 0x000A
	ExtCapIDVendorSpecific     uint16 = 0x000B
	ExtCapIDCAC                uint16 = 0x000C
	ExtCapIDACS                uint16 = 0x000D
	ExtCapIDARI                uint16 = 0x000E
	ExtCapIDATS                uint16 = 0x000F
	ExtCapIDSRIOV              uint16 = 0x0010
	ExtCapIDMRIOV              uint16 = 0x0011
	ExtCapIDMulticast          uint16 = 0x0012
	ExtCapIDPageRequest        uint16 = 0x0013
	ExtCapIDResizableBAR       uint16 = 0x0015
	ExtCapIDDPA                uint16 = 0x0016
	ExtCapIDTPHRequester       uint16 = 0x0017
	ExtCapIDLTR                uint16 = 0x0018
	ExtCapIDSecondaryPCIe      uint16 = 0x0019
	ExtCapIDPMUX               uint16 = 0x001A
	ExtCapIDPASID              uint16 = 0x001B
	ExtCapIDLNR                uint16 = 0x001C
	ExtCapIDDPC                uint16 = 0x001D
	ExtCapIDL1PMSubstates      uint16 = 0x001E
	ExtCapIDPTM                uint16 = 0x001F
)

// Capability represents a standard PCI capability in the capability list.
type Capability struct {
	ID     uint8  `json:"id"`
	Offset int    `json:"offset"`
	Data   []byte `json:"data"`
}

// ExtCapability represents a PCIe extended capability.
type ExtCapability struct {
	ID      uint16 `json:"id"`
	Version uint8  `json:"version"`
	Offset  int    `json:"offset"`
	Data    []byte `json:"data"`
}

// CapabilityName returns the human-readable name for a standard PCI capability ID.
func CapabilityName(id uint8) string {
	switch id {
	case CapIDPowerManagement:
		return "Power Management"
	case CapIDAGP:
		return "AGP"
	case CapIDVPD:
		return "Vital Product Data"
	case CapIDSlotID:
		return "Slot Identification"
	case CapIDMSI:
		return "MSI"
	case CapIDCompactPCIHotSwap:
		return "CompactPCI HotSwap"
	case CapIDPCIX:
		return "PCI-X"
	case CapIDHyperTransport:
		return "HyperTransport"
	case CapIDVendorSpecific:
		return "Vendor Specific"
	case CapIDDebugPort:
		return "Debug Port"
	case CapIDCompactPCI:
		return "CompactPCI"
	case CapIDPCIHotPlug:
		return "PCI Hot-Plug"
	case CapIDBridgeSubsysVID:
		return "Bridge Subsystem VID"
	case CapIDAGP8x:
		return "AGP 8x"
	case CapIDSecureDevice:
		return "Secure Device"
	case CapIDPCIExpress:
		return "PCI Express"
	case CapIDMSIX:
		return "MSI-X"
	case CapIDSATADataIndex:
		return "SATA Data/Index"
	case CapIDAdvancedFeatures:
		return "Advanced Features"
	case CapIDEnhancedAlloc:
		return "Enhanced Allocation"
	case CapIDFlatteningPortal:
		return "Flattening Portal Bridge"
	default:
		return "Unknown"
	}
}

// ExtCapabilityName returns the human-readable name for an extended capability ID.
func ExtCapabilityName(id uint16) string {
	switch id {
	case ExtCapIDAER:
		return "Advanced Error Reporting"
	case ExtCapIDVCNoMFVC:
		return "Virtual Channel (No MFVC)"
	case ExtCapIDDeviceSerialNumber:
		return "Device Serial Number"
	case ExtCapIDPowerBudgeting:
		return "Power Budgeting"
	case ExtCapIDRCLinkDeclaration:
		return "Root Complex Link Declaration"
	case ExtCapIDVendorSpecific:
		return "Vendor Specific"
	case ExtCapIDACS:
		return "Access Control Services"
	case ExtCapIDARI:
		return "Alternative Routing-ID Interpretation"
	case ExtCapIDATS:
		return "Address Translation Services"
	case ExtCapIDSRIOV:
		return "Single Root I/O Virtualization"
	case ExtCapIDResizableBAR:
		return "Resizable BAR"
	case ExtCapIDLTR:
		return "Latency Tolerance Reporting"
	case ExtCapIDSecondaryPCIe:
		return "Secondary PCI Express"
	case ExtCapIDL1PMSubstates:
		return "L1 PM Substates"
	case ExtCapIDPTM:
		return "Precision Time Measurement"
	case ExtCapIDDPC:
		return "Downstream Port Containment"
	case ExtCapIDPASID:
		return "Process Address Space ID"
	default:
		return "Unknown"
	}
}

// ParseCapabilities walks the standard PCI capability linked list from config space.
func ParseCapabilities(cs *ConfigSpace) []Capability {
	if !cs.HasCapabilities() {
		return nil
	}

	var caps []Capability
	visited := make(map[int]bool)

	ptr := int(cs.CapabilityPointer()) & 0xFC // must be DWORD-aligned
	for ptr != 0 && ptr < ConfigSpaceLegacySize && !visited[ptr] {
		visited[ptr] = true

		capID := cs.ReadU8(ptr)
		nextPtr := int(cs.ReadU8(ptr+1)) & 0xFC

		// Determine capability size (minimum 2 bytes for id+next)
		capSize := 2
		if nextPtr > ptr {
			capSize = nextPtr - ptr
		} else if nextPtr == 0 {
			// Last capability, extends to end of standard config space or next boundary
			capSize = ConfigSpaceLegacySize - ptr
		}

		data := make([]byte, capSize)
		copy(data, cs.Data[ptr:ptr+capSize])

		caps = append(caps, Capability{
			ID:     capID,
			Offset: ptr,
			Data:   data,
		})

		ptr = nextPtr
	}

	return caps
}

// ParseExtCapabilities walks the PCIe extended capability linked list.
func ParseExtCapabilities(cs *ConfigSpace) []ExtCapability {
	if cs.Size < ConfigSpaceSize {
		return nil
	}

	var caps []ExtCapability
	visited := make(map[int]bool)

	offset := 0x100 // Extended capabilities start at offset 0x100
	for offset >= 0x100 && offset < ConfigSpaceSize && !visited[offset] {
		visited[offset] = true

		header := cs.ReadU32(offset)
		if header == 0 || header == 0xFFFFFFFF {
			break
		}

		capID := uint16(header & 0xFFFF)
		version := uint8((header >> 16) & 0xF)
		nextOffset := int((header >> 20) & 0xFFC)

		// Determine capability size
		capSize := 4 // minimum: the header itself
		if nextOffset > offset {
			capSize = nextOffset - offset
		} else if nextOffset == 0 {
			capSize = ConfigSpaceSize - offset
		}

		data := make([]byte, capSize)
		copy(data, cs.Data[offset:offset+capSize])

		caps = append(caps, ExtCapability{
			ID:      capID,
			Version: version,
			Offset:  offset,
			Data:    data,
		})

		if nextOffset == 0 {
			break
		}
		offset = nextOffset
	}

	return caps
}
