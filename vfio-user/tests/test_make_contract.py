#!/usr/bin/env python3
"""Static contracts for the required, hermetic libFuzzer gate."""

from __future__ import annotations

import unittest
from pathlib import Path


VFIO_ROOT = Path(__file__).resolve().parents[1]
MAKEFILE = (VFIO_ROOT / "Makefile").read_text(encoding="utf-8")
HARNESS = (VFIO_ROOT / "tests" / "fuzz_behavior.c").read_text(encoding="utf-8")
CORPUS = VFIO_ROOT / "tests" / "fuzz-corpus"


class FuzzMakeContractTests(unittest.TestCase):
    def test_required_gate_creates_summary_before_fixture_generation(self) -> None:
        target = MAKEFILE.split("fuzz-ci:", 1)[1].split("\n\n", 1)[0]
        self.assertIn("FUZZ_SUMMARY", target)
        self.assertIn("fuzz-fixtures", target)
        self.assertLess(target.index("FUZZ_SUMMARY"), target.index("fuzz-fixtures"))

    def test_gate_requires_compiler_and_uses_isolated_fixtures(self) -> None:
        self.assertIn("FUZZ_CC ?= clang", MAKEFILE)
        self.assertIn('command -v "$(FUZZ_CC)"', MAKEFILE)
        self.assertIn("FUZZ_FIXTURE_ROOT", MAKEFILE)
        self.assertIn("--board ac701_ft601", MAKEFILE)
        self.assertIn("--board PCIeSquirrel", MAKEFILE)
        self.assertIn('--lib-dir "$(ROOT_DIR)/lib/pcileech-fpga"', MAKEFILE)
        self.assertIn("verify-manifest", MAKEFILE)
        self.assertIn("fuzz-fixture-root-check", MAKEFILE)
        self.assertIn("fuzz_fixture_guard.py", MAKEFILE)

    def test_gate_copies_only_committed_source_seeds(self) -> None:
        self.assertIn("tests/fuzz-corpus/*.bin", MAKEFILE)
        self.assertIn("fuzz-corpus.XXXXXX", MAKEFILE)
        self.assertIn("mktemp -d", MAKEFILE)
        self.assertEqual(sorted(path.name for path in CORPUS.glob("*.bin")), ["ethernet.bin", "nvme.bin"])

    def test_harness_uses_fixture_environment_not_cocotb_outputs(self) -> None:
        self.assertIn('getenv("VFIO_FUZZ_FIXTURE_ROOT")', HARNESS)
        self.assertNotIn("tests/cocotb/out", HARNESS)
        self.assertIn("static struct fuzz_fixture", HARNESS)
        self.assertIn("fixture_reset", HARNESS)

    def test_gate_fails_on_sanitizer_or_crash_artifacts(self) -> None:
        self.assertIn("AddressSanitizer", MAKEFILE)
        self.assertIn("UndefinedBehaviorSanitizer", MAKEFILE)
        self.assertIn("crash-*", MAKEFILE)
        self.assertIn("timeout-*", MAKEFILE)
        self.assertIn("LSAN_OPTIONS", MAKEFILE)
        self.assertIn("lsan_json_c.supp", MAKEFILE)

    def test_lsan_suppression_is_narrowly_scoped_to_json_c_parser(self) -> None:
        suppression = (VFIO_ROOT / "tests" / "lsan_json_c.supp").read_text(encoding="utf-8")
        self.assertIn("leak:json_tokener_parse_ex", suppression)
        self.assertIn("leak:StartRssThread", suppression)
        self.assertNotIn("leak:*", suppression)


if __name__ == "__main__":
    unittest.main()
