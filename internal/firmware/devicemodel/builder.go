package devicemodel

import (
	"encoding/binary"
	"fmt"
	"sort"

	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

func New(ctx *donor.DeviceContext) (*Model, error) { return Build(ctx) }

func Build(ctx *donor.DeviceContext) (*Model, error) {
	if ctx == nil {
		return nil, fmt.Errorf("device model: nil donor context")
	}
	if ctx.ConfigSpace == nil {
		return nil, fmt.Errorf("device model: donor config space is missing")
	}

	m := &Model{
		SchemaVersion:       CurrentSchemaVersion,
		Name:                fmt.Sprintf("pci-%04x-%04x", ctx.ConfigSpace.VendorID(), ctx.ConfigSpace.DeviceID()),
		Functions:           []Function{buildFunction(ctx)},
		ConfigSpace:         buildConfigSpace(ctx),
		Capabilities:        buildCapabilities(ctx),
		BARs:                buildBARs(ctx),
		Registers:           buildRegisters(ctx),
		Interrupts:          []InterruptDescriptor{},
		Transformations:     []Transformation{},
		UnsupportedFeatures: buildUnsupported(ctx.ExtCapabilities),
		Confidence:          buildConfidence(ctx),
		Provenance: Provenance{
			Source:      "donor.DeviceContext",
			ToolVersion: ctx.ToolVersion,
			CollectedAt: ctx.CollectedAt,
			DonorBDF:    ctx.Device.BDF.String(),
			Host:        ctx.Hostname,
		},
	}
	m.Interrupts, m.MSIX = buildInterrupts(ctx, m.Capabilities)
	if err := m.Validate(); err != nil {
		return nil, fmt.Errorf("device model: donor context produced invalid model: %w", err)
	}
	return m, nil
}

func buildFunction(ctx *donor.DeviceContext) Function {
	cs := ctx.ConfigSpace
	return Function{
		BDF:               ctx.Device.BDF.String(),
		VendorID:          cs.VendorID(),
		DeviceID:          cs.DeviceID(),
		SubsystemVendorID: cs.SubsysVendorID(),
		SubsystemDeviceID: cs.SubsysDeviceID(),
		RevisionID:        cs.RevisionID(),
		ClassCode:         cs.ClassCode(),
		HeaderType:        cs.HeaderType(),
	}
}

func buildConfigSpace(ctx *donor.DeviceContext) ConfigSpace {
	size := ctx.ConfigSpace.Size
	if size < 0 {
		size = 0
	}
	if size > pci.ConfigSpaceSize {
		size = pci.ConfigSpaceSize
	}
	image := append([]byte(nil), ctx.ConfigSpace.Data[:size]...)
	command := uint64(ctx.ConfigSpace.ReadU16(0x04))
	status := uint64(ctx.ConfigSpace.ReadU16(0x06))
	fields := []ConfigField{
		{Name: "vendor_id", Offset: 0x00, Width: 2, Mask: 0xffff, Access: AccessRO, ResetValue: uint64(ctx.ConfigSpace.ReadU16(0x00))},
		{Name: "device_id", Offset: 0x02, Width: 2, Mask: 0xffff, Access: AccessRO, ResetValue: uint64(ctx.ConfigSpace.ReadU16(0x02))},
		{Name: "command_rw", Offset: 0x04, Width: 2, Mask: 0x077f, Access: AccessRW, ResetValue: command & 0x077f},
		{Name: "command_reserved", Offset: 0x04, Width: 2, Mask: 0xf880, Access: AccessReserved, ResetValue: command & 0xf880},
		{Name: "status_rw1c", Offset: 0x06, Width: 2, Mask: 0xf900, Access: AccessRW1C, ResetValue: status & 0xf900},
		{Name: "status_ro", Offset: 0x06, Width: 2, Mask: 0x06f8, Access: AccessRO, ResetValue: status & 0x06f8},
		{Name: "status_reserved", Offset: 0x06, Width: 2, Mask: 0x0007, Access: AccessReserved, ResetValue: status & 0x0007},
		{Name: "revision_class", Offset: 0x08, Width: 4, Mask: 0xffffffff, Access: AccessRO, ResetValue: uint64(ctx.ConfigSpace.ReadU32(0x08))},
		{Name: "header", Offset: 0x0c, Width: 4, Mask: 0xffffffff, Access: AccessRO, ResetValue: uint64(ctx.ConfigSpace.ReadU32(0x0c))},
	}
	return ConfigSpace{Size: uint32(size), ResetImage: image, Fields: fields}
}

func buildCapabilities(ctx *donor.DeviceContext) []Capability {
	caps := make([]Capability, 0, len(ctx.Capabilities)+len(ctx.ExtCapabilities))
	standardStarts := make([]int, 0, len(ctx.Capabilities))
	for _, capability := range ctx.Capabilities {
		standardStarts = append(standardStarts, capability.Offset)
	}
	sort.Ints(standardStarts)
	for _, c := range ctx.Capabilities {
		data := normalizeCapabilityData(c.Offset, c.Data, standardStarts)
		next := uint16(0)
		if len(data) >= 2 {
			next = uint16(data[1] & 0xfc)
		}
		caps = append(caps, Capability{
			ID: uint16(c.ID), Name: pci.CapabilityName(c.ID), Offset: uint16(c.Offset),
			NextOffset: next, Length: uint16(len(data)), Data: data,
		})
	}
	extendedStarts := make([]int, 0, len(ctx.ExtCapabilities))
	for _, capability := range ctx.ExtCapabilities {
		extendedStarts = append(extendedStarts, capability.Offset)
	}
	sort.Ints(extendedStarts)
	for _, c := range ctx.ExtCapabilities {
		data := normalizeCapabilityData(c.Offset, c.Data, extendedStarts)
		next := uint16(0)
		if len(data) >= 4 {
			next = uint16((binary.LittleEndian.Uint32(data[:4]) >> 20) & 0xffc)
		}
		caps = append(caps, Capability{
			ID: c.ID, Name: pci.ExtCapabilityName(c.ID), Version: c.Version,
			Offset: uint16(c.Offset), NextOffset: next, Length: uint16(len(data)),
			Extended: true, Data: data,
		})
	}
	sort.Slice(caps, func(i, j int) bool {
		if caps[i].Extended != caps[j].Extended {
			return !caps[i].Extended
		}
		return caps[i].Offset < caps[j].Offset
	})
	return caps
}

func normalizeCapabilityData(offset int, data []byte, starts []int) []byte {
	length := len(data)
	for _, start := range starts {
		if start > offset && start-offset < length {
			length = start - offset
			break
		}
	}
	return append([]byte(nil), data[:length]...)
}

func buildBARs(ctx *donor.DeviceContext) []BAR {
	sources := append([]pci.BAR(nil), ctx.BARs...)
	sort.Slice(sources, func(i, j int) bool { return sources[i].Index < sources[j].Index })
	bars := make([]BAR, 0, len(sources))
	for _, source := range sources {
		if source.Type == pci.BARTypeDisabled {
			continue
		}
		t := BARType(source.Type)
		width := uint8(32)
		var pair *int
		if source.Type == pci.BARTypeMem64 || source.Is64Bit && source.Type != pci.BARTypeDisabled {
			t = BARTypeMem64
			width = 64
			upper := source.Index + 1
			pair = &upper
		}
		size := source.Size
		if captured := uint64(len(ctx.BARContents[source.Index])); captured > size {
			size = captured
		}
		if profile := ctx.BARProfiles[source.Index]; profile != nil && profile.Size > 0 && uint64(profile.Size) > size {
			size = uint64(profile.Size)
		}
		image := append([]byte(nil), ctx.BARContents[source.Index]...)
		bars = append(bars, BAR{
			BIR: source.Index, Type: t, Size: size, SizeKnown: size > 0, Prefetchable: source.Prefetchable,
			AddressWidth: width, PairBIR: pair, ResetImage: image,
		})
	}
	sort.Slice(bars, func(i, j int) bool { return bars[i].BIR < bars[j].BIR })
	return bars
}

func buildRegisters(ctx *donor.DeviceContext) []Register {
	regs := buildConfigRegisters(ctx)
	barIndexes := make([]int, 0, len(ctx.BARProfiles))
	for index := range ctx.BARProfiles {
		barIndexes = append(barIndexes, index)
	}
	sort.Ints(barIndexes)
	for _, index := range barIndexes {
		profile := ctx.BARProfiles[index]
		if profile == nil || len(profile.Probes) == 0 {
			continue
		}
		probes := append([]donor.BARProbeResult(nil), profile.Probes...)
		sort.Slice(probes, func(i, j int) bool { return probes[i].Offset < probes[j].Offset })
		for _, probe := range probes {
			regs = append(regs, registerFromProbe(index, probe, ctx.BARContents[index]))
		}
	}
	return regs
}

func buildConfigRegisters(ctx *donor.DeviceContext) []Register {
	cs := ctx.ConfigSpace
	return []Register{
		{Name: "vendor_device", Space: SpaceConfig, BIR: ConfigBIR, Offset: 0, Width: 4, ResetDomain: ResetPowerOn, ResetValue: uint64(cs.ReadU32(0)), Confidence: ConfidenceSpecified, Fields: []RegisterField{{Name: "identity", Mask: 0xffffffff, Access: AccessRO, ResetValue: uint64(cs.ReadU32(0))}}},
		{Name: "command_status", Space: SpaceConfig, BIR: ConfigBIR, Offset: 4, Width: 4, ResetDomain: ResetFundamental, ResetValue: uint64(cs.ReadU32(4)), Confidence: ConfidenceSpecified, Fields: []RegisterField{{Name: "command_rw", Mask: 0x0000077f, Access: AccessRW, ResetValue: uint64(cs.ReadU32(4)) & 0x0000077f}, {Name: "command_reserved", Mask: 0x0000f880, Access: AccessReserved, ResetValue: uint64(cs.ReadU32(4)) & 0x0000f880}, {Name: "status_rw1c", Mask: 0xf9000000, Access: AccessRW1C, ResetValue: uint64(cs.ReadU32(4)) & 0xf9000000}, {Name: "status_ro", Mask: 0x06f80000, Access: AccessRO, ResetValue: uint64(cs.ReadU32(4)) & 0x06f80000}, {Name: "status_reserved", Mask: 0x00070000, Access: AccessReserved, ResetValue: uint64(cs.ReadU32(4)) & 0x00070000}}},
		{Name: "revision_class", Space: SpaceConfig, BIR: ConfigBIR, Offset: 8, Width: 4, ResetDomain: ResetPowerOn, ResetValue: uint64(cs.ReadU32(8)), Confidence: ConfidenceSpecified, Fields: []RegisterField{{Name: "value", Mask: 0xffffffff, Access: AccessRO, ResetValue: uint64(cs.ReadU32(8))}}},
	}
}

func registerFromProbe(index int, probe donor.BARProbeResult, resetImage []byte) Register {
	w1c := uint64(probe.W1CMask)
	rw := uint64(probe.RWMask &^ probe.W1CMask)
	ro := uint64(0xffffffff) &^ (rw | w1c)
	reset := uint64(probe.Original)
	for lane := range 4 {
		imageOffset := int(probe.Offset) + lane
		if imageOffset >= len(resetImage) {
			continue
		}
		mask := uint64(0xff) << (lane * 8)
		reset = reset&^mask | uint64(resetImage[imageOffset])<<(lane*8)
	}
	fields := make([]RegisterField, 0, 3)
	if rw != 0 {
		fields = append(fields, RegisterField{Name: "rw", Mask: rw, Access: AccessRW, ResetValue: reset & rw})
	}
	if w1c != 0 {
		fields = append(fields, RegisterField{Name: "rw1c", Mask: w1c, Access: AccessRW1C, ResetValue: reset & w1c})
	}
	if ro != 0 {
		fields = append(fields, RegisterField{Name: "ro", Mask: ro, Access: AccessRO, ResetValue: reset & ro})
	}
	return Register{
		Name: fmt.Sprintf("bar%d_%04x", index, probe.Offset), Space: SpaceBAR, BIR: index,
		Offset: uint64(probe.Offset), Width: 4, ResetDomain: ResetPowerOn,
		ResetValue: reset, Fields: fields, Confidence: ConfidenceMeasured,
	}
}

func buildInterrupts(ctx *donor.DeviceContext, caps []Capability) ([]InterruptDescriptor, *MSIXDescriptor) {
	interrupts := make([]InterruptDescriptor, 0, 3)
	if pin := ctx.ConfigSpace.InterruptPin(); pin != 0 {
		interrupts = append(interrupts, InterruptDescriptor{Kind: "intx", Vectors: 1, Pin: pin})
	}
	var msixCapOffset uint16
	for _, cap := range caps {
		if cap.Extended {
			continue
		}
		switch uint8(cap.ID) {
		case pci.CapIDMSI:
			vectors := uint16(1)
			if len(cap.Data) >= 4 {
				vectors <<= (binary.LittleEndian.Uint16(cap.Data[2:4]) >> 1) & 0x7
			}
			interrupts = append(interrupts, InterruptDescriptor{Kind: "msi", CapabilityOffset: cap.Offset, Vectors: vectors})
		case pci.CapIDMSIX:
			msixCapOffset = cap.Offset
		}
	}
	if ctx.MSIXData == nil {
		return interrupts, nil
	}
	d := &MSIXDescriptor{
		CapabilityOffset: msixCapOffset, TableSize: uint16(ctx.MSIXData.TableSize),
		TableBIR: ctx.MSIXData.TableBIR, TableOffset: uint64(ctx.MSIXData.TableOffset),
		PBABIR: ctx.MSIXData.PBABIR, PBAOffset: uint64(ctx.MSIXData.PBAOffset),
	}
	interrupts = append(interrupts, InterruptDescriptor{
		Kind: "msix", CapabilityOffset: msixCapOffset, Vectors: d.TableSize,
		BIR: d.TableBIR, TableOffset: d.TableOffset, PBAOffset: d.PBAOffset,
	})
	return interrupts, d
}

func buildUnsupported(caps []pci.ExtCapability) []UnsupportedFeature {
	known := map[uint16][2]string{
		pci.ExtCapIDSRIOV:        {"SR-IOV", "virtual-function lifecycle is not modeled"},
		pci.ExtCapIDMRIOV:        {"MR-IOV", "multi-root routing is not modeled"},
		pci.ExtCapIDResizableBAR: {"Resizable BAR", "dynamic BAR resizing is not modeled"},
		pci.ExtCapIDATS:          {"ATS", "translated DMA requests are not modeled"},
		pci.ExtCapIDPageRequest:  {"Page Request", "page-fault requests are not modeled"},
		pci.ExtCapIDPASID:        {"PASID", "process address spaces are not modeled"},
		pci.ExtCapIDDPC:          {"DPC", "containment signaling is not modeled"},
		pci.ExtCapIDPTM:          {"PTM", "precision time synchronization is not modeled"},
		pci.ExtCapIDMulticast:    {"Multicast", "TLP multicast routing is not modeled"},
	}
	out := make([]UnsupportedFeature, 0)
	for _, cap := range caps {
		if item, ok := known[cap.ID]; ok {
			out = append(out, UnsupportedFeature{Name: item[0], Reason: item[1], Source: fmt.Sprintf("extended capability 0x%04x", cap.ID)})
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Name < out[j].Name })
	return out
}

func buildConfidence(ctx *donor.DeviceContext) Confidence {
	level := ConfidenceInferred
	evidence := []string{"configuration space captured from donor"}
	if len(ctx.BARProfiles) > 0 {
		level = ConfidenceMeasured
		evidence = append(evidence, "BAR access behavior measured by probe")
	} else if len(ctx.BARContents) > 0 {
		evidence = append(evidence, "BAR reset images captured without access probing")
	}
	return Confidence{Overall: level, Evidence: evidence}
}
