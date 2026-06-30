package svgen

import (
	"strings"
	"testing"
)

// TestGenerateNVMeBRAMDiskSV_RendersRuntimeSnoop confirms the BRAM disk cache
// no longer takes layout parameters: it discovers the partition start and
// NTFS $MFT LBA at runtime by snooping the boot sector.  No pinned-LBA
// parameters should appear; the snoop machinery must.
func TestGenerateNVMeBRAMDiskSV_RendersRuntimeSnoop(t *testing.T) {
	cfg := testConfig()

	result, err := GenerateNVMeBRAMDiskSV(cfg)
	if err != nil {
		t.Fatalf("GenerateNVMeBRAMDiskSV failed: %v", err)
	}

	for _, want := range []string{
		"module pcileech_bram_disk",
		"parameter integer DISK_WORDS",
		// runtime snoop machinery (code identifiers, not comment prose)
		"boot_detect_now",
		"pin_boot_lba",
		"pin_mft_lba",
		"mft_armed",
		"fmt_mft_lcn", // BPB $MFT cluster field
		"16'hAA55",    // boot signature check
	} {
		if !strings.Contains(result, want) {
			t.Fatalf("bram_disk output should contain %q", want)
		}
	}

	for _, mustNot := range []string{
		"PIN_PARTITION_START_LBA", // no static layout params
		"PIN_FORMAT_METADATA_LBA",
		"PIN_NTFS_MFT_LBA",
		"DISK_LBAS", // dead param removed
	} {
		if strings.Contains(result, mustNot) {
			t.Fatalf("bram_disk output should NOT contain removed param %q", mustNot)
		}
	}
}

// TestNVMeDiskWordsForBRAM36 pins the board->cache-size policy so a 35T gets a
// cache that fits (~8 KiB) and bigger boards get the full 32 KiB.
func TestNVMeDiskWordsForBRAM36(t *testing.T) {
	cases := []struct{ bram36, want int }{
		{0, 0},     // unknown / unsupported part -> refuse
		{25, 0},    // 15T-class: too small
		{50, 8192}, // 35T: ~8 KiB (fits 50 RAMB36 with room for the base design)
		{65, 16384},// 50T
		{135, 32768}, // 75T/100T: full 32 KiB
		{140, 32768}, // 200T
	}
	for _, c := range cases {
		if got := NVMeDiskWordsForBRAM36(c.bram36); got != c.want {
			t.Errorf("NVMeDiskWordsForBRAM36(%d) = %d, want %d", c.bram36, got, c.want)
		}
	}
}
