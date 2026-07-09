//go:build !linux

package donor

import (
	"fmt"
	"runtime"
)

func RequireLiveCollection() error {
	return fmt.Errorf("live donor collection is unsupported on %s; use --from-json", runtime.GOOS)
}
