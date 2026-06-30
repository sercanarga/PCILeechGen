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
		"localparam integer DISK_WORDS",
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
