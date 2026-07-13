#!/usr/bin/env python3
import json
import re
import sys


def parse_log(text: str, platform: str) -> dict:
    events = []
    errors = []
    if platform == "linux":
        for line in text.splitlines():
            try:
                record = json.loads(line)
            except json.JSONDecodeError:
                continue
            if record.get("event") in {"stage", "result"}:
                events.append(record)
                if record.get("status") == "fail":
                    errors.append(record.get("detail", "probe-failed"))
    elif platform == "windows":
        patterns = {
            "code10": r"(?i)code\s*10|CM_PROB_FAILED_START",
            "driver": r"(?i)(?:driver|service)\s+(?:started|bound|loaded)",
            "bar": r"(?i)\bBAR[0-5]\b",
            "msi": r"(?i)MSI(?:-X)?",
            "reset": r"(?i)\b(?:FLR|reset|restart)\b",
        }
        for line in text.splitlines():
            for stage, pattern in patterns.items():
                if re.search(pattern, line):
                    events.append({"event": "stage", "stage": stage, "source": "windows", "text": line})
            if re.search(patterns["code10"], line):
                errors.append("code10")
    else:
        raise ValueError(f"unsupported platform: {platform}")
    stages = {event.get("stage") for event in events}
    if not events:
        return {"status": "blocked", "platform": platform, "events": [], "errors": ["no-probe-data"]}
    status = "fail" if errors else "pass"
    if platform == "linux" and "result" not in {event.get("event") for event in events}:
        status = "blocked"
        errors.append("no-terminal-result")
    return {"status": status, "platform": platform, "events": events,
            "stages": sorted(stage for stage in stages if stage), "errors": errors}


def main() -> int:
    platform = sys.argv[1] if len(sys.argv) > 1 else "linux"
    report = parse_log(sys.stdin.read(), platform)
    print(json.dumps(report, sort_keys=True))
    return 0 if report["status"] == "pass" else 1


if __name__ == "__main__":
    raise SystemExit(main())
