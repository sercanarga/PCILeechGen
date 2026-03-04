// Stub for non-Linux — LiveTrace always returns an error.
//
//go:build !linux

package mmio

import (
	"fmt"
	"time"
)

// LiveTrace is not supported on non-Linux platforms.
func LiveTrace(bdf string, duration time.Duration) (*TraceResult, error) {
	return nil, fmt.Errorf("live MMIO tracing requires Linux with CONFIG_MMIOTRACE=y")
}
