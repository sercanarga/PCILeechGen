// Stub for non-Linux platforms — VFIO is Linux-only.
//
//go:build !linux

package vfio

import "fmt"

// Collect is not available on non-Linux platforms.
func Collect(bdf string) (*DeviceDump, error) {
	return nil, fmt.Errorf("VFIO is only available on Linux")
}

// CollectToFile is not available on non-Linux platforms.
func CollectToFile(bdf, outputPath string) error {
	return fmt.Errorf("VFIO is only available on Linux")
}
