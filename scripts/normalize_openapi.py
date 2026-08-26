#!/usr/bin/env python3
"""Normalize free-form JSON unions for OpenAPI Generator's Go client.

OpenAPI Generator 7.25.0 emits invalid Go identifiers for OAS 3.1 schemas whose
value may be an object, an array, or null. Those fields intentionally represent
arbitrary provider JSON. Replacing only that exact union with an unconstrained
schema makes the generated type `interface{}` without narrowing accepted input.
The signed source contract remains byte-for-byte unchanged in spec/openapi.json.
"""

from __future__ import annotations

import argparse
import json
from pathlib import Path
from typing import Any


def normalize(value: Any) -> int:
    changed = 0
    if isinstance(value, dict):
        variants = value.get("anyOf")
        if isinstance(variants, list):
            types = {
                variant.get("type")
                for variant in variants
                if isinstance(variant, dict) and isinstance(variant.get("type"), str)
            }
            if types == {"object", "array", "null"} and len(variants) == 3:
                description = value.get("description")
                value.clear()
                if description is not None:
                    value["description"] = description
                changed += 1
        for child in value.values():
            changed += normalize(child)
    elif isinstance(value, list):
        for child in value:
            changed += normalize(child)
    return changed


def main() -> None:
    parser = argparse.ArgumentParser()
    parser.add_argument("source", type=Path)
    parser.add_argument("destination", type=Path)
    args = parser.parse_args()

    document = json.loads(args.source.read_text(encoding="utf-8"))
    changed = normalize(document)
    if changed == 0:
        raise SystemExit("expected at least one free-form object/array/null union")

    args.destination.write_text(json.dumps(document, indent=2) + "\n", encoding="utf-8")
    print(f"Normalized {changed} free-form JSON unions")


if __name__ == "__main__":
    main()
