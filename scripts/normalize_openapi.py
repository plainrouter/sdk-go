#!/usr/bin/env python3
"""Normalize free-form JSON unions for OpenAPI Generator's Go client.

OpenAPI Generator 7.25.0 emits invalid Go identifiers for OAS 3.1 schemas whose
value may be an object, an array, or null, and for free-form object-or-array
unions (including an empty-array branch). Those fields intentionally represent arbitrary JSON. Replacing only those
exact unions with an unconstrained schema makes the generated type `interface{}`
without narrowing accepted input. It also emits the invalid type `nil` for a
schema that is only ever null, which receives the same replacement.
The signed source contract remains byte-for-byte unchanged in spec/openapi.json.
"""

from __future__ import annotations

import argparse
import copy
import hashlib
import json
import re
from pathlib import Path
from typing import Any


def is_free_form_container(variants: list[Any]) -> bool:
    """An object with arbitrary members or an array with arbitrary items."""
    if len(variants) != 2 or not all(isinstance(variant, dict) for variant in variants):
        return False
    by_type = {variant.get("type"): variant for variant in variants}
    if set(by_type) != {"object", "array"}:
        return False
    return (
        by_type["object"].get("additionalProperties") in ({}, True)
        and set(by_type["object"]) <= {"type", "additionalProperties"}
        and by_type["array"].get("items", {}) == {}
        and set(by_type["array"]) <= {"type", "items", "maxItems"}
        and by_type["array"].get("maxItems", 0) == 0
    )


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
            if (types == {"object", "array", "null"} and len(variants) == 3) or is_free_form_container(variants):
                description = value.get("description")
                value.clear()
                if description is not None:
                    value["description"] = description
                changed += 1
        if value.get("type") == "null" and set(value) <= {"type", "description"}:
            del value["type"]
            changed += 1
        for child in value.values():
            changed += normalize(child)
    elif isinstance(value, list):
        for child in value:
            changed += normalize(child)
    return changed


def hoist_shared_closed_objects(document: dict[str, Any]) -> None:
    """Share structurally identical closed objects used by distinct operations.

    Derive both discovery and names from the spec. Common operation/property
    names retain readable public model names; unrelated operations use a stable
    structural hash. Rewrite nested duplicates as well as top-level bodies.
    """
    occurrences: dict[str, list[tuple[str, list[str]]]] = {}
    schemas: dict[str, Any] = {}

    def collect(value: Any, operation: str, path: list[str]) -> None:
        if isinstance(value, dict):
            if value.get("additionalProperties") is False:
                signature = json.dumps(value, sort_keys=True)
                occurrences.setdefault(signature, []).append((operation, path))
                schemas[signature] = value
            for key, child in value.items():
                collect(child, operation, [*path, key])
        elif isinstance(value, list):
            for index, child in enumerate(value):
                collect(child, operation, [*path, str(index)])

    operations = [
        operation
        for item in document["paths"].values()
        for method, operation in item.items()
        if method.lower() in {"get", "put", "post", "delete", "patch", "options", "head", "trace"}
        and isinstance(operation, dict)
    ]
    for operation in operations:
        collect(operation, operation["operationId"], [])

    components = document.setdefault("components", {}).setdefault("schemas", {})
    names: dict[str, str] = {}
    for signature, uses in occurrences.items():
        if len({operation for operation, _ in uses}) < 2:
            continue
        tokens = [re.split(r"[^a-zA-Z0-9]+", operation) for operation, _ in uses]
        common = []
        for column in zip(*tokens):
            if len(set(column)) != 1:
                break
            common.append(column[0])
        path = uses[0][1]
        properties = [path[index + 1] for index, key in enumerate(path[:-1]) if key == "properties"]
        # Singular common operation nouns preserve established shared type names.
        parts = [token.removesuffix("s") for token in common] + properties
        name = "".join(part[:1].upper() + part[1:] for part in parts)
        if len(path) >= 2 and path[-2] in {"anyOf", "oneOf"}:
            name += "Variant" + str(int(path[-1]) + 1)
        digest = hashlib.sha256(signature.encode()).hexdigest()[:12]
        if not common or not name or name in components or name in names.values():
            name = "SharedClosedObject" + digest
        names[signature] = name

    def rewrite(value: Any, root: bool = False) -> Any:
        if isinstance(value, dict):
            signature = json.dumps(value, sort_keys=True)
            if not root and signature in names:
                return {"$ref": "#/components/schemas/" + names[signature]}
            return {key: rewrite(child) for key, child in value.items()}
        if isinstance(value, list):
            return [rewrite(child) for child in value]
        return value

    shared = {name: rewrite(copy.deepcopy(schemas[signature]), root=True) for signature, name in names.items()}
    for operation in operations:
        rewritten = rewrite(operation, root=True)
        operation.clear()
        operation.update(rewritten)
    components.update(shared)


def main() -> None:
    parser = argparse.ArgumentParser()
    parser.add_argument("source", type=Path)
    parser.add_argument("destination", type=Path)
    args = parser.parse_args()

    document = json.loads(args.source.read_text(encoding="utf-8"))
    changed = normalize(document)

    # Go defaults unformatted integers to int32, even when a bound exceeds it.
    # Optional query defaults otherwise generate an explicit page=1 fallback.
    for item in document["paths"].values():
        for operation in item.values():
            if not isinstance(operation, dict):
                continue
            for parameter in operation.get("parameters", []):
                schema = parameter.get("schema", {})
                if parameter.get("name") == "page" and schema.get("type") == "integer":
                    schema["format"] = "int64"
                    schema.pop("default", None)

    # Inline-model deduplication can discard closure on the first copy.
    hoist_shared_closed_objects(document)

    # An empty closed object is otherwise emitted as map[string]interface{}.
    # Schema mapping plus the supporting template makes it a Go struct{}, which can
    # neither retain nor serialize unknown keys. The generation checker audits
    # all closed objects, including nonempty inline union variants.
    def close_empty(value: Any) -> None:
        if isinstance(value, dict):
            if value.get("type") == "object" and value.get("additionalProperties") is False and not value.get("properties"):
                value.clear()
                value["$ref"] = "#/components/schemas/EmptyClosedObject"
            else:
                for child in value.values():
                    close_empty(child)
        elif isinstance(value, list):
            for child in value:
                close_empty(child)

    close_empty(document)
    document["components"]["schemas"]["EmptyClosedObject"] = {
        "type": "object", "additionalProperties": False, "properties": {}
    }
    if changed == 0:
        raise SystemExit("expected at least one free-form object/array/null union")

    args.destination.write_text(json.dumps(document, indent=2) + "\n", encoding="utf-8")
    print(f"Normalized {changed} free-form JSON unions")


if __name__ == "__main__":
    main()
