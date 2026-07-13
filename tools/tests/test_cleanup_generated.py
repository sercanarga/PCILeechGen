#!/usr/bin/env python3
"""Unit tests for the restricted generated-artifact cleanup helper."""

from __future__ import annotations

import tempfile
import unittest
from pathlib import Path
import sys

sys.path.insert(0, str(Path(__file__).resolve().parents[1]))
from cleanup_generated import CleanupError, cleanup_top_level, cleanup_vfio_build


class CleanupGeneratedTests(unittest.TestCase):
    def setUp(self) -> None:
        self.tempdir = tempfile.TemporaryDirectory()
        self.root = Path(self.tempdir.name) / "PCILeechGen"
        self.root.mkdir()

    def tearDown(self) -> None:
        self.tempdir.cleanup()

    def test_top_level_cleanup_removes_only_known_regular_outputs(self) -> None:
        (self.root / "bin").mkdir()
        (self.root / "bin" / "pcileechgen").write_text("binary", encoding="utf-8")
        (self.root / "coverage.out").write_text("coverage", encoding="utf-8")
        archive = self.root / "PCILeechGen-EarlyAccess-071326.zip"
        archive.write_text("archive", encoding="utf-8")
        keep = self.root / "keep.txt"
        keep.write_text("keep", encoding="utf-8")

        cleanup_top_level(self.root)

        self.assertFalse((self.root / "bin" / "pcileechgen").exists())
        self.assertFalse((self.root / "coverage.out").exists())
        self.assertFalse(archive.exists())
        self.assertTrue(keep.exists())

    def test_top_level_cleanup_rejects_a_symlinked_parent(self) -> None:
        outside = Path(self.tempdir.name) / "outside"
        outside.mkdir()
        (self.root / "bin").symlink_to(outside, target_is_directory=True)

        with self.assertRaisesRegex(CleanupError, "must not be a symlink"):
            cleanup_top_level(self.root)

    def test_vfio_cleanup_removes_known_binaries_and_preserves_fixtures(self) -> None:
        build = self.root / "vfio-user" / "build"
        build.mkdir(parents=True)
        binary = build / "test_behavior_nvme"
        binary.write_text("binary", encoding="utf-8")
        fixture = build / "test-fixtures"
        fixture.mkdir()
        (fixture / "user-visible-artifact").write_text("fixture", encoding="utf-8")

        cleanup_vfio_build(self.root)

        self.assertFalse(binary.exists())
        self.assertTrue((fixture / "user-visible-artifact").exists())

    def test_vfio_cleanup_rejects_a_symlinked_build_directory(self) -> None:
        vfio = self.root / "vfio-user"
        vfio.mkdir()
        outside = Path(self.tempdir.name) / "outside"
        outside.mkdir()
        (vfio / "build").symlink_to(outside, target_is_directory=True)

        with self.assertRaisesRegex(CleanupError, "must not be a symlink"):
            cleanup_vfio_build(self.root)


if __name__ == "__main__":
    unittest.main()
