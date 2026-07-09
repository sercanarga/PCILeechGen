//go:build !linux

package vfio

import (
	"fmt"
	"runtime"
)

func Collect(bdf string) (*DeviceDump, error) {
	return nil, fmt.Errorf("VFIO collection is unsupported on %s; use --from-json", runtime.GOOS)
}

func CollectToFile(bdf, outputPath string) error {
	_, err := Collect(bdf)
	return err
}

