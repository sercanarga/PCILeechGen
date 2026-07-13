import contextlib
import importlib.util
import io
import json
import sys
import tempfile
import unittest
from pathlib import Path
from types import SimpleNamespace
from unittest import mock


MODULE_PATH = Path(__file__).resolve().parents[1] / "matrix.py"


def load_matrix():
    spec = importlib.util.spec_from_file_location("vfio_matrix_authority", MODULE_PATH)
    module = importlib.util.module_from_spec(spec)
    sys.modules[spec.name] = module
    spec.loader.exec_module(module)
    return module


def fixture_contract(case):
    fixture = json.loads(case.fixture.read_text(encoding="utf-8"))
    device = fixture["device"]
    return {
        "identity": {
            "vendor": f"{device['vendor_id']:04x}",
            "device": f"{device['device_id']:04x}",
            "class": f"{device['class_code']:06x}",
        },
        "bars": [
            {
                "bir": bar["index"],
                "size": bar["size"],
                "type": bar["type"],
                "prefetchable": bar["prefetchable"],
                "address_width": 64 if bar["is_64bit"] else 32,
            }
            for bar in fixture["bars"]
        ],
    }


def guest_log(case, *, status="pass", detail="enumeration", bdf="0000:03:00.0",
              vendor=None, device=None, class_code=None, bars=None, driver="none",
              driver_status=None):
    contract = fixture_contract(case)
    identity = contract["identity"]
    if bars is None:
        bars = len(contract["bars"])
    bar_count = len(bars) if isinstance(bars, list) else bars
    if driver_status is None:
        driver_status = status
    records = [
        {"event": "stage", "case": case.name, "stage": "enumerate", "bdf": bdf},
        {"event": "stage", "case": case.name, "stage": "bars", "count": bar_count},
        {
            "event": "stage",
            "case": case.name,
            "stage": "driver",
            "driver": driver,
            "status": driver_status,
        },
        {
            "event": "result",
            "case": case.name,
            "status": status,
            "bdf": bdf,
            "vendor": identity["vendor"] if vendor is None else vendor,
            "device": identity["device"] if device is None else device,
            "class": identity["class"] if class_code is None else class_code,
            "driver": driver,
            "bars": bars,
            "detail": detail,
        },
    ]
    return "".join(json.dumps(record) + "\n" for record in records)


def write_device_model(path, case):
    contract = fixture_contract(case)
    identity = contract["identity"]
    model = {
        "functions": [
            {
                "vendor_id": int(identity["vendor"], 16),
                "device_id": int(identity["device"], 16),
                "class_code": int(identity["class"], 16),
            }
        ],
        "bars": contract["bars"],
        "capabilities": [{"id": 1}],
    }
    path.mkdir(parents=True, exist_ok=True)
    (path / "device_model.json").write_text(json.dumps(model), encoding="utf-8")


class GuestEvidenceAuthorityTests(unittest.TestCase):
    def setUp(self):
        self.matrix = load_matrix()

    def test_terminal_result_requires_ordered_stage_evidence(self):
        case = self.matrix.CASES["generic"]
        terminal_only = guest_log(case).splitlines()[-1] + "\n"

        with self.assertRaisesRegex(self.matrix.CaseFailure, "stage evidence"):
            self.matrix.parse_guest_results(terminal_only, case)

    def test_identity_is_checked_against_case_contract(self):
        case = self.matrix.CASES["generic"]
        for field, kwargs in (
            ("vendor", {"vendor": "ffff"}),
            ("device", {"device": "ffff"}),
            ("class", {"class_code": "ffffff"}),
        ):
            with self.subTest(field=field):
                with self.assertRaisesRegex(self.matrix.CaseFailure, f"{field} mismatch"):
                    self.matrix.parse_guest_results(guest_log(case, **kwargs), case)

    def test_bdf_must_be_canonical_and_match_enumeration_stage(self):
        case = self.matrix.CASES["generic"]
        with self.assertRaisesRegex(self.matrix.CaseFailure, "invalid guest BDF"):
            self.matrix.parse_guest_results(guest_log(case, bdf="not-a-bdf"), case)

        records = [json.loads(line) for line in guest_log(case).splitlines()]
        records[-1]["bdf"] = "0000:04:00.0"
        text = "".join(json.dumps(record) + "\n" for record in records)
        with self.assertRaisesRegex(self.matrix.CaseFailure, "BDF.*stage"):
            self.matrix.parse_guest_results(text, case)

    def test_bar_count_must_match_stage_and_all_generated_birs(self):
        case = self.matrix.CASES["multibar"]
        with self.assertRaisesRegex(self.matrix.CaseFailure, "BAR count mismatch"):
            self.matrix.parse_guest_results(guest_log(case, bars=2), case)

        records = [json.loads(line) for line in guest_log(case).splitlines()]
        records[1]["count"] = 2
        text = "".join(json.dumps(record) + "\n" for record in records)
        with self.assertRaisesRegex(self.matrix.CaseFailure, "BAR.*stage"):
            self.matrix.parse_guest_results(text, case)

    def test_indexed_bar_evidence_rejects_wrong_size(self):
        case = self.matrix.CASES["multibar"]
        bars = fixture_contract(case)["bars"]
        self.matrix.parse_guest_results(guest_log(case, bars=bars), case)

        corrupt = [dict(bar) for bar in bars]
        corrupt[1]["size"] *= 2
        with self.assertRaisesRegex(self.matrix.CaseFailure, "BAR2 size mismatch"):
            self.matrix.parse_guest_results(guest_log(case, bars=corrupt), case)

    def test_missing_duplicate_and_corrupt_terminal_records_fail(self):
        case = self.matrix.CASES["generic"]
        stages = "\n".join(guest_log(case).splitlines()[:-1]) + "\n"
        with self.assertRaisesRegex(self.matrix.CaseFailure, "terminal guest result, got 0"):
            self.matrix.parse_guest_results(stages, case)

        terminal = guest_log(case).splitlines()[-1]
        with self.assertRaisesRegex(self.matrix.CaseFailure, "terminal guest result, got 2"):
            self.matrix.parse_guest_results(guest_log(case) + terminal + "\n", case)

        with self.assertRaisesRegex(self.matrix.CaseFailure, "invalid guest result JSON"):
            self.matrix.parse_guest_results(stages + "{corrupt-json\n", case)

    def test_driver_stage_and_terminal_must_agree(self):
        case = self.matrix.CASES["audio"]
        records = [json.loads(line) for line in guest_log(case, driver="snd_hda_intel").splitlines()]
        records[2]["driver"] = "other"
        text = "".join(json.dumps(record) + "\n" for record in records)

        with self.assertRaisesRegex(self.matrix.CaseFailure, "driver.*stage"):
            self.matrix.parse_guest_results(text, case)

    def test_fatal_kernel_message_after_bind_invalidates_pass(self):
        case = self.matrix.CASES["audio"]
        lines = guest_log(case, driver="snd_hda_intel").splitlines()
        lines.insert(3, "[    2.1] BUG: unable to handle kernel paging request")

        with self.assertRaisesRegex(self.matrix.CaseFailure, "fatal guest kernel"):
            self.matrix.parse_guest_results("\n".join(lines) + "\n", case)

    def test_only_spec_permitted_guest_dependency_skip_is_accepted(self):
        allowed = self.matrix.CASES["audio"]
        result = self.matrix.parse_guest_results(
            guest_log(allowed, status="skip", detail="optional-driver-not-bound"),
            allowed,
        )
        self.assertEqual(result.status, "skip")

        mandatory = self.matrix.CASES["ethernet"]
        with self.assertRaisesRegex(self.matrix.CaseFailure, "skip reason is not permitted"):
            self.matrix.parse_guest_results(
                guest_log(mandatory, status="skip", detail="optional-driver-not-bound"),
                mandatory,
            )

        with self.assertRaisesRegex(self.matrix.CaseFailure, "skip reason is not permitted"):
            self.matrix.parse_guest_results(
                guest_log(allowed, status="skip", detail="device-not-found"),
                allowed,
            )


class PipelineAuthorityTests(unittest.TestCase):
    def setUp(self):
        self.matrix = load_matrix()

    def test_case_blocked_rejects_unlisted_reason(self):
        with self.assertRaisesRegex(ValueError, "not allowlisted"):
            self.matrix.CaseBlocked("pipeline-error", "must fail")

    def test_qemu_nonzero_exit_is_failure_even_with_pass_record(self):
        case = self.matrix.CASES["generic"]
        with tempfile.TemporaryDirectory() as tmp:
            root = Path(tmp)
            artifacts = root / "artifacts"
            write_device_model(artifacts, case)

            def fake_run(_command, *, stdout, **_kwargs):
                stdout.write(guest_log(case))
                return SimpleNamespace(returncode=9)

            with (
                mock.patch.object(
                    self.matrix,
                    "start_server",
                    return_value=(
                        object(),
                        object(),
                        root / "device.sock",
                        {
                            "event": "ready",
                            "vendor_id": "1234",
                            "device_id": "5678",
                            "class_code": "000000",
                            "bar_count": 1,
                        },
                    ),
                ),
                mock.patch.object(self.matrix, "stop_server"),
                mock.patch.object(self.matrix.subprocess, "run", side_effect=fake_run),
            ):
                with self.assertRaisesRegex(self.matrix.CaseFailure, "QEMU exited 9"):
                    self.matrix.run_qemu_case(
                        case,
                        artifacts,
                        root,
                        Path("qemu-system-aarch64"),
                        Path("kernel"),
                        Path("initrd"),
                    )

            self.assertTrue((root / "guest-results.jsonl").is_file())

    def test_qemu_version_probe_error_is_failure_not_blocked(self):
        case = self.matrix.CASES["generic"]
        with tempfile.TemporaryDirectory() as tmp:
            root = Path(tmp)
            artifacts = root / "artifacts"
            write_device_model(artifacts, case)
            with mock.patch.object(
                self.matrix.subprocess,
                "run",
                return_value=SimpleNamespace(returncode=2, stdout="", stderr="broken"),
            ):
                with self.assertRaises(self.matrix.CaseFailure):
                    self.matrix.run_qemu_case(
                        case,
                        artifacts,
                        root,
                        Path("qemu-system-aarch64"),
                        Path("kernel"),
                        Path("initrd"),
                        expected_version="9.2",
                    )

    def test_invalid_readiness_is_failure_and_still_stops_server(self):
        case = self.matrix.CASES["generic"]
        with tempfile.TemporaryDirectory() as tmp:
            root = Path(tmp)
            artifacts = root / "artifacts"
            write_device_model(artifacts, case)
            process = object()
            output = object()
            with (
                mock.patch.object(
                    self.matrix,
                    "start_server",
                    return_value=(
                        process,
                        output,
                        root / "device.sock",
                        {"event": "ready", "vendor_id": "ffff"},
                    ),
                ),
                mock.patch.object(self.matrix, "stop_server") as stop_server,
            ):
                with self.assertRaises(self.matrix.CaseFailure):
                    self.matrix.run_qemu_case(
                        case,
                        artifacts,
                        root,
                        Path("qemu-system-aarch64"),
                        Path("kernel"),
                        Path("initrd"),
                    )

            stop_server.assert_called_once_with(process, output)

    def test_rebind_requires_positive_reset_and_rebind_evidence(self):
        case = self.matrix.CASES["audio"]
        result = self.matrix.parse_guest_results(
            guest_log(case, driver="snd_hda_intel", detail="driver-bound"),
            case,
        )
        with self.assertRaisesRegex(self.matrix.CaseFailure, "rebind evidence"):
            self.matrix.validate_guest_result(result, case, rebind=True)

        result = self.matrix.parse_guest_results(
            guest_log(
                case,
                driver="snd_hda_intel",
                detail="rebind-reset-driver-bound",
            ),
            case,
        )
        self.matrix.validate_guest_result(result, case, rebind=True)

    def test_native_nvme_kvm_gate_does_not_misclassify_cross_arch_tcg(self):
        case = self.matrix.CASES["nvme"]
        missing = Path("/definitely/missing/kvm")
        self.assertTrue(
            self.matrix.qemu_requires_kvm(
                case,
                missing,
                qemu=Path("qemu-system-x86_64"),
                host_machine="x86_64",
                host_system="Linux",
            )
        )
        self.assertFalse(
            self.matrix.qemu_requires_kvm(
                case,
                missing,
                qemu=Path("qemu-system-aarch64"),
                host_machine="x86_64",
                host_system="Linux",
            )
        )
        self.assertFalse(
            self.matrix.qemu_requires_kvm(
                case,
                missing,
                qemu=Path("qemu-system-x86_64"),
                host_machine="x86_64",
                host_system="Darwin",
            )
        )

    def test_pipeline_failure_wins_over_blocked_and_summary_exit_is_authoritative(self):
        with tempfile.TemporaryDirectory() as tmp:
            work_dir = Path(tmp)
            cases = {
                "audio": self.matrix.CASES["audio"],
                "generic": self.matrix.CASES["generic"],
            }
            results = [
                {
                    "case": "audio",
                    "status": "blocked",
                    "reason": "guest-optional-driver",
                    "detail": "optional-driver-not-bound",
                },
                {"case": "generic", "status": "fail", "detail": "generation failed"},
            ]
            with (
                mock.patch.object(self.matrix, "CASES", cases),
                mock.patch.object(self.matrix, "run_case", side_effect=results),
                contextlib.redirect_stdout(io.StringIO()),
            ):
                status = self.matrix.main(["--all", "--work-dir", str(work_dir)])

            self.assertEqual(status, 1)
            summary = json.loads((work_dir / "summary.json").read_text(encoding="utf-8"))
            self.assertEqual(summary["status"], "fail")
            self.assertEqual(summary["exit_code"], status)
            self.assertEqual(summary["counts"], {"pass": 0, "blocked": 1, "fail": 1})
            self.assertEqual([item["case"] for item in summary["results"]], list(cases))

    def test_run_case_never_converts_pipeline_failure_to_blocked(self):
        case = self.matrix.CASES["generic"]
        with tempfile.TemporaryDirectory() as tmp:
            root = Path(tmp)
            with mock.patch.object(
                self.matrix,
                "run_command",
                side_effect=self.matrix.CaseFailure("generation stage failed"),
            ):
                result = self.matrix.run_case(
                    case,
                    root,
                    qemu=Path("qemu-system-aarch64"),
                    kernel=Path("kernel"),
                    initrd=Path("initrd"),
                )

            self.assertEqual(result["status"], "fail")
            self.assertIn("generation stage failed", result["detail"])
            persisted = json.loads(
                (root / case.name / "result.json").read_text(encoding="utf-8")
            )
            self.assertEqual(persisted, result)

    def test_missing_qemu_guest_evidence_is_failure(self):
        case = self.matrix.CASES["generic"]
        with tempfile.TemporaryDirectory() as tmp:
            with mock.patch.object(self.matrix, "run_command") as run_command:
                result = self.matrix.run_case(case, Path(tmp))

        self.assertEqual(result["status"], "fail")
        self.assertIn("guest evidence is required", result["detail"])
        run_command.assert_not_called()

    def test_summary_rejects_forged_blocked_result(self):
        summary = self.matrix.summarize_results(
            [
                {
                    "case": "generic",
                    "status": "blocked",
                    "reason": "pipeline-error",
                    "detail": "forged",
                }
            ]
        )

        self.assertEqual(summary["status"], "fail")
        self.assertEqual(summary["exit_code"], 1)
        self.assertEqual(summary["results"][0]["status"], "fail")

    def test_allowlisted_external_block_keeps_matrix_successful(self):
        summary = self.matrix.summarize_results(
            [
                {"case": "generic", "status": "pass"},
                {
                    "case": "nvme",
                    "status": "blocked",
                    "reason": "kvm-unavailable",
                    "detail": "/dev/kvm is unavailable",
                },
            ]
        )

        self.assertEqual(summary["status"], "pass")
        self.assertEqual(summary["exit_code"], 0)
        self.assertEqual(summary["counts"], {"pass": 1, "blocked": 1, "fail": 0})


if __name__ == "__main__":
    unittest.main()
