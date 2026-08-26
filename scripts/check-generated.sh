#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
WORK="$(mktemp -d "${TMPDIR:-/tmp}/plainrouter-go-check.XXXXXX")"
trap 'rm -rf "$WORK"' EXIT

"$ROOT/scripts/generate.sh" "$WORK"
"${PYTHON_BIN:-python3}" "$ROOT/scripts/check-generated.py" "$ROOT" "$WORK"
