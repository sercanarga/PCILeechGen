#!/usr/bin/env python3
"""Fail closed before the fuzz gate lets the generator replace fixture files."""

from __future__ import annotations

import argparse
from pathlib import Path


class FixturePathError(ValueError):
    """The generated fixture directory is not a safe project-local directory."""


def validate_fixture_root(project_root: Path, fixture_root: Path) -> Path:
    """Return the canonical safe fixture root or raise ``FixturePathError``.

    The generator rewrites its output's ``src`` directory, so accepting a
    caller-controlled output path or a symlink would let a normal fuzz target
    mutate files outside this checkout.  Keep the output location fixed below
    ``vfio-user/build`` and refuse symlinks before anything creates it.
    """

    project = project_root.resolve(strict=True)
    raw_build = project_root.absolute() / "build"
    raw_expected = raw_build / "fuzz-fixtures"
    build = project / "build"
    expected = build / "fuzz-fixtures"

    for path in (raw_build, raw_expected):
        if path.is_symlink():
            raise FixturePathError(f"fixture path must not traverse a symlink: {path}")
    canonical = fixture_root.resolve(strict=False)
    if canonical != expected.resolve(strict=False):
        raise FixturePathError("fixture root must be vfio-user/build/fuzz-fixtures")
    try:
        canonical.relative_to(build.resolve(strict=False))
    except ValueError as exc:
        raise FixturePathError("fixture root escapes vfio-user/build") from exc
    if canonical == build.resolve(strict=False):
        raise FixturePathError("fixture root must be a child of vfio-user/build")
    return canonical


def main() -> int:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--project-root", required=True, type=Path)
    parser.add_argument("--fixture-root", required=True, type=Path)
    args = parser.parse_args()
    try:
        validate_fixture_root(args.project_root, args.fixture_root)
    except (FixturePathError, OSError) as exc:
        parser.error(str(exc))
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
