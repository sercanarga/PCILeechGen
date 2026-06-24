# Full Donor Emulation Foundation Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Add multi-BAR artifacts, explicit register semantics, and deterministic trace reports as the foundation for class-functional donor emulation.

**Architecture:** Keep existing BAR0-compatible firmware generation intact. Add per-BAR COE artifacts and evidence/reporting APIs that later class engines can consume without claiming unsupported behavior.

**Tech Stack:** Go 1.26+, standard library only, existing `go test` package tests.

---

### Task 1: Multi-BAR COE generation

**Files:**
- Modify: `internal/firmware/codegen/codegen.go`
- Modify: `internal/firmware/codegen/codegen_test.go`
- Modify: `internal/firmware/output/writer.go`
- Modify: `internal/firmware/output/writer_test.go`

- [ ] Write failing tests for `GenerateSingleBarContentCOE` and `writeBARContentArtifacts`.
- [ ] Run targeted tests and confirm missing-symbol failures.
- [ ] Implement `GenerateSingleBarContentCOE(barIndex int, data []byte, size int) string`.
- [ ] Keep `GenerateBarContentCOE(map[int][]byte, size int)` as compatibility wrapper.
- [ ] Implement output writer helper that emits `pcileech_bar<N>.coe` for each captured BAR plus `pcileech_bar_zero4k.coe`.
- [ ] Run targeted codegen/output tests.

### Task 2: Register access semantics

**Files:**
- Modify: `internal/firmware/barmodel/model.go`
- Modify: `internal/firmware/barmodel/model_test.go`

- [ ] Write failing test for `BARRegister.AccessKind()`.
- [ ] Run test and confirm missing-symbol failure.
- [ ] Add `RegisterAccessKind` constants and JSON tags on `BARRegister`/`BARModel`.
- [ ] Implement access-kind derivation from existing `RWMask`, `IsRW1C`, and `IsFSMDriven` fields.
- [ ] Run barmodel tests.

### Task 3: MMIO trace reports

**Files:**
- Modify: `internal/donor/mmio/tracer.go`
- Modify: `internal/donor/mmio/tracer_test.go`

- [ ] Write failing test for `BuildReport(sampleTrace())` classification.
- [ ] Run test and confirm missing-symbol failure.
- [ ] Add `TraceReport`, `RegisterSummary`, and `RegisterClassification` types.
- [ ] Implement deterministic sorted register summaries using existing `Analyze` output.
- [ ] Run mmio tests.

### Task 4: Verification

**Files:**
- No new files.

- [ ] Run `go test ./internal/firmware/codegen ./internal/firmware/output ./internal/firmware/barmodel ./internal/donor/mmio`.
- [ ] Run `go test ./...`.
- [ ] Review generated docs/spec for contradictions with implementation.
