#!/usr/bin/env python3
"""Static contracts for hermetic VFIO-user unit-test fixtures."""

from __future__ import annotations

import unittest
from pathlib import Path


VFIO_ROOT = Path(__file__).resolve().parents[1]
MAKEFILE = (VFIO_ROOT / "Makefile").read_text(encoding="utf-8")
TESTS = VFIO_ROOT / "tests"
FIXTURE_DEPENDENT_C_TESTS = (
    "test_device_model.c",
    "test_behavior_nvme.c",
    "test_behavior_hda.c",
    "test_behavior_gpu.c",
    "test_behavior_storage.c",
)


class FixtureContractTests(unittest.TestCase):
    def test_makefile_owns_a_fixed_non_overridable_fixture_root(self) -> None:
        self.assertIn(
            "override VFIO_TEST_FIXTURE_ROOT := $(abspath $(CURDIR)/build/test-fixtures)",
            MAKEFILE,
        )
        self.assertIn("test-fixture-root-check", MAKEFILE)
        self.assertIn("test_fixture_root_guard.py", MAKEFILE)
        self.assertIn("--fixture-name", MAKEFILE)
        self.assertIn("--lib-dir \"$(ROOT_DIR)/lib/pcileech-fpga\"", MAKEFILE)
        self.assertIn("verify-manifest", MAKEFILE)

    def test_fixture_generators_cover_all_committed_donors(self) -> None:
        expected = {
            path.stem for path in (VFIO_ROOT.parent / "testdata" / "donors").glob("*.json")
        }
        declared = set(
            MAKEFILE.split("VFIO_TEST_FIXTURE_NAMES :=", 1)[1]
            .split("\n", 1)[0]
            .split()
        )
        self.assertEqual(declared, expected)

    def test_fixture_dependent_c_tests_do_not_fall_back_to_cocotb_outputs(self) -> None:
        for name in FIXTURE_DEPENDENT_C_TESTS:
            source = (TESTS / name).read_text(encoding="utf-8")
            self.assertIn('"fixture_path.h"', source)
            self.assertIn("vfio_test_fixture_path", source)
            self.assertNotIn("tests/cocotb/out", source)

    def test_server_cli_requires_explicit_isolated_artifacts(self) -> None:
        source = (TESTS / "test_server_cli.py").read_text(encoding="utf-8")
        self.assertIn('os.environ.get("VFIO_ARTIFACTS")', source)
        self.assertIn("must name an isolated generated fixture", source)
        self.assertNotIn("tests/cocotb/out", source)


if __name__ == "__main__":
    unittest.main()
