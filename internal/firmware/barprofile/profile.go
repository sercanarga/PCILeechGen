package barprofile

import (
	"encoding/binary"
	"fmt"
	"sort"

	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/firmware/devclass"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

const (
	HintProbeModel         = "probe_model"
	HintSpecModel          = "spec_model"
	HintShadowOnly         = "shadow_only"
	HintIOSpaceUnsupported = "io_space_unsupported"

	RegisterDead     = "dead"
	RegisterStatic   = "static"
	RegisterWritable = "writable"
	RegisterRW1C     = "rw1c"
	RegisterAllOnes  = "all_ones"
	RegisterUnknown  = "unknown"
)

type Profile struct {
	ClassCode   uint32       `json:"class_code"`
	DeviceClass string       `json:"device_class,omitempty"`
	BARs        []BARSummary `json:"bars"`
}

type BARSummary struct {
	Index             int               `json:"index"`
	Type              string            `json:"type"`
	Size              uint64            `json:"size"`
	CapturedBytes     int               `json:"captured_bytes"`
	Prefetchable      bool              `json:"prefetchable"`
	Is64Bit           bool              `json:"is_64bit"`
	EmulationHint     string            `json:"emulation_hint"`
	Density           string            `json:"density"`
	NonZeroBytes      int               `json:"non_zero_bytes"`
	NonFFBytes        int               `json:"non_ff_bytes"`
	DeadRegisters     int               `json:"dead_registers"`
	StaticRegisters   int               `json:"static_registers"`
	WritableRegisters int               `json:"writable_registers"`
	RW1CRegisters     int               `json:"rw1c_registers"`
	AllOnesRegisters  int               `json:"all_ones_registers"`
	Registers         []RegisterSummary `json:"registers,omitempty"`
}

type RegisterSummary struct {
	Offset   uint32 `json:"offset"`
	Original uint32 `json:"original"`
	RWMask   uint32 `json:"rw_mask"`
	Kind     string `json:"kind"`
}

func Build(ctx *donor.DeviceContext) *Profile {
	profile := &Profile{}
	if ctx == nil {
		return profile
	}

	profile.ClassCode = ctx.Device.ClassCode
	if strategy := devclass.StrategyForClass(ctx.Device.ClassCode); strategy != nil {
		profile.DeviceClass = strategy.DeviceClass()
	}

	bars := append([]pci.BAR(nil), ctx.BARs...)
	sort.Slice(bars, func(i, j int) bool { return bars[i].Index < bars[j].Index })
	for _, bar := range bars {
		profile.BARs = append(profile.BARs, buildBARSummary(ctx, bar))
	}
	return profile
}

func buildBARSummary(ctx *donor.DeviceContext, bar pci.BAR) BARSummary {
	data := ctx.BARContents[bar.Index]
	probe := ctx.BARProfiles[bar.Index]
	summary := BARSummary{
		Index:         bar.Index,
		Type:          bar.Type,
		Size:          bar.Size,
		CapturedBytes: len(data),
		Prefetchable:  bar.Prefetchable,
		Is64Bit:       bar.Is64Bit,
		Density:       densityFor(data),
		NonZeroBytes:  countNotEqual(data, 0x00),
		NonFFBytes:    countNotEqual(data, 0xFF),
	}
	summary.EmulationHint = emulationHint(bar, data, probe)
	summary.Registers = registerSummaries(data, probe)
	for _, reg := range summary.Registers {
		addRegisterKind(&summary, reg.Kind)
	}
	return summary
}

func emulationHint(bar pci.BAR, data []byte, probe *donor.BARProfile) string {
	if bar.Type == pci.BARTypeIO {
		return HintIOSpaceUnsupported
	}
	if probe != nil && len(probe.Probes) > 0 {
		return HintProbeModel
	}
	if len(data) > 0 {
		return HintSpecModel
	}
	return HintShadowOnly
}

func registerSummaries(data []byte, probe *donor.BARProfile) []RegisterSummary {
	if probe != nil && len(probe.Probes) > 0 {
		return summariesFromProbe(probe)
	}
	return summariesFromData(data)
}

func summariesFromProbe(profile *donor.BARProfile) []RegisterSummary {
	regs := make([]RegisterSummary, 0, len(profile.Probes))
	for _, probe := range profile.Probes {
		regs = append(regs, RegisterSummary{
			Offset:   probe.Offset,
			Original: probe.Original,
			RWMask:   probe.RWMask,
			Kind:     classifyRegister(probe.Original, probe.RWMask, probe.MaybeRW1C),
		})
	}
	sort.Slice(regs, func(i, j int) bool { return regs[i].Offset < regs[j].Offset })
	return regs
}

func summariesFromData(data []byte) []RegisterSummary {
	count := len(data) / 4
	regs := make([]RegisterSummary, 0, count)
	for i := 0; i < count; i++ {
		offset := uint32(i * 4)
		value := binary.LittleEndian.Uint32(data[i*4 : i*4+4])
		regs = append(regs, RegisterSummary{
			Offset:   offset,
			Original: value,
			Kind:     classifyRegister(value, 0, false),
		})
	}
	return regs
}

func classifyRegister(original, rwMask uint32, maybeRW1C bool) string {
	switch {
	case maybeRW1C:
		return RegisterRW1C
	case original == 0 && rwMask == 0:
		return RegisterDead
	case original == 0xFFFFFFFF && rwMask == 0:
		return RegisterAllOnes
	case rwMask != 0:
		return RegisterWritable
	case rwMask == 0:
		return RegisterStatic
	default:
		return RegisterUnknown
	}
}

func addRegisterKind(summary *BARSummary, kind string) {
	switch kind {
	case RegisterDead:
		summary.DeadRegisters++
	case RegisterStatic:
		summary.StaticRegisters++
	case RegisterWritable:
		summary.WritableRegisters++
	case RegisterRW1C:
		summary.RW1CRegisters++
	case RegisterAllOnes:
		summary.AllOnesRegisters++
	}
}

func densityFor(data []byte) string {
	if len(data) == 0 {
		return "empty"
	}
	active := 0
	for _, b := range data {
		if b != 0 && b != 0xFF {
			active++
		}
	}
	ratio := float64(active) / float64(len(data))
	switch {
	case ratio == 0:
		return "blank"
	case ratio < 0.05:
		return "sparse"
	case ratio > 0.75:
		return "dense"
	default:
		return "mixed"
	}
}

func countNotEqual(data []byte, value byte) int {
	count := 0
	for _, b := range data {
		if b != value {
			count++
		}
	}
	return count
}

func (p *Profile) String() string {
	if p == nil {
		return "BAR behavior profile: <nil>"
	}
	return fmt.Sprintf("BAR behavior profile: class=0x%06X bars=%d", p.ClassCode, len(p.BARs))
}
