package scrub

import (
	"github.com/sercanarga/pcileechgen/internal/firmware/overlay"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

// vendorQuirk ties a vendor+class combo to a fixup.
// ClassCode 0 = match any class.
type vendorQuirk struct {
	VendorID  uint16
	ClassCode uint32
	Name      string
	Apply     func(cs *pci.ConfigSpace, om *overlay.Map)
}

var vendorQuirks = []vendorQuirk{
	{
		VendorID:  0x1912,
		ClassCode: 0x0C0330,
		Name:      "Renesas xHCI FW status",
		Apply:     fixRenesasFirmwareStatus,
	},
	// add new vendors here
}

func applyVendorQuirks(cs *pci.ConfigSpace, om *overlay.Map) {
	vid := cs.VendorID()
	cc := cs.ClassCode()
	for _, q := range vendorQuirks {
		if q.VendorID != vid {
			continue
		}
		if q.ClassCode != 0 && q.ClassCode != cc {
			continue
		}
		q.Apply(cs, om)
	}
}

// Renesas uPD720201/202: mark FW as loaded.
// Without this the driver starts a FW download handshake -> Code 10.
func fixRenesasFirmwareStatus(cs *pci.ConfigSpace, om *overlay.Map) {
	const (
		fwStatusOff   = 0xF4
		romStatusOff  = 0xF6
		fwSuccess     = 0x10 // bit 4
		fwLock        = 0x80 // bit 7
		romResultMask = 0x0070
	)

	om.WriteU8(fwStatusOff, fwSuccess|fwLock, "Renesas: mark FW as loaded")

	rs := cs.ReadU16(romStatusOff)
	rs = (rs &^ uint16(romResultMask)) | uint16(fwSuccess)
	om.WriteU16(romStatusOff, rs, "Renesas: set ROM status success")
}
