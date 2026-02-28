package donor

import (
	"fmt"
	"os"
	"time"

	"github.com/sercanarga/pcileechgen/internal/pci"
	"github.com/sercanarga/pcileechgen/internal/version"
)

// Collector reads donor PCI device data via sysfs and VFIO.
type Collector struct {
	sysfs *SysfsReader
	vfio  *VFIOManager
}

// NewCollector creates a new Collector.
func NewCollector() *Collector {
	return &Collector{
		sysfs: NewSysfsReader(),
		vfio:  NewVFIOManager(),
	}
}

// NewCollectorWithSysfs creates a Collector with a custom sysfs reader (for testing).
func NewCollectorWithSysfs(sr *SysfsReader) *Collector {
	return &Collector{
		sysfs: sr,
		vfio:  NewVFIOManager(),
	}
}

// Collect reads config space, BARs, and capabilities from the given device.
func (c *Collector) Collect(bdf pci.BDF) (*DeviceContext, error) {
	ctx := &DeviceContext{
		CollectedAt: time.Now(),
		ToolVersion: version.Version,
	}

	hostname, _ := os.Hostname()
	ctx.Hostname = hostname

	// Read basic device info
	dev, err := c.sysfs.ReadDeviceInfo(bdf)
	if err != nil {
		return nil, fmt.Errorf("failed to read device info for %s: %w", bdf, err)
	}
	ctx.Device = *dev

	// Read config space
	cs, err := c.sysfs.ReadConfigSpace(bdf)
	if err != nil {
		return nil, fmt.Errorf("failed to read config space for %s: %w", bdf, err)
	}
	ctx.ConfigSpace = cs

	// Read BAR information from sysfs resource file
	bars, err := c.sysfs.ReadResourceFile(bdf)
	if err != nil {
		// Fall back to parsing from config space
		bars = pci.ParseBARsFromConfigSpace(cs)
	}
	ctx.BARs = bars

	// Parse capabilities
	ctx.Capabilities = pci.ParseCapabilities(cs)
	ctx.ExtCapabilities = pci.ParseExtCapabilities(cs)

	return ctx, nil
}

// SaveContext saves a DeviceContext to a JSON file.
func SaveContext(ctx *DeviceContext, path string) error {
	data, err := ctx.ToJSON()
	if err != nil {
		return fmt.Errorf("failed to marshal device context: %w", err)
	}
	return os.WriteFile(path, data, 0644)
}

// LoadContext loads a DeviceContext from a JSON file.
func LoadContext(path string) (*DeviceContext, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read device context file: %w", err)
	}
	return FromJSON(data)
}
