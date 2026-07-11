package session

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/sercanarga/pcileechgen/internal/donor/mmio"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

// CaptureSource provides host data for one initialization session.
type CaptureSource interface {
	Device() (*pci.PCIDevice, error)
	Config() (*pci.ConfigSpace, error)
	BARs() ([]pci.BAR, error)
	Resources() ([]byte, error)
	Interrupts() ([]byte, error)
	Trace() (*mmio.TraceResult, error)
	Driver() string
	Kernel() string
}

// Capture records before/after state and a trace in one checksummed session directory.
func Capture(dir string, scenario Scenario, source CaptureSource) (*Manifest, error) {
	if source == nil {
		return nil, fmt.Errorf("capture source is nil")
	}
	started := time.Now().UTC()
	device, err := source.Device()
	if err != nil {
		return nil, fmt.Errorf("read device: %w", err)
	}
	before, err := source.Config()
	if err != nil {
		return nil, fmt.Errorf("read config before trace: %w", err)
	}
	bars, err := source.BARs()
	if err != nil {
		return nil, fmt.Errorf("read BARs: %w", err)
	}
	resources, err := source.Resources()
	if err != nil {
		return nil, fmt.Errorf("read resources: %w", err)
	}
	interruptsBefore, err := source.Interrupts()
	if err != nil {
		return nil, fmt.Errorf("read interrupts before trace: %w", err)
	}
	trace, err := source.Trace()
	if err != nil {
		return nil, fmt.Errorf("capture trace: %w", err)
	}
	after, err := source.Config()
	if err != nil {
		return nil, fmt.Errorf("read config after trace: %w", err)
	}
	interruptsAfter, err := source.Interrupts()
	if err != nil {
		return nil, fmt.Errorf("read interrupts after trace: %w", err)
	}
	traceJSON, err := json.MarshalIndent(trace, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("encode trace: %w", err)
	}
	manifest := &Manifest{
		Version: 1, Scenario: scenario, BDF: device.BDF.String(), Driver: source.Driver(), Kernel: source.Kernel(),
		StartedAt: started, Duration: time.Since(started), Device: *device, BARs: bars,
	}
	files := map[string][]byte{
		"config-before.bin":     append([]byte(nil), before.Data[:before.Size]...),
		"config-after.bin":      append([]byte(nil), after.Data[:after.Size]...),
		"resources.txt":         resources,
		"interrupts-before.txt": interruptsBefore,
		"interrupts-after.txt":  interruptsAfter,
		"trace.json":            append(traceJSON, '\n'),
	}
	if err := Save(dir, manifest, files); err != nil {
		return nil, err
	}
	return manifest, nil
}

// LoadCapture verifies and loads one captured initialization session.
func LoadCapture(dir string) (CaptureData, error) {
	manifest, err := Load(dir)
	if err != nil {
		return CaptureData{}, err
	}
	if verifyErr := Verify(dir, manifest); verifyErr != nil {
		return CaptureData{}, verifyErr
	}
	before, err := os.ReadFile(filepath.Join(dir, "config-before.bin"))
	if err != nil {
		return CaptureData{}, fmt.Errorf("read config before: %w", err)
	}
	after, err := os.ReadFile(filepath.Join(dir, "config-after.bin"))
	if err != nil {
		return CaptureData{}, fmt.Errorf("read config after: %w", err)
	}
	traceData, err := os.ReadFile(filepath.Join(dir, "trace.json"))
	if err != nil {
		return CaptureData{}, fmt.Errorf("read trace: %w", err)
	}
	var trace mmio.TraceResult
	if err := json.Unmarshal(traceData, &trace); err != nil {
		return CaptureData{}, fmt.Errorf("decode trace: %w", err)
	}
	return CaptureData{
		Manifest: manifest, Trace: &trace,
		ConfigBefore: pci.NewConfigSpaceFromBytes(before),
		ConfigAfter:  pci.NewConfigSpaceFromBytes(after),
	}, nil
}
