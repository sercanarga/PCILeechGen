import importlib.util
import socket
import sys
import tempfile
import unittest
from pathlib import Path


MODULE_PATH = Path(__file__).resolve().parents[1] / "matrix.py"


def load_matrix():
    spec = importlib.util.spec_from_file_location("vfio_matrix", MODULE_PATH)
    module = importlib.util.module_from_spec(spec)
    sys.modules[spec.name] = module
    spec.loader.exec_module(module)
    return module


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

    def test_parse_guest_results_requires_passing_terminal_record(self):
        matrix = load_matrix()

        result = matrix.parse_guest_results(
            '{"event":"result","case":"generic","status":"pass","bdf":"0000:03:00.0","vendor":"1234","device":"5678","class":"000000","driver":"none"}\n',
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
            "[    1.2] kernel message\n"
            '{"event":"result","case":"generic","status":"pass",'
            '"bdf":"0000:03:00.0","vendor":"1234","device":"5678",'
            '"class":"000000","driver":"none"}\n',
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
        for case in matrix.CASES.values():
            artifacts = Path(__file__).resolve().parents[2] / "tests" / "cocotb" / f"out_{case.name}"
            contract = matrix.build_contract(case, artifacts)
            self.assertEqual(contract["case"], case.name)
            self.assertTrue(contract["bars"])
            self.assertTrue(contract["capabilities"])
            self.assertEqual(contract["probe"][:3], ["enumerate", "bars", "reset"])

    def test_nvme_qemu_requires_kvm(self):
        matrix = load_matrix()

        self.assertTrue(matrix.qemu_requires_kvm(matrix.CASES["nvme"], Path("/missing/kvm")))
        self.assertFalse(matrix.qemu_requires_kvm(matrix.CASES["generic"], Path("/missing/kvm")))

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

    def test_prepare_socket_path_removes_stale_socket(self):
        matrix = load_matrix()
        with tempfile.TemporaryDirectory() as tmp:
            path = Path(tmp) / "device.sock"
            listener = socket.socket(socket.AF_UNIX)
            listener.bind(str(path))
            listener.close()
            matrix.prepare_socket_path(path)
            self.assertFalse(path.exists())

    def test_prepare_socket_path_rejects_non_socket(self):
        matrix = load_matrix()
        with tempfile.TemporaryDirectory() as tmp:
            path = Path(tmp) / "device.sock"
            path.write_text("occupied", encoding="utf-8")
            with self.assertRaisesRegex(matrix.CaseFailure, "not a socket"):
                matrix.prepare_socket_path(path)

    def test_prepare_socket_path_rejects_active_socket(self):
        matrix = load_matrix()
        with tempfile.TemporaryDirectory() as tmp:
            path = Path(tmp) / "device.sock"
            listener = socket.socket(socket.AF_UNIX)
            listener.bind(str(path))
            listener.listen(1)
            try:
                with self.assertRaisesRegex(matrix.CaseFailure, "already active"):
                    matrix.prepare_socket_path(path)
            finally:
                listener.close()

    def test_run_command_timeout_is_reported(self):
        matrix = load_matrix()
        with tempfile.TemporaryDirectory() as tmp:
            with self.assertRaisesRegex(matrix.CaseFailure, "timed out"):
                matrix.run_command(
                    [sys.executable, "-c", "import time; time.sleep(30)"],
                    Path(tmp) / "command.log",
                    timeout=1,
                )


if __name__ == "__main__":
    unittest.main()
