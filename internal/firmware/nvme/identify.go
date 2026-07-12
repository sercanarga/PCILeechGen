// Package nvme builds NVMe Identify Controller/Namespace 4KB responses
// for the Admin Responder FSM.
package nvme

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"strings"

	"github.com/sercanarga/pcileechgen/internal/firmware"
)

// IdentifyData holds pre-built 4KB Identify responses.
type IdentifyData struct {
	Controller [4096]byte // CNS=1
	Namespace  [4096]byte // CNS=0, NSID=1
}

// ControllerIdentity holds donor-captured NVMe controller strings; empty fields fall back to synthesis.
type ControllerIdentity struct {
	Serial string
	Model  string
	FWRev  string

	RawControllerIdent []byte
	RawNamespaceIdent  []byte
}

// BuildIdentifyData constructs both Identify structures from donor IDs.
func BuildIdentifyData(ids firmware.DeviceIDs, barData []byte, identity *ControllerIdentity) *IdentifyData {
	d := &IdentifyData{}

	// Use raw donor data if available, otherwise synthesize.
	if identity != nil && len(identity.RawControllerIdent) == 4096 {
		d.Controller = [4096]byte(identity.RawControllerIdent)
		// Force emulator-critical fields: VID/SSVID, MDTS clamp, VER==BAR VS.
		binary.LittleEndian.PutUint16(d.Controller[0x000:], ids.VendorID)
		binary.LittleEndian.PutUint16(d.Controller[0x002:], ids.SubsysVendorID)
		d.Controller[0x04D] = deriveMDTS(barData)
		binary.LittleEndian.PutUint32(d.Controller[0x050:], deriveVER(barData))
		binary.LittleEndian.PutUint32(d.Controller[0x05C:], 0) // OAES
		binary.LittleEndian.PutUint32(d.Controller[0x060:], 0) // CTRATT
	} else {
		d.Controller = buildIdentifyController(ids, barData, identity)
	}

	if identity != nil && len(identity.RawNamespaceIdent) == 4096 {
		d.Namespace = [4096]byte(identity.RawNamespaceIdent)
	} else {
		d.Namespace = buildIdentifyNamespace(barData)
	}

	return d
}

// deriveMDTS clamps MDTS to ≤5 from CAP.MPSMIN (backend-safe).
func deriveMDTS(barData []byte) uint8 {
	mdts := uint8(5)
	if len(barData) >= 0x08 {
		capHi := binary.LittleEndian.Uint32(barData[0x04:0x08])
		mpsmin := (capHi >> 16) & 0x0F // CAP.MPSMIN: bits 51:48
		if mpsmin > 0 && mpsmin < 5 && mdts > uint8(5-mpsmin) {
			mdts = uint8(5 - mpsmin)
			if mdts == 0 {
				mdts = 1
			}
		}
	}
	return mdts
}

// deriveVER returns VER equal to BAR VS (stornvme Code 10 guard).
func deriveVER(barData []byte) uint32 {
	if len(barData) >= 0x0C {
		if v := binary.LittleEndian.Uint32(barData[0x08:0x0C]); v != 0 {
			return v
		}
	}
	return 0x00010400 // NVMe 1.4
}

// buildIdentifyController - CNS=1, NVMe 1.4 spec Figure 247.
func buildIdentifyController(ids firmware.DeviceIDs, barData []byte, identity *ControllerIdentity) [4096]byte {
	var data [4096]byte

	// VID / SSVID
	binary.LittleEndian.PutUint16(data[0x000:], ids.VendorID)
	binary.LittleEndian.PutUint16(data[0x002:], ids.SubsysVendorID)

	// SN (20B ASCII)
	sn := ""
	if identity != nil {
		sn = strings.TrimSpace(identity.Serial)
	}
	if sn == "" {
		sn = generateSerialNumber(ids)
	}
	copy(data[0x004:0x018], padASCII(sn, 20))

	// MN (40B ASCII)
	mn := ""
	if identity != nil {
		mn = strings.TrimSpace(identity.Model)
	}
	if mn == "" {
		mn = modelNumberForVendor(ids.VendorID)
	}
	copy(data[0x018:0x040], padASCII(mn, 40))

	// FR (8B ASCII)
	fr := ""
	if identity != nil {
		fr = strings.TrimSpace(identity.FWRev)
	}
	if fr == "" {
		fr = firmwareRevisionForVendor(ids.VendorID)
	}
	copy(data[0x040:0x048], padASCII(fr, 8))

	data[0x048] = 6 // RAB

	// IEEE OUI
	oui := ouiForVendor(ids.VendorID)
	data[0x049] = oui[2]
	data[0x04A] = oui[1]
	data[0x04B] = oui[0]

	data[0x04C] = 0x00                                              // CMIC
	data[0x04D] = deriveMDTS(barData)                               // MDTS — clamped to backend-safe
	binary.LittleEndian.PutUint16(data[0x04E:], 0x0001)             // CNTLID
	binary.LittleEndian.PutUint32(data[0x050:], deriveVER(barData)) // VER — must match BAR VS
	binary.LittleEndian.PutUint32(data[0x054:], 0x00000064)         // RTD3 Resume Latency (100µs)
	binary.LittleEndian.PutUint32(data[0x058:], 0x00000064)         // RTD3 Entry Latency (100µs)
	binary.LittleEndian.PutUint32(data[0x05C:], 0x00000000)         // OAES
	binary.LittleEndian.PutUint32(data[0x060:], 0x00000000)         // CTRATT

	// Admin Command Set Attributes (0x100)
	binary.LittleEndian.PutUint16(data[0x100:], 0x0003) // OACS - Format NVM + AER
	data[0x102] = 3                                     // ACL
	data[0x103] = 7                                     // AERL
	data[0x104] = 0x02                                  // FRMW
	data[0x105] = 0x00                                  // LPA
	data[0x106] = 0x00                                  // ELPE
	data[0x107] = 0x00                                  // NPSS (no extra power states)
	data[0x108] = 0x00                                  // AVSCC
	data[0x109] = 0x00                                  // APSTA
	binary.LittleEndian.PutUint16(data[0x10A:], 343)    // WCTEMP
	binary.LittleEndian.PutUint16(data[0x10C:], 358)    // CCTEMP (matches SMART critical threshold)
	data[0x15C] = 0x01                                  // CNTRLTYPE - I/O Controller (NVMe 1.4)

	// NVM Command Set Attributes (0x200)
	data[0x200] = 0x66                                  // SQES min=max=64B
	data[0x201] = 0x44                                  // CQES min=max=16B
	binary.LittleEndian.PutUint16(data[0x202:], 0x0000) // MAXCMD
	binary.LittleEndian.PutUint32(data[0x204:], 1)      // NN - 1 namespace
	binary.LittleEndian.PutUint16(data[0x208:], 0x000C) // ONCS - DSM + Write Zeroes only
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

	// 1 TB / 512 B sectors; kept in sync with the templates' NVME_ADVERTISED_LBAS.
	var nsze uint64 = 2000409264
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
		return [3]byte{0x00, 0x25, 0x38}
	case 0x8086:
		return [3]byte{0x5C, 0xD2, 0xE4}
	case 0x15B7:
		return [3]byte{0x00, 0x1B, 0x44}
	case 0x1C5C:
		return [3]byte{0xAC, 0xE4, 0x2E}
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

// DoorbellStrideFromCAP returns DSTRD (CAP bits 35:32 = low nibble of the high dword at BAR0+0x04).
func DoorbellStrideFromCAP(capHi uint32) uint32 {
	return capHi & 0x0F
}

// IdentifyDataToHex emits the 4 KB Identify Controller ROM as Xilinx HEX init text (namespace is runtime-generated).
func IdentifyDataToHex(id *IdentifyData) string {
	var b strings.Builder
	b.Grow(4096 / 4 * 9)
	for i := 0; i < 4096; i += 4 {
		fmt.Fprintf(&b, "%08X\n", binary.LittleEndian.Uint32(id.Controller[i:i+4]))
	}
	return b.String()
}
