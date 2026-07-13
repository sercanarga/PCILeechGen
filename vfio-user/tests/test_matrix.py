import concurrent.futures
import io
import importlib.util
import json
import os
import socket
import stat
import sys
import tempfile
import unittest
from pathlib import Path
from unittest import mock


MODULE_PATH = Path(__file__).resolve().parents[1] / "matrix.py"


def load_matrix():
    spec = importlib.util.spec_from_file_location("vfio_matrix", MODULE_PATH)
    module = importlib.util.module_from_spec(spec)
    sys.modules[spec.name] = module
    spec.loader.exec_module(module)
    return module


def generic_guest_log(kernel_line=""):
    lines = [
        '{"event":"stage","case":"generic","stage":"enumerate","bdf":"0000:03:00.0"}',
        '{"event":"stage","case":"generic","stage":"bars","count":1}',
        '{"event":"stage","case":"generic","stage":"driver","driver":"none","status":"pass"}',
    ]
    if kernel_line:
        lines.append(kernel_line)
    lines.append(
        '{"event":"result","case":"generic","status":"pass",'
        '"bdf":"0000:03:00.0","vendor":"1234","device":"5678",'
        '"class":"000000","driver":"none","bars":1,"detail":"enumeration"}'
    )
    return "\n".join(lines) + "\n"


class FakeServerProcess:
    def __init__(self, stdout=None, on_terminate=None):
        self.stdout = stdout if stdout is not None else io.StringIO()
        self.returncode = None
        self.on_terminate = on_terminate

    def poll(self):
        return self.returncode

    def terminate(self):
        if self.on_terminate is not None:
            self.on_terminate()
        self.returncode = -15

    def wait(self, timeout=None):
        del timeout
        if self.returncode is None:
            self.returncode = 0
        return self.returncode

    def kill(self):
        self.returncode = -9


class MatrixTests(unittest.TestCase):
    def test_matrix_covers_every_generated_fixture(self):
        matrix = load_matrix()

        self.assertEqual(
            set(matrix.CASES),
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
            },
        )
        self.assertEqual(matrix.CASES["nvme"].board, "ac701_ft601")
        self.assertTrue(
            all(case.board == "PCIeSquirrel" for name, case in matrix.CASES.items() if name != "nvme")
        )

    def test_run_command_rejects_nonzero_exit(self):
        matrix = load_matrix()

        with tempfile.TemporaryDirectory() as tmp:
            with self.assertRaisesRegex(matrix.CaseFailure, "exit status 7"):
                matrix.run_command(
                    [sys.executable, "-c", "raise SystemExit(7)"],
                    Path(tmp) / "command.log",
                    timeout=5,
                )

    def test_work_root_rejects_symlink_without_touching_target(self):
        matrix = load_matrix()

        with tempfile.TemporaryDirectory() as tmp:
            root = Path(tmp)
            target = root / "target"
            target.mkdir()
            valuable = target / "valuable.txt"
            valuable.write_text("preserve", encoding="utf-8")
            work_root = root / "matrix"
            work_root.symlink_to(target, target_is_directory=True)

            with self.assertRaisesRegex(matrix.CaseFailure, "must not be a symlink"):
                matrix.prepare_work_root(work_root)

            self.assertEqual(valuable.read_text(encoding="utf-8"), "preserve")
            self.assertFalse((target / matrix.WORK_ROOT_SENTINEL).exists())

    def test_work_root_refuses_unowned_existing_tree(self):
        matrix = load_matrix()

        with tempfile.TemporaryDirectory() as tmp:
            work_root = Path(tmp) / "matrix"
            work_root.mkdir(mode=0o700)
            valuable = work_root / "valuable.txt"
            valuable.write_text("preserve", encoding="utf-8")

            with self.assertRaisesRegex(matrix.CaseFailure, "unowned, non-empty"):
                matrix.prepare_work_root(work_root)

            self.assertEqual(valuable.read_text(encoding="utf-8"), "preserve")
            self.assertFalse((work_root / matrix.WORK_ROOT_SENTINEL).exists())

    def test_owned_case_rerun_clears_only_the_owned_case(self):
        matrix = load_matrix()

        with tempfile.TemporaryDirectory() as tmp:
            work_root = matrix.prepare_work_root(Path(tmp) / "matrix")
            case_dir = matrix.prepare_case_directory(
                work_root,
                matrix.CASES["generic"],
            )
            nested = case_dir / "artifacts" / "nested"
            nested.mkdir(parents=True)
            (nested / "generated.bin").write_bytes(b"generated")
            summary = work_root / "summary.json"
            matrix._atomic_write(summary, "{}\n")

            rerun = matrix.prepare_case_directory(
                work_root,
                matrix.CASES["generic"],
            )

            self.assertEqual(rerun, case_dir)
            self.assertEqual(
                {entry.name for entry in case_dir.iterdir()},
                {matrix.CASE_SENTINEL},
            )
            self.assertEqual(summary.read_text(encoding="utf-8"), "{}\n")
            self.assertEqual(stat.S_IMODE(work_root.stat().st_mode), 0o700)
            self.assertEqual(stat.S_IMODE(case_dir.stat().st_mode), 0o700)

    def test_owned_case_rejects_symlink_before_cleanup(self):
        matrix = load_matrix()

        with tempfile.TemporaryDirectory() as tmp:
            root = Path(tmp)
            work_root = matrix.prepare_work_root(root / "matrix")
            case_dir = matrix.prepare_case_directory(
                work_root,
                matrix.CASES["generic"],
            )
            valuable = root / "valuable.txt"
            valuable.write_text("preserve", encoding="utf-8")
            linked_log = case_dir / "generation.log"
            linked_log.symlink_to(valuable)

            with self.assertRaisesRegex(matrix.CaseFailure, "symlink.*not allowed"):
                matrix.prepare_case_directory(
                    work_root,
                    matrix.CASES["generic"],
                )

            self.assertEqual(valuable.read_text(encoding="utf-8"), "preserve")
            self.assertTrue(linked_log.is_symlink())

    def test_run_command_rejects_symlink_log_without_truncating_target(self):
        matrix = load_matrix()

        with tempfile.TemporaryDirectory() as tmp:
            root = Path(tmp)
            logs = root / "logs"
            logs.mkdir(mode=0o700)
            valuable = root / "valuable.txt"
            valuable.write_text("preserve", encoding="utf-8")
            log_path = logs / "command.log"
            log_path.symlink_to(valuable)

            with self.assertRaisesRegex(matrix.CaseFailure, "refusing symlink"):
                matrix.run_command(
                    [sys.executable, "-c", "print('must not run')"],
                    log_path,
                    timeout=5,
                )

            self.assertEqual(valuable.read_text(encoding="utf-8"), "preserve")
            self.assertTrue(log_path.is_symlink())

    def test_run_command_replaces_existing_hardlink_without_truncating_source(self):
        matrix = load_matrix()

        with tempfile.TemporaryDirectory() as tmp:
            root = Path(tmp)
            logs = root / "logs"
            logs.mkdir(mode=0o700)
            valuable = root / "valuable.txt"
            valuable.write_text("preserve", encoding="utf-8")
            log_path = logs / "command.log"
            os.link(valuable, log_path)

            matrix.run_command(
                [sys.executable, "-c", "print('new log')"],
                log_path,
                timeout=5,
            )

            self.assertEqual(valuable.read_text(encoding="utf-8"), "preserve")
            self.assertEqual(log_path.read_text(encoding="utf-8"), "new log\n")
            self.assertNotEqual(valuable.stat().st_ino, log_path.stat().st_ino)

    def test_atomic_result_write_rejects_symlink_target(self):
        matrix = load_matrix()

        with tempfile.TemporaryDirectory() as tmp:
            root = Path(tmp)
            work_root = matrix.prepare_work_root(root / "matrix")
            valuable = root / "valuable.txt"
            valuable.write_text("preserve", encoding="utf-8")
            summary = work_root / "summary.json"
            summary.symlink_to(valuable)

            with self.assertRaisesRegex(matrix.CaseFailure, "refusing symlink"):
                matrix._atomic_write(summary, "unsafe\n")

            self.assertEqual(valuable.read_text(encoding="utf-8"), "preserve")
            self.assertTrue(summary.is_symlink())

    def test_parse_guest_results_requires_passing_terminal_record(self):
        matrix = load_matrix()

        result = matrix.parse_guest_results(
            generic_guest_log(),
            matrix.CASES["generic"],
        )

        self.assertEqual(result.status, "pass")
        self.assertEqual(result.bdf, "0000:03:00.0")

    def test_parse_guest_results_rejects_wrong_case(self):
        matrix = load_matrix()

        with self.assertRaisesRegex(matrix.CaseFailure, "case mismatch"):
            matrix.parse_guest_results(
                '{"event":"result","case":"nvme","status":"pass"}\n',
                matrix.CASES["generic"],
            )

    def test_parse_guest_results_ignores_kernel_log_lines(self):
        matrix = load_matrix()
        result = matrix.parse_guest_results(
            generic_guest_log("[    1.2] ordinary kernel message"),
            matrix.CASES["generic"],
        )
        self.assertEqual(result.status, "pass")

    def test_build_command_always_skips_vivado(self):
        matrix = load_matrix()

        command = matrix.build_command(matrix.CASES["nvme"], Path("/tmp/work"))

        self.assertIn("--skip-vivado", command)
        self.assertIn("--from-json", command)
        self.assertEqual(command[command.index("--board") + 1], "ac701_ft601")

    def test_contract_covers_generated_artifact_stages(self):
        matrix = load_matrix()
        with tempfile.TemporaryDirectory() as tmp:
            root = Path(tmp)
            for case in matrix.CASES.values():
                fixture = json.loads(case.fixture.read_text(encoding="utf-8"))
                device = fixture["device"]
                artifacts = root / case.name
                artifacts.mkdir()
                model = {
                    "functions": [
                        {
                            "vendor_id": device["vendor_id"],
                            "device_id": device["device_id"],
                            "class_code": device["class_code"],
                        }
                    ],
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
                    "capabilities": fixture["capabilities"],
                }
                (artifacts / "device_model.json").write_text(
                    json.dumps(model), encoding="utf-8"
                )

                contract = matrix.build_contract(case, artifacts)
                self.assertEqual(contract["case"], case.name)
                self.assertTrue(contract["bars"])
                self.assertTrue(contract["capabilities"])
                self.assertEqual(contract["probe"][:3], ["enumerate", "bars", "reset"])

    def test_nvme_qemu_requires_kvm(self):
        matrix = load_matrix()

        self.assertTrue(
            matrix.qemu_requires_kvm(
                matrix.CASES["nvme"],
                Path("/missing/kvm"),
                qemu=Path("qemu-system-x86_64"),
                host_machine="x86_64",
                host_system="Linux",
            )
        )
        self.assertFalse(
            matrix.qemu_requires_kvm(
                matrix.CASES["generic"],
                Path("/missing/kvm"),
                qemu=Path("qemu-system-x86_64"),
                host_machine="x86_64",
                host_system="Linux",
            )
        )

    def test_qemu_rebind_mode_is_explicit(self):
        matrix = load_matrix()
        with tempfile.TemporaryDirectory() as tmp:
            artifacts = Path(tmp)
            (artifacts / "device_model.json").write_text(
                '{"functions":[{"vendor_id":4660,"device_id":22136}]}\n',
                encoding="utf-8",
            )
            command = matrix.build_qemu_command(
                matrix.CASES["sata"], Path("/tmp/work/device.sock"),
                Path("/tmp/kernel"), Path("/tmp/initrd"), artifacts,
                rebind=True,
            )

        append = command[command.index("-append") + 1]
        self.assertIn("vfio_rebind=1", append)

    def test_socket_leases_are_private_unique_and_never_reuse_a_stale_path(self):
        matrix = load_matrix()
        with tempfile.TemporaryDirectory() as tmp:
            work_dir = Path(tmp)
            stale_path = work_dir / "device.sock"
            stale_listener = socket.socket(socket.AF_UNIX)
            stale_listener.bind(str(stale_path))
            stale_listener.close()
            try:
                with concurrent.futures.ThreadPoolExecutor(max_workers=8) as executor:
                    leases = list(
                        executor.map(
                            lambda _index: matrix.allocate_socket_lease(
                                matrix.CASES["generic"], work_dir
                            ),
                            range(16),
                        )
                    )

                self.assertEqual(len({lease.path for lease in leases}), len(leases))
                self.assertTrue(os.path.lexists(stale_path))
                for lease in leases:
                    self.assertFalse(os.path.lexists(lease.path))
                    self.assertTrue(stat.S_ISDIR(lease.directory.lstat().st_mode))
                    self.assertEqual(stat.S_IMODE(lease.directory.lstat().st_mode), 0o700)
                    lease.directory.rmdir()
            finally:
                stale_path.unlink()

    def test_readiness_socket_evidence_rejects_non_socket(self):
        matrix = load_matrix()
        with tempfile.TemporaryDirectory() as tmp:
            lease = matrix.allocate_socket_lease(matrix.CASES["generic"], Path(tmp))
            lease.path.write_text("not a socket", encoding="utf-8")

            with self.assertRaisesRegex(matrix.CaseFailure, "not a Unix socket"):
                matrix.wait_for_unix_socket(
                    lease,
                    FakeServerProcess(),
                    timeout=0.05,
                )

            lease.path.unlink()
            lease.directory.rmdir()

    def test_stop_server_requires_server_to_remove_socket(self):
        matrix = load_matrix()
        with tempfile.TemporaryDirectory() as tmp:
            lease = matrix.allocate_socket_lease(matrix.CASES["generic"], Path(tmp))
            listener = socket.socket(socket.AF_UNIX)
            listener.bind(str(lease.path))
            try:
                with self.assertRaisesRegex(matrix.CaseFailure, "did not remove Unix socket"):
                    matrix.stop_server(
                        FakeServerProcess(),
                        io.StringIO(),
                        lease,
                        socket_timeout=0.02,
                    )
            finally:
                listener.close()

            self.assertFalse(os.path.lexists(lease.path))
            self.assertFalse(lease.directory.exists())

    def test_start_server_returns_only_after_real_socket_is_ready(self):
        matrix = load_matrix()
        with tempfile.TemporaryDirectory() as tmp:
            root = Path(tmp)
            binary = root / "vfio-user" / "build" / "vfio-device"
            binary.parent.mkdir(parents=True)
            binary.write_text("fake", encoding="utf-8")
            work_dir = root / "work"
            work_dir.mkdir(mode=0o700)
            artifacts = root / "artifacts"
            artifacts.mkdir()
            listener = socket.socket(socket.AF_UNIX)

            def fake_popen(command, **_kwargs):
                socket_path = Path(command[command.index("--socket") + 1])
                listener.bind(str(socket_path))

                def remove_socket():
                    listener.close()
                    socket_path.unlink()

                return FakeServerProcess(
                    io.StringIO('{"event":"ready"}\n'),
                    on_terminate=remove_socket,
                )

            with (
                mock.patch.object(matrix, "ROOT", root),
                mock.patch.object(matrix.subprocess, "Popen", side_effect=fake_popen),
                mock.patch.object(
                    matrix.select,
                    "select",
                    side_effect=lambda reads, _writes, _errors, _timeout: (reads, [], []),
                ),
            ):
                process, output, lease, record = matrix.start_server(
                    matrix.CASES["generic"], artifacts, work_dir, timeout=0.2
                )
                self.assertEqual(record["event"], "ready")
                self.assertTrue(stat.S_ISSOCK(lease.path.lstat().st_mode))
                matrix.stop_server(process, output, lease, socket_timeout=0.05)

            self.assertFalse(lease.directory.exists())

    def test_ready_record_without_socket_fails_and_releases_private_lease(self):
        matrix = load_matrix()
        with tempfile.TemporaryDirectory() as tmp:
            root = Path(tmp)
            binary = root / "vfio-user" / "build" / "vfio-device"
            binary.parent.mkdir(parents=True)
            binary.write_text("fake", encoding="utf-8")
            work_dir = root / "work"
            work_dir.mkdir(mode=0o700)
            artifacts = root / "artifacts"
            artifacts.mkdir()
            processes = []

            def fake_popen(_command, **_kwargs):
                process = FakeServerProcess(io.StringIO('{"event":"ready"}\n'))
                processes.append(process)
                return process

            with (
                mock.patch.object(matrix, "ROOT", root),
                mock.patch.object(matrix.subprocess, "Popen", side_effect=fake_popen),
                mock.patch.object(
                    matrix.select,
                    "select",
                    side_effect=lambda reads, _writes, _errors, _timeout: (reads, [], []),
                ),
            ):
                with self.assertRaisesRegex(matrix.CaseFailure, "did not publish a Unix socket"):
                    matrix.start_server(
                        matrix.CASES["generic"], artifacts, work_dir, timeout=0.02
                    )

            self.assertEqual(processes[0].returncode, -15)
            self.assertEqual(list(work_dir.glob(".v-*")), [])

    def test_server_spawn_failure_releases_private_lease(self):
        matrix = load_matrix()
        with tempfile.TemporaryDirectory() as tmp:
            root = Path(tmp)
            binary = root / "vfio-user" / "build" / "vfio-device"
            binary.parent.mkdir(parents=True)
            binary.write_text("fake", encoding="utf-8")
            work_dir = root / "work"
            work_dir.mkdir(mode=0o700)
            artifacts = root / "artifacts"
            artifacts.mkdir()

            with (
                mock.patch.object(matrix, "ROOT", root),
                mock.patch.object(
                    matrix.subprocess,
                    "Popen",
                    side_effect=OSError("spawn failed"),
                ),
            ):
                with self.assertRaisesRegex(OSError, "spawn failed"):
                    matrix.start_server(
                        matrix.CASES["generic"], artifacts, work_dir, timeout=0.02
                    )

            self.assertEqual(list(work_dir.glob(".v-*")), [])

    def test_run_command_timeout_is_reported(self):
        matrix = load_matrix()
        with tempfile.TemporaryDirectory() as tmp:
            with self.assertRaisesRegex(matrix.CaseFailure, "timed out"):
                matrix.run_command(
                    [sys.executable, "-c", "import time; time.sleep(30)"],
                    Path(tmp) / "command.log",
                    timeout=1,
                )

    def test_nvme_firmware_and_vfio_contracts_expose_same_core_features(self):
        root = Path(__file__).resolve().parents[2]
        c_source = (root / "vfio-user" / "src" / "behavior_nvme.c").read_text(encoding="utf-8")
        sv_source = (root / "internal" / "firmware" / "svgen" / "templates" /
                     "nvme_admin_responder.sv.tmpl").read_text(encoding="utf-8")
        for token in ("0x02", "0x06", "0x80", "0xc0", "0x09"):
            self.assertIn(token, c_source)
        for token in ("8'h02", "8'h06", "8'h80", "8'hC0", "8'h09"):
            self.assertIn(token, sv_source)
        for token in ("LOG_PAGE_VENDOR", "PRP_LIST", "stat_dma_mrd_tlps",
                      "stat_transport_errors"):
            self.assertIn(token, sv_source)


if __name__ == "__main__":
    unittest.main()
