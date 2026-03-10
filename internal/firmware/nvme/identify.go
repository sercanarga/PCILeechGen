// Package nvme builds NVMe Identify Controller/Namespace 4KB responses
// for the Admin Responder FSM.
package nvme

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"

	"github.com/sercanarga/pcileechgen/internal/firmware"
)

// IdentifyData holds pre-built 4KB Identify responses.
type IdentifyData struct {
	Controller [4096]byte // CNS=1
	Namespace  [4096]byte // CNS=0, NSID=1
}

// BuildIdentifyData constructs both Identify structures from donor IDs.
func BuildIdentifyData(ids firmware.DeviceIDs, barData []byte) *IdentifyData {
	d := &IdentifyData{}
	d.Controller = buildIdentifyController(ids, barData)
	d.Namespace = buildIdentifyNamespace(barData)
	return d
}

// buildIdentifyController - CNS=1, NVMe 1.4 spec Figure 247.
func buildIdentifyController(ids firmware.DeviceIDs, barData []byte) [4096]byte {
	var data [4096]byte

	// VID / SSVID
	binary.LittleEndian.PutUint16(data[0x000:], ids.VendorID)
	binary.LittleEndian.PutUint16(data[0x002:], ids.SubsysVendorID)

	// SN (20B ASCII)
	sn := generateSerialNumber(ids)
	copy(data[0x004:0x018], padASCII(sn, 20))

	// MN (40B ASCII)
	mn := modelNumberForVendor(ids.VendorID)
	copy(data[0x018:0x040], padASCII(mn, 40))

	// FR (8B ASCII) - vendor-appropriate firmware revision
	fr := firmwareRevisionForVendor(ids.VendorID)
	copy(data[0x040:0x048], padASCII(fr, 8))

	data[0x048] = 6 // RAB

	// IEEE OUI
	oui := ouiForVendor(ids.VendorID)
	data[0x049] = oui[0]
	data[0x04A] = oui[1]
	data[0x04B] = oui[2]

	data[0x04C] = 0x00 // CMIC
	// MDTS - derive from donor CAP.MPSMIN when available.
	// Most NVMe drives use MDTS=5 (2^5 * MPSMIN pages = 128KB at 4KB pages).
	// Some use higher values; we stay conservative to avoid host-side overruns
	// that our fake controller can't actually service.
	mdts := uint8(5)
	if len(barData) >= 0x08 {
		capHi := binary.LittleEndian.Uint32(barData[0x04:0x08])
		mpsmin := (capHi >> 16) & 0x0F // CAP.MPSMIN: bits 51:48
		// MPSMIN=0 -> 4KB pages -> MDTS=5 is fine (128KB)
		// MPSMIN=1 -> 8KB pages -> MDTS=4 keeps total ≤128KB
		if mpsmin > 0 && mpsmin < 5 && mdts > uint8(5-mpsmin) {
			mdts = uint8(5 - mpsmin)
			if mdts == 0 {
				mdts = 1
			}
		}
	}
	data[0x04D] = mdts                                  // MDTS
	binary.LittleEndian.PutUint16(data[0x04E:], 0x0001) // CNTLID
	// VER must match BAR VS register - stornvme.sys compares the two and
	// triggers Code 10 on mismatch (e.g. BAR says 1.3, identify says 1.4).
	nvmeVer := uint32(0x00010400) // default: NVMe 1.4
	if len(barData) >= 0x0C {
		nvmeVer = binary.LittleEndian.Uint32(barData[0x08:0x0C])
		if nvmeVer == 0 {
			nvmeVer = 0x00010400
		}
	}
	binary.LittleEndian.PutUint32(data[0x050:], nvmeVer)    // VER
	binary.LittleEndian.PutUint32(data[0x054:], 0x00000064) // RTD3 Resume Latency (100µs)
	binary.LittleEndian.PutUint32(data[0x058:], 0x00000064) // RTD3 Entry Latency (100µs)
	binary.LittleEndian.PutUint32(data[0x05C:], 0x00000000) // OAES
	binary.LittleEndian.PutUint32(data[0x060:], 0x00000000) // CTRATT

	// Admin Command Set Attributes (0x100)
	binary.LittleEndian.PutUint16(data[0x100:], 0x0006) // OACS - Format + FW Download
	data[0x102] = 3                                     // ACL
	data[0x103] = 7                                     // AERL
	data[0x104] = 0x14                                  // FRMW
	data[0x105] = 0x0E                                  // LPA
	data[0x106] = 0x3F                                  // ELPE
	data[0x107] = 0                                     // NPSS (1 power state)
	data[0x108] = 0x01                                  // AVSCC
	data[0x111] = 0x01                                  // CNTRLTYPE - I/O Controller

	// NVM Command Set Attributes (0x200)
	data[0x200] = 0x66                                  // SQES min=max=64B
	data[0x201] = 0x44                                  // CQES min=max=16B
	binary.LittleEndian.PutUint16(data[0x202:], 0x0000) // MAXCMD
	binary.LittleEndian.PutUint32(data[0x204:], 1)      // NN - 1 namespace
	binary.LittleEndian.PutUint16(data[0x208:], 0x001F) // ONCS
	binary.LittleEndian.PutUint16(data[0x20A:], 0x0000) // FUSES
	data[0x20C] = 0x00                                  // FNA
	data[0x20D] = 0x01                                  // VWC present
	binary.LittleEndian.PutUint16(data[0x20E:], 0x0000) // AWUN
	binary.LittleEndian.PutUint16(data[0x210:], 0x0000) // AWUPF

	// NVMe Qualified Name - helps identify the controller to host software
	subnqn := fmt.Sprintf("nqn.2014.08.org.nvmexpress:%04x-%04x", ids.VendorID, ids.DeviceID)
	if len(subnqn) > 256 {
		subnqn = subnqn[:256]
	}
	copy(data[0x300:0x300+256], padASCII(subnqn, 256))

	// Power State 0 (0x800)
	binary.LittleEndian.PutUint16(data[0x800:], 500) // MP 5.00W
	data[0x803] = 0x00                               // flags
	binary.LittleEndian.PutUint32(data[0x804:], 0)   // ENLAT
	binary.LittleEndian.PutUint32(data[0x808:], 0)   // EXLAT
	data[0x80C] = 0                                  // RRT
	data[0x80D] = 0                                  // RRL
	data[0x80E] = 0                                  // RWT
	data[0x80F] = 0                                  // RWL

	return data
}

// buildIdentifyNamespace - CNS=0 NSID=1, NVMe 1.4 spec Figure 245.
func buildIdentifyNamespace(barData []byte) [4096]byte {
	var data [4096]byte

	// 1TB / 512B sectors
	var nsze uint64 = 1953525168
	_ = barData // reserved for future donor extraction

	binary.LittleEndian.PutUint64(data[0x000:], nsze) // NSZE
	binary.LittleEndian.PutUint64(data[0x008:], nsze) // NCAP
	binary.LittleEndian.PutUint64(data[0x010:], nsze) // NUSE

	data[0x018] = 0x00 // NSFEAT
	data[0x019] = 0    // NLBAF (1 format)
	data[0x01A] = 0x00 // FLBAS (format 0 active)
	data[0x01B] = 0x00 // MC
	data[0x01C] = 0x00 // DPC
	data[0x01D] = 0x00 // DPS
	data[0x01E] = 0x00 // NMIC
	data[0x01F] = 0x00 // RESCAP
	data[0x020] = 0x00 // FPI

	generateNGUID(data[0x080:0x090]) // NGUID (16B)
	// EUI64 at 0x098 stays zero (optional)

	// LBAF0: LBADS=9 (512B), MS=0, RP=0
	binary.LittleEndian.PutUint32(data[0x0C0:], 0x00090000)

	return data
}

// generateSerialNumber creates a vendor-prefixed random serial.
func generateSerialNumber(ids firmware.DeviceIDs) string {
	prefix := vendorSNPrefix(ids.VendorID)
	const chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	buf := make([]byte, 12)
	randBytes := make([]byte, 12)
	_, _ = rand.Read(randBytes)
	for i := range buf {
		buf[i] = chars[randBytes[i]%byte(len(chars))]
	}
	return prefix + string(buf)
}

func vendorSNPrefix(vid uint16) string {
	switch vid {
	case 0x144D:
		return "S6PXNG0T"
	case 0x1179:
		return "Y9SF7"
	case 0x1987:
		return "PHS"
	case 0x15B7:
		return "WD-"
	case 0x1C5C:
		return "SKH"
	case 0x8086:
		return "BTLH"
	case 0x1E0F:
		return "KX"
	default:
		return "NVME"
	}
}

func modelNumberForVendor(vid uint16) string {
	switch vid {
	case 0x144D:
		return "Samsung SSD 980 PRO 1TB"
	case 0x1179:
		return "TOSHIBA KXG60ZNV1T02"
	case 0x1987:
		return "Sabrent Rocket 4.0 1TB"
	case 0x15B7:
		return "WD Black SN850X 1TB"
	case 0x1C5C:
		return "SK hynix PC801 NVMe 1TB"
	case 0x8086:
		return "Intel SSD 670p 1TB"
	case 0x1E0F:
		return "KIOXIA EXCERIA PRO 1TB"
	case 0x126F:
		return "NVMe SSD 1TB"
	default:
		return "NVMe SSD Drive 1TB"
	}
}

func ouiForVendor(vid uint16) [3]byte {
	switch vid {
	case 0x144D:
		return [3]byte{0x00, 0x26, 0x2D}
	case 0x8086:
		return [3]byte{0x5C, 0xD2, 0xE4}
	case 0x15B7:
		return [3]byte{0x00, 0x1B, 0x44}
	case 0x1C5C:
		return [3]byte{0x00, 0xAD, 0x00}
	default:
		return [3]byte{0x00, 0x00, 0x00}
	}
}

func firmwareRevisionForVendor(vid uint16) string {
	switch vid {
	case 0x144D:
		return "5B2QGXA7"
	case 0x1179:
		return "ADHA0101"
	case 0x1987:
		return "RKT4CB.2"
	case 0x15B7:
		return "620311WD"
	case 0x1C5C:
		return "41062C20"
	case 0x8086:
		return "002C"
	case 0x1E0F:
		return "1102"
	default:
		return "1.0"
	}
}

func padASCII(s string, n int) []byte {
	buf := make([]byte, n)
	for i := range buf {
		buf[i] = ' '
	}
	copy(buf, []byte(s))
	return buf
}

func generateNGUID(buf []byte) {
	if len(buf) < 16 {
		return
	}
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	b[6] = (b[6] & 0x0F) | 0x40 // UUID v4
	b[8] = (b[8] & 0x3F) | 0x80
	copy(buf, b)
}

// IdentifyDataToHex converts 8KB identify data to Xilinx HEX init format.
func IdentifyDataToHex(id *IdentifyData) string {
	var result string
	result += "// NVMe Identify ROM - 8KB (4KB Controller + 4KB Namespace)\n"
	result += "// Auto-generated by PCILeechGen\n"

	for i := 0; i < 4096; i += 4 {
		word := binary.LittleEndian.Uint32(id.Controller[i : i+4])
		result += fmt.Sprintf("%08X\n", word)
	}
	for i := 0; i < 4096; i += 4 {
		word := binary.LittleEndian.Uint32(id.Namespace[i : i+4])
		result += fmt.Sprintf("%08X\n", word)
	}

	return result
}
