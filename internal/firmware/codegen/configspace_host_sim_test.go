package codegen

import (
	"strconv"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/pci"
)

// buildNVMeHostSimConfigSpace builds a Samsung 990 Pro-like scrubbed config space
// mirroring testdata/donors/nvme.json (VID=0x144D, class=0x010802) but with a full
// PM/MSI-X/PCIe capability chain, non-zero subsystem IDs and a 16KB Mem32 BAR0 so
// every Windows enumeration check is meaningful.
func buildNVMeHostSimConfigSpace() *pci.ConfigSpace {
	cs := pci.NewConfigSpace()
	cs.Size = pci.ConfigSpaceSize

	// Identity DWORDs (pci.sys reads these first).
	cs.WriteU16(0x00, 0x144D) // VID Samsung
	cs.WriteU16(0x02, 0xA80A) // DID 990 Pro
	cs.WriteU16(0x04, 0x0000) // Command
	cs.WriteU16(0x06, 0x0010) // Status: cap list present
	cs.WriteU8(0x08, 0x00)    // Revision
	cs.WriteU8(0x09, 0x02)    // Prog IF (NVM Express)
	cs.WriteU8(0x0A, 0x08)    // Subclass NVM
	cs.WriteU8(0x0B, 0x01)    // Base class mass storage -> 0x010802
	cs.WriteU8(0x0E, 0x00)    // Header type 0, single function
	cs.WriteU16(0x2C, 0x144D) // Subsys VID (non-zero)
	cs.WriteU16(0x2E, 0xA80A) // Subsys DID (non-zero)
	cs.WriteU8(0x34, 0x40)    // CapPtr -> PM capability
	cs.WriteU8(0x3D, 0x01)    // IntPin INTA

	// BAR0: 16KB Mem32, scrubbed size-encoding (address bits all 1, size bits 0).
	cs.WriteU32(0x10, 0xFFFFC000)

	// PM capability @0x40 (8 bytes) -> next MSI-X.
	cs.WriteU8(0x40, byte(pci.CapIDPowerManagement))
	cs.WriteU8(0x41, 0x48)
	cs.WriteU16(0x42, 0x0003) // PMC
	cs.WriteU32(0x44, 0x00000000)

	// MSI-X capability @0x48 (12 bytes), 8 vectors -> next PCIe.
	cs.WriteU8(0x48, byte(pci.CapIDMSIX))
	cs.WriteU8(0x49, 0x58)
	cs.WriteU16(0x4A, 0x0007)     // Message Control: TableSize=7 (8 vectors)
	cs.WriteU32(0x4C, 0x00002000) // Table Offset/BIR: offset 0x2000 in BAR0
	cs.WriteU32(0x50, 0x00003000) // PBA Offset/BIR: offset 0x3000 in BAR0

	// PCIe capability @0x58, last in chain (next=0).
	cs.WriteU8(0x58, byte(pci.CapIDPCIExpress))
	cs.WriteU8(0x59, 0x00)
	cs.WriteU16(0x5A, 0x0002) // PCIe Cap Register: version 2
	cs.WriteU32(0x5C, 0x00000000)
	cs.WriteU16(0x60, 0x0000) // Device Control
	cs.WriteU16(0x62, 0x0000) // Device Status

	return cs
}

// writemaskDWORDs renders the writemask COE and returns the 1024 DWORD masks.
func writemaskDWORDs(t *testing.T, cs *pci.ConfigSpace) []uint32 {
	t.Helper()
	strs := parseCOEDwords(t, GenerateWritemaskCOE(cs))
	dw := make([]uint32, len(strs))
	for i, s := range strs {
		v, err := strconv.ParseUint(s, 16, 32)
		if err != nil {
			t.Fatalf("parse writemask DWORD[%d] %q: %v", i, s, err)
		}
		dw[i] = uint32(v)
	}
	return dw
}

// TestConfigSpace_HostSim_Identity mirrors pci.sys reading the identity DWORDs.
func TestConfigSpace_HostSim_Identity(t *testing.T) {
	cs := buildNVMeHostSimConfigSpace()

	if vid, did := cs.ReadU16(0x00), cs.ReadU16(0x02); vid != 0x144D || did != 0xA80A {
		t.Fatalf("VID:DID = %04X:%04X, want 144D:A80A", vid, did)
	}
	// DWORD 0 little-endian = DID:VID.
	if got := cs.ReadU32(0x00); got != 0xA80A144D {
		t.Errorf("DWORD0 = %08X, want A80A144D", got)
	}
	// Revision:ClassCode at 0x08 = class(010802) << 8 | revision(00) = 0x01080200.
	if got := cs.ReadU32(0x08); got != 0x01080200 {
		t.Errorf("Rev:Class DWORD = %08X, want 01080200", got)
	}
	// Subsystem IDs must be non-zero or stornvme/nvme.sys will reject the function.
	if cs.ReadU32(0x2C) == 0 {
		t.Error("subsys VID/DID at 0x2C is zero; host storage driver refuses zero subsystem")
	}
}

// TestConfigSpace_HostSim_CapabilityChain walks the cap list like the PCI bus driver.
func TestConfigSpace_HostSim_CapabilityChain(t *testing.T) {
	cs := buildNVMeHostSimConfigSpace()
	if !cs.HasCapabilities() {
		t.Fatal("Status cap-list bit clear; pci.sys treats device as cap-less")
	}

	visited := make(map[int]bool)
	ptr := int(cs.CapabilityPointer())
	if ptr == 0 || ptr >= pci.ConfigSpaceLegacySize {
		t.Fatalf("CapPtr = 0x%02X out of range", ptr)
	}
	hasPCIe := false
	for ptr != 0 {
		if ptr&0x3 != 0 {
			t.Fatalf("cap ptr 0x%02X not DWORD-aligned", ptr)
		}
		if visited[ptr] {
			t.Fatalf("capability cycle detected at 0x%02X", ptr)
		}
		visited[ptr] = true
		if ptr >= pci.ConfigSpaceLegacySize {
			t.Fatalf("cap ptr 0x%02X beyond legacy config space", ptr)
		}
		capID := cs.ReadU8(ptr)
		next := int(cs.ReadU8(ptr+1)) & 0xFC
		if capID == byte(pci.CapIDPCIExpress) {
			hasPCIe = true
		}
		if next != 0 && next <= ptr {
			t.Fatalf("cap 0x%02X (id=0x%02X) next 0x%02X not forward", ptr, capID, next)
		}
		ptr = next
	}
	if !hasPCIe {
		t.Error("no PCIe capability (0x10) in chain; pciexpress.sys cannot bind")
	}

	// Cross-check against the package parser.
	caps := pci.ParseCapabilities(cs)
	if len(caps) == 0 {
		t.Fatal("ParseCapabilities returned empty")
	}
	ids := map[uint8]bool{}
	for _, c := range caps {
		ids[c.ID] = true
	}
	for _, want := range []uint8{pci.CapIDPowerManagement, pci.CapIDMSIX, pci.CapIDPCIExpress} {
		if !ids[want] {
			t.Errorf("missing capability id 0x%02X in parsed chain", want)
		}
	}
}

// TestConfigSpace_HostSim_BARTopology simulates the host writing 0xFFFFFFFF to each
// BAR through the shadow-BRAM writemask and decoding the reported size.
func TestConfigSpace_HostSim_BARTopology(t *testing.T) {
	cs := buildNVMeHostSimConfigSpace()
	masks := writemaskDWORDs(t, cs)

	type memBAR struct {
		idx  int
		size uint64
	}
	var memBARs []memBAR
	for i := 0; i < 6; i++ {
		off := 0x10 + i*4
		orig := cs.ReadU32(off)
		if orig == 0 {
			continue // unused BAR
		}
		isIO := orig&0x01 != 0
		isMem64 := orig&0x06 == 0x04
		dw := off / 4
		mask := masks[dw]

		// Host sizing probe: write all-1s; shadow BRAM returns write&mask | ro bits.
		readback := uint32(0xFFFFFFFF)&mask | orig&^mask
		if isIO {
			continue
		}
		// 64-bit: also probe upper dword, combine into a 64-bit size mask.
		var size uint64
		if isMem64 && i+1 < 6 {
			upperMask := masks[(off+4)/4]
			upperReadback := uint64(0xFFFFFFFF)&uint64(upperMask) | uint64(cs.ReadU32(off+4))&^uint64(upperMask)
			combined := uint64(readback&^uint32(0xF)) | upperReadback<<32
			size = ^combined + 1
			i++ // consume upper dword of the pair
		} else {
			addr := readback & ^uint32(0xF) // keep decode in 32-bit space
			if addr == 0 {
				size = 1 << 32 // fully writable 32-bit BAR = 4GB
			} else {
				size = uint64(^addr) + 1
			}
		}
		memBARs = append(memBARs, memBAR{idx: i, size: size})
	}
	if len(memBARs) == 0 {
		t.Fatal("no memory BAR present; NVMe BAR0 must exist")
	}

	// BAR0 must be present and >= 4KB (NVMe CAP/MQ registers need at least one page).
	var bar0 uint64
	for _, b := range memBARs {
		if b.idx == 0 {
			bar0 = b.size
		}
	}
	if bar0 == 0 {
		t.Fatal("BAR0 is not a memory BAR")
	}
	if bar0 < 4096 {
		t.Errorf("BAR0 size = %d, want >= 4096 (NVMe requires >= 1 page)", bar0)
	}
	if bar0 != 16*1024 {
		t.Errorf("BAR0 size = %d, want 16384 (fixture)", bar0)
	}
}

// TestConfigSpace_HostSim_MSIXLayout validates MSI-X Table/PBA placement in BAR0.
func TestConfigSpace_HostSim_MSIXLayout(t *testing.T) {
	cs := buildNVMeHostSimConfigSpace()
	info := pci.ParseMSIXCap(cs)
	if info == nil {
		t.Skip("no MSI-X capability; nothing to validate")
	}

	// BAR0 aperture (decode from scrubbed size-encoding).
	b0 := cs.ReadU32(0x10)
	if b0 == 0 {
		t.Fatal("BAR0 disabled but MSI-X references it")
	}
	aperture := uint64(^(b0 & ^uint32(0xF)) + 1)

	for _, bir := range []int{info.TableBIR, info.PBABIR} {
		if bir < 0 || bir > 5 {
			t.Fatalf("BIR %d out of range", bir)
		}
		if cs.ReadU32(0x10+bir*4) == 0 {
			t.Errorf("BIR %d references a disabled BAR", bir)
		}
	}

	tableBytes := uint64(info.TableSize) * 16
	tableEnd := uint64(info.TableOffset) + tableBytes
	if tableEnd > aperture {
		t.Errorf("MSI-X Table [%X,%X) exceeds BAR0 aperture %X", info.TableOffset, tableEnd, aperture)
	}

	pbaBytes := uint64((info.TableSize+63)/64) * 8
	pbaEnd := uint64(info.PBAOffset) + pbaBytes
	if pbaEnd > aperture {
		t.Errorf("MSI-X PBA [%X,%X) exceeds BAR0 aperture %X", info.PBAOffset, pbaEnd, aperture)
	}

	// Table and PBA must not overlap.
	tStart, tEnd := uint64(info.TableOffset), tableEnd
	pStart, pEnd := uint64(info.PBAOffset), pbaEnd
	if tStart < pEnd && pStart < tEnd {
		t.Errorf("MSI-X Table [%X,%X) overlaps PBA [%X,%X)", tStart, tEnd, pStart, pEnd)
	}
}

// TestConfigSpace_HostSim_Writemask verifies identity DWORDs are read-only and the
// Command:Status DWORD exposes W1C status bits while protecting RO status bits.
func TestConfigSpace_HostSim_Writemask(t *testing.T) {
	cs := buildNVMeHostSimConfigSpace()
	dw := writemaskDWORDs(t, cs)

	if dw[0] != 0x00000000 {
		t.Errorf("VID:DID (DWORD0) mask = %08X, want 00000000 (fully read-only)", dw[0])
	}
	if dw[2] != 0x00000000 {
		t.Errorf("Rev:Class (DWORD2) mask = %08X, want 00000000 (read-only)", dw[2])
	}
	// DWORD1 = Command (low 16) | Status (high 16).
	const want = uint32(0xF900FFFF)
	if dw[1] != want {
		t.Errorf("Command:Status mask = %08X, want %08X", dw[1], want)
	}
	// Low half = Command, all writable.
	if dw[1]&0xFFFF != 0xFFFF {
		t.Errorf("Command bits mask = %04X, want FFFF (fully writable)", dw[1]&0xFFFF)
	}
	// High half = Status: W1C bits writable, hard-wired RO bits protected.
	statusMask := dw[1] >> 16
	if statusMask == 0xFFFF {
		t.Error("Status fully writable; RO bits (INTx, CAP list, CLS) must be protected")
	}
	// The W1C error bits (15:11) and DEVSEL[9] must be writable.
	if statusMask&0xF900 != 0xF900 {
		t.Errorf("Status W1C mask = %04X, want F900 bits set", statusMask)
	}
}
