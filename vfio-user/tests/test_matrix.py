import importlib.util
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


if __name__ == "__main__":
    unittest.main()
