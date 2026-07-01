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
	// ponytail: KNOWN_ISSUES also names ASMedia (0x1B21, e.g. ASM1042A/ASM3142)
	// and VIA (0x1106, e.g. VL805/VL806) as xHCI Code-10 offenders, but neither
	// gets an entry here. Checked upstream drivers/usb/host/xhci-pci.c and
	// pci-quirks.c: every real quirk for both vendors (XHCI_NO_64BIT_SUPPORT,
	// XHCI_RESET_ON_RESUME, XHCI_ASMEDIA_MODIFY_FLOWCONTROL, XHCI_LPM_SUPPORT,
	// XHCI_TRB_OVERFETCH, ...) is a driver-side runtime behavior flag with no
	// PCI config-space counterpart - unlike Renesas's FW-status bits, there is no
	// static "mark as done" register to poke. The one thing ASMedia does do in
	// config space (usb_asmedia_modifyflowcontrol, regs 0xE0/0xF8/0xFC) is a live
	// read-poll-write mailbox to real silicon that a static overlay write can't
	// reproduce, and it fixes Ethernet-over-USB flow control, not enumeration -
	// unrelated to Code 10. VIA VL805's only known FW handshake is an out-of-band
	// SPI/SoC-mailbox load (Raspberry Pi's rpi_firmware_init_vl805), nothing in
	// PCI config space either. No grounded register write found for either
	// vendor - add one here if a real config-space fixup ever surfaces.
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
