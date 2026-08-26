#!/usr/bin/env python3
"""Compare generated files with a fresh, clean generation directory."""

from __future__ import annotations

import argparse
from pathlib import Path


REPOSITORY_OWNED = {
    ".gitignore",
    ".openapi-generator-ignore",
    ".travis.yml",
    "README.md",
    "api/openapi.yaml",
    "git_push.sh",
    "go.sum",
}


def manifest(root: Path) -> set[str]:
    return {
        line.strip()
        for line in (root / ".openapi-generator" / "FILES").read_text(encoding="utf-8").splitlines()
        if line.strip() and line.strip() not in REPOSITORY_OWNED
    }


def main() -> None:
    parser = argparse.ArgumentParser()
    parser.add_argument("expected", type=Path)
    parser.add_argument("actual", type=Path)
    args = parser.parse_args()

    expected_files = manifest(args.expected)
    actual_files = manifest(args.actual)
    if expected_files != actual_files:
        missing = sorted(expected_files - actual_files)
        extra = sorted(actual_files - expected_files)
        raise SystemExit(f"generated manifest drift; missing={missing}, extra={extra}")

    changed = [
        relative
        for relative in sorted(expected_files)
        if (args.expected / relative).read_bytes() != (args.actual / relative).read_bytes()
    ]
    if changed:
        raise SystemExit("generated file drift: " + ", ".join(changed))

    print(f"Verified {len(expected_files)} generator-owned files")


if __name__ == "__main__":
    main()
