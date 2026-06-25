# PCILeechGen Feature Gap Map

Generated: 2026-06-24

## Scope

This map compares PCILeechGen's current local feature surface with adjacent
public PCIe/FPGA/PCILeech tooling. It is intentionally scoped to legitimate
firmware generation, diagnostics, reproducibility, board support, and lab
validation. It excludes anti-cheat bypass, fake activation, evasion, malware,
game cheating, and arbitrary unauthorized memory access features.

## Current Baseline

PCILeechGen currently provides:

- CLI: `scan`, `check`, `build`, `validate`, `verify-manifest`, `boards`,
  `version`.
- Donor collection through Linux sysfs/VFIO with BAR capture, MSI-X table
  capture, BAR profiling, VFIO diagnostics, and D0 wake attempts.
- Offline builds from `device_context.json`.
- 17 board definitions across PCIeSquirrel, ScreamerM2, CaptainDMA, ZDMA,
  GBOX, NeTV2, ac701, acorn, litefury, and sp605 variants.
- Config-space scrub pipeline, capability pruning, PCIe cap injection,
  write-mask generation, and config-space diff reporting.
- BAR0 generation with dynamic size capping against board BRAM.
- Device-class profiles for NVMe, xHCI, audio, ethernet, GPU, SATA, Wi-Fi,
  MediaTek Wi-Fi, and Thunderbolt.
- SV/TCL/COE/HEX generation, Vivado project/build scripts, optional Vivado
  synthesis, output validation, and build manifest verification.
- New local additions from the prior pass: `boards --json` and manifest
  `board_profile` metadata.

## External Signals Used

- `ret2c/MMIO2Verilog`: mmiotrace log to PCILeech BAR controller generator.
- `Simonrak/verilog-generator`: modular MMIO-log-to-Verilog generator with
  ROM, counter, response-logic, static, and address-check generator concepts.
- `Shocka-Zulu/vfio2verilog`: VFIO/QEMU trace log to BAR controller tool.
- `Crump3tte/bettermmiotrace`: QEMU VFIO tracing setup guide.
- `Ap3x/PCIeConfigSpace`: TLScan config-space parser and write-mask generator.
- `dom0ng/openFPGALoader` fork of upstream openFPGALoader: cross-platform FPGA
  programming utility.
- `WangXuan95/Xilinx-FPGA-PCIe-XDMA-Tutorial`: board bring-up, PCIe pin/XDC,
  XDMA, AXI, and host interaction tutorial material.
- `rick-heig/xilinx_version_ip`: build/version metadata IP concept.
- `rick-heig/eNVMe` and related endpoint/NVMe projects: endpoint/NVMe
  capability breadth, used here only for standards/conformance gap awareness,
  not attack behavior.

## Gap Priority Matrix

| ID | Gap | Priority | Value | Risk | Primary Area |
| --- | --- | --- | --- | --- | --- |
| G1 | Add a real `flash` command | P0 | Completes advertised workflow | Medium | CLI, board metadata |
| G2 | Import external donor captures | P0 | Removes Linux/VFIO-only bottleneck | Medium | donor, pci, codegen |
| G3 | Trace-to-BAR modeling pipeline | P1 | Better BAR behavior from real driver traces | High | donor/mmio, barmodel, svgen |
| G4 | Multi-BAR artifact generation | P1 | Supports devices whose useful state is not BAR0 | High | firmware helpers, barmodel, svgen |
| G5 | Board bring-up wizard and board metadata expansion | P1 | Faster, safer board onboarding | Medium | board, CLI, docs |
| G6 | Simulation and lint harness for generated SV/TCL | P1 | Catches broken generated artifacts before Vivado | Medium | svgen, tclgen, CI |
| G7 | Vivado report summarizer and build diagnostics | P1 | Makes synthesis failures actionable | Low | vivado, CLI |
| G8 | Standards-based PCIe capability coverage roadmap | P2 | Fewer stripped capabilities, better fidelity | High | scrub, pci |
| G9 | NVMe/HDA conformance fixtures | P2 | Turns device-class support into measurable contracts | Medium | svgen, nvme, tests |
| G10 | Reproducible build provenance and version IP | P2 | Better artifact auditability | Low | manifest, svgen, tclgen |
| G11 | Release/CI quality gate cleanup | P2 | Makes `make check` trustworthy | Medium | tests, lint, CI |
| G12 | UX workflow helpers | P3 | Fewer operator mistakes | Low | CLI, docs |

## Detailed Gap Map

### G1. Real `flash` command

Evidence:

- README describes the workflow as `scan -> check -> build -> flash`.
- The current CLI has no `flash` subcommand.
- openFPGALoader provides cross-vendor board/cable/flash programming flows.

Safe scope:

- Add `pcileechgen flash --board <name> --bitstream <file>`.
- Support a small provider interface:
  - `openFPGALoader` provider for supported boards/cables.
  - Vivado Hardware Manager provider as an optional later backend.
  - Dry-run mode that prints exact command.
- Extend `boards.json` with optional flash metadata:
  - `flash_tool`
  - `openfpgaloader_board`
  - `cable`
  - `flash_target`

Non-goals:

- No firmware activation, lock bypass, serial spoofing, or board anti-crack
  logic.

Acceptance criteria:

- `pcileechgen flash --dry-run` produces the expected command for at least
  one board.
- Missing external flasher produces a clear install/path error.
- `boards --json` exposes flash metadata when present.

### G2. Import external donor captures

Evidence:

- PCILeechGen can build from `device_context.json`.
- Ap3x/PCIeConfigSpace parses `.TLScan` config-space captures and emits COE
  and write masks.
- Current `donor.FromJSON` rejects unsupported JSON and there is no importer
  command for TLScan/RWEverything/MMIO dump formats.

Safe scope:

- Add `pcileechgen import` with subcommands:
  - `import tlscan --input donor.tlscan --output device_context.json`
  - `import coe --cfgspace pcileech_cfgspace.coe --writemask ...`
  - `import raw-config --input config.bin`
- Parse imports into `donor.DeviceContext`, not directly into generated SV.
- Preserve provenance in the manifest.

Non-goals:

- No import path should claim donor behavior fidelity when BAR contents or
  runtime traces are absent.

Acceptance criteria:

- Imported contexts pass `validate --json`.
- The importer emits warnings for missing BAR content, MSI/MSI-X data, or
  incomplete extended config space.

### G3. Trace-to-BAR modeling pipeline

Evidence:

- PCILeechGen has `internal/donor/mmio` tracing support and a BAR profiler.
- ret2c/MMIO2Verilog, Simonrak/verilog-generator, and Shocka-Zulu/vfio2verilog
  all focus on turning MMIO traces into BAR response logic.
- Crump3tte/bettermmiotrace documents a QEMU/VFIO tracing path.

Safe scope:

- Add a neutral `trace` workflow:
  - `trace collect` documents/coordinates local Linux trace capture.
  - `trace import --input trace.log --format qemu|mmiotrace`
  - `trace summarize` shows read/write address coverage, repeated patterns,
    volatile registers, and candidate static/register models.
- Feed trace summaries into `barmodel` as evidence, not unchecked generated
  behavior.
- Generate a human-readable `bar_model_report.json` and optional conservative
  SV stubs.

Non-goals:

- Do not auto-generate bypass-oriented device personalities.
- Do not infer malicious host/driver manipulation logic.

Acceptance criteria:

- A fixture trace produces deterministic coverage statistics.
- Static read registers, RW registers, counters, and volatile regions are
  classified separately.
- Generated code has tests and a conservative fallback for unknown addresses.

### G4. Multi-BAR artifact generation

Evidence:

- README and file names focus on `pcileech_bar_zero4k.coe` and BAR0.
- MMIO2Verilog notes a single-BAR limitation and points at dual-BAR controller
  work.
- PCILeechGen captures `BARContents` as a map by BAR index.

Safe scope:

- Generalize output names:
  - `pcileech_bar0.coe`
  - `pcileech_bar1.coe`
  - `bar_model_0.json`
  - `bar_model_1.json`
- Extend SV config with a slice/map of BAR models.
- Generate BAR routing in the controller for multiple implemented BARs.

Non-goals:

- Do not bypass board BRAM limits. Oversized BARs remain gated by board
  capacity and explicit `--force` semantics.

Acceptance criteria:

- Existing BAR0-only outputs remain backward-compatible or have migration
  aliases.
- A synthetic two-BAR donor produces two COEs and a controller route test.

### G5. Board bring-up wizard and board metadata expansion

Evidence:

- Board definitions are embedded JSON with FPGA part, lanes, BRAM, top module,
  project dir, subdir, and TCL files.
- Xilinx XDMA tutorials emphasize PCIe pin assignment, XDC correctness, board
  schematics, lane width, and host-side testing.
- Current board metadata does not include connector/form factor, flash method,
  clock/reset hints, required submodule files, or board-source health checks.

Safe scope:

- Add `pcileechgen board inspect --board <name> --lib-dir <dir>`.
- Add `pcileechgen board scaffold --name ...` for a metadata skeleton only.
- Extend `boards.json` with optional:
  - `form_factor`
  - `flash`
  - `required_sources`
  - `required_ip`
  - `xdc_files`
  - `notes`

Non-goals:

- Do not generate unverified XDC constraints from guesses.

Acceptance criteria:

- `board inspect` checks required source/IP/TCL files and reports missing
  artifacts.
- `boards --json` can serve as machine-readable board inventory for tooling.

### G6. Simulation and lint harness for generated SV/TCL

Evidence:

- The project has strong Go tests but no visible generated-SV lint/simulation
  gate.
- Generated files are complex enough that syntax errors or module conflicts
  can surface late in Vivado.

Safe scope:

- Add optional `pcileechgen validate --hdl` or `pcileechgen hdl-check`.
- Use available tools if installed:
  - `verilator --lint-only`
  - `svlint`
  - `yosys` for simple syntax checks where compatible
  - Vivado TCL parse/dry-run where Vivado exists
- Keep it optional and environment-detected.

Acceptance criteria:

- Missing tools produce skipped checks, not failures.
- Generated fixture SV passes syntax/lint checks in CI when tools are present.

### G7. Vivado report summarizer and build diagnostics

Evidence:

- `internal/vivado/reporter.go` exists, but the CLI build path could expose
  more direct timing/utilization/error summaries.
- Users need actionable failure causes after long synthesis runs.

Safe scope:

- Emit `vivado_report.json` with:
  - timing summary
  - utilization summary
  - critical warnings
  - generated bit/bin paths
  - Vivado version and command line
- Add `pcileechgen report --output pcileech_datastore`.

Acceptance criteria:

- Fixture Vivado logs parse into stable JSON.
- Build failures print the top relevant error/warning lines.

### G8. Standards-based PCIe capability coverage roadmap

Evidence:

- `internal/firmware/scrub/extcap.go` strips unsupported extended
  capabilities including MR-IOV, Resizable BAR, ATS, Page Request, PASID, DPC,
  PTM, and Multicast.

Safe scope:

- Add explicit capability coverage reporting:
  - supported
  - stripped with reason
  - preserved read-only
  - future candidate
- Prioritize safe, passive/read-only handling first, such as reporting PTM/DPC
  presence clearly without pretending to support behavior.

Non-goals:

- Do not claim functional support for IOMMU translations, PASID, page request,
  or containment signaling without a tested implementation.

Acceptance criteria:

- `check` and `validate` report a structured capability coverage table.
- Stripped caps appear in `scrub_diff_report.txt` and manifest metadata.

### G9. NVMe/HDA conformance fixtures

Evidence:

- PCILeechGen has NVMe Identify/Admin responder and HDA DMA/MSI generation.
- Endpoint/NVMe research projects show the breadth of NVMe command and device
  behavior, but many of those behaviors are outside this generator's safe
  scope.

Safe scope:

- Build fixture-driven conformance tests:
  - NVMe Identify fields
  - admin queue command responses
  - Create IO CQ/SQ behavior
  - CC.EN/CSTS.RDY transition timing
  - HDA CORB/RIRB status transitions
- Include generated ROM/HEX consistency checks.

Non-goals:

- No storage attack behaviors, altered reads, remote triggers, or host
  manipulation features.

Acceptance criteria:

- Tests assert command/response bytes from generated artifacts.
- Manifest records enabled device-class features.

### G10. Reproducible build provenance and version IP

Evidence:

- Build manifests now record board profile and hashes.
- `xilinx_version_ip` demonstrates the value of embedding build/version
  metadata in hardware designs.

Safe scope:

- Add optional generated `build_info.sv` with:
  - tool version
  - git commit
  - board name
  - build timestamp
  - manifest digest
- Add deterministic-build mode:
  - fixed entropy seed
  - fixed timestamp
  - manifest comparison

Acceptance criteria:

- `--deterministic --seed <hex>` produces byte-identical generated text
  artifacts for the same input.
- Manifest includes source commit and submodule commit.

### G11. Release/CI quality gate cleanup

Evidence:

- `make check` includes lint, but local `golangci-lint` v2 cannot load the
  current config and v1 reports existing issues.

Safe scope:

- Pin a known-good golangci-lint version or migrate config to v2.
- Fix existing lint failures in small batches.
- Add CI matrix for Go version, tests, vet, build, and optional HDL lint.

Acceptance criteria:

- Fresh checkout can run `make check` without manual tool-version guessing.
- CI publishes binaries and checksums for tagged releases.

### G12. UX workflow helpers

Evidence:

- Common flows require several commands and root/VFIO/Vivado prerequisites.

Safe scope:

- Add `pcileechgen doctor`:
  - Go binary version
  - OS/IOMMU/VFIO status
  - Vivado availability
  - submodule status
  - supported board-source presence
- Add `pcileechgen quickstart --bdf ... --board ... --skip-vivado` as a
  guided dry-run that prints the next command.

Acceptance criteria:

- `doctor` is read-only and safe for non-root execution.
- Root-only checks degrade to warnings with suggested commands.

## Rejected or Deferred Items

Reject:

- Anti-cheat bypass, fake connection, firmware activation, anti-crack,
  arbitrary DMA memory service, game cheat integration, malware analysis
  payload delivery, stealth/evasion feature expansion.

Defer:

- Full ATS/PASID/Page Request support. It requires real IOMMU semantics and
  cannot be safely represented as a cosmetic config-space feature.
- Full NVMe storage semantics beyond testable admin/identify/controller
  behavior. This should be scoped as standards conformance, not attack logic.

## Recommended Delivery Order

1. P0: `flash` command with dry-run and openFPGALoader provider.
2. P0: donor capture importers for TLScan/raw config/COE.
3. P1: `doctor` plus board source/IP inspector.
4. P1: trace import and BAR model report, without auto-aggressive SV behavior.
5. P1: multi-BAR output model.
6. P1: HDL lint/simulation optional gate.
7. P2: Vivado report summarizer and deterministic provenance.
8. P2: capability coverage report and conformance fixtures.
9. P2: CI/lint cleanup and release artifacts.

## First Implementation Slice

The best first slice is G1 plus part of G5:

- Add board flash metadata to `internal/board/boards.json`.
- Add `internal/flash` package with provider interface and openFPGALoader
  implementation.
- Add `pcileechgen flash --dry-run`.
- Add tests for command generation and missing-tool errors.

Why first:

- It closes a README workflow gap.
- It is low-risk because dry-run can be implemented before hardware writes.
- It does not touch generated firmware behavior.
