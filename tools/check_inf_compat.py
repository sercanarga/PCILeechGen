#!/usr/bin/env python3
"""Check whether a Windows PCI INF contains a model's hardware ID."""

import argparse
import json
import re
import sys
from pathlib import Path

PCI_ID = re.compile(r"PCI\\VEN_[0-9A-F]{4}&DEV_[0-9A-F]{4}(?:&(?:SUBSYS_[0-9A-F]{8}|REV_[0-9A-F]{2}|CC_[0-9A-F]{6}))?", re.I)


def model_ids(model: dict) -> list[str]:
    functions = model.get("functions") or []
    if not functions:
        raise ValueError("model has no PCI function")
    fn = functions[0]
    vendor = int(fn["vendor_id"])
    device = int(fn["device_id"])
    subsystem = (int(fn["subsystem_device_id"]) << 16) | int(fn["subsystem_vendor_id"])
    revision = int(fn["revision_id"])
    class_code = int(fn["class_code"])
    prefix = f"PCI\\VEN_{vendor:04X}&DEV_{device:04X}"
    return [
        f"{prefix}&SUBSYS_{subsystem:08X}",
        f"{prefix}&REV_{revision:02X}",
        f"{prefix}&CC_{class_code:06X}",
        prefix,
        f"PCI\\CC_{class_code:06X}",
    ]


def check(model_path: Path, inf_path: Path) -> tuple[str, list[str]]:
    if not model_path.is_file() or not inf_path.is_file():
        return "blocked", ["model or INF file is missing"]
    try:
        model = json.loads(model_path.read_text(encoding="utf-8"))
        expected = model_ids(model)
    except (OSError, ValueError, KeyError, TypeError, json.JSONDecodeError) as exc:
        return "blocked", [f"invalid model: {exc}"]
    try:
        text = inf_path.read_text(encoding="utf-8-sig", errors="replace")
    except OSError as exc:
        return "blocked", [f"cannot read INF: {exc}"]
    found = {match.upper() for match in PCI_ID.findall(text)}
    matches = [identifier for identifier in expected if identifier.upper() in found]
    if matches:
        return "pass", matches
    return "fail", [f"no INF model match; expected one of: {', '.join(expected)}"]


def main() -> int:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--model", type=Path, required=True)
    parser.add_argument("--inf", type=Path, required=True)
    args = parser.parse_args()
    status, details = check(args.model, args.inf)
    print(json.dumps({"status": status, "details": details}, sort_keys=True))
    return {"pass": 0, "fail": 1, "blocked": 2}[status]


if __name__ == "__main__":
    sys.exit(main())
