# VFIO-user Firmware End-to-End Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Replace the hard-coded NVMe-only VFIO-user prototype with a generated-artifact-driven harness that validates all ten donor fixtures through firmware generation, HDL simulation, QEMU enumeration, and class-specific Linux driver initialization.

**Architecture:** A Python matrix runner builds real PCILeechGen artifacts and launches a generic libvfio-user C server. The server loads `device_model.json`, delegates dynamic MMIO to one behavior module selected by PCI class, and exposes generated config space and BAR layout to QEMU. A deterministic guest probe emits JSON Lines consumed by the runner; Cocotb remains the generated-HDL execution layer.

**Tech Stack:** Go 1.26.2 CLI, C11, libvfio-user, json-c, cmocka, Python 3 standard library, QEMU, BusyBox initramfs, Cocotb, Verilator, Docker.

## Global Constraints

- Cover exactly `audio`, `ethernet`, `generic`, `gpu`, `multibar`, `nvme`, `sata`, `thunderbolt`, `wifi`, and `xhci`.
- Do not invoke or detect Vivado.
- Treat generated `device_model.json`, `build_manifest.json`, and config-space reset data as runtime sources of truth.
- Do not duplicate fixture vendor IDs, device IDs, class codes, BAR sizes, or capability offsets in C.
- Preserve user-owned untracked work and review the existing `vfio-user/` files before replacing or moving them.
- Do not hide command failures with unconditional `|| true`.
- Keep only comments that explain protocol constraints, ABI requirements, or non-obvious rationale.
- Follow red-green-refactor for every behavior change.

## File Structure

- `vfio-user/include/device_model.h`: parsed generated-model types and loader API.
- `vfio-user/include/device_behavior.h`: class behavior contract.
- `vfio-user/include/vfio_device.h`: generic server lifecycle contract.
- `vfio-user/src/device_model.c`: JSON/base64/config/capability/BAR validation.
- `vfio-user/src/vfio_device.c`: libvfio-user setup, regions, DMA, IRQs, event loop.
- `vfio-user/src/behavior_static.c`: generated static BAR fallback.
- `vfio-user/src/behavior_nvme.c`: extracted NVMe queue/register behavior.
- `vfio-user/src/behavior_ahci.c`: AHCI reset and port-0 initialization behavior.
- `vfio-user/src/behavior_xhci.c`: xHCI capability/operational reset behavior.
- `vfio-user/src/behavior_profiled.c`: bounded register-state behavior for audio, Ethernet, Wi-Fi, GPU, and Thunderbolt fixtures.
- `vfio-user/src/main.c`: argument parsing and composition only.
- `vfio-user/tests/test_device_model.c`: malformed/generated model tests.
- `vfio-user/tests/test_behavior_*.c`: driver-sequence unit tests.
- `vfio-user/tests/test_result_parser.py`: structured guest-result tests.
- `vfio-user/guest/init`: minimal guest boot/probe entrypoint.
- `vfio-user/guest/probe.py`: JSON Lines PCI/driver/device checks.
- `vfio-user/matrix.py`: fixture inventory, build/sim/server/QEMU orchestration, reports.
- `vfio-user/Dockerfile`: pinned Linux test environment.
- `vfio-user/Makefile`: unit, image, single-device, and matrix targets.
- `vfio-user/README.md`: prerequisites, boundaries, and commands.

---

### Task 1: Freeze the fixture matrix and failure semantics

**Files:**
- Create: `vfio-user/matrix.py`
- Create: `vfio-user/tests/test_matrix.py`
- Modify: `vfio-user/Makefile`

**Interfaces:**
- Produces: `Case(name: str, fixture: Path, board: str, behavior: str, mandatory_probe: str)`.
- Produces: `run_command(argv: Sequence[str], log_path: Path, timeout: int) -> None`, raising `CaseFailure` on nonzero exit or timeout.
- Produces: `CASES`, the ten immutable cases; NVMe uses `ac701_ft601`, all others use `PCIeSquirrel`.

- [ ] **Step 1: Write failing matrix and failure-propagation tests**

```python
def test_matrix_covers_every_generated_fixture(self):
    self.assertEqual(
        set(matrix.CASES),
        {"audio", "ethernet", "generic", "gpu", "multibar", "nvme",
         "sata", "thunderbolt", "wifi", "xhci"},
    )

def test_run_command_rejects_nonzero_exit(self):
    with tempfile.TemporaryDirectory() as tmp:
        with self.assertRaisesRegex(matrix.CaseFailure, "exit status 7"):
            matrix.run_command(
                [sys.executable, "-c", "raise SystemExit(7)"],
                Path(tmp) / "command.log",
                timeout=5,
            )
```

- [ ] **Step 2: Run RED**

Run: `python3 -m unittest vfio-user/tests/test_matrix.py -v`

Expected: import failure because `vfio-user/matrix.py` does not exist.

- [ ] **Step 3: Implement the typed matrix and strict command runner**

Use a frozen dataclass, `subprocess.run(..., check=False, timeout=timeout)`, explicit log files, and `start_new_session=True`. On timeout, terminate the process group before raising `CaseFailure`. Do not use `shell=True`.

- [ ] **Step 4: Run GREEN**

Run: `python3 -m unittest vfio-user/tests/test_matrix.py -v`

Expected: both tests pass.

- [ ] **Step 5: Commit**

```bash
git add vfio-user/matrix.py vfio-user/tests/test_matrix.py vfio-user/Makefile
git commit -m "test: define vfio-user device matrix"
```

### Task 2: Load and validate generated device artifacts

**Files:**
- Create: `vfio-user/include/device_model.h`
- Create: `vfio-user/src/device_model.c`
- Create: `vfio-user/tests/test_device_model.c`
- Create: `vfio-user/tests/fixtures/invalid_bar.json`
- Modify: `vfio-user/Makefile`

**Interfaces:**
- Produces: `int device_model_load(const char *artifact_dir, struct device_model **out, char *err, size_t err_len)`.
- Produces: `void device_model_free(struct device_model *model)`.
- Produces: `int device_model_validate(const struct device_model *model, char *err, size_t err_len)`.
- `struct device_model` owns a 4096-byte config image, up to six `device_bar` entries, interrupt metadata, and the generated PCI function identity.

- [ ] **Step 1: Write failing loader tests**

```c
static void loads_generated_nvme_model(void **state)
{
    struct device_model *model = NULL;
    char err[256] = {0};
    assert_int_equal(device_model_load("tests/cocotb/out_nvme", &model, err, sizeof(err)), 0);
    assert_int_equal(model->vendor_id, 0x144d);
    assert_int_equal(model->device_id, 0xa809);
    assert_int_equal(model->bars[0].size, 16384);
    assert_memory_equal(model->config_space, "\x4d\x14\x09\xa8", 4);
    device_model_free(model);
}

static void rejects_overlapping_or_non_power_of_two_bars(void **state)
{
    struct device_model *model = NULL;
    char err[256] = {0};
    assert_true(device_model_load("vfio-user/tests/fixtures/invalid_bar", &model, err, sizeof(err)) < 0);
    assert_non_null(strstr(err, "BAR"));
}
```

- [ ] **Step 2: Run RED**

Run: `make -C vfio-user test-device-model`

Expected: compile failure for missing `device_model.h`.

- [ ] **Step 3: Implement strict loading**

Parse `build_manifest.json` first and verify the declared size and SHA-256 of `device_model.json`. Parse schema version 1, decode `config_space.reset_image`, map BIR 0-5, and validate power-of-two BAR sizes, 64-bit BAR pairing, MSI-X table/PBA bounds, capability alignment, unique offsets, and finite next pointers. Fail closed on unknown schema versions or truncated config images.

- [ ] **Step 4: Run GREEN and generated-fixture sweep**

Run: `make -C vfio-user test-device-model && for d in tests/cocotb/out_{audio,ethernet,generic,gpu,multibar,nvme,sata,thunderbolt,wifi,xhci}; do vfio-user/build/model-check "$d"; done`

Expected: unit tests pass and all ten generated directories print `valid`.

- [ ] **Step 5: Commit**

```bash
git add vfio-user/include/device_model.h vfio-user/src/device_model.c vfio-user/tests vfio-user/Makefile
git commit -m "feat: load generated vfio device models"
```

### Task 3: Introduce the behavior boundary and static BAR model

**Files:**
- Create: `vfio-user/include/device_behavior.h`
- Create: `vfio-user/src/behavior_static.c`
- Create: `vfio-user/tests/test_behavior_static.c`
- Modify: `vfio-user/Makefile`

**Interfaces:**
- Produces:

```c
struct device_behavior {
    void *state;
    int (*reset)(void *state);
    ssize_t (*read)(void *state, unsigned bir, uint64_t offset, void *buf, size_t len);
    ssize_t (*write)(void *state, unsigned bir, uint64_t offset, const void *buf, size_t len);
    void (*destroy)(void *state);
};

int behavior_create(const struct device_model *model,
                    struct device_behavior *out,
                    char *err, size_t err_len);
```

- [ ] **Step 1: Write RED tests for bounds, byte enables, and reset**

Test that an in-range 32-bit write is readable, an out-of-range access returns `-EINVAL`, a two-byte write changes only those bytes, and reset restores generated BAR reset bytes.

- [ ] **Step 2: Run RED**

Run: `make -C vfio-user test-behavior-static`

Expected: compile failure for missing behavior API.

- [ ] **Step 3: Implement the static module and class factory table**

Keep the dispatch table data-driven by full 24-bit class code. Return static behavior for class `0x000000` and multibar models; reserve factory entries for NVMe `0x010802`, AHCI `0x010601`, and xHCI `0x0c0330`.

- [ ] **Step 4: Run GREEN**

Run: `make -C vfio-user test-behavior-static`

Expected: all static behavior tests pass under AddressSanitizer and UndefinedBehaviorSanitizer.

- [ ] **Step 5: Commit**

```bash
git add vfio-user/include/device_behavior.h vfio-user/src/behavior_static.c vfio-user/tests/test_behavior_static.c vfio-user/Makefile
git commit -m "feat: add bounded vfio device behaviors"
```

### Task 4: Build the generic libvfio-user server

**Files:**
- Create: `vfio-user/include/vfio_device.h`
- Create: `vfio-user/src/vfio_device.c`
- Create: `vfio-user/src/main.c`
- Create: `vfio-user/tests/test_server_cli.py`
- Modify: `vfio-user/Makefile`

**Interfaces:**
- Produces: `int vfio_device_run(const struct device_model *model, struct device_behavior *behavior, const char *socket_path, volatile sig_atomic_t *stop)`.
- CLI: `vfio-device --artifacts DIR --socket PATH`; readiness is a single JSON line on stdout.

- [ ] **Step 1: Write RED CLI tests**

Test missing artifact directory, malformed model, stale socket replacement, SIGTERM cleanup, and readiness JSON containing the generated VID/DID/class/BAR count.

- [ ] **Step 2: Run RED**

Run: `python3 -m unittest vfio-user/tests/test_server_cli.py -v`

Expected: failure because `vfio-user/build/vfio-device` does not exist.

- [ ] **Step 3: Implement generic setup**

Initialize PCI Express config space from the model; register each implemented BIR with exact size and memory/I/O flags; register DMA callbacks and generated MSI/MSI-X counts; delegate BAR callbacks to `device_behavior`; stop cleanly on disconnect, SIGINT, or SIGTERM. Check every libvfio-user return code and preserve the first error.

- [ ] **Step 4: Run GREEN**

Run: `make -C vfio-user test-server`

Expected: CLI tests pass with no leaked socket and no sanitizer findings.

- [ ] **Step 5: Commit**

```bash
git add vfio-user/include/vfio_device.h vfio-user/src/vfio_device.c vfio-user/src/main.c vfio-user/tests/test_server_cli.py vfio-user/Makefile
git commit -m "feat: add generated-model vfio-user server"
```

### Task 5: Extract and correct NVMe behavior

**Files:**
- Create: `vfio-user/src/behavior_nvme.c`
- Create: `vfio-user/tests/test_behavior_nvme.c`
- Delete after parity is proven: `vfio-user/nvme_vfio_server.c`
- Modify: `vfio-user/Makefile`

**Interfaces:**
- Consumes: generated CAP/VS/reset BAR values, MSI/MSI-X metadata, and the generic DMA/IRQ callbacks.
- Produces: `int behavior_nvme_create(const struct device_model *, const struct behavior_host_ops *, struct device_behavior *, char *, size_t)`.

- [ ] **Step 1: Write RED driver-sequence tests**

Cover `CC.EN -> CSTS.RDY`, `CC.SHN -> CSTS.SHST`, DSTRD-aware SQ/CQ doorbells, admin queue bounds, Identify Controller/Namespace, required log pages/features, CQ-full refusal, DMA error completion, reset, MSI mask state, and MSI-X pending delivery on unmask.

- [ ] **Step 2: Run RED**

Run: `make -C vfio-user test-behavior-nvme`

Expected: tests fail because the behavior factory has no NVMe module.

- [ ] **Step 3: Extract minimal proven behavior**

Move logic from the prototype only when a RED test requires it. Read identity and register reset values from generated artifacts. Use checked DMA host operations, queue sizes derived from AQA, and generated MSI-X table/PBA locations. Unknown admin opcodes return Invalid Opcode; they never return success by default.

- [ ] **Step 4: Run GREEN and parity check**

Run: `make -C vfio-user test-behavior-nvme test-server`

Expected: all NVMe and generic server tests pass; sanitizers are clean.

- [ ] **Step 5: Delete the obsolete server and remove narration comments**

Run: `rg -n '^/\* ----|NVMe VFIO-user Device Server|our generated firmware|This runs:' vfio-user`

Expected: no matches in touched production files.

- [ ] **Step 6: Commit**

```bash
git add -A vfio-user
git commit -m "feat: extract nvme vfio behavior"
```

### Task 6: Add AHCI, xHCI, and profiled-class behaviors

**Files:**
- Create: `vfio-user/src/behavior_ahci.c`
- Create: `vfio-user/src/behavior_xhci.c`
- Create: `vfio-user/src/behavior_profiled.c`
- Create: `vfio-user/tests/test_behavior_ahci.c`
- Create: `vfio-user/tests/test_behavior_xhci.c`
- Create: `vfio-user/tests/test_behavior_profiled.c`
- Modify: `vfio-user/src/behavior_static.c`
- Modify: `vfio-user/Makefile`

**Interfaces:**
- Produces one factory per module with the same model/host-ops/behavior/error signature as NVMe.
- `behavior_profiled` consumes generated `bar_behavior_profile.json` and model reset data; it cannot invent unprofiled writable registers.

- [ ] **Step 1: Write AHCI RED tests**

Replay the Linux `ahci` sequence: CAP/GHC/PI/VS reads, HBA reset self-clear, port command start/stop, signature/status reads, interrupt status write-one-to-clear, and reset restoration.

- [ ] **Step 2: Implement AHCI and run GREEN**

Run: `make -C vfio-user test-behavior-ahci`

Expected: AHCI tests pass under sanitizers.

- [ ] **Step 3: Write xHCI RED tests**

Replay capability discovery, `USBCMD.HCRST`, `USBSTS.CNR/HCH`, page-size read, command-ring and event-ring programming, run/stop, and reset restoration.

- [ ] **Step 4: Implement xHCI and run GREEN**

Run: `make -C vfio-user test-behavior-xhci`

Expected: xHCI tests pass under sanitizers.

- [ ] **Step 5: Write profiled-device RED tests**

For audio `0x040300`, Ethernet `0x020000`, Wi-Fi `0x028000`, GPU `0x030000`, and Thunderbolt `0x080700`, verify static, writable, W1C, dead, and all-ones register semantics from the generated profile plus reset behavior. Reject a profile whose class or BAR size differs from `device_model.json`.

- [ ] **Step 6: Implement the bounded profile interpreter and run GREEN**

Run: `make -C vfio-user test-behavior-profiled test`

Expected: all C unit tests pass; no class-specific constants appear outside behavior factory tests.

- [ ] **Step 7: Commit**

```bash
git add vfio-user/src/behavior_ahci.c vfio-user/src/behavior_xhci.c vfio-user/src/behavior_profiled.c vfio-user/src/behavior_static.c vfio-user/tests vfio-user/Makefile
git commit -m "feat: add class-specific vfio behaviors"
```

### Task 7: Build the deterministic guest probe

**Files:**
- Create: `vfio-user/guest/init`
- Create: `vfio-user/guest/probe.py`
- Create: `vfio-user/tests/test_result_parser.py`
- Modify: `vfio-user/Dockerfile`

**Interfaces:**
- Guest emits records with `event`, `case`, `status`, `bdf`, `vendor`, `device`, `class`, `bars`, `driver`, and `detail`.
- Host parser: `parse_guest_results(text: str, case: Case) -> GuestResult`; malformed, duplicate terminal, or missing terminal records fail.

- [ ] **Step 1: Write RED parser tests**

Cover a passing generic record, wrong VID/DID, wrong BAR size, missing terminal record, explicit dependency skip, and a `dmesg` fatal signature after driver bind.

- [ ] **Step 2: Run RED**

Run: `python3 -m unittest vfio-user/tests/test_result_parser.py -v`

Expected: import failure for missing parser.

- [ ] **Step 3: Implement probe and parser**

The probe locates the VFIO-user device by generated VID/DID/class, reads `/sys/bus/pci/devices`, records BAR resources and driver symlink, and runs fixture-specific checks: NVMe controller state/namespace, AHCI host/port, xHCI host registration, Ethernet interface, and bind status for remaining classes. Always emit a terminal record before poweroff.

- [ ] **Step 4: Build and inspect initramfs**

Run: `make -C vfio-user guest && gzip -dc vfio-user/build/initramfs.cpio.gz | cpio -it | sort`

Expected: archive contains `/init`, probe runtime, required modules, and no build-host absolute paths.

- [ ] **Step 5: Run GREEN**

Run: `python3 -m unittest vfio-user/tests/test_result_parser.py -v`

Expected: parser tests pass.

- [ ] **Step 6: Commit**

```bash
git add vfio-user/guest vfio-user/tests/test_result_parser.py vfio-user/Dockerfile
git commit -m "test: add structured vfio guest probe"
```

### Task 8: Connect generation, Cocotb, VFIO-user, and QEMU

**Files:**
- Modify: `vfio-user/matrix.py`
- Modify: `vfio-user/tests/test_matrix.py`
- Create: `vfio-user/tests/test_qemu_smoke.py`
- Modify: `vfio-user/Makefile`

**Interfaces:**
- CLI: `python3 vfio-user/matrix.py --case nvme --work-dir PATH` and `python3 vfio-user/matrix.py --all --work-dir PATH`.
- Per-case outputs: `artifacts/`, `generation.log`, `cocotb.log`, `server.log`, `qemu.log`, `guest-results.jsonl`, and `result.json`.
- Matrix output: `summary.json`; process exit is 0 only when every mandatory stage passes.

- [ ] **Step 1: Write RED orchestration tests with executable fakes**

Verify exact build arguments include `--skip-vivado`, board selection, socket uniqueness, timeout cleanup, stage ordering, log retention, dependency skip handling, and nonzero matrix exit after one injected failing case.

- [ ] **Step 2: Run RED**

Run: `python3 -m unittest vfio-user/tests/test_matrix.py -v`

Expected: new orchestration assertions fail.

- [ ] **Step 3: Implement the pipeline**

Build `./bin/pcileechgen` once, generate each fixture into its case directory, verify the manifest, invoke the existing device-appropriate Cocotb target, start `vfio-device`, wait for readiness JSON and socket creation, run QEMU with a hard timeout, parse guest output, and write result files atomically.

- [ ] **Step 4: Run GREEN with fakes**

Run: `python3 -m unittest discover -s vfio-user/tests -p 'test_*.py' -v`

Expected: all Python tests pass.

- [ ] **Step 5: Run the first real generic smoke test**

Run: `make -C vfio-user e2e CASE=generic`

Expected: `generic` passes generation, Cocotb generic enumeration, QEMU enumeration, identity, capability, and BAR checks.

- [ ] **Step 6: Commit**

```bash
git add vfio-user/matrix.py vfio-user/tests vfio-user/Makefile
git commit -m "test: orchestrate vfio firmware e2e pipeline"
```

### Task 9: Verify every device-specific E2E contract

**Files:**
- Test: `vfio-user/tests/test_behavior_nvme.c`
- Test: `vfio-user/tests/test_behavior_ahci.c`
- Test: `vfio-user/tests/test_behavior_xhci.c`
- Test: `vfio-user/tests/test_behavior_profiled.c`
- Test: `vfio-user/tests/test_qemu_smoke.py`

**Interfaces:**
- Consumes the matrix pipeline from Task 8.
- Produces passing mandatory generic checks for all ten cases and passing class smoke checks where the generated firmware implements the required behavior.

- [ ] **Step 1: Run each case independently and capture RED evidence**

Run: `for case in nvme sata xhci audio ethernet wifi gpu thunderbolt multibar generic; do make -C vfio-user e2e CASE="$case" || break; done`

Expected: every case reaches a terminal pass or an explicitly permitted external-dependency skip after its mandatory generic checks pass.

- [ ] **Step 2: Run all class behavior tests together**

Run: `make -C vfio-user test-behavior-nvme test-behavior-ahci test-behavior-xhci test-behavior-profiled`

Expected: all driver-sequence, reset, bounds, interrupt, and error-path tests pass under sanitizers.

- [ ] **Step 3: Run the complete matrix as one process**

Run: `make -C vfio-user e2e-all`

Expected: `summary.json` contains ten unique terminal results and the command exits zero.

- [ ] **Step 4: Commit the matrix acceptance test**

```bash
git add vfio-user/tests/test_qemu_smoke.py
git commit -m "test: require all vfio device contracts"
```

### Task 10: Documentation, cleanup, and full verification

**Files:**
- Create: `vfio-user/README.md`
- Modify: `vfio-user/Dockerfile`
- Modify: `vfio-user/Makefile`
- Delete: `vfio-user/run.sh` after its supported commands exist in `Makefile` and `matrix.py`.

**Interfaces:**
- `make -C vfio-user test`: all local unit tests.
- `make -C vfio-user image`: reproducible Docker image.
- `make -C vfio-user e2e CASE=nvme`: one case.
- `make -C vfio-user e2e-all`: all ten cases.

- [ ] **Step 1: Document exact prerequisites, boundaries, commands, outputs, skips, and troubleshooting**

State explicitly that VFIO-user does not execute FPGA RTL and that Cocotb is the HDL behavior stage. Do not mention Vivado as a runtime option.

- [ ] **Step 2: Run comment and failure-swallow audits**

Run: `rg -n '^/\* ----|This runs:|our generated firmware|AI|TO[D]O|FIX[M]E|\|\| true' vfio-user`

Expected: no production-code narration, placeholder comments, or unconditional failure swallowing. Legitimate test strings must be narrowly excluded rather than weakening the audit.

- [ ] **Step 3: Run formatters and static checks**

Run: `make -C vfio-user format-check lint`

Expected: `clang-format --dry-run --Werror`, compiler `-Wall -Wextra -Werror`, sanitizers, and Python compile checks pass.

- [ ] **Step 4: Run repository regressions**

Run: `go test ./... && make -C vfio-user test && make -C tests/cocotb`

Expected: all Go, VFIO-user unit, and default Cocotb tests pass.

- [ ] **Step 5: Run the complete matrix**

Run: `make -C vfio-user e2e-all`

Expected: all ten fixtures pass mandatory artifact, HDL, PCI identity, capability, and BAR checks; class probes pass or report only spec-permitted external-dependency skips; `summary.json` contains ten terminal results.

- [ ] **Step 6: Prove failure propagation**

Run: `python3 vfio-user/matrix.py --all --work-dir vfio-user/build/failure-check --inject-corrupt-model generic`

Expected: nonzero exit, `generic/result.json` reports artifact validation failure, remaining completed case logs are retained, and the summary cannot report success.

- [ ] **Step 7: Commit**

```bash
git add -A vfio-user
git commit -m "docs: finalize vfio firmware e2e suite"
```
