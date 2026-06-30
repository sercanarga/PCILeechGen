package ahci

import (
	"strings"
	"testing"
)

func ataString(w [256]uint16, start, words int) string {
	var b []byte
	for i := 0; i < words; i++ {
		b = append(b, byte(w[start+i]>>8), byte(w[start+i]))
	}
	return strings.TrimRight(string(b), " ")
}

func TestBuildIdentify(t *testing.T) {
	const sectors = uint64(0x1_0000_0000) // 4G sectors -> 48-bit range
	w := BuildIdentify("PCILeech SATA SSD", "SN0001", "1.0", sectors)

	if w[0] != 0x0040 {
		t.Errorf("word0 = %04x, want 0040", w[0])
	}
	if w[49]&0x0300 != 0x0300 {
		t.Errorf("word49 LBA+DMA not set: %04x", w[49])
	}
	if w[83]&0x0400 == 0 {
		t.Errorf("word83 LBA48 not advertised: %04x", w[83])
	}
	if got := ataString(w, 27, 20); got != "PCILeech SATA SSD" {
		t.Errorf("model = %q", got)
	}
	if got := ataString(w, 10, 10); got != "SN0001" {
		t.Errorf("serial = %q", got)
	}
	got := uint64(w[100]) | uint64(w[101])<<16 | uint64(w[102])<<32 | uint64(w[103])<<48
	if got != sectors {
		t.Errorf("LBA48 sectors = %x, want %x", got, sectors)
	}

	// integrity: signature 0xA5 + checksum so the byte sum is 0 mod 256
	if byte(w[255]) != 0xA5 {
		t.Errorf("integrity signature = %02x, want a5", byte(w[255]))
	}
	var sum uint8
	for i := 0; i < 256; i++ {
		sum += uint8(w[i]) + uint8(w[i]>>8)
	}
	if sum != 0 {
		t.Errorf("checksum byte-sum = %d, want 0", sum)
	}
}

func TestIdentifyHex(t *testing.T) {
	w := BuildIdentify("M", "S", "F", 2048)
	hex := IdentifyHex(w)
	lines := strings.Split(strings.TrimRight(hex, "\n"), "\n")
	if len(lines) != 128 {
		t.Fatalf("want 128 dword lines, got %d", len(lines))
	}
	// dword0 = word1<<16 | word0 ; word0 = 0x0040
	if !strings.HasSuffix(lines[0], "0040") {
		t.Errorf("dword0 = %q, want low word 0040", lines[0])
	}
}
