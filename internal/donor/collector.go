package donor

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/sercanarga/pcileechgen/internal/donor/vfio"
	"github.com/sercanarga/pcileechgen/internal/pci"
	"github.com/sercanarga/pcileechgen/internal/version"
)

const (
	nativeDriverInitDelay  = 5 * time.Second
	nativeDriverRetryDelay = 1 * time.Second
	nativeDriverMaxRetries = 3
	memSpaceSettleDelay    = 100 * time.Millisecond
	maxBARReadSize         = 4096
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
	ctx.BARProfiles = c.collectBARProfiles(bdf, bars, ctx.BARContents)
	ctx.Capabilities = pci.ParseCapabilities(cs)
	ctx.ExtCapabilities = pci.ParseExtCapabilities(cs)
	ctx.MSIXData = c.collectMSIXData(cs, ctx.BARContents)

	if err := c.validateBARContents(ctx); err != nil {
		return nil, err
	}

	warnDeviceCompatibility(ctx)

	return ctx, nil
}

// barCriticalClass returns true for device classes where empty BAR contents
// will cause the Windows driver to fail with Code 10.
func barCriticalClass(classCode uint32) bool {
	baseClass := (classCode >> 16) & 0xFF
	subClass := (classCode >> 8) & 0xFF
	switch {
	case baseClass == 0x01 && subClass == 0x08: // NVMe
		return true
	case baseClass == 0x0C && subClass == 0x03: // xHCI USB 3.0
		return true
	case baseClass == 0x02 && subClass == 0x00: // Ethernet
		return true
	case baseClass == 0x02 && subClass == 0x80: // WiFi / CNVi
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
	allFF := false
	for _, bar := range eligible {
		data, ok := ctx.BARContents[bar.Index]
		if !ok || len(data) == 0 {
			continue
		}
		hasContent = true
		if isAllFF(data) {
			allFF = true
			slog.Warn("BAR content is all 0xFF - device may be inaccessible or in D3 power state",
				"bar", bar.Index, "bytes", len(data),
				"class", fmt.Sprintf("0x%06X", ctx.Device.ClassCode),
			)
		}
	}

	if allFF && barCriticalClass(ctx.Device.ClassCode) {
		return fmt.Errorf(
			"BAR content for class 0x%06X is all 0xFF (driver %q). "+
				"The device is not responding - possible causes:\n"+
				"  • device is in D3 (sleep) power state\n"+
				"  • IOMMU/VT-d not enabled or misconfigured\n"+
				"  • PCI Command Register memory space not enabled\n"+
				"  • device requires native driver initialization (e.g. CNVi WiFi, firmware-dependent devices)\n"+
				"Without valid BAR data, Windows will produce Code 10 and DMA will not work\n\n"+
				"Workarounds:\n"+
				"  1. Try: sudo setpci -s %s COMMAND=0x06\n"+
				"  2. For CNVi/WiFi: boot with native driver, dump BAR with a tool, then use --from-json",
			ctx.Device.ClassCode, ctx.Device.Driver, ctx.Device.BDF,
		)
	}

	if !hasContent {
		msg := fmt.Sprintf(
			"BAR content collection failed for %d eligible BAR(s) (class 0x%06X, driver %q)",
			len(eligible), ctx.Device.ClassCode, ctx.Device.Driver,
		)
		if barCriticalClass(ctx.Device.ClassCode) {
			return fmt.Errorf(
				"%s. This device class requires BAR data - without it Windows will produce Code 10. "+
					"Make sure the device is bound to vfio-pci and retry", msg,
			)
		}
		slog.Warn("proceeding without BAR content - firmware will use zeroed BAR registers",
			"eligible_bars", len(eligible),
			"class", fmt.Sprintf("0x%06X", ctx.Device.ClassCode),
		)
	}

	return nil
}

// warnDeviceCompatibility checks for device traits that may cause
// DMA issues and logs warnings so the user knows before synthesis.
func warnDeviceCompatibility(ctx *DeviceContext) {
	hasMSI := false
	hasPCIe := false
	for _, cap := range ctx.Capabilities {
		switch cap.ID {
		case pci.CapIDMSI, pci.CapIDMSIX:
			hasMSI = true
		case pci.CapIDPCIExpress:
			hasPCIe = true
		}
	}

	hasMemBAR := false
	for _, bar := range ctx.BARs {
		if !bar.IsDisabled() && !bar.IsIO() && bar.Size > 0 {
			hasMemBAR = true
			break
		}
	}

	if !hasMemBAR {
		slog.Warn("device has no memory BARs (IO-only) - FPGA firmware cannot emulate IO space, DMA will likely not work")
	}

	if !hasMSI {
		slog.Warn("device has no MSI/MSI-X capability - legacy INTx interrupts may not work over PCIe DMA")
	}

	if !hasPCIe {
		slog.Warn("device has no PCIe capability - this is a legacy PCI device behind a bridge, compatibility may be limited")
	}
}

// isAllFF returns true when every byte in the slice is 0xFF.
func isAllFF(data []byte) bool {
	if len(data) == 0 {
		return false
	}
	for _, b := range data {
		if b != 0xFF {
			return false
		}
	}
	return true
}

// memBARsAllFF checks whether all eligible memory BAR contents are 0xFF.
// it ignores IO BARs so a valid IO BAR doesn't mask broken memory BARs.
func memBARsAllFF(contents map[int][]byte, bars []pci.BAR) bool {
	checked := 0
	for _, bar := range eligibleBARs(bars) {
		data, ok := contents[bar.Index]
		if !ok || len(data) == 0 {
			continue
		}
		checked++
		if !isAllFF(data) {
			return false
		}
	}
	return checked > 0
}

// tryNativeDriverRebind temporarily switches to the native driver so the
// device can initialize its BARs, then switches back to vfio-pci.
func (c *Collector) tryNativeDriverRebind(bdf pci.BDF, bars []pci.BAR, current map[int][]byte) map[int][]byte {
	if err := vfio.UnbindFromVFIO(bdf.String()); err != nil {
		slog.Warn("native driver rebind: unbind failed", "error", err)
		return current
	}

	slog.Info("waiting for native driver to initialize device...",
		"delay", nativeDriverInitDelay)
	time.Sleep(nativeDriverInitDelay)

	recovered := make(map[int][]byte)
	eligible := eligibleBARs(bars)

	for attempt := 0; attempt <= nativeDriverMaxRetries; attempt++ {
		if attempt > 0 {
			slog.Info("native driver rebind: retrying BAR reads",
				"attempt", attempt, "max", nativeDriverMaxRetries)
			time.Sleep(nativeDriverRetryDelay)
		}

		for _, bar := range eligible {
			if _, ok := recovered[bar.Index]; ok {
				continue // already recovered
			}
			data, err := c.sysfs.ReadBARContent(bdf, bar.Index, maxBARReadSize)
			if err != nil {
				slog.Warn("native driver rebind: BAR read failed",
					"bar", bar.Index, "error", err)
				continue
			}
			if !isAllFF(data) && len(data) > 0 {
				recovered[bar.Index] = data
				slog.Info("native driver rebind: BAR read success",
					"bar", bar.Index, "bytes", len(data))
			}
		}

		if len(recovered) == len(eligible) {
			break // all BARs recovered
		}
	}

	// rebind to vfio-pci; discard recovered data on failure
	if err := vfio.BindToVFIO(bdf.String()); err != nil {
		slog.Warn("native driver rebind: re-bind to vfio-pci failed, discarding recovered BARs",
			"error", err)
		return current
	}

	// vfio-pci starts with memory space disabled
	if err := vfio.EnableMemorySpace(bdf.String()); err != nil {
		slog.Warn("native driver rebind: could not re-enable memory space", "error", err)
	} else {
		time.Sleep(memSpaceSettleDelay)
	}

	if len(recovered) > 0 {
		for idx, data := range recovered {
			current[idx] = data
		}
		slog.Info("native driver rebind cycle succeeded",
			"bars_recovered", len(recovered))
	} else {
		slog.Warn("native driver rebind cycle did not recover any BAR data")
	}

	return current
}

func (c *Collector) collectConfigSpace(bdf pci.BDF) (*pci.ConfigSpace, error) {
	cs, err := c.sysfs.ReadConfigSpace(bdf)
	if err != nil {
		slog.Info("sysfs config read failed, trying VFIO", "error", err)
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
	eligible := eligibleBARs(bars)
	if len(eligible) == 0 {
		return nil
	}

	contents := make(map[int][]byte)

	// try reading without touching config space first.
	// works on most devices since BIOS/previous driver left memory space on.
	c.readBARs(bdf, eligible, contents)
	if !memBARsAllFF(contents, bars) {
		return contents
	}

	// didn't work, try enabling memory space (vfio-pci clears it on bind).
	if vfio.IsBoundToVFIO(bdf.String()) {
		if err := vfio.EnableMemorySpace(bdf.String()); err != nil {
			slog.Warn("could not enable PCI memory space",
				"bdf", bdf, "error", err)
		} else {
			slog.Info("PCI memory space enabled", "bdf", bdf)
			time.Sleep(memSpaceSettleDelay)
		}
		c.readBARs(bdf, eligible, contents)
		if !memBARsAllFF(contents, bars) {
			return contents
		}
	}

	// last resort: let native driver init the device.
	// avoid VFIO sessions here, session close can trigger FLR.
	if vfio.IsBoundToVFIO(bdf.String()) {
		slog.Info("all memory BAR contents are 0xFF, attempting native driver rebind cycle")
		contents = c.tryNativeDriverRebind(bdf, bars, contents)
	}

	return contents
}

// readBARs reads eligible BARs via sysfs mmap, skipping already-valid entries.
func (c *Collector) readBARs(bdf pci.BDF, eligible []pci.BAR, contents map[int][]byte) {
	for _, bar := range eligible {
		if data, ok := contents[bar.Index]; ok && !isAllFF(data) {
			continue // already have valid data
		}
		data, err := c.sysfs.ReadBARContent(bdf, bar.Index, maxBARReadSize)
		if err != nil {
			slog.Warn("could not read BAR via sysfs", "bar", bar.Index, "error", err)
			continue
		}
		if isAllFF(data) {
			slog.Warn("sysfs BAR read returned all 0xFF", "bar", bar.Index)
		}
		contents[bar.Index] = data
		slog.Info("BAR read complete", "bar", bar.Index, "bytes", len(data))
	}
}

func (c *Collector) collectBARProfiles(bdf pci.BDF, bars []pci.BAR, barContents map[int][]byte) map[int]*BARProfile {
	profiler := NewBARProfiler()
	profiles := make(map[int]*BARProfile)

	for _, bar := range eligibleBARs(bars) {
		// don't probe unresponsive BARs, writes can brick the device
		if data, ok := barContents[bar.Index]; ok && isAllFF(data) {
			slog.Info("BAR profiling skipped: content is all 0xFF", "bar", bar.Index)
			continue
		}

		resourcePath := fmt.Sprintf("%s/%s/resource%d", c.sysfs.basePath, bdf.String(), bar.Index)
		profile, err := profiler.ProfileBAR(resourcePath, bar.Index, maxBARReadSize)
		if err != nil {
			slog.Info("BAR profiling skipped", "bar", bar.Index, "error", err)
			continue
		}
		profiles[bar.Index] = profile
		activeRegs := 0
		for _, p := range profile.Probes {
			if p.Original != 0 || p.RWMask != 0 {
				activeRegs++
			}
		}
		slog.Info("BAR profiled", "bar", bar.Index, "active_registers", activeRegs)
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

	slog.Info("MSI-X detected",
		"vectors", info.TableSize,
		"table_bar", info.TableBIR,
		"table_offset", fmt.Sprintf("0x%X", info.TableOffset),
		"pba_bar", info.PBABIR,
		"pba_offset", fmt.Sprintf("0x%X", info.PBAOffset),
	)

	return &MSIXData{
		TableSize:   info.TableSize,
		TableBIR:    info.TableBIR,
		TableOffset: info.TableOffset,
		PBABIR:      info.PBABIR,
		PBAOffset:   info.PBAOffset,
		Entries:     entries,
	}
}
