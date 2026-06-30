// Package ahci builds the ATA IDENTIFY DEVICE data an AHCI/SATA clone returns
// to storahci. A real block (model, serial, LBA count, feature words) is what
// lets the disk enumerate instead of failing init.
package ahci

import (
	"fmt"
	"strings"
)

// BuildIdentify returns the 256-word ATA IDENTIFY DEVICE block for a SATA disk
// of the given size in 512-byte sectors. Strings are byte-swapped per the ATA
// word convention.
func BuildIdentify(model, serial, fwRev string, sectors uint64) [256]uint16 {
	var w [256]uint16

	w[0] = 0x0040 // ATA device, non-removable
	putATAString(&w, 10, 10, serial)
	putATAString(&w, 23, 4, fwRev)
	putATAString(&w, 27, 20, model)
	w[47] = 0x8000 | 0x10 // max sectors per READ/WRITE MULTIPLE
	w[49] = 0x0300        // LBA + DMA supported
	w[50] = 0x4000
	w[53] = 0x0006 // words 70:64 and 88 valid

	// 28-bit total sectors (capped); 48-bit in words 100-103.
	lba28 := sectors
	if lba28 > 0x0FFFFFFF {
		lba28 = 0x0FFFFFFF
	}
	w[60] = uint16(lba28)
	w[61] = uint16(lba28 >> 16)

	w[64] = 0x0003 // PIO modes 3,4
	w[80] = 0x01F0 // ATA-4..8
	w[81] = 0x0000
	w[82] = 0x0000
	w[83] = 0x4400 // LBA48 + flush-ext supported
	w[84] = 0x4000
	w[85] = 0x0000
	w[86] = 0x4400 // LBA48 + flush-ext enabled
	w[87] = 0x4000
	w[88] = 0x203F // UDMA modes 0-5, mode 5 selected

	w[100] = uint16(sectors)
	w[101] = uint16(sectors >> 16)
	w[102] = uint16(sectors >> 32)
	w[103] = uint16(sectors >> 48)

	w[106] = 0x4000
	w[217] = 0x0001 // non-rotating (SSD)

	// Integrity word: signature 0xA5 in low byte, checksum in high byte so the
	// two's-complement sum of all 512 bytes is zero.
	w[255] = 0x00A5
	var sum uint8
	for i := 0; i < 255; i++ {
		sum += uint8(w[i]) + uint8(w[i]>>8)
	}
	sum += 0xA5
	w[255] |= uint16(uint8(-int8(sum))) << 8

	return w
}

// IdentifyHex renders the block as 128 little-endian 32-bit words, one hex
// dword per line, for $readmemh into the AHCI engine's identify ROM.
func IdentifyHex(w [256]uint16) string {
	var b strings.Builder
	for i := 0; i < 256; i += 2 {
		dw := uint32(w[i]) | uint32(w[i+1])<<16
		fmt.Fprintf(&b, "%08x\n", dw)
	}
	return b.String()
}

func putATAString(w *[256]uint16, start, words int, s string) {
	buf := make([]byte, words*2)
	for i := range buf {
		buf[i] = ' '
	}
	copy(buf, []byte(s))
	for i := 0; i < words; i++ {
		w[start+i] = uint16(buf[2*i])<<8 | uint16(buf[2*i+1])
	}
}
