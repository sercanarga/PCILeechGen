package firmware

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// PatchResult tracks what was patched in a file.
type PatchResult struct {
	File    string
	Patches []string
}

// SVPatcher patches pcileech-fpga SystemVerilog source files with donor device data.
type SVPatcher struct {
	ids     DeviceIDs
	srcDir  string // path to board's src/ directory
	results []PatchResult
}

// NewSVPatcher creates a patcher for the given device IDs and board source directory.
func NewSVPatcher(ids DeviceIDs, srcDir string) *SVPatcher {
	return &SVPatcher{ids: ids, srcDir: srcDir}
}

// Results returns all patch results.
func (p *SVPatcher) Results() []PatchResult {
	return p.results
}

// PatchAll applies all patches to the board's source directory.
// Files are modified in-place within the output working copy.
func (p *SVPatcher) PatchAll() error {
	if err := p.patchCfgSV(); err != nil {
		return fmt.Errorf("patching pcileech_pcie_cfg_a7.sv: %w", err)
	}

	if err := p.patchFifoSV(); err != nil {
		return fmt.Errorf("patching pcileech_fifo.sv: %w", err)
	}

	return nil
}

// svRegexPatch defines a single regex-based patch operation.
type svRegexPatch struct {
	pattern     string         // regex pattern with capture groups
	replacement string         // replacement string using $1, $2 etc.
	label       string         // human-readable description
	re          *regexp.Regexp // compiled pattern (lazy)
}

// compile returns the compiled regex, compiling on first use.
func (p *svRegexPatch) compile() *regexp.Regexp {
	if p.re == nil {
		p.re = regexp.MustCompile(p.pattern)
	}
	return p.re
}

// applyRegexPatches applies a list of regex patches to content, returning modified content and patch labels.
func applyRegexPatches(content string, patches []svRegexPatch) (string, []string) {
	modified := content
	var applied []string

	for i := range patches {
		re := patches[i].compile()
		if re.MatchString(modified) {
			modified = re.ReplaceAllString(modified, patches[i].replacement)
			applied = append(applied, patches[i].label)
		}
	}

	return modified, applied
}

// patchFile reads a file, applies patches, and writes back only if changed.
func (p *SVPatcher) patchFile(filename string, patches []svRegexPatch) error {
	path := filepath.Join(p.srcDir, filename)

	content, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("reading %s: %w", filename, err)
	}

	original := string(content)
	modified, applied := applyRegexPatches(original, patches)

	if len(applied) > 0 && modified != original {
		if err := os.WriteFile(path, []byte(modified), 0644); err != nil {
			return err
		}
		p.results = append(p.results, PatchResult{File: filename, Patches: applied})
	}

	return nil
}

// patchCfgSV patches pcileech_pcie_cfg_a7.sv:
//   - DSN (Device Serial Number) from donor's extended capabilities
func (p *SVPatcher) patchCfgSV() error {
	var patches []svRegexPatch

	if p.ids.HasDSN {
		dsnHex := DSNToSVHex(p.ids.DSN)
		patches = append(patches, svRegexPatch{
			pattern:     `(rw\[127:64\]\s*<=\s*64'h)[0-9a-fA-F]+(\s*;\s*//.*cfg_dsn)`,
			replacement: fmt.Sprintf("${1}%s${2}", dsnHex),
			label:       fmt.Sprintf("DSN: 0x%s", dsnHex),
		})
	}

	return p.patchFile("pcileech_pcie_cfg_a7.sv", patches)
}

// patchFifoSV patches pcileech_fifo.sv with donor device identity.
func (p *SVPatcher) patchFifoSV() error {
	ids := p.ids

	patches := []svRegexPatch{
		// Shadow config space: CFGTLP ZERO DATA -> 0
		{
			pattern:     `(rw\[203\]\s*<=\s*)1'b1(\s*;\s*//\s*CFGTLP ZERO DATA)`,
			replacement: "${1}1'b0${2}",
			label:       "Shadow config space: ENABLED (CFGTLP ZERO DATA -> 0)",
		},
		// CFG_SUBSYS_VEND_ID
		{
			pattern:     `(rw\[143:128\]\s*<=\s*16'h)[0-9a-fA-F]{4}(\s*;\s*//.*CFG_SUBSYS_VEND_ID)`,
			replacement: fmt.Sprintf("${1}%04X${2}", ids.SubsysVendorID),
			label:       fmt.Sprintf("CFG_SUBSYS_VEND_ID: 0x%04X", ids.SubsysVendorID),
		},
		// CFG_SUBSYS_ID
		{
			pattern:     `(rw\[159:144\]\s*<=\s*16'h)[0-9a-fA-F]{4}(\s*;\s*//.*CFG_SUBSYS_ID)`,
			replacement: fmt.Sprintf("${1}%04X${2}", ids.SubsysDeviceID),
			label:       fmt.Sprintf("CFG_SUBSYS_ID: 0x%04X", ids.SubsysDeviceID),
		},
		// CFG_VEND_ID
		{
			pattern:     `(rw\[175:160\]\s*<=\s*16'h)[0-9a-fA-F]{4}(\s*;\s*//.*CFG_VEND_ID)`,
			replacement: fmt.Sprintf("${1}%04X${2}", ids.VendorID),
			label:       fmt.Sprintf("CFG_VEND_ID: 0x%04X", ids.VendorID),
		},
		// CFG_DEV_ID
		{
			pattern:     `(rw\[191:176\]\s*<=\s*16'h)[0-9a-fA-F]{4}(\s*;\s*//.*CFG_DEV_ID)`,
			replacement: fmt.Sprintf("${1}%04X${2}", ids.DeviceID),
			label:       fmt.Sprintf("CFG_DEV_ID: 0x%04X", ids.DeviceID),
		},
		// CFG_REV_ID
		{
			pattern:     `(rw\[199:192\]\s*<=\s*8'h)[0-9a-fA-F]{2}(\s*;\s*//.*CFG_REV_ID)`,
			replacement: fmt.Sprintf("${1}%02X${2}", ids.RevisionID),
			label:       fmt.Sprintf("CFG_REV_ID: 0x%02X", ids.RevisionID),
		},
		// _pcie_core_config initial register (all IDs packed)
		{
			pattern: `(_pcie_core_config\s*=\s*\{\s*4'hf,\s*1'b1,\s*1'b1,\s*1'b0,\s*1'b0,\s*8'h)[0-9a-fA-F]{2}(,\s*16'h)[0-9a-fA-F]{4}(,\s*16'h)[0-9a-fA-F]{4}(,\s*16'h)[0-9a-fA-F]{4}(,\s*16'h)[0-9a-fA-F]{4}`,
			replacement: fmt.Sprintf("${1}%02X${2}%04X${3}%04X${4}%04X${5}%04X",
				ids.RevisionID, ids.DeviceID, ids.VendorID,
				ids.SubsysDeviceID, ids.SubsysVendorID),
			label: "_pcie_core_config: all IDs updated",
		},
	}

	return p.patchFile("pcileech_fifo.sv", patches)
}

// FormatPatchSummary returns a human-readable summary of all patches applied.
func FormatPatchSummary(results []PatchResult) string {
	if len(results) == 0 {
		return "  (no patches applied)"
	}

	var sb strings.Builder
	for _, r := range results {
		sb.WriteString(fmt.Sprintf("  %s:\n", r.File))
		for _, p := range r.Patches {
			sb.WriteString(fmt.Sprintf("    -> %s\n", p))
		}
	}
	return sb.String()
}
