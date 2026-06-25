package donor

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/sercanarga/pcileechgen/internal/pci"
	"github.com/sercanarga/pcileechgen/internal/version"
)

// Warning is a structured, non-fatal import advisory. Importers return
// warnings for missing optional donor fields (BAR contents, MSI-X table,
// capabilities) so the existing `validate` command — not the importer —
// decides whether a context is build-ready.
type Warning struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (w Warning) String() string { return fmt.Sprintf("%s: %s", w.Code, w.Message) }

// Accepted config-space capture lengths.
const (
	configSpaceLegacyLen    = pci.ConfigSpaceLegacySize // 256 bytes
	configSpaceExtendedLen  = pci.ConfigSpaceSize       // 4096 bytes
)

// ImportConfigSpace converts a raw little-endian PCI config-space capture
// (256 bytes standard or 4096 bytes extended) into a donor.DeviceContext.
//
// Errors are returned only for malformed/truncated INPUT. Missing optional
// donor data (BAR contents, MSI-X entries, capabilities) is reported via the
// returned warnings slice, not as an error — the caller (typically the
// `validate` command) decides whether the context is build-ready.
//
// The returned context never carries HDL; import stops at DeviceContext JSON.
func ImportConfigSpace(raw []byte) (*DeviceContext, []Warning, error) {
	n := len(raw)
	if n != configSpaceLegacyLen && n != configSpaceExtendedLen {
		return nil, nil, fmt.Errorf("import: config space truncated: got %d bytes, want %d or %d", n, configSpaceLegacyLen, configSpaceExtendedLen)
	}

	cs := pci.NewConfigSpaceFromBytes(raw)
	ctx, warnings := buildContextFromConfigSpace(cs)
	return ctx, warnings, nil
}

// ImportCOE parses a COE write-mask / config-space capture emitted by
// codegen.GenerateConfigSpaceCOE (radix=16, little-endian DWORD vector) and
// converts it into a donor.DeviceContext. Only the generator's own COE format
// is supported — round-trip only. Ambiguous third-party COE variants must be
// added with their own dedicated parser.
//
// ponytail: ceiling = per-format importer registry; upgrade path = add a
// `format` flag and dispatch table when a second COE dialect lands.
func ImportCOE(path string) (*DeviceContext, []Warning, error) {
	words, err := parseCOEFile(path)
	if err != nil {
		return nil, nil, err
	}
	if len(words) != configSpaceExtendedLen/4 {
		return nil, nil, fmt.Errorf("import: COE vector length %d DWORDs does not match 4KB config space", len(words))
	}

	raw := make([]byte, configSpaceExtendedLen)
	for i, w := range words {
		// Config-space words are little-endian on PCI; the generator emits
		// them as %08x of the host-endian uint32 read via ReadU32, so we
		// reconstruct with the same endianness WriteU32 uses.
		off := i * 4
		raw[off] = byte(w)
		raw[off+1] = byte(w >> 8)
		raw[off+2] = byte(w >> 16)
		raw[off+3] = byte(w >> 24)
	}
	ctx, warnings, err := ImportConfigSpace(raw)
	if err != nil {
		return nil, nil, err
	}
	// COE captures are always 4KB; mark the context size explicitly so the
	// extended-capability parser in buildContextFromConfigSpace ran.
	return ctx, warnings, nil
}

// buildContextFromConfigSpace populates a DeviceContext from a parsed
// ConfigSpace, mirroring the field population in collector.Collect without
// touching sysfs, BAR memory, or live traces. Missing BAR contents and
// MSI-X table data are reported as warnings, not errors.
func buildContextFromConfigSpace(cs *pci.ConfigSpace) (*DeviceContext, []Warning) {
	dev := pci.PCIDevice{
		VendorID:       cs.VendorID(),
		DeviceID:       cs.DeviceID(),
		SubsysVendorID: cs.SubsysVendorID(),
		SubsysDeviceID: cs.SubsysDeviceID(),
		RevisionID:     cs.RevisionID(),
		ClassCode:      cs.ClassCode(),
		HeaderType:     cs.HeaderType(),
	}

	ctx := &DeviceContext{
		CollectedAt:  time.Now().UTC(),
		ToolVersion:  version.Version,
		Device:       dev,
		ConfigSpace:  cs,
		BARs:         pci.ParseBARsFromConfigSpace(cs),
		Capabilities: pci.ParseCapabilities(cs),
	}

	if cs.Size >= pci.ConfigSpaceSize {
		ctx.ExtCapabilities = pci.ParseExtCapabilities(cs)
	}

	var warnings []Warning

	// BAR contents cannot be recovered from config space alone.
	hasImplementedBAR := false
	for _, b := range ctx.BARs {
		if b.Type != pci.BARTypeDisabled {
			hasImplementedBAR = true
			break
		}
	}
	if hasImplementedBAR {
		warnings = append(warnings, Warning{
			Code:    "bar_contents_missing",
			Message: "config-space capture has no BAR memory; BAR contents will synthesize defaults",
		})
	}

	// MSI-X capability presence implies a table we cannot populate without
	// BAR memory. Surface a warning so `validate` can flag the gap.
	for _, c := range ctx.Capabilities {
		if c.ID == pci.CapIDMSIX {
			warnings = append(warnings, Warning{
				Code:    "msix_table_missing",
				Message: "MSI-X capability present but table entries not captured (no BAR memory)",
			})
			break
		}
	}

	if len(ctx.Capabilities) == 0 {
		warnings = append(warnings, Warning{
			Code:    "capabilities_missing",
			Message: "no standard capabilities parsed from config space",
		})
	}

	if ctx.ExtCapabilities == nil {
		warnings = append(warnings, Warning{
			Code:    "ext_capabilities_missing",
			Message: fmt.Sprintf("only legacy config space (%d bytes) captured; extended caps not populated", cs.Size),
		})
	}

	return ctx, warnings
}

// parseCOEFile reads a COE file and extracts its memory_initialization_vector
// DWORDs. Supports the generator's own format (radix=16). The parser is
// intentionally strict: unknown radices or a missing vector header are
// rejected with a typed error rather than best-effort guessing.
//
// ponytail: ceiling = multi-format COE reader; upgrade path = split into a
// dedicated codegen.ParseCOE helper reused by validate when a second emitter
// appears.
func parseCOEFile(path string) ([]uint32, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("import: failed to open COE: %w", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Buffer(make([]byte, 64*1024), 1024*1024)

	radix := 0
	inVector := false
	var words []uint32

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, ";") {
			continue
		}

		switch {
		case strings.HasPrefix(line, "memory_initialization_radix="):
			r := strings.TrimPrefix(line, "memory_initialization_radix=")
			r = strings.TrimSuffix(r, ";")
			radix, err = strconv.Atoi(strings.TrimSpace(r))
			if err != nil || radix != 16 {
				return nil, fmt.Errorf("import: COE radix %q unsupported (only 16)", line)
			}
		case strings.HasPrefix(line, "memory_initialization_vector="):
			inVector = true
			rest := strings.TrimPrefix(line, "memory_initialization_vector=")
			rest = strings.TrimSpace(rest)
			if rest != "" {
				if w, ok := parseCOEWord(rest, radix); ok {
					words = append(words, w)
				}
			}
		default:
			if !inVector {
				continue
			}
			if w, ok := parseCOEWord(line, radix); ok {
				words = append(words, w)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("import: COE read error: %w", err)
	}
	if radix == 0 {
		return nil, fmt.Errorf("import: COE missing memory_initialization_radix")
	}
	if !inVector {
		return nil, fmt.Errorf("import: COE missing memory_initialization_vector")
	}
	if len(words) == 0 {
		return nil, fmt.Errorf("import: COE vector is empty")
	}
	return words, nil
}

// parseCOEWord parses a single vector entry, tolerating a trailing comma or
// the terminating semicolon on the final entry.
func parseCOEWord(s string, radix int) (uint32, bool) {
	s = strings.TrimSpace(s)
	s = strings.TrimSuffix(s, ",")
	s = strings.TrimSuffix(s, ";")
	s = strings.TrimSpace(s)
	if s == "" {
		return 0, false
	}
	v, err := strconv.ParseUint(s, radix, 32)
	if err != nil {
		return 0, false
	}
	return uint32(v), true
}