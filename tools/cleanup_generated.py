#!/usr/bin/env python3
"""Remove only fixed, generated files after validating their project paths."""

from __future__ import annotations

import argparse
import os
import re
import stat
from pathlib import Path


TOP_LEVEL_FILES = (
    Path("bin/pcileechgen"),
    Path("hdl-lint-report.tsv"),
    Path("coverage.out"),
    Path("coverage.html"),
)
VFIO_BUILD_FILES = (
    Path("vfio-device"),
    Path("test_device_model"),
    Path("test_behavior_static"),
    Path("test_behavior_nvme"),
    Path("test_behavior_hda"),
    Path("test_behavior_gpu"),
    Path("test_behavior_storage"),
    Path("fuzz_behavior"),
    Path("fuzz-summary.txt"),
    Path("fuzz-ci.log"),
    Path("fuzz-fixtures.log"),
)
EARLY_ACCESS_ARCHIVE = re.compile(r"^PCILeechGen-EarlyAccess-[0-9]{6}\.zip$")


class CleanupError(ValueError):
    """A requested generated file is not safe to remove."""


def _lstat(path: Path) -> os.stat_result | None:
    try:
        return path.lstat()
    except FileNotFoundError:
        return None


def _require_real_directory(path: Path, label: str, *, required: bool) -> bool:
    metadata = _lstat(path)
    if metadata is None:
        if required:
            raise CleanupError(f"{label} is missing: {path}")
        return False
    if stat.S_ISLNK(metadata.st_mode):
        raise CleanupError(f"{label} must not be a symlink: {path}")
    if not stat.S_ISDIR(metadata.st_mode):
        raise CleanupError(f"{label} must be a directory: {path}")
    return True


def _canonical_root(raw_root: Path) -> Path:
    root = raw_root.absolute()
    if root.is_symlink():
        raise CleanupError(f"repository root must not be a symlink: {root}")
    _require_real_directory(root, "repository root", required=True)
    return root.resolve(strict=True)


def _validated_file(root: Path, relative: Path) -> Path:
    if relative.is_absolute() or ".." in relative.parts or not relative.parts:
        raise CleanupError(f"unsafe generated-file path: {relative}")
    parent = root
    for component in relative.parts[:-1]:
        parent = parent / component
        metadata = _lstat(parent)
        if metadata is None:
            return root / relative
        if stat.S_ISLNK(metadata.st_mode):
            raise CleanupError(f"generated-file parent must not be a symlink: {parent}")
        if not stat.S_ISDIR(metadata.st_mode):
            raise CleanupError(f"generated-file parent is not a directory: {parent}")
    return root / relative


def _remove_regular_file(root: Path, relative: Path) -> None:
    candidate = _validated_file(root, relative)
    metadata = _lstat(candidate)
    if metadata is None:
        return
    if stat.S_ISLNK(metadata.st_mode):
        raise CleanupError(f"refusing to remove generated-file symlink: {candidate}")
    if not stat.S_ISREG(metadata.st_mode):
        raise CleanupError(f"refusing to remove non-file generated output: {candidate}")
    candidate.unlink()


def cleanup_top_level(repo_root: Path) -> None:
    root = _canonical_root(repo_root)
    for relative in TOP_LEVEL_FILES:
        _remove_regular_file(root, relative)
    cleanup_early_access_archives(root)


def cleanup_early_access_archives(repo_root: Path) -> None:
    root = _canonical_root(repo_root)
    for entry in root.iterdir():
        if EARLY_ACCESS_ARCHIVE.fullmatch(entry.name):
            _remove_regular_file(root, Path(entry.name))


def cleanup_vfio_build(repo_root: Path) -> None:
    root = _canonical_root(repo_root)
    build = root / "vfio-user" / "build"
    if not _require_real_directory(build, "VFIO-user build directory", required=False):
        return
    for relative in VFIO_BUILD_FILES:
        _remove_regular_file(build, relative)


def main() -> int:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--repo-root", required=True, type=Path)
    group = parser.add_mutually_exclusive_group(required=True)
    group.add_argument("--top-level", action="store_true")
    group.add_argument("--vfio-build", action="store_true")
    group.add_argument("--early-access-archives", action="store_true")
    args = parser.parse_args()
    try:
        if args.top_level:
            cleanup_top_level(args.repo_root)
        elif args.vfio_build:
            cleanup_vfio_build(args.repo_root)
        else:
            cleanup_early_access_archives(args.repo_root)
    except (CleanupError, OSError) as exc:
        parser.error(str(exc))
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
