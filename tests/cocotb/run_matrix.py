#!/usr/bin/env python3
"""Generate isolated fixtures and run the supported cocotb matrix."""

from __future__ import annotations

import argparse
import os
import re
import shutil
import subprocess
import sys
import tempfile
import xml.etree.ElementTree as ET
from dataclasses import dataclass
from pathlib import Path
from typing import Iterable, Optional, Sequence


ROOT = Path(__file__).resolve().parents[2]
COCOTB_DIR = ROOT / "tests" / "cocotb"
OUTPUT_RELATIVE = Path("tests/cocotb/out_matrix")
VENV_RELATIVE = Path("bin/cocotb-venv")
OUTPUT_SENTINEL = ".pcileechgen-cocotb-output-v1"
OUTPUT_SENTINEL_CONTENT = "owned by PCILeechGen cocotb matrix\n"
MIN_PYTHON = (3, 9)
MAX_PYTHON = (3, 13)
NVME_TIMEOUT_CYCLES = "000400"


class MatrixError(RuntimeError):
    """A matrix precondition or test result failed."""


@dataclass(frozen=True)
class Case:
    name: str
    fixture: str
    module: str
    tb_top: str = "tb_top"
    timeout_cycles: Optional[str] = None


CASES: tuple[Case, ...] = (
    Case("nvme", "nvme", "test_pcie_host", timeout_cycles=NVME_TIMEOUT_CYCLES),
    Case("generic", "generic", "test_pio", tb_top="tb_top_generic"),
    Case("audio", "audio", "test_audio_init"),
    Case("xhci", "xhci", "test_xhci_init"),
    Case("ethernet", "ethernet", "test_eth_init"),
    Case("wifi", "wifi", "test_wifi_init"),
    Case("sata", "sata", "test_sata_init"),
    Case("gpu", "gpu", "test_gpu_init"),
)

CASE_NAMES = frozenset(case.name for case in CASES)
KNOWN_OUTPUT_FILES = frozenset({"generation.log", "summary.tsv"})


def ensure_supported_python(version_info: Sequence[int] = sys.version_info) -> None:
    version = tuple(version_info[:2])
    if version < MIN_PYTHON or version > MAX_PYTHON:
        raise MatrixError(
            "cocotb requires Python "
            f"{MIN_PYTHON[0]}.{MIN_PYTHON[1]} through "
            f"{MAX_PYTHON[0]}.{MAX_PYTHON[1]}; got {version[0]}.{version[1]}"
        )


def _canonical_repo(repo_root: Path) -> Path:
    repo = repo_root.absolute()
    if repo.is_symlink() or repo.resolve() != repo:
        raise MatrixError(f"repository root must be a canonical non-symlink path: {repo}")
    return repo


def _reject_symlink_components(repo: Path, relative: Path) -> None:
    current = repo
    for component in relative.parts:
        current = current / component
        if current.is_symlink():
            raise MatrixError(f"symlink path component is not allowed: {current}")


def _validate_exact_repo_relative(
    raw_path: Path,
    expected: Path,
    *,
    repo_root: Path = ROOT,
) -> Path:
    raw_text = os.fspath(raw_path)
    expected_text = expected.as_posix()
    if raw_path.is_absolute() or raw_text != expected_text:
        raise MatrixError(
            f"path must be the canonical repository-relative path {expected_text!r}"
        )

    repo = _canonical_repo(repo_root)
    _reject_symlink_components(repo, expected)
    candidate = repo / expected
    if candidate.resolve(strict=False) != candidate:
        raise MatrixError(f"path does not resolve canonically: {candidate}")
    return candidate


def validate_output_root(raw_path: Path, *, repo_root: Path = ROOT) -> Path:
    output_root = _validate_exact_repo_relative(
        raw_path, OUTPUT_RELATIVE, repo_root=repo_root
    )
    repo = _canonical_repo(repo_root)
    forbidden = {
        Path("/").resolve(),
        Path.home().resolve(),
        repo,
        repo / "tests",
        repo / "tests" / "cocotb",
    }
    if output_root in forbidden or output_root.parent != repo / "tests" / "cocotb":
        raise MatrixError(f"refusing unsafe cocotb output root: {output_root}")
    return output_root


def validate_venv_path(raw_path: Path, *, repo_root: Path = ROOT) -> Path:
    venv = _validate_exact_repo_relative(raw_path, VENV_RELATIVE, repo_root=repo_root)
    repo = _canonical_repo(repo_root)
    if venv.parent != repo / "bin":
        raise MatrixError(f"refusing unsafe cocotb virtualenv path: {venv}")
    return venv


def _require_output_sentinel(output_root: Path) -> None:
    sentinel = output_root / OUTPUT_SENTINEL
    if sentinel.is_symlink() or not sentinel.is_file():
        raise MatrixError(
            f"refusing unowned output directory without sentinel: {output_root}"
        )
    try:
        content = sentinel.read_text(encoding="utf-8")
    except OSError as exc:
        raise MatrixError(f"cannot read output sentinel {sentinel}: {exc}") from exc
    if content != OUTPUT_SENTINEL_CONTENT:
        raise MatrixError(f"output sentinel has unexpected content: {sentinel}")


def _validate_owned_output_contents(output_root: Path) -> None:
    if output_root.is_symlink() or not output_root.is_dir():
        raise MatrixError(f"cocotb output root is not a real directory: {output_root}")
    _require_output_sentinel(output_root)

    allowed = CASE_NAMES | KNOWN_OUTPUT_FILES | {OUTPUT_SENTINEL}
    entries = list(output_root.iterdir())
    unexpected = sorted(entry.name for entry in entries if entry.name not in allowed)
    if unexpected:
        raise MatrixError(
            "refusing to clear output directory with unrecognized entries: "
            + ", ".join(unexpected)
        )

    for entry in entries:
        if entry.is_symlink():
            raise MatrixError(f"symlink in cocotb output is not allowed: {entry}")
        if entry.name in CASE_NAMES:
            if not entry.is_dir():
                raise MatrixError(f"expected case directory, got non-directory: {entry}")
            for descendant in entry.rglob("*"):
                if descendant.is_symlink():
                    raise MatrixError(
                        f"symlink in cocotb case output is not allowed: {descendant}"
                    )
        elif entry.name in KNOWN_OUTPUT_FILES and not entry.is_file():
            raise MatrixError(f"expected output file, got non-file: {entry}")


def _clear_known_output_contents(output_root: Path) -> None:
    """Clear only prevalidated, known children; never recursively remove the root."""
    for case_name in sorted(CASE_NAMES):
        case_dir = output_root / case_name
        if case_dir.exists():
            shutil.rmtree(case_dir)
    for filename in sorted(KNOWN_OUTPUT_FILES):
        output_file = output_root / filename
        if output_file.exists():
            output_file.unlink()


def prepare_output_root(raw_path: Path, *, repo_root: Path = ROOT) -> Path:
    output_root = validate_output_root(raw_path, repo_root=repo_root)
    if not output_root.exists():
        if not output_root.parent.is_dir():
            raise MatrixError(f"cocotb source directory is missing: {output_root.parent}")
        output_root.mkdir()
        (output_root / OUTPUT_SENTINEL).write_text(
            OUTPUT_SENTINEL_CONTENT, encoding="utf-8"
        )
        return output_root

    _validate_owned_output_contents(output_root)
    _clear_known_output_contents(output_root)
    return output_root


def clean_owned_output_root(raw_path: Path, *, repo_root: Path = ROOT) -> None:
    output_root = validate_output_root(raw_path, repo_root=repo_root)
    if not output_root.exists():
        return
    _validate_owned_output_contents(output_root)
    _clear_known_output_contents(output_root)
    (output_root / OUTPUT_SENTINEL).unlink()
    output_root.rmdir()


def _validate_fixture_tree(source: Path) -> None:
    if source.is_symlink() or not source.is_dir():
        raise MatrixError(f"generated fixture is not a real directory: {source}")
    for item in source.rglob("*"):
        if item.is_symlink():
            raise MatrixError(f"generated fixture contains a symlink: {item}")


def stage_fixture(source: Path, destination: Path) -> Path:
    _validate_fixture_tree(source)
    if destination.exists() or destination.is_symlink():
        raise MatrixError(f"staged fixture destination already exists: {destination}")
    shutil.copytree(source, destination, symlinks=False)
    return destination


TIMEOUT_PATTERN = re.compile(
    r"(\.TIMEOUT_CYCLES\s*\(\s*24'h)[0-9A-Fa-f]+(\s*\))"
)


def patch_nvme_timeout(staged_fixture: Path, timeout_cycles: str) -> Path:
    if not re.fullmatch(r"[0-9A-Fa-f]{1,6}", timeout_cycles):
        raise MatrixError(f"invalid 24-bit timeout value: {timeout_cycles!r}")
    bridge = staged_fixture / "pcileech_nvme_dma_bridge.sv"
    if bridge.is_symlink() or not bridge.is_file():
        raise MatrixError(f"NVMe DMA bridge is missing from staged fixture: {bridge}")
    original = bridge.read_text(encoding="utf-8")
    updated, replacements = TIMEOUT_PATTERN.subn(
        rf"\g<1>{timeout_cycles.upper()}\g<2>", original
    )
    if replacements != 1:
        raise MatrixError(
            f"expected one NVMe timeout parameter in {bridge}, found {replacements}"
        )
    bridge.write_text(updated, encoding="utf-8")
    return bridge


def active_venv_tools(executable: Optional[Path] = None) -> tuple[Path, Path]:
    """Return tools next to the invoked interpreter without resolving venv symlinks."""
    python = executable if executable is not None else Path(sys.executable)
    return python, python.parent / "cocotb-config"


def find_cocotb_makefiles(executable: Optional[Path] = None) -> Path:
    _, cocotb_config = active_venv_tools(executable)
    if not cocotb_config.is_file() or not os.access(cocotb_config, os.X_OK):
        raise MatrixError(
            f"cocotb-config is missing next to the active interpreter: {cocotb_config}"
        )
    try:
        completed = subprocess.run(
            [str(cocotb_config), "--makefiles"],
            check=True,
            text=True,
            stdout=subprocess.PIPE,
            stderr=subprocess.PIPE,
        )
    except (OSError, subprocess.CalledProcessError) as exc:
        raise MatrixError(f"cannot locate cocotb makefiles: {exc}") from exc
    makefiles = Path(completed.stdout.strip())
    if not (makefiles / "Makefile.sim").is_file():
        raise MatrixError(f"invalid cocotb makefiles directory: {makefiles}")
    return makefiles


def generate_fixtures(generator: Path, output_root: Path, log_path: Path) -> None:
    if not generator.is_file() or not os.access(generator, os.X_OK):
        raise MatrixError(f"generator is not executable: {generator}")
    script = COCOTB_DIR / "generate_all.sh"
    command = [
        "/bin/bash",
        str(script),
        "--generator",
        str(generator),
        "--output-root",
        str(output_root),
    ]
    try:
        with log_path.open("w", encoding="utf-8") as log_file:
            completed = subprocess.run(
                command,
                cwd=ROOT,
                text=True,
                stdout=log_file,
                stderr=subprocess.STDOUT,
            )
    except OSError as exc:
        raise MatrixError(f"fixture generation could not start: {exc}") from exc
    if completed.returncode != 0:
        raise MatrixError(
            f"fixture generation exited with status {completed.returncode}; "
            f"see {log_path}"
        )


def validate_results_xml(results_file: Path) -> tuple[int, int]:
    if not results_file.is_file():
        raise MatrixError(f"cocotb result XML is missing: {results_file}")
    try:
        root = ET.parse(results_file).getroot()
    except (OSError, ET.ParseError) as exc:
        raise MatrixError(f"cocotb result XML is unreadable: {results_file}: {exc}") from exc

    testcases = root.findall(".//testcase")
    if not testcases:
        raise MatrixError(f"cocotb result XML contains no test cases: {results_file}")
    failed = sum(
        1
        for testcase in testcases
        if testcase.find("failure") is not None or testcase.find("error") is not None
    )
    if failed:
        raise MatrixError(
            f"cocotb reported {failed} failed/error test case(s) in {results_file}"
        )
    return len(testcases), failed


def case_paths(output_root: Path, case: Case) -> tuple[Path, Path, Path, Path]:
    case_dir = output_root / case.name
    return (
        case_dir,
        case_dir / "fixture",
        case_dir / "sim_build",
        case_dir / "results.xml",
    )


def run_case(
    case: Case,
    generated_root: Path,
    output_root: Path,
    cocotb_makefiles: Path,
    *,
    sim: str,
) -> int:
    case_dir, staged_fixture, sim_build, results_file = case_paths(output_root, case)
    case_dir.mkdir()
    stage_fixture(generated_root / case.fixture, staged_fixture)
    if case.timeout_cycles is not None:
        patch_nvme_timeout(staged_fixture, case.timeout_cycles)

    log_path = case_dir / "simulation.log"
    command = [
        "make",
        "--no-print-directory",
        "-f",
        str(COCOTB_DIR / "Makefile"),
        f"SIM={sim}",
        f"FIXTURE={staged_fixture}",
        f"TB_TOP={case.tb_top}",
        f"COCOTB_TEST_MODULES={case.module}",
        f"SIM_BUILD={sim_build}",
        f"COCOTB_RESULTS_FILE={results_file}",
        f"COCOTB_MAKEFILES={cocotb_makefiles}",
    ]
    python, _ = active_venv_tools()
    env = os.environ.copy()
    env.pop("PYTHONPATH", None)
    env["PATH"] = str(python.parent) + os.pathsep + env.get("PATH", "")
    try:
        with log_path.open("w", encoding="utf-8") as log_file:
            completed = subprocess.run(
                command,
                cwd=case_dir,
                env=env,
                text=True,
                stdout=log_file,
                stderr=subprocess.STDOUT,
            )
    except OSError as exc:
        try:
            log_path.write_text(f"simulator could not start: {exc}\n", encoding="utf-8")
        except OSError:
            pass
        raise MatrixError(f"{case.name}: simulator could not start: {exc}") from exc
    if completed.returncode != 0:
        raise MatrixError(
            f"{case.name}: simulator exited with status {completed.returncode}; "
            f"see {log_path}"
        )
    tests, _ = validate_results_xml(results_file)
    return tests


def selected_cases(names: Optional[Iterable[str]]) -> list[Case]:
    if not names:
        return list(CASES)
    selected = set(names)
    return [case for case in CASES if case.name in selected]


def write_summary(
    output_root: Path, rows: Sequence[tuple[str, str, str]]
) -> Path:
    summary = output_root / "summary.tsv"
    lines = ["case\tstatus\tdetail"]
    lines.extend("\t".join(row) for row in rows)
    summary.write_text("\n".join(lines) + "\n", encoding="utf-8")
    return summary


def run_matrix(args: argparse.Namespace) -> int:
    ensure_supported_python()
    cases = selected_cases(args.case)
    generator = args.generator if args.generator.is_absolute() else ROOT / args.generator
    if not generator.is_file() or not os.access(generator, os.X_OK):
        raise MatrixError(f"generator is not executable: {generator}")
    cocotb_makefiles = find_cocotb_makefiles()
    output_root = prepare_output_root(args.output_root)
    rows: list[tuple[str, str, str]] = []

    with tempfile.TemporaryDirectory(prefix="pcileechgen-cocotb-") as temp_dir:
        generated_root = Path(temp_dir) / "fixtures"
        generate_fixtures(
            generator,
            generated_root,
            output_root / "generation.log",
        )
        for case in cases:
            try:
                tests = run_case(
                    case,
                    generated_root,
                    output_root,
                    cocotb_makefiles,
                    sim=args.sim,
                )
            except MatrixError as exc:
                detail = str(exc).replace("\t", " ").replace("\n", " ")
                rows.append((case.name, "FAIL", detail))
                print(f"{case.name}: FAIL - {detail}")
            else:
                rows.append((case.name, "PASS", f"{tests} tests"))
                print(f"{case.name}: PASS ({tests} tests)")

    summary = write_summary(output_root, rows)
    print(f"matrix summary: {summary}")
    return 1 if any(status == "FAIL" for _, status, _ in rows) else 0


def build_parser() -> argparse.ArgumentParser:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument(
        "--generator",
        type=Path,
        default=ROOT / "bin" / "pcileechgen",
        help="path to the built pcileechgen binary",
    )
    parser.add_argument(
        "--output-root",
        type=Path,
        default=OUTPUT_RELATIVE,
        help=f"must be exactly {OUTPUT_RELATIVE.as_posix()}",
    )
    parser.add_argument("--sim", default="verilator")
    parser.add_argument("--case", action="append", choices=sorted(CASE_NAMES))
    parser.add_argument("--check-python", action="store_true")
    parser.add_argument("--check-venv", type=Path)
    parser.add_argument("--clean", action="store_true")
    return parser


def main(argv: Optional[Sequence[str]] = None) -> int:
    args = build_parser().parse_args(argv)
    try:
        if args.check_python:
            ensure_supported_python()
        if args.check_venv is not None:
            validate_venv_path(args.check_venv)
        if args.clean:
            clean_owned_output_root(args.output_root)
            return 0
        if args.check_python or args.check_venv is not None:
            return 0
        return run_matrix(args)
    except (MatrixError, OSError) as exc:
        print(f"cocotb matrix: {exc}", file=sys.stderr)
        return 2


if __name__ == "__main__":
    raise SystemExit(main())
