package svgen

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/sercanarga/pcileechgen/internal/firmware"
)

// PatchResult tracks what was patched in a file.
type PatchResult struct {
	File     string
	Patches  []string
	Warnings []string
}

// SVPatcher injects donor IDs into pcileech-fpga SV sources via regex.
type SVPatcher struct {
	ids     firmware.DeviceIDs
	srcDir  string // path to board's src/ directory
	results []PatchResult
}

func NewSVPatcher(ids firmware.DeviceIDs, srcDir string) *SVPatcher {
	return &SVPatcher{ids: ids, srcDir: srcDir}
}

// Results returns all patch results.
func (p *SVPatcher) Results() []PatchResult {
	return p.results
}

// PatchAll patches cfg + fifo SV files in-place.
func (p *SVPatcher) PatchAll() error {
	if err := p.patchCfgSV(); err != nil {
		return fmt.Errorf("patching pcileech_pcie_cfg_a7.sv: %w", err)
	}
	tlpData, err := os.ReadFile(filepath.Join(p.srcDir, "pcileech_pcie_tlp_a7.sv"))
	if err == nil {
		modern, classifyErr := classifyServiceTLP(string(tlpData))
		if classifyErr != nil {
			return fmt.Errorf("patching pcileech_pcie_tlp_a7.sv services: %w", classifyErr)
		}
		if modern {
			requestPort := interruptRequestPort(string(tlpData))
			if requestPort == "" {
				requestPort = "generated_bar_intr_req"
			}
			if perr := p.patchCfgInterruptHandshake(requestPort); perr != nil {
				return fmt.Errorf("patching pcileech_pcie_cfg_a7.sv interrupt handshake: %w", perr)
			}
		}
	} else if !os.IsNotExist(err) {
		return err
	}
	if err := p.patchTLPServiceWiring(); err != nil {
		return fmt.Errorf("patching pcileech_pcie_tlp_a7.sv services: %w", err)
	}
	if err := p.patchPCIeWrapperServiceWiring(); err != nil {
		return fmt.Errorf("patching PCIe wrapper services: %w", err)
	}
	if err := p.validateServiceWiring(); err != nil {
		return err
	}

	if err := p.patchFifoSV(); err != nil {
		return fmt.Errorf("patching pcileech_fifo.sv: %w", err)
	}

	// Validate that critical patches were applied
	if err := p.validatePatchResults(); err != nil {
		return err
	}

	return nil
}

// validatePatchResults returns an error if fewer patches landed than expected.
// Catches silent breakage when upstream SV format changes.
func (p *SVPatcher) validatePatchResults() error {
	fifoPatched := 0
	cfgPatched := 0
	for _, r := range p.results {
		switch r.File {
		case "pcileech_fifo.sv":
			fifoPatched = len(r.Patches)
		case "pcileech_pcie_cfg_a7.sv":
			cfgPatched = len(r.Patches)
		}
	}

	// pcileech_fifo.sv should have at least VendorID + DeviceID patches
	const minFifoPatches = 2
	if fifoPatched < minFifoPatches {
		return fmt.Errorf("pcileech_fifo.sv: only %d/%d minimum patches applied - "+
			"upstream SV format may have changed (VendorID/DeviceID patches are critical)", fifoPatched, minFifoPatches)
	}

	// cfg SV: if device has DSN, at least 1 patch expected
	if p.ids.HasDSN && cfgPatched == 0 {
		w := "pcileech_pcie_cfg_a7.sv: DSN patch expected but not applied - " +
			"upstream SV format may have changed"
		p.results = append(p.results, PatchResult{
			File:     "pcileech_pcie_cfg_a7.sv",
			Warnings: []string{w},
		})
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

// patchCfgSV patches pcileech_pcie_cfg_a7.sv with DSN and power management
// settings. In addition to DSN, it forces the IP core to reject ASPM L0s/L1
// transitions so the link stays in L0 even if the root complex requests it.
func (p *SVPatcher) patchCfgSV() error {
	var patches []svRegexPatch

	if p.ids.HasDSN {
		dsnHex := firmware.DSNToSVHex(p.ids.DSN)
		patches = append(patches, svRegexPatch{
			pattern:     `(rw\[127:64\]\s*<=\s*64'h)[0-9a-fA-F]+(\s*;\s*//.*cfg_dsn)`,
			replacement: fmt.Sprintf("${1}%s${2}", dsnHex),
			label:       fmt.Sprintf("DSN: 0x%s", dsnHex),
		})
	} else {
		// No DSN on the donor - clear the default value so the FPGA
		// won't show a serial number that doesn't exist on the real device.
		patches = append(patches, svRegexPatch{
			pattern:     `(rw\[127:64\]\s*<=\s*64'h)[0-9a-fA-F]+(\s*;\s*//.*cfg_dsn)`,
			replacement: "${1}0000000000000000${2}",
			label:       "DSN: cleared (donor has no DSN)",
		})
	}

	// halt ASPM L0s at IP core level - prevents root complex from
	// transitioning the link to L0s even if it requests it.
	patches = append(patches, svRegexPatch{
		pattern:     `(rw\[211\]\s*<=\s*)0(\s*;\s*//\s*cfg_pm_halt_aspm_l0s)`,
		replacement: "${1}1${2}",
		label:       "PM: halt ASPM L0s (cfg_pm_halt_aspm_l0s -> 1)",
	})

	// halt ASPM L1 at IP core level - prevents root complex from
	// transitioning the link to L1, which would stop TLP processing
	// and cause the device to appear dead after ~5 minutes of idle.
	patches = append(patches, svRegexPatch{
		pattern:     `(rw\[212\]\s*<=\s*)0(\s*;\s*//\s*cfg_pm_halt_aspm_l1)`,
		replacement: "${1}1${2}",
		label:       "PM: halt ASPM L1 (cfg_pm_halt_aspm_l1 -> 1)",
	})

	patches = append(patches, svRegexPatch{
		pattern:     `(rw\[210\]\s*<=\s*)0(\s*;\s*//\s*cfg_pm_force_state_en)`,
		replacement: "${1}1${2}",
		label:       "PM: force D0 state (cfg_pm_force_state_en -> 1)",
	})

	patches = append(patches, svRegexPatch{
		pattern:     `(rw\[21\]\s*<=\s*)0(\s*;\s*//\s*CFGSPACE_COMMAND_REGISTER_AUTO_SET)`,
		replacement: "${1}1${2}",
		label:       "CMD: auto-set BME+MSE+IOSE (CFGSPACE_COMMAND_REGISTER_AUTO_SET -> 1)",
	})

	// enable periodic status register error bit clearing. every ~1ms
	// this clears W1C error bits (master abort, target abort, parity)
	// in the status register. prevents error accumulation that would
	// cause Windows to disable the device.
	patches = append(patches, svRegexPatch{
		pattern:     `(rw\[20\]\s*<=\s*)0(\s*;\s*//\s*CFGSPACE_STATUS_REGISTER_AUTO_CLEAR)`,
		replacement: "${1}1${2}",
		label:       "STATUS: auto-clear error bits (CFGSPACE_STATUS_REGISTER_AUTO_CLEAR -> 1)",
	})

	return p.patchFile("pcileech_pcie_cfg_a7.sv", patches)
}

func (p *SVPatcher) patchCfgInterruptHandshake(requestPort string) error {
	path := filepath.Join(p.srcDir, "pcileech_pcie_cfg_a7.sv")
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	content := string(data)
	if requestPort == "" {
		return fmt.Errorf("interrupt request port is empty")
	}
	var patches []svRegexPatch
	if !hasInterruptRequestPort(content, requestPort) {
		patches = append(patches, svRegexPatch{
			pattern:     `(module\s+pcileech_pcie_cfg_a7\s*\(\s*\r?\n)`,
			replacement: "${1}    input                   " + requestPort + ",\n",
			label:       "services: accept interrupt request",
		})
	}
	if !strings.Contains(content, "intr_req_pending") {
		patches = append(patches, svRegexPatch{
			pattern: `(assign\s+ctx\.cfg_interrupt\s*=\s*)rw\[206\](\s*;)`,
			replacement: "reg intr_req_pending = 1'b0;\n" +
				"    always @(posedge clk_pcie) begin\n" +
				"        if (rst)\n" +
				"            intr_req_pending <= 1'b0;\n" +
				"        else if (" + requestPort + ")\n" +
				"            intr_req_pending <= 1'b1;\n" +
				"        else if (intr_req_pending && ctx.cfg_interrupt_rdy)\n" +
				"            intr_req_pending <= 1'b0;\n" +
				"    end\n\n" +
				"    ${1}rw[206] | intr_req_pending${2}",
			label: "services: hold interrupt request until ready",
		})
	}
	escapedPort := regexp.QuoteMeta(requestPort)
	patches = append(patches,
		svRegexPatch{
			pattern:     `(assign\s+ctx\.cfg_interrupt_assert\s*=\s*)rw\[205\](\s*;)`,
			replacement: "${1}rw[205] | intr_req_pending${2}",
			label:       "services: route held interrupt assertion",
		},
		svRegexPatch{
			pattern: `(else\s+if\s*\(\s*intr_req_pending\s*&&\s*ctx\.cfg_interrupt_rdy\s*\)\s*begin\s*\r?\n\s*intr_req_pending\s*<=\s*1'b0;\s*\r?\n\s*end\s+else\s+if\s*\(\s*` + escapedPort + `\s*\)\s*begin\s*\r?\n\s*intr_req_pending\s*<=\s*1'b1;\s*\r?\n\s*end)`,
			replacement: "else if (" + requestPort + ") begin\n" +
				"            intr_req_pending <= 1'b1;\n" +
				"        end else if (intr_req_pending && ctx.cfg_interrupt_rdy) begin\n" +
				"            intr_req_pending <= 1'b0;\n" +
				"        end",
			label: "services: preserve interrupt arriving with acknowledgement",
		},
	)
	return p.patchFile("pcileech_pcie_cfg_a7.sv", patches)
}

func classifyServiceTLP(content string) (bool, error) {
	if strings.Contains(content, "pcileech_tlps128_bar_controller") {
		return true, nil
	}
	if strings.Contains(content, "IfTlp64") || strings.Contains(content, "IfPCIeTlpRxTx") {
		return false, nil
	}
	return false, fmt.Errorf("unsupported TLP service architecture")
}

func interruptRequestPort(content string) string {
	for _, name := range []string{"generated_bar_intr_req", "intr_req"} {
		if hasInterruptRequestPort(content, name) {
			return name
		}
	}
	return ""
}

func hasInterruptRequestPort(content, name string) bool {
	re := regexp.MustCompile(`\b(?:input|output)\s+(?:wire\s+)?` + regexp.QuoteMeta(name) + `\b`)
	return re.MatchString(content)
}

func instanceHasPort(content, moduleName, portName string) bool {
	re := regexp.MustCompile(`(?s)\b` + regexp.QuoteMeta(moduleName) + `\s+\w+\s*\((.*?)\);`)
	match := re.FindStringSubmatch(content)
	return len(match) == 2 && strings.Contains(match[1], "."+portName)
}

func instancePortPatch(moduleName, portName, connection, label string) svRegexPatch {
	return svRegexPatch{
		pattern:     `(` + regexp.QuoteMeta(moduleName) + `\s+\w+\s*\(\s*\r?\n)`,
		replacement: "${1}        ." + portName + connection + ",\n",
		label:       label,
	}
}

func (p *SVPatcher) patchTLPServiceWiring() error {
	path := filepath.Join(p.srcDir, "pcileech_pcie_tlp_a7.sv")
	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return err
	}
	content := string(data)
	modern, err := classifyServiceTLP(content)
	if err != nil {
		return err
	}
	if !modern {
		return nil
	}

	var patches []svRegexPatch
	if !regexp.MustCompile(`\bIfPCIeSignals(?:\.\w+)?\s+ctx\b`).MatchString(content) {
		patches = append(patches, svRegexPatch{
			pattern:     `(module\s+pcileech_pcie_tlp_a7\s*\(\s*\r?\n)`,
			replacement: "${1}    IfPCIeSignals           ctx,\n",
			label:       "services: add PCIe context port",
		})
	}
	requestPort := interruptRequestPort(content)
	if requestPort == "" {
		requestPort = "generated_bar_intr_req"
		patches = append(patches, svRegexPatch{
			pattern:     `(module\s+pcileech_pcie_tlp_a7\s*\(\s*\r?\n)`,
			replacement: "${1}    output wire             " + requestPort + ",\n",
			label:       "services: add interrupt request port",
		})
	}
	type portConn struct{ port, conn string }
	desired := []portConn{
		{"cfg_command", "ctx.cfg_command"},
		{"cfg_power_state", "ctx.cfg_pmcsr_powerstate"},
		{"cfg_flr_in_process", "ctx.cfg_received_func_lvl_rst"},
		{"cfg_to_turnoff", "ctx.cfg_to_turnoff"},
		{"cfg_link_up", "ctx.pl_phy_lnk_up"},
		{"cfg_msi_enable", "ctx.cfg_interrupt_msienable"},
		{"cfg_msix_enable", "ctx.cfg_interrupt_msixenable"},
		{"cfg_msix_function_mask", "ctx.cfg_interrupt_msixfm"},
	}
	var ports string
	for _, pc := range desired {
		if !instanceHasPort(content, "pcileech_tlps128_bar_controller", pc.port) {
			ports += fmt.Sprintf("        .%-22s ( %-28s ),\n", pc.port, pc.conn)
		}
	}
	if !instanceHasPort(content, "pcileech_tlps128_bar_controller", "intr_req") {
		ports += fmt.Sprintf("        .intr_req        ( %-28s ),\n", requestPort)
	}
	if ports != "" {
		patches = append(patches, svRegexPatch{
			pattern:     `(pcileech_tlps128_bar_controller\s+\w+\s*\(\s*\r?\n)`,
			replacement: "${1}" + ports,
			label:       "services: wire lifecycle and interrupt BAR ports",
		})
	}
	return p.patchFile("pcileech_pcie_tlp_a7.sv", patches)
}

func (p *SVPatcher) patchPCIeWrapperServiceWiring() error {
	tlpData, err := os.ReadFile(filepath.Join(p.srcDir, "pcileech_pcie_tlp_a7.sv"))
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return err
	}
	modern, err := classifyServiceTLP(string(tlpData))
	if err != nil {
		return err
	}
	if !modern {
		return nil
	}
	requestPort := interruptRequestPort(string(tlpData))
	if requestPort == "" {
		return fmt.Errorf("modern TLP wrapper has no interrupt request port")
	}

	for _, filename := range []string{"pcileech_pcie_a7.sv", "pcileech_pcie_a7x4.sv"} {
		path := filepath.Join(p.srcDir, filename)
		data, err := os.ReadFile(path)
		if os.IsNotExist(err) {
			continue
		}
		if err != nil {
			return err
		}
		content := string(data)
		var patches []svRegexPatch
		wirePattern := regexp.MustCompile(`\bwire\s+` + regexp.QuoteMeta(requestPort) + `\b`)
		if !wirePattern.MatchString(content) {
			patches = append(patches, svRegexPatch{
				pattern:     `(IfPCIeSignals\s+ctx\(\);\s*\r?\n)`,
				replacement: "${1}    wire                    " + requestPort + ";\n",
				label:       "services: declare interrupt wire",
			})
		}
		if !instanceHasPort(content, "pcileech_pcie_cfg_a7", requestPort) {
			patches = append(patches, instancePortPatch(
				"pcileech_pcie_cfg_a7", requestPort, "   ( "+requestPort+"    )",
				"services: connect interrupt to cfg wrapper",
			))
		}
		if !instanceHasPort(content, "pcileech_pcie_tlp_a7", "ctx") {
			patches = append(patches, instancePortPatch(
				"pcileech_pcie_tlp_a7", "ctx", "                        ( ctx                       )",
				"services: connect lifecycle context to TLP wrapper",
			))
		}
		if !instanceHasPort(content, "pcileech_pcie_tlp_a7", requestPort) {
			patches = append(patches, instancePortPatch(
				"pcileech_pcie_tlp_a7", requestPort, "   ( "+requestPort+"    )",
				"services: connect interrupt from TLP wrapper",
			))
		}
		if err := p.patchFile(filename, patches); err != nil {
			return err
		}
	}
	return nil
}

func (p *SVPatcher) validateServiceWiring() error {
	tlpPath := filepath.Join(p.srcDir, "pcileech_pcie_tlp_a7.sv")
	tlpData, err := os.ReadFile(tlpPath)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return err
	}
	tlpContent := string(tlpData)
	modern, err := classifyServiceTLP(tlpContent)
	if err != nil {
		return fmt.Errorf("pcileech_pcie_tlp_a7.sv: %w", err)
	}
	requestPort := interruptRequestPort(tlpContent)

	if modern {
		cfgData, err := os.ReadFile(filepath.Join(p.srcDir, "pcileech_pcie_cfg_a7.sv"))
		if err != nil {
			return err
		}
		cfgContent := string(cfgData)
		cfgChecks := []*regexp.Regexp{
			regexp.MustCompile(`\binput\s+(?:wire\s+)?` + regexp.QuoteMeta(requestPort) + `\b`),
			regexp.MustCompile(`\bintr_req_pending\b`),
			regexp.MustCompile(`intr_req_pending\s*&&\s*ctx\.cfg_interrupt_rdy`),
			regexp.MustCompile(`ctx\.cfg_interrupt\s*=\s*rw\[206\]\s*\|\s*intr_req_pending`),
		}
		for _, check := range cfgChecks {
			if !check.MatchString(cfgContent) {
				return fmt.Errorf("pcileech_pcie_cfg_a7.sv: required interrupt handshake %q not applied", check.String())
			}
		}
		required := []string{
			".cfg_command",
			".cfg_power_state",
			".cfg_flr_in_process",
			".cfg_to_turnoff",
			".cfg_link_up",
			".cfg_msi_enable",
			".cfg_msix_enable",
			".cfg_msix_function_mask",
			".intr_req",
		}
		if requestPort == "" {
			return fmt.Errorf("pcileech_pcie_tlp_a7.sv: interrupt request port not applied")
		}
		if !regexp.MustCompile(`\bIfPCIeSignals(?:\.\w+)?\s+ctx\b`).MatchString(tlpContent) {
			return fmt.Errorf("pcileech_pcie_tlp_a7.sv: required service context not applied")
		}
		for _, token := range required {
			if !strings.Contains(tlpContent, token) {
				return fmt.Errorf("pcileech_pcie_tlp_a7.sv: required service wiring %q not applied", token)
			}
		}
	}

	foundWrapper := false
	for _, filename := range []string{"pcileech_pcie_a7.sv", "pcileech_pcie_a7x4.sv"} {
		data, err := os.ReadFile(filepath.Join(p.srcDir, filename))
		if os.IsNotExist(err) {
			continue
		}
		if err != nil {
			return err
		}
		foundWrapper = true
		if modern {
			content := string(data)
			if !instanceHasPort(content, "pcileech_pcie_cfg_a7", requestPort) ||
				!instanceHasPort(content, "pcileech_pcie_tlp_a7", requestPort) {
				return fmt.Errorf("%s: interrupt request path not connected", filename)
			}
			if !instanceHasPort(content, "pcileech_pcie_tlp_a7", "ctx") {
				return fmt.Errorf("%s: lifecycle context not connected", filename)
			}
		}
	}
	if !foundWrapper {
		return fmt.Errorf("no supported Xilinx PCIe wrapper found")
	}
	return nil
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
		// Shadow config space: CFGTLP PCIE WRITE ENABLE -> 1
		// without this, host config writes (BAR sizing, command reg, PM)
		// never reach the shadow BRAM. The writemask DROM filters which
		// bits are actually stored.
		{
			pattern:     `(rw\[206\]\s*<=\s*)1'b0(\s*;\s*//\s*CFGTLP PCIE WRITE ENABLE)`,
			replacement: "${1}1'b1${2}",
			label:       "Shadow config space: PCIe WRITE ENABLED (cfgtlp_wren -> 1)",
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
		if len(r.Patches) > 0 {
			sb.WriteString(fmt.Sprintf("  %s:\n", r.File))
			for _, p := range r.Patches {
				sb.WriteString(fmt.Sprintf("    -> %s\n", p))
			}
		}
		for _, w := range r.Warnings {
			sb.WriteString(fmt.Sprintf("  ⚠ %s\n", w))
		}
	}
	return sb.String()
}
