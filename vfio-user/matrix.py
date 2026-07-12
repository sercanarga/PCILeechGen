#!/usr/bin/env python3
from __future__ import annotations

import os
import json
import argparse
import select
import signal
import subprocess
from dataclasses import dataclass
from pathlib import Path
from typing import Sequence


ROOT = Path(__file__).resolve().parents[1]
GENERATOR = Path(os.environ.get("PCILEECHGEN_BIN", ROOT / "bin" / "pcileechgen"))


class CaseFailure(RuntimeError):
    pass


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
    detail: str


def parse_guest_results(text: str, case: Case) -> GuestResult:
    records = []
    for line in text.splitlines():
        line = line.strip()
        if not line or not line.startswith("{"):
            continue
        try:
            record = json.loads(line)
        except json.JSONDecodeError as exc:
            raise CaseFailure(f"invalid guest result JSON: {exc}") from exc
        if record.get("event") == "result":
            records.append(record)
    if len(records) != 1:
        raise CaseFailure(f"expected one terminal guest result, got {len(records)}")
    record = records[0]
    if record.get("case") != case.name:
        raise CaseFailure(f"guest result case mismatch: {record.get('case')}")
    if record.get("status") not in {"pass", "skip", "fail"}:
        raise CaseFailure("guest result has invalid status")
    required = ("bdf", "vendor", "device", "class", "driver")
    missing = [key for key in required if key not in record]
    if missing:
        raise CaseFailure(f"guest result missing fields: {','.join(missing)}")
    return GuestResult(
        case=case.name,
        status=record["status"],
        bdf=record["bdf"],
        vendor=record["vendor"],
        device=record["device"],
        class_code=record["class"],
        driver=record["driver"],
        detail=record.get("detail", ""),
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


def start_server(case: Case, artifacts: Path, work_dir: Path, timeout: int = 10):
    binary = ROOT / "vfio-user" / "build" / "vfio-device"
    if not binary.is_file():
        raise CaseFailure(f"VFIO server binary is missing: {binary}")
    socket_path = work_dir / "device.sock"
    log_path = work_dir / "server.log"
    work_dir.mkdir(parents=True, exist_ok=True)
    with log_path.open("w", encoding="utf-8") as log:
        process = subprocess.Popen(
            [str(binary), "--artifacts", str(artifacts), "--socket", str(socket_path)],
            cwd=ROOT / "vfio-user",
            stdout=subprocess.PIPE,
            stderr=log,
            text=True,
            start_new_session=True,
        )
        try:
            ready, _, _ = select.select([process.stdout], [], [], timeout)
            if not ready:
                raise CaseFailure(f"{case.name}: VFIO server did not become ready")
            line = process.stdout.readline()
            record = json.loads(line)
            if record.get("event") != "ready":
                raise CaseFailure(f"{case.name}: invalid readiness record")
            return process, process.stdout, socket_path, record
        finally:
            if process.poll() is not None:
                process.stdout.close()


def stop_server(process: subprocess.Popen, output) -> None:
    if process.poll() is None:
        process.terminate()
        try:
            process.wait(timeout=5)
        except subprocess.TimeoutExpired:
            process.kill()
            process.wait()
    output.close()


def qemu_requires_kvm(case: Case, kvm_path: Path = Path("/dev/kvm")) -> bool:
    return case.name == "nvme" and not kvm_path.exists()


def run_server_smoke(case: Case, artifacts: Path, work_dir: Path, timeout: int = 10) -> dict:
    process, output, _, record = start_server(case, artifacts, work_dir, timeout)
    stop_server(process, output)
    return record


def run_qemu_case(case: Case, artifacts: Path, work_dir: Path,
                  qemu: Path, kernel: Path, initrd: Path, timeout: int = 30,
                  shared_memory: bool = False, rebind: bool = False) -> GuestResult:
    if qemu_requires_kvm(case):
        raise CaseFailure(f"{case.name}: /dev/kvm is required for QEMU MSI-X E2E")
    process, output, socket_path, _ = start_server(case, artifacts, work_dir)
    qemu_log = work_dir / "qemu.log"
    command = build_qemu_command(case, socket_path, kernel, initrd, artifacts,
                                 qemu=qemu, shared_memory=shared_memory,
                                 rebind=rebind)
    try:
        with qemu_log.open("w", encoding="utf-8") as log:
            try:
                result = subprocess.run(command, stdout=log, stderr=subprocess.STDOUT,
                                        timeout=timeout, check=False)
            except subprocess.TimeoutExpired as exc:
                raise CaseFailure(f"{case.name}: QEMU timed out after {timeout}s") from exc
        text = qemu_log.read_text(encoding="utf-8")
        guest = parse_guest_results(text, case)
        if result.returncode != 0 and guest.status != "pass":
            raise CaseFailure(f"{case.name}: QEMU exited {result.returncode}")
        return guest
    finally:
        stop_server(process, output)


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


def run_case(case: Case, work_root: Path, timeout: int = 120,
             qemu: Path | None = None, kernel: Path | None = None,
             initrd: Path | None = None, shared_memory: bool = False,
             rebind: bool = False) -> dict:
    case_dir = work_root / case.name
    artifacts = case_dir / "artifacts"
    case_dir.mkdir(parents=True, exist_ok=True)
    generation_log = case_dir / "generation.log"
    try:
        run_command(build_command(case, artifacts), generation_log, timeout)
        readiness = run_server_smoke(case, artifacts, case_dir)
        result = {"case": case.name, "status": "pass", "readiness": readiness}
        if qemu is not None:
            if kernel is None or initrd is None:
                raise CaseFailure("QEMU mode requires --kernel and --initrd")
            guest = run_qemu_case(case, artifacts, case_dir, qemu, kernel, initrd,
                                  shared_memory=shared_memory, rebind=rebind)
            if guest.status == "fail" or (case.mandatory_probe == "driver" and guest.status == "fail"):
                raise CaseFailure(f"{case.name}: guest probe failed: {guest.detail}")
            result["guest"] = guest.__dict__
    except (CaseFailure, OSError, json.JSONDecodeError) as exc:
        result = {"case": case.name, "status": "fail", "detail": str(exc)}
    result_path = case_dir / "result.json"
    result_path.write_text(json.dumps(result, indent=2) + "\n", encoding="utf-8")
    return result


def main(argv: Sequence[str] | None = None) -> int:
    parser = argparse.ArgumentParser()
    group = parser.add_mutually_exclusive_group(required=True)
    group.add_argument("--case", choices=sorted(CASES))
    group.add_argument("--all", action="store_true")
    parser.add_argument("--work-dir", type=Path, default=ROOT / "vfio-user" / "build" / "matrix")
    parser.add_argument("--qemu", type=Path)
    parser.add_argument("--kernel", type=Path)
    parser.add_argument("--initrd", type=Path)
    parser.add_argument("--shared-memory", action="store_true")
    parser.add_argument("--rebind", action="store_true")
    args = parser.parse_args(argv)
    selected = [CASES[args.case]] if args.case else list(CASES.values())
    results = [run_case(case, args.work_dir, qemu=args.qemu, kernel=args.kernel,
                        initrd=args.initrd, shared_memory=args.shared_memory,
                        rebind=args.rebind)
               for case in selected]
    print(json.dumps(results, indent=2))
    return 0 if all(result["status"] == "pass" for result in results) else 1


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
