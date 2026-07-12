#!/usr/bin/env python3
from __future__ import annotations

import os
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


def _case(name: str, behavior: str, mandatory_probe: str, board: str = "PCIeSquirrel") -> Case:
    return Case(
        name=name,
        fixture=ROOT / "testdata" / "donors" / f"{name}.json",
        board=board,
        behavior=behavior,
        mandatory_probe=mandatory_probe,
    )


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
    raise SystemExit("matrix orchestration is not implemented yet; use the unit-test targets")
