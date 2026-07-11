package donor

import "fmt"

// ValidateDeviceLayout reports invalid BAR pairing and MSI-X table placement.
func ValidateDeviceLayout(ctx *DeviceContext) []string {
	if ctx == nil {
		return []string{"device context is nil"}
	}
	var issues []string
	barSizes := make(map[int]uint64, len(ctx.BARs))
	for _, bar := range ctx.BARs {
		barSizes[bar.Index] = bar.Size
		if bar.Is64Bit && bar.Index == 5 {
			issues = append(issues, "64-bit BAR5 has no upper BAR")
		}
	}
	if ctx.MSIXData == nil || ctx.MSIXData.TableSize <= 0 {
		return issues
	}

	msix := ctx.MSIXData
	tableEnd := uint64(msix.TableOffset) + uint64(msix.TableSize)*16
	issues = appendMSIXBoundsIssue(issues, "table", msix.TableBIR, tableEnd, barSizes)
	pbaBytes := uint64((msix.TableSize+63)/64) * 8
	pbaEnd := uint64(msix.PBAOffset) + pbaBytes
	issues = appendMSIXBoundsIssue(issues, "PBA", msix.PBABIR, pbaEnd, barSizes)
	return issues
}

func appendMSIXBoundsIssue(issues []string, kind string, barIndex int, end uint64, barSizes map[int]uint64) []string {
	size, ok := barSizes[barIndex]
	if !ok || size == 0 {
		return append(issues, fmt.Sprintf("MSI-X %s references unavailable BAR%d", kind, barIndex))
	}
	if end > size {
		return append(issues, fmt.Sprintf("MSI-X %s exceeds BAR%d: end 0x%x, size 0x%x", kind, barIndex, end, size))
	}
	return issues
}
