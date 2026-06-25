package output

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/sercanarga/pcileechgen/internal/firmware/devclass"
	"github.com/sercanarga/pcileechgen/internal/firmware/svgen"
)

func (ow *OutputWriter) logSVSummary(cfg *svgen.SVGeneratorConfig) {
	var features []string
	switch cfg.DeviceClass {
	case devclass.ClassNVMe:
		features = append(features, "NVMe FSM")
		if cfg.NVMeIdentify != nil {
			features = append(features, "NVMe Admin Responder", "NVMe DMA Bridge")
		}
	case devclass.ClassXHCI:
		features = append(features, "xHCI FSM")
	case devclass.ClassAudio:
		features = append(features, "HD Audio FSM", "RIRB DMA Bridge")
		if cfg.MSIConfig != nil {
			features = append(features, "MSI Interrupt Gen")
		}
	}
	if cfg.MSIXConfig != nil {
		features = append(features, fmt.Sprintf("MSI-X %d vectors", cfg.MSIXConfig.NumVectors))
	}
	if cfg.BARModel != nil {
		features = append(features, fmt.Sprintf("%d registers", len(cfg.BARModel.Registers)))
	} else {
		features = append(features, "BRAM fallback")
	}
	features = append(features, "latency emulator", "interrupt controller")
	slog.Info("SV modules generated", "features", strings.Join(features, ", "))
}
