// Package util provides common utility functions.
package util

import (
	"encoding/binary"
	"fmt"
	"strings"
)

// HexToBytes converts a hex string (with or without spaces) to a byte slice.
func HexToBytes(hex string) ([]byte, error) {
	hex = strings.ReplaceAll(hex, " ", "")
	hex = strings.ReplaceAll(hex, "\n", "")
	hex = strings.ReplaceAll(hex, "\r", "")

	if len(hex)%2 != 0 {
		return nil, fmt.Errorf("hex string has odd length: %d", len(hex))
	}

	result := make([]byte, len(hex)/2)
	for i := 0; i < len(result); i++ {
		_, err := fmt.Sscanf(hex[i*2:i*2+2], "%02x", &result[i])
		if err != nil {
			return nil, fmt.Errorf("invalid hex at position %d: %w", i*2, err)
		}
	}
	return result, nil
}

// BytesToHex converts a byte slice to a hex string with spaces between bytes.
func BytesToHex(data []byte) string {
	parts := make([]string, len(data))
	for i, b := range data {
		parts[i] = fmt.Sprintf("%02x", b)
	}
	return strings.Join(parts, " ")
}

// BytesToHexNoSpaces converts a byte slice to a compact hex string.
func BytesToHexNoSpaces(data []byte) string {
	var sb strings.Builder
	for _, b := range data {
		fmt.Fprintf(&sb, "%02x", b)
	}
	return sb.String()
}

// U32ToLEBytes converts a uint32 to a 4-byte little-endian slice.
func U32ToLEBytes(v uint32) []byte {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, v)
	return b
}

// U16ToLEBytes converts a uint16 to a 2-byte little-endian slice.
func U16ToLEBytes(v uint16) []byte {
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, v)
	return b
}

// LEBytesToU32 converts a 4-byte little-endian slice to uint32.
func LEBytesToU32(b []byte) uint32 {
	if len(b) < 4 {
		return 0
	}
	return binary.LittleEndian.Uint32(b)
}

// LEBytesToU16 converts a 2-byte little-endian slice to uint16.
func LEBytesToU16(b []byte) uint16 {
	if len(b) < 2 {
		return 0
	}
	return binary.LittleEndian.Uint16(b)
}

// SwapEndian32 swaps the byte order of a 32-bit value.
func SwapEndian32(v uint32) uint32 {
	return (v>>24)&0xFF | (v>>8)&0xFF00 | (v<<8)&0xFF0000 | (v<<24)&0xFF000000
}
