# Full Donor Emulation Foundation Design

## Goal
Build the shared substrate needed for class-functional donor emulation across NVMe, xHCI, HDA/audio, Ethernet, SATA, Wi-Fi, Thunderbolt, and limited GPU facades.

## Scope for this slice
This slice does not claim full class behavior. It adds the reusable foundation every class engine needs:

1. Multi-BAR artifact support: emit per-BAR COE files for captured donor BAR contents while keeping the legacy BAR0-compatible filename.
2. Register model metadata: expose register access semantics as explicit model data instead of implicit flags.
3. Trace report data model: turn MMIO traces into deterministic per-register classifications suitable for later BAR model synthesis.

## Non-goals
- No anti-cheat bypass, stealth/evasion feature, arbitrary DMA service, firmware activation, or spoofing claims.
- No claim that a generated bitstream fully emulates an unsupported class.
- No broad class engines in this slice; NVMe/xHCI/HDA/NIC/SATA/Wi-Fi engines come after this foundation.

## Architecture

### Multi-BAR artifacts
`internal/firmware/output` writes all captured `DeviceContext.BARContents` entries as `pcileech_bar<N>.coe`. Existing `pcileech_bar_zero4k.coe` remains as a compatibility alias using the current largest-BAR behavior, so current templates and users do not break.

### Register model metadata
`internal/firmware/barmodel.BARRegister` gains JSON tags and an `AccessKind()` method derived from current flags:

- `read_only`: `RWMask == 0`
- `read_write`: writable normal register
- `rw1c`: write-one-to-clear
- `fsm`: driven by a dedicated device FSM

This keeps current template behavior unchanged while making model output machine-readable.

### Trace reports
`internal/donor/mmio` gains `BuildReport(trace)` with sorted register summaries. `internal/firmware/output` writes `bar_model_report.json` when a donor context contains MMIO traces. `pcileechgen trace import` converts a local mmiotrace log into the same report format, and `pcileechgen build --trace BAR=trace.log` carries trace evidence into `DeviceContext.MMIOTraces` so normal builds emit the report. Trace import can use a BAR physical base to preserve offsets beyond 4KB, which is required for NVMe doorbells. Classification is deliberately conservative:

- `static_read`: reads only, one observed value
- `volatile_read`: reads only, multiple values
- `polled`: repeated reads detected
- `write_only`: writes only
- `read_write`: both reads and writes

Trace-observed registers are merged into BAR models only when the class/profile model does not already cover that offset. Existing class semantics win; trace registers get reset values from the last observed value, read-only masks for read classifications, and full write masks only when writes were observed.

### NVMe trace evidence
`internal/firmware/nvme` derives small, explicit evidence from traces: Admin Queue register writes (`AQA`, `ASQ`, `ACQ`) and observed SQ0/CQ0 doorbell spacing. When BAR data does not provide CAP.DSTRD, `output.buildSVConfig` uses trace-derived doorbell stride so generated NVMe doorbell hit logic can match real donor traces. This does not add new NVMe command behavior; it only preserves observed initialization geometry.

## Testing
- Codegen test proves single-BAR COE generation uses the requested BAR data instead of the largest BAR.
- Output test proves multiple captured BARs produce multiple deterministic artifact names.
- BAR model test proves access-kind derivation.
- MMIO test proves trace report classification is deterministic.
- Output test proves `bar_model_report.json` is deterministic and parseable when traces are present.
- CLI test proves `trace import` writes a parseable report from a local mmiotrace fixture.
- Build CLI test proves `--trace BAR=path` loads trace records into the donor context with BAR metadata.
- BAR model/output tests prove trace-only registers become conservative SV model registers without overriding existing class registers.
- MMIO/build tests prove BAR base-aware trace parsing preserves offsets beyond 4KB.
- NVMe/output tests prove trace evidence extracts Admin Queue writes and doorbell stride, then wires stride into SV config.
