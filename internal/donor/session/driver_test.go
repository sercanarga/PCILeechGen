package session

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunUnbindBindRestoresDriverAfterCaptureError(t *testing.T) {
	root := t.TempDir()
	bdf := "0000:02:00.0"
	driverDir := filepath.Join(root, "drivers", "r8169")
	deviceDir := filepath.Join(root, "devices", bdf)
	mustMkdir(t, driverDir)
	mustMkdir(t, filepath.Join(deviceDir, "net", "eth0"))
	mustWrite(t, filepath.Join(driverDir, "unbind"), nil)
	mustWrite(t, filepath.Join(driverDir, "bind"), nil)
	if err := os.Symlink(driverDir, filepath.Join(deviceDir, "driver")); err != nil {
		t.Fatal(err)
	}
	route := filepath.Join(root, "route")
	mustWrite(t, route, []byte("Iface\tDestination\tGateway\tFlags\neth1\t00000000\t00000000\t0003\n"))
	controller := NewDriverController(filepath.Join(root, "devices"), route)

	err := controller.RunUnbindBind(bdf, false, func() error { return errors.New("capture failed") })
	if err == nil || !strings.Contains(err.Error(), "capture failed") {
		t.Fatalf("RunUnbindBind() error = %v", err)
	}
	bound, err := os.ReadFile(filepath.Join(driverDir, "bind"))
	if err != nil {
		t.Fatal(err)
	}
	if string(bound) != bdf {
		t.Fatalf("bind = %q, want %q", bound, bdf)
	}
}

func TestRunUnbindBindRefusesDefaultRouteInterface(t *testing.T) {
	root := t.TempDir()
	bdf := "0000:02:00.0"
	deviceDir := filepath.Join(root, "devices", bdf)
	mustMkdir(t, filepath.Join(deviceDir, "net", "eth0"))
	route := filepath.Join(root, "route")
	mustWrite(t, route, []byte("Iface\tDestination\tGateway\tFlags\neth0\t00000000\t00000000\t0003\n"))
	controller := NewDriverController(filepath.Join(root, "devices"), route)

	err := controller.RunUnbindBind(bdf, false, func() error { return nil })
	if err == nil || !strings.Contains(err.Error(), "default-route interface") {
		t.Fatalf("RunUnbindBind() error = %v", err)
	}
}

func TestRunUnbindBindCaptureCanRebindDuringCapture(t *testing.T) {
	root := t.TempDir()
	bdf := "0000:02:00.0"
	driverDir := filepath.Join(root, "drivers", "r8169")
	deviceDir := filepath.Join(root, "devices", bdf)
	mustMkdir(t, driverDir)
	mustMkdir(t, filepath.Join(deviceDir, "net"))
	mustWrite(t, filepath.Join(driverDir, "unbind"), nil)
	mustWrite(t, filepath.Join(driverDir, "bind"), nil)
	if err := os.Symlink(driverDir, filepath.Join(deviceDir, "driver")); err != nil {
		t.Fatal(err)
	}
	route := filepath.Join(root, "route")
	mustWrite(t, route, []byte("Iface\tDestination\tGateway\tFlags\n"))
	controller := NewDriverController(filepath.Join(root, "devices"), route)

	err := controller.RunUnbindBindCapture(bdf, false, func(bind func() error) error {
		if err := bind(); err != nil {
			return err
		}
		bound, err := os.ReadFile(filepath.Join(driverDir, "bind"))
		if err != nil {
			return err
		}
		if string(bound) != bdf {
			return errors.New("driver was not rebound during capture")
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}

func mustMkdir(t *testing.T, path string) {
	t.Helper()
	if err := os.MkdirAll(path, 0o755); err != nil {
		t.Fatal(err)
	}
}

func mustWrite(t *testing.T, path string, data []byte) {
	t.Helper()
	if err := os.WriteFile(path, data, 0o644); err != nil {
		t.Fatal(err)
	}
}
