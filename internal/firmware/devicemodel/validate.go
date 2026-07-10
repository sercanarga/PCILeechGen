package devicemodel

import (
	"fmt"
	"sort"
	"strings"
)

func (m *Model) Validate() error {
	if m == nil {
		return fmt.Errorf("nil model")
	}
	if m.SchemaVersion != CurrentSchemaVersion {
		return fmt.Errorf("unsupported schema_version %d (supported: %d)", m.SchemaVersion, CurrentSchemaVersion)
	}
	if m.Name == "" {
		return fmt.Errorf("name is required")
	}
	if len(m.Functions) != 1 {
		return fmt.Errorf("exactly one function is required")
	}
	if err := validateConfigSpace(m.ConfigSpace); err != nil {
		return err
	}
	if err := validateFunctions(m.Functions, m.ConfigSpace); err != nil {
		return err
	}
	if err := validateCapabilities(m.Capabilities, m.ConfigSpace.Size); err != nil {
		return err
	}
	bars, err := validateBARs(m.BARs)
	if err != nil {
		return err
	}
	if err := validateRegisters(m.Registers, m.ConfigSpace, bars); err != nil {
		return err
	}
	if err := validateInterrupts(m.Interrupts, m.MSIX, m.Capabilities, bars); err != nil {
		return err
	}
	if !validConfidence(m.Confidence.Overall) {
		return fmt.Errorf("invalid confidence level %q", m.Confidence.Overall)
	}
	if strings.TrimSpace(m.Provenance.Source) == "" {
		return fmt.Errorf("provenance source is required")
	}
	if strings.TrimSpace(m.Provenance.ToolVersion) == "" {
		return fmt.Errorf("provenance tool_version is required")
	}
	if strings.TrimSpace(m.Provenance.DonorBDF) == "" {
		return fmt.Errorf("provenance donor_bdf is required")
	}
	if m.Provenance.CollectedAt.IsZero() {
		return fmt.Errorf("provenance collected_at is required")
	}
	return nil
}

func validateFunctions(functions []Function, config ConfigSpace) error {
	seen := make(map[string]bool)
	for i, fn := range functions {
		if fn.BDF == "" {
			return fmt.Errorf("functions[%d]: bdf is required", i)
		}
		if seen[fn.BDF] {
			return fmt.Errorf("functions[%d]: duplicate bdf %q", i, fn.BDF)
		}
		seen[fn.BDF] = true
		if fn.ClassCode > 0xffffff {
			return fmt.Errorf("functions[%d]: class_code exceeds 24 bits", i)
		}
		if config.Size >= 4 {
			identity := readLittleEndian(config.ResetImage, 0, 4)
			if fn.VendorID != uint16(identity) || fn.DeviceID != uint16(identity>>16) {
				return fmt.Errorf("functions[%d]: vendor or device identity does not match config reset image", i)
			}
		}
		if config.Size >= 12 {
			classRevision := readLittleEndian(config.ResetImage, 8, 4)
			if fn.RevisionID != uint8(classRevision) || fn.ClassCode != uint32(classRevision>>8) {
				return fmt.Errorf("functions[%d]: revision or class identity does not match config reset image", i)
			}
		}
		if config.Size >= 15 && fn.HeaderType != config.ResetImage[0x0e] {
			return fmt.Errorf("functions[%d]: header type does not match config reset image", i)
		}
		if config.Size >= 48 {
			subsystem := readLittleEndian(config.ResetImage, 0x2c, 4)
			if fn.SubsystemVendorID != uint16(subsystem) || fn.SubsystemDeviceID != uint16(subsystem>>16) {
				return fmt.Errorf("functions[%d]: subsystem identity does not match config reset image", i)
			}
		}
	}
	return nil
}

func validateConfigSpace(config ConfigSpace) error {
	if config.Size == 0 || config.Size > 4096 {
		return fmt.Errorf("config_space.size %d is outside 1..4096", config.Size)
	}
	if uint32(len(config.ResetImage)) != config.Size {
		return fmt.Errorf("config_space.reset_image length %d does not match size %d", len(config.ResetImage), config.Size)
	}
	occupied := make(map[uint32]uint8)
	for i, field := range config.Fields {
		if field.Name == "" {
			return fmt.Errorf("config_space.fields[%d]: name is required", i)
		}
		if !validWidth(field.Width) {
			return fmt.Errorf("config_space.fields[%d]: width %d is not 1, 2, 4, or 8 bytes", i, field.Width)
		}
		end := uint32(field.Offset) + uint32(field.Width)
		if end > config.Size {
			return fmt.Errorf("config_space.fields[%d]: range exceeds config space", i)
		}
		if field.Mask == 0 || field.Mask&^widthMask(field.Width) != 0 {
			return fmt.Errorf("config_space.fields[%d]: mask 0x%x is invalid for width %d", i, field.Mask, field.Width)
		}
		if !validAccess(field.Access) {
			return fmt.Errorf("config_space.fields[%d]: invalid access policy %q", i, field.Access)
		}
		if field.ResetValue&^field.Mask != 0 {
			return fmt.Errorf("config_space.fields[%d]: reset value sets bits outside mask", i)
		}
		var imageValue uint64
		for lane := range int(field.Width) {
			imageValue |= uint64(config.ResetImage[int(field.Offset)+lane]) << (lane * 8)
		}
		if field.ResetValue != imageValue&field.Mask {
			return fmt.Errorf("config_space.fields[%d]: reset value does not match reset image", i)
		}
		for lane := range int(field.Width) {
			byteMask := uint8(field.Mask >> (lane * 8))
			offset := uint32(field.Offset) + uint32(lane)
			if occupied[offset]&byteMask != 0 {
				return fmt.Errorf("config_space field %q overlaps another field at byte 0x%x", field.Name, offset)
			}
			occupied[offset] |= byteMask
		}
	}
	return nil
}

func validateCapabilities(caps []Capability, configSize uint32) error {
	standard := make(map[uint16]Capability)
	extended := make(map[uint16]Capability)
	for i, cap := range caps {
		if cap.Length == 0 || int(cap.Length) != len(cap.Data) {
			return fmt.Errorf("capabilities[%d]: length %d does not match data length %d", i, cap.Length, len(cap.Data))
		}
		if cap.Offset&3 != 0 {
			return fmt.Errorf("capabilities[%d]: offset 0x%x is not DWORD aligned", i, cap.Offset)
		}
		set := standard
		minimum, limit, minLength := uint16(0x40), configSize, uint16(2)
		if limit > 0x100 {
			limit = 0x100
		}
		if cap.Extended {
			set = extended
			minimum, limit, minLength = 0x100, configSize, 4
			if cap.Version == 0 {
				return fmt.Errorf("capabilities[%d]: extended capability version must be nonzero", i)
			}
		}
		if cap.Offset < minimum || uint32(cap.Offset)+uint32(cap.Length) > limit || cap.Length < minLength {
			return fmt.Errorf("capabilities[%d]: range 0x%x+0x%x is invalid", i, cap.Offset, cap.Length)
		}
		if _, exists := set[cap.Offset]; exists {
			return fmt.Errorf("capabilities[%d]: duplicate offset 0x%x", i, cap.Offset)
		}
		set[cap.Offset] = cap
	}
	if err := validateCapabilityRanges("standard", standard); err != nil {
		return err
	}
	if err := validateCapabilityRanges("extended", extended); err != nil {
		return err
	}

	if err := validateCapabilityLinks("standard", standard); err != nil {
		return err
	}
	if err := validateCapabilityLinks("extended", extended); err != nil {
		return err
	}
	return nil
}

func validateCapabilityRanges(kind string, caps map[uint16]Capability) error {
	ordered := make([]Capability, 0, len(caps))
	for _, cap := range caps {
		ordered = append(ordered, cap)
	}
	sort.Slice(ordered, func(i, j int) bool { return ordered[i].Offset < ordered[j].Offset })
	for i := 1; i < len(ordered); i++ {
		previousEnd := uint32(ordered[i-1].Offset) + uint32(ordered[i-1].Length)
		if uint32(ordered[i].Offset) < previousEnd {
			return fmt.Errorf("%s capabilities at 0x%x and 0x%x overlap", kind, ordered[i-1].Offset, ordered[i].Offset)
		}
	}
	return nil
}

func validateCapabilityLinks(kind string, caps map[uint16]Capability) error {
	for offset, cap := range caps {
		if cap.NextOffset == 0 {
			continue
		}
		if cap.NextOffset&3 != 0 {
			return fmt.Errorf("%s capability at 0x%x has unaligned next pointer 0x%x", kind, offset, cap.NextOffset)
		}
		if _, ok := caps[cap.NextOffset]; !ok {
			return fmt.Errorf("%s capability at 0x%x points to missing capability 0x%x", kind, offset, cap.NextOffset)
		}
	}
	state := make(map[uint16]uint8, len(caps))
	var visit func(uint16) error
	visit = func(offset uint16) error {
		switch state[offset] {
		case 1:
			return fmt.Errorf("%s capability chain contains a cycle at 0x%x", kind, offset)
		case 2:
			return nil
		}
		state[offset] = 1
		if next := caps[offset].NextOffset; next != 0 {
			if err := visit(next); err != nil {
				return err
			}
		}
		state[offset] = 2
		return nil
	}
	for offset := range caps {
		if err := visit(offset); err != nil {
			return err
		}
	}
	return nil
}

func validateBARs(bars []BAR) (map[int]BAR, error) {
	byBIR := make(map[int]BAR, len(bars))
	for i, bar := range bars {
		if bar.BIR < 0 || bar.BIR > 5 {
			return nil, fmt.Errorf("bars[%d]: BIR %d is outside 0..5", i, bar.BIR)
		}
		if _, exists := byBIR[bar.BIR]; exists {
			return nil, fmt.Errorf("bars[%d]: duplicate BIR %d", i, bar.BIR)
		}
		if !validBARType(bar.Type) {
			return nil, fmt.Errorf("bars[%d]: invalid type %q", i, bar.Type)
		}
		if bar.Type == BARTypeDisabled {
			if bar.Size != 0 || bar.SizeKnown || bar.PairBIR != nil {
				return nil, fmt.Errorf("bars[%d]: disabled BAR must have unknown zero size and no pair", i)
			}
		} else if bar.SizeKnown {
			if bar.Size == 0 || bar.Size&(bar.Size-1) != 0 {
				return nil, fmt.Errorf("bars[%d]: known enabled BAR size %d must be a nonzero power of two", i, bar.Size)
			}
			if uint64(len(bar.ResetImage)) > bar.Size {
				return nil, fmt.Errorf("bars[%d]: reset image is larger than BAR", i)
			}
		} else if bar.Size != 0 || len(bar.ResetImage) != 0 {
			return nil, fmt.Errorf("bars[%d]: unknown BAR size must be zero with no reset image", i)
		}
		if bar.Type == BARTypeMem64 {
			if bar.AddressWidth != 64 || bar.PairBIR == nil || *bar.PairBIR != bar.BIR+1 || *bar.PairBIR > 5 {
				return nil, fmt.Errorf("bars[%d]: 64-bit BAR must claim the immediately following BIR", i)
			}
		} else if bar.PairBIR != nil {
			return nil, fmt.Errorf("bars[%d]: only a 64-bit BAR may claim a pair", i)
		} else if bar.Type != BARTypeDisabled && bar.AddressWidth != 32 {
			return nil, fmt.Errorf("bars[%d]: non-64-bit BAR address_width must be 32", i)
		}
		byBIR[bar.BIR] = bar
	}
	for _, bar := range bars {
		if bar.PairBIR == nil {
			continue
		}
		if _, exists := byBIR[*bar.PairBIR]; exists {
			return nil, fmt.Errorf("64-bit BAR%d upper BIR %d must not have an independent descriptor", bar.BIR, *bar.PairBIR)
		}
	}
	return byBIR, nil
}

func validateRegisters(registers []Register, config ConfigSpace, bars map[int]BAR) error {
	type key struct {
		space AddressSpace
		bir   int
	}
	type span struct {
		start, end uint64
		index      int
	}
	groups := make(map[key][]span)
	for i, reg := range registers {
		if reg.Name == "" {
			return fmt.Errorf("registers[%d]: name is required", i)
		}
		if !validWidth(reg.Width) {
			return fmt.Errorf("registers[%d]: width %d is not 1, 2, 4, or 8 bytes", i, reg.Width)
		}
		if reg.Offset%uint64(reg.Width) != 0 {
			return fmt.Errorf("registers[%d]: offset 0x%x is not aligned to width %d", i, reg.Offset, reg.Width)
		}
		if reg.ResetValue&^widthMask(reg.Width) != 0 {
			return fmt.Errorf("registers[%d]: reset value exceeds width", i)
		}
		if !validResetDomain(reg.ResetDomain) {
			return fmt.Errorf("registers[%d]: invalid reset domain %q", i, reg.ResetDomain)
		}
		if !validConfidence(reg.Confidence) {
			return fmt.Errorf("registers[%d]: invalid confidence %q", i, reg.Confidence)
		}
		limit := uint64(config.Size)
		resetImage := config.ResetImage
		switch reg.Space {
		case SpaceConfig:
			if reg.BIR != ConfigBIR {
				return fmt.Errorf("registers[%d]: config register BIR must be %d", i, ConfigBIR)
			}
		case SpaceBAR:
			bar, ok := bars[reg.BIR]
			if !ok || bar.Type == BARTypeDisabled {
				return fmt.Errorf("registers[%d]: references missing or disabled BAR%d", i, reg.BIR)
			}
			limit = bar.Size
			resetImage = bar.ResetImage
		default:
			return fmt.Errorf("registers[%d]: invalid address space %q", i, reg.Space)
		}
		end := reg.Offset + uint64(reg.Width)
		if end < reg.Offset || end > limit {
			return fmt.Errorf("registers[%d]: range exceeds its address space", i)
		}
		for lane := range int(reg.Width) {
			imageOffset := int(reg.Offset) + lane
			if imageOffset >= len(resetImage) {
				continue
			}
			actual := uint8(reg.ResetValue >> (lane * 8))
			if actual != resetImage[imageOffset] {
				return fmt.Errorf("registers[%d]: reset value does not match address-space reset image", i)
			}
		}
		var occupied uint64
		for j, field := range reg.Fields {
			if field.Name == "" || field.Mask == 0 || field.Mask&^widthMask(reg.Width) != 0 {
				return fmt.Errorf("registers[%d].fields[%d]: invalid name or mask", i, j)
			}
			if occupied&field.Mask != 0 {
				return fmt.Errorf("registers[%d].fields[%d]: field mask overlaps another field", i, j)
			}
			if !validAccess(field.Access) {
				return fmt.Errorf("registers[%d].fields[%d]: invalid access policy %q", i, j, field.Access)
			}
			if field.ResetValue&^field.Mask != 0 {
				return fmt.Errorf("registers[%d].fields[%d]: reset value sets bits outside mask", i, j)
			}
			if field.ResetValue != reg.ResetValue&field.Mask {
				return fmt.Errorf("registers[%d].fields[%d]: reset value does not match register reset value", i, j)
			}
			occupied |= field.Mask
		}
		groups[key{reg.Space, reg.BIR}] = append(groups[key{reg.Space, reg.BIR}], span{reg.Offset, end, i})
	}
	for _, spans := range groups {
		sort.Slice(spans, func(i, j int) bool { return spans[i].start < spans[j].start })
		for i := 1; i < len(spans); i++ {
			if spans[i].start < spans[i-1].end {
				return fmt.Errorf("registers[%d] overlaps registers[%d]", spans[i].index, spans[i-1].index)
			}
		}
	}
	return nil
}

func validateInterrupts(interrupts []InterruptDescriptor, msix *MSIXDescriptor, caps []Capability, bars map[int]BAR) error {
	for i, irq := range interrupts {
		switch irq.Kind {
		case "intx", "msi", "msix":
		default:
			return fmt.Errorf("interrupts[%d]: unsupported kind %q", i, irq.Kind)
		}
		if irq.Vectors == 0 {
			return fmt.Errorf("interrupts[%d]: vectors must be nonzero", i)
		}
	}
	if msix == nil {
		return nil
	}
	if msix.TableSize == 0 || msix.TableSize > 2048 {
		return fmt.Errorf("msix.table_size %d is outside 1..2048", msix.TableSize)
	}
	if msix.CapabilityOffset != 0 {
		found := false
		for _, cap := range caps {
			if !cap.Extended && cap.ID == 0x11 && cap.Offset == msix.CapabilityOffset {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("msix capability offset 0x%x does not reference an MSI-X capability", msix.CapabilityOffset)
		}
	}
	tableBAR, ok := bars[msix.TableBIR]
	if !ok || (tableBAR.Type != BARTypeMem32 && tableBAR.Type != BARTypeMem64) {
		return fmt.Errorf("msix table BIR %d does not reference a memory BAR", msix.TableBIR)
	}
	pbaBAR, ok := bars[msix.PBABIR]
	if !ok || (pbaBAR.Type != BARTypeMem32 && pbaBAR.Type != BARTypeMem64) {
		return fmt.Errorf("msix PBA BIR %d does not reference a memory BAR", msix.PBABIR)
	}
	if msix.TableOffset&7 != 0 || msix.PBAOffset&7 != 0 {
		return fmt.Errorf("msix table and PBA offsets must be 8-byte aligned")
	}
	tableEnd := msix.TableOffset + uint64(msix.TableSize)*16
	pbaBytes := uint64((uint32(msix.TableSize)+63)/64) * 8
	pbaEnd := msix.PBAOffset + pbaBytes
	if tableEnd < msix.TableOffset || (tableBAR.SizeKnown && tableEnd > tableBAR.Size) {
		return fmt.Errorf("msix table range exceeds BAR%d", msix.TableBIR)
	}
	if pbaEnd < msix.PBAOffset || (pbaBAR.SizeKnown && pbaEnd > pbaBAR.Size) {
		return fmt.Errorf("msix PBA range exceeds BAR%d", msix.PBABIR)
	}
	if msix.TableBIR == msix.PBABIR && msix.TableOffset < pbaEnd && msix.PBAOffset < tableEnd {
		return fmt.Errorf("msix table and PBA ranges overlap")
	}
	return nil
}

func readLittleEndian(data []byte, offset, width int) uint64 {
	var value uint64
	for lane := range width {
		value |= uint64(data[offset+lane]) << (lane * 8)
	}
	return value
}

func validWidth(width uint8) bool { return width == 1 || width == 2 || width == 4 || width == 8 }

func widthMask(width uint8) uint64 {
	if width == 8 {
		return ^uint64(0)
	}
	return uint64(1)<<(width*8) - 1
}

func validAccess(access AccessPolicy) bool {
	switch access {
	case AccessRO, AccessRW, AccessRW1C, AccessW1S, AccessW0C, AccessRC, AccessReserved:
		return true
	default:
		return false
	}
}

func validResetDomain(domain ResetDomain) bool {
	switch domain {
	case ResetPowerOn, ResetFundamental, ResetFunction, ResetSoftware:
		return true
	default:
		return false
	}
}

func validConfidence(level ConfidenceLevel) bool {
	switch level {
	case ConfidenceUnknown, ConfidenceInferred, ConfidenceMeasured, ConfidenceSpecified:
		return true
	default:
		return false
	}
}

func validBARType(t BARType) bool {
	return t == BARTypeIO || t == BARTypeMem32 || t == BARTypeMem64 || t == BARTypeDisabled
}
