#!/usr/bin/env python3
"""Select generated root-level SystemVerilog files from a build manifest."""

from __future__ import annotations

import json
import stat
import sys
from pathlib import Path, PurePosixPath
from typing import Any


class SelectionError(ValueError):
    """Raised when a manifest cannot safely drive HDL lint selection."""


def _is_within(root: Path, candidate: Path) -> bool:
    try:
        candidate.relative_to(root)
    except ValueError:
        return False
    return True


def _validated_name(value: Any, index: int) -> PurePosixPath:
    if not isinstance(value, str) or not value:
        raise SelectionError(f"manifest entry {index} has an invalid name")
    if "\x00" in value or "\n" in value or "\r" in value:
        raise SelectionError(f"manifest entry {index} contains a control character")
    if "\\" in value:
        raise SelectionError(
            f"manifest entry {index} must use forward slashes: {value!r}"
        )

    relative = PurePosixPath(value)
    if (
        relative.is_absolute()
        or value in {".", ".."}
        or ".." in relative.parts
        or relative.as_posix() != value
    ):
        raise SelectionError(f"manifest entry {index} has an unsafe path: {value!r}")
    return relative


def _load_entries(manifest_path: Path) -> list[Any]:
    try:
        payload = json.loads(manifest_path.read_text(encoding="utf-8"))
    except json.JSONDecodeError as exc:
        raise SelectionError(f"invalid JSON in {manifest_path}: {exc}") from exc
    if not isinstance(payload, dict):
        raise SelectionError("manifest root must be a JSON object")
    entries = payload.get("files")
    if not isinstance(entries, list):
        raise SelectionError("manifest field 'files' must be an array")
    return entries


def select_sv_files(manifest: Path, output_dir: Path) -> list[Path]:
    try:
        output_root = output_dir.resolve(strict=True)
    except OSError as exc:
        raise SelectionError(f"cannot resolve output directory {output_dir}: {exc}") from exc
    if not output_root.is_dir():
        raise SelectionError(f"output path is not a directory: {output_dir}")

    if manifest.is_symlink():
        raise SelectionError(f"manifest must not be a symlink: {manifest}")
    try:
        manifest_path = manifest.resolve(strict=True)
    except OSError as exc:
        raise SelectionError(f"cannot resolve manifest {manifest}: {exc}") from exc
    if not manifest_path.is_file():
        raise SelectionError(f"manifest is not a regular file: {manifest}")
    if not _is_within(output_root, manifest_path):
        raise SelectionError("manifest is outside the output directory")

    validated: list[tuple[str, PurePosixPath, Path]] = []
    seen: set[str] = set()
    for index, entry in enumerate(_load_entries(manifest_path)):
        if not isinstance(entry, dict):
            raise SelectionError(f"manifest entry {index} must be an object")
        relative = _validated_name(entry.get("name"), index)
        name = relative.as_posix()
        if name in seen:
            raise SelectionError(f"manifest entry {index} duplicates path {name!r}")
        seen.add(name)

        candidate = output_root.joinpath(*relative.parts)
        try:
            resolved = candidate.resolve(strict=False)
        except (OSError, RuntimeError) as exc:
            raise SelectionError(f"cannot resolve manifest path {name!r}: {exc}") from exc
        if not _is_within(output_root, resolved):
            raise SelectionError(f"manifest path escapes output directory: {name!r}")
        validated.append((name, relative, candidate))

    selected: list[Path] = []
    for name, relative, candidate in validated:
        if len(relative.parts) != 1 or relative.suffix != ".sv":
            continue
        if candidate.is_symlink():
            raise SelectionError(f"selected SystemVerilog file is a symlink: {name!r}")
        try:
            mode = candidate.lstat().st_mode
        except OSError as exc:
            raise SelectionError(
                f"cannot inspect selected SystemVerilog file {name!r}: {exc}"
            ) from exc
        if not stat.S_ISREG(mode):
            raise SelectionError(
                f"selected SystemVerilog path is not a regular file: {name!r}"
            )
        try:
            resolved = candidate.resolve(strict=True)
        except OSError as exc:
            raise SelectionError(
                f"cannot resolve selected SystemVerilog file {name!r}: {exc}"
            ) from exc
        if not _is_within(output_root, resolved):
            raise SelectionError(
                f"selected SystemVerilog file escapes output directory: {name!r}"
            )
        selected.append(resolved)

    return sorted(selected, key=lambda item: item.name)


def main(argv: list[str]) -> int:
    if len(argv) != 3:
        print(
            "usage: manifest_sv_files.py MANIFEST OUTPUT_DIR",
            file=sys.stderr,
        )
        return 2
    try:
        selected = select_sv_files(Path(argv[1]), Path(argv[2]))
    except (OSError, SelectionError) as exc:
        print(f"manifest SV selection failed: {exc}", file=sys.stderr)
        return 1
    for file_path in selected:
        print(file_path)
    return 0


if __name__ == "__main__":
    raise SystemExit(main(sys.argv))
