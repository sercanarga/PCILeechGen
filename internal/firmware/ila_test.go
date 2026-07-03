package firmware

import (
	"fmt"
	"strings"
	"testing"
)

func TestILAProbes_Valid(t *testing.T) {
	probes := ILAProbes()
	if len(probes) == 0 {
		t.Fatal("expected at least one probe")
	}
	for _, p := range probes {
		if p.Signal == "" || p.Width <= 0 {
			t.Errorf("invalid probe: %+v", p)
		}
	}
}

func TestILACreateIPTCL(t *testing.T) {
	tcl := ILACreateIPTCL(0)
	n := len(ILAProbes())
	for _, want := range []string{
		"create_ip -name ila -vendor xilinx.com",
		"-module_name ila_0",
		fmt.Sprintf("CONFIG.C_NUM_OF_PROBES {%d}", n),
		"CONFIG.C_DATA_DEPTH {1024}",
		"CONFIG.C_PROBE0_WIDTH {32}",
		"CONFIG.C_ADV_TRIGGER {true}",
		"CONFIG.C_EN_STRG_QUAL {1}",
		"CONFIG.ALL_PROBE_SAME_MU_CNT {2}",
		"generate_target all [get_ips ila_0]",
	} {
		if !strings.Contains(tcl, want) {
			t.Errorf("create_ip TCL missing %q:\n%s", want, tcl)
		}
	}
	if !strings.Contains(ILACreateIPTCL(4096), "CONFIG.C_DATA_DEPTH {4096}") {
		t.Error("custom depth not honored")
	}
}

func TestILAInstanceSV(t *testing.T) {
	sv := ILAInstanceSV()
	probes := ILAProbes()
	if !strings.Contains(sv, "ila_0 u_ila_dbg") || !strings.Contains(sv, ".clk(clk)") {
		t.Errorf("instance missing module/clk:\n%s", sv)
	}
	if !strings.Contains(sv, fmt.Sprintf(".probe0(%s)", probes[0].Signal)) {
		t.Errorf("first probe not wired:\n%s", sv)
	}
	last := fmt.Sprintf(".probe%d(%s)\n    );", len(probes)-1, probes[len(probes)-1].Signal)
	if !strings.Contains(sv, last) {
		t.Errorf("last probe / closing malformed:\n%s", sv)
	}
}

func TestILADebugDoc(t *testing.T) {
	doc := ILADebugDoc()
	for _, want := range []string{"JTAG", "trigger", "pre-trigger", "warm reboot", "probe0"} {
		if !strings.Contains(doc, want) {
			t.Errorf("ILA debug doc missing %q", want)
		}
	}
	for _, p := range ILAProbes() {
		if !strings.Contains(doc, p.Signal) {
			t.Errorf("ILA debug doc should list probe signal %q", p.Signal)
		}
	}
}
