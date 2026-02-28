package pci

import (
	"bufio"
	"os"
	"strings"
)

// PCIDB holds vendor and device name mappings parsed from pci.ids.
type PCIDB struct {
	Vendors map[uint16]string // vendor ID -> name
	Devices map[uint32]string // (vendor<<16 | device) -> name
}

// pci.ids search paths (same as lspci)
var pciIDPaths = []string{
	"/usr/share/hwdata/pci.ids",
	"/usr/share/misc/pci.ids",
	"/usr/share/pci.ids",
}

// LoadPCIDB loads the PCI ID database from the system.
func LoadPCIDB() *PCIDB {
	for _, path := range pciIDPaths {
		db, err := parsePCIIDs(path)
		if err == nil {
			return db
		}
	}
	return &PCIDB{
		Vendors: make(map[uint16]string),
		Devices: make(map[uint32]string),
	}
}

// VendorName returns the vendor name or hex fallback.
func (db *PCIDB) VendorName(vendorID uint16) string {
	if name, ok := db.Vendors[vendorID]; ok {
		return name
	}
	return ""
}

// DeviceName returns the device name or empty string.
func (db *PCIDB) DeviceName(vendorID, deviceID uint16) string {
	key := uint32(vendorID)<<16 | uint32(deviceID)
	if name, ok := db.Devices[key]; ok {
		return name
	}
	return ""
}

// parsePCIIDs parses a pci.ids file.
// Format:
//
//	VVVV  Vendor Name
//	\tDDDD  Device Name
func parsePCIIDs(path string) (*PCIDB, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	db := &PCIDB{
		Vendors: make(map[uint16]string),
		Devices: make(map[uint32]string),
	}

	var currentVendor uint16
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		// skip comments and empty lines
		if len(line) == 0 || line[0] == '#' {
			continue
		}

		// stop at class definitions
		if line[0] == 'C' && len(line) > 1 && line[1] == ' ' {
			break
		}

		if line[0] == '\t' && len(line) > 1 && line[1] == '\t' {
			// subsystem line - skip
			continue
		}

		if line[0] == '\t' {
			// device line: \tDDDD  Device Name
			line = line[1:] // trim tab
			if len(line) < 6 {
				continue
			}
			devID := parseHex4(line[:4])
			if devID >= 0 {
				name := strings.TrimSpace(line[4:])
				key := uint32(currentVendor)<<16 | uint32(devID)
				db.Devices[key] = name
			}
		} else {
			// vendor line: VVVV  Vendor Name
			if len(line) < 6 {
				continue
			}
			vid := parseHex4(line[:4])
			if vid >= 0 {
				currentVendor = uint16(vid)
				db.Vendors[currentVendor] = strings.TrimSpace(line[4:])
			}
		}
	}

	return db, nil
}

// parseHex4 parses a 4-char hex string, returns -1 on failure.
func parseHex4(s string) int {
	if len(s) != 4 {
		return -1
	}
	var val int
	for _, c := range s {
		val <<= 4
		switch {
		case c >= '0' && c <= '9':
			val |= int(c - '0')
		case c >= 'a' && c <= 'f':
			val |= int(c-'a') + 10
		case c >= 'A' && c <= 'F':
			val |= int(c-'A') + 10
		default:
			return -1
		}
	}
	return val
}
