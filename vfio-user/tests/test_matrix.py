import importlib.util
import json
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
