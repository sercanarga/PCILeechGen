package firmware

import (
	"fmt"
	"strings"
)

type ILAProbe struct {
	Signal string
	Width  int
}

const DefaultILADepth = 1024

func ILAProbes() []ILAProbe {
	return []ILAProbe{
		{"wr_addr", 32},
		{"rd_req_addr", 32},
		{"bar0_raw_data", 32},
		{"bar0_raw_valid", 1},
		{"rd_req_valid", 1},
		{"intr_req", 1},
		{"in_is_wr_ready", 1},
	}
}

func ILACreateIPTCL(depth int) string {
	if depth <= 0 {
		depth = DefaultILADepth
	}
	probes := ILAProbes()
	var b strings.Builder
	b.WriteString("create_ip -name ila -vendor xilinx.com -library ip -module_name ila_0\n")
	b.WriteString("set_property -dict [list \\\n")
	fmt.Fprintf(&b, "  CONFIG.C_NUM_OF_PROBES {%d} \\\n", len(probes))
	fmt.Fprintf(&b, "  CONFIG.C_DATA_DEPTH {%d} \\\n", depth)
	b.WriteString("  CONFIG.C_INPUT_PIPE_STAGES {1} \\\n")
	b.WriteString("  CONFIG.C_ADV_TRIGGER {true} \\\n")
	b.WriteString("  CONFIG.C_EN_STRG_QUAL {1} \\\n")
	b.WriteString("  CONFIG.C_TRIGIN_EN {false} \\\n")
	b.WriteString("  CONFIG.C_TRIGOUT_EN {false} \\\n")
	b.WriteString("  CONFIG.ALL_PROBE_SAME_MU {true} \\\n")
	b.WriteString("  CONFIG.ALL_PROBE_SAME_MU_CNT {2} \\\n")
	for i, p := range probes {
		fmt.Fprintf(&b, "  CONFIG.C_PROBE%d_WIDTH {%d} \\\n", i, p.Width)
	}
	b.WriteString("] [get_ips ila_0]\n")
	b.WriteString("generate_target all [get_ips ila_0]\n")
	return b.String()
}

func ILADebugDoc() string {
	probes := ILAProbes()
	var b strings.Builder
	b.WriteString("ILA debug core - capture workflow\n")
	b.WriteString("=================================\n\n")
	b.WriteString("The ILA is wired into pcileech_tlps128_bar_controller and works over JTAG\n")
	b.WriteString("only (not the data/USB-C port). Connect the board's JTAG to a 2nd PC running\n")
	b.WriteString("Vivado Hardware Manager; the target PC (with the card inserted) is separate.\n\n")
	b.WriteString("Probes:\n")
	for i, p := range probes {
		fmt.Fprintf(&b, "  probe%d  %-15s (%d bit)\n", i, p.Signal, p.Width)
	}
	b.WriteString("\nArming a pre-trigger capture (example: NVMe CQ0 init):\n")
	b.WriteString("  1. open_hw_manager, connect to the board over JTAG, refresh hw_device.\n")
	b.WriteString("  2. In the ILA dashboard set a trigger on a probe - e.g. wr_addr == the CQ0\n")
	b.WriteString("     doorbell offset, or intr_req == 1 for the first completion interrupt.\n")
	b.WriteString("  3. Set the trigger position early in the window (a pre-trigger position) so\n")
	b.WriteString("     the buffer captures what happened BEFORE the trigger as well as after.\n")
	b.WriteString("  4. Enable capture control (storage qualification) to only store samples of\n")
	b.WriteString("     interest - this stretches the small on-card buffer much further.\n")
	b.WriteString("  5. Arm the core, then do a warm reboot of the target PC. When the trigger\n")
	b.WriteString("     arrives during init, the ILA captures the window; export it to a .ila/CSV.\n\n")
	b.WriteString("Buffers are small, so expect several capture runs, each narrowed to one signal\n")
	b.WriteString("or window, to reconstruct the full init sequence around a Code 10.\n")
	return b.String()
}

func ILAInstanceSV() string {
	probes := ILAProbes()
	var b strings.Builder
	b.WriteString("    ila_0 u_ila_dbg (\n")
	b.WriteString("        .clk(clk),\n")
	for i, p := range probes {
		comma := ","
		if i == len(probes)-1 {
			comma = ""
		}
		fmt.Fprintf(&b, "        .probe%d(%s)%s\n", i, p.Signal, comma)
	}
	b.WriteString("    );\n")
	return b.String()
}
