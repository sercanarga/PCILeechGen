package output

import (
	"fmt"
	"log/slog"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/donor/mmio"
	"github.com/sercanarga/pcileechgen/internal/firmware"
	"github.com/sercanarga/pcileechgen/internal/firmware/barmodel"
	"github.com/sercanarga/pcileechgen/internal/firmware/codegen"
	"github.com/sercanarga/pcileechgen/internal/firmware/devclass"
	"github.com/sercanarga/pcileechgen/internal/firmware/nvme"
	"github.com/sercanarga/pcileechgen/internal/firmware/svgen"
	"github.com/sercanarga/pcileechgen/internal/pci"
	"github.com/sercanarga/pcileechgen/internal/util"
)

type svArtifact struct {
	filename string
	generate func(cfg *svgen.SVGeneratorConfig) (string, error)
}

var coreSVArtifacts = []svArtifact{
	{"pcileech_bar_impl_device.sv", svgen.GenerateBarImplDeviceSV},
	{"pcileech_tlps128_bar_controller.sv", svgen.GenerateBarControllerSV},
	{"pcileech_bar_impl_msi.sv", svgen.GenerateBarImplMSISV},
	{"tlp_latency_emulator.sv", svgen.GenerateLatencyEmulatorSV},
	{"device_config.sv", svgen.GenerateDeviceConfigSV},
}

func (ow *OutputWriter) writeSVModules(ctx *donor.DeviceContext, scrubbedCS *pci.ConfigSpace, ids firmware.DeviceIDs, entropy uint32, b *board.Board) error {
	cfg, err := ow.buildSVConfig(ctx, scrubbedCS, ids, entropy, b)
	if err != nil {
		return err
	}

	if err := ow.writeCoreSVArtifacts(cfg, scrubbedCS); err != nil {
		return err
	}
	if err := ow.writeConditionalArtifacts(cfg, ctx); err != nil {
		return err
	}

	ow.logSVSummary(cfg)
	return nil
}

func extractMSIInfo(cs *pci.ConfigSpace) *svgen.MSIConfig {
	ptr := int(cs.CapabilityPointer()) & 0xFC
	for ptr != 0 && ptr < 0x100 {
		capID := cs.ReadU8(ptr)
		nextPtr := int(cs.ReadU8(ptr+1)) & 0xFC

		if capID == 0x05 {
			msgCtl := cs.ReadU16(ptr + 2)
			is64bit := (msgCtl & (1 << 7)) != 0

			addrLo := cs.ReadU32(ptr + 4)
			var data uint16
			if is64bit {
				data = cs.ReadU16(ptr + 12)
			} else {
				data = cs.ReadU16(ptr + 8)
			}

			if addrLo == 0 {
				addrLo = 0xFEE00000
			}
			if data == 0 {
				data = 0x0000
			}

			return &svgen.MSIConfig{
				Enabled: (msgCtl & (1 << 0)) != 0,
				AddrLo:  addrLo,
				Data:    data,
			}
		}

		ptr = nextPtr
	}
	return nil
}

// interruptPinToINTxLine maps an interrupt pin (1-4 = INTA#-INTD#) to the
// legacy INTx line number (0-3). It is retained for the future board-top INTx
// wiring path; codegen does not advertise INTx until that integration exists.
func interruptPinToINTxLine(pin uint8) uint8 {
	if pin >= 1 && pin <= 4 {
		return pin - 1
	}
	return 0
}

func (ow *OutputWriter) buildSVConfig(ctx *donor.DeviceContext, scrubbedCS *pci.ConfigSpace, ids firmware.DeviceIDs, entropy uint32, b *board.Board) (*svgen.SVGeneratorConfig, error) {
	barIdx := chooseBARIndexForSV(ctx)
	barData := ctx.BARContents[barIdx]
	var barProfile *donor.BARProfile
	if ctx.BARProfiles != nil {
		if p, ok := ctx.BARProfiles[barIdx]; ok {
			barProfile = p
		}
	}
	var barTraceOverlay *mmio.TraceBAROverlay
	if ctx.BARTraceOverlays != nil {
		barTraceOverlay = ctx.BARTraceOverlays[barIdx]
	}
	slog.Info("BAR selection for SV codegen",
		"bar_index", barIdx,
		"bar_size", len(barData),
		"has_profile", barProfile != nil,
		"has_trace_overlay", barTraceOverlay != nil,
		"class_code", fmt.Sprintf("0x%06X", ctx.Device.ClassCode),
	)
	bm := barmodel.BuildBARModelForDeviceIDWithOverlay(barData, ctx.Device.ClassCode, ids.VendorID, ids.DeviceID, barProfile, barTraceOverlay)
	slog.Info("BAR model built",
		"model_nil", bm == nil,
		"reg_count", func() int {
			if bm != nil {
				return len(bm.Registers)
			}
			return 0
		}(),
	)

	strategy := devclass.StrategyForDevice(ctx.Device.ClassCode, ids.VendorID, ids.DeviceID)
	devClass := ""
	if strategy != nil {
		devClass = strategy.DeviceClass()
	}
	slog.Info("device class resolution", "class", devClass)

	msixTableSize := 0
	if ctx.MSIXData != nil && ctx.MSIXData.TableSize > 0 {
		msixTableSize = ctx.MSIXData.TableSize
	}
	donorBar := firmware.DonorBAR0Demand(ctx, b, msixTableSize)
	bar0Size := firmware.CappedBAR0Size(ctx, b, msixTableSize)
	if bm != nil && bm.Size < bar0Size {
		bm.Size = bar0Size
	}
	if bm == nil && bar0Size > board.DefaultBRAMSize {
		bm = &barmodel.BARModel{Size: bar0Size}
	}

	cfg := &svgen.SVGeneratorConfig{
		DeviceIDs:     ids,
		BARModel:      bm,
		ClassCode:     ctx.Device.ClassCode,
		LatencyConfig: svgen.DefaultLatencyConfig(ctx.Device.ClassCode),
		HasMSIX:       bm != nil,
		BuildEntropy:  entropy,
		PRNGSeeds:     svgen.BuildPRNGSeeds(ids.VendorID, ids.DeviceID, entropy),
		DeviceClass:   devClass,
		Bar0Size:      bar0Size,
	}

	if devClass == devclass.ClassNVMe {
		// Prefer donor-captured Identify pages when present so model/serial/FRU
		// match the real donor exactly; identity fields (VID/SSVID/VER/MDTS) are
		// merged back for config-space consistency.
		if ctx.NVMeIdentify != nil && (ctx.NVMeIdentify.HasController || ctx.NVMeIdentify.HasNamespace) {
			cfg.NVMeIdentify = nvme.BuildIdentifyDataFromCapture(ids, barData, &nvme.Capture{
				Controller:    ctx.NVMeIdentify.Controller,
				Namespace:     ctx.NVMeIdentify.Namespace,
				HasController: ctx.NVMeIdentify.HasController,
				HasNamespace:  ctx.NVMeIdentify.HasNamespace,
			})
		} else {
			cfg.NVMeIdentify = nvme.BuildIdentifyData(ids, barData)
		}
		if len(barData) >= 0x08 {
			capHI := util.ReadLE32(barData, 0x04)
			cfg.NVMeDoorbellStride = capHI & 0x0F
		}
		// Advertise 7 IO SQ + 7 IO CQ pairs. stornvme requests a count via
		// Set Features 0x07 and reads it back; the responder clamps the grant
		// to this value so the readback stays consistent. 0 -> template default.
		cfg.NVMeNumIOQueues = 0x00070007
	}

	if ctx.MSIXData != nil && ctx.MSIXData.TableSize > 0 {
		class := uint32(0)
		if devClass == devclass.ClassNVMe {
			class = 0x0108
		}
		dstrd := cfg.NVMeDoorbellStride
		tableOff, pbaOffset, _ := firmware.MSIXPlacement(cfg.Bar0Size, ctx.MSIXData.TableSize, class, dstrd)
		cfg.MSIXConfig = &svgen.MSIXConfig{
			NumVectors:  ctx.MSIXData.TableSize,
			TableOffset: tableOff,
			PBAOffset:   pbaOffset,
		}
		cfg.HasMSIX = true
		if m := pci.ParseMSIXCap(scrubbedCS); m != nil {
			scrubbedCS.WriteU32(m.CapOffset+4, tableOff&0xFFFFFFF8)
			scrubbedCS.WriteU32(m.CapOffset+8, pbaOffset&0xFFFFFFF8)
		}
	}

	if msiInfo := extractMSIInfo(scrubbedCS); msiInfo != nil {
		cfg.MSIConfig = msiInfo
	}

	// Legacy INTx is intentionally not auto-enabled here. The generator has an
	// INTx module, but the board-top PCIe IP ports are not wired by codegen yet;
	// advertising INTx before that integration would produce a dead artifact.

	bram := b.BRAMSizeOrDefault()
	if issues := ValidateBARSize(donorBar, bram, 0); len(issues) > 0 {
		if !ow.Force {
			return nil, fmt.Errorf("%s", issues[0])
		}
	}
	if cfg.MSIXConfig != nil {
		if issues := ValidateBARSize(donorBar, bram, cfg.MSIXConfig.TableOffset); len(issues) > 0 {
			if !ow.Force {
				return nil, fmt.Errorf("%s", issues[0])
			}
		}
	}

	return cfg, nil
}

func chooseBARIndexForSV(ctx *donor.DeviceContext) int {
	if ctx == nil {
		return 0
	}

	candidates := map[int]struct{}{}
	for idx := range ctx.BARContents {
		candidates[idx] = struct{}{}
	}
	for _, bar := range ctx.BARs {
		if bar.Size > 0 {
			candidates[bar.Index] = struct{}{}
		}
	}
	for idx := range ctx.BARTraceOverlays {
		candidates[idx] = struct{}{}
	}
	for idx := range ctx.MMIOTraces {
		candidates[idx] = struct{}{}
	}

	if len(candidates) == 0 {
		return 0
	}

	bestIdx := -1
	bestScore := -1
	bestSize := -1
	bestNZ := -1
	hasEvidence := false

	for idx := range candidates {
		size := 0
		nz := 0
		if data, ok := ctx.BARContents[idx]; ok {
			size = len(data)
			nz = countNonZeroBytes(data)
		} else {
			for _, bar := range ctx.BARs {
				if bar.Index == idx {
					size = int(bar.Size)
					break
				}
			}
		}

		score := 0
		if trace := ctx.MMIOTraces[idx]; trace != nil {
			score += len(trace.Records)
		}
		if overlay := ctx.BARTraceOverlays[idx]; overlay != nil {
			score += len(overlay.Static) * 3
			score += len(overlay.Sequential) * 6
			score += len(overlay.WriteMask) * 8
			score += len(overlay.RW1CMask) * 12
		}
		if score > 0 {
			hasEvidence = true
		}

		tie := score == bestScore && (size > bestSize || (size == bestSize && (nz > bestNZ || (nz == bestNZ && (bestIdx < 0 || idx < bestIdx)))))
		if score > bestScore || tie {
			bestIdx = idx
			bestScore = score
			bestSize = size
			bestNZ = nz
		}
	}
	if !hasEvidence {
		return firmware.LargestBarIndex(ctx.BARContents)
	}
	return bestIdx
}

func countNonZeroBytes(data []byte) int {
	n := 0
	for _, v := range data {
		if v != 0 {
			n++
		}
	}
	return n
}

func (ow *OutputWriter) writeCoreSVArtifacts(cfg *svgen.SVGeneratorConfig, scrubbedCS *pci.ConfigSpace) error {
	for _, art := range coreSVArtifacts {
		content, err := art.generate(cfg)
		if err != nil {
			return fmt.Errorf("generating %s: %w", art.filename, err)
		}
		if err := ow.writeFile(art.filename, content); err != nil {
			return err
		}
	}

	hex := codegen.GenerateConfigSpaceHex(scrubbedCS)
	return ow.writeFile("config_space_init.hex", hex)
}

func (ow *OutputWriter) writeConditionalArtifacts(cfg *svgen.SVGeneratorConfig, ctx *donor.DeviceContext) error {
	if cfg.MSIXConfig != nil {
		msixSV, err := svgen.GenerateMSIXTableSV(cfg)
		if err != nil {
			return fmt.Errorf("generating pcileech_msix_table.sv: %w", err)
		}
		if err := ow.writeFile("pcileech_msix_table.sv", msixSV); err != nil {
			return err
		}

		entries := ctx.MSIXData.Entries
		if entries == nil {
			entries = make([]pci.MSIXEntry, cfg.MSIXConfig.NumVectors)
			for i := range entries {
				entries[i].Control = 0x01
			}
		}
		if err := ow.writeFile("msix_table_init.hex", codegen.GenerateMSIXTableHex(entries)); err != nil {
			return err
		}
	}

	if cfg.NVMeIdentify != nil {
		nvmeSV, err := svgen.GenerateNVMeResponderSV(cfg)
		if err != nil {
			return fmt.Errorf("generating pcileech_nvme_admin_responder.sv: %w", err)
		}
		if writeErr := ow.writeFile("pcileech_nvme_admin_responder.sv", nvmeSV); writeErr != nil {
			return writeErr
		}

		bridgeSV, err := svgen.GenerateNVMeDMABridgeSV(cfg)
		if err != nil {
			return fmt.Errorf("generating pcileech_nvme_dma_bridge.sv: %w", err)
		}
		if err := ow.writeFile("pcileech_nvme_dma_bridge.sv", bridgeSV); err != nil {
			return err
		}

		if err := ow.writeFile("identify_init.hex", nvme.IdentifyDataToHex(cfg.NVMeIdentify)); err != nil {
			return err
		}
	}

	if cfg.DeviceClass == devclass.ClassAudio && cfg.BARModel != nil {
		hdaSV, err := svgen.GenerateHDARIRBDMASV(cfg)
		if err != nil {
			return fmt.Errorf("generating pcileech_hda_rirb_dma.sv: %w", err)
		}
		if writeErr := ow.writeFile("pcileech_hda_rirb_dma.sv", hdaSV); writeErr != nil {
			return writeErr
		}

		hdaMSISV, err := svgen.GenerateHDAMSISV(cfg)
		if err != nil {
			return fmt.Errorf("generating pcileech_hda_msi.sv: %w", err)
		}
		if err := ow.writeFile("pcileech_hda_msi.sv", hdaMSISV); err != nil {
			return err
		}
	}

	return nil
}
