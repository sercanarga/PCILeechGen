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
	nativeDriverBindTimeout   = 10 * time.Second
	nvmeLiveTimeout           = 15 * time.Second
	nativeDriverPollInterval  = 200 * time.Millisecond
	nativeDriverBARRetryDelay = 500 * time.Millisecond
	nativeDriverBARRetries    = 3
	memSpaceSettleDelay       = 200 * time.Millisecond
	maxBARReadSize            = 65536
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

	// One shared native-driver visit for the whole Collect() (see runNativeVisit):
	// the BAR fallback and the NVMe identity capture reuse a single rebind cycle.
	visit := &nativeVisitCache{bdf: bdf, bars: bars, nvme: isNVMeClass(ctx.Device.ClassCode)}

	ctx.BARContents = c.collectBARMemory(bdf, bars, visit)
	ctx.BARProfiles = c.collectBARProfiles(bdf, bars, ctx.BARContents)
	ctx.Capabilities = pci.ParseCapabilities(cs)
	ctx.ExtCapabilities = pci.ParseExtCapabilities(cs)
	ctx.MSIXData = c.collectMSIXData(cs, ctx.BARContents)
	ctx.NVMeIdentity = c.collectNVMeIdentity(bdf, ctx.Device.ClassCode, visit)

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

	barsWithData := 0
	barsWithFF := 0
	hasValidData := false
	for _, bar := range eligible {
		data, ok := ctx.BARContents[bar.Index]
		if !ok || len(data) == 0 {
			continue
		}
		barsWithData++
		if isAllFF(data) {
			barsWithFF++
			if barCriticalClass(ctx.Device.ClassCode) {
				slog.Warn("BAR content is all 0xFF - device may be inaccessible or in D3 power state",
					"bar", bar.Index, "bytes", len(data),
					"class", fmt.Sprintf("0x%06X", ctx.Device.ClassCode),
				)
			} else {
				slog.Warn("BAR content is all 0xFF DMA will continue with zeroed registers",
					"bar", bar.Index, "bytes", len(data),
				)
			}
		} else {
			hasValidData = true
		}
	}

	allMemoryBARsFF := barsWithData > 0 && barsWithFF == barsWithData

	if allMemoryBARsFF && barCriticalClass(ctx.Device.ClassCode) {
		return fmt.Errorf(
			"All %d eligible memory BAR(s) returned 0xFF for class 0x%06X (driver %q). "+
				"The device is not responding - possible causes:\n"+
				"  • device is in D3 (sleep) power state\n"+
				"  • IOMMU/VT-d not enabled or misconfigured\n"+
				"  • PCI Command Register memory space not enabled\n"+
				"  • device requires native driver initialization (e.g. CNVi WiFi, firmware-dependent devices)\n"+
				"Without valid BAR data, Windows will produce Code 10 and DMA will not work\n\n"+
				"Workarounds:\n"+
				"  1. Try: sudo setpci -s %s COMMAND=0x06\n"+
				"  2. For CNVi/WiFi: boot with native driver, dump BAR with a tool, then use --from-json",
			barsWithData, ctx.Device.ClassCode, ctx.Device.Driver, ctx.Device.BDF,
		)
	}

	if barsWithFF > 0 && hasValidData && barCriticalClass(ctx.Device.ClassCode) {
		slog.Warn("some BARs returned 0xFF but valid data exists on other BARs proceeding with valid BAR data",
			"ff_bars", barsWithFF, "total_bars", barsWithData,
			"class", fmt.Sprintf("0x%06X", ctx.Device.ClassCode),
		)
	}

	if barsWithData == 0 {
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

type nativeVisitResult struct {
	bars   map[int][]byte
	nvmeID *NVMeIdentity
}

// nativeVisitCache runs the native-driver visit at most once per Collect, so
// the BAR fallback and the NVMe identity capture share a single rebind cycle.
type nativeVisitCache struct {
	bdf  pci.BDF
	bars []pci.BAR
	nvme bool
	done bool
	res  nativeVisitResult
	err  error
}

func isNVMeClass(classCode uint32) bool {
	return classCode>>16 == 0x01 && (classCode>>8)&0xFF == 0x08
}

func (c *Collector) runNativeVisit(vc *nativeVisitCache) (nativeVisitResult, error) {
	if !vc.done {
		vc.done = true
		vc.res, vc.err = c.captureViaNativeDriver(vc.bdf, vc.bars, vc.nvme)
	}
	return vc.res, vc.err
}

// restoreVFIO rebinds vfio-pci and re-enables memory space + D0. Only a rebind
// failure is returned; the rest is best-effort.
func (c *Collector) restoreVFIO(bdf pci.BDF) error {
	if err := vfio.BindToVFIO(bdf.String()); err != nil {
		return fmt.Errorf("rebind vfio-pci: %w", err)
	}
	if err := vfio.EnableMemorySpace(bdf.String()); err != nil {
		slog.Warn("could not re-enable memory space after rebind", "bdf", bdf, "error", err)
	} else {
		time.Sleep(memSpaceSettleDelay)
	}
	if err := vfio.WakeToD0(bdf.String()); err != nil {
		slog.Warn("could not wake device to D0 after rebind", "bdf", bdf, "error", err)
	} else {
		time.Sleep(memSpaceSettleDelay)
	}
	return nil
}

// restoreVFIOOrWarn rebinds vfio-pci and warns on failure, so the device is not
// left unbound silently.
func (c *Collector) restoreVFIOOrWarn(bdf pci.BDF) {
	if err := c.restoreVFIO(bdf); err != nil {
		slog.Warn("could not restore vfio-pci; device may be unbound", "bdf", bdf, "error", err)
	}
}

// captureViaNativeDriver runs the single native-driver visit per Collect so the
// native driver initializes the device (live BAR/identity reads need it under
// vfio-pci). Returns an error only when no native driver binds.
func (c *Collector) captureViaNativeDriver(bdf pci.BDF, bars []pci.BAR, nvme bool) (nativeVisitResult, error) {
	res := nativeVisitResult{bars: make(map[int][]byte)}

	if err := vfio.UnbindFromVFIO(bdf.String()); err != nil {
		c.restoreVFIOOrWarn(bdf)
		return res, fmt.Errorf("native visit: %w", err)
	}

	drv, err := vfio.WaitForNativeDriver(bdf.String(), nativeDriverBindTimeout, nativeDriverPollInterval)
	if err != nil {
		c.restoreVFIOOrWarn(bdf)
		return res, fmt.Errorf("native visit: %w", err)
	}
	slog.Info("native driver bound", "bdf", bdf, "driver", drv)

	if nvme {
		if liveErr := vfio.WaitForNVMeLive(bdf.String(), nvmeLiveTimeout, nativeDriverPollInterval); liveErr != nil {
			slog.Warn("nvme controller not live; identity capture may fail", "bdf", bdf, "error", liveErr)
		}
	}

	for _, bar := range eligibleBARs(bars) {
		readLen := int(bar.Size)
		if readLen > maxBARReadSize {
			readLen = maxBARReadSize
		}
		data, readErr := c.readBARUntilValid(bdf, bar.Index, readLen)
		if readErr != nil {
			slog.Warn("native visit: BAR read failed", "bar", bar.Index, "error", readErr)
			continue
		}
		if isAllFF(data) {
			slog.Warn("native visit: BAR still 0xFF after native bind", "bar", bar.Index)
			continue
		}
		res.bars[bar.Index] = data
		slog.Info("native visit: BAR captured", "bar", bar.Index, "bytes", len(data))
	}

	if nvme {
		if id, idErr := c.sysfs.ReadNVMeIdentity(bdf); idErr != nil {
			slog.Warn("native visit: NVMe identity read failed; firmware will use synthesized strings", "error", idErr)
		} else {
			res.nvmeID = id
			slog.Info("native visit: NVMe identity captured",
				"model", id.Model, "serial", id.Serial, "firmware", id.FWRev)
		}
	}

	if err := c.restoreVFIO(bdf); err != nil {
		return res, fmt.Errorf("native visit: %w", err)
	}
	return res, nil
}

// readBARUntilValid reads a BAR a few times with a settle delay between tries;
// BAR registers can lag slightly behind driver bind on cold controllers.
func (c *Collector) readBARUntilValid(bdf pci.BDF, index, readLen int) ([]byte, error) {
	var data []byte
	var err error
	for attempt := 0; attempt < nativeDriverBARRetries; attempt++ {
		if attempt > 0 {
			time.Sleep(nativeDriverBARRetryDelay)
		}
		data, err = c.sysfs.ReadBARContent(bdf, index, readLen)
		if err == nil && len(data) > 0 && !isAllFF(data) {
			return data, nil
		}
	}
	return data, err
}

// collectNVMeIdentity captures NVMe controller strings for NVMe-class devices.
// Returns nil for non-NVMe devices or when capture fails (firmware then uses
// synthesized strings).
func (c *Collector) collectNVMeIdentity(bdf pci.BDF, classCode uint32, vc *nativeVisitCache) *NVMeIdentity {
	if !isNVMeClass(classCode) {
		return nil
	}

	if id, err := c.sysfs.ReadNVMeIdentity(bdf); err == nil {
		slog.Info("captured NVMe identity from bound driver",
			"model", id.Model, "serial", id.Serial, "firmware", id.FWRev)
		return id
	}

	if vc == nil {
		slog.Warn("NVMe identity unavailable (nvme driver not bound); firmware will use synthesized strings")
		return nil
	}
	res, err := c.runNativeVisit(vc)
	if err != nil {
		slog.Warn("NVMe identity capture via native visit failed; firmware will use synthesized strings", "error", err)
		return nil
	}
	if res.nvmeID == nil {
		slog.Warn("NVMe identity not captured; firmware will use synthesized strings")
		return nil
	}
	slog.Info("captured NVMe identity via native visit",
		"model", res.nvmeID.Model, "serial", res.nvmeID.Serial, "firmware", res.nvmeID.FWRev)
	return res.nvmeID
}

func (c *Collector) collectConfigSpace(bdf pci.BDF) (*pci.ConfigSpace, error) {
	cs, err := c.sysfs.ReadConfigSpace(bdf)
	if err != nil {
		slog.Info("sysfs config read failed, trying VFIO", "error", err)
		if bindErr := vfio.BindToVFIO(bdf.String()); bindErr != nil {
			return nil, fmt.Errorf("config space read failed for %s (sysfs: %w, VFIO: %w)", bdf, err, bindErr)
		}
		cs, err = c.sysfs.ReadConfigSpace(bdf)
		if err != nil {
			return nil, fmt.Errorf("config space read failed for %s even after VFIO bind: %w", bdf, err)
		}
	}
	return cs, nil
}

func (c *Collector) collectBARMemory(bdf pci.BDF, bars []pci.BAR, vc *nativeVisitCache) map[int][]byte {
	eligible := eligibleBARs(bars)
	if len(eligible) == 0 {
		return nil
	}

	contents := make(map[int][]byte)

	// wake device from D3 if needed, otherwise BAR reads return 0xFF
	if err := vfio.WakeToD0(bdf.String()); err != nil {
		slog.Warn("could not wake device to D0", "bdf", bdf, "error", err)
	}

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
		_ = vfio.WakeToD0(bdf.String())
		c.readBARs(bdf, eligible, contents)
		if !memBARsAllFF(contents, bars) {
			return contents
		}
	}

	// last resort: let the native driver init the device (single shared visit).
	// avoid VFIO sessions here, session close can trigger FLR.
	if vfio.IsBoundToVFIO(bdf.String()) {
		slog.Info("all memory BAR contents are 0xFF, attempting native driver visit")
		res, vErr := c.runNativeVisit(vc)
		if vErr != nil {
			slog.Warn("native driver visit failed; keeping prior BAR contents", "error", vErr)
		} else {
			for idx, data := range res.bars {
				contents[idx] = data
			}
			if len(res.bars) > 0 {
				slog.Info("native visit recovered BAR data", "bars_recovered", len(res.bars))
			} else {
				slog.Warn("native visit did not recover any BAR data")
			}
		}
	}

	return contents
}

// readBARs reads eligible BARs via sysfs mmap, skipping already-valid entries.
// Contents are capped to maxBARReadSize to avoid extremely slow/ hanging reads
// on devices with large BAR apertures (hundreds of MB+). The low registers
// needed for emulation are captured; higher areas are not required for the
// initial snapshot.
func (c *Collector) readBARs(bdf pci.BDF, eligible []pci.BAR, contents map[int][]byte) {
	for _, bar := range eligible {
		if data, ok := contents[bar.Index]; ok && !isAllFF(data) {
			continue // already have valid data
		}
		readLen := int(bar.Size)
		if readLen > maxBARReadSize {
			readLen = maxBARReadSize
		}
		data, err := c.sysfs.ReadBARContent(bdf, bar.Index, readLen)
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
		profile, err := profiler.ProfileBAR(resourcePath, bar.Index, min(int(bar.Size), maxBARReadSize))
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
