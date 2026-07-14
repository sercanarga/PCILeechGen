package output

import (
	"encoding/binary"
	"fmt"
	"log/slog"
	"strings"

	"github.com/sercanarga/pcileechgen/internal/board"
	"github.com/sercanarga/pcileechgen/internal/donor"
	"github.com/sercanarga/pcileechgen/internal/firmware"
	"github.com/sercanarga/pcileechgen/internal/firmware/barmodel"
	"github.com/sercanarga/pcileechgen/internal/firmware/codegen"
	"github.com/sercanarga/pcileechgen/internal/firmware/devclass"
	"github.com/sercanarga/pcileechgen/internal/firmware/nvme"
	"github.com/sercanarga/pcileechgen/internal/firmware/svgen"
	"github.com/sercanarga/pcileechgen/internal/pci"
	"github.com/sercanarga/pcileechgen/internal/util"
)

// svArtifact describes one SV file to generate.
type svArtifact struct {
	filename string
	generate func(cfg *svgen.SVGeneratorConfig) (string, error)
}

// coreSVArtifacts are always generated.
var coreSVArtifacts = []svArtifact{
	{"pcileech_lifecycle_service.sv", svgen.GenerateLifecycleServiceSV},
	{"pcileech_dma_tag_service.sv", svgen.GenerateDMATagServiceSV},
	{"pcileech_interrupt_service.sv", svgen.GenerateInterruptServiceSV},
	{"pcileech_bar_impl_device.sv", svgen.GenerateBarImplDeviceSV},
	{"pcileech_tlps128_bar_controller.sv", svgen.GenerateBarControllerSV},
	{"pcileech_tlp_normalizer.sv", svgen.GenerateTransactionNormalizerSV},
	{"pcileech_tlps128_bar_rdengine.sv", svgen.GenerateBarReadEngineSV},
	{"pcileech_tlp_ur_completer.sv", svgen.GenerateURCompleterSV},
	{"pcileech_bar_rsp_arbiter.sv", svgen.GenerateBarRspArbiterSV},
	{"tlp_latency_emulator.sv", svgen.GenerateLatencyEmulatorSV},
	{"device_config.sv", svgen.GenerateDeviceConfigSV},
}

func (ow *OutputWriter) writeSVModules(ctx *donor.DeviceContext, scrubbedCS *pci.ConfigSpace, cfg *svgen.SVGeneratorConfig) error {
	if err := ow.writeCoreSVArtifacts(cfg, scrubbedCS); err != nil {
		return err
	}
	if err := ow.writeConditionalArtifacts(cfg, ctx); err != nil {
		return err
	}
	ow.logSVSummary(cfg)
	return nil
}

// extractMSIInfo parses the MSI capability from scrubbed config space.
// Returns nil if MSI is absent or not yet programmed by the host.
func extractMSIInfo(cs *pci.ConfigSpace) *svgen.MSIConfig {
	// Walk the capability linked list
	ptr := int(cs.CapabilityPointer()) & 0xFC
	for ptr != 0 && ptr < 0x100 {
		capID := cs.ReadU8(ptr)
		nextPtr := int(cs.ReadU8(ptr+1)) & 0xFC

		if capID == 0x05 { // CapIDMSI
			msgCtl := cs.ReadU16(ptr + 2)
			return &svgen.MSIConfig{
				Enabled: (msgCtl & (1 << 0)) != 0,
			}
		}

		ptr = nextPtr
	}
	return nil
}

type barInterval struct {
	start uint64
	end   uint64
}

func alignUp(value, alignment uint64) uint64 {
	if alignment <= 1 {
		return value
	}
	return (value + alignment - 1) &^ (alignment - 1)
}

func placeBARRegion(model *barmodel.BARModel, preferred, length, alignment uint64, reserved []barInterval) (uint32, error) {
	if model == nil || model.Size <= 0 {
		return 0, fmt.Errorf("BAR endpoint is absent")
	}
	limit := uint64(model.Size)
	occupied := append([]barInterval(nil), reserved...)
	for _, reg := range model.Registers {
		width := reg.Width
		if width <= 0 {
			width = 4
		}
		occupied = append(occupied, barInterval{start: uint64(reg.Offset), end: uint64(reg.Offset) + uint64(width)})
	}
	fits := func(start uint64) bool {
		if start%alignment != 0 || start > limit || length > limit-start {
			return false
		}
		end := start + length
		for _, used := range occupied {
			if start < used.end && used.start < end {
				return false
			}
		}
		return start <= uint64(^uint32(0)) && end <= uint64(^uint32(0))+1
	}

	candidates := []uint64{preferred, 0}
	for _, used := range occupied {
		candidates = append(candidates, alignUp(used.end, alignment))
	}
	for _, candidate := range candidates {
		if fits(candidate) {
			return uint32(candidate), nil
		}
	}
	return 0, fmt.Errorf("BAR%d aperture %#x has no %#x-byte aligned gap for MSI-X region",
		model.BIR, limit, length)
}

func (ow *OutputWriter) buildSVConfig(ctx *donor.DeviceContext, scrubbedCS *pci.ConfigSpace, ids firmware.DeviceIDs, entropy uint32, b *board.Board) (*svgen.SVGeneratorConfig, error) {
	strategy := devclass.StrategyForClassAndVendor(ctx.Device.ClassCode, ids.VendorID)
	devClass := ""
	preferredBIR := 0
	if strategy != nil {
		devClass = strategy.DeviceClass()
		if profile := strategy.Profile(); profile != nil {
			preferredBIR = profile.PreferredBAR
		}
	}

	models, err := barmodel.BuildBARModels(ctx.BARs, ctx.BARContents, ctx.BARProfiles,
		ctx.Device.ClassCode, preferredBIR)
	if err != nil {
		return nil, fmt.Errorf("invalid donor BAR topology: %w", err)
	}
	primary := barmodel.ModelForBIR(models, preferredBIR)
	if primary == nil && len(models) > 0 {
		primary = models[0]
	}

	msixTableSize := 0
	if ctx.MSIXData != nil && ctx.MSIXData.TableSize > 0 {
		msixTableSize = ctx.MSIXData.TableSize
	}
	donorBar := firmware.DonorBAR0Demand(ctx, b, msixTableSize)
	bar0Size := firmware.CappedBAR0Size(ctx, b, msixTableSize)
	if bar0 := barmodel.ModelForBIR(models, 0); bar0 != nil && bar0.Size < bar0Size {
		bar0.Size = bar0Size
	}

	barData := []byte(nil)
	if primary != nil {
		barData = ctx.BARContents[primary.BIR]
	}
	slog.Info("BAR topology for SV codegen",
		"models", len(models),
		"preferred_bir", preferredBIR,
		"primary_bir", func() int {
			if primary != nil {
				return primary.BIR
			}
			return -1
		}(),
		"class_specific", primary != nil && primary.ClassSpecific,
		"class_code", fmt.Sprintf("0x%06X", ctx.Device.ClassCode),
	)

	var compiledRules *svgen.CompiledBehavior
	if ctx.BehaviorRules != nil {
		if ownershipErr := svgen.ValidateBehaviorRuleOwnership(ctx.BehaviorRules, devClass); ownershipErr != nil {
			return nil, ownershipErr
		}
		ruleModel := barmodel.ModelForBIR(models, ctx.BehaviorRules.BARIndex)
		if ruleModel == nil {
			return nil, fmt.Errorf("behavior rules target BAR%d but no generated endpoint exists at that BIR", ctx.BehaviorRules.BARIndex)
		}
		if ctx.BehaviorRules.BARSize > ruleModel.Size {
			return nil, fmt.Errorf("behavior BAR size %d exceeds generated BAR%d size %d", ctx.BehaviorRules.BARSize, ruleModel.BIR, ruleModel.Size)
		}
		applied, applyErr := barmodel.ApplyBehaviorRules(ruleModel, ctx.BehaviorRules)
		if applyErr != nil {
			return nil, fmt.Errorf("apply behavior rules: %w", applyErr)
		}
		for i := range models {
			if models[i].BIR == applied.BIR {
				models[i] = applied
				break
			}
		}
		if primary != nil && primary.BIR == applied.BIR {
			primary = applied
		}
		var compileErr error
		compiledRules, compileErr = svgen.CompileBehaviorRules(ctx.BehaviorRules, applied)
		if compileErr != nil {
			return nil, fmt.Errorf("compile behavior rules: %w", compileErr)
		}
	}

	cfg := &svgen.SVGeneratorConfig{
		DeviceIDs:                   ids,
		DonorCapabilities:           extractDonorCapabilities(scrubbedCS),
		BARModels:                   models,
		DonorBARTopology:            true,
		BARModel:                    primary,
		ClassCode:                   ctx.Device.ClassCode,
		LatencyConfig:               svgen.DefaultLatencyConfig(ctx.Device.ClassCode),
		HasMSIX:                     ctx.MSIXData != nil && ctx.MSIXData.TableSize > 0,
		BuildEntropy:                entropy,
		PRNGSeeds:                   svgen.BuildPRNGSeeds(ids.VendorID, ids.DeviceID, entropy),
		DeviceClass:                 devClass,
		Bar0Size:                    bar0Size,
		ReadCompletionBoundaryBytes: 64,
		MaxPayloadBytes:             128,
		BehaviorRules:               ctx.BehaviorRules,
		CompiledBehavior:            compiledRules,
	}
	if primary == nil {
		cfg.LatencyConfig = nil
	}

	if devClass == devclass.ClassNVMe && cfg.HasClassEndpoint() {
		var identity *nvme.ControllerIdentity
		if ctx.NVMeIdentity != nil {
			identity = &nvme.ControllerIdentity{
				Serial: ctx.NVMeIdentity.Serial,
				Model:  ctx.NVMeIdentity.Model,
				FWRev:  ctx.NVMeIdentity.FWRev,

				RawControllerIdent: ctx.NVMeIdentity.RawControllerIdent,
				RawNamespaceIdent:  ctx.NVMeIdentity.RawNamespaceIdent,
			}
		}
		cfg.NVMeIdentify = nvme.BuildIdentifyData(ids, barData, identity)
		cfg.NVMeSMART = nvme.BuildSMART()

		// Normalize donor NSZE to 512B units: firmware always serves LBADS=9.
		if nsze := binary.LittleEndian.Uint64(cfg.NVMeIdentify.Namespace[0x000:0x008]); nsze > 0 {
			lbaf0 := binary.LittleEndian.Uint32(cfg.NVMeIdentify.Namespace[0x0C0:0x0C4])
			if donorLBADS := uint((lbaf0 >> 16) & 0xFF); donorLBADS >= 9 {
				nsze <<= (donorLBADS - 9)
			}
			cfg.NVMeAdvertisedLBAs = nsze
		}
		if len(barData) >= 0x08 {
			cfg.NVMeDoorbellStride = nvme.DoorbellStrideFromCAP(util.ReadLE32(barData, 0x04))
		}
		// refuse early if board can't fit a cache
		cfg.NVMeDiskWords = svgen.NVMeDiskWordsForBRAM36(b.BRAM36Capacity())
		if cfg.NVMeDiskWords == 0 {
			return nil, fmt.Errorf(
				"board %q (%s) has insufficient block RAM for NVMe disk emulation; "+
					"use an Artix-7 75T/100T/200T board (35T/50T lack the block RAM), "+
					"200T for the full disk cache",
				b.Name, b.FPGAPart)
		}
	}

	if ctx.MSIXData != nil && ctx.MSIXData.TableSize > 0 {
		tableModel := barmodel.ModelForBIR(models, ctx.MSIXData.TableBIR)
		if tableModel == nil {
			return nil, fmt.Errorf("MSI-X table advertises BIR %d, but that BIR has no generated memory endpoint",
				ctx.MSIXData.TableBIR)
		}
		pbaModel := barmodel.ModelForBIR(models, ctx.MSIXData.PBABIR)
		if pbaModel == nil {
			return nil, fmt.Errorf("MSI-X PBA advertises BIR %d, but that BIR has no generated memory endpoint",
				ctx.MSIXData.PBABIR)
		}

		var tableReserved []barInterval
		if devClass == devclass.ClassNVMe && tableModel.BIR == preferredBIR {
			stride := uint64(4) << cfg.NVMeDoorbellStride
			tableReserved = append(tableReserved, barInterval{
				start: uint64(board.DefaultBRAMSize),
				end:   uint64(board.DefaultBRAMSize) + 4*stride,
			})
		}
		tableBytes := uint64(ctx.MSIXData.TableSize * 16)
		tableOff, err := placeBARRegion(tableModel, uint64(ctx.MSIXData.TableOffset),
			tableBytes, 16, tableReserved)
		if err != nil {
			return nil, fmt.Errorf("placing MSI-X table: %w", err)
		}

		pbaBytes := uint64((ctx.MSIXData.TableSize + 63) / 64 * 8)
		if pbaBytes < 8 {
			pbaBytes = 8
		}
		var pbaReserved []barInterval
		if pbaModel.BIR == tableModel.BIR {
			pbaReserved = append(pbaReserved, barInterval{
				start: uint64(tableOff),
				end:   uint64(tableOff) + tableBytes,
			})
		}
		if devClass == devclass.ClassNVMe && pbaModel.BIR == preferredBIR {
			stride := uint64(4) << cfg.NVMeDoorbellStride
			pbaReserved = append(pbaReserved, barInterval{
				start: uint64(board.DefaultBRAMSize),
				end:   uint64(board.DefaultBRAMSize) + 4*stride,
			})
		}
		pbaOff, err := placeBARRegion(pbaModel, uint64(ctx.MSIXData.PBAOffset),
			pbaBytes, 8, pbaReserved)
		if err != nil {
			return nil, fmt.Errorf("placing MSI-X PBA: %w", err)
		}

		cfg.MSIXConfig = &svgen.MSIXConfig{
			NumVectors:  ctx.MSIXData.TableSize,
			TableBIR:    tableModel.BIR,
			TableOffset: tableOff,
			PBABIR:      pbaModel.BIR,
			PBAOffset:   pbaOff,
		}
		cfg.HasMSIX = true
		if m := pci.ParseMSIXCap(scrubbedCS); m != nil {
			scrubbedCS.WriteU32(m.CapOffset+4, (tableOff&0xFFFFFFF8)|uint32(tableModel.BIR))
			scrubbedCS.WriteU32(m.CapOffset+8, (pbaOff&0xFFFFFFF8)|uint32(pbaModel.BIR))
		}
	}

	// Extract MSI capability for interrupt generation.
	// MSI is the primary interrupt mechanism for HDA devices.
	if msiInfo := extractMSIInfo(scrubbedCS); msiInfo != nil {
		cfg.MSIConfig = msiInfo
	}

	bram := b.BRAMSizeOrDefault()
	// Validate *donor demand* (may exceed) against board BRAM; error unless --force.
	// (bar0Size is the Capped value actually used for artifacts/scrub/SV.)
	if issues := ValidateBARSize(donorBar, bram, 0); len(issues) > 0 {
		if !ow.Force {
			return nil, fmt.Errorf("%s", issues[0])
		}
	}

	return cfg, nil
}

// writeCoreSVArtifacts generates core SV modules and config space HEX.
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

// writeConditionalArtifacts generates MSI-X and NVMe artifacts when applicable.
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
				entries[i].Control = 0x01 // masked
			}
		}
		if err := ow.writeFile("msix_table_init.hex", codegen.GenerateMSIXTableHex(entries)); err != nil {
			return err
		}
	}

	if cfg.MSIConfig != nil {
		msiEpSV, err := svgen.GenerateBarImplMSISV(cfg)
		if err != nil {
			return fmt.Errorf("generating pcileech_bar_impl_msi.sv: %w", err)
		}
		if err := ow.writeFile("pcileech_bar_impl_msi.sv", msiEpSV); err != nil {
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
		if werr := ow.writeFile("pcileech_nvme_dma_bridge.sv", bridgeSV); werr != nil {
			return werr
		}

		diskSV, err := svgen.GenerateNVMeBRAMDiskSV(cfg)
		if err != nil {
			return fmt.Errorf("generating pcileech_bram_disk.sv: %w", err)
		}
		if err := ow.writeFile("pcileech_bram_disk.sv", diskSV); err != nil {
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
	}

	return nil
}

// logSVSummary prints a summary of generated SV features.
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
	if cfg.DonorCapabilities.HasPMCap {
		features = append(features, "donor PM cap")
	}
	if cfg.DonorCapabilities.HasMSICap {
		features = append(features, "donor MSI cap")
	}
	if cfg.DonorCapabilities.HasMSIXCap {
		features = append(features, "donor MSI-X cap")
	}
	if cfg.DonorCapabilities.HasPCIeCap {
		features = append(features, "donor PCIe cap")
	}
	features = append(features, "latency emulator", "interrupt controller")
	slog.Info("SV modules generated", "features", strings.Join(features, ", "))
}
