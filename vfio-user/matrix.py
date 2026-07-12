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
        if not line.strip():
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
        str(ROOT / "bin" / "pcileechgen"),
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


def run_server_smoke(case: Case, artifacts: Path, work_dir: Path, timeout: int = 10) -> dict:
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
            return record
        finally:
            if process.poll() is None:
                process.terminate()
                try:
                    process.wait(timeout=5)
                except subprocess.TimeoutExpired:
                    process.kill()
                    process.wait()
            if process.stdout is not None:
                process.stdout.close()


def run_case(case: Case, work_root: Path, timeout: int = 120) -> dict:
    case_dir = work_root / case.name
    artifacts = case_dir / "artifacts"
    case_dir.mkdir(parents=True, exist_ok=True)
    generation_log = case_dir / "generation.log"
    try:
        run_command(build_command(case, artifacts), generation_log, timeout)
        readiness = run_server_smoke(case, artifacts, case_dir)
        result = {"case": case.name, "status": "pass", "readiness": readiness}
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
    args = parser.parse_args(argv)
    selected = [CASES[args.case]] if args.case else list(CASES.values())
    results = [run_case(case, args.work_dir) for case in selected]
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
