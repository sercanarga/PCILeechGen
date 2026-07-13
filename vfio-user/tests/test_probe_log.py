import importlib.util
import unittest
from pathlib import Path


PATH = Path(__file__).resolve().parents[1] / "tools" / "parse_probe_log.py"
SPEC = importlib.util.spec_from_file_location("parse_probe_log", PATH)
MODULE = importlib.util.module_from_spec(SPEC)
SPEC.loader.exec_module(MODULE)


class ProbeLogTests(unittest.TestCase):
    def test_linux_stages_and_result(self):
        report = MODULE.parse_log(
            '{"event":"stage","stage":"enumerate"}\n'
            '{"event":"stage","stage":"bars","count":2}\n'
            '{"event":"stage","stage":"driver","driver":"nvme","status":"pass"}\n'
            '{"event":"result","status":"pass"}\n',
            "linux",
        )
        self.assertEqual(report["status"], "pass")
        self.assertIn("bars", report["stages"])

    def test_windows_code10_is_failure(self):
        report = MODULE.parse_log("Device reported CM_PROB_FAILED_START (Code 10)\nMSI-X enabled\n", "windows")
        self.assertEqual(report["status"], "fail")
        self.assertIn("code10", report["errors"])

    def test_missing_data_is_blocked(self):
        report = MODULE.parse_log("", "linux")
        self.assertEqual(report["status"], "blocked")
        self.assertEqual(report["errors"], ["no-probe-data"])


if __name__ == "__main__":
    unittest.main()
