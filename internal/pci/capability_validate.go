package pci

import "fmt"

// ValidateCapabilityChains reports malformed standard and extended capability links.
func ValidateCapabilityChains(cs *ConfigSpace) []string {
	if cs == nil {
		return []string{"configuration space is nil"}
	}

	var issues []string
	if cs.HasCapabilities() {
		issues = append(issues, validateStandardCapabilityChain(cs)...)
	}
	if cs.Size >= ConfigSpaceSize {
		issues = append(issues, validateExtendedCapabilityChain(cs)...)
	}
	return issues
}

func validateStandardCapabilityChain(cs *ConfigSpace) []string {
	visited := make(map[int]bool)
	ptr := int(cs.CapabilityPointer())
	for ptr != 0 {
		if ptr < 0x40 || ptr > 0xFC {
			return []string{fmt.Sprintf("standard capability pointer 0x%03x outside 0x040-0x0fc", ptr)}
		}
		if ptr&3 != 0 {
			return []string{fmt.Sprintf("standard capability pointer 0x%03x is not DWORD-aligned", ptr)}
		}
		if visited[ptr] {
			return []string{fmt.Sprintf("standard capability loop at 0x%03x", ptr)}
		}
		visited[ptr] = true
		ptr = int(cs.ReadU8(ptr + 1))
	}
	return nil
}

func validateExtendedCapabilityChain(cs *ConfigSpace) []string {
	visited := make(map[int]bool)
	ptr := 0x100
	for ptr != 0 {
		if ptr < 0x100 || ptr > 0xFFC {
			return []string{fmt.Sprintf("extended capability pointer 0x%03x outside 0x100-0xffc", ptr)}
		}
		if ptr&3 != 0 {
			return []string{fmt.Sprintf("extended capability pointer 0x%03x is not DWORD-aligned", ptr)}
		}
		if visited[ptr] {
			return []string{fmt.Sprintf("extended capability loop at 0x%03x", ptr)}
		}
		visited[ptr] = true

		header := cs.ReadU32(ptr)
		if header == 0 || header == 0xFFFFFFFF {
			return nil
		}
		ptr = int((header >> 20) & 0xFFF)
	}
	return nil
}
