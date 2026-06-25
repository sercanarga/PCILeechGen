# External Adaptation Policy

PCILeechGen generates FPGA firmware for PCIe DMA devices in the ufrisk/pcileech
ecosystem. This document governs how feature and profile ideas from external
PCIe/FPGA projects may be proposed for adoption. The goal is to harden
PCILeechGen by treating external repos as *evidence*, not as code sources.

## Guardrails — what we do not adopt

External adaptation must never add the following. These categories are rejected
on intake regardless of source, license, or apparent usefulness:

- Anti-cheat bypasses, game cheats, or evasion of anti-cheat / detection
  mechanisms.
- Ransomware, malware, or persistence mechanisms.
- Cracks, license bypasses, or serial spoofing.
- Arbitrary unauthorized memory access features.
- Unverified hard-coded spoof profiles copied from third-party repos.
- Copied third-party source without recorded license compatibility,
  attribution, and exact provenance.
- Unverified XDC generation or board pin guesses.
- Behavior-changing profile or HDL generation without tests and a
  manifest/provenance trail.

Intake records whose `scope` or `imported_artifacts` touch these categories
must set `safety_classification: rejected` with a `rejection_reason` that names
the violated category.

## Default posture

No external code or profile is blessed as safe by default. Every intake record
begins at `safety_classification: needs-review` until a human reviewer with
context on the target board and use case promotes it to `safe` or `rejected`.
The validator enforces that required provenance fields are present; it does not
judge safety.

## Intake record schema

Each external idea is recorded as one JSON intake record under
`internal/intake/testdata` (for fixtures) or an operator-chosen intake
directory. Fields:

| Field                  | Required | Description |
| ---------------------- | -------- | ----------- |
| `source_url`           | yes      | Canonical URL of the upstream repo/commit/file. |
| `license`              | yes      | SPDX identifier (e.g. `MIT`, `GPL-2.0-only`) or the literal `unknown`. `unknown` is permitted but must be resolved before any artifact import. |
| `commit_ref`           | no       | Git commit, tag, or branch the record pins. |
| `scope`                | no       | Free-text description of what aspect of PCILeechGen the idea touches. |
| `safety_classification` | no      | One of `safe`, `needs-review`, `rejected`. Defaults to `needs-review` when omitted. |
| `imported_artifacts`   | no       | List of artifact paths or descriptors brought in from the source. |
| `attribution`          | no       | Credit line or upstream author/maintainer reference. |
| `rejection_reason`     | required when `safety_classification == rejected` | Names the guardrail category that was violated. |

### Validation rules

- `source_url` and `license` are required. Missing either is a validation
  error.
- `rejection_reason` must be non-empty when `safety_classification` is
  `rejected`.
- The validator does NOT mark anything `safe`. It only checks structural
  completeness. Promotion to `safe` is a human act recorded outside this
  schema.

## ponytail: ceiling and upgrade path

This policy is deliberately a single markdown file plus a tiny validator. The
ceiling: it does not encode a machine-readable enumeration of guardrail
keywords, nor does it sign records. Upgrade path: if intake volume grows,
replace the free-text `scope` with a controlled vocabulary, add a registry of
banned source domains, and introduce signed review records. Do not add any of
that until a second intake record actually arrives.