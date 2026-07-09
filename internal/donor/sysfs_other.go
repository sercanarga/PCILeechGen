//go:build !linux

package donor

import (
	"fmt"
	"os"
	"runtime"

	"github.com/sercanarga/pcileechgen/internal/pci"
)

func (sr *SysfsReader) readBARViaMmap(f *os.File, size int) ([]byte, error) {
	return nil, fmt.Errorf("live donor BAR access is unsupported on %s", runtime.GOOS)
}

func (sr *SysfsReader) ReadNVMeRawIdentify(bdf pci.BDF, cns uint32, nsid uint32) ([]byte, error) {
	return nil, fmt.Errorf("live NVMe identity collection is unsupported on %s; use --from-json", runtime.GOOS)
}
