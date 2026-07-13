#!/usr/bin/env python3
"""Fail closed before unit fixtures let the generator rewrite source trees."""

from __future__ import annotations

import argparse
import os
import stat
from pathlib import Path


FIXTURE_NAMES = frozenset(
    {
        "audio",
        "ethernet",
        "generic",
        "gpu",
        "multibar",
        "nvme",
        "sata",
        "thunderbolt",
        "wifi",
        "xhci",
    }
)
SENTINEL = ".pcileech-vfio-test-fixtures-v1"
SENTINEL_CONTENT = "owned by PCILeechGen VFIO-user tests\n"
GENERATED_OUTPUT_MARKER = ".pcileechgen-output-v1"
GENERATED_OUTPUT_CONTENT = "owned by PCILeechGen generated output\n"


class FixturePathError(ValueError):
    """The fixture output tree is not an owned, safe project-local tree."""


def _lstat(path: Path) -> os.stat_result | None:
    try:
        return path.lstat()
    except FileNotFoundError:
        return None


def _require_real_directory(path: Path, label: str, *, required: bool) -> bool:
    metadata = _lstat(path)
    if metadata is None:
        if required:
            raise FixturePathError(f"{label} is missing: {path}")
        return False
    if stat.S_ISLNK(metadata.st_mode):
        raise FixturePathError(f"{label} must not be a symlink: {path}")
    if not stat.S_ISDIR(metadata.st_mode):
        raise FixturePathError(f"{label} must be a directory: {path}")
    return True


def _reject_symlinks_below(directory: Path) -> None:
    for parent, directories, files in os.walk(directory, followlinks=False):
        for name in directories + files:
            path = Path(parent) / name
            if path.is_symlink():
                raise FixturePathError(f"fixture tree contains a symlink: {path}")


def _expected_fixture_root(project_root: Path, fixture_root: Path) -> tuple[Path, Path]:
    raw_project = project_root.absolute()
    if raw_project.is_symlink():
        raise FixturePathError(f"project root must not be a symlink: {raw_project}")
    if not raw_project.is_dir():
        raise FixturePathError(f"project root is missing: {raw_project}")
    project = raw_project.resolve(strict=True)
    raw_build = raw_project / "build"
    raw_root = raw_build / "test-fixtures"
    requested = fixture_root.absolute()

    if requested != raw_root:
        raise FixturePathError(
            "fixture root must be vfio-user/build/test-fixtures"
        )
    for path, label in ((raw_build, "fixture build directory"), (raw_root, "fixture root")):
        metadata = _lstat(path)
        if metadata is not None and stat.S_ISLNK(metadata.st_mode):
            raise FixturePathError(f"{label} must not be a symlink: {path}")
    return project, project / "build" / "test-fixtures"


def _validate_owned_root(root: Path) -> None:
    _require_real_directory(root, "fixture root", required=True)
    sentinel = root / SENTINEL
    metadata = _lstat(sentinel)
    if metadata is None or stat.S_ISLNK(metadata.st_mode) or not stat.S_ISREG(metadata.st_mode):
        raise FixturePathError(f"fixture root is not owned by this test target: {root}")
    if sentinel.read_text(encoding="utf-8") != SENTINEL_CONTENT:
        raise FixturePathError(f"fixture root has an unexpected ownership marker: {root}")

    entries = list(root.iterdir())
    unexpected = sorted(
        entry.name for entry in entries if entry.name not in FIXTURE_NAMES | {SENTINEL}
    )
    if unexpected:
        raise FixturePathError(
            "fixture root contains unexpected entries: " + ", ".join(unexpected)
        )
    for entry in entries:
        if entry.name == SENTINEL:
            continue
        _require_real_directory(entry, f"fixture {entry.name}", required=True)
        _reject_symlinks_below(entry)


def _adopt_generated_output(candidate: Path) -> None:
    manifest = candidate / "build_manifest.json"
    manifest_metadata = _lstat(manifest)
    if (
        manifest_metadata is None
        or stat.S_ISLNK(manifest_metadata.st_mode)
        or not stat.S_ISREG(manifest_metadata.st_mode)
    ):
        raise FixturePathError(
            f"fixture cannot be adopted without a regular build manifest: {candidate}"
        )
    marker = candidate / GENERATED_OUTPUT_MARKER
    marker_metadata = _lstat(marker)
    if marker_metadata is None:
        marker.write_text(GENERATED_OUTPUT_CONTENT, encoding="utf-8")
        return
    if stat.S_ISLNK(marker_metadata.st_mode) or not stat.S_ISREG(marker_metadata.st_mode):
        raise FixturePathError(f"fixture output marker is not a regular file: {marker}")
    if marker.read_text(encoding="utf-8") != GENERATED_OUTPUT_CONTENT:
        raise FixturePathError(f"fixture output marker has unexpected content: {marker}")


def validate_fixture_root(
    project_root: Path,
    fixture_root: Path,
    *,
    prepare: bool = False,
    fixture_name: str | None = None,
    adopt_generated_output: bool = False,
) -> Path:
    """Validate the fixed output root, optionally creating its ownership marker.

    The generator's writer replaces ``<output>/src``.  This guard permits only
    a named fixture below ``vfio-user/build/test-fixtures`` and rejects every
    symlink in an existing fixture tree before the generator is invoked.
    """

    project, root = _expected_fixture_root(project_root, fixture_root)
    build = project / "build"
    if prepare:
        _require_real_directory(build, "fixture build directory", required=False)
        if not build.exists():
            build.mkdir()
        _require_real_directory(build, "fixture build directory", required=True)
        if not root.exists():
            root.mkdir()
            (root / SENTINEL).write_text(SENTINEL_CONTENT, encoding="utf-8")
    _validate_owned_root(root)

    if fixture_name is not None:
        if fixture_name not in FIXTURE_NAMES:
            raise FixturePathError(f"unknown fixture name: {fixture_name}")
        candidate = root / fixture_name
        if candidate.exists():
            _require_real_directory(candidate, f"fixture {fixture_name}", required=True)
            _reject_symlinks_below(candidate)
            if adopt_generated_output:
                _adopt_generated_output(candidate)
        elif adopt_generated_output:
            raise FixturePathError(f"fixture does not exist to adopt: {fixture_name}")
    return root


def main() -> int:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--project-root", required=True, type=Path)
    parser.add_argument("--fixture-root", required=True, type=Path)
    parser.add_argument("--prepare", action="store_true")
    parser.add_argument("--fixture-name", choices=sorted(FIXTURE_NAMES))
    parser.add_argument("--adopt-generated-output", action="store_true")
    args = parser.parse_args()
    try:
        validate_fixture_root(
            args.project_root,
            args.fixture_root,
            prepare=args.prepare,
            fixture_name=args.fixture_name,
            adopt_generated_output=args.adopt_generated_output,
        )
    except (FixturePathError, OSError) as exc:
        parser.error(str(exc))
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
