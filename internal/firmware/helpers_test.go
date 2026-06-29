package firmware

import (
	"testing"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/pci"
)

func TestLowestBar_ByteSlice_EmptyMap(t *testing.T) {
	result := LowestBar(map[int][]byte(nil))
	if result != nil {
		t.Error("nil map should return nil")
	}

	result = LowestBar(map[int][]byte{})
	if result != nil {
		t.Error("empty map should return nil")
	}
}

func TestLowestBar_ByteSlice_SingleEntry(t *testing.T) {
	data := []byte{0x01, 0x02, 0x03, 0x04}
	result := LowestBar(map[int][]byte{2: data})
	if result == nil || len(result) != 4 {
		t.Error("should return single entry data")
	}
}

func TestLowestBar_ByteSlice_MultipleEntries(t *testing.T) {
	bar0 := []byte{0xAA}
	bar2 := []byte{0xBB}
	bar4 := []byte{0xCC}

	result := LowestBar(map[int][]byte{4: bar4, 0: bar0, 2: bar2})
	if result == nil || result[0] != 0xAA {
		t.Errorf("should pick BAR0 (lowest index), got %v", result)
	}
}

func TestLowestBar_Profile_EmptyMap(t *testing.T) {
	result := LowestBar(map[int]*donor.BARProfile(nil))
	if result != nil {
		t.Error("nil map should return nil")
	}
}

func TestLowestBar_Profile_SingleEntry(t *testing.T) {
	p := &donor.BARProfile{}
	result := LowestBar(map[int]*donor.BARProfile{1: p})
	if result == nil {
		t.Error("should return the single profile")
	}
}

func TestLowestBar_Profile_PicksLowest(t *testing.T) {
	p0 := &donor.BARProfile{}
	p2 := &donor.BARProfile{}

	result := LowestBar(map[int]*donor.BARProfile{2: p2, 0: p0})
	if result != p0 {
		t.Error("should pick BAR0 (lowest index)")
	}
}

func TestLargestBar_EmptyMap(t *testing.T) {
	if LargestBar(nil) != nil {
		t.Error("nil map should return nil")
	}
	if LargestBar(map[int][]byte{}) != nil {
		t.Error("empty map should return nil")
	}
}

func TestLargestBar_SingleEntry(t *testing.T) {
	data := []byte{0x01, 0x02, 0x03}
	result := LargestBar(map[int][]byte{2: data})
	if len(result) != 3 {
		t.Errorf("should return single entry, got len=%d", len(result))
	}
}

func TestLargestBar_PicksLargest(t *testing.T) {
	bar0 := make([]byte, 256)   // IO BAR (small)
	bar2 := make([]byte, 65536) // MMIO BAR (large)
	bar0[0] = 0xAA
	bar2[0] = 0xBB

	result := LargestBar(map[int][]byte{0: bar0, 2: bar2})
	if len(result) != 65536 || result[0] != 0xBB {
		t.Errorf("should pick BAR2 (largest), got len=%d first=0x%02X", len(result), result[0])
	}
}

func TestLargestBar_EqualSize_PicksNonZero(t *testing.T) {
	bar0 := make([]byte, 4096) // all zeros
	bar2 := make([]byte, 4096)
	bar2[0] = 0x22
	bar2[1] = 0x79
	bar2[4] = 0x01

	result := LargestBar(map[int][]byte{0: bar0, 2: bar2})
	if result[0] != 0x22 || result[1] != 0x79 {
		t.Error("equal size BARs should pick the one with more non-zero bytes")
	}
}

func TestLargestBarIndex_EqualSize_PicksNonZero(t *testing.T) {
	bar0 := make([]byte, 4096) // all zeros
	bar2 := make([]byte, 4096)
	bar2[0] = 0x22
	bar2[1] = 0x79

	idx := LargestBarIndex(map[int][]byte{0: bar0, 2: bar2})
	if idx != 2 {
		t.Errorf("equal size BARs should pick index 2 (non-zero), got %d", idx)
	}
}

func TestCountNonZero(t *testing.T) {
	if countNonZero(nil) != 0 {
		t.Error("nil should return 0")
	}
	if countNonZero([]byte{0, 0, 0}) != 0 {
		t.Error("all zeros should return 0")
	}
	if countNonZero([]byte{1, 0, 2, 0, 3}) != 3 {
		t.Errorf("expected 3 non-zero, got %d", countNonZero([]byte{1, 0, 2, 0, 3}))
	}
}

func TestExtractDeviceIDs_ClassCode(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	cs.WriteU16(0x00, 0x144D) // VendorID
	cs.WriteU16(0x02, 0xA808) // DeviceID
	cs.WriteU8(0x09, 0x02)    // ProgIF
	cs.WriteU8(0x0A, 0x08)    // SubClass
	cs.WriteU8(0x0B, 0x01)    // BaseClass

	ids := ExtractDeviceIDs(cs, nil)

	if ids.VendorID != 0x144D {
		t.Errorf("VendorID: got 0x%04X, want 0x144D", ids.VendorID)
	}
	if ids.DeviceID != 0xA808 {
		t.Errorf("DeviceID: got 0x%04X, want 0xA808", ids.DeviceID)
	}
	if ids.ClassCode != 0x010802 {
		t.Errorf("ClassCode: got 0x%06X, want 0x010802", ids.ClassCode)
	}
}

func TestExtractDeviceIDs_SubsystemIDs(t *testing.T) {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize
	cs.WriteU16(0x00, 0x8086)
	cs.WriteU16(0x02, 0x1533)
	cs.WriteU16(0x2C, 0x1234) // Subsystem Vendor
	cs.WriteU16(0x2E, 0x5678) // Subsystem Device

	ids := ExtractDeviceIDs(cs, nil)
	if ids.SubsysVendorID != 0x1234 {
		t.Errorf("SubsysVendorID: got 0x%04X, want 0x1234", ids.SubsysVendorID)
	}
	if ids.SubsysDeviceID != 0x5678 {
		t.Errorf("SubsysDeviceID: got 0x%04X, want 0x5678", ids.SubsysDeviceID)
	}
}

// New positive tests for large/dynamic BAR (donor/BRAM driven, Size not forced 4k)
func TestCappedBAR0Size_LargeDonorNVMe(t *testing.T) {
	// default board + no-msix + large donor contents -> capped to board bram (4k default)
	ctx := &donor.DeviceContext{BARContents: map[int][]byte{0: make([]byte, 65536)}}
	b := &board.Board{}
	got := CappedBAR0Size(ctx, b, 0)
	if got != 4096 {
		t.Errorf("default+large donor got %d want 4096", got)
	}
	// large board + small msix + donor content -> donor size (within board)
	b = &board.Board{BRAMSize: 32768}
	ctx = &donor.DeviceContext{BARContents: map[int][]byte{0: make([]byte, 16384)}}
	got = CappedBAR0Size(ctx, b, 1)
	if got != 16384 {
		t.Errorf("donor-override got %d want 16384", got)
	}
	// small donor (small contents, no bars, no msix) on large board -> small demand (not expanded to full board)
	ctxSmall := &donor.DeviceContext{BARContents: map[int][]byte{0: make([]byte, 256)}}
	bLarge := &board.Board{BRAMSize: 131072}
	gotSmall := CappedBAR0Size(ctxSmall, bLarge, 0)
	if gotSmall != 256 {
		t.Errorf("small donor no-msix large board got %d want 256 (donor demand, not auto-expanded)", gotSmall)
	}
}

func TestMSIXPlacement_PostDoorbellNVMe_16k(t *testing.T) {
	tableOff, pbaOff, _ := MSIXPlacement(16384, 16, 0x010802, 0)
	if tableOff < 0x1000 {
		t.Errorf("NVMe tableOff %d not post-doorbell", tableOff)
	}
	if tableOff+uint32(16*16+16) > 16384 {
		t.Error("does not fit")
	}
	if pbaOff <= tableOff {
		t.Error("pba after table")
	}
}

func TestMSIXPlacement_LargeBAR_NoClamp(t *testing.T) {
	tableOff, _, _ := MSIXPlacement(32768, 4, 0, 0)
	if tableOff == 0 || tableOff < 0x2000 {
		t.Errorf("large placement %d suspicious", tableOff)
	}
}

// TestLargeBAR_DemandAndCapped (local to firmware pkg to avoid import cycles).
// Exercises DonorBAR0Demand + CappedBAR0Size (the core calls underlying
// cmd/pcileechgen check/build/validate large-BAR >4k donor vs board BRAM gates,
// force, Capped for scrub/writer, and tclgen Bar0 reporting). Full repro with
// ValidateBARSize(donor), tclgen stock/Bar0ByteSize, sim of build/check/validate
// logic lives in validator_test.go (can import donor/pci/tclgen/output).
func TestLargeBAR_DemandAndCapped(t *testing.T) {
	board4k := &board.Board{Name: "SmallBoard", FPGAPart: "xc7a35t", PCIeLanes: 1} // defaults to 4096
	board16k := &board.Board{Name: "LargeBoard", FPGAPart: "xc7a75t", PCIeLanes: 1, BRAMSize: 16384}

	data16k := make([]byte, 16384)
	data16k[0] = 0xAB
	ctx16k := &donor.DeviceContext{
		Device:      pci.PCIDevice{VendorID: 0x1234, DeviceID: 0x5678, ClassCode: 0x010802},
		BARs:        []pci.BAR{{Index: 0, Size: 16384, Type: pci.BARTypeMem32}},
		BARContents: map[int][]byte{0: data16k},
	}

	demand := DonorBAR0Demand(ctx16k, board4k, 0)
	capped := CappedBAR0Size(ctx16k, board4k, 0)
	if demand != 16384 {
		t.Errorf("Donor demand for 16k ctx on 4k board = %d, want 16384 (to trigger gate)", demand)
	}
	if capped != 4096 || capped > board4k.BRAMSizeOrDefault() {
		t.Errorf("Capped must be 4096 (board BRAM) for oversized donor; got %d", capped)
	}

	demandFit := DonorBAR0Demand(ctx16k, board16k, 0)
	cappedFit := CappedBAR0Size(ctx16k, board16k, 0)
	if demandFit != 16384 || cappedFit != 16384 {
		t.Errorf("fitting donor on 16k board: demand/capped want 16384 got %d/%d", demandFit, cappedFit)
	}

	// size consistency: Capped never > board, used for scrub/writer/validate COE etc.
	for _, tc := range []struct{ dmd, br, want int }{{16384, 4096, 4096}, {8192, 16384, 8192}, {4096, 4096, 4096}, {0, 4096, 4096}} {
		c := CappedBAR0Size(&donor.DeviceContext{BARContents: map[int][]byte{0: make([]byte, tc.dmd)}}, &board.Board{BRAMSize: tc.br}, 0)
		if c != tc.want || c > tc.br {
			t.Errorf("Capped(demand=%d,br=%d)=%d inconsistent (want<=%d)", tc.dmd, tc.br, c, tc.br)
		}
	}
	t.Logf("LargeBAR demand/capped OK: 16k donor on 4k board demand=16384 capped=4096; fits on larger; artifacts size always <=BRAM")
}

func TestDonorMSIXPlacement(t *testing.T) {
	const nvme = uint32(0x010802)
	// donor table in BAR0, fits with room for donor PBA -> kept verbatim
	if tbl, pba, ok := DonorMSIXPlacement(0x4000, 16, 0, 0x3000, 0, 0x3800, nvme, 0); !ok || tbl != 0x3000 || pba != 0x3800 {
		t.Errorf("fit case: got tbl=0x%X pba=0x%X ok=%v, want 0x3000/0x3800/true", tbl, pba, ok)
	}
	// donor PBA doesn't fit -> PBA packed right after the table, table still donor's
	if tbl, pba, ok := DonorMSIXPlacement(0x3200, 16, 0, 0x3000, 0, 0x3800, nvme, 0); !ok || tbl != 0x3000 || pba != 0x3100 {
		t.Errorf("packed-PBA case: got tbl=0x%X pba=0x%X ok=%v, want 0x3000/0x3100/true", tbl, pba, ok)
	}
	// table doesn't fit the BAR -> relocate (ok=false)
	if _, _, ok := DonorMSIXPlacement(0x2000, 16, 0, 0x3000, 0, 0x3800, nvme, 0); ok {
		t.Error("table beyond BAR should not be donor-faithful")
	}
	// table in a different BAR (BIR != 0) -> can't match, relocate
	if _, _, ok := DonorMSIXPlacement(0x4000, 16, 2, 0x3000, 2, 0x3800, nvme, 0); ok {
		t.Error("BIR!=0 should not be donor-faithful")
	}
	// donor table overlapping the NVMe doorbell window -> relocate
	if _, _, ok := DonorMSIXPlacement(0x4000, 16, 0, 0x1000, 0, 0x1800, nvme, 0); ok {
		t.Error("table over NVMe doorbell window should not be donor-faithful")
	}
	// non-NVMe device: same doorbell offset is fine (no doorbell window)
	if _, _, ok := DonorMSIXPlacement(0x4000, 16, 0, 0x1000, 0, 0x1800, 0x020000, 0); !ok {
		t.Error("non-NVMe at 0x1000 should be donor-faithful")
	}
}
