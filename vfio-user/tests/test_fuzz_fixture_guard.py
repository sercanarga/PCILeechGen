#!/usr/bin/env python3
"""Unit tests for the fail-closed fuzz fixture output guard."""

from __future__ import annotations

import tempfile
import unittest
from pathlib import Path
import sys

sys.path.insert(0, str(Path(__file__).resolve().parent))
from fuzz_fixture_guard import (
    FixturePathError,
    GENERATED_OUTPUT_CONTENT,
    GENERATED_OUTPUT_MARKER,
    validate_fixture_root,
)


class FuzzFixtureGuardTests(unittest.TestCase):
    def setUp(self) -> None:
        self.tempdir = tempfile.TemporaryDirectory()
        self.project = Path(self.tempdir.name) / "vfio-user"
        self.project.mkdir()
        self.fixture = self.project / "build" / "fuzz-fixtures"

    def tearDown(self) -> None:
        self.tempdir.cleanup()

    def test_accepts_the_fixed_project_build_subdirectory(self) -> None:
        self.assertEqual(
            validate_fixture_root(self.project, self.fixture),
            self.fixture.resolve(strict=False),
        )

    def test_rejects_an_arbitrary_output_directory(self) -> None:
        with self.assertRaises(FixturePathError):
            validate_fixture_root(self.project, Path(self.tempdir.name) / "outside")

    def test_rejects_a_symlinked_build_directory(self) -> None:
        outside = Path(self.tempdir.name) / "outside"
        outside.mkdir()
        (self.project / "build").symlink_to(outside, target_is_directory=True)
        with self.assertRaises(FixturePathError):
            validate_fixture_root(self.project, self.fixture)

    def test_adopts_a_manifested_legacy_fixture(self) -> None:
        self.fixture.mkdir(parents=True)
        generated = self.fixture / "nvme"
        generated.mkdir()
        (generated / "build_manifest.json").write_text("{}\n", encoding="utf-8")

        validate_fixture_root(
            self.project,
            self.fixture,
            fixture_name="nvme",
            adopt_generated_output=True,
        )

        self.assertEqual(
            (generated / GENERATED_OUTPUT_MARKER).read_text(encoding="utf-8"),
            GENERATED_OUTPUT_CONTENT,
        )


if __name__ == "__main__":
    unittest.main()
