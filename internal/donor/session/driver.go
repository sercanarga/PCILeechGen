package session

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// DriverController performs one bounded PCI driver unbind/bind cycle.
type DriverController struct {
	devicesPath string
	routePath   string
}

// NewDriverController creates a controller with injectable Linux paths.
func NewDriverController(devicesPath, routePath string) *DriverController {
	return &DriverController{devicesPath: devicesPath, routePath: routePath}
}

// NewLinuxDriverController creates a controller for the running Linux host.
func NewLinuxDriverController() *DriverController {
	return NewDriverController("/sys/bus/pci/devices", "/proc/net/route")
}

// RunUnbindBind unbinds the device, runs capture, and always restores its driver.
func (c *DriverController) RunUnbindBind(bdf string, force bool, capture func() error) error {
	if capture == nil {
		return fmt.Errorf("capture callback is nil")
	}
	return c.RunUnbindBindCapture(bdf, force, func(_ func() error) error {
		return capture()
	})
}

// RunUnbindBindCapture lets capture rebind the device while tracing remains active.
func (c *DriverController) RunUnbindBindCapture(bdf string, force bool, capture func(bind func() error) error) (resultErr error) {
	if capture == nil {
		return fmt.Errorf("capture callback is nil")
	}
	if err := c.CheckSafe(bdf, force); err != nil {
		return err
	}

	driverPath, err := filepath.EvalSymlinks(filepath.Join(c.devicesPath, bdf, "driver"))
	if err != nil {
		return fmt.Errorf("resolve driver for %s: %w", bdf, err)
	}
	if err := os.WriteFile(filepath.Join(driverPath, "unbind"), []byte(bdf), 0o200); err != nil {
		return fmt.Errorf("unbind %s: %w", bdf, err)
	}
	bound := false
	bind := func() error {
		if bound {
			return nil
		}
		if err := os.WriteFile(filepath.Join(driverPath, "bind"), []byte(bdf), 0o200); err != nil {
			return fmt.Errorf("bind %s: %w", bdf, err)
		}
		bound = true
		return nil
	}
	defer func() {
		if err := bind(); err != nil {
			if resultErr == nil {
				resultErr = fmt.Errorf("restore driver for %s: %w", bdf, err)
			} else {
				resultErr = fmt.Errorf("capture failed: %w; restore driver for %s: %w", resultErr, bdf, err)
			}
		}
	}()
	return capture(bind)
}

// CheckSafe refuses disruptive operations on the active default-route NIC.
func (c *DriverController) CheckSafe(bdf string, force bool) error {
	if force {
		return nil
	}
	active, err := c.isDefaultRouteDevice(bdf)
	if err != nil {
		return err
	}
	if active {
		return fmt.Errorf("refusing to disrupt default-route interface at %s; use --force only from a recoverable local session", bdf)
	}
	return nil
}

// Interfaces returns network interfaces exposed by a PCI device.
func (c *DriverController) Interfaces(bdf string) ([]string, error) {
	entries, err := os.ReadDir(filepath.Join(c.devicesPath, bdf, "net"))
	if err != nil {
		return nil, fmt.Errorf("read network interfaces for %s: %w", bdf, err)
	}
	names := make([]string, 0, len(entries))
	for _, entry := range entries {
		names = append(names, entry.Name())
	}
	return names, nil
}

func (c *DriverController) isDefaultRouteDevice(bdf string) (bool, error) {
	entries, err := os.ReadDir(filepath.Join(c.devicesPath, bdf, "net"))
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, fmt.Errorf("read network interfaces for %s: %w", bdf, err)
	}
	interfaces := make(map[string]bool, len(entries))
	for _, entry := range entries {
		interfaces[entry.Name()] = true
	}
	file, err := os.Open(c.routePath)
	if err != nil {
		return false, fmt.Errorf("read routes: %w", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) >= 2 && fields[1] == "00000000" && interfaces[fields[0]] {
			return true, nil
		}
	}
	if err := scanner.Err(); err != nil {
		return false, fmt.Errorf("scan routes: %w", err)
	}
	return false, nil
}
