#!/usr/bin/env python3
"""Unit tests for the fail-closed VFIO-user unit-fixture guard."""

from __future__ import annotations

import tempfile
import unittest
from pathlib import Path
import sys

sys.path.insert(0, str(Path(__file__).resolve().parent))
from test_fixture_root_guard import (
    FixturePathError,
    GENERATED_OUTPUT_CONTENT,
    GENERATED_OUTPUT_MARKER,
    SENTINEL,
    validate_fixture_root,
)


class TestFixtureRootGuardTests(unittest.TestCase):
    def setUp(self) -> None:
        self.tempdir = tempfile.TemporaryDirectory()
        self.project = Path(self.tempdir.name) / "vfio-user"
        self.project.mkdir()
        self.fixture = self.project / "build" / "test-fixtures"

    def tearDown(self) -> None:
        self.tempdir.cleanup()

    def test_prepare_marks_the_fixed_project_build_subdirectory(self) -> None:
        self.assertEqual(
            validate_fixture_root(self.project, self.fixture, prepare=True),
            self.fixture.resolve(strict=False),
        )
        self.assertTrue((self.fixture / SENTINEL).is_file())

    def test_rejects_an_arbitrary_output_directory(self) -> None:
        with self.assertRaises(FixturePathError):
            validate_fixture_root(
                self.project,
                Path(self.tempdir.name) / "outside",
                prepare=True,
            )

    def test_rejects_an_unowned_existing_root(self) -> None:
        self.fixture.mkdir(parents=True)
        with self.assertRaisesRegex(FixturePathError, "not owned"):
            validate_fixture_root(self.project, self.fixture)

    def test_rejects_a_symlinked_build_directory(self) -> None:
        outside = Path(self.tempdir.name) / "outside"
        outside.mkdir()
        (self.project / "build").symlink_to(outside, target_is_directory=True)
        with self.assertRaises(FixturePathError):
            validate_fixture_root(self.project, self.fixture, prepare=True)

    def test_rejects_a_symlink_in_an_existing_fixture(self) -> None:
        validate_fixture_root(self.project, self.fixture, prepare=True)
        generated = self.fixture / "nvme"
        generated.mkdir()
        (generated / "src").symlink_to(Path(self.tempdir.name), target_is_directory=True)
        with self.assertRaisesRegex(FixturePathError, "symlink"):
            validate_fixture_root(self.project, self.fixture, fixture_name="nvme")

    def test_rejects_unexpected_root_entries(self) -> None:
        validate_fixture_root(self.project, self.fixture, prepare=True)
        (self.fixture / "unrelated").mkdir()
        with self.assertRaisesRegex(FixturePathError, "unexpected entries"):
            validate_fixture_root(self.project, self.fixture)

    def test_adopts_a_verified_legacy_fixture_after_its_manifest_is_checked(self) -> None:
        validate_fixture_root(self.project, self.fixture, prepare=True)
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
