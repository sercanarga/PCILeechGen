package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/sercanarga/pcileechgen/internal/color"
	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/donor/mmio"
	"github.com/sercanarga/pcileechgen/internal/donor/session"
	"github.com/sercanarga/pcileechgen/internal/pci"
	"golang.org/x/sys/unix"
)

type linuxSessionSource struct {
	reader   *donor.SysfsReader
	bdf      pci.BDF
	opts     mmioTraceOptions
	barBase  uint64
	scenario session.Scenario
	force    bool
	driver   string
	kernel   string
}

func newLinuxSessionSource(bdf pci.BDF, opts mmioTraceOptions, barBase uint64, scenario session.Scenario, force bool) *linuxSessionSource {
	kernel, _ := os.ReadFile("/proc/sys/kernel/osrelease")
	return &linuxSessionSource{reader: donor.NewSysfsReader(), bdf: bdf, opts: opts, barBase: barBase, scenario: scenario, force: force, kernel: strings.TrimSpace(string(kernel))}
}

func (s *linuxSessionSource) Device() (*pci.PCIDevice, error) {
	device, err := s.reader.ReadDeviceInfo(s.bdf)
	if err == nil {
		s.driver = device.Driver
	}
	return device, err
}

func (s *linuxSessionSource) Config() (*pci.ConfigSpace, error) {
	return s.reader.ReadConfigSpace(s.bdf)
}

func (s *linuxSessionSource) BARs() ([]pci.BAR, error) {
	return s.reader.ReadResourceFile(s.bdf)
}

func (s *linuxSessionSource) Resources() ([]byte, error) {
	return os.ReadFile(filepath.Join("/sys/bus/pci/devices", s.bdf.String(), "resource"))
}

func (s *linuxSessionSource) Interrupts() ([]byte, error) {
	return os.ReadFile("/proc/interrupts")
}

func (s *linuxSessionSource) Trace() (*mmio.TraceResult, error) {
	if s.scenario == session.ScenarioTrace {
		return loadMMIOTrace(s.opts, s.barBase)
	}
	if s.opts.traceFile != "" {
		return nil, fmt.Errorf("--scenario %s requires live tracing", s.scenario)
	}
	controller := session.NewLinuxDriverController()
	switch s.scenario {
	case session.ScenarioUnbindBind:
		var trace *mmio.TraceResult
		err := controller.RunUnbindBindCapture(s.bdf.String(), s.force, func(bind func() error) error {
			var err error
			trace, err = s.captureTraceWithAction(bind)
			return err
		})
		return trace, err
	case session.ScenarioInterfaceUp:
		if err := controller.CheckSafe(s.bdf.String(), s.force); err != nil {
			return nil, err
		}
		interfaces, err := controller.Interfaces(s.bdf.String())
		if err != nil {
			return nil, err
		}
		sort.Strings(interfaces)
		if len(interfaces) != 1 {
			return nil, fmt.Errorf("interface-up scenario requires exactly one network interface, found %d", len(interfaces))
		}
		flags, err := interfaceFlags(interfaces[0])
		if err != nil {
			return nil, err
		}
		if err := setInterfaceFlags(interfaces[0], flags&^uint16(unix.IFF_UP)); err != nil {
			return nil, err
		}
		defer func() { _ = setInterfaceFlags(interfaces[0], flags) }()
		return s.captureTraceWithAction(func() error {
			return setInterfaceFlags(interfaces[0], flags|uint16(unix.IFF_UP))
		})
	default:
		return nil, fmt.Errorf("unsupported capture scenario %q", s.scenario)
	}
}

func (s *linuxSessionSource) captureTraceWithAction(action func() error) (*mmio.TraceResult, error) {
	type result struct {
		trace *mmio.TraceResult
		err   error
	}
	done := make(chan result, 1)
	go func() {
		trace, err := loadMMIOTrace(s.opts, s.barBase)
		done <- result{trace: trace, err: err}
	}()
	time.Sleep(100 * time.Millisecond)
	actionErr := action()
	captured := <-done
	if actionErr != nil {
		return nil, actionErr
	}
	return captured.trace, captured.err
}

func interfaceFlags(name string) (uint16, error) {
	fd, err := unix.Socket(unix.AF_INET, unix.SOCK_DGRAM, 0)
	if err != nil {
		return 0, fmt.Errorf("open interface control socket: %w", err)
	}
	defer unix.Close(fd)
	request, err := unix.NewIfreq(name)
	if err != nil {
		return 0, err
	}
	if err := unix.IoctlIfreq(fd, unix.SIOCGIFFLAGS, request); err != nil {
		return 0, fmt.Errorf("read flags for %s: %w", name, err)
	}
	return request.Uint16(), nil
}

func setInterfaceFlags(name string, flags uint16) error {
	fd, err := unix.Socket(unix.AF_INET, unix.SOCK_DGRAM, 0)
	if err != nil {
		return fmt.Errorf("open interface control socket: %w", err)
	}
	defer unix.Close(fd)
	request, err := unix.NewIfreq(name)
	if err != nil {
		return err
	}
	request.SetUint16(flags)
	if err := unix.IoctlIfreq(fd, unix.SIOCSIFFLAGS, request); err != nil {
		return fmt.Errorf("set flags for %s: %w", name, err)
	}
	return nil
}

func (s *linuxSessionSource) Driver() string { return s.driver }
func (s *linuxSessionSource) Kernel() string { return s.kernel }

func runSessionAnalysis(out io.Writer, dirs []string, jsonOutput bool) error {
	captures := make([]session.CaptureData, 0, len(dirs))
	for _, dir := range dirs {
		capture, err := session.LoadCapture(dir)
		if err != nil {
			return fmt.Errorf("load session %q: %w", dir, err)
		}
		captures = append(captures, capture)
	}
	analysis := session.AnalyzeInitialization(captures)
	if jsonOutput {
		encoder := json.NewEncoder(out)
		encoder.SetIndent("", "  ")
		if err := encoder.Encode(analysis); err != nil {
			return fmt.Errorf("render session analysis: %w", err)
		}
		return nil
	}
	printInitializationAnalysis(out, analysis)
	return nil
}

func printInitializationAnalysis(out io.Writer, analysis session.Analysis) {
	fmt.Fprintln(out, color.Header("Initialization Session Analysis"))
	fmt.Fprintf(out, "Sessions: %d | phases: %d | registers: %d\n\n", analysis.SessionCount, len(analysis.Phases), len(analysis.Registers))
	fmt.Fprintln(out, "Offset      Reads  Writes  Class       Confidence  Effect")
	for _, register := range analysis.Registers {
		fmt.Fprintf(out, "0x%08x  %-6d %-7d %-11s %-11s %s\n", register.Offset, register.Reads, register.Writes, register.Classification, register.Confidence, register.WriteEffect)
	}
	if len(analysis.Dependencies) > 0 {
		fmt.Fprintln(out)
		fmt.Fprintln(out, "Observed write/poll dependencies:")
		for _, dependency := range analysis.Dependencies {
			fmt.Fprintf(out, "  0x%08x -> 0x%08x (%d sessions)\n", dependency.WriteOffset, dependency.ReadOffset, dependency.Occurrences)
		}
	}
	if len(analysis.ConfigChanges) > 0 {
		fmt.Fprintln(out)
		fmt.Fprintln(out, "Configuration changes:")
		for _, change := range analysis.ConfigChanges {
			fmt.Fprintf(out, "  0x%03x: 0x%08x -> 0x%08x (mask 0x%08x)\n", change.Offset, change.Before, change.After, change.Mask)
		}
	}
}
