package svgen

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestGenerateEthernetDMAEngineSV(t *testing.T) {
	sv, err := GenerateEthernetDMAEngineSV(&SVGeneratorConfig{})
	if err != nil {
		t.Fatalf("GenerateEthernetDMAEngineSV: %v", err)
	}
	for _, want := range []string{
		"module pcileech_ethernet_dma_engine",
		"rdt_write",
		"dma_rd_req",
		"dma_wr_valid",
		"32'h0000003C",
		"32'h00000003",
		"icmp_word",
		"packet_profile",
	} {
		if !strings.Contains(sv, want) {
			t.Errorf("Ethernet DMA engine missing %q", want)
		}
	}
}

func TestGenerateEthernetDMAEngineSVVerilatorLint(t *testing.T) {
	if _, err := exec.LookPath("verilator"); err != nil {
		t.Skip("verilator unavailable")
	}
	sv, err := GenerateEthernetDMAEngineSV(&SVGeneratorConfig{})
	if err != nil {
		t.Fatalf("GenerateEthernetDMAEngineSV: %v", err)
	}
	dir := t.TempDir()
	path := filepath.Join(dir, "ethernet_dma_engine.sv")
	if err := os.WriteFile(path, []byte(sv), 0o600); err != nil {
		t.Fatalf("write generated SV: %v", err)
	}
	cmd := exec.Command("verilator", "--lint-only", "--language", "1800-2012", path)
	if output, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("verilator lint: %v\n%s", err, output)
	}
}

func TestGenerateEthernetDMABridgeSVUsesEthernetModuleName(t *testing.T) {
	sv, err := GenerateEthernetDMABridgeSV(&SVGeneratorConfig{})
	if err != nil {
		t.Fatalf("GenerateEthernetDMABridgeSV: %v", err)
	}
	if !strings.Contains(sv, "module pcileech_ethernet_dma_bridge") {
		t.Fatal("Ethernet bridge module name missing")
	}
	if strings.Contains(sv, "module pcileech_nvme_dma_bridge") {
		t.Fatal("Ethernet bridge retained NVMe module declaration")
	}
}
