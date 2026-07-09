//go:build !linux

package mmio

import (
	"fmt"
	"runtime"
	"time"
)

func RequireLiveTrace() error {
	return fmt.Errorf("live MMIO tracing is unsupported on %s; use --trace-file", runtime.GOOS)
}

func LiveTrace(bdf string, duration time.Duration) (*TraceResult, error) {
	return nil, RequireLiveTrace()
}
