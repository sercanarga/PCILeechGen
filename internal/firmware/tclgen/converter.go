package tclgen

import (
	"fmt"

	"github.com/sercanarga/pcileechgen/internal/firmware"
)

// linkSpeedToTCL converts a numeric link speed to Vivado TCL format.
func linkSpeedToTCL(speed uint8) string {
	switch speed {
	case firmware.LinkSpeedGen1:
		return "2.5_GT/s"
	case firmware.LinkSpeedGen3:
		return "8.0_GT/s"
	default:
		return "5.0_GT/s" // Gen2 default
	}
}

// linkSpeedToTrgt converts a numeric link speed to Trgt_Link_Speed TCL property.
func linkSpeedToTrgt(speed uint8) string {
	switch speed {
	case firmware.LinkSpeedGen1:
		return "4'h1"
	case firmware.LinkSpeedGen3:
		return "4'h3"
	default:
		return "4'h2"
	}
}

// linkWidthToTCL converts a numeric link width to Vivado TCL format.
func linkWidthToTCL(width uint8) string {
	switch width {
	case 2:
		return "X2"
	case 4:
		return "X4"
	case 8:
		return "X8"
	default:
		return "X1"
	}
}

// clampLinkWidth limits donor link width to board's physical lane count.
func clampLinkWidth(donorWidth uint8, boardLanes int) uint8 {
	if int(donorWidth) > boardLanes {
		return uint8(boardLanes)
	}
	if donorWidth == 0 {
		return uint8(boardLanes)
	}
	return donorWidth
}

// barSizeToTCL converts a BAR size in bytes to Vivado TCL scale and size values.
func barSizeToTCL(sizeBytes uint64) (scale string, size string) {
	if sizeBytes == 0 {
		return "Kilobytes", "4"
	}
	if sizeBytes >= 1024*1024 {
		mb := sizeBytes / (1024 * 1024)
		return "Megabytes", fmt.Sprintf("%d", mb)
	}
	kb := sizeBytes / 1024
	if kb < 4 {
		kb = 4 // Minimum 4KB
	}
	return "Kilobytes", fmt.Sprintf("%d", kb)
}
