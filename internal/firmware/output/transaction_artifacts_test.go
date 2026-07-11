package output

import (
	"strings"
	"testing"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/firmware"
	"github.com/sercanarga/pcileechgen/internal/firmware/svgen"
)

func TestTransactionRTLArtifactsAreGenerated(t *testing.T) {
	want := map[string]string{
		"pcileech_tlp_normalizer.sv":         "module pcileech_tlp_normalizer",
		"pcileech_tlps128_bar_rdengine.sv":   "module pcileech_tlps128_bar_rdengine",
		"pcileech_tlps128_bar_controller.sv": ") i_tlp_normalizer(",
		"pcileech_tlp_ur_completer.sv":       "module pcileech_tlp_ur_completer",
	}
	seen := make(map[string]bool, len(want))
	cfg := &svgen.SVGeneratorConfig{}
	for _, artifact := range coreSVArtifacts {
		module, required := want[artifact.filename]
		if !required {
			continue
		}
		seen[artifact.filename] = true
		source, err := artifact.generate(cfg)
		if err != nil {
			t.Fatalf("generate %s: %v", artifact.filename, err)
		}
		if !strings.Contains(source, module) {
			t.Errorf("generated %s does not contain %q", artifact.filename, module)
		}
	}
	for filename := range want {
		if !seen[filename] {
			t.Errorf("coreSVArtifacts does not emit required transaction RTL %s", filename)
		}
	}

	listed := make(map[string]bool)
	for _, filename := range ListOutputFiles() {
		listed[filename] = true
	}
	for filename := range want {
		if !listed[filename] {
			t.Errorf("ListOutputFiles omits generated transaction RTL %s", filename)
		}
	}
}

func TestBuildSVConfigUsesConservativeTransactionLimits(t *testing.T) {
	ctx := makeDonorContext(0x8086, 0x15B7, 0x020000)
	ids := firmware.ExtractDeviceIDs(ctx.ConfigSpace, ctx.ExtCapabilities)
	ow := NewOutputWriter(t.TempDir(), "", 0, 0)
	for _, b := range []*board.Board{
		{},
		{Name: "other", FPGAPart: "xc7a200tsbg484-1"},
	} {
		cfg, err := ow.buildSVConfig(ctx, ctx.ConfigSpace, ids, 0x1234, b)
		if err != nil {
			t.Fatalf("buildSVConfig() error = %v", err)
		}
		if cfg.ReadCompletionBoundaryBytes != 64 {
			t.Errorf("ReadCompletionBoundaryBytes = %d, want 64", cfg.ReadCompletionBoundaryBytes)
		}
		if cfg.MaxPayloadBytes != 128 {
			t.Errorf("MaxPayloadBytes = %d, want 128", cfg.MaxPayloadBytes)
		}
	}
}
