//go:build !linux

package donor

import (
	"fmt"
	"runtime"
)

func (p *BARProfiler) ProfileBAR(resourcePath string, barIndex, maxSize int) (*BARProfile, error) {
	return nil, fmt.Errorf("live donor BAR profiling is unsupported on %s; use --from-json", runtime.GOOS)
}
