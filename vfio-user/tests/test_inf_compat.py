import importlib.util
import json
import tempfile
import unittest
from pathlib import Path


ROOT = Path(__file__).resolve().parents[2]
SPEC = importlib.util.spec_from_file_location("check_inf_compat", ROOT / "tools" / "check_inf_compat.py")
CHECKER = importlib.util.module_from_spec(SPEC)
assert SPEC.loader is not None
SPEC.loader.exec_module(CHECKER)


class InfCompatibilityTests(unittest.TestCase):
    def setUp(self):
        self.tmp = tempfile.TemporaryDirectory()
        root = Path(self.tmp.name)
        self.model = root / "device_model.json"
        self.model.write_text(json.dumps({
            "functions": [{
                "vendor_id": 0x144D,
                "device_id": 0xA809,
                "subsystem_vendor_id": 0x144D,
                "subsystem_device_id": 0xA809,
                "revision_id": 1,
                "class_code": 0x010802,
            }],
        }), encoding="utf-8")
        self.inf = root / "device.inf"

    def tearDown(self):
        self.tmp.cleanup()

    def test_accepts_exact_subsystem_id(self):
        self.inf.write_text("%PCI\\VEN_144D&DEV_A809&SUBSYS_A809144D% = Install\n", encoding="utf-8")
        self.assertEqual(CHECKER.check(self.model, self.inf)[0], "pass")

    def test_accepts_case_insensitive_base_id(self):
        self.inf.write_text("pci\\ven_144d&dev_a809 = Install\n", encoding="utf-8")
        self.assertEqual(CHECKER.check(self.model, self.inf)[0], "pass")

    def test_rejects_unmatched_id(self):
        self.inf.write_text("PCI\\VEN_8086&DEV_1234 = Install\n", encoding="utf-8")
        status, details = CHECKER.check(self.model, self.inf)
        self.assertEqual(status, "fail")
        self.assertIn("no INF model match", details[0])

    def test_missing_inf_is_blocked(self):
        status, _ = CHECKER.check(self.model, self.inf)
        self.assertEqual(status, "blocked")


if __name__ == "__main__":
    unittest.main()
