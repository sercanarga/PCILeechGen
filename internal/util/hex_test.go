package util

import (
	"testing"
)

func TestHexToBytes(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    []byte
		wantErr bool
	}{
		{"simple", "0102", []byte{0x01, 0x02}, false},
		{"with spaces", "01 02 ff", []byte{0x01, 0x02, 0xff}, false},
		{"uppercase", "AABB", []byte{0xaa, 0xbb}, false},
		{"odd length", "012", nil, true},
		{"invalid hex", "zz", nil, true},
		{"empty", "", []byte{}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HexToBytes(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("HexToBytes(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if len(got) != len(tt.want) {
					t.Errorf("HexToBytes(%q) len = %d, want %d", tt.input, len(got), len(tt.want))
					return
				}
				for i := range got {
					if got[i] != tt.want[i] {
						t.Errorf("HexToBytes(%q)[%d] = 0x%02x, want 0x%02x", tt.input, i, got[i], tt.want[i])
					}
				}
			}
		})
	}
}

func TestBytesToHex(t *testing.T) {
	got := BytesToHex([]byte{0x01, 0x02, 0xff})
	want := "01 02 ff"
	if got != want {
		t.Errorf("BytesToHex() = %q, want %q", got, want)
	}
}

func TestBytesToHexNoSpaces(t *testing.T) {
	got := BytesToHexNoSpaces([]byte{0x01, 0x02, 0xff})
	want := "0102ff"
	if got != want {
		t.Errorf("BytesToHexNoSpaces() = %q, want %q", got, want)
	}
}

func TestU32Conversion(t *testing.T) {
	original := uint32(0x12345678)
	bytes := U32ToLEBytes(original)
	result := LEBytesToU32(bytes)
	if result != original {
		t.Errorf("U32 roundtrip: got 0x%08x, want 0x%08x", result, original)
	}

	// Verify little-endian byte order
	if bytes[0] != 0x78 || bytes[1] != 0x56 || bytes[2] != 0x34 || bytes[3] != 0x12 {
		t.Errorf("U32ToLEBytes byte order wrong: %v", bytes)
	}
}

func TestU16Conversion(t *testing.T) {
	original := uint16(0xABCD)
	bytes := U16ToLEBytes(original)
	result := LEBytesToU16(bytes)
	if result != original {
		t.Errorf("U16 roundtrip: got 0x%04x, want 0x%04x", result, original)
	}
}

func TestSwapEndian32(t *testing.T) {
	if SwapEndian32(0x12345678) != 0x78563412 {
		t.Errorf("SwapEndian32(0x12345678) = 0x%08x", SwapEndian32(0x12345678))
	}
}

func TestLEBytesToU32Short(t *testing.T) {
	if LEBytesToU32([]byte{0x01}) != 0 {
		t.Error("LEBytesToU32 with short slice should return 0")
	}
}

func TestReadLE32(t *testing.T) {
	data := []byte{0x78, 0x56, 0x34, 0x12}
	got := ReadLE32(data, 0)
	if got != 0x12345678 {
		t.Errorf("ReadLE32 = 0x%08x, want 0x12345678", got)
	}
}

func TestReadLE32_OutOfBounds(t *testing.T) {
	data := []byte{0x01, 0x02}
	got := ReadLE32(data, 0)
	if got != 0 {
		t.Errorf("ReadLE32 OOB = 0x%08x, want 0", got)
	}
}

func TestReadLE32_NegativeOffset(t *testing.T) {
	data := []byte{0x01, 0x02, 0x03, 0x04}
	got := ReadLE32(data, -1)
	if got != 0 {
		t.Errorf("ReadLE32 negative = 0x%08x, want 0", got)
	}
}

func TestWriteLE32(t *testing.T) {
	data := make([]byte, 8)
	WriteLE32(data, 0, 0xDEADBEEF)
	if data[0] != 0xEF || data[1] != 0xBE || data[2] != 0xAD || data[3] != 0xDE {
		t.Errorf("WriteLE32 = %02x %02x %02x %02x", data[0], data[1], data[2], data[3])
	}
}

func TestWriteLE32_OutOfBounds(t *testing.T) {
	data := make([]byte, 2)
	WriteLE32(data, 0, 0x12345678) // should be no-op
	if data[0] != 0 || data[1] != 0 {
		t.Error("WriteLE32 OOB should not write")
	}
}

func TestWriteLE32_NegativeOffset(t *testing.T) {
	data := make([]byte, 4)
	WriteLE32(data, -1, 0x12345678) // should be no-op
	for i, b := range data {
		if b != 0 {
			t.Errorf("data[%d] = %02x, want 0", i, b)
		}
	}
}

func TestLEBytesToU16_Short(t *testing.T) {
	got := LEBytesToU16([]byte{0x01})
	if got != 0 {
		t.Errorf("LEBytesToU16 short = %d, want 0", got)
	}
}

func TestSwapEndian32_Roundtrip(t *testing.T) {
	got := SwapEndian32(0x12345678)
	if got != 0x78563412 {
		t.Errorf("SwapEndian32 = 0x%08x, want 0x78563412", got)
	}
}

func TestCopyFile_NonExistent(t *testing.T) {
	err := CopyFile("/nonexistent/src", "/tmp/dst")
	if err == nil {
		t.Error("CopyFile should fail for non-existent source")
	}
}

func TestCopyDir_NonExistent(t *testing.T) {
	err := CopyDir("/nonexistent/src", "/tmp/dst")
	if err == nil {
		t.Error("CopyDir should fail for non-existent source")
	}
}
