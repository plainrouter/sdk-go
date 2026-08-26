#!/usr/bin/env python3
"""Verify the vendored PlainRouter contract against its provenance record."""

from __future__ import annotations

import hashlib
import json
from pathlib import Path


ROOT = Path(__file__).resolve().parents[1]
SPEC = ROOT / "spec" / "openapi.json"
CHECKSUM = ROOT / "spec" / "CHECKSUM"


def read_record() -> dict[str, str]:
    record: dict[str, str] = {}
    for line in CHECKSUM.read_text(encoding="utf-8").splitlines():
        key, separator, value = line.partition(": ")
        if separator:
            record[key] = value
    return record


def main() -> None:
    record = read_record()
    raw = SPEC.read_bytes()
    document = json.loads(raw)
    actual_sha = hashlib.sha256(raw).hexdigest()

    assert actual_sha == record["sha256"], "spec/openapi.json checksum mismatch"
    assert document["info"]["version"] == record["info.version"], "API version mismatch"
    assert document.get("x-signed") is True, "OpenAPI contract is not signed"
    assert record["x-signed"] == "true", "provenance record does not require a signed contract"
    assert record["source"].startswith("https://plainrouter.com/"), "unexpected contract source"

    print(f"Verified signed OpenAPI {record['info.version']} ({actual_sha})")


if __name__ == "__main__":
    main()
