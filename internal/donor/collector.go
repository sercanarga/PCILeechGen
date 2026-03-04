package donor

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/sercanarga/pcileechgen/internal/donor/vfio"
	"github.com/sercanarga/pcileechgen/internal/pci"
	"github.com/sercanarga/pcileechgen/internal/version"
)

// Collector gathers donor PCI data from sysfs.
type Collector struct {
	sysfs *SysfsReader
}

func NewCollector() *Collector {
	return &Collector{
		sysfs: NewSysfsReader(),
	}
}

// NewCollectorWithSysfs lets tests inject a fake sysfs reader.
func NewCollectorWithSysfs(sr *SysfsReader) *Collector {
	return &Collector{
		sysfs: sr,
	}
}

// Collect reads config space, BARs, and caps from the device.
func (c *Collector) Collect(bdf pci.BDF) (*DeviceContext, error) {
	ctx := &DeviceContext{
		CollectedAt: time.Now(),
		ToolVersion: version.Version,
	}

	hostname, _ := os.Hostname()
	ctx.Hostname = hostname

	// basic PCI info
	dev, err := c.sysfs.ReadDeviceInfo(bdf)
	if err != nil {
		return nil, fmt.Errorf("failed to read device info for %s: %w", bdf, err)
	}
	ctx.Device = *dev

	// config space — try sysfs, fall back to VFIO
	cs, err := c.sysfs.ReadConfigSpace(bdf)
	if err != nil {
		log.Printf("[donor] sysfs config read failed, trying VFIO: %v", err)
		if bindErr := vfio.BindToVFIO(bdf.String()); bindErr != nil {
			return nil, fmt.Errorf("failed to read config space for %s (sysfs: %v, VFIO bind: %v)", bdf, err, bindErr)
		}
		cs, err = c.sysfs.ReadConfigSpace(bdf)
		if err != nil {
			return nil, fmt.Errorf("failed to read config space for %s even after VFIO bind: %w", bdf, err)
		}
	}
	ctx.ConfigSpace = cs

	// BAR layout from sysfs resource file
	bars, err := c.sysfs.ReadResourceFile(bdf)
	if err != nil {
		// fall back to parsing from config space
		bars = pci.ParseBARsFromConfigSpace(cs)
	}
	ctx.BARs = bars

	// dump BAR memory — 4KB cap to match pcileech BRAM
	const maxBARReadSize = 4096
	ctx.BARContents = make(map[int][]byte)
	sysfsBarFailed := false
	for _, bar := range bars {
		if bar.IsDisabled() || bar.IsIO() || bar.Size == 0 {
			continue
		}
		data, err := c.sysfs.ReadBARContent(bdf, bar.Index, maxBARReadSize)
		if err != nil {
			log.Printf("[donor] Warning: could not read BAR%d via sysfs: %v", bar.Index, err)
			sysfsBarFailed = true
			continue
		}
		ctx.BARContents[bar.Index] = data
		log.Printf("[donor] Read %d bytes from BAR%d", len(data), bar.Index)
	}

	// VFIO fallback for failed BAR reads
	if sysfsBarFailed && vfio.IsBoundToVFIO(bdf.String()) {
		log.Println("[donor] Trying VFIO for failed BAR reads...")
		if dump, err := vfio.Collect(bdf.String()); err == nil {
			for idx, data := range dump.BARContents {
				if _, already := ctx.BARContents[idx]; !already && len(data) > 0 {
					ctx.BARContents[idx] = data
					log.Printf("[donor] Read %d bytes from BAR%d via VFIO", len(data), idx)
				}
			}
		} else {
			log.Printf("[donor] VFIO BAR fallback failed: %v", err)
		}
	}

	// probe BAR regs for RW/RO masks (Linux only, best-effort)
	profiler := NewBARProfiler()
	ctx.BARProfiles = make(map[int]*BARProfile)
	for _, bar := range bars {
		if bar.IsDisabled() || bar.IsIO() || bar.Size == 0 {
			continue
		}
		resourcePath := fmt.Sprintf("%s/%s/resource%d", c.sysfs.basePath, bdf.String(), bar.Index)
		profile, err := profiler.ProfileBAR(resourcePath, bar.Index, maxBARReadSize)
		if err != nil {
			log.Printf("[donor] BAR%d profiling skipped: %v", bar.Index, err)
			continue
		}
		ctx.BARProfiles[bar.Index] = profile
		activeRegs := 0
		for _, p := range profile.Probes {
			if p.Original != 0 || p.RWMask != 0 {
				activeRegs++
			}
		}
		log.Printf("[donor] Profiled BAR%d: %d active registers detected", bar.Index, activeRegs)
	}

	// caps
	ctx.Capabilities = pci.ParseCapabilities(cs)
	ctx.ExtCapabilities = pci.ParseExtCapabilities(cs)

	return ctx, nil
}

// SaveContext dumps a DeviceContext to JSON on disk.
func SaveContext(ctx *DeviceContext, path string) error {
	data, err := ctx.ToJSON()
	if err != nil {
		return fmt.Errorf("failed to marshal device context: %w", err)
	}
	return os.WriteFile(path, data, 0644)
}

// LoadContext restores a DeviceContext from a JSON file.
func LoadContext(path string) (*DeviceContext, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read device context file: %w", err)
	}
	return FromJSON(data)
}
