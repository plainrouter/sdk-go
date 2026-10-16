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
import json
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

    # Inline-model deduplication loses additionalProperties:false on the store
    # definitions when identical update variants exist. Explicit shared refs
    # preserve closure on both operations before the generator flattens them.
    rule_definitions: dict[str, str] = {}
    for item in document["paths"].values():
        for operation in item.values():
            if not isinstance(operation, dict) or operation.get("operationId") not in {"actions.rules.store", "actions.rules.update"}:
                continue
            definition = operation["requestBody"]["content"]["application/json"]["schema"]["properties"]["definition"]
            for index, variant in enumerate(definition["anyOf"]):
                signature = json.dumps(variant, sort_keys=True)
                if signature not in rule_definitions:
                    name = f"ActionRuleDefinitionVariant{len(rule_definitions) + 1}"
                    document["components"]["schemas"][name] = variant
                    rule_definitions[signature] = name
                definition["anyOf"][index] = {"$ref": "#/components/schemas/" + rule_definitions[signature]}

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
