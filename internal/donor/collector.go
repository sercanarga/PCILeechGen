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

func (c *Collector) Collect(bdf pci.BDF) (*DeviceContext, error) {
	ctx := &DeviceContext{
		CollectedAt: time.Now(),
		ToolVersion: version.Version,
	}

	hostname, _ := os.Hostname()
	ctx.Hostname = hostname

	dev, err := c.sysfs.ReadDeviceInfo(bdf)
	if err != nil {
		return nil, fmt.Errorf("failed to read device info for %s: %w", bdf, err)
	}
	ctx.Device = *dev

	cs, err := c.collectConfigSpace(bdf)
	if err != nil {
		return nil, err
	}
	ctx.ConfigSpace = cs

	bars, err := c.sysfs.ReadResourceFile(bdf)
	if err != nil {
		bars = pci.ParseBARsFromConfigSpace(cs)
	}
	ctx.BARs = bars

	ctx.BARContents = c.collectBARMemory(bdf, bars)
	ctx.BARProfiles = c.collectBARProfiles(bdf, bars)
	ctx.Capabilities = pci.ParseCapabilities(cs)
	ctx.ExtCapabilities = pci.ParseExtCapabilities(cs)
	ctx.MSIXData = c.collectMSIXData(cs, ctx.BARContents)

	if err := c.validateBARContents(ctx); err != nil {
		return nil, err
	}

	return ctx, nil
}

// barCriticalClass returns true for device classes where empty BAR contents
// will cause the Windows driver to fail with Code 10.
func barCriticalClass(classCode uint32) bool {
	switch classCode {
	case 0x010802: // NVMe
		return true
	case 0x0C0330: // xHCI USB 3.0
		return true
	}
	return false
}

func (c *Collector) validateBARContents(ctx *DeviceContext) error {
	eligible := eligibleBARs(ctx.BARs)
	if len(eligible) == 0 {
		return nil
	}

	hasContent := false
	for _, bar := range eligible {
		if data, ok := ctx.BARContents[bar.Index]; ok && len(data) > 0 {
			hasContent = true
			break
		}
	}

	if !hasContent {
		hint := "Make sure the device is bound to vfio-pci and retry"
		if barCriticalClass(ctx.Device.ClassCode) {
			hint += " - this device class is known to produce Code 10 without BAR data"
		}
		return fmt.Errorf(
			"BAR content collection failed for %d eligible BAR(s) (class 0x%06X, driver %q). "+
				"Without BAR register data the firmware will likely produce Code 10 in Windows. %s",
			len(eligible), ctx.Device.ClassCode, ctx.Device.Driver, hint,
		)
	}

	return nil
}

func (c *Collector) collectConfigSpace(bdf pci.BDF) (*pci.ConfigSpace, error) {
	cs, err := c.sysfs.ReadConfigSpace(bdf)
	if err != nil {
		log.Printf("[donor] sysfs config read failed, trying VFIO: %v", err)
		if bindErr := vfio.BindToVFIO(bdf.String()); bindErr != nil {
			return nil, fmt.Errorf("config space read failed for %s (sysfs: %v, VFIO: %v)", bdf, err, bindErr)
		}
		cs, err = c.sysfs.ReadConfigSpace(bdf)
		if err != nil {
			return nil, fmt.Errorf("config space read failed for %s even after VFIO bind: %w", bdf, err)
		}
	}
	return cs, nil
}

func (c *Collector) collectBARMemory(bdf pci.BDF, bars []pci.BAR) map[int][]byte {
	const maxBARReadSize = 4096
	contents := make(map[int][]byte)
	sysfsBarFailed := false

	for _, bar := range eligibleBARs(bars) {
		data, err := c.sysfs.ReadBARContent(bdf, bar.Index, maxBARReadSize)
		if err != nil {
			log.Printf("[donor] Warning: could not read BAR%d via sysfs: %v", bar.Index, err)
			sysfsBarFailed = true
			continue
		}
		contents[bar.Index] = data
		log.Printf("[donor] Read %d bytes from BAR%d", len(data), bar.Index)
	}

	if sysfsBarFailed && vfio.IsBoundToVFIO(bdf.String()) {
		log.Println("[donor] Trying VFIO for failed BAR reads...")
		if dump, err := vfio.Collect(bdf.String()); err == nil {
			for idx, data := range dump.BARContents {
				if _, already := contents[idx]; !already && len(data) > 0 {
					contents[idx] = data
					log.Printf("[donor] Read %d bytes from BAR%d via VFIO", len(data), idx)
				}
			}
		} else {
			log.Printf("[donor] VFIO BAR fallback failed: %v", err)
		}
	}

	return contents
}

func (c *Collector) collectBARProfiles(bdf pci.BDF, bars []pci.BAR) map[int]*BARProfile {
	const maxBARReadSize = 4096
	profiler := NewBARProfiler()
	profiles := make(map[int]*BARProfile)

	for _, bar := range eligibleBARs(bars) {
		resourcePath := fmt.Sprintf("%s/%s/resource%d", c.sysfs.basePath, bdf.String(), bar.Index)
		profile, err := profiler.ProfileBAR(resourcePath, bar.Index, maxBARReadSize)
		if err != nil {
			log.Printf("[donor] BAR%d profiling skipped: %v", bar.Index, err)
			continue
		}
		profiles[bar.Index] = profile
		activeRegs := 0
		for _, p := range profile.Probes {
			if p.Original != 0 || p.RWMask != 0 {
				activeRegs++
			}
		}
		log.Printf("[donor] Profiled BAR%d: %d active registers detected", bar.Index, activeRegs)
	}

	return profiles
}

// eligibleBARs filters out disabled, IO, and zero-size BARs.
func eligibleBARs(bars []pci.BAR) []pci.BAR {
	var result []pci.BAR
	for _, bar := range bars {
		if !bar.IsDisabled() && !bar.IsIO() && bar.Size > 0 {
			result = append(result, bar)
		}
	}
	return result
}

// collectMSIXData reads MSI-X table entries from BAR memory.
func (c *Collector) collectMSIXData(cs *pci.ConfigSpace, barContents map[int][]byte) *MSIXData {
	info := pci.ParseMSIXCap(cs)
	if info == nil {
		return nil
	}

	barData := barContents[info.TableBIR]
	entries := pci.ReadMSIXTable(barData, info)

	log.Printf("[donor] MSI-X: %d vectors, table BAR%d+0x%X, PBA BAR%d+0x%X",
		info.TableSize, info.TableBIR, info.TableOffset, info.PBABIR, info.PBAOffset)

	return &MSIXData{
		TableSize:   info.TableSize,
		TableBIR:    info.TableBIR,
		TableOffset: info.TableOffset,
		PBABIR:      info.PBABIR,
		PBAOffset:   info.PBAOffset,
		Entries:     entries,
	}
}
