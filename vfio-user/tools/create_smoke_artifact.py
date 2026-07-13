#!/usr/bin/env python3
from __future__ import annotations

import argparse
import base64
import hashlib
import json
from pathlib import Path


def b64_zeros(size: int) -> str:
    return base64.b64encode(bytes(size)).decode("ascii")


def build_model() -> dict:
    config = bytearray(256)
    config[0:2] = (0x1234).to_bytes(2, "little")
    config[2:4] = (0x5678).to_bytes(2, "little")
    config[0x0A:0x0D] = bytes([0x00, 0x00, 0x00])
    config[0x0E] = 0x00
    config[0x10:0x14] = (0xfffff000).to_bytes(4, "little")
    return {
        "schema_version": 1,
        "functions": [
            {
                "vendor_id": 0x1234,
                "device_id": 0x5678,
                "subsystem_vendor_id": 0,
                "subsystem_device_id": 0,
                "revision_id": 0,
                "class_code": 0x000000,
                "header_type": 0,
            }
        ],
        "config_space": {
            "size": 256,
            "reset_image": base64.b64encode(config).decode("ascii"),
        },
        "bars": [
            {
                "bir": 0,
                "type": "mem32",
                "size": 4096,
                "prefetchable": False,
                "address_width": 32,
                "reset_image": b64_zeros(4096),
            }
        ],
        "interrupts": [],
        "msix": None,
    }


def write_artifact(output: Path) -> None:
    output.mkdir(parents=True, exist_ok=True)
    model_data = json.dumps(build_model(), indent=2).encode("utf-8")
    (output / "device_model.json").write_bytes(model_data)
    manifest = {
        "files": [
            {
                "name": "device_model.json",
                "size": len(model_data),
                "sha256": hashlib.sha256(model_data).hexdigest(),
            }
        ]
    }
    (output / "build_manifest.json").write_text(json.dumps(manifest, indent=2) + "\n", encoding="utf-8")


def main() -> int:
    parser = argparse.ArgumentParser()
    parser.add_argument("--output", type=Path, required=True)
    args = parser.parse_args()
    write_artifact(args.output)
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
