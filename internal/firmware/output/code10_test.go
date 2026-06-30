package output

import (
	"strings"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/firmware"
)

const (
	ccNVMe = 0x010802
	ccXHCI = 0x0C0330
	ccSATA = 0x010601
	ccWiFi = 0x028000
	ccGPU  = 0x030000
	ccEth  = 0x020000
)

func risksFor(t *testing.T, classCode uint32, stockBar bool) []string {
	t.Helper()
	ctx := makeDonorContext(0x1234, 0x5678, classCode)
	b, err := board.Find("PCIeSquirrel")
	if err != nil {
		t.Fatalf("board.Find: %v", err)
	}
	ids := firmware.ExtractDeviceIDs(ctx.ConfigSpace, ctx.ExtCapabilities)
	ow := NewOutputWriter(t.TempDir(), "", 0, 0)
	ow.StockBar = stockBar
	return ow.code10Risks(ctx, b, ids)
}

func joinLower(rs []string) string { return strings.ToLower(strings.Join(rs, "\n")) }

func mustContain(t *testing.T, rs []string, sub string) {
	t.Helper()
	if !strings.Contains(joinLower(rs), strings.ToLower(sub)) {
		t.Errorf("expected a risk mentioning %q, got:\n%s", sub, strings.Join(rs, "\n"))
	}
}

func mustNotContain(t *testing.T, rs []string, sub string) {
	t.Helper()
	if strings.Contains(joinLower(rs), strings.ToLower(sub)) {
		t.Errorf("did NOT expect %q in risks, got:\n%s", sub, strings.Join(rs, "\n"))
	}
}

func TestCode10_NVMe_BackingStoreNotStaleText(t *testing.T) {
	rs := risksFor(t, ccNVMe, false)
	mustContain(t, rs, "backing store")
	mustNotContain(t, rs, "io queues + bram cache are not implemented")
}

func TestCode10_XHCI_NoTransferEngine(t *testing.T) {
	rs := risksFor(t, ccXHCI, false)
	mustContain(t, rs, "transfer engine")
}

func TestCode10_SATA_NoFSM(t *testing.T) {
	rs := risksFor(t, ccSATA, false)
	mustContain(t, rs, "no command-list")
}

func TestCode10_WiFi_NoFSM(t *testing.T) {
	rs := risksFor(t, ccWiFi, false)
	mustContain(t, rs, "wifi")
}

func TestCode10_GPU_NoFSM(t *testing.T) {
	rs := risksFor(t, ccGPU, false)
	mustContain(t, rs, "gpu")
}

func TestCode10_InterruptRoutingVerificationNote(t *testing.T) {
	rs := risksFor(t, ccNVMe, false)
	mustContain(t, rs, "intr_req")

	stock := risksFor(t, ccNVMe, true)
	if strings.Contains(joinLower(stock), "intr_req") {
		t.Errorf("stock-bar mode should not emit the intr_req routing note, got:\n%s", strings.Join(stock, "\n"))
	}
}

func TestCode10_Ethernet_NoPanic(t *testing.T) {
	_ = risksFor(t, ccEth, false)
}

func TestCode10_MSIX_BIRRelocationWarning(t *testing.T) {
	ctx := makeDonorContext(0x1234, 0x5678, ccEth)
	ctx.MSIXData = &donor.MSIXData{TableSize: 4, TableBIR: 2, PBABIR: 2}
	b, _ := board.Find("PCIeSquirrel")
	ids := firmware.ExtractDeviceIDs(ctx.ConfigSpace, ctx.ExtCapabilities)
	ow := NewOutputWriter(t.TempDir(), "", 0, 0)
	rs := ow.code10Risks(ctx, b, ids)
	mustContain(t, rs, "relocated to bar0")

	ctx.MSIXData = &donor.MSIXData{TableSize: 4, TableBIR: 0, PBABIR: 0}
	rs0 := ow.code10Risks(ctx, b, ids)
	mustNotContain(t, rs0, "relocated to bar0")
}
