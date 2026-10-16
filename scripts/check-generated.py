#!/usr/bin/env python3
"""Compare generated files with a fresh, clean generation directory."""

from __future__ import annotations

import argparse
import json
import os
import re
import shutil
import subprocess
import tempfile
from pathlib import Path


REPOSITORY_OWNED = {
    ".gitignore",
    "doc.go",
    ".openapi-generator-ignore",
    ".travis.yml",
    "README.md",
    "api/openapi.yaml",
    "git_push.sh",
    "go.sum",
}


def manifest(root: Path) -> set[str]:
    return {
        f"openapi/{line.strip()}"
        for line in (root / "openapi" / ".openapi-generator" / "FILES")
        .read_text(encoding="utf-8")
        .splitlines()
        if line.strip() and line.strip() not in REPOSITORY_OWNED
    }


def check_closed_objects(output: Path, source: Path) -> None:
    """Audit every signed closed schema through its actual generated Go type.

    Follow JSON-tagged fields and union members instead of assuming the
    generator's inline names or keeping a list of closed models. Execute the
    decoders as well: source-token checks alone cannot prove strict behavior.
    """
    spec = json.loads((source / "spec/openapi.json").read_text())
    models = {}
    for file in (output / "openapi").glob("*.go"):
        text = file.read_text()
        structures = re.findall(r"^type (\w+) struct \{(.*?)^\}", text, re.M | re.S)
        structures += [(name, "") for name in re.findall(r"^type (\w+) struct\{\}", text, re.M)]
        for name, body in structures:
            fields = {}
            members = []
            for line in body.splitlines():
                match = re.match(r'\s*(\w+)\s+([^`\s]+)(?:\s+`json:"([^",]+)[^`]*`)?', line)
                if match:
                    field, kind, tag = match.groups()
                    if tag:
                        fields[tag] = kind
                    members.append((field, kind))
            models[name] = (fields, members)

    def resolve(schema, path):
        if "$ref" in schema:
            path = schema["$ref"][2:]
            schema = spec
            for key in path.split("/"):
                schema = schema[key]
        return schema, path

    def has_closed(schema):
        if not isinstance(schema, dict):
            return False
        schema, _ = resolve(schema, "")
        if schema.get("additionalProperties") is False:
            return True
        children = list(schema.get("properties", {}).values())
        children += [schema.get("items"), schema.get("additionalProperties")]
        for union in ("anyOf", "oneOf", "allOf"):
            children += schema.get(union, [])
        return any(has_closed(child) for child in children)

    def sample(schema):
        schema, _ = resolve(schema, "")
        if "enum" in schema:
            return next((value for value in schema["enum"] if value is not None), None)
        for union in ("anyOf", "oneOf"):
            if union in schema:
                return sample(next(child for child in schema[union] if child.get("type") != "null"))
        kind = schema.get("type")
        if isinstance(kind, list):
            kind = next(value for value in kind if value != "null")
        if kind == "object" or "properties" in schema:
            return {key: sample(schema["properties"][key]) for key in schema.get("required", [])}
        if kind == "array":
            return [sample(schema.get("items", {})) for _ in range(schema.get("minItems", 0))]
        if kind in {"integer", "number"}:
            return schema.get("minimum", 1)
        if kind == "boolean":
            return False
        return "fixture"

    rows = {}
    covered = set()

    def visit(schema, kind, path):
        schema, path = resolve(schema, path)
        if not has_closed(schema):
            return
        kind = kind.lstrip("*")
        if "items" in schema:
            if not kind.startswith("[]"):
                raise SystemExit(f"closed schema audit: {path} has non-array Go type {kind}")
            visit(schema["items"], kind[2:], path + "/items")
            return
        if kind.startswith("Nullable") and kind not in models:
            kind = kind[len("Nullable"):]
        if kind.startswith("Nullable") and kind in models:
            values = [value for field, value in models[kind][1] if field == "value"]
            if values:
                kind = values[0].lstrip("*")
        if kind not in models:
            raise SystemExit(f"closed schema audit: {path} has no generated struct ({kind})")
        fields, members = models[kind]
        for union in ("anyOf", "oneOf", "allOf"):
            for index, branch in enumerate(schema.get(union, [])):
                if not has_closed(branch):
                    continue
                candidates = [value.lstrip("*") for _, value in members if value.lstrip("*") in models]
                inline = kind + union[0].upper() + union[1:] + (str(index) if index else "")
                reference = branch.get("$ref", "").split("/")[-1]
                resolved, _ = resolve(branch, "")
                if inline in candidates:
                    candidates = [inline]
                elif reference in candidates:
                    candidates = [reference]
                else:
                    candidates = [name for name in candidates if set(models[name][0]) == set(resolved.get("properties", {}))]
                if not candidates:
                    raise SystemExit(f"closed schema audit: cannot map {path}/{union}/{index} through {kind}")
                for candidate in candidates:
                    visit(branch, candidate, f"{path}/{union}/{index}")
        if schema.get("additionalProperties") is False:
            covered.add(path)
            if set(fields) != set(schema.get("properties", {})):
                raise SystemExit(f"closed schema audit: {path} properties differ from {kind}")
            rows.setdefault(kind, (path, sample(schema)))
        for key, child in schema.get("properties", {}).items():
            if has_closed(child):
                if key not in fields:
                    raise SystemExit(f"closed schema audit: missing {kind}.{key} for {path}")
                visit(child, fields[key], path + "/properties/" + key)

    def go_name(value):
        return "".join(part[:1].upper() + part[1:] for part in re.split(r"[^a-zA-Z0-9]+", value))

    for name, schema in spec["components"]["schemas"].items():
        if has_closed(schema):
            visit(schema, go_name(name), "components/schemas/" + name)
    for path, item in spec["paths"].items():
        for method, operation in item.items():
            if not isinstance(operation, dict):
                continue
            for media, body in operation.get("requestBody", {}).get("content", {}).items():
                if not has_closed(body.get("schema", {})):
                    continue
                request = "Api" + go_name(operation["operationId"]) + "Request"
                candidates = [kind.lstrip("*") for _, kind in models.get(request, ({}, []))[1] if kind.lstrip("*") in models]
                if len(candidates) != 1:
                    raise SystemExit(f"closed schema audit: cannot identify body for {request}")
                visit(body["schema"], candidates[0], f"paths/{path}/{method}/requestBody/content/{media}/schema")

    expected = set()

    def closed_paths(value, path=""):
        if isinstance(value, dict):
            if value.get("additionalProperties") is False:
                expected.add(path)
            for key, child in value.items():
                closed_paths(child, path + "/" + key if path else key)
        elif isinstance(value, list):
            for index, child in enumerate(value):
                closed_paths(child, path + "/" + str(index))

    closed_paths(spec)
    if expected - covered:
        raise SystemExit("closed schema audit: unmapped signed schemas: " + ", ".join(sorted(expected - covered)))
    failures = [name for name in rows if any(field == "AdditionalProperties" for field, _ in models[name][1])]
    if failures:
        raise SystemExit("closed schema audit: AdditionalProperties on " + ", ".join(sorted(failures)))

    # Isolate the generated subpackage so the fresh-output audit never imports
    # committed models by mistake, and never writes files into the source tree.
    with tempfile.TemporaryDirectory(prefix="plainrouter-closed-audit-") as temporary:
        work = Path(temporary)
        (work / "openapi").symlink_to((output / "openapi").resolve(), target_is_directory=True)
        for name in ("go.mod", "go.sum"):
            shutil.copyfile(source / name, work / name)
        cases = "\n".join(
            "{" + json.dumps(name) + ", &openapi." + name + "{}, " + json.dumps(json.dumps(fixture)) + "},"
            for name, (_, fixture) in sorted(rows.items())
        )
        program = '''package main
import ("encoding/json"; "fmt"; "os"; "reflect"; "strings"; "github.com/plainrouter/sdk-go/openapi")
func main() {
    failed := false
    for _, row := range []struct {name string; value interface{}; fixture string} {
CASES
    } {
        reject := func(message string) { fmt.Fprintln(os.Stderr, row.name+": "+message); failed = true }
        if _, ok := reflect.TypeOf(row.value).Elem().FieldByName("AdditionalProperties"); ok { reject("AdditionalProperties field") }
        if err := json.Unmarshal([]byte(row.fixture), row.value); err != nil { reject("valid fixture: "+err.Error()); continue }
        var input map[string]interface{}
        json.Unmarshal([]byte(row.fixture), &input)
        input["zz_unexpected"] = true
        raw, _ := json.Marshal(input)
        if err := json.Unmarshal(raw, row.value); err == nil || !strings.Contains(err.Error(), "zz_unexpected") { reject(fmt.Sprintf("unknown key decoder: %v", err)) }
        encoded, err := json.Marshal(row.value)
        if err != nil { reject("encode: "+err.Error()) } else if strings.Contains(string(encoded), "zz_unexpected") { reject("unknown key retained on encode") }
    }
    if failed { os.Exit(1) }
}
'''.replace("CASES", cases)
        (work / "main.go").write_text(program)
        subprocess.run([os.environ.get("GO_BIN", "go"), "run", "."], cwd=work, check=True)
    print(f"Verified {len(expected)} signed closed schemas through {len(rows)} generated structs")


def main() -> None:
    parser = argparse.ArgumentParser()
    parser.add_argument("--closed-only", action="store_true")
    parser.add_argument("expected", type=Path)
    parser.add_argument("actual", type=Path)
    args = parser.parse_args()

    check_closed_objects(args.actual, args.expected)
    if args.closed_only:
        return

    expected_files = manifest(args.expected) | {"request_contract_generated.go"}
    actual_files = manifest(args.actual) | {"request_contract_generated.go"}
    if expected_files != actual_files:
        missing = sorted(expected_files - actual_files)
        extra = sorted(actual_files - expected_files)
        raise SystemExit(f"generated manifest drift; missing={missing}, extra={extra}")

    unexpected = {
        str(path.relative_to(args.expected))
        for pattern in ("*.go", "docs/*.md")
        for path in (args.expected / "openapi").glob(pattern)
        if path.name not in REPOSITORY_OWNED
    } - expected_files
    if unexpected:
        raise SystemExit("unmanaged generated files: " + ", ".join(sorted(unexpected)))

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
