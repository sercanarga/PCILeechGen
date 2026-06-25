package output

import (
	"github.com/sercanarga/pcileechgen/internal/firmware/svgen"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

func extractDonorCapabilities(cs *pci.ConfigSpace) svgen.DonorCapabilities {
	if cs == nil {
		return svgen.DonorCapabilities{}
	}

	var caps svgen.DonorCapabilities
	for _, cap := range pci.ParseCapabilities(cs) {
		switch cap.ID {
		case pci.CapIDPowerManagement:
			caps.HasPMCap = true
			if legacyFieldFits(cap.Offset+2, 2) {
				pmc := cs.ReadU16(cap.Offset + 2)
				caps.PMESupportMask = uint8((pmc >> 11) & 0x1F)
			}
			if legacyFieldFits(cap.Offset+4, 2) {
				caps.PMEDefault = cs.ReadU16(cap.Offset+4)&0x0100 != 0
			}
		case pci.CapIDMSI:
			caps.HasMSICap = true
			if legacyFieldFits(cap.Offset+2, 2) {
				msgCtl := cs.ReadU16(cap.Offset + 2)
				caps.MSIDisable64Bit = msgCtl&0x0080 == 0
				caps.MSIMultipleMsg = uint8((msgCtl >> 4) & 0x07)
			}
		case pci.CapIDMSIX:
			caps.HasMSIXCap = true
		case pci.CapIDPCIExpress:
			caps.HasPCIeCap = true
			if legacyFieldFits(cap.Offset+0x0C, 4) {
				linkCap := cs.ReadU32(cap.Offset + 0x0C)
				caps.PCIELinkSpeed = uint8(linkCap & 0x0F)
				caps.PCIELinkWidth = uint8((linkCap >> 4) & 0x3F)
				caps.PCIeASPMCap = uint8((linkCap >> 10) & 0x03)
			}
			if legacyFieldFits(cap.Offset+0x10, 2) {
				linkCtl := cs.ReadU16(cap.Offset + 0x10)
				caps.PCIeASPMEnable = uint8(linkCtl & 0x3)
			}
		}
	}

	for _, cap := range pci.ParseExtCapabilities(cs) {
		switch cap.ID {
		case pci.ExtCapIDAER:
			caps.HasAERCap = true
		case pci.ExtCapIDLTR:
			caps.HasLTRCap = true
		case pci.ExtCapIDL1PMSubstates:
			caps.HasL1PMSubstates = true
		case pci.ExtCapIDDeviceSerialNumber:
			caps.HasDSNCap = true
		}
	}

	return caps
}

func legacyFieldFits(offset int, width int) bool {
	return offset >= 0 && width > 0 && offset+width <= pci.ConfigSpaceLegacySize
}
