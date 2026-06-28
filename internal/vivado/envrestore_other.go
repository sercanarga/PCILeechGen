//go:build !linux

package vivado

import (
	"context"
	"errors"
)

func probeShellAsUser(context.Context, string, string, uint32, uint32) ([]byte, error) {
	return nil, errors.New("credential drop not supported on this platform")
}
