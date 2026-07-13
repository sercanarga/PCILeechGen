#!/usr/bin/env python3
"""Unit tests for the hermetic cocotb matrix runner."""

from __future__ import annotations

import tempfile
import unittest
from pathlib import Path

import run_matrix


TEST_DIR = Path(__file__).resolve().parent
ROOT = TEST_DIR.parents[1]


class TemporaryRepo:
    def __init__(self) -> None:
        self._temporary = tempfile.TemporaryDirectory(prefix="cocotb-runner-test-")
        self.root = (Path(self._temporary.name).resolve() / "repo")
        (self.root / "tests" / "cocotb").mkdir(parents=True)

    def close(self) -> None:
        self._temporary.cleanup()


class OutputGuardTests(unittest.TestCase):
    def setUp(self) -> None:
        self.repo = TemporaryRepo()

    def tearDown(self) -> None:
        self.repo.close()

    @property
    def output(self) -> Path:
        return self.repo.root / run_matrix.OUTPUT_RELATIVE

    def write_sentinel(self) -> None:
        (self.output / run_matrix.OUTPUT_SENTINEL).write_text(
            run_matrix.OUTPUT_SENTINEL_CONTENT, encoding="utf-8"
        )

    def test_only_exact_canonical_repo_relative_output_is_accepted(self) -> None:
        expected = run_matrix.validate_output_root(
            run_matrix.OUTPUT_RELATIVE, repo_root=self.repo.root
        )
        self.assertEqual(expected, self.output)

        unsafe = (
            Path("."),
            Path("tests"),
            Path("tests/cocotb"),
            Path("tests/cocotb/out_matrix/.."),
            Path("tests/cocotb/out_matrix/child"),
            self.output,
            Path("/"),
            Path.home(),
        )
        for candidate in unsafe:
            with self.subTest(candidate=candidate):
                with self.assertRaises(run_matrix.MatrixError):
                    run_matrix.validate_output_root(
                        candidate, repo_root=self.repo.root
                    )

    def test_output_symlink_is_refused_without_touching_target(self) -> None:
        with tempfile.TemporaryDirectory(prefix="cocotb-external-") as temp_dir:
            external = Path(temp_dir).resolve()
            marker = external / "user-data"
            marker.write_text("keep", encoding="utf-8")
            self.output.symlink_to(external, target_is_directory=True)

            with self.assertRaises(run_matrix.MatrixError):
                run_matrix.validate_output_root(
                    run_matrix.OUTPUT_RELATIVE, repo_root=self.repo.root
                )
            self.assertEqual(marker.read_text(encoding="utf-8"), "keep")

    def test_existing_unowned_output_is_refused_and_preserved(self) -> None:
        self.output.mkdir()
        marker = self.output / "user-data"
        marker.write_text("keep", encoding="utf-8")

        with self.assertRaises(run_matrix.MatrixError):
            run_matrix.prepare_output_root(
                run_matrix.OUTPUT_RELATIVE, repo_root=self.repo.root
            )
        self.assertEqual(marker.read_text(encoding="utf-8"), "keep")

    def test_invalid_sentinel_is_refused_and_preserved(self) -> None:
        self.output.mkdir()
        sentinel = self.output / run_matrix.OUTPUT_SENTINEL
        sentinel.write_text("not the ownership token\n", encoding="utf-8")
        marker = self.output / "generation.log"
        marker.write_text("keep", encoding="utf-8")

        with self.assertRaises(run_matrix.MatrixError):
            run_matrix.prepare_output_root(
                run_matrix.OUTPUT_RELATIVE, repo_root=self.repo.root
            )
        self.assertEqual(marker.read_text(encoding="utf-8"), "keep")

    def test_unknown_entry_blocks_owned_reset_before_any_deletion(self) -> None:
        self.output.mkdir()
        self.write_sentinel()
        known = self.output / "nvme"
        known.mkdir()
        (known / "simulation.log").write_text("old", encoding="utf-8")
        unknown = self.output / "user-data"
        unknown.write_text("keep", encoding="utf-8")

        with self.assertRaises(run_matrix.MatrixError):
            run_matrix.prepare_output_root(
                run_matrix.OUTPUT_RELATIVE, repo_root=self.repo.root
            )
        self.assertTrue(known.is_dir())
        self.assertEqual(unknown.read_text(encoding="utf-8"), "keep")

    def test_owned_reset_clears_only_known_children_and_keeps_sentinel(self) -> None:
        self.output.mkdir()
        self.write_sentinel()
        case_dir = self.output / "nvme"
        case_dir.mkdir()
        (case_dir / "simulation.log").write_text("old", encoding="utf-8")
        (self.output / "summary.tsv").write_text("old", encoding="utf-8")

        result = run_matrix.prepare_output_root(
            run_matrix.OUTPUT_RELATIVE, repo_root=self.repo.root
        )
        self.assertEqual(result, self.output)
        self.assertFalse(case_dir.exists())
        self.assertFalse((self.output / "summary.tsv").exists())
        self.assertEqual(
            (self.output / run_matrix.OUTPUT_SENTINEL).read_text(encoding="utf-8"),
            run_matrix.OUTPUT_SENTINEL_CONTENT,
        )

    def test_clean_requires_sentinel_and_preserves_unowned_directory(self) -> None:
        self.output.mkdir()
        marker = self.output / "user-data"
        marker.write_text("keep", encoding="utf-8")
        with self.assertRaises(run_matrix.MatrixError):
            run_matrix.clean_owned_output_root(
                run_matrix.OUTPUT_RELATIVE, repo_root=self.repo.root
            )
        self.assertEqual(marker.read_text(encoding="utf-8"), "keep")

    def test_clean_removes_only_validated_owned_output(self) -> None:
        self.output.mkdir()
        self.write_sentinel()
        case_dir = self.output / "audio"
        case_dir.mkdir()
        (case_dir / "simulation.log").write_text("old", encoding="utf-8")

        run_matrix.clean_owned_output_root(
            run_matrix.OUTPUT_RELATIVE, repo_root=self.repo.root
        )
        self.assertFalse(self.output.exists())

    def test_symlink_inside_owned_case_blocks_reset(self) -> None:
        self.output.mkdir()
        self.write_sentinel()
        case_dir = self.output / "gpu"
        case_dir.mkdir()
        with tempfile.TemporaryDirectory(prefix="cocotb-external-") as temp_dir:
            external = Path(temp_dir).resolve()
            marker = external / "keep"
            marker.write_text("keep", encoding="utf-8")
            (case_dir / "bridge").symlink_to(external, target_is_directory=True)
            with self.assertRaises(run_matrix.MatrixError):
                run_matrix.prepare_output_root(
                    run_matrix.OUTPUT_RELATIVE, repo_root=self.repo.root
                )
            self.assertEqual(marker.read_text(encoding="utf-8"), "keep")


class FixtureStagingTests(unittest.TestCase):
    def test_timeout_patch_only_mutates_regular_staged_copy(self) -> None:
        with tempfile.TemporaryDirectory(prefix="cocotb-stage-") as temp_dir:
            root = Path(temp_dir).resolve()
            source = root / "source"
            destination = root / "destination"
            source.mkdir()
            bridge = source / "pcileech_nvme_dma_bridge.sv"
            original = ".TIMEOUT_CYCLES ( 24'h4C4B40 )\n"
            bridge.write_text(original, encoding="utf-8")

            run_matrix.stage_fixture(source, destination)
            run_matrix.patch_nvme_timeout(destination, "000400")

            self.assertEqual(bridge.read_text(encoding="utf-8"), original)
            staged = destination / bridge.name
            self.assertFalse(staged.is_symlink())
            self.assertIn("24'h000400", staged.read_text(encoding="utf-8"))

    def test_fixture_symlink_is_rejected_without_mutating_external_target(self) -> None:
        with tempfile.TemporaryDirectory(prefix="cocotb-stage-") as temp_dir:
            root = Path(temp_dir).resolve()
            source = root / "source"
            destination = root / "destination"
            source.mkdir()
            external = root / "external.sv"
            original = ".TIMEOUT_CYCLES ( 24'h4C4B40 )\n"
            external.write_text(original, encoding="utf-8")
            (source / "pcileech_nvme_dma_bridge.sv").symlink_to(external)

            with self.assertRaises(run_matrix.MatrixError):
                run_matrix.stage_fixture(source, destination)
            self.assertEqual(external.read_text(encoding="utf-8"), original)
            self.assertFalse(destination.exists())


class ResultValidationTests(unittest.TestCase):
    def test_missing_and_malformed_result_files_fail(self) -> None:
        with tempfile.TemporaryDirectory(prefix="cocotb-results-") as temp_dir:
            result = Path(temp_dir) / "results.xml"
            with self.assertRaises(run_matrix.MatrixError):
                run_matrix.validate_results_xml(result)
            result.write_text("<testsuites>", encoding="utf-8")
            with self.assertRaises(run_matrix.MatrixError):
                run_matrix.validate_results_xml(result)

    def test_failure_node_fails_and_passing_xml_reports_count(self) -> None:
        with tempfile.TemporaryDirectory(prefix="cocotb-results-") as temp_dir:
            result = Path(temp_dir) / "results.xml"
            result.write_text(
                '<testsuites><testsuite><testcase name="bad">'
                '<failure message="boom"/></testcase></testsuite></testsuites>',
                encoding="utf-8",
            )
            with self.assertRaises(run_matrix.MatrixError):
                run_matrix.validate_results_xml(result)
            result.write_text(
                '<testsuites><testsuite><testcase name="a"/>'
                '<testcase name="b"/></testsuite></testsuites>',
                encoding="utf-8",
            )
            self.assertEqual(run_matrix.validate_results_xml(result), (2, 0))


class MatrixContractTests(unittest.TestCase):
    def test_matrix_is_the_required_eight_isolated_cases(self) -> None:
        self.assertEqual(
            [case.name for case in run_matrix.CASES],
            ["nvme", "generic", "audio", "xhci", "ethernet", "wifi", "sata", "gpu"],
        )
        self.assertEqual(len({case.name for case in run_matrix.CASES}), 8)

    def test_each_case_has_unique_fixture_build_result_and_log_paths(self) -> None:
        root = Path("tests/cocotb/out_matrix")
        all_paths: set[Path] = set()
        for case in run_matrix.CASES:
            case_dir, fixture, sim_build, results = run_matrix.case_paths(root, case)
            paths = {case_dir, fixture, sim_build, results, case_dir / "simulation.log"}
            self.assertEqual(len(paths), 5)
            self.assertTrue(all_paths.isdisjoint(paths))
            all_paths.update(paths)

    def test_supported_python_range_rejects_314(self) -> None:
        run_matrix.ensure_supported_python((3, 13, 9))
        with self.assertRaises(run_matrix.MatrixError):
            run_matrix.ensure_supported_python((3, 14, 0))

    def test_active_venv_tool_path_is_not_resolved(self) -> None:
        invoked = Path("/repo/bin/cocotb-venv/bin/python")
        python, config = run_matrix.active_venv_tools(invoked)
        self.assertEqual(python, invoked)
        self.assertEqual(config, invoked.parent / "cocotb-config")

    def test_venv_guard_accepts_only_exact_repo_relative_path(self) -> None:
        repo = TemporaryRepo()
        try:
            self.assertEqual(
                run_matrix.validate_venv_path(
                    run_matrix.VENV_RELATIVE, repo_root=repo.root
                ),
                repo.root / run_matrix.VENV_RELATIVE,
            )
            for unsafe in (Path("bin"), repo.root / "bin/cocotb-venv", Path("/tmp/venv")):
                with self.assertRaises(run_matrix.MatrixError):
                    run_matrix.validate_venv_path(unsafe, repo_root=repo.root)
        finally:
            repo.close()

    def test_venv_guard_rejects_symlink_destination(self) -> None:
        repo = TemporaryRepo()
        try:
            (repo.root / "bin").mkdir()
            with tempfile.TemporaryDirectory(prefix="cocotb-external-venv-") as temp_dir:
                (repo.root / run_matrix.VENV_RELATIVE).symlink_to(
                    Path(temp_dir).resolve(), target_is_directory=True
                )
                with self.assertRaises(run_matrix.MatrixError):
                    run_matrix.validate_venv_path(
                        run_matrix.VENV_RELATIVE, repo_root=repo.root
                    )
        finally:
            repo.close()


class SourceContractTests(unittest.TestCase):
    def test_cocotb_makefile_has_no_parse_time_shell_or_ambient_pythonpath(self) -> None:
        source = (TEST_DIR / "Makefile").read_text(encoding="utf-8")
        self.assertNotIn("$(shell", source)
        self.assertNotIn("MODULE ?=", source)
        self.assertIn("COCOTB_TEST_MODULES", source)
        self.assertIn("PYTHONPATH := $(COCOTB_TEST_DIR)", source)
        self.assertNotIn("override export PYTHONPATH", source)
        self.assertIn("COCOTB_RESULTS_FILE", source)
        self.assertIn("prepare-config", source)

    def test_generate_script_requires_external_paths_and_has_no_recursive_delete(self) -> None:
        source = (TEST_DIR / "generate_all.sh").read_text(encoding="utf-8")
        self.assertIn("--generator", source)
        self.assertIn("--output-root", source)
        self.assertNotIn("rm -rf", source)
        self.assertNotIn("tests/cocotb/out_", source)
        self.assertNotIn(">/dev/null 2>&1", source)
        self.assertIn("--skip-vivado --force", source)

    def test_requirements_are_pinned(self) -> None:
        requirements = (TEST_DIR / "requirements.txt").read_text(encoding="utf-8").splitlines()
        self.assertEqual(requirements, ["cocotb==2.0.1", "find_libpython==0.5.1"])

    def test_root_makefile_checks_interpreters_before_pip_and_uses_safe_clean(self) -> None:
        source = (ROOT / "Makefile").read_text(encoding="utf-8")
        self.assertIn("COCOTB_PYTHON ?= python3.13", source)
        self.assertIn("--check-venv", source)
        self.assertIn("--clean --output-root", source)
        self.assertNotIn('rm -rf "$(COCOTB_OUTPUT_ROOT)"', source)
        self.assertLess(source.index("--check-venv"), source.index("-m venv"))
        self.assertLess(source.index("--check-python", source.index("-m venv")), source.index("-m pip"))

    def test_default_output_is_ignored(self) -> None:
        source = (ROOT / ".gitignore").read_text(encoding="utf-8")
        self.assertIn("tests/cocotb/out_*/", source)

    def test_runner_retains_generation_and_simulation_diagnostics(self) -> None:
        source = (TEST_DIR / "run_matrix.py").read_text(encoding="utf-8")
        self.assertIn('output_root / "generation.log"', source)
        self.assertIn('case_dir / "simulation.log"', source)
        self.assertNotIn("subprocess.DEVNULL", source)


if __name__ == "__main__":
    unittest.main(verbosity=2)
