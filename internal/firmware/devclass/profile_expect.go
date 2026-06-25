package devclass

import (
	"fmt"

	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

// SchemaVersion is the version of the profile-expectation schema. Bump when
// the Expectation shape or validation rules change incompatibly so callers
// can detect stale expectations.
const SchemaVersion = 1

// unsupportedBehaviorCaps lists PCIe extended capabilities the codegen does
// not emulate today. If a donor advertises any of these, the validator emits
// a warning so operators know the generated firmware will not replicate the
// behavior. This list is conservative and must shrink only when a capability
// is actually implemented and tested.
var unsupportedBehaviorCaps = []uint16{
	pci.ExtCapIDATS,
	pci.ExtCapIDPASID,
	pci.ExtCapIDPageRequest,
	pci.ExtCapIDDPC,
	pci.ExtCapIDPTM,
}

// Expectation describes what the generator expects a donor device to look
// like for a given class profile. It is derived from DeviceProfile plus the
// class code that selected the profile. Expectations are conservative: they
// reflect what the generator actually supports today, never unimplemented
// protocols.
type Expectation struct {
	SchemaVersion int
	ClassName     string

	// BaseClass/SubClass are the class bytes the strategy dispatched on.
	// ponytail: prog-if is ignored by the strategy table today; ceiling is
	// per-prog-if profiles when a device family needs them.
	BaseClass uint8
	SubClass  uint8

	// BAR expectations. ExpectedBARCount is the minimum enabled BARs the
	// driver expects. PreferredBAR is the register window the profile
	// targets (BAR0 by default).
	ExpectedBARCount   int
	PreferredBAR       int
	MinBARSize         int
	Expect64BitBAR     bool
	ExpectPrefetchable bool

	// PrefersMSIX means the driver path expects MSI-X; its absence is warned.
	PrefersMSIX bool

	// ExpectedCaps / ExpectedExtCaps are capabilities the profile assumes.
	ExpectedCaps    []uint8
	ExpectedExtCaps []uint16

	// UnsupportedBehaviors lists PCIe protocol features the codegen does not
	// emulate. Their presence on a donor is warned, never silently dropped.
	UnsupportedBehaviors []uint16
}

// Warning is a non-blocking profile validation finding. Severity is always
// "warning"; the field exists so future info-level notes can reuse the type
// without reshaping callers.
type Warning struct {
	Code     string
	Severity string
	Message  string
}

// String renders a stable, grep-friendly single-line form.
func (w Warning) String() string {
	return fmt.Sprintf("[profile:%s] %s", w.Code, w.Message)
}

// ExpectationForDevice derives a conservative profile expectation from the
// devclass strategy table. Returns nil for the generic fallback (unknown
// classes), which means "no profile, skipping" — the validator then produces
// no warnings for that donor.
func ExpectationForDevice(classCode uint32, vendorID, deviceID uint16) *Expectation {
	s := StrategyForDevice(classCode, vendorID, deviceID)
	if s == nil || s.DeviceClass() == ClassGeneric {
		return nil
	}
	p := s.Profile()
	if p == nil {
		return nil
	}
	return &Expectation{
		SchemaVersion:        SchemaVersion,
		ClassName:            p.ClassName,
		BaseClass:            uint8((classCode >> 16) & 0xFF),
		SubClass:             uint8((classCode >> 8) & 0xFF),
		ExpectedBARCount:     1,
		PreferredBAR:         p.PreferredBAR,
		MinBARSize:           p.MinBARSize,
		Expect64BitBAR:       p.Uses64BitBAR,
		ExpectPrefetchable:   p.BARIsPrefetchable,
		PrefersMSIX:          p.PrefersMSIX,
		ExpectedCaps:         p.ExpectedCaps,
		ExpectedExtCaps:      p.ExpectedExtCaps,
		UnsupportedBehaviors: unsupportedBehaviorCaps,
	}
}

// Validate compares a donor DeviceContext against the matching profile
// expectation and returns non-blocking warnings. It never returns an error:
// profile mismatches are advisory only and must not block builds. Unknown
// device classes (no profile) produce zero warnings.
func Validate(ctx *donor.DeviceContext) []Warning {
	if ctx == nil {
		return nil
	}
	exp := ExpectationForDevice(ctx.Device.ClassCode, ctx.Device.VendorID, ctx.Device.DeviceID)
	if exp == nil {
		return nil
	}
	var ws []Warning

	// Class match.
	if ctx.Device.BaseClass() != exp.BaseClass || ctx.Device.SubClass() != exp.SubClass {
		ws = append(ws, Warning{
			Code:     "profile.class.mismatch",
			Severity: "warning",
			Message: fmt.Sprintf("device class %02x%02x does not match %s profile expectation %02x%02x",
				ctx.Device.BaseClass(), ctx.Device.SubClass(),
				exp.ClassName, exp.BaseClass, exp.SubClass),
		})
	}

	// BAR count + preferred BAR type/size.
	enabled := 0
	var preferred *pci.BAR
	for i := range ctx.BARs {
		b := &ctx.BARs[i]
		if b.IsDisabled() {
			continue
		}
		enabled++
		if b.Index == exp.PreferredBAR {
			preferred = b
		}
	}
	if exp.ExpectedBARCount > 0 && enabled < exp.ExpectedBARCount {
		ws = append(ws, Warning{
			Code:     "profile.bar.count",
			Severity: "warning",
			Message:  fmt.Sprintf("%s expects at least %d enabled BAR(s), found %d", exp.ClassName, exp.ExpectedBARCount, enabled),
		})
	}
	if preferred != nil {
		if exp.Expect64BitBAR && preferred.Type != pci.BARTypeMem64 {
			ws = append(ws, Warning{
				Code:     "profile.bar.type",
				Severity: "warning",
				Message:  fmt.Sprintf("%s expects 64-bit BAR%d, found %s", exp.ClassName, exp.PreferredBAR, preferred.Type),
			})
		}
		if exp.MinBARSize > 0 && preferred.Size < uint64(exp.MinBARSize) {
			ws = append(ws, Warning{
				Code:     "profile.bar.size",
				Severity: "warning",
				Message:  fmt.Sprintf("%s BAR%d size %d below profile minimum %d", exp.ClassName, exp.PreferredBAR, preferred.Size, exp.MinBARSize),
			})
		}
	}

	// Capabilities.
	have := make(map[uint8]bool, len(ctx.Capabilities))
	for _, c := range ctx.Capabilities {
		have[c.ID] = true
	}
	for _, id := range exp.ExpectedCaps {
		if !have[id] {
			ws = append(ws, Warning{
				Code:     "profile.cap.missing",
				Severity: "warning",
				Message:  fmt.Sprintf("%s expects capability 0x%02x (%s) not advertised by donor", exp.ClassName, id, pci.CapabilityName(id)),
			})
		}
	}

	// MSI-X expectation.
	if exp.PrefersMSIX {
		hasMSIXCap := have[pci.CapIDMSIX]
		hasMSIXData := ctx.MSIXData != nil && ctx.MSIXData.TableSize > 0
		if !hasMSIXCap && !hasMSIXData {
			ws = append(ws, Warning{
				Code:     "profile.msix.missing",
				Severity: "warning",
				Message:  fmt.Sprintf("%s prefers MSI-X but donor does not advertise it", exp.ClassName),
			})
		}
	}

	// Unsupported behaviors: warn if donor advertises protocols codegen does
	// not emulate. These are never silently dropped.
	haveExt := make(map[uint16]bool, len(ctx.ExtCapabilities))
	for _, c := range ctx.ExtCapabilities {
		haveExt[c.ID] = true
	}
	for _, id := range exp.UnsupportedBehaviors {
		if haveExt[id] {
			ws = append(ws, Warning{
				Code:     "profile.unsupported.behavior",
				Severity: "warning",
				Message: fmt.Sprintf("donor advertises %s (0x%04x) which the generator does not emulate; firmware will not replicate this behavior",
					pci.ExtCapabilityName(id), id),
			})
		}
	}

	return ws
}