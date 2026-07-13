#!/usr/bin/env python3
from __future__ import annotations

import os
import json
import argparse
import platform
import re
import select
import signal
import stat
import subprocess
import tempfile
import time
from dataclasses import dataclass
from pathlib import Path
from typing import Sequence


ROOT = Path(__file__).resolve().parents[1]
GENERATOR = Path(os.environ.get("PCILEECHGEN_BIN", ROOT / "bin" / "pcileechgen"))
REQUIRED_GUEST_STAGES = ("enumerate", "bars", "driver")
ALLOWED_BLOCK_REASONS = frozenset(
    {
        "guest-optional-driver",
        "kvm-unavailable",
        "qemu-version-mismatch",
    }
)
ALLOWED_GUEST_SKIPS = {
    "audio": frozenset({"optional-driver-not-bound"}),
    "gpu": frozenset({"optional-driver-not-bound"}),
    "thunderbolt": frozenset({"optional-driver-not-bound"}),
    "wifi": frozenset({"optional-driver-not-bound"}),
}
FATAL_GUEST_SIGNATURES = (
    "bug:",
    "general protection fault",
    "kernel panic",
    "oops:",
    "unable to handle kernel",
)
PCI_BDF = re.compile(
    r"^(?P<domain>[0-9a-fA-F]{4}):(?P<bus>[0-9a-fA-F]{2}):"
    r"(?P<device>[0-9a-fA-F]{2})\.(?P<function>[0-7])$"
)
UNIX_SOCKET_PATH_LIMIT = 103


class CaseFailure(RuntimeError):
    pass


class CaseBlocked(RuntimeError):
    def __init__(self, reason: str, detail: str):
        if reason not in ALLOWED_BLOCK_REASONS:
            raise ValueError(f"blocked reason is not allowlisted: {reason}")
        super().__init__(detail)
        self.reason = reason


@dataclass(frozen=True)
class SocketLease:
    directory: Path
    path: Path


@dataclass(frozen=True)
class Case:
    name: str
    fixture: Path
    board: str
    behavior: str
    mandatory_probe: str


@dataclass(frozen=True)
class GuestResult:
    case: str
    status: str
    bdf: str
    vendor: str
    device: str
    class_code: str
    driver: str
    bars: int
    detail: str


def _normalize_hex(value: object, width: int, field: str) -> str:
    if not isinstance(value, str):
        raise CaseFailure(f"guest {field} is not a hexadecimal string")
    normalized = value.removeprefix("0x").removeprefix("0X")
    if len(normalized) != width or any(char not in "0123456789abcdefABCDEF" for char in normalized):
        raise CaseFailure(f"guest {field} is not {width} hexadecimal digits")
    return normalized.lower()


def _validate_bdf(value: object) -> str:
    if not isinstance(value, str):
        raise CaseFailure("invalid guest BDF")
    match = PCI_BDF.fullmatch(value)
    if match is None or int(match.group("device"), 16) > 0x1F:
        raise CaseFailure(f"invalid guest BDF: {value}")
    return value.lower()


def _fixture_contract(case: Case) -> dict:
    try:
        fixture = json.loads(case.fixture.read_text(encoding="utf-8"))
        device = fixture["device"]
        bars = fixture["bars"]
        return {
            "identity": {
                "vendor": f"{device['vendor_id']:04x}",
                "device": f"{device['device_id']:04x}",
                "class": f"{device['class_code']:06x}",
            },
            "bars": [
                {
                    "bir": bar["index"],
                    "size": bar["size"],
                    "type": bar["type"],
                    "prefetchable": bar["prefetchable"],
                    "address_width": 64 if bar["is_64bit"] else 32,
                }
                for bar in bars
            ],
        }
    except (KeyError, TypeError, ValueError, OSError, json.JSONDecodeError) as exc:
        raise CaseFailure(f"{case.name}: invalid fixture contract: {exc}") from exc


def _validated_expected_bars(contract: dict, case: Case) -> list[dict]:
    try:
        bars = contract["bars"]
        if not isinstance(bars, list) or len(bars) > 6:
            raise ValueError("BAR collection must contain at most six entries")
        seen = set()
        for bar in bars:
            bir = bar["bir"]
            if isinstance(bir, bool) or not isinstance(bir, int) or not 0 <= bir <= 5:
                raise ValueError(f"invalid BAR index {bir!r}")
            if bir in seen:
                raise ValueError(f"duplicate BAR index {bir}")
            seen.add(bir)
            if isinstance(bar["size"], bool) or not isinstance(bar["size"], int) or bar["size"] <= 0:
                raise ValueError(f"invalid BAR{bir} size")
            if bar["type"] not in {"io", "mem32", "mem64"}:
                raise ValueError(f"invalid BAR{bir} type")
            if not isinstance(bar["prefetchable"], bool):
                raise ValueError(f"invalid BAR{bir} prefetchability")
            if bar["address_width"] not in {32, 64}:
                raise ValueError(f"invalid BAR{bir} address width")
        return bars
    except (KeyError, TypeError, ValueError) as exc:
        raise CaseFailure(f"{case.name}: invalid expected BAR contract: {exc}") from exc


def _validate_guest_bar_evidence(value: object, expected: list[dict]) -> int:
    if isinstance(value, bool):
        raise CaseFailure("guest BAR evidence has invalid type")
    if isinstance(value, int):
        return value
    if not isinstance(value, list):
        raise CaseFailure("guest BAR evidence must be a count or indexed records")

    expected_by_bir = {bar["bir"]: bar for bar in expected}
    observed_birs = set()
    for observed in value:
        if not isinstance(observed, dict):
            raise CaseFailure("guest BAR evidence contains a non-record")
        try:
            bir = observed["bir"]
            expected_bar = expected_by_bir[bir]
        except (KeyError, TypeError) as exc:
            raise CaseFailure("guest BAR evidence contains an unexpected BIR") from exc
        if bir in observed_birs:
            raise CaseFailure(f"guest BAR evidence duplicates BAR{bir}")
        observed_birs.add(bir)
        for field in ("size", "type", "prefetchable", "address_width"):
            if field not in observed:
                raise CaseFailure(f"guest BAR{bir} evidence is missing {field}")
            if observed[field] != expected_bar[field]:
                raise CaseFailure(f"guest BAR{bir} {field} mismatch")
    if observed_birs != set(expected_by_bir):
        raise CaseFailure("guest BAR evidence does not cover every generated BIR")
    return len(value)


def parse_guest_results(text: str, case: Case, contract: dict | None = None) -> GuestResult:
    records = []
    raw_lines = text.splitlines()
    for line_number, line in enumerate(raw_lines):
        line = line.strip()
        if not line or not line.startswith("{"):
            continue
        try:
            record = json.loads(line)
        except json.JSONDecodeError as exc:
            raise CaseFailure(f"invalid guest result JSON: {exc}") from exc
        if not isinstance(record, dict):
            raise CaseFailure("guest record is not a JSON object")
        if record.get("event") in {"stage", "result"}:
            if record.get("case") != case.name:
                raise CaseFailure(f"guest result case mismatch: {record.get('case')}")
            records.append((line_number, record))

    terminal = [(line_number, record) for line_number, record in records
                if record.get("event") == "result"]
    if len(terminal) != 1:
        raise CaseFailure(f"expected one terminal guest result, got {len(terminal)}")
    terminal_line, record = terminal[0]
    if record.get("case") != case.name:
        raise CaseFailure(f"guest result case mismatch: {record.get('case')}")
    if record.get("status") not in {"pass", "skip", "fail"}:
        raise CaseFailure("guest result has invalid status")
    required = ("bdf", "vendor", "device", "class", "driver", "bars", "detail")
    missing = [key for key in required if key not in record]
    if missing:
        raise CaseFailure(f"guest result missing fields: {','.join(missing)}")

    stage_records = {}
    for stage in REQUIRED_GUEST_STAGES:
        matches = [(line_number, item) for line_number, item in records
                   if item.get("event") == "stage" and item.get("stage") == stage]
        if len(matches) != 1:
            raise CaseFailure(
                f"missing or duplicate guest stage evidence for {stage}: {len(matches)}"
            )
        stage_records[stage] = matches[0]
    stage_lines = [stage_records[stage][0] for stage in REQUIRED_GUEST_STAGES]
    if stage_lines != sorted(stage_lines) or stage_lines[-1] >= terminal_line:
        raise CaseFailure("guest stage evidence is out of order")

    expected = contract if contract is not None else _fixture_contract(case)
    try:
        identity = expected["identity"]
    except (KeyError, TypeError) as exc:
        raise CaseFailure(f"{case.name}: invalid expected identity contract") from exc
    observed_identity = {
        "vendor": _normalize_hex(record["vendor"], 4, "vendor"),
        "device": _normalize_hex(record["device"], 4, "device"),
        "class": _normalize_hex(record["class"], 6, "class"),
    }
    for field, observed in observed_identity.items():
        if observed != str(identity.get(field, "")).lower():
            raise CaseFailure(
                f"guest {field} mismatch: got {observed}, expected {identity.get(field)}"
            )

    bdf = _validate_bdf(record["bdf"])
    enumerate_stage = stage_records["enumerate"][1]
    if _validate_bdf(enumerate_stage.get("bdf")) != bdf:
        raise CaseFailure("guest BDF does not match enumerate stage")

    expected_bars = _validated_expected_bars(expected, case)
    bar_count = _validate_guest_bar_evidence(record["bars"], expected_bars)
    stage_count = stage_records["bars"][1].get("count")
    if isinstance(stage_count, bool) or not isinstance(stage_count, int):
        raise CaseFailure("guest BAR stage count is invalid")
    if stage_count != bar_count:
        raise CaseFailure("guest BAR terminal count does not match BAR stage")
    if bar_count != len(expected_bars):
        raise CaseFailure(
            f"guest BAR count mismatch: got {bar_count}, expected {len(expected_bars)}"
        )

    driver = record["driver"]
    if not isinstance(driver, str) or not driver:
        raise CaseFailure("guest driver evidence is invalid")
    driver_stage = stage_records["driver"][1]
    if driver_stage.get("driver") != driver:
        raise CaseFailure("guest driver does not match driver stage")
    if driver_stage.get("status") != record["status"]:
        raise CaseFailure("guest status does not match driver stage")

    detail = record["detail"]
    if not isinstance(detail, str) or not detail:
        raise CaseFailure("guest result detail is missing")
    if record["status"] == "skip" and detail not in ALLOWED_GUEST_SKIPS.get(case.name, ()):
        raise CaseFailure(f"guest skip reason is not permitted for {case.name}: {detail}")
    if record["status"] == "pass" and case.mandatory_probe not in {"enumeration", "bar-layout"}:
        if driver == "none":
            raise CaseFailure(f"{case.name}: passing guest result has no bound driver")

    driver_line = stage_records["driver"][0]
    for line_number in range(driver_line + 1, len(raw_lines)):
        line = raw_lines[line_number].strip()
        if line.startswith("{"):
            continue
        lowered = line.lower()
        if any(signature in lowered for signature in FATAL_GUEST_SIGNATURES):
            raise CaseFailure(f"fatal guest kernel message after driver bind: {line}")

    return GuestResult(
        case=case.name,
        status=record["status"],
        bdf=bdf,
        vendor=observed_identity["vendor"],
        device=observed_identity["device"],
        class_code=observed_identity["class"],
        driver=driver,
        bars=bar_count,
        detail=detail,
    )


def _case(name: str, behavior: str, mandatory_probe: str, board: str = "PCIeSquirrel") -> Case:
    return Case(
        name=name,
        fixture=ROOT / "testdata" / "donors" / f"{name}.json",
        board=board,
        behavior=behavior,
        mandatory_probe=mandatory_probe,
    )


def build_command(case: Case, output: Path) -> list[str]:
    return [
        str(GENERATOR),
        "build",
        "--from-json",
        str(case.fixture),
        "--board",
        case.board,
        "--skip-vivado",
        "--output",
        str(output),
        "--force",
    ]


def build_contract(case: Case, artifacts: Path) -> dict:
    try:
        model = json.loads((artifacts / "device_model.json").read_text(encoding="utf-8"))
        functions = model["functions"]
        if not isinstance(functions, list) or len(functions) != 1:
            raise ValueError("device model must contain exactly one function")
        function = functions[0]
        contract = {
            "case": case.name,
            "identity": {
                "vendor": f"{function['vendor_id']:04x}",
                "device": f"{function['device_id']:04x}",
                "class": f"{function['class_code']:06x}",
            },
            "bars": [
                {
                    "bir": bar["bir"],
                    "size": bar["size"],
                    "type": bar["type"],
                    "prefetchable": bar["prefetchable"],
                    "address_width": bar["address_width"],
                }
                for bar in model["bars"]
            ],
            "capabilities": [cap["id"] for cap in model["capabilities"]],
            "reset": {"vfio_callback": True, "bar_reset_image": True},
            "probe": ["enumerate", "bars", "reset", case.mandatory_probe],
        }
        for field, width in (("vendor", 4), ("device", 4), ("class", 6)):
            _normalize_hex(contract["identity"][field], width, field)
        _validated_expected_bars(contract, case)
        return contract
    except CaseFailure:
        raise
    except (KeyError, TypeError, ValueError, OSError, json.JSONDecodeError) as exc:
        raise CaseFailure(f"{case.name}: invalid generated device contract: {exc}") from exc


def allocate_socket_lease(case: Case, work_dir: Path) -> SocketLease:
    work_dir.mkdir(parents=True, exist_ok=True)
    directory = Path(
        tempfile.mkdtemp(prefix=".v-", dir=work_dir.resolve())
    )
    try:
        directory.chmod(0o700)
        lease = SocketLease(directory=directory, path=directory / "s")
        if len(os.fsencode(lease.path)) > UNIX_SOCKET_PATH_LIMIT:
            raise CaseFailure(
                f"{case.name}: VFIO Unix socket path exceeds "
                f"{UNIX_SOCKET_PATH_LIMIT} bytes: {lease.path}"
            )
        return lease
    except BaseException:
        directory.rmdir()
        raise


def wait_for_unix_socket(lease: SocketLease, process: subprocess.Popen,
                         timeout: float) -> None:
    if timeout <= 0:
        raise CaseFailure("VFIO server did not publish a Unix socket before readiness timeout")
    deadline = time.monotonic() + timeout
    while True:
        returncode = process.poll()
        if returncode is not None:
            raise CaseFailure(
                f"VFIO server exited with status {returncode} before publishing its Unix socket"
            )
        try:
            metadata = lease.path.lstat()
        except FileNotFoundError:
            remaining = deadline - time.monotonic()
            if remaining <= 0:
                raise CaseFailure(
                    "VFIO server did not publish a Unix socket before readiness timeout"
                )
            time.sleep(min(0.01, remaining))
            continue
        # A connect probe can consume libvfio-user's QEMU client slot. The private
        # lease directory plus lstat proves the published node without touching it.
        if not stat.S_ISSOCK(metadata.st_mode):
            raise CaseFailure(f"VFIO readiness path is not a Unix socket: {lease.path}")
        if process.poll() is not None:
            raise CaseFailure("VFIO server exited while publishing its Unix socket")
        return


def _path_is_present(path: Path) -> bool:
    try:
        path.lstat()
    except FileNotFoundError:
        return False
    return True


def _remove_socket_lease_directory(lease: SocketLease, errors: list[str]) -> None:
    try:
        lease.directory.rmdir()
    except FileNotFoundError:
        return
    except OSError as exc:
        errors.append(f"failed to remove private socket directory {lease.directory}: {exc}")


def start_server(case: Case, artifacts: Path, work_dir: Path, timeout: int = 10):
    binary = ROOT / "vfio-user" / "build" / "vfio-device"
    if not binary.is_file():
        raise CaseFailure(f"VFIO server binary is missing: {binary}")
    if timeout <= 0:
        raise CaseFailure("VFIO server readiness timeout must be positive")
    log_path = work_dir / "server.log"
    work_dir.mkdir(parents=True, exist_ok=True)
    lease = allocate_socket_lease(case, work_dir)
    try:
        log = log_path.open("w", encoding="utf-8")
    except BaseException:
        lease.directory.rmdir()
        raise
    with log:
        try:
            process = subprocess.Popen(
                [
                    str(binary),
                    "--artifacts",
                    str(artifacts),
                    "--socket",
                    str(lease.path),
                ],
                cwd=ROOT / "vfio-user",
                stdout=subprocess.PIPE,
                stderr=log,
                text=True,
                start_new_session=True,
            )
        except BaseException:
            lease.directory.rmdir()
            raise
        try:
            deadline = time.monotonic() + timeout
            ready, _, _ = select.select([process.stdout], [], [], timeout)
            if not ready:
                raise CaseFailure(f"{case.name}: VFIO server did not become ready")
            line = process.stdout.readline()
            try:
                record = json.loads(line)
            except json.JSONDecodeError as exc:
                raise CaseFailure(f"{case.name}: malformed readiness record") from exc
            if not isinstance(record, dict) or record.get("event") != "ready":
                raise CaseFailure(f"{case.name}: invalid readiness record")
            wait_for_unix_socket(lease, process, deadline - time.monotonic())
            return process, process.stdout, lease, record
        except BaseException as exc:
            try:
                stop_server(process, process.stdout, lease)
            except CaseFailure as cleanup_exc:
                raise CaseFailure(
                    f"{case.name}: server startup failed ({exc}); cleanup failed: {cleanup_exc}"
                ) from exc
            raise


def stop_server(process: subprocess.Popen, output, lease: SocketLease, *,
                process_timeout: float = 5, socket_timeout: float = 1) -> None:
    errors = []
    try:
        if process.poll() is None:
            process.terminate()
            try:
                process.wait(timeout=process_timeout)
            except subprocess.TimeoutExpired:
                process.kill()
                process.wait(timeout=process_timeout)
    except (OSError, subprocess.SubprocessError) as exc:
        errors.append(f"failed to stop VFIO server process: {exc}")

    try:
        output.close()
    except OSError as exc:
        errors.append(f"failed to close VFIO server output: {exc}")

    if process.poll() is None:
        errors.append("VFIO server process is still running; socket lease was retained")
    else:
        deadline = time.monotonic() + max(0, socket_timeout)
        while _path_is_present(lease.path) and time.monotonic() < deadline:
            time.sleep(min(0.01, max(0, deadline - time.monotonic())))
        if _path_is_present(lease.path):
            errors.append(f"VFIO server did not remove Unix socket: {lease.path}")
            try:
                lease.path.unlink()
            except OSError as exc:
                errors.append(f"failed to remove abandoned Unix socket {lease.path}: {exc}")
        _remove_socket_lease_directory(lease, errors)

    if errors:
        raise CaseFailure("; ".join(errors))


def _normalized_architecture(value: str) -> str:
    return {
        "amd64": "x86_64",
        "arm64": "aarch64",
    }.get(value.lower(), value.lower())


def _qemu_architecture(qemu: Path) -> str | None:
    name = qemu.name.lower()
    for architecture in ("x86_64", "aarch64"):
        if name.endswith(architecture):
            return architecture
    return None


def qemu_requires_kvm(case: Case, kvm_path: Path = Path("/dev/kvm"), *,
                      qemu: Path | None = None, host_machine: str | None = None,
                      host_system: str | None = None) -> bool:
    if case.name != "nvme":
        return False
    system = platform.system() if host_system is None else host_system
    if system != "Linux":
        return False
    host_arch = _normalized_architecture(
        platform.machine() if host_machine is None else host_machine
    )
    guest_arch = _qemu_architecture(qemu) if qemu is not None else host_arch
    if guest_arch is not None and guest_arch != host_arch:
        return False
    return not kvm_path.exists()


def run_server_smoke(case: Case, artifacts: Path, work_dir: Path, timeout: int = 10) -> dict:
    process, output, lease, record = start_server(case, artifacts, work_dir, timeout)
    stop_server(process, output, lease)
    return record


def run_qemu_case(case: Case, artifacts: Path, work_dir: Path,
                  qemu: Path, kernel: Path, initrd: Path, timeout: int = 30,
                  shared_memory: bool = False, rebind: bool = False,
                  expected_version: str | None = None) -> GuestResult:
    if qemu_requires_kvm(case, qemu=qemu):
        raise CaseBlocked(
            "kvm-unavailable",
            f"{case.name}: /dev/kvm is required for native QEMU MSI-X E2E",
        )
    if expected_version is not None:
        version = subprocess.run([str(qemu), "--version"], capture_output=True,
                                 text=True, check=False)
        if version.returncode != 0:
            raise CaseFailure(
                f"{case.name}: QEMU version probe exited {version.returncode}"
            )
        if expected_version not in version.stdout:
            raise CaseBlocked(
                "qemu-version-mismatch",
                f"{case.name}: QEMU version mismatch; expected {expected_version!r}",
            )
    contract = build_contract(case, artifacts)
    process, output, lease, readiness = start_server(case, artifacts, work_dir)
    try:
        validate_readiness(readiness, contract, case)
        qemu_log = work_dir / "qemu.log"
        command = build_qemu_command(case, lease.path, kernel, initrd, artifacts,
                                     qemu=qemu, shared_memory=shared_memory,
                                     rebind=rebind)
        try:
            with qemu_log.open("w", encoding="utf-8") as log:
                result = subprocess.run(command, stdout=log, stderr=subprocess.STDOUT,
                                        timeout=timeout, check=False)
        except subprocess.TimeoutExpired as exc:
            _retain_guest_records(qemu_log, work_dir / "guest-results.jsonl")
            raise CaseFailure(f"{case.name}: QEMU timed out after {timeout}s") from exc
        text = qemu_log.read_text(encoding="utf-8")
        _retain_guest_records(qemu_log, work_dir / "guest-results.jsonl")
        if result.returncode != 0:
            raise CaseFailure(f"{case.name}: QEMU exited {result.returncode}")
        guest = parse_guest_results(text, case, contract)
        validate_guest_result(guest, case, rebind=rebind)
        return guest
    finally:
        stop_server(process, output, lease)


def validate_guest_result(guest: GuestResult, case: Case, *, rebind: bool = False) -> None:
    if guest.status == "fail":
        raise CaseFailure(f"{case.name}: guest probe failed: {guest.detail}")
    if rebind:
        if guest.status != "pass" or guest.driver == "none" or (
            guest.detail != "rebind-reset-driver-bound"
        ):
            raise CaseFailure(f"{case.name}: guest result lacks reset/rebind evidence")
    if guest.status == "skip":
        raise CaseBlocked(
            "guest-optional-driver",
            f"{case.name}: optional guest dependency unavailable: {guest.detail}",
        )


def build_qemu_command(case: Case, socket_path: Path, kernel: Path,
                       initrd: Path, artifacts: Path, qemu: Path = Path("qemu"),
                       shared_memory: bool = False, rebind: bool = False) -> list[str]:
    machine_args = ["-machine", "virt", "-cpu", "max", "-m", "512"]
    console = "ttyAMA0"
    if shared_memory:
        machine_args = [
            "-machine", "q35,memory-backend=mem",
            "-object", "memory-backend-memfd,id=mem,size=512M,share=on",
            "-cpu", "max", "-m", "512",
        ]
        console = "ttyS0"
    append = f"console={console} rdinit=/init vfio_case={case.name} "
    append += f"vfio_vendor={vendor_for(artifacts)} vfio_device={device_for(artifacts)}"
    if rebind:
        append += " vfio_rebind=1"
    return [
        str(qemu), *machine_args, "-nographic",
        "-kernel", str(kernel), "-initrd", str(initrd),
        "-append", append,
        "-device", json.dumps({"driver": "vfio-user-pci",
                               "socket": {"path": str(socket_path), "type": "unix"}}),
    ]


def vendor_for(artifacts: Path) -> str:
    data = json.loads((artifacts / "device_model.json").read_text(encoding="utf-8"))
    return f"{data['functions'][0]['vendor_id']:04x}"


def device_for(artifacts: Path) -> str:
    data = json.loads((artifacts / "device_model.json").read_text(encoding="utf-8"))
    return f"{data['functions'][0]['device_id']:04x}"


def _atomic_write(path: Path, text: str) -> None:
    path.parent.mkdir(parents=True, exist_ok=True)
    descriptor, temporary_name = tempfile.mkstemp(
        dir=path.parent,
        prefix=f".{path.name}.",
        suffix=".tmp",
    )
    temporary = Path(temporary_name)
    try:
        with os.fdopen(descriptor, "w", encoding="utf-8") as output:
            output.write(text)
            output.flush()
            os.fsync(output.fileno())
        os.replace(temporary, path)
    except BaseException:
        temporary.unlink(missing_ok=True)
        raise


def _atomic_write_json(path: Path, value: object) -> None:
    _atomic_write(path, json.dumps(value, indent=2, sort_keys=True) + "\n")


def _retain_guest_records(qemu_log: Path, destination: Path) -> None:
    if not qemu_log.is_file():
        _atomic_write(destination, "")
        return
    records = []
    for line in qemu_log.read_text(encoding="utf-8").splitlines():
        stripped = line.strip()
        if stripped.startswith("{"):
            records.append(stripped)
    _atomic_write(destination, "".join(record + "\n" for record in records))


def validate_readiness(record: object, contract: dict, case: Case) -> None:
    if not isinstance(record, dict) or record.get("event") != "ready":
        raise CaseFailure(f"{case.name}: invalid VFIO readiness record")
    try:
        observed = {
            "vendor": _normalize_hex(record["vendor_id"], 4, "readiness vendor"),
            "device": _normalize_hex(record["device_id"], 4, "readiness device"),
            "class": _normalize_hex(record["class_code"], 6, "readiness class"),
        }
        expected = contract["identity"]
        for field, value in observed.items():
            if value != expected[field]:
                raise CaseFailure(
                    f"{case.name}: VFIO readiness {field} mismatch: "
                    f"got {value}, expected {expected[field]}"
                )
        bar_count = record["bar_count"]
        if isinstance(bar_count, bool) or not isinstance(bar_count, int):
            raise CaseFailure(f"{case.name}: VFIO readiness BAR count is invalid")
        if bar_count != len(_validated_expected_bars(contract, case)):
            raise CaseFailure(f"{case.name}: VFIO readiness BAR count mismatch")
    except KeyError as exc:
        raise CaseFailure(
            f"{case.name}: VFIO readiness record is missing {exc.args[0]}"
        ) from exc


def run_case(case: Case, work_root: Path, timeout: int = 120,
             qemu: Path | None = None, kernel: Path | None = None,
             initrd: Path | None = None, shared_memory: bool = False,
             rebind: bool = False, qemu_version: str | None = None) -> dict:
    case_dir = work_root / case.name
    artifacts = case_dir / "artifacts"
    case_dir.mkdir(parents=True, exist_ok=True)
    generation_log = case_dir / "generation.log"
    try:
        if qemu is None:
            raise CaseFailure(f"{case.name}: QEMU guest evidence is required")
        if kernel is None or initrd is None:
            raise CaseFailure("QEMU mode requires --kernel and --initrd")
        run_command(build_command(case, artifacts), generation_log, timeout)
        contract = build_contract(case, artifacts)
        readiness = run_server_smoke(case, artifacts, case_dir)
        validate_readiness(readiness, contract, case)
        result = {"case": case.name, "status": "pass", "readiness": readiness}
        guest = run_qemu_case(case, artifacts, case_dir, qemu, kernel, initrd,
                              shared_memory=shared_memory, rebind=rebind,
                              expected_version=qemu_version)
        result["guest"] = guest.__dict__
    except CaseBlocked as exc:
        result = {
            "case": case.name,
            "status": "blocked",
            "reason": exc.reason,
            "detail": str(exc),
        }
    except (CaseFailure, OSError, json.JSONDecodeError) as exc:
        result = {"case": case.name, "status": "fail", "detail": str(exc)}
    result_path = case_dir / "result.json"
    _atomic_write_json(result_path, result)
    return result


def summarize_results(results: list[dict]) -> dict:
    counts = {"pass": 0, "blocked": 0, "fail": 0}
    seen = set()
    authoritative_results = []
    for result in results:
        case_name = result.get("case")
        status = result.get("status")
        if not isinstance(case_name, str) or not case_name or case_name in seen:
            status = "fail"
            result = {
                "case": case_name if isinstance(case_name, str) else "<invalid>",
                "status": "fail",
                "detail": "matrix produced a missing or duplicate case result",
            }
        elif status not in counts:
            status = "fail"
            result = {
                "case": case_name,
                "status": "fail",
                "detail": "matrix produced an invalid terminal status",
            }
        elif status == "blocked" and result.get("reason") not in ALLOWED_BLOCK_REASONS:
            status = "fail"
            result = {
                "case": case_name,
                "status": "fail",
                "detail": "matrix produced a non-allowlisted blocked result",
            }
        seen.add(result["case"])
        counts[status] += 1
        authoritative_results.append(result)
    exit_code = 1 if counts["fail"] else 0
    return {
        "status": "fail" if exit_code else "pass",
        "exit_code": exit_code,
        "counts": counts,
        "results": authoritative_results,
    }


def main(argv: Sequence[str] | None = None) -> int:
    parser = argparse.ArgumentParser()
    group = parser.add_mutually_exclusive_group(required=True)
    group.add_argument("--case", choices=sorted(CASES))
    group.add_argument("--all", action="store_true")
    parser.add_argument("--work-dir", type=Path, default=ROOT / "vfio-user" / "build" / "matrix")
    parser.add_argument("--qemu", type=Path)
    parser.add_argument("--qemu-version", help="required substring in qemu --version output")
    parser.add_argument("--kernel", type=Path)
    parser.add_argument("--initrd", type=Path)
    parser.add_argument("--shared-memory", action="store_true")
    parser.add_argument("--rebind", action="store_true")
    args = parser.parse_args(argv)
    selected = [CASES[args.case]] if args.case else list(CASES.values())
    results = [run_case(case, args.work_dir, qemu=args.qemu, kernel=args.kernel,
                        initrd=args.initrd, shared_memory=args.shared_memory,
                        rebind=args.rebind, qemu_version=args.qemu_version)
               for case in selected]
    summary = summarize_results(results)
    _atomic_write_json(args.work_dir / "summary.json", summary)
    print(json.dumps(summary, indent=2, sort_keys=True))
    return summary["exit_code"]


CASES = {
    "audio": _case("audio", "profiled", "driver"),
    "ethernet": _case("ethernet", "profiled", "network-interface"),
    "generic": _case("generic", "static", "enumeration"),
    "gpu": _case("gpu", "profiled", "driver"),
    "multibar": _case("multibar", "static", "bar-layout"),
    "nvme": _case("nvme", "nvme", "identify", board="ac701_ft601"),
    "sata": _case("sata", "ahci", "ahci-port"),
    "thunderbolt": _case("thunderbolt", "profiled", "driver"),
    "wifi": _case("wifi", "profiled", "driver"),
    "xhci": _case("xhci", "xhci", "host-controller"),
}


def run_command(argv: Sequence[str], log_path: Path, timeout: int) -> None:
    if not argv:
        raise ValueError("command must not be empty")
    if timeout <= 0:
        raise ValueError("timeout must be positive")

    log_path.parent.mkdir(parents=True, exist_ok=True)
    with log_path.open("wb") as log:
        process = subprocess.Popen(
            list(argv),
            cwd=ROOT,
            stdout=log,
            stderr=subprocess.STDOUT,
            start_new_session=True,
        )
        try:
            status = process.wait(timeout=timeout)
        except subprocess.TimeoutExpired as exc:
            os.killpg(process.pid, signal.SIGTERM)
            try:
                process.wait(timeout=5)
            except subprocess.TimeoutExpired:
                os.killpg(process.pid, signal.SIGKILL)
                process.wait()
            raise CaseFailure(f"command timed out after {timeout}s: {argv[0]}") from exc

    if status != 0:
        raise CaseFailure(f"command exited with exit status {status}: {argv[0]}")


if __name__ == "__main__":
    raise SystemExit(main())
