import json
import os
import select
import signal
import subprocess
import tempfile
import unittest
from pathlib import Path


class ServerCliTests(unittest.TestCase):
    @classmethod
    def setUpClass(cls):
        cls.binary = Path(os.environ.get("VFIO_DEVICE_BIN", "build/vfio-device")).resolve()
        if not cls.binary.is_file():
            raise unittest.SkipTest(f"server binary is missing: {cls.binary}")
        artifacts = os.environ.get("VFIO_ARTIFACTS")
        if not artifacts:
            raise RuntimeError("VFIO_ARTIFACTS must name an isolated generated fixture")
        cls.artifacts = Path(artifacts).resolve()
        if not cls.artifacts.is_dir():
            raise RuntimeError(f"server fixture is missing: {cls.artifacts}")

    def test_missing_artifacts_fail(self):
        result = subprocess.run(
            [self.binary, "--artifacts", "/missing", "--socket", "/tmp/missing.sock"],
            text=True,
            capture_output=True,
            check=False,
        )

        self.assertNotEqual(result.returncode, 0)
        self.assertIn("device model", result.stderr)

    def test_existing_socket_path_is_not_removed(self):
        with tempfile.TemporaryDirectory() as tmp:
            socket_path = Path(tmp) / "occupied.sock"
            socket_path.write_text("do-not-remove", encoding="utf-8")
            result = subprocess.run(
                [self.binary, "--artifacts", self.artifacts, "--socket", socket_path],
                text=True,
                capture_output=True,
                check=False,
            )

            self.assertNotEqual(result.returncode, 0)
            self.assertEqual(socket_path.read_text(encoding="utf-8"), "do-not-remove")

    def test_readiness_and_sigterm_cleanup(self):
        with tempfile.TemporaryDirectory() as tmp:
            socket_path = Path(tmp) / "device.sock"
            process = subprocess.Popen(
                [self.binary, "--artifacts", self.artifacts, "--socket", socket_path],
                text=True,
                stdout=subprocess.PIPE,
                stderr=subprocess.PIPE,
            )
            ready, _, _ = select.select([process.stdout], [], [], 5)
            self.assertTrue(ready, "server did not emit readiness")
            record = json.loads(process.stdout.readline())
            self.assertEqual(record["event"], "ready")
            self.assertEqual(record["vendor_id"], "1234")
            self.assertEqual(record["device_id"], "5678")
            self.assertEqual(record["bar_count"], 1)
            process.send_signal(signal.SIGTERM)
            process.communicate(timeout=5)
            self.assertEqual(process.returncode, 0)
            self.assertFalse(socket_path.exists())


if __name__ == "__main__":
    unittest.main()
